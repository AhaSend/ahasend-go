package webhooks

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// specSchema is the fragment of a schema in openapi/webhooks.yaml that these
// tests read: its properties, the schemas it is composed of, and the values it
// declares.
type specSchema struct {
	Ref         string                `yaml:"$ref"`
	Properties  map[string]specSchema `yaml:"properties"`
	AllOf       []specSchema          `yaml:"allOf"`
	Enum        []string              `yaml:"enum"`
	KnownValues []string              `yaml:"x-known-values"`
	Items       *specSchema           `yaml:"items"`
}

// specSchemas is the components.schemas section of openapi/webhooks.yaml.
type specSchemas struct {
	Components struct {
		Schemas map[string]specSchema `yaml:"schemas"`
	} `yaml:"components"`
}

func loadSpecSchemas(t *testing.T) map[string]specSchema {
	t.Helper()

	var spec specSchemas
	decodeWebhookSpec(t, &spec)
	require.NotEmpty(t, spec.Components.Schemas, "no schemas found in the specification")
	return spec.Components.Schemas
}

// schemaProperties returns the property names of the named schema, including
// the properties of the schemas its allOf composes.
func schemaProperties(t *testing.T, schemas map[string]specSchema, name string) []string {
	t.Helper()

	schema, ok := schemas[name]
	require.True(t, ok, "schema %s is not in the specification", name)

	var names []string
	for property := range schema.Properties {
		names = append(names, property)
	}
	for _, part := range schema.AllOf {
		if part.Ref != "" {
			names = append(names, schemaProperties(t, schemas, strings.TrimPrefix(part.Ref, "#/components/schemas/"))...)
			continue
		}
		for property := range part.Properties {
			names = append(names, property)
		}
	}
	return names
}

// jsonFieldNames returns the JSON keys a struct type decodes, including the
// keys of the structs it embeds.
func jsonFieldNames(typ reflect.Type) []string {
	var names []string
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Anonymous && field.Type.Kind() == reflect.Struct {
			names = append(names, jsonFieldNames(field.Type)...)
			continue
		}
		name := strings.Split(field.Tag.Get("json"), ",")[0]
		if name == "" || name == "-" {
			continue
		}
		names = append(names, name)
	}
	return names
}

// TestSpecSchemasMatchGoTypes checks that each data schema in the
// specification and the Go type that decodes it name the same keys. A key the
// spec has and the type lacks is a field the SDK drops; a key the type has and
// the spec lacks is a field nobody sends. Unlike TestSpecExamplesDecodeStrictly,
// this also sees properties that no example carries.
func TestSpecSchemasMatchGoTypes(t *testing.T) {
	schemas := loadSpecSchemas(t)

	types := map[string]reflect.Type{
		"MessageWebhookData":        reflect.TypeOf(MessageEventData{}),
		"MessageClickedWebhookData": reflect.TypeOf(MessageClickedEventData{}),
		"DeliveryAttempt":           reflect.TypeOf(DeliveryAttempt{}),
		"SuppressionWebhookData":    reflect.TypeOf(SuppressionEventData{}),
		"DomainWebhookData":         reflect.TypeOf(DomainEventData{}),
		"RouteWebhookData":          reflect.TypeOf(RouteEventData{}),
		"RouteAttachment":           reflect.TypeOf(RouteAttachment{}),
		"WebhookUnsubscribe":        reflect.TypeOf(WebhookUnsubscribe{}),
		"WebhookComplaint":          reflect.TypeOf(WebhookComplaint{}),
		"WebhookList":               reflect.TypeOf(WebhookList{}),
		"WebhookListContact":        reflect.TypeOf(WebhookListContact{}),
		"Contact":                   reflect.TypeOf(Contact{}),
		"ContactWebhookData":        reflect.TypeOf(ContactEventData{}),
		"ContactUpdatedWebhookData": reflect.TypeOf(ContactUpdatedEventData{}),
		"ContactListWebhookData":    reflect.TypeOf(ContactListEventData{}),
	}

	for name, typ := range types {
		name, typ := name, typ
		t.Run(name, func(t *testing.T) {
			assert.ElementsMatch(t, schemaProperties(t, schemas, name), jsonFieldNames(typ),
				"the properties of %s and the json tags of %s differ", name, typ)
		})
	}
}

