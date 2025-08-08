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

// WebhooksAPIService WebhooksAPI service
type WebhooksAPIService service

type ApiCreateWebhookRequest struct {
	ctx                  context.Context
	ApiService           *WebhooksAPIService
	createWebhookRequest *CreateWebhookRequest
	accountId            uuid.UUID
	idempotencyKey       *string
}

func (r ApiCreateWebhookRequest) CreateWebhookRequest(createWebhookRequest CreateWebhookRequest) ApiCreateWebhookRequest {
	r.createWebhookRequest = &createWebhookRequest
	return r
}

// Optional idempotency key for safe request retries. Must be a unique string for each logical request. Requests with the same key will return the same response. Keys expire after 24 hours.
func (r ApiCreateWebhookRequest) IdempotencyKey(idempotencyKey string) ApiCreateWebhookRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiCreateWebhookRequest) Execute() (*Webhook, *http.Response, error) {
	return r.ApiService.CreateWebhookExecute(r)
}

/*
CreateWebhook Create Webhook

Creates a new webhook for event notifications

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@return ApiCreateWebhookRequest
*/
func (a *WebhooksAPIService) CreateWebhook(ctx context.Context, accountId uuid.UUID) ApiCreateWebhookRequest {
	return ApiCreateWebhookRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
	}
}

