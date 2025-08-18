package ahasend

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/AhaSend/ahasend-go/models/responses"
	"github.com/stretchr/testify/assert"
)

func TestDeliveryTime_JSONMarshaling(t *testing.T) {
	// Test with complete delivery time data
	recipientDomain := "example.com"
	deliveryTime := 45.5

	dt := &responses.DeliveryTime{
		RecipientDomain: &recipientDomain,
		DeliveryTime:    &deliveryTime,
	}

	// Test marshaling
	data, err := json.Marshal(dt)
	assert.NoError(t, err)
	assert.Contains(t, string(data), "example.com")
	assert.Contains(t, string(data), "45.5")

	// Test unmarshaling
	var unmarshaled responses.DeliveryTime
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, recipientDomain, *unmarshaled.RecipientDomain)
	assert.Equal(t, deliveryTime, *unmarshaled.DeliveryTime)
}

func TestDeliveryTime_OptionalFields(t *testing.T) {
	// Test with no fields set
	dt := &responses.DeliveryTime{}

	// Test marshaling - optional fields should be omitted
	data, err := json.Marshal(dt)
	assert.NoError(t, err)
	assert.NotContains(t, string(data), "recipient_domain")
	assert.NotContains(t, string(data), "delivery_time")

	// Test unmarshaling
	var unmarshaled responses.DeliveryTime
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Nil(t, unmarshaled.RecipientDomain)
	assert.Nil(t, unmarshaled.DeliveryTime)
}

func TestDeliveryTimeStatistics_JSONMarshaling(t *testing.T) {
	domain1 := "example.com"
	time1 := 30.0
	domain2 := "test.com"
	time2 := 60.0

	deliveryTimes := []responses.DeliveryTime{
		{RecipientDomain: &domain1, DeliveryTime: &time1},
		{RecipientDomain: &domain2, DeliveryTime: &time2},
	}

	stats := &responses.DeliveryTimeStatistics{
		FromTimestamp:   time.Date(2023, 10, 1, 10, 0, 0, 0, time.UTC),
		ToTimestamp:     time.Date(2023, 10, 1, 11, 0, 0, 0, time.UTC),
		AvgDeliveryTime: 45.0,
		DeliveredCount:  100,
		DeliveryTimes:   deliveryTimes,
	}

	// Test marshaling
	data, err := json.Marshal(stats)
	assert.NoError(t, err)
	assert.Contains(t, string(data), "45")
	assert.Contains(t, string(data), "100")
	assert.Contains(t, string(data), "delivery_times")

	// Test unmarshaling
	var unmarshaled responses.DeliveryTimeStatistics
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, stats.AvgDeliveryTime, unmarshaled.AvgDeliveryTime)
	assert.Equal(t, stats.DeliveredCount, unmarshaled.DeliveredCount)
	assert.Len(t, unmarshaled.DeliveryTimes, 2)
	assert.Equal(t, domain1, *unmarshaled.DeliveryTimes[0].RecipientDomain)
	assert.Equal(t, time1, *unmarshaled.DeliveryTimes[0].DeliveryTime)
}

func TestDeliveryTimeStatistics_OptionalDeliveryTimes(t *testing.T) {
	// Test without delivery times
	stats := &responses.DeliveryTimeStatistics{
		FromTimestamp:   time.Date(2023, 10, 1, 10, 0, 0, 0, time.UTC),
		ToTimestamp:     time.Date(2023, 10, 1, 11, 0, 0, 0, time.UTC),
		AvgDeliveryTime: 30.0,
		DeliveredCount:  50,
	}

	// Test marshaling - delivery_times should be omitted
	data, err := json.Marshal(stats)
	assert.NoError(t, err)
	assert.NotContains(t, string(data), "delivery_times")

	// Test unmarshaling
	var unmarshaled responses.DeliveryTimeStatistics
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, stats.AvgDeliveryTime, unmarshaled.AvgDeliveryTime)
	assert.Equal(t, stats.DeliveredCount, unmarshaled.DeliveredCount)
	assert.Nil(t, unmarshaled.DeliveryTimes)
}

func TestDeliveryTimeStatisticsResponse_JSONMarshaling(t *testing.T) {
	stats1 := responses.DeliveryTimeStatistics{
		FromTimestamp:   time.Date(2023, 10, 1, 10, 0, 0, 0, time.UTC),
		ToTimestamp:     time.Date(2023, 10, 1, 11, 0, 0, 0, time.UTC),
		AvgDeliveryTime: 45.0,
		DeliveredCount:  100,
	}

	stats2 := responses.DeliveryTimeStatistics{
		FromTimestamp:   time.Date(2023, 10, 1, 11, 0, 0, 0, time.UTC),
		ToTimestamp:     time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC),
		AvgDeliveryTime: 50.0,
		DeliveredCount:  75,
	}

	response := &responses.DeliveryTimeStatisticsResponse{
		Object: "list",
		Data:   []responses.DeliveryTimeStatistics{stats1, stats2},
	}

	// Test marshaling
	data, err := json.Marshal(response)
	assert.NoError(t, err)
	assert.Contains(t, string(data), "list")
	assert.Contains(t, string(data), "data")

	// Test unmarshaling
	var unmarshaled responses.DeliveryTimeStatisticsResponse
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, response.Object, unmarshaled.Object)
	assert.Len(t, unmarshaled.Data, 2)
	assert.Equal(t, stats1.AvgDeliveryTime, unmarshaled.Data[0].AvgDeliveryTime)
	assert.Equal(t, stats2.AvgDeliveryTime, unmarshaled.Data[1].AvgDeliveryTime)
}
