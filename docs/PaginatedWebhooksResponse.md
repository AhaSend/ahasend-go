# PaginatedWebhooksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Object type identifier | 
**Data** | [**[]Webhook**](Webhook.md) | Array of webhooks | 
**Pagination** | [**PaginationInfo**](PaginationInfo.md) |  | 

## Methods

### NewPaginatedWebhooksResponse

`func NewPaginatedWebhooksResponse(object string, data []Webhook, pagination PaginationInfo, ) *PaginatedWebhooksResponse`

NewPaginatedWebhooksResponse instantiates a new PaginatedWebhooksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedWebhooksResponseWithDefaults

`func NewPaginatedWebhooksResponseWithDefaults() *PaginatedWebhooksResponse`

NewPaginatedWebhooksResponseWithDefaults instantiates a new PaginatedWebhooksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *PaginatedWebhooksResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaginatedWebhooksResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaginatedWebhooksResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetData

`func (o *PaginatedWebhooksResponse) GetData() []Webhook`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaginatedWebhooksResponse) GetDataOk() (*[]Webhook, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaginatedWebhooksResponse) SetData(v []Webhook)`

SetData sets Data field to given value.


### GetPagination

`func (o *PaginatedWebhooksResponse) GetPagination() PaginationInfo`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedWebhooksResponse) GetPaginationOk() (*PaginationInfo, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedWebhooksResponse) SetPagination(v PaginationInfo)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


