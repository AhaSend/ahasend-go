# DeliverabilityStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromTimestamp** | **time.Time** | Start time of the statistics bucket | 
**ToTimestamp** | **time.Time** | End time of the statistics bucket | 
**ReceptionCount** | Pointer to **int32** | Number of messages accepted for delivery | [optional] 
**DeliveredCount** | Pointer to **int32** | Number of messages delivered | [optional] 
**DeferredCount** | Pointer to **int32** | Number of messages deferred | [optional] 
**BouncedCount** | Pointer to **int32** | Number of messages bounced | [optional] 
**FailedCount** | Pointer to **int32** | Number of messages failed | [optional] 
**SuppressedCount** | Pointer to **int32** | Number of messages suppressed | [optional] 
**OpenedCount** | Pointer to **int32** | Number of messages opened at least once | [optional] 
**ClickedCount** | Pointer to **int32** | Number of messages that have at least one link in them clicked. | [optional] 

## Methods

### NewDeliverabilityStatistics

`func NewDeliverabilityStatistics(fromTimestamp time.Time, toTimestamp time.Time, ) *DeliverabilityStatistics`

NewDeliverabilityStatistics instantiates a new DeliverabilityStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliverabilityStatisticsWithDefaults

`func NewDeliverabilityStatisticsWithDefaults() *DeliverabilityStatistics`

NewDeliverabilityStatisticsWithDefaults instantiates a new DeliverabilityStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromTimestamp

`func (o *DeliverabilityStatistics) GetFromTimestamp() time.Time`

GetFromTimestamp returns the FromTimestamp field if non-nil, zero value otherwise.

### GetFromTimestampOk

`func (o *DeliverabilityStatistics) GetFromTimestampOk() (*time.Time, bool)`

GetFromTimestampOk returns a tuple with the FromTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromTimestamp

`func (o *DeliverabilityStatistics) SetFromTimestamp(v time.Time)`

SetFromTimestamp sets FromTimestamp field to given value.


### GetToTimestamp

`func (o *DeliverabilityStatistics) GetToTimestamp() time.Time`

GetToTimestamp returns the ToTimestamp field if non-nil, zero value otherwise.

### GetToTimestampOk

`func (o *DeliverabilityStatistics) GetToTimestampOk() (*time.Time, bool)`

GetToTimestampOk returns a tuple with the ToTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToTimestamp

`func (o *DeliverabilityStatistics) SetToTimestamp(v time.Time)`

SetToTimestamp sets ToTimestamp field to given value.


### GetReceptionCount

`func (o *DeliverabilityStatistics) GetReceptionCount() int32`

GetReceptionCount returns the ReceptionCount field if non-nil, zero value otherwise.

### GetReceptionCountOk

`func (o *DeliverabilityStatistics) GetReceptionCountOk() (*int32, bool)`

GetReceptionCountOk returns a tuple with the ReceptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceptionCount

`func (o *DeliverabilityStatistics) SetReceptionCount(v int32)`

SetReceptionCount sets ReceptionCount field to given value.

### HasReceptionCount

`func (o *DeliverabilityStatistics) HasReceptionCount() bool`

HasReceptionCount returns a boolean if a field has been set.

### GetDeliveredCount

`func (o *DeliverabilityStatistics) GetDeliveredCount() int32`

GetDeliveredCount returns the DeliveredCount field if non-nil, zero value otherwise.

### GetDeliveredCountOk

`func (o *DeliverabilityStatistics) GetDeliveredCountOk() (*int32, bool)`

GetDeliveredCountOk returns a tuple with the DeliveredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredCount

`func (o *DeliverabilityStatistics) SetDeliveredCount(v int32)`

SetDeliveredCount sets DeliveredCount field to given value.

### HasDeliveredCount

`func (o *DeliverabilityStatistics) HasDeliveredCount() bool`

HasDeliveredCount returns a boolean if a field has been set.

### GetDeferredCount

`func (o *DeliverabilityStatistics) GetDeferredCount() int32`

GetDeferredCount returns the DeferredCount field if non-nil, zero value otherwise.

### GetDeferredCountOk

`func (o *DeliverabilityStatistics) GetDeferredCountOk() (*int32, bool)`

GetDeferredCountOk returns a tuple with the DeferredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferredCount

`func (o *DeliverabilityStatistics) SetDeferredCount(v int32)`

SetDeferredCount sets DeferredCount field to given value.

### HasDeferredCount

`func (o *DeliverabilityStatistics) HasDeferredCount() bool`

HasDeferredCount returns a boolean if a field has been set.

### GetBouncedCount

`func (o *DeliverabilityStatistics) GetBouncedCount() int32`

GetBouncedCount returns the BouncedCount field if non-nil, zero value otherwise.

### GetBouncedCountOk

`func (o *DeliverabilityStatistics) GetBouncedCountOk() (*int32, bool)`

GetBouncedCountOk returns a tuple with the BouncedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBouncedCount

`func (o *DeliverabilityStatistics) SetBouncedCount(v int32)`

SetBouncedCount sets BouncedCount field to given value.

### HasBouncedCount

`func (o *DeliverabilityStatistics) HasBouncedCount() bool`

HasBouncedCount returns a boolean if a field has been set.

### GetFailedCount

`func (o *DeliverabilityStatistics) GetFailedCount() int32`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *DeliverabilityStatistics) GetFailedCountOk() (*int32, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCount

`func (o *DeliverabilityStatistics) SetFailedCount(v int32)`

SetFailedCount sets FailedCount field to given value.

### HasFailedCount

`func (o *DeliverabilityStatistics) HasFailedCount() bool`

HasFailedCount returns a boolean if a field has been set.

### GetSuppressedCount

`func (o *DeliverabilityStatistics) GetSuppressedCount() int32`

GetSuppressedCount returns the SuppressedCount field if non-nil, zero value otherwise.

### GetSuppressedCountOk

`func (o *DeliverabilityStatistics) GetSuppressedCountOk() (*int32, bool)`

GetSuppressedCountOk returns a tuple with the SuppressedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressedCount

`func (o *DeliverabilityStatistics) SetSuppressedCount(v int32)`

SetSuppressedCount sets SuppressedCount field to given value.

### HasSuppressedCount

`func (o *DeliverabilityStatistics) HasSuppressedCount() bool`

HasSuppressedCount returns a boolean if a field has been set.

### GetOpenedCount

`func (o *DeliverabilityStatistics) GetOpenedCount() int32`

GetOpenedCount returns the OpenedCount field if non-nil, zero value otherwise.

### GetOpenedCountOk

`func (o *DeliverabilityStatistics) GetOpenedCountOk() (*int32, bool)`

GetOpenedCountOk returns a tuple with the OpenedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedCount

`func (o *DeliverabilityStatistics) SetOpenedCount(v int32)`

SetOpenedCount sets OpenedCount field to given value.

### HasOpenedCount

`func (o *DeliverabilityStatistics) HasOpenedCount() bool`

HasOpenedCount returns a boolean if a field has been set.

### GetClickedCount

`func (o *DeliverabilityStatistics) GetClickedCount() int32`

GetClickedCount returns the ClickedCount field if non-nil, zero value otherwise.

### GetClickedCountOk

`func (o *DeliverabilityStatistics) GetClickedCountOk() (*int32, bool)`

GetClickedCountOk returns a tuple with the ClickedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickedCount

`func (o *DeliverabilityStatistics) SetClickedCount(v int32)`

SetClickedCount sets ClickedCount field to given value.

### HasClickedCount

`func (o *DeliverabilityStatistics) HasClickedCount() bool`

HasClickedCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


