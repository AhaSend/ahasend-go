package ahasend_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/AhaSend/ahasend-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestStatisticsModelsUpdatedFields tests that the statistics models use the new timestamp fields
func TestStatisticsModelsUpdatedFields(t *testing.T) {
	fromTime := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	toTime := time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC)

	t.Run("DeliverabilityStatistics uses from_timestamp and to_timestamp", func(t *testing.T) {
		stats := ahasend.NewDeliverabilityStatistics(fromTime, toTime)
		stats.SetReceptionCount(100)
		stats.SetDeliveredCount(95)

		// Test getters
		assert.Equal(t, fromTime, stats.GetFromTimestamp())
		assert.Equal(t, toTime, stats.GetToTimestamp())
		assert.Equal(t, int32(100), stats.GetReceptionCount())
		assert.Equal(t, int32(95), stats.GetDeliveredCount())

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
		bounces := []ahasend.Bounce{*ahasend.NewBounce("hard", 10)}
		stats := ahasend.NewBounceStatistics(fromTime, toTime, bounces)

		// Test getters
		assert.Equal(t, fromTime, stats.GetFromTimestamp())
		assert.Equal(t, toTime, stats.GetToTimestamp())
		assert.Equal(t, bounces, stats.GetBounces())
		assert.Equal(t, "hard", stats.GetBounces()[0].GetClassification())
		assert.Equal(t, int32(10), stats.GetBounces()[0].GetCount())

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
		stats := ahasend.NewDeliveryTimeStatistics(fromTime, toTime, 1.25, 50)

		// Test getters
		assert.Equal(t, fromTime, stats.GetFromTimestamp())
		assert.Equal(t, toTime, stats.GetToTimestamp())
		assert.Equal(t, 1.25, stats.GetAvgDeliveryTime())
		assert.Equal(t, int32(50), stats.GetDeliveredCount())

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

		var stats ahasend.DeliverabilityStatistics
		err := json.Unmarshal([]byte(jsonData), &stats)
		require.NoError(t, err)

		assert.Equal(t, int32(100), stats.GetReceptionCount())
		assert.Equal(t, int32(95), stats.GetDeliveredCount())
		assert.Equal(t, int32(2), stats.GetBouncedCount())
		assert.Equal(t, int32(1), stats.GetFailedCount())

		// Test timestamps
		expectedFrom := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
		expectedTo := time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC)
		assert.Equal(t, expectedFrom, stats.GetFromTimestamp())
		assert.Equal(t, expectedTo, stats.GetToTimestamp())
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

		var stats ahasend.BounceStatistics
		err := json.Unmarshal([]byte(jsonData), &stats)
		require.NoError(t, err)

		bounces := stats.GetBounces()
		assert.Len(t, bounces, 2)
		assert.Equal(t, "hard", bounces[0].GetClassification())
		assert.Equal(t, int32(10), bounces[0].GetCount())
		assert.Equal(t, "soft", bounces[1].GetClassification())
		assert.Equal(t, int32(5), bounces[1].GetCount())
	})

	t.Run("DeliveryTimeStatistics JSON deserialization", func(t *testing.T) {
		jsonData := `{
			"from_timestamp": "2024-01-01T12:00:00Z",
			"to_timestamp": "2024-01-01T13:00:00Z",
			"avg_delivery_time": 1.25,
			"delivered_count": 50
		}`

		var stats ahasend.DeliveryTimeStatistics
		err := json.Unmarshal([]byte(jsonData), &stats)
		require.NoError(t, err)

		assert.Equal(t, 1.25, stats.GetAvgDeliveryTime())
		assert.Equal(t, int32(50), stats.GetDeliveredCount())
	})
}

// TestStatisticsModelsRequiredFields tests that the new timestamp fields are properly validated as required
func TestStatisticsModelsRequiredFields(t *testing.T) {
	t.Run("DeliverabilityStatistics requires both timestamps", func(t *testing.T) {
		// Missing from_timestamp
		jsonData := `{
			"to_timestamp": "2024-01-01T13:00:00Z"
		}`

		var stats ahasend.DeliverabilityStatistics
		err := json.Unmarshal([]byte(jsonData), &stats)
		assert.Error(t, err, "Should error when from_timestamp is missing")
		assert.Contains(t, err.Error(), "from_timestamp")

		// Missing to_timestamp
		jsonData = `{
			"from_timestamp": "2024-01-01T12:00:00Z"
		}`

		err = json.Unmarshal([]byte(jsonData), &stats)
		assert.Error(t, err, "Should error when to_timestamp is missing")
		assert.Contains(t, err.Error(), "to_timestamp")
	})
}
