// Error types for the AhaSend Go SDK.
package api

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/AhaSend/ahasend-go/models/common"
)

// ErrorType represents the category of error
type ErrorType string

const (
	ErrorTypeAuthentication ErrorType = "authentication"
	ErrorTypePermission     ErrorType = "permission"
	ErrorTypeValidation     ErrorType = "validation"
	ErrorTypeNotFound       ErrorType = "not_found"
	ErrorTypeConflict       ErrorType = "conflict"
	ErrorTypeRateLimit      ErrorType = "rate_limit"
	ErrorTypeIdempotency    ErrorType = "idempotency"
	// ErrorTypeIdempotencyConflict is the one retryable 409: an earlier
	// request with the same idempotency key is still in flight, so the API
	// has not decided the outcome yet. APIError.RetryAfter is what remains of
	// that request's execution lease; retrying with the same key after it
	// elapses replays the stored result if the original finished, and
	// otherwise takes the lease over and executes again. Every other 409 is a
	// terminal conflict (ErrorTypeConflict).
	ErrorTypeIdempotencyConflict ErrorType = "idempotency_conflict"
	ErrorTypeServer              ErrorType = "server"
	ErrorTypeNetwork             ErrorType = "network"
	ErrorTypeUnknown             ErrorType = "unknown"
)

// APIError represents an error from the AhaSend API
type APIError struct {
	Type       ErrorType `json:"-"`
	StatusCode int       `json:"status_code,omitempty"`
	Message    string    `json:"message"`
	Code       string    `json:"code,omitempty"`
	RequestID  string    `json:"request_id,omitempty"`
	RetryAfter int       `json:"retry_after,omitempty"`
	// Errors are the problems of AhaSend Markup (AhaML) that a write refused,
	// in document order, up to the first 100; Message says how many there
	// are in all. It is nil for every other error.
	Errors []common.MarkupProblem `json:"errors,omitempty"`
	Raw    []byte                 `json:"-"`
}

// Error implements the error interface
func (e *APIError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("%s error (HTTP %d): %s [code: %s]", e.Type, e.StatusCode, e.Message, e.Code)
	}
	return fmt.Sprintf("%s error (HTTP %d): %s", e.Type, e.StatusCode, e.Message)
}

// IsRetryable returns true if the error is retryable.
//
// Note that the SDK's transport retries no 4xx on its own other than 429
// (unless RetryConfig.RetryClientErrors is set), so an
// ErrorTypeIdempotencyConflict is reported as retryable but is not retried
// automatically: retrying it is the caller's decision, after waiting
// RetryAfter seconds.
func (e *APIError) IsRetryable() bool {
	switch e.Type {
	case ErrorTypeRateLimit, ErrorTypeServer, ErrorTypeNetwork, ErrorTypeIdempotencyConflict:
		return true
	default:
		return false
	}
}

// ParseAPIError creates an APIError from an HTTP response
func ParseAPIError(resp *http.Response, body []byte) *APIError {
	apiErr := &APIError{
		StatusCode: resp.StatusCode,
		Type:       classifyErrorType(resp),
		Raw:        body,
	}

	// Extract request ID from headers if available
	if reqID := resp.Header.Get("X-Request-Id"); reqID != "" {
		apiErr.RequestID = reqID
	}

	apiErr.RetryAfter = retryAfterSeconds(resp, apiErr.Type)

	// Set a basic message - let the API response provide the details
	apiErr.Message = http.StatusText(resp.StatusCode)
	if len(body) > 0 && len(body) < 1000 {
		// Use the raw body as the message if it's reasonably sized
		apiErr.Message = string(body)
	}

	var errorResp errorResponseBody
	if err := json.Unmarshal(body, &errorResp); err == nil {
		apiErr.Errors = markupProblems(errorResp.Errors)
	}

	return apiErr
}

// errorResponseBody is the body of an error response: a message, and for
// refused AhaML its problems. Errors stays raw, so an errors field of another
// shape does not keep the message from decoding.
type errorResponseBody struct {
	common.ErrorResponse
	Errors json.RawMessage `json:"errors"`
}

// markupProblems decodes the errors of a refused AhaML write. It returns nil
// when there are none, or when they are not a list of problems, each with a
// code.
func markupProblems(raw json.RawMessage) []common.MarkupProblem {
	if len(raw) == 0 {
		return nil
	}
	var problems []common.MarkupProblem
	if err := json.Unmarshal(raw, &problems); err != nil {
		return nil
	}
	for _, problem := range problems {
		if problem.Code == "" {
			return nil
		}
	}
	return problems
}

// classifyErrorType maps an HTTP response to an error type. It is
// determineErrorType plus the refinements a bare status code cannot express.
func classifyErrorType(resp *http.Response) ErrorType {
	if resp.StatusCode == http.StatusConflict {
		if _, ok := idempotencyInProgressRetryAfter(resp); ok {
			return ErrorTypeIdempotencyConflict
		}
	}
	return determineErrorType(resp.StatusCode)
}

