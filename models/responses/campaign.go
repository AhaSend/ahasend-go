package responses

import (
	"time"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/google/uuid"
)

// Campaign issue codes: the reasons a campaign cannot be sent yet.
//
//   - CampaignIssueSenderMissing: the campaign has no complete sender, or its
//     domain was deleted.
//   - CampaignIssueListsMissing: the campaign has no list or segment that
//     still exists.
//   - CampaignIssueSubjectMissing: the campaign has no subject.
//   - CampaignIssuePreheaderMissing: the campaign has no preview text.
//   - CampaignIssueDesignMissing: the campaign has no design, or its design
//     has no HTML.
//   - CampaignIssueTextBroken: the subject, the preview text or the design
//     holds template syntax that cannot be rendered, so every message would
//     fail.
//   - CampaignIssueTextUnchecked: the subject, the preview text and the design
//     could not be checked right now. Nothing is wrong with them that is
//     known; try again in a moment. Under load it can come and go between
//     reads.
//   - CampaignIssueUnsubscribeLinkMissing: the HTML uses neither
//     {{ unsubscribe_url }} nor {{ unsubscribe_link }}, with or without
//     filters.
const (
	CampaignIssueSenderMissing          = "sender_missing"
	CampaignIssueListsMissing           = "lists_missing"
	CampaignIssueSubjectMissing         = "subject_missing"
	CampaignIssuePreheaderMissing       = "preheader_missing"
	CampaignIssueDesignMissing          = "design_missing"
	CampaignIssueTextBroken             = "text_broken"
	CampaignIssueTextUnchecked          = "text_unchecked"
	CampaignIssueUnsubscribeLinkMissing = "unsubscribe_link_missing"
)

// CampaignIssue represents a reason a campaign cannot be sent yet.
type CampaignIssue struct {
	// Code is one of the CampaignIssue constants.
	Code string `json:"code"`
	// Message says what to do, in the words the dashboard uses.
	Message string `json:"message"`
}

// CampaignListItem represents a campaign as ListCampaigns returns it, without
// its audience, content and issues. A field documented as nil is sent as null.
type CampaignListItem struct {
	Object    string    `json:"object"`
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is when the campaign's own fields last changed.
	UpdatedAt time.Time `json:"updated_at"`
	// Name is the campaign name the dashboard shows. Recipients do not see
	// it.
	Name string `json:"name"`
	// Subject is the subject line, nil until one is set.
	Subject *string `json:"subject"`
	// Preheader is the preview text, nil until one is set.
	Preheader *string `json:"preheader"`
	// From is the sender. It is nil until a complete sender is set, and when
	// its domain was deleted. Its Name points at "" when the campaign has no
	// sender name, so it can be sent back as it is.
	From *common.SenderAddress `json:"from"`
	// ReplyTo is the reply-to address, nil when none is set.
	ReplyTo *string `json:"reply_to"`
	// Status is one of the common.CampaignStatus constants.
	Status string `json:"status"`
	// StatusReason tells why the campaign is in its status, when there is a
	// reason to give, and is nil otherwise.
	StatusReason *string `json:"status_reason"`
	// SendAt is when the campaign sends or sent, nil until it is scheduled.
	SendAt *time.Time `json:"send_at"`
	// SendResumesAt is when a send that reached the account's daily marketing
	// limit continues, and is nil otherwise.
	SendResumesAt *time.Time `json:"send_resumes_at"`
	// FinishedAt is when the send ended, nil until then.
	FinishedAt *time.Time `json:"finished_at"`
}

