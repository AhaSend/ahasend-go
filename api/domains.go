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

// DomainsAPIService DomainsAPI service
type DomainsAPIService service

/*
CreateDomain Create Domain

# Creates a new domain

Validation Requirements:
- `domain` must be a valid domain name
- `domain` must not already exist in the account
- `dkim_private_key` is optional and must be a valid DKIM RSA private key with a minimum key length of 2048 bits if provided.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param request CreateDomainRequest - The domain details to create
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return Domain, *http.Response, error
*/
func (a *DomainsAPIService) CreateDomain(
	ctx context.Context,
	accountId uuid.UUID,
	request requests.CreateDomainRequest,
	opts ...RequestOption,
) (*responses.Domain, *http.Response, error) {
	var result responses.Domain

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/domains",
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
DeleteDomain Delete Domain

Deletes a domain

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param domain Domain name
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SuccessResponse, *http.Response, error
*/
func (a *DomainsAPIService) DeleteDomain(
	ctx context.Context,
	accountId uuid.UUID,
	domain string,
	opts ...RequestOption,
) (*common.SuccessResponse, *http.Response, error) {
	var result common.SuccessResponse

	config := RequestConfig{
		Method:       http.MethodDelete,
		PathTemplate: "/v2/accounts/{account_id}/domains/{domain}",
		PathParams: map[string]string{
			"account_id": accountId.String(),
			"domain":     domain,
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
GetDomain Get Domain

Returns a specific domain by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param domain Domain name
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return Domain, *http.Response, error
*/
func (a *DomainsAPIService) GetDomain(
	ctx context.Context,
	accountId uuid.UUID,
	domain string,
	opts ...RequestOption,
) (*responses.Domain, *http.Response, error) {
	var result responses.Domain

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/domains/{domain}",
		PathParams: map[string]string{
			"account_id": accountId.String(),
			"domain":     domain,
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
GetDomains Get Domains

Returns a list of domains for the account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param dnsValid Filter results by DNS validation status (optional)
	@param pagination Pagination parameters (optional)
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return PaginatedDomainsResponse, *http.Response, error
*/
func (a *DomainsAPIService) GetDomains(
	ctx context.Context,
	accountId uuid.UUID,
	dnsValid *bool,
	pagination *common.PaginationParams,
	opts ...RequestOption,
) (*responses.PaginatedDomainsResponse, *http.Response, error) {
	var result responses.PaginatedDomainsResponse

	// Build query parameters
	queryParams := url.Values{}
	if dnsValid != nil {
		queryParams.Set("dns_valid", fmt.Sprintf("%t", *dnsValid))
	}

	// Handle pagination parameters
	if pagination != nil {
		if pagination.Limit != nil {
			queryParams.Set("limit", fmt.Sprintf("%d", *pagination.Limit))
		} else {
			queryParams.Set("limit", "100") // Default value
		}

		// Handle pagination parameters - prioritize after/before over cursor for backward compatibility
		if pagination.After != nil {
			queryParams.Set("after", *pagination.After)
		} else if pagination.Before != nil {
			queryParams.Set("before", *pagination.Before)
		} else if pagination.Cursor != nil {
			queryParams.Set("cursor", *pagination.Cursor)
		}
	} else {
		queryParams.Set("limit", "100") // Default value
	}

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/domains",
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
