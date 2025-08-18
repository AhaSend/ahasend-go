package common

// Tracking represents email tracking settings for opens and clicks.
type Tracking struct {
	// Whether to track opens
	Open *bool `json:"open,omitempty"`

	// Whether to track clicks
	Click *bool `json:"click,omitempty"`
}
