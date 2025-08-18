package responses

import (
	"time"

	"github.com/google/uuid"
)

// Account represents an AhaSend account
type Account struct {
	Object    string    `json:"object"`
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`

	// Optional account information
	Website *string `json:"website,omitempty"`
	About   *string `json:"about,omitempty"`

	// Email behavior settings
	TrackOpens               *bool  `json:"track_opens,omitempty"`
	TrackClicks              *bool  `json:"track_clicks,omitempty"`
	RejectBadRecipients      *bool  `json:"reject_bad_recipients,omitempty"`
	RejectMistypedRecipients *bool  `json:"reject_mistyped_recipients,omitempty"`
	MessageMetadataRetention *int32 `json:"message_metadata_retention,omitempty"`
	MessageDataRetention     *int32 `json:"message_data_retention,omitempty"`
}

// UserAccount represents a user's relationship to an account
type UserAccount struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	AccountID uuid.UUID `json:"account_id"`
	Role      string    `json:"role"`
}

// AccountMembersResponse represents the API response containing account members
type AccountMembersResponse struct {
	Object string        `json:"object"`
	Data   []UserAccount `json:"data"`
}
