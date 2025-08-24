package responses

import (
	"time"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/google/uuid"
)

// Message represents an AhaSend message
type Message struct {
	Object               string          `json:"object"`
	ID                   uuid.UUID       `json:"id"`
	MessageID            string          `json:"message_id"`
	CreatedAt            time.Time       `json:"created_at"`
	UpdatedAt            time.Time       `json:"updated_at"`
	SentAt               *time.Time      `json:"sent_at,omitempty"`
	DeliveredAt          *time.Time      `json:"delivered_at,omitempty"`
	RetainUntil          time.Time       `json:"retain_until"`
	Subject              string          `json:"subject"`
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
}

// DeliveryEvent represents a single delivery attempt for a message
type DeliveryEvent struct {
	Time   time.Time `json:"time"`
	Log    string    `json:"log"`
	Status string    `json:"status"`
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
	Object    string           `json:"object"`
	ID        *string          `json:"id,omitempty"`
	Recipient common.Recipient `json:"recipient"`
	Status    string           `json:"status"`
	Error     *string          `json:"error,omitempty"`
	Schedule  *MessageSchedule `json:"schedule,omitempty"`
}
