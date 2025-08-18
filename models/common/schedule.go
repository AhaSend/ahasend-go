package common

import (
	"time"
)

// MessageSchedule represents message delivery scheduling settings.
type MessageSchedule struct {
	// The time to make the first attempt for delivering the message (RFC3339 format)
	FirstAttempt *time.Time `json:"first_attempt,omitempty"`

	// Expire and drop the message if not delivered by this time (RFC3339 format)
	Expires *time.Time `json:"expires,omitempty"`
}
