package requests

// CreateAPIKeyRequest represents a request to create a new API key.
type CreateAPIKeyRequest struct {
	Label  string   `json:"label"`
	Scopes []string `json:"scopes"`
	// IPAllowList optionally restricts the source IPs allowed to authenticate
	// with this key. Each entry is a CIDR block (e.g. "203.0.113.0/24") or a
	// bare IPv4/IPv6 address (stored as a /32 or /128). Leave nil or empty to
	// keep the key usable from any IP.
	IPAllowList []string `json:"ip_allow_list,omitempty"`
}

// UpdateAPIKeyRequest represents a request to update an existing API key.
type UpdateAPIKeyRequest struct {
	Label  *string   `json:"label,omitempty"`
	Scopes *[]string `json:"scopes,omitempty"`
	// IPAllowList replaces the key's allowed source IPs. The pointer encodes a
	// tri-state: leave it nil to keep the existing list unchanged, point it at
	// an empty slice (&[]string{}) to clear the list (key usable from any IP),
	// or point it at a non-empty slice of CIDR blocks / bare addresses to
	// replace the list.
	IPAllowList *[]string `json:"ip_allow_list,omitempty"`
}

// Validate checks UpdateAPIKeyRequest client-side constraints.
func (r UpdateAPIKeyRequest) Validate() error {
	if err := validateAtLeastOneField(r.Label != nil || r.Scopes != nil || r.IPAllowList != nil); err != nil {
		return err
	}

	if err := validateOptionalString("label", r.Label, maxAPIKeyLabelLength); err != nil {
		return err
	}

	return validateOptionalNonEmptyStringSlice("scopes", r.Scopes)
}
