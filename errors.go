// Custom error types and error handling utilities for the AhaSend Go SDK.
//
// This file provides structured error types that make it easier to handle
// different error scenarios programmatically.

package ahasend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// ErrorType represents the category of error
type ErrorType string

const (
	// ErrorTypeAuthentication indicates authentication failure (401)
	ErrorTypeAuthentication ErrorType = "authentication"
	// ErrorTypePermission indicates insufficient permissions (403)
	ErrorTypePermission ErrorType = "permission"
	// ErrorTypeValidation indicates invalid request parameters (400)
	ErrorTypeValidation ErrorType = "validation"
	// ErrorTypeNotFound indicates resource not found (404)
	ErrorTypeNotFound ErrorType = "not_found"
	// ErrorTypeConflict indicates resource conflict (409)
	ErrorTypeConflict ErrorType = "conflict"
	// ErrorTypeRateLimit indicates rate limit exceeded (429)
	ErrorTypeRateLimit ErrorType = "rate_limit"
	// ErrorTypeIdempotency indicates idempotency error (412)
	ErrorTypeIdempotency ErrorType = "idempotency"
	// ErrorTypeServer indicates server error (500+)
	ErrorTypeServer ErrorType = "server"
	// ErrorTypeNetwork indicates network/connection error
	ErrorTypeNetwork ErrorType = "network"
	// ErrorTypeUnknown indicates unknown error type
	ErrorTypeUnknown ErrorType = "unknown"
)

// APIError represents a structured error from the AhaSend API
type APIError struct {
	// Type categorizes the error for programmatic handling
	Type ErrorType `json:"-"`
	// StatusCode is the HTTP status code
	StatusCode int `json:"status_code,omitempty"`
	// Message is the error message from the API
	Message string `json:"message"`
	// RequestID is the request ID for tracking (if available)
	RequestID string `json:"request_id,omitempty"`
	// RetryAfter indicates when to retry (for rate limit errors)
	RetryAfter int `json:"retry_after,omitempty"`
	// Field indicates which field caused validation error (if applicable)
	Field string `json:"field,omitempty"`
	// Resource indicates which resource was not found or conflicted
	Resource string `json:"resource,omitempty"`
	// Action indicates what action was attempted
	Action string `json:"action,omitempty"`
	// Raw contains the raw error response body
	Raw []byte `json:"-"`
}

// Error implements the error interface
func (e *APIError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("%s error (HTTP %d): %s (field: %s)", e.Type, e.StatusCode, e.Message, e.Field)
	}
	if e.Resource != "" {
		return fmt.Sprintf("%s error (HTTP %d): %s (resource: %s)", e.Type, e.StatusCode, e.Message, e.Resource)
	}
	return fmt.Sprintf("%s error (HTTP %d): %s", e.Type, e.StatusCode, e.Message)
}

// IsRetryable returns true if the error is retryable
func (e *APIError) IsRetryable() bool {
	switch e.Type {
	case ErrorTypeRateLimit, ErrorTypeServer, ErrorTypeNetwork:
		return true
	case ErrorTypeConflict:
		// 409 conflicts are retryable in AhaSend's idempotency system
		return strings.Contains(strings.ToLower(e.Message), "in progress") ||
			strings.Contains(strings.ToLower(e.Message), "processing")
	default:
		return false
	}
}

// IsAuthError returns true if this is an authentication error
func (e *APIError) IsAuthError() bool {
	return e.Type == ErrorTypeAuthentication
}

// IsPermissionError returns true if this is a permission/scope error
func (e *APIError) IsPermissionError() bool {
	return e.Type == ErrorTypePermission
}

// IsValidationError returns true if this is a validation error
func (e *APIError) IsValidationError() bool {
	return e.Type == ErrorTypeValidation
}

// IsNotFoundError returns true if this is a not found error
func (e *APIError) IsNotFoundError() bool {
	return e.Type == ErrorTypeNotFound
}

// IsRateLimitError returns true if this is a rate limit error
func (e *APIError) IsRateLimitError() bool {
	return e.Type == ErrorTypeRateLimit
}

// IsIdempotencyError returns true if this is an idempotency error
func (e *APIError) IsIdempotencyError() bool {
	return e.Type == ErrorTypeIdempotency
}

