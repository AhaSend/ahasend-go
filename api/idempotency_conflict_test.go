package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync/atomic"
	"testing"

	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// conflictResponse builds a 409 response as if it came from a request that
// carried the given idempotency key.
func conflictResponse(t *testing.T, idempotencyKey string, headers map[string]string) *http.Response {
	t.Helper()

	req, err := http.NewRequest(http.MethodPost, "https://api.ahasend.com/v2/accounts/acct/domains", nil)
	require.NoError(t, err)
	if idempotencyKey != "" {
		req.Header.Set("Idempotency-Key", idempotencyKey)
	}

	resp := &http.Response{
		StatusCode: http.StatusConflict,
		Header:     http.Header{},
		Request:    req,
	}
	for name, value := range headers {
		resp.Header.Set(name, value)
	}
	return resp
}

func TestParseAPIErrorClassifies409(t *testing.T) {
	tests := []struct {
		name           string
		idempotencyKey string
		headers        map[string]string
		wantType       ErrorType
		wantRetryAfter int
	}{
		{
			name:           "in-progress conflict is retryable",
			idempotencyKey: "key-1",
			headers:        map[string]string{"Idempotent-Replayed": "false", "Retry-After": "2"},
			wantType:       ErrorTypeIdempotencyConflict,
			wantRetryAfter: 2,
		},
		{
			// A duplicate resource carries neither header: the request is
			// decided and retrying it will fail the same way forever.
			name:           "duplicate resource is terminal",
			idempotencyKey: "key-1",
			headers:        nil,
			wantType:       ErrorTypeConflict,
		},
		{
			// A replayed result means the original request finished.
			name:           "replayed conflict is terminal",
			idempotencyKey: "key-1",
			headers:        map[string]string{"Idempotent-Replayed": "true", "Retry-After": "2"},
			wantType:       ErrorTypeConflict,
		},
		{
			name:           "in-progress marker without retry-after is terminal",
			idempotencyKey: "key-1",
			headers:        map[string]string{"Idempotent-Replayed": "false"},
			wantType:       ErrorTypeConflict,
		},
		{
			name:           "non-positive retry-after is terminal",
			idempotencyKey: "key-1",
			headers:        map[string]string{"Idempotent-Replayed": "false", "Retry-After": "0"},
			wantType:       ErrorTypeConflict,
		},
		{
			// Retry-After may also be an HTTP date, which is not the
			// in-progress signal.
			name:           "http-date retry-after is terminal",
			idempotencyKey: "key-1",
			headers:        map[string]string{"Idempotent-Replayed": "false", "Retry-After": "Wed, 21 Oct 2026 07:28:00 GMT"},
			wantType:       ErrorTypeConflict,
		},
		{
			// The in-progress state only exists per idempotency key, so an
			// unkeyed request cannot be in it.
			name:           "unkeyed request is terminal",
			idempotencyKey: "",
			headers:        map[string]string{"Idempotent-Replayed": "false", "Retry-After": "2"},
			wantType:       ErrorTypeConflict,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiErr := ParseAPIError(conflictResponse(t, tt.idempotencyKey, tt.headers), []byte(`{"message":"conflict"}`))

			assert.Equal(t, tt.wantType, apiErr.Type)
			assert.Equal(t, http.StatusConflict, apiErr.StatusCode)
			assert.Equal(t, tt.wantRetryAfter, apiErr.RetryAfter)
			assert.Equal(t, tt.wantType == ErrorTypeIdempotencyConflict, apiErr.IsRetryable())
		})
	}
}

func TestParseAPIErrorLeaves422Unchanged(t *testing.T) {
	resp := conflictResponse(t, "key-1", map[string]string{"Idempotent-Replayed": "false", "Retry-After": "2"})
	resp.StatusCode = http.StatusUnprocessableEntity

	apiErr := ParseAPIError(resp, []byte(`{"message":"idempotency key reused"}`))

	assert.Equal(t, ErrorTypeIdempotency, apiErr.Type)
	assert.False(t, apiErr.IsRetryable())
	assert.Equal(t, 0, apiErr.RetryAfter)
}

func TestInProgressConflictSurfacedButNotAutoRetried(t *testing.T) {
	var requestCount int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&requestCount, 1)

		w.Header().Set("Idempotent-Replayed", "false")
		w.Header().Set("Retry-After", "3")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		_, _ = w.Write([]byte(`{"message":"a request with this idempotency key is in progress"}`))
	}))
	defer server.Close()

	serverURL, err := url.Parse(server.URL)
	require.NoError(t, err)

	cfg := NewConfiguration()
	cfg.Host = serverURL.Host
	cfg.Scheme = serverURL.Scheme
	cfg.APIKey = "test-key"
	// Retries are on: the point of this test is that the transport still does
	// not act on the retryable classification by itself.
	cfg.RetryConfig.Enabled = true
	cfg.RetryConfig.MaxRetries = 3
	client := NewAPIClientWithConfig(cfg)

	_, httpResp, err := client.DomainsAPI.CreateDomain(context.Background(), uuid.New(), requests.CreateDomainRequest{
		Domain: "example.com",
	}, WithIdempotencyKey("in-progress-key"))

	require.Error(t, err)

	apiErr, ok := err.(*APIError)
	require.True(t, ok, "expected an *APIError, got %T", err)
	assert.Equal(t, ErrorTypeIdempotencyConflict, apiErr.Type)
	assert.True(t, apiErr.IsRetryable())
	assert.Equal(t, 3, apiErr.RetryAfter)

	assert.Equal(t, int32(1), atomic.LoadInt32(&requestCount), "the transport must not retry a 409 on its own")

	// The discriminator is also readable from the response the call returns.
	require.NotNil(t, httpResp)
	assert.Equal(t, "false", httpResp.Header.Get("Idempotent-Replayed"))
}
