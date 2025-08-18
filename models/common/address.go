package common

// SenderAddress represents a sender email address with optional display name.
type SenderAddress struct {
	Email string  `json:"email"`
	Name  *string `json:"name,omitempty"`
}

// Recipient represents an email recipient with optional display name and substitution data.
type Recipient struct {
	Email         string                 `json:"email"`
	Name          *string                `json:"name,omitempty"`
	Substitutions map[string]interface{} `json:"substitutions,omitempty"`
}
