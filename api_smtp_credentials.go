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

// SMTPCredentialsAPIService SMTPCredentialsAPI service
type SMTPCredentialsAPIService service

type ApiCreateSMTPCredentialRequest struct {
	ctx                         context.Context
	ApiService                  *SMTPCredentialsAPIService
	createSMTPCredentialRequest *CreateSMTPCredentialRequest
	accountId                   uuid.UUID
	idempotencyKey              *string
}

func (r ApiCreateSMTPCredentialRequest) CreateSMTPCredentialRequest(createSMTPCredentialRequest CreateSMTPCredentialRequest) ApiCreateSMTPCredentialRequest {
	r.createSMTPCredentialRequest = &createSMTPCredentialRequest
	return r
}

// Optional idempotency key for safe request retries. Must be a unique string for each logical request. Requests with the same key will return the same response. Keys expire after 24 hours.
func (r ApiCreateSMTPCredentialRequest) IdempotencyKey(idempotencyKey string) ApiCreateSMTPCredentialRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiCreateSMTPCredentialRequest) Execute() (*SMTPCredential, *http.Response, error) {
	return r.ApiService.CreateSMTPCredentialExecute(r)
}

/*
CreateSMTPCredential Create SMTP Credential

Creates a new SMTP credential for SMTP authentication

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@return ApiCreateSMTPCredentialRequest
*/
func (a *SMTPCredentialsAPIService) CreateSMTPCredential(ctx context.Context, accountId uuid.UUID) ApiCreateSMTPCredentialRequest {
	return ApiCreateSMTPCredentialRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
	}
}

