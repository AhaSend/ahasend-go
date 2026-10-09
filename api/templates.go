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

// TemplatesAPIService TemplatesAPI service
type TemplatesAPIService service

/*
GetTemplates List Templates

# Returns the account's transactional templates in newest-first order

Query Parameters:
- `limit`: Maximum number of templates to return (1-100, default: 100)
- `after`: Cursor for the next page
- `before`: Cursor for the previous page

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param params GetTemplatesParams - query parameters
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return PaginatedTemplatesResponse, *http.Response, error
*/
func (a *TemplatesAPIService) GetTemplates(
	ctx context.Context,
	accountId uuid.UUID,
	params requests.GetTemplatesParams,
	opts ...RequestOption,
) (*responses.PaginatedTemplatesResponse, *http.Response, error) {
	var result responses.PaginatedTemplatesResponse

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
		PathTemplate: "/v2/accounts/{account_id}/templates",
		PathParams:   map[string]string{"account_id": accountId.String()},
		QueryParams:  queryParams,
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
GetTemplate Get Template

# Returns one transactional template, including the variables a send must supply

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param templateId Template ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return Template, *http.Response, error
*/
func (a *TemplatesAPIService) GetTemplate(
	ctx context.Context,
	accountId uuid.UUID,
	templateId uuid.UUID,
	opts ...RequestOption,
) (*responses.Template, *http.Response, error) {
	var result responses.Template

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/templates/{template_id}",
		PathParams: map[string]string{
			"account_id":  accountId.String(),
			"template_id": templateId.String(),
		},
		Result: &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
CreateTemplate Create Template

# Creates a transactional template

The request's fields go to the template's draft, so sends fail until the
draft is published. Set Publish to publish it in the same request.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param request CreateTemplateRequest - the template to create
	@param opts ...RequestOption - optional request options, including WithIdempotencyKey
	@return Template, *http.Response, error
*/
func (a *TemplatesAPIService) CreateTemplate(
	ctx context.Context,
	accountId uuid.UUID,
	request requests.CreateTemplateRequest,
	opts ...RequestOption,
) (*responses.Template, *http.Response, error) {
	var result responses.Template

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/templates",
		PathParams:   map[string]string{"account_id": accountId.String()},
		Body:         request,
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
UpdateTemplate Update Template

# Changes a transactional template

Every field but the name goes to the template's draft, the same draft the
dashboard edits; the name changes at once. Set Publish to publish the whole
draft. A request that changes nothing changes nothing, so it can be repeated.
It returns 503 when the template keeps changing while the request writes it.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param templateId Template ID
	@param request UpdateTemplateRequest - the fields to change
	@param opts ...RequestOption - optional request options, including WithIdempotencyKey
	@return Template, *http.Response, error
*/
func (a *TemplatesAPIService) UpdateTemplate(
	ctx context.Context,
	accountId uuid.UUID,
	templateId uuid.UUID,
	request requests.UpdateTemplateRequest,
	opts ...RequestOption,
) (*responses.Template, *http.Response, error) {
	var result responses.Template

	config := RequestConfig{
		Method:       http.MethodPut,
		PathTemplate: "/v2/accounts/{account_id}/templates/{template_id}",
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
DeleteTemplate Delete Template

# Deletes a transactional template

Sends of it fail from then on.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param templateId Template ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SuccessResponse, *http.Response, error
*/
func (a *TemplatesAPIService) DeleteTemplate(
	ctx context.Context,
	accountId uuid.UUID,
	templateId uuid.UUID,
	opts ...RequestOption,
) (*common.SuccessResponse, *http.Response, error) {
	var result common.SuccessResponse

	config := RequestConfig{
		Method:       http.MethodDelete,
		PathTemplate: "/v2/accounts/{account_id}/templates/{template_id}",
		PathParams:   templatePathParams(accountId, templateId),
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
GetTemplateDraft Get Template Draft

# Returns a template's draft, the changes that are not published yet

A template with no draft returns 404.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param templateId Template ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return TemplateDraft, *http.Response, error
*/
func (a *TemplatesAPIService) GetTemplateDraft(
	ctx context.Context,
	accountId uuid.UUID,
	templateId uuid.UUID,
	opts ...RequestOption,
) (*responses.TemplateDraft, *http.Response, error) {
	var result responses.TemplateDraft

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/templates/{template_id}/draft",
		PathParams:   templatePathParams(accountId, templateId),
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
DiscardTemplateDraft Discard Template Draft

# Discards a template's draft

A template with no draft is returned as it is, so the request can be repeated.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param templateId Template ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return Template, *http.Response, error
*/
func (a *TemplatesAPIService) DiscardTemplateDraft(
	ctx context.Context,
	accountId uuid.UUID,
	templateId uuid.UUID,
	opts ...RequestOption,
) (*responses.Template, *http.Response, error) {
	var result responses.Template

	config := RequestConfig{
		Method:       http.MethodDelete,
		PathTemplate: "/v2/accounts/{account_id}/templates/{template_id}/draft",
		PathParams:   templatePathParams(accountId, templateId),
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
PublishTemplate Publish Template

# Publishes a template's draft and keeps it as a new version

The whole draft is published, including changes made in the dashboard. A
template with no draft is returned as it is when its published copy has a
body, so the request can be repeated.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param templateId Template ID
	@param opts ...RequestOption - optional request options, including WithIdempotencyKey
	@return Template, *http.Response, error
*/
func (a *TemplatesAPIService) PublishTemplate(
	ctx context.Context,
	accountId uuid.UUID,
	templateId uuid.UUID,
	opts ...RequestOption,
) (*responses.Template, *http.Response, error) {
	var result responses.Template

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/templates/{template_id}/publish",
		PathParams:   templatePathParams(accountId, templateId),
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
GetTemplateVersions List Template Versions

# Returns a template's published versions, newest first, without their content

A template keeps its 50 most recent versions, so the list is not paginated.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param templateId Template ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return TemplateVersionsResponse, *http.Response, error
*/
func (a *TemplatesAPIService) GetTemplateVersions(
	ctx context.Context,
	accountId uuid.UUID,
	templateId uuid.UUID,
	opts ...RequestOption,
) (*responses.TemplateVersionsResponse, *http.Response, error) {
	var result responses.TemplateVersionsResponse

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/templates/{template_id}/versions",
		PathParams:   templatePathParams(accountId, templateId),
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
GetTemplateVersion Get Template Version

# Returns one published version of a template with its content

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param templateId Template ID
	@param versionId Version ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return TemplateVersionDetail, *http.Response, error
*/
func (a *TemplatesAPIService) GetTemplateVersion(
	ctx context.Context,
	accountId uuid.UUID,
	templateId uuid.UUID,
	versionId uuid.UUID,
	opts ...RequestOption,
) (*responses.TemplateVersionDetail, *http.Response, error) {
	var result responses.TemplateVersionDetail

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/templates/{template_id}/versions/{version_id}",
		PathParams:   templateVersionPathParams(accountId, templateId, versionId),
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
RestoreTemplateVersion Restore Template Version

# Restores a published version into the template's draft

Set Publish to publish the draft once the version is restored. A version the
publish checks refuse is refused before the restore, and the draft is left as
it was.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param templateId Template ID
	@param versionId Version ID
	@param request RestoreTemplateVersionRequest - the restore options
	@param opts ...RequestOption - optional request options, including WithIdempotencyKey
	@return Template, *http.Response, error
*/
func (a *TemplatesAPIService) RestoreTemplateVersion(
	ctx context.Context,
	accountId uuid.UUID,
	templateId uuid.UUID,
	versionId uuid.UUID,
	request requests.RestoreTemplateVersionRequest,
	opts ...RequestOption,
) (*responses.Template, *http.Response, error) {
	var result responses.Template

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/templates/{template_id}/versions/{version_id}/restore",
		PathParams:   templateVersionPathParams(accountId, templateId, versionId),
		Body:         request,
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

func templatePathParams(accountId, templateId uuid.UUID) map[string]string {
	return map[string]string{
		"account_id":  accountId.String(),
		"template_id": templateId.String(),
	}
}

func templateVersionPathParams(accountId, templateId, versionId uuid.UUID) map[string]string {
	params := templatePathParams(accountId, templateId)
	params["version_id"] = versionId.String()
	return params
}
