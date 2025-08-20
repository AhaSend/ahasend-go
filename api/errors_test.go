package api

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIError_Error(t *testing.T) {
	tests := []struct {
		name     string
		apiErr   *APIError
		expected string
	}{
		{
			name: "Basic error",
			apiErr: &APIError{
				Type:       ErrorTypeValidation,
				StatusCode: 400,
				Message:    "Invalid request",
			},
			expected: "validation error (HTTP 400): Invalid request",
		},
		{
			name: "Error with code",
			apiErr: &APIError{
				Type:       ErrorTypeAuthentication,
				StatusCode: 401,
				Message:    "Invalid API key",
				Code:       "AUTH_001",
			},
			expected: "authentication error (HTTP 401): Invalid API key [code: AUTH_001]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.apiErr.Error())
		})
	}
}

func TestAPIError_IsRetryable(t *testing.T) {
	tests := []struct {
		name      string
		errorType ErrorType
		retryable bool
	}{
		{"RateLimit", ErrorTypeRateLimit, true},
		{"Server", ErrorTypeServer, true},
		{"Network", ErrorTypeNetwork, true},
		{"Validation", ErrorTypeValidation, false},
		{"Authentication", ErrorTypeAuthentication, false},
		{"Permission", ErrorTypePermission, false},
		{"NotFound", ErrorTypeNotFound, false},
		{"Conflict", ErrorTypeConflict, false},
		{"Idempotency", ErrorTypeIdempotency, false},
		{"Unknown", ErrorTypeUnknown, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiErr := &APIError{Type: tt.errorType}
			assert.Equal(t, tt.retryable, apiErr.IsRetryable())
		})
	}
}

func TestParseAPIError(t *testing.T) {
	t.Run("Basic error response", func(t *testing.T) {
		resp := httptest.NewRecorder()
		resp.WriteHeader(http.StatusBadRequest)
		resp.Body.WriteString(`{"message": "Invalid email format"}`)

		httpResp := resp.Result()
		body := []byte(`{"message": "Invalid email format"}`)

		apiErr := ParseAPIError(httpResp, body)

		assert.Equal(t, ErrorTypeValidation, apiErr.Type)
		assert.Equal(t, 400, apiErr.StatusCode)
		assert.Equal(t, `{"message": "Invalid email format"}`, apiErr.Message)
		assert.Equal(t, body, apiErr.Raw)
	})

	t.Run("Rate limit with retry-after", func(t *testing.T) {
		resp := httptest.NewRecorder()
		resp.Header().Set("Retry-After", "60")
		resp.Header().Set("X-Request-Id", "req-123")
		resp.WriteHeader(http.StatusTooManyRequests)

		httpResp := resp.Result()
		body := []byte(`Rate limit exceeded`)

		apiErr := ParseAPIError(httpResp, body)

		assert.Equal(t, ErrorTypeRateLimit, apiErr.Type)
		assert.Equal(t, 429, apiErr.StatusCode)
		assert.Equal(t, 60, apiErr.RetryAfter)
		assert.Equal(t, "req-123", apiErr.RequestID)
	})

	t.Run("Empty body", func(t *testing.T) {
		resp := httptest.NewRecorder()
		resp.WriteHeader(http.StatusInternalServerError)

		httpResp := resp.Result()
		body := []byte{}

		apiErr := ParseAPIError(httpResp, body)

		assert.Equal(t, ErrorTypeServer, apiErr.Type)
		assert.Equal(t, 500, apiErr.StatusCode)
		assert.Equal(t, "Internal Server Error", apiErr.Message)
	})
}

func TestDetermineErrorType(t *testing.T) {
	tests := []struct {
		statusCode int
		expected   ErrorType
	}{
		{http.StatusBadRequest, ErrorTypeValidation},
		{http.StatusUnauthorized, ErrorTypeAuthentication},
		{http.StatusForbidden, ErrorTypePermission},
		{http.StatusNotFound, ErrorTypeNotFound},
		{http.StatusConflict, ErrorTypeConflict},
		{http.StatusPreconditionFailed, ErrorTypeIdempotency},
		{http.StatusTooManyRequests, ErrorTypeRateLimit},
		{http.StatusInternalServerError, ErrorTypeServer},
		{http.StatusBadGateway, ErrorTypeServer},
		{http.StatusServiceUnavailable, ErrorTypeServer},
		{http.StatusGatewayTimeout, ErrorTypeServer},
		{505, ErrorTypeServer},  // Other 5xx
		{418, ErrorTypeUnknown}, // I'm a teapot
	}

	for _, tt := range tests {
		t.Run(http.StatusText(tt.statusCode), func(t *testing.T) {
			assert.Equal(t, tt.expected, determineErrorType(tt.statusCode))
		})
	}
}

func TestNetworkError(t *testing.T) {
	t.Run("Error with operation", func(t *testing.T) {
		netErr := &NetworkError{
			Op:  "GET /api/messages",
			Err: errors.New("connection timeout"),
		}

		assert.Equal(t, "network error during GET /api/messages: connection timeout", netErr.Error())
		assert.True(t, netErr.IsRetryable())
	})

	t.Run("Error without operation", func(t *testing.T) {
		netErr := &NetworkError{
			Err: errors.New("DNS lookup failed"),
		}

		assert.Equal(t, "network error: DNS lookup failed", netErr.Error())
		assert.True(t, netErr.IsRetryable())
	})
}
