package responses

import (
	"time"

	"github.com/google/uuid"
)

// SubAccount represents an AhaSend sub account.
type SubAccount struct {
	Object          string     `json:"object"`
	ID              uuid.UUID  `json:"id"`
	ParentAccountID uuid.UUID  `json:"parent_account_id"`
	CreatedAt       time.Time  `json:"created_at"`
	Name            string     `json:"name"`
	Website         string     `json:"website"`
	Status          string     `json:"status"`
	MonthlyCredit   int64      `json:"monthly_credit"`
	DomainCount     int64      `json:"domain_count"`
	MemberCount     int64      `json:"member_count"`
	LastActivityAt  *time.Time `json:"last_activity_at"`
}

// SubAccountUsageBreakdown represents one usage allocation bucket.
type SubAccountUsageBreakdown struct {
	AccountID      *uuid.UUID `json:"account_id,omitempty"`
	Name           *string    `json:"name,omitempty"`
	ReceptionCount int64      `json:"reception_count"`
	AllocatedCost  float64    `json:"allocated_cost"`
}

// SubAccountUsageBillingPeriod represents the billing period for usage allocation.
type SubAccountUsageBillingPeriod struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

// SubAccountUsageResponse represents allocated parent/sub-account usage.
type SubAccountUsageResponse struct {
	BillingPeriod      SubAccountUsageBillingPeriod `json:"billing_period"`
	Currency           string                       `json:"currency"`
	AllocationMethod   string                       `json:"allocation_method"`
	AllocationNote     string                       `json:"allocation_note"`
	Parent             SubAccountUsageBreakdown     `json:"parent"`
	SubAccounts        []SubAccountUsageBreakdown   `json:"sub_accounts"`
	RemovedSubAccounts SubAccountUsageBreakdown     `json:"removed_sub_accounts"`
	Total              SubAccountUsageBreakdown     `json:"total"`
}