// RequiresScope checks if the error message indicates a missing scope
func (e *APIError) RequiresScope() (bool, string) {
	if e.Type != ErrorTypePermission {
		return false, ""
	}

	// Check for scope-related messages
	msg := strings.ToLower(e.Message)
	if strings.Contains(msg, "scope") || strings.Contains(msg, "permission") {
		// Try to extract the required scope from the message
		// Common patterns: "requires scope: messages:send:all"
		// or "missing permission: domains:write"

		// Look for patterns like "scope: xxx" or "permission: xxx"
		if idx := strings.Index(msg, "scope:"); idx >= 0 {
			after := strings.TrimSpace(e.Message[idx+6:]) // Skip "scope:"
			return true, after
		}
		if idx := strings.Index(msg, "permission:"); idx >= 0 {
			after := strings.TrimSpace(e.Message[idx+11:]) // Skip "permission:"
			return true, after
		}

		return true, ""
	}
	return false, ""
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
			// Parse retry-after (assuming seconds as integer)
			var seconds int
			fmt.Sscanf(retryAfter, "%d", &seconds)
			apiErr.RetryAfter = seconds
		}
	}

	// Parse the error response body
	var errorResp struct {
		Message string `json:"message"`
	}

	if err := json.Unmarshal(body, &errorResp); err == nil && errorResp.Message != "" {
		apiErr.Message = errorResp.Message

		// Try to extract additional context from the message
		apiErr.parseMessageContext()
	} else {
		// Fallback to HTTP status text if parsing fails
		apiErr.Message = http.StatusText(resp.StatusCode)
		if len(body) > 0 && len(body) < 1000 {
			// Include raw body if it's reasonably sized and not JSON
			apiErr.Message = string(body)
		}
	}

	return apiErr
}

// parseMessageContext extracts additional context from error messages
func (e *APIError) parseMessageContext() {
	msg := strings.ToLower(e.Message)

	// Extract field information for validation errors
	if e.Type == ErrorTypeValidation {
		if strings.Contains(msg, "invalid") || strings.Contains(msg, "missing") {
			// Common patterns: "invalid email", "missing field: subject"
			parts := strings.Split(e.Message, ":")
			if len(parts) > 1 {
				e.Field = strings.TrimSpace(parts[1])
			} else if strings.Contains(msg, "invalid ") {
				// Extract field from "invalid {field}" pattern
				after := strings.TrimPrefix(msg, "invalid ")
				if space := strings.Index(after, " "); space > 0 {
					e.Field = after[:space]
				} else {
					e.Field = after
				}
			}
		}
	}

	// Extract resource information
	if e.Type == ErrorTypeNotFound {
		// Common patterns: "domain not found", "message {id} not found"
		if strings.Contains(msg, "domain") {
			e.Resource = "domain"
		} else if strings.Contains(msg, "message") {
			e.Resource = "message"
		} else if strings.Contains(msg, "api key") || strings.Contains(msg, "api_key") {
			e.Resource = "api_key"
		} else if strings.Contains(msg, "webhook") {
			e.Resource = "webhook"
		} else if strings.Contains(msg, "route") {
			e.Resource = "route"
		} else if strings.Contains(msg, "account") {
			e.Resource = "account"
		}
	}
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

// IsError checks if an error is an APIError
func IsAPIError(err error) (*APIError, bool) {
	apiErr, ok := err.(*APIError)
	return apiErr, ok
}

// IsNetworkError checks if an error is a NetworkError
func IsNetworkError(err error) (*NetworkError, bool) {
	netErr, ok := err.(*NetworkError)
	return netErr, ok
}

// IsRetryableError checks if any error is retryable
func IsRetryableError(err error) bool {
	if apiErr, ok := IsAPIError(err); ok {
		return apiErr.IsRetryable()
	}
	if netErr, ok := IsNetworkError(err); ok {
		return netErr.IsRetryable()
	}
	return false
}

// GetErrorMessage extracts a user-friendly message from any error
func GetErrorMessage(err error) string {
	if apiErr, ok := IsAPIError(err); ok {
		return apiErr.Message
	}
	if netErr, ok := IsNetworkError(err); ok {
		return netErr.Error()
	}
	return err.Error()
}

// ValidationError represents multiple field validation errors
type ValidationError struct {
	Fields map[string]string // Field name -> error message
}

// Error implements the error interface
func (e *ValidationError) Error() string {
	if len(e.Fields) == 0 {
		return "validation error"
	}

	var errors []string
	for field, msg := range e.Fields {
		errors = append(errors, fmt.Sprintf("%s: %s", field, msg))
	}
	return fmt.Sprintf("validation errors: %s", strings.Join(errors, ", "))
}

// AddFieldError adds a field validation error
func (e *ValidationError) AddFieldError(field, message string) {
	if e.Fields == nil {
		e.Fields = make(map[string]string)
	}
	e.Fields[field] = message
}

// HasErrors returns true if there are validation errors
func (e *ValidationError) HasErrors() bool {
	return len(e.Fields) > 0
}
