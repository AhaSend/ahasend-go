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

// APIKeysAPIService APIKeysAPI service
type APIKeysAPIService service

type ApiCreateAPIKeyRequest struct {
	ctx                 context.Context
	ApiService          *APIKeysAPIService
	createAPIKeyRequest *CreateAPIKeyRequest
	accountId           uuid.UUID
	idempotencyKey      *string
}

func (r ApiCreateAPIKeyRequest) CreateAPIKeyRequest(createAPIKeyRequest CreateAPIKeyRequest) ApiCreateAPIKeyRequest {
	r.createAPIKeyRequest = &createAPIKeyRequest
	return r
}

// Optional idempotency key for safe request retries. Must be a unique string for each logical request. Requests with the same key will return the same response. Keys expire after 24 hours.
func (r ApiCreateAPIKeyRequest) IdempotencyKey(idempotencyKey string) ApiCreateAPIKeyRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiCreateAPIKeyRequest) Execute() (*ModelAPIKey, *http.Response, error) {
	return r.ApiService.CreateAPIKeyExecute(r)
}

/*
CreateAPIKey Create API Key

Creates a new API key with the specified scopes

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@return ApiCreateAPIKeyRequest
*/
func (a *APIKeysAPIService) CreateAPIKey(ctx context.Context, accountId uuid.UUID) ApiCreateAPIKeyRequest {
	return ApiCreateAPIKeyRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
	}
}

// Execute executes the request
//
//	@return ModelAPIKey
func (a *APIKeysAPIService) CreateAPIKeyExecute(r ApiCreateAPIKeyRequest) (*ModelAPIKey, *http.Response, error) {
	var (
		method = http.MethodPost
		body   interface{}

		returnValue *ModelAPIKey
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APIKeysAPIService.CreateAPIKey")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/api-keys"
	path = strings.ReplaceAll(
		path,
		"{account_id}",
		url.PathEscape(parameterValueToString(r.accountId, "accountId")),
	)

	headers := make(map[string]string)
	params := url.Values{}
	formParams := url.Values{}
	if r.createAPIKeyRequest == nil {
		return returnValue, nil, NewRequiredFieldError("createAPIKeyRequest")
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
	body = r.createAPIKeyRequest
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
			err = a.client.decode(&v, responseBody)
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
			err = a.client.decode(&v, responseBody)
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
			err = a.client.decode(&v, responseBody)
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
			err = a.client.decode(&v, responseBody)
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
			err = a.client.decode(&v, responseBody)
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
			err = a.client.decode(&v, responseBody)
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
		}
		return returnValue, response, newErr
	}

	err = a.client.decode(&returnValue, responseBody)
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  responseBody,
			error: err.Error(),
		}
		return returnValue, response, newErr
	}

	return returnValue, response, nil
}

type ApiDeleteAPIKeyRequest struct {
	ctx        context.Context
	ApiService *APIKeysAPIService
	accountId  uuid.UUID
	keyId      uuid.UUID
}

func (r ApiDeleteAPIKeyRequest) Execute() (*SuccessResponse, *http.Response, error) {
	return r.ApiService.DeleteAPIKeyExecute(r)
}

/*
DeleteAPIKey Delete API Key

Deletes an API key

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param keyId API Key ID
	@return ApiDeleteAPIKeyRequest
*/
func (a *APIKeysAPIService) DeleteAPIKey(ctx context.Context, accountId uuid.UUID, keyId uuid.UUID) ApiDeleteAPIKeyRequest {
	return ApiDeleteAPIKeyRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
		keyId:      keyId,
	}
}

