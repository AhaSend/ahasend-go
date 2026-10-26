package webhooks

import (
	"time"
)

// MessageEventData contains the common data for message webhook events.
// For a campaign message, From includes the sender's name, as "Name <address>".
type MessageEventData struct {
	AccountID       string  `json:"account_id"`
	Event           string  `json:"event"`
	From            string  `json:"from"`
	Recipient       string  `json:"recipient"`
	Subject         string  `json:"subject"`
	MessageIDHeader string  `json:"message_id_header"`
	ID              string  `json:"id"`
	UserAgent       *string `json:"user_agent,omitempty"`
	IP              *string `json:"ip,omitempty"`
	// IsBot reports whether the open was attributed to an automated client.
	// Bot detection runs on every open and always yields a definite true or
	// false, so there is no unknown state; the field is only meaningful for
	// message.opened and is false on every other message event. Deliveries
	// predating the always-present field omit it and decode to false.
	IsBot bool `json:"is_bot"`
	// DeliveryAttempt carries diagnostics for the delivery attempt this event
	// reports on. It is optional and absent more often than present, and nil
	// covers both a missing field and an explicit null. Only events sharing this
	// data shape can carry one; MessageClickedEventData never does. See
	// DeliveryAttempt.
	DeliveryAttempt *DeliveryAttempt `json:"delivery_attempt,omitempty"`
	// Unsubscribe says how the recipient unsubscribed. It is set only on
	// message.unsubscribed and is nil on every other message event.
	Unsubscribe *WebhookUnsubscribe `json:"unsubscribe,omitempty"`
	// Complaint is the spam complaint report. It is set only on
	// message.complained and is nil on every other message event.
	Complaint *WebhookComplaint `json:"complaint,omitempty"`
	// CampaignID is the campaign that sent the message. It is set only on the
	// events of a campaign message and is nil for a transactional message.
	CampaignID *string `json:"campaign_id,omitempty"`
	// ListID is the list through which the campaign reached the recipient. It
	// is nil when the campaign did not reach the recipient through a list, and
	// on the events of a transactional message. A JSON null and a missing field
	// both decode to nil.
	ListID *string `json:"list_id,omitempty"`
	// SegmentID is the segment through which the campaign reached the
	// recipient. It is nil when the campaign did not reach the recipient
	// through a segment, and on the events of a transactional message. A JSON
	// null and a missing field both decode to nil.
	SegmentID *string `json:"segment_id,omitempty"`
}

// MessageClickedEventData contains data specific to message clicked events.
// For a campaign message, From includes the sender's name, as "Name <address>".
// CampaignID, ListID, and SegmentID follow the same rules as on
// MessageEventData.
type MessageClickedEventData struct {
	AccountID       string  `json:"account_id"`
	Event           string  `json:"event"`
	From            string  `json:"from"`
	Recipient       string  `json:"recipient"`
	Subject         string  `json:"subject"`
	MessageIDHeader string  `json:"message_id_header"`
	URL             string  `json:"url"`
	UserAgent       string  `json:"user_agent"`
	IP              string  `json:"ip"`
	ID              string  `json:"id"`
	IsBot           bool    `json:"is_bot"`
	CampaignID      *string `json:"campaign_id,omitempty"`
	ListID          *string `json:"list_id,omitempty"`
	SegmentID       *string `json:"segment_id,omitempty"`
}

// UnsubscribeMethod says how a recipient unsubscribed.
type UnsubscribeMethod string

const (
	// UnsubscribeMethodOneClick is a one-click unsubscribe from the
	// recipient's email client (RFC 8058).
	UnsubscribeMethodOneClick UnsubscribeMethod = "one_click"
	// UnsubscribeMethodLink is any other unsubscribe, for example through the
	// unsubscribe link in the email.
	UnsubscribeMethodLink UnsubscribeMethod = "link"
)

// WebhookUnsubscribe says how the recipient unsubscribed. It is carried by
// message.unsubscribed.
type WebhookUnsubscribe struct {
	Method UnsubscribeMethod `json:"method"`
}

// ComplaintReporter is the feedback loop that sent a spam complaint report.
//
// The set is open. AhaSend can add feedback loops at any time, so the
// constants below describe the known values rather than everything that may
// arrive. The underlying type is string so that an unfamiliar value decodes
// without error: keep a default branch for the rest, and never reject a
// delivery because the value is unrecognized.
type ComplaintReporter string

const (
	ComplaintReporterYahoo     ComplaintReporter = "yahoo"
	ComplaintReporterMicrosoft ComplaintReporter = "microsoft"
	ComplaintReporterComcast   ComplaintReporter = "comcast"
)

