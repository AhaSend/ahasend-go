# DeliveryEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **time.Time** | Timestamp of the delivery event | 
**Log** | **string** | Log message for the delivery event | 
**Status** | **string** | Status of the delivery event | 

## Methods

### NewDeliveryEvent

`func NewDeliveryEvent(time time.Time, log string, status string, ) *DeliveryEvent`

NewDeliveryEvent instantiates a new DeliveryEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryEventWithDefaults

`func NewDeliveryEventWithDefaults() *DeliveryEvent`

NewDeliveryEventWithDefaults instantiates a new DeliveryEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *DeliveryEvent) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *DeliveryEvent) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *DeliveryEvent) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetLog

`func (o *DeliveryEvent) GetLog() string`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *DeliveryEvent) GetLogOk() (*string, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *DeliveryEvent) SetLog(v string)`

SetLog sets Log field to given value.


### GetStatus

`func (o *DeliveryEvent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeliveryEvent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeliveryEvent) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


