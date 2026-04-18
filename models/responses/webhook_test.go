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
			Scope:     "global",
			Domains:   []string{},
		}

		// Marshal to JSON
		jsonData, err := json.Marshal(webhook)
		require.NoError(t, err)

		assert.Contains(t, string(jsonData), "scope")
		assert.Contains(t, string(jsonData), "domains")

		// Unmarshal and verify
		var decoded Webhook
		err = json.Unmarshal(jsonData, &decoded)
		require.NoError(t, err)

		assert.Equal(t, webhook.Object, decoded.Object)
		assert.Equal(t, webhook.ID, decoded.ID)
		assert.Equal(t, webhook.Name, decoded.Name)
		assert.Equal(t, webhook.URL, decoded.URL)
		assert.Equal(t, webhook.Enabled, decoded.Enabled)

		assert.Equal(t, false, decoded.OnReception)
		assert.Equal(t, "global", decoded.Scope)
		assert.Equal(t, []string{}, decoded.Domains)
	})

	t.Run("complete webhook with all optional fields", func(t *testing.T) {
		scope := "scoped"
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
			Scope:     "global",
			Domains:   []string{},
		}

		data, err := json.Marshal(webhook)
		require.NoError(t, err)

		// Parse as generic map to check omitempty behavior
		var result map[string]interface{}
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)

		assert.Contains(t, result, "scope")
		assert.Contains(t, result, "domains")

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
