package api

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// trackedBody is a response body that records whether it was closed.
type trackedBody struct {
	*strings.Reader
	mu     sync.Mutex
	closed bool
}

func newTrackedBody(s string) *trackedBody {
	return &trackedBody{Reader: strings.NewReader(s)}
}

func (b *trackedBody) Close() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.closed = true
	return nil
}

func (b *trackedBody) isClosed() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.closed
}

// fastRetryConfig retries quickly so the tests do not wait on backoff.
func fastRetryConfig(maxDelay time.Duration) RetryConfig {
	return RetryConfig{
		Enabled:         true,
		MaxRetries:      3,
		BackoffStrategy: BackoffConstant,
		BaseDelay:       time.Millisecond,
		MaxDelay:        maxDelay,
	}
}

// retryTestClient returns a client whose requests go to transport.
func retryTestClient(transport http.RoundTripper, retry RetryConfig) *APIClient {
	config := NewConfiguration()
	config.APIKey = "test-key"
	config.EnableRateLimit = false
	config.RetryConfig = retry
	config.HTTPClient = &http.Client{Transport: transport}
	return NewAPIClientWithConfig(config)
}

func statusResponse(req *http.Request, status int, header http.Header, body *trackedBody) *http.Response {
	if header == nil {
		header = make(http.Header)
	}
	return &http.Response{StatusCode: status, Header: header, Body: body, Request: req}
}

func TestRetryClosesEveryRetriedBody(t *testing.T) {
	var bodies []*trackedBody
	var mu sync.Mutex
	transport := roundTripFunc(func(req *http.Request) (*http.Response, error) {
		mu.Lock()
		defer mu.Unlock()
		body := newTrackedBody(`{"message":"unavailable"}`)
		bodies = append(bodies, body)
		if len(bodies) < 3 {
			return statusResponse(req, http.StatusServiceUnavailable, nil, body), nil
		}
		return statusResponse(req, http.StatusOK, nil, body), nil
	})

	client := retryTestClient(transport, fastRetryConfig(10*time.Millisecond))
	_, _, err := client.UtilityAPI.Ping(context.Background())
	require.NoError(t, err)

	mu.Lock()
	defer mu.Unlock()
	require.Len(t, bodies, 3)
	assert.True(t, bodies[0].isClosed(), "first retried body must be closed")
	assert.True(t, bodies[1].isClosed(), "second retried body must be closed")
}

func TestRetryWaitsForRetryAfter(t *testing.T) {
	var attempts int32
	transport := roundTripFunc(func(req *http.Request) (*http.Response, error) {
		if atomic.AddInt32(&attempts, 1) == 1 {
			header := http.Header{"Retry-After": []string{"1"}}
			return statusResponse(req, http.StatusTooManyRequests, header, newTrackedBody(`{}`)), nil
		}
		return statusResponse(req, http.StatusOK, nil, newTrackedBody(`{}`)), nil
	})

	client := retryTestClient(transport, fastRetryConfig(5*time.Second))
	started := time.Now()
	_, _, err := client.UtilityAPI.Ping(context.Background())
	require.NoError(t, err)

	assert.Equal(t, int32(2), atomic.LoadInt32(&attempts))
	assert.GreaterOrEqual(t, time.Since(started), 900*time.Millisecond, "the retry must wait for Retry-After, not the 1ms backoff")
}

func TestRetryAfterLongerThanMaxDelayIsNotRetried(t *testing.T) {
	var attempts int32
	transport := roundTripFunc(func(req *http.Request) (*http.Response, error) {
		atomic.AddInt32(&attempts, 1)
		header := http.Header{"Retry-After": []string{"60"}}
		return statusResponse(req, http.StatusTooManyRequests, header, newTrackedBody(`{"message":"slow down"}`)), nil
	})

	client := retryTestClient(transport, fastRetryConfig(time.Second))
	_, _, err := client.UtilityAPI.Ping(context.Background())
	require.Error(t, err)

	assert.Equal(t, int32(1), atomic.LoadInt32(&attempts))
	var apiErr *APIError
	require.True(t, errors.As(err, &apiErr), "expected an *APIError, got %T", err)
	assert.Equal(t, http.StatusTooManyRequests, apiErr.StatusCode)
	assert.Equal(t, 60, apiErr.RetryAfter)
}

func TestRetryAfterIsNotRetriedPastTheMaxDelayLimit(t *testing.T) {
	tests := []struct {
		name       string
		maxDelay   time.Duration
		retryAfter string
		want       int
	}{
		// MaxDelay 0 caps every wait at 0, as it caps the backoff
		{name: "zero MaxDelay", maxDelay: 0, retryAfter: "120", want: 120},
		// Too large for a Duration: still past the limit, reported as sent
		{name: "beyond a Duration", maxDelay: time.Second, retryAfter: "9999999999", want: 9999999999},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var attempts int32
			transport := roundTripFunc(func(req *http.Request) (*http.Response, error) {
				atomic.AddInt32(&attempts, 1)
				header := http.Header{"Retry-After": []string{tt.retryAfter}}
				return statusResponse(req, http.StatusTooManyRequests, header, newTrackedBody(`{}`)), nil
			})

			client := retryTestClient(transport, fastRetryConfig(tt.maxDelay))
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			_, _, err := client.UtilityAPI.Ping(ctx)

			var apiErr *APIError
			require.True(t, errors.As(err, &apiErr), "expected an *APIError, got %v", err)
			assert.Equal(t, int32(1), atomic.LoadInt32(&attempts))
			assert.Equal(t, tt.want, apiErr.RetryAfter)
		})
	}
}

