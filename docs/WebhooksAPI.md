# \WebhooksAPI

All URIs are relative to *https://api.ahasend.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhook**](WebhooksAPI.md#CreateWebhook) | **Post** /v2/accounts/{account_id}/webhooks | Create Webhook
[**DeleteWebhook**](WebhooksAPI.md#DeleteWebhook) | **Delete** /v2/accounts/{account_id}/webhooks/{webhook_id} | Delete Webhook
[**GetWebhook**](WebhooksAPI.md#GetWebhook) | **Get** /v2/accounts/{account_id}/webhooks/{webhook_id} | Get Webhook
[**GetWebhooks**](WebhooksAPI.md#GetWebhooks) | **Get** /v2/accounts/{account_id}/webhooks | Get Webhooks
[**UpdateWebhook**](WebhooksAPI.md#UpdateWebhook) | **Put** /v2/accounts/{account_id}/webhooks/{webhook_id} | Update Webhook



## CreateWebhook

> Webhook CreateWebhook(ctx, accountId).CreateWebhookRequest(createWebhookRequest).IdempotencyKey(idempotencyKey).Execute()

Create Webhook



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
	createWebhookRequest := *openapiclient.NewCreateWebhookRequest("Name_example", "Url_example") // CreateWebhookRequest | 
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | Account ID
	idempotencyKey := "user-12345-create-domain-20240101" // string | Optional idempotency key for safe request retries. Must be a unique string for each logical request. Requests with the same key will return the same response. Keys expire after 24 hours.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.CreateWebhook(context.Background(), accountId).CreateWebhookRequest(createWebhookRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.CreateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebhook`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.CreateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWebhookRequest** | [**CreateWebhookRequest**](CreateWebhookRequest.md) |  | 

 **idempotencyKey** | **string** | Optional idempotency key for safe request retries. Must be a unique string for each logical request. Requests with the same key will return the same response. Keys expire after 24 hours.  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhook

> SuccessResponse DeleteWebhook(ctx, accountId, webhookId).Execute()

Delete Webhook



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
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | Webhook ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.DeleteWebhook(context.Background(), accountId, webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.DeleteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWebhook`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.DeleteWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 
**webhookId** | **uuid.UUID** | Webhook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookRequest struct via the builder pattern


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


## GetWebhook

> Webhook GetWebhook(ctx, accountId, webhookId).Execute()

Get Webhook



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
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | Webhook ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.GetWebhook(context.Background(), accountId, webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhook`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 
**webhookId** | **uuid.UUID** | Webhook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Webhook**](Webhook.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooks

> PaginatedWebhooksResponse GetWebhooks(ctx, accountId).Enabled(enabled).OnReception(onReception).OnDelivered(onDelivered).OnTransientError(onTransientError).OnFailed(onFailed).OnBounced(onBounced).OnSuppressed(onSuppressed).OnOpened(onOpened).OnClicked(onClicked).OnSuppressionCreated(onSuppressionCreated).OnDnsError(onDnsError).Limit(limit).Cursor(cursor).Execute()

Get Webhooks



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
	enabled := true // bool | Filter by enabled status (optional)
	onReception := true // bool | Filter by reception event trigger (optional)
	onDelivered := true // bool | Filter by delivery event trigger (optional)
	onTransientError := true // bool | Filter by transient error event trigger (optional)
	onFailed := true // bool | Filter by failure event trigger (optional)
	onBounced := true // bool | Filter by bounce event trigger (optional)
	onSuppressed := true // bool | Filter by suppression event trigger (optional)
	onOpened := true // bool | Filter by open event trigger (optional)
	onClicked := true // bool | Filter by click event trigger (optional)
	onSuppressionCreated := true // bool | Filter by suppression creation event trigger (optional)
	onDnsError := true // bool | Filter by DNS error event trigger (optional)
	limit := int32(56) // int32 | Maximum number of items to return (1-100) (optional) (default to 100)
	cursor := "cursor_example" // string | Pagination cursor for the next page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.GetWebhooks(context.Background(), accountId).Enabled(enabled).OnReception(onReception).OnDelivered(onDelivered).OnTransientError(onTransientError).OnFailed(onFailed).OnBounced(onBounced).OnSuppressed(onSuppressed).OnOpened(onOpened).OnClicked(onClicked).OnSuppressionCreated(onSuppressionCreated).OnDnsError(onDnsError).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhooks`: PaginatedWebhooksResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enabled** | **bool** | Filter by enabled status | 
 **onReception** | **bool** | Filter by reception event trigger | 
 **onDelivered** | **bool** | Filter by delivery event trigger | 
 **onTransientError** | **bool** | Filter by transient error event trigger | 
 **onFailed** | **bool** | Filter by failure event trigger | 
 **onBounced** | **bool** | Filter by bounce event trigger | 
 **onSuppressed** | **bool** | Filter by suppression event trigger | 
 **onOpened** | **bool** | Filter by open event trigger | 
 **onClicked** | **bool** | Filter by click event trigger | 
 **onSuppressionCreated** | **bool** | Filter by suppression creation event trigger | 
 **onDnsError** | **bool** | Filter by DNS error event trigger | 
 **limit** | **int32** | Maximum number of items to return (1-100) | [default to 100]
 **cursor** | **string** | Pagination cursor for the next page | 

### Return type

[**PaginatedWebhooksResponse**](PaginatedWebhooksResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhook

> Webhook UpdateWebhook(ctx, accountId, webhookId).UpdateWebhookRequest(updateWebhookRequest).Execute()

Update Webhook



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
	updateWebhookRequest := *openapiclient.NewUpdateWebhookRequest() // UpdateWebhookRequest | 
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | Account ID
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | Webhook ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.UpdateWebhook(context.Background(), accountId, webhookId).UpdateWebhookRequest(updateWebhookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.UpdateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhook`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.UpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 
**webhookId** | **uuid.UUID** | Webhook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateWebhookRequest** | [**UpdateWebhookRequest**](UpdateWebhookRequest.md) |  | 



### Return type

[**Webhook**](Webhook.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

