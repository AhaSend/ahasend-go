# \MessagesAPI

All URIs are relative to *https://api.ahasend.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelMessage**](MessagesAPI.md#CancelMessage) | **Delete** /v2/accounts/{account_id}/messages/{message_id}/cancel | Cancel Message
[**CreateMessage**](MessagesAPI.md#CreateMessage) | **Post** /v2/accounts/{account_id}/messages | Create Message
[**GetMessages**](MessagesAPI.md#GetMessages) | **Get** /v2/accounts/{account_id}/messages | Get Messages



## CancelMessage

> SuccessResponse CancelMessage(ctx, accountId, messageId).Execute()

Cancel Message



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
	messageId := "messageId_example" // string | Message ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.CancelMessage(context.Background(), accountId, messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.CancelMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelMessage`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.CancelMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 
**messageId** | **string** | Message ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelMessageRequest struct via the builder pattern


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


## CreateMessage

> CreateMessageResponse CreateMessage(ctx, accountId).CreateMessageRequest(createMessageRequest).IdempotencyKey(idempotencyKey).Execute()

Create Message



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
	createMessageRequest := *openapiclient.NewCreateMessageRequest(*openapiclient.NewSenderAddress("Email_example"), []openapiclient.Recipient{*openapiclient.NewRecipient("Email_example")}, "Subject_example") // CreateMessageRequest | 
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | Account ID
	idempotencyKey := "user-12345-create-domain-20240101" // string | Optional idempotency key for safe request retries. Must be a unique string for each logical request. Requests with the same key will return the same response. Keys expire after 24 hours.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.CreateMessage(context.Background(), accountId).CreateMessageRequest(createMessageRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.CreateMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMessage`: CreateMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.CreateMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMessageRequest** | [**CreateMessageRequest**](CreateMessageRequest.md) |  | 

 **idempotencyKey** | **string** | Optional idempotency key for safe request retries. Must be a unique string for each logical request. Requests with the same key will return the same response. Keys expire after 24 hours.  | 

### Return type

[**CreateMessageResponse**](CreateMessageResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessages

> PaginatedMessagesResponse GetMessages(ctx, accountId).Sender(sender).Recipient(recipient).Subject(subject).MessageIdHeader(messageIdHeader).FromTime(fromTime).ToTime(toTime).Limit(limit).Cursor(cursor).Execute()

Get Messages



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
	sender := "sender_example" // string | Sender email address (must be from domain in API key scopes)
	recipient := "recipient_example" // string | Recipient email address (optional)
	subject := "subject_example" // string | Filter by subject text (optional)
	messageIdHeader := "messageIdHeader_example" // string | Filter by message ID header (same ID returned by CreateMessage API) (optional)
	fromTime := time.Now() // time.Time | Filter messages created after this time (RFC3339 format) (optional)
	toTime := time.Now() // time.Time | Filter messages created before this time (RFC3339 format) (optional)
	limit := int32(56) // int32 | Maximum number of items to return (1-100) (optional) (default to 100)
	cursor := "cursor_example" // string | Pagination cursor for the next page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.GetMessages(context.Background(), accountId).Sender(sender).Recipient(recipient).Subject(subject).MessageIdHeader(messageIdHeader).FromTime(fromTime).ToTime(toTime).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.GetMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessages`: PaginatedMessagesResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.GetMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sender** | **string** | Sender email address (must be from domain in API key scopes) | 
 **recipient** | **string** | Recipient email address | 
 **subject** | **string** | Filter by subject text | 
 **messageIdHeader** | **string** | Filter by message ID header (same ID returned by CreateMessage API) | 
 **fromTime** | **time.Time** | Filter messages created after this time (RFC3339 format) | 
 **toTime** | **time.Time** | Filter messages created before this time (RFC3339 format) | 
 **limit** | **int32** | Maximum number of items to return (1-100) | [default to 100]
 **cursor** | **string** | Pagination cursor for the next page | 

### Return type

[**PaginatedMessagesResponse**](PaginatedMessagesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

