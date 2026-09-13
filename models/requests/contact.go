package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// CreateContactRequest represents a request to create an account-global contact.
type CreateContactRequest struct {
	Email        string         `json:"email"`
	FirstName    *string        `json:"first_name,omitempty"`
	LastName     *string        `json:"last_name,omitempty"`
	Status       *string        `json:"status,omitempty"`
	StatusReason *string        `json:"status_reason,omitempty"`
	Unsubscribed *bool          `json:"unsubscribed,omitempty"`
	Attributes   map[string]any `json:"attributes,omitempty"`
}

// Validate rejects null create attributes. Null attribute values have mutation
// semantics only for update and batch requests.
func (r CreateContactRequest) Validate() error {
	for key, value := range r.Attributes {
		encoded, err := json.Marshal(value)
		if err == nil && bytes.Equal(bytes.TrimSpace(encoded), []byte("null")) {
			return fmt.Errorf("attribute %q must not be null", key)
		}
	}

	return nil
}

// UpdateContactRequest represents a partial update to a contact. A nil
// attribute value removes that attribute from the contact.
type UpdateContactRequest struct {
	Email        *string        `json:"email,omitempty"`
	FirstName    *string        `json:"first_name,omitempty"`
	LastName     *string        `json:"last_name,omitempty"`
	Status       *string        `json:"status,omitempty"`
	StatusReason *string        `json:"status_reason,omitempty"`
	Unsubscribed *bool          `json:"unsubscribed,omitempty"`
	Attributes   map[string]any `json:"attributes,omitempty"`
}

// BatchUpsertContactInput represents one contact in a batch upsert. A nil
// attribute value removes that attribute when the contact already exists.
type BatchUpsertContactInput struct {
	Email        string         `json:"email"`
	FirstName    *string        `json:"first_name,omitempty"`
	LastName     *string        `json:"last_name,omitempty"`
	Status       *string        `json:"status,omitempty"`
	StatusReason *string        `json:"status_reason,omitempty"`
	Unsubscribed *bool          `json:"unsubscribed,omitempty"`
	Attributes   map[string]any `json:"attributes,omitempty"`
}

// BatchUpsertContactsRequest represents a synchronous batch upsert request.
type BatchUpsertContactsRequest struct {
	Data []BatchUpsertContactInput `json:"data"`
}

// GetContactsParams represents the filters and pagination cursors accepted by
// the contacts list endpoint.
type GetContactsParams struct {
	Limit      *int32
	After      *string
	Before     *string
	Email      *string
	Status     *string
	Subscribed *bool
	FromTime   *time.Time
	ToTime     *time.Time
}
