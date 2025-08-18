package ahasend

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/AhaSend/ahasend-go/models/responses"
	"github.com/stretchr/testify/assert"
)

func TestBounce_JSONMarshaling(t *testing.T) {
	// Test with complete bounce data
	bounce := &responses.Bounce{
		Classification: "hard_bounce",
		Count:          15,
	}

	// Test marshaling
	data, err := json.Marshal(bounce)
	assert.NoError(t, err)
	assert.Contains(t, string(data), "hard_bounce")
	assert.Contains(t, string(data), "15")

	// Test unmarshaling
	var unmarshaled responses.Bounce
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, bounce.Classification, unmarshaled.Classification)
	assert.Equal(t, bounce.Count, unmarshaled.Count)
}

func TestBounceStatistics_JSONMarshaling(t *testing.T) {
	// Test with complete bounce statistics
	bounces := []responses.Bounce{
		{Classification: "hard_bounce", Count: 15},
		{Classification: "soft_bounce", Count: 25},
		{Classification: "spam", Count: 3},
	}

	stats := &responses.BounceStatistics{
		FromTimestamp: time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
		ToTimestamp:   time.Date(2023, 10, 1, 1, 0, 0, 0, time.UTC),
		Bounces:       bounces,
	}

	// Test marshaling
	data, err := json.Marshal(stats)
	assert.NoError(t, err)
	assert.Contains(t, string(data), "bounces")
	assert.Contains(t, string(data), "hard_bounce")

	// Test unmarshaling
	var unmarshaled responses.BounceStatistics
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, stats.FromTimestamp, unmarshaled.FromTimestamp)
	assert.Equal(t, stats.ToTimestamp, unmarshaled.ToTimestamp)
	assert.Len(t, unmarshaled.Bounces, 3)
	assert.Equal(t, stats.Bounces[0].Classification, unmarshaled.Bounces[0].Classification)
	assert.Equal(t, stats.Bounces[0].Count, unmarshaled.Bounces[0].Count)
}

func TestBounceStatisticsResponse_JSONMarshaling(t *testing.T) {
	stats1 := responses.BounceStatistics{
		FromTimestamp: time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
		ToTimestamp:   time.Date(2023, 10, 1, 1, 0, 0, 0, time.UTC),
		Bounces: []responses.Bounce{
			{Classification: "hard_bounce", Count: 15},
			{Classification: "soft_bounce", Count: 25},
		},
	}

	stats2 := responses.BounceStatistics{
		FromTimestamp: time.Date(2023, 10, 1, 1, 0, 0, 0, time.UTC),
		ToTimestamp:   time.Date(2023, 10, 1, 2, 0, 0, 0, time.UTC),
		Bounces: []responses.Bounce{
			{Classification: "spam", Count: 3},
		},
	}

	response := &responses.BounceStatisticsResponse{
		Object: "list",
		Data:   []responses.BounceStatistics{stats1, stats2},
	}

	// Test marshaling
	data, err := json.Marshal(response)
	assert.NoError(t, err)
	assert.Contains(t, string(data), "list")
	assert.Contains(t, string(data), "data")

	// Test unmarshaling
	var unmarshaled responses.BounceStatisticsResponse
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, response.Object, unmarshaled.Object)
	assert.Len(t, unmarshaled.Data, 2)
	assert.Len(t, unmarshaled.Data[0].Bounces, 2)
	assert.Len(t, unmarshaled.Data[1].Bounces, 1)
}
