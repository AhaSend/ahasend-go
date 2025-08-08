# BounceStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeBucket** | **time.Time** | Time bucket for the statistics | 
**Classification** | **string** | Bounce classification | 
**Count** | **int32** | Number of bounces | 

## Methods

### NewBounceStatistics

`func NewBounceStatistics(timeBucket time.Time, classification string, count int32, ) *BounceStatistics`

NewBounceStatistics instantiates a new BounceStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBounceStatisticsWithDefaults

`func NewBounceStatisticsWithDefaults() *BounceStatistics`

NewBounceStatisticsWithDefaults instantiates a new BounceStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeBucket

`func (o *BounceStatistics) GetTimeBucket() time.Time`

GetTimeBucket returns the TimeBucket field if non-nil, zero value otherwise.

### GetTimeBucketOk

`func (o *BounceStatistics) GetTimeBucketOk() (*time.Time, bool)`

GetTimeBucketOk returns a tuple with the TimeBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBucket

`func (o *BounceStatistics) SetTimeBucket(v time.Time)`

SetTimeBucket sets TimeBucket field to given value.


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