// WebhookComplaint is a spam complaint report from a mailbox provider's
// feedback loop. It is carried by message.complained and contact.complained.
type WebhookComplaint struct {
	// FeedbackType is the feedback type of the report (the ARF Feedback-Type),
	// in lower case: "abuse" or "fraud".
	FeedbackType string `json:"feedback_type"`
	// Reporter is the feedback loop that sent the report. The set of values is
	// open — see ComplaintReporter.
	Reporter ComplaintReporter `json:"reporter"`
}

// WebhookOrigin says who made the change that a contact or suppression event
// reports. Message events, route events and domain.dns_error carry no
// origin.
type WebhookOrigin string

const (
	// OriginRecipient is a change the contact or recipient made, through an
	// unsubscribe link, a one-click unsubscribe, or a spam complaint.
	OriginRecipient WebhookOrigin = "recipient"
	// OriginAPI is a change you made through the AhaSend API.
	OriginAPI WebhookOrigin = "api"
	// OriginDashboard is a change you made through the AhaSend dashboard.
	OriginDashboard WebhookOrigin = "dashboard"
	// OriginSystem is a change AhaSend made, for example after a bounce, an
	// expired message or an email validation, or by AhaSend support.
	OriginSystem WebhookOrigin = "system"
)

// DeliveryAttemptClassification is the bucket the bounce classifier assigned to
// a delivery attempt.
//
// The set is open. New buckets can be introduced at any time, so the constants
// below describe what AhaSend emits today rather than everything that may ever
// arrive. The underlying type is string precisely so that an unfamiliar value
// decodes without error: switch on the values you handle and keep a default
// branch for the rest. Never reject a delivery because the value is
// unrecognized — when more than 100 attempts in a row fail, retries included,
// the webhook or route is automatically disabled.
type DeliveryAttemptClassification string

const (
	ClassificationInvalidRecipient DeliveryAttemptClassification = "InvalidRecipient"
	ClassificationBadDomain        DeliveryAttemptClassification = "BadDomain"
	ClassificationInactiveMailbox  DeliveryAttemptClassification = "InactiveMailbox"
	ClassificationInvalidSender    DeliveryAttemptClassification = "InvalidSender"
	ClassificationQuotaIssues      DeliveryAttemptClassification = "QuotaIssues"
	ClassificationNoAnswerFromHost DeliveryAttemptClassification = "NoAnswerFromHost"
	ClassificationBadConnection    DeliveryAttemptClassification = "BadConnection"
	ClassificationDNSFailure       DeliveryAttemptClassification = "DNSFailure"
	ClassificationRoutingErrors    DeliveryAttemptClassification = "RoutingErrors"
	// ClassificationTransientFailure is a bounce classification and is
	// unrelated to the message.transient_error event type and to the event
	// field: a delivery carrying it is not necessarily a transient-error event,
	// and a transient-error event does not necessarily carry it.
	ClassificationTransientFailure     DeliveryAttemptClassification = "TransientFailure"
	ClassificationMessageExpired       DeliveryAttemptClassification = "MessageExpired"
	ClassificationProtocolErrors       DeliveryAttemptClassification = "ProtocolErrors"
	ClassificationAuthenticationFailed DeliveryAttemptClassification = "AuthenticationFailed"
	ClassificationPolicyRelated        DeliveryAttemptClassification = "PolicyRelated"
	ClassificationUncategorized        DeliveryAttemptClassification = "Uncategorized"
)

