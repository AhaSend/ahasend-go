package api

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConfigurationDefaults(t *testing.T) {
	defaults := GetDefaults()

	assert.Equal(t, "api.ahasend.com", defaults.DefaultHost)
	assert.Equal(t, "https", defaults.DefaultScheme)
	assert.Equal(t, "AhaSend-Go-SDK/1.0", defaults.DefaultUserAgent)
	assert.False(t, defaults.DefaultDebug)
	assert.Equal(t, 30*time.Second, defaults.DefaultTimeout)
	assert.True(t, defaults.DefaultEnableRateLimit)
	assert.Equal(t, 3, defaults.DefaultMaxRetries)
	assert.Equal(t, 100, defaults.DefaultGeneralRPS)
	assert.Equal(t, 200, defaults.DefaultGeneralBurst)
	assert.Equal(t, 1, defaults.DefaultStatisticsRPS)
	assert.Equal(t, 1, defaults.DefaultStatisticsBurst)
	assert.True(t, defaults.DefaultIdempotencyAutoGenerate)
}

func TestValidateConfiguration(t *testing.T) {
	t.Run("Nil configuration", func(t *testing.T) {
		result := ValidateConfiguration(nil)
		assert.False(t, result.Valid)
		assert.True(t, result.HasErrors())
		assert.Contains(t, result.Error(), "configuration cannot be nil")
	})

	t.Run("Valid configuration", func(t *testing.T) {
		cfg := NewConfiguration()
		result := ValidateConfiguration(cfg)
		assert.True(t, result.Valid)
		assert.False(t, result.HasErrors())
	})

	t.Run("Invalid scheme", func(t *testing.T) {
		cfg := NewConfiguration()
		cfg.Scheme = "ftp"
		result := ValidateConfiguration(cfg)
		assert.False(t, result.Valid)
		assert.True(t, result.HasErrors())
		assert.Contains(t, result.Error(), "must be 'http' or 'https'")
	})

	t.Run("Host with scheme", func(t *testing.T) {
		cfg := NewConfiguration()
		cfg.Host = "https://api.example.com"
		result := ValidateConfiguration(cfg)
		assert.False(t, result.Valid)
		assert.True(t, result.HasErrors())
		assert.Contains(t, result.Error(), "should not include scheme")
	})

	t.Run("Invalid hostname characters", func(t *testing.T) {
		cfg := NewConfiguration()
		cfg.Host = "api@example.com"
		result := ValidateConfiguration(cfg)
		assert.False(t, result.Valid)
		assert.True(t, result.HasErrors())
		assert.Contains(t, result.Error(), "contains invalid characters")
	})

	t.Run("Negative max retries", func(t *testing.T) {
		cfg := NewConfiguration()
		cfg.RetryConfig.MaxRetries = -1
		result := ValidateConfiguration(cfg)
		assert.False(t, result.Valid)
		assert.True(t, result.HasErrors())
		assert.Contains(t, result.Error(), "cannot be negative")
	})

	t.Run("High max retries warning", func(t *testing.T) {
		cfg := NewConfiguration()
		cfg.RetryConfig.MaxRetries = 15
		result := ValidateConfiguration(cfg)
		assert.True(t, result.Valid)
		assert.True(t, result.HasWarnings())
		assert.Contains(t, result.Warnings[0], "very high")
	})

	t.Run("Invalid rate limit config", func(t *testing.T) {
		cfg := NewConfiguration()
		cfg.DefaultGeneralRateLimit = &RateLimitConfig{
			RequestsPerSecond: -1,
			BurstCapacity:     0,
		}
		result := ValidateConfiguration(cfg)
		assert.False(t, result.Valid)
		assert.True(t, result.HasErrors())
	})

	t.Run("Invalid idempotency prefix", func(t *testing.T) {
		cfg := NewConfiguration()
		cfg.IdempotencyConfig.KeyPrefix = "invalid@prefix"
		result := ValidateConfiguration(cfg)
		assert.False(t, result.Valid)
		assert.True(t, result.HasErrors())
		assert.Contains(t, result.Error(), "can only contain letters, numbers, hyphens, and underscores")
	})

	t.Run("Invalid default headers", func(t *testing.T) {
		cfg := NewConfiguration()
		cfg.DefaultHeader = map[string]string{
			"":          "value",
			"key\n":     "value",
			"valid-key": "value\r\n",
		}
		result := ValidateConfiguration(cfg)
		assert.False(t, result.Valid)
		assert.True(t, result.HasErrors())
		assert.Len(t, result.Errors, 3)
	})
}

