package ahasend_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/AhaSend/ahasend-go"
	"github.com/stretchr/testify/assert"
)

func TestAPIErrorTypes(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		body       string
		wantType   ahasend.ErrorType
		wantMsg    string
	}{
		{
			name:       "Authentication Error",
			statusCode: 401,
			body:       `{"message": "Missing or malformed bearer token"}`,
			wantType:   ahasend.ErrorTypeAuthentication,
			wantMsg:    "Missing or malformed bearer token",
		},
		{
			name:       "Permission Error",
			statusCode: 403,
			body:       `{"message": "Insufficient permissions: requires scope messages:send:all"}`,
			wantType:   ahasend.ErrorTypePermission,
			wantMsg:    "Insufficient permissions: requires scope messages:send:all",
		},
		{
			name:       "Validation Error",
			statusCode: 400,
			body:       `{"message": "invalid email format"}`,
			wantType:   ahasend.ErrorTypeValidation,
			wantMsg:    "invalid email format",
		},
		{
			name:       "Not Found Error",
			statusCode: 404,
			body:       `{"message": "domain not found"}`,
			wantType:   ahasend.ErrorTypeNotFound,
			wantMsg:    "domain not found",
		},
		{
			name:       "Conflict Error",
			statusCode: 409,
			body:       `{"message": "Request is currently being processed"}`,
			wantType:   ahasend.ErrorTypeConflict,
			wantMsg:    "Request is currently being processed",
		},
		{
			name:       "Rate Limit Error",
			statusCode: 429,
			body:       `{"message": "Rate limit exceeded"}`,
			wantType:   ahasend.ErrorTypeRateLimit,
			wantMsg:    "Rate limit exceeded",
		},
		{
			name:       "Idempotency Error",
			statusCode: 412,
			body:       `{"message": "Idempotency key already used"}`,
			wantType:   ahasend.ErrorTypeIdempotency,
			wantMsg:    "Idempotency key already used",
		},
		{
			name:       "Server Error",
			statusCode: 500,
			body:       `{"message": "Internal server error"}`,
			wantType:   ahasend.ErrorTypeServer,
			wantMsg:    "Internal server error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Header:     make(http.Header),
			}

			apiErr := ahasend.ParseAPIError(resp, []byte(tt.body))

			assert.Equal(t, tt.wantType, apiErr.Type, "Error type mismatch")
			assert.Equal(t, tt.statusCode, apiErr.StatusCode, "Status code mismatch")
			assert.Equal(t, tt.wantMsg, apiErr.Message, "Message mismatch")
		})
	}
}

func TestAPIErrorRetryable(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		body       string
		wantRetry  bool
	}{
		{
			name:       "Rate limit is retryable",
			statusCode: 429,
			body:       `{"message": "Rate limit exceeded"}`,
			wantRetry:  true,
		},
		{
			name:       "Server error is retryable",
			statusCode: 503,
			body:       `{"message": "Service unavailable"}`,
			wantRetry:  true,
		},
		{
			name:       "Conflict in progress is retryable",
			statusCode: 409,
			body:       `{"message": "Request is currently in progress"}`,
			wantRetry:  true,
		},
		{
			name:       "Validation error is not retryable",
			statusCode: 400,
			body:       `{"message": "Invalid email"}`,
			wantRetry:  false,
		},
		{
			name:       "Auth error is not retryable",
			statusCode: 401,
			body:       `{"message": "Invalid token"}`,
			wantRetry:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Header:     make(http.Header),
			}

			apiErr := ahasend.ParseAPIError(resp, []byte(tt.body))
			assert.Equal(t, tt.wantRetry, apiErr.IsRetryable(), "Retryable mismatch")
		})
	}
}

func TestAPIErrorScopeDetection(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		body       string
		wantScope  bool
		scopeValue string
	}{
		{
			name:       "Detects missing scope",
			statusCode: 403,
			body:       `{"message": "requires scope: messages:send:all"}`,
			wantScope:  true,
			scopeValue: "messages:send:all",
		},
		{
			name:       "Detects missing permission",
			statusCode: 403,
			body:       `{"message": "missing permission: domains:write"}`,
			wantScope:  true,
			scopeValue: "domains:write",
		},
		{
			name:       "No scope in non-permission error",
			statusCode: 400,
			body:       `{"message": "invalid request"}`,
			wantScope:  false,
			scopeValue: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Header:     make(http.Header),
			}

			apiErr := ahasend.ParseAPIError(resp, []byte(tt.body))
			hasScope, scope := apiErr.RequiresScope()

			assert.Equal(t, tt.wantScope, hasScope, "Scope detection mismatch")
			if tt.wantScope {
				assert.Equal(t, tt.scopeValue, scope, "Scope value mismatch")
			}
		})
	}
}

