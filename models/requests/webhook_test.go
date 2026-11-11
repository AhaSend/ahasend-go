package requests

import (
	"encoding/json"
	"testing"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateWebhookRequestSendsEventSettingsAndOrigins(t *testing.T) {
	t.Run("no origins means every origin, so the key is left out", func(t *testing.T) {
		body, err := json.Marshal(CreateWebhookRequest{Name: "Contacts", URL: "https://example.com/hook", Scope: "global", OnContactCreated: true})
		require.NoError(t, err)

		var fields map[string]any
		require.NoError(t, json.Unmarshal(body, &fields))
		assert.Equal(t, true, fields["on_contact_created"])
		assert.Equal(t, false, fields["on_contact_complained"])
		assert.NotContains(t, fields, "origins")
	})

	t.Run("origins are sent", func(t *testing.T) {
		body, err := json.Marshal(CreateWebhookRequest{Name: "Contacts", URL: "https://example.com/hook", Scope: "global", Origins: []string{common.WebhookOriginRecipient}})
		require.NoError(t, err)

		assert.Contains(t, string(body), `"origins":["recipient"]`)
	})
}

func TestCreateWebhookRequestSendsSendingTypeChanged(t *testing.T) {
	testCases := []struct {
		name string
		set  bool
	}{
		{name: "on", set: true},
		{name: "off", set: false},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			body, err := json.Marshal(CreateWebhookRequest{Name: "Domains", URL: "https://example.com/hook", Scope: "global", OnSendingTypeChanged: tc.set})
			require.NoError(t, err)

			var fields map[string]any
			require.NoError(t, json.Unmarshal(body, &fields))
			assert.Equal(t, tc.set, fields["on_sending_type_changed"])
		})
	}
}

func TestUpdateWebhookRequestSendsOnlySetFields(t *testing.T) {
	onContactDeleted := false
	body, err := json.Marshal(UpdateWebhookRequest{OnContactDeleted: &onContactDeleted, Origins: []string{common.WebhookOriginAPI, common.WebhookOriginDashboard}})
	require.NoError(t, err)

	assert.JSONEq(t, `{"on_contact_deleted":false,"origins":["api","dashboard"]}`, string(body))
}

func TestUpdateWebhookRequestSendsSendingTypeChangedOnlyWhenSet(t *testing.T) {
	t.Run("nil is left out", func(t *testing.T) {
		body, err := json.Marshal(UpdateWebhookRequest{})
		require.NoError(t, err)

		assert.NotContains(t, string(body), "on_sending_type_changed")
	})

	t.Run("false is sent", func(t *testing.T) {
		off := false
		body, err := json.Marshal(UpdateWebhookRequest{OnSendingTypeChanged: &off})
		require.NoError(t, err)

		assert.JSONEq(t, `{"on_sending_type_changed":false}`, string(body))
	})
}
