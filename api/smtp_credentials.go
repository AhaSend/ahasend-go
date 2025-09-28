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

// SMTPCredentialsAPIService SMTPCredentialsAPI service
type SMTPCredentialsAPIService service

/*
CreateSMTPCredential Create SMTP Credential

# Creates a new SMTP credential for SMTP authentication

Validation Requirements:
- `name` must be a unique credential name within the account
- `password` will be auto-generated and returned in the response

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param request CreateSMTPCredentialRequest - The SMTP credential details to create
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SMTPCredential, *http.Response, error
*/
func (a *SMTPCredentialsAPIService) CreateSMTPCredential(
	ctx context.Context,
	accountId uuid.UUID,
	request requests.CreateSMTPCredentialRequest,
	opts ...RequestOption,
) (*responses.SMTPCredential, *http.Response, error) {
	var result responses.SMTPCredential

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/smtp-credentials",
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
DeleteSMTPCredential Delete SMTP Credential

Deletes an SMTP credential

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param smtpCredentialId SMTP Credential ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SuccessResponse, *http.Response, error
*/
func (a *SMTPCredentialsAPIService) DeleteSMTPCredential(
	ctx context.Context,
	accountId uuid.UUID,
	smtpCredentialId uint64,
	opts ...RequestOption,
) (*common.SuccessResponse, *http.Response, error) {
	var result common.SuccessResponse

	config := RequestConfig{
		Method:       http.MethodDelete,
		PathTemplate: "/v2/accounts/{account_id}/smtp-credentials/{smtp_credential_id}",
		PathParams: map[string]string{
			"account_id":         accountId.String(),
			"smtp_credential_id": fmt.Sprintf("%d", smtpCredentialId),
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
GetSMTPCredential Get SMTP Credential

Returns a specific SMTP credential by ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param smtpCredentialId SMTP Credential ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SMTPCredential, *http.Response, error
*/
func (a *SMTPCredentialsAPIService) GetSMTPCredential(
	ctx context.Context,
	accountId uuid.UUID,
	smtpCredentialId uint64,
	opts ...RequestOption,
) (*responses.SMTPCredential, *http.Response, error) {
	var result responses.SMTPCredential

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/smtp-credentials/{smtp_credential_id}",
		PathParams: map[string]string{
			"account_id":         accountId.String(),
			"smtp_credential_id": fmt.Sprintf("%d", smtpCredentialId),
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
GetSMTPCredentials Get SMTP Credentials

Returns a list of SMTP credentials for the account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param pagination Pagination parameters (optional)
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return PaginatedSMTPCredentialsResponse, *http.Response, error
*/
func (a *SMTPCredentialsAPIService) GetSMTPCredentials(
	ctx context.Context,
	accountId uuid.UUID,
	pagination *common.PaginationParams,
	opts ...RequestOption,
) (*responses.PaginatedSMTPCredentialsResponse, *http.Response, error) {
	var result responses.PaginatedSMTPCredentialsResponse

	// Build query parameters
	queryParams := url.Values{}

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
		PathTemplate: "/v2/accounts/{account_id}/smtp-credentials",
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