// DeliveryAttempt holds the diagnostics for a single delivery attempt: the SMTP
// status code recorded for it, the response text, and — on failures — the
// bucket the bounce classifier assigned.
//
// Absence is normal rather than an error. Today message.delivered,
// message.bounced, and message.transient_error carry an attempt, and only when
// an SMTP attempt was actually recorded; retry exhaustion, out-of-band bounces
// where the DSN arrives after the destination already accepted the message,
// non-SMTP routing, and empty responses all arrive without one. That list is
// not closed. A nil pointer covers both a missing field and an explicit null,
// which mean the same thing, so check for nil before reading and never treat
// the absence as a failure.
//
// Sandbox sends are the exception to the non-SMTP case: rather than omitting
// the object they synthesize a representative one. On sandbox deliveries and
// test webhooks alike the values are representative rather than observed, and a
// sandbox simulation can substitute the classification of the outcome it was
// asked to simulate, which then may not agree with the representative SMTPCode
// and Response beside it. Do not calibrate a Classification switch against
// either.
//
// Response and Description are free-form diagnostic text — usually the
// destination's words, sometimes AhaSend's — and routinely embed the recipient
// address and parts of the message. Neither belongs in a log or a metric
// label. Log SMTPCode, EnhancedStatusCode, Classification, and Command
// instead: those are short values drawn from small vocabularies — low
// cardinality in practice, even though the classification vocabulary is open
// rather than closed — and Command excludes its own arguments, so no envelope
// address rides along in it.
type DeliveryAttempt struct {
	// Classification is the bucket the bounce classifier assigned. It is absent
	// on successful deliveries and on a failure the classifier did not label.
	// The set of values is open — see DeliveryAttemptClassification.
	Classification *DeliveryAttemptClassification `json:"classification,omitempty"`
	// SMTPCode is the SMTP status code recorded for the attempt, usually the
	// one the destination returned. It is always present when the attempt
	// itself is, and 0 is a real value: it means AhaSend recorded response
	// content without an SMTP code, typically an internal error. Do not test
	// this field for truthiness and do not read 0 as the field being absent.
	SMTPCode int `json:"smtp_code"`
	// EnhancedStatusCode is the RFC 3463 enhanced status code in
	// class.subject.detail form (for example "5.1.1"), when one was recorded
	// for the attempt.
	EnhancedStatusCode *string `json:"enhanced_status_code,omitempty"`
	// Response is the outcome text as the MTA recorded it, present when
	// non-empty. It is not reliably code-free — some destinations repeat the
	// reply code, the enhanced code, or both at the front of the text — and it
	// is not always the destination's own words, since a failure raised inside
	// AhaSend carries AhaSend's description instead. Read it as human-readable
	// diagnostics and branch on Classification, SMTPCode, and
	// EnhancedStatusCode instead. Not safe to log.
	Response *string `json:"response,omitempty"`
	// Description is a human-readable translation of a complex Response,
	// present only when that translation differs from Response itself. It is
	// display prose rather than an identifier and its wording changes as the
	// translations improve, so never compare it to a literal and never parse
	// it. Not safe to log.
	Description *string `json:"description,omitempty"`
	// Command is the normalized name of the SMTP command that was in flight
	// when the attempt was recorded (for example "RCPT TO" or "DATA"), when the
	// MTA reported one. The command's arguments are not included, so unlike
	// Response it carries no envelope address and is safe to log.
	Command *string `json:"command,omitempty"`
}

// SuppressionEventData contains data for suppression webhook events. It is the
// suppression as it was right after the change, or right before the delete for
// suppression.deleted.
type SuppressionEventData struct {
	AccountID string `json:"account_id"`
	// Origin says who made the change: who created, changed or deleted the
	// suppression.
	Origin        WebhookOrigin `json:"origin"`
	Recipient     string        `json:"recipient"`
	CreatedAt     time.Time     `json:"created_at"`
	ExpiresAt     time.Time     `json:"expires_at"`
	Reason        string        `json:"reason"`
	SendingDomain string        `json:"sending_domain"`
}

// DomainEventData contains data for domain webhook events
type DomainEventData struct {
	Domain           string    `json:"domain"`
	AccountID        string    `json:"account_id"`
	SPFValid         bool      `json:"spf_valid"`
	DKIMValid        bool      `json:"dkim_valid"`
	DMARCValid       bool      `json:"dmarc_valid"`
	DNSLastCheckedAt time.Time `json:"dns_last_checked_at"`
}

// RouteAttachment represents a conventional, inline, or inferred email attachment in route events.
type RouteAttachment struct {
	Filename    string  `json:"filename"`
	ContentType string  `json:"content_type"`
	ContentID   *string `json:"content_id,omitempty"`
	// Disposition is the Content-Disposition type of the MIME part, usually
	// "attachment" or "inline", and "" when the part carried no
	// Content-Disposition header. The value is deliberately unconstrained —
	// it is whatever token the sending mail server wrote, lowercased and
	// trimmed of parameters — so tokens other than "attachment" and "inline"
	// do arrive in practice. Do not validate it against a fixed list; treat
	// an unrecognized value as "attachment". Note that "" does not mean "not
	// inline": embedded images from Gmail and Outlook arrive with an empty
	// disposition and a populated ContentID, which is what separates them
	// from real attachments.
	Disposition string `json:"disposition"`
	Data        string `json:"data"`
}

