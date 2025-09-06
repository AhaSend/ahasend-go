// Configuration and context management for the AhaSend Go SDK.
//
// This file contains configuration structures and context helpers for API client setup,
// server configuration, and authentication management.

package api

import (
	"context"
	"fmt"
	"math/rand/v2"
	"net/http"
	"strings"
	"time"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextServerIndex uses a server configuration from the index.
	ContextServerIndex = contextKey("serverIndex")

	// ContextOperationServerIndices uses a server configuration from the index mapping.
	ContextOperationServerIndices = contextKey("serverOperationIndices")

	// ContextServerVariables overrides a server configuration variables.
	ContextServerVariables = contextKey("serverVariables")

	// ContextOperationServerVariables overrides a server configuration variables using operation specific values.
	ContextOperationServerVariables = contextKey("serverOperationVariables")

	// ContextSkipRateLimit skips rate limiting for this request (used internally by Execute method)
	ContextSkipRateLimit = contextKey("skipRateLimit")

	// ContextIdempotencyKey sets an idempotency key for the request
	ContextIdempotencyKey = contextKey("idempotencyKey")
)

// ServerVariable stores the information about a server variable
type ServerVariable struct {
	Description  string
	DefaultValue string
	EnumValues   []string
}

// ServerConfiguration stores the information about a server
type ServerConfiguration struct {
	URL         string
	Description string
	Variables   map[string]ServerVariable
}

// ServerConfigurations stores multiple ServerConfiguration items
type ServerConfigurations []ServerConfiguration

// BackoffStrategy defines the retry backoff strategy
type BackoffStrategy string

const (
	// BackoffExponential uses exponential backoff with jitter
	BackoffExponential BackoffStrategy = "exponential"
	// BackoffLinear uses linear backoff
	BackoffLinear BackoffStrategy = "linear"
	// BackoffConstant uses constant delay between retries
	BackoffConstant BackoffStrategy = "constant"
)

// RetryConfig provides comprehensive retry configuration options
type RetryConfig struct {
	// Enabled controls whether retries are enabled at all
	Enabled bool `json:"enabled"`
	// MaxRetries is the maximum number of retry attempts (0 means no retries)
	MaxRetries int `json:"max_retries"`
	// RetryClientErrors controls whether 4xx client errors are retried (default: false)
	RetryClientErrors bool `json:"retry_client_errors"`
	// RetryValidationErrors controls whether client-side validation errors are retried (default: false)
	RetryValidationErrors bool `json:"retry_validation_errors"`
	// BackoffStrategy determines the delay strategy between retries
	BackoffStrategy BackoffStrategy `json:"backoff_strategy"`
	// BaseDelay is the initial delay for the first retry (used by all strategies)
	BaseDelay time.Duration `json:"base_delay"`
	// MaxDelay is the maximum delay between retries (prevents exponential backoff from growing too large)
	MaxDelay time.Duration `json:"max_delay"`
}

// DefaultRetryConfig returns sensible default retry configuration
func DefaultRetryConfig() RetryConfig {
	return RetryConfig{
		Enabled:               true,
		MaxRetries:            3,
		RetryClientErrors:     false, // Never retry 4xx errors by default
		RetryValidationErrors: false, // Never retry validation errors by default
		BackoffStrategy:       BackoffExponential,
		BaseDelay:             1 * time.Second,
		MaxDelay:              30 * time.Second,
	}
}

// IsRetryEnabled returns true if retries are enabled and MaxRetries > 0
func (rc RetryConfig) IsRetryEnabled() bool {
	return rc.Enabled && rc.MaxRetries > 0
}

// GetDelay calculates the delay for a specific retry attempt
func (rc RetryConfig) GetDelay(attempt int) time.Duration {
	if attempt <= 0 {
		return rc.BaseDelay
	}

	var delay time.Duration
	switch rc.BackoffStrategy {
	case BackoffExponential:
		delay = time.Duration(1<<uint(attempt-1)) * rc.BaseDelay
		// Add jitter (±25% randomization)
		jitter := time.Duration(rand.Int64N(int64(delay / 2)))
		delay = delay + jitter - delay/4
	case BackoffLinear:
		delay = time.Duration(attempt) * rc.BaseDelay
	case BackoffConstant:
		delay = rc.BaseDelay
	default:
		delay = rc.BaseDelay
	}

	// Cap at MaxDelay
	if delay > rc.MaxDelay {
		delay = rc.MaxDelay
	}

	return delay
}

