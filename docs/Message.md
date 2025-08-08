# Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Object type identifier | 
**CreatedAt** | **time.Time** | When the message was created | 
**UpdatedAt** | **time.Time** | When the message was last updated | 
**SentAt** | Pointer to **time.Time** | When the message was sent | [optional] 
**DeliveredAt** | Pointer to **time.Time** | When the message was delivered | [optional] 
**RetainUntil** | **time.Time** | When the message data will be purged | 
**Direction** | **string** | Message direction | 
**IsBounceNotification** | **bool** | Whether this is a bounce notification | 
**BounceClassification** | Pointer to **string** | Classification of bounce if applicable | [optional] 
**DeliveryAttempts** | [**[]DeliveryEvent**](DeliveryEvent.md) | List of delivery attempts for this message | 
**MessageId** | **string** | Message-ID header value | 
**ApiId** | Pointer to [**uuid.UUID**](uuid.UUID.md) | API-generated message ID | [optional] 
**AhasendId** | **string** | Internal AhaSend message ID | 
**Subject** | **string** | Message subject | 
**Tags** | **[]string** | Tags associated with the message | 
**Sender** | **string** | Sender email address | 
**Recipient** | **string** | Recipient email address | 
**Status** | **string** | Current message status | 
**NumAttempts** | **int32** | Number of delivery attempts | 
**ClickCount** | **int32** | Number of clicks tracked for this message | 
**OpenCount** | **int32** | Number of opens tracked for this message | 
**ReferenceMessageId** | Pointer to **int64** | ID of the original message (for bounce messages) | [optional] 
**DomainId** | [**uuid.UUID**](uuid.UUID.md) | Domain ID this message was sent from | 
**AccountId** | [**uuid.UUID**](uuid.UUID.md) | Account ID this message belongs to | 

## Methods

### NewMessage

`func NewMessage(object string, createdAt time.Time, updatedAt time.Time, retainUntil time.Time, direction string, isBounceNotification bool, deliveryAttempts []DeliveryEvent, messageId string, ahasendId string, subject string, tags []string, sender string, recipient string, status string, numAttempts int32, clickCount int32, openCount int32, domainId uuid.UUID, accountId uuid.UUID, ) *Message`

NewMessage instantiates a new Message object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageWithDefaults

`func NewMessageWithDefaults() *Message`

NewMessageWithDefaults instantiates a new Message object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Message) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Message) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Message) SetObject(v string)`

SetObject sets Object field to given value.


### GetCreatedAt

`func (o *Message) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Message) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Message) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Message) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Message) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Message) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetSentAt

`func (o *Message) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *Message) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *Message) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *Message) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.

### GetDeliveredAt

`func (o *Message) GetDeliveredAt() time.Time`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *Message) GetDeliveredAtOk() (*time.Time, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *Message) SetDeliveredAt(v time.Time)`

SetDeliveredAt sets DeliveredAt field to given value.

### HasDeliveredAt

`func (o *Message) HasDeliveredAt() bool`

HasDeliveredAt returns a boolean if a field has been set.

### GetRetainUntil

`func (o *Message) GetRetainUntil() time.Time`

GetRetainUntil returns the RetainUntil field if non-nil, zero value otherwise.

### GetRetainUntilOk

`func (o *Message) GetRetainUntilOk() (*time.Time, bool)`

GetRetainUntilOk returns a tuple with the RetainUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainUntil

`func (o *Message) SetRetainUntil(v time.Time)`

SetRetainUntil sets RetainUntil field to given value.


### GetDirection