func TestAPIErrorResourceExtraction(t *testing.T) {
	tests := []struct {
		name         string
		statusCode   int
		body         string
		wantResource string
	}{
		{
			name:         "Domain not found",
			statusCode:   404,
			body:         `{"message": "domain not found"}`,
			wantResource: "domain",
		},
		{
			name:         "Message not found",
			statusCode:   404,
			body:         `{"message": "message abc123 not found"}`,
			wantResource: "message",
		},
		{
			name:         "API key not found",
			statusCode:   404,
			body:         `{"message": "api key not found"}`,
			wantResource: "api_key",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Header:     make(http.Header),
			}

			apiErr := ahasend.ParseAPIError(resp, []byte(tt.body))
			assert.Equal(t, tt.wantResource, apiErr.Resource, "Resource mismatch")
		})
	}
}

func TestAPIErrorFieldExtraction(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		body       string
		wantField  string
	}{
		{
			name:       "Invalid field with colon",
			statusCode: 400,
			body:       `{"message": "invalid field: email"}`,
			wantField:  "email",
		},
		{
			name:       "Invalid field without colon",
			statusCode: 400,
			body:       `{"message": "invalid email format"}`,
			wantField:  "email",
		},
		{
			name:       "Missing field",
			statusCode: 400,
			body:       `{"message": "missing field: subject"}`,
			wantField:  "subject",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Header:     make(http.Header),
			}

			apiErr := ahasend.ParseAPIError(resp, []byte(tt.body))
			assert.Equal(t, tt.wantField, apiErr.Field, "Field mismatch")
		})
	}
}

func TestAPIErrorHeaders(t *testing.T) {
	t.Run("Extract Request ID", func(t *testing.T) {
		resp := &http.Response{
			StatusCode: 400,
			Header: http.Header{
				"X-Request-Id": []string{"req-12345"},
			},
		}

		apiErr := ahasend.ParseAPIError(resp, []byte(`{"message": "error"}`))
		assert.Equal(t, "req-12345", apiErr.RequestID, "Request ID mismatch")
	})

	t.Run("Extract Retry-After", func(t *testing.T) {
		resp := &http.Response{
			StatusCode: 429,
			Header: http.Header{
				"Retry-After": []string{"60"},
			},
		}

		apiErr := ahasend.ParseAPIError(resp, []byte(`{"message": "rate limited"}`))
		assert.Equal(t, 60, apiErr.RetryAfter, "Retry-After mismatch")
	})
}

func TestNetworkError(t *testing.T) {
	netErr := &ahasend.NetworkError{
		Op:  "POST /v2/messages",
		Err: errors.New("connection refused"),
	}

	assert.True(t, netErr.IsRetryable(), "Network errors should be retryable")
	assert.Contains(t, netErr.Error(), "POST /v2/messages", "Should include operation")
	assert.Contains(t, netErr.Error(), "connection refused", "Should include underlying error")
}

func TestErrorHelpers(t *testing.T) {
	t.Run("IsAPIError", func(t *testing.T) {
		apiErr := &ahasend.APIError{
			Type:       ahasend.ErrorTypeValidation,
			StatusCode: 400,
			Message:    "test error",
		}

		result, ok := ahasend.IsAPIError(apiErr)
		assert.True(t, ok, "Should recognize APIError")
		assert.Equal(t, apiErr, result, "Should return same error")

		_, ok = ahasend.IsAPIError(errors.New("regular error"))
		assert.False(t, ok, "Should not recognize regular error")
	})

	t.Run("IsNetworkError", func(t *testing.T) {
		netErr := &ahasend.NetworkError{
			Err: errors.New("timeout"),
		}

		result, ok := ahasend.IsNetworkError(netErr)
		assert.True(t, ok, "Should recognize NetworkError")
		assert.Equal(t, netErr, result, "Should return same error")

		_, ok = ahasend.IsNetworkError(errors.New("regular error"))
		assert.False(t, ok, "Should not recognize regular error")
	})

	t.Run("IsRetryableError", func(t *testing.T) {
		retryableAPI := &ahasend.APIError{
			Type: ahasend.ErrorTypeRateLimit,
		}
		assert.True(t, ahasend.IsRetryableError(retryableAPI), "Rate limit should be retryable")

		nonRetryableAPI := &ahasend.APIError{
			Type: ahasend.ErrorTypeValidation,
		}
		assert.False(t, ahasend.IsRetryableError(nonRetryableAPI), "Validation should not be retryable")

		netErr := &ahasend.NetworkError{
			Err: errors.New("timeout"),
		}
		assert.True(t, ahasend.IsRetryableError(netErr), "Network error should be retryable")
	})

	t.Run("GetErrorMessage", func(t *testing.T) {
		apiErr := &ahasend.APIError{
			Message: "API error message",
		}
		assert.Equal(t, "API error message", ahasend.GetErrorMessage(apiErr))

		netErr := &ahasend.NetworkError{
			Op:  "GET /test",
			Err: errors.New("timeout"),
		}
		assert.Contains(t, ahasend.GetErrorMessage(netErr), "timeout")

		regularErr := errors.New("regular error")
		assert.Equal(t, "regular error", ahasend.GetErrorMessage(regularErr))
	})
}

