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

// TokenBucket implements a token bucket rate limiter.
//
// Callers that find no token wait in a first-come, first-served queue. Only
// the caller at the front waits for the next token; when it takes it, the
// next caller moves to the front. Callers that wait at the same time
// therefore leave one after another at the configured rate, never together,
// and a caller that stops waiting leaves the queue and frees its place.
type TokenBucket struct {
	tokens          float64
	maxTokens       int
	tokensPerSecond float64
	lastRefill      time.Time
	config          RateLimitConfig
	mu              sync.Mutex
	// queue holds the callers waiting for a token, the front one first
	queue []*tokenWaiter
	// changed is closed, and replaced, when the configuration changes, to
	// wake the caller at the front of the queue
	changed chan struct{}
}

// tokenWaiter is a caller in a TokenBucket's queue. Its turn channel is
// closed when it reaches the front.
type tokenWaiter struct {
	turn chan struct{}
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

// WaitForTokenWithContext blocks until a token is available, respecting the
// rate limit and context cancellation: a caller whose context ends while it
// waits leaves the queue and returns the context's error. A disabled bucket,
// or one whose rate is zero or less (which can be set at run time, and would
// never refill), does not limit: disabling it, or setting such a rate, also
// lets the waiting callers go.
func (tb *TokenBucket) WaitForTokenWithContext(ctx context.Context) error {
	tb.mu.Lock()
	if !tb.limits() {
		tb.mu.Unlock()
		return nil
	}
	now := time.Now()
	tb.refill(now)
	if len(tb.queue) == 0 && tb.tokens >= 1 {
		tb.tokens--
		tb.mu.Unlock()
		return nil
	}
	w := &tokenWaiter{turn: make(chan struct{})}
	tb.queue = append(tb.queue, w)
	if len(tb.queue) == 1 {
		close(w.turn)
	}
	tb.mu.Unlock()

	select {
	case <-w.turn:
	case <-ctx.Done():
		tb.leave(w)
		return ctx.Err()
	}

	// At the front: wait for the next token
	for {
		tb.mu.Lock()
		if !tb.limits() {
			tb.popFront()
			tb.mu.Unlock()
			return nil
		}
		tb.refill(time.Now())
		if tb.tokens >= 1 {
			tb.tokens--
			tb.popFront()
			tb.mu.Unlock()
			return nil
		}
		wait := tb.durationFor(1 - tb.tokens)
		changed := tb.changed
		tb.mu.Unlock()

		timer := time.NewTimer(wait)
		select {
		case <-timer.C:
		case <-changed:
			timer.Stop()
		case <-ctx.Done():
			timer.Stop()
			tb.mu.Lock()
			tb.popFront()
			tb.mu.Unlock()
			return ctx.Err()
		}
	}
}

// limits reports whether the bucket limits its callers: it is enabled and
// refills. The caller holds tb.mu.
func (tb *TokenBucket) limits() bool {
	return tb.config.Enabled && tb.tokensPerSecond > 0
}

// durationFor returns how long the bucket takes to refill the given number
// of tokens. The caller holds tb.mu.
func (tb *TokenBucket) durationFor(tokens float64) time.Duration {
	return time.Duration(tokens / tb.tokensPerSecond * float64(time.Second))
}

// capacity is the most tokens the bucket holds: its burst capacity, and at
// least one, so a burst of zero or less, which the configuration validation
// refuses but can be set at run time, still lets one caller through per
// interval. The caller holds tb.mu.
func (tb *TokenBucket) capacity() float64 {
	return math.Max(float64(tb.maxTokens), 1)
}

// refill adds the tokens earned since the last refill, up to the bucket's
// capacity. The caller holds tb.mu.
func (tb *TokenBucket) refill(now time.Time) {
	if tb.tokensPerSecond > 0 {
		elapsed := now.Sub(tb.lastRefill)
		tb.tokens = math.Min(tb.tokens+elapsed.Seconds()*tb.tokensPerSecond, tb.capacity())
	}
	tb.lastRefill = now
}

// popFront removes the caller at the front of the queue and gives the next
// one its turn. The caller holds tb.mu.
func (tb *TokenBucket) popFront() {
	tb.queue[0] = nil
	tb.queue = tb.queue[1:]
	if len(tb.queue) > 0 {
		close(tb.queue[0].turn)
	}
}

// leave removes a caller that stopped waiting for its turn from the queue.
// A caller that reached the front meanwhile hands its turn on.
func (tb *TokenBucket) leave(w *tokenWaiter) {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	for i, queued := range tb.queue {
		if queued != w {
			continue
		}
		if i == 0 {
			tb.popFront()
			return
		}
		tb.queue = append(tb.queue[:i], tb.queue[i+1:]...)
		return
	}
}

// notifyChanged wakes the caller at the front of the queue, so it sees a new
// configuration. The caller holds tb.mu.
func (tb *TokenBucket) notifyChanged() {
	if tb.changed != nil {
		close(tb.changed)
	}
	tb.changed = make(chan struct{})
}

// UpdateConfig updates the rate limit configuration. The tokens earned so far
// are counted at the old rate; the callers that wait follow the new rate from
// then on.
func (tb *TokenBucket) UpdateConfig(config RateLimitConfig) {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill(time.Now())
	tb.config = config
	tb.tokensPerSecond = float64(config.RequestsPerSecond)
	tb.maxTokens = config.BurstCapacity

	// Adjust current tokens if new max is lower
	if tb.tokens > tb.capacity() {
		tb.tokens = tb.capacity()
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

// GetStatus returns the current status of the token bucket. A bucket that
// does not limit (disabled, or with a rate of zero or less) reports no next
// refill time.
func (tb *TokenBucket) GetStatus(endpointType EndpointType) RateLimitStatus {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	currentTokens := tb.tokens
	if tb.tokensPerSecond > 0 {
		elapsed := now.Sub(tb.lastRefill)
		currentTokens = math.Min(currentTokens+elapsed.Seconds()*tb.tokensPerSecond, tb.capacity())
	}

	var nextRefill time.Time
	if tb.limits() && currentTokens < tb.capacity() {
		nextRefill = now.Add(tb.durationFor(tb.capacity() - currentTokens))
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
