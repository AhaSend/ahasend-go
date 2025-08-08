# PaginatedMessagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Object type identifier | 
**Data** | [**[]Message**](Message.md) | Array of messages | 
**Pagination** | [**PaginationInfo**](PaginationInfo.md) |  | 

## Methods

### NewPaginatedMessagesResponse

`func NewPaginatedMessagesResponse(object string, data []Message, pagination PaginationInfo, ) *PaginatedMessagesResponse`

NewPaginatedMessagesResponse instantiates a new PaginatedMessagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedMessagesResponseWithDefaults

`func NewPaginatedMessagesResponseWithDefaults() *PaginatedMessagesResponse`

NewPaginatedMessagesResponseWithDefaults instantiates a new PaginatedMessagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *PaginatedMessagesResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaginatedMessagesResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaginatedMessagesResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetData

`func (o *PaginatedMessagesResponse) GetData() []Message`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaginatedMessagesResponse) GetDataOk() (*[]Message, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaginatedMessagesResponse) SetData(v []Message)`

SetData sets Data field to given value.


### GetPagination

`func (o *PaginatedMessagesResponse) GetPagination() PaginationInfo`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedMessagesResponse) GetPaginationOk() (*PaginationInfo, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedMessagesResponse) SetPagination(v PaginationInfo)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


