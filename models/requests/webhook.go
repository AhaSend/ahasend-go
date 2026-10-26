package requests

// CreateWebhookRequest represents a request to create a new webhook.
type CreateWebhookRequest struct {
	Name                 string `json:"name"`
	URL                  string `json:"url"`
	Enabled              *bool  `json:"enabled,omitempty"`
	OnReception          bool   `json:"on_reception"`
	OnDelivered          bool   `json:"on_delivered"`
	OnTransientError     bool   `json:"on_transient_error"`
	OnFailed             bool   `json:"on_failed"`
	OnBounced            bool   `json:"on_bounced"`
	OnSuppressed         bool   `json:"on_suppressed"`
	OnOpened             bool   `json:"on_opened"`
	OnClicked            bool   `json:"on_clicked"`
	OnUnsubscribed       bool   `json:"on_unsubscribed"`
	OnComplained         bool   `json:"on_complained"`
	OnSuppressionCreated bool   `json:"on_suppression_created"`
	OnSuppressionDeleted bool   `json:"on_suppression_deleted"`
	OnDnsError           bool   `json:"on_dns_error"`
	// The contact events need the contact read scopes on the API key that
	// creates the webhook.
	OnContactCreated         bool `json:"on_contact_created"`
	OnContactUpdated         bool `json:"on_contact_updated"`
	OnContactDeleted         bool `json:"on_contact_deleted"`
	OnContactAddedToList     bool `json:"on_contact_added_to_list"`
	OnContactRemovedFromList bool `json:"on_contact_removed_from_list"`
	OnContactSubscribed      bool `json:"on_contact_subscribed"`
	OnContactUnsubscribed    bool `json:"on_contact_unsubscribed"`
	OnContactComplained      bool `json:"on_contact_complained"`
	// Origins lists the origins, common.WebhookOrigin*, of the contact and
	// suppression events the webhook gets. An empty Origins is left out of
	// the request and means every origin; to stop contact or suppression
	// events, turn them off. All four origins are saved as every origin, so
	// an origin added later reaches the webhook too.
	Origins []string  `json:"origins,omitempty"`
	Scope   string    `json:"scope"`
	Domains *[]string `json:"domains,omitempty"`
}

// UpdateWebhookRequest represents a request to update an existing webhook.
type UpdateWebhookRequest struct {
	Name                 *string `json:"name,omitempty"`
	URL                  *string `json:"url,omitempty"`
	Enabled              *bool   `json:"enabled,omitempty"`
	OnReception          *bool   `json:"on_reception,omitempty"`
	OnDelivered          *bool   `json:"on_delivered,omitempty"`
	OnTransientError     *bool   `json:"on_transient_error,omitempty"`
	OnFailed             *bool   `json:"on_failed,omitempty"`
	OnBounced            *bool   `json:"on_bounced,omitempty"`
	OnSuppressed         *bool   `json:"on_suppressed,omitempty"`
	OnOpened             *bool   `json:"on_opened,omitempty"`
	OnClicked            *bool   `json:"on_clicked,omitempty"`
	OnUnsubscribed       *bool   `json:"on_unsubscribed,omitempty"`
	OnComplained         *bool   `json:"on_complained,omitempty"`
	OnSuppressionCreated *bool   `json:"on_suppression_created,omitempty"`
	OnSuppressionDeleted *bool   `json:"on_suppression_deleted,omitempty"`
	OnDnsError           *bool   `json:"on_dns_error,omitempty"`
	// The contact events need the contact read scopes on the API key.
	OnContactCreated         *bool `json:"on_contact_created,omitempty"`
	OnContactUpdated         *bool `json:"on_contact_updated,omitempty"`
	OnContactDeleted         *bool `json:"on_contact_deleted,omitempty"`
	OnContactAddedToList     *bool `json:"on_contact_added_to_list,omitempty"`
	OnContactRemovedFromList *bool `json:"on_contact_removed_from_list,omitempty"`
	OnContactSubscribed      *bool `json:"on_contact_subscribed,omitempty"`
	OnContactUnsubscribed    *bool `json:"on_contact_unsubscribed,omitempty"`
	OnContactComplained      *bool `json:"on_contact_complained,omitempty"`
	// Origins replaces the origins, common.WebhookOrigin*, of the contact and
	// suppression events the webhook gets. An empty Origins is left out of
	// the request and keeps the stored value; send all four origins to get
	// every origin again.
	Origins []string  `json:"origins,omitempty"`
	Scope   *string   `json:"scope,omitempty"`
	Domains *[]string `json:"domains,omitempty"`
}
