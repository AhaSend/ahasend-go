// Custom error types and error handling utilities for the AhaSend Go SDK.
//
// This file provides structured error types that make it easier to handle
// different error scenarios programmatically.

package api

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
	// Code is the error code for programmatic identification
	Code string `json:"code,omitempty"`
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
	// Endpoint is the API endpoint that generated the error
	Endpoint string `json:"endpoint,omitempty"`
	// Method is the HTTP method used
	Method string `json:"method,omitempty"`
	// Suggestions provides actionable guidance for fixing the error
	Suggestions []string `json:"suggestions,omitempty"`
	// Raw contains the raw error response body
	Raw []byte `json:"-"`
}

// Error implements the error interface
func (e *APIError) Error() string {
	var parts []string

	// Base error message
	baseMsg := fmt.Sprintf("%s error (HTTP %d): %s", e.Type, e.StatusCode, e.Message)
	parts = append(parts, baseMsg)

	// Add contextual information
	if e.Field != "" {
		parts = append(parts, fmt.Sprintf("field: %s", e.Field))
	}
	if e.Resource != "" {
		parts = append(parts, fmt.Sprintf("resource: %s", e.Resource))
	}
	if e.Method != "" && e.Endpoint != "" {
		parts = append(parts, fmt.Sprintf("%s %s", e.Method, e.Endpoint))
	}
	if e.Code != "" {
		parts = append(parts, fmt.Sprintf("code: %s", e.Code))
	}

	return strings.Join(parts, " | ")
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

// GetSuggestions returns actionable suggestions for fixing the error
func (e *APIError) GetSuggestions() []string {
	return e.Suggestions
}

// HasSuggestions returns true if the error has actionable suggestions
func (e *APIError) HasSuggestions() bool {
	return len(e.Suggestions) > 0
}

// GetEndpointInfo returns the endpoint and method information if available
func (e *APIError) GetEndpointInfo() (method, endpoint string) {
	return e.Method, e.Endpoint
}

// AddSuggestion adds a suggestion to the error
func (e *APIError) AddSuggestion(suggestion string) {
	if e.Suggestions == nil {
		e.Suggestions = make([]string, 0)
	}
	e.Suggestions = append(e.Suggestions, suggestion)
}

// SetEndpointInfo sets the endpoint and method information
func (e *APIError) SetEndpointInfo(method, endpoint string) {
	e.Method = method
	e.Endpoint = endpoint
}

// SetCode sets the error code for programmatic identification
func (e *APIError) SetCode(code string) {
	e.Code = code
}

// GetFormattedSuggestions returns suggestions formatted for CLI display
func (e *APIError) GetFormattedSuggestions() string {
	if len(e.Suggestions) == 0 {
		return ""
	}

	var formatted strings.Builder
	formatted.WriteString("Suggestions:\n")
	for i, suggestion := range e.Suggestions {
		formatted.WriteString(fmt.Sprintf("  %d. %s\n", i+1, suggestion))
	}
	return formatted.String()
}

// ParseAPIError creates an APIError from an HTTP response
func ParseAPIError(resp *http.Response, body []byte) *APIError {
	return ParseAPIErrorWithContext(resp, body, "", "")
}

// ParseAPIErrorWithContext creates an APIError from an HTTP response with additional context
func ParseAPIErrorWithContext(resp *http.Response, body []byte, method, endpoint string) *APIError {
	apiErr := &APIError{
		StatusCode: resp.StatusCode,
		Type:       determineErrorType(resp.StatusCode),
		Method:     method,
		Endpoint:   endpoint,
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
		// Generate helpful suggestions based on error type and context
		apiErr.generateSuggestions()
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

// generateSuggestions creates helpful suggestions based on error type and context
func (e *APIError) generateSuggestions() {
	switch e.Type {
	case ErrorTypeAuthentication:
		e.Suggestions = []string{
			"Check that your API key is valid and properly formatted",
			"Ensure the API key starts with 'aha-sk-'",
			"Verify the Authorization header is set: 'Bearer <your-api-key>'",
		}

	case ErrorTypePermission:
		e.Suggestions = []string{
			"Check that your API key has the required scopes for this operation",
			"Review the API documentation for required permissions",
		}
		if hasScope, scope := e.RequiresScope(); hasScope && scope != "" {
			e.AddSuggestion(fmt.Sprintf("Required scope: %s", scope))
		}

	case ErrorTypeValidation:
		if e.Field != "" {
			e.Suggestions = []string{
				fmt.Sprintf("Check the '%s' field in your request", e.Field),
				"Review the API documentation for field requirements",
			}
			switch e.Field {
			case "sender", "from":
				e.AddSuggestion("Ensure the sender email is from a domain you own")
				e.AddSuggestion("Example: noreply@yourdomain.com")
			case "to", "recipient":
				e.AddSuggestion("Verify the recipient email address format")
				e.AddSuggestion("Example: user@example.com")
			}
		} else {
			e.Suggestions = []string{
				"Review your request parameters for missing or invalid values",
				"Check the API documentation for required fields",
			}
		}

	case ErrorTypeNotFound:
		if e.Resource != "" {
			e.Suggestions = []string{
				fmt.Sprintf("Verify the %s ID is correct", e.Resource),
				fmt.Sprintf("Check that the %s exists and you have access to it", e.Resource),
			}
			if e.Resource == "domain" {
				e.AddSuggestion("Ensure the domain is properly configured with DNS records")
			}
		} else {
			e.Suggestions = []string{
				"Double-check the resource ID or identifier",
				"Ensure the resource exists and you have access to it",
			}
		}

	case ErrorTypeRateLimit:
		e.Suggestions = []string{
			"Reduce the frequency of your API requests",
			"Implement exponential backoff in your retry logic",
			"Consider upgrading your plan for higher rate limits",
		}
		if e.RetryAfter > 0 {
			e.AddSuggestion(fmt.Sprintf("Wait %d seconds before retrying", e.RetryAfter))
		}

	case ErrorTypeServer:
		e.Suggestions = []string{
			"This is a temporary server issue - try again in a few moments",
			"If the issue persists, contact AhaSend support",
		}
		if e.RequestID != "" {
			e.AddSuggestion(fmt.Sprintf("Include this request ID when contacting support: %s", e.RequestID))
		}

	case ErrorTypeConflict:
		e.Suggestions = []string{
			"The resource may already exist or be in a conflicting state",
			"Try using a different identifier or check existing resources",
		}

	case ErrorTypeIdempotency:
		e.Suggestions = []string{
			"Use a unique idempotency key for each logical request",
			"Don't reuse idempotency keys across different operations",
			"Idempotency keys expire after 24 hours",
		}
	}

	// Add endpoint-specific suggestions
	if e.Endpoint != "" {
		if strings.Contains(e.Endpoint, "/messages") && e.Method == "POST" {
			if e.Type == ErrorTypeValidation {
				e.AddSuggestion("Ensure either 'text_content' or 'html_content' is provided")
			}
		} else if strings.Contains(e.Endpoint, "/messages") && e.Method == "GET" {
			if e.Type == ErrorTypeValidation {
				e.AddSuggestion("Check your query parameters (sender, recipient, status, etc.)")
			}
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
