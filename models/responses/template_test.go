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

const templateFixture = `{
	"object":"template",
	"id":"11111111-1111-4111-8111-111111111111",
	"created_at":"2026-09-10T10:00:00Z",
	"updated_at":"2026-09-10T11:00:00Z",
	"name":"Password reset",
	"subject":"Reset your password",
	"preheader":"It expires in an hour",
	"variables":[{"name":"first_name","required":true,"kind":"value"},{"name":"items","required":true,"kind":"list","item_fields":["name","price"]}],
	"from":{"email":"hello@example.com","name":"Example"},
	"reply_to":{"email":"support@example.com","name":""},
	"editor":"advanced",
	"has_draft":true,
	"content":{"mjml":"<mjml><mj-body></mj-body></mjml>","text":"Reset it","text_is_custom":false}
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
	assert.Equal(t, TemplateVariableKindValue, template.Variables[0].Kind)
	assert.Nil(t, template.Variables[0].ItemFields)
	assert.Equal(t, "items", template.Variables[1].Name)
	assert.True(t, template.Variables[1].Required)
	assert.Equal(t, TemplateVariableKindList, template.Variables[1].Kind)
	assert.Equal(t, []string{"name", "price"}, template.Variables[1].ItemFields)
	require.NotNil(t, template.From)
	assert.Equal(t, "hello@example.com", template.From.Email)
	require.NotNil(t, template.From.Name)
	assert.Equal(t, "Example", *template.From.Name)
	require.NotNil(t, template.ReplyTo)
	assert.Equal(t, "support@example.com", template.ReplyTo.Email)
	assert.Equal(t, common.TemplateEditorAdvanced, template.Editor)
	assert.True(t, template.HasDraft)
	require.NotNil(t, template.Content)
	assert.Equal(t, "<mjml><mj-body></mj-body></mjml>", template.Content.MJML)
	assert.Empty(t, template.Content.HTML)
	assert.Equal(t, "Reset it", template.Content.Text)
	assert.False(t, template.Content.TextIsCustom)
}

func TestTemplateDecodesNullSenderAndReplyTo(t *testing.T) {
	payload := []byte(`{
		"object":"template",
		"id":"33333333-3333-4333-8333-333333333333",
		"created_at":"2026-09-10T10:00:00Z",
		"updated_at":"2026-09-10T10:00:00Z",
		"name":"No sender",
		"subject":"Hello",
		"preheader":"",
		"variables":[],
		"from":null,
		"reply_to":null
	}`)

	var template Template
	require.NoError(t, json.Unmarshal(payload, &template))

	assert.Nil(t, template.From)
	assert.Nil(t, template.ReplyTo)
}

func TestTemplateDecodesSenderWithoutName(t *testing.T) {
	payload := []byte(`{
		"object":"template",
		"id":"44444444-4444-4444-8444-444444444444",
		"created_at":"2026-09-10T10:00:00Z",
		"updated_at":"2026-09-10T10:00:00Z",
		"name":"Sender without a name",
		"subject":"Hello",
		"preheader":"",
		"variables":[],
		"from":{"email":"hello@example.com","name":""},
		"reply_to":null
	}`)

	var template Template
	require.NoError(t, json.Unmarshal(payload, &template))

	require.NotNil(t, template.From)
	assert.Equal(t, "hello@example.com", template.From.Email)
	require.NotNil(t, template.From.Name)
	assert.Empty(t, *template.From.Name)
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
		"variables":[],
		"editor":"simple",
		"has_draft":true
	}`)

	var template Template
	require.NoError(t, json.Unmarshal(payload, &template))

	assert.Equal(t, "Draft", template.Name)
	assert.Nil(t, template.Content)
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
	assert.Equal(t, TemplateVariableKindValue, page.Data[0].Variables[0].Kind)
	assert.Equal(t, TemplateVariableKindList, page.Data[0].Variables[1].Kind)
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

func TestTemplateDecodesEachContentShape(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    *TemplateContent
	}{
		{name: "never published", content: `null`},
		{name: "only a custom text", content: `{"text":"Hi","text_is_custom":true}`, want: &TemplateContent{Text: "Hi", TextIsCustom: true}},
		{name: "html template", content: `{"html":"<p>Hi</p>","text":"Hi","text_is_custom":false}`, want: &TemplateContent{HTML: "<p>Hi</p>", Text: "Hi"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var template Template
			require.NoError(t, json.Unmarshal([]byte(`{"object":"template","editor":"html","has_draft":false,"content":`+tt.content+`}`), &template))

			assert.Equal(t, tt.want, template.Content)
		})
	}
}