// TestKnownValuesMatchSpec checks the constants of each open or enumerated
// string field against the values the specification lists, the same way
// TestDeliveryAttemptClassificationMatchesSpec does for classifications.
func TestKnownValuesMatchSpec(t *testing.T) {
	schemas := loadSpecSchemas(t)

	var changedFields specSchema
	for _, part := range schemas["ContactUpdatedWebhookData"].AllOf {
		if field, ok := part.Properties["changed_fields"]; ok {
			changedFields = field
		}
	}
	require.NotNil(t, changedFields.Items, "ContactUpdatedWebhookData has no changed_fields property")

	testCases := []struct {
		name     string
		spec     []string
		declared []string
	}{
		{
			name: "complaint reporter",
			spec: schemas["WebhookComplaint"].Properties["reporter"].KnownValues,
			declared: []string{
				string(ComplaintReporterYahoo),
				string(ComplaintReporterMicrosoft),
				string(ComplaintReporterComcast),
			},
		},
		{
			name: "changed fields",
			spec: changedFields.KnownValues,
			declared: []string{
				string(ChangedFieldEmail),
				string(ChangedFieldFirstName),
				string(ChangedFieldLastName),
				string(ChangedFieldStatus),
				string(ChangedFieldStatusReason),
				string(ChangedFieldAttributes),
			},
		},
		{
			name: "origin",
			spec: schemas["WebhookOrigin"].KnownValues,
			declared: []string{
				string(OriginRecipient),
				string(OriginAPI),
				string(OriginDashboard),
				string(OriginSystem),
			},
		},
		{
			name: "unsubscribe method",
			spec: schemas["WebhookUnsubscribe"].Properties["method"].Enum,
			declared: []string{
				string(UnsubscribeMethodOneClick),
				string(UnsubscribeMethodLink),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			require.NotEmpty(t, tc.spec, "the specification lists no values")
			assert.ElementsMatch(t, tc.spec, tc.declared)
		})
	}

	t.Run("open sets have no enum", func(t *testing.T) {
		assert.Empty(t, schemas["WebhookComplaint"].Properties["reporter"].Enum)
		assert.Empty(t, schemas["WebhookOrigin"].Enum)
		assert.Empty(t, changedFields.Enum)
		require.NotNil(t, changedFields.Items)
		assert.Empty(t, changedFields.Items.Enum)
	})
}

// contactJSON is the contact of the specification's examples, with attributes
// that hold every kind of JSON value.
const contactJSON = `{
	"object": "contact",
	"id": "e43e4167-f186-4a7f-9b6e-ac9fdc6426e1",
	"created_at": "2026-09-21T10:15:00Z",
	"updated_at": "2026-10-21T10:15:00Z",
	"email": "someone@example.com",
	"first_name": "Ada",
	"last_name": "Lovelace",
	"status": "enabled",
	"status_reason": "",
	"unsubscribed": false,
	"unsubscribed_at": null,
	"attributes": {
		"plan": "pro",
		"seats": 3,
		"trial": true,
		"referrer": null,
		"tags": ["a", 1, {"deep": [null]}],
		"address": {"city": "London", "geo": {"lat": 51.5}}
	},
	"validation_status": "valid",
	"last_validated_at": "2026-10-20T10:15:00Z"
}`

