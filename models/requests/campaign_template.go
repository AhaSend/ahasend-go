package requests

// ListCampaignTemplatesParams represents the pagination cursors accepted by
// the campaign templates list endpoint.
type ListCampaignTemplatesParams struct {
	Limit  *int32
	After  *string
	Before *string
}

// CreateCampaignTemplateRequest represents a request to create a campaign
// template. A campaign template has no draft: what the request sets is what a
// campaign made from the template copies.
type CreateCampaignTemplateRequest struct {
	// Name is the campaign template name.
	Name string `json:"name"`
	// Editor is common.TemplateEditorAdvanced, common.TemplateEditorSimple or
	// common.TemplateEditorHTML. It is required when Content holds none of
	// Markup, MJML and HTML: Markup means simple, also beside HTML; MJML means
	// advanced and HTML alone means html. Without Content the campaign
	// template gets an empty design of this editor.
	Editor *string `json:"editor,omitempty"`
	// Content is the content, by the rules of TemplateContentInput. Where those
	// rules name the draft, they mean the campaign template itself.
	Content *TemplateContentInput `json:"content,omitempty"`
}

// UpdateCampaignTemplateRequest represents a partial update to a campaign
// template. A nil field is not changed or checked. The change applies at
// once; campaigns already made from the template keep their own copies. The
// editor cannot change.
type UpdateCampaignTemplateRequest struct {
	// Name is the campaign template name.
	Name *string `json:"name,omitempty"`
	// Content changes the content, by the rules of TemplateContentInput. A
	// simple campaign template's design is written with Markup. Leave it nil
	// to keep the content.
	Content *TemplateContentInput `json:"content,omitempty"`
}
