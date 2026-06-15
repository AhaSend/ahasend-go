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

// SubAccountsAPIService SubAccountsAPI service
type SubAccountsAPIService service

/*
ListSubAccounts List Sub Accounts

Returns a cursor-paginated list of sub accounts under the parent account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Parent account ID
	@param pagination Pagination parameters (optional)
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return PaginatedSubAccountsResponse, *http.Response, error
*/
func (a *SubAccountsAPIService) ListSubAccounts(
	ctx context.Context,
	accountId uuid.UUID,
	pagination *common.PaginationParams,
	opts ...RequestOption,
) (*responses.PaginatedSubAccountsResponse, *http.Response, error) {
	var result responses.PaginatedSubAccountsResponse

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
		PathTemplate: "/v2/accounts/{account_id}/sub-accounts",
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
CreateSubAccount Create Sub Account

Creates a new sub account under the parent account.

Validation Requirements:
- `name` must be provided and non-empty
- `website` must be provided and non-empty
- `monthly_credit` must be between 0 and 1000000000 when provided

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Parent account ID
	@param request CreateSubAccountRequest - The sub-account details to create
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SubAccount, *http.Response, error
*/
func (a *SubAccountsAPIService) CreateSubAccount(
	ctx context.Context,
	accountId uuid.UUID,
	request requests.CreateSubAccountRequest,
	opts ...RequestOption,
) (*responses.SubAccount, *http.Response, error) {
	var result responses.SubAccount

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/sub-accounts",
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
GetSubAccountsUsage Get Sub-Account Usage

Returns current billing-period usage and proportional allocated cost for the parent and active sub accounts.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Parent account ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SubAccountUsageResponse, *http.Response, error
*/
func (a *SubAccountsAPIService) GetSubAccountsUsage(
	ctx context.Context,
	accountId uuid.UUID,
	opts ...RequestOption,
) (*responses.SubAccountUsageResponse, *http.Response, error) {
	var result responses.SubAccountUsageResponse

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/sub-accounts/usage",
		PathParams: map[string]string{
			"account_id": accountId.String(),
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
GetSubAccount Get Sub Account

Returns a specific sub account under the parent account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Parent account ID
	@param subAccountId Sub account ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SubAccount, *http.Response, error
*/
func (a *SubAccountsAPIService) GetSubAccount(
	ctx context.Context,
	accountId uuid.UUID,
	subAccountId uuid.UUID,
	opts ...RequestOption,
) (*responses.SubAccount, *http.Response, error) {
	var result responses.SubAccount

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/sub-accounts/{sub_account_id}",
		PathParams: map[string]string{
			"account_id":     accountId.String(),
			"sub_account_id": subAccountId.String(),
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
UpdateSubAccount Update Sub Account

Updates a sub account's editable settings.

Validation Requirements:
- at least one editable field must be provided
- `name` and `website` must be non-empty when provided
- `monthly_credit` must be between 0 and 1000000000 when provided

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Parent account ID
	@param subAccountId Sub account ID
	@param request UpdateSubAccountRequest - The sub-account details to update
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SubAccount, *http.Response, error
*/
func (a *SubAccountsAPIService) UpdateSubAccount(
	ctx context.Context,
	accountId uuid.UUID,
	subAccountId uuid.UUID,
	request requests.UpdateSubAccountRequest,
	opts ...RequestOption,
) (*responses.SubAccount, *http.Response, error) {
	var result responses.SubAccount

	config := RequestConfig{
		Method:       http.MethodPut,
		PathTemplate: "/v2/accounts/{account_id}/sub-accounts/{sub_account_id}",
		PathParams: map[string]string{
			"account_id":     accountId.String(),
			"sub_account_id": subAccountId.String(),
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
DeleteSubAccount Delete Sub Account

Soft-deletes a sub account under the parent account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Parent account ID
	@param subAccountId Sub account ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SuccessResponse, *http.Response, error
*/
func (a *SubAccountsAPIService) DeleteSubAccount(
	ctx context.Context,
	accountId uuid.UUID,
	subAccountId uuid.UUID,
	opts ...RequestOption,
) (*common.SuccessResponse, *http.Response, error) {
	var result common.SuccessResponse

	config := RequestConfig{
		Method:       http.MethodDelete,
		PathTemplate: "/v2/accounts/{account_id}/sub-accounts/{sub_account_id}",
		PathParams: map[string]string{
			"account_id":     accountId.String(),
			"sub_account_id": subAccountId.String(),
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
SuspendSubAccount Suspend Sub Account

Suspends a sub account under the parent account.

Validation Requirements:
- `reason` must be provided and non-empty

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Parent account ID
	@param subAccountId Sub account ID
	@param request SuspendSubAccountRequest - The suspension details
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SubAccount, *http.Response, error
*/
func (a *SubAccountsAPIService) SuspendSubAccount(
	ctx context.Context,
	accountId uuid.UUID,
	subAccountId uuid.UUID,
	request requests.SuspendSubAccountRequest,
	opts ...RequestOption,
) (*responses.SubAccount, *http.Response, error) {
	var result responses.SubAccount

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/sub-accounts/{sub_account_id}/suspend",
		PathParams: map[string]string{
			"account_id":     accountId.String(),
			"sub_account_id": subAccountId.String(),
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
UnsuspendSubAccount Unsuspend Sub Account

Unsuspends a sub account under the parent account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Parent account ID
	@param subAccountId Sub account ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SubAccount, *http.Response, error
*/
func (a *SubAccountsAPIService) UnsuspendSubAccount(
	ctx context.Context,
	accountId uuid.UUID,
	subAccountId uuid.UUID,
	opts ...RequestOption,
) (*responses.SubAccount, *http.Response, error) {
	var result responses.SubAccount

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/sub-accounts/{sub_account_id}/unsuspend",
		PathParams: map[string]string{
			"account_id":     accountId.String(),
			"sub_account_id": subAccountId.String(),
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
ListSubAccountAPIKeys List Sub-Account API Keys

Returns a cursor-paginated list of API keys owned by a sub account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Parent account ID
	@param subAccountId Sub account ID
	@param pagination Pagination parameters (optional)
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return PaginatedAPIKeysResponse, *http.Response, error
*/
func (a *SubAccountsAPIService) ListSubAccountAPIKeys(
	ctx context.Context,
	accountId uuid.UUID,
	subAccountId uuid.UUID,
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
		PathTemplate: "/v2/accounts/{account_id}/sub-accounts/{sub_account_id}/api-keys",
		PathParams: map[string]string{
			"account_id":     accountId.String(),
			"sub_account_id": subAccountId.String(),
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
CreateSubAccountAPIKey Create Sub-Account API Key

Creates an API key owned by the sub account. SecretKey is one-time and present only on create responses and exact idempotent replay responses.

Validation Requirements:
- `label` must be provided and non-empty
- `scopes` must contain at least one valid scope

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Parent account ID
	@param subAccountId Sub account ID
	@param request CreateAPIKeyRequest - The API key details to create
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return APIKey, *http.Response, error
*/
func (a *SubAccountsAPIService) CreateSubAccountAPIKey(
	ctx context.Context,
	accountId uuid.UUID,
	subAccountId uuid.UUID,
	request requests.CreateAPIKeyRequest,
	opts ...RequestOption,
) (*responses.APIKey, *http.Response, error) {
	var result responses.APIKey

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/sub-accounts/{sub_account_id}/api-keys",
		PathParams: map[string]string{
			"account_id":     accountId.String(),
			"sub_account_id": subAccountId.String(),
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
GetSubAccountAPIKey Get Sub-Account API Key

Returns a specific API key owned by a sub account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Parent account ID
	@param subAccountId Sub account ID
	@param keyId API key ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return APIKey, *http.Response, error
*/
func (a *SubAccountsAPIService) GetSubAccountAPIKey(
	ctx context.Context,
	accountId uuid.UUID,
	subAccountId uuid.UUID,
	keyId uuid.UUID,
	opts ...RequestOption,
) (*responses.APIKey, *http.Response, error) {
	var result responses.APIKey

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/sub-accounts/{sub_account_id}/api-keys/{key_id}",
		PathParams: map[string]string{
			"account_id":     accountId.String(),
			"sub_account_id": subAccountId.String(),
			"key_id":         keyId.String(),
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
UpdateSubAccountAPIKey Update Sub-Account API Key

Updates the label and scopes for an API key owned by a sub account.

Validation Requirements:
- at least one editable field must be provided
- `label` must be non-empty when provided
- `scopes` must contain at least one valid scope when provided

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Parent account ID
	@param subAccountId Sub account ID
	@param keyId API key ID
	@param request UpdateAPIKeyRequest - The API key details to update
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return APIKey, *http.Response, error
*/
func (a *SubAccountsAPIService) UpdateSubAccountAPIKey(
	ctx context.Context,
	accountId uuid.UUID,
	subAccountId uuid.UUID,
	keyId uuid.UUID,
	request requests.UpdateAPIKeyRequest,
	opts ...RequestOption,
) (*responses.APIKey, *http.Response, error) {
	var result responses.APIKey

	config := RequestConfig{
		Method:       http.MethodPut,
		PathTemplate: "/v2/accounts/{account_id}/sub-accounts/{sub_account_id}/api-keys/{key_id}",
		PathParams: map[string]string{
			"account_id":     accountId.String(),
			"sub_account_id": subAccountId.String(),
			"key_id":         keyId.String(),
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
DeleteSubAccountAPIKey Delete Sub-Account API Key

Deletes an API key owned by a sub account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Parent account ID
	@param subAccountId Sub account ID
	@param keyId API key ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SuccessResponse, *http.Response, error
*/
func (a *SubAccountsAPIService) DeleteSubAccountAPIKey(
	ctx context.Context,
	accountId uuid.UUID,
	subAccountId uuid.UUID,
	keyId uuid.UUID,
	opts ...RequestOption,
) (*common.SuccessResponse, *http.Response, error) {
	var result common.SuccessResponse

	config := RequestConfig{
		Method:       http.MethodDelete,
		PathTemplate: "/v2/accounts/{account_id}/sub-accounts/{sub_account_id}/api-keys/{key_id}",
		PathParams: map[string]string{
			"account_id":     accountId.String(),
			"sub_account_id": subAccountId.String(),
			"key_id":         keyId.String(),
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
