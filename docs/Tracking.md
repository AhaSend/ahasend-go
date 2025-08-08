# Tracking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Open** | Pointer to **bool** | Whether to track opens | [optional] 
**Click** | Pointer to **bool** | Whether to track clicks | [optional] 

## Methods

### NewTracking

`func NewTracking() *Tracking`

NewTracking instantiates a new Tracking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackingWithDefaults

`func NewTrackingWithDefaults() *Tracking`

NewTrackingWithDefaults instantiates a new Tracking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpen

`func (o *Tracking) GetOpen() bool`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *Tracking) GetOpenOk() (*bool, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *Tracking) SetOpen(v bool)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *Tracking) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### GetClick

`func (o *Tracking) GetClick() bool`

GetClick returns the Click field if non-nil, zero value otherwise.

### GetClickOk

`func (o *Tracking) GetClickOk() (*bool, bool)`

GetClickOk returns a tuple with the Click field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClick

`func (o *Tracking) SetClick(v bool)`

SetClick sets Click field to given value.

### HasClick

`func (o *Tracking) HasClick() bool`

HasClick returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


