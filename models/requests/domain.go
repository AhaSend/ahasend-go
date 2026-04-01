package requests

// CreateDomainRequest represents a request to create a new domain.
type CreateDomainRequest struct {
	// Domain is the fully qualified domain name to create.
	Domain string `json:"domain"`
	// DKIMPrivateKey is optional and must be a valid DKIM RSA private key with a minimum key length of 2048 bits if provided.
	DKIMPrivateKey *string `json:"dkim_private_key,omitempty"`
	// TrackingSubdomain is an optional custom tracking subdomain. Omit to use the default.
	TrackingSubdomain *string `json:"tracking_subdomain,omitempty"`
	// ReturnPathSubdomain is an optional custom return-path subdomain. Omit to use the default.
	ReturnPathSubdomain *string `json:"return_path_subdomain,omitempty"`
	// SubscriptionSubdomain is an optional custom subscription management subdomain. Omit to use the default.
	SubscriptionSubdomain *string `json:"subscription_subdomain,omitempty"`
	// MediaSubdomain is an optional custom media subdomain. Omit to use the default.
	MediaSubdomain *string `json:"media_subdomain,omitempty"`
	// DKIMRotationIntervalDays is an optional custom DKIM rotation interval in days. Only supported for managed DNS domains on eligible plans.
	DKIMRotationIntervalDays *int `json:"dkim_rotation_interval_days,omitempty"`
}

// UpdateDomainRequest represents a request to update domain settings.
type UpdateDomainRequest struct {
	// TrackingSubdomain is a custom tracking subdomain. Omit to leave unchanged.
	TrackingSubdomain *string `json:"tracking_subdomain,omitempty"`
	// ReturnPathSubdomain is a custom return-path subdomain. Omit to leave unchanged.
	ReturnPathSubdomain *string `json:"return_path_subdomain,omitempty"`
	// SubscriptionSubdomain is a custom subscription management subdomain. Omit to leave unchanged.
	SubscriptionSubdomain *string `json:"subscription_subdomain,omitempty"`
	// MediaSubdomain is a custom media subdomain. Omit to leave unchanged.
	MediaSubdomain *string `json:"media_subdomain,omitempty"`
	// DKIMRotationIntervalDays is a custom DKIM rotation interval in days. Omit to leave unchanged. Only supported for managed DNS domains on eligible plans.
	DKIMRotationIntervalDays *int `json:"dkim_rotation_interval_days,omitempty"`
}
