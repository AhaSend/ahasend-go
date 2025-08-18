// Package ahasend provides a Go SDK for the AhaSend transactional email API.
//
// AhaSend enables developers to send transactional emails with high deliverability,
// comprehensive tracking, and powerful management features.
//
// Key Features:
//   - Send transactional emails with attachments and scheduling
//   - Comprehensive domain, webhook, and route management
//   - Built-in rate limiting and automatic retry with exponential backoff
//   - Statistics and analytics for email performance
//   - Support for suppression lists and SMTP credentials
//
// Authentication:
// All API requests require a Bearer token in the Authorization header:
//   Authorization: Bearer aha-sk-64-CHARACTER-RANDOM-STRING
//
// Basic Usage:
//   import "github.com/AhaSend/ahasend-go"
//
//   cfg := ahasend.NewConfiguration()
//   client := ahasend.NewAPIClient(cfg)
//
//   ctx := context.WithValue(context.Background(),
//     ahasend.ContextAccessToken, "your-api-key")
//
//   // Send a message
//   messageReq := &ahasend.CreateMessageRequest{
//     From:    "sender@yourdomain.com",
//     To:      []string{"recipient@example.com"},
//     Subject: "Test Email",
//     Html:    "<p>Hello from AhaSend!</p>",
//   }
//
//   response, _, err := client.MessagesAPI.
//     CreateMessage(ctx, accountID).
//     CreateMessageRequest(*messageReq).
//     Execute()
//
// For more information visit: https://ahasend.com/docs

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/AhaSend/ahasend-go/models/common"
)

var (
	// pathParamsRegex is compiled once for performance in validatePathParams
	pathParamsRegex = regexp.MustCompile(`\{([^}]+)\}`)
)

// RequestConfig contains all information needed to execute an API request
type RequestConfig struct {
	// HTTP method (GET, POST, PUT, DELETE, PATCH)
	Method string

	// URL path template with placeholders like "/v2/accounts/{account_id}/messages/{message_id}"
	PathTemplate string

	// Path parameters to replace in template (e.g., {"account_id": "123", "message_id": "456"})
	PathParams map[string]string

	// Query parameters to append to URL
	QueryParams url.Values

	// Additional headers to include (Auth and standard headers added automatically)
	Headers map[string]string

	// Request body (will be JSON encoded if not nil)
	Body interface{}

	// Pointer to result struct where response will be decoded
	Result interface{}

	// Optional: Skip rate limiting for this request
	SkipRateLimit bool

	// Optional: Custom timeout for this specific request
	CustomTimeout *time.Duration

	// Optional: Custom retry configuration for this request
	CustomRetry *RetryConfig

	// Internal: Endpoint type for rate limiting classification
	endpointType EndpointType
}

// RequestOption allows modifying RequestConfig using functional options pattern
type RequestOption func(*RequestConfig)

// WithTimeout sets a custom timeout for this specific request
func WithTimeout(timeout time.Duration) RequestOption {
	return func(rc *RequestConfig) {
		rc.CustomTimeout = &timeout
	}
}

// WithRetry sets custom retry configuration for this request
func WithRetry(retry RetryConfig) RequestOption {
	return func(rc *RequestConfig) {
		rc.CustomRetry = &retry
	}
}

// WithHeaders adds additional headers to the request
func WithHeaders(headers map[string]string) RequestOption {
	return func(rc *RequestConfig) {
		for k, v := range headers {
			if rc.Headers == nil {
				rc.Headers = make(map[string]string)
			}
			rc.Headers[k] = v
		}
	}
}

// WithIdempotencyKey sets an idempotency key for the request
func WithIdempotencyKey(key string) RequestOption {
	return func(rc *RequestConfig) {
		if rc.Headers == nil {
			rc.Headers = make(map[string]string)
		}
		rc.Headers["Idempotency-Key"] = key
	}
}

// WithRequestAPIKey sets an API key for this specific request (overrides client-level auth)
func WithRequestAPIKey(key string) RequestOption {
	return func(rc *RequestConfig) {
		if rc.Headers == nil {
			rc.Headers = make(map[string]string)
		}
		rc.Headers["Authorization"] = "Bearer " + key
	}
}

