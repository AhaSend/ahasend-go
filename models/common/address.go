package common

// SenderAddress represents a sender email address with optional display name.
type SenderAddress struct {
	Email string  `json:"email"`
	Name  *string `json:"name,omitempty"`
}

// Recipient represents an email recipient with optional display name and substitution data.
type Recipient struct {
	Email string  `json:"email"`
	Name  *string `json:"name,omitempty"`
	// Substitutions are this recipient's template values. In the HTML and AMP
	// bodies each value is HTML-escaped; {{ name|safe }} inserts it as HTML.
	// The subject, the headers and the text body get values as they are.
	Substitutions map[string]interface{} `json:"substitutions,omitempty"`
}
