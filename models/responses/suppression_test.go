package responses

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSuppression_JSONMarshaling(t *testing.T) {
	id := uint64(1)
	accountID := uuid.MustParse("01234567-89ab-cdef-0123-456789abcdef")
	createdAt := time.Now().UTC().Truncate(time.Second)
	updatedAt := createdAt.Add(time.Hour)
	expiresAt := createdAt.Add(24 * time.Hour)

	t.Run("minimal suppression without optional fields", func(t *testing.T) {
		suppression := Suppression{
			Object:    "suppression",
			ID:        id,
			AccountID: accountID,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Email:     "user@example.com",
			ExpiresAt: expiresAt,
		}

		// Marshal to JSON
		jsonData, err := json.Marshal(suppression)
		require.NoError(t, err)

		// Should not contain optional fields
		assert.NotContains(t, string(jsonData), "domain")
		assert.NotContains(t, string(jsonData), "reason")

		// Unmarshal and verify
		var decoded Suppression
		err = json.Unmarshal(jsonData, &decoded)
		require.NoError(t, err)

		assert.Equal(t, suppression.Object, decoded.Object)
		assert.Equal(t, suppression.ID, decoded.ID)
		assert.Equal(t, suppression.Email, decoded.Email)
		assert.True(t, suppression.CreatedAt.Equal(decoded.CreatedAt))
		assert.True(t, suppression.UpdatedAt.Equal(decoded.UpdatedAt))
		assert.True(t, suppression.ExpiresAt.Equal(decoded.ExpiresAt))

		// Optional fields should be empty
		assert.Equal(t, "", decoded.Domain)
		assert.Equal(t, "", decoded.Reason)
	})

	t.Run("complete suppression with all optional fields", func(t *testing.T) {
		domain := "example.com"
		reason := "bounce"

		suppression := Suppression{
			Object:    "suppression",
			ID:        id,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Email:     "blocked@example.com",
			ExpiresAt: expiresAt,
			Domain:    domain,
			Reason:    reason,
		}

		// Marshal to JSON
		jsonData, err := json.Marshal(suppression)
		require.NoError(t, err)

		// Should contain optional fields
		assert.Contains(t, string(jsonData), "domain")
		assert.Contains(t, string(jsonData), "reason")

		// Unmarshal and verify
		var decoded Suppression
		err = json.Unmarshal(jsonData, &decoded)
		require.NoError(t, err)

		assert.Equal(t, suppression.Email, decoded.Email)
		assert.True(t, suppression.CreatedAt.Equal(decoded.CreatedAt))
		assert.True(t, suppression.UpdatedAt.Equal(decoded.UpdatedAt))
		assert.True(t, suppression.ExpiresAt.Equal(decoded.ExpiresAt))
		assert.Equal(t, domain, decoded.Domain)
		assert.Equal(t, reason, decoded.Reason)
	})

	t.Run("suppression with omitempty behavior", func(t *testing.T) {
		suppression := Suppression{
			Object:    "suppression",
			ID:        id,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Email:     "user@example.com",
			ExpiresAt: expiresAt,
			// Domain and Reason are nil
		}

		data, err := json.Marshal(suppression)
		require.NoError(t, err)

		// Parse as generic map to check omitempty behavior
		var result map[string]interface{}
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)

		// These should be omitted
		assert.NotContains(t, result, "domain")
		assert.NotContains(t, result, "reason")

		// These should be present
		assert.Contains(t, result, "object")
		assert.Contains(t, result, "id")
		assert.Contains(t, result, "created_at")
		assert.Contains(t, result, "updated_at")
		assert.Contains(t, result, "email")
		assert.Contains(t, result, "expires_at")
	})
}