// listJSON and listContactJSON are the membership of the specification's
// examples.
const (
	listJSON = `{
		"id": "8b5a3112-8fa1-4bde-9c20-8af76484daf4",
		"name": "Newsletter"
	}`
	listContactJSON = `{
		"object": "list_contact",
		"list_id": "8b5a3112-8fa1-4bde-9c20-8af76484daf4",
		"contact_id": "e43e4167-f186-4a7f-9b6e-ac9fdc6426e1",
		"email": "someone@example.com",
		"subscription_status": "unsubscribed",
		"subscribed_at": "2026-09-21T10:15:00Z",
		"unsubscribed_at": "2026-10-21T10:15:00Z",
		"created_at": "2026-09-21T10:15:00Z",
		"updated_at": "2026-10-21T10:15:00Z"
	}`
)

// contactEventPayload renders a contact webhook payload of the given type with
// the supplied JSON fragment spliced into its data object. The fragment carries
// its own leading comma, so "" produces a payload with no extra fields.
func contactEventPayload(eventType, origin, fragment string) string {
	return fmt.Sprintf(`{
		"type": %q,
		"webhook_id": "abe11757-2886-4b55-96f1-0e0afc95795a",
		"timestamp": "2026-10-21T10:15:00Z",
		"data": {
			"account_id": "4cdd7bdd-294e-4762-892f-83d40abf5a87",
			"origin": %q,
			"contact": %s%s
		}
	}`, eventType, origin, contactJSON, fragment)
}

// parseSigned signs and parses a payload, and fails the test on an error.
func parseSigned(t *testing.T, verifier *WebhookVerifier, payload string) WebhookEvent {
	t.Helper()

	event, err := verifier.Parse([]byte(payload), signedHeaders(t, verifier, payload))
	require.NoError(t, err)
	require.NotNil(t, event)
	return event
}

