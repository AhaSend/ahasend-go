// Configuration validation and defaults for the AhaSend Go SDK.
//
// This file provides comprehensive validation, default value management,
// and configuration utilities to ensure robust SDK behavior.

package api

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

// ConfigurationDefaults contains all default values used by the SDK
type ConfigurationDefaults struct {
	// Server Configuration
	DefaultHost   string
	DefaultScheme string

	// Client Configuration
	DefaultUserAgent   string
	DefaultDebug       bool
	DefaultTimeout     time.Duration
	DefaultIdleTimeout time.Duration

	// Rate Limiting Defaults
	DefaultEnableRateLimit bool
	DefaultMaxRetries      int

	// Rate Limit Configurations
	DefaultGeneralRPS       int
	DefaultGeneralBurst     int
	DefaultStatisticsRPS    int
	DefaultStatisticsBurst  int
	DefaultSendMessageRPS   int
	DefaultSendMessageBurst int

	// Idempotency Defaults
	DefaultIdempotencyAutoGenerate bool
	DefaultIdempotencyPrefix       string

	// Connection Limits
	DefaultMaxIdleConns        int
	DefaultMaxIdleConnsPerHost int
	DefaultMaxConnsPerHost     int
}

// GetDefaults returns the default configuration values
func GetDefaults() ConfigurationDefaults {
	return ConfigurationDefaults{
		// Server Configuration
		DefaultHost:   "api.ahasend.com",
		DefaultScheme: "https",

		// Client Configuration
		DefaultUserAgent:   "AhaSend-Go-SDK/1.0",
		DefaultDebug:       false,
		DefaultTimeout:     30 * time.Second,
		DefaultIdleTimeout: 90 * time.Second,

		// Rate Limiting Defaults
		DefaultEnableRateLimit: true,
		DefaultMaxRetries:      3,

		// Rate Limit Configurations (per second)
		DefaultGeneralRPS:       100,
		DefaultGeneralBurst:     200,
		DefaultStatisticsRPS:    1,
		DefaultStatisticsBurst:  1,
		DefaultSendMessageRPS:   100,
		DefaultSendMessageBurst: 200,

		// Idempotency Defaults
		DefaultIdempotencyAutoGenerate: true,
		DefaultIdempotencyPrefix:       "",

		// Connection Limits
		DefaultMaxIdleConns:        100,
		DefaultMaxIdleConnsPerHost: 2,
		DefaultMaxConnsPerHost:     10,
	}
}

// ValidationError represents configuration validation errors
type ConfigurationValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

func (e ConfigurationValidationError) Error() string {
	return fmt.Sprintf("configuration validation error for %s (%v): %s", e.Field, e.Value, e.Message)
}

// ValidationResult contains the results of configuration validation
type ValidationResult struct {
	Valid    bool
	Errors   []ConfigurationValidationError
	Warnings []string
}

// HasErrors returns true if there are validation errors
func (r ValidationResult) HasErrors() bool {
	return len(r.Errors) > 0
}

// HasWarnings returns true if there are validation warnings
func (r ValidationResult) HasWarnings() bool {
	return len(r.Warnings) > 0
}

// Error returns a formatted error message with all validation errors
func (r ValidationResult) Error() string {
	if !r.HasErrors() {
		return ""
	}

	var messages []string
	for _, err := range r.Errors {
		messages = append(messages, err.Error())
	}
	return fmt.Sprintf("configuration validation failed: %s", strings.Join(messages, "; "))
}

// ValidateConfiguration performs comprehensive validation of a Configuration object
func ValidateConfiguration(cfg *Configuration) ValidationResult {
	result := ValidationResult{Valid: true}

	if cfg == nil {
		result.Valid = false
		result.Errors = append(result.Errors, ConfigurationValidationError{
			Field:   "configuration",
			Value:   nil,
			Message: "configuration cannot be nil",
		})
		return result
	}

	// Validate server configuration
	validateServerConfig(cfg, &result)

	// Validate client configuration
	validateClientConfig(cfg, &result)

	// Validate rate limiting configuration
	validateRateLimitConfig(cfg, &result)

	// Validate idempotency configuration
	validateIdempotencyConfig(cfg, &result)

	// Validate HTTP client configuration
	validateHTTPClientConfig(cfg, &result)

	// Set overall validity
	result.Valid = !result.HasErrors()

	return result
}

