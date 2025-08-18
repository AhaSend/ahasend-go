package api

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Helper to set and clean up environment variables
func withEnv(t *testing.T, key, value string, fn func()) {
	old := os.Getenv(key)
	defer func() {
		if old == "" {
			os.Unsetenv(key)
		} else {
			os.Setenv(key, old)
		}
	}()

	if value == "" {
		os.Unsetenv(key)
	} else {
		os.Setenv(key, value)
	}

	fn()
}

func TestConfigFromEnv(t *testing.T) {
	t.Run("Default values when no env vars", func(t *testing.T) {
		// Ensure no relevant env vars are set
		envVars := []string{
			"AHASEND_API_KEY", "AHASEND_TOKEN", "AHASEND_BASE_URL",
			"AHASEND_HOST", "AHASEND_SCHEME", "AHASEND_DEBUG",
		}
		for _, key := range envVars {
			withEnv(t, key, "", func() {})
		}

		cfg := ConfigFromEnv()

		// Should have default values
		assert.False(t, cfg.Debug, "Debug should default to false")
		assert.Equal(t, "AhaSend-Go-SDK/1.0", cfg.UserAgent, "Should have default user agent")
		assert.True(t, cfg.EnableRateLimit, "Should enable rate limiting by default")
		assert.Equal(t, 3, cfg.RetryConfig.MaxRetries, "Should have default max retries")
		assert.True(t, cfg.IdempotencyConfig.AutoGenerate, "Should auto-generate idempotency keys by default")
	})

	t.Run("Load server configuration", func(t *testing.T) {
		withEnv(t, "AHASEND_BASE_URL", "https://api.staging.ahasend.com", func() {
			cfg := ConfigFromEnv()
			assert.Equal(t, "https", cfg.Scheme, "Should extract scheme from base URL")
			assert.Equal(t, "api.staging.ahasend.com", cfg.Host, "Should extract host from base URL")
		})

		withEnv(t, "AHASEND_HOST", "custom.host.com", func() {
			withEnv(t, "AHASEND_SCHEME", "http", func() {
				cfg := ConfigFromEnv()
				assert.Equal(t, "http", cfg.Scheme, "Should use scheme from env")
				assert.Equal(t, "custom.host.com", cfg.Host, "Should use host from env")
			})
		})
	})

	t.Run("Load debug configuration", func(t *testing.T) {
		testCases := []struct {
			value    string
			expected bool
		}{
			{"true", true},
			{"1", true},
			{"yes", true},
			{"on", true},
			{"enable", true},
			{"enabled", true},
			{"false", false},
			{"0", false},
			{"no", false},
			{"off", false},
			{"disable", false},
			{"disabled", false},
		}

		for _, tc := range testCases {
			withEnv(t, "AHASEND_DEBUG", tc.value, func() {
				cfg := ConfigFromEnv()
				assert.Equal(t, tc.expected, cfg.Debug, "Debug value for %s", tc.value)
			})
		}
	})

	t.Run("Load rate limiting configuration", func(t *testing.T) {
		withEnv(t, "AHASEND_ENABLE_RATE_LIMIT", "false", func() {
			withEnv(t, "AHASEND_MAX_RETRIES", "5", func() {
				cfg := ConfigFromEnv()
				assert.False(t, cfg.EnableRateLimit, "Should disable rate limiting")
				assert.Equal(t, 5, cfg.RetryConfig.MaxRetries, "Should set max retries")
			})
		})
	})

	t.Run("Load idempotency configuration", func(t *testing.T) {
		withEnv(t, "AHASEND_IDEMPOTENCY_AUTO_GENERATE", "false", func() {
			withEnv(t, "AHASEND_IDEMPOTENCY_PREFIX", "test", func() {
				cfg := ConfigFromEnv()
				assert.False(t, cfg.IdempotencyConfig.AutoGenerate, "Should disable auto-generation")
				assert.Equal(t, "test", cfg.IdempotencyConfig.KeyPrefix, "Should set prefix")
			})
		})
	})

	t.Run("Custom user agent", func(t *testing.T) {
		withEnv(t, "AHASEND_USER_AGENT", "MyApp/1.0", func() {
			cfg := ConfigFromEnv()
			assert.Equal(t, "MyApp/1.0", cfg.UserAgent, "Should use custom user agent")
		})
	})
}

func TestGetAPIKeyFromEnv(t *testing.T) {
	t.Run("No API key set", func(t *testing.T) {
		withEnv(t, "AHASEND_API_KEY", "", func() {
			withEnv(t, "AHASEND_TOKEN", "", func() {
				key := GetAPIKeyFromEnv()
				assert.Empty(t, key, "Should return empty string when no key set")
			})
		})
	})

	t.Run("API key takes precedence", func(t *testing.T) {
		withEnv(t, "AHASEND_API_KEY", "aha-sk-key", func() {
			withEnv(t, "AHASEND_TOKEN", "aha-sk-token", func() {
				key := GetAPIKeyFromEnv()
				assert.Equal(t, "aha-sk-key", key, "Should prefer API_KEY over TOKEN")
			})
		})
	})

	t.Run("Fallback to token", func(t *testing.T) {
		withEnv(t, "AHASEND_API_KEY", "", func() {
			withEnv(t, "AHASEND_TOKEN", "aha-sk-token", func() {
				key := GetAPIKeyFromEnv()
				assert.Equal(t, "aha-sk-token", key, "Should fallback to TOKEN")
			})
		})
	})
}

