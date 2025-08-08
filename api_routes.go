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

// RoutesAPIService RoutesAPI service
type RoutesAPIService service

type ApiCreateRouteRequest struct {
	ctx                context.Context
	ApiService         *RoutesAPIService
	createRouteRequest *CreateRouteRequest
	accountId          uuid.UUID
	idempotencyKey     *string
}

func (r ApiCreateRouteRequest) CreateRouteRequest(createRouteRequest CreateRouteRequest) ApiCreateRouteRequest {
	r.createRouteRequest = &createRouteRequest
	return r
}

// Optional idempotency key for safe request retries. Must be a unique string for each logical request. Requests with the same key will return the same response. Keys expire after 24 hours.
func (r ApiCreateRouteRequest) IdempotencyKey(idempotencyKey string) ApiCreateRouteRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiCreateRouteRequest) Execute() (*Route, *http.Response, error) {
	return r.ApiService.CreateRouteExecute(r)
}

/*
CreateRoute Create Route

Creates a new route for inbound email routing

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@return ApiCreateRouteRequest
*/
func (a *RoutesAPIService) CreateRoute(ctx context.Context, accountId uuid.UUID) ApiCreateRouteRequest {
	return ApiCreateRouteRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
	}
}

// Execute executes the request
//
//	@return Route
func (a *RoutesAPIService) CreateRouteExecute(r ApiCreateRouteRequest) (*Route, *http.Response, error) {
	var (
		method = http.MethodPost
		body   interface{}

		returnValue *Route
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoutesAPIService.CreateRoute")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/routes"
	path = strings.ReplaceAll(path, "{account_id}", url.PathEscape(parameterValueToString(r.accountId, "accountId")))

	headers := make(map[string]string)
	params := url.Values{}
	formParams := url.Values{}
	if r.createRouteRequest == nil {
		return returnValue, nil, reportError("createRouteRequest is required and must be specified")
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
	body = r.createRouteRequest
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

type ApiDeleteRouteRequest struct {
	ctx        context.Context
	ApiService *RoutesAPIService
	accountId  uuid.UUID
	routeId    uuid.UUID
}

func (r ApiDeleteRouteRequest) Execute() (*SuccessResponse, *http.Response, error) {
	return r.ApiService.DeleteRouteExecute(r)
}

/*
DeleteRoute Delete Route

Deletes a route

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param routeId Route ID
	@return ApiDeleteRouteRequest
*/
func (a *RoutesAPIService) DeleteRoute(ctx context.Context, accountId uuid.UUID, routeId uuid.UUID) ApiDeleteRouteRequest {
	return ApiDeleteRouteRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
		routeId:    routeId,
	}
}

// Execute executes the request
//
//	@return SuccessResponse
func (a *RoutesAPIService) DeleteRouteExecute(r ApiDeleteRouteRequest) (*SuccessResponse, *http.Response, error) {
	var (
		method = http.MethodDelete
		body   interface{}

		returnValue *SuccessResponse
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoutesAPIService.DeleteRoute")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/routes/{route_id}"
	path = strings.ReplaceAll(path, "{account_id}", url.PathEscape(parameterValueToString(r.accountId, "accountId")))
	path = strings.ReplaceAll(path, "{route_id}", url.PathEscape(parameterValueToString(r.routeId, "routeId")))

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

type ApiGetRouteRequest struct {
	ctx        context.Context
	ApiService *RoutesAPIService
	accountId  uuid.UUID
	routeId    uuid.UUID
}

func (r ApiGetRouteRequest) Execute() (*Route, *http.Response, error) {
	return r.ApiService.GetRouteExecute(r)
}

/*
GetRoute Get Route

Returns a specific route by ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param routeId Route ID
	@return ApiGetRouteRequest
*/
func (a *RoutesAPIService) GetRoute(ctx context.Context, accountId uuid.UUID, routeId uuid.UUID) ApiGetRouteRequest {
	return ApiGetRouteRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
		routeId:    routeId,
	}
}

// Execute executes the request
//
//	@return Route
func (a *RoutesAPIService) GetRouteExecute(r ApiGetRouteRequest) (*Route, *http.Response, error) {
	var (
		method = http.MethodGet
		body   interface{}

		returnValue *Route
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoutesAPIService.GetRoute")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/routes/{route_id}"
	path = strings.ReplaceAll(path, "{account_id}", url.PathEscape(parameterValueToString(r.accountId, "accountId")))
	path = strings.ReplaceAll(path, "{route_id}", url.PathEscape(parameterValueToString(r.routeId, "routeId")))

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

type ApiGetRoutesRequest struct {
	ctx        context.Context
	ApiService *RoutesAPIService
	accountId  uuid.UUID
	limit      *int32
	cursor     *string
}

// Maximum number of items to return (1-100)
func (r ApiGetRoutesRequest) Limit(limit int32) ApiGetRoutesRequest {
	r.limit = &limit
	return r
}

// Pagination cursor for the next page
func (r ApiGetRoutesRequest) Cursor(cursor string) ApiGetRoutesRequest {
	r.cursor = &cursor
	return r
}

func (r ApiGetRoutesRequest) Execute() (*PaginatedRoutesResponse, *http.Response, error) {
	return r.ApiService.GetRoutesExecute(r)
}

/*
GetRoutes Get Routes

Returns a list of routes for the account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@return ApiGetRoutesRequest
*/
func (a *RoutesAPIService) GetRoutes(ctx context.Context, accountId uuid.UUID) ApiGetRoutesRequest {
	return ApiGetRoutesRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
	}
}

// Execute executes the request
//
//	@return PaginatedRoutesResponse
func (a *RoutesAPIService) GetRoutesExecute(r ApiGetRoutesRequest) (*PaginatedRoutesResponse, *http.Response, error) {
	var (
		method = http.MethodGet
		body   interface{}

		returnValue *PaginatedRoutesResponse
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoutesAPIService.GetRoutes")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/routes"
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

type ApiUpdateRouteRequest struct {
	ctx                context.Context
	ApiService         *RoutesAPIService
	updateRouteRequest *UpdateRouteRequest
	accountId          uuid.UUID
	routeId            uuid.UUID
}

func (r ApiUpdateRouteRequest) UpdateRouteRequest(updateRouteRequest UpdateRouteRequest) ApiUpdateRouteRequest {
	r.updateRouteRequest = &updateRouteRequest
	return r
}

func (r ApiUpdateRouteRequest) Execute() (*Route, *http.Response, error) {
	return r.ApiService.UpdateRouteExecute(r)
}

/*
UpdateRoute Update Route

Updates an existing route

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param routeId Route ID
	@return ApiUpdateRouteRequest
*/
func (a *RoutesAPIService) UpdateRoute(ctx context.Context, accountId uuid.UUID, routeId uuid.UUID) ApiUpdateRouteRequest {
	return ApiUpdateRouteRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
		routeId:    routeId,
	}
}

// Execute executes the request
//
//	@return Route
func (a *RoutesAPIService) UpdateRouteExecute(r ApiUpdateRouteRequest) (*Route, *http.Response, error) {
	var (
		method = http.MethodPut
		body   interface{}

		returnValue *Route
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoutesAPIService.UpdateRoute")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/routes/{route_id}"
	path = strings.ReplaceAll(path, "{account_id}", url.PathEscape(parameterValueToString(r.accountId, "accountId")))
	path = strings.ReplaceAll(path, "{route_id}", url.PathEscape(parameterValueToString(r.routeId, "routeId")))

	headers := make(map[string]string)
	params := url.Values{}
	formParams := url.Values{}
	if r.updateRouteRequest == nil {
		return returnValue, nil, reportError("updateRouteRequest is required and must be specified")
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
	body = r.updateRouteRequest
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
