package requests

// CreateAPIKeyRequest represents a request to create a new API key.
type CreateAPIKeyRequest struct {
	Label  string   `json:"label"`
	Scopes []string `json:"scopes"`
}

// UpdateAPIKeyRequest represents a request to update an existing API key.
type UpdateAPIKeyRequest struct {
	Label  *string   `json:"label,omitempty"`
	Scopes *[]string `json:"scopes,omitempty"`
}

// Validate checks UpdateAPIKeyRequest client-side constraints.
func (r UpdateAPIKeyRequest) Validate() error {
	if err := validateAtLeastOneField(r.Label != nil || r.Scopes != nil); err != nil {
		return err
	}

	if err := validateOptionalString("label", r.Label, maxAPIKeyLabelLength); err != nil {
		return err
	}

	return validateOptionalNonEmptyStringSlice("scopes", r.Scopes)
}
