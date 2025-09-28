package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/AhaSend/ahasend-go/models/responses"
	"github.com/google/uuid"
)

// WebhooksAPIService WebhooksAPI service
type WebhooksAPIService service

/*
CreateWebhook Create Webhook

# Creates a new webhook for event notifications

Validation Requirements:
- `url` must be a valid HTTPS URL
- `events` array must contain at least one event type

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param request CreateWebhookRequest - The webhook details to create
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return Webhook, *http.Response, error
*/
func (a *WebhooksAPIService) CreateWebhook(
	ctx context.Context,
	accountId uuid.UUID,
	request requests.CreateWebhookRequest,
	opts ...RequestOption,
) (*responses.Webhook, *http.Response, error) {
	var result responses.Webhook

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/webhooks",
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
DeleteWebhook Delete Webhook

Deletes a webhook

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param webhookId Webhook ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SuccessResponse, *http.Response, error
*/
func (a *WebhooksAPIService) DeleteWebhook(
	ctx context.Context,
	accountId uuid.UUID,
	webhookId uuid.UUID,
	opts ...RequestOption,
) (*common.SuccessResponse, *http.Response, error) {
	var result common.SuccessResponse

	config := RequestConfig{
		Method:       http.MethodDelete,
		PathTemplate: "/v2/accounts/{account_id}/webhooks/{webhook_id}",
		PathParams: map[string]string{
			"account_id": accountId.String(),
			"webhook_id": webhookId.String(),
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
GetWebhook Get Webhook

Returns a specific webhook by ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param webhookId Webhook ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return Webhook, *http.Response, error
*/
func (a *WebhooksAPIService) GetWebhook(
	ctx context.Context,
	accountId uuid.UUID,
	webhookId uuid.UUID,
	opts ...RequestOption,
) (*responses.Webhook, *http.Response, error) {
	var result responses.Webhook

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/webhooks/{webhook_id}",
		PathParams: map[string]string{
			"account_id": accountId.String(),
			"webhook_id": webhookId.String(),
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

type GetWebhooksParams struct {
	Enabled              *bool
	OnReception          *bool
	OnDelivered          *bool
	OnTransientError     *bool
	OnFailed             *bool
	OnBounced            *bool
	OnSuppressed         *bool
	OnOpened             *bool
	OnClicked            *bool
	OnSuppressionCreated *bool
	OnDnsError           *bool
	common.PaginationParams
}

/*
GetWebhooks Get Webhooks

# Returns a list of webhooks for the account

**Query Parameters:**
- `enabled`: Filter by enabled status
- Event filters: `on_reception`, `on_delivered`, `on_transient_error`, `on_failed`, `on_bounced`, `on_suppressed`, `on_opened`, `on_clicked`, `on_suppression_created`, `on_dns_error`
- `limit`: Maximum number of items to return (1-100, default: 100)
- `cursor`: Pagination cursor for the next page (backward compatibility)
- `after`: Get items after this cursor (forward pagination)
- `before`: Get items before this cursor (backward pagination)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param params GetWebhooksParams - query parameters
	@return PaginatedWebhooksResponse, *http.Response, error
*/
func (a *WebhooksAPIService) GetWebhooks(
	ctx context.Context,
	accountId uuid.UUID,
	params GetWebhooksParams,
	opts ...RequestOption,
) (*responses.PaginatedWebhooksResponse, *http.Response, error) {
	var result responses.PaginatedWebhooksResponse

	// Build query parameters
	queryParams := url.Values{}
	if params.Enabled != nil {
		queryParams.Set("enabled", fmt.Sprintf("%t", *params.Enabled))
	}
	if params.OnReception != nil {
		queryParams.Set("on_reception", fmt.Sprintf("%t", *params.OnReception))
	}
	if params.OnDelivered != nil {
		queryParams.Set("on_delivered", fmt.Sprintf("%t", *params.OnDelivered))
	}
	if params.OnTransientError != nil {
		queryParams.Set("on_transient_error", fmt.Sprintf("%t", *params.OnTransientError))
	}
	if params.OnFailed != nil {
		queryParams.Set("on_failed", fmt.Sprintf("%t", *params.OnFailed))
	}
	if params.OnBounced != nil {
		queryParams.Set("on_bounced", fmt.Sprintf("%t", *params.OnBounced))
	}
	if params.OnSuppressed != nil {
		queryParams.Set("on_suppressed", fmt.Sprintf("%t", *params.OnSuppressed))
	}
	if params.OnOpened != nil {
		queryParams.Set("on_opened", fmt.Sprintf("%t", *params.OnOpened))
	}
	if params.OnClicked != nil {
		queryParams.Set("on_clicked", fmt.Sprintf("%t", *params.OnClicked))
	}
	if params.OnSuppressionCreated != nil {
		queryParams.Set("on_suppression_created", fmt.Sprintf("%t", *params.OnSuppressionCreated))
	}
	if params.OnDnsError != nil {
		queryParams.Set("on_dns_error", fmt.Sprintf("%t", *params.OnDnsError))
	}
	// Handle pagination parameters
	if params.Limit != nil {
		queryParams.Set("limit", fmt.Sprintf("%d", *params.Limit))
	} else {
		queryParams.Set("limit", "100") // Default value
	}

	// Handle pagination parameters - prioritize after/before over cursor for backward compatibility
	if params.After != nil {
		queryParams.Set("after", *params.After)
	} else if params.Before != nil {
		queryParams.Set("before", *params.Before)
	} else if params.Cursor != nil {
		queryParams.Set("cursor", *params.Cursor)
	}

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/webhooks",
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

/*
UpdateWebhook Update Webhook

# Updates an existing webhook

**Note:** The webhook secret is not updatable

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param webhookId Webhook ID
	@param request UpdateWebhookRequest - The webhook details to update
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return Webhook, *http.Response, error
*/
func (a *WebhooksAPIService) UpdateWebhook(
	ctx context.Context,
	accountId uuid.UUID,
	webhookId uuid.UUID,
	request requests.UpdateWebhookRequest,
	opts ...RequestOption,
) (*responses.Webhook, *http.Response, error) {
	var result responses.Webhook

	config := RequestConfig{
		Method:       http.MethodPut,
		PathTemplate: "/v2/accounts/{account_id}/webhooks/{webhook_id}",
		PathParams: map[string]string{
			"account_id": accountId.String(),
			"webhook_id": webhookId.String(),
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
