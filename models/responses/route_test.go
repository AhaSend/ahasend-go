package responses

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRoute_JSONMarshaling(t *testing.T) {
	routeID := uuid.MustParse("01234567-89ab-cdef-0123-456789abcdef")
	createdAt := time.Now().UTC().Truncate(time.Second)
	updatedAt := createdAt.Add(time.Hour)

	t.Run("minimal route without optional fields", func(t *testing.T) {
		route := Route{
			Object:           "route",
			ID:               routeID,
			CreatedAt:        createdAt,
			UpdatedAt:        updatedAt,
			Name:             "Test Route",
			URL:              "https://example.com/webhook",
			Attachments:      false,
			Headers:          false,
			GroupByMessageID: false,
			StripReplies:     false,
			Enabled:          true,
		}

		// Marshal to JSON
		jsonData, err := json.Marshal(route)
		require.NoError(t, err)

		// Should not contain recipient (omitempty)
		assert.NotContains(t, string(jsonData), "recipient")

		// Unmarshal and verify
		var decoded Route
		err = json.Unmarshal(jsonData, &decoded)
		require.NoError(t, err)

		assert.Equal(t, route.Object, decoded.Object)
		assert.Equal(t, route.ID, decoded.ID)
		assert.Equal(t, route.Name, decoded.Name)
		assert.Equal(t, route.URL, decoded.URL)
		assert.Equal(t, route.Enabled, decoded.Enabled)
		assert.Equal(t, route.Attachments, decoded.Attachments)
		assert.Equal(t, route.Headers, decoded.Headers)
		assert.Equal(t, route.GroupByMessageID, decoded.GroupByMessageID)
		assert.Equal(t, route.StripReplies, decoded.StripReplies)

		// Optional fields should be empty
		assert.Equal(t, "", decoded.Recipient)
	})

	t.Run("complete route with all optional fields", func(t *testing.T) {
		recipient := "user@example.com"

		route := Route{
			Object:           "route",
			ID:               routeID,
			CreatedAt:        createdAt,
			UpdatedAt:        updatedAt,
			Name:             "Complete Route",
			URL:              "https://api.example.com/routes/handler",
			Recipient:        recipient,
			Attachments:      true,
			Headers:          false,
			GroupByMessageID: true,
			StripReplies:     false,
			Enabled:          true,
		}

		// Marshal to JSON
		jsonData, err := json.Marshal(route)
		require.NoError(t, err)

		// Should contain all fields
		assert.Contains(t, string(jsonData), "recipient")
		assert.Contains(t, string(jsonData), "attachments")
		assert.Contains(t, string(jsonData), "headers")
		assert.Contains(t, string(jsonData), "group_by_message_id")
		assert.Contains(t, string(jsonData), "strip_replies")

		// Unmarshal and verify
		var decoded Route
		err = json.Unmarshal(jsonData, &decoded)
		require.NoError(t, err)

		assert.Equal(t, route.Name, decoded.Name)
		assert.True(t, route.CreatedAt.Equal(decoded.CreatedAt))
		assert.True(t, route.UpdatedAt.Equal(decoded.UpdatedAt))
		assert.Equal(t, recipient, decoded.Recipient)
		assert.Equal(t, true, decoded.Attachments)
		assert.Equal(t, false, decoded.Headers)
		assert.Equal(t, true, decoded.GroupByMessageID)
		assert.Equal(t, false, decoded.StripReplies)
	})

	t.Run("route with omitempty behavior", func(t *testing.T) {
		route := Route{
			Object:           "route",
			ID:               routeID,
			CreatedAt:        createdAt,
			UpdatedAt:        updatedAt,
			Name:             "Test Route",
			URL:              "https://example.com/webhook",
			Attachments:      false,
			Headers:          false,
			GroupByMessageID: false,
			StripReplies:     false,
			Enabled:          true,
			// Recipient is nil
		}

		data, err := json.Marshal(route)
		require.NoError(t, err)

		// Parse as generic map to check omitempty behavior
		var result map[string]interface{}
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)

		// This should be omitted
		assert.NotContains(t, result, "recipient")

		// These should be present (boolean fields are not omitempty)
		assert.Contains(t, result, "object")
		assert.Contains(t, result, "name")
		assert.Contains(t, result, "url")
		assert.Contains(t, result, "enabled")
		assert.Contains(t, result, "attachments")
		assert.Contains(t, result, "headers")
		assert.Contains(t, result, "group_by_message_id")
		assert.Contains(t, result, "strip_replies")
	})
}
