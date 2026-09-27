package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/AhaSend/ahasend-go/models/responses"
	"github.com/google/uuid"
)

// ListsAPIService ListsAPI service
type ListsAPIService service

/*
GetLists List Lists

# Returns the account's lists in newest-first order

Query Parameters:
- `limit`: Maximum number of lists to return (1-100, default: 100)
- `after`: Cursor for the next page
- `before`: Cursor for the previous page
- `name`: List name filter

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param params GetListsParams - query parameters
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return PaginatedContactListsResponse, *http.Response, error
*/
func (a *ListsAPIService) GetLists(
	ctx context.Context,
	accountId uuid.UUID,
	params requests.GetListsParams,
	opts ...RequestOption,
) (*responses.PaginatedContactListsResponse, *http.Response, error) {
	var result responses.PaginatedContactListsResponse

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
	if params.Name != nil {
		queryParams.Set("name", *params.Name)
	}

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/lists",
		PathParams:   map[string]string{"account_id": accountId.String()},
		QueryParams:  queryParams,
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
CreateList Create List

# Creates one list with no members

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param request CreateListRequest - the list to create
	@param opts ...RequestOption - optional request options, including WithIdempotencyKey
	@return ContactList, *http.Response, error
*/
func (a *ListsAPIService) CreateList(
	ctx context.Context,
	accountId uuid.UUID,
	request requests.CreateListRequest,
	opts ...RequestOption,
) (*responses.ContactList, *http.Response, error) {
	var result responses.ContactList

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/lists",
		PathParams:   map[string]string{"account_id": accountId.String()},
		Body:         request,
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
GetList Get List

# Returns one account-scoped list, including its contact count

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param listId List ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return ContactList, *http.Response, error
*/
func (a *ListsAPIService) GetList(
	ctx context.Context,
	accountId uuid.UUID,
	listId uuid.UUID,
	opts ...RequestOption,
) (*responses.ContactList, *http.Response, error) {
	var result responses.ContactList

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/lists/{list_id}",
		PathParams: map[string]string{
			"account_id": accountId.String(),
			"list_id":    listId.String(),
		},
		Result: &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
UpdateList Update List

# Partially updates one account-scoped list

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param listId List ID
	@param request UpdateListRequest - the fields to update
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return ContactList, *http.Response, error
*/
func (a *ListsAPIService) UpdateList(
	ctx context.Context,
	accountId uuid.UUID,
	listId uuid.UUID,
	request requests.UpdateListRequest,
	opts ...RequestOption,
) (*responses.ContactList, *http.Response, error) {
	var result responses.ContactList

	config := RequestConfig{
		Method:       http.MethodPut,
		PathTemplate: "/v2/accounts/{account_id}/lists/{list_id}",
		PathParams: map[string]string{
			"account_id": accountId.String(),
			"list_id":    listId.String(),
		},
		Body:   request,
		Result: &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
DeleteList Delete List

# Soft-deletes one account-scoped list and keeps its memberships

Returns 409 while a campaign that still needs the list references it.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param listId List ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SuccessResponse, *http.Response, error
*/
func (a *ListsAPIService) DeleteList(
	ctx context.Context,
	accountId uuid.UUID,
	listId uuid.UUID,
	opts ...RequestOption,
) (*common.SuccessResponse, *http.Response, error) {
	var result common.SuccessResponse

	config := RequestConfig{
		Method:       http.MethodDelete,
		PathTemplate: "/v2/accounts/{account_id}/lists/{list_id}",
		PathParams: map[string]string{
			"account_id": accountId.String(),
			"list_id":    listId.String(),
		},
		Result: &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
GetListContacts List Contacts on a List

# Returns one list's memberships in newest-first order

Query Parameters:
- `limit`: Maximum number of memberships to return (1-100, default: 100)
- `after`: Cursor for the next page
- `before`: Cursor for the previous page
- `subscription_status`: Membership status filter
- `email`: Normalized exact email filter
- `include_contacts`: Embed each whole contact; requires the contacts:read scope

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param listId List ID
	@param params GetListContactsParams - query parameters
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return PaginatedListContactsResponse, *http.Response, error
*/
func (a *ListsAPIService) GetListContacts(
	ctx context.Context,
	accountId uuid.UUID,
	listId uuid.UUID,
	params requests.GetListContactsParams,
	opts ...RequestOption,
) (*responses.PaginatedListContactsResponse, *http.Response, error) {
	var result responses.PaginatedListContactsResponse

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
	if params.SubscriptionStatus != nil {
		queryParams.Set("subscription_status", *params.SubscriptionStatus)
	}
	if params.Email != nil {
		queryParams.Set("email", *params.Email)
	}
	if params.IncludeContacts != nil {
		queryParams.Set("include_contacts", strconv.FormatBool(*params.IncludeContacts))
	}

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/lists/{list_id}/contacts",
		PathParams: map[string]string{
			"account_id": accountId.String(),
			"list_id":    listId.String(),
		},
		QueryParams: queryParams,
		Result:      &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
BatchAddListContacts Batch Add Contacts to a List

# Synchronously adds up to 1,000 existing contacts to one list

Existing memberships are reported already_member and left unchanged. The
response is 200 even when entries fail; check Failed and each outcome.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param listId List ID
	@param request BatchAddListContactsRequest - ordered contacts to add
	@param opts ...RequestOption - optional request options, including WithIdempotencyKey
	@return BatchAddListContactsResponse, *http.Response, error
*/
func (a *ListsAPIService) BatchAddListContacts(
	ctx context.Context,
	accountId uuid.UUID,
	listId uuid.UUID,
	request requests.BatchAddListContactsRequest,
	opts ...RequestOption,
) (*responses.BatchAddListContactsResponse, *http.Response, error) {
	var result responses.BatchAddListContactsResponse

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/lists/{list_id}/contacts/batch",
		PathParams: map[string]string{
			"account_id": accountId.String(),
			"list_id":    listId.String(),
		},
		Body:   request,
		Result: &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
UpsertListContact Add or Update a List Contact

# Adds one existing contact to one list, or updates its membership

Returns 409 when the membership is complained, which cannot be changed.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param listId List ID
	@param idOrEmail Contact UUID or raw email address
	@param request UpsertListContactRequest - the membership status to write, if any
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return ListContact, *http.Response, error
*/
func (a *ListsAPIService) UpsertListContact(
	ctx context.Context,
	accountId uuid.UUID,
	listId uuid.UUID,
	idOrEmail string,
	request requests.UpsertListContactRequest,
	opts ...RequestOption,
) (*responses.ListContact, *http.Response, error) {
	var result responses.ListContact

	config := RequestConfig{
		Method:       http.MethodPut,
		PathTemplate: "/v2/accounts/{account_id}/lists/{list_id}/contacts/{id_or_email}",
		PathParams: map[string]string{
			"account_id":  accountId.String(),
			"list_id":     listId.String(),
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
DeleteListContact Remove a Contact from a List

# Deletes one membership, discarding any unsubscribe it records

To stop marketing mail on one list reversibly, use UpsertListContact with
the unsubscribed status instead. A complained membership cannot be removed
and answers 409.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param listId List ID
	@param idOrEmail Contact UUID or raw email address
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SuccessResponse, *http.Response, error
*/
func (a *ListsAPIService) DeleteListContact(
	ctx context.Context,
	accountId uuid.UUID,
	listId uuid.UUID,
	idOrEmail string,
	opts ...RequestOption,
) (*common.SuccessResponse, *http.Response, error) {
	var result common.SuccessResponse

	config := RequestConfig{
		Method:       http.MethodDelete,
		PathTemplate: "/v2/accounts/{account_id}/lists/{list_id}/contacts/{id_or_email}",
		PathParams: map[string]string{
			"account_id":  accountId.String(),
			"list_id":     listId.String(),
			"id_or_email": idOrEmail,
		},
		Result: &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
GetContactLists List a Contact's Lists

# Returns one contact's memberships, each with its list embedded

Query Parameters:
- `limit`: Maximum number of memberships to return (1-100, default: 100)
- `after`: Cursor for the next page
- `before`: Cursor for the previous page
- `subscription_status`: Membership status filter

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param idOrEmail Contact UUID or raw email address
	@param params GetContactListsParams - query parameters
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return PaginatedListContactsResponse, *http.Response, error
*/
func (a *ListsAPIService) GetContactLists(
	ctx context.Context,
	accountId uuid.UUID,
	idOrEmail string,
	params requests.GetContactListsParams,
	opts ...RequestOption,
) (*responses.PaginatedListContactsResponse, *http.Response, error) {
	var result responses.PaginatedListContactsResponse

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
	if params.SubscriptionStatus != nil {
		queryParams.Set("subscription_status", *params.SubscriptionStatus)
	}

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/contacts/{id_or_email}/lists",
		PathParams: map[string]string{
			"account_id":  accountId.String(),
			"id_or_email": idOrEmail,
		},
		QueryParams: queryParams,
		Result:      &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}
