package api

import (
	"context"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestConcurrentWaitersLeaveAtTheConfiguredRate checks that callers waiting
// at the same time are let through one after another at the bucket's rate,
// not all together when the next token is due.
func TestConcurrentWaitersLeaveAtTheConfiguredRate(t *testing.T) {
	const (
		rate    = 20
		burst   = 2
		callers = 12
	)
	bucket := NewTokenBucket(RateLimitConfig{RequestsPerSecond: rate, BurstCapacity: burst, Enabled: true})

	start := time.Now()
	var mu sync.Mutex
	var released []time.Duration
	var wg sync.WaitGroup
	for i := 0; i < callers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			require.NoError(t, bucket.WaitForToken())
			mu.Lock()
			released = append(released, time.Since(start))
			mu.Unlock()
		}()
	}
	wg.Wait()

	sort.Slice(released, func(i, j int) bool { return released[i] < released[j] })
	// The burst leaves at once; caller k after it leaves no sooner than
	// k/rate seconds in, give or take timer slack
	const slack = 15 * time.Millisecond
	for k := burst; k < callers; k++ {
		due := time.Duration(k-burst+1) * time.Second / rate
		assert.GreaterOrEqual(t, released[k], due-slack, "caller %d left at %v, before its turn at %v", k, released[k], due)
	}
}

// TestTheLastWaiterGivesItsTokenBack checks that the last caller in the
// queue, cancelled while it waits, does not hold back the callers after it.
func TestTheLastWaiterGivesItsTokenBack(t *testing.T) {
	bucket := NewTokenBucket(RateLimitConfig{RequestsPerSecond: 5, BurstCapacity: 1, Enabled: true})
	require.NoError(t, bucket.WaitForToken())

	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(20*time.Millisecond, cancel)
	assert.ErrorIs(t, bucket.WaitForTokenWithContext(ctx), context.Canceled)

	// The next token is due 200ms after the first; had the cancelled caller
	// kept its place, this one would wait until 400ms
	start := time.Now()
	require.NoError(t, bucket.WaitForToken())
	assert.Less(t, time.Since(start), 300*time.Millisecond)
}

// TestAWaiterInTheMiddleKeepsTheQueueSpaced checks that a caller cancelled
// with others queued after it does not let the next caller leave together
// with the last one.
func TestAWaiterInTheMiddleKeepsTheQueueSpaced(t *testing.T) {
	bucket := NewTokenBucket(RateLimitConfig{RequestsPerSecond: 10, BurstCapacity: 1, Enabled: true})
	require.NoError(t, bucket.WaitForToken())

	middle, cancel := context.WithCancel(context.Background())
	middleDone := make(chan error, 1)
	go func() { middleDone <- bucket.WaitForTokenWithContext(middle) }()
	time.Sleep(5 * time.Millisecond)

	lastLeft := make(chan time.Time, 1)
	go func() {
		assert.NoError(t, bucket.WaitForToken())
		lastLeft <- time.Now()
	}()
	time.Sleep(5 * time.Millisecond)

	cancel()
	assert.ErrorIs(t, <-middleDone, context.Canceled)
	require.NoError(t, bucket.WaitForToken())
	newcomerLeft := time.Now()

	gap := newcomerLeft.Sub(<-lastLeft)
	if gap < 0 {
		gap = -gap
	}
	assert.GreaterOrEqual(t, gap, 90*time.Millisecond, "two callers left %v apart at 10 requests a second", gap)
}

// TestAWaiterWhoseDeadlineComesFirstFailsAtOnce checks that a caller whose
// context ends before its turn returns at once instead of holding a place.
func TestAWaiterWhoseDeadlineComesFirstFailsAtOnce(t *testing.T) {
	bucket := NewTokenBucket(RateLimitConfig{RequestsPerSecond: 1, BurstCapacity: 1, Enabled: true})
	require.NoError(t, bucket.WaitForToken())

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	start := time.Now()
	assert.ErrorIs(t, bucket.WaitForTokenWithContext(ctx), context.DeadlineExceeded)
	assert.Less(t, time.Since(start), 50*time.Millisecond)
}

// TestDisablingLetsTheWaitingCallersGo checks that turning a bucket off
// releases the callers already waiting in it.
func TestDisablingLetsTheWaitingCallersGo(t *testing.T) {
	bucket := NewTokenBucket(RateLimitConfig{RequestsPerSecond: 1, BurstCapacity: 1, Enabled: true})
	require.NoError(t, bucket.WaitForToken())

	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			assert.NoError(t, bucket.WaitForToken())
		}()
	}
	time.Sleep(20 * time.Millisecond)
	bucket.setEnabled(false)
	wg.Wait()
	assert.Less(t, time.Since(start), 500*time.Millisecond)
}

