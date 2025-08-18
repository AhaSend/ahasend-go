package api

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/AhaSend/ahasend-go/models/responses"
	"github.com/google/uuid"
)

// StatisticsAPIService StatisticsAPI service
type StatisticsAPIService service

/*
GetBounceStatistics Get Bounce Statistics

# Returns transactional bounce statistics grouped by classification

**Query Parameters:**
- `from_time`: Filter statistics after this datetime (RFC3339 format)
- `to_time`: Filter statistics before this datetime (RFC3339 format)
- `sender_domain`: Filter by sender domain
- `recipient_domains`: Filter by a comma separated list of recipient domains
- `tags`: Filter by tags (comma-separated)
- `group_by`: Group by time period (hour, day, week, month)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param params The query parameters for filtering the statistics
	@param opts Optional request options
	@return BounceStatisticsResponse
*/
func (a *StatisticsAPIService) GetBounceStatistics(
	ctx context.Context,
	accountId uuid.UUID,
	params requests.GetBounceStatisticsParams,
	opts ...RequestOption,
) (*responses.BounceStatisticsResponse, *http.Response, error) {
	var result responses.BounceStatisticsResponse

	// Build query parameters
	queryParams := url.Values{}
	if params.FromTime != nil {
		queryParams.Set("from_time", params.FromTime.Format(time.RFC3339))
	}
	if params.ToTime != nil {
		queryParams.Set("to_time", params.ToTime.Format(time.RFC3339))
	}
	if params.SenderDomain != nil {
		queryParams.Set("sender_domain", *params.SenderDomain)
	}
	if params.RecipientDomains != nil {
		queryParams.Set("recipient_domains", *params.RecipientDomains)
	}
	if params.Tags != nil {
		queryParams.Set("tags", *params.Tags)
	}
	if params.GroupBy != nil {
		queryParams.Set("group_by", *params.GroupBy)
	} else {
		queryParams.Set("group_by", "day") // Default value
	}

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/statistics/transactional/bounce",
		PathParams: map[string]string{
			"account_id": accountId.String(),
		},
		QueryParams: queryParams,
		Result:      &result,
	}

	for _, opt := range opts {
		opt(&config)
	}

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
GetDeliverabilityStatistics Get Deliverability Statistics

# Returns transactional deliverability statistics grouped by status

**Query Parameters:**
- `from_time`: Filter statistics after this datetime (RFC3339 format)
- `to_time`: Filter statistics before this datetime (RFC3339 format)
- `sender_domain`: Filter by sender domain
- `recipient_domains`: Filter by a comma separated list of recipient domains
- `tags`: Filter by tags (comma-separated)
- `group_by`: Group by time period (hour, day, week, month)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param params The filtering query params
	@param opts Optional request options
	@return DeliverabilityStatisticsResponse
*/
func (a *StatisticsAPIService) GetDeliverabilityStatistics(
	ctx context.Context,
	accountId uuid.UUID,
	params requests.GetDeliverabilityStatisticsParams,
	opts ...RequestOption,
) (*responses.DeliverabilityStatisticsResponse, *http.Response, error) {
	var result responses.DeliverabilityStatisticsResponse

	// Build query parameters
	queryParams := url.Values{}
	if params.FromTime != nil {
		queryParams.Set("from_time", params.FromTime.Format(time.RFC3339))
	}
	if params.ToTime != nil {
		queryParams.Set("to_time", params.ToTime.Format(time.RFC3339))
	}
	if params.SenderDomain != nil {
		queryParams.Set("sender_domain", *params.SenderDomain)
	}
	if params.RecipientDomains != nil {
		queryParams.Set("recipient_domains", *params.RecipientDomains)
	}
	if params.Tags != nil {
		queryParams.Set("tags", *params.Tags)
	}
	if params.GroupBy != nil {
		queryParams.Set("group_by", *params.GroupBy)
	} else {
		queryParams.Set("group_by", "day") // Default value
	}

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/statistics/transactional/deliverability",
		PathParams: map[string]string{
			"account_id": accountId.String(),
		},
		QueryParams: queryParams,
		Result:      &result,
	}

	for _, opt := range opts {
		opt(&config)
	}

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
GetDeliveryTimeStatistics Get Delivery Time Statistics

# Returns transactional delivery time statistics grouped by recipient domain

**Query Parameters:**
- `from_time`: Filter statistics after this datetime (RFC3339 format)
- `to_time`: Filter statistics before this datetime (RFC3339 format)
- `sender_domain`: Filter by sender domain
- `recipient_domains`: Filter by a comma separated list of recipient domains
- `tags`: Filter by tags (comma-separated)
- `group_by`: Group by time period (hour, day, week, month)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param params Query parameters for filtering the results
	@param opts Optional request options
	@return DeliveryTimeStatisticsResponse
*/
func (a *StatisticsAPIService) GetDeliveryTimeStatistics(
	ctx context.Context,
	accountId uuid.UUID,
	params requests.GetDeliveryTimeStatisticsParams,
	opts ...RequestOption,
) (*responses.DeliveryTimeStatisticsResponse, *http.Response, error) {
	var result responses.DeliveryTimeStatisticsResponse

	// Build query parameters
	queryParams := url.Values{}
	if params.FromTime != nil {
		queryParams.Set("from_time", params.FromTime.Format(time.RFC3339))
	}
	if params.ToTime != nil {
		queryParams.Set("to_time", params.ToTime.Format(time.RFC3339))
	}
	if params.SenderDomain != nil {
		queryParams.Set("sender_domain", *params.SenderDomain)
	}
	if params.RecipientDomains != nil {
		queryParams.Set("recipient_domains", *params.RecipientDomains)
	}
	if params.Tags != nil {
		queryParams.Set("tags", *params.Tags)
	}
	if params.GroupBy != nil {
		queryParams.Set("group_by", *params.GroupBy)
	} else {
		queryParams.Set("group_by", "day") // Default value
	}

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/statistics/transactional/delivery-time",
		PathParams: map[string]string{
			"account_id": accountId.String(),
		},
		QueryParams: queryParams,
		Result:      &result,
	}

	for _, opt := range opts {
		opt(&config)
	}

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}
