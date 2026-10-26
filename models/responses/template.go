package responses

import (
	"time"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/google/uuid"
)

// Kinds of a template variable. They say what shape of value a send gives
// for the variable in a recipient's Substitutions.
//
//   - TemplateVariableKindValue: text or a number that appears in the email,
//     such as a first name.
//   - TemplateVariableKindCondition: used only to decide whether a part of
//     the email is shown. It never appears in the email. A value that is not
//     empty, such as true or "yes", makes it set. Leaving it out or sending
//     "", false, 0, nil, an empty slice or an empty map makes it not set. The
//     strings "false" and "0" are not empty, so they make it set. When the
//     design compares it, for example "show this section if plan equals
//     gold", send a value that meets the rule to show the section.
//   - TemplateVariableKindList: a slice. A section of the email is repeated
//     once for each element, for example one row per product in an order.
//
// Kind is a plain string, so a kind this version of the SDK does not know
// still decodes.
const (
	TemplateVariableKindValue     = "value"
	TemplateVariableKindCondition = "condition"
	TemplateVariableKindList      = "list"
)

// TemplateVariable represents one variable a template's design uses.
type TemplateVariable struct {
	Name string `json:"name"`
	// Required reports whether a send has to supply a value for this
	// variable. A variable the design wraps in a fallback is not required,
	// and neither are the names the send supplies itself.
	Required bool `json:"required"`
	// Kind is TemplateVariableKindValue, TemplateVariableKindCondition or
	// TemplateVariableKindList.
	Kind string `json:"kind"`
	// ItemFields is set only for a TemplateVariableKindList variable. It
	// names the fields the repeated section uses from each element: send
	// each element as a map with these fields. A field an element does not
	// have is taken from the recipient's substitution of the same name, if
	// there is one. When ItemFields is ["value"], the section shows each
	// element as it is, so plain values such as []string{"new", "sale"}
	// work. It is empty when the section uses nothing from its elements:
	// send one element, of any kind, for each time the section repeats.
	ItemFields []string `json:"item_fields,omitempty"`
}

// Codes of a MarkupError. They say why the design of a simple template could
// not be given as AhaSend Markup (AhaML).
//
//   - MarkupErrorNoDesign: the design has no blocks yet.
//   - MarkupErrorHTMLBlock: the design holds an HTML block, which AhaML cannot
//     show. Remove the block in the dashboard to read and write the design as
//     AhaML.
//   - MarkupErrorUnmappedBlock: the design holds a block or a mark that AhaML
//     has no tag for. The message names its type.
//   - MarkupErrorBadStoredValue: the design holds a value that AhaML cannot
//     write. The message names the block and the attribute.
//   - MarkupErrorRendererUnavailable: AhaML cannot be made now. Read the
//     design again later.
//   - MarkupErrorRateLimited: the account reads and writes AhaML more often
//     than its limit allows. Read the design again in a second.
//
// Code is a plain string, so a code this version of the SDK does not know
// still decodes.
const (
	MarkupErrorNoDesign            = "no_design"
	MarkupErrorHTMLBlock           = "html_block"
	MarkupErrorUnmappedBlock       = "unmapped_block"
	MarkupErrorBadStoredValue      = "bad_stored_value"
	MarkupErrorRendererUnavailable = "renderer_unavailable"
	MarkupErrorRateLimited         = "rate_limited"
)

// MarkupError tells why TemplateContent.Markup is nil for a simple design.
type MarkupError struct {
	// Code is one of the MarkupError constants.
	Code string `json:"code"`
	// Message says what went wrong, in words for a person.
	Message string `json:"message"`
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
	// Markup is the design of a simple template as AhaSend Markup (AhaML),
	// which TemplateContentInput.Markup writes back. It is nil for other
	// editors, and for a simple design when AhaML cannot be made; MarkupError
	// then says why.
	Markup *string `json:"markup,omitempty"`
	// MarkupError tells why Markup is nil for a simple design. It is nil when
	// Markup is set, and for other editors.
	MarkupError *MarkupError `json:"markup_error,omitempty"`
}

// Template represents an account's transactional template. GetTemplate and
// the writes answer its working copy, the draft while one exists;
// GetPublishedTemplate and GetTemplates answer its published copy, which sends
// use.
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
	// published yet. GetTemplate returns them; PublishTemplate publishes them
	// and DiscardTemplateDraft drops them.
	HasDraft bool `json:"has_draft"`
	// Content is this copy's content. It is nil when the copy has no body, as
	// for the published copy of a template that was never published, and in
	// the items of GetTemplates.
	Content *TemplateContent `json:"content,omitempty"`
	// MarkupWarnings are the warnings about the AhaML a write stored, up to
	// the first 100. Only the answer to a write that carried
	// TemplateContentInput.Markup sets it, as an empty slice when there are
	// no warnings; it is nil in every other answer.
	MarkupWarnings []common.MarkupProblem `json:"markup_warnings,omitempty"`
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
	// Variables are every variable the version uses, as saved when the
	// version was published. A version published before kinds were recorded
	// shows every variable as TemplateVariableKindValue, with no ItemFields;
	// save and publish the template again to get a version that records them.
	Variables []TemplateVariable `json:"variables"`
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
