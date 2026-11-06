package responses

import (
	"time"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/google/uuid"
)

type ContentParsed struct {
	Parts       []ContentPart       `json:"parts"`
	Attachments []ContentAttachment `json:"attachments"`
	Headers     map[string][]string `json:"headers"`
}

type ContentPart struct {
	ContentType string `json:"content_type"`
	Content     string `json:"content"`
}

type ContentAttachment struct {
	Filename    string `json:"filename"`
	ContentType string `json:"content_type"`
	Content     string `json:"content"`
	ContentID   string `json:"content_id"`
}

// Message represents an AhaSend message
type Message struct {
	Object string `json:"object"`
	// ID is the message's ID.
	ID uuid.UUID `json:"id"`
	// MessageID is the message's Message-ID header.
	MessageID            string          `json:"message_id"`
	CreatedAt            time.Time       `json:"created_at"`
	UpdatedAt            time.Time       `json:"updated_at"`
	SentAt               *time.Time      `json:"sent_at,omitempty"`
	DeliveredAt          *time.Time      `json:"delivered_at,omitempty"`
	RetainUntil          time.Time       `json:"retain_until"`
	Subject              string          `json:"subject"`
	Content              *string         `json:"content,omitempty"`
	ContentParsed        *ContentParsed  `json:"content_parsed,omitempty"`
	Tags                 []string        `json:"tags"`
	Sender               string          `json:"sender"`
	Recipient            string          `json:"recipient"`
	Direction            string          `json:"direction"`
	Status               string          `json:"status"`
	NumAttempts          int32           `json:"num_attempts"`
	DeliveryAttempts     []DeliveryEvent `json:"delivery_attempts"`
	IsBounceNotification bool            `json:"is_bounce_notification"`
	BounceClassification *string         `json:"bounce_classification,omitempty"`
	ClickCount           int32           `json:"click_count"`
	OpenCount            int32           `json:"open_count"`
	ReferenceMessageID   *int64          `json:"reference_message_id,omitempty"`
	DomainID             uuid.UUID       `json:"domain_id"`
	AccountID            uuid.UUID       `json:"account_id"`
	// TemplateID is the transactional template the message was sent from, or
	// nil for a message sent without one.
	TemplateID *uuid.UUID `json:"template_id,omitempty"`
}

// DeliveryEvent represents a single delivery attempt for a message
type DeliveryEvent struct {
	Time   time.Time `json:"time"`
	Log    string    `json:"log"`
	Status string    `json:"status"`
	// ReportMessageID is the ID of the bounce report message that the
	// recipient's mail server sent after it accepted this message. Fetch the
	// report with GetMessage. It is set only on the delivery event of an
	// out-of-band bounce.
	ReportMessageID *uuid.UUID `json:"report_message_id,omitempty"`
}

// MessageSchedule represents scheduling information for a message
type MessageSchedule struct {
	FirstAttempt *time.Time `json:"first_attempt,omitempty"`
	Expires      *time.Time `json:"expires,omitempty"`
}

// CreateMessageResponse represents the response when creating messages
type CreateMessageResponse struct {
	Object string                        `json:"object"`
	Data   []CreateSingleMessageResponse `json:"data"`
}

// CreateSingleMessageResponse represents the result of creating a single message
type CreateSingleMessageResponse struct {
	Object string `json:"object"`
	// ID is the ID of this recipient's message. GetMessage,
	// GetMessageByAPIID and CancelMessage take it. It is nil when the message
	// was not sent.
	ID *string `json:"id,omitempty"`
	// MessageIDHeader is the Message-ID header of this recipient's email, such
	// as <0199f1c2-7d3a-7b4e-9a1f-2c3d4e5f6a7b@example.com>. All copies of a
	// conversation send share it. It is nil when the message was not sent.
	MessageIDHeader *string          `json:"message_id_header,omitempty"`
	Recipient       common.Recipient `json:"recipient"`
	Status          string           `json:"status"`
	Error           *string          `json:"error,omitempty"`
	Schedule        *MessageSchedule `json:"schedule,omitempty"`
}
