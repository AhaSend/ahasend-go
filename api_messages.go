package ahasend

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
)

// MessagesAPIService MessagesAPI service
type MessagesAPIService service

type ApiCancelMessageRequest struct {
	ctx        context.Context
	ApiService *MessagesAPIService
	accountId  uuid.UUID
	messageId  string
}

func (r ApiCancelMessageRequest) Execute() (*SuccessResponse, *http.Response, error) {
	return r.ApiService.CancelMessageExecute(r)
}

/*
CancelMessage Cancel Message

Cancels a scheduled message

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param messageId Message ID
	@return ApiCancelMessageRequest
*/
func (a *MessagesAPIService) CancelMessage(ctx context.Context, accountId uuid.UUID, messageId string) ApiCancelMessageRequest {
	return ApiCancelMessageRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
		messageId:  messageId,
	}
}

// Execute executes the request
//
//	@return SuccessResponse
func (a *MessagesAPIService) CancelMessageExecute(r ApiCancelMessageRequest) (*SuccessResponse, *http.Response, error) {
	var (
		method      = http.MethodDelete
		body        interface{}
		returnValue *SuccessResponse
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagesAPIService.CancelMessage")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/messages/{message_id}/cancel"
	path = strings.ReplaceAll(path, "{account_id}", url.PathEscape(parameterValueToString(r.accountId, "accountId")))
	path = strings.ReplaceAll(path, "{message_id}", url.PathEscape(parameterValueToString(r.messageId, "messageId")))

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

type ApiCreateMessageRequest struct {
	ctx                  context.Context
	ApiService           *MessagesAPIService
	createMessageRequest *CreateMessageRequest
	accountId            uuid.UUID
	idempotencyKey       *string
}

func (r ApiCreateMessageRequest) CreateMessageRequest(createMessageRequest CreateMessageRequest) ApiCreateMessageRequest {
	r.createMessageRequest = &createMessageRequest
	return r
}

// Optional idempotency key for safe request retries. Must be a unique string for each logical request. Requests with the same key will return the same response. Keys expire after 24 hours.
func (r ApiCreateMessageRequest) IdempotencyKey(idempotencyKey string) ApiCreateMessageRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiCreateMessageRequest) Execute() (*CreateMessageResponse, *http.Response, error) {
	return r.ApiService.CreateMessageExecute(r)
}

/*
CreateMessage Create Message

Creates and sends a message to one or more recipients.

**Validation Requirements:**
- Either `text_content` or `html_content` is required
- `from.email` must be from a domain you own with valid DNS records
- `retention.metadata` must be between 1 and 30 days
- `retention.data` must be between 0 and 30 days
- If `reply_to` is provided, do not include `reply-to` in headers
- `message-id` header will be ignored and automatically generated
- Schedule times must be in RFC3339 format

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@return ApiCreateMessageRequest
*/
func (a *MessagesAPIService) CreateMessage(ctx context.Context, accountId uuid.UUID) ApiCreateMessageRequest {
	return ApiCreateMessageRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
	}
}

