# \SMTPCredentialsAPI

All URIs are relative to *https://api.ahasend.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSMTPCredential**](SMTPCredentialsAPI.md#CreateSMTPCredential) | **Post** /v2/accounts/{account_id}/smtp-credentials | Create SMTP Credential
[**DeleteSMTPCredential**](SMTPCredentialsAPI.md#DeleteSMTPCredential) | **Delete** /v2/accounts/{account_id}/smtp-credentials/{smtp_credential_id} | Delete SMTP Credential
[**GetSMTPCredential**](SMTPCredentialsAPI.md#GetSMTPCredential) | **Get** /v2/accounts/{account_id}/smtp-credentials/{smtp_credential_id} | Get SMTP Credential
[**GetSMTPCredentials**](SMTPCredentialsAPI.md#GetSMTPCredentials) | **Get** /v2/accounts/{account_id}/smtp-credentials | Get SMTP Credentials



## CreateSMTPCredential

> SMTPCredential CreateSMTPCredential(ctx, accountId).CreateSMTPCredentialRequest(createSMTPCredentialRequest).IdempotencyKey(idempotencyKey).Execute()

Create SMTP Credential



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
	createSMTPCredentialRequest := *openapiclient.NewCreateSMTPCredentialRequest("Name_example", "Username_example", "Password_example", "Scope_example") // CreateSMTPCredentialRequest | 
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | Account ID
	idempotencyKey := "user-12345-create-domain-20240101" // string | Optional idempotency key for safe request retries. Must be a unique string for each logical request. Requests with the same key will return the same response. Keys expire after 24 hours.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMTPCredentialsAPI.CreateSMTPCredential(context.Background(), accountId).CreateSMTPCredentialRequest(createSMTPCredentialRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMTPCredentialsAPI.CreateSMTPCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSMTPCredential`: SMTPCredential
	fmt.Fprintf(os.Stdout, "Response from `SMTPCredentialsAPI.CreateSMTPCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSMTPCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSMTPCredentialRequest** | [**CreateSMTPCredentialRequest**](CreateSMTPCredentialRequest.md) |  | 

 **idempotencyKey** | **string** | Optional idempotency key for safe request retries. Must be a unique string for each logical request. Requests with the same key will return the same response. Keys expire after 24 hours.  | 

### Return type

[**SMTPCredential**](SMTPCredential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSMTPCredential

> SuccessResponse DeleteSMTPCredential(ctx, accountId, smtpCredentialId).Execute()

Delete SMTP Credential



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
	smtpCredentialId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | SMTP Credential ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMTPCredentialsAPI.DeleteSMTPCredential(context.Background(), accountId, smtpCredentialId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMTPCredentialsAPI.DeleteSMTPCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSMTPCredential`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `SMTPCredentialsAPI.DeleteSMTPCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 
**smtpCredentialId** | **uuid.UUID** | SMTP Credential ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSMTPCredentialRequest struct via the builder pattern


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


## GetSMTPCredential

> SMTPCredential GetSMTPCredential(ctx, accountId, smtpCredentialId).Execute()

Get SMTP Credential



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
	smtpCredentialId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | SMTP Credential ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SMTPCredentialsAPI.GetSMTPCredential(context.Background(), accountId, smtpCredentialId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMTPCredentialsAPI.GetSMTPCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSMTPCredential`: SMTPCredential
	fmt.Fprintf(os.Stdout, "Response from `SMTPCredentialsAPI.GetSMTPCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 
**smtpCredentialId** | **uuid.UUID** | SMTP Credential ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSMTPCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SMTPCredential**](SMTPCredential.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSMTPCredentials

> PaginatedSMTPCredentialsResponse GetSMTPCredentials(ctx, accountId).Limit(limit).Cursor(cursor).Execute()

Get SMTP Credentials



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
	resp, r, err := apiClient.SMTPCredentialsAPI.GetSMTPCredentials(context.Background(), accountId).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SMTPCredentialsAPI.GetSMTPCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSMTPCredentials`: PaginatedSMTPCredentialsResponse
	fmt.Fprintf(os.Stdout, "Response from `SMTPCredentialsAPI.GetSMTPCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSMTPCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of items to return (1-100) | [default to 100]
 **cursor** | **string** | Pagination cursor for the next page | 

### Return type

[**PaginatedSMTPCredentialsResponse**](PaginatedSMTPCredentialsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

