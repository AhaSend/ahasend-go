# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Object type identifier | 
**Id** | [**uuid.UUID**](uuid.UUID.md) | Unique identifier for the account | 
**CreatedAt** | **time.Time** | When the account was created | 
**UpdatedAt** | **time.Time** | When the account was last updated | 
**Name** | **string** | Account name | 
**Website** | Pointer to **string** | Account website URL | [optional] 
**About** | Pointer to **string** | Account description | [optional] 
**TrackOpens** | Pointer to **bool** | Default open tracking setting | [optional] 
**TrackClicks** | Pointer to **bool** | Default click tracking setting | [optional] 
**RejectBadRecipients** | Pointer to **bool** | Whether to reject bad recipients | [optional] 
**RejectMistypedRecipients** | Pointer to **bool** | Whether to reject mistyped recipients | [optional] 
**MessageMetadataRetention** | Pointer to **int32** | Default message metadata retention in days | [optional] 
**MessageDataRetention** | Pointer to **int32** | Default message data retention in days | [optional] 

## Methods

### NewAccount

`func NewAccount(object string, id uuid.UUID, createdAt time.Time, updatedAt time.Time, name string, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Account) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Account) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Account) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *Account) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v uuid.UUID)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Account) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Account) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Account) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Account) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Account) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Account) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetName

`func (o *Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Account) SetName(v string)`

SetName sets Name field to given value.


### GetWebsite

`func (o *Account) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *Account) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *Account) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *Account) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetAbout

`func (o *Account) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *Account) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *Account) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *Account) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### GetTrackOpens

`func (o *Account) GetTrackOpens() bool`

GetTrackOpens returns the TrackOpens field if non-nil, zero value otherwise.

### GetTrackOpensOk

`func (o *Account) GetTrackOpensOk() (*bool, bool)`

GetTrackOpensOk returns a tuple with the TrackOpens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackOpens

`func (o *Account) SetTrackOpens(v bool)`

SetTrackOpens sets TrackOpens field to given value.

### HasTrackOpens

`func (o *Account) HasTrackOpens() bool`

HasTrackOpens returns a boolean if a field has been set.

### GetTrackClicks

`func (o *Account) GetTrackClicks() bool`

GetTrackClicks returns the TrackClicks field if non-nil, zero value otherwise.

### GetTrackClicksOk

`func (o *Account) GetTrackClicksOk() (*bool, bool)`

GetTrackClicksOk returns a tuple with the TrackClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackClicks

`func (o *Account) SetTrackClicks(v bool)`

SetTrackClicks sets TrackClicks field to given value.

### HasTrackClicks

`func (o *Account) HasTrackClicks() bool`

HasTrackClicks returns a boolean if a field has been set.

### GetRejectBadRecipients

`func (o *Account) GetRejectBadRecipients() bool`

GetRejectBadRecipients returns the RejectBadRecipients field if non-nil, zero value otherwise.

### GetRejectBadRecipientsOk

`func (o *Account) GetRejectBadRecipientsOk() (*bool, bool)`

GetRejectBadRecipientsOk returns a tuple with the RejectBadRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectBadRecipients

`func (o *Account) SetRejectBadRecipients(v bool)`

SetRejectBadRecipients sets RejectBadRecipients field to given value.

### HasRejectBadRecipients

`func (o *Account) HasRejectBadRecipients() bool`

HasRejectBadRecipients returns a boolean if a field has been set.

### GetRejectMistypedRecipients

`func (o *Account) GetRejectMistypedRecipients() bool`

GetRejectMistypedRecipients returns the RejectMistypedRecipients field if non-nil, zero value otherwise.

### GetRejectMistypedRecipientsOk

`func (o *Account) GetRejectMistypedRecipientsOk() (*bool, bool)`

GetRejectMistypedRecipientsOk returns a tuple with the RejectMistypedRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectMistypedRecipients

`func (o *Account) SetRejectMistypedRecipients(v bool)`

SetRejectMistypedRecipients sets RejectMistypedRecipients field to given value.

### HasRejectMistypedRecipients

`func (o *Account) HasRejectMistypedRecipients() bool`

HasRejectMistypedRecipients returns a boolean if a field has been set.

### GetMessageMetadataRetention

`func (o *Account) GetMessageMetadataRetention() int32`

GetMessageMetadataRetention returns the MessageMetadataRetention field if non-nil, zero value otherwise.

### GetMessageMetadataRetentionOk

`func (o *Account) GetMessageMetadataRetentionOk() (*int32, bool)`

GetMessageMetadataRetentionOk returns a tuple with the MessageMetadataRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageMetadataRetention

`func (o *Account) SetMessageMetadataRetention(v int32)`

SetMessageMetadataRetention sets MessageMetadataRetention field to given value.

### HasMessageMetadataRetention

`func (o *Account) HasMessageMetadataRetention() bool`

HasMessageMetadataRetention returns a boolean if a field has been set.

### GetMessageDataRetention

`func (o *Account) GetMessageDataRetention() int32`

GetMessageDataRetention returns the MessageDataRetention field if non-nil, zero value otherwise.

### GetMessageDataRetentionOk

`func (o *Account) GetMessageDataRetentionOk() (*int32, bool)`

GetMessageDataRetentionOk returns a tuple with the MessageDataRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageDataRetention

`func (o *Account) SetMessageDataRetention(v int32)`

SetMessageDataRetention sets MessageDataRetention field to given value.

### HasMessageDataRetention

`func (o *Account) HasMessageDataRetention() bool`

HasMessageDataRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


