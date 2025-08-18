package responses

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebhook_JSONMarshaling(t *testing.T) {
	webhookID := uuid.MustParse("01234567-89ab-cdef-0123-456789abcdef")
	createdAt := time.Now().UTC().Truncate(time.Second)
	updatedAt := createdAt.Add(time.Hour)

	t.Run("minimal webhook without optional fields", func(t *testing.T) {
		webhook := Webhook{
			Object:    "webhook",
			ID:        webhookID,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Name:      "Test Webhook",
			URL:       "https://example.com/webhook",
			Enabled:   true,
		}

		// Marshal to JSON
		jsonData, err := json.Marshal(webhook)
		require.NoError(t, err)

		// Should not contain empty optional fields (only string/slice fields with omitempty work)
		assert.NotContains(t, string(jsonData), "scope")
		assert.NotContains(t, string(jsonData), "domains")
		// Boolean fields with false values will still appear in JSON

		// Unmarshal and verify
		var decoded Webhook
		err = json.Unmarshal(jsonData, &decoded)
		require.NoError(t, err)

		assert.Equal(t, webhook.Object, decoded.Object)
		assert.Equal(t, webhook.ID, decoded.ID)
		assert.Equal(t, webhook.Name, decoded.Name)
		assert.Equal(t, webhook.URL, decoded.URL)
		assert.Equal(t, webhook.Enabled, decoded.Enabled)

		// Optional fields should have default values
		assert.Equal(t, false, decoded.OnReception)
		assert.Equal(t, "", decoded.Scope)
		assert.Nil(t, decoded.Domains)
	})

	t.Run("complete webhook with all optional fields", func(t *testing.T) {
		scope := "domain"
		onReception := true
		onDelivered := false
		onTransientError := true
		onFailed := true
		onBounced := false
		onSuppressed := true
		onOpened := false
		onClicked := true
		onSuppressionCreated := false
		onDNSError := true

		webhook := Webhook{
			Object:               "webhook",
			ID:                   webhookID,
			CreatedAt:            createdAt,
			UpdatedAt:            updatedAt,
			Name:                 "Complete Webhook",
			URL:                  "https://api.example.com/webhooks/handler",
			Enabled:              true,
			OnReception:          onReception,
			OnDelivered:          onDelivered,
			OnTransientError:     onTransientError,
			OnFailed:             onFailed,
			OnBounced:            onBounced,
			OnSuppressed:         onSuppressed,
			OnOpened:             onOpened,
			OnClicked:            onClicked,
			OnSuppressionCreated: onSuppressionCreated,
			OnDNSError:           onDNSError,
			Scope:                scope,
			Domains:              []string{"example.com", "test.com"},
		}

		// Marshal to JSON
		jsonData, err := json.Marshal(webhook)
		require.NoError(t, err)

		// Should contain all optional fields
		assert.Contains(t, string(jsonData), "on_reception")
		assert.Contains(t, string(jsonData), "on_delivered")
		assert.Contains(t, string(jsonData), "scope")
		assert.Contains(t, string(jsonData), "domains")

		// Unmarshal and verify
		var decoded Webhook
		err = json.Unmarshal(jsonData, &decoded)
		require.NoError(t, err)

		assert.Equal(t, webhook.Name, decoded.Name)
		assert.True(t, webhook.CreatedAt.Equal(decoded.CreatedAt))
		assert.True(t, webhook.UpdatedAt.Equal(decoded.UpdatedAt))
		assert.Equal(t, onReception, decoded.OnReception)
		assert.Equal(t, onDelivered, decoded.OnDelivered)
		assert.Equal(t, scope, decoded.Scope)
		assert.Equal(t, webhook.Domains, decoded.Domains)
	})

	t.Run("webhook with omitempty behavior", func(t *testing.T) {
		webhook := Webhook{
			Object:    "webhook",
			ID:        webhookID,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Name:      "Test Webhook",
			URL:       "https://example.com/webhook",
			Enabled:   true,
			// All optional fields are nil/empty
		}

		data, err := json.Marshal(webhook)
		require.NoError(t, err)

		// Parse as generic map to check omitempty behavior
		var result map[string]interface{}
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)

		// These string/slice fields should be omitted when empty
		assert.NotContains(t, result, "scope")
		assert.NotContains(t, result, "domains")
		// Boolean fields will appear even when false

		// These should be present
		assert.Contains(t, result, "object")
		assert.Contains(t, result, "id")
		assert.Contains(t, result, "created_at")
		assert.Contains(t, result, "updated_at")
		assert.Contains(t, result, "name")
		assert.Contains(t, result, "url")
		assert.Contains(t, result, "enabled")
	})
}