func TestValidationError(t *testing.T) {
	valErr := &ahasend.ValidationError{}

	assert.False(t, valErr.HasErrors(), "Should have no errors initially")

	valErr.AddFieldError("email", "invalid format")
	valErr.AddFieldError("subject", "required")

	assert.True(t, valErr.HasErrors(), "Should have errors after adding")
	assert.Contains(t, valErr.Error(), "email: invalid format")
	assert.Contains(t, valErr.Error(), "subject: required")
	assert.Equal(t, 2, len(valErr.Fields), "Should have 2 field errors")
}

func TestGenericOpenAPIErrorCompatibility(t *testing.T) {
	// Test that GenericOpenAPIError wraps APIError correctly
	resp := &http.Response{
		StatusCode: 403,
		Status:     "403 Forbidden",
		Header:     make(http.Header),
	}
	body := []byte(`{"message": "Insufficient permissions"}`)

	genErr := ahasend.NewGenericOpenAPIError(resp, body, nil)

	assert.NotNil(t, genErr.APIError, "Should have APIError")
	assert.Equal(t, ahasend.ErrorTypePermission, genErr.Type(), "Should have correct type")
	assert.Equal(t, 403, genErr.StatusCode(), "Should have correct status code")
	assert.False(t, genErr.IsRetryable(), "Permission error should not be retryable")
	assert.Equal(t, body, genErr.Body(), "Should preserve body")
	assert.Contains(t, genErr.Error(), "Insufficient permissions", "Should have correct error message")
}

func TestErrorFormatting(t *testing.T) {
	tests := []struct {
		name     string
		apiErr   *ahasend.APIError
		expected string
	}{
		{
			name: "Error with field",
			apiErr: &ahasend.APIError{
				Type:       ahasend.ErrorTypeValidation,
				StatusCode: 400,
				Message:    "invalid format",
				Field:      "email",
			},
			expected: "validation error (HTTP 400): invalid format (field: email)",
		},
		{
			name: "Error with resource",
			apiErr: &ahasend.APIError{
				Type:       ahasend.ErrorTypeNotFound,
				StatusCode: 404,
				Message:    "not found",
				Resource:   "domain",
			},
			expected: "not_found error (HTTP 404): not found (resource: domain)",
		},
		{
			name: "Simple error",
			apiErr: &ahasend.APIError{
				Type:       ahasend.ErrorTypeAuthentication,
				StatusCode: 401,
				Message:    "Invalid token",
			},
			expected: "authentication error (HTTP 401): Invalid token",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.apiErr.Error(), "Error format mismatch")
		})
	}
}

func TestErrorTypeHelpers(t *testing.T) {
	apiErr := &ahasend.APIError{}

	// Test each helper method
	apiErr.Type = ahasend.ErrorTypeAuthentication
	assert.True(t, apiErr.IsAuthError())
	assert.False(t, apiErr.IsPermissionError())

	apiErr.Type = ahasend.ErrorTypePermission
	assert.True(t, apiErr.IsPermissionError())
	assert.False(t, apiErr.IsValidationError())

	apiErr.Type = ahasend.ErrorTypeValidation
	assert.True(t, apiErr.IsValidationError())
	assert.False(t, apiErr.IsNotFoundError())

	apiErr.Type = ahasend.ErrorTypeNotFound
	assert.True(t, apiErr.IsNotFoundError())
	assert.False(t, apiErr.IsRateLimitError())

	apiErr.Type = ahasend.ErrorTypeRateLimit
	assert.True(t, apiErr.IsRateLimitError())
	assert.False(t, apiErr.IsIdempotencyError())

	apiErr.Type = ahasend.ErrorTypeIdempotency
	assert.True(t, apiErr.IsIdempotencyError())
	assert.False(t, apiErr.IsAuthError())
}
