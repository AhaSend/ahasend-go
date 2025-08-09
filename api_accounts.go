package ahasend

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
)

// AccountsAPIService AccountsAPI service
type AccountsAPIService service

type ApiAddAccountMemberRequest struct {
	ctx              context.Context
	ApiService       *AccountsAPIService
	addMemberRequest *AddMemberRequest
	accountId        uuid.UUID
	idempotencyKey   *string
}

func (r ApiAddAccountMemberRequest) AddMemberRequest(addMemberRequest AddMemberRequest) ApiAddAccountMemberRequest {
	r.addMemberRequest = &addMemberRequest
	return r
}

// Optional idempotency key for safe request retries. Must be a unique string for each logical request. Requests with the same key will return the same response. Keys expire after 24 hours.
func (r ApiAddAccountMemberRequest) IdempotencyKey(idempotencyKey string) ApiAddAccountMemberRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiAddAccountMemberRequest) Execute() (*UserAccount, *http.Response, error) {
	return r.ApiService.AddAccountMemberExecute(r)
}

/*
AddAccountMember Add Account Member

# Adds a new member to the account

**Validation Requirements:**
- `email` must be a valid email address
- `role` must be one of: Administrator, Developer, Analyst, Billing Manager

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@return ApiAddAccountMemberRequest
*/
func (a *AccountsAPIService) AddAccountMember(ctx context.Context, accountId uuid.UUID) ApiAddAccountMemberRequest {
	return ApiAddAccountMemberRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
	}
}

// Execute executes the request
//
//	@return UserAccount
func (a *AccountsAPIService) AddAccountMemberExecute(r ApiAddAccountMemberRequest) (*UserAccount, *http.Response, error) {
	var (
		method = http.MethodPost
		body   interface{}

		returnValue *UserAccount
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.AddAccountMember")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/members"
	path = strings.ReplaceAll(
		path,
		"{account_id}",
		url.PathEscape(parameterValueToString(r.accountId, "accountId")),
	)

	headers := make(map[string]string)
	params := url.Values{}
	formParams := url.Values{}
	if r.addMemberRequest == nil {
		return returnValue, nil, NewRequiredFieldError("addMemberRequest")
	}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headers["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headers["Accept"] = headerAccept
	}
	if r.idempotencyKey != nil {
		parameterAddToHeaderOrQuery(headers, "Idempotency-Key", r.idempotencyKey, "simple", "")
	}
	// body params
	body = r.addMemberRequest
	req, err := a.client.prepareRequest(r.ctx, path, method, body, headers, params, formParams)
	if err != nil {
		return returnValue, nil, err
	}

	response, err := a.client.callAPI(req)
	if err != nil || response == nil {
		return returnValue, response, err
	}

	responseBody, err := io.ReadAll(response.Body)
	response.Body.Close()
	response.Body = io.NopCloser(bytes.NewBuffer(responseBody))
	if err != nil {
		return returnValue, response, err
	}

	if response.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  responseBody,
			error: response.Status,
		}
		if response.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
			return returnValue, response, newErr
		}
		if response.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
			return returnValue, response, newErr
		}
		if response.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
			return returnValue, response, newErr
		}
		if response.StatusCode == 409 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
			return returnValue, response, newErr
		}
		if response.StatusCode == 412 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
			return returnValue, response, newErr
		}
		if response.StatusCode == 500 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
		}
		return returnValue, response, newErr
	}

	err = a.client.decode(&returnValue, responseBody, response.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  responseBody,
			error: err.Error(),
		}
		return returnValue, response, newErr
	}

	return returnValue, response, nil
}

type ApiGetAccountRequest struct {
	ctx        context.Context
	ApiService *AccountsAPIService
	accountId  uuid.UUID
}

func (r ApiGetAccountRequest) Execute() (*Account, *http.Response, error) {
	return r.ApiService.GetAccountExecute(r)
}

/*
GetAccount Get Account

Returns account information

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@return ApiGetAccountRequest
*/
func (a *AccountsAPIService) GetAccount(ctx context.Context, accountId uuid.UUID) ApiGetAccountRequest {
	return ApiGetAccountRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
	}
}

