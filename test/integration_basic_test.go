package ahasend_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/AhaSend/ahasend-go/api"
	"github.com/stretchr/testify/assert"
)

// Basic integration test to verify the mock server infrastructure works
func TestBasicIntegrationInfrastructure(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION_TESTS") == "true" {
		t.Skip("Integration tests are disabled")
	}

	// This is a basic test to verify our integration test infrastructure
	// without depending on specific API models that may change
	t.Run("Client creation and configuration", func(t *testing.T) {
		// Test that we can create a client with test configuration
		cfg := api.NewConfiguration()
		cfg.Host = "localhost:4010"
		cfg.Scheme = "http"
		cfg.EnableRateLimit = false
		cfg.Debug = false

		client := api.NewAPIClientWithConfig(cfg)
		assert.NotNil(t, client)

		// Verify configuration
		assert.Equal(t, "localhost:4010", client.GetConfig().Host)
		assert.Equal(t, "http", client.GetConfig().Scheme)
		assert.False(t, client.GetConfig().EnableRateLimit)
	})

	t.Run("Authentication context creation", func(t *testing.T) {
		// Test that we can create authenticated contexts
		testAPIKey := "aha-sk-test-key-for-integration-testing-12345678901234567890"
		ctx := context.WithValue(context.Background(), api.ContextAccessToken, testAPIKey)

		// Verify the context contains the API key
		value := ctx.Value(api.ContextAccessToken)
		assert.Equal(t, testAPIKey, value)
	})

	t.Run("Rate limiting configuration", func(t *testing.T) {
		cfg := api.NewConfiguration()
		cfg.EnableRateLimit = true

		client := api.NewAPIClientWithConfig(cfg)

		// Test rate limit configuration methods
		client.SetGeneralRateLimit(50, 100)
		client.SetStatisticsRateLimit(1, 2)
		client.SetSendMessageRateLimit(25, 50)

		// Verify we can get rate limit status
		generalStatus := client.GetRateLimitStatus(api.GeneralAPI)
		assert.Equal(t, api.GeneralAPI, generalStatus.EndpointType)
		assert.True(t, generalStatus.Enabled)
		assert.Equal(t, 50, generalStatus.RequestsPerSecond)
		assert.Equal(t, 100, generalStatus.BurstCapacity)

		statisticsStatus := client.GetRateLimitStatus(api.StatisticsAPI)
		assert.Equal(t, api.StatisticsAPI, statisticsStatus.EndpointType)
		assert.Equal(t, 1, statisticsStatus.RequestsPerSecond)
		assert.Equal(t, 2, statisticsStatus.BurstCapacity)
	})

	t.Run("Idempotency key generation", func(t *testing.T) {
		client := api.NewAPIClientWithConfig(api.NewConfiguration())

		// Test basic key generation
		key1 := client.GenerateIdempotencyKey()
		key2 := client.GenerateIdempotencyKey()

		assert.NotEmpty(t, key1)
		assert.NotEmpty(t, key2)
		assert.NotEqual(t, key1, key2)
		assert.Len(t, key1, 36) // UUID length
		assert.Len(t, key2, 36) // UUID length

		// Test key builder
		builder := client.NewIdempotencyKeyBuilder()
		builtKey := builder.WithSuffix("test")
		assert.NotEmpty(t, builtKey)
		assert.Contains(t, builtKey, "-test")
	})

	t.Run("Configuration validation", func(t *testing.T) {
		cfg := api.NewConfiguration()

		// Test validation
		result := api.ValidateConfiguration(cfg)
		assert.False(t, result.HasErrors())

		// Test with invalid configuration
		cfg.Scheme = "invalid-scheme"
		result = api.ValidateConfiguration(cfg)
		assert.True(t, result.HasErrors())
	})

	t.Run("Environment configuration", func(t *testing.T) {
		// Test environment variable support
		os.Setenv("AHASEND_API_KEY", "test-key-123")
		os.Setenv("AHASEND_DEBUG", "true")
		os.Setenv("AHASEND_MAX_RETRIES", "5")
		defer func() {
			os.Unsetenv("AHASEND_API_KEY")
			os.Unsetenv("AHASEND_DEBUG")
			os.Unsetenv("AHASEND_MAX_RETRIES")
		}()

		cfg := api.ConfigFromEnv()
		assert.True(t, cfg.Debug)
		assert.Equal(t, 5, cfg.RetryConfig.MaxRetries)

		// Test authenticated context from environment
		ctx := api.ContextWithEnvAuth(context.Background())
		apiKey := ctx.Value(api.ContextAccessToken)
		assert.Equal(t, "test-key-123", apiKey)
	})
}

