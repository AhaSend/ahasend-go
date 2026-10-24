package responses

import (
	"time"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/google/uuid"
)

// CampaignTemplate represents a design a campaign copies when it is made from
// it. A campaign template has no draft, no versions, no subject and no
// sender; the campaign owns its subject and sender.
type CampaignTemplate struct {
	Object    string    `json:"object"`
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Name is the campaign template name the dashboard shows.
	Name string `json:"name"`
	// Editor is common.TemplateEditorAdvanced, common.TemplateEditorSimple or
	// common.TemplateEditorHTML.
	Editor string `json:"editor"`
	// Content is the content. It is nil when the campaign template has
	// neither HTML nor text, and in the items of ListCampaignTemplates. A
	// campaign can be made from the campaign template only once it has HTML.
	Content *TemplateContent `json:"content,omitempty"`
	// MarkupWarnings are the warnings about the AhaML a write stored, as on
	// Template.
	MarkupWarnings []common.MarkupProblem `json:"markup_warnings,omitempty"`
}

// PaginatedCampaignTemplatesResponse represents a page of campaign
// templates, each without its content.
type PaginatedCampaignTemplatesResponse struct {
	Object     string                `json:"object"`
	Data       []CampaignTemplate    `json:"data"`
	Pagination common.PaginationInfo `json:"pagination"`
}