// WithoutRateLimit skips rate limiting for this request
func WithoutRateLimit() RequestOption {
	return func(rc *RequestConfig) {
		rc.SkipRateLimit = true
	}
}

// Security and Validation Functions

// validatePathParam checks for path traversal and other security issues
func validatePathParam(value string) error {
	if strings.Contains(value, "..") || strings.Contains(value, "//") {
		return fmt.Errorf("potential path traversal detected: %s", value)
	}
	if strings.ContainsAny(value, "\n\r\t") {
		return fmt.Errorf("invalid characters in parameter: %s", value)
	}
	return nil
}

// buildPath replaces path parameters in a template string with security validation
// Example: buildPath("/v2/accounts/{account_id}/messages", {"account_id": "123"})
// Returns: "/v2/accounts/123/messages"
func buildPath(template string, params map[string]string) (string, error) {
	path := template
	for key, value := range params {
		// Security: Validate against path traversal
		if err := validatePathParam(value); err != nil {
			return "", fmt.Errorf("invalid path parameter %s: %w", key, err)
		}

		placeholder := "{" + key + "}"
		// URL escape the value to handle special characters
		escapedValue := url.PathEscape(value)
		path = strings.ReplaceAll(path, placeholder, escapedValue)
	}
	return path, nil
}

// validatePathParams ensures all required parameters are provided
func validatePathParams(template string, params map[string]string) error {
	// Find all placeholders in template using compiled regex
	matches := pathParamsRegex.FindAllStringSubmatch(template, -1)

	for _, match := range matches {
		if len(match) > 1 {
			paramName := match[1]
			if _, exists := params[paramName]; !exists {
				return fmt.Errorf("missing required path parameter: %s", paramName)
			}
		}
	}
	return nil
}

// classifyEndpoint determines the endpoint type for rate limiting
func classifyEndpoint(path string, method string) EndpointType {
	// Statistics endpoints
	if strings.Contains(path, "/statistics") {
		return StatisticsAPI
	}

	// Send message endpoint
	if strings.Contains(path, "/messages") && method == "POST" {
		return SendMessageAPI
	}

	// Default to general API
	return GeneralAPI
}

// APIClient manages communication with the AhaSend API v2 API v2.0.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg               *Configuration
	common            service // Reuse a single struct instead of allocating one for each service on the heap.
	rateLimiter       *RateLimiter
	idempotencyHelper *IdempotencyHelper

	// API Services

	APIKeysAPI *APIKeysAPIService

	AccountsAPI *AccountsAPIService

	DomainsAPI *DomainsAPIService

	MessagesAPI *MessagesAPIService

	RoutesAPI *RoutesAPIService

	SMTPCredentialsAPI *SMTPCredentialsAPIService

	StatisticsAPI *StatisticsAPIService

	SuppressionsAPI *SuppressionsAPIService

	UtilityAPI *UtilityAPIService

	WebhooksAPI *WebhooksAPIService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client with functional options.
