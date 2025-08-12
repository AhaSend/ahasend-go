# BounceStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromTimestamp** | **time.Time** | Start time of the statistics bucket | 
**ToTimestamp** | **time.Time** | End time of the statistics bucket | 
**Classification** | **string** | Bounce classification | 
**Count** | **int32** | Number of bounces | 

## Methods

### NewBounceStatistics

`func NewBounceStatistics(fromTimestamp time.Time, toTimestamp time.Time, classification string, count int32, ) *BounceStatistics`

NewBounceStatistics instantiates a new BounceStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBounceStatisticsWithDefaults

`func NewBounceStatisticsWithDefaults() *BounceStatistics`

NewBounceStatisticsWithDefaults instantiates a new BounceStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromTimestamp

`func (o *BounceStatistics) GetFromTimestamp() time.Time`

GetFromTimestamp returns the FromTimestamp field if non-nil, zero value otherwise.

### GetFromTimestampOk

`func (o *BounceStatistics) GetFromTimestampOk() (*time.Time, bool)`

GetFromTimestampOk returns a tuple with the FromTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromTimestamp

`func (o *BounceStatistics) SetFromTimestamp(v time.Time)`

SetFromTimestamp sets FromTimestamp field to given value.


### GetToTimestamp

`func (o *BounceStatistics) GetToTimestamp() time.Time`

GetToTimestamp returns the ToTimestamp field if non-nil, zero value otherwise.

### GetToTimestampOk

`func (o *BounceStatistics) GetToTimestampOk() (*time.Time, bool)`

GetToTimestampOk returns a tuple with the ToTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToTimestamp

`func (o *BounceStatistics) SetToTimestamp(v time.Time)`

SetToTimestamp sets ToTimestamp field to given value.


### GetClassification

`func (o *BounceStatistics) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *BounceStatistics) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *BounceStatistics) SetClassification(v string)`

SetClassification sets Classification field to given value.


### GetCount

`func (o *BounceStatistics) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BounceStatistics) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BounceStatistics) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


