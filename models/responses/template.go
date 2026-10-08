package responses

import (
	"time"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/google/uuid"
)

// TemplateVariable represents one variable a template's design uses.
type TemplateVariable struct {
	Name string `json:"name"`
	// Required reports whether a send has to supply a value for this
	// variable. A variable the design wraps in a fallback is not required,
	// and neither are the names the send supplies itself.
	Required bool `json:"required"`
}

// TemplateContent is the content of a template, a draft or a version. A
// field that does not apply to the template's editor is empty, and so are
// MJML, HTML and Text when there is nothing to show.
type TemplateContent struct {
	// MJML is the MJML source of an advanced template.
	MJML string `json:"mjml,omitempty"`
	// HTML is the HTML of an html template, or the HTML made from the design
	// of a simple template.
	HTML string `json:"html,omitempty"`
	// Text is the plain text version.
	Text string `json:"text,omitempty"`
	// TextIsCustom reports whether the text was written rather than made from
	// the HTML. A text made from the HTML is made again when the design
	// changes.
	TextIsCustom bool `json:"text_is_custom"`
}

// Template represents an account's transactional template. Its fields are
// those of the published copy, which sends use.
type Template struct {
	Object    string             `json:"object"`
	ID        uuid.UUID          `json:"id"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	Name      string             `json:"name"`
	Subject   string             `json:"subject"`
	Preheader string             `json:"preheader"`
	Variables []TemplateVariable `json:"variables"`
	// From is the default sender, used by a send that names no sender. It is
	// nil when the template has none, and then every send must name one.
	From *common.SenderAddress `json:"from,omitempty"`
	// ReplyTo is the default reply-to address, used by a send that sets no
	// reply-to, also one that names its own sender. It is nil when the template
	// has none, and its Name is always empty.
	ReplyTo *common.SenderAddress `json:"reply_to,omitempty"`
	// Editor is common.TemplateEditorAdvanced, common.TemplateEditorSimple or
	// common.TemplateEditorHTML.
	Editor string `json:"editor"`
	// HasDraft reports whether the template has changes that are not
	// published yet. GetTemplateDraft returns them.
	HasDraft bool `json:"has_draft"`
	// Content is the published content. It is nil when the published copy has
	// no body, and in the items of GetTemplates.
	Content *TemplateContent `json:"content,omitempty"`
}

// TemplateDraft represents a template's draft: the changes saved in the
// dashboard or through the API that are not published yet.
type TemplateDraft struct {
	Object     string             `json:"object"`
	TemplateID uuid.UUID          `json:"template_id"`
	UpdatedAt  time.Time          `json:"updated_at"`
	Subject    string             `json:"subject"`
	Preheader  string             `json:"preheader"`
	Variables  []TemplateVariable `json:"variables"`
	// From is the draft's default sender, nil when it has none.
	From *common.SenderAddress `json:"from,omitempty"`
	// ReplyTo is the draft's default reply-to address, nil when it has none.
	ReplyTo *common.SenderAddress `json:"reply_to,omitempty"`
	// Content is the draft's content, nil when the draft has no body.
	Content *TemplateContent `json:"content,omitempty"`
}

// Publisher types of a template version.
const (
	TemplatePublisherTypeUser   = "user"
	TemplatePublisherTypeAPIKey = "api_key"
)

// TemplatePublisher tells who published a template version.
type TemplatePublisher struct {
	// Type is TemplatePublisherTypeUser for a version published in the
	// dashboard, or TemplatePublisherTypeAPIKey for one published through the
	// API.
	Type string `json:"type"`
	// ID is the user's ID, as GetAccountMembers shows it, or the API key's ID.
	ID uuid.UUID `json:"id"`
}

// TemplateVersion represents one published version of a template, without
// its content.
type TemplateVersion struct {
	Object string    `json:"object"`
	ID     uuid.UUID `json:"id"`
	// Version is the version number, counted from 1 for each template.
	Version     int       `json:"version"`
	PublishedAt time.Time `json:"published_at"`
	// PublishedBy is nil when the publisher is not known, or was deleted.
	PublishedBy *TemplatePublisher `json:"published_by,omitempty"`
}

// TemplateVersionDetail represents one published version of a template with
// its content.
type TemplateVersionDetail struct {
	Object string    `json:"object"`
	ID     uuid.UUID `json:"id"`
	// Version is the version number, counted from 1 for each template.
	Version     int       `json:"version"`
	PublishedAt time.Time `json:"published_at"`
	// PublishedBy is nil when the publisher is not known, or was deleted.
	PublishedBy *TemplatePublisher `json:"published_by,omitempty"`
	Subject     string             `json:"subject"`
	Preheader   string             `json:"preheader"`
	Variables   []TemplateVariable `json:"variables"`
	// From is the version's default sender, nil when it has none.
	From *common.SenderAddress `json:"from,omitempty"`
	// ReplyTo is the version's default reply-to address, nil when it has none.
	ReplyTo *common.SenderAddress `json:"reply_to,omitempty"`
	// Content is the version's content, nil when the version has no body.
	Content *TemplateContent `json:"content,omitempty"`
}

// TemplateVersionsResponse lists a template's versions, newest first. A
// template keeps its 50 most recent versions, so the list is not paginated.
type TemplateVersionsResponse struct {
	Object string            `json:"object"`
	Data   []TemplateVersion `json:"data"`
}

// PaginatedTemplatesResponse represents a page of transactional templates.
type PaginatedTemplatesResponse struct {
	Object     string                `json:"object"`
	Data       []Template            `json:"data"`
	Pagination common.PaginationInfo `json:"pagination"`
}