// Example: client := NewAPIClient(WithAPIKey("aha-sk-..."), WithDebug(true))
func NewAPIClient(opts ...ClientOption) *APIClient {
	cfg := NewConfiguration()
	for _, opt := range opts {
		opt(cfg)
	}
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

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *Configuration {
	return c.cfg
}

// Rate Limiting Public API Methods

// SetGeneralRateLimit sets the rate limit for general API endpoints
func (c *APIClient) SetGeneralRateLimit(requestsPerSecond, burstCapacity int) {
	c.rateLimiter.SetRateLimit(GeneralAPI, requestsPerSecond, burstCapacity)
}

// SetStatisticsRateLimit sets the rate limit for statistics API endpoints
func (c *APIClient) SetStatisticsRateLimit(requestsPerSecond, burstCapacity int) {
	c.rateLimiter.SetRateLimit(StatisticsAPI, requestsPerSecond, burstCapacity)
}

// SetSendMessageRateLimit sets the rate limit for send message API endpoint
func (c *APIClient) SetSendMessageRateLimit(requestsPerSecond, burstCapacity int) {
	c.rateLimiter.SetRateLimit(SendMessageAPI, requestsPerSecond, burstCapacity)
}

// SetCustomRateLimit sets the rate limit for a specific endpoint type
func (c *APIClient) SetCustomRateLimit(endpointType EndpointType, requestsPerSecond, burstCapacity int) {
	c.rateLimiter.SetRateLimit(endpointType, requestsPerSecond, burstCapacity)
}

// EnableRateLimit enables or disables rate limiting for a specific endpoint type
func (c *APIClient) EnableRateLimit(endpointType EndpointType, enabled bool) {
	c.rateLimiter.EnableRateLimit(endpointType, enabled)
}

// SetGlobalRateLimit enables or disables rate limiting globally
func (c *APIClient) SetGlobalRateLimit(enabled bool) {
	c.rateLimiter.SetGlobalEnabled(enabled)
}

// GetRateLimitStatus returns the current rate limit status for a specific endpoint type
func (c *APIClient) GetRateLimitStatus(endpointType EndpointType) RateLimitStatus {
	return c.rateLimiter.GetStatus(endpointType)
}

// ConfigureCustomerRateLimits applies a complete customer rate limit configuration
// This is a convenience method for enterprise customers to set all limits at once
func (c *APIClient) ConfigureCustomerRateLimits(config CustomerRateLimitConfig) {
	c.rateLimiter.ConfigureFromCustomerConfig(config)
}

// Idempotency Public API Methods

// GenerateIdempotencyKey generates a new UUID-based idempotency key
func (c *APIClient) GenerateIdempotencyKey() string {
	return c.idempotencyHelper.GenerateKey()
}

// SetIdempotencyConfig updates the idempotency configuration
func (c *APIClient) SetIdempotencyConfig(config IdempotencyConfig) {
	c.cfg.IdempotencyConfig = config
	c.idempotencyHelper = NewIdempotencyHelper(config)
}

// GetIdempotencyConfig returns the current idempotency configuration
func (c *APIClient) GetIdempotencyConfig() IdempotencyConfig {
	return c.cfg.IdempotencyConfig
}

// NewIdempotencyKeyBuilder creates a new idempotency key builder for related operations
func (c *APIClient) NewIdempotencyKeyBuilder(baseKey ...string) *IdempotencyKeyBuilder {
	var base string
	if len(baseKey) > 0 && baseKey[0] != "" {
		base = baseKey[0]
	} else {
		base = c.GenerateIdempotencyKey()
	}
	return NewIdempotencyKeyBuilder(base)
}

// Core Execute Method and Supporting Functions

// Execute is the centralized method for executing all API requests
func (c *APIClient) Execute(ctx context.Context, config RequestConfig) (*http.Response, error) {
	// Step 1: Validate and build the path
	if err := validatePathParams(config.PathTemplate, config.PathParams); err != nil {
		return nil, &APIError{
			Type:    ErrorTypeValidation,
			Message: fmt.Sprintf("Invalid path parameters: %v", err),
		}
	}

	path, err := buildPath(config.PathTemplate, config.PathParams)
	if err != nil {
		return nil, &APIError{
			Type:    ErrorTypeValidation,
			Message: fmt.Sprintf("Failed to build path: %v", err),
		}
	}

	// Step 2: Classify endpoint for rate limiting
	config.endpointType = classifyEndpoint(path, config.Method)

	// Step 3: Build the full URL
	var baseURL string

	// Check if Host and Scheme are explicitly set (for testing/custom configurations)
	if c.cfg.Host != "" && c.cfg.Scheme != "" {
		baseURL = fmt.Sprintf("%s://%s", c.cfg.Scheme, c.cfg.Host)
	} else {
		// Fall back to ServerURLWithContext for default server selection
		var err error
		baseURL, err = c.cfg.ServerURLWithContext(ctx, "")
		if err != nil {
			return nil, &NetworkError{Op: "server selection", Err: err}
		}
	}

	fullURL := baseURL + path
	if len(config.QueryParams) > 0 {
		fullURL += "?" + config.QueryParams.Encode()
	}

	// Step 4: Create the request body
	var bodyReader io.Reader
	if config.Body != nil {
		jsonBody, err := json.Marshal(config.Body)
		if err != nil {
			return nil, &APIError{
				Type:    ErrorTypeValidation,
				Message: fmt.Sprintf("Failed to encode request body: %v", err),
			}
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	// Step 5: Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, config.Method, fullURL, bodyReader)
	if err != nil {
		return nil, &NetworkError{Op: "request creation", Err: err}
	}

	// Step 6: Apply headers
	if err := c.applyHeaders(ctx, req, config); err != nil {
		return nil, err
	}

	// Step 7: Apply rate limiting (unless skipped)
	if !config.SkipRateLimit && c.rateLimiter != nil {
		if err := c.applyRateLimit(ctx, config.endpointType); err != nil {
			return nil, err
		}
	}

	// Step 8: Execute with retry logic
	resp, err := c.executeWithRetry(ctx, req, config)
	if err != nil {
		return nil, err
	}

	// Step 9: Read response body
	responseBody, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return resp, &NetworkError{Op: "reading response", Err: err}
	}

	// Replace body with new reader so it can be read again if needed
	resp.Body = io.NopCloser(bytes.NewBuffer(responseBody))

	// Step 10: Handle errors for non-2xx responses
	if resp.StatusCode >= 300 {
		return resp, c.handleErrorResponse(resp, responseBody, config.Method, path)
	}

	// Step 11: Decode successful response into result
	if config.Result != nil && len(responseBody) > 0 {
		if err := json.Unmarshal(responseBody, config.Result); err != nil {
			return resp, &APIError{
				Type:    ErrorTypeUnknown,
				Message: fmt.Sprintf("Failed to decode response: %v", err),
				Raw:     responseBody,
			}
		}
	}

	return resp, nil
}

// applyHeaders applies headers to the request with authentication hierarchy
func (c *APIClient) applyHeaders(ctx context.Context, req *http.Request, config RequestConfig) error {
	// Set default headers
	req.Header.Set("Accept", "application/json")
	if config.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Apply user agent
	if c.cfg.UserAgent != "" {
		req.Header.Set("User-Agent", c.cfg.UserAgent)
	}

	// Apply authentication - check multiple sources in order of precedence
	authApplied := false

	// 1. Check request-specific headers (highest priority)
	if auth, exists := config.Headers["Authorization"]; exists {
		req.Header.Set("Authorization", auth)
		authApplied = true
	}

	// 2. Check context override (medium priority)
	if !authApplied {
		if auth := ctx.Value(ContextAccessToken); auth != nil {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", auth))
			authApplied = true
		}
	}

	// 3. Use client-level configuration (lowest priority)
	if !authApplied && c.cfg.APIKey != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.cfg.APIKey))
		authApplied = true
	}

	// Fail if no authentication provided
	if !authApplied {
		return &APIError{
			Type:    ErrorTypeAuthentication,
			Message: "No API key provided. Set via client configuration, context, or request options",
		}
	}

	// Apply idempotency key from context (if present)
	if idempotencyKey := ctx.Value(ContextIdempotencyKey); idempotencyKey != nil {
		req.Header.Set("Idempotency-Key", idempotencyKey.(string))
	}

	// Auto-generate idempotency key for POST requests if not already set and auto-generation is enabled
	if config.Method == http.MethodPost {
		if req.Header.Get("Idempotency-Key") == "" && c.cfg.IdempotencyConfig.AutoGenerate {
			req.Header.Set("Idempotency-Key", c.idempotencyHelper.GenerateKey())
		}
	}

	// Apply any default headers from configuration
	for key, value := range c.cfg.DefaultHeader {
		req.Header.Set(key, value)
	}

	// Apply request-specific headers (these override defaults)
	for key, value := range config.Headers {
		// Skip Authorization as we handled it above
		if key != "Authorization" {
			req.Header.Set(key, value)
		}
	}

	return nil
}