func TestContextWithEnvAuth(t *testing.T) {
	t.Run("With API key", func(t *testing.T) {
		withEnv(t, "AHASEND_API_KEY", "test-key", func() {
			ctx := context.Background()
			authCtx := ContextWithEnvAuth(ctx)

			value := authCtx.Value(ContextAccessToken)
			assert.Equal(t, "test-key", value, "Should add API key to context")
		})
	})

	t.Run("Without API key", func(t *testing.T) {
		withEnv(t, "AHASEND_API_KEY", "", func() {
			withEnv(t, "AHASEND_TOKEN", "", func() {
				ctx := context.Background()
				authCtx := ContextWithEnvAuth(ctx)

				value := authCtx.Value(ContextAccessToken)
				assert.Nil(t, value, "Should not modify context when no key")
			})
		})
	})
}

func TestGetTimeoutFromEnv(t *testing.T) {
	t.Run("No timeouts set", func(t *testing.T) {
		withEnv(t, "AHASEND_TIMEOUT", "", func() {
			withEnv(t, "AHASEND_CONNECT_TIMEOUT", "", func() {
				timeout, connectTimeout := GetTimeoutFromEnv()
				assert.Equal(t, time.Duration(0), timeout, "Should return zero timeout")
				assert.Equal(t, time.Duration(0), connectTimeout, "Should return zero connect timeout")
			})
		})
	})

	t.Run("Set timeouts", func(t *testing.T) {
		withEnv(t, "AHASEND_TIMEOUT", "30", func() {
			withEnv(t, "AHASEND_CONNECT_TIMEOUT", "10", func() {
				timeout, connectTimeout := GetTimeoutFromEnv()
				assert.Equal(t, 30*time.Second, timeout, "Should parse timeout")
				assert.Equal(t, 10*time.Second, connectTimeout, "Should parse connect timeout")
			})
		})
	})

	t.Run("Invalid values", func(t *testing.T) {
		withEnv(t, "AHASEND_TIMEOUT", "invalid", func() {
			withEnv(t, "AHASEND_CONNECT_TIMEOUT", "-5", func() {
				timeout, connectTimeout := GetTimeoutFromEnv()
				assert.Equal(t, time.Duration(0), timeout, "Should ignore invalid timeout")
				assert.Equal(t, time.Duration(0), connectTimeout, "Should ignore negative timeout")
			})
		})
	})
}

func TestLoadEnvIntoConfig(t *testing.T) {
	t.Run("Load into existing config", func(t *testing.T) {
		cfg := NewConfiguration()
		cfg.Debug = true               // Start with debug enabled
		cfg.RetryConfig.MaxRetries = 1 // Start with 1 retry

		withEnv(t, "AHASEND_DEBUG", "false", func() {
			withEnv(t, "AHASEND_MAX_RETRIES", "5", func() {
				LoadEnvIntoConfig(cfg)

				assert.False(t, cfg.Debug, "Should override debug setting")
				assert.Equal(t, 5, cfg.RetryConfig.MaxRetries, "Should override max retries")
			})
		})
	})

	t.Run("Nil config handling", func(t *testing.T) {
		// Should not panic
		LoadEnvIntoConfig(nil)
	})
}

func TestNewConfigurationFromEnv(t *testing.T) {
	withEnv(t, "AHASEND_DEBUG", "true", func() {
		cfg := NewConfigurationFromEnv()
		assert.True(t, cfg.Debug, "Should load from environment")
		assert.NotNil(t, cfg, "Should return valid configuration")
	})
}

func TestNewAPIClientFromEnv(t *testing.T) {
	withEnv(t, "AHASEND_DEBUG", "true", func() {
		client := NewAPIClientFromEnv()
		assert.NotNil(t, client, "Should create client")
		assert.True(t, client.GetConfig().Debug, "Should have debug enabled from env")
	})
}

func TestGetEnvDocumentation(t *testing.T) {
	docs := GetEnvDocumentation()

	assert.NotEmpty(t, docs, "Should return documentation")
	assert.Contains(t, docs, "AHASEND_API_KEY", "Should document API key")
	assert.Contains(t, docs, "AHASEND_DEBUG", "Should document debug flag")
	assert.Contains(t, docs, "AHASEND_MAX_RETRIES", "Should document max retries")

	// Verify all keys have descriptions
	for key, desc := range docs {
		assert.NotEmpty(t, key, "Key should not be empty")
		assert.NotEmpty(t, desc, "Description should not be empty")
	}
}

