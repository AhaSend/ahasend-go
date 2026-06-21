package responses

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPIKey_JSONMarshaling(t *testing.T) {
	apiKeyID := uuid.MustParse("01234567-89ab-cdef-0123-456789abcdef")
	accountID := uuid.MustParse("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee")
	scopeID1 := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	scopeID2 := uuid.MustParse("22222222-2222-2222-2222-222222222222")
	domainID := uuid.MustParse("dddddddd-dddd-dddd-dddd-dddddddddddd")
	createdAt := time.Now().UTC().Truncate(time.Second)
	updatedAt := createdAt.Add(time.Hour)
	lastUsedAt := updatedAt.Add(time.Minute * 30)

	t.Run("minimal API key without optional fields", func(t *testing.T) {
		apiKey := APIKey{
			Object:      "api_key",
			ID:          apiKeyID,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			AccountID:   accountID,
			Label:       "Test Key",
			PublicKey:   "aha-pk-test123",
			IPAllowList: []string{},
			Scopes: []APIKeyScope{
				{
					ID:        scopeID1,
					CreatedAt: createdAt,
					UpdatedAt: updatedAt,
					APIKeyID:  apiKeyID,
					Scope:     "messages:read:all",
				},
			},
		}

		// Marshal to JSON
		jsonData, err := json.Marshal(apiKey)
		require.NoError(t, err)

		// Should not contain optional fields
		assert.NotContains(t, string(jsonData), "last_used_at")
		assert.NotContains(t, string(jsonData), "secret_key")

		// ip_allow_list is required/always-present, even when empty
		assert.Contains(t, string(jsonData), `"ip_allow_list":[]`)

		// Unmarshal and verify
		var decoded APIKey
		err = json.Unmarshal(jsonData, &decoded)
		require.NoError(t, err)

		assert.Equal(t, apiKey.Object, decoded.Object)
		assert.Equal(t, apiKey.ID, decoded.ID)
		assert.Equal(t, apiKey.Label, decoded.Label)
		assert.Equal(t, apiKey.PublicKey, decoded.PublicKey)
		assert.Len(t, decoded.Scopes, 1)
		assert.Equal(t, apiKey.Scopes[0].Scope, decoded.Scopes[0].Scope)
		assert.Empty(t, decoded.IPAllowList)

		// Optional fields should be nil
		assert.Nil(t, decoded.LastUsedAt)
		assert.Nil(t, decoded.SecretKey)
	})

	t.Run("complete API key with all optional fields", func(t *testing.T) {
		secretKey := "aha-sk-secret123"

		apiKey := APIKey{
			Object:      "api_key",
			ID:          apiKeyID,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			AccountID:   accountID,
			Label:       "Complete Key",
			PublicKey:   "aha-pk-complete123",
			IPAllowList: []string{"203.0.113.0/24", "198.51.100.7"},
			LastUsedAt:  &lastUsedAt,
			SecretKey:   &secretKey,
			Scopes: []APIKeyScope{
				{
					ID:        scopeID1,
					CreatedAt: createdAt,
					UpdatedAt: updatedAt,
					APIKeyID:  apiKeyID,
					Scope:     "messages:send:all",
				},
				{
					ID:        scopeID2,
					CreatedAt: createdAt,
					UpdatedAt: updatedAt,
					APIKeyID:  apiKeyID,
					Scope:     "routes:read:{example.com}",
					DomainID:  &domainID,
				},
			},
		}

		// Marshal to JSON
		jsonData, err := json.Marshal(apiKey)
		require.NoError(t, err)

		// Should contain all optional fields
		assert.Contains(t, string(jsonData), "last_used_at")
		assert.Contains(t, string(jsonData), "secret_key")

		// Unmarshal and verify
		var decoded APIKey
		err = json.Unmarshal(jsonData, &decoded)
		require.NoError(t, err)

		assert.Equal(t, apiKey.Label, decoded.Label)
		assert.True(t, apiKey.CreatedAt.Equal(decoded.CreatedAt))
		assert.True(t, lastUsedAt.Equal(*decoded.LastUsedAt))
		assert.Equal(t, secretKey, *decoded.SecretKey)
		assert.Len(t, decoded.Scopes, 2)
		assert.Equal(t, domainID, *decoded.Scopes[1].DomainID)
		assert.Equal(t, []string{"203.0.113.0/24", "198.51.100.7"}, decoded.IPAllowList)
	})
}