// TestLoweringTheRateKeepsTheQueueItsSchedule checks that a caller arriving
// after the rate is lowered waits for the queue built at the old rate and
// then one interval of the new rate, not for that queue at the new rate.
func TestLoweringTheRateKeepsTheQueueItsSchedule(t *testing.T) {
	bucket := NewTokenBucket(RateLimitConfig{RequestsPerSecond: 100, BurstCapacity: 1, Enabled: true})
	var wg sync.WaitGroup
	for i := 0; i < 21; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			assert.NoError(t, bucket.WaitForToken())
		}()
	}
	time.Sleep(10 * time.Millisecond)
	bucket.UpdateConfig(RateLimitConfig{RequestsPerSecond: 10, BurstCapacity: 1, Enabled: true})

	// The queue clears about 200ms in; the next turn at 10/s comes 100ms
	// later. At the new rate the same queue would take 2s
	start := time.Now()
	require.NoError(t, bucket.WaitForToken())
	assert.Less(t, time.Since(start), 600*time.Millisecond)
	wg.Wait()
}

func TestRateLimiterCreation(t *testing.T) {
	rl := NewRateLimiter()
	require.NotNil(t, rl)

	// Test default configuration
	assert.True(t, rl.IsEnabled())

	// Test status for different endpoint types
	generalStatus := rl.GetStatus(GeneralAPI)
	assert.Equal(t, GeneralAPI, generalStatus.EndpointType)
	assert.True(t, generalStatus.Enabled)
	assert.Equal(t, 100, generalStatus.RequestsPerSecond)
	assert.Equal(t, 200, generalStatus.BurstCapacity)

	statsStatus := rl.GetStatus(StatisticsAPI)
	assert.Equal(t, StatisticsAPI, statsStatus.EndpointType)
	assert.True(t, statsStatus.Enabled)
	assert.Equal(t, 1, statsStatus.RequestsPerSecond)
	assert.Equal(t, 1, statsStatus.BurstCapacity)

	sendMessageStatus := rl.GetStatus(SendMessageAPI)
	assert.Equal(t, SendMessageAPI, sendMessageStatus.EndpointType)
	assert.True(t, sendMessageStatus.Enabled)
	assert.Equal(t, 100, sendMessageStatus.RequestsPerSecond)
	assert.Equal(t, 200, sendMessageStatus.BurstCapacity)
}

func TestEndpointDetection(t *testing.T) {
	rl := NewRateLimiter()

	tests := []struct {
		method       string
		path         string
		expectedType EndpointType
	}{
		{"POST", "/v2/accounts/123/messages", SendMessageAPI},
		{"GET", "/v2/accounts/123/messages/456", GeneralAPI}, // Specific message, not send
		{"GET", "/v2/accounts/123/statistics/bounce", StatisticsAPI},
		{"GET", "/v2/accounts/123/statistics/deliverability", StatisticsAPI},
		{"GET", "/v2/accounts/123/domains", GeneralAPI},
		{"POST", "/v2/accounts/123/domains", GeneralAPI},
		{"GET", "/v2/accounts/123/api-keys", GeneralAPI},
		{"GET", "/v2/accounts/123/sub-accounts", GeneralAPI},
		{"POST", "/v2/accounts/123/sub-accounts", GeneralAPI},
	}

	for _, test := range tests {
		t.Run(test.method+" "+test.path, func(t *testing.T) {
			endpointType := rl.DetectEndpointType(test.method, test.path)
			assert.Equal(t, test.expectedType, endpointType)
		})
	}
}

func TestRateLimitConfiguration(t *testing.T) {
	rl := NewRateLimiter()

	// Test setting custom rate limits
	rl.SetRateLimit(GeneralAPI, 200, 400)
	status := rl.GetStatus(GeneralAPI)
	assert.Equal(t, 200, status.RequestsPerSecond)
	assert.Equal(t, 400, status.BurstCapacity)

	// Test enabling/disabling
	rl.EnableRateLimit(StatisticsAPI, false)
	status = rl.GetStatus(StatisticsAPI)
	assert.False(t, status.Enabled)

	rl.EnableRateLimit(StatisticsAPI, true)
	status = rl.GetStatus(StatisticsAPI)
	assert.True(t, status.Enabled)
}

func TestCustomerConfiguration(t *testing.T) {
	rl := NewRateLimiter()

	config := CustomerRateLimitConfig{
		General: &RateLimitConfig{
			RequestsPerSecond: 150,
			BurstCapacity:     300,
			Enabled:           true,
		},
		SendMessage: &RateLimitConfig{
			RequestsPerSecond: 500,
			BurstCapacity:     1000,
			Enabled:           true,
		},
		Statistics: &RateLimitConfig{
			RequestsPerSecond: 5,
			BurstCapacity:     10,
			Enabled:           true,
		},
	}

	rl.ConfigureFromCustomerConfig(config)

	// Verify all configurations were applied
	generalStatus := rl.GetStatus(GeneralAPI)
	assert.Equal(t, 150, generalStatus.RequestsPerSecond)
	assert.Equal(t, 300, generalStatus.BurstCapacity)

	sendMessageStatus := rl.GetStatus(SendMessageAPI)
	assert.Equal(t, 500, sendMessageStatus.RequestsPerSecond)
	assert.Equal(t, 1000, sendMessageStatus.BurstCapacity)

	statsStatus := rl.GetStatus(StatisticsAPI)
	assert.Equal(t, 5, statsStatus.RequestsPerSecond)
	assert.Equal(t, 10, statsStatus.BurstCapacity)
}

