package responses

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const embeddedListMembershipFixture = `{
	"object":"list_contact",
	"list_id":"22222222-2222-4222-8222-222222222222",
	"contact_id":"11111111-1111-4111-8111-111111111111",
	"email":"person@example.com",
	"subscription_status":"unconfirmed",
	"subscribed_at":null,
	"unsubscribed_at":null,
	"created_at":"2026-09-10T10:00:00Z",
	"updated_at":"2026-09-10T10:00:00Z",
	"list":{
		"object":"contact_list",
		"id":"22222222-2222-4222-8222-222222222222",
		"created_at":"2026-09-10T10:00:00Z",
		"updated_at":"2026-09-10T10:00:00Z",
		"name":"Newsletter",
		"description":"",
		"tags":[]
	}
}`

func TestListContactDecodesNullTimestampsAndEmbeddedList(t *testing.T) {
	var membership ListContact
	require.NoError(t, json.Unmarshal([]byte(embeddedListMembershipFixture), &membership))

	assert.Nil(t, membership.SubscribedAt)
	assert.Nil(t, membership.UnsubscribedAt)
	assert.Nil(t, membership.Contact)
	require.NotNil(t, membership.List)
	assert.Equal(t, "Newsletter", membership.List.Name)
	assert.Equal(t, []string{}, membership.List.Tags)
}

func TestContactListEncodesContactCountEvenWhenZero(t *testing.T) {
	encoded, err := json.Marshal(ContactList{Tags: []string{}})
	require.NoError(t, err)

	var body map[string]any
	require.NoError(t, json.Unmarshal(encoded, &body))
	assert.Equal(t, float64(0), body["contact_count"])
}

func TestBatchListContactResultOmitsIdentifiersItDidNotEcho(t *testing.T) {
	var result BatchListContactResult
	require.NoError(t, json.Unmarshal([]byte(`{"position":3,"outcome":"invalid","reason":"entry must name exactly one of email or id"}`), &result))

	assert.Equal(t, 3, result.Position)
	assert.Equal(t, BatchListContactOutcomeInvalid, result.Outcome)
	assert.Empty(t, result.Email)
	assert.Nil(t, result.ID)
	assert.Nil(t, result.Membership)
}
