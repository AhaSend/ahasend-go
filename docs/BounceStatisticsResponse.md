# BounceStatisticsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Object type identifier | 
**Data** | [**[]BounceStatistics**](BounceStatistics.md) | Array of bounce statistics | 

## Methods

### NewBounceStatisticsResponse

`func NewBounceStatisticsResponse(object string, data []BounceStatistics, ) *BounceStatisticsResponse`

NewBounceStatisticsResponse instantiates a new BounceStatisticsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBounceStatisticsResponseWithDefaults

`func NewBounceStatisticsResponseWithDefaults() *BounceStatisticsResponse`

NewBounceStatisticsResponseWithDefaults instantiates a new BounceStatisticsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *BounceStatisticsResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BounceStatisticsResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BounceStatisticsResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetData

`func (o *BounceStatisticsResponse) GetData() []BounceStatistics`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BounceStatisticsResponse) GetDataOk() (*[]BounceStatistics, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BounceStatisticsResponse) SetData(v []BounceStatistics)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