// RouteEventData contains data for route webhook events.
//
// ReplyTo is the message's Reply-To header as received, which may include a
// display name or several addresses. To is the recipient address the route
// matched; on a route that groups by Message-ID it is the message's To header
// as received, which may include display names and several addresses. Parse
// both as address lists, not as single bare addresses.
//
// SpamScore is the spam filter's score for the message. Higher means more
// likely spam; it can be below 0 or above 10. It is 0 when the message was not
// checked.
type RouteEventData struct {
	ID                 string   `json:"id"`
	From               string   `json:"from"`
	ReplyTo            *string  `json:"reply_to,omitempty"`
	To                 string   `json:"to"`
	Subject            string   `json:"subject"`
	MessageID          string   `json:"message_id"`
	Size               int      `json:"size"`
	SpamScore          *float64 `json:"spam_score,omitempty"`
	Bounce             bool     `json:"bounce"`
	CC                 *string  `json:"cc,omitempty"`
	Date               *string  `json:"date,omitempty"`
	InReplyTo          *string  `json:"in_reply_to,omitempty"`
	References         *string  `json:"references,omitempty"`
	AutoSubmitted      *string  `json:"auto_submitted,omitempty"`
	HTMLBody           string   `json:"html_body"`
	PlainBody          string   `json:"plain_body"`
	ReplyFromPlainBody *string  `json:"reply_from_plain_body,omitempty"`
	// Attachments includes conventional attachments, inline MIME parts, and
	// filename-bearing MIME parts without a Content-Disposition header.
	Attachments []RouteAttachment `json:"attachments,omitempty"`
	Headers     map[string]string `json:"headers,omitempty"`
}

// Contact is a contact in the shape that the Contacts API returns.
type Contact struct {
	// Object is always "contact".
	Object string `json:"object"`
	// ID is the contact's unique identifier.
	ID string `json:"id"`
	// CreatedAt is when the contact was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is when the contact was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// Email is the contact's normalized email address.
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	// Status is "enabled", "disabled", or "blocked".
	Status       string `json:"status"`
	StatusReason string `json:"status_reason"`
	// Unsubscribed reports the contact's unsubscribe from all lists, and
	// UnsubscribedAt is when it was set; nil when it is not set.
	Unsubscribed   bool       `json:"unsubscribed"`
	UnsubscribedAt *time.Time `json:"unsubscribed_at"`
	// Attributes holds the contact's custom fields. A value can be any JSON
	// value: null, a boolean, a number, a string, an array, or a nested
	// object, decoded as encoding/json decodes into any. New writes accept only
	// the values the dashboard defines, but older values can have any of these
	// shapes, so do not assume the type of a field.
	Attributes map[string]any `json:"attributes"`
	// ValidationStatus is "unvalidated", "valid", "invalid", "risky", or
	// "unknown".
	ValidationStatus string `json:"validation_status"`
	// LastValidatedAt is when the address was last validated; nil when it
	// never was.
	LastValidatedAt *time.Time `json:"last_validated_at"`
}

// WebhookList is the list of a membership that a contact event reports on.
type WebhookList struct {
	// ID is the list ID.
	ID string `json:"id"`
	// Name is the list name.
	Name string `json:"name"`
}

// WebhookListContact is one contact's membership of one list, in the shape
// that the Lists API returns without the embedded contact and list. It has no
// ID of its own: a membership is addressed by the ListID and ContactID pair.
type WebhookListContact struct {
	// Object is always "list_contact".
	Object string `json:"object"`
	// ListID is the list of the membership.
	ListID string `json:"list_id"`
	// ContactID is the contact of the membership.
	ContactID string `json:"contact_id"`
	// Email is the contact's email address.
	Email string `json:"email"`
	// SubscriptionStatus is the per-list consent: "unconfirmed", "confirmed",
	// "unsubscribed", or "complained". It is independent of the contact's
	// account-wide Unsubscribed flag.
	SubscriptionStatus string `json:"subscription_status"`
	// SubscribedAt is when consent was given. It is set when the membership
	// becomes confirmed and is not changed afterwards.
	SubscribedAt *time.Time `json:"subscribed_at"`
	// UnsubscribedAt is when the membership became unsubscribed, or when the
	// complaint arrived for a complained membership. It is cleared when the
	// membership becomes confirmed again.
	UnsubscribedAt *time.Time `json:"unsubscribed_at"`
	// CreatedAt is when the membership was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is when the membership was last updated.
	UpdatedAt time.Time `json:"updated_at"`
}

// ContactEventData contains the common data for contact webhook events.
// Contact events are available to accounts with the marketing features.
type ContactEventData struct {
	// AccountID is the account of the contact.
	AccountID string `json:"account_id"`
	// Origin says who made the change.
	Origin WebhookOrigin `json:"origin"`
	// CampaignID is the campaign whose email caused the change, when a
	// recipient unsubscribed from a campaign email or reported it as spam. It
	// is nil otherwise.
	CampaignID *string `json:"campaign_id,omitempty"`
	// SegmentID is the segment through which that campaign reached the
	// recipient, when the campaign was sent to a segment. It is nil otherwise.
	SegmentID *string `json:"segment_id,omitempty"`
	// Contact is the contact as it was right after the change, or right
	// before the delete for contact.deleted.
	Contact Contact `json:"contact"`
}

