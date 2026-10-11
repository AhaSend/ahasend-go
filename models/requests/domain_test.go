package requests

import (
	"encoding/json"
	"testing"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func marshalToMap(t *testing.T, v interface{}) map[string]interface{} {
	t.Helper()

	data, err := json.Marshal(v)
	require.NoError(t, err)

	var result map[string]interface{}
	require.NoError(t, json.Unmarshal(data, &result))
	return result
}

func TestCreateDomainRequest_JSON(t *testing.T) {
	t.Run("omits sending type and DKIM selector when nil", func(t *testing.T) {
		body := marshalToMap(t, CreateDomainRequest{Domain: "example.com"})

		assert.Equal(t, map[string]interface{}{"domain": "example.com"}, body)
	})

	t.Run("sends sending type and DKIM selector", func(t *testing.T) {
		sendingType := common.DomainSendingTypeMarketing
		selector := "partner1"

		body := marshalToMap(t, CreateDomainRequest{
			Domain:       "example.com",
			SendingType:  &sendingType,
			DKIMSelector: &selector,
		})

		assert.Equal(t, "marketing", body["sending_type"])
		assert.Equal(t, "partner1", body["dkim_selector"])
	})

	t.Run("round trips", func(t *testing.T) {
		sendingType := common.DomainSendingTypeTransactional
		request := CreateDomainRequest{Domain: "example.com", SendingType: &sendingType}

		data, err := json.Marshal(request)
		require.NoError(t, err)
		var decoded CreateDomainRequest
		require.NoError(t, json.Unmarshal(data, &decoded))

		assert.Equal(t, request, decoded)
	})
}

func TestUpdateDomainRequest_JSON(t *testing.T) {
	t.Run("nil fields are omitted and leave the values unchanged", func(t *testing.T) {
		body := marshalToMap(t, UpdateDomainRequest{})

		assert.Empty(t, body)
	})

	t.Run("an empty DKIM selector is sent to clear the override", func(t *testing.T) {
		empty := ""

		body := marshalToMap(t, UpdateDomainRequest{DKIMSelector: &empty})

		require.Contains(t, body, "dkim_selector")
		assert.Equal(t, "", body["dkim_selector"])
		assert.NotContains(t, body, "sending_type")
	})

	t.Run("sends sending type", func(t *testing.T) {
		sendingType := common.DomainSendingTypeMarketing

		body := marshalToMap(t, UpdateDomainRequest{SendingType: &sendingType})

		assert.Equal(t, map[string]interface{}{"sending_type": "marketing"}, body)
	})
}
