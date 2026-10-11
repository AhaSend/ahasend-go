package responses

import (
	"time"

	"github.com/google/uuid"
)

// Domain pause reasons. The set is open: AhaSend can add reasons at any time,
// so keep a default branch for values you do not know.
const (
	// DomainPauseReasonBounceRate means too many recent emails from the domain bounced.
	DomainPauseReasonBounceRate = "bounce_rate"
)

// Domain represents an AhaSend domain
type Domain struct {
	Object                   string      `json:"object"`
	ID                       uuid.UUID   `json:"id"`
	CreatedAt                time.Time   `json:"created_at"`
	UpdatedAt                time.Time   `json:"updated_at"`
	Domain                   string      `json:"domain"`
	AccountID                uuid.UUID   `json:"account_id"`
	DNSRecords               []DNSRecord `json:"dns_records"`
	LastDNSCheckAt           *time.Time  `json:"last_dns_check_at,omitempty"`
	DNSValid                 bool        `json:"dns_valid"`
	TrackingSubdomain        *string     `json:"tracking_subdomain,omitempty"`
	ReturnPathSubdomain      *string     `json:"return_path_subdomain,omitempty"`
	SubscriptionSubdomain    *string     `json:"subscription_subdomain,omitempty"`
	MediaSubdomain           *string     `json:"media_subdomain,omitempty"`
	DKIMRotationIntervalDays *int        `json:"dkim_rotation_interval_days,omitempty"`
	// DKIMSelector is the per-domain DKIM selector override, or nil when none
	// is set. It is not always the selector used for signing: a selector locked
	// after DNS verification takes precedence.
	DKIMSelector  *string `json:"dkim_selector,omitempty"`
	RotationReady bool    `json:"rotation_ready"`
	DSNRecipient  *string `json:"dsn_recipient,omitempty"`
	// SendingType is common.DomainSendingTypeTransactional or
	// common.DomainSendingTypeMarketing. It is "" when the server does not
	// return the field, as servers before domain sending types do not.
	SendingType string `json:"sending_type"`
	// Paused reports whether sending from the domain is paused. A paused domain
	// refuses new email and cannot be deleted or renamed.
	Paused bool `json:"paused"`
	// PausedAt is when sending from the domain was paused, or nil when it is not paused.
	PausedAt *time.Time `json:"paused_at,omitempty"`
	// PauseReason is why sending from the domain is paused, or nil when it is
	// not paused. See DomainPauseReasonBounceRate; other values can appear.
	PauseReason *string `json:"pause_reason,omitempty"`
}

// DNSRecord represents a DNS record required for domain verification
type DNSRecord struct {
	Type       string  `json:"type"`
	Label      *string `json:"label,omitempty"`
	Host       string  `json:"host"`
	Content    string  `json:"content"`
	Required   bool    `json:"required"`
	Propagated bool    `json:"propagated"`
}
