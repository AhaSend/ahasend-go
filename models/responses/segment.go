package responses

import (
	"time"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/google/uuid"
)

// Segment represents a saved set of rules over the account's contacts, and
// the contacts that match them. A field documented as nil is sent as null.
type Segment struct {
	Object    string    `json:"object"`
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is when the segment was last saved.
	UpdatedAt time.Time `json:"updated_at"`
	// Name is the segment name. Names are not unique within an account.
	Name string `json:"name"`
	// Description is empty when none was given.
	Description string `json:"description"`
	// Definition holds the segment's rules.
	Definition common.SegmentDefinition `json:"definition"`
	// Timezone is the IANA time zone that relative times and days in the
	// definition are read in.
	Timezone string `json:"timezone"`
	// State is one of the common.SegmentState constants.
	State string `json:"state"`
	// FailureReason says why the segment failed. It is nil unless State is
	// common.SegmentStateFailed.
	FailureReason *string `json:"failure_reason"`
	// MemberCount is the number of contacts that match the rules, whatever
	// their consent. A campaign reaches only those of them who may get
	// marketing email; GetSegment counts them as MarketableCount.
	MemberCount int `json:"member_count"`
	// LastEvaluatedAt is when the members were last checked against the
	// whole account, nil until the first build ends.
	LastEvaluatedAt *time.Time `json:"last_evaluated_at"`
}

// SegmentDetail represents a segment as GetSegment returns it, with the
// number of members a campaign to it alone would reach.
type SegmentDetail struct {
	Segment
	// MarketableCount is the number of members who are enabled, not
	// unsubscribed, not found invalid, and either confirmed on at least one
	// list or on no list at all. It is nil while State is not
	// common.SegmentStateReady, and when the members cannot be counted
	// within 10 seconds.
	MarketableCount *int `json:"marketable_count"`
}

// SegmentPreview represents what a definition matches now.
type SegmentPreview struct {
	Object string `json:"object"`
	// Count is the number of contacts the definition matches, whatever their
	// consent. It is nil when they cannot be counted within 10 seconds.
	Count *int `json:"count"`
	// Contacts holds up to 10 of the matching contacts, in no particular
	// order. It is returned also when Count is nil, and is empty when the
	// contacts cannot be read within 10 seconds.
	Contacts []SegmentPreviewContact `json:"contacts"`
}

// SegmentPreviewContact represents one contact that a previewed definition
// matches.
type SegmentPreviewContact struct {
	ContactID uuid.UUID `json:"contact_id"`
	// Email is the contact's normalized email address.
	Email string `json:"email"`
}

// SegmentContact represents one member of a segment.
type SegmentContact struct {
	Object    string    `json:"object"`
	SegmentID uuid.UUID `json:"segment_id"`
	ContactID uuid.UUID `json:"contact_id"`
	// Email is the contact's normalized email address.
	Email string `json:"email"`
	// EnteredAt is when the contact became a member.
	EnteredAt time.Time `json:"entered_at"`
}

// PaginatedSegmentsResponse represents a page of segments.
type PaginatedSegmentsResponse struct {
	Object     string            `json:"object"`
	Data       []Segment         `json:"data"`
	Pagination ContactPagination `json:"pagination"`
}

// PaginatedSegmentContactsResponse represents a page of segment members.
type PaginatedSegmentContactsResponse struct {
	Object     string            `json:"object"`
	Data       []SegmentContact  `json:"data"`
	Pagination ContactPagination `json:"pagination"`
}
