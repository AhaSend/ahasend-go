package requests

// CreateSMTPCredentialRequest represents a request to create a new SMTP credential.
type CreateSMTPCredentialRequest struct {
	Name    string   `json:"name"`
	Scope   string   `json:"scope"`
	Sandbox bool     `json:"sandbox"`
	Domains []string `json:"domains,omitempty"`
}
