package responses

import (
	"time"

	"github.com/google/uuid"
)

// Route represents an AhaSend route
type Route struct {
	Object           string    `json:"object"`
	ID               uuid.UUID `json:"id"`
	AccountID        uuid.UUID `json:"account_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Name             string    `json:"name"`
	URL              string    `json:"url"`
	Recipient        string    `json:"recipient,omitempty"`
	Attachments      bool      `json:"attachments"`
	Headers          bool      `json:"headers"`
	GroupByMessageID bool      `json:"group_by_message_id"`
	StripReplies     bool      `json:"strip_replies"`
	Secret           string    `json:"secret,omitempty"`
	Enabled          bool      `json:"enabled"`
}