// TestAPIKey_DecodeServerResponse decodes raw server-shaped JSON (not SDK
// round-tripped) to pin down how ip_allow_list arrives over the wire: the
// server always sends it as an array (populated or empty), but we also document
// the defensive null/absent cases.
func TestAPIKey_DecodeServerResponse(t *testing.T) {
	const base = `"object":"api_key","id":"01234567-89ab-cdef-0123-456789abcdef",` +
		`"created_at":"2024-01-01T00:00:00Z","updated_at":"2024-01-01T00:00:00Z",` +
		`"account_id":"aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee","label":"Server Key",` +
		`"public_key":"aha-pk-server","scopes":[]`

	t.Run("populated ip_allow_list", func(t *testing.T) {
		var key APIKey
		err := json.Unmarshal([]byte(`{`+base+`,"ip_allow_list":["203.0.113.0/24","198.51.100.7"]}`), &key)
		require.NoError(t, err)
		assert.Equal(t, []string{"203.0.113.0/24", "198.51.100.7"}, key.IPAllowList)
	})

	t.Run("empty ip_allow_list means allow-all", func(t *testing.T) {
		var key APIKey
		err := json.Unmarshal([]byte(`{`+base+`,"ip_allow_list":[]}`), &key)
		require.NoError(t, err)
		assert.NotNil(t, key.IPAllowList)
		assert.Empty(t, key.IPAllowList)
	})

	t.Run("null ip_allow_list decodes to nil", func(t *testing.T) {
		var key APIKey
		err := json.Unmarshal([]byte(`{`+base+`,"ip_allow_list":null}`), &key)
		require.NoError(t, err)
		assert.Nil(t, key.IPAllowList)
	})

	t.Run("absent ip_allow_list decodes to nil", func(t *testing.T) {
		var key APIKey
		err := json.Unmarshal([]byte(`{`+base+`}`), &key)
		require.NoError(t, err)
		assert.Nil(t, key.IPAllowList)
	})
}

func TestAPIKeyScope_JSONMarshaling(t *testing.T) {
	scopeID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	apiKeyID := uuid.MustParse("01234567-89ab-cdef-0123-456789abcdef")
	domainID := uuid.MustParse("dddddddd-dddd-dddd-dddd-dddddddddddd")
	createdAt := time.Now().UTC().Truncate(time.Second)
	updatedAt := createdAt.Add(time.Hour)

	t.Run("scope without domain ID", func(t *testing.T) {
		scope := APIKeyScope{
			ID:        scopeID,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			APIKeyID:  apiKeyID,
			Scope:     "messages:send:all",
		}

		// Marshal to JSON
		jsonData, err := json.Marshal(scope)
		require.NoError(t, err)

		// Should not contain domain_id
		assert.NotContains(t, string(jsonData), "domain_id")

		// Unmarshal and verify
		var decoded APIKeyScope
		err = json.Unmarshal(jsonData, &decoded)
		require.NoError(t, err)

		assert.Equal(t, scope.ID, decoded.ID)
		assert.Equal(t, scope.Scope, decoded.Scope)
		assert.Nil(t, decoded.DomainID)
	})

	t.Run("scope with domain ID", func(t *testing.T) {
		scope := APIKeyScope{
			ID:        scopeID,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			APIKeyID:  apiKeyID,
			Scope:     "routes:read:{example.com}",
			DomainID:  &domainID,
		}

		// Marshal to JSON
		jsonData, err := json.Marshal(scope)
		require.NoError(t, err)

		// Should contain domain_id
		assert.Contains(t, string(jsonData), "domain_id")

		// Unmarshal and verify
		var decoded APIKeyScope
		err = json.Unmarshal(jsonData, &decoded)
		require.NoError(t, err)

		assert.Equal(t, scope.ID, decoded.ID)
		assert.Equal(t, scope.Scope, decoded.Scope)
		assert.Equal(t, domainID, *decoded.DomainID)
	})
}