// validateServerConfig validates server-related configuration
func validateServerConfig(cfg *Configuration, result *ValidationResult) {
	// Validate scheme
	if cfg.Scheme != "" && cfg.Scheme != "http" && cfg.Scheme != "https" {
		result.Errors = append(result.Errors, ConfigurationValidationError{
			Field:   "Scheme",
			Value:   cfg.Scheme,
			Message: "must be 'http' or 'https'",
		})
	}

	// Validate host
	if cfg.Host == "" {
		result.Warnings = append(result.Warnings, "Host is empty, will use default")
	} else {
		// Check for valid hostname format
		if strings.Contains(cfg.Host, "://") {
			result.Errors = append(result.Errors, ConfigurationValidationError{
				Field:   "Host",
				Value:   cfg.Host,
				Message: "should not include scheme (use Scheme field instead)",
			})
		}

		// Basic hostname validation
		if matched, _ := regexp.MatchString(`^[a-zA-Z0-9.-]+$`, cfg.Host); !matched {
			result.Errors = append(result.Errors, ConfigurationValidationError{
				Field:   "Host",
				Value:   cfg.Host,
				Message: "contains invalid characters for hostname",
			})
		}
	}

	// Validate servers configuration
	if len(cfg.Servers) == 0 {
		result.Warnings = append(result.Warnings, "No servers configured, will use defaults")
	} else {
		for i, server := range cfg.Servers {
			if _, err := url.Parse(server.URL); err != nil {
				result.Errors = append(result.Errors, ConfigurationValidationError{
					Field:   fmt.Sprintf("Servers[%d].URL", i),
					Value:   server.URL,
					Message: fmt.Sprintf("invalid URL: %v", err),
				})
			}
		}
	}
}

// validateClientConfig validates client-related configuration
func validateClientConfig(cfg *Configuration, result *ValidationResult) {
	// Validate user agent
	if cfg.UserAgent == "" {
		result.Warnings = append(result.Warnings, "UserAgent is empty, will use default")
	} else if len(cfg.UserAgent) > 200 {
		result.Warnings = append(result.Warnings, "UserAgent is very long (>200 characters)")
	}

	// Validate default headers
	if cfg.DefaultHeader != nil {
		for key, value := range cfg.DefaultHeader {
			if key == "" {
				result.Errors = append(result.Errors, ConfigurationValidationError{
					Field:   "DefaultHeader",
					Value:   key,
					Message: "header key cannot be empty",
				})
			}
			if strings.ContainsAny(key, "\r\n") {
				result.Errors = append(result.Errors, ConfigurationValidationError{
					Field:   "DefaultHeader",
					Value:   key,
					Message: "header key cannot contain newline characters",
				})
			}
			if strings.ContainsAny(value, "\r\n") {
				result.Errors = append(result.Errors, ConfigurationValidationError{
					Field:   "DefaultHeader",
					Value:   fmt.Sprintf("%s: %s", key, value),
					Message: "header value cannot contain newline characters",
				})
			}
		}
	}
}

// validateRateLimitConfig validates rate limiting configuration
func validateRateLimitConfig(cfg *Configuration, result *ValidationResult) {
	// Validate retry configuration
	validateRetryConfig(cfg, result)

	// Validate rate limit configurations
	validateRateLimitConfigStruct := func(config *RateLimitConfig, name string) {
		if config != nil {
			if config.RequestsPerSecond <= 0 {
				result.Errors = append(result.Errors, ConfigurationValidationError{
					Field:   fmt.Sprintf("%s.RequestsPerSecond", name),
					Value:   config.RequestsPerSecond,
					Message: "must be positive",
				})
			}
			if config.BurstCapacity <= 0 {
				result.Errors = append(result.Errors, ConfigurationValidationError{
					Field:   fmt.Sprintf("%s.BurstCapacity", name),
					Value:   config.BurstCapacity,
					Message: "must be positive",
				})
			}
			if config.BurstCapacity < config.RequestsPerSecond {
				result.Warnings = append(result.Warnings,
					fmt.Sprintf("%s.BurstCapacity (%d) is less than RequestsPerSecond (%d)",
						name, config.BurstCapacity, config.RequestsPerSecond))
			}
		}
	}

	validateRateLimitConfigStruct(cfg.DefaultGeneralRateLimit, "DefaultGeneralRateLimit")
	validateRateLimitConfigStruct(cfg.DefaultStatisticsRateLimit, "DefaultStatisticsRateLimit")
	validateRateLimitConfigStruct(cfg.DefaultSendMessageRateLimit, "DefaultSendMessageRateLimit")

	// Validate customer rate limits if present
	if cfg.CustomerRateLimits != nil {
		validateRateLimitConfigStruct(cfg.CustomerRateLimits.General, "CustomerRateLimits.General")
		validateRateLimitConfigStruct(cfg.CustomerRateLimits.Statistics, "CustomerRateLimits.Statistics")
		validateRateLimitConfigStruct(cfg.CustomerRateLimits.SendMessage, "CustomerRateLimits.SendMessage")
	}
}

