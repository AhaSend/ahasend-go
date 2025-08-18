package responses

import (
	"time"

	"github.com/google/uuid"
)

// APIKey represents an AhaSend API key
type APIKey struct {
	Object     string        `json:"object"`
	ID         uuid.UUID     `json:"id"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at"`
	AccountID  uuid.UUID     `json:"account_id"`
	Label      string        `json:"label"`
	PublicKey  string        `json:"public_key"`
	Scopes     []APIKeyScope `json:"scopes"`
	LastUsedAt *time.Time    `json:"last_used_at,omitempty"`
	SecretKey  *string       `json:"secret_key,omitempty"`
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
