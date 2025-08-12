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
		stats.SetSent(100)
		stats.SetDelivered(95)

		// Test getters
		assert.Equal(t, fromTime, stats.GetFromTimestamp())
		assert.Equal(t, toTime, stats.GetToTimestamp())

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

	t.Run("BounceStatistics uses from_timestamp and to_timestamp", func(t *testing.T) {
		stats := ahasend.NewBounceStatistics(fromTime, toTime, "hard", 10)

		// Test getters
		assert.Equal(t, fromTime, stats.GetFromTimestamp())
		assert.Equal(t, toTime, stats.GetToTimestamp())
		assert.Equal(t, "hard", stats.GetClassification())
		assert.Equal(t, int32(10), stats.GetCount())

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
			"sent": 100,
			"delivered": 95
		}`

		var stats ahasend.DeliverabilityStatistics
		err := json.Unmarshal([]byte(jsonData), &stats)
		require.NoError(t, err)

		assert.Equal(t, int32(100), stats.GetSent())
		assert.Equal(t, int32(95), stats.GetDelivered())

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
			"classification": "hard",
			"count": 10
		}`

		var stats ahasend.BounceStatistics
		err := json.Unmarshal([]byte(jsonData), &stats)
		require.NoError(t, err)

		assert.Equal(t, "hard", stats.GetClassification())
		assert.Equal(t, int32(10), stats.GetCount())
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
