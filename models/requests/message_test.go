package requests

import (
	"testing"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func templateSendRequest(templateID uuid.UUID) CreateTemplateMessageRequest {
	return CreateTemplateMessageRequest{
		TemplateID: templateID,
		Recipients: []common.Recipient{{Email: "recipient@example.com"}},
	}
}

func TestCreateTemplateMessageRequestMarshalsTemplateID(t *testing.T) {
	templateID := uuid.MustParse("11111111-1111-4111-8111-111111111111")

	request := marshalToMap(t, templateSendRequest(templateID))

	assert.Equal(t, templateID.String(), request["template_id"])
}

func TestCreateTemplateMessageRequestOmitsFromWhenNil(t *testing.T) {
	// The API reads a request with no from as "use the template's sender".
	request := marshalToMap(t, templateSendRequest(uuid.New()))

	assert.NotContains(t, request, "from")
}

func TestCreateTemplateMessageRequestMarshalsFromWhenSet(t *testing.T) {
	message := templateSendRequest(uuid.New())
	message.From = &common.SenderAddress{Email: "sender@example.com"}

	request := marshalToMap(t, message)

	assert.Equal(t, map[string]interface{}{"email": "sender@example.com"}, request["from"])
}

func TestCreateTemplateMessageRequestOmitsSubjectWhenEmpty(t *testing.T) {
	// The API uses the template's subject when the request gives none.
	request := marshalToMap(t, templateSendRequest(uuid.New()))

	assert.NotContains(t, request, "subject")
}

func TestCreateMessageRequestAlwaysMarshalsFrom(t *testing.T) {
	request := marshalToMap(t, CreateMessageRequest{
		Recipients: []common.Recipient{{Email: "recipient@example.com"}},
		Subject:    "Hello",
	})

	assert.Contains(t, request, "from")
}
