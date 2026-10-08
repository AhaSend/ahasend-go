package requests

import (
	"encoding/json"

	"github.com/AhaSend/ahasend-go/models/common"
)

// GetTemplatesParams represents the pagination cursors accepted by the
// transactional templates list endpoint.
type GetTemplatesParams struct {
	Limit  *int32
	After  *string
	Before *string
}

// TemplateContentInput is the content of a template write. It holds at most
// one of MJML and HTML, and at least one of MJML, HTML and Text.
//
// The content of a draft, or of a template with no draft, can be sent back
// with its non-empty fields set: an MJML or HTML equal to the stored one keeps
// the stored HTML. While
// a draft exists, a template's content is the published one, and sending it
// back replaces the draft's design, so leave Content nil when it does not
// change. For a simple template, send only Text: its HTML can be changed only
// in the dashboard.
type TemplateContentInput struct {
	// MJML is the MJML source, only for advanced templates. It is compiled in
	// strict mode. Images must use full URLs, such as https://.
	MJML *string `json:"mjml,omitempty"`
	// HTML is the HTML, only for html templates. Images and stylesheets must
	// use full URLs, such as https://.
	HTML *string `json:"html,omitempty"`
	// Text is the plain text version. Beside a new MJML or HTML, a nil Text
	// makes the text from it; otherwise a nil Text keeps the text. Point it at
	// "" to make the text again from the HTML. Any other value sets the text,
	// except that beside a new MJML or HTML a text equal to the text made from
	// the old HTML is made again from the new one.
	Text *string `json:"text,omitempty"`
}

// MarshalJSON sends a Text that points at "" as null, which makes the text
// again from the HTML.
func (c TemplateContentInput) MarshalJSON() ([]byte, error) {
	wire := struct {
		MJML *string         `json:"mjml,omitempty"`
		HTML *string         `json:"html,omitempty"`
		Text json.RawMessage `json:"text,omitempty"`
	}{MJML: c.MJML, HTML: c.HTML}

	var err error
	if wire.Text, err = encodeText(c.Text); err != nil {
		return nil, err
	}

	return json.Marshal(wire)
}

// CreateTemplateRequest represents a request to create a transactional
// template. Its fields go to the template's draft; set Publish to publish the
// draft in the same request, so sends use it.
type CreateTemplateRequest struct {
	// Name is the template name.
	Name string `json:"name"`
	// Editor is common.TemplateEditorAdvanced, common.TemplateEditorSimple or
	// common.TemplateEditorHTML. It is required when Content holds neither
	// MJML nor HTML: MJML means advanced and HTML means html.
	Editor *string `json:"editor,omitempty"`
	// Subject is the subject.
	Subject *string `json:"subject,omitempty"`
	// Preheader is the preview text.
	Preheader *string `json:"preheader,omitempty"`
	// From is the default sender. Its domain must be one of the account's
	// domains with valid DNS records that is not paused. Leave it nil for a
	// template with no sender.
	From *common.SenderAddress `json:"from,omitempty"`
	// ReplyTo is the default reply-to address. It needs a From, and its Name
	// must be empty.
	ReplyTo *common.SenderAddress `json:"reply_to,omitempty"`
	// Content is the template's content.
	Content *TemplateContentInput `json:"content,omitempty"`
	// Publish publishes the draft once it is written.
	Publish bool `json:"publish,omitempty"`
}

// UpdateTemplateRequest represents a partial update to a transactional
// template. A nil field is not changed or checked. Every field but Name goes
// to the template's draft, the same draft the dashboard edits; Name changes
// at once. The editor cannot change.
type UpdateTemplateRequest struct {
	// Name is the template name.
	Name *string `json:"name,omitempty"`
	// Subject is the subject. A pointer to "" clears it.
	Subject *string `json:"subject,omitempty"`
	// Preheader is the preview text. A pointer to "" clears it.
	Preheader *string `json:"preheader,omitempty"`
	// From is the default sender. It replaces the address and the name
	// together, so a nil Name clears the name; the reply-to is kept. A pointer
	// to an empty common.SenderAddress clears the sender and the reply-to.
	From *common.SenderAddress `json:"from,omitempty"`
	// ReplyTo is the default reply-to address. Its Name must be empty. A
	// pointer to an empty common.SenderAddress clears it.
	ReplyTo *common.SenderAddress `json:"reply_to,omitempty"`
	// Content changes the template's content. Leave it nil to keep it.
	Content *TemplateContentInput `json:"content,omitempty"`
	// Publish publishes the draft once it is written, including changes made
	// in the dashboard. With no other field, the request only publishes.
	Publish bool `json:"publish,omitempty"`
}

// MarshalJSON sends a From or ReplyTo that points at an empty
// common.SenderAddress as null, which clears it.
func (r UpdateTemplateRequest) MarshalJSON() ([]byte, error) {
	wire := struct {
		Name      *string               `json:"name,omitempty"`
		Subject   *string               `json:"subject,omitempty"`
		Preheader *string               `json:"preheader,omitempty"`
		From      json.RawMessage       `json:"from,omitempty"`
		ReplyTo   json.RawMessage       `json:"reply_to,omitempty"`
		Content   *TemplateContentInput `json:"content,omitempty"`
		Publish   bool                  `json:"publish,omitempty"`
	}{
		Name:      r.Name,
		Subject:   r.Subject,
		Preheader: r.Preheader,
		Content:   r.Content,
		Publish:   r.Publish,
	}

	var err error
	if wire.From, err = encodeAddress(r.From); err != nil {
		return nil, err
	}
	if wire.ReplyTo, err = encodeAddress(r.ReplyTo); err != nil {
		return nil, err
	}

	return json.Marshal(wire)
}

// RestoreTemplateVersionRequest represents the options of a version restore.
type RestoreTemplateVersionRequest struct {
	// Publish publishes the draft once the version is restored into it.
	Publish bool `json:"publish,omitempty"`
}

// encodeText leaves a nil text out, sends a text that points at "" as null,
// and any other text as itself.
func encodeText(text *string) (json.RawMessage, error) {
	switch {
	case text == nil:
		return nil, nil
	case *text == "":
		return json.RawMessage("null"), nil
	}
	return json.Marshal(*text)
}

// encodeAddress leaves a nil address out, sends an empty address as null, and
// any other address as itself.
func encodeAddress(address *common.SenderAddress) (json.RawMessage, error) {
	switch {
	case address == nil:
		return nil, nil
	case *address == (common.SenderAddress{}):
		return json.RawMessage("null"), nil
	}
	return json.Marshal(address)
}