func TestApplyDefaults(t *testing.T) {
	t.Run("Apply to empty config", func(t *testing.T) {
		cfg := &Configuration{}
		ApplyDefaults(cfg)

		assert.Equal(t, "api.ahasend.com", cfg.Host)
		assert.Equal(t, "https", cfg.Scheme)
		assert.Equal(t, "AhaSend-Go-SDK/1.0", cfg.UserAgent)
		assert.NotNil(t, cfg.DefaultGeneralRateLimit)
		assert.NotNil(t, cfg.DefaultStatisticsRateLimit)
		assert.NotNil(t, cfg.DefaultSendMessageRateLimit)
		assert.NotNil(t, cfg.HTTPClient)
		assert.NotEmpty(t, cfg.Servers)
	})

	t.Run("Don't override existing values", func(t *testing.T) {
		cfg := &Configuration{
			Host:        "custom.host.com",
			Scheme:      "http",
			UserAgent:   "CustomAgent/1.0",
			RetryConfig: RetryConfig{MaxRetries: 5},
		}
		ApplyDefaults(cfg)

		assert.Equal(t, "custom.host.com", cfg.Host)
		assert.Equal(t, "http", cfg.Scheme)
		assert.Equal(t, "CustomAgent/1.0", cfg.UserAgent)
		assert.Equal(t, 5, cfg.RetryConfig.MaxRetries)
	})

	t.Run("Nil config handling", func(t *testing.T) {
		// Should not panic
		ApplyDefaults(nil)
	})
}

func TestValidateAndApplyDefaults(t *testing.T) {
	cfg := &Configuration{}
	result := ValidateAndApplyDefaults(cfg)

	assert.True(t, result.Valid)
	assert.Equal(t, "api.ahasend.com", cfg.Host)
	assert.Equal(t, 3, cfg.RetryConfig.MaxRetries)
}

func TestNewValidatedConfiguration(t *testing.T) {
	cfg, result := NewValidatedConfiguration()

	assert.NotNil(t, cfg)
	assert.True(t, result.Valid)
	assert.Equal(t, "api.ahasend.com", cfg.Host)
}

func TestNewValidatedConfigurationFromEnv(t *testing.T) {
	withEnv(t, "AHASEND_DEBUG", "true", func() {
		cfg, result := NewValidatedConfigurationFromEnv()

		assert.NotNil(t, cfg)
		assert.True(t, result.Valid)
		assert.True(t, cfg.Debug)
	})
}

func TestGetConfigurationSummary(t *testing.T) {
	cfg := NewConfiguration()
	summary := GetConfigurationSummary(cfg)

	assert.Equal(t, "https://api.ahasend.com", summary.ServerURL)
	assert.Equal(t, "AhaSend-Go-SDK/1.0", summary.UserAgent)
	assert.False(t, summary.Debug)
	assert.Equal(t, 3, summary.MaxRetries)
	assert.True(t, summary.RateLimitEnabled)
	assert.True(t, summary.IdempotencyEnabled)
	assert.Equal(t, 30, summary.TimeoutSeconds)
}

func TestIsProductionReady(t *testing.T) {
	t.Run("Development config", func(t *testing.T) {
		cfg := NewConfiguration()
		cfg.Debug = true
		cfg.EnableRateLimit = false
		cfg.RetryConfig = RetryConfig{MaxRetries: 0}
		cfg.RetryConfig.Enabled = false // Disable retries to trigger warning
		cfg.Scheme = "http"

		ready, issues := IsProductionReady(cfg)
		assert.False(t, ready)
		assert.NotEmpty(t, issues)
		assert.Len(t, issues, 4) // Debug, rate limit, retries disabled, scheme
	})

	t.Run("Production config", func(t *testing.T) {
		cfg := NewConfiguration()
		cfg.UserAgent = "MyApp/1.0"

		ready, issues := IsProductionReady(cfg)
		assert.True(t, ready)
		assert.Empty(t, issues)
	})
}

func TestOptimizeForProduction(t *testing.T) {
	cfg := &Configuration{
		Debug:           true,
		EnableRateLimit: false,
		RetryConfig:     RetryConfig{MaxRetries: 0},
		Scheme:          "http",
	}

	OptimizeForProduction(cfg)

	assert.False(t, cfg.Debug)
	assert.True(t, cfg.EnableRateLimit)
	assert.Equal(t, 3, cfg.RetryConfig.MaxRetries)
	assert.Equal(t, "https", cfg.Scheme)
	assert.True(t, cfg.IdempotencyConfig.AutoGenerate)
}

func TestValidationResultMethods(t *testing.T) {
	t.Run("HasErrors and HasWarnings", func(t *testing.T) {
		result := ValidationResult{}
		assert.False(t, result.HasErrors())
		assert.False(t, result.HasWarnings())

		result.Errors = append(result.Errors, ConfigurationValidationError{
			Field: "test", Value: "value", Message: "error",
		})
		result.Warnings = append(result.Warnings, "warning")

		assert.True(t, result.HasErrors())
		assert.True(t, result.HasWarnings())
	})

	t.Run("Error formatting", func(t *testing.T) {
		result := ValidationResult{
			Errors: []ConfigurationValidationError{
				{Field: "field1", Value: "value1", Message: "error1"},
				{Field: "field2", Value: "value2", Message: "error2"},
			},
		}

		errMsg := result.Error()
		assert.Contains(t, errMsg, "configuration validation failed")
		assert.Contains(t, errMsg, "field1")
		assert.Contains(t, errMsg, "field2")
	})

	t.Run("No errors", func(t *testing.T) {
		result := ValidationResult{}
		assert.Empty(t, result.Error())
	})
}

func TestConfigurationValidationError(t *testing.T) {
	err := ConfigurationValidationError{
		Field:   "TestField",
		Value:   "test-value",
		Message: "test error message",
	}

	expected := "configuration validation error for TestField (test-value): test error message"
	assert.Equal(t, expected, err.Error())
}
