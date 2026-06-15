package requests

// CreateSubAccountRequest represents a request to create a sub account.
type CreateSubAccountRequest struct {
	Name          string `json:"name"`
	Website       string `json:"website"`
	MonthlyCredit *int64 `json:"monthly_credit,omitempty"`
}

// Validate checks CreateSubAccountRequest client-side constraints.
func (r CreateSubAccountRequest) Validate() error {
	if err := validateRequiredString("name", r.Name, maxRequestNameLength); err != nil {
		return err
	}

	if err := validateRequiredString("website", r.Website, maxRequestWebsiteLength); err != nil {
		return err
	}

	return validateOptionalInt64Range("monthly_credit", r.MonthlyCredit, minMonthlyCredit, maxMonthlyCredit)
}

// UpdateSubAccountRequest represents a request to update a sub account.
type UpdateSubAccountRequest struct {
	Name          *string `json:"name,omitempty"`
	Website       *string `json:"website,omitempty"`
	MonthlyCredit *int64  `json:"monthly_credit,omitempty"`
}

// Validate checks UpdateSubAccountRequest client-side constraints.
func (r UpdateSubAccountRequest) Validate() error {
	if err := validateAtLeastOneField(r.Name != nil || r.Website != nil || r.MonthlyCredit != nil); err != nil {
		return err
	}

	if err := validateOptionalString("name", r.Name, maxRequestNameLength); err != nil {
		return err
	}

	if err := validateOptionalString("website", r.Website, maxRequestWebsiteLength); err != nil {
		return err
	}

	return validateOptionalInt64Range("monthly_credit", r.MonthlyCredit, minMonthlyCredit, maxMonthlyCredit)
}

// SuspendSubAccountRequest represents a request to suspend a sub account.
type SuspendSubAccountRequest struct {
	Reason string `json:"reason"`
}

// Validate checks SuspendSubAccountRequest client-side constraints.
func (r SuspendSubAccountRequest) Validate() error {
	return validateRequiredString("reason", r.Reason, maxSuspensionReasonLength)
}
