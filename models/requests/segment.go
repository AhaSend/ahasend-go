package requests

import "github.com/AhaSend/ahasend-go/models/common"

// GetSegmentsParams represents the pagination cursors accepted by the
// segments endpoint. After and Before cannot be combined.
type GetSegmentsParams struct {
	Limit  *int32
	After  *string
	Before *string
}

// GetSegmentContactsParams represents the pagination cursors accepted by the
// segment members endpoint. After and Before cannot be combined.
type GetSegmentContactsParams struct {
	Limit  *int32
	After  *string
	Before *string
}

// GetContactSegmentsParams represents the pagination cursors accepted by the
// contact's segments endpoint. After and Before cannot be combined.
type GetContactSegmentsParams struct {
	Limit  *int32
	After  *string
	Before *string
}

// CreateSegmentRequest represents a request to create a segment.
type CreateSegmentRequest struct {
	// Name is the segment name, 1 to 255 characters after trimming. Names
	// are not unique within an account.
	Name string `json:"name"`
	// Description is trimmed and at most 1000 characters. A nil Description
	// stores an empty one.
	Description *string `json:"description,omitempty"`
	// Definition holds the segment's rules.
	Definition common.SegmentDefinition `json:"definition"`
	// Timezone is the IANA time zone, such as "Europe/Berlin", that relative
	// times and days are read in. A nil Timezone or a pointer to "" stores
	// "UTC".
	Timezone *string `json:"timezone,omitempty"`
}

// UpdateSegmentRequest represents a partial update to a segment. A nil field
// is omitted and leaves the stored value unchanged. A change of Definition or
// Timezone builds the segment's members again, and the segment is building
// until that ends; a change of only Name or Description leaves its members
// and its state as they are. Any update of a failed segment, an empty one
// included, builds it again.
type UpdateSegmentRequest struct {
	// Name is trimmed and bounded as on create.
	Name *string `json:"name,omitempty"`
	// Description is trimmed and bounded as on create.
	Description *string `json:"description,omitempty"`
	// Definition replaces the stored rules as a whole.
	Definition *common.SegmentDefinition `json:"definition,omitempty"`
	// Timezone is the IANA time zone, as on create. A pointer to "" stores
	// "UTC".
	Timezone *string `json:"timezone,omitempty"`
}

// PreviewSegmentRequest represents a definition to try. Nothing is saved.
type PreviewSegmentRequest struct {
	// Definition holds the rules to try.
	Definition common.SegmentDefinition `json:"definition"`
	// Timezone is the IANA time zone that relative times and days are read
	// in. A nil Timezone or a pointer to "" reads them in "UTC".
	Timezone *string `json:"timezone,omitempty"`
}
