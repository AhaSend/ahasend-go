package ahasend_test

import (
	"testing"
	"time"

	"github.com/AhaSend/ahasend-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRateLimiterCreation(t *testing.T) {
	rl := ahasend.NewRateLimiter()
	require.NotNil(t, rl)

	// Test default configuration
	assert.True(t, rl.IsEnabled())

	// Test status for different endpoint types
	generalStatus := rl.GetStatus(ahasend.GeneralAPI)
	assert.Equal(t, ahasend.GeneralAPI, generalStatus.EndpointType)
	assert.True(t, generalStatus.Enabled)
	assert.Equal(t, 100, generalStatus.RequestsPerSecond)
	assert.Equal(t, 200, generalStatus.BurstCapacity)

	statsStatus := rl.GetStatus(ahasend.StatisticsAPI)
	assert.Equal(t, ahasend.StatisticsAPI, statsStatus.EndpointType)
	assert.True(t, statsStatus.Enabled)
	assert.Equal(t, 1, statsStatus.RequestsPerSecond)
	assert.Equal(t, 1, statsStatus.BurstCapacity)

	sendMessageStatus := rl.GetStatus(ahasend.SendMessageAPI)
	assert.Equal(t, ahasend.SendMessageAPI, sendMessageStatus.EndpointType)
	assert.True(t, sendMessageStatus.Enabled)
	assert.Equal(t, 100, sendMessageStatus.RequestsPerSecond)
	assert.Equal(t, 200, sendMessageStatus.BurstCapacity)
}

func TestEndpointDetection(t *testing.T) {
	rl := ahasend.NewRateLimiter()

	tests := []struct {
		method       string
		path         string
		expectedType ahasend.EndpointType
	}{
		{"POST", "/v2/accounts/123/messages", ahasend.SendMessageAPI},
		{"GET", "/v2/accounts/123/messages/456", ahasend.GeneralAPI}, // Specific message, not send
		{"GET", "/v2/accounts/123/statistics/bounce", ahasend.StatisticsAPI},
		{"GET", "/v2/accounts/123/statistics/deliverability", ahasend.StatisticsAPI},
		{"GET", "/v2/accounts/123/domains", ahasend.GeneralAPI},
		{"POST", "/v2/accounts/123/domains", ahasend.GeneralAPI},
		{"GET", "/v2/accounts/123/api-keys", ahasend.GeneralAPI},
	}

	for _, test := range tests {
		t.Run(test.method+" "+test.path, func(t *testing.T) {
			endpointType := rl.DetectEndpointType(test.method, test.path)
			assert.Equal(t, test.expectedType, endpointType)
		})
	}
}

func TestRateLimitConfiguration(t *testing.T) {
	rl := ahasend.NewRateLimiter()

	// Test setting custom rate limits
	rl.SetRateLimit(ahasend.GeneralAPI, 200, 400)
	status := rl.GetStatus(ahasend.GeneralAPI)
	assert.Equal(t, 200, status.RequestsPerSecond)
	assert.Equal(t, 400, status.BurstCapacity)

	// Test enabling/disabling
	rl.EnableRateLimit(ahasend.StatisticsAPI, false)
	status = rl.GetStatus(ahasend.StatisticsAPI)
	assert.False(t, status.Enabled)

	rl.EnableRateLimit(ahasend.StatisticsAPI, true)
	status = rl.GetStatus(ahasend.StatisticsAPI)
	assert.True(t, status.Enabled)
}

func TestCustomerConfiguration(t *testing.T) {
	rl := ahasend.NewRateLimiter()

	config := ahasend.CustomerRateLimitConfig{
		General: &ahasend.RateLimitConfig{
			RequestsPerSecond: 150,
			BurstCapacity:     300,
			Enabled:           true,
		},
		SendMessage: &ahasend.RateLimitConfig{
			RequestsPerSecond: 500,
			BurstCapacity:     1000,
			Enabled:           true,
		},
		Statistics: &ahasend.RateLimitConfig{
			RequestsPerSecond: 5,
			BurstCapacity:     10,
			Enabled:           true,
		},
	}

	rl.ConfigureFromCustomerConfig(config)

	// Verify all configurations were applied
	generalStatus := rl.GetStatus(ahasend.GeneralAPI)
	assert.Equal(t, 150, generalStatus.RequestsPerSecond)
	assert.Equal(t, 300, generalStatus.BurstCapacity)

	sendMessageStatus := rl.GetStatus(ahasend.SendMessageAPI)
	assert.Equal(t, 500, sendMessageStatus.RequestsPerSecond)
	assert.Equal(t, 1000, sendMessageStatus.BurstCapacity)

	statsStatus := rl.GetStatus(ahasend.StatisticsAPI)
	assert.Equal(t, 5, statsStatus.RequestsPerSecond)
	assert.Equal(t, 10, statsStatus.BurstCapacity)
}

func TestTokenBucketBasicOperation(t *testing.T) {
	config := ahasend.RateLimitConfig{
		RequestsPerSecond: 2, // 2 requests per second
		BurstCapacity:     3, // Can burst up to 3
		Enabled:           true,
	}

	bucket := ahasend.NewTokenBucket(config)

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
	config := ahasend.NewConfiguration()
	client := ahasend.NewAPIClient(config)

	// Test public API methods
	client.SetGeneralRateLimit(150, 300)
	client.SetStatisticsRateLimit(5, 10)
	client.SetSendMessageRateLimit(500, 1000)

	// Verify settings were applied
	generalStatus := client.GetRateLimitStatus(ahasend.GeneralAPI)
	assert.Equal(t, 150, generalStatus.RequestsPerSecond)
	assert.Equal(t, 300, generalStatus.BurstCapacity)

	statsStatus := client.GetRateLimitStatus(ahasend.StatisticsAPI)
	assert.Equal(t, 5, statsStatus.RequestsPerSecond)
	assert.Equal(t, 10, statsStatus.BurstCapacity)

	sendMessageStatus := client.GetRateLimitStatus(ahasend.SendMessageAPI)
	assert.Equal(t, 500, sendMessageStatus.RequestsPerSecond)
	assert.Equal(t, 1000, sendMessageStatus.BurstCapacity)

	// Test global enable/disable
	client.SetGlobalRateLimit(false)
	client.SetGlobalRateLimit(true)

	// Test endpoint-specific enable/disable
	client.EnableRateLimit(ahasend.GeneralAPI, false)
	generalStatus = client.GetRateLimitStatus(ahasend.GeneralAPI)
	assert.False(t, generalStatus.Enabled)

	client.EnableRateLimit(ahasend.GeneralAPI, true)
	generalStatus = client.GetRateLimitStatus(ahasend.GeneralAPI)
	assert.True(t, generalStatus.Enabled)
}

func TestConfigurationIntegration(t *testing.T) {
	config := ahasend.NewConfiguration()

	// Test default rate limiting configuration
	assert.True(t, config.EnableRateLimit)
	assert.Equal(t, 3, config.MaxRetries)

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
	config.CustomerRateLimits = &ahasend.CustomerRateLimitConfig{
		General: &ahasend.RateLimitConfig{
			RequestsPerSecond: 200,
			BurstCapacity:     400,
			Enabled:           true,
		},
	}

	client := ahasend.NewAPIClient(config)

	// Verify custom configuration was applied
	generalStatus := client.GetRateLimitStatus(ahasend.GeneralAPI)
	assert.Equal(t, 200, generalStatus.RequestsPerSecond)
	assert.Equal(t, 400, generalStatus.BurstCapacity)
}
