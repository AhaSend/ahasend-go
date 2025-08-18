package responses

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDomain_JSONMarshaling(t *testing.T) {
	now := time.Now().Truncate(time.Second) // Truncate for JSON precision
	lastCheck := now.Add(-1 * time.Hour)
	domainID := uuid.MustParse("01234567-89ab-cdef-0123-456789abcdef")
	accountID := uuid.MustParse("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee")

	t.Run("Domain with all fields", func(t *testing.T) {
		domain := Domain{
			Object:         "domain",
			ID:             domainID,
			CreatedAt:      now,
			UpdatedAt:      now,
			LastDNSCheckAt: &lastCheck,
			Domain:         "example.com",
			AccountID:      accountID,
			DNSRecords: []DNSRecord{
				{
					Type:       "CNAME",
					Host:       "mail.example.com",
					Content:    "mail.ahasend.com",
					Required:   true,
					Propagated: true,
				},
			},
			DNSValid: true,
		}

		// Marshal to JSON
		data, err := json.Marshal(domain)
		require.NoError(t, err)

		// Unmarshal back
		var unmarshaled Domain
		err = json.Unmarshal(data, &unmarshaled)
		require.NoError(t, err)

		// Verify fields
		assert.Equal(t, domain.Object, unmarshaled.Object)
		assert.Equal(t, domain.ID, unmarshaled.ID)
		assert.Equal(t, domain.Domain, unmarshaled.Domain)
		assert.Equal(t, domain.AccountID, unmarshaled.AccountID)
		assert.Equal(t, domain.DNSValid, unmarshaled.DNSValid)
		assert.True(t, unmarshaled.CreatedAt.Equal(now))
		assert.True(t, unmarshaled.UpdatedAt.Equal(now))

		// Verify pointer field
		assert.NotNil(t, unmarshaled.LastDNSCheckAt)
		assert.True(t, unmarshaled.LastDNSCheckAt.Equal(lastCheck))

		// Verify DNS records
		assert.Len(t, unmarshaled.DNSRecords, 1)
		record := unmarshaled.DNSRecords[0]
		assert.Equal(t, "CNAME", record.Type)
		assert.Equal(t, "mail.example.com", record.Host)
		assert.Equal(t, "mail.ahasend.com", record.Content)
		assert.True(t, record.Required)
		assert.True(t, record.Propagated)
	})

	t.Run("Domain with optional fields omitted", func(t *testing.T) {
		domain := Domain{
			Object:     "domain",
			ID:         domainID,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			Domain:     "example.com",
			AccountID:  accountID,
			DNSRecords: []DNSRecord{},
			DNSValid:   false,
		}

		data, err := json.Marshal(domain)
		require.NoError(t, err)

		// Parse as generic map to check omitempty behavior
		var result map[string]interface{}
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)

		// LastDNSCheckAt should be omitted
		assert.NotContains(t, result, "last_dns_check_at")

		// These should be present
		assert.Contains(t, result, "object")
		assert.Contains(t, result, "domain")
		assert.Contains(t, result, "dns_valid")
		assert.Contains(t, result, "dns_records")
	})
}

func TestDNSRecord_JSONMarshaling(t *testing.T) {
	record := DNSRecord{
		Type:       "TXT",
		Host:       "_ahasend.example.com",
		Content:    "v=ah1; verification=abc123",
		Required:   true,
		Propagated: false,
	}

	// Marshal to JSON
	data, err := json.Marshal(record)
	require.NoError(t, err)

	// Unmarshal back
	var unmarshaled DNSRecord
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err)

	// Verify all fields
	assert.Equal(t, record.Type, unmarshaled.Type)
	assert.Equal(t, record.Host, unmarshaled.Host)
	assert.Equal(t, record.Content, unmarshaled.Content)
	assert.Equal(t, record.Required, unmarshaled.Required)
	assert.Equal(t, record.Propagated, unmarshaled.Propagated)
}
