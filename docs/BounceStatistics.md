# BounceStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromTimestamp** | **time.Time** | Start time of the statistics bucket | 
**ToTimestamp** | **time.Time** | End time of the statistics bucket | 
**Bounces** | [**[]Bounce**](Bounce.md) | Bounce count per bounce classification | 

## Methods

### NewBounceStatistics

`func NewBounceStatistics(fromTimestamp time.Time, toTimestamp time.Time, bounces []Bounce, ) *BounceStatistics`

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


### GetBounces

`func (o *BounceStatistics) GetBounces() []Bounce`

GetBounces returns the Bounces field if non-nil, zero value otherwise.

### GetBouncesOk

`func (o *BounceStatistics) GetBouncesOk() (*[]Bounce, bool)`

GetBouncesOk returns a tuple with the Bounces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounces

`func (o *BounceStatistics) SetBounces(v []Bounce)`

SetBounces sets Bounces field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


