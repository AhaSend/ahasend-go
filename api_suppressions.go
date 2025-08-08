package ahasend

import (
	"bytes"
	"context"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// SuppressionsAPIService SuppressionsAPI service
type SuppressionsAPIService service

type ApiCreateSuppressionRequest struct {
	ctx                      context.Context
	ApiService               *SuppressionsAPIService
	createSuppressionRequest *CreateSuppressionRequest
	accountId                uuid.UUID
	idempotencyKey           *string
}

func (r ApiCreateSuppressionRequest) CreateSuppressionRequest(createSuppressionRequest CreateSuppressionRequest) ApiCreateSuppressionRequest {
	r.createSuppressionRequest = &createSuppressionRequest
	return r
}

// Optional idempotency key for safe request retries. Must be a unique string for each logical request. Requests with the same key will return the same response. Keys expire after 24 hours.
func (r ApiCreateSuppressionRequest) IdempotencyKey(idempotencyKey string) ApiCreateSuppressionRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiCreateSuppressionRequest) Execute() (*Suppression, *http.Response, error) {
	return r.ApiService.CreateSuppressionExecute(r)
}

/*
CreateSuppression Create Suppression

# Creates a new suppression for an email address

**Validation Requirements:**
- `email` must be a valid email address
- `expires_at` must be in RFC3339 format
- `domain` is optional - if not provided, applies to all account domains

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@return ApiCreateSuppressionRequest
*/
func (a *SuppressionsAPIService) CreateSuppression(ctx context.Context, accountId uuid.UUID) ApiCreateSuppressionRequest {
	return ApiCreateSuppressionRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
	}
}