func TestMessageCampaignEventParsing(t *testing.T) {
	verifier, err := NewWebhookVerifier(testWebhookSecret)
	require.NoError(t, err)

	t.Run("message.unsubscribed", func(t *testing.T) {
		payload := messageEventPayload("message.unsubscribed", "on_unsubscribed", `,
			"is_bot": false,
			"unsubscribe": {"method": "one_click"},
			"campaign_id": "8b0f4c2e-5d1a-4f6b-9c3e-2a7d1e9f0b44",
			"list_id": "c3e1a9d2-7f4b-4e8a-b5c6-1d2e3f4a5b6c",
			"segment_id": null`)

		event, ok := parseSigned(t, verifier, payload).(*MessageUnsubscribedEvent)
		require.True(t, ok)
		assert.Equal(t, "message.unsubscribed", event.GetType())
		assert.Equal(t, "on_unsubscribed", event.Data.Event)
		require.NotNil(t, event.Data.Unsubscribe)
		assert.Equal(t, UnsubscribeMethodOneClick, event.Data.Unsubscribe.Method)
		assert.Nil(t, event.Data.Complaint)
		require.NotNil(t, event.Data.CampaignID)
		assert.Equal(t, "8b0f4c2e-5d1a-4f6b-9c3e-2a7d1e9f0b44", *event.Data.CampaignID)
		require.NotNil(t, event.Data.ListID)
		assert.Equal(t, "c3e1a9d2-7f4b-4e8a-b5c6-1d2e3f4a5b6c", *event.Data.ListID)
		assert.Nil(t, event.Data.SegmentID)
	})

	t.Run("message.complained with an unknown reporter", func(t *testing.T) {
		payload := messageEventPayload("message.complained", "on_complained", `,
			"complaint": {"feedback_type": "abuse", "reporter": "some_future_loop"},
			"campaign_id": "8b0f4c2e-5d1a-4f6b-9c3e-2a7d1e9f0b44",
			"list_id": null,
			"segment_id": "5f1d2c3b-4a59-4e6f-8a7b-9c0d1e2f3a4b"`)

		event, ok := parseSigned(t, verifier, payload).(*MessageComplainedEvent)
		require.True(t, ok)
		require.NotNil(t, event.Data.Complaint)
		assert.Equal(t, "abuse", event.Data.Complaint.FeedbackType)
		assert.Equal(t, ComplaintReporter("some_future_loop"), event.Data.Complaint.Reporter)
		assert.Nil(t, event.Data.Unsubscribe)
		assert.Nil(t, event.Data.ListID)
		require.NotNil(t, event.Data.SegmentID)
		assert.Equal(t, "5f1d2c3b-4a59-4e6f-8a7b-9c0d1e2f3a4b", *event.Data.SegmentID)
	})

	t.Run("list_id and segment_id null and absent decode the same", func(t *testing.T) {
		fragments := map[string]string{
			"null":   `, "campaign_id": "8b0f4c2e-5d1a-4f6b-9c3e-2a7d1e9f0b44", "list_id": null, "segment_id": null`,
			"absent": `, "campaign_id": "8b0f4c2e-5d1a-4f6b-9c3e-2a7d1e9f0b44"`,
		}
		for name, fragment := range fragments {
			fragment := fragment
			t.Run(name, func(t *testing.T) {
				payload := messageEventPayload("message.delivered", "on_delivered", fragment)

				data := GetMessageEventData(parseSigned(t, verifier, payload))
				require.NotNil(t, data)
				require.NotNil(t, data.CampaignID)
				assert.Nil(t, data.ListID)
				assert.Nil(t, data.SegmentID)
			})
		}
	})

	t.Run("a transactional message has no campaign fields", func(t *testing.T) {
		payload := messageEventPayload("message.delivered", "on_delivered", "")

		data := GetMessageEventData(parseSigned(t, verifier, payload))
		require.NotNil(t, data)
		assert.Nil(t, data.CampaignID)
		assert.Nil(t, data.ListID)
		assert.Nil(t, data.SegmentID)
		assert.Nil(t, data.Unsubscribe)
		assert.Nil(t, data.Complaint)
	})

	t.Run("message.clicked carries the campaign fields", func(t *testing.T) {
		payload := `{
			"type": "message.clicked",
			"webhook_id": "abe11757-2886-4b55-96f1-0e0afc95795a",
			"timestamp": "2024-05-06T10:20:16.687031577Z",
			"data": {
				"account_id": "4cdd7bdd-294e-4762-892f-83d40abf5a87",
				"event": "on_clicked",
				"from": "Example News <news@example.com>",
				"recipient": "recipient@example.com",
				"subject": "Our October news",
				"message_id_header": "<message-id-12345@localhost>",
				"id": "407926766d2711f09b30960002cafe7c",
				"url": "https://example.com/landing-page",
				"user_agent": "Mozilla/5.0",
				"ip": "192.168.1.100",
				"is_bot": false,
				"campaign_id": "8b0f4c2e-5d1a-4f6b-9c3e-2a7d1e9f0b44",
				"list_id": "c3e1a9d2-7f4b-4e8a-b5c6-1d2e3f4a5b6c",
				"segment_id": null
			}
		}`

		event, ok := parseSigned(t, verifier, payload).(*MessageClickedEvent)
		require.True(t, ok)
		require.NotNil(t, event.Data.CampaignID)
		assert.Equal(t, "8b0f4c2e-5d1a-4f6b-9c3e-2a7d1e9f0b44", *event.Data.CampaignID)
		require.NotNil(t, event.Data.ListID)
		assert.Equal(t, "c3e1a9d2-7f4b-4e8a-b5c6-1d2e3f4a5b6c", *event.Data.ListID)
		assert.Nil(t, event.Data.SegmentID)
	})
}

func TestSuppressionDeletedParsing(t *testing.T) {
	verifier, err := NewWebhookVerifier(testWebhookSecret)
	require.NoError(t, err)

	payload := `{
		"type": "suppression.deleted",
		"webhook_id": "abe11757-2886-4b55-96f1-0e0afc95795a",
		"timestamp": "2026-10-21T10:15:00Z",
		"data": {
			"account_id": "4cdd7bdd-294e-4762-892f-83d40abf5a87",
			"origin": "api",
			"recipient": "user@example.org",
			"created_at": "2026-10-01T08:00:00Z",
			"expires_at": "2026-10-31T08:00:00Z",
			"reason": "Too many hard bounces",
			"sending_domain": "example.com"
		}
	}`

	event, ok := parseSigned(t, verifier, payload).(*SuppressionDeletedEvent)
	require.True(t, ok)
	assert.Equal(t, "suppression.deleted", event.GetType())
	assert.Equal(t, OriginAPI, event.Data.Origin)
	assert.Equal(t, "user@example.org", event.Data.Recipient)
	assert.Equal(t, "example.com", event.Data.SendingDomain)
	assert.Equal(t, time.Date(2026, 10, 31, 8, 0, 0, 0, time.UTC), event.Data.ExpiresAt.UTC())
}