// applyRateLimit applies rate limiting for the specified endpoint type
func (c *APIClient) applyRateLimit(ctx context.Context, endpointType EndpointType) error {
	if c.rateLimiter == nil {
		return nil
	}

	bucket := c.rateLimiter.GetBucket(endpointType)
	if bucket == nil {
		return nil
	}

	// Wait for rate limit token
	if err := bucket.WaitForTokenWithContext(ctx); err != nil {
		if errors.Is(err, context.Canceled) {
			return &NetworkError{Op: "rate limiting", Err: err}
		}
		return &APIError{
			Type:    ErrorTypeRateLimit,
			Message: "Rate limit exceeded",
		}
	}

	return nil
}

// executeWithRetry executes the request with retry logic and proper body preservation
func (c *APIClient) executeWithRetry(ctx context.Context, req *http.Request, config RequestConfig) (*http.Response, error) {
	// Determine retry configuration
	var retryConfig RetryConfig
	if config.CustomRetry != nil {
		retryConfig = *config.CustomRetry
	} else {
		retryConfig = c.cfg.RetryConfig
	}

	// If retries are disabled, execute once
	if !retryConfig.IsRetryEnabled() {
		return c.cfg.HTTPClient.Do(req)
	}

	// CRITICAL FIX: Store original body bytes for retry attempts
	var bodyBytes []byte
	if req.Body != nil {
		var err error
		bodyBytes, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, &NetworkError{Op: "reading request body", Err: err}
		}
		// Reset the original request body
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	}

	var lastErr error
	var lastResp *http.Response

	for attempt := 0; attempt <= retryConfig.MaxRetries; attempt++ {
		// Clone the request for each attempt
		reqClone := req.Clone(ctx)
		if bodyBytes != nil {
			// Recreate body from stored bytes
			reqClone.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		}

		// Execute the request
		resp, err := c.cfg.HTTPClient.Do(reqClone)

		// Check if we should retry
		if !c.shouldRetry(err, resp, retryConfig, attempt) {
			return resp, err
		}

		lastErr = err
		lastResp = resp

		// Calculate delay before next attempt
		if attempt < retryConfig.MaxRetries {
			delay := retryConfig.GetDelay(attempt + 1)

			select {
			case <-ctx.Done():
				return lastResp, ctx.Err()
			case <-time.After(delay):
				// Continue to next attempt
			}
		}
	}

	// All retries exhausted
	if lastErr != nil {
		return lastResp, lastErr
	}

	return lastResp, &APIError{
		Type:       ErrorTypeServer,
		StatusCode: lastResp.StatusCode,
		Message:    fmt.Sprintf("Request failed after %d retries", retryConfig.MaxRetries),
	}
}

