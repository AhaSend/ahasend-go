# DeliveryTimeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromTimestamp** | **time.Time** | Start time of the statistics bucket | 
**ToTimestamp** | **time.Time** | End time of the statistics bucket | 
**AvgDeliveryTime** | **float64** | Average delivery time in seconds | 
**DeliveredCount** | **int32** | Number of messages | 
**DeliveryTimes** | Pointer to [**[]DeliveryTime**](DeliveryTime.md) |  | [optional] 

## Methods

### NewDeliveryTimeStatistics

`func NewDeliveryTimeStatistics(fromTimestamp time.Time, toTimestamp time.Time, avgDeliveryTime float64, deliveredCount int32, ) *DeliveryTimeStatistics`

NewDeliveryTimeStatistics instantiates a new DeliveryTimeStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryTimeStatisticsWithDefaults

`func NewDeliveryTimeStatisticsWithDefaults() *DeliveryTimeStatistics`

NewDeliveryTimeStatisticsWithDefaults instantiates a new DeliveryTimeStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromTimestamp

`func (o *DeliveryTimeStatistics) GetFromTimestamp() time.Time`

GetFromTimestamp returns the FromTimestamp field if non-nil, zero value otherwise.

### GetFromTimestampOk

`func (o *DeliveryTimeStatistics) GetFromTimestampOk() (*time.Time, bool)`

GetFromTimestampOk returns a tuple with the FromTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromTimestamp

`func (o *DeliveryTimeStatistics) SetFromTimestamp(v time.Time)`

SetFromTimestamp sets FromTimestamp field to given value.


### GetToTimestamp

`func (o *DeliveryTimeStatistics) GetToTimestamp() time.Time`

GetToTimestamp returns the ToTimestamp field if non-nil, zero value otherwise.

### GetToTimestampOk

`func (o *DeliveryTimeStatistics) GetToTimestampOk() (*time.Time, bool)`

GetToTimestampOk returns a tuple with the ToTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToTimestamp

`func (o *DeliveryTimeStatistics) SetToTimestamp(v time.Time)`

SetToTimestamp sets ToTimestamp field to given value.


### GetAvgDeliveryTime

`func (o *DeliveryTimeStatistics) GetAvgDeliveryTime() float64`

GetAvgDeliveryTime returns the AvgDeliveryTime field if non-nil, zero value otherwise.

### GetAvgDeliveryTimeOk

`func (o *DeliveryTimeStatistics) GetAvgDeliveryTimeOk() (*float64, bool)`

GetAvgDeliveryTimeOk returns a tuple with the AvgDeliveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDeliveryTime

`func (o *DeliveryTimeStatistics) SetAvgDeliveryTime(v float64)`

SetAvgDeliveryTime sets AvgDeliveryTime field to given value.


### GetDeliveredCount

`func (o *DeliveryTimeStatistics) GetDeliveredCount() int32`

GetDeliveredCount returns the DeliveredCount field if non-nil, zero value otherwise.

### GetDeliveredCountOk

`func (o *DeliveryTimeStatistics) GetDeliveredCountOk() (*int32, bool)`

GetDeliveredCountOk returns a tuple with the DeliveredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredCount

`func (o *DeliveryTimeStatistics) SetDeliveredCount(v int32)`

SetDeliveredCount sets DeliveredCount field to given value.


### GetDeliveryTimes

`func (o *DeliveryTimeStatistics) GetDeliveryTimes() []DeliveryTime`

GetDeliveryTimes returns the DeliveryTimes field if non-nil, zero value otherwise.

### GetDeliveryTimesOk

`func (o *DeliveryTimeStatistics) GetDeliveryTimesOk() (*[]DeliveryTime, bool)`

GetDeliveryTimesOk returns a tuple with the DeliveryTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTimes

`func (o *DeliveryTimeStatistics) SetDeliveryTimes(v []DeliveryTime)`

SetDeliveryTimes sets DeliveryTimes field to given value.

### HasDeliveryTimes

`func (o *DeliveryTimeStatistics) HasDeliveryTimes() bool`

HasDeliveryTimes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