// Execute executes the request
//
//	@return Suppression
func (a *SuppressionsAPIService) CreateSuppressionExecute(r ApiCreateSuppressionRequest) (*Suppression, *http.Response, error) {
	var (
		method = http.MethodPost
		body   interface{}

		returnValue *Suppression
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SuppressionsAPIService.CreateSuppression")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/suppressions"
	path = strings.Replace(path, "{"+"account_id"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

	headers := make(map[string]string)
	params := url.Values{}
	formParams := url.Values{}
	if r.createSuppressionRequest == nil {
		return returnValue, nil, reportError("createSuppressionRequest is required and must be specified")
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
	body = r.createSuppressionRequest
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

type ApiDeleteAllSuppressionsRequest struct {
	ctx        context.Context
	ApiService *SuppressionsAPIService
	accountId  uuid.UUID
	domain     *string
}

// Optional domain filter
func (r ApiDeleteAllSuppressionsRequest) Domain(domain string) ApiDeleteAllSuppressionsRequest {
	r.domain = &domain
	return r
}

func (r ApiDeleteAllSuppressionsRequest) Execute() (*SuccessResponse, *http.Response, error) {
	return r.ApiService.DeleteAllSuppressionsExecute(r)
}

/*
DeleteAllSuppressions Delete All Suppressions

# Deletes all suppressions for the account

**Query Parameters:**
- `domain`: Optional domain filter to delete suppressions for specific domain only

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@return ApiDeleteAllSuppressionsRequest
*/
func (a *SuppressionsAPIService) DeleteAllSuppressions(ctx context.Context, accountId uuid.UUID) ApiDeleteAllSuppressionsRequest {
	return ApiDeleteAllSuppressionsRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
	}
}

// Execute executes the request
//
//	@return SuccessResponse
func (a *SuppressionsAPIService) DeleteAllSuppressionsExecute(r ApiDeleteAllSuppressionsRequest) (*SuccessResponse, *http.Response, error) {
	var (
		method = http.MethodDelete
		body   interface{}

		returnValue *SuccessResponse
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SuppressionsAPIService.DeleteAllSuppressions")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/suppressions"
	path = strings.Replace(path, "{"+"account_id"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

	headers := make(map[string]string)
	params := url.Values{}
	formParams := url.Values{}

	if r.domain != nil {
		parameterAddToHeaderOrQuery(params, "domain", r.domain, "form", "")
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

type ApiDeleteSuppressionRequest struct {
	ctx        context.Context
	ApiService *SuppressionsAPIService
	accountId  uuid.UUID
	email      string
	domain     *string
}

// Optional domain filter
func (r ApiDeleteSuppressionRequest) Domain(domain string) ApiDeleteSuppressionRequest {
	r.domain = &domain
	return r
}

func (r ApiDeleteSuppressionRequest) Execute() (*SuccessResponse, *http.Response, error) {
	return r.ApiService.DeleteSuppressionExecute(r)
}

/*
DeleteSuppression Delete Suppression

# Deletes a specific suppression by email address

**Query Parameters:**
- `domain`: Optional domain filter to delete suppression for specific domain only

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param email Email address
	@return ApiDeleteSuppressionRequest
*/
func (a *SuppressionsAPIService) DeleteSuppression(ctx context.Context, accountId uuid.UUID, email string) ApiDeleteSuppressionRequest {
	return ApiDeleteSuppressionRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
		email:      email,
	}
}

// Execute executes the request
//
//	@return SuccessResponse
func (a *SuppressionsAPIService) DeleteSuppressionExecute(r ApiDeleteSuppressionRequest) (*SuccessResponse, *http.Response, error) {
	var (
		method = http.MethodDelete
		body   interface{}

		returnValue *SuccessResponse
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SuppressionsAPIService.DeleteSuppression")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/suppressions/{email}"
	path = strings.Replace(path, "{"+"account_id"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	path = strings.Replace(path, "{"+"email"+"}", url.PathEscape(parameterValueToString(r.email, "email")), -1)

	headers := make(map[string]string)
	params := url.Values{}
	formParams := url.Values{}

	if r.domain != nil {
		parameterAddToHeaderOrQuery(params, "domain", r.domain, "form", "")
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

type ApiGetSuppressionsRequest struct {
	ctx        context.Context
	ApiService *SuppressionsAPIService
	accountId  uuid.UUID
	domain     *string
	fromDate   *time.Time
	toDate     *time.Time
	limit      *int32
	cursor     *string
}

// Filter by domain
func (r ApiGetSuppressionsRequest) Domain(domain string) ApiGetSuppressionsRequest {
	r.domain = &domain
	return r
}

// Filter suppressions created after this date (RFC3339 format)
func (r ApiGetSuppressionsRequest) FromDate(fromDate time.Time) ApiGetSuppressionsRequest {
	r.fromDate = &fromDate
	return r
}

// Filter suppressions created before this date (RFC3339 format)
func (r ApiGetSuppressionsRequest) ToDate(toDate time.Time) ApiGetSuppressionsRequest {
	r.toDate = &toDate
	return r
}

// Maximum number of items to return (1-100)
func (r ApiGetSuppressionsRequest) Limit(limit int32) ApiGetSuppressionsRequest {
	r.limit = &limit
	return r
}

// Pagination cursor for the next page
func (r ApiGetSuppressionsRequest) Cursor(cursor string) ApiGetSuppressionsRequest {
	r.cursor = &cursor
	return r
}

func (r ApiGetSuppressionsRequest) Execute() (*PaginatedSuppressionsResponse, *http.Response, error) {
	return r.ApiService.GetSuppressionsExecute(r)
}

/*
GetSuppressions Get Suppressions

# Returns a list of suppressions for the account

**Query Parameters:**
- `domain`: Filter by domain (optional)
- `from_date`: Filter suppressions created after this date (RFC3339 format)
- `to_date`: Filter suppressions created before this date (RFC3339 format)
- `limit`: Maximum number of items to return (1-100, default: 100)
- `cursor`: Pagination cursor for the next page

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@return ApiGetSuppressionsRequest
*/
func (a *SuppressionsAPIService) GetSuppressions(ctx context.Context, accountId uuid.UUID) ApiGetSuppressionsRequest {
	return ApiGetSuppressionsRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
	}
}

// Execute executes the request
//
//	@return PaginatedSuppressionsResponse
func (a *SuppressionsAPIService) GetSuppressionsExecute(r ApiGetSuppressionsRequest) (*PaginatedSuppressionsResponse, *http.Response, error) {
	var (
		method = http.MethodGet
		body   interface{}

		returnValue *PaginatedSuppressionsResponse
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SuppressionsAPIService.GetSuppressions")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/suppressions"
	path = strings.Replace(path, "{"+"account_id"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

	headers := make(map[string]string)
	params := url.Values{}
	formParams := url.Values{}

	if r.domain != nil {
		parameterAddToHeaderOrQuery(params, "domain", r.domain, "form", "")
	}
	if r.fromDate != nil {
		parameterAddToHeaderOrQuery(params, "from_date", r.fromDate, "form", "")
	}
	if r.toDate != nil {
		parameterAddToHeaderOrQuery(params, "to_date", r.toDate, "form", "")
	}
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