func TestRetryWaitEndsWhenContextIsCancelled(t *testing.T) {
	var attempts int32
	var first *trackedBody
	transport := roundTripFunc(func(req *http.Request) (*http.Response, error) {
		atomic.AddInt32(&attempts, 1)
		first = newTrackedBody(`{}`)
		header := http.Header{"Retry-After": []string{"10"}}
		return statusResponse(req, http.StatusServiceUnavailable, header, first), nil
	})

	client := retryTestClient(transport, fastRetryConfig(30*time.Second))
	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(50*time.Millisecond, cancel)

	started := time.Now()
	_, httpResp, err := client.UtilityAPI.Ping(ctx)

	require.Error(t, err)
	assert.ErrorIs(t, err, context.Canceled)
	assert.Nil(t, httpResp)
	assert.Less(t, time.Since(started), time.Second, "cancelling must end the 10s wait")
	assert.Equal(t, int32(1), atomic.LoadInt32(&attempts))
	assert.True(t, first.isClosed(), "the response held during the wait must be closed")
}

func TestEndedContextIsNotRetried(t *testing.T) {
	var attempts int32
	transport := roundTripFunc(func(req *http.Request) (*http.Response, error) {
		atomic.AddInt32(&attempts, 1)
		<-req.Context().Done()
		return nil, errors.New("connection reset")
	})

	// A zero backoff would let the loop start another attempt before it
	// notices the context. The request ends with its own error, after one
	// attempt.
	retry := fastRetryConfig(0)
	retry.BaseDelay = 0
	client := retryTestClient(transport, retry)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	_, _, err := client.UtilityAPI.Ping(ctx)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "connection reset")
	assert.Equal(t, int32(1), atomic.LoadInt32(&attempts))
}

// countingTransport counts the requests it passes to next.
type countingTransport struct {
	next  http.RoundTripper
	count int32
}

func (c *countingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	atomic.AddInt32(&c.count, 1)
	return c.next.RoundTrip(req)
}

func TestUntrustedCertificateIsNotRetried(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	serverURL, err := url.Parse(server.URL)
	require.NoError(t, err)

	// A transport that does not trust the test server's self-signed certificate.
	transport := &countingTransport{next: &http.Transport{}}
	client := retryTestClient(transport, fastRetryConfig(10*time.Millisecond))
	client.cfg.Host = serverURL.Host
	client.cfg.Scheme = serverURL.Scheme

	_, _, err = client.UtilityAPI.Ping(context.Background())

	require.Error(t, err)
	assert.True(t, isCertificateError(err), "expected a certificate error, got %v", err)
	assert.Equal(t, int32(1), atomic.LoadInt32(&transport.count))
}

func TestRetryAfterParsing(t *testing.T) {
	future := time.Now().Add(90 * time.Second).UTC().Format(http.TimeFormat)
	past := time.Now().Add(-time.Hour).UTC().Format(http.TimeFormat)

	tests := []struct {
		name     string
		status   int
		header   string
		wantOK   bool
		wantWait func(time.Duration) bool
	}{
		{"seconds on 429", http.StatusTooManyRequests, "7", true, func(d time.Duration) bool { return d == 7*time.Second }},
		{"seconds on 503", http.StatusServiceUnavailable, "2", true, func(d time.Duration) bool { return d == 2*time.Second }},
		{"HTTP date in the future", http.StatusTooManyRequests, future, true, func(d time.Duration) bool { return d > 80*time.Second && d <= 90*time.Second }},
		{"HTTP date in the past", http.StatusTooManyRequests, past, true, func(d time.Duration) bool { return d == 0 }},
		{"negative seconds", http.StatusTooManyRequests, "-1", false, nil},
		{"malformed", http.StatusTooManyRequests, "soon", false, nil},
		{"missing", http.StatusTooManyRequests, "", false, nil},
		{"other status", http.StatusInternalServerError, "5", false, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{StatusCode: tt.status, Header: http.Header{}}
			if tt.header != "" {
				resp.Header.Set("Retry-After", tt.header)
			}
			wait, ok := retryAfter(resp)
			assert.Equal(t, tt.wantOK, ok)
			if tt.wantWait != nil {
				assert.True(t, tt.wantWait(wait), "unexpected wait %v", wait)
			}
		})
	}
}

func TestServiceUnavailableErrorCarriesRetryAfter(t *testing.T) {
	resp := &http.Response{StatusCode: http.StatusServiceUnavailable, Header: http.Header{"Retry-After": []string{"4"}}}
	apiErr := ParseAPIError(resp, []byte(`{"message":"busy"}`))
	assert.Equal(t, ErrorTypeServer, apiErr.Type)
	assert.Equal(t, 4, apiErr.RetryAfter)
}
