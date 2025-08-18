// Environment variable support for the AhaSend Go SDK.
//
// This file provides utilities for loading configuration from environment variables,
// making it easier to configure the SDK in different deployment environments.

package api

import (
	"context"
	"os"
	"strconv"
	"strings"
	"time"
)

// Environment variable names
const (
	// API Authentication
	EnvAPIKey   = "AHASEND_API_KEY"
	EnvAPIToken = "AHASEND_TOKEN" // Alternative to API_KEY

	// Server Configuration
	EnvBaseURL = "AHASEND_BASE_URL"
	EnvHost    = "AHASEND_HOST"   // Host without scheme
	EnvScheme  = "AHASEND_SCHEME" // http or https

	// Debug and Logging
	EnvDebug     = "AHASEND_DEBUG"
	EnvUserAgent = "AHASEND_USER_AGENT"

	// Rate Limiting
	EnvEnableRateLimit = "AHASEND_ENABLE_RATE_LIMIT"
	EnvMaxRetries      = "AHASEND_MAX_RETRIES"

	// Timeouts (in seconds)
	EnvTimeout        = "AHASEND_TIMEOUT"
	EnvConnectTimeout = "AHASEND_CONNECT_TIMEOUT"

	// Idempotency
	EnvIdempotencyAutoGenerate = "AHASEND_IDEMPOTENCY_AUTO_GENERATE"
	EnvIdempotencyPrefix       = "AHASEND_IDEMPOTENCY_PREFIX"
)

// ConfigFromEnv creates a new Configuration with values loaded from environment variables.
// Environment variables override default values but can still be overridden programmatically.
func ConfigFromEnv() *Configuration {
	cfg := NewConfiguration()

	// Load environment variables into configuration
	loadEnvIntoConfig(cfg)

	return cfg
}

// LoadEnvIntoConfig loads environment variables into an existing configuration.
// This allows you to start with a custom config and overlay environment variables.
func LoadEnvIntoConfig(cfg *Configuration) {
	if cfg == nil {
		return
	}
	loadEnvIntoConfig(cfg)
}

// loadEnvIntoConfig is the internal implementation
func loadEnvIntoConfig(cfg *Configuration) {
	// Server Configuration
	if baseURL := getEnv(EnvBaseURL); baseURL != "" {
		// Parse base URL to extract host and scheme
		if strings.HasPrefix(baseURL, "http://") {
			cfg.Scheme = "http"
			cfg.Host = strings.TrimPrefix(baseURL, "http://")
		} else if strings.HasPrefix(baseURL, "https://") {
			cfg.Scheme = "https"
			cfg.Host = strings.TrimPrefix(baseURL, "https://")
		} else {
			cfg.Host = baseURL
		}
	}

	// Individual host/scheme override base URL
	if host := getEnv(EnvHost); host != "" {
		cfg.Host = host
	}
	if scheme := getEnv(EnvScheme); scheme != "" {
		cfg.Scheme = scheme
	}

	// Debug Configuration
	if debug := getEnvBool(EnvDebug); debug != nil {
		cfg.Debug = *debug
	}

	if userAgent := getEnv(EnvUserAgent); userAgent != "" {
		cfg.UserAgent = userAgent
	}

	// Rate Limiting Configuration
	if enableRL := getEnvBool(EnvEnableRateLimit); enableRL != nil {
		cfg.EnableRateLimit = *enableRL
	}

	if maxRetries := getEnvInt(EnvMaxRetries); maxRetries != nil && *maxRetries >= 0 {
		cfg.RetryConfig.MaxRetries = *maxRetries
	}

	// Idempotency Configuration
	if autoGen := getEnvBool(EnvIdempotencyAutoGenerate); autoGen != nil {
		cfg.IdempotencyConfig.AutoGenerate = *autoGen
	}

	if prefix := getEnv(EnvIdempotencyPrefix); prefix != "" {
		cfg.IdempotencyConfig.KeyPrefix = prefix
	}
}

// GetAPIKeyFromEnv returns the API key from environment variables.
// It checks AHASEND_API_KEY first, then falls back to AHASEND_TOKEN.
func GetAPIKeyFromEnv() string {
	if key := getEnv(EnvAPIKey); key != "" {
		return key
	}
	return getEnv(EnvAPIToken)
}

// ContextWithEnvAuth creates a context with API key from environment variables.
// This is a convenience function for quick setup.
func ContextWithEnvAuth(ctx context.Context) context.Context {
	apiKey := GetAPIKeyFromEnv()
	if apiKey == "" {
		return ctx
	}
	return context.WithValue(ctx, ContextAccessToken, apiKey)
}

// GetTimeoutFromEnv returns timeout configuration from environment variables
func GetTimeoutFromEnv() (timeout time.Duration, connectTimeout time.Duration) {
	if timeoutSec := getEnvInt(EnvTimeout); timeoutSec != nil && *timeoutSec > 0 {
		timeout = time.Duration(*timeoutSec) * time.Second
	}

	if connectSec := getEnvInt(EnvConnectTimeout); connectSec != nil && *connectSec > 0 {
		connectTimeout = time.Duration(*connectSec) * time.Second
	}

	return
}

// Environment variable helper functions

// getEnv gets an environment variable, returning empty string if not set
func getEnv(key string) string {
	return strings.TrimSpace(os.Getenv(key))
}