func TestTokenBucketBasicOperation(t *testing.T) {
	config := RateLimitConfig{
		RequestsPerSecond: 2, // 2 requests per second
		BurstCapacity:     3, // Can burst up to 3
		Enabled:           true,
	}

	bucket := NewTokenBucket(config)

	// Should be able to make burst capacity requests immediately
	for i := 0; i < 3; i++ {
		start := time.Now()
		err := bucket.WaitForToken()
		duration := time.Since(start)

		assert.NoError(t, err)
		// Should be essentially immediate (less than 10ms)
		assert.Less(t, duration, 10*time.Millisecond)
	}

	// Next request should block for ~500ms (1/2 second for 2 req/s)
	start := time.Now()
	err := bucket.WaitForToken()
	duration := time.Since(start)

	assert.NoError(t, err)
	// Should wait approximately 500ms, allowing for some variance
	assert.Greater(t, duration, 400*time.Millisecond)
	assert.Less(t, duration, 600*time.Millisecond)
}

func TestAPIClientRateLimitIntegration(t *testing.T) {
	config := NewConfiguration()
	client := NewAPIClientWithConfig(config)

	// Test public API methods
	client.SetGeneralRateLimit(150, 300)
	client.SetStatisticsRateLimit(5, 10)
	client.SetSendMessageRateLimit(500, 1000)

	// Verify settings were applied
	generalStatus := client.GetRateLimitStatus(GeneralAPI)
	assert.Equal(t, 150, generalStatus.RequestsPerSecond)
	assert.Equal(t, 300, generalStatus.BurstCapacity)

	statsStatus := client.GetRateLimitStatus(StatisticsAPI)
	assert.Equal(t, 5, statsStatus.RequestsPerSecond)
	assert.Equal(t, 10, statsStatus.BurstCapacity)

	sendMessageStatus := client.GetRateLimitStatus(SendMessageAPI)
	assert.Equal(t, 500, sendMessageStatus.RequestsPerSecond)
	assert.Equal(t, 1000, sendMessageStatus.BurstCapacity)

	// Test global enable/disable
	client.SetGlobalRateLimit(false)
	client.SetGlobalRateLimit(true)

	// Test endpoint-specific enable/disable
	client.EnableRateLimit(GeneralAPI, false)
	generalStatus = client.GetRateLimitStatus(GeneralAPI)
	assert.False(t, generalStatus.Enabled)

	client.EnableRateLimit(GeneralAPI, true)
	generalStatus = client.GetRateLimitStatus(GeneralAPI)
	assert.True(t, generalStatus.Enabled)
}

func TestConfigurationIntegration(t *testing.T) {
	config := NewConfiguration()

	// Test default rate limiting configuration
	assert.True(t, config.EnableRateLimit)
	assert.Equal(t, 3, config.RetryConfig.MaxRetries)

	require.NotNil(t, config.DefaultGeneralRateLimit)
	assert.Equal(t, 100, config.DefaultGeneralRateLimit.RequestsPerSecond)
	assert.Equal(t, 200, config.DefaultGeneralRateLimit.BurstCapacity)

	require.NotNil(t, config.DefaultStatisticsRateLimit)
	assert.Equal(t, 1, config.DefaultStatisticsRateLimit.RequestsPerSecond)
	assert.Equal(t, 1, config.DefaultStatisticsRateLimit.BurstCapacity)

	require.NotNil(t, config.DefaultSendMessageRateLimit)
	assert.Equal(t, 100, config.DefaultSendMessageRateLimit.RequestsPerSecond)
	assert.Equal(t, 200, config.DefaultSendMessageRateLimit.BurstCapacity)

	// Test creating client with custom configuration
	config.CustomerRateLimits = &CustomerRateLimitConfig{
		General: &RateLimitConfig{
			RequestsPerSecond: 200,
			BurstCapacity:     400,
			Enabled:           true,
		},
	}

	client := NewAPIClientWithConfig(config)

	// Verify custom configuration was applied
	generalStatus := client.GetRateLimitStatus(GeneralAPI)
	assert.Equal(t, 200, generalStatus.RequestsPerSecond)
	assert.Equal(t, 400, generalStatus.BurstCapacity)
}
