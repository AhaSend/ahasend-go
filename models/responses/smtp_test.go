package responses

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSMTPCredential_JSONMarshaling(t *testing.T) {
	credentialID := uint64(1)
	createdAt := time.Now().UTC().Truncate(time.Second)
	updatedAt := createdAt.Add(time.Hour)

	t.Run("minimal SMTP credential without optional fields", func(t *testing.T) {
		credential := SMTPCredential{
			Object:    "smtp_credential",
			ID:        credentialID,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Name:      "Test Credential",
			Username:  "test_user",
			Sandbox:   false,
			Scope:     "global",
			Domains:   []string{}, // Empty array, not nil
		}

		// Marshal to JSON
		jsonData, err := json.Marshal(credential)
		require.NoError(t, err)

		// Should contain domains field (even if empty)
		assert.Contains(t, string(jsonData), "domains")

		// Unmarshal and verify
		var decoded SMTPCredential
		err = json.Unmarshal(jsonData, &decoded)
		require.NoError(t, err)

		assert.Equal(t, credential.Object, decoded.Object)
		assert.Equal(t, credential.ID, decoded.ID)
		assert.Equal(t, credential.Name, decoded.Name)
		assert.Equal(t, credential.Username, decoded.Username)
		assert.Equal(t, credential.Sandbox, decoded.Sandbox)
		assert.Equal(t, credential.Scope, decoded.Scope)
		assert.Equal(t, []string{}, decoded.Domains)
	})

	t.Run("complete SMTP credential with domains", func(t *testing.T) {
		domains := []string{"example.com", "test.com"}

		credential := SMTPCredential{
			Object:    "smtp_credential",
			ID:        credentialID,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Name:      "Scoped Credential",
			Username:  "scoped_user",
			Sandbox:   true,
			Scope:     "scoped",
			Domains:   domains,
		}

		// Marshal to JSON
		jsonData, err := json.Marshal(credential)
		require.NoError(t, err)

		// Should contain domains field
		assert.Contains(t, string(jsonData), "domains")

		// Unmarshal and verify
		var decoded SMTPCredential
		err = json.Unmarshal(jsonData, &decoded)
		require.NoError(t, err)

		assert.Equal(t, credential.Name, decoded.Name)
		assert.True(t, credential.CreatedAt.Equal(decoded.CreatedAt))
		assert.True(t, credential.UpdatedAt.Equal(decoded.UpdatedAt))
		assert.Equal(t, credential.Sandbox, decoded.Sandbox)
		assert.Equal(t, credential.Scope, decoded.Scope)
		assert.Equal(t, domains, decoded.Domains)
	})

	t.Run("SMTP credential with all fields", func(t *testing.T) {
		credential := SMTPCredential{
			Object:    "smtp_credential",
			ID:        credentialID,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Name:      "Full Credential",
			Username:  "full_user",
			Sandbox:   false,
			Scope:     "scoped",
			Domains:   []string{"example.com", "api.example.com", "mail.example.com"},
		}

		// Marshal to JSON
		data, err := json.Marshal(credential)
		require.NoError(t, err)

		// Parse as generic map to check field presence
		var result map[string]interface{}
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)

		// All fields should be present
		assert.Contains(t, result, "object")
		assert.Contains(t, result, "id")
		assert.Contains(t, result, "created_at")
		assert.Contains(t, result, "updated_at")
		assert.Contains(t, result, "name")
		assert.Contains(t, result, "username")
		assert.Contains(t, result, "sandbox")
		assert.Contains(t, result, "scope")
		assert.Contains(t, result, "scoped_domains")

		// Unmarshal back and verify
		var decoded SMTPCredential
		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)

		assert.Equal(t, credential.Object, decoded.Object)
		assert.Equal(t, credential.ID, decoded.ID)
		assert.Equal(t, credential.Name, decoded.Name)
		assert.Equal(t, credential.Username, decoded.Username)
		assert.Equal(t, credential.Sandbox, decoded.Sandbox)
		assert.Equal(t, credential.Scope, decoded.Scope)
		assert.Len(t, decoded.Domains, 3)
		assert.Equal(t, credential.Domains, decoded.Domains)
	})

	t.Run("different scope values", func(t *testing.T) {
		// Global scope
		globalCred := SMTPCredential{
			Object:    "smtp_credential",
			ID:        credentialID,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Name:      "Global Credential",
			Username:  "global_user",
			Sandbox:   false,
			Scope:     "global",
			Domains:   []string{},
		}

		data, err := json.Marshal(globalCred)
		require.NoError(t, err)

		var decoded SMTPCredential
		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)

		assert.Equal(t, "global", decoded.Scope)
		assert.Empty(t, decoded.Domains)

		// Scoped credential
		scopedCred := SMTPCredential{
			Object:    "smtp_credential",
			ID:        credentialID,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Name:      "Scoped Credential",
			Username:  "scoped_user",
			Sandbox:   true,
			Scope:     "scoped",
			Domains:   []string{"domain1.com", "domain2.com"},
		}

		data, err = json.Marshal(scopedCred)
		require.NoError(t, err)

		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)

		assert.Equal(t, "scoped", decoded.Scope)
		assert.True(t, decoded.Sandbox)
		assert.Len(t, decoded.Domains, 2)
		assert.Contains(t, decoded.Domains, "domain1.com")
		assert.Contains(t, decoded.Domains, "domain2.com")
	})
}
