# UpdateAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Account name | [optional] 
**Website** | Pointer to **string** | Account website URL | [optional] 
**About** | Pointer to **string** | Account description (used for account verification) | [optional] 
**TrackOpens** | Pointer to **bool** | Default open tracking setting | [optional] 
**TrackClicks** | Pointer to **bool** | Default click tracking setting | [optional] 
**RejectBadRecipients** | Pointer to **bool** | Whether to reject bad recipients | [optional] 
**RejectMistypedRecipients** | Pointer to **bool** | Whether to reject mistyped recipients | [optional] 
**MessageMetadataRetention** | Pointer to **int32** | Default message metadata retention in days | [optional] 
**MessageDataRetention** | Pointer to **int32** | Default message data retention in days | [optional] 

## Methods

### NewUpdateAccountRequest

`func NewUpdateAccountRequest() *UpdateAccountRequest`

NewUpdateAccountRequest instantiates a new UpdateAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountRequestWithDefaults

`func NewUpdateAccountRequestWithDefaults() *UpdateAccountRequest`

NewUpdateAccountRequestWithDefaults instantiates a new UpdateAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAccountRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAccountRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAccountRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAccountRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWebsite

`func (o *UpdateAccountRequest) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *UpdateAccountRequest) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *UpdateAccountRequest) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *UpdateAccountRequest) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetAbout

`func (o *UpdateAccountRequest) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *UpdateAccountRequest) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *UpdateAccountRequest) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *UpdateAccountRequest) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### GetTrackOpens

`func (o *UpdateAccountRequest) GetTrackOpens() bool`

GetTrackOpens returns the TrackOpens field if non-nil, zero value otherwise.

### GetTrackOpensOk

`func (o *UpdateAccountRequest) GetTrackOpensOk() (*bool, bool)`

GetTrackOpensOk returns a tuple with the TrackOpens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackOpens

`func (o *UpdateAccountRequest) SetTrackOpens(v bool)`

SetTrackOpens sets TrackOpens field to given value.

### HasTrackOpens

`func (o *UpdateAccountRequest) HasTrackOpens() bool`

HasTrackOpens returns a boolean if a field has been set.

### GetTrackClicks

`func (o *UpdateAccountRequest) GetTrackClicks() bool`

GetTrackClicks returns the TrackClicks field if non-nil, zero value otherwise.

### GetTrackClicksOk

`func (o *UpdateAccountRequest) GetTrackClicksOk() (*bool, bool)`

GetTrackClicksOk returns a tuple with the TrackClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackClicks

`func (o *UpdateAccountRequest) SetTrackClicks(v bool)`

SetTrackClicks sets TrackClicks field to given value.

### HasTrackClicks

`func (o *UpdateAccountRequest) HasTrackClicks() bool`

HasTrackClicks returns a boolean if a field has been set.

### GetRejectBadRecipients

`func (o *UpdateAccountRequest) GetRejectBadRecipients() bool`

GetRejectBadRecipients returns the RejectBadRecipients field if non-nil, zero value otherwise.

### GetRejectBadRecipientsOk

`func (o *UpdateAccountRequest) GetRejectBadRecipientsOk() (*bool, bool)`

GetRejectBadRecipientsOk returns a tuple with the RejectBadRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectBadRecipients

`func (o *UpdateAccountRequest) SetRejectBadRecipients(v bool)`

SetRejectBadRecipients sets RejectBadRecipients field to given value.

### HasRejectBadRecipients

`func (o *UpdateAccountRequest) HasRejectBadRecipients() bool`

HasRejectBadRecipients returns a boolean if a field has been set.

### GetRejectMistypedRecipients

`func (o *UpdateAccountRequest) GetRejectMistypedRecipients() bool`

GetRejectMistypedRecipients returns the RejectMistypedRecipients field if non-nil, zero value otherwise.

### GetRejectMistypedRecipientsOk

`func (o *UpdateAccountRequest) GetRejectMistypedRecipientsOk() (*bool, bool)`

GetRejectMistypedRecipientsOk returns a tuple with the RejectMistypedRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectMistypedRecipients

`func (o *UpdateAccountRequest) SetRejectMistypedRecipients(v bool)`

SetRejectMistypedRecipients sets RejectMistypedRecipients field to given value.

### HasRejectMistypedRecipients

`func (o *UpdateAccountRequest) HasRejectMistypedRecipients() bool`

HasRejectMistypedRecipients returns a boolean if a field has been set.

### GetMessageMetadataRetention

`func (o *UpdateAccountRequest) GetMessageMetadataRetention() int32`

GetMessageMetadataRetention returns the MessageMetadataRetention field if non-nil, zero value otherwise.

### GetMessageMetadataRetentionOk

`func (o *UpdateAccountRequest) GetMessageMetadataRetentionOk() (*int32, bool)`

GetMessageMetadataRetentionOk returns a tuple with the MessageMetadataRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageMetadataRetention

`func (o *UpdateAccountRequest) SetMessageMetadataRetention(v int32)`

SetMessageMetadataRetention sets MessageMetadataRetention field to given value.

### HasMessageMetadataRetention

`func (o *UpdateAccountRequest) HasMessageMetadataRetention() bool`

HasMessageMetadataRetention returns a boolean if a field has been set.

### GetMessageDataRetention

`func (o *UpdateAccountRequest) GetMessageDataRetention() int32`

GetMessageDataRetention returns the MessageDataRetention field if non-nil, zero value otherwise.

### GetMessageDataRetentionOk

`func (o *UpdateAccountRequest) GetMessageDataRetentionOk() (*int32, bool)`

GetMessageDataRetentionOk returns a tuple with the MessageDataRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageDataRetention

`func (o *UpdateAccountRequest) SetMessageDataRetention(v int32)`

SetMessageDataRetention sets MessageDataRetention field to given value.

### HasMessageDataRetention

`func (o *UpdateAccountRequest) HasMessageDataRetention() bool`

HasMessageDataRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


