package api

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
