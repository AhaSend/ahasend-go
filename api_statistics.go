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

// StatisticsAPIService StatisticsAPI service
type StatisticsAPIService service

type ApiGetBounceStatisticsRequest struct {
	ctx             context.Context
	ApiService      *StatisticsAPIService
	accountId       uuid.UUID
	fromDate        *time.Time
	toDate          *time.Time
	senderDomain    *string
	recipientDomain *string
	tags            *string
	groupBy         *string
}

// Filter statistics after this date (RFC3339 format)
func (r ApiGetBounceStatisticsRequest) FromDate(fromDate time.Time) ApiGetBounceStatisticsRequest {
	r.fromDate = &fromDate
	return r
}

// Filter statistics before this date (RFC3339 format)
func (r ApiGetBounceStatisticsRequest) ToDate(toDate time.Time) ApiGetBounceStatisticsRequest {
	r.toDate = &toDate
	return r
}

// Filter by sender domain
func (r ApiGetBounceStatisticsRequest) SenderDomain(senderDomain string) ApiGetBounceStatisticsRequest {
	r.senderDomain = &senderDomain
	return r
}

// Filter by recipient domain
func (r ApiGetBounceStatisticsRequest) RecipientDomain(recipientDomain string) ApiGetBounceStatisticsRequest {
	r.recipientDomain = &recipientDomain
	return r
}

// Filter by tags (comma-separated)
func (r ApiGetBounceStatisticsRequest) Tags(tags string) ApiGetBounceStatisticsRequest {
	r.tags = &tags
	return r
}

// Group by time period
func (r ApiGetBounceStatisticsRequest) GroupBy(groupBy string) ApiGetBounceStatisticsRequest {
	r.groupBy = &groupBy
	return r
}

func (r ApiGetBounceStatisticsRequest) Execute() (*BounceStatisticsResponse, *http.Response, error) {
	return r.ApiService.GetBounceStatisticsExecute(r)
}

/*
GetBounceStatistics Get Bounce Statistics

# Returns transactional bounce statistics grouped by classification

**Query Parameters:**
- `from_date`: Filter statistics after this date (RFC3339 format)
- `to_date`: Filter statistics before this date (RFC3339 format)
- `sender_domain`: Filter by sender domain
- `recipient_domain`: Filter by recipient domain
- `tags`: Filter by tags (comma-separated)
- `group_by`: Group by time period (hour, day, week, month)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@return ApiGetBounceStatisticsRequest
*/
func (a *StatisticsAPIService) GetBounceStatistics(ctx context.Context, accountId uuid.UUID) ApiGetBounceStatisticsRequest {
	return ApiGetBounceStatisticsRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
	}
}

