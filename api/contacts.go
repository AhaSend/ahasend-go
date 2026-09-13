package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/AhaSend/ahasend-go/models/responses"
	"github.com/google/uuid"
)

// ContactsAPIService ContactsAPI service
type ContactsAPIService service

/*
GetContacts List Contacts

# Returns account-global contacts in newest-first order

Query Parameters:
- `limit`: Maximum number of contacts to return (1-100, default: 100)
- `after`: Cursor for the next page
- `before`: Cursor for the previous page
- `email`: Normalized exact email filter
- `status`: Contact status filter
- `subscribed`: Global subscription-state filter
- `from_time`: Include contacts created at or after this RFC3339 time
- `to_time`: Include contacts created at or before this RFC3339 time

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param params GetContactsParams - query parameters
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return PaginatedContactsResponse, *http.Response, error
*/
func (a *ContactsAPIService) GetContacts(
	ctx context.Context,
	accountId uuid.UUID,
	params requests.GetContactsParams,
	opts ...RequestOption,
) (*responses.PaginatedContactsResponse, *http.Response, error) {
	var result responses.PaginatedContactsResponse

	queryParams := url.Values{}
	if params.Limit != nil {
		queryParams.Set("limit", fmt.Sprintf("%d", *params.Limit))
	} else {
		queryParams.Set("limit", "100")
	}
	if params.After != nil {
		queryParams.Set("after", *params.After)
	}
	if params.Before != nil {
		queryParams.Set("before", *params.Before)
	}
	if params.Email != nil {
		queryParams.Set("email", *params.Email)
	}
	if params.Status != nil {
		queryParams.Set("status", *params.Status)
	}
	if params.Subscribed != nil {
		queryParams.Set("subscribed", strconv.FormatBool(*params.Subscribed))
	}
	if params.FromTime != nil {
		queryParams.Set("from_time", params.FromTime.Format(time.RFC3339))
	}
	if params.ToTime != nil {
		queryParams.Set("to_time", params.ToTime.Format(time.RFC3339))
	}

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/contacts",
		PathParams:   map[string]string{"account_id": accountId.String()},
		QueryParams:  queryParams,
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
GetContact Get Contact

# Returns one account-scoped contact by UUID or email

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param idOrEmail Contact UUID or raw email address
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return Contact, *http.Response, error
*/
func (a *ContactsAPIService) GetContact(
	ctx context.Context,
	accountId uuid.UUID,
	idOrEmail string,
	opts ...RequestOption,
) (*responses.Contact, *http.Response, error) {
	var result responses.Contact

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/contacts/{id_or_email}",
		PathParams: map[string]string{
			"account_id":  accountId.String(),
			"id_or_email": idOrEmail,
		},
		Result: &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
CreateContact Create Contact

# Creates one account-global contact with no list membership

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param request CreateContactRequest - the contact to create
	@param opts ...RequestOption - optional request options, including WithIdempotencyKey
	@return Contact, *http.Response, error
*/
func (a *ContactsAPIService) CreateContact(
	ctx context.Context,
	accountId uuid.UUID,
	request requests.CreateContactRequest,
	opts ...RequestOption,
) (*responses.Contact, *http.Response, error) {
	var result responses.Contact

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/contacts",
		PathParams:   map[string]string{"account_id": accountId.String()},
		Body:         request,
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
UpdateContact Update Contact

# Partially updates one account-scoped contact by UUID or email

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param idOrEmail Contact UUID or raw email address
	@param request UpdateContactRequest - the fields to update
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return Contact, *http.Response, error
*/
func (a *ContactsAPIService) UpdateContact(
	ctx context.Context,
	accountId uuid.UUID,
	idOrEmail string,
	request requests.UpdateContactRequest,
	opts ...RequestOption,
) (*responses.Contact, *http.Response, error) {
	var result responses.Contact

	config := RequestConfig{
		Method:       http.MethodPut,
		PathTemplate: "/v2/accounts/{account_id}/contacts/{id_or_email}",
		PathParams: map[string]string{
			"account_id":  accountId.String(),
			"id_or_email": idOrEmail,
		},
		Body:   request,
		Result: &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
DeleteContact Delete Contact

# Permanently deletes one account-scoped contact by UUID or email

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param idOrEmail Contact UUID or raw email address
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SuccessResponse, *http.Response, error
*/
func (a *ContactsAPIService) DeleteContact(
	ctx context.Context,
	accountId uuid.UUID,
	idOrEmail string,
	opts ...RequestOption,
) (*common.SuccessResponse, *http.Response, error) {
	var result common.SuccessResponse

	config := RequestConfig{
		Method:       http.MethodDelete,
		PathTemplate: "/v2/accounts/{account_id}/contacts/{id_or_email}",
		PathParams: map[string]string{
			"account_id":  accountId.String(),
			"id_or_email": idOrEmail,
		},
		Result: &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
BatchUpsertContacts Batch Upsert Contacts

# Synchronously creates or updates up to 1,000 account-global contacts

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param request BatchUpsertContactsRequest - ordered contacts to upsert
	@param opts ...RequestOption - optional request options, including WithIdempotencyKey
	@return BatchUpsertContactsResponse, *http.Response, error
*/
func (a *ContactsAPIService) BatchUpsertContacts(
	ctx context.Context,
	accountId uuid.UUID,
	request requests.BatchUpsertContactsRequest,
	opts ...RequestOption,
) (*responses.BatchUpsertContactsResponse, *http.Response, error) {
	var result responses.BatchUpsertContactsResponse

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/contacts/batch",
		PathParams:   map[string]string{"account_id": accountId.String()},
		Body:         request,
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}
