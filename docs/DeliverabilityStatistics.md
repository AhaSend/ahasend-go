# DeliverabilityStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromTimestamp** | **time.Time** | Start time of the statistics bucket | 
**ToTimestamp** | **time.Time** | End time of the statistics bucket | 
**Direction** | **string** | Message direction | 
**Sent** | Pointer to **int32** | Number of messages sent | [optional] 
**Delivered** | Pointer to **int32** | Number of messages delivered | [optional] 
**Bounced** | Pointer to **int32** | Number of messages bounced | [optional] 
**Failed** | Pointer to **int32** | Number of messages failed | [optional] 
**Suppressed** | Pointer to **int32** | Number of messages suppressed | [optional] 

## Methods

### NewDeliverabilityStatistics

`func NewDeliverabilityStatistics(fromTimestamp time.Time, toTimestamp time.Time, direction string, ) *DeliverabilityStatistics`

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


### GetDirection

`func (o *DeliverabilityStatistics) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *DeliverabilityStatistics) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *DeliverabilityStatistics) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetSent

`func (o *DeliverabilityStatistics) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *DeliverabilityStatistics) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *DeliverabilityStatistics) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *DeliverabilityStatistics) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetDelivered

`func (o *DeliverabilityStatistics) GetDelivered() int32`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *DeliverabilityStatistics) GetDeliveredOk() (*int32, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *DeliverabilityStatistics) SetDelivered(v int32)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *DeliverabilityStatistics) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetBounced

`func (o *DeliverabilityStatistics) GetBounced() int32`

GetBounced returns the Bounced field if non-nil, zero value otherwise.

### GetBouncedOk

`func (o *DeliverabilityStatistics) GetBouncedOk() (*int32, bool)`

GetBouncedOk returns a tuple with the Bounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounced

`func (o *DeliverabilityStatistics) SetBounced(v int32)`

SetBounced sets Bounced field to given value.

### HasBounced

`func (o *DeliverabilityStatistics) HasBounced() bool`

HasBounced returns a boolean if a field has been set.

### GetFailed

`func (o *DeliverabilityStatistics) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *DeliverabilityStatistics) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *DeliverabilityStatistics) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *DeliverabilityStatistics) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetSuppressed

`func (o *DeliverabilityStatistics) GetSuppressed() int32`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *DeliverabilityStatistics) GetSuppressedOk() (*int32, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *DeliverabilityStatistics) SetSuppressed(v int32)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *DeliverabilityStatistics) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


