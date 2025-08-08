# \UtilityAPI

All URIs are relative to *https://api.ahasend.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Ping**](UtilityAPI.md#Ping) | **Get** /v2/ping | Ping



## Ping

> SuccessResponse Ping(ctx).Execute()

Ping



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilityAPI.Ping(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilityAPI.Ping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Ping`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `UtilityAPI.Ping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPingRequest struct via the builder pattern


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

