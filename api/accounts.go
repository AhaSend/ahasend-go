package api

import (
	"context"
	"net/http"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/AhaSend/ahasend-go/models/responses"
	"github.com/google/uuid"
)

// AccountsAPIService AccountsAPI service
type AccountsAPIService service

/*
AddAccountMember Add Account Member

# Adds a new member to the account

Validation Requirements:
- `email` must be a valid email address
- `role` must be one of: Administrator, Developer, Analyst, Billing Manager

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param request AddMemberRequest - The member details to add
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return UserAccount, *http.Response, error
*/
func (a *AccountsAPIService) AddAccountMember(
	ctx context.Context,
	accountId uuid.UUID,
	request requests.AddMemberRequest,
	opts ...RequestOption,
) (*responses.UserAccount, *http.Response, error) {
	var result responses.UserAccount

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/members",
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
GetAccount Get Account

Retrieves detailed information about a specific account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return Account, *http.Response, error
*/
func (a *AccountsAPIService) GetAccount(
	ctx context.Context,
	accountId uuid.UUID,
	opts ...RequestOption,
) (*responses.Account, *http.Response, error) {
	var result responses.Account

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}",
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
GetAccountMembers Get Account Members

Retrieves a list of all members in the account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return AccountMembersResponse, *http.Response, error
*/
func (a *AccountsAPIService) GetAccountMembers(
	ctx context.Context,
	accountId uuid.UUID,
	opts ...RequestOption,
) (*responses.AccountMembersResponse, *http.Response, error) {
	var result responses.AccountMembersResponse

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/members",
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
RemoveAccountMember Remove Account Member

Removes a member from the account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param userId User ID to remove
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SuccessResponse, *http.Response, error
*/
func (a *AccountsAPIService) RemoveAccountMember(
	ctx context.Context,
	accountId uuid.UUID,
	userId uuid.UUID,
	opts ...RequestOption,
) (*common.SuccessResponse, *http.Response, error) {
	var result common.SuccessResponse

	config := RequestConfig{
		Method:       http.MethodDelete,
		PathTemplate: "/v2/accounts/{account_id}/members/{user_id}",
		PathParams: map[string]string{
			"account_id": accountId.String(),
			"user_id":    userId.String(),
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
UpdateAccount Update Account

Updates account information

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param request UpdateAccountRequest - The account details to update
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return Account, *http.Response, error
*/
func (a *AccountsAPIService) UpdateAccount(
	ctx context.Context,
	accountId uuid.UUID,
	request requests.UpdateAccountRequest,
	opts ...RequestOption,
) (*responses.Account, *http.Response, error) {
	var result responses.Account

	config := RequestConfig{
		Method:       http.MethodPut,
		PathTemplate: "/v2/accounts/{account_id}",
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
