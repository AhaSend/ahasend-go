package responses

import (
	"time"

	"github.com/google/uuid"
)

// Batch add outcomes. added and already_member carry the membership;
// not_found and invalid carry a reason and count as failed.
const (
	BatchListContactOutcomeAdded         = "added"
	BatchListContactOutcomeAlreadyMember = "already_member"
	BatchListContactOutcomeNotFound      = "not_found"
	BatchListContactOutcomeInvalid       = "invalid"
)

// ContactList represents a static, private, single opt-in list of contacts,
// the unit a campaign targets.
type ContactList struct {
	Object      string    `json:"object"`
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Tags        []string  `json:"tags"`
	// ContactCount is the number of members a campaign to this list would
	// reach: confirmed memberships whose contact is enabled, not unsubscribed
	// from everything, and not found invalid.
	ContactCount int `json:"contact_count"`
}

// EmbeddedContactList is the list a membership carries. It has no
// contact_count.
type EmbeddedContactList struct {
	Object      string    `json:"object"`
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Tags        []string  `json:"tags"`
}

// ListContact represents one contact's membership of one list. It has no ID
// of its own and is addressed by the list and contact pair. Contact is set
// only by GetListContacts with IncludeContacts, and List only by
// GetContactLists.
type ListContact struct {
	Object             string               `json:"object"`
	ListID             uuid.UUID            `json:"list_id"`
	ContactID          uuid.UUID            `json:"contact_id"`
	Email              string               `json:"email"`
	SubscriptionStatus string               `json:"subscription_status"`
	SubscribedAt       *time.Time           `json:"subscribed_at"`
	UnsubscribedAt     *time.Time           `json:"unsubscribed_at"`
	CreatedAt          time.Time            `json:"created_at"`
	UpdatedAt          time.Time            `json:"updated_at"`
	Contact            *Contact             `json:"contact,omitempty"`
	List               *EmbeddedContactList `json:"list,omitempty"`
}

// PaginatedContactListsResponse represents a page of lists.
type PaginatedContactListsResponse struct {
	Object     string            `json:"object"`
	Data       []ContactList     `json:"data"`
	Pagination ContactPagination `json:"pagination"`
}

// PaginatedListContactsResponse represents a page of list memberships.
type PaginatedListContactsResponse struct {
	Object     string            `json:"object"`
	Data       []ListContact     `json:"data"`
	Pagination ContactPagination `json:"pagination"`
}

// BatchListContactResult represents one ordered outcome from a batch add.
// Email or ID echoes the identifier the entry used; an entry that named no
// single identifier carries only Position.
type BatchListContactResult struct {
	Position   int          `json:"position"`
	Email      string       `json:"email,omitempty"`
	ID         *uuid.UUID   `json:"id,omitempty"`
	Outcome    string       `json:"outcome"`
	Membership *ListContact `json:"membership,omitempty"`
	Reason     string       `json:"reason,omitempty"`
}

// BatchAddListContactsResponse represents the counts and ordered outcomes
// from a synchronous batch add. Skipped counts already_member entries, and
// Failed counts not_found and invalid ones.
type BatchAddListContactsResponse struct {
	Object  string                   `json:"object"`
	Added   int                      `json:"added"`
	Skipped int                      `json:"skipped"`
	Failed  int                      `json:"failed"`
	Data    []BatchListContactResult `json:"data"`
}
