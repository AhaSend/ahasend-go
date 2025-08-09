package ahasend_test

import (
	"testing"
	"time"

	"github.com/AhaSend/ahasend-go"
	"github.com/stretchr/testify/assert"
)

func TestRetryConfig(t *testing.T) {
	t.Run("DefaultRetryConfig", func(t *testing.T) {
		config := ahasend.DefaultRetryConfig()

		assert.True(t, config.Enabled)
		assert.Equal(t, 3, config.MaxRetries)
		assert.False(t, config.RetryClientErrors)
		assert.False(t, config.RetryValidationErrors)
		assert.Equal(t, ahasend.BackoffExponential, config.BackoffStrategy)
		assert.Equal(t, time.Second, config.BaseDelay)
		assert.Equal(t, 30*time.Second, config.MaxDelay)
	})

	t.Run("IsRetryEnabled", func(t *testing.T) {
		config := ahasend.RetryConfig{Enabled: true, MaxRetries: 3}
		assert.True(t, config.IsRetryEnabled())

		config.Enabled = false
		assert.False(t, config.IsRetryEnabled())

		config.Enabled = true
		config.MaxRetries = 0
		assert.False(t, config.IsRetryEnabled())
	})

	t.Run("GetDelay", func(t *testing.T) {
		t.Run("Exponential backoff", func(t *testing.T) {
			config := ahasend.RetryConfig{
				BackoffStrategy: ahasend.BackoffExponential,
				BaseDelay:       time.Second,
				MaxDelay:        30 * time.Second,
			}

			delay1 := config.GetDelay(1)
			delay2 := config.GetDelay(2)
			delay3 := config.GetDelay(3)

			// First delay should be around BaseDelay (with jitter)
			assert.GreaterOrEqual(t, delay1, time.Second*3/4)
			assert.LessOrEqual(t, delay1, time.Second*2)

			// Second delay should be larger than first
			assert.Greater(t, delay2, delay1)

			// Third delay should be larger than second
			assert.Greater(t, delay3, delay2)

			// No delay should exceed MaxDelay
			assert.LessOrEqual(t, delay1, 30*time.Second)
			assert.LessOrEqual(t, delay2, 30*time.Second)
			assert.LessOrEqual(t, delay3, 30*time.Second)
		})

		t.Run("Linear backoff", func(t *testing.T) {
			config := ahasend.RetryConfig{
				BackoffStrategy: ahasend.BackoffLinear,
				BaseDelay:       time.Second,
				MaxDelay:        30 * time.Second,
			}

			delay1 := config.GetDelay(1)
			delay2 := config.GetDelay(2)
			delay3 := config.GetDelay(3)

			assert.Equal(t, time.Second, delay1)
			assert.Equal(t, 2*time.Second, delay2)
			assert.Equal(t, 3*time.Second, delay3)
		})

		t.Run("Constant backoff", func(t *testing.T) {
			config := ahasend.RetryConfig{
				BackoffStrategy: ahasend.BackoffConstant,
				BaseDelay:       time.Second,
				MaxDelay:        30 * time.Second,
			}

			delay1 := config.GetDelay(1)
			delay2 := config.GetDelay(2)
			delay3 := config.GetDelay(3)

			assert.Equal(t, time.Second, delay1)
			assert.Equal(t, time.Second, delay2)
			assert.Equal(t, time.Second, delay3)
		})
	})
}

func TestConfigurationWithRetryConfig(t *testing.T) {
	t.Run("NewConfiguration applies RetryConfig defaults", func(t *testing.T) {
		cfg := ahasend.NewConfiguration()

		assert.True(t, cfg.RetryConfig.Enabled)
		assert.Equal(t, 3, cfg.RetryConfig.MaxRetries)
		assert.False(t, cfg.RetryConfig.RetryClientErrors)
		assert.False(t, cfg.RetryConfig.RetryValidationErrors)
	})

	t.Run("RetryConfig migration from MaxRetries", func(t *testing.T) {
		cfg := &ahasend.Configuration{MaxRetries: 5}
		ahasend.ApplyDefaults(cfg)

		// Should have both the legacy field and the new RetryConfig
		assert.Equal(t, 5, cfg.MaxRetries)
		assert.Equal(t, 5, cfg.RetryConfig.MaxRetries)
		assert.True(t, cfg.RetryConfig.Enabled)
	})
}

func TestClientRetryLogicWithRetryConfig(t *testing.T) {
	t.Run("APIClient uses RetryConfig for retry decisions", func(t *testing.T) {
		cfg := ahasend.NewConfiguration()
		cfg.RetryConfig.RetryClientErrors = true // Enable client error retries
		client := ahasend.NewAPIClient(cfg)

		// With RetryClientErrors enabled, should retry 4xx errors
		shouldRetry := client.GetConfig().RetryConfig.RetryClientErrors
		assert.True(t, shouldRetry)
	})

	t.Run("APIClient respects disabled retries", func(t *testing.T) {
		cfg := ahasend.NewConfiguration()
		cfg.RetryConfig.Enabled = false
		client := ahasend.NewAPIClient(cfg)

		assert.False(t, client.GetConfig().RetryConfig.IsRetryEnabled())
	})
}
