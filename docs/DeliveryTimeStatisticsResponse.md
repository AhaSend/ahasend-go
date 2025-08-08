# DeliveryTimeStatisticsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Object type identifier | 
**Data** | [**[]DeliveryTimeStatistics**](DeliveryTimeStatistics.md) | Array of delivery time statistics | 

## Methods

### NewDeliveryTimeStatisticsResponse

`func NewDeliveryTimeStatisticsResponse(object string, data []DeliveryTimeStatistics, ) *DeliveryTimeStatisticsResponse`

NewDeliveryTimeStatisticsResponse instantiates a new DeliveryTimeStatisticsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryTimeStatisticsResponseWithDefaults

`func NewDeliveryTimeStatisticsResponseWithDefaults() *DeliveryTimeStatisticsResponse`

NewDeliveryTimeStatisticsResponseWithDefaults instantiates a new DeliveryTimeStatisticsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *DeliveryTimeStatisticsResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *DeliveryTimeStatisticsResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *DeliveryTimeStatisticsResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetData

`func (o *DeliveryTimeStatisticsResponse) GetData() []DeliveryTimeStatistics`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeliveryTimeStatisticsResponse) GetDataOk() (*[]DeliveryTimeStatistics, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeliveryTimeStatisticsResponse) SetData(v []DeliveryTimeStatistics)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