// retryAfterSeconds returns the Retry-After value in whole seconds, rounded
// up, for the responses that carry one (a 429, a 503, and a 409 for an
// idempotent request still in flight), and 0 otherwise.
func retryAfterSeconds(resp *http.Response, errType ErrorType) int {
	switch errType {
	case ErrorTypeRateLimit, ErrorTypeServer:
		if wait, ok := retryAfter(resp); ok {
			// Seconds are reported as sent, even past what a Duration holds
			if seconds, err := strconv.Atoi(strings.TrimSpace(resp.Header.Get("Retry-After"))); err == nil {
				return seconds
			}
			seconds := wait / time.Second
			if wait%time.Second != 0 {
				seconds++
			}
			return int(seconds)
		}
	case ErrorTypeIdempotencyConflict:
		if seconds, ok := idempotencyInProgressRetryAfter(resp); ok {
			return seconds
		}
	}
	return 0
}

// retryAfter returns the wait a 429 or 503 response asks for in its
// Retry-After header, given as seconds or as an HTTP date. A date in the past
// is a wait of zero, and a wait longer than a Duration holds is the longest
// Duration. Other responses, and a missing or malformed header, report false.
func retryAfter(resp *http.Response) (time.Duration, bool) {
	if resp.StatusCode != http.StatusTooManyRequests && resp.StatusCode != http.StatusServiceUnavailable {
		return 0, false
	}
	value := strings.TrimSpace(resp.Header.Get("Retry-After"))
	if value == "" {
		return 0, false
	}
	if seconds, err := strconv.ParseInt(value, 10, 64); err == nil {
		if seconds < 0 {
			return 0, false
		}
		if seconds > int64(math.MaxInt64/time.Second) {
			return time.Duration(math.MaxInt64), true
		}
		return time.Duration(seconds) * time.Second, true
	}
	if at, err := http.ParseTime(value); err == nil {
		if wait := time.Until(at); wait > 0 {
			return wait, true
		}
		return 0, true
	}
	return 0, false
}

// idempotencyInProgressRetryAfter reports whether a 409 response describes an
// idempotent request that is still in flight, and if so how many seconds to
// wait before retrying.
//
// The API marks that state with two headers together: `Idempotent-Replayed:
// false`, meaning no stored result was returned, and a positive integer
// `Retry-After`. A duplicate-resource 409 carries neither, and a replayed
// success carries `Idempotent-Replayed: true`. The request must also have
// carried an idempotency key, since the state only exists per key - checked
// when the response still references its request, which is always the case
// for responses produced by an http.Client.
func idempotencyInProgressRetryAfter(resp *http.Response) (int, bool) {
	if resp.Request != nil && resp.Request.Header.Get("Idempotency-Key") == "" {
		return 0, false
	}
	if !strings.EqualFold(strings.TrimSpace(resp.Header.Get("Idempotent-Replayed")), "false") {
		return 0, false
	}
	seconds, err := strconv.Atoi(strings.TrimSpace(resp.Header.Get("Retry-After")))
	if err != nil || seconds <= 0 {
		return 0, false
	}
	return seconds, true
}

// determineErrorType maps HTTP status codes to error types
func determineErrorType(statusCode int) ErrorType {
	switch statusCode {
	case http.StatusBadRequest:
		return ErrorTypeValidation
	case http.StatusUnauthorized:
		return ErrorTypeAuthentication
	case http.StatusForbidden:
		return ErrorTypePermission
	case http.StatusNotFound:
		return ErrorTypeNotFound
	case http.StatusConflict:
		return ErrorTypeConflict
	case http.StatusUnprocessableEntity:
		return ErrorTypeIdempotency
	case http.StatusPreconditionFailed:
		return ErrorTypeIdempotency
	case http.StatusTooManyRequests:
		return ErrorTypeRateLimit
	case http.StatusInternalServerError,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout:
		return ErrorTypeServer
	default:
		if statusCode >= 500 {
			return ErrorTypeServer
		}
		return ErrorTypeUnknown
	}
}

// NetworkError represents network-level errors
type NetworkError struct {
	Op  string // Operation attempted
	Err error  // Underlying error
}

// Error implements the error interface
func (e *NetworkError) Error() string {
	if e.Op != "" {
		return fmt.Sprintf("network error during %s: %v", e.Op, e.Err)
	}
	return fmt.Sprintf("network error: %v", e.Err)
}

// Unwrap returns the underlying network error.
func (e *NetworkError) Unwrap() error {
	return e.Err
}

// IsRetryable returns true as network errors are generally retryable
func (e *NetworkError) IsRetryable() bool {
	return true
}
