package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/AhaSend/ahasend-go/models/responses"
	"github.com/google/uuid"
)

// RoutesAPIService RoutesAPI service
type RoutesAPIService service

/*
CreateRoute Create Route

# Creates a new route for inbound email routing

Validation Requirements:
- `name` must be a unique route name within the account
- `destination` must be a valid webhook URL or email address
- `patterns` array must contain at least one routing pattern

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param request CreateRouteRequest - The route details to create
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return Route, *http.Response, error
*/
func (a *RoutesAPIService) CreateRoute(
	ctx context.Context,
	accountId uuid.UUID,
	request requests.CreateRouteRequest,
	opts ...RequestOption,
) (*responses.Route, *http.Response, error) {
	var result responses.Route

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/routes",
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
DeleteRoute Delete Route

Deletes a route

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param routeId Route ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SuccessResponse, *http.Response, error
*/
func (a *RoutesAPIService) DeleteRoute(
	ctx context.Context,
	accountId uuid.UUID,
	routeId uuid.UUID,
	opts ...RequestOption,
) (*common.SuccessResponse, *http.Response, error) {
	var result common.SuccessResponse

	config := RequestConfig{
		Method:       http.MethodDelete,
		PathTemplate: "/v2/accounts/{account_id}/routes/{route_id}",
		PathParams: map[string]string{
			"account_id": accountId.String(),
			"route_id":   routeId.String(),
		},
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
GetRoute Get Route

Returns a specific route by ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param routeId Route ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return Route, *http.Response, error
*/
func (a *RoutesAPIService) GetRoute(
	ctx context.Context,
	accountId uuid.UUID,
	routeId uuid.UUID,
	opts ...RequestOption,
) (*responses.Route, *http.Response, error) {
	var result responses.Route

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/routes/{route_id}",
		PathParams: map[string]string{
			"account_id": accountId.String(),
			"route_id":   routeId.String(),
		},
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
GetRoutes Get Routes

Returns a list of routes for the account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param pagination Pagination parameters (optional)
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return PaginatedRoutesResponse, *http.Response, error
*/
func (a *RoutesAPIService) GetRoutes(
	ctx context.Context,
	accountId uuid.UUID,
	pagination *common.PaginationParams,
	opts ...RequestOption,
) (*responses.PaginatedRoutesResponse, *http.Response, error) {
	var result responses.PaginatedRoutesResponse

	// Build query parameters
	queryParams := url.Values{}

	// Handle pagination parameters
	if pagination != nil {
		if pagination.Limit != nil {
			queryParams.Set("limit", fmt.Sprintf("%d", *pagination.Limit))
		} else {
			queryParams.Set("limit", "100") // Default value
		}

		// Handle pagination parameters - prioritize after/before over cursor for backward compatibility
		if pagination.After != nil {
			queryParams.Set("after", *pagination.After)
		} else if pagination.Before != nil {
			queryParams.Set("before", *pagination.Before)
		} else if pagination.Cursor != nil {
			queryParams.Set("cursor", *pagination.Cursor)
		}
	} else {
		queryParams.Set("limit", "100") // Default value
	}

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/routes",
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
UpdateRoute Update Route

Updates an existing route

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param routeId Route ID
	@param request UpdateRouteRequest - The route details to update
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return Route, *http.Response, error
*/
func (a *RoutesAPIService) UpdateRoute(
	ctx context.Context,
	accountId uuid.UUID,
	routeId uuid.UUID,
	request requests.UpdateRouteRequest,
	opts ...RequestOption,
) (*responses.Route, *http.Response, error) {
	var result responses.Route

	config := RequestConfig{
		Method:       http.MethodPut,
		PathTemplate: "/v2/accounts/{account_id}/routes/{route_id}",
		PathParams: map[string]string{
			"account_id": accountId.String(),
			"route_id":   routeId.String(),
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
