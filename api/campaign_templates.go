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

// CampaignTemplatesAPIService CampaignTemplatesAPI service
type CampaignTemplatesAPIService service

/*
ListCampaignTemplates List Campaign Templates

# Returns the account's campaign templates in newest-first order, each without its content

Transactional templates are a different resource and never appear here.

Query Parameters:
- `limit`: Maximum number of campaign templates to return (1-100, default: 100)
- `after`: Cursor for the next page
- `before`: Cursor for the previous page

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param params ListCampaignTemplatesParams - query parameters
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return PaginatedCampaignTemplatesResponse, *http.Response, error
*/
func (a *CampaignTemplatesAPIService) ListCampaignTemplates(
	ctx context.Context,
	accountId uuid.UUID,
	params requests.ListCampaignTemplatesParams,
	opts ...RequestOption,
) (*responses.PaginatedCampaignTemplatesResponse, *http.Response, error) {
	var result responses.PaginatedCampaignTemplatesResponse

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

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/campaign-templates",
		PathParams:   map[string]string{"account_id": accountId.String()},
		QueryParams:  queryParams,
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
GetCampaignTemplate Get Campaign Template

# Returns one of the account's campaign templates with its content

A transactional template's ID returns 404.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param templateId Campaign template ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return CampaignTemplate, *http.Response, error
*/
func (a *CampaignTemplatesAPIService) GetCampaignTemplate(
	ctx context.Context,
	accountId uuid.UUID,
	templateId uuid.UUID,
	opts ...RequestOption,
) (*responses.CampaignTemplate, *http.Response, error) {
	var result responses.CampaignTemplate

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/campaign-templates/{template_id}",
		PathParams:   templatePathParams(accountId, templateId),
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
CreateCampaignTemplate Create Campaign Template

# Creates a campaign template

A campaign template has no draft: what the request sets is what a campaign
made from the template copies. Without Content it is created with an empty
design of its Editor, which the dashboard's editor opens. It returns 503 when
the AhaML in Content.Markup cannot be read now; retry after Retry-After.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param request CreateCampaignTemplateRequest - the campaign template to create
	@param opts ...RequestOption - optional request options, including WithIdempotencyKey
	@return CampaignTemplate, *http.Response, error
*/
func (a *CampaignTemplatesAPIService) CreateCampaignTemplate(
	ctx context.Context,
	accountId uuid.UUID,
	request requests.CreateCampaignTemplateRequest,
	opts ...RequestOption,
) (*responses.CampaignTemplate, *http.Response, error) {
	var result responses.CampaignTemplate

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/campaign-templates",
		PathParams:   map[string]string{"account_id": accountId.String()},
		Body:         request,
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
UpdateCampaignTemplate Update Campaign Template

# Changes a campaign template's name, its content, or both

The change applies at once; campaigns already made from the template keep
their own copies. A request that changes nothing changes nothing, so it can be
repeated. A design saved later in the dashboard, also from an editor opened
before this request, replaces the design this request wrote. It returns 503
when the AhaML in Content.Markup cannot be read now; retry after Retry-After.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param templateId Campaign template ID
	@param request UpdateCampaignTemplateRequest - the fields to change
	@param opts ...RequestOption - optional request options, including WithIdempotencyKey
	@return CampaignTemplate, *http.Response, error
*/
func (a *CampaignTemplatesAPIService) UpdateCampaignTemplate(
	ctx context.Context,
	accountId uuid.UUID,
	templateId uuid.UUID,
	request requests.UpdateCampaignTemplateRequest,
	opts ...RequestOption,
) (*responses.CampaignTemplate, *http.Response, error) {
	var result responses.CampaignTemplate

	config := RequestConfig{
		Method:       http.MethodPut,
		PathTemplate: "/v2/accounts/{account_id}/campaign-templates/{template_id}",
		PathParams:   templatePathParams(accountId, templateId),
		Body:         request,
		Result:       &result,
		idempotent:   true,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
DeleteCampaignTemplate Delete Campaign Template

# Deletes a campaign template

Campaigns made from it keep their own copies, and a second delete returns 404.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param templateId Campaign template ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SuccessResponse, *http.Response, error
*/
func (a *CampaignTemplatesAPIService) DeleteCampaignTemplate(
	ctx context.Context,
	accountId uuid.UUID,
	templateId uuid.UUID,
	opts ...RequestOption,
) (*common.SuccessResponse, *http.Response, error) {
	var result common.SuccessResponse

	config := RequestConfig{
		Method:       http.MethodDelete,
		PathTemplate: "/v2/accounts/{account_id}/campaign-templates/{template_id}",
		PathParams:   templatePathParams(accountId, templateId),
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}