// RequestMonitor provides hooks for monitoring HTTP requests
type RequestMonitor interface {
	// OnRequestStart is called before a request is sent
	OnRequestStart(ctx context.Context, method, url string, headers map[string]string)
	// OnRequestComplete is called after a request completes (success or failure)
	OnRequestComplete(ctx context.Context, method, url string, statusCode int, duration time.Duration, err error)
	// OnRetry is called when a request is being retried
	OnRetry(ctx context.Context, method, url string, attempt int, err error)
	// OnRateLimitWait is called when a request is delayed due to rate limiting
	OnRateLimitWait(ctx context.Context, endpointType string, waitDuration time.Duration)
}

// ClientOption is a function that modifies a Configuration
type ClientOption func(*Configuration)

// Configuration stores the configuration of the API client
type Configuration struct {
	Host             string            `json:"host,omitempty"`
	Scheme           string            `json:"scheme,omitempty"`
	DefaultHeader    map[string]string `json:"defaultHeader,omitempty"`
	UserAgent        string            `json:"userAgent,omitempty"`
	Debug            bool              `json:"debug,omitempty"`
	Servers          ServerConfigurations
	OperationServers map[string]ServerConfigurations
	HTTPClient       *http.Client

	// Authentication configuration
	APIKey string `json:"apiKey,omitempty"`

	// Rate limiting configuration
	EnableRateLimit    bool                     `json:"enableRateLimit,omitempty"`
	CustomerRateLimits *CustomerRateLimitConfig `json:"customerRateLimits,omitempty"`

	// Retry configuration
	RetryConfig RetryConfig `json:"retryConfig"`

	// Default rate limits (can be overridden per customer)
	DefaultGeneralRateLimit     *RateLimitConfig `json:"defaultGeneralRateLimit,omitempty"`
	DefaultStatisticsRateLimit  *RateLimitConfig `json:"defaultStatisticsRateLimit,omitempty"`
	DefaultSendMessageRateLimit *RateLimitConfig `json:"defaultSendMessageRateLimit,omitempty"`

	// Idempotency configuration
	IdempotencyConfig IdempotencyConfig `json:"idempotencyConfig,omitempty"`

	// Monitoring configuration
	RequestMonitor RequestMonitor `json:"-"` // Not serialized - runtime configuration only
}

// NewConfiguration returns a new Configuration object with default settings
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		DefaultHeader:    make(map[string]string),
		OperationServers: map[string]ServerConfigurations{},
	}

	// Apply all defaults through the validation system
	ApplyDefaults(cfg)

	return cfg
}

// NewConfigurationFromEnv returns a new Configuration object with values loaded from environment variables.
// This is a convenience function equivalent to ConfigFromEnv().
func NewConfigurationFromEnv() *Configuration {
	return ConfigFromEnv()
}

// NewAPIClientFromEnv creates a new API client with configuration loaded from environment variables.
// This is the easiest way to get started - just set AHASEND_API_KEY and go.
func NewAPIClientFromEnv() *APIClient {
	cfg := ConfigFromEnv()
	// Convert configuration to functional options
	opts := []ClientOption{}

	if cfg.APIKey != "" {
		opts = append(opts, func(c *Configuration) { c.APIKey = cfg.APIKey })
	}
	if cfg.Debug {
		opts = append(opts, WithDebug(cfg.Debug))
	}
	if cfg.UserAgent != "" {
		opts = append(opts, WithUserAgent(cfg.UserAgent))
	}
	if cfg.HTTPClient != nil {
		opts = append(opts, WithHTTPClient(cfg.HTTPClient))
	}
	opts = append(opts, WithRateLimit(cfg.EnableRateLimit))
	opts = append(opts, WithRetryConfig(cfg.RetryConfig))
	if cfg.CustomerRateLimits != nil {
		opts = append(opts, WithCustomerRateLimits(*cfg.CustomerRateLimits))
	}
	opts = append(opts, WithIdempotencyConfig(cfg.IdempotencyConfig))

	// Add default headers
	for key, value := range cfg.DefaultHeader {
		opts = append(opts, WithDefaultHeader(key, value))
	}

	return NewAPIClient(opts...)
}

