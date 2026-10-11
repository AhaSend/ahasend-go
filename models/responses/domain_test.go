package responses

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/AhaSend/ahasend-go/models/common"
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
		trackingSub := "click"
		returnPathSub := "mail"
		subscriptionSub := "preferences"
		mediaSub := "media"
		rotationDays := 45

		domain := Domain{
			Object:                   "domain",
			ID:                       domainID,
			CreatedAt:                now,
			UpdatedAt:                now,
			LastDNSCheckAt:           &lastCheck,
			Domain:                   "example.com",
			AccountID:                accountID,
			TrackingSubdomain:        &trackingSub,
			ReturnPathSubdomain:      &returnPathSub,
			SubscriptionSubdomain:    &subscriptionSub,
			MediaSubdomain:           &mediaSub,
			DKIMRotationIntervalDays: &rotationDays,
			RotationReady:            true,
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

		// Verify DNS settings fields
		require.NotNil(t, unmarshaled.TrackingSubdomain)
		assert.Equal(t, "click", *unmarshaled.TrackingSubdomain)
		require.NotNil(t, unmarshaled.ReturnPathSubdomain)
		assert.Equal(t, "mail", *unmarshaled.ReturnPathSubdomain)
		require.NotNil(t, unmarshaled.SubscriptionSubdomain)
		assert.Equal(t, "preferences", *unmarshaled.SubscriptionSubdomain)
		require.NotNil(t, unmarshaled.MediaSubdomain)
		assert.Equal(t, "media", *unmarshaled.MediaSubdomain)
		require.NotNil(t, unmarshaled.DKIMRotationIntervalDays)
		assert.Equal(t, 45, *unmarshaled.DKIMRotationIntervalDays)
		assert.True(t, unmarshaled.RotationReady)

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

		// Optional pointer fields should be omitted when nil
		assert.NotContains(t, result, "last_dns_check_at")
		assert.NotContains(t, result, "tracking_subdomain")
		assert.NotContains(t, result, "return_path_subdomain")
		assert.NotContains(t, result, "subscription_subdomain")
		assert.NotContains(t, result, "media_subdomain")
		assert.NotContains(t, result, "dkim_rotation_interval_days")
		assert.NotContains(t, result, "dkim_selector")
		assert.NotContains(t, result, "paused_at")
		assert.NotContains(t, result, "pause_reason")

		// These should be present
		assert.Contains(t, result, "object")
		assert.Contains(t, result, "domain")
		assert.Contains(t, result, "dns_valid")
		assert.Contains(t, result, "dns_records")
		assert.Contains(t, result, "rotation_ready")
		assert.Contains(t, result, "sending_type")
		assert.Contains(t, result, "paused")
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

func TestDomain_MarketingFieldsRoundTrip(t *testing.T) {
	pausedAt := time.Date(2026, 10, 1, 12, 30, 0, 0, time.UTC)
	selector := "partner1"
	reason := DomainPauseReasonBounceRate

	domain := Domain{
		Object:       "domain",
		ID:           uuid.MustParse("01234567-89ab-cdef-0123-456789abcdef"),
		Domain:       "example.com",
		AccountID:    uuid.MustParse("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
		DNSRecords:   []DNSRecord{},
		DKIMSelector: &selector,
		SendingType:  common.DomainSendingTypeMarketing,
		Paused:       true,
		PausedAt:     &pausedAt,
		PauseReason:  &reason,
	}

	data, err := json.Marshal(domain)
	require.NoError(t, err)

	var raw map[string]interface{}
	require.NoError(t, json.Unmarshal(data, &raw))
	assert.Equal(t, "marketing", raw["sending_type"])
	assert.Equal(t, true, raw["paused"])
	assert.Equal(t, "2026-10-01T12:30:00Z", raw["paused_at"])
	assert.Equal(t, "bounce_rate", raw["pause_reason"])
	assert.Equal(t, "partner1", raw["dkim_selector"])

	var decoded Domain
	require.NoError(t, json.Unmarshal(data, &decoded))
	assert.Equal(t, common.DomainSendingTypeMarketing, decoded.SendingType)
	assert.True(t, decoded.Paused)
	require.NotNil(t, decoded.PausedAt)
	assert.True(t, decoded.PausedAt.Equal(pausedAt))
	require.NotNil(t, decoded.PauseReason)
	assert.Equal(t, DomainPauseReasonBounceRate, *decoded.PauseReason)
	require.NotNil(t, decoded.DKIMSelector)
	assert.Equal(t, "partner1", *decoded.DKIMSelector)
}

func TestDomain_DecodesSpecPayload(t *testing.T) {
	// Every required field, with the nullable ones null, as the API sends a
	// transactional domain that is not paused.
	payload := `{
		"object": "domain",
		"id": "01234567-89ab-cdef-0123-456789abcdef",
		"created_at": "2026-10-01T12:00:00Z",
		"updated_at": "2026-10-01T12:00:00Z",
		"domain": "example.com",
		"account_id": "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
		"dns_records": [],
		"last_dns_check_at": null,
		"dns_valid": true,
		"tracking_subdomain": null,
		"return_path_subdomain": null,
		"subscription_subdomain": null,
		"media_subdomain": null,
		"dkim_rotation_interval_days": null,
		"dkim_selector": null,
		"rotation_ready": false,
		"dsn_recipient": null,
		"sending_type": "transactional",
		"paused": false,
		"paused_at": null,
		"pause_reason": null
	}`

	var domain Domain
	require.NoError(t, json.Unmarshal([]byte(payload), &domain))
	assert.Equal(t, common.DomainSendingTypeTransactional, domain.SendingType)
	assert.False(t, domain.Paused)
	assert.Nil(t, domain.PausedAt)
	assert.Nil(t, domain.PauseReason)
	assert.Nil(t, domain.DKIMSelector)
}

func TestDomain_DecodesPayloadWithoutMarketingFields(t *testing.T) {
	// A server that predates domain sending types and pauses sends none of
	// these fields. They decode to zero values, not to an error.
	payload := `{
		"object": "domain",
		"id": "01234567-89ab-cdef-0123-456789abcdef",
		"created_at": "2026-10-01T12:00:00Z",
		"updated_at": "2026-10-01T12:00:00Z",
		"domain": "example.com",
		"account_id": "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
		"dns_records": [],
		"dns_valid": true,
		"rotation_ready": false
	}`

	var domain Domain
	require.NoError(t, json.Unmarshal([]byte(payload), &domain))
	assert.Equal(t, "example.com", domain.Domain)
	assert.Empty(t, domain.SendingType)
	assert.False(t, domain.Paused)
	assert.Nil(t, domain.PausedAt)
	assert.Nil(t, domain.PauseReason)
	assert.Nil(t, domain.DKIMSelector)
}

func TestDomain_DecodesUnknownPauseReason(t *testing.T) {
	payload := `{"object":"domain","sending_type":"marketing","paused":true,"paused_at":"2026-10-01T12:30:00Z","pause_reason":"complaint_rate"}`

	var domain Domain
	require.NoError(t, json.Unmarshal([]byte(payload), &domain))
	assert.True(t, domain.Paused)
	require.NotNil(t, domain.PauseReason)
	assert.Equal(t, "complaint_rate", *domain.PauseReason)
}