// Test error handling without making actual API calls
func TestErrorHandlingInfrastructure(t *testing.T) {
	t.Run("API Error types", func(t *testing.T) {
		// Test that we can create and work with different error types
		authErr := &api.APIError{
			Type:       api.ErrorTypeAuthentication,
			StatusCode: 401,
			Message:    "Invalid API key",
		}

		assert.Equal(t, api.ErrorTypeAuthentication, authErr.Type)
		assert.False(t, authErr.IsRetryable())
		assert.Equal(t, 401, authErr.StatusCode)

		rateLimitErr := &api.APIError{
			Type:       api.ErrorTypeRateLimit,
			StatusCode: 429,
			Message:    "Rate limit exceeded",
		}

		assert.Equal(t, api.ErrorTypeRateLimit, rateLimitErr.Type)
		assert.True(t, rateLimitErr.IsRetryable())
	})

	t.Run("Network Error", func(t *testing.T) {
		netErr := &api.NetworkError{
			Op:  "GET /test",
			Err: context.DeadlineExceeded,
		}

		assert.Contains(t, netErr.Error(), "GET /test")
		assert.Contains(t, netErr.Error(), "context deadline exceeded")
	})

	t.Run("Error types", func(t *testing.T) {
		apiErr := &api.APIError{Type: api.ErrorTypeValidation}
		netErr := &api.NetworkError{Op: "test", Err: context.Canceled}

		// Test APIError type assertion
		var testAPIErr *api.APIError
		assert.True(t, errors.As(apiErr, &testAPIErr))
		assert.False(t, errors.As(netErr, &testAPIErr))

		// Test NetworkError type assertion
		var testNetErr *api.NetworkError
		assert.True(t, errors.As(netErr, &testNetErr))
		assert.False(t, errors.As(apiErr, &testNetErr))

		// Test IsRetryable method
		assert.True(t, netErr.IsRetryable())
		assert.False(t, apiErr.IsRetryable())
	})
}

// Test context cancellation functionality
func TestContextCancellationInfrastructure(t *testing.T) {
	t.Run("Context cancellation with rate limiter", func(t *testing.T) {
		// Create rate limiter and consume burst capacity first
		rateLimiter := api.NewRateLimiter()
		rateLimiter.SetRateLimit(api.GeneralAPI, 1, 1) // 1 req/s, 1 burst

		// Consume the single token
		err := rateLimiter.WaitForToken("GET", "/test")
		assert.NoError(t, err)

		// Now create cancelled context
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel immediately

		// This should return quickly with context canceled error since no tokens available
		err = rateLimiter.WaitForTokenWithContext(ctx, "GET", "/test")
		assert.Error(t, err)
		assert.Equal(t, context.Canceled, err)
	})
}

// Benchmark basic SDK operations
func BenchmarkBasicOperations(b *testing.B) {
	client := api.NewAPIClientWithConfig(api.NewConfiguration())

	b.Run("GenerateIdempotencyKey", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = client.GenerateIdempotencyKey()
		}
	})

	b.Run("ConfigValidation", func(b *testing.B) {
		cfg := api.NewConfiguration()
		for i := 0; i < b.N; i++ {
			_ = api.ValidateConfiguration(cfg)
		}
	})

	b.Run("ContextCreation", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = context.WithValue(context.Background(), api.ContextAccessToken, "test-key")
		}
	})
}
