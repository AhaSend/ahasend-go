package requests

// CreateRouteRequest represents a request to create a new email route.
type CreateRouteRequest struct {
	Name             string `json:"name"`
	URL              string `json:"url"`
	Recipient        string `json:"recipient"`
	Attachments      bool   `json:"attachments"`
	Headers          bool   `json:"headers"`
	GroupByMessageId bool   `json:"group_by_message_id"`
	StripReplies     bool   `json:"strip_replies"`
	Enabled          bool   `json:"enabled"`
}

// UpdateRouteRequest represents a request to update an existing email route.
type UpdateRouteRequest struct {
	Name             *string `json:"name,omitempty"`
	URL              *string `json:"url,omitempty"`
	Recipient        *string `json:"recipient,omitempty"`
	Attachments      *bool   `json:"attachments,omitempty"`
	Headers          *bool   `json:"headers,omitempty"`
	GroupByMessageId *bool   `json:"group_by_message_id,omitempty"`
	StripReplies     *bool   `json:"strip_replies,omitempty"`
	Enabled          *bool   `json:"enabled,omitempty"`
}
