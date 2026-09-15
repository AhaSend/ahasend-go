package responses

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const templateFixture = `{
	"object":"template",
	"id":"11111111-1111-4111-8111-111111111111",
	"created_at":"2026-09-10T10:00:00Z",
	"updated_at":"2026-09-10T11:00:00Z",
	"name":"Password reset",
	"subject":"Reset your password",
	"preheader":"It expires in an hour",
	"variables":[{"name":"first_name","required":true},{"name":"unsubscribe_url","required":false}]
}`

func TestTemplateDecodesAPIFixture(t *testing.T) {
	var template Template
	require.NoError(t, json.Unmarshal([]byte(templateFixture), &template))

	assert.Equal(t, "template", template.Object)
	assert.Equal(t, uuid.MustParse("11111111-1111-4111-8111-111111111111"), template.ID)
	assert.Equal(t, time.Date(2026, 9, 10, 10, 0, 0, 0, time.UTC), template.CreatedAt.UTC())
	assert.Equal(t, time.Date(2026, 9, 10, 11, 0, 0, 0, time.UTC), template.UpdatedAt.UTC())
	assert.Equal(t, "Password reset", template.Name)
	assert.Equal(t, "Reset your password", template.Subject)
	assert.Equal(t, "It expires in an hour", template.Preheader)
	require.Len(t, template.Variables, 2)
	assert.Equal(t, "first_name", template.Variables[0].Name)
	assert.True(t, template.Variables[0].Required)
	assert.Equal(t, "unsubscribe_url", template.Variables[1].Name)
	assert.False(t, template.Variables[1].Required)
}

func TestTemplateRoundTripsWithoutLoss(t *testing.T) {
	var decoded Template
	require.NoError(t, json.Unmarshal([]byte(templateFixture), &decoded))

	encoded, err := json.Marshal(decoded)
	require.NoError(t, err)
	assert.JSONEq(t, templateFixture, string(encoded))

	var reDecoded Template
	require.NoError(t, json.Unmarshal(encoded, &reDecoded))
	assert.Equal(t, decoded, reDecoded)
}

func TestTemplateDecodesEmptySubjectPreheaderAndVariables(t *testing.T) {
	payload := []byte(`{
		"object":"template",
		"id":"22222222-2222-4222-8222-222222222222",
		"created_at":"2026-09-10T10:00:00Z",
		"updated_at":"2026-09-10T10:00:00Z",
		"name":"Draft",
		"subject":"",
		"preheader":"",
		"variables":[]
	}`)

	var template Template
	require.NoError(t, json.Unmarshal(payload, &template))

	assert.Equal(t, "Draft", template.Name)
	assert.Empty(t, template.Subject)
	assert.Empty(t, template.Preheader)
	require.NotNil(t, template.Variables)
	assert.Empty(t, template.Variables)

	encoded, err := json.Marshal(template)
	require.NoError(t, err)
	assert.JSONEq(t, string(payload), string(encoded))
}

func TestPaginatedTemplatesResponseDecodesPage(t *testing.T) {
	payload := []byte(`{
		"object":"list",
		"data":[` + templateFixture + `],
		"pagination":{"has_more":true,"next_cursor":"next","previous_cursor":null}
	}`)

	var page PaginatedTemplatesResponse
	require.NoError(t, json.Unmarshal(payload, &page))

	assert.Equal(t, "list", page.Object)
	require.Len(t, page.Data, 1)
	assert.Equal(t, "Password reset", page.Data[0].Name)
	require.Len(t, page.Data[0].Variables, 2)
	assert.True(t, page.Data[0].Variables[0].Required)
	assert.False(t, page.Data[0].Variables[1].Required)
	assert.True(t, page.Pagination.HasMore)
	require.NotNil(t, page.Pagination.NextCursor)
	assert.Equal(t, "next", *page.Pagination.NextCursor)
	assert.Nil(t, page.Pagination.PreviousCursor)
}

func TestPaginatedTemplatesResponseDecodesEmptyPage(t *testing.T) {
	payload := []byte(`{"object":"list","data":[],"pagination":{"has_more":false,"next_cursor":null,"previous_cursor":null}}`)

	var page PaginatedTemplatesResponse
	require.NoError(t, json.Unmarshal(payload, &page))

	assert.Equal(t, "list", page.Object)
	require.NotNil(t, page.Data)
	assert.Empty(t, page.Data)
	assert.False(t, page.Pagination.HasMore)
	assert.Nil(t, page.Pagination.NextCursor)
	assert.Nil(t, page.Pagination.PreviousCursor)

	encoded, err := json.Marshal(page)
	require.NoError(t, err)
	assert.JSONEq(t, `{"object":"list","data":[],"pagination":{"has_more":false}}`, string(encoded))
}