// ContactChangedField is the name of a contact field that contact.updated
// reports as changed.
//
// The set is open: a new contact field can appear at any time, so the
// constants below describe the known values rather than everything that may
// arrive. The underlying type is string so that an unfamiliar value decodes
// without error.
type ContactChangedField string

const (
	ChangedFieldEmail        ContactChangedField = "email"
	ChangedFieldFirstName    ContactChangedField = "first_name"
	ChangedFieldLastName     ContactChangedField = "last_name"
	ChangedFieldStatus       ContactChangedField = "status"
	ChangedFieldStatusReason ContactChangedField = "status_reason"
	// ChangedFieldAttributes reports a change to any key of Attributes.
	ChangedFieldAttributes ContactChangedField = "attributes"
)

// ContactUpdatedEventData contains the data for contact.updated events.
type ContactUpdatedEventData struct {
	ContactEventData
	// ChangedFields lists the Contact fields whose values changed. It never
	// lists unsubscribed or unsubscribed_at: the unsubscribe from all lists
	// sends contact.unsubscribed and contact.subscribed instead. The set of
	// values is open — see ContactChangedField.
	ChangedFields []ContactChangedField `json:"changed_fields"`
}

// ContactListEventData contains the data for the contact events that can be
// about a list membership: contact.added_to_list, contact.removed_from_list,
// contact.subscribed, contact.unsubscribed, and contact.complained.
type ContactListEventData struct {
	ContactEventData
	// List is the list of the membership. It is nil when the event is about
	// the contact's unsubscribe from all lists: contact.unsubscribed when it is
	// set, contact.subscribed when it is cleared, and contact.complained for a
	// complaint about a campaign sent to a segment. A JSON null and a missing
	// field both decode to nil.
	List *WebhookList `json:"list"`
	// ListContact is the membership as it was right after the change, or
	// right before the removal for contact.removed_from_list. It is nil when
	// List is nil.
	ListContact *WebhookListContact `json:"list_contact"`
	// PreviousSubscriptionStatus is the membership's subscription status
	// before the change. It is set only on contact.subscribed,
	// contact.unsubscribed, and contact.complained, when List is not nil.
	PreviousSubscriptionStatus *string `json:"previous_subscription_status,omitempty"`
	// Complaint is the spam complaint report. It is set only on
	// contact.complained.
	Complaint *WebhookComplaint `json:"complaint,omitempty"`
}

// Message Events

// MessageReceptionEvent is triggered when an email has been received and queued
type MessageReceptionEvent struct {
	Type      string           `json:"type"`
	WebhookID *string          `json:"webhook_id,omitempty"`
	Timestamp time.Time        `json:"timestamp"`
	Data      MessageEventData `json:"data"`
}

func (e *MessageReceptionEvent) GetType() string         { return e.Type }
func (e *MessageReceptionEvent) GetTimestamp() time.Time { return e.Timestamp }

// MessageDeliveredEvent is triggered when an email has been successfully delivered
type MessageDeliveredEvent struct {
	Type      string           `json:"type"`
	WebhookID *string          `json:"webhook_id,omitempty"`
	Timestamp time.Time        `json:"timestamp"`
	Data      MessageEventData `json:"data"`
}

func (e *MessageDeliveredEvent) GetType() string         { return e.Type }
func (e *MessageDeliveredEvent) GetTimestamp() time.Time { return e.Timestamp }

// MessageTransientErrorEvent is triggered when email delivery is delayed due to a temporary issue
type MessageTransientErrorEvent struct {
	Type      string           `json:"type"`
	WebhookID *string          `json:"webhook_id,omitempty"`
	Timestamp time.Time        `json:"timestamp"`
	Data      MessageEventData `json:"data"`
}

func (e *MessageTransientErrorEvent) GetType() string         { return e.Type }
func (e *MessageTransientErrorEvent) GetTimestamp() time.Time { return e.Timestamp }

// MessageFailedEvent is triggered when a message exhausts its retry window
// without being delivered. It reports retry exhaustion rather than one specific
// SMTP exchange, so Data.DeliveryAttempt is always nil here. An immediate
// permanent rejection by the destination is not retry exhaustion and arrives as
// MessageBouncedEvent instead.
type MessageFailedEvent struct {
	Type      string           `json:"type"`
	WebhookID *string          `json:"webhook_id,omitempty"`
	Timestamp time.Time        `json:"timestamp"`
	Data      MessageEventData `json:"data"`
}

func (e *MessageFailedEvent) GetType() string         { return e.Type }
func (e *MessageFailedEvent) GetTimestamp() time.Time { return e.Timestamp }