// validateRetryConfig validates retry configuration
func validateRetryConfig(cfg *Configuration, result *ValidationResult) {
	retryConfig := cfg.RetryConfig

	// Validate MaxRetries
	if retryConfig.MaxRetries < 0 {
		result.Errors = append(result.Errors, ConfigurationValidationError{
			Field:   "RetryConfig.MaxRetries",
			Value:   retryConfig.MaxRetries,
			Message: "cannot be negative",
		})
	} else if retryConfig.MaxRetries > 10 {
		result.Warnings = append(result.Warnings, "RetryConfig.MaxRetries is very high (>10), this may cause long delays")
	}

	// Validate BaseDelay
	if retryConfig.BaseDelay < 0 {
		result.Errors = append(result.Errors, ConfigurationValidationError{
			Field:   "RetryConfig.BaseDelay",
			Value:   retryConfig.BaseDelay,
			Message: "cannot be negative",
		})
	} else if retryConfig.BaseDelay > 60*time.Second {
		result.Warnings = append(result.Warnings, "RetryConfig.BaseDelay is very high (>60s), this may cause long delays")
	}

	// Validate MaxDelay
	if retryConfig.MaxDelay < 0 {
		result.Errors = append(result.Errors, ConfigurationValidationError{
			Field:   "RetryConfig.MaxDelay",
			Value:   retryConfig.MaxDelay,
			Message: "cannot be negative",
		})
	} else if retryConfig.MaxDelay < retryConfig.BaseDelay {
		result.Warnings = append(result.Warnings, "RetryConfig.MaxDelay is less than BaseDelay, MaxDelay will be used as the constant delay")
	}

	// Validate BackoffStrategy
	validStrategies := map[BackoffStrategy]bool{
		BackoffExponential: true,
		BackoffLinear:      true,
		BackoffConstant:    true,
	}
	if !validStrategies[retryConfig.BackoffStrategy] {
		result.Errors = append(result.Errors, ConfigurationValidationError{
			Field:   "RetryConfig.BackoffStrategy",
			Value:   retryConfig.BackoffStrategy,
			Message: "must be 'exponential', 'linear', or 'constant'",
		})
	}

	// Warn about potentially problematic configurations
	if retryConfig.RetryClientErrors {
		result.Warnings = append(result.Warnings, "RetryConfig.RetryClientErrors is enabled - 4xx errors will be retried, which is usually not recommended")
	}

	if retryConfig.RetryValidationErrors {
		result.Warnings = append(result.Warnings, "RetryConfig.RetryValidationErrors is enabled - validation errors will be retried, which is usually not recommended")
	}
}

// validateIdempotencyConfig validates idempotency configuration
func validateIdempotencyConfig(cfg *Configuration, result *ValidationResult) {
	// Validate idempotency key prefix
	if len(cfg.IdempotencyConfig.KeyPrefix) > 50 {
		result.Warnings = append(result.Warnings, "IdempotencyConfig.KeyPrefix is very long (>50 characters)")
	}

	// Check for invalid characters in prefix
	if cfg.IdempotencyConfig.KeyPrefix != "" {
		if matched, _ := regexp.MatchString(`^[a-zA-Z0-9-_]+$`, cfg.IdempotencyConfig.KeyPrefix); !matched {
			result.Errors = append(result.Errors, ConfigurationValidationError{
				Field:   "IdempotencyConfig.KeyPrefix",
				Value:   cfg.IdempotencyConfig.KeyPrefix,
				Message: "can only contain letters, numbers, hyphens, and underscores",
			})
		}
	}
}