func TestTemplateVersionsDecodeEachPublisher(t *testing.T) {
	userID := uuid.MustParse("33333333-3333-4333-8333-333333333333")
	keyID := uuid.MustParse("44444444-4444-4444-8444-444444444444")
	payload := []byte(`{"object":"list","data":[
		{"object":"template_version","id":"55555555-5555-4555-8555-555555555555","version":3,"published_at":"2026-09-12T10:00:00Z","published_by":{"type":"api_key","id":"44444444-4444-4444-8444-444444444444"}},
		{"object":"template_version","id":"66666666-6666-4666-8666-666666666666","version":2,"published_at":"2026-09-11T10:00:00Z","published_by":{"type":"user","id":"33333333-3333-4333-8333-333333333333"}},
		{"object":"template_version","id":"77777777-7777-4777-8777-777777777777","version":1,"published_at":"2026-09-10T10:00:00Z","published_by":null}
	]}`)

	var versions TemplateVersionsResponse
	require.NoError(t, json.Unmarshal(payload, &versions))

	assert.Equal(t, "list", versions.Object)
	require.Len(t, versions.Data, 3)
	assert.Equal(t, 3, versions.Data[0].Version)
	assert.Equal(t, &TemplatePublisher{Type: TemplatePublisherTypeAPIKey, ID: keyID}, versions.Data[0].PublishedBy)
	assert.Equal(t, &TemplatePublisher{Type: TemplatePublisherTypeUser, ID: userID}, versions.Data[1].PublishedBy)
	assert.Nil(t, versions.Data[2].PublishedBy)
	assert.Equal(t, time.Date(2026, 9, 10, 10, 0, 0, 0, time.UTC), versions.Data[2].PublishedAt.UTC())
}

func TestTemplateVersionDetailDecodesAPIFixture(t *testing.T) {
	payload := []byte(`{
		"object":"template_version",
		"id":"55555555-5555-4555-8555-555555555555",
		"version":3,
		"published_at":"2026-09-12T10:00:00Z",
		"published_by":{"type":"user","id":"33333333-3333-4333-8333-333333333333"},
		"subject":"Reset your password",
		"preheader":"It expires in an hour",
		"variables":[],
		"from":{"email":"hello@example.com","name":""},
		"reply_to":null,
		"content":{"mjml":"<mjml></mjml>","html":"<html></html>","text":"Reset it","text_is_custom":false}
	}`)

	var version TemplateVersionDetail
	require.NoError(t, json.Unmarshal(payload, &version))

	assert.Equal(t, 3, version.Version)
	require.NotNil(t, version.PublishedBy)
	assert.Equal(t, TemplatePublisherTypeUser, version.PublishedBy.Type)
	assert.Equal(t, "Reset your password", version.Subject)
	assert.Equal(t, "It expires in an hour", version.Preheader)
	require.NotNil(t, version.From)
	assert.Equal(t, "hello@example.com", version.From.Email)
	assert.Nil(t, version.ReplyTo)
	require.NotNil(t, version.Content)
	assert.Equal(t, "<mjml></mjml>", version.Content.MJML)
}

