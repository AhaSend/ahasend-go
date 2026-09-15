package requests

import (
	"encoding/json"
	"testing"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func templatedSendRequest(templateID uuid.UUID) CreateMessageRequest {
	return CreateMessageRequest{
		From:       common.SenderAddress{Email: "sender@example.com"},
		Recipients: []common.Recipient{{Email: "recipient@example.com"}},
		TemplateID: &templateID,
	}
}

func TestCreateMessageRequestMarshalsTemplateID(t *testing.T) {
	templateID := uuid.MustParse("11111111-1111-4111-8111-111111111111")

	encoded, err := json.Marshal(templatedSendRequest(templateID))
	require.NoError(t, err)

	var request map[string]any
	require.NoError(t, json.Unmarshal(encoded, &request))
	assert.Equal(t, templateID.String(), request["template_id"])
}

func TestCreateMessageRequestOmitsTemplateIDWhenUnset(t *testing.T) {
	textContent := "Hello"

	encoded, err := json.Marshal(CreateMessageRequest{
		From:        common.SenderAddress{Email: "sender@example.com"},
		Recipients:  []common.Recipient{{Email: "recipient@example.com"}},
		Subject:     "Hello",
		TextContent: &textContent,
	})
	require.NoError(t, err)

	var request map[string]any
	require.NoError(t, json.Unmarshal(encoded, &request))
	assert.NotContains(t, request, "template_id")
}

func TestCreateMessageRequestSendsEmptySubjectWhenTheTemplateSuppliesOne(t *testing.T) {
	encoded, err := json.Marshal(templatedSendRequest(uuid.New()))
	require.NoError(t, err)

	// The API reads an empty subject as "no subject given, use the template's",
	// so the key has to reach the wire rather than being omitted.
	var request map[string]any
	require.NoError(t, json.Unmarshal(encoded, &request))
	require.Contains(t, request, "subject")
	assert.Equal(t, "", request["subject"])
}