// validateHTTPClientConfig validates HTTP client configuration
func validateHTTPClientConfig(cfg *Configuration, result *ValidationResult) {
	if cfg.HTTPClient != nil {
		// Check for reasonable timeout values
		if cfg.HTTPClient.Timeout > 0 {
			if cfg.HTTPClient.Timeout < time.Second {
				result.Warnings = append(result.Warnings, "HTTPClient.Timeout is very short (<1 second)")
			} else if cfg.HTTPClient.Timeout > 5*time.Minute {
				result.Warnings = append(result.Warnings, "HTTPClient.Timeout is very long (>5 minutes)")
			}
		}

		// Validate transport settings if accessible
		if transport, ok := cfg.HTTPClient.Transport.(*http.Transport); ok && transport != nil {
			if transport.MaxIdleConns > 1000 {
				result.Warnings = append(result.Warnings, "HTTPClient Transport MaxIdleConns is very high (>1000)")
			}
			if transport.IdleConnTimeout > 0 && transport.IdleConnTimeout < 30*time.Second {
				result.Warnings = append(result.Warnings, "HTTPClient Transport IdleConnTimeout is very short (<30s)")
			}
		}
	}
}

// ApplyDefaults applies default values to a configuration
func ApplyDefaults(cfg *Configuration) {
	if cfg == nil {
		return
	}

	defaults := GetDefaults()

	// Apply server defaults
	if cfg.Host == "" {
		cfg.Host = defaults.DefaultHost
	}
	if cfg.Scheme == "" {
		cfg.Scheme = defaults.DefaultScheme
	}

	// Apply client defaults
	if cfg.UserAgent == "" {
		cfg.UserAgent = defaults.DefaultUserAgent
	}
	if cfg.DefaultHeader == nil {
		cfg.DefaultHeader = make(map[string]string)
	}

	if cfg.RetryConfig == (RetryConfig{}) {
		// No configuration set - apply defaults
		cfg.RetryConfig = DefaultRetryConfig()
	}

	// Set EnableRateLimit to default if not explicitly set
	cfg.EnableRateLimit = defaults.DefaultEnableRateLimit

	// Apply default rate limit configurations if not set
	if cfg.DefaultGeneralRateLimit == nil {
		cfg.DefaultGeneralRateLimit = &RateLimitConfig{
			RequestsPerSecond: defaults.DefaultGeneralRPS,
			BurstCapacity:     defaults.DefaultGeneralBurst,
			Enabled:           defaults.DefaultEnableRateLimit,
		}
	}

	if cfg.DefaultStatisticsRateLimit == nil {
		cfg.DefaultStatisticsRateLimit = &RateLimitConfig{
			RequestsPerSecond: defaults.DefaultStatisticsRPS,
			BurstCapacity:     defaults.DefaultStatisticsBurst,
			Enabled:           defaults.DefaultEnableRateLimit,
		}
	}

	if cfg.DefaultSendMessageRateLimit == nil {
		cfg.DefaultSendMessageRateLimit = &RateLimitConfig{
			RequestsPerSecond: defaults.DefaultSendMessageRPS,
			BurstCapacity:     defaults.DefaultSendMessageBurst,
			Enabled:           defaults.DefaultEnableRateLimit,
		}
	}

	// Apply idempotency defaults (already done in NewConfiguration, but ensure consistency)
	if cfg.IdempotencyConfig == (IdempotencyConfig{}) {
		cfg.IdempotencyConfig = DefaultIdempotencyConfig()
	}

	// Apply HTTP client defaults if not set
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = createDefaultHTTPClient(defaults)
	}

	// Ensure servers are configured
	if len(cfg.Servers) == 0 {
		cfg.Servers = ServerConfigurations{
			{
				URL:         fmt.Sprintf("%s://%s", cfg.Scheme, cfg.Host),
				Description: "Default server",
			},
		}
	}
}

// createDefaultHTTPClient creates a properly configured default HTTP client
func createDefaultHTTPClient(defaults ConfigurationDefaults) *http.Client {
	transport := &http.Transport{
		MaxIdleConns:        defaults.DefaultMaxIdleConns,
		MaxIdleConnsPerHost: defaults.DefaultMaxIdleConnsPerHost,
		MaxConnsPerHost:     defaults.DefaultMaxConnsPerHost,
		IdleConnTimeout:     defaults.DefaultIdleTimeout,
	}

	return &http.Client{
		Transport: transport,
		Timeout:   defaults.DefaultTimeout,
	}
}

