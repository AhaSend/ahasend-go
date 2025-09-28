package requests

import (
	"time"

	"github.com/AhaSend/ahasend-go/models/common"
)

// CreateSuppressionRequest represents a request to create a new email suppression.
type CreateSuppressionRequest struct {
	Email     string    `json:"email"`
	ExpiresAt time.Time `json:"expires_at"`
	Domain    *string   `json:"domain,omitempty"`
	Reason    *string   `json:"reason,omitempty"`
}

type GetSuppressionsParams struct {
	Email    *string
	Domain   *string
	FromDate *time.Time
	ToDate   *time.Time
	common.PaginationParams
}
