# UpdateRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Route name | [optional] 
**Url** | Pointer to **string** | Webhook URL for the route | [optional] 
**Recipient** | Pointer to **string** | Recipient filter | [optional] 
**IncludeAttachments** | Pointer to **bool** | Whether to include attachments in webhooks | [optional] 
**IncludeHeaders** | Pointer to **bool** | Whether to include headers in webhooks | [optional] 
**GroupByMessageId** | Pointer to **bool** | Whether to group by message ID | [optional] 
**StripReplies** | Pointer to **bool** | Whether to strip reply content | [optional] 
**Enabled** | Pointer to **bool** | Whether the route is enabled | [optional] 

## Methods

### NewUpdateRouteRequest

`func NewUpdateRouteRequest() *UpdateRouteRequest`

NewUpdateRouteRequest instantiates a new UpdateRouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRouteRequestWithDefaults

`func NewUpdateRouteRequestWithDefaults() *UpdateRouteRequest`

NewUpdateRouteRequestWithDefaults instantiates a new UpdateRouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateRouteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRouteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRouteRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateRouteRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateRouteRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateRouteRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateRouteRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateRouteRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRecipient

`func (o *UpdateRouteRequest) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *UpdateRouteRequest) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *UpdateRouteRequest) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *UpdateRouteRequest) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetIncludeAttachments

`func (o *UpdateRouteRequest) GetIncludeAttachments() bool`

GetIncludeAttachments returns the IncludeAttachments field if non-nil, zero value otherwise.

### GetIncludeAttachmentsOk

`func (o *UpdateRouteRequest) GetIncludeAttachmentsOk() (*bool, bool)`

GetIncludeAttachmentsOk returns a tuple with the IncludeAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttachments

`func (o *UpdateRouteRequest) SetIncludeAttachments(v bool)`

SetIncludeAttachments sets IncludeAttachments field to given value.

### HasIncludeAttachments

`func (o *UpdateRouteRequest) HasIncludeAttachments() bool`

HasIncludeAttachments returns a boolean if a field has been set.

### GetIncludeHeaders

`func (o *UpdateRouteRequest) GetIncludeHeaders() bool`

GetIncludeHeaders returns the IncludeHeaders field if non-nil, zero value otherwise.

### GetIncludeHeadersOk

`func (o *UpdateRouteRequest) GetIncludeHeadersOk() (*bool, bool)`

GetIncludeHeadersOk returns a tuple with the IncludeHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeHeaders

`func (o *UpdateRouteRequest) SetIncludeHeaders(v bool)`

SetIncludeHeaders sets IncludeHeaders field to given value.

### HasIncludeHeaders

`func (o *UpdateRouteRequest) HasIncludeHeaders() bool`

HasIncludeHeaders returns a boolean if a field has been set.

### GetGroupByMessageId

`func (o *UpdateRouteRequest) GetGroupByMessageId() bool`

GetGroupByMessageId returns the GroupByMessageId field if non-nil, zero value otherwise.

### GetGroupByMessageIdOk

`func (o *UpdateRouteRequest) GetGroupByMessageIdOk() (*bool, bool)`

GetGroupByMessageIdOk returns a tuple with the GroupByMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByMessageId

`func (o *UpdateRouteRequest) SetGroupByMessageId(v bool)`

SetGroupByMessageId sets GroupByMessageId field to given value.

### HasGroupByMessageId

`func (o *UpdateRouteRequest) HasGroupByMessageId() bool`

HasGroupByMessageId returns a boolean if a field has been set.

### GetStripReplies

`func (o *UpdateRouteRequest) GetStripReplies() bool`

GetStripReplies returns the StripReplies field if non-nil, zero value otherwise.

### GetStripRepliesOk

`func (o *UpdateRouteRequest) GetStripRepliesOk() (*bool, bool)`

GetStripRepliesOk returns a tuple with the StripReplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripReplies

`func (o *UpdateRouteRequest) SetStripReplies(v bool)`

SetStripReplies sets StripReplies field to given value.

### HasStripReplies

`func (o *UpdateRouteRequest) HasStripReplies() bool`

HasStripReplies returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateRouteRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateRouteRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateRouteRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateRouteRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