// Execute executes the request
//
//	@return BounceStatisticsResponse
func (a *StatisticsAPIService) GetBounceStatisticsExecute(r ApiGetBounceStatisticsRequest) (*BounceStatisticsResponse, *http.Response, error) {
	var (
		method = http.MethodGet
		body   interface{}

		returnValue *BounceStatisticsResponse
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatisticsAPIService.GetBounceStatistics")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/statistics/transactional/bounce"
	path = strings.ReplaceAll(path, "{account_id}", url.PathEscape(parameterValueToString(r.accountId, "accountId")))

	headers := make(map[string]string)
	params := url.Values{}
	formParams := url.Values{}

	if r.fromDate != nil {
		parameterAddToHeaderOrQuery(params, "from_date", r.fromDate, "form", "")
	}
	if r.toDate != nil {
		parameterAddToHeaderOrQuery(params, "to_date", r.toDate, "form", "")
	}
	if r.senderDomain != nil {
		parameterAddToHeaderOrQuery(params, "sender_domain", r.senderDomain, "form", "")
	}
	if r.recipientDomain != nil {
		parameterAddToHeaderOrQuery(params, "recipient_domain", r.recipientDomain, "form", "")
	}
	if r.tags != nil {
		parameterAddToHeaderOrQuery(params, "tags", r.tags, "form", "")
	}
	if r.groupBy != nil {
		parameterAddToHeaderOrQuery(params, "group_by", r.groupBy, "form", "")
	} else {
		defaultValue := "day"
		r.groupBy = &defaultValue
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
		if response.StatusCode == 429 {
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

type ApiGetDeliverabilityStatisticsRequest struct {
	ctx             context.Context
	ApiService      *StatisticsAPIService
	accountId       uuid.UUID
	fromDate        *time.Time
	toDate          *time.Time
	senderDomain    *string
	recipientDomain *string
	tags            *string
	groupBy         *string
}

// Filter statistics after this date (RFC3339 format)
func (r ApiGetDeliverabilityStatisticsRequest) FromDate(fromDate time.Time) ApiGetDeliverabilityStatisticsRequest {
	r.fromDate = &fromDate
	return r
}

// Filter statistics before this date (RFC3339 format)
func (r ApiGetDeliverabilityStatisticsRequest) ToDate(toDate time.Time) ApiGetDeliverabilityStatisticsRequest {
	r.toDate = &toDate
	return r
}

// Filter by sender domain
func (r ApiGetDeliverabilityStatisticsRequest) SenderDomain(senderDomain string) ApiGetDeliverabilityStatisticsRequest {
	r.senderDomain = &senderDomain
	return r
}

// Filter by recipient domain
func (r ApiGetDeliverabilityStatisticsRequest) RecipientDomain(recipientDomain string) ApiGetDeliverabilityStatisticsRequest {
	r.recipientDomain = &recipientDomain
	return r
}

// Filter by tags (comma-separated)
func (r ApiGetDeliverabilityStatisticsRequest) Tags(tags string) ApiGetDeliverabilityStatisticsRequest {
	r.tags = &tags
	return r
}

// Group by time period
func (r ApiGetDeliverabilityStatisticsRequest) GroupBy(groupBy string) ApiGetDeliverabilityStatisticsRequest {
	r.groupBy = &groupBy
	return r
}

func (r ApiGetDeliverabilityStatisticsRequest) Execute() (*DeliverabilityStatisticsResponse, *http.Response, error) {
	return r.ApiService.GetDeliverabilityStatisticsExecute(r)
}

/*
GetDeliverabilityStatistics Get Deliverability Statistics

# Returns transactional deliverability statistics grouped by status

**Query Parameters:**
- `from_date`: Filter statistics after this date (RFC3339 format)
- `to_date`: Filter statistics before this date (RFC3339 format)
- `sender_domain`: Filter by sender domain
- `recipient_domain`: Filter by recipient domain
- `tags`: Filter by tags (comma-separated)
- `group_by`: Group by time period (hour, day, week, month)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@return ApiGetDeliverabilityStatisticsRequest
*/
func (a *StatisticsAPIService) GetDeliverabilityStatistics(ctx context.Context, accountId uuid.UUID) ApiGetDeliverabilityStatisticsRequest {
	return ApiGetDeliverabilityStatisticsRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
	}
}

// Execute executes the request
//
//	@return DeliverabilityStatisticsResponse
func (a *StatisticsAPIService) GetDeliverabilityStatisticsExecute(r ApiGetDeliverabilityStatisticsRequest) (*DeliverabilityStatisticsResponse, *http.Response, error) {
	var (
		method = http.MethodGet
		body   interface{}

		returnValue *DeliverabilityStatisticsResponse
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatisticsAPIService.GetDeliverabilityStatistics")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/statistics/transactional/deliverability"
	path = strings.ReplaceAll(path, "{account_id}", url.PathEscape(parameterValueToString(r.accountId, "accountId")))

	headers := make(map[string]string)
	params := url.Values{}
	formParams := url.Values{}

	if r.fromDate != nil {
		parameterAddToHeaderOrQuery(params, "from_date", r.fromDate, "form", "")
	}
	if r.toDate != nil {
		parameterAddToHeaderOrQuery(params, "to_date", r.toDate, "form", "")
	}
	if r.senderDomain != nil {
		parameterAddToHeaderOrQuery(params, "sender_domain", r.senderDomain, "form", "")
	}
	if r.recipientDomain != nil {
		parameterAddToHeaderOrQuery(params, "recipient_domain", r.recipientDomain, "form", "")
	}
	if r.tags != nil {
		parameterAddToHeaderOrQuery(params, "tags", r.tags, "form", "")
	}
	if r.groupBy != nil {
		parameterAddToHeaderOrQuery(params, "group_by", r.groupBy, "form", "")
	} else {
		defaultValue := "day"
		r.groupBy = &defaultValue
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
		if response.StatusCode == 429 {
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

type ApiGetDeliveryTimeStatisticsRequest struct {
	ctx          context.Context
	ApiService   *StatisticsAPIService
	accountId    uuid.UUID
	fromDate     *time.Time
	toDate       *time.Time
	senderDomain *string
	tags         *string
	groupBy      *string
}

// Filter statistics after this date (RFC3339 format)
func (r ApiGetDeliveryTimeStatisticsRequest) FromDate(fromDate time.Time) ApiGetDeliveryTimeStatisticsRequest {
	r.fromDate = &fromDate
	return r
}

// Filter statistics before this date (RFC3339 format)
func (r ApiGetDeliveryTimeStatisticsRequest) ToDate(toDate time.Time) ApiGetDeliveryTimeStatisticsRequest {
	r.toDate = &toDate
	return r
}

// Filter by sender domain
func (r ApiGetDeliveryTimeStatisticsRequest) SenderDomain(senderDomain string) ApiGetDeliveryTimeStatisticsRequest {
	r.senderDomain = &senderDomain
	return r
}

// Filter by tags (comma-separated)
func (r ApiGetDeliveryTimeStatisticsRequest) Tags(tags string) ApiGetDeliveryTimeStatisticsRequest {
	r.tags = &tags
	return r
}

// Group by time period
func (r ApiGetDeliveryTimeStatisticsRequest) GroupBy(groupBy string) ApiGetDeliveryTimeStatisticsRequest {
	r.groupBy = &groupBy
	return r
}

func (r ApiGetDeliveryTimeStatisticsRequest) Execute() (*DeliveryTimeStatisticsResponse, *http.Response, error) {
	return r.ApiService.GetDeliveryTimeStatisticsExecute(r)
}

/*
GetDeliveryTimeStatistics Get Delivery Time Statistics

# Returns transactional delivery time statistics grouped by recipient domain

**Query Parameters:**
- `from_date`: Filter statistics after this date (RFC3339 format)
- `to_date`: Filter statistics before this date (RFC3339 format)
- `sender_domain`: Filter by sender domain
- `tags`: Filter by tags (comma-separated)
- `group_by`: Group by time period (hour, day, week, month)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@return ApiGetDeliveryTimeStatisticsRequest
*/
func (a *StatisticsAPIService) GetDeliveryTimeStatistics(ctx context.Context, accountId uuid.UUID) ApiGetDeliveryTimeStatisticsRequest {
	return ApiGetDeliveryTimeStatisticsRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
	}
}

// Execute executes the request
//
//	@return DeliveryTimeStatisticsResponse
func (a *StatisticsAPIService) GetDeliveryTimeStatisticsExecute(r ApiGetDeliveryTimeStatisticsRequest) (*DeliveryTimeStatisticsResponse, *http.Response, error) {
	var (
		method = http.MethodGet
		body   interface{}

		returnValue *DeliveryTimeStatisticsResponse
	)

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatisticsAPIService.GetDeliveryTimeStatistics")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/v2/accounts/{account_id}/statistics/transactional/delivery-time"
	path = strings.ReplaceAll(path, "{account_id}", url.PathEscape(parameterValueToString(r.accountId, "accountId")))

	headers := make(map[string]string)
	params := url.Values{}
	formParams := url.Values{}

	if r.fromDate != nil {
		parameterAddToHeaderOrQuery(params, "from_date", r.fromDate, "form", "")
	}
	if r.toDate != nil {
		parameterAddToHeaderOrQuery(params, "to_date", r.toDate, "form", "")
	}
	if r.senderDomain != nil {
		parameterAddToHeaderOrQuery(params, "sender_domain", r.senderDomain, "form", "")
	}
	if r.tags != nil {
		parameterAddToHeaderOrQuery(params, "tags", r.tags, "form", "")
	}
	if r.groupBy != nil {
		parameterAddToHeaderOrQuery(params, "group_by", r.groupBy, "form", "")
	} else {
		defaultValue := "day"
		r.groupBy = &defaultValue
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
		if response.StatusCode == 429 {
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
