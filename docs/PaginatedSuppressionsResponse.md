# PaginatedSuppressionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Object type identifier | 
**Data** | [**[]Suppression**](Suppression.md) | Array of suppressions | 
**Pagination** | [**PaginationInfo**](PaginationInfo.md) |  | 

## Methods

### NewPaginatedSuppressionsResponse

`func NewPaginatedSuppressionsResponse(object string, data []Suppression, pagination PaginationInfo, ) *PaginatedSuppressionsResponse`

NewPaginatedSuppressionsResponse instantiates a new PaginatedSuppressionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedSuppressionsResponseWithDefaults

`func NewPaginatedSuppressionsResponseWithDefaults() *PaginatedSuppressionsResponse`

NewPaginatedSuppressionsResponseWithDefaults instantiates a new PaginatedSuppressionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *PaginatedSuppressionsResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaginatedSuppressionsResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaginatedSuppressionsResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetData

`func (o *PaginatedSuppressionsResponse) GetData() []Suppression`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaginatedSuppressionsResponse) GetDataOk() (*[]Suppression, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaginatedSuppressionsResponse) SetData(v []Suppression)`

SetData sets Data field to given value.


### GetPagination

`func (o *PaginatedSuppressionsResponse) GetPagination() PaginationInfo`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedSuppressionsResponse) GetPaginationOk() (*PaginationInfo, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedSuppressionsResponse) SetPagination(v PaginationInfo)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


