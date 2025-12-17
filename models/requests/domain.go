package requests

// CreateDomainRequest represents a request to create a new domain.
type CreateDomainRequest struct {
	// Domain is the fully qualified domain name to create.
	Domain string `json:"domain"`
	// DKIM private key is optional and must be a valid DKIM RSA private key with a minimum key length of 2048 bits if provided.
	DKIMPrivateKey *string `json:"dkim_private_key,omitempty"`
}
