package responses

import (
	"time"

	"github.com/google/uuid"
)

// Contact represents an account-global AhaSend contact.
type Contact struct {
	Object           string         `json:"object"`
	ID               uuid.UUID      `json:"id"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	Email            string         `json:"email"`
	FirstName        string         `json:"first_name"`
	LastName         string         `json:"last_name"`
	Status           string         `json:"status"`
	StatusReason     string         `json:"status_reason"`
	Unsubscribed     bool           `json:"unsubscribed"`
	UnsubscribedAt   *time.Time     `json:"unsubscribed_at"`
	Attributes       map[string]any `json:"attributes"`
	ValidationStatus string         `json:"validation_status"`
	LastValidatedAt  *time.Time     `json:"last_validated_at"`
}

// ContactPagination represents bidirectional contact pagination metadata.
type ContactPagination struct {
	HasMore        bool    `json:"has_more"`
	NextCursor     *string `json:"next_cursor"`
	PreviousCursor *string `json:"previous_cursor"`
}

// PaginatedContactsResponse represents a page of contacts.
type PaginatedContactsResponse struct {
	Object     string            `json:"object"`
	Data       []Contact         `json:"data"`
	Pagination ContactPagination `json:"pagination"`
}

// BatchContactResult represents one ordered outcome from a batch upsert.
type BatchContactResult struct {
	Position int      `json:"position"`
	Email    string   `json:"email"`
	Outcome  string   `json:"outcome"`
	Contact  *Contact `json:"contact,omitempty"`
	Reason   string   `json:"reason,omitempty"`
}

// BatchUpsertContactsResponse represents the counts and ordered outcomes from
// a synchronous contact batch upsert.
type BatchUpsertContactsResponse struct {
	Object  string               `json:"object"`
	Created int                  `json:"created"`
	Updated int                  `json:"updated"`
	Failed  int                  `json:"failed"`
	Data    []BatchContactResult `json:"data"`
}