// Execute executes the request
//
//	@return SuccessResponse
func (a *APIKeysAPIService) DeleteAPIKeyExecute(r ApiDeleteAPIKeyRequest) (*SuccessResponse, *http.Response, error) {
	var (
		method = http.MethodDelete
		body   interface{}

		returnValue *SuccessResponse
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APIKeysAPIService.DeleteAPIKey")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/api-keys/{key_id}"
	path = strings.ReplaceAll(path, "{account_id}", url.PathEscape(parameterValueToString(r.accountId, "accountId")))
	path = strings.ReplaceAll(path, "{key_id}", url.PathEscape(parameterValueToString(r.keyId, "keyId")))

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
			err = a.client.decode(&v, responseBody)
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
			err = a.client.decode(&v, responseBody)
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
			err = a.client.decode(&v, responseBody)
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
			err = a.client.decode(&v, responseBody)
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
			err = a.client.decode(&v, responseBody)
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
		}
		return returnValue, response, newErr
	}

	err = a.client.decode(&returnValue, responseBody)
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  responseBody,
			error: err.Error(),
		}
		return returnValue, response, newErr
	}

	return returnValue, response, nil
}

type ApiGetAPIKeyRequest struct {
	ctx        context.Context
	ApiService *APIKeysAPIService
	accountId  uuid.UUID
	keyId      uuid.UUID
}

func (r ApiGetAPIKeyRequest) Execute() (*ModelAPIKey, *http.Response, error) {
	return r.ApiService.GetAPIKeyExecute(r)
}

/*
GetAPIKey Get API Key

Returns a specific API key by ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param keyId API Key ID
	@return ApiGetAPIKeyRequest
*/
func (a *APIKeysAPIService) GetAPIKey(ctx context.Context, accountId uuid.UUID, keyId uuid.UUID) ApiGetAPIKeyRequest {
	return ApiGetAPIKeyRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
		keyId:      keyId,
	}
}

// Execute executes the request
//
//	@return ModelAPIKey
func (a *APIKeysAPIService) GetAPIKeyExecute(r ApiGetAPIKeyRequest) (*ModelAPIKey, *http.Response, error) {
	var (
		method = http.MethodGet
		body   interface{}

		returnValue *ModelAPIKey
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APIKeysAPIService.GetAPIKey")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/api-keys/{key_id}"
	path = strings.ReplaceAll(path, "{account_id}", url.PathEscape(parameterValueToString(r.accountId, "accountId")))
	path = strings.ReplaceAll(path, "{key_id}", url.PathEscape(parameterValueToString(r.keyId, "keyId")))

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
			err = a.client.decode(&v, responseBody)
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
			err = a.client.decode(&v, responseBody)
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
			err = a.client.decode(&v, responseBody)
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
			err = a.client.decode(&v, responseBody)
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
			err = a.client.decode(&v, responseBody)
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
		}
		return returnValue, response, newErr
	}

	err = a.client.decode(&returnValue, responseBody)
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  responseBody,
			error: err.Error(),
		}
		return returnValue, response, newErr
	}

	return returnValue, response, nil
}

type ApiGetAPIKeysRequest struct {
	ctx        context.Context
	ApiService *APIKeysAPIService
	accountId  uuid.UUID
	limit      *int32
	cursor     *string
}

// Maximum number of items to return
func (r ApiGetAPIKeysRequest) Limit(limit int32) ApiGetAPIKeysRequest {
	r.limit = &limit
	return r
}

// Pagination cursor for the next page
func (r ApiGetAPIKeysRequest) Cursor(cursor string) ApiGetAPIKeysRequest {
	r.cursor = &cursor
	return r
}

func (r ApiGetAPIKeysRequest) Execute() (*PaginatedAPIKeysResponse, *http.Response, error) {
	return r.ApiService.GetAPIKeysExecute(r)
}

/*
GetAPIKeys Get API Keys

Returns a list of API keys for the account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@return ApiGetAPIKeysRequest
*/
func (a *APIKeysAPIService) GetAPIKeys(ctx context.Context, accountId uuid.UUID) ApiGetAPIKeysRequest {
	return ApiGetAPIKeysRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
	}
}