// getEnvBool parses a boolean environment variable
// Returns nil if not set or invalid, otherwise returns pointer to bool
func getEnvBool(key string) *bool {
	val := strings.TrimSpace(strings.ToLower(os.Getenv(key)))
	if val == "" {
		return nil
	}

	switch val {
	case "true", "1", "yes", "on", "enable", "enabled":
		result := true
		return &result
	case "false", "0", "no", "off", "disable", "disabled":
		result := false
		return &result
	default:
		return nil
	}
}

// getEnvInt parses an integer environment variable
// Returns nil if not set or invalid, otherwise returns pointer to int
func getEnvInt(key string) *int {
	val := strings.TrimSpace(os.Getenv(key))
	if val == "" {
		return nil
	}

	if num, err := strconv.Atoi(val); err == nil {
		return &num
	}
	return nil
}

// getEnvDuration parses a duration environment variable (supports suffixes like "30s", "5m")
// Returns zero duration if not set or invalid
func getEnvDuration(key string) time.Duration {
	val := strings.TrimSpace(os.Getenv(key))
	if val == "" {
		return 0
	}

	// Try parsing as duration first
	if duration, err := time.ParseDuration(val); err == nil {
		return duration
	}

	// Fall back to seconds if it's just a number
	if seconds, err := strconv.Atoi(val); err == nil && seconds > 0 {
		return time.Duration(seconds) * time.Second
	}

	return 0
}

// EnvConfig represents environment variable configuration for documentation
type EnvConfig struct {
	// API Authentication
	APIKey   string `env:"AHASEND_API_KEY" description:"AhaSend API key for authentication"`
	APIToken string `env:"AHASEND_TOKEN" description:"Alternative to AHASEND_API_KEY"`

	// Server Configuration
	BaseURL string `env:"AHASEND_BASE_URL" description:"Full base URL (e.g. https://api.ahasend.com)"`
	Host    string `env:"AHASEND_HOST" description:"API host without scheme (e.g. api.ahasend.com)"`
	Scheme  string `env:"AHASEND_SCHEME" description:"URL scheme: http or https"`

	// Debug and Logging
	Debug     bool   `env:"AHASEND_DEBUG" description:"Enable debug logging"`
	UserAgent string `env:"AHASEND_USER_AGENT" description:"Custom user agent string"`

	// Rate Limiting
	EnableRateLimit bool `env:"AHASEND_ENABLE_RATE_LIMIT" description:"Enable rate limiting"`
	MaxRetries      int  `env:"AHASEND_MAX_RETRIES" description:"Maximum number of retries"`

	// Timeouts
	Timeout        int `env:"AHASEND_TIMEOUT" description:"Request timeout in seconds"`
	ConnectTimeout int `env:"AHASEND_CONNECT_TIMEOUT" description:"Connection timeout in seconds"`

	// Idempotency
	IdempotencyAutoGenerate bool   `env:"AHASEND_IDEMPOTENCY_AUTO_GENERATE" description:"Auto-generate idempotency keys"`
	IdempotencyPrefix       string `env:"AHASEND_IDEMPOTENCY_PREFIX" description:"Prefix for generated idempotency keys"`
}

// GetEnvDocumentation returns a list of all supported environment variables with descriptions
func GetEnvDocumentation() map[string]string {
	return map[string]string{
		EnvAPIKey:   "AhaSend API key for authentication (aha-sk-...)",
		EnvAPIToken: "Alternative to AHASEND_API_KEY",

		EnvBaseURL: "Full base URL including scheme (e.g. https://api.ahasend.com)",
		EnvHost:    "API host without scheme (e.g. api.ahasend.com)",
		EnvScheme:  "URL scheme: 'http' or 'https'",

		EnvDebug:     "Enable debug logging: true/false, 1/0, yes/no",
		EnvUserAgent: "Custom user agent string for requests",

		EnvEnableRateLimit: "Enable rate limiting: true/false, 1/0, yes/no",
		EnvMaxRetries:      "Maximum number of retries for failed requests (default: 3)",

		EnvTimeout:        "Request timeout in seconds (default: 30)",
		EnvConnectTimeout: "Connection timeout in seconds (default: 10)",

		EnvIdempotencyAutoGenerate: "Auto-generate idempotency keys: true/false (default: true)",
		EnvIdempotencyPrefix:       "Prefix for generated idempotency keys",
	}
}

// ValidateEnvConfig validates environment variables and returns any issues
func ValidateEnvConfig() []string {
	var issues []string

	// Check for API key
	if GetAPIKeyFromEnv() == "" {
		issues = append(issues, "No API key found in AHASEND_API_KEY or AHASEND_TOKEN")
	}

	// Validate scheme if provided
	if scheme := getEnv(EnvScheme); scheme != "" && scheme != "http" && scheme != "https" {
		issues = append(issues, "AHASEND_SCHEME must be 'http' or 'https'")
	}

	// Validate numeric values
	if maxRetries := getEnv(EnvMaxRetries); maxRetries != "" {
		if getEnvInt(EnvMaxRetries) == nil {
			issues = append(issues, "AHASEND_MAX_RETRIES must be a valid integer")
		}
	}

	if timeout := getEnv(EnvTimeout); timeout != "" {
		if getEnvInt(EnvTimeout) == nil && getEnvDuration(EnvTimeout) == 0 {
			issues = append(issues, "AHASEND_TIMEOUT must be a valid number or duration")
		}
	}

	// Validate boolean values
	boolVars := []string{EnvDebug, EnvEnableRateLimit, EnvIdempotencyAutoGenerate}
	for _, key := range boolVars {
		if val := getEnv(key); val != "" && getEnvBool(key) == nil {
			issues = append(issues, key+" must be a valid boolean (true/false, 1/0, yes/no)")
		}
	}

	return issues
}
