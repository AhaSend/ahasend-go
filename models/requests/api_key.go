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