// Execute executes the request
//
//	@return PaginatedAPIKeysResponse
func (a *APIKeysAPIService) GetAPIKeysExecute(r ApiGetAPIKeysRequest) (*PaginatedAPIKeysResponse, *http.Response, error) {
	var (
		method = http.MethodGet
		body   interface{}

		returnValue *PaginatedAPIKeysResponse
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APIKeysAPIService.GetAPIKeys")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/api-keys"
	path = strings.ReplaceAll(path, "{account_id}", url.PathEscape(parameterValueToString(r.accountId, "accountId")))

	headers := make(map[string]string)
	params := url.Values{}
	formParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(params, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 100
		r.limit = &defaultValue
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(params, "cursor", r.cursor, "form", "")
	}
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
			err = a.client.decode(&v, responseBody)
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
			err = a.client.decode(&v, responseBody)
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
			err = a.client.decode(&v, responseBody)
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
			err = a.client.decode(&v, responseBody)
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
		}
		return returnValue, response, newErr
	}

	err = a.client.decode(&returnValue, responseBody)
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  responseBody,
			error: err.Error(),
		}
		return returnValue, response, newErr
	}

	return returnValue, response, nil
}

type ApiUpdateAPIKeyRequest struct {
	ctx                 context.Context
	ApiService          *APIKeysAPIService
	updateAPIKeyRequest *UpdateAPIKeyRequest
	accountId           uuid.UUID
	keyId               uuid.UUID
}

func (r ApiUpdateAPIKeyRequest) UpdateAPIKeyRequest(updateAPIKeyRequest UpdateAPIKeyRequest) ApiUpdateAPIKeyRequest {
	r.updateAPIKeyRequest = &updateAPIKeyRequest
	return r
}

func (r ApiUpdateAPIKeyRequest) Execute() (*ModelAPIKey, *http.Response, error) {
	return r.ApiService.UpdateAPIKeyExecute(r)
}

/*
UpdateAPIKey Update API Key

Updates an existing API key's label and scopes

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param keyId API Key ID
	@return ApiUpdateAPIKeyRequest
*/
func (a *APIKeysAPIService) UpdateAPIKey(ctx context.Context, accountId uuid.UUID, keyId uuid.UUID) ApiUpdateAPIKeyRequest {
	return ApiUpdateAPIKeyRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
		keyId:      keyId,
	}
}

// Execute executes the request
//
//	@return ModelAPIKey
func (a *APIKeysAPIService) UpdateAPIKeyExecute(r ApiUpdateAPIKeyRequest) (*ModelAPIKey, *http.Response, error) {
	var (
		method = http.MethodPut
		body   interface{}

		returnValue *ModelAPIKey
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APIKeysAPIService.UpdateAPIKey")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/api-keys/{key_id}"
	path = strings.ReplaceAll(path, "{account_id}", url.PathEscape(parameterValueToString(r.accountId, "accountId")))
	path = strings.ReplaceAll(path, "{key_id}", url.PathEscape(parameterValueToString(r.keyId, "keyId")))

	headers := make(map[string]string)
	params := url.Values{}
	formParams := url.Values{}
	if r.updateAPIKeyRequest == nil {
		return returnValue, nil, NewRequiredFieldError("updateAPIKeyRequest")
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
	body = r.updateAPIKeyRequest
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
			err = a.client.decode(&v, responseBody)
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
			err = a.client.decode(&v, responseBody)
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
			err = a.client.decode(&v, responseBody)
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
			err = a.client.decode(&v, responseBody)
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
			err = a.client.decode(&v, responseBody)
			if err != nil {
				newErr.error = err.Error()
				return returnValue, response, newErr
			}
			newErr.error = formatErrorMessage(response.Status, &v)
			newErr.model = v
		}
		return returnValue, response, newErr
	}

	err = a.client.decode(&returnValue, responseBody)
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  responseBody,
			error: err.Error(),
		}
		return returnValue, response, newErr
	}

	return returnValue, response, nil
}