// MessageBouncedEvent is triggered when a message reaches a bounced outcome,
// which covers two cases: an immediate permanent SMTP rejection, and an
// out-of-band delivery status notification arriving after the destination had
// already accepted the message. They differ in what they carry — a direct
// rejection normally includes Data.DeliveryAttempt, while an out-of-band bounce
// describes no single SMTP attempt and so carries none.
type MessageBouncedEvent struct {
	Type      string           `json:"type"`
	WebhookID *string          `json:"webhook_id,omitempty"`
	Timestamp time.Time        `json:"timestamp"`
	Data      MessageEventData `json:"data"`
}

func (e *MessageBouncedEvent) GetType() string         { return e.Type }
func (e *MessageBouncedEvent) GetTimestamp() time.Time { return e.Timestamp }

// MessageSuppressedEvent is triggered when no delivery attempt was made because the recipient is suppressed
type MessageSuppressedEvent struct {
	Type      string           `json:"type"`
	WebhookID *string          `json:"webhook_id,omitempty"`
	Timestamp time.Time        `json:"timestamp"`
	Data      MessageEventData `json:"data"`
}

func (e *MessageSuppressedEvent) GetType() string         { return e.Type }
func (e *MessageSuppressedEvent) GetTimestamp() time.Time { return e.Timestamp }

// MessageOpenedEvent is triggered when the recipient opens your email
type MessageOpenedEvent struct {
	Type      string           `json:"type"`
	WebhookID *string          `json:"webhook_id,omitempty"`
	Timestamp time.Time        `json:"timestamp"`
	Data      MessageEventData `json:"data"`
}

func (e *MessageOpenedEvent) GetType() string         { return e.Type }
func (e *MessageOpenedEvent) GetTimestamp() time.Time { return e.Timestamp }

// MessageClickedEvent is triggered when the recipient clicks a tracked link in your email
type MessageClickedEvent struct {
	Type      string                  `json:"type"`
	WebhookID *string                 `json:"webhook_id,omitempty"`
	Timestamp time.Time               `json:"timestamp"`
	Data      MessageClickedEventData `json:"data"`
}

func (e *MessageClickedEvent) GetType() string         { return e.Type }
func (e *MessageClickedEvent) GetTimestamp() time.Time { return e.Timestamp }

// MessageUnsubscribedEvent is triggered when the recipient of a campaign email
// unsubscribes through its unsubscribe link or through the one-click
// unsubscribe of their email client. Data.Unsubscribe says which.
//
// It is sent for every unsubscribe request that AhaSend accepts for the
// message, also when the recipient is already unsubscribed or suppressed, so
// drop repeats by Type and Data.ID. Repeats are rate limited per message while
// the contact events of the same requests are not, so a contact event can
// arrive without a message.unsubscribed of its own.
type MessageUnsubscribedEvent struct {
	Type      string           `json:"type"`
	WebhookID *string          `json:"webhook_id,omitempty"`
	Timestamp time.Time        `json:"timestamp"`
	Data      MessageEventData `json:"data"`
}

func (e *MessageUnsubscribedEvent) GetType() string         { return e.Type }
func (e *MessageUnsubscribedEvent) GetTimestamp() time.Time { return e.Timestamp }

// MessageComplainedEvent is triggered when the recipient's mailbox provider
// reports a campaign email as spam through its feedback loop. Data.Complaint
// is the report.
//
// It is sent for every complaint report that AhaSend matches to the message,
// also when the recipient is already unsubscribed or suppressed, so drop
// repeats by Type and Data.ID.
type MessageComplainedEvent struct {
	Type      string           `json:"type"`
	WebhookID *string          `json:"webhook_id,omitempty"`
	Timestamp time.Time        `json:"timestamp"`
	Data      MessageEventData `json:"data"`
}

func (e *MessageComplainedEvent) GetType() string         { return e.Type }
func (e *MessageComplainedEvent) GetTimestamp() time.Time { return e.Timestamp }

// Suppression Events

// SuppressionCreatedEvent is triggered when a suppression is created for an
// email address, or when an existing suppression gets a new reason. For a new
// reason Data.CreatedAt stays the time the suppression was first created, so
// key stored suppressions by Data.Recipient and Data.SendingDomain.
type SuppressionCreatedEvent struct {
	Type      string               `json:"type"`
	WebhookID *string              `json:"webhook_id,omitempty"`
	Timestamp time.Time            `json:"timestamp"`
	Data      SuppressionEventData `json:"data"`
}

func (e *SuppressionCreatedEvent) GetType() string         { return e.Type }
func (e *SuppressionCreatedEvent) GetTimestamp() time.Time { return e.Timestamp }

