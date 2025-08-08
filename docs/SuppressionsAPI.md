# \SuppressionsAPI

All URIs are relative to *https://api.ahasend.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSuppression**](SuppressionsAPI.md#CreateSuppression) | **Post** /v2/accounts/{account_id}/suppressions | Create Suppression
[**DeleteAllSuppressions**](SuppressionsAPI.md#DeleteAllSuppressions) | **Delete** /v2/accounts/{account_id}/suppressions | Delete All Suppressions
[**DeleteSuppression**](SuppressionsAPI.md#DeleteSuppression) | **Delete** /v2/accounts/{account_id}/suppressions/{email} | Delete Suppression
[**GetSuppressions**](SuppressionsAPI.md#GetSuppressions) | **Get** /v2/accounts/{account_id}/suppressions | Get Suppressions



## CreateSuppression

> Suppression CreateSuppression(ctx, accountId).CreateSuppressionRequest(createSuppressionRequest).IdempotencyKey(idempotencyKey).Execute()

Create Suppression



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/AhaSend/ahasend-go"
)

func main() {
	createSuppressionRequest := *openapiclient.NewCreateSuppressionRequest("Email_example", time.Now()) // CreateSuppressionRequest | 
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | Account ID
	idempotencyKey := "user-12345-create-domain-20240101" // string | Optional idempotency key for safe request retries. Must be a unique string for each logical request. Requests with the same key will return the same response. Keys expire after 24 hours.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppressionsAPI.CreateSuppression(context.Background(), accountId).CreateSuppressionRequest(createSuppressionRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppressionsAPI.CreateSuppression``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSuppression`: Suppression
	fmt.Fprintf(os.Stdout, "Response from `SuppressionsAPI.CreateSuppression`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSuppressionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSuppressionRequest** | [**CreateSuppressionRequest**](CreateSuppressionRequest.md) |  | 

 **idempotencyKey** | **string** | Optional idempotency key for safe request retries. Must be a unique string for each logical request. Requests with the same key will return the same response. Keys expire after 24 hours.  | 

### Return type

[**Suppression**](Suppression.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllSuppressions

> SuccessResponse DeleteAllSuppressions(ctx, accountId).Domain(domain).Execute()

Delete All Suppressions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/AhaSend/ahasend-go"
)

func main() {
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | Account ID
	domain := "domain_example" // string | Optional domain filter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppressionsAPI.DeleteAllSuppressions(context.Background(), accountId).Domain(domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppressionsAPI.DeleteAllSuppressions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAllSuppressions`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `SuppressionsAPI.DeleteAllSuppressions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllSuppressionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domain** | **string** | Optional domain filter | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSuppression

> SuccessResponse DeleteSuppression(ctx, accountId, email).Domain(domain).Execute()

Delete Suppression



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/AhaSend/ahasend-go"
)

func main() {
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | Account ID
	email := "email_example" // string | Email address
	domain := "domain_example" // string | Optional domain filter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppressionsAPI.DeleteSuppression(context.Background(), accountId, email).Domain(domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppressionsAPI.DeleteSuppression``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSuppression`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `SuppressionsAPI.DeleteSuppression`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 
**email** | **string** | Email address | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSuppressionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **domain** | **string** | Optional domain filter | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSuppressions

> PaginatedSuppressionsResponse GetSuppressions(ctx, accountId).Domain(domain).FromDate(fromDate).ToDate(toDate).Limit(limit).Cursor(cursor).Execute()

Get Suppressions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/AhaSend/ahasend-go"
)

func main() {
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | Account ID
	domain := "domain_example" // string | Filter by domain (optional)
	fromDate := time.Now() // time.Time | Filter suppressions created after this date (RFC3339 format) (optional)
	toDate := time.Now() // time.Time | Filter suppressions created before this date (RFC3339 format) (optional)
	limit := int32(56) // int32 | Maximum number of items to return (1-100) (optional) (default to 100)
	cursor := "cursor_example" // string | Pagination cursor for the next page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppressionsAPI.GetSuppressions(context.Background(), accountId).Domain(domain).FromDate(fromDate).ToDate(toDate).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppressionsAPI.GetSuppressions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSuppressions`: PaginatedSuppressionsResponse
	fmt.Fprintf(os.Stdout, "Response from `SuppressionsAPI.GetSuppressions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSuppressionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domain** | **string** | Filter by domain | 
 **fromDate** | **time.Time** | Filter suppressions created after this date (RFC3339 format) | 
 **toDate** | **time.Time** | Filter suppressions created before this date (RFC3339 format) | 
 **limit** | **int32** | Maximum number of items to return (1-100) | [default to 100]
 **cursor** | **string** | Pagination cursor for the next page | 

### Return type

[**PaginatedSuppressionsResponse**](PaginatedSuppressionsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