func TestValidateEnvConfig(t *testing.T) {
	t.Run("Valid configuration", func(t *testing.T) {
		withEnv(t, "AHASEND_API_KEY", "aha-sk-valid-key", func() {
			withEnv(t, "AHASEND_SCHEME", "https", func() {
				withEnv(t, "AHASEND_DEBUG", "true", func() {
					withEnv(t, "AHASEND_MAX_RETRIES", "3", func() {
						issues := ValidateEnvConfig()
						assert.Empty(t, issues, "Should have no validation issues")
					})
				})
			})
		})
	})

	t.Run("Missing API key", func(t *testing.T) {
		withEnv(t, "AHASEND_API_KEY", "", func() {
			withEnv(t, "AHASEND_TOKEN", "", func() {
				issues := ValidateEnvConfig()
				assert.Len(t, issues, 1, "Should have one validation issue")
				assert.Contains(t, issues[0], "No API key found", "Should mention missing API key")
			})
		})
	})

	t.Run("Invalid scheme", func(t *testing.T) {
		withEnv(t, "AHASEND_API_KEY", "aha-sk-key", func() {
			withEnv(t, "AHASEND_SCHEME", "ftp", func() {
				issues := ValidateEnvConfig()
				assert.Contains(t, issues[0], "AHASEND_SCHEME must be", "Should validate scheme")
			})
		})
	})

	t.Run("Invalid boolean", func(t *testing.T) {
		withEnv(t, "AHASEND_API_KEY", "aha-sk-key", func() {
			withEnv(t, "AHASEND_DEBUG", "maybe", func() {
				issues := ValidateEnvConfig()
				found := false
				for _, issue := range issues {
					if strings.Contains(issue, "AHASEND_DEBUG") {
						found = true
						break
					}
				}
				assert.True(t, found, "Should validate boolean values")
			})
		})
	})

	t.Run("Invalid integer", func(t *testing.T) {
		withEnv(t, "AHASEND_API_KEY", "aha-sk-key", func() {
			withEnv(t, "AHASEND_MAX_RETRIES", "not-a-number", func() {
				issues := ValidateEnvConfig()
				found := false
				for _, issue := range issues {
					if strings.Contains(issue, "AHASEND_MAX_RETRIES") {
						found = true
						break
					}
				}
				assert.True(t, found, "Should validate integer values")
			})
		})
	})
}

func TestBooleanParsing(t *testing.T) {
	testCases := []struct {
		input    string
		expected *bool
	}{
		{"", nil},
		{"true", boolPtr(true)},
		{"TRUE", boolPtr(true)},
		{"1", boolPtr(true)},
		{"yes", boolPtr(true)},
		{"YES", boolPtr(true)},
		{"on", boolPtr(true)},
		{"enable", boolPtr(true)},
		{"enabled", boolPtr(true)},
		{"false", boolPtr(false)},
		{"FALSE", boolPtr(false)},
		{"0", boolPtr(false)},
		{"no", boolPtr(false)},
		{"off", boolPtr(false)},
		{"disable", boolPtr(false)},
		{"disabled", boolPtr(false)},
		{"invalid", nil},
		{"maybe", nil},
	}

	for _, tc := range testCases {
		withEnv(t, "TEST_BOOL", tc.input, func() {
			// We can't directly test the private getEnvBool function,
			// but we can test through the public API
			withEnv(t, "AHASEND_DEBUG", tc.input, func() {
				cfg := ConfigFromEnv()

				if tc.expected == nil {
					// Should use default value (false) when invalid
					assert.False(t, cfg.Debug, "Should use default for invalid value: %s", tc.input)
				} else {
					assert.Equal(t, *tc.expected, cfg.Debug, "Should parse boolean: %s", tc.input)
				}
			})
		})
	}
}

func TestIntegerParsing(t *testing.T) {
	testCases := []struct {
		input    string
		expected *int
	}{
		{"", nil},
		{"0", intPtr(0)},
		{"1", intPtr(1)},
		{"123", intPtr(123)},
		{"-1", intPtr(-1)},
		{"invalid", nil},
		{"1.5", nil},
		{"1a", nil},
	}

	for _, tc := range testCases {
		withEnv(t, "AHASEND_MAX_RETRIES", tc.input, func() {
			cfg := ConfigFromEnv()

			if tc.expected == nil {
				// Should use default value when invalid
				assert.Equal(t, 3, cfg.RetryConfig.MaxRetries, "Should use default for invalid value: %s", tc.input)
			} else if *tc.expected >= 0 {
				assert.Equal(t, *tc.expected, cfg.RetryConfig.MaxRetries, "Should parse integer: %s", tc.input)
			} else {
				// Negative values should use default
				assert.Equal(t, 3, cfg.RetryConfig.MaxRetries, "Should use default for negative value: %s", tc.input)
			}
		})
	}
}

// Helper functions for test cases
func boolPtr(b bool) *bool {
	return &b
}

func intPtr(i int) *int {
	return &i
}