`func (o *Message) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *Message) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *Message) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetIsBounceNotification

`func (o *Message) GetIsBounceNotification() bool`

GetIsBounceNotification returns the IsBounceNotification field if non-nil, zero value otherwise.

### GetIsBounceNotificationOk

`func (o *Message) GetIsBounceNotificationOk() (*bool, bool)`

GetIsBounceNotificationOk returns a tuple with the IsBounceNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBounceNotification

`func (o *Message) SetIsBounceNotification(v bool)`

SetIsBounceNotification sets IsBounceNotification field to given value.


### GetBounceClassification

`func (o *Message) GetBounceClassification() string`

GetBounceClassification returns the BounceClassification field if non-nil, zero value otherwise.

### GetBounceClassificationOk

`func (o *Message) GetBounceClassificationOk() (*string, bool)`

GetBounceClassificationOk returns a tuple with the BounceClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounceClassification

`func (o *Message) SetBounceClassification(v string)`

SetBounceClassification sets BounceClassification field to given value.

### HasBounceClassification

`func (o *Message) HasBounceClassification() bool`

HasBounceClassification returns a boolean if a field has been set.

### GetDeliveryAttempts

`func (o *Message) GetDeliveryAttempts() []DeliveryEvent`

GetDeliveryAttempts returns the DeliveryAttempts field if non-nil, zero value otherwise.

### GetDeliveryAttemptsOk

`func (o *Message) GetDeliveryAttemptsOk() (*[]DeliveryEvent, bool)`

GetDeliveryAttemptsOk returns a tuple with the DeliveryAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAttempts

`func (o *Message) SetDeliveryAttempts(v []DeliveryEvent)`

SetDeliveryAttempts sets DeliveryAttempts field to given value.


### GetMessageId

`func (o *Message) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *Message) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *Message) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.


### GetApiId

`func (o *Message) GetApiId() uuid.UUID`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *Message) GetApiIdOk() (*uuid.UUID, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *Message) SetApiId(v uuid.UUID)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *Message) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetAhasendId

`func (o *Message) GetAhasendId() string`

GetAhasendId returns the AhasendId field if non-nil, zero value otherwise.

### GetAhasendIdOk

`func (o *Message) GetAhasendIdOk() (*string, bool)`

GetAhasendIdOk returns a tuple with the AhasendId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAhasendId

`func (o *Message) SetAhasendId(v string)`

SetAhasendId sets AhasendId field to given value.


### GetSubject

`func (o *Message) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Message) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Message) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetTags

`func (o *Message) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Message) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Message) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetSender

`func (o *Message) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *Message) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *Message) SetSender(v string)`

SetSender sets Sender field to given value.


### GetRecipient

`func (o *Message) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *Message) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *Message) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetStatus

`func (o *Message) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Message) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Message) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetNumAttempts

`func (o *Message) GetNumAttempts() int32`

GetNumAttempts returns the NumAttempts field if non-nil, zero value otherwise.

### GetNumAttemptsOk

`func (o *Message) GetNumAttemptsOk() (*int32, bool)`

GetNumAttemptsOk returns a tuple with the NumAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAttempts

`func (o *Message) SetNumAttempts(v int32)`

SetNumAttempts sets NumAttempts field to given value.


### GetClickCount

`func (o *Message) GetClickCount() int32`

GetClickCount returns the ClickCount field if non-nil, zero value otherwise.

### GetClickCountOk

`func (o *Message) GetClickCountOk() (*int32, bool)`

GetClickCountOk returns a tuple with the ClickCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickCount

`func (o *Message) SetClickCount(v int32)`

SetClickCount sets ClickCount field to given value.


### GetOpenCount

`func (o *Message) GetOpenCount() int32`

GetOpenCount returns the OpenCount field if non-nil, zero value otherwise.

### GetOpenCountOk

`func (o *Message) GetOpenCountOk() (*int32, bool)`

GetOpenCountOk returns a tuple with the OpenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenCount

`func (o *Message) SetOpenCount(v int32)`

SetOpenCount sets OpenCount field to given value.


### GetReferenceMessageId

`func (o *Message) GetReferenceMessageId() int64`

GetReferenceMessageId returns the ReferenceMessageId field if non-nil, zero value otherwise.

### GetReferenceMessageIdOk

`func (o *Message) GetReferenceMessageIdOk() (*int64, bool)`

GetReferenceMessageIdOk returns a tuple with the ReferenceMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceMessageId

`func (o *Message) SetReferenceMessageId(v int64)`

SetReferenceMessageId sets ReferenceMessageId field to given value.

### HasReferenceMessageId

`func (o *Message) HasReferenceMessageId() bool`

HasReferenceMessageId returns a boolean if a field has been set.

### GetDomainId

`func (o *Message) GetDomainId() uuid.UUID`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *Message) GetDomainIdOk() (*uuid.UUID, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *Message) SetDomainId(v uuid.UUID)`

SetDomainId sets DomainId field to given value.


### GetAccountId

`func (o *Message) GetAccountId() uuid.UUID`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Message) GetAccountIdOk() (*uuid.UUID, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Message) SetAccountId(v uuid.UUID)`

SetAccountId sets AccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


