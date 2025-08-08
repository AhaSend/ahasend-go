// Package ahasend provides a Go SDK for the AhaSend transactional email API.
//
// AhaSend enables developers to send transactional emails with high deliverability,
// comprehensive tracking, and powerful management features.
//
// Key Features:
//   - Send transactional emails with attachments and scheduling
//   - Comprehensive domain, webhook, and route management
//   - Built-in rate limiting and automatic retry with exponential backoff
//   - Statistics and analytics for email performance
//   - Support for suppression lists and SMTP credentials
//
// Authentication:
// All API requests require a Bearer token in the Authorization header:
//   Authorization: Bearer aha-sk-64-CHARACTER-RANDOM-STRING
//
// Basic Usage:
//   import "github.com/AhaSend/ahasend-go"
//
//   cfg := ahasend.NewConfiguration()
//   client := ahasend.NewAPIClient(cfg)
//
//   ctx := context.WithValue(context.Background(),
//     ahasend.ContextAccessToken, "your-api-key")
//
//   // Send a message
//   messageReq := &ahasend.CreateMessageRequest{
//     From:    "sender@yourdomain.com",
//     To:      []string{"recipient@example.com"},
//     Subject: "Test Email",
//     Html:    "<p>Hello from AhaSend!</p>",
//   }
//
//   response, _, err := client.MessagesAPI.
//     CreateMessage(ctx, accountID).
//     CreateMessageRequest(*messageReq).
//     Execute()
//
// For more information visit: https://ahasend.com/docs

