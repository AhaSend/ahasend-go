# MessageSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstAttempt** | Pointer to **time.Time** | The time to make the first attempt for delivering the message (RFC3339 format) | [optional] 
**Expires** | Pointer to **time.Time** | Expire and drop the message if not delivered by this time (RFC3339 format) | [optional] 

## Methods

### NewMessageSchedule

`func NewMessageSchedule() *MessageSchedule`

NewMessageSchedule instantiates a new MessageSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageScheduleWithDefaults

`func NewMessageScheduleWithDefaults() *MessageSchedule`

NewMessageScheduleWithDefaults instantiates a new MessageSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstAttempt

`func (o *MessageSchedule) GetFirstAttempt() time.Time`

GetFirstAttempt returns the FirstAttempt field if non-nil, zero value otherwise.

### GetFirstAttemptOk

`func (o *MessageSchedule) GetFirstAttemptOk() (*time.Time, bool)`

GetFirstAttemptOk returns a tuple with the FirstAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAttempt

`func (o *MessageSchedule) SetFirstAttempt(v time.Time)`

SetFirstAttempt sets FirstAttempt field to given value.

### HasFirstAttempt

`func (o *MessageSchedule) HasFirstAttempt() bool`

HasFirstAttempt returns a boolean if a field has been set.

### GetExpires

`func (o *MessageSchedule) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *MessageSchedule) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *MessageSchedule) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *MessageSchedule) HasExpires() bool`

HasExpires returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


