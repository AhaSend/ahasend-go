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
	IsBot           *string `json:"is_bot,omitempty"`
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

// RouteAttachment represents an email attachment in route events
type RouteAttachment struct {
	Filename    string  `json:"filename"`
	ContentType string  `json:"content_type"`
	ContentID   *string `json:"content_id,omitempty"`
	Data        string  `json:"data"`
}

// RouteEventData contains data for route webhook events
type RouteEventData struct {
	ID                 string            `json:"id"`
	From               string            `json:"from"`
	ReplyTo            *string           `json:"reply_to,omitempty"`
	To                 string            `json:"to"`
	Subject            string            `json:"subject"`
	MessageID          string            `json:"message_id"`
	Size               int               `json:"size"`
	SpamScore          *float32          `json:"spam_score,omitempty"`
	Bounce             bool              `json:"bounce"`
	CC                 *string           `json:"cc,omitempty"`
	Date               *string           `json:"date,omitempty"`
	InReplyTo          *string           `json:"in_reply_to,omitempty"`
	References         *string           `json:"references,omitempty"`
	AutoSubmitted      *string           `json:"auto_submitted,omitempty"`
	HTMLBody           string            `json:"html_body"`
	PlainBody          string            `json:"plain_body"`
	ReplyFromPlainBody *string           `json:"reply_from_plain_body,omitempty"`
	Attachments        []RouteAttachment `json:"attachments,omitempty"`
	Headers            map[string]string `json:"headers,omitempty"`
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

// MessageFailedEvent is triggered when an email cannot be delivered due to repeated failures
type MessageFailedEvent struct {
	Type      string           `json:"type"`
	WebhookID *string          `json:"webhook_id,omitempty"`
	Timestamp time.Time        `json:"timestamp"`
	Data      MessageEventData `json:"data"`
}

func (e *MessageFailedEvent) GetType() string         { return e.Type }
func (e *MessageFailedEvent) GetTimestamp() time.Time { return e.Timestamp }

// MessageBouncedEvent is triggered when a bounce notification is received
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
	Timestamp time.Time               `json:"timestamp"`
	Data      MessageClickedEventData `json:"data"`
}

func (e *MessageClickedEvent) GetType() string         { return e.Type }
func (e *MessageClickedEvent) GetTimestamp() time.Time { return e.Timestamp }

// Suppression Events

// SuppressionCreatedEvent is triggered when a suppression is created for an email address
type SuppressionCreatedEvent struct {
	Type      string               `json:"type"`
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
