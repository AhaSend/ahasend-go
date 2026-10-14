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
	// lastAct is when the last caller in the queue leaves
	lastAct time.Time
	// changed is closed, and replaced, when the configuration changes, to
	// wake the callers that wait
	changed chan struct{}
}

// NewTokenBucket creates a new token bucket with the specified configuration
func NewTokenBucket(config RateLimitConfig) *TokenBucket {
	return &TokenBucket{
		tokens:          float64(config.BurstCapacity),
		maxTokens:       config.BurstCapacity,
		tokensPerSecond: float64(config.RequestsPerSecond),
		lastRefill:      time.Now(),
		config:          config,
		changed:         make(chan struct{}),
	}
}

// WaitForToken blocks until a token is available, respecting the rate limit
func (tb *TokenBucket) WaitForToken() error {
	return tb.WaitForTokenWithContext(context.Background())
}

// WaitForTokenWithContext blocks until a token is available, respecting the rate limit and context cancellation.
//
// Each caller takes its token when it arrives, even when none is left: the
// bucket then goes below zero, and the caller waits until the refill has
// covered its place in the queue. Callers that wait at the same time
// therefore leave one after another at the configured rate, never together.
// A caller whose context would end before its turn returns at once with
// context.DeadlineExceeded. A caller whose context ends while it waits gives
// its token back when it is the last in the queue; a place in the middle of
// the queue stays used, so the callers after it never leave early. Disabling
// the bucket lets the waiting callers go.
func (tb *TokenBucket) WaitForTokenWithContext(ctx context.Context) error {
	tb.mu.Lock()

	// A rate of zero or less, which can be set at run time, would never
	// refill the bucket: it is treated as no limit
	if !tb.config.Enabled || tb.tokensPerSecond <= 0 {
		tb.mu.Unlock()
		return nil
	}

	now := time.Now()
	tb.refill(now)
	tb.tokens--
	if tb.tokens >= 0 {
		tb.mu.Unlock()
		return nil
	}
	actAt := now.Add(time.Duration(-tb.tokens / tb.tokensPerSecond * float64(time.Second)))
	if deadline, ok := ctx.Deadline(); ok && actAt.After(deadline) {
		// This caller is the last in the queue: its token goes back whole
		tb.tokens++
		tb.mu.Unlock()
		return context.DeadlineExceeded
	}
	tb.lastAct = actAt
	changed := tb.changed
	tb.mu.Unlock()

	timer := time.NewTimer(time.Until(actAt))
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			return nil
		case <-ctx.Done():
			tb.giveBack(actAt)
			return ctx.Err()
		case <-changed:
			tb.mu.Lock()
			enabled := tb.config.Enabled
			changed = tb.changed
			tb.mu.Unlock()
			if !enabled {
				return nil
			}
		}
	}
}

// refill adds the tokens earned since the last refill, up to the burst
// capacity. The caller holds tb.mu.
func (tb *TokenBucket) refill(now time.Time) {
	elapsed := now.Sub(tb.lastRefill)
	tb.tokens = math.Min(tb.tokens+elapsed.Seconds()*tb.tokensPerSecond, float64(tb.maxTokens))
	tb.lastRefill = now
}

// giveBack returns the token of a caller that stopped waiting for its turn
// at actAt, when that caller is the last in the queue. A caller with others
// after it keeps its place used: giving its token back would let the next
// caller leave together with the last one.
func (tb *TokenBucket) giveBack(actAt time.Time) {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	if !actAt.Equal(tb.lastAct) || tb.tokensPerSecond <= 0 {
		return
	}
	tb.tokens = math.Min(tb.tokens+1, float64(tb.maxTokens))
	tb.lastAct = actAt.Add(-time.Duration(float64(time.Second) / tb.tokensPerSecond))
}

// notifyChanged wakes the callers that wait, so they see a new
// configuration. The caller holds tb.mu.
func (tb *TokenBucket) notifyChanged() {
	if tb.changed != nil {
		close(tb.changed)
	}
	tb.changed = make(chan struct{})
}

// UpdateConfig updates the rate limit configuration. A queue of callers
// built up at the old rate keeps its schedule; the callers after it wait for
// it to clear and then follow the new rate.
func (tb *TokenBucket) UpdateConfig(config RateLimitConfig) {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	oldRate := tb.tokensPerSecond
	if oldRate > 0 {
		tb.refill(time.Now())
	} else {
		tb.lastRefill = time.Now()
	}
	newRate := float64(config.RequestsPerSecond)
	// A shortfall counts the callers queued at the old rate: expressed at
	// the new rate, it lasts as long as that queue does
	if tb.tokens < 0 && oldRate > 0 && newRate > 0 {
		tb.tokens = tb.tokens * newRate / oldRate
	}

	tb.config = config
	tb.tokensPerSecond = newRate
	tb.maxTokens = config.BurstCapacity

	// Adjust current tokens if new max is lower
	if tb.tokens > float64(tb.maxTokens) {
		tb.tokens = float64(tb.maxTokens)
	}
	tb.notifyChanged()
}

// setEnabled turns the bucket on or off. Turning it off lets the callers
// that wait go.
func (tb *TokenBucket) setEnabled(enabled bool) {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	tb.config.Enabled = enabled
	tb.notifyChanged()
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
		TokensAvailable:   int(math.Max(math.Floor(currentTokens), 0)),
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
		bucket.setEnabled(enabled)
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