// Execute executes the request
//
//	@return Account
func (a *AccountsAPIService) GetAccountExecute(r ApiGetAccountRequest) (*Account, *http.Response, error) {
	var (
		method = http.MethodGet
		body   interface{}

		returnValue *Account
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.GetAccount")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}"
	path = strings.ReplaceAll(
		path,
		"{account_id}",
		url.PathEscape(parameterValueToString(r.accountId, "accountId")),
	)

	headers := make(map[string]string)
	params := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headers["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headers["Accept"] = headerAccept
	}
	req, err := a.client.prepareRequest(r.ctx, path, method, body, headers, params, formParams)
	if err != nil {
		return returnValue, nil, err
	}

	response, err := a.client.callAPI(req)
	if err != nil || response == nil {
		return returnValue, response, err
	}

	responseBody, err := io.ReadAll(response.Body)
	response.Body.Close()
	response.Body = io.NopCloser(bytes.NewBuffer(responseBody))
	if err != nil {
		return returnValue, response, err
	}

	if response.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  responseBody,
			error: response.Status,
		}
		if response.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
			return returnValue, response, newErr
		}
		if response.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
			return returnValue, response, newErr
		}
		if response.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
			return returnValue, response, newErr
		}
		if response.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
			return returnValue, response, newErr
		}
		if response.StatusCode == 500 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
		}
		return returnValue, response, newErr
	}

	err = a.client.decode(&returnValue, responseBody, response.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  responseBody,
			error: err.Error(),
		}
		return returnValue, response, newErr
	}

	return returnValue, response, nil
}

type ApiGetAccountMembersRequest struct {
	ctx        context.Context
	ApiService *AccountsAPIService
	accountId  uuid.UUID
}

func (r ApiGetAccountMembersRequest) Execute() (*AccountMembersResponse, *http.Response, error) {
	return r.ApiService.GetAccountMembersExecute(r)
}

/*
GetAccountMembers Get Account Members

Returns a list of account members

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@return ApiGetAccountMembersRequest
*/
func (a *AccountsAPIService) GetAccountMembers(ctx context.Context, accountId uuid.UUID) ApiGetAccountMembersRequest {
	return ApiGetAccountMembersRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
	}
}

// Execute executes the request
//
//	@return AccountMembersResponse
func (a *AccountsAPIService) GetAccountMembersExecute(r ApiGetAccountMembersRequest) (*AccountMembersResponse, *http.Response, error) {
	var (
		method = http.MethodGet
		body   interface{}

		returnValue *AccountMembersResponse
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.GetAccountMembers")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/members"
	path = strings.ReplaceAll(
		path,
		"{account_id}",
		url.PathEscape(parameterValueToString(r.accountId, "accountId")),
	)

	headers := make(map[string]string)
	params := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headers["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headers["Accept"] = headerAccept
	}
	req, err := a.client.prepareRequest(r.ctx, path, method, body, headers, params, formParams)
	if err != nil {
		return returnValue, nil, err
	}

	response, err := a.client.callAPI(req)
	if err != nil || response == nil {
		return returnValue, response, err
	}

	responseBody, err := io.ReadAll(response.Body)
	response.Body.Close()
	response.Body = io.NopCloser(bytes.NewBuffer(responseBody))
	if err != nil {
		return returnValue, response, err
	}

	if response.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  responseBody,
			error: response.Status,
		}
		if response.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
			return returnValue, response, newErr
		}
		if response.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
			return returnValue, response, newErr
		}
		if response.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
			return returnValue, response, newErr
		}
		if response.StatusCode == 500 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
		}
		return returnValue, response, newErr
	}

	err = a.client.decode(&returnValue, responseBody, response.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  responseBody,
			error: err.Error(),
		}
		return returnValue, response, newErr
	}

	return returnValue, response, nil
}

type ApiRemoveAccountMemberRequest struct {
	ctx        context.Context
	ApiService *AccountsAPIService
	accountId  uuid.UUID
	userId     uuid.UUID
}

func (r ApiRemoveAccountMemberRequest) Execute() (*SuccessResponse, *http.Response, error) {
	return r.ApiService.RemoveAccountMemberExecute(r)
}

/*
RemoveAccountMember Remove Account Member

# Removes a member from the account

**Restrictions:**
- You cannot delete yourself
- You cannot delete the account owner

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param userId User ID
	@return ApiRemoveAccountMemberRequest
*/
func (a *AccountsAPIService) RemoveAccountMember(ctx context.Context, accountId uuid.UUID, userId uuid.UUID) ApiRemoveAccountMemberRequest {
	return ApiRemoveAccountMemberRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
		userId:     userId,
	}
}

