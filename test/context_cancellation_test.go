package ahasend_test

import (
	"context"
	"testing"
	"time"

	"github.com/AhaSend/ahasend-go/api"
	"github.com/stretchr/testify/assert"
)

func TestRateLimiterContextCancellation(t *testing.T) {
	t.Run("Context cancellation stops rate limiter wait", func(t *testing.T) {
		// Create a token bucket with very slow refill rate (1 token per 10 seconds)
		// Since RequestsPerSecond is int, use 1 RPS and rely on timing for the test
		config := api.RateLimitConfig{
			RequestsPerSecond: 1, // 1 token per second
			BurstCapacity:     1,
			Enabled:           true,
		}
		bucket := api.NewTokenBucket(config)

		// Consume the only token
		err := bucket.WaitForToken()
		assert.NoError(t, err)

		// Create a context that will cancel after 100ms
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		// Try to get another token - should timeout due to context cancellation
		start := time.Now()
		err = bucket.WaitForTokenWithContext(ctx)
		elapsed := time.Since(start)

		assert.Error(t, err)
		assert.Equal(t, context.DeadlineExceeded, err)
		assert.Less(t, elapsed, 200*time.Millisecond, "Should timeout quickly, not wait for token refill")
	})

	t.Run("Context cancellation with cancelled context", func(t *testing.T) {
		config := api.RateLimitConfig{
			RequestsPerSecond: 1, // 1 per second
			BurstCapacity:     1,
			Enabled:           true,
		}
		bucket := api.NewTokenBucket(config)

		// Consume the only token
		err := bucket.WaitForToken()
		assert.NoError(t, err)

		// Create and immediately cancel context
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		// Should immediately return with cancellation error
		start := time.Now()
		err = bucket.WaitForTokenWithContext(ctx)
		elapsed := time.Since(start)

		assert.Error(t, err)
		assert.Equal(t, context.Canceled, err)
		assert.Less(t, elapsed, 10*time.Millisecond, "Should return immediately")
	})

	t.Run("Rate limiter works normally without context cancellation", func(t *testing.T) {
		config := api.RateLimitConfig{
			RequestsPerSecond: 1000, // Very fast refill
			BurstCapacity:     5,
			Enabled:           true,
		}
		bucket := api.NewTokenBucket(config)

		ctx := context.Background()

		// Should get token immediately
		err := bucket.WaitForTokenWithContext(ctx)
		assert.NoError(t, err)
	})
}

func TestRateLimiterIntegrationContextCancellation(t *testing.T) {
	t.Run("RateLimiter WaitForTokenWithContext respects cancellation", func(t *testing.T) {
		rl := api.NewRateLimiter()

		// Set very slow rate limit for general API
		rl.SetRateLimit(api.GeneralAPI, 1, 1) // 1 RPS, 1 burst - will need to wait 1 second for next token

		// Use up the burst capacity
		err := rl.WaitForToken("GET", "/v2/accounts/123/messages")
		assert.NoError(t, err)

		// Create context that cancels quickly (before the 1-second refill)
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		// Should timeout due to context cancellation
		start := time.Now()
		err = rl.WaitForTokenWithContext(ctx, "GET", "/v2/accounts/123/messages")
		elapsed := time.Since(start)

		assert.Error(t, err)
		assert.Equal(t, context.DeadlineExceeded, err)
		assert.Less(t, elapsed, 200*time.Millisecond)
	})

	t.Run("Disabled rate limiter ignores context", func(t *testing.T) {
		rl := api.NewRateLimiter()
		rl.SetGlobalEnabled(false)

		// Even with cancelled context, should succeed immediately
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		err := rl.WaitForTokenWithContext(ctx, "GET", "/v2/accounts/123/messages")
		assert.NoError(t, err)
	})
}

func TestTokenBucketEdgeCases(t *testing.T) {
	t.Run("Disabled token bucket with context", func(t *testing.T) {
		config := api.RateLimitConfig{
			RequestsPerSecond: 1,
			BurstCapacity:     1,
			Enabled:           false, // Disabled
		}
		bucket := api.NewTokenBucket(config)

		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel immediately

		// Should succeed even with cancelled context when disabled
		err := bucket.WaitForTokenWithContext(ctx)
		assert.NoError(t, err)
	})

	t.Run("Token bucket with background context", func(t *testing.T) {
		config := api.RateLimitConfig{
			RequestsPerSecond: 1000,
			BurstCapacity:     5,
			Enabled:           true,
		}
		bucket := api.NewTokenBucket(config)

		// Should work with background context
		err := bucket.WaitForTokenWithContext(context.Background())
		assert.NoError(t, err)
	})

	t.Run("Token available immediately with context", func(t *testing.T) {
		config := api.RateLimitConfig{
			RequestsPerSecond: 1,
			BurstCapacity:     5,
			Enabled:           true,
		}
		bucket := api.NewTokenBucket(config)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancel()

		// Should succeed immediately without waiting (token available)
		err := bucket.WaitForTokenWithContext(ctx)
		assert.NoError(t, err)
	})
}
