package responses

import (
	"time"

	"github.com/google/uuid"
)

// Domain represents an AhaSend domain
type Domain struct {
	Object         string      `json:"object"`
	ID             uuid.UUID   `json:"id"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	Domain         string      `json:"domain"`
	AccountID      uuid.UUID   `json:"account_id"`
	DNSRecords     []DNSRecord `json:"dns_records"`
	LastDNSCheckAt *time.Time  `json:"last_dns_check_at,omitempty"`
	DNSValid       bool        `json:"dns_valid"`
}

// DNSRecord represents a DNS record required for domain verification
type DNSRecord struct {
	Type       string `json:"type"`
	Host       string `json:"host"`
	Content    string `json:"content"`
	Required   bool   `json:"required"`
	Propagated bool   `json:"propagated"`
}
