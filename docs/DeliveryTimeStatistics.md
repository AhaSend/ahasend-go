# DeliveryTimeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeBucket** | **time.Time** | Time bucket for the statistics | 
**RecipientDomain** | Pointer to **string** | Recipient domain | [optional] 
**AvgDeliveryTime** | **float32** | Average delivery time in seconds | 
**MedianDeliveryTime** | Pointer to **float32** | Median delivery time in seconds | [optional] 
**Count** | **int32** | Number of messages | 

## Methods

### NewDeliveryTimeStatistics

`func NewDeliveryTimeStatistics(timeBucket time.Time, avgDeliveryTime float32, count int32, ) *DeliveryTimeStatistics`

NewDeliveryTimeStatistics instantiates a new DeliveryTimeStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryTimeStatisticsWithDefaults

`func NewDeliveryTimeStatisticsWithDefaults() *DeliveryTimeStatistics`

NewDeliveryTimeStatisticsWithDefaults instantiates a new DeliveryTimeStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeBucket

`func (o *DeliveryTimeStatistics) GetTimeBucket() time.Time`

GetTimeBucket returns the TimeBucket field if non-nil, zero value otherwise.

### GetTimeBucketOk

`func (o *DeliveryTimeStatistics) GetTimeBucketOk() (*time.Time, bool)`

GetTimeBucketOk returns a tuple with the TimeBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBucket

`func (o *DeliveryTimeStatistics) SetTimeBucket(v time.Time)`

SetTimeBucket sets TimeBucket field to given value.


### GetRecipientDomain

`func (o *DeliveryTimeStatistics) GetRecipientDomain() string`

GetRecipientDomain returns the RecipientDomain field if non-nil, zero value otherwise.

### GetRecipientDomainOk

`func (o *DeliveryTimeStatistics) GetRecipientDomainOk() (*string, bool)`

GetRecipientDomainOk returns a tuple with the RecipientDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientDomain

`func (o *DeliveryTimeStatistics) SetRecipientDomain(v string)`

SetRecipientDomain sets RecipientDomain field to given value.

### HasRecipientDomain

`func (o *DeliveryTimeStatistics) HasRecipientDomain() bool`

HasRecipientDomain returns a boolean if a field has been set.

### GetAvgDeliveryTime

`func (o *DeliveryTimeStatistics) GetAvgDeliveryTime() float32`

GetAvgDeliveryTime returns the AvgDeliveryTime field if non-nil, zero value otherwise.

### GetAvgDeliveryTimeOk

`func (o *DeliveryTimeStatistics) GetAvgDeliveryTimeOk() (*float32, bool)`

GetAvgDeliveryTimeOk returns a tuple with the AvgDeliveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDeliveryTime

`func (o *DeliveryTimeStatistics) SetAvgDeliveryTime(v float32)`

SetAvgDeliveryTime sets AvgDeliveryTime field to given value.


### GetMedianDeliveryTime

`func (o *DeliveryTimeStatistics) GetMedianDeliveryTime() float32`

GetMedianDeliveryTime returns the MedianDeliveryTime field if non-nil, zero value otherwise.

### GetMedianDeliveryTimeOk

`func (o *DeliveryTimeStatistics) GetMedianDeliveryTimeOk() (*float32, bool)`

GetMedianDeliveryTimeOk returns a tuple with the MedianDeliveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianDeliveryTime

`func (o *DeliveryTimeStatistics) SetMedianDeliveryTime(v float32)`

SetMedianDeliveryTime sets MedianDeliveryTime field to given value.

### HasMedianDeliveryTime

`func (o *DeliveryTimeStatistics) HasMedianDeliveryTime() bool`

HasMedianDeliveryTime returns a boolean if a field has been set.

### GetCount

`func (o *DeliveryTimeStatistics) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DeliveryTimeStatistics) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DeliveryTimeStatistics) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


