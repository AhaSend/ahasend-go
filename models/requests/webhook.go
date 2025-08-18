package requests

// CreateWebhookRequest represents a request to create a new webhook.
type CreateWebhookRequest struct {
	Name                 string    `json:"name"`
	URL                  string    `json:"url"`
	Enabled              bool      `json:"enabled"`
	OnReception          bool      `json:"on_reception"`
	OnDelivered          bool      `json:"on_delivered"`
	OnTransientError     bool      `json:"on_transient_error"`
	OnFailed             bool      `json:"on_failed"`
	OnBounced            bool      `json:"on_bounced"`
	OnSuppressed         bool      `json:"on_suppressed"`
	OnOpened             bool      `json:"on_opened"`
	OnClicked            bool      `json:"on_clicked"`
	OnSuppressionCreated bool      `json:"on_suppression_created"`
	OnDnsError           bool      `json:"on_dns_error"`
	Scope                string    `json:"scope"`
	Domains              *[]string `json:"domains,omitempty"`
}

// UpdateWebhookRequest represents a request to update an existing webhook.
type UpdateWebhookRequest struct {
	Name                 *string   `json:"name,omitempty"`
	URL                  *string   `json:"url,omitempty"`
	Enabled              *bool     `json:"enabled,omitempty"`
	OnReception          *bool     `json:"on_reception,omitempty"`
	OnDelivered          *bool     `json:"on_delivered,omitempty"`
	OnTransientError     *bool     `json:"on_transient_error,omitempty"`
	OnFailed             *bool     `json:"on_failed,omitempty"`
	OnBounced            *bool     `json:"on_bounced,omitempty"`
	OnSuppressed         *bool     `json:"on_suppressed,omitempty"`
	OnOpened             *bool     `json:"on_opened,omitempty"`
	OnClicked            *bool     `json:"on_clicked,omitempty"`
	OnSuppressionCreated *bool     `json:"on_suppression_created,omitempty"`
	OnDnsError           *bool     `json:"on_dns_error,omitempty"`
	Scope                *string   `json:"scope,omitempty"`
	Domains              *[]string `json:"domains,omitempty"`
}
