package responses

import (
	"time"

	"github.com/google/uuid"
)

// APIKey represents an AhaSend API key
type APIKey struct {
	Object    string        `json:"object"`
	ID        uuid.UUID     `json:"id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	AccountID uuid.UUID     `json:"account_id"`
	Label     string        `json:"label"`
	PublicKey string        `json:"public_key"`
	Scopes    []APIKeyScope `json:"scopes"`
	// IPAllowList holds the source IPs allowed to authenticate with this key,
	// as canonical CIDR blocks (a bare address is stored as a /32 for IPv4 or
	// /128 for IPv6). Always present; an empty list means the key may be used
	// from any source IP.
	IPAllowList []string   `json:"ip_allow_list"`
	LastUsedAt  *time.Time `json:"last_used_at,omitempty"`
	SecretKey   *string    `json:"secret_key,omitempty"`
}

// APIKeyScope represents a scope granted to an API key
type APIKeyScope struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	APIKeyID  uuid.UUID  `json:"api_key_id"`
	Scope     string     `json:"scope"`
	DomainID  *uuid.UUID `json:"domain_id,omitempty"`
}
