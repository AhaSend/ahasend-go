package api

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/AhaSend/ahasend-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func noContentResponse(req *http.Request) *http.Response {
	return &http.Response{
		StatusCode: http.StatusNoContent,
		Header:     make(http.Header),
		Body:       http.NoBody,
		Request:    req,
	}
}

// TestClientCreation tests that we can create API client instances
func TestClientCreation(t *testing.T) {
	config := NewConfiguration()
	require.NotNil(t, config)

	client := NewAPIClientWithConfig(config)
	require.NotNil(t, client)

	// Test that all API services are initialized
	assert.NotNil(t, client.AccountsAPI)
	assert.NotNil(t, client.APIKeysAPI)
	assert.NotNil(t, client.ContactsAPI)
	assert.NotNil(t, client.DomainsAPI)
	assert.NotNil(t, client.MessagesAPI)
	assert.NotNil(t, client.RoutesAPI)
	assert.NotNil(t, client.SMTPCredentialsAPI)
	assert.NotNil(t, client.StatisticsAPI)
	assert.NotNil(t, client.SubAccountsAPI)
	assert.NotNil(t, client.SuppressionsAPI)
	assert.NotNil(t, client.UtilityAPI)
	assert.NotNil(t, client.WebhooksAPI)

	defaultClient := NewAPIClient()
	require.NotNil(t, defaultClient)
	assert.NotNil(t, defaultClient.ContactsAPI)
	assert.NotNil(t, defaultClient.SubAccountsAPI)
}

func TestExecuteHonorsRequestTimeout(t *testing.T) {
	const requestTimeout = 25 * time.Millisecond

	var requestDeadline time.Time
	transport := roundTripFunc(func(req *http.Request) (*http.Response, error) {
		requestDeadline, _ = req.Context().Deadline()

		select {
		case <-req.Context().Done():
			return nil, req.Context().Err()
		case <-time.After(time.Second):
			return nil, errors.New("request context was not canceled")
		}
	})

	config := NewConfiguration()
	config.APIKey = "test-key"
	config.EnableRateLimit = false
	config.RetryConfig = RetryConfig{}
	config.HTTPClient = &http.Client{Transport: transport}
	client := NewAPIClientWithConfig(config)

	started := time.Now()
	_, _, err := client.UtilityAPI.Ping(context.Background(), WithTimeout(requestTimeout))
	elapsed := time.Since(started)

	require.Error(t, err)
	assert.ErrorIs(t, err, context.DeadlineExceeded)
	assert.False(t, requestDeadline.IsZero(), "request context should have a deadline")
	assert.Less(t, elapsed, 500*time.Millisecond)
}

func TestExecuteRequestTimeoutPreservesEarlierCallerDeadline(t *testing.T) {
	const callerTimeout = 25 * time.Millisecond

	var requestDeadline time.Time
	transport := roundTripFunc(func(req *http.Request) (*http.Response, error) {
		requestDeadline, _ = req.Context().Deadline()
		<-req.Context().Done()
		return nil, req.Context().Err()
	})

	config := NewConfiguration()
	config.APIKey = "test-key"
	config.EnableRateLimit = false
	config.RetryConfig = RetryConfig{}
	config.HTTPClient = &http.Client{Transport: transport}
	client := NewAPIClientWithConfig(config)

	ctx, cancel := context.WithTimeout(context.Background(), callerTimeout)
	defer cancel()
	callerDeadline, ok := ctx.Deadline()
	require.True(t, ok)

	_, _, err := client.UtilityAPI.Ping(ctx, WithTimeout(time.Second))

	require.Error(t, err)
	assert.ErrorIs(t, err, context.DeadlineExceeded)
	assert.Equal(t, callerDeadline, requestDeadline)
}

func TestExecuteRequestTimeoutOverridesHTTPClientTimeout(t *testing.T) {
	const (
		clientTimeout  = 10 * time.Millisecond
		requestTimeout = time.Second
	)

	var requestDeadline time.Time
	transport := roundTripFunc(func(req *http.Request) (*http.Response, error) {
		requestDeadline, _ = req.Context().Deadline()
		return noContentResponse(req), nil
	})

	config := NewConfiguration()
	config.APIKey = "test-key"
	config.EnableRateLimit = false
	config.RetryConfig = RetryConfig{}
	config.HTTPClient = &http.Client{
		Transport: transport,
		Timeout:   clientTimeout,
	}
	client := NewAPIClientWithConfig(config)

	started := time.Now()
	_, _, err := client.UtilityAPI.Ping(context.Background(), WithTimeout(requestTimeout))

	require.NoError(t, err)
	assert.Greater(t, requestDeadline.Sub(started), requestTimeout/2)
	assert.Equal(t, clientTimeout, config.HTTPClient.Timeout, "shared client timeout should remain unchanged")
}

func TestExecuteRequestTimeoutDuringRateLimit(t *testing.T) {
	transportCalls := 0
	transport := roundTripFunc(func(req *http.Request) (*http.Response, error) {
		transportCalls++
		return noContentResponse(req), nil
	})

	config := NewConfiguration()
	config.APIKey = "test-key"
	config.RetryConfig = RetryConfig{}
	config.HTTPClient = &http.Client{Transport: transport}
	client := NewAPIClientWithConfig(config)
	client.SetGeneralRateLimit(1, 1)

	_, _, err := client.UtilityAPI.Ping(context.Background())
	require.NoError(t, err)

	_, _, err = client.UtilityAPI.Ping(context.Background(), WithTimeout(25*time.Millisecond))

	require.Error(t, err)
	assert.ErrorIs(t, err, context.DeadlineExceeded)
	var networkErr *NetworkError
	require.ErrorAs(t, err, &networkErr)
	assert.Equal(t, "rate limiting", networkErr.Op)
	assert.Equal(t, 1, transportCalls, "timed-out request should not reach the transport")
}

// TestUtilityFunctions tests the pointer utility functions
func TestUtilityFunctions(t *testing.T) {
	// Test PtrString
	str := "test"
	ptrStr := ahasend.String(str)
	assert.NotNil(t, ptrStr)
	assert.Equal(t, str, *ptrStr)

	// Test PtrBool
	b := true
	ptrBool := ahasend.Bool(b)
	assert.NotNil(t, ptrBool)
	assert.Equal(t, b, *ptrBool)

	// Test PtrInt
	i := 42
	ptrInt := ahasend.Int(i)
	assert.NotNil(t, ptrInt)
	assert.Equal(t, i, *ptrInt)
}
