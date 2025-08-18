/*
AhaSend Go SDK - Rate Limiting

Provides intelligent rate limiting for the AhaSend API with customer-configurable limits.
Supports different rate limits for general API, statistics, and send message endpoints.
*/

package api

import (
	"context"
	"math"
	"strings"
	"sync"
	"time"
)

// EndpointType represents different categories of API endpoints with distinct rate limits
type EndpointType int

const (
	GeneralAPI     EndpointType = iota // General API endpoints (default: 100 req/s, 200 burst)
	StatisticsAPI                      // Statistics endpoints (default: 1 req/s, 1 burst)
	SendMessageAPI                     // Send message endpoint (default: 100 req/s, 200 burst, often customized)
)

// String returns the string representation of EndpointType
func (e EndpointType) String() string {
	switch e {
	case GeneralAPI:
		return "general"
	case StatisticsAPI:
		return "statistics"
	case SendMessageAPI:
		return "send_message"
	default:
		return "unknown"
	}
}

// RateLimitConfig defines the configuration for a rate limiter
type RateLimitConfig struct {
	RequestsPerSecond int  `json:"requests_per_second"`
	BurstCapacity     int  `json:"burst_capacity"`
	Enabled           bool `json:"enabled"`
}

// CustomerRateLimitConfig allows configuring all rate limits at once
type CustomerRateLimitConfig struct {
	General     *RateLimitConfig `json:"general,omitempty"`
	Statistics  *RateLimitConfig `json:"statistics,omitempty"`
	SendMessage *RateLimitConfig `json:"send_message,omitempty"`
}

// RateLimitStatus provides current status information about a rate limiter
type RateLimitStatus struct {
	EndpointType      EndpointType `json:"endpoint_type"`
	Enabled           bool         `json:"enabled"`
	RequestsPerSecond int          `json:"requests_per_second"`
	BurstCapacity     int          `json:"burst_capacity"`
	TokensAvailable   int          `json:"tokens_available"`
	NextRefillTime    time.Time    `json:"next_refill_time"`
}

// TokenBucket implements a token bucket rate limiter
type TokenBucket struct {
	tokens          float64
	maxTokens       int
	tokensPerSecond float64
	lastRefill      time.Time
	config          RateLimitConfig
	mu              sync.Mutex
}

// NewTokenBucket creates a new token bucket with the specified configuration
func NewTokenBucket(config RateLimitConfig) *TokenBucket {
	return &TokenBucket{
		tokens:          float64(config.BurstCapacity),
		maxTokens:       config.BurstCapacity,
		tokensPerSecond: float64(config.RequestsPerSecond),
		lastRefill:      time.Now(),
		config:          config,
	}
}

// WaitForToken blocks until a token is available, respecting the rate limit
func (tb *TokenBucket) WaitForToken() error {
	return tb.WaitForTokenWithContext(context.Background())
}

// WaitForTokenWithContext blocks until a token is available, respecting the rate limit and context cancellation
func (tb *TokenBucket) WaitForTokenWithContext(ctx context.Context) error {
	tb.mu.Lock()

	if !tb.config.Enabled {
		tb.mu.Unlock()
		return nil
	}

	// Refill tokens based on elapsed time
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill)
	tokensToAdd := elapsed.Seconds() * tb.tokensPerSecond

	tb.tokens = math.Min(tb.tokens+tokensToAdd, float64(tb.maxTokens))
	tb.lastRefill = now

	// If we have tokens available, consume one and return
	if tb.tokens >= 1.0 {
		tb.tokens -= 1.0
		tb.mu.Unlock()
		return nil
	}

	// Calculate wait time for the next token
	waitTime := time.Duration((1.0-tb.tokens)/tb.tokensPerSecond*float64(time.Second)) + time.Millisecond
	tb.mu.Unlock()

	// Wait for either the timeout or context cancellation
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(waitTime):
		// Continue to acquire token
	}

	// Re-acquire lock and consume the token
	tb.mu.Lock()
	defer tb.mu.Unlock()

	// Refill again after waiting (in case more time passed than expected)
	now = time.Now()
	elapsed = now.Sub(tb.lastRefill)
	tokensToAdd = elapsed.Seconds() * tb.tokensPerSecond
	tb.tokens = math.Min(tb.tokens+tokensToAdd, float64(tb.maxTokens))
	tb.lastRefill = now

	// Consume a token
	if tb.tokens >= 1.0 {
		tb.tokens -= 1.0
	} else {
		// This shouldn't happen but handle gracefully
		tb.tokens = 0
	}

	return nil
}

// UpdateConfig updates the rate limit configuration
func (tb *TokenBucket) UpdateConfig(config RateLimitConfig) {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.config = config
	tb.tokensPerSecond = float64(config.RequestsPerSecond)
	tb.maxTokens = config.BurstCapacity

	// Adjust current tokens if new max is lower
	if tb.tokens > float64(tb.maxTokens) {
		tb.tokens = float64(tb.maxTokens)
	}
}