// shouldRetry determines whether a request should be retried based on error and response
func (c *APIClient) shouldRetry(err error, resp *http.Response, config RetryConfig, attempt int) bool {
	// Don't retry if we've exhausted attempts
	if attempt >= config.MaxRetries {
		return false
	}

	// Network errors are always retryable
	if err != nil {
		return true
	}

	if resp == nil {
		return false
	}

	// Check status code
	switch {
	case resp.StatusCode == 429: // Rate limit
		return true
	case resp.StatusCode >= 500: // Server errors
		return true
	case resp.StatusCode >= 400 && resp.StatusCode < 500: // Client errors
		return config.RetryClientErrors
	default:
		return false
	}
}

// handleErrorResponse creates structured error information from HTTP error responses
func (c *APIClient) handleErrorResponse(resp *http.Response, body []byte, method, path string) error {
	// Create base API error
	apiErr := &APIError{
		StatusCode: resp.StatusCode,
		Type:       determineErrorType(resp.StatusCode),
		Method:     method,
		Endpoint:   path,
		Raw:        body,
	}

	// Extract request ID from headers
	if reqID := resp.Header.Get("X-Request-Id"); reqID != "" {
		apiErr.RequestID = reqID
	}

	// Extract retry-after for rate limit errors
	if apiErr.Type == ErrorTypeRateLimit {
		if retryAfter := resp.Header.Get("Retry-After"); retryAfter != "" {
			if seconds, err := strconv.Atoi(retryAfter); err == nil {
				apiErr.RetryAfter = seconds
			}
		}
	}

	// Try to parse error response body
	var errorResp common.ErrorResponse
	if err := json.Unmarshal(body, &errorResp); err == nil {
		apiErr.Message = errorResp.Message
		// Note: ErrorResponse doesn't have a Code field, using Message for now

		// Note: ErrorResponse doesn't have Field or Resource fields
	} else {
		// Fallback to HTTP status text if parsing fails
		apiErr.Message = http.StatusText(resp.StatusCode)
	}

	// Parse message context and generate suggestions
	apiErr.parseMessageContext()
	apiErr.generateSuggestions()

	return apiErr
}
