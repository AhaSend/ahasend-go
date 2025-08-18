package responses

import (
	"time"
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