// NewValidatedAPIClient creates a new API client with a validated configuration.
// Returns the client and any validation results.
// Deprecated: Use NewAPIClient with functional options instead.
func NewValidatedAPIClient(cfg *Configuration) (*APIClient, ValidationResult) {
	result := ValidateAndApplyDefaults(cfg)
	return NewAPIClientWithConfig(cfg), result
}

// NewAPIClientWithConfig creates a new API client with a pre-configured Configuration.
// This is provided for compatibility with existing validation functions.
func NewAPIClientWithConfig(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// Initialize rate limiter
	c.rateLimiter = NewRateLimiter()
	c.rateLimiter.SetGlobalEnabled(cfg.EnableRateLimit)

	// Apply custom rate limits if provided
	if cfg.CustomerRateLimits != nil {
		c.rateLimiter.ConfigureFromCustomerConfig(*cfg.CustomerRateLimits)
	}

	// Initialize idempotency helper
	c.idempotencyHelper = NewIdempotencyHelper(cfg.IdempotencyConfig)

	// API Services
	c.APIKeysAPI = (*APIKeysAPIService)(&c.common)
	c.AccountsAPI = (*AccountsAPIService)(&c.common)
	c.DomainsAPI = (*DomainsAPIService)(&c.common)
	c.MessagesAPI = (*MessagesAPIService)(&c.common)
	c.RoutesAPI = (*RoutesAPIService)(&c.common)
	c.SMTPCredentialsAPI = (*SMTPCredentialsAPIService)(&c.common)
	c.StatisticsAPI = (*StatisticsAPIService)(&c.common)
	c.SuppressionsAPI = (*SuppressionsAPIService)(&c.common)
	c.UtilityAPI = (*UtilityAPIService)(&c.common)
	c.WebhooksAPI = (*WebhooksAPIService)(&c.common)

	return c
}

// NewValidatedAPIClientFromEnv creates a new API client from environment variables with validation.
// Returns the client and any validation results.
func NewValidatedAPIClientFromEnv() (*APIClient, ValidationResult) {
	cfg, result := NewValidatedConfigurationFromEnv()
	return NewAPIClientWithConfig(cfg), result
}

// AddDefaultHeader adds a new HTTP header to the default header in the request
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

// URL formats template on a index using given variables
func (sc ServerConfigurations) URL(index int, variables map[string]string) (string, error) {
	if index < 0 || len(sc) <= index {
		return "", fmt.Errorf("index %v out of range %v", index, len(sc)-1)
	}
	server := sc[index]
	url := server.URL

	// go through variables and replace placeholders
	for name, variable := range server.Variables {
		if value, ok := variables[name]; ok {
			found := bool(len(variable.EnumValues) == 0)
			for _, enumValue := range variable.EnumValues {
				if value == enumValue {
					found = true
				}
			}
			if !found {
				return "", fmt.Errorf("the variable %s in the server URL has invalid value %v. Must be %v", name, value, variable.EnumValues)
			}
			url = strings.ReplaceAll(url, "{"+name+"}", value)
		} else {
			url = strings.ReplaceAll(url, "{"+name+"}", variable.DefaultValue)
		}
	}
	return url, nil
}

// ServerURL returns URL based on server settings
func (c *Configuration) ServerURL(index int, variables map[string]string) (string, error) {
	return c.Servers.URL(index, variables)
}

func getServerIndex(ctx context.Context) (int, error) {
	si := ctx.Value(ContextServerIndex)
	if si != nil {
		if index, ok := si.(int); ok {
			return index, nil
		}
		return 0, reportError("Invalid type %T should be int", si)
	}
	return 0, nil
}

