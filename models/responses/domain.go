package responses

import (
	"time"

	"github.com/google/uuid"
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
	RotationReady            bool        `json:"rotation_ready"`
}

// DNSRecord represents a DNS record required for domain verification
type DNSRecord struct {
	Type       string `json:"type"`
	Host       string `json:"host"`
	Content    string `json:"content"`
	Required   bool   `json:"required"`
	Propagated bool   `json:"propagated"`
}
