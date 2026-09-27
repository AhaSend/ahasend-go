package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

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