// Execute executes the request
//
//	@return CreateMessageResponse
func (a *MessagesAPIService) CreateMessageExecute(r ApiCreateMessageRequest) (*CreateMessageResponse, *http.Response, error) {
	var (
		method      = http.MethodPost
		body        interface{}
		returnValue *CreateMessageResponse
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagesAPIService.CreateMessage")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/messages"
	path = strings.ReplaceAll(path, "{account_id}", url.PathEscape(parameterValueToString(r.accountId, "accountId")))

	headers := make(map[string]string)
	params := url.Values{}
	formParams := url.Values{}
	if r.createMessageRequest == nil {
		return returnValue, nil, reportError("createMessageRequest is required and must be specified")
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
	body = r.createMessageRequest
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

type ApiGetMessagesRequest struct {
	ctx             context.Context
	ApiService      *MessagesAPIService
	accountId       uuid.UUID
	sender          *string
	recipient       *string
	subject         *string
	messageIdHeader *string
	fromTime        *time.Time
	toTime          *time.Time
	limit           *int32
	cursor          *string
}

// Sender email address (must be from domain in API key scopes)
func (r ApiGetMessagesRequest) Sender(sender string) ApiGetMessagesRequest {
	r.sender = &sender
	return r
}

// Recipient email address
func (r ApiGetMessagesRequest) Recipient(recipient string) ApiGetMessagesRequest {
	r.recipient = &recipient
	return r
}

// Filter by subject text
func (r ApiGetMessagesRequest) Subject(subject string) ApiGetMessagesRequest {
	r.subject = &subject
	return r
}

// Filter by message ID header (same ID returned by CreateMessage API)
func (r ApiGetMessagesRequest) MessageIdHeader(messageIdHeader string) ApiGetMessagesRequest {
	r.messageIdHeader = &messageIdHeader
	return r
}

// Filter messages created after this time (RFC3339 format)
func (r ApiGetMessagesRequest) FromTime(fromTime time.Time) ApiGetMessagesRequest {
	r.fromTime = &fromTime
	return r
}

// Filter messages created before this time (RFC3339 format)
func (r ApiGetMessagesRequest) ToTime(toTime time.Time) ApiGetMessagesRequest {
	r.toTime = &toTime
	return r
}

// Maximum number of items to return (1-100)
func (r ApiGetMessagesRequest) Limit(limit int32) ApiGetMessagesRequest {
	r.limit = &limit
	return r
}

// Pagination cursor for the next page
func (r ApiGetMessagesRequest) Cursor(cursor string) ApiGetMessagesRequest {
	r.cursor = &cursor
	return r
}

func (r ApiGetMessagesRequest) Execute() (*PaginatedMessagesResponse, *http.Response, error) {
	return r.ApiService.GetMessagesExecute(r)
}

/*
GetMessages Get Messages

Returns a list of messages for the account. Can be filtered by various parameters.

**Query Parameters:**
- `sender`: Filter by sender email (required, must be from domain in API key scopes)
- `recipient`: Filter by recipient email
- `subject`: Filter by subject text
- `message_id_header`: Filter by message ID header (same ID returned by CreateMessage API)
- `from_time`: Filter messages created after this time (RFC3339 format)
- `to_time`: Filter messages created before this time (RFC3339 format)
- `limit`: Maximum number of items to return (1-100, default: 100)
- `cursor`: Pagination cursor for the next page

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@return ApiGetMessagesRequest
*/
func (a *MessagesAPIService) GetMessages(ctx context.Context, accountId uuid.UUID) ApiGetMessagesRequest {
	return ApiGetMessagesRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
	}
}

// Execute executes the request
//
//	@return PaginatedMessagesResponse
func (a *MessagesAPIService) GetMessagesExecute(r ApiGetMessagesRequest) (*PaginatedMessagesResponse, *http.Response, error) {
	var (
		method      = http.MethodGet
		body        interface{}
		returnValue *PaginatedMessagesResponse
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagesAPIService.GetMessages")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/messages"
	path = strings.ReplaceAll(path, "{account_id}", url.PathEscape(parameterValueToString(r.accountId, "accountId")))

	headers := make(map[string]string)
	params := url.Values{}
	formParams := url.Values{}
	if r.sender == nil {
		return returnValue, nil, reportError("sender is required and must be specified")
	}

	parameterAddToHeaderOrQuery(params, "sender", r.sender, "form", "")
	if r.recipient != nil {
		parameterAddToHeaderOrQuery(params, "recipient", r.recipient, "form", "")
	}
	if r.subject != nil {
		parameterAddToHeaderOrQuery(params, "subject", r.subject, "form", "")
	}
	if r.messageIdHeader != nil {
		parameterAddToHeaderOrQuery(params, "message_id_header", r.messageIdHeader, "form", "")
	}
	if r.fromTime != nil {
		parameterAddToHeaderOrQuery(params, "from_time", r.fromTime, "form", "")
	}
	if r.toTime != nil {
		parameterAddToHeaderOrQuery(params, "to_time", r.toTime, "form", "")
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
