package common

// Retention represents data retention settings for email messages.
type Retention struct {
	// Number of days to retain metadata
	Metadata *int32 `json:"metadata,omitempty"`

	// Number of days to retain data
	Data *int32 `json:"data,omitempty"`
}