// GetStatus returns the current status of the token bucket
func (tb *TokenBucket) GetStatus(endpointType EndpointType) RateLimitStatus {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	// Calculate next refill time
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill)
	tokensToAdd := elapsed.Seconds() * tb.tokensPerSecond
	currentTokens := math.Min(tb.tokens+tokensToAdd, float64(tb.maxTokens))

	var nextRefill time.Time
	if currentTokens < float64(tb.maxTokens) {
		tokensNeeded := float64(tb.maxTokens) - currentTokens
		secondsToWait := tokensNeeded / tb.tokensPerSecond
		nextRefill = now.Add(time.Duration(secondsToWait * float64(time.Second)))
	}

	return RateLimitStatus{
		EndpointType:      endpointType,
		Enabled:           tb.config.Enabled,
		RequestsPerSecond: tb.config.RequestsPerSecond,
		BurstCapacity:     tb.config.BurstCapacity,
		TokensAvailable:   int(math.Floor(currentTokens)),
		NextRefillTime:    nextRefill,
	}
}

// RateLimiter manages rate limiting for different endpoint types
type RateLimiter struct {
	general       *TokenBucket
	statistics    *TokenBucket
	sendMessage   *TokenBucket
	globalEnabled bool
	mu            sync.RWMutex
}

// NewRateLimiter creates a new rate limiter with default configurations
func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		general: NewTokenBucket(RateLimitConfig{
			RequestsPerSecond: 100,
			BurstCapacity:     200,
			Enabled:           true,
		}),
		statistics: NewTokenBucket(RateLimitConfig{
			RequestsPerSecond: 1,
			BurstCapacity:     1,
			Enabled:           true,
		}),
		sendMessage: NewTokenBucket(RateLimitConfig{
			RequestsPerSecond: 100,
			BurstCapacity:     200,
			Enabled:           true,
		}),
		globalEnabled: true,
	}
}

// DetectEndpointType determines the endpoint type based on HTTP method and path
func (rl *RateLimiter) DetectEndpointType(method, path string) EndpointType {
	// Send message endpoint: POST /v2/accounts/{id}/messages (not /messages/{id})
	if method == "POST" && strings.Contains(path, "/messages") &&
		!strings.Contains(path, "/messages/") {
		return SendMessageAPI
	}

	// Statistics endpoints: /v2/accounts/{id}/statistics/*
	if strings.Contains(path, "/statistics/") {
		return StatisticsAPI
	}

	return GeneralAPI
}

// GetBucket returns the appropriate token bucket for the endpoint type
func (rl *RateLimiter) GetBucket(endpointType EndpointType) *TokenBucket {
	switch endpointType {
	case GeneralAPI:
		return rl.general
	case StatisticsAPI:
		return rl.statistics
	case SendMessageAPI:
		return rl.sendMessage
	default:
		return rl.general
	}
}

// WaitForToken blocks until a token is available for the specified request
func (rl *RateLimiter) WaitForToken(method, path string) error {
	return rl.WaitForTokenWithContext(context.Background(), method, path)
}

// WaitForTokenWithContext blocks until a token is available for the specified request, respecting context cancellation
func (rl *RateLimiter) WaitForTokenWithContext(ctx context.Context, method, path string) error {
	if !rl.IsEnabled() {
		return nil
	}

	endpointType := rl.DetectEndpointType(method, path)
	bucket := rl.GetBucket(endpointType)

	return bucket.WaitForTokenWithContext(ctx)
}

// SetRateLimit updates the rate limit for a specific endpoint type
func (rl *RateLimiter) SetRateLimit(endpointType EndpointType, requestsPerSecond, burstCapacity int) {
	config := RateLimitConfig{
		RequestsPerSecond: requestsPerSecond,
		BurstCapacity:     burstCapacity,
		Enabled:           true,
	}

	bucket := rl.GetBucket(endpointType)
	if bucket != nil {
		bucket.UpdateConfig(config)
	}
}

// EnableRateLimit enables or disables rate limiting for a specific endpoint type
func (rl *RateLimiter) EnableRateLimit(endpointType EndpointType, enabled bool) {
	bucket := rl.GetBucket(endpointType)
	if bucket != nil {
		bucket.mu.Lock()
		bucket.config.Enabled = enabled
		bucket.mu.Unlock()
	}
}

// SetGlobalEnabled enables or disables rate limiting globally
func (rl *RateLimiter) SetGlobalEnabled(enabled bool) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	rl.globalEnabled = enabled
}

// IsEnabled returns whether rate limiting is globally enabled
func (rl *RateLimiter) IsEnabled() bool {
	rl.mu.RLock()
	defer rl.mu.RUnlock()
	return rl.globalEnabled
}

// GetStatus returns the current status for a specific endpoint type
func (rl *RateLimiter) GetStatus(endpointType EndpointType) RateLimitStatus {
	bucket := rl.GetBucket(endpointType)
	if bucket != nil {
		return bucket.GetStatus(endpointType)
	}

	return RateLimitStatus{
		EndpointType: endpointType,
		Enabled:      false,
	}
}

// ConfigureFromCustomerConfig applies a customer rate limit configuration
func (rl *RateLimiter) ConfigureFromCustomerConfig(config CustomerRateLimitConfig) {
	if config.General != nil {
		rl.general.UpdateConfig(*config.General)
	}
	if config.Statistics != nil {
		rl.statistics.UpdateConfig(*config.Statistics)
	}
	if config.SendMessage != nil {
		rl.sendMessage.UpdateConfig(*config.SendMessage)
	}
}
