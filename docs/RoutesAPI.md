# \RoutesAPI

All URIs are relative to *https://api.ahasend.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoute**](RoutesAPI.md#CreateRoute) | **Post** /v2/accounts/{account_id}/routes | Create Route
[**DeleteRoute**](RoutesAPI.md#DeleteRoute) | **Delete** /v2/accounts/{account_id}/routes/{route_id} | Delete Route
[**GetRoute**](RoutesAPI.md#GetRoute) | **Get** /v2/accounts/{account_id}/routes/{route_id} | Get Route
[**GetRoutes**](RoutesAPI.md#GetRoutes) | **Get** /v2/accounts/{account_id}/routes | Get Routes
[**UpdateRoute**](RoutesAPI.md#UpdateRoute) | **Put** /v2/accounts/{account_id}/routes/{route_id} | Update Route



## CreateRoute

> Route CreateRoute(ctx, accountId).CreateRouteRequest(createRouteRequest).IdempotencyKey(idempotencyKey).Execute()

Create Route



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
	createRouteRequest := *openapiclient.NewCreateRouteRequest("Name_example", "Url_example") // CreateRouteRequest | 
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | Account ID
	idempotencyKey := "user-12345-create-domain-20240101" // string | Optional idempotency key for safe request retries. Must be a unique string for each logical request. Requests with the same key will return the same response. Keys expire after 24 hours.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.CreateRoute(context.Background(), accountId).CreateRouteRequest(createRouteRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.CreateRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoute`: Route
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.CreateRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRouteRequest** | [**CreateRouteRequest**](CreateRouteRequest.md) |  | 

 **idempotencyKey** | **string** | Optional idempotency key for safe request retries. Must be a unique string for each logical request. Requests with the same key will return the same response. Keys expire after 24 hours.  | 

### Return type

[**Route**](Route.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoute

> SuccessResponse DeleteRoute(ctx, accountId, routeId).Execute()

Delete Route



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
	routeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | Route ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.DeleteRoute(context.Background(), accountId, routeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.DeleteRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRoute`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.DeleteRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 
**routeId** | **uuid.UUID** | Route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetRoute

> Route GetRoute(ctx, accountId, routeId).Execute()

Get Route



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
	routeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | Route ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.GetRoute(context.Background(), accountId, routeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.GetRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoute`: Route
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.GetRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 
**routeId** | **uuid.UUID** | Route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Route**](Route.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoutes

> PaginatedRoutesResponse GetRoutes(ctx, accountId).Limit(limit).Cursor(cursor).Execute()

Get Routes



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
	limit := int32(56) // int32 | Maximum number of items to return (1-100) (optional) (default to 100)
	cursor := "cursor_example" // string | Pagination cursor for the next page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.GetRoutes(context.Background(), accountId).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.GetRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoutes`: PaginatedRoutesResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.GetRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of items to return (1-100) | [default to 100]
 **cursor** | **string** | Pagination cursor for the next page | 

### Return type

[**PaginatedRoutesResponse**](PaginatedRoutesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoute

> Route UpdateRoute(ctx, accountId, routeId).UpdateRouteRequest(updateRouteRequest).Execute()

Update Route



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
	updateRouteRequest := *openapiclient.NewUpdateRouteRequest() // UpdateRouteRequest | 
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | Account ID
	routeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | Route ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.UpdateRoute(context.Background(), accountId, routeId).UpdateRouteRequest(updateRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.UpdateRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoute`: Route
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.UpdateRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 
**routeId** | **uuid.UUID** | Route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRouteRequest** | [**UpdateRouteRequest**](UpdateRouteRequest.md) |  | 



### Return type

[**Route**](Route.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