package ahasend

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	JsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?json)`)
)

// APIClient manages communication with the AhaSend API v2 API v2.0.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg               *Configuration
	common            service // Reuse a single struct instead of allocating one for each service on the heap.
	rateLimiter       *RateLimiter
	idempotencyHelper *IdempotencyHelper

	// API Services

	APIKeysAPI *APIKeysAPIService

	AccountsAPI *AccountsAPIService

	DomainsAPI *DomainsAPIService

	MessagesAPI *MessagesAPIService

	RoutesAPI *RoutesAPIService

	SMTPCredentialsAPI *SMTPCredentialsAPIService

	StatisticsAPI *StatisticsAPIService

	SuppressionsAPI *SuppressionsAPIService

	UtilityAPI *UtilityAPIService

	WebhooksAPI *WebhooksAPIService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// Initialize rate limiter
	c.rateLimiter = NewRateLimiter()
	c.rateLimiter.SetGlobalEnabled(cfg.EnableRateLimit)

	// Apply custom rate limits if provided
	if cfg.CustomerRateLimits != nil {
		c.rateLimiter.ConfigureFromCustomerConfig(*cfg.CustomerRateLimits)
	}

	// Initialize idempotency helper
	c.idempotencyHelper = NewIdempotencyHelper(cfg.IdempotencyConfig)

	// API Services
	c.APIKeysAPI = (*APIKeysAPIService)(&c.common)
	c.AccountsAPI = (*AccountsAPIService)(&c.common)
	c.DomainsAPI = (*DomainsAPIService)(&c.common)
	c.MessagesAPI = (*MessagesAPIService)(&c.common)
	c.RoutesAPI = (*RoutesAPIService)(&c.common)
	c.SMTPCredentialsAPI = (*SMTPCredentialsAPIService)(&c.common)
	c.StatisticsAPI = (*StatisticsAPIService)(&c.common)
	c.SuppressionsAPI = (*SuppressionsAPIService)(&c.common)
	c.UtilityAPI = (*UtilityAPIService)(&c.common)
	c.WebhooksAPI = (*WebhooksAPIService)(&c.common)

	return c
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insensitive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.EqualFold(a, needle) {
			return true
		}
	}
	return false
}

func parameterValueToString(obj interface{}, key string) string {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		if actualObj, ok := obj.(interface{ GetActualInstanceValue() interface{} }); ok {
			return fmt.Sprintf("%v", actualObj.GetActualInstanceValue())
		}

		return fmt.Sprintf("%v", obj)
	}
	var param, ok = obj.(MappedNullable)
	if !ok {
		return ""
	}
	dataMap, err := param.ToMap()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%v", dataMap[key])
}

// parameterAddToHeaderOrQuery adds the provided object to the request header or url query
// supporting deep object syntax
func parameterAddToHeaderOrQuery(headerOrQueryParams interface{}, keyPrefix string, obj interface{}, style string, collectionType string) {
	var v = reflect.ValueOf(obj)
	var value = ""
	if v == reflect.ValueOf(nil) {
		value = "null"
	} else {
		switch v.Kind() {
		case reflect.Invalid:
			value = "invalid"

		case reflect.Struct:
			if t, ok := obj.(MappedNullable); ok {
				dataMap, err := t.ToMap()
				if err != nil {
					return
				}
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, dataMap, style, collectionType)
				return
			}
			if t, ok := obj.(time.Time); ok {
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, t.Format(time.RFC3339Nano), style, collectionType)
				return
			}
			value = v.Type().String() + " value"
		case reflect.Slice:
			var indValue = reflect.ValueOf(obj)
			if indValue == reflect.ValueOf(nil) {
				return
			}
			var lenIndValue = indValue.Len()
			for i := 0; i < lenIndValue; i++ {
				var arrayValue = indValue.Index(i)
				var keyPrefixForCollectionType = keyPrefix
				if style == "deepObject" {
					keyPrefixForCollectionType = keyPrefix + "[" + strconv.Itoa(i) + "]"
				}
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefixForCollectionType, arrayValue.Interface(), style, collectionType)
			}
			return

		case reflect.Map:
			var indValue = reflect.ValueOf(obj)
			if indValue == reflect.ValueOf(nil) {
				return
			}
			iter := indValue.MapRange()
			for iter.Next() {
				k, v := iter.Key(), iter.Value()
				parameterAddToHeaderOrQuery(headerOrQueryParams, fmt.Sprintf("%s[%s]", keyPrefix, k.String()), v.Interface(), style, collectionType)
			}
			return

		case reflect.Interface:
			fallthrough
		case reflect.Ptr:
			parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, v.Elem().Interface(), style, collectionType)
			return

		case reflect.Int, reflect.Int8, reflect.Int16,
			reflect.Int32, reflect.Int64:
			value = strconv.FormatInt(v.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16,
			reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			value = strconv.FormatUint(v.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			value = strconv.FormatFloat(v.Float(), 'g', -1, 32)
		case reflect.Bool:
			value = strconv.FormatBool(v.Bool())
		case reflect.String:
			value = v.String()
		default:
			value = v.Type().String() + " value"
		}
	}

	switch valuesMap := headerOrQueryParams.(type) {
	case url.Values:
		if collectionType == "csv" && valuesMap.Get(keyPrefix) != "" {
			valuesMap.Set(keyPrefix, valuesMap.Get(keyPrefix)+","+value)
		} else {
			valuesMap.Add(keyPrefix, value)
		}

	case map[string]string:
		valuesMap[keyPrefix] = value
	}
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	// Apply rate limiting before making the request, using the request's context for cancellation
	ctx := request.Context()
	if err := c.rateLimiter.WaitForTokenWithContext(ctx, request.Method, request.URL.Path); err != nil {
		return nil, err
	}

	if c.cfg.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	resp, err := c.cfg.HTTPClient.Do(request)
	if err != nil {
		// Wrap in NetworkError for better error handling
		netErr := &NetworkError{
			Op:  fmt.Sprintf("%s %s", request.Method, request.URL.Path),
			Err: err,
		}
		// Handle network errors with retry
		if c.shouldRetryOnError(err) {
			return c.handleRetryableError(request, netErr)
		}
		return nil, netErr
	}

	// Handle retryable HTTP status codes
	if c.shouldRetryOnStatus(resp.StatusCode) {
		return c.handleRetryableResponse(request, resp)
	}

	if c.cfg.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
		}
		log.Printf("\n%s\n", string(dump))
	}
	return resp, err
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *Configuration {
	return c.cfg
}

// shouldRetryOnError determines if an error is retryable
func (c *APIClient) shouldRetryOnError(err error) bool {
	if err == nil {
		return false
	}
	// Retry on network errors (connection refused, timeout, etc.)
	return true
}

// shouldRetryOnStatus determines if an HTTP status code is retryable
func (c *APIClient) shouldRetryOnStatus(statusCode int) bool {
	switch statusCode {
	case 429, // Too Many Requests
		502, // Bad Gateway
		503, // Service Unavailable
		504: // Gateway Timeout
		return true
	default:
		return false
	}
}

// handleRetryableError handles network errors with exponential backoff retry logic
func (c *APIClient) handleRetryableError(request *http.Request, originalErr error) (*http.Response, error) {
	ctx := request.Context()

	for attempt := 0; attempt < c.cfg.MaxRetries; attempt++ {
		// Exponential backoff with jitter
		backoffTime := time.Duration(1<<uint(attempt)) * time.Second
		jitter := time.Duration(rand.Intn(1000)) * time.Millisecond
		totalWait := backoffTime + jitter

		// Wait with context cancellation support
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(totalWait):
			// Continue with retry
		}

		// Clone the request for retry
		retryReq := request.Clone(ctx)

		// Make the retry request
		resp, err := c.cfg.HTTPClient.Do(retryReq)
		if err != nil {
			continue // Try again on error
		}

		// If successful or non-retryable status, return the response
		if !c.shouldRetryOnStatus(resp.StatusCode) {
			return resp, nil
		}

		resp.Body.Close() // Close retry response body before next attempt
	}

	// All retries failed, return the original error
	return nil, originalErr
}

// handleRetryableResponse handles retryable HTTP responses with exponential backoff
func (c *APIClient) handleRetryableResponse(request *http.Request, originalResp *http.Response) (*http.Response, error) {
	ctx := request.Context()
	originalStatus := originalResp.StatusCode
	originalResp.Body.Close() // Close the original response body

	for attempt := 0; attempt < c.cfg.MaxRetries; attempt++ {
		// Exponential backoff with jitter
		backoffTime := time.Duration(1<<uint(attempt)) * time.Second
		jitter := time.Duration(rand.Intn(1000)) * time.Millisecond
		totalWait := backoffTime + jitter

		// Wait with context cancellation support
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(totalWait):
			// Continue with retry
		}

		// Clone the request for retry
		retryReq := request.Clone(ctx)

		// Make the retry request
		resp, err := c.cfg.HTTPClient.Do(retryReq)
		if err != nil {
			if c.shouldRetryOnError(err) {
				continue // Try again on network error
			}
			return nil, err
		}

		// If not a retryable status anymore, return the response
		if !c.shouldRetryOnStatus(resp.StatusCode) {
			return resp, nil
		}

		resp.Body.Close() // Close retry response body before next attempt
	}

	// All retries failed, return a response with the original status
	return &http.Response{
		StatusCode: originalStatus,
		Status:     http.StatusText(originalStatus),
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader(fmt.Sprintf("Request failed after %d retries", c.cfg.MaxRetries))),
	}, nil
}

// Rate Limiting Public API Methods

// SetGeneralRateLimit sets the rate limit for general API endpoints
func (c *APIClient) SetGeneralRateLimit(requestsPerSecond, burstCapacity int) {
	c.rateLimiter.SetRateLimit(GeneralAPI, requestsPerSecond, burstCapacity)
}

// SetStatisticsRateLimit sets the rate limit for statistics API endpoints
func (c *APIClient) SetStatisticsRateLimit(requestsPerSecond, burstCapacity int) {
	c.rateLimiter.SetRateLimit(StatisticsAPI, requestsPerSecond, burstCapacity)
}

// SetSendMessageRateLimit sets the rate limit for send message API endpoint
func (c *APIClient) SetSendMessageRateLimit(requestsPerSecond, burstCapacity int) {
	c.rateLimiter.SetRateLimit(SendMessageAPI, requestsPerSecond, burstCapacity)
}

// SetCustomRateLimit sets the rate limit for a specific endpoint type
func (c *APIClient) SetCustomRateLimit(endpointType EndpointType, requestsPerSecond, burstCapacity int) {
	c.rateLimiter.SetRateLimit(endpointType, requestsPerSecond, burstCapacity)
}

// EnableRateLimit enables or disables rate limiting for a specific endpoint type
func (c *APIClient) EnableRateLimit(endpointType EndpointType, enabled bool) {
	c.rateLimiter.EnableRateLimit(endpointType, enabled)
}

// SetGlobalRateLimit enables or disables rate limiting globally
func (c *APIClient) SetGlobalRateLimit(enabled bool) {
	c.rateLimiter.SetGlobalEnabled(enabled)
}

// GetRateLimitStatus returns the current rate limit status for a specific endpoint type
func (c *APIClient) GetRateLimitStatus(endpointType EndpointType) RateLimitStatus {
	return c.rateLimiter.GetStatus(endpointType)
}

// ConfigureCustomerRateLimits applies a complete customer rate limit configuration
// This is a convenience method for enterprise customers to set all limits at once
func (c *APIClient) ConfigureCustomerRateLimits(config CustomerRateLimitConfig) {
	c.rateLimiter.ConfigureFromCustomerConfig(config)
}

// Idempotency Public API Methods

// GenerateIdempotencyKey generates a new UUID-based idempotency key
func (c *APIClient) GenerateIdempotencyKey() string {
	return c.idempotencyHelper.GenerateKey()
}

// SetIdempotencyConfig updates the idempotency configuration
func (c *APIClient) SetIdempotencyConfig(config IdempotencyConfig) {
	c.cfg.IdempotencyConfig = config
	c.idempotencyHelper = NewIdempotencyHelper(config)
}

// GetIdempotencyConfig returns the current idempotency configuration
func (c *APIClient) GetIdempotencyConfig() IdempotencyConfig {
	return c.cfg.IdempotencyConfig
}

// NewIdempotencyKeyBuilder creates a new idempotency key builder for related operations
func (c *APIClient) NewIdempotencyKeyBuilder(baseKey ...string) *IdempotencyKeyBuilder {
	var base string
	if len(baseKey) > 0 && baseKey[0] != "" {
		base = baseKey[0]
	} else {
		base = c.GenerateIdempotencyKey()
	}
	return NewIdempotencyKeyBuilder(base)
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values) (request *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		url.Host = c.cfg.Host
	}

	// Override request scheme, if applicable
	if c.cfg.Scheme != "" {
		url.Scheme = c.cfg.Scheme
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		request, err = http.NewRequest(method, url.String(), body)
	} else {
		request, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers[h] = []string{v}
		}
		request.Header = headers
	}

	// Add the user agent to the request.
	request.Header.Add("User-Agent", c.cfg.UserAgent)

	// Auto-generate idempotency key for POST requests if not already set
	if method == http.MethodPost {
		if _, exists := headerParams["Idempotency-Key"]; !exists && c.cfg.IdempotencyConfig.AutoGenerate {
			request.Header.Add("Idempotency-Key", c.idempotencyHelper.GenerateKey())
		}
	}

	if ctx != nil {
		// add context to the request
		request = request.WithContext(ctx)

		// Walk through any authentication.

		// AccessToken Authentication
		if auth, ok := ctx.Value(ContextAccessToken).(string); ok {
			request.Header.Add("Authorization", "Bearer "+auth)
		}

	}

	for header, value := range c.cfg.DefaultHeader {
		request.Header.Add(header, value)
	}
	return request, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if JsonCheck.MatchString(contentType) {
		if actualObj, ok := v.(interface{ GetActualInstance() interface{} }); ok { // oneOf, anyOf schemas
			if unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON([]byte) error }); ok { // make sure it has UnmarshalJSON defined
				if err = unmarshalObj.UnmarshalJSON(b); err != nil {
					return err
				}
			} else {
				return errors.New("unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
			}
		} else if err = json.Unmarshal(b, v); err != nil { // simple model
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	bodyBuf = &bytes.Buffer{}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if JsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type %s", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	// AhaSend API primarily uses JSON, with plain text for simple strings
	switch reflect.TypeOf(body).Kind() {
	case reflect.Struct, reflect.Map, reflect.Ptr, reflect.Slice:
		return "application/json; charset=utf-8"
	case reflect.String:
		return "text/plain; charset=utf-8"
	default:
		return "application/json; charset=utf-8"
	}
}

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
// This type wraps the new APIError for backward compatibility.
type GenericOpenAPIError struct {
	body  []byte
	error string
	model interface{}
	// APIError provides structured error information
	APIError *APIError
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	if e.APIError != nil {
		return e.APIError.Error()
	}
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() interface{} {
	return e.model
}

// Type returns the error type for programmatic handling
func (e GenericOpenAPIError) Type() ErrorType {
	if e.APIError != nil {
		return e.APIError.Type
	}
	return ErrorTypeUnknown
}

// IsRetryable returns true if the error is retryable
func (e GenericOpenAPIError) IsRetryable() bool {
	if e.APIError != nil {
		return e.APIError.IsRetryable()
	}
	return false
}

// StatusCode returns the HTTP status code if available
func (e GenericOpenAPIError) StatusCode() int {
	if e.APIError != nil {
		return e.APIError.StatusCode
	}
	return 0
}

// format error message using title and detail when model implements rfc7807
func formatErrorMessage(status string, v interface{}) string {
	str := ""
	metaValue := reflect.ValueOf(v).Elem()

	if metaValue.Kind() == reflect.Struct {
		field := metaValue.FieldByName("Title")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s", field.Interface())
		}

		field = metaValue.FieldByName("Detail")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s (%s)", str, field.Interface())
		}
	}

	return strings.TrimSpace(fmt.Sprintf("%s %s", status, str))
}

// NewGenericOpenAPIError creates a new GenericOpenAPIError with structured error information
func NewGenericOpenAPIError(response *http.Response, body []byte, model interface{}) *GenericOpenAPIError {
	// Parse the structured API error
	apiErr := ParseAPIError(response, body)

	// Create the generic error for backward compatibility
	genErr := &GenericOpenAPIError{
		body:     body,
		error:    response.Status,
		model:    model,
		APIError: apiErr,
	}

	// If we have a model, try to extract additional information
	if model != nil {
		genErr.error = formatErrorMessage(response.Status, model)
	} else if apiErr != nil {
		genErr.error = apiErr.Error()
	}

	return genErr
}