// Execute executes the request
//
//	@return Webhook
func (a *WebhooksAPIService) CreateWebhookExecute(r ApiCreateWebhookRequest) (*Webhook, *http.Response, error) {
	var (
		method = http.MethodPost
		body   interface{}

		returnValue *Webhook
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.CreateWebhook")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/webhooks"
	path = strings.ReplaceAll(path, "{account_id}", url.PathEscape(parameterValueToString(r.accountId, "accountId")))

	headers := make(map[string]string)
	params := url.Values{}
	formParams := url.Values{}
	if r.createWebhookRequest == nil {
		return returnValue, nil, reportError("createWebhookRequest is required and must be specified")
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
	body = r.createWebhookRequest
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

type ApiDeleteWebhookRequest struct {
	ctx        context.Context
	ApiService *WebhooksAPIService
	accountId  uuid.UUID
	webhookId  uuid.UUID
}

func (r ApiDeleteWebhookRequest) Execute() (*SuccessResponse, *http.Response, error) {
	return r.ApiService.DeleteWebhookExecute(r)
}

/*
DeleteWebhook Delete Webhook

Deletes a webhook

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param webhookId Webhook ID
	@return ApiDeleteWebhookRequest
*/
func (a *WebhooksAPIService) DeleteWebhook(ctx context.Context, accountId uuid.UUID, webhookId uuid.UUID) ApiDeleteWebhookRequest {
	return ApiDeleteWebhookRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
		webhookId:  webhookId,
	}
}

// Execute executes the request
//
//	@return SuccessResponse
func (a *WebhooksAPIService) DeleteWebhookExecute(r ApiDeleteWebhookRequest) (*SuccessResponse, *http.Response, error) {
	var (
		method = http.MethodDelete
		body   interface{}

		returnValue *SuccessResponse
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.DeleteWebhook")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/webhooks/{webhook_id}"
	path = strings.ReplaceAll(path, "{account_id}", url.PathEscape(parameterValueToString(r.accountId, "accountId")))
	path = strings.ReplaceAll(path, "{webhook_id}", url.PathEscape(parameterValueToString(r.webhookId, "webhookId")))

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

type ApiGetWebhookRequest struct {
	ctx        context.Context
	ApiService *WebhooksAPIService
	accountId  uuid.UUID
	webhookId  uuid.UUID
}

func (r ApiGetWebhookRequest) Execute() (*Webhook, *http.Response, error) {
	return r.ApiService.GetWebhookExecute(r)
}

/*
GetWebhook Get Webhook

Returns a specific webhook by ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param webhookId Webhook ID
	@return ApiGetWebhookRequest
*/
func (a *WebhooksAPIService) GetWebhook(ctx context.Context, accountId uuid.UUID, webhookId uuid.UUID) ApiGetWebhookRequest {
	return ApiGetWebhookRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
		webhookId:  webhookId,
	}
}

// Execute executes the request
//
//	@return Webhook
func (a *WebhooksAPIService) GetWebhookExecute(r ApiGetWebhookRequest) (*Webhook, *http.Response, error) {
	var (
		method = http.MethodGet
		body   interface{}

		returnValue *Webhook
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.GetWebhook")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/webhooks/{webhook_id}"
	path = strings.ReplaceAll(path, "{account_id}", url.PathEscape(parameterValueToString(r.accountId, "accountId")))
	path = strings.ReplaceAll(path, "{webhook_id}", url.PathEscape(parameterValueToString(r.webhookId, "webhookId")))

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

type ApiGetWebhooksRequest struct {
	ctx                  context.Context
	ApiService           *WebhooksAPIService
	accountId            uuid.UUID
	enabled              *bool
	onReception          *bool
	onDelivered          *bool
	onTransientError     *bool
	onFailed             *bool
	onBounced            *bool
	onSuppressed         *bool
	onOpened             *bool
	onClicked            *bool
	onSuppressionCreated *bool
	onDnsError           *bool
	limit                *int32
	cursor               *string
}

// Filter by enabled status
func (r ApiGetWebhooksRequest) Enabled(enabled bool) ApiGetWebhooksRequest {
	r.enabled = &enabled
	return r
}

// Filter by reception event trigger
func (r ApiGetWebhooksRequest) OnReception(onReception bool) ApiGetWebhooksRequest {
	r.onReception = &onReception
	return r
}

// Filter by delivery event trigger
func (r ApiGetWebhooksRequest) OnDelivered(onDelivered bool) ApiGetWebhooksRequest {
	r.onDelivered = &onDelivered
	return r
}

// Filter by transient error event trigger
func (r ApiGetWebhooksRequest) OnTransientError(onTransientError bool) ApiGetWebhooksRequest {
	r.onTransientError = &onTransientError
	return r
}

// Filter by failure event trigger
func (r ApiGetWebhooksRequest) OnFailed(onFailed bool) ApiGetWebhooksRequest {
	r.onFailed = &onFailed
	return r
}

// Filter by bounce event trigger
func (r ApiGetWebhooksRequest) OnBounced(onBounced bool) ApiGetWebhooksRequest {
	r.onBounced = &onBounced
	return r
}

// Filter by suppression event trigger
func (r ApiGetWebhooksRequest) OnSuppressed(onSuppressed bool) ApiGetWebhooksRequest {
	r.onSuppressed = &onSuppressed
	return r
}

// Filter by open event trigger
func (r ApiGetWebhooksRequest) OnOpened(onOpened bool) ApiGetWebhooksRequest {
	r.onOpened = &onOpened
	return r
}

// Filter by click event trigger
func (r ApiGetWebhooksRequest) OnClicked(onClicked bool) ApiGetWebhooksRequest {
	r.onClicked = &onClicked
	return r
}

// Filter by suppression creation event trigger
func (r ApiGetWebhooksRequest) OnSuppressionCreated(onSuppressionCreated bool) ApiGetWebhooksRequest {
	r.onSuppressionCreated = &onSuppressionCreated
	return r
}

// Filter by DNS error event trigger
func (r ApiGetWebhooksRequest) OnDnsError(onDnsError bool) ApiGetWebhooksRequest {
	r.onDnsError = &onDnsError
	return r
}

// Maximum number of items to return (1-100)
func (r ApiGetWebhooksRequest) Limit(limit int32) ApiGetWebhooksRequest {
	r.limit = &limit
	return r
}

// Pagination cursor for the next page
func (r ApiGetWebhooksRequest) Cursor(cursor string) ApiGetWebhooksRequest {
	r.cursor = &cursor
	return r
}

func (r ApiGetWebhooksRequest) Execute() (*PaginatedWebhooksResponse, *http.Response, error) {
	return r.ApiService.GetWebhooksExecute(r)
}

/*
GetWebhooks Get Webhooks

# Returns a list of webhooks for the account

**Query Parameters:**
- `enabled`: Filter by enabled status
- Event filters: `on_reception`, `on_delivered`, `on_transient_error`, `on_failed`, `on_bounced`, `on_suppressed`, `on_opened`, `on_clicked`, `on_suppression_created`, `on_dns_error`
- `limit`: Maximum number of items to return (1-100, default: 100)
- `cursor`: Pagination cursor for the next page

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@return ApiGetWebhooksRequest
*/
func (a *WebhooksAPIService) GetWebhooks(ctx context.Context, accountId uuid.UUID) ApiGetWebhooksRequest {
	return ApiGetWebhooksRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
	}
}

// Execute executes the request
//
//	@return PaginatedWebhooksResponse
func (a *WebhooksAPIService) GetWebhooksExecute(r ApiGetWebhooksRequest) (*PaginatedWebhooksResponse, *http.Response, error) {
	var (
		method = http.MethodGet
		body   interface{}

		returnValue *PaginatedWebhooksResponse
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.GetWebhooks")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/webhooks"
	path = strings.ReplaceAll(path, "{account_id}", url.PathEscape(parameterValueToString(r.accountId, "accountId")))

	headers := make(map[string]string)
	params := url.Values{}
	formParams := url.Values{}

	if r.enabled != nil {
		parameterAddToHeaderOrQuery(params, "enabled", r.enabled, "form", "")
	}
	if r.onReception != nil {
		parameterAddToHeaderOrQuery(params, "on_reception", r.onReception, "form", "")
	}
	if r.onDelivered != nil {
		parameterAddToHeaderOrQuery(params, "on_delivered", r.onDelivered, "form", "")
	}
	if r.onTransientError != nil {
		parameterAddToHeaderOrQuery(params, "on_transient_error", r.onTransientError, "form", "")
	}
	if r.onFailed != nil {
		parameterAddToHeaderOrQuery(params, "on_failed", r.onFailed, "form", "")
	}
	if r.onBounced != nil {
		parameterAddToHeaderOrQuery(params, "on_bounced", r.onBounced, "form", "")
	}
	if r.onSuppressed != nil {
		parameterAddToHeaderOrQuery(params, "on_suppressed", r.onSuppressed, "form", "")
	}
	if r.onOpened != nil {
		parameterAddToHeaderOrQuery(params, "on_opened", r.onOpened, "form", "")
	}
	if r.onClicked != nil {
		parameterAddToHeaderOrQuery(params, "on_clicked", r.onClicked, "form", "")
	}
	if r.onSuppressionCreated != nil {
		parameterAddToHeaderOrQuery(params, "on_suppression_created", r.onSuppressionCreated, "form", "")
	}
	if r.onDnsError != nil {
		parameterAddToHeaderOrQuery(params, "on_dns_error", r.onDnsError, "form", "")
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

type ApiUpdateWebhookRequest struct {
	ctx                  context.Context
	ApiService           *WebhooksAPIService
	updateWebhookRequest *UpdateWebhookRequest
	accountId            uuid.UUID
	webhookId            uuid.UUID
}

func (r ApiUpdateWebhookRequest) UpdateWebhookRequest(updateWebhookRequest UpdateWebhookRequest) ApiUpdateWebhookRequest {
	r.updateWebhookRequest = &updateWebhookRequest
	return r
}

func (r ApiUpdateWebhookRequest) Execute() (*Webhook, *http.Response, error) {
	return r.ApiService.UpdateWebhookExecute(r)
}

/*
UpdateWebhook Update Webhook

# Updates an existing webhook

**Note:** The webhook secret is not updatable

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param webhookId Webhook ID
	@return ApiUpdateWebhookRequest
*/
func (a *WebhooksAPIService) UpdateWebhook(ctx context.Context, accountId uuid.UUID, webhookId uuid.UUID) ApiUpdateWebhookRequest {
	return ApiUpdateWebhookRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
		webhookId:  webhookId,
	}
}

// Execute executes the request
//
//	@return Webhook
func (a *WebhooksAPIService) UpdateWebhookExecute(r ApiUpdateWebhookRequest) (*Webhook, *http.Response, error) {
	var (
		method = http.MethodPut
		body   interface{}

		returnValue *Webhook
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.UpdateWebhook")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/webhooks/{webhook_id}"
	path = strings.ReplaceAll(path, "{account_id}", url.PathEscape(parameterValueToString(r.accountId, "accountId")))
	path = strings.ReplaceAll(path, "{webhook_id}", url.PathEscape(parameterValueToString(r.webhookId, "webhookId")))

	headers := make(map[string]string)
	params := url.Values{}
	formParams := url.Values{}
	if r.updateWebhookRequest == nil {
		return returnValue, nil, reportError("updateWebhookRequest is required and must be specified")
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
	body = r.updateWebhookRequest
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