// Campaign represents a campaign with its audience, its content and, for a
// draft, the issues that keep it from being sent.
type Campaign struct {
	CampaignListItem
	// ListIDs are the lists the campaign sends to, in the order they were
	// saved. A list that was deleted is left out.
	ListIDs []uuid.UUID `json:"list_ids"`
	// SegmentIDs are the segments the campaign sends to. A segment that was
	// deleted is left out.
	SegmentIDs []uuid.UUID `json:"segment_ids"`
	// ExcludeListIDs are the lists whose contacts the campaign does not send
	// to. A contact with any membership of one of them, whatever its status,
	// is left out. A list that was deleted is left out.
	ExcludeListIDs []uuid.UUID `json:"exclude_list_ids"`
	// ExcludeSegmentIDs are the segments whose members the campaign does not
	// send to. A segment that was deleted is left out.
	ExcludeSegmentIDs []uuid.UUID `json:"exclude_segment_ids"`
	// Editor is the design's editor, common.TemplateEditorAdvanced,
	// common.TemplateEditorSimple or common.TemplateEditorHTML. It is nil
	// while the campaign has no design.
	Editor *string `json:"editor"`
	// Content is the design's content. It is nil while the campaign has no
	// design, or its design has neither HTML nor text.
	Content *TemplateContent `json:"content"`
	// Ready reports whether the draft has no issue, so it can be sent. It is
	// nil unless the campaign is a draft.
	Ready *bool `json:"ready"`
	// Issues are the reasons the draft cannot be sent yet, in a fixed order,
	// and empty when it is ready. It is nil unless the campaign is a draft.
	Issues []CampaignIssue `json:"issues"`
}

// CampaignSend represents a campaign whose send was accepted. The campaign is
// in common.CampaignStatusValidating, as the send left it.
type CampaignSend struct {
	Campaign
	// Warnings are things the send accepted but you should know, such as an
	// account that cannot send right now while the campaign is sent later, or
	// a design its view-in-browser page cannot show. It is empty when there is
	// none.
	Warnings []string `json:"warnings"`
}

// CampaignNotReady is the body of a 400 answer to SendCampaign or
// TestCampaign. A campaign with issues is refused with the message
// "campaign is not ready to send" and the issues that block the request;
// every other refusal has only Message.
type CampaignNotReady struct {
	Message string `json:"message"`
	// Issues are the issues that block the request, in the order of the
	// CampaignIssue constants. A test is blocked by every issue but
	// CampaignIssueUnsubscribeLinkMissing.
	Issues []CampaignIssue `json:"issues,omitempty"`
}

// CampaignTestFailure represents an address the mail server refused.
type CampaignTestFailure struct {
	// Email is the address as the mail server reported it, which can be one
	// the request did not hold.
	Email   string `json:"email"`
	Message string `json:"message"`
}

// CampaignTest represents what a test send did.
type CampaignTest struct {
	Object string `json:"object"`
	// SentTo are the addresses the mail server accepted.
	SentTo []string `json:"sent_to"`
	// Failed are the addresses the mail server refused.
	Failed []CampaignTestFailure `json:"failed"`
	// ResultUnknown is true when the mail server took the request but its
	// answer could not be read. The test may have gone out, and SentTo and
	// Failed are empty. Check the inboxes before you send another.
	ResultUnknown bool `json:"result_unknown"`
}

// CampaignStatsReport represents a campaign's stored report, which is
// recounted from its recipients from time to time.
type CampaignStatsReport struct {
	// ComputedAt is when the report was counted.
	ComputedAt time.Time `json:"computed_at"`
	// Recipients are the recipients the report counted, the base of
	// BounceRate and DeferralRate.
	Recipients int `json:"recipients"`
	// Openers are the contacts who opened, each once, without bots.
	Openers int `json:"openers"`
	// Clickers are the contacts who clicked, each once, without bots.
	Clickers int `json:"clickers"`
	// Bounces are the recipients that bounced or failed.
	Bounces int `json:"bounces"`
	// Deferrals are the recipients whose delivery was deferred.
	Deferrals int `json:"deferrals"`
	// ByStatus counts the recipients by status, under the names the
	// dashboard's report shows, such as "Queued", "Delivered", "Bounced" or
	// "Administratively Bounced". Each name is a status of
	// ListCampaignRecipients.
	ByStatus map[string]int `json:"by_status"`
	// ByBounceClassification counts the bounces by bounce classification.
	// "Unknown" counts the bounces with no classification and has no
	// recipient page; every other key is a bounce classification of
	// ListCampaignRecipients.
	ByBounceClassification map[string]int `json:"by_bounce_classification"`
}

