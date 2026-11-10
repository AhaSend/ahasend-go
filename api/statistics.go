package api

import (
	"context"
	"fmt"
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

	applyRequestOptions(&config, opts)

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

	applyRequestOptions(&config, opts)

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

	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
GetEmailStatistics Get Email Statistics

# Returns the email counts of the account, or of each of its domains

The counts come in hourly, daily, weekly or monthly periods, for any time in
the account's history; they are not limited to the account's message metadata
retention period. They cover transactional, campaign and inbound mail. Inbound
mail is counted under the domain that received it. Sandbox mail is not
counted. The fields of responses.EmailStatisticsTotals say what each counter
counts.

A domain is identified by its name, and deleted domains are included. The API
key needs statistics-transactional:read:all or at least one
statistics-transactional:read:{domain} scope. Without SenderDomain, the
request selects every domain the key can read.

Periods are in UTC, and weeks start on Monday. The range includes FromTime
and excludes ToTime. It can be at most 31 days for hour, 366 days for day and
week, and 5 calendar years for month, before rounding. The API then rounds it
out to whole periods and returns the rounded range.

With Breakdown set to domain, a page holds up to Limit domains, each with all
its periods in the range. While Pagination.HasMore is true, pass
Pagination.NextCursor as After to fetch the next page. Limit and After are
accepted only with a breakdown. Totals covers the whole request on every page,
and is not always the sum of the rows.

This endpoint has the general rate limit, not the lower limit of the
transactional statistics, and the SDK applies its general rate limiter to it.

**Query Parameters:**
- `from_time`: Start of the range, included (RFC3339 format). Defaults to 30 days before `to_time`
- `to_time`: End of the range, excluded (RFC3339 format). Defaults to now
- `group_by`: Length of each period (hour, day, week, month). Defaults to day
- `sender_domain`: Comma-separated domain names of the account
- `breakdown`: `domain` returns the rows of each domain, paged by domain
- `limit`: Maximum number of domains on a page of a breakdown (1-50, default: 25)
- `after`: Cursor for the next page of a breakdown

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param params The query parameters; a nil field is not sent
	@param opts Optional request options
	@return EmailStatisticsResponse
*/
func (a *StatisticsAPIService) GetEmailStatistics(
	ctx context.Context,
	accountId uuid.UUID,
	params requests.GetEmailStatisticsParams,
	opts ...RequestOption,
) (*responses.EmailStatisticsResponse, *http.Response, error) {
	var result responses.EmailStatisticsResponse

	// Build query parameters. Unset ones are not sent, so the API defaults
	// apply; limit and after are refused without a breakdown.
	queryParams := url.Values{}
	if params.FromTime != nil {
		queryParams.Set("from_time", params.FromTime.Format(time.RFC3339))
	}
	if params.ToTime != nil {
		queryParams.Set("to_time", params.ToTime.Format(time.RFC3339))
	}
	if params.GroupBy != nil {
		queryParams.Set("group_by", *params.GroupBy)
	}
	if params.SenderDomain != nil {
		queryParams.Set("sender_domain", *params.SenderDomain)
	}
	if params.Breakdown != nil {
		queryParams.Set("breakdown", *params.Breakdown)
	}
	if params.Limit != nil {
		queryParams.Set("limit", fmt.Sprintf("%d", *params.Limit))
	}
	if params.After != nil {
		queryParams.Set("after", *params.After)
	}

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/statistics/emails",
		PathParams: map[string]string{
			"account_id": accountId.String(),
		},
		QueryParams: queryParams,
		Result:      &result,
	}

	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}
