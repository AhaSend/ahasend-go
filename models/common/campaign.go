package common

// Campaign statuses: where a campaign is in its life. Only a
// CampaignStatusDraft campaign can be changed or deleted.
//
//   - CampaignStatusDraft: being prepared.
//   - CampaignStatusValidating: scheduled, and its lists are being checked.
//   - CampaignStatusWaitingForReview: held for a review before it sends.
//   - CampaignStatusScheduled: accepted, waiting for its send time.
//   - CampaignStatusSending: sending now.
//   - CampaignStatusPaused: sending stopped for now; the status reason says
//     why.
//   - CampaignStatusSent: every message was sent.
//   - CampaignStatusCanceled: stopped before it finished; the status reason
//     says why.
//   - CampaignStatusFailed: could not be sent; the status reason says why.
const (
	CampaignStatusDraft            = "draft"
	CampaignStatusValidating       = "validating"
	CampaignStatusWaitingForReview = "waiting_for_review"
	CampaignStatusScheduled        = "scheduled"
	CampaignStatusSending          = "sending"
	CampaignStatusPaused           = "paused"
	CampaignStatusSent             = "sent"
	CampaignStatusCanceled         = "canceled"
	CampaignStatusFailed           = "failed"
)
