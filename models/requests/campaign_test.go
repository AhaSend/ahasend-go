package requests

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCampaignRequestsEncodeOnlyNamedFields(t *testing.T) {
	empty := ""
	name := "October newsletter"
	subject := "Our news"
	replyTo := "support@example.com"
	html := "<p>News</p>"
	editor := common.TemplateEditorAdvanced
	listID := uuid.MustParse("66666666-6666-4666-8666-666666666666")
	templateID := uuid.MustParse("44444444-4444-4444-8444-444444444444")
	sendAt := time.Date(2026, 10, 20, 9, 0, 0, 0, time.UTC)

	tests := []struct {
		name    string
		request any
		want    string
	}{
		{
			name:    "campaign template create with name and editor",
			request: CreateCampaignTemplateRequest{Name: name, Editor: &editor},
			want:    `{"name":"October newsletter","editor":"advanced"}`,
		},
		{
			name:    "campaign template create with content",
			request: CreateCampaignTemplateRequest{Name: name, Content: &TemplateContentInput{HTML: &html}},
			want:    `{"name":"October newsletter","content":{"html":"<p>News</p>"}}`,
		},
		{
			name:    "campaign template update naming nothing",
			request: UpdateCampaignTemplateRequest{},
			want:    `{}`,
		},
		{
			name:    "campaign template update making the text again",
			request: UpdateCampaignTemplateRequest{Content: &TemplateContentInput{Text: &empty}},
			want:    `{"content":{"text":null}}`,
		},
		{
			name:    "campaign create with only a name",
			request: CreateCampaignRequest{Name: name},
			want:    `{"name":"October newsletter"}`,
		},
		{
			name: "campaign create from a campaign template",
			request: CreateCampaignRequest{
				Name:       name,
				Subject:    &subject,
				From:       &common.SenderAddress{Email: "news@example.com"},
				ReplyTo:    &replyTo,
				ListIDs:    []uuid.UUID{listID},
				TemplateID: &templateID,
			},
			want: `{"name":"October newsletter","subject":"Our news","from":{"email":"news@example.com"},"reply_to":"support@example.com","list_ids":["66666666-6666-4666-8666-666666666666"],"template_id":"44444444-4444-4444-8444-444444444444"}`,
		},
		{
			name:    "campaign create with an editor and content",
			request: CreateCampaignRequest{Name: name, Editor: &editor, Content: &TemplateContentInput{MJML: &html}},
			want:    `{"name":"October newsletter","editor":"advanced","content":{"mjml":"<p>News</p>"}}`,
		},
		{
			name:    "campaign update naming nothing",
			request: UpdateCampaignRequest{},
			want:    `{}`,
		},
		{
			name:    "campaign update clearing subject, preheader, sender and reply-to",
			request: UpdateCampaignRequest{Subject: &empty, Preheader: &empty, From: &common.SenderAddress{}, ReplyTo: &empty},
			want:    `{"subject":null,"preheader":null,"from":null,"reply_to":null}`,
		},
		{
			name:    "campaign update sending a read sender back",
			request: UpdateCampaignRequest{From: &common.SenderAddress{Email: "news@example.com", Name: &empty}},
			want:    `{"from":{"email":"news@example.com","name":""}}`,
		},
		{
			name:    "campaign update setting fields",
			request: UpdateCampaignRequest{Name: &name, Subject: &subject, ReplyTo: &replyTo, ListIDs: []uuid.UUID{listID}, Content: &TemplateContentInput{HTML: &html}},
			want:    `{"name":"October newsletter","subject":"Our news","reply_to":"support@example.com","list_ids":["66666666-6666-4666-8666-666666666666"],"content":{"html":"<p>News</p>"}}`,
		},
		{
			name:    "campaign update sending an empty list of lists",
			request: UpdateCampaignRequest{ListIDs: []uuid.UUID{}},
			want:    `{"list_ids":[]}`,
		},
		{
			name:    "duplicate with the default name",
			request: DuplicateCampaignRequest{},
			want:    `{}`,
		},
		{
			name:    "duplicate with a name",
			request: DuplicateCampaignRequest{Name: &name},
			want:    `{"name":"October newsletter"}`,
		},
		{
			name:    "send now",
			request: SendCampaignRequest{},
			want:    `{}`,
		},
		{
			name:    "send later",
			request: SendCampaignRequest{SendAt: &sendAt},
			want:    `{"send_at":"2026-10-20T09:00:00Z"}`,
		},
		{
			name:    "test with no address",
			request: TestCampaignRequest{},
			want:    `{}`,
		},
		{
			name:    "test",
			request: TestCampaignRequest{To: []string{"Ann <ann@example.com>"}},
			want:    `{"to":["Ann <ann@example.com>"]}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded, err := json.Marshal(tt.request)

			require.NoError(t, err)
			assert.JSONEq(t, tt.want, string(encoded))
		})
	}
}