func TestContactEventParsing(t *testing.T) {
	verifier, err := NewWebhookVerifier(testWebhookSecret)
	require.NoError(t, err)

	t.Run("contact.created with nested attributes", func(t *testing.T) {
		payload := contactEventPayload("contact.created", "api", "")

		event, ok := parseSigned(t, verifier, payload).(*ContactCreatedEvent)
		require.True(t, ok)
		assert.Equal(t, "contact.created", event.GetType())
		assert.Equal(t, OriginAPI, event.Data.Origin)
		assert.Nil(t, event.Data.CampaignID)
		assert.Nil(t, event.Data.SegmentID)

		contact := event.Data.Contact
		assert.Equal(t, "contact", contact.Object)
		assert.Equal(t, "e43e4167-f186-4a7f-9b6e-ac9fdc6426e1", contact.ID)
		assert.Equal(t, time.Date(2026, 9, 21, 10, 15, 0, 0, time.UTC), contact.CreatedAt.UTC())
		assert.Equal(t, time.Date(2026, 10, 21, 10, 15, 0, 0, time.UTC), contact.UpdatedAt.UTC())
		assert.Equal(t, "enabled", contact.Status)
		assert.Equal(t, "", contact.StatusReason)
		assert.Equal(t, "valid", contact.ValidationStatus)
		assert.Equal(t, "someone@example.com", contact.Email)
		assert.Equal(t, "Ada", contact.FirstName)
		assert.Equal(t, "Lovelace", contact.LastName)
		assert.Equal(t, "enabled", contact.Status)
		assert.False(t, contact.Unsubscribed)
		assert.Nil(t, contact.UnsubscribedAt)
		require.NotNil(t, contact.LastValidatedAt)
		assert.Equal(t, time.Date(2026, 10, 20, 10, 15, 0, 0, time.UTC), contact.LastValidatedAt.UTC())

		assert.Equal(t, map[string]any{
			"plan":     "pro",
			"seats":    float64(3),
			"trial":    true,
			"referrer": nil,
			"tags":     []any{"a", float64(1), map[string]any{"deep": []any{nil}}},
			"address":  map[string]any{"city": "London", "geo": map[string]any{"lat": 51.5}},
		}, contact.Attributes)
	})

	t.Run("contact.updated with an unknown changed field", func(t *testing.T) {
		payload := contactEventPayload("contact.updated", "dashboard",
			`, "changed_fields": ["first_name", "attributes", "some_future_field"]`)

		event, ok := parseSigned(t, verifier, payload).(*ContactUpdatedEvent)
		require.True(t, ok)
		assert.Equal(t, OriginDashboard, event.Data.Origin)
		assert.Equal(t, "Ada", event.Data.Contact.FirstName)
		assert.Equal(t, []ContactChangedField{
			ChangedFieldFirstName,
			ChangedFieldAttributes,
			ContactChangedField("some_future_field"),
		}, event.Data.ChangedFields)
	})

	t.Run("contact.deleted", func(t *testing.T) {
		payload := contactEventPayload("contact.deleted", "system", "")

		event, ok := parseSigned(t, verifier, payload).(*ContactDeletedEvent)
		require.True(t, ok)
		assert.Equal(t, OriginSystem, event.Data.Origin)
		assert.Equal(t, "someone@example.com", event.Data.Contact.Email)
	})

	t.Run("membership events carry the list and the membership", func(t *testing.T) {
		fragment := fmt.Sprintf(`, "list": %s, "list_contact": %s`, listJSON, listContactJSON)

		// The two events share their data type, so only the event type tells
		// them apart: each must come back as its own.
		testCases := []struct {
			eventType string
			want      WebhookEvent
		}{
			{eventType: "contact.added_to_list", want: &ContactAddedToListEvent{}},
			{eventType: "contact.removed_from_list", want: &ContactRemovedFromListEvent{}},
		}
		for _, tc := range testCases {
			eventType := tc.eventType
			want := tc.want
			t.Run(eventType, func(t *testing.T) {
				event := parseSigned(t, verifier, contactEventPayload(eventType, "api", fragment))
				require.IsType(t, want, event)

				var data ContactListEventData
				switch e := event.(type) {
				case *ContactAddedToListEvent:
					data = e.Data
				case *ContactRemovedFromListEvent:
					data = e.Data
				}
				assert.Equal(t, eventType, event.GetType())
				require.NotNil(t, data.List)
				assert.Equal(t, "8b5a3112-8fa1-4bde-9c20-8af76484daf4", data.List.ID)
				assert.Equal(t, "Newsletter", data.List.Name)
				require.NotNil(t, data.ListContact)
				assert.Equal(t, "8b5a3112-8fa1-4bde-9c20-8af76484daf4", data.ListContact.ListID)
				assert.Equal(t, "e43e4167-f186-4a7f-9b6e-ac9fdc6426e1", data.ListContact.ContactID)
				assert.Equal(t, "list_contact", data.ListContact.Object)
				assert.Equal(t, "someone@example.com", data.ListContact.Email)
				assert.Equal(t, time.Date(2026, 9, 21, 10, 15, 0, 0, time.UTC), data.ListContact.CreatedAt.UTC())
				assert.Equal(t, time.Date(2026, 10, 21, 10, 15, 0, 0, time.UTC), data.ListContact.UpdatedAt.UTC())
				assert.Equal(t, "unsubscribed", data.ListContact.SubscriptionStatus)
				require.NotNil(t, data.ListContact.SubscribedAt)
				require.NotNil(t, data.ListContact.UnsubscribedAt)
				assert.Equal(t, time.Date(2026, 10, 21, 10, 15, 0, 0, time.UTC), data.ListContact.UnsubscribedAt.UTC())
				assert.Nil(t, data.PreviousSubscriptionStatus)
				assert.Nil(t, data.Complaint)
			})
		}
	})

	t.Run("contact.subscribed for a membership", func(t *testing.T) {
		fragment := fmt.Sprintf(`, "list": %s, "list_contact": %s, "previous_subscription_status": "unsubscribed"`,
			listJSON, listContactJSON)

		event, ok := parseSigned(t, verifier, contactEventPayload("contact.subscribed", "api", fragment)).(*ContactSubscribedEvent)
		require.True(t, ok)
		require.NotNil(t, event.Data.List)
		require.NotNil(t, event.Data.PreviousSubscriptionStatus)
		assert.Equal(t, "unsubscribed", *event.Data.PreviousSubscriptionStatus)
	})

	t.Run("list and list_contact null and absent decode the same", func(t *testing.T) {
		fragments := map[string]string{
			"null":   `, "list": null, "list_contact": null`,
			"absent": "",
		}
		for name, fragment := range fragments {
			fragment := fragment
			t.Run(name, func(t *testing.T) {
				event, ok := parseSigned(t, verifier, contactEventPayload("contact.subscribed", "api", fragment)).(*ContactSubscribedEvent)
				require.True(t, ok)
				assert.Nil(t, event.Data.List)
				assert.Nil(t, event.Data.ListContact)
				assert.Nil(t, event.Data.PreviousSubscriptionStatus)
			})
		}
	})

	t.Run("contact.unsubscribed from a campaign sent to a segment", func(t *testing.T) {
		fragment := `,
			"campaign_id": "3af77d52-096f-42a2-8cc7-b6eccdccff8e",
			"segment_id": "5f1d2c3b-4a59-4e6f-8a7b-9c0d1e2f3a4b",
			"list": null,
			"list_contact": null`

		event, ok := parseSigned(t, verifier, contactEventPayload("contact.unsubscribed", "recipient", fragment)).(*ContactUnsubscribedEvent)
		require.True(t, ok)
		assert.Equal(t, OriginRecipient, event.Data.Origin)
		require.NotNil(t, event.Data.CampaignID)
		assert.Equal(t, "3af77d52-096f-42a2-8cc7-b6eccdccff8e", *event.Data.CampaignID)
		require.NotNil(t, event.Data.SegmentID)
		assert.Equal(t, "5f1d2c3b-4a59-4e6f-8a7b-9c0d1e2f3a4b", *event.Data.SegmentID)
		assert.Nil(t, event.Data.List)
		assert.Nil(t, event.Data.ListContact)
	})

	t.Run("an unknown origin decodes", func(t *testing.T) {
		event, ok := parseSigned(t, verifier, contactEventPayload("contact.created", "some_future_origin", "")).(*ContactCreatedEvent)
		require.True(t, ok)
		assert.Equal(t, WebhookOrigin("some_future_origin"), event.Data.Origin)
	})

	t.Run("contact.complained about a campaign sent to a segment", func(t *testing.T) {
		fragment := `,
			"campaign_id": "3af77d52-096f-42a2-8cc7-b6eccdccff8e",
			"segment_id": "5f1d2c3b-4a59-4e6f-8a7b-9c0d1e2f3a4b",
			"list": null,
			"list_contact": null,
			"complaint": {"feedback_type": "abuse", "reporter": "yahoo"}`

		event, ok := parseSigned(t, verifier, contactEventPayload("contact.complained", "recipient", fragment)).(*ContactComplainedEvent)
		require.True(t, ok)
		require.NotNil(t, event.Data.SegmentID)
		assert.Equal(t, "5f1d2c3b-4a59-4e6f-8a7b-9c0d1e2f3a4b", *event.Data.SegmentID)
		assert.Nil(t, event.Data.List)
		assert.Nil(t, event.Data.ListContact)
		assert.Nil(t, event.Data.PreviousSubscriptionStatus)
		require.NotNil(t, event.Data.Complaint)
		assert.Equal(t, ComplaintReporterYahoo, event.Data.Complaint.Reporter)
	})

	t.Run("contact.complained with an unknown reporter", func(t *testing.T) {
		fragment := fmt.Sprintf(`,
			"campaign_id": "3af77d52-096f-42a2-8cc7-b6eccdccff8e",
			"list": %s,
			"list_contact": %s,
			"previous_subscription_status": "confirmed",
			"complaint": {"feedback_type": "fraud", "reporter": "some_future_loop"}`, listJSON, listContactJSON)

		event, ok := parseSigned(t, verifier, contactEventPayload("contact.complained", "recipient", fragment)).(*ContactComplainedEvent)
		require.True(t, ok)
		require.NotNil(t, event.Data.Complaint)
		assert.Equal(t, "fraud", event.Data.Complaint.FeedbackType)
		assert.Equal(t, ComplaintReporter("some_future_loop"), event.Data.Complaint.Reporter)
		require.NotNil(t, event.Data.PreviousSubscriptionStatus)
		assert.Equal(t, "confirmed", *event.Data.PreviousSubscriptionStatus)
		assert.Nil(t, event.Data.SegmentID)
	})

	t.Run("an unknown event type keeps returning ErrUnknownEventType", func(t *testing.T) {
		payload := contactEventPayload("contact.some_future_event", "api", "")

		event, err := verifier.Parse([]byte(payload), signedHeaders(t, verifier, payload))
		assert.ErrorIs(t, err, ErrUnknownEventType)
		assert.Nil(t, event)
	})
}