func TestTemplateVariableDecodesEachKind(t *testing.T) {
	tests := []struct {
		name    string
		payload string
		want    TemplateVariable
	}{
		{
			name:    "value",
			payload: `{"name":"first_name","required":true,"kind":"value"}`,
			want:    TemplateVariable{Name: "first_name", Required: true, Kind: TemplateVariableKindValue},
		},
		{
			name:    "condition",
			payload: `{"name":"vip","required":false,"kind":"condition"}`,
			want:    TemplateVariable{Name: "vip", Kind: TemplateVariableKindCondition},
		},
		{
			name:    "list with item fields",
			payload: `{"name":"items","required":true,"kind":"list","item_fields":["name","price"]}`,
			want:    TemplateVariable{Name: "items", Required: true, Kind: TemplateVariableKindList, ItemFields: []string{"name", "price"}},
		},
		{
			name:    "list with empty item fields",
			payload: `{"name":"rows","required":true,"kind":"list","item_fields":[]}`,
			want:    TemplateVariable{Name: "rows", Required: true, Kind: TemplateVariableKindList, ItemFields: []string{}},
		},
		{
			name:    "list without item fields",
			payload: `{"name":"rows","required":true,"kind":"list"}`,
			want:    TemplateVariable{Name: "rows", Required: true, Kind: TemplateVariableKindList},
		},
		{
			name:    "kind the SDK does not know",
			payload: `{"name":"when","required":true,"kind":"date","item_fields":["day"]}`,
			want:    TemplateVariable{Name: "when", Required: true, Kind: "date", ItemFields: []string{"day"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var variable TemplateVariable
			require.NoError(t, json.Unmarshal([]byte(tt.payload), &variable))

			assert.Equal(t, tt.want, variable)
		})
	}
}

// simpleTemplateWriteFixture is the answer to a write that carried markup.
const simpleTemplateWriteFixture = `{
	"object":"template",
	"id":"55555555-5555-4555-8555-555555555555",
	"created_at":"2026-09-10T10:00:00Z",
	"updated_at":"2026-09-10T11:00:00Z",
	"name":"Welcome",
	"subject":"Welcome",
	"preheader":"",
	"variables":[],
	"editor":"simple",
	"has_draft":false,
	"content":{"html":"<p>Price: {{ price }}</p>","text":"Price: {{ price }}","text_is_custom":false,"markup":"<Email version=\"1\">\n  <Paragraph>Price: {{ price }}</Paragraph>\n</Email>\n"},
	"markup_warnings":[{"code":"jinja_in_text","message":"Text holds {{, which stays text","line":2,"column":22,"path":"Email > Paragraph[1]"}]
}`

func TestTemplateDecodesMarkupAndWarnings(t *testing.T) {
	var template Template
	require.NoError(t, json.Unmarshal([]byte(simpleTemplateWriteFixture), &template))

	require.NotNil(t, template.Content)
	require.NotNil(t, template.Content.Markup)
	assert.Equal(t, "<Email version=\"1\">\n  <Paragraph>Price: {{ price }}</Paragraph>\n</Email>\n", *template.Content.Markup)
	assert.Nil(t, template.Content.MarkupError)
	assert.Equal(t, []common.MarkupProblem{{
		Code:    common.MarkupWarningJinjaInText,
		Message: "Text holds {{, which stays text",
		Line:    2,
		Column:  22,
		Path:    "Email > Paragraph[1]",
	}}, template.MarkupWarnings)

	encoded, err := json.Marshal(template)
	require.NoError(t, err)
	assert.JSONEq(t, simpleTemplateWriteFixture, string(encoded))
}

func TestTemplateContentDecodesNullMarkupWithItsError(t *testing.T) {
	payload := []byte(`{"html":"<p>Hi</p>","text":"Hi","text_is_custom":false,"markup":null,"markup_error":{"code":"html_block","message":"The design holds an HTML block"}}`)

	var content TemplateContent
	require.NoError(t, json.Unmarshal(payload, &content))

	assert.Nil(t, content.Markup)
	require.NotNil(t, content.MarkupError)
	assert.Equal(t, MarkupError{Code: MarkupErrorHTMLBlock, Message: "The design holds an HTML block"}, *content.MarkupError)

	// A nil Markup is left out, so the null does not come back, but the
	// content decodes again to the same value.
	encoded, err := json.Marshal(content)
	require.NoError(t, err)
	assert.JSONEq(t, `{"html":"<p>Hi</p>","text":"Hi","text_is_custom":false,"markup_error":{"code":"html_block","message":"The design holds an HTML block"}}`, string(encoded))
	var reDecoded TemplateContent
	require.NoError(t, json.Unmarshal(encoded, &reDecoded))
	assert.Equal(t, content, reDecoded)
}

func TestTemplateContentOfOtherEditorsHasNoMarkup(t *testing.T) {
	var template Template
	require.NoError(t, json.Unmarshal([]byte(templateFixture), &template))

	require.NotNil(t, template.Content)
	assert.Nil(t, template.Content.Markup)
	assert.Nil(t, template.Content.MarkupError)
	assert.Nil(t, template.MarkupWarnings)
}

func TestTemplateDecodesEmptyMarkupWarnings(t *testing.T) {
	var template Template
	require.NoError(t, json.Unmarshal([]byte(`{"object":"template","editor":"simple","markup_warnings":[]}`), &template))

	// An empty list tells a write that carried markup apart from any other
	// answer, which has no markup_warnings.
	assert.NotNil(t, template.MarkupWarnings)
	assert.Empty(t, template.MarkupWarnings)
}

func TestTemplateVersionDecodesMarkup(t *testing.T) {
	content := `{"html":"<p>Hi</p>","text":"Hi","text_is_custom":false,"markup":"<Email version=\"1\"><Paragraph>Hi</Paragraph></Email>"}`

	var version TemplateVersionDetail
	require.NoError(t, json.Unmarshal([]byte(`{"object":"template_version","content":`+content+`}`), &version))
	require.NotNil(t, version.Content)
	require.NotNil(t, version.Content.Markup)
	assert.Equal(t, `<Email version="1"><Paragraph>Hi</Paragraph></Email>`, *version.Content.Markup)
}
