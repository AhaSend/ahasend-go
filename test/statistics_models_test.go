package ahasend_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/AhaSend/ahasend-go/models/responses"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestStatisticsModelsUpdatedFields tests that the statistics models use the new timestamp fields
func TestStatisticsModelsUpdatedFields(t *testing.T) {
	fromTime := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	toTime := time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC)

	t.Run("DeliverabilityStatistics uses from_timestamp and to_timestamp", func(t *testing.T) {
		receptionCount := 100
		deliveredCount := 95
		stats := responses.DeliverabilityStatistics{
			FromTimestamp:  fromTime,
			ToTimestamp:    toTime,
			ReceptionCount: receptionCount,
			DeliveredCount: deliveredCount,
		}

		// Test direct field access
		assert.Equal(t, fromTime, stats.FromTimestamp)
		assert.Equal(t, toTime, stats.ToTimestamp)
		assert.Equal(t, 100, stats.ReceptionCount)
		assert.Equal(t, 95, stats.DeliveredCount)

		// Test JSON serialization contains correct fields
		jsonBytes, err := json.Marshal(stats)
		require.NoError(t, err)

		var jsonMap map[string]interface{}
		err = json.Unmarshal(jsonBytes, &jsonMap)
		require.NoError(t, err)

		assert.Contains(t, jsonMap, "from_timestamp")
		assert.Contains(t, jsonMap, "to_timestamp")
		assert.Contains(t, jsonMap, "reception_count")
		assert.Contains(t, jsonMap, "delivered_count")
		assert.NotContains(t, jsonMap, "time_bucket", "Should not contain old time_bucket field")
	})

	t.Run("BounceStatistics uses from_timestamp and to_timestamp", func(t *testing.T) {
		bounces := []responses.Bounce{{Classification: "hard", Count: 10}}
		stats := responses.BounceStatistics{
			FromTimestamp: fromTime,
			ToTimestamp:   toTime,
			Bounces:       bounces,
		}

		// Test direct field access
		assert.Equal(t, fromTime, stats.FromTimestamp)
		assert.Equal(t, toTime, stats.ToTimestamp)
		assert.Equal(t, bounces, stats.Bounces)
		assert.Equal(t, "hard", stats.Bounces[0].Classification)
		assert.Equal(t, 10, stats.Bounces[0].Count)

		// Test JSON serialization contains correct fields
		jsonBytes, err := json.Marshal(stats)
		require.NoError(t, err)

		var jsonMap map[string]interface{}
		err = json.Unmarshal(jsonBytes, &jsonMap)
		require.NoError(t, err)

		assert.Contains(t, jsonMap, "from_timestamp")
		assert.Contains(t, jsonMap, "to_timestamp")
		assert.Contains(t, jsonMap, "bounces")
		assert.NotContains(t, jsonMap, "time_bucket", "Should not contain old time_bucket field")
	})

	t.Run("DeliveryTimeStatistics uses from_timestamp and to_timestamp", func(t *testing.T) {
		stats := responses.DeliveryTimeStatistics{
			FromTimestamp:   fromTime,
			ToTimestamp:     toTime,
			AvgDeliveryTime: 1.25,
			DeliveredCount:  50,
		}

		// Test direct field access
		assert.Equal(t, fromTime, stats.FromTimestamp)
		assert.Equal(t, toTime, stats.ToTimestamp)
		assert.Equal(t, 1.25, stats.AvgDeliveryTime)
		assert.Equal(t, 50, stats.DeliveredCount)

		// Test JSON serialization contains correct fields
		jsonBytes, err := json.Marshal(stats)
		require.NoError(t, err)

		var jsonMap map[string]interface{}
		err = json.Unmarshal(jsonBytes, &jsonMap)
		require.NoError(t, err)

		assert.Contains(t, jsonMap, "from_timestamp")
		assert.Contains(t, jsonMap, "to_timestamp")
		assert.NotContains(t, jsonMap, "time_bucket", "Should not contain old time_bucket field")
	})
}

// TestStatisticsModelsJSONDeserialization tests that the models can deserialize the new JSON format
func TestStatisticsModelsJSONDeserialization(t *testing.T) {
	t.Run("DeliverabilityStatistics JSON deserialization", func(t *testing.T) {
		jsonData := `{
			"from_timestamp": "2024-01-01T12:00:00Z",
			"to_timestamp": "2024-01-01T13:00:00Z",
			"reception_count": 100,
			"delivered_count": 95,
			"bounced_count": 2,
			"failed_count": 1
		}`

		var stats responses.DeliverabilityStatistics
		err := json.Unmarshal([]byte(jsonData), &stats)
		require.NoError(t, err)

		assert.Equal(t, 100, stats.ReceptionCount)
		assert.Equal(t, 95, stats.DeliveredCount)
		assert.Equal(t, 2, stats.BouncedCount)
		assert.Equal(t, 1, stats.FailedCount)

		// Test timestamps
		expectedFrom := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
		expectedTo := time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC)
		assert.Equal(t, expectedFrom, stats.FromTimestamp)
		assert.Equal(t, expectedTo, stats.ToTimestamp)
	})

	t.Run("BounceStatistics JSON deserialization", func(t *testing.T) {
		jsonData := `{
			"from_timestamp": "2024-01-01T12:00:00Z",
			"to_timestamp": "2024-01-01T13:00:00Z",
			"bounces": [
				{"classification": "hard", "count": 10},
				{"classification": "soft", "count": 5}
			]
		}`

		var stats responses.BounceStatistics
		err := json.Unmarshal([]byte(jsonData), &stats)
		require.NoError(t, err)

		bounces := stats.Bounces
		assert.Len(t, bounces, 2)
		assert.Equal(t, "hard", bounces[0].Classification)
		assert.Equal(t, 10, bounces[0].Count)
		assert.Equal(t, "soft", bounces[1].Classification)
		assert.Equal(t, 5, bounces[1].Count)
	})

	t.Run("DeliveryTimeStatistics JSON deserialization", func(t *testing.T) {
		jsonData := `{
			"from_timestamp": "2024-01-01T12:00:00Z",
			"to_timestamp": "2024-01-01T13:00:00Z",
			"avg_delivery_time": 1.25,
			"delivered_count": 50
		}`

		var stats responses.DeliveryTimeStatistics
		err := json.Unmarshal([]byte(jsonData), &stats)
		require.NoError(t, err)

		assert.Equal(t, 1.25, stats.AvgDeliveryTime)
		assert.Equal(t, 50, stats.DeliveredCount)
	})
}

// TestStatisticsModelsRequiredFields tests that the new timestamp fields are properly handled
func TestStatisticsModelsRequiredFields(t *testing.T) {
	t.Run("DeliverabilityStatistics handles missing timestamps gracefully", func(t *testing.T) {
		// With only to_timestamp, from_timestamp should be zero value
		jsonData := `{
			"to_timestamp": "2024-01-01T13:00:00Z",
			"reception_count": 100
		}`

		var stats responses.DeliverabilityStatistics
		err := json.Unmarshal([]byte(jsonData), &stats)
		require.NoError(t, err)

		// from_timestamp should be zero value
		assert.True(t, stats.FromTimestamp.IsZero())
		assert.False(t, stats.ToTimestamp.IsZero())
		assert.Equal(t, 100, stats.ReceptionCount)
	})
}
