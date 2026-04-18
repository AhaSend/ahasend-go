package responses

import (
	"time"

	"github.com/google/uuid"
)

// Webhook represents an AhaSend webhook
type Webhook struct {
	Object                 string     `json:"object"`
	ID                     uuid.UUID  `json:"id"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              time.Time  `json:"updated_at"`
	Name                   string     `json:"name"`
	URL                    string     `json:"url"`
	Enabled                bool       `json:"enabled"`
	Secret                 string     `json:"secret,omitempty"`
	OnReception            bool       `json:"on_reception"`
	OnDelivered            bool       `json:"on_delivered"`
	OnTransientError       bool       `json:"on_transient_error"`
	OnFailed               bool       `json:"on_failed"`
	OnBounced              bool       `json:"on_bounced"`
	OnSuppressed           bool       `json:"on_suppressed"`
	OnOpened               bool       `json:"on_opened"`
	OnClicked              bool       `json:"on_clicked"`
	OnSuppressionCreated   bool       `json:"on_suppression_created"`
	OnDNSError             bool       `json:"on_dns_error"`
	Scope                  string     `json:"scope"`
	Domains                []string   `json:"domains"`
	ErrorCount             uint64     `json:"error_count"`
	SuccessCount           uint64     `json:"success_count"`
	ErrorsSinceLastSuccess int        `json:"errors_since_last_success"`
	LastRequestAt          *time.Time `json:"last_request_at"`
}
