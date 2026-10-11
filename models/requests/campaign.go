package requests

import (
	"encoding/json"
	"time"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/google/uuid"
)

// ListCampaignsParams represents the filter and pagination cursors accepted
// by the campaigns list endpoint.
type ListCampaignsParams struct {
	// Status lists only the campaigns in this status, one of the
	// common.CampaignStatus constants.
	Status *string
	Limit  *int32
	After  *string
	Before *string
}

// CreateCampaignRequest represents a request to create a draft campaign.
// Every field but Name can be set later with UpdateCampaign. The design comes
// from at most one source: Editor, Content, both, or TemplateID. With none of
// them the campaign has no design yet.
type CreateCampaignRequest struct {
	// Name is the campaign name the dashboard shows. Recipients do not see it.
	Name string `json:"name"`
	// Subject is the subject line. Line breaks become spaces and spaces around
	// it are dropped; an empty subject is no subject.
	Subject *string `json:"subject,omitempty"`
	// Preheader is the preview text, cleaned as Subject is.
	Preheader *string `json:"preheader,omitempty"`
	// From is the sender. Its domain must be one of the account's domains with
	// valid DNS records. A paused domain is accepted here and refused when the
	// campaign is sent.
	From *common.SenderAddress `json:"from,omitempty"`
	// ReplyTo is the reply-to address, an address only, without a display
	// name.
	ReplyTo *string `json:"reply_to,omitempty"`
	// ListIDs are the lists to send to, in order: 1 to 50 of the account's
	// lists. A list named twice counts once.
	ListIDs []uuid.UUID `json:"list_ids,omitempty"`
	// Editor is the design's editor, common.TemplateEditorAdvanced,
	// common.TemplateEditorSimple or common.TemplateEditorHTML. Without
	// Content the campaign gets an empty design of this editor, which the
	// dashboard's editor opens. With Content, it is required when the content
	// holds neither MJML nor HTML, and must agree with it otherwise. It cannot
	// be set with TemplateID.
	Editor *string `json:"editor,omitempty"`
	// Content is the design's content, by the rules of TemplateContentInput.
	// Where those rules name the draft, they mean the campaign's design. It
	// cannot be set with TemplateID.
	Content *TemplateContentInput `json:"content,omitempty"`
	// TemplateID is a campaign template whose design the campaign copies. The
	// copy is the campaign's own: later changes to the template do not reach
	// it. A campaign template with no HTML is refused.
	TemplateID *uuid.UUID `json:"template_id,omitempty"`
}

// UpdateCampaignRequest represents a partial update to a draft campaign. A
// nil field is not changed or checked. The design's editor cannot change.
type UpdateCampaignRequest struct {
	// Name is the campaign name.
	Name *string `json:"name,omitempty"`
	// Subject is the subject line, cleaned as on create. A pointer to ""
	// clears it.
	Subject *string `json:"subject,omitempty"`
	// Preheader is the preview text, cleaned as on create. A pointer to ""
	// clears it.
	Preheader *string `json:"preheader,omitempty"`
	// From is the sender, as on create. A pointer to an empty
	// common.SenderAddress clears it; the reply-to is kept.
	From *common.SenderAddress `json:"from,omitempty"`
	// ReplyTo is the reply-to address, as on create. A pointer to "" clears
	// it.
	ReplyTo *string `json:"reply_to,omitempty"`
	// ListIDs replace the campaign's lists, as on create. A non-nil empty
	// slice is sent, and the API refuses it.
	ListIDs []uuid.UUID `json:"list_ids,omitempty"`
	// Content changes the design's content, by the rules of
	// TemplateContentInput, in the design's editor. For a simple design, send
	// only Text. A campaign with no design yet needs MJML or HTML, which sets
	// its editor. Leave it nil to keep the content.
	Content *TemplateContentInput `json:"content,omitempty"`
}

// MarshalJSON sends a Subject, Preheader or ReplyTo that points at "", and a
// From that points at an empty common.SenderAddress, as null, which clears
// it. It sends ListIDs whenever it is not nil, also when it is empty.
func (r UpdateCampaignRequest) MarshalJSON() ([]byte, error) {
	wire := struct {
		Name      *string               `json:"name,omitempty"`
		Subject   json.RawMessage       `json:"subject,omitempty"`
		Preheader json.RawMessage       `json:"preheader,omitempty"`
		From      json.RawMessage       `json:"from,omitempty"`
		ReplyTo   json.RawMessage       `json:"reply_to,omitempty"`
		ListIDs   *[]uuid.UUID          `json:"list_ids,omitempty"`
		Content   *TemplateContentInput `json:"content,omitempty"`
	}{
		Name:    r.Name,
		Content: r.Content,
	}
	if r.ListIDs != nil {
		wire.ListIDs = &r.ListIDs
	}

	var err error
	if wire.Subject, err = encodeString(r.Subject); err != nil {
		return nil, err
	}
	if wire.Preheader, err = encodeString(r.Preheader); err != nil {
		return nil, err
	}
	if wire.From, err = encodeAddress(r.From); err != nil {
		return nil, err
	}
	if wire.ReplyTo, err = encodeString(r.ReplyTo); err != nil {
		return nil, err
	}

	return json.Marshal(wire)
}

// DuplicateCampaignRequest represents the options of a campaign duplicate.
type DuplicateCampaignRequest struct {
	// Name is the copy's name. Leave it nil to name the copy "Copy of" and
	// the source's name, cut to 255 characters.
	Name *string `json:"name,omitempty"`
}

// SendCampaignRequest represents when to send a campaign.
type SendCampaignRequest struct {
	// SendAt is the time to send, at least 1 minute ahead. Leave it nil to
	// send now.
	SendAt *time.Time `json:"send_at,omitempty"`
}

// TestCampaignRequest represents the addresses to send a test of a campaign
// to.
type TestCampaignRequest struct {
	// To holds 1 to 5 addresses, counted after addresses that differ only in
	// case are dropped. An entry may carry a display name, such as
	// "Ann <ann@example.com>"; the test goes to the address alone.
	To []string `json:"to,omitempty"`
}

// ListCampaignRecipientsParams represents the filter and pagination cursors
// accepted by the campaign recipients list endpoint. Set exactly one of
// Status and BounceClassification.
type ListCampaignRecipientsParams struct {
	// Status is a key of responses.CampaignStatsReport.ByStatus, such as
	// "Queued", "Delivered" or "Administratively Bounced".
	Status *string
	// BounceClassification is a key of
	// responses.CampaignStatsReport.ByBounceClassification other than
	// "Unknown", which has no recipient page.
	BounceClassification *string
	Limit                *int32
	After                *string
	Before               *string
}
