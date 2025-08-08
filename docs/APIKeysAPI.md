# \APIKeysAPI

All URIs are relative to *https://api.ahasend.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAPIKey**](APIKeysAPI.md#CreateAPIKey) | **Post** /v2/accounts/{account_id}/api-keys | Create API Key
[**DeleteAPIKey**](APIKeysAPI.md#DeleteAPIKey) | **Delete** /v2/accounts/{account_id}/api-keys/{key_id} | Delete API Key
[**GetAPIKey**](APIKeysAPI.md#GetAPIKey) | **Get** /v2/accounts/{account_id}/api-keys/{key_id} | Get API Key
[**GetAPIKeys**](APIKeysAPI.md#GetAPIKeys) | **Get** /v2/accounts/{account_id}/api-keys | Get API Keys
[**UpdateAPIKey**](APIKeysAPI.md#UpdateAPIKey) | **Put** /v2/accounts/{account_id}/api-keys/{key_id} | Update API Key



## CreateAPIKey

> ModelAPIKey CreateAPIKey(ctx, accountId).CreateAPIKeyRequest(createAPIKeyRequest).IdempotencyKey(idempotencyKey).Execute()

Create API Key



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
	createAPIKeyRequest := *openapiclient.NewCreateAPIKeyRequest("Label_example", []string{"Scopes_example"}) // CreateAPIKeyRequest | 
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | Account ID
	idempotencyKey := "user-12345-create-domain-20240101" // string | Optional idempotency key for safe request retries. Must be a unique string for each logical request. Requests with the same key will return the same response. Keys expire after 24 hours.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIKeysAPI.CreateAPIKey(context.Background(), accountId).CreateAPIKeyRequest(createAPIKeyRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.CreateAPIKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAPIKey`: ModelAPIKey
	fmt.Fprintf(os.Stdout, "Response from `APIKeysAPI.CreateAPIKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAPIKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAPIKeyRequest** | [**CreateAPIKeyRequest**](CreateAPIKeyRequest.md) |  | 

 **idempotencyKey** | **string** | Optional idempotency key for safe request retries. Must be a unique string for each logical request. Requests with the same key will return the same response. Keys expire after 24 hours.  | 

### Return type

[**ModelAPIKey**](ModelAPIKey.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAPIKey

> SuccessResponse DeleteAPIKey(ctx, accountId, keyId).Execute()

Delete API Key



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
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | API Key ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIKeysAPI.DeleteAPIKey(context.Background(), accountId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.DeleteAPIKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAPIKey`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `APIKeysAPI.DeleteAPIKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 
**keyId** | **uuid.UUID** | API Key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAPIKeyRequest struct via the builder pattern


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


## GetAPIKey

> ModelAPIKey GetAPIKey(ctx, accountId, keyId).Execute()

Get API Key



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
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | API Key ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIKeysAPI.GetAPIKey(context.Background(), accountId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.GetAPIKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAPIKey`: ModelAPIKey
	fmt.Fprintf(os.Stdout, "Response from `APIKeysAPI.GetAPIKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 
**keyId** | **uuid.UUID** | API Key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAPIKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ModelAPIKey**](ModelAPIKey.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAPIKeys

> PaginatedAPIKeysResponse GetAPIKeys(ctx, accountId).Limit(limit).Cursor(cursor).Execute()

Get API Keys



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
	limit := int32(56) // int32 | Maximum number of items to return (optional) (default to 100)
	cursor := "cursor_example" // string | Pagination cursor for the next page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIKeysAPI.GetAPIKeys(context.Background(), accountId).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.GetAPIKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAPIKeys`: PaginatedAPIKeysResponse
	fmt.Fprintf(os.Stdout, "Response from `APIKeysAPI.GetAPIKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAPIKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of items to return | [default to 100]
 **cursor** | **string** | Pagination cursor for the next page | 

### Return type

[**PaginatedAPIKeysResponse**](PaginatedAPIKeysResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAPIKey

> ModelAPIKey UpdateAPIKey(ctx, accountId, keyId).UpdateAPIKeyRequest(updateAPIKeyRequest).Execute()

Update API Key



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
	updateAPIKeyRequest := *openapiclient.NewUpdateAPIKeyRequest() // UpdateAPIKeyRequest | 
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | Account ID
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | API Key ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIKeysAPI.UpdateAPIKey(context.Background(), accountId, keyId).UpdateAPIKeyRequest(updateAPIKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.UpdateAPIKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAPIKey`: ModelAPIKey
	fmt.Fprintf(os.Stdout, "Response from `APIKeysAPI.UpdateAPIKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 
**keyId** | **uuid.UUID** | API Key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAPIKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAPIKeyRequest** | [**UpdateAPIKeyRequest**](UpdateAPIKeyRequest.md) |  | 



### Return type

[**ModelAPIKey**](ModelAPIKey.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

