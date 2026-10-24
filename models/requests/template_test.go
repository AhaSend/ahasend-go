package requests

import (
	"encoding/json"
	"testing"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTemplateRequestsEncodeOnlyNamedFields(t *testing.T) {
	empty := ""
	name := "Password reset"
	subject := "Reset your password"
	html := "<p>Hi</p>"
	text := "Hi"
	markup := `<Email version="1"><Paragraph>Hi</Paragraph></Email>`
	editor := common.TemplateEditorSimple
	senderName := "Example"

	tests := []struct {
		name    string
		request any
		want    string
	}{
		{
			name:    "create with name and editor",
			request: CreateTemplateRequest{Name: name, Editor: &editor},
			want:    `{"name":"Password reset","editor":"simple"}`,
		},
		{
			name: "create with content and publish",
			request: CreateTemplateRequest{
				Name:    name,
				Subject: &subject,
				From:    &common.SenderAddress{Email: "hello@example.com", Name: &senderName},
				Content: &TemplateContentInput{HTML: &html},
				Publish: true,
			},
			want: `{"name":"Password reset","subject":"Reset your password","from":{"email":"hello@example.com","name":"Example"},"content":{"html":"<p>Hi</p>"},"publish":true}`,
		},
		{
			name:    "update naming nothing",
			request: UpdateTemplateRequest{},
			want:    `{}`,
		},
		{
			name:    "update that only publishes",
			request: UpdateTemplateRequest{Publish: true},
			want:    `{"publish":true}`,
		},
		{
			name:    "update clearing subject and preheader",
			request: UpdateTemplateRequest{Subject: &empty, Preheader: &empty},
			want:    `{"subject":"","preheader":""}`,
		},
		{
			name:    "update clearing sender and reply-to",
			request: UpdateTemplateRequest{From: &common.SenderAddress{}, ReplyTo: &common.SenderAddress{}},
			want:    `{"from":null,"reply_to":null}`,
		},
		{
			name:    "update setting a sender without a name",
			request: UpdateTemplateRequest{From: &common.SenderAddress{Email: "hello@example.com"}},
			want:    `{"from":{"email":"hello@example.com"}}`,
		},
		{
			name:    "update setting a reply-to read back with an empty name",
			request: UpdateTemplateRequest{ReplyTo: &common.SenderAddress{Email: "support@example.com", Name: &empty}},
			want:    `{"reply_to":{"email":"support@example.com","name":""}}`,
		},
		{
			name:    "update setting the text",
			request: UpdateTemplateRequest{Content: &TemplateContentInput{Text: &text}},
			want:    `{"content":{"text":"Hi"}}`,
		},
		{
			name:    "update making the text again",
			request: UpdateTemplateRequest{Content: &TemplateContentInput{Text: &empty}},
			want:    `{"content":{"text":null}}`,
		},
		{
			name:    "update with a design and no text",
			request: UpdateTemplateRequest{Name: &name, Content: &TemplateContentInput{HTML: &html}},
			want:    `{"name":"Password reset","content":{"html":"<p>Hi</p>"}}`,
		},
		{
			name:    "update with markup and text",
			request: UpdateTemplateRequest{Content: &TemplateContentInput{Markup: &markup, Text: &text}},
			want:    `{"content":{"text":"Hi","markup":"<Email version=\"1\"><Paragraph>Hi</Paragraph></Email>"}}`,
		},
		{
			// An empty markup is not AhaML: it is sent, and the API refuses
			// it, rather than taken as no markup.
			name:    "update with empty markup",
			request: UpdateTemplateRequest{Content: &TemplateContentInput{Markup: &empty}},
			want:    `{"content":{"markup":""}}`,
		},
		{
			name:    "create with markup and no editor",
			request: CreateTemplateRequest{Name: name, Content: &TemplateContentInput{Markup: &markup}},
			want:    `{"name":"Password reset","content":{"markup":"<Email version=\"1\"><Paragraph>Hi</Paragraph></Email>"}}`,
		},
		{
			name:    "restore without publishing",
			request: RestoreTemplateVersionRequest{},
			want:    `{}`,
		},
		{
			name:    "restore and publish",
			request: RestoreTemplateVersionRequest{Publish: true},
			want:    `{"publish":true}`,
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
