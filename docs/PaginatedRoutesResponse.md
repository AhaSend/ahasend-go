# PaginatedRoutesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Object type identifier | 
**Data** | [**[]Route**](Route.md) | Array of routes | 
**Pagination** | [**PaginationInfo**](PaginationInfo.md) |  | 

## Methods

### NewPaginatedRoutesResponse

`func NewPaginatedRoutesResponse(object string, data []Route, pagination PaginationInfo, ) *PaginatedRoutesResponse`

NewPaginatedRoutesResponse instantiates a new PaginatedRoutesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedRoutesResponseWithDefaults

`func NewPaginatedRoutesResponseWithDefaults() *PaginatedRoutesResponse`

NewPaginatedRoutesResponseWithDefaults instantiates a new PaginatedRoutesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *PaginatedRoutesResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaginatedRoutesResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaginatedRoutesResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetData

`func (o *PaginatedRoutesResponse) GetData() []Route`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaginatedRoutesResponse) GetDataOk() (*[]Route, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaginatedRoutesResponse) SetData(v []Route)`

SetData sets Data field to given value.


### GetPagination

`func (o *PaginatedRoutesResponse) GetPagination() PaginationInfo`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedRoutesResponse) GetPaginationOk() (*PaginationInfo, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedRoutesResponse) SetPagination(v PaginationInfo)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


