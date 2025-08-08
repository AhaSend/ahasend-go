# PaginatedAPIKeysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Object type identifier | 
**Data** | [**[]ModelAPIKey**](ModelAPIKey.md) | Array of API keys | 
**Pagination** | [**PaginationInfo**](PaginationInfo.md) |  | 

## Methods

### NewPaginatedAPIKeysResponse

`func NewPaginatedAPIKeysResponse(object string, data []ModelAPIKey, pagination PaginationInfo, ) *PaginatedAPIKeysResponse`

NewPaginatedAPIKeysResponse instantiates a new PaginatedAPIKeysResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedAPIKeysResponseWithDefaults

`func NewPaginatedAPIKeysResponseWithDefaults() *PaginatedAPIKeysResponse`

NewPaginatedAPIKeysResponseWithDefaults instantiates a new PaginatedAPIKeysResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *PaginatedAPIKeysResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaginatedAPIKeysResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaginatedAPIKeysResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetData

`func (o *PaginatedAPIKeysResponse) GetData() []ModelAPIKey`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaginatedAPIKeysResponse) GetDataOk() (*[]ModelAPIKey, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaginatedAPIKeysResponse) SetData(v []ModelAPIKey)`

SetData sets Data field to given value.


### GetPagination

`func (o *PaginatedAPIKeysResponse) GetPagination() PaginationInfo`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedAPIKeysResponse) GetPaginationOk() (*PaginationInfo, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedAPIKeysResponse) SetPagination(v PaginationInfo)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


