// Error types for the AhaSend Go SDK.
package api

import (
	"fmt"
	"net/http"
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
	ErrorTypeServer         ErrorType = "server"
	ErrorTypeNetwork        ErrorType = "network"
	ErrorTypeUnknown        ErrorType = "unknown"
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

// IsRetryable returns true if the error is retryable
func (e *APIError) IsRetryable() bool {
	switch e.Type {
	case ErrorTypeRateLimit, ErrorTypeServer, ErrorTypeNetwork:
		return true
	default:
		return false
	}
}

// ParseAPIError creates an APIError from an HTTP response
func ParseAPIError(resp *http.Response, body []byte) *APIError {
	apiErr := &APIError{
		StatusCode: resp.StatusCode,
		Type:       determineErrorType(resp.StatusCode),
		Raw:        body,
	}

	// Extract request ID from headers if available
	if reqID := resp.Header.Get("X-Request-Id"); reqID != "" {
		apiErr.RequestID = reqID
	}

	// Extract retry-after for rate limit errors
	if apiErr.Type == ErrorTypeRateLimit {
		if retryAfter := resp.Header.Get("Retry-After"); retryAfter != "" {
			var seconds int
			fmt.Sscanf(retryAfter, "%d", &seconds)
			apiErr.RetryAfter = seconds
		}
	}

	// Set a basic message - let the API response provide the details
	apiErr.Message = http.StatusText(resp.StatusCode)
	if len(body) > 0 && len(body) < 1000 {
		// Use the raw body as the message if it's reasonably sized
		apiErr.Message = string(body)
	}

	return apiErr
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
