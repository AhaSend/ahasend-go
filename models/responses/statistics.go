package responses

import (
	"time"

	"github.com/AhaSend/ahasend-go/models/common"
)

// DeliverabilityStatistics represents AhaSend email deliverability statistics for a time bucket
type DeliverabilityStatistics struct {
	// Time bucket boundaries (required)
	FromTimestamp time.Time `json:"from_timestamp"`
	ToTimestamp   time.Time `json:"to_timestamp"`

	// Email delivery pipeline metrics (optional - use pointer types for response models)
	ReceptionCount  int `json:"reception_count"`
	DeliveredCount  int `json:"delivered_count"`
	DeferredCount   int `json:"deferred_count"`
	BouncedCount    int `json:"bounced_count"`
	FailedCount     int `json:"failed_count"`
	SuppressedCount int `json:"suppressed_count"`
	OpenedCount     int `json:"opened_count"`
	ClickedCount    int `json:"clicked_count"`
}

// DeliverabilityStatisticsResponse represents the API response containing deliverability statistics
type DeliverabilityStatisticsResponse struct {
	Object string                     `json:"object"`
	Data   []DeliverabilityStatistics `json:"data"`
}

// Bounce represents a bounce classification with count data
type Bounce struct {
	Classification string `json:"classification"`
	Count          int    `json:"count"`
}

// BounceStatistics represents AhaSend bounce statistics for a time bucket
type BounceStatistics struct {
	// Time bucket boundaries (required)
	FromTimestamp time.Time `json:"from_timestamp"`
	ToTimestamp   time.Time `json:"to_timestamp"`

	// Bounce classifications and counts (required)
	Bounces []Bounce `json:"bounces"`
}

// BounceStatisticsResponse represents the API response containing bounce statistics
type BounceStatisticsResponse struct {
	Object string             `json:"object"`
	Data   []BounceStatistics `json:"data"`
}

// DeliveryTime represents delivery time statistics for a specific recipient domain
type DeliveryTime struct {
	// The recipient domain (optional)
	RecipientDomain *string `json:"recipient_domain,omitempty"`

	// The average time from reception to delivery in seconds (optional)
	DeliveryTime *float64 `json:"delivery_time,omitempty"`

	// Number of delivered messages included in the average.
	Count int64 `json:"count"`
}

// DeliveryTimeStatistics represents delivery time statistics for a time bucket
type DeliveryTimeStatistics struct {
	// Time bucket boundaries (required)
	FromTimestamp time.Time `json:"from_timestamp"`
	ToTimestamp   time.Time `json:"to_timestamp"`

	// Overall delivery metrics (required)
	AvgDeliveryTime float64 `json:"avg_delivery_time"`
	DeliveredCount  int     `json:"delivered_count"`

	// Per-domain delivery times (optional)
	DeliveryTimes []DeliveryTime `json:"delivery_times,omitempty"`
}

// DeliveryTimeStatisticsResponse represents the API response containing delivery time statistics
type DeliveryTimeStatisticsResponse struct {
	Object string                   `json:"object"`
	Data   []DeliveryTimeStatistics `json:"data"`
}

// EmailStatisticsTotals holds the email counts of one period, or of the whole
// request. Every counter is at least 0.
type EmailStatisticsTotals struct {
	// ReceptionCount is the number of messages that AhaSend accepted, inbound
	// mail included. Billing counts the same messages.
	ReceptionCount int64 `json:"reception_count"`
	// DeliveredCount is the number of messages delivered.
	DeliveredCount int64 `json:"delivered_count"`
	// BouncedCount is the number of messages bounced.
	BouncedCount int64 `json:"bounced_count"`
	// SuppressedCount is the number of messages not sent because the
	// recipient was suppressed.
	SuppressedCount int64 `json:"suppressed_count"`
	// ComplainedCount is the number of spam complaints.
	ComplainedCount int64 `json:"complained_count"`
	// DeferredCount is the number of temporary delivery failures. It counts
	// events, and one message can add more than one, so it differs from
	// DeliverabilityStatistics.DeferredCount, which counts messages.
	DeferredCount int64 `json:"deferred_count"`
	// OpenedCount is the number of opens. It counts events, and one message
	// can add more than one, so it differs from
	// DeliverabilityStatistics.OpenedCount, which counts messages. Opens are
	// not fully accurate; AhaSend makes a reasonable effort to filter out bots
	// and scanners.
	OpenedCount int64 `json:"opened_count"`
	// ClickedCount is the number of clicks. It counts events, and one message
	// can add more than one, so it differs from
	// DeliverabilityStatistics.ClickedCount, which counts messages. Clicks are
	// not fully accurate; AhaSend makes a reasonable effort to filter out bots
	// and scanners.
	ClickedCount int64 `json:"clicked_count"`
}

// EmailStatistics holds the email counts of one period, and of one domain in
// a breakdown.
type EmailStatistics struct {
	// FromTimestamp is the start of the period.
	FromTimestamp time.Time `json:"from_timestamp"`
	// ToTimestamp is the end of the period, excluded, which is the start of
	// the next period.
	ToTimestamp time.Time `json:"to_timestamp"`
	// SenderDomain is the domain name of the row in a breakdown, and nil
	// without a breakdown.
	SenderDomain *string `json:"sender_domain"`
	EmailStatisticsTotals
}

// EmailStatisticsResponse represents the API response of GetEmailStatistics.
type EmailStatisticsResponse struct {
	// Object is "email_statistics".
	Object string `json:"object"`
	// FromTime is the start of the rounded range, included.
	FromTime time.Time `json:"from_time"`
	// ToTime is the end of the rounded range, excluded.
	ToTime time.Time `json:"to_time"`
	// GroupBy is the length of each period: hour, day, week or month.
	GroupBy string `json:"group_by"`
	// Breakdown is "domain" for a breakdown by domain, and nil without a
	// breakdown.
	Breakdown *string `json:"breakdown"`
	// Totals holds the counts of the whole request: the whole rounded range
	// and every domain the request selects, on every page of a breakdown.
	Totals EmailStatisticsTotals `json:"totals"`
	// Data holds, without a breakdown, one row for each period, in time
	// order. With a breakdown, it holds the rows of the domains of this page,
	// by domain name, then by time.
	Data []EmailStatistics `json:"data"`
	// Pagination pages the domains of a breakdown. NextCursor is set when
	// HasMore is true; PreviousCursor is never set. Without a breakdown,
	// HasMore is false.
	Pagination common.PaginationInfo `json:"pagination"`
}
