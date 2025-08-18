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
