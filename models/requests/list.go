package requests

import "github.com/google/uuid"

// Membership subscription statuses. Every status can be read and filtered on;
// only confirmed, unconfirmed, and unsubscribed can be written. complained is
// the mailbox provider's verdict, and a membership holding it cannot be changed.
const (
	ListContactStatusUnconfirmed  = "unconfirmed"
	ListContactStatusConfirmed    = "confirmed"
	ListContactStatusUnsubscribed = "unsubscribed"
	ListContactStatusComplained   = "complained"
)

// CreateListRequest represents a request to create a contact list. An omitted
// description stores an empty one, and omitted tags store an empty array.
type CreateListRequest struct {
	Name        string   `json:"name"`
	Description *string  `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

// UpdateListRequest represents a partial update to a contact list. A nil
// field is omitted and leaves the stored value unchanged. A pointer to ""
// clears the description, and a pointer to an empty slice clears the tags;
// a non-empty slice replaces them.
type UpdateListRequest struct {
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	Tags        *[]string `json:"tags,omitempty"`
}

// UpsertListContactRequest adds a contact to a list or updates its membership.
// A nil SubscriptionStatus names no status: a new membership is confirmed, and
// an existing one keeps its status. complained cannot be written.
type UpsertListContactRequest struct {
	SubscriptionStatus *string `json:"subscription_status,omitempty"`
}

// BatchAddListContactInput names one existing contact by exactly one of Email
// or ID. An entry naming neither or both is reported invalid in the response.
type BatchAddListContactInput struct {
	Email *string    `json:"email,omitempty"`
	ID    *uuid.UUID `json:"id,omitempty"`
}

// BatchAddListContactsRequest adds up to 1,000 existing contacts to one list.
type BatchAddListContactsRequest struct {
	Data []BatchAddListContactInput `json:"data"`
}

// GetListsParams represents the filters and pagination cursors accepted by
// the lists endpoint.
type GetListsParams struct {
	Limit  *int32
	After  *string
	Before *string
	Name   *string
}

// GetListContactsParams represents the filters and pagination cursors
// accepted by the list members endpoint. IncludeContacts embeds each whole
// contact and additionally requires the contacts:read scope.
type GetListContactsParams struct {
	Limit              *int32
	After              *string
	Before             *string
	SubscriptionStatus *string
	Email              *string
	IncludeContacts    *bool
}

// GetContactListsParams represents the filters and pagination cursors
// accepted by the contact's lists endpoint.
type GetContactListsParams struct {
	Limit              *int32
	After              *string
	Before             *string
	SubscriptionStatus *string
}
