package responses

import (
	"time"

	"github.com/google/uuid"
)

// Suppression represents an AhaSend suppression
type Suppression struct {
	Object    string    `json:"object"`
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Email     string    `json:"email"`
	ExpiresAt time.Time `json:"expires_at"`
	Domain    string    `json:"domain,omitempty"`
	Reason    string    `json:"reason,omitempty"`
	// Protected is set by the API when the recipient made the decision
	// themselves (unsubscribed or reported spam). Deleting all suppressions
	// keeps protected ones. It does not affect whether the address can be mailed.
	Protected bool `json:"protected"`
}

// CreateSuppressionResponse represents the response when creating suppressions
type CreateSuppressionResponse struct {
	Object string        `json:"object"`
	Data   []Suppression `json:"data"`
}