// SuppressionDeletedEvent is triggered when a suppression is deleted, through
// the API, the dashboard, or by AhaSend support. Future emails to the address
// are attempted again. Data is the suppression as it was right before the
// delete. Wiping all suppressions of an account or a domain and the removal of
// expired suppressions send no event.
type SuppressionDeletedEvent struct {
	Type      string               `json:"type"`
	WebhookID *string              `json:"webhook_id,omitempty"`
	Timestamp time.Time            `json:"timestamp"`
	Data      SuppressionEventData `json:"data"`
}

func (e *SuppressionDeletedEvent) GetType() string         { return e.Type }
func (e *SuppressionDeletedEvent) GetTimestamp() time.Time { return e.Timestamp }

// Domain Events

// DomainDNSErrorEvent is triggered when DNS configuration issues are detected
type DomainDNSErrorEvent struct {
	Type      string          `json:"type"`
	WebhookID *string         `json:"webhook_id,omitempty"`
	Timestamp time.Time       `json:"timestamp"`
	Data      DomainEventData `json:"data"`
}

func (e *DomainDNSErrorEvent) GetType() string         { return e.Type }
func (e *DomainDNSErrorEvent) GetTimestamp() time.Time { return e.Timestamp }

// Contact Events

// ContactCreatedEvent is triggered when a contact is created through the API
// or the dashboard. Contacts created by a CSV import send no event.
type ContactCreatedEvent struct {
	Type      string           `json:"type"`
	WebhookID *string          `json:"webhook_id,omitempty"`
	Timestamp time.Time        `json:"timestamp"`
	Data      ContactEventData `json:"data"`
}

func (e *ContactCreatedEvent) GetType() string         { return e.Type }
func (e *ContactCreatedEvent) GetTimestamp() time.Time { return e.Timestamp }

// ContactUpdatedEvent is triggered when a field of a contact changes.
// Data.ChangedFields lists the fields whose values changed. A handler that
// writes the contact back through the API receives its own write as a new
// contact.updated with OriginAPI, so skip the write when the contact already
// has the values the handler wants.
type ContactUpdatedEvent struct {
	Type      string                  `json:"type"`
	WebhookID *string                 `json:"webhook_id,omitempty"`
	Timestamp time.Time               `json:"timestamp"`
	Data      ContactUpdatedEventData `json:"data"`
}

func (e *ContactUpdatedEvent) GetType() string         { return e.Type }
func (e *ContactUpdatedEvent) GetTimestamp() time.Time { return e.Timestamp }

// ContactDeletedEvent is triggered when a contact is deleted. Its list
// memberships are deleted with it and send no contact.removed_from_list.
type ContactDeletedEvent struct {
	Type      string           `json:"type"`
	WebhookID *string          `json:"webhook_id,omitempty"`
	Timestamp time.Time        `json:"timestamp"`
	Data      ContactEventData `json:"data"`
}

func (e *ContactDeletedEvent) GetType() string         { return e.Type }
func (e *ContactDeletedEvent) GetTimestamp() time.Time { return e.Timestamp }

// ContactAddedToListEvent is triggered when a contact is added to a list, with
// any subscription status. A new membership sends only this event.
type ContactAddedToListEvent struct {
	Type      string               `json:"type"`
	WebhookID *string              `json:"webhook_id,omitempty"`
	Timestamp time.Time            `json:"timestamp"`
	Data      ContactListEventData `json:"data"`
}

func (e *ContactAddedToListEvent) GetType() string         { return e.Type }
func (e *ContactAddedToListEvent) GetTimestamp() time.Time { return e.Timestamp }

// ContactRemovedFromListEvent is triggered when a contact is removed from a
// list. Data.ListContact is the membership as it was right before the
// removal.
type ContactRemovedFromListEvent struct {
	Type      string               `json:"type"`
	WebhookID *string              `json:"webhook_id,omitempty"`
	Timestamp time.Time            `json:"timestamp"`
	Data      ContactListEventData `json:"data"`
}

func (e *ContactRemovedFromListEvent) GetType() string         { return e.Type }
func (e *ContactRemovedFromListEvent) GetTimestamp() time.Time { return e.Timestamp }

// ContactSubscribedEvent is triggered when a membership's subscription status
// becomes confirmed, or when the contact's unsubscribe from all lists is
// cleared. In the second case Data.List and Data.ListContact are nil.
type ContactSubscribedEvent struct {
	Type      string               `json:"type"`
	WebhookID *string              `json:"webhook_id,omitempty"`
	Timestamp time.Time            `json:"timestamp"`
	Data      ContactListEventData `json:"data"`
}

