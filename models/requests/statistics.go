package requests

import "time"

type GetDeliverabilityStatisticsParams struct {
	FromTime         *time.Time
	ToTime           *time.Time
	SenderDomain     *string
	RecipientDomains *string
	Tags             *string
	GroupBy          *string
}

type GetBounceStatisticsParams struct {
	FromTime         *time.Time
	ToTime           *time.Time
	SenderDomain     *string
	RecipientDomains *string
	Tags             *string
	GroupBy          *string
}

type GetDeliveryTimeStatisticsParams struct {
	FromTime         *time.Time
	ToTime           *time.Time
	SenderDomain     *string
	RecipientDomains *string
	Tags             *string
	GroupBy          *string
}

// GetEmailStatisticsParams represents the query parameters of the email
// statistics endpoint. A nil field is not sent, so the API default applies.
// The endpoint pages only the domains of a breakdown, forward only, so the
// params have Limit and After and no Before.
type GetEmailStatisticsParams struct {
	// FromTime is the start of the range, included. It is rounded down to the
	// start of its period. Defaults to 30 days before ToTime. It cannot be
	// before 2000-01-01.
	FromTime *time.Time
	// ToTime is the end of the range, excluded. It is rounded up to the start
	// of the next period unless it is already at the start of a period.
	// Defaults to now; a time after now is now.
	ToTime *time.Time
	// GroupBy is the length of each period, in UTC: hour, day, week or month.
	// Weeks start on Monday. The range can be at most 31 days for hour, 366
	// days for day and week, and 5 calendar years for month. Defaults to day.
	GroupBy *string
	// SenderDomain holds comma-separated domain names of the account, at most
	// 50, deleted domains included; the key must be able to read every name.
	// Defaults to every domain the key can read.
	SenderDomain *string
	// Breakdown set to domain returns the rows of each domain, paged by
	// domain. Without it, each row covers all the domains the request
	// selects.
	Breakdown *string
	// Limit is the maximum number of domains on a page of a breakdown (1-50,
	// default 25). Only with Breakdown set to domain.
	Limit *int32
	// After is the cursor for the next page of a breakdown: the NextCursor of
	// the previous response's Pagination. Only with Breakdown set to domain.
	After *string
}