// Execute executes the request
//
//	@return SMTPCredential
func (a *SMTPCredentialsAPIService) CreateSMTPCredentialExecute(r ApiCreateSMTPCredentialRequest) (*SMTPCredential, *http.Response, error) {
	var (
		method = http.MethodPost
		body   interface{}

		returnValue *SMTPCredential
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SMTPCredentialsAPIService.CreateSMTPCredential")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/smtp-credentials"
	path = strings.ReplaceAll(path, "{account_id}", url.PathEscape(parameterValueToString(r.accountId, "accountId")))

	headers := make(map[string]string)
	params := url.Values{}
	formParams := url.Values{}
	if r.createSMTPCredentialRequest == nil {
		return returnValue, nil, NewRequiredFieldError("createSMTPCredentialRequest")
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
	body = r.createSMTPCredentialRequest
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

type ApiDeleteSMTPCredentialRequest struct {
	ctx              context.Context
	ApiService       *SMTPCredentialsAPIService
	accountId        uuid.UUID
	smtpCredentialId uuid.UUID
}

func (r ApiDeleteSMTPCredentialRequest) Execute() (*SuccessResponse, *http.Response, error) {
	return r.ApiService.DeleteSMTPCredentialExecute(r)
}

/*
DeleteSMTPCredential Delete SMTP Credential

Deletes an SMTP credential

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param smtpCredentialId SMTP Credential ID
	@return ApiDeleteSMTPCredentialRequest
*/
func (a *SMTPCredentialsAPIService) DeleteSMTPCredential(ctx context.Context, accountId uuid.UUID, smtpCredentialId uuid.UUID) ApiDeleteSMTPCredentialRequest {
	return ApiDeleteSMTPCredentialRequest{
		ApiService:       a,
		ctx:              ctx,
		accountId:        accountId,
		smtpCredentialId: smtpCredentialId,
	}
}

// Execute executes the request
//
//	@return SuccessResponse
func (a *SMTPCredentialsAPIService) DeleteSMTPCredentialExecute(r ApiDeleteSMTPCredentialRequest) (*SuccessResponse, *http.Response, error) {
	var (
		method = http.MethodDelete
		body   interface{}

		returnValue *SuccessResponse
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SMTPCredentialsAPIService.DeleteSMTPCredential")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/smtp-credentials/{smtp_credential_id}"
	path = strings.ReplaceAll(path, "{account_id}", url.PathEscape(parameterValueToString(r.accountId, "accountId")))
	path = strings.ReplaceAll(path, "{smtp_credential_id}", url.PathEscape(parameterValueToString(r.smtpCredentialId, "smtpCredentialId")))

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

type ApiGetSMTPCredentialRequest struct {
	ctx              context.Context
	ApiService       *SMTPCredentialsAPIService
	accountId        uuid.UUID
	smtpCredentialId uuid.UUID
}

func (r ApiGetSMTPCredentialRequest) Execute() (*SMTPCredential, *http.Response, error) {
	return r.ApiService.GetSMTPCredentialExecute(r)
}

/*
GetSMTPCredential Get SMTP Credential

Returns a specific SMTP credential by ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param smtpCredentialId SMTP Credential ID
	@return ApiGetSMTPCredentialRequest
*/
func (a *SMTPCredentialsAPIService) GetSMTPCredential(ctx context.Context, accountId uuid.UUID, smtpCredentialId uuid.UUID) ApiGetSMTPCredentialRequest {
	return ApiGetSMTPCredentialRequest{
		ApiService:       a,
		ctx:              ctx,
		accountId:        accountId,
		smtpCredentialId: smtpCredentialId,
	}
}

// Execute executes the request
//
//	@return SMTPCredential
func (a *SMTPCredentialsAPIService) GetSMTPCredentialExecute(r ApiGetSMTPCredentialRequest) (*SMTPCredential, *http.Response, error) {
	var (
		method = http.MethodGet
		body   interface{}

		returnValue *SMTPCredential
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SMTPCredentialsAPIService.GetSMTPCredential")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/smtp-credentials/{smtp_credential_id}"
	path = strings.ReplaceAll(path, "{account_id}", url.PathEscape(parameterValueToString(r.accountId, "accountId")))
	path = strings.ReplaceAll(path, "{smtp_credential_id}", url.PathEscape(parameterValueToString(r.smtpCredentialId, "smtpCredentialId")))

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

type ApiGetSMTPCredentialsRequest struct {
	ctx        context.Context
	ApiService *SMTPCredentialsAPIService
	accountId  uuid.UUID
	limit      *int32
	cursor     *string
}

// Maximum number of items to return (1-100)
func (r ApiGetSMTPCredentialsRequest) Limit(limit int32) ApiGetSMTPCredentialsRequest {
	r.limit = &limit
	return r
}

// Pagination cursor for the next page
func (r ApiGetSMTPCredentialsRequest) Cursor(cursor string) ApiGetSMTPCredentialsRequest {
	r.cursor = &cursor
	return r
}

func (r ApiGetSMTPCredentialsRequest) Execute() (*PaginatedSMTPCredentialsResponse, *http.Response, error) {
	return r.ApiService.GetSMTPCredentialsExecute(r)
}

/*
GetSMTPCredentials Get SMTP Credentials

Returns a list of SMTP credentials for the account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@return ApiGetSMTPCredentialsRequest
*/
func (a *SMTPCredentialsAPIService) GetSMTPCredentials(ctx context.Context, accountId uuid.UUID) ApiGetSMTPCredentialsRequest {
	return ApiGetSMTPCredentialsRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
	}
}

// Execute executes the request
//
//	@return PaginatedSMTPCredentialsResponse
func (a *SMTPCredentialsAPIService) GetSMTPCredentialsExecute(r ApiGetSMTPCredentialsRequest) (*PaginatedSMTPCredentialsResponse, *http.Response, error) {
	var (
		method = http.MethodGet
		body   interface{}

		returnValue *PaginatedSMTPCredentialsResponse
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SMTPCredentialsAPIService.GetSMTPCredentials")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/smtp-credentials"
	path = strings.ReplaceAll(path, "{account_id}", url.PathEscape(parameterValueToString(r.accountId, "accountId")))

	headers := make(map[string]string)
	params := url.Values{}
	formParams := url.Values{}

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
