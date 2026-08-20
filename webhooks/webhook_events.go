package webhooks

import (
	"time"
)

// MessageEventData contains the common data for message webhook events
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
}

// MessageClickedEventData contains data specific to message clicked events
type MessageClickedEventData struct {
	AccountID       string `json:"account_id"`
	Event           string `json:"event"`
	From            string `json:"from"`
	Recipient       string `json:"recipient"`
	Subject         string `json:"subject"`
	MessageIDHeader string `json:"message_id_header"`
	URL             string `json:"url"`
	UserAgent       string `json:"user_agent"`
	IP              string `json:"ip"`
	ID              string `json:"id"`
	IsBot           bool   `json:"is_bot"`
}

// DeliveryAttemptClassification is the bucket the bounce classifier assigned to
// a delivery attempt.
//
// The set is open. New buckets can be introduced at any time, so the constants
// below describe what AhaSend emits today rather than everything that may ever
// arrive. The underlying type is string precisely so that an unfamiliar value
// decodes without error: switch on the values you handle and keep a default
// branch for the rest. Never reject a delivery because the value is
// unrecognized — a webhook endpoint that answers 400 is disabled after 100
// consecutive failures.
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

// SuppressionEventData contains data for suppression webhook events
type SuppressionEventData struct {
	AccountID     string    `json:"account_id"`
	Recipient     string    `json:"recipient"`
	CreatedAt     time.Time `json:"created_at"`
	ExpiresAt     time.Time `json:"expires_at"`
	Reason        string    `json:"reason"`
	SendingDomain string    `json:"sending_domain"`
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

// RouteEventData contains data for route webhook events
type RouteEventData struct {
	ID                 string   `json:"id"`
	From               string   `json:"from"`
	ReplyTo            *string  `json:"reply_to,omitempty"`
	To                 string   `json:"to"`
	Subject            string   `json:"subject"`
	MessageID          string   `json:"message_id"`
	Size               int      `json:"size"`
	SpamScore          *float32 `json:"spam_score,omitempty"`
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

// Suppression Events

// SuppressionCreatedEvent is triggered when a suppression is created for an email address
type SuppressionCreatedEvent struct {
	Type      string               `json:"type"`
	WebhookID *string              `json:"webhook_id,omitempty"`
	Timestamp time.Time            `json:"timestamp"`
	Data      SuppressionEventData `json:"data"`
}

func (e *SuppressionCreatedEvent) GetType() string         { return e.Type }
func (e *SuppressionCreatedEvent) GetTimestamp() time.Time { return e.Timestamp }

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
		*MessageOpenedEvent, *MessageClickedEvent:
		return true
	}
	return false
}

// IsSuppressionEvent checks if an event is a suppression-related event
func IsSuppressionEvent(event WebhookEvent) bool {
	_, ok := event.(*SuppressionCreatedEvent)
	return ok
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
	}
	return nil
}