func getServerOperationIndex(ctx context.Context, endpoint string) (int, error) {
	osi := ctx.Value(ContextOperationServerIndices)
	if osi != nil {
		if operationIndices, ok := osi.(map[string]int); !ok {
			return 0, reportError("Invalid type %T should be map[string]int", osi)
		} else {
			index, ok := operationIndices[endpoint]
			if ok {
				return index, nil
			}
		}
	}
	return getServerIndex(ctx)
}

func getServerVariables(ctx context.Context) (map[string]string, error) {
	sv := ctx.Value(ContextServerVariables)
	if sv != nil {
		if variables, ok := sv.(map[string]string); ok {
			return variables, nil
		}
		return nil, reportError("ctx value of ContextServerVariables has invalid type %T should be map[string]string", sv)
	}
	return nil, nil
}

func getServerOperationVariables(ctx context.Context, endpoint string) (map[string]string, error) {
	osv := ctx.Value(ContextOperationServerVariables)
	if osv != nil {
		if operationVariables, ok := osv.(map[string]map[string]string); !ok {
			return nil, reportError("ctx value of ContextOperationServerVariables has invalid type %T should be map[string]map[string]string", osv)
		} else {
			variables, ok := operationVariables[endpoint]
			if ok {
				return variables, nil
			}
		}
	}
	return getServerVariables(ctx)
}

// ServerURLWithContext returns a new server URL given an endpoint
func (c *Configuration) ServerURLWithContext(ctx context.Context, endpoint string) (string, error) {
	sc, ok := c.OperationServers[endpoint]
	if !ok {
		sc = c.Servers
	}

	if ctx == nil {
		return sc.URL(0, nil)
	}

	index, err := getServerOperationIndex(ctx, endpoint)
	if err != nil {
		return "", err
	}

	variables, err := getServerOperationVariables(ctx, endpoint)
	if err != nil {
		return "", err
	}

	return sc.URL(index, variables)
}

// Client Option Functions

// WithAPIKey sets the API key for authentication
func WithAPIKey(key string) ClientOption {
	return func(cfg *Configuration) {
		cfg.APIKey = key
	}
}

// WithRateLimit enables or disables rate limiting globally
func WithRateLimit(enabled bool) ClientOption {
	return func(cfg *Configuration) {
		cfg.EnableRateLimit = enabled
	}
}

// WithRetryConfig sets the retry configuration
func WithRetryConfig(retryConfig RetryConfig) ClientOption {
	return func(cfg *Configuration) {
		cfg.RetryConfig = retryConfig
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(client *http.Client) ClientOption {
	return func(cfg *Configuration) {
		cfg.HTTPClient = client
	}
}

// WithRequestMonitor sets a request monitor for observability
func WithRequestMonitor(monitor RequestMonitor) ClientOption {
	return func(cfg *Configuration) {
		cfg.RequestMonitor = monitor
	}
}

// WithDebug enables or disables debug mode
func WithDebug(debug bool) ClientOption {
	return func(cfg *Configuration) {
		cfg.Debug = debug
	}
}

// WithUserAgent sets a custom user agent string
func WithUserAgent(userAgent string) ClientOption {
	return func(cfg *Configuration) {
		cfg.UserAgent = userAgent
	}
}

// WithDefaultHeader adds a default header to all requests
func WithDefaultHeader(key, value string) ClientOption {
	return func(cfg *Configuration) {
		if cfg.DefaultHeader == nil {
			cfg.DefaultHeader = make(map[string]string)
		}
		cfg.DefaultHeader[key] = value
	}
}

// WithCustomerRateLimits sets customer-specific rate limits
func WithCustomerRateLimits(rateLimits CustomerRateLimitConfig) ClientOption {
	return func(cfg *Configuration) {
		cfg.CustomerRateLimits = &rateLimits
	}
}

// WithIdempotencyConfig sets the idempotency configuration
func WithIdempotencyConfig(config IdempotencyConfig) ClientOption {
	return func(cfg *Configuration) {
		cfg.IdempotencyConfig = config
	}
}

func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}