// CampaignStatsRates represents a campaign's rates: percentages from 0 to 100
// with two decimals. Each is nil while the report is nil or its base is zero.
type CampaignStatsRates struct {
	// OpenRate is the openers out of the delivered messages, at most 100.
	OpenRate *float64 `json:"open_rate"`
	// ClickRate is the clickers out of the delivered messages, at most 100.
	ClickRate *float64 `json:"click_rate"`
	// BounceRate is the bounces out of the report's recipients.
	BounceRate *float64 `json:"bounce_rate"`
	// DeferralRate is the deferrals out of the report's recipients.
	DeferralRate *float64 `json:"deferral_rate"`
}

// CampaignStats represents the numbers of a campaign. The counters are live;
// Report is the stored report, which is recounted from time to time. A
// request never counts: when the report is missing or old, it asks for a
// recount and answers what is stored, so read again later for the new
// numbers.
type CampaignStats struct {
	Object     string    `json:"object"`
	CampaignID uuid.UUID `json:"campaign_id"`
	// Recipients are the recipients the campaign sends to.
	Recipients int `json:"recipients"`
	// Delivered are the messages delivered.
	Delivered int `json:"delivered"`
	// Opens are the opens, without bots.
	Opens int `json:"opens"`
	// Clicks are the clicks, without bots.
	Clicks int `json:"clicks"`
	// Unsubscribes are the recipients who unsubscribed.
	Unsubscribes int `json:"unsubscribes"`
	// Complaints are the recipients who marked the message as spam.
	Complaints int `json:"complaints"`
	// Report is the stored report, nil until the campaign's first recount.
	Report *CampaignStatsReport `json:"report"`
	Rates  CampaignStatsRates   `json:"rates"`
	// RecountFailed is true while the last recount of the report failed, so
	// Report may be old.
	RecountFailed bool `json:"recount_failed"`
}

// CampaignRecipient represents a recipient of a campaign.
type CampaignRecipient struct {
	Object    string    `json:"object"`
	ContactID uuid.UUID `json:"contact_id"`
	// Email is the address the campaign sent to.
	Email string `json:"email"`
	// ListName is the name of the list the contact was sent to from, nil
	// when the contact was sent to from a segment.
	ListName *string `json:"list_name"`
	// SegmentID is the segment the contact was sent to from, nil when the
	// contact was sent to from a list.
	SegmentID *uuid.UUID `json:"segment_id"`
	// SegmentName is the name of the segment the contact was sent to from,
	// as it was at the send, nil when the contact was sent to from a list.
	SegmentName *string `json:"segment_name"`
	// Status is the status under the name the dashboard's report shows, as a
	// key of CampaignStatsReport.ByStatus.
	Status string `json:"status"`
	// BounceClassification is nil when the recipient has none.
	BounceClassification *string `json:"bounce_classification"`
	// BounceReason is nil when the recipient has none.
	BounceReason *string `json:"bounce_reason"`
	// BounceMessage is the receiving server's bounce message, nil when it has
	// none.
	BounceMessage *string `json:"bounce_message"`
	// Attempts are the delivery attempts.
	Attempts int `json:"attempts"`
	// CreatedAt is when the recipient was added to the send.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is when the recipient last changed.
	UpdatedAt time.Time `json:"updated_at"`
}

// PaginatedCampaignsResponse represents a page of campaigns.
type PaginatedCampaignsResponse struct {
	Object     string                `json:"object"`
	Data       []CampaignListItem    `json:"data"`
	Pagination common.PaginationInfo `json:"pagination"`
}

// PaginatedCampaignRecipientsResponse represents a page of a campaign's
// recipients, by contact ID, the highest first. An empty page past the last
// recipient has no Pagination.PreviousCursor; read the first page with no
// cursor.
type PaginatedCampaignRecipientsResponse struct {
	Object     string                `json:"object"`
	Data       []CampaignRecipient   `json:"data"`
	Pagination common.PaginationInfo `json:"pagination"`
}
