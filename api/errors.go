// Error types for the AhaSend Go SDK.
package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	Raw        []byte    `json:"-"`
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
// Note that the SDK's transport never retries a 4xx on its own, so an
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

	return apiErr
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

// retryAfterSeconds returns the Retry-After value in seconds for the error
// types that carry one, and 0 otherwise.
func retryAfterSeconds(resp *http.Response, errType ErrorType) int {
	switch errType {
	case ErrorTypeRateLimit:
		if seconds, err := strconv.Atoi(strings.TrimSpace(resp.Header.Get("Retry-After"))); err == nil {
			return seconds
		}
	case ErrorTypeIdempotencyConflict:
		if seconds, ok := idempotencyInProgressRetryAfter(resp); ok {
			return seconds
		}
	}
	return 0
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

// IsRetryable returns true as network errors are generally retryable
func (e *NetworkError) IsRetryable() bool {
	return true
}
