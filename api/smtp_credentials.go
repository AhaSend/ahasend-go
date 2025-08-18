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
	smtpCredentialId uuid.UUID,
	opts ...RequestOption,
) (*common.SuccessResponse, *http.Response, error) {
	var result common.SuccessResponse

	config := RequestConfig{
		Method:       http.MethodDelete,
		PathTemplate: "/v2/accounts/{account_id}/smtp-credentials/{smtp_credential_id}",
		PathParams: map[string]string{
			"account_id":         accountId.String(),
			"smtp_credential_id": smtpCredentialId.String(),
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
	smtpCredentialId uuid.UUID,
	opts ...RequestOption,
) (*responses.SMTPCredential, *http.Response, error) {
	var result responses.SMTPCredential

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/smtp-credentials/{smtp_credential_id}",
		PathParams: map[string]string{
			"account_id":         accountId.String(),
			"smtp_credential_id": smtpCredentialId.String(),
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
	@param limit Maximum number of items to return (1-100, default: 100) (optional)
	@param cursor Pagination cursor for the next page (optional)
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return PaginatedSMTPCredentialsResponse, *http.Response, error
*/
func (a *SMTPCredentialsAPIService) GetSMTPCredentials(
	ctx context.Context,
	accountId uuid.UUID,
	limit *int32,
	cursor *string,
	opts ...RequestOption,
) (*responses.PaginatedSMTPCredentialsResponse, *http.Response, error) {
	var result responses.PaginatedSMTPCredentialsResponse

	// Build query parameters
	queryParams := url.Values{}
	if limit != nil {
		queryParams.Set("limit", fmt.Sprintf("%d", *limit))
	} else {
		queryParams.Set("limit", "100") // Default value
	}
	if cursor != nil {
		queryParams.Set("cursor", *cursor)
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
