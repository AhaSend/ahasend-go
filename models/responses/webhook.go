package responses

import (
	"time"

	"github.com/google/uuid"
)

// Webhook represents an AhaSend webhook
type Webhook struct {
	Object                   string    `json:"object"`
	ID                       uuid.UUID `json:"id"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
	Name                     string    `json:"name"`
	URL                      string    `json:"url"`
	Enabled                  bool      `json:"enabled"`
	Secret                   string    `json:"secret,omitempty"`
	OnReception              bool      `json:"on_reception"`
	OnDelivered              bool      `json:"on_delivered"`
	OnTransientError         bool      `json:"on_transient_error"`
	OnFailed                 bool      `json:"on_failed"`
	OnBounced                bool      `json:"on_bounced"`
	OnSuppressed             bool      `json:"on_suppressed"`
	OnOpened                 bool      `json:"on_opened"`
	OnClicked                bool      `json:"on_clicked"`
	OnUnsubscribed           bool      `json:"on_unsubscribed"`
	OnComplained             bool      `json:"on_complained"`
	OnSuppressionCreated     bool      `json:"on_suppression_created"`
	OnSuppressionDeleted     bool      `json:"on_suppression_deleted"`
	OnDNSError               bool      `json:"on_dns_error"`
	OnContactCreated         bool      `json:"on_contact_created"`
	OnContactUpdated         bool      `json:"on_contact_updated"`
	OnContactDeleted         bool      `json:"on_contact_deleted"`
	OnContactAddedToList     bool      `json:"on_contact_added_to_list"`
	OnContactRemovedFromList bool      `json:"on_contact_removed_from_list"`
	OnContactSubscribed      bool      `json:"on_contact_subscribed"`
	OnContactUnsubscribed    bool      `json:"on_contact_unsubscribed"`
	OnContactComplained      bool      `json:"on_contact_complained"`
	// Origins lists the origins, common.WebhookOrigin*, of the contact and
	// suppression events the webhook gets. Nil means every origin.
	Origins                []string   `json:"origins"`
	Scope                  string     `json:"scope"`
	Domains                []string   `json:"domains"`
	ErrorCount             uint64     `json:"error_count"`
	SuccessCount           uint64     `json:"success_count"`
	ErrorsSinceLastSuccess int        `json:"errors_since_last_success"`
	LastRequestAt          *time.Time `json:"last_request_at"`
}
