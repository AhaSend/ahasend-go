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

// APIKeysAPIService APIKeysAPI service
type APIKeysAPIService service

/*
CreateAPIKey Create API Key

# Creates a new API key with the specified scopes

Validation Requirements:
- `label` must be provided and non-empty
- `scopes` must contain at least one valid scope

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param request CreateAPIKeyRequest - The API key details to create
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return ModelAPIKey, *http.Response, error
*/
func (a *APIKeysAPIService) CreateAPIKey(
	ctx context.Context,
	accountId uuid.UUID,
	request requests.CreateAPIKeyRequest,
	opts ...RequestOption,
) (*responses.APIKey, *http.Response, error) {
	var result responses.APIKey

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/api-keys",
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
DeleteAPIKey Delete API Key

Deletes an API key

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param keyId API Key ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SuccessResponse, *http.Response, error
*/
func (a *APIKeysAPIService) DeleteAPIKey(
	ctx context.Context,
	accountId uuid.UUID,
	keyId uuid.UUID,
	opts ...RequestOption,
) (*common.SuccessResponse, *http.Response, error) {
	var result common.SuccessResponse

	config := RequestConfig{
		Method:       http.MethodDelete,
		PathTemplate: "/v2/accounts/{account_id}/api-keys/{key_id}",
		PathParams: map[string]string{
			"account_id": accountId.String(),
			"key_id":     keyId.String(),
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
GetAPIKey Get API Key

Returns a specific API key by ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param keyId API Key ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return ModelAPIKey, *http.Response, error
*/
func (a *APIKeysAPIService) GetAPIKey(
	ctx context.Context,
	accountId uuid.UUID,
	keyId uuid.UUID,
	opts ...RequestOption,
) (*responses.APIKey, *http.Response, error) {
	var result responses.APIKey

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/api-keys/{key_id}",
		PathParams: map[string]string{
			"account_id": accountId.String(),
			"key_id":     keyId.String(),
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
GetAPIKeys Get API Keys

Returns a list of API keys for the account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param pagination Pagination parameters (optional)
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return PaginatedAPIKeysResponse, *http.Response, error
*/
func (a *APIKeysAPIService) GetAPIKeys(
	ctx context.Context,
	accountId uuid.UUID,
	pagination *common.PaginationParams,
	opts ...RequestOption,
) (*responses.PaginatedAPIKeysResponse, *http.Response, error) {
	var result responses.PaginatedAPIKeysResponse

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
		PathTemplate: "/v2/accounts/{account_id}/api-keys",
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
UpdateAPIKey Update API Key

Updates an existing API key's label and scopes

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param keyId API Key ID
	@param request UpdateAPIKeyRequest - The API key details to update
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return ModelAPIKey, *http.Response, error
*/
func (a *APIKeysAPIService) UpdateAPIKey(
	ctx context.Context,
	accountId uuid.UUID,
	keyId uuid.UUID,
	request requests.UpdateAPIKeyRequest,
	opts ...RequestOption,
) (*responses.APIKey, *http.Response, error) {
	var result responses.APIKey

	config := RequestConfig{
		Method:       http.MethodPut,
		PathTemplate: "/v2/accounts/{account_id}/api-keys/{key_id}",
		PathParams: map[string]string{
			"account_id": accountId.String(),
			"key_id":     keyId.String(),
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