// Execute executes the request
//
//	@return SuccessResponse
func (a *AccountsAPIService) RemoveAccountMemberExecute(r ApiRemoveAccountMemberRequest) (*SuccessResponse, *http.Response, error) {
	var (
		method = http.MethodDelete
		body   interface{}

		returnValue *SuccessResponse
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.RemoveAccountMember")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/members/{user_id}"
	path = strings.ReplaceAll(
		path,
		"{account_id}",
		url.PathEscape(parameterValueToString(r.accountId, "accountId")),
	)
	path = strings.ReplaceAll(
		path,
		"{user_id}",
		url.PathEscape(parameterValueToString(r.userId, "userId")),
	)

	headers := make(map[string]string)
	params := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headers["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headers["Accept"] = headerAccept
	}
	req, err := a.client.prepareRequest(r.ctx, path, method, body, headers, params, formParams)
	if err != nil {
		return returnValue, nil, err
	}

	response, err := a.client.callAPI(req)
	if err != nil || response == nil {
		return returnValue, response, err
	}

	responseBody, err := io.ReadAll(response.Body)
	response.Body.Close()
	response.Body = io.NopCloser(bytes.NewBuffer(responseBody))
	if err != nil {
		return returnValue, response, err
	}

	if response.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  responseBody,
			error: response.Status,
		}
		if response.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
			return returnValue, response, newErr
		}
		if response.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
			return returnValue, response, newErr
		}
		if response.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
			return returnValue, response, newErr
		}
		if response.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
			return returnValue, response, newErr
		}
		if response.StatusCode == 500 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
		}
		return returnValue, response, newErr
	}

	err = a.client.decode(&returnValue, responseBody, response.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  responseBody,
			error: err.Error(),
		}
		return returnValue, response, newErr
	}

	return returnValue, response, nil
}

type ApiUpdateAccountRequest struct {
	ctx                  context.Context
	ApiService           *AccountsAPIService
	updateAccountRequest *UpdateAccountRequest
	accountId            uuid.UUID
}

func (r ApiUpdateAccountRequest) UpdateAccountRequest(updateAccountRequest UpdateAccountRequest) ApiUpdateAccountRequest {
	r.updateAccountRequest = &updateAccountRequest
	return r
}

func (r ApiUpdateAccountRequest) Execute() (*Account, *http.Response, error) {
	return r.ApiService.UpdateAccountExecute(r)
}

/*
UpdateAccount Update Account

# Updates account settings

**Validation Requirements:**
- `name` must be maximum 255 characters
- `website` must be a valid URL
- `message_metadata_retention` must be between 1 and 30 days
- `message_data_retention` must be between 0 and 30 days

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@return ApiUpdateAccountRequest
*/
func (a *AccountsAPIService) UpdateAccount(ctx context.Context, accountId uuid.UUID) ApiUpdateAccountRequest {
	return ApiUpdateAccountRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
	}
}

// Execute executes the request
//
//	@return Account
func (a *AccountsAPIService) UpdateAccountExecute(r ApiUpdateAccountRequest) (*Account, *http.Response, error) {
	var (
		method = http.MethodPut
		body   interface{}

		returnValue *Account
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.UpdateAccount")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}"
	path = strings.ReplaceAll(
		path,
		"{account_id}",
		url.PathEscape(parameterValueToString(r.accountId, "accountId")),
	)

	headers := make(map[string]string)
	params := url.Values{}
	formParams := url.Values{}
	if r.updateAccountRequest == nil {
		return returnValue, nil, NewRequiredFieldError("updateAccountRequest")
	}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headers["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headers["Accept"] = headerAccept
	}
	// body params
	body = r.updateAccountRequest
	req, err := a.client.prepareRequest(r.ctx, path, method, body, headers, params, formParams)
	if err != nil {
		return returnValue, nil, err
	}

	response, err := a.client.callAPI(req)
	if err != nil || response == nil {
		return returnValue, response, err
	}

	responseBody, err := io.ReadAll(response.Body)
	response.Body.Close()
	response.Body = io.NopCloser(bytes.NewBuffer(responseBody))
	if err != nil {
		return returnValue, response, err
	}

	if response.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  responseBody,
			error: response.Status,
		}
		if response.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
			return returnValue, response, newErr
		}
		if response.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
			return returnValue, response, newErr
		}
		if response.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
			return returnValue, response, newErr
		}
		if response.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
			return returnValue, response, newErr
		}
		if response.StatusCode == 500 {
			var v ErrorResponse
			err = a.client.decode(&v, responseBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
		}
		return returnValue, response, newErr
	}

	err = a.client.decode(&returnValue, responseBody, response.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  responseBody,
			error: err.Error(),
		}
		return returnValue, response, newErr
	}

	return returnValue, response, nil
}
