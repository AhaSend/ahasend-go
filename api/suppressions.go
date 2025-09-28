package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/AhaSend/ahasend-go/models/responses"
	"github.com/google/uuid"
)

// SuppressionsAPIService SuppressionsAPI service
type SuppressionsAPIService service

/*
CreateSuppression Create Suppression

# Creates a new suppression for an email address

Validation Requirements:
- `email` must be a valid email address
- `expires_at` must be in RFC3339 format
- `domain` is optional - if not provided, applies to all account domains

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param request CreateSuppressionRequest - The suppression details to create
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return CreateSuppressionResponse, *http.Response, error
*/
func (a *SuppressionsAPIService) CreateSuppression(
	ctx context.Context,
	accountId uuid.UUID,
	request requests.CreateSuppressionRequest,
	opts ...RequestOption,
) (*responses.CreateSuppressionResponse, *http.Response, error) {
	var result responses.CreateSuppressionResponse

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/suppressions",
		PathParams: map[string]string{
			"account_id": accountId.String(),
		},
		Body:   request,
		Result: &result,
	}

	// Apply options
	for _, opt := range opts {
		opt(&config)
	}

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
DeleteAllSuppressions Delete All Suppressions

# Deletes all suppressions for the account

Query Parameters:
- `domain`: Optional domain filter to delete suppressions for specific domain only

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param domain Optional domain filter (optional)
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SuccessResponse, *http.Response, error
*/
func (a *SuppressionsAPIService) DeleteAllSuppressions(
	ctx context.Context,
	accountId uuid.UUID,
	domain *string,
	opts ...RequestOption,
) (*common.SuccessResponse, *http.Response, error) {
	var result common.SuccessResponse

	// Build query parameters
	queryParams := url.Values{}
	if domain != nil {
		queryParams.Set("domain", *domain)
	}

	config := RequestConfig{
		Method:       http.MethodDelete,
		PathTemplate: "/v2/accounts/{account_id}/suppressions/all",
		PathParams: map[string]string{
			"account_id": accountId.String(),
		},
		QueryParams: queryParams,
		Result:      &result,
	}

	// Apply options
	for _, opt := range opts {
		opt(&config)
	}

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
DeleteSuppression Delete Suppression

# Deletes a specific suppression by email address

Query Parameters:
- `domain`: Optional domain filter to delete suppression for specific domain only

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param email Email address
	@param domain Optional domain filter (optional)
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SuccessResponse, *http.Response, error
*/
func (a *SuppressionsAPIService) DeleteSuppression(
	ctx context.Context,
	accountId uuid.UUID,
	email string,
	domain *string,
	opts ...RequestOption,
) (*common.SuccessResponse, *http.Response, error) {
	var result common.SuccessResponse

	// Build query parameters
	queryParams := url.Values{}
	queryParams.Set("email", email)
	if domain != nil {
		queryParams.Set("domain", *domain)
	}

	config := RequestConfig{
		Method:       http.MethodDelete,
		PathTemplate: "/v2/accounts/{account_id}/suppressions",
		PathParams: map[string]string{
			"account_id": accountId.String(),
		},
		QueryParams: queryParams,
		Result:      &result,
	}

	// Apply options
	for _, opt := range opts {
		opt(&config)
	}

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
GetSuppressions Get Suppressions

# Returns a list of suppressions for the account

Query Parameters:
- `domain`: Filter by domain (optional)
- `from_date`: Filter suppressions created after this date (RFC3339 format)
- `to_date`: Filter suppressions created before this date (RFC3339 format)
- `limit`: Maximum number of items to return (1-100, default: 100)
- `cursor`: Pagination cursor for the next page (backward compatibility)
- `after`: Get items after this cursor (forward pagination)
- `before`: Get items before this cursor (backward pagination)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param params GetSuppressionsParams - query parameters
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return PaginatedSuppressionsResponse, *http.Response, error
*/
func (a *SuppressionsAPIService) GetSuppressions(
	ctx context.Context,
	accountId uuid.UUID,
	params requests.GetSuppressionsParams,
	opts ...RequestOption,
) (*responses.PaginatedSuppressionsResponse, *http.Response, error) {
	var result responses.PaginatedSuppressionsResponse

	// Build query parameters
	queryParams := url.Values{}
	if params.Email != nil {
		queryParams.Set("email", *params.Email)
	}
	if params.Domain != nil {
		queryParams.Set("domain", *params.Domain)
	}
	if params.FromDate != nil {
		queryParams.Set("from_date", params.FromDate.Format(time.RFC3339))
	}
	if params.ToDate != nil {
		queryParams.Set("to_date", params.ToDate.Format(time.RFC3339))
	}
	// Handle pagination parameters
	if params.Limit != nil {
		queryParams.Set("limit", fmt.Sprintf("%d", *params.Limit))
	} else {
		queryParams.Set("limit", "100") // Default value
	}

	// Handle pagination parameters - prioritize after/before over cursor for backward compatibility
	if params.After != nil {
		queryParams.Set("after", *params.After)
	} else if params.Before != nil {
		queryParams.Set("before", *params.Before)
	} else if params.Cursor != nil {
		queryParams.Set("cursor", *params.Cursor)
	}

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/suppressions",
		PathParams: map[string]string{
			"account_id": accountId.String(),
		},
		QueryParams: queryParams,
		Result:      &result,
	}

	// Apply options
	for _, opt := range opts {
		opt(&config)
	}

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}
