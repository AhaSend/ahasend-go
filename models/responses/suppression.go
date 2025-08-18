package responses

import (
	"time"

	"github.com/google/uuid"
)

// Suppression represents an AhaSend suppression
type Suppression struct {
	Object    string    `json:"object"`
	ID        uint64    `json:"id"`
	AccountID uuid.UUID `json:"account_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	ExpiresAt time.Time `json:"expires_at"`
	Domain    string    `json:"domain,omitempty"`
	Reason    string    `json:"reason,omitempty"`
}
