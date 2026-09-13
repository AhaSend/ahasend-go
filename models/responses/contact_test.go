package responses

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContactDecodesLegacyJSONAttributesWithoutLoss(t *testing.T) {
	payload := []byte(`{
		"object":"contact",
		"id":"11111111-1111-4111-8111-111111111111",
		"created_at":"2026-09-10T10:00:00Z",
		"updated_at":"2026-09-10T11:00:00Z",
		"email":"person@example.com",
		"first_name":"Pat",
		"last_name":"Example",
		"status":"enabled",
		"status_reason":"",
		"unsubscribed":false,
		"unsubscribed_at":null,
		"attributes":{
			"null_value":null,
			"boolean":true,
			"number":12.5,
			"string":"legacy",
			"array":[1,"two",false,null],
			"object":{"nested":[{"ok":true}]}
		},
		"validation_status":"unvalidated",
		"last_validated_at":null
	}`)

	var contact Contact
	require.NoError(t, json.Unmarshal(payload, &contact))

	assert.Contains(t, contact.Attributes, "null_value")
	assert.Nil(t, contact.Attributes["null_value"])
	assert.Equal(t, true, contact.Attributes["boolean"])
	assert.Equal(t, 12.5, contact.Attributes["number"])
	assert.Equal(t, "legacy", contact.Attributes["string"])
	assert.Equal(t, []any{float64(1), "two", false, nil}, contact.Attributes["array"])
	assert.Equal(t, map[string]any{"nested": []any{map[string]any{"ok": true}}}, contact.Attributes["object"])
}

func TestBatchUpsertContactsResponseDecodesOrderedOutcomes(t *testing.T) {
	payload := []byte(`{
		"object":"list","created":1,"updated":1,"failed":1,
		"data":[
			{"position":0,"email":"new@example.com","outcome":"created","contact":{"object":"contact","id":"11111111-1111-4111-8111-111111111111","created_at":"2026-09-10T10:00:00Z","updated_at":"2026-09-10T10:00:00Z","email":"new@example.com","first_name":"","last_name":"","status":"enabled","status_reason":"","unsubscribed":false,"unsubscribed_at":null,"attributes":{"source":"api"},"validation_status":"unvalidated","last_validated_at":null}},
			{"position":1,"email":"old@example.com","outcome":"updated","contact":{"object":"contact","id":"22222222-2222-4222-8222-222222222222","created_at":"2026-09-09T10:00:00Z","updated_at":"2026-09-10T10:00:00Z","email":"old@example.com","first_name":"","last_name":"","status":"enabled","status_reason":"","unsubscribed":false,"unsubscribed_at":null,"attributes":{},"validation_status":"unvalidated","last_validated_at":null}},
			{"position":2,"email":"bad@example.com","outcome":"failed","reason":"attribute score has the wrong type"}
		]
	}`)

	var response BatchUpsertContactsResponse
	require.NoError(t, json.Unmarshal(payload, &response))

	assert.Equal(t, 1, response.Created)
	assert.Equal(t, 1, response.Updated)
	assert.Equal(t, 1, response.Failed)
	require.Len(t, response.Data, 3)
	assert.Equal(t, []int{0, 1, 2}, []int{response.Data[0].Position, response.Data[1].Position, response.Data[2].Position})
	assert.Equal(t, "created", response.Data[0].Outcome)
	require.NotNil(t, response.Data[0].Contact)
	assert.Equal(t, "api", response.Data[0].Contact.Attributes["source"])
	assert.Equal(t, "attribute score has the wrong type", response.Data[2].Reason)
	assert.Nil(t, response.Data[2].Contact)
}
