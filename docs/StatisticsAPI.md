# \StatisticsAPI

All URIs are relative to *https://api.ahasend.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBounceStatistics**](StatisticsAPI.md#GetBounceStatistics) | **Get** /v2/accounts/{account_id}/statistics/transactional/bounce | Get Bounce Statistics
[**GetDeliverabilityStatistics**](StatisticsAPI.md#GetDeliverabilityStatistics) | **Get** /v2/accounts/{account_id}/statistics/transactional/deliverability | Get Deliverability Statistics
[**GetDeliveryTimeStatistics**](StatisticsAPI.md#GetDeliveryTimeStatistics) | **Get** /v2/accounts/{account_id}/statistics/transactional/delivery-time | Get Delivery Time Statistics



## GetBounceStatistics

> BounceStatisticsResponse GetBounceStatistics(ctx, accountId).FromDate(fromDate).ToDate(toDate).SenderDomain(senderDomain).RecipientDomain(recipientDomain).Tags(tags).GroupBy(groupBy).Execute()

Get Bounce Statistics



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
	fromDate := time.Now() // time.Time | Filter statistics after this date (RFC3339 format) (optional)
	toDate := time.Now() // time.Time | Filter statistics before this date (RFC3339 format) (optional)
	senderDomain := "senderDomain_example" // string | Filter by sender domain (optional)
	recipientDomain := "recipientDomain_example" // string | Filter by recipient domain (optional)
	tags := "tags_example" // string | Filter by tags (comma-separated) (optional)
	groupBy := "groupBy_example" // string | Group by time period (optional) (default to "day")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatisticsAPI.GetBounceStatistics(context.Background(), accountId).FromDate(fromDate).ToDate(toDate).SenderDomain(senderDomain).RecipientDomain(recipientDomain).Tags(tags).GroupBy(groupBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticsAPI.GetBounceStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBounceStatistics`: BounceStatisticsResponse
	fmt.Fprintf(os.Stdout, "Response from `StatisticsAPI.GetBounceStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBounceStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **time.Time** | Filter statistics after this date (RFC3339 format) | 
 **toDate** | **time.Time** | Filter statistics before this date (RFC3339 format) | 
 **senderDomain** | **string** | Filter by sender domain | 
 **recipientDomain** | **string** | Filter by recipient domain | 
 **tags** | **string** | Filter by tags (comma-separated) | 
 **groupBy** | **string** | Group by time period | [default to &quot;day&quot;]

### Return type

[**BounceStatisticsResponse**](BounceStatisticsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeliverabilityStatistics

> DeliverabilityStatisticsResponse GetDeliverabilityStatistics(ctx, accountId).FromDate(fromDate).ToDate(toDate).SenderDomain(senderDomain).RecipientDomain(recipientDomain).Tags(tags).GroupBy(groupBy).Execute()

Get Deliverability Statistics



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
	fromDate := time.Now() // time.Time | Filter statistics after this date (RFC3339 format) (optional)
	toDate := time.Now() // time.Time | Filter statistics before this date (RFC3339 format) (optional)
	senderDomain := "senderDomain_example" // string | Filter by sender domain (optional)
	recipientDomain := "recipientDomain_example" // string | Filter by recipient domain (optional)
	tags := "tags_example" // string | Filter by tags (comma-separated) (optional)
	groupBy := "groupBy_example" // string | Group by time period (optional) (default to "day")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatisticsAPI.GetDeliverabilityStatistics(context.Background(), accountId).FromDate(fromDate).ToDate(toDate).SenderDomain(senderDomain).RecipientDomain(recipientDomain).Tags(tags).GroupBy(groupBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticsAPI.GetDeliverabilityStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeliverabilityStatistics`: DeliverabilityStatisticsResponse
	fmt.Fprintf(os.Stdout, "Response from `StatisticsAPI.GetDeliverabilityStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeliverabilityStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **time.Time** | Filter statistics after this date (RFC3339 format) | 
 **toDate** | **time.Time** | Filter statistics before this date (RFC3339 format) | 
 **senderDomain** | **string** | Filter by sender domain | 
 **recipientDomain** | **string** | Filter by recipient domain | 
 **tags** | **string** | Filter by tags (comma-separated) | 
 **groupBy** | **string** | Group by time period | [default to &quot;day&quot;]

### Return type

[**DeliverabilityStatisticsResponse**](DeliverabilityStatisticsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeliveryTimeStatistics

> DeliveryTimeStatisticsResponse GetDeliveryTimeStatistics(ctx, accountId).FromDate(fromDate).ToDate(toDate).SenderDomain(senderDomain).Tags(tags).GroupBy(groupBy).Execute()

Get Delivery Time Statistics



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
	fromDate := time.Now() // time.Time | Filter statistics after this date (RFC3339 format) (optional)
	toDate := time.Now() // time.Time | Filter statistics before this date (RFC3339 format) (optional)
	senderDomain := "senderDomain_example" // string | Filter by sender domain (optional)
	tags := "tags_example" // string | Filter by tags (comma-separated) (optional)
	groupBy := "groupBy_example" // string | Group by time period (optional) (default to "day")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatisticsAPI.GetDeliveryTimeStatistics(context.Background(), accountId).FromDate(fromDate).ToDate(toDate).SenderDomain(senderDomain).Tags(tags).GroupBy(groupBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatisticsAPI.GetDeliveryTimeStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeliveryTimeStatistics`: DeliveryTimeStatisticsResponse
	fmt.Fprintf(os.Stdout, "Response from `StatisticsAPI.GetDeliveryTimeStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **uuid.UUID** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeliveryTimeStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **time.Time** | Filter statistics after this date (RFC3339 format) | 
 **toDate** | **time.Time** | Filter statistics before this date (RFC3339 format) | 
 **senderDomain** | **string** | Filter by sender domain | 
 **tags** | **string** | Filter by tags (comma-separated) | 
 **groupBy** | **string** | Group by time period | [default to &quot;day&quot;]

### Return type

[**DeliveryTimeStatisticsResponse**](DeliveryTimeStatisticsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

