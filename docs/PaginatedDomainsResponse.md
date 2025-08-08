# PaginatedDomainsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Object type identifier | 
**Data** | [**[]Domain**](Domain.md) | Array of domains | 
**Pagination** | [**PaginationInfo**](PaginationInfo.md) |  | 

## Methods

### NewPaginatedDomainsResponse

`func NewPaginatedDomainsResponse(object string, data []Domain, pagination PaginationInfo, ) *PaginatedDomainsResponse`

NewPaginatedDomainsResponse instantiates a new PaginatedDomainsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedDomainsResponseWithDefaults

`func NewPaginatedDomainsResponseWithDefaults() *PaginatedDomainsResponse`

NewPaginatedDomainsResponseWithDefaults instantiates a new PaginatedDomainsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *PaginatedDomainsResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaginatedDomainsResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaginatedDomainsResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetData

`func (o *PaginatedDomainsResponse) GetData() []Domain`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaginatedDomainsResponse) GetDataOk() (*[]Domain, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaginatedDomainsResponse) SetData(v []Domain)`

SetData sets Data field to given value.


### GetPagination

`func (o *PaginatedDomainsResponse) GetPagination() PaginationInfo`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedDomainsResponse) GetPaginationOk() (*PaginationInfo, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedDomainsResponse) SetPagination(v PaginationInfo)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