func (e *ContactSubscribedEvent) GetType() string         { return e.Type }
func (e *ContactSubscribedEvent) GetTimestamp() time.Time { return e.Timestamp }

// ContactUnsubscribedEvent is triggered when a membership's subscription
// status becomes unsubscribed, or when the contact's unsubscribe from all lists
// is set. In the second case Data.List and Data.ListContact are nil. A
// recipient who unsubscribes from a campaign email sends it with
// OriginRecipient and Data.CampaignID. A recipient of a campaign sent to a
// segment unsubscribes from all lists, and Data.SegmentID is then set too.
type ContactUnsubscribedEvent struct {
	Type      string               `json:"type"`
	WebhookID *string              `json:"webhook_id,omitempty"`
	Timestamp time.Time            `json:"timestamp"`
	Data      ContactListEventData `json:"data"`
}

func (e *ContactUnsubscribedEvent) GetType() string         { return e.Type }
func (e *ContactUnsubscribedEvent) GetTimestamp() time.Time { return e.Timestamp }

// ContactComplainedEvent is triggered when a spam complaint about a campaign
// email changes the contact's membership of the campaign's list to
// complained, or is recorded for a recipient of a campaign sent to a segment.
// In the second case Data.List and Data.ListContact are nil, and the event is
// sent also when the contact is already unsubscribed; Data.Contact.Unsubscribed
// says whether the contact is unsubscribed from all lists after the complaint.
type ContactComplainedEvent struct {
	Type      string               `json:"type"`
	WebhookID *string              `json:"webhook_id,omitempty"`
	Timestamp time.Time            `json:"timestamp"`
	Data      ContactListEventData `json:"data"`
}

func (e *ContactComplainedEvent) GetType() string         { return e.Type }
func (e *ContactComplainedEvent) GetTimestamp() time.Time { return e.Timestamp }

// Route Events

// RouteMessageEvent is triggered when an inbound email is received and processed through a configured route
type RouteMessageEvent struct {
	Type      string         `json:"type"`
	RouteID   *string        `json:"route_id,omitempty"`
	Timestamp time.Time      `json:"timestamp"`
	Data      RouteEventData `json:"data"`
}

func (e *RouteMessageEvent) GetType() string         { return e.Type }
func (e *RouteMessageEvent) GetTimestamp() time.Time { return e.Timestamp }

// Helper functions for working with webhook events

// IsMessageEvent checks if an event is a message-related event
func IsMessageEvent(event WebhookEvent) bool {
	switch event.(type) {
	case *MessageReceptionEvent, *MessageDeliveredEvent, *MessageTransientErrorEvent,
		*MessageFailedEvent, *MessageBouncedEvent, *MessageSuppressedEvent,
		*MessageOpenedEvent, *MessageClickedEvent, *MessageUnsubscribedEvent,
		*MessageComplainedEvent:
		return true
	}
	return false
}

// IsSuppressionEvent checks if an event is a suppression-related event
func IsSuppressionEvent(event WebhookEvent) bool {
	switch event.(type) {
	case *SuppressionCreatedEvent, *SuppressionDeletedEvent:
		return true
	}
	return false
}

// IsContactEvent checks if an event is a contact-related event
func IsContactEvent(event WebhookEvent) bool {
	switch event.(type) {
	case *ContactCreatedEvent, *ContactUpdatedEvent, *ContactDeletedEvent,
		*ContactAddedToListEvent, *ContactRemovedFromListEvent,
		*ContactSubscribedEvent, *ContactUnsubscribedEvent, *ContactComplainedEvent:
		return true
	}
	return false
}

// IsDomainEvent checks if an event is a domain-related event
func IsDomainEvent(event WebhookEvent) bool {
	_, ok := event.(*DomainDNSErrorEvent)
	return ok
}

// IsRouteEvent checks if an event is a route-related event
func IsRouteEvent(event WebhookEvent) bool {
	_, ok := event.(*RouteMessageEvent)
	return ok
}

// GetMessageEventData extracts the MessageEventData from a message event
// Returns nil if the event is not a message event or doesn't contain MessageEventData
func GetMessageEventData(event WebhookEvent) *MessageEventData {
	switch e := event.(type) {
	case *MessageReceptionEvent:
		return &e.Data
	case *MessageDeliveredEvent:
		return &e.Data
	case *MessageTransientErrorEvent:
		return &e.Data
	case *MessageFailedEvent:
		return &e.Data
	case *MessageBouncedEvent:
		return &e.Data
	case *MessageSuppressedEvent:
		return &e.Data
	case *MessageOpenedEvent:
		return &e.Data
	case *MessageUnsubscribedEvent:
		return &e.Data
	case *MessageComplainedEvent:
		return &e.Data
	}
	return nil
}
