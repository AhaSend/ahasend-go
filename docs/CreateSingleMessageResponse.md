# CreateSingleMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Object type identifier | 
**Id** | Pointer to **string** | Message ID (null if the message was not sent) | [optional] 
**Recipient** | [**Recipient**](Recipient.md) |  | 
**Status** | **string** | Status of the message | 
**Error** | Pointer to **string** | Error message if the message was not sent | [optional] 
**Schedule** | Pointer to [**MessageSchedule**](MessageSchedule.md) | Provided if the request contained a schedule | [optional] 

## Methods

### NewCreateSingleMessageResponse

`func NewCreateSingleMessageResponse(object string, recipient Recipient, status string, ) *CreateSingleMessageResponse`

NewCreateSingleMessageResponse instantiates a new CreateSingleMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSingleMessageResponseWithDefaults

`func NewCreateSingleMessageResponseWithDefaults() *CreateSingleMessageResponse`

NewCreateSingleMessageResponseWithDefaults instantiates a new CreateSingleMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *CreateSingleMessageResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CreateSingleMessageResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CreateSingleMessageResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *CreateSingleMessageResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateSingleMessageResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateSingleMessageResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateSingleMessageResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecipient

`func (o *CreateSingleMessageResponse) GetRecipient() Recipient`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *CreateSingleMessageResponse) GetRecipientOk() (*Recipient, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *CreateSingleMessageResponse) SetRecipient(v Recipient)`

SetRecipient sets Recipient field to given value.


### GetStatus

`func (o *CreateSingleMessageResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateSingleMessageResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateSingleMessageResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetError

`func (o *CreateSingleMessageResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CreateSingleMessageResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CreateSingleMessageResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CreateSingleMessageResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetSchedule

`func (o *CreateSingleMessageResponse) GetSchedule() MessageSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CreateSingleMessageResponse) GetScheduleOk() (*MessageSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CreateSingleMessageResponse) SetSchedule(v MessageSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CreateSingleMessageResponse) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


