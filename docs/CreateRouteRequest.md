# CreateRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Route name | 
**Url** | **string** | Webhook URL for the route | 
**Recipient** | Pointer to **string** | Recipient filter | [optional] 
**IncludeAttachments** | Pointer to **bool** | Whether to include attachments in webhooks | [optional] [default to false]
**IncludeHeaders** | Pointer to **bool** | Whether to include headers in webhooks | [optional] [default to false]
**GroupByMessageId** | Pointer to **bool** | Whether to group by message ID | [optional] [default to false]
**StripReplies** | Pointer to **bool** | Whether to strip reply content | [optional] [default to false]
**Enabled** | Pointer to **bool** | Whether the route is enabled | [optional] [default to true]

## Methods

### NewCreateRouteRequest

`func NewCreateRouteRequest(name string, url string, ) *CreateRouteRequest`

NewCreateRouteRequest instantiates a new CreateRouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRouteRequestWithDefaults

`func NewCreateRouteRequestWithDefaults() *CreateRouteRequest`

NewCreateRouteRequestWithDefaults instantiates a new CreateRouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRouteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRouteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRouteRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *CreateRouteRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateRouteRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateRouteRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRecipient

`func (o *CreateRouteRequest) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *CreateRouteRequest) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *CreateRouteRequest) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *CreateRouteRequest) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetIncludeAttachments

`func (o *CreateRouteRequest) GetIncludeAttachments() bool`

GetIncludeAttachments returns the IncludeAttachments field if non-nil, zero value otherwise.

### GetIncludeAttachmentsOk

`func (o *CreateRouteRequest) GetIncludeAttachmentsOk() (*bool, bool)`

GetIncludeAttachmentsOk returns a tuple with the IncludeAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttachments

`func (o *CreateRouteRequest) SetIncludeAttachments(v bool)`

SetIncludeAttachments sets IncludeAttachments field to given value.

### HasIncludeAttachments

`func (o *CreateRouteRequest) HasIncludeAttachments() bool`

HasIncludeAttachments returns a boolean if a field has been set.

### GetIncludeHeaders

`func (o *CreateRouteRequest) GetIncludeHeaders() bool`

GetIncludeHeaders returns the IncludeHeaders field if non-nil, zero value otherwise.

### GetIncludeHeadersOk

`func (o *CreateRouteRequest) GetIncludeHeadersOk() (*bool, bool)`

GetIncludeHeadersOk returns a tuple with the IncludeHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeHeaders

`func (o *CreateRouteRequest) SetIncludeHeaders(v bool)`

SetIncludeHeaders sets IncludeHeaders field to given value.

### HasIncludeHeaders

`func (o *CreateRouteRequest) HasIncludeHeaders() bool`

HasIncludeHeaders returns a boolean if a field has been set.

### GetGroupByMessageId

`func (o *CreateRouteRequest) GetGroupByMessageId() bool`

GetGroupByMessageId returns the GroupByMessageId field if non-nil, zero value otherwise.

### GetGroupByMessageIdOk

`func (o *CreateRouteRequest) GetGroupByMessageIdOk() (*bool, bool)`

GetGroupByMessageIdOk returns a tuple with the GroupByMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByMessageId

`func (o *CreateRouteRequest) SetGroupByMessageId(v bool)`

SetGroupByMessageId sets GroupByMessageId field to given value.

### HasGroupByMessageId

`func (o *CreateRouteRequest) HasGroupByMessageId() bool`

HasGroupByMessageId returns a boolean if a field has been set.

### GetStripReplies

`func (o *CreateRouteRequest) GetStripReplies() bool`

GetStripReplies returns the StripReplies field if non-nil, zero value otherwise.

### GetStripRepliesOk

`func (o *CreateRouteRequest) GetStripRepliesOk() (*bool, bool)`

GetStripRepliesOk returns a tuple with the StripReplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripReplies

`func (o *CreateRouteRequest) SetStripReplies(v bool)`

SetStripReplies sets StripReplies field to given value.

### HasStripReplies

`func (o *CreateRouteRequest) HasStripReplies() bool`

HasStripReplies returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateRouteRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateRouteRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateRouteRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateRouteRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