// ValidateAndApplyDefaults validates configuration and applies defaults
func ValidateAndApplyDefaults(cfg *Configuration) ValidationResult {
	// Apply defaults first
	ApplyDefaults(cfg)

	// Then validate
	return ValidateConfiguration(cfg)
}

// NewValidatedConfiguration creates a new configuration with defaults applied and validation performed
func NewValidatedConfiguration() (*Configuration, ValidationResult) {
	cfg := NewConfiguration()
	result := ValidateAndApplyDefaults(cfg)
	return cfg, result
}

// NewValidatedConfigurationFromEnv creates a configuration from environment variables with validation
func NewValidatedConfigurationFromEnv() (*Configuration, ValidationResult) {
	cfg := ConfigFromEnv()
	result := ValidateAndApplyDefaults(cfg)
	return cfg, result
}

// ConfigurationSummary provides a summary of the current configuration
type ConfigurationSummary struct {
	ServerURL          string
	UserAgent          string
	Debug              bool
	MaxRetries         int
	RateLimitEnabled   bool
	IdempotencyEnabled bool
	TimeoutSeconds     int
}

// GetConfigurationSummary returns a summary of the configuration
func GetConfigurationSummary(cfg *Configuration) ConfigurationSummary {
	summary := ConfigurationSummary{
		ServerURL:          fmt.Sprintf("%s://%s", cfg.Scheme, cfg.Host),
		UserAgent:          cfg.UserAgent,
		Debug:              cfg.Debug,
		MaxRetries:         cfg.RetryConfig.MaxRetries,
		RateLimitEnabled:   cfg.EnableRateLimit,
		IdempotencyEnabled: cfg.IdempotencyConfig.AutoGenerate,
	}

	if cfg.HTTPClient != nil && cfg.HTTPClient.Timeout > 0 {
		summary.TimeoutSeconds = int(cfg.HTTPClient.Timeout.Seconds())
	}

	return summary
}

// IsProductionReady checks if a configuration is suitable for production use
func IsProductionReady(cfg *Configuration) (bool, []string) {
	var issues []string

	// Check debug mode
	if cfg.Debug {
		issues = append(issues, "Debug mode is enabled - should be disabled in production")
	}

	// Check rate limiting
	if !cfg.EnableRateLimit {
		issues = append(issues, "Rate limiting is disabled - should be enabled in production")
	}

	// Check retry configuration
	if !cfg.RetryConfig.IsRetryEnabled() {
		issues = append(issues, "Retries are disabled - should be enabled for production resilience")
	} else if cfg.RetryConfig.MaxRetries > 5 {
		issues = append(issues, "RetryConfig.MaxRetries is very high - may cause excessive delays in production")
	}

	// Check timeout configuration
	if cfg.HTTPClient != nil && cfg.HTTPClient.Timeout == 0 {
		issues = append(issues, "No HTTP timeout configured - should set reasonable timeout for production")
	}

	// Check user agent
	if cfg.UserAgent == "AhaSend-Go-SDK" || cfg.UserAgent == "" {
		issues = append(issues, "Using default UserAgent - consider setting a more specific one for production")
	}

	// Check scheme
	if cfg.Scheme != "https" {
		issues = append(issues, "Using non-HTTPS scheme - should use HTTPS in production")
	}

	return len(issues) == 0, issues
}

// OptimizeForProduction applies production-ready settings to a configuration
func OptimizeForProduction(cfg *Configuration) {
	cfg.Debug = false
	cfg.EnableRateLimit = true

	// Ensure RetryConfig has production-ready settings
	if !cfg.RetryConfig.IsRetryEnabled() {
		cfg.RetryConfig.Enabled = true
		cfg.RetryConfig.MaxRetries = 3
	}
	cfg.RetryConfig.RetryClientErrors = false     // Never retry 4xx in production
	cfg.RetryConfig.RetryValidationErrors = false // Never retry validation errors

	// Force HTTPS in production
	cfg.Scheme = "https"

	// Set reasonable timeout if not configured
	if cfg.HTTPClient != nil && cfg.HTTPClient.Timeout == 0 {
		cfg.HTTPClient.Timeout = 30 * time.Second
	}

	// Ensure idempotency is enabled
	cfg.IdempotencyConfig.AutoGenerate = true
}
