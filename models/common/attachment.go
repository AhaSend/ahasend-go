package common

// Attachment represents a file attachment for email messages.
type Attachment struct {
	// If true, data must be encoded using base64. Otherwise, data will be interpreted as UTF-8
	Base64 bool `json:"base64,omitempty"`

	// Either plaintext or base64 encoded attachment data (depending on base64 field)
	Data string `json:"data"`

	// The MIME type of the attachment
	ContentType string `json:"content_type"`

	// The Content-ID of the attachment for inline images
	ContentID *string `json:"content_id,omitempty"`

	// The filename of the attachment
	FileName string `json:"file_name"`
}
