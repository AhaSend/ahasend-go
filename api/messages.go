package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/AhaSend/ahasend-go/models/responses"
	"github.com/google/uuid"
)

// MessagesAPIService MessagesAPI service
type MessagesAPIService service

/*
CancelMessage Cancel Message

Cancels a scheduled message

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param messageId Message ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SuccessResponse, *http.Response, error
*/
func (a *MessagesAPIService) CancelMessage(
	ctx context.Context,
	accountId uuid.UUID,
	messageId string,
	opts ...RequestOption,
) (*common.SuccessResponse, *http.Response, error) {
	var result common.SuccessResponse

	config := RequestConfig{
		Method:       http.MethodDelete,
		PathTemplate: "/v2/accounts/{account_id}/messages/{message_id}/cancel",
		PathParams: map[string]string{
			"account_id": accountId.String(),
			"message_id": messageId,
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
	@param request CreateMessageRequest - The message details to create
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return CreateMessageResponse, *http.Response, error
*/
func (a *MessagesAPIService) CreateMessage(
	ctx context.Context,
	accountId uuid.UUID,
	request requests.CreateMessageRequest,
	opts ...RequestOption,
) (*responses.CreateMessageResponse, *http.Response, error) {
	var result responses.CreateMessageResponse

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/messages",
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
GetMessages Get Messages

Returns a list of messages for the account. Can be filtered by various parameters.

**Query Parameters:**
- `status`: Filter by comma-separated list of message statuses
- `sender`: Filter by sender email (must be from domain in API key scopes)
- `recipient`: Filter by recipient email
- `subject`: Filter by subject text
- `message_id_header`: Filter by message ID header (same ID returned by CreateMessage API)
- `from_time`: Filter messages created after this time (RFC3339 format)
- `to_time`: Filter messages created before this time (RFC3339 format)
- `limit`: Maximum number of items to return (1-100, default: 100)
- `cursor`: Pagination cursor for the next page

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param params GetMessagesParams - query parameters
	@return PaginatedMessagesResponse, *http.Response, error
*/
func (a *MessagesAPIService) GetMessages(
	ctx context.Context,
	accountId uuid.UUID,
	params requests.GetMessagesParams,
	opts ...RequestOption,
) (*responses.PaginatedMessagesResponse, *http.Response, error) {
	var result responses.PaginatedMessagesResponse

	// Build query parameters
	queryParams := url.Values{}
	if params.Status != nil {
		queryParams.Set("status", *params.Status)
	}
	if params.Sender != nil {
		queryParams.Set("sender", *params.Sender)
	}
	if params.Recipient != nil {
		queryParams.Set("recipient", *params.Recipient)
	}
	if params.Subject != nil {
		queryParams.Set("subject", *params.Subject)
	}
	if params.MessageIDHeader != nil {
		queryParams.Set("message_id_header", *params.MessageIDHeader)
	}
	if params.FromTime != nil {
		queryParams.Set("from_time", params.FromTime.Format(time.RFC3339))
	}
	if params.ToTime != nil {
		queryParams.Set("to_time", params.ToTime.Format(time.RFC3339))
	}
	if params.Limit != nil {
		queryParams.Set("limit", fmt.Sprintf("%d", *params.Limit))
	} else {
		queryParams.Set("limit", "100") // Default value
	}
	if params.Cursor != nil {
		queryParams.Set("cursor", *params.Cursor)
	}

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/messages",
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

func (a *MessagesAPIService) GetMessage(
	ctx context.Context,
	accountId uuid.UUID,
	messageId string,
	opts ...RequestOption,
) (*responses.Message, *http.Response, error) {
	var result responses.Message

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/messages/{message_id}",
		PathParams: map[string]string{
			"account_id": accountId.String(),
			"message_id": messageId,
		},
	}

	// Apply options
	for _, opt := range opts {
		opt(&config)
	}

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}
