# Route

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Object type identifier | 
**Id** | [**uuid.UUID**](uuid.UUID.md) | Unique identifier for the route | 
**CreatedAt** | **time.Time** | When the route was created | 
**UpdatedAt** | **time.Time** | When the route was last updated | 
**Name** | **string** | Route name | 
**Url** | **string** | Webhook URL for the route | 
**Recipient** | Pointer to **string** | Recipient filter | [optional] 
**IncludeAttachments** | Pointer to **bool** | Whether to include attachments in webhooks | [optional] 
**IncludeHeaders** | Pointer to **bool** | Whether to include headers in webhooks | [optional] 
**GroupByMessageId** | Pointer to **bool** | Whether to group by message ID | [optional] 
**StripReplies** | Pointer to **bool** | Whether to strip reply content | [optional] 
**Enabled** | **bool** | Whether the route is enabled | 

## Methods

### NewRoute

`func NewRoute(object string, id uuid.UUID, createdAt time.Time, updatedAt time.Time, name string, url string, enabled bool, ) *Route`

NewRoute instantiates a new Route object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteWithDefaults

`func NewRouteWithDefaults() *Route`

NewRouteWithDefaults instantiates a new Route object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Route) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Route) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Route) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *Route) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Route) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Route) SetId(v uuid.UUID)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Route) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Route) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Route) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Route) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Route) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Route) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetName

`func (o *Route) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Route) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Route) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *Route) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Route) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Route) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRecipient

`func (o *Route) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *Route) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *Route) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *Route) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetIncludeAttachments

`func (o *Route) GetIncludeAttachments() bool`

GetIncludeAttachments returns the IncludeAttachments field if non-nil, zero value otherwise.

### GetIncludeAttachmentsOk

`func (o *Route) GetIncludeAttachmentsOk() (*bool, bool)`

GetIncludeAttachmentsOk returns a tuple with the IncludeAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttachments

`func (o *Route) SetIncludeAttachments(v bool)`

SetIncludeAttachments sets IncludeAttachments field to given value.

### HasIncludeAttachments

`func (o *Route) HasIncludeAttachments() bool`

HasIncludeAttachments returns a boolean if a field has been set.

### GetIncludeHeaders

`func (o *Route) GetIncludeHeaders() bool`

GetIncludeHeaders returns the IncludeHeaders field if non-nil, zero value otherwise.

### GetIncludeHeadersOk

`func (o *Route) GetIncludeHeadersOk() (*bool, bool)`

GetIncludeHeadersOk returns a tuple with the IncludeHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeHeaders

`func (o *Route) SetIncludeHeaders(v bool)`

SetIncludeHeaders sets IncludeHeaders field to given value.

### HasIncludeHeaders

`func (o *Route) HasIncludeHeaders() bool`

HasIncludeHeaders returns a boolean if a field has been set.

### GetGroupByMessageId

`func (o *Route) GetGroupByMessageId() bool`

GetGroupByMessageId returns the GroupByMessageId field if non-nil, zero value otherwise.

### GetGroupByMessageIdOk

`func (o *Route) GetGroupByMessageIdOk() (*bool, bool)`

GetGroupByMessageIdOk returns a tuple with the GroupByMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByMessageId

`func (o *Route) SetGroupByMessageId(v bool)`

SetGroupByMessageId sets GroupByMessageId field to given value.

### HasGroupByMessageId

`func (o *Route) HasGroupByMessageId() bool`

HasGroupByMessageId returns a boolean if a field has been set.

### GetStripReplies

`func (o *Route) GetStripReplies() bool`

GetStripReplies returns the StripReplies field if non-nil, zero value otherwise.

### GetStripRepliesOk

`func (o *Route) GetStripRepliesOk() (*bool, bool)`

GetStripRepliesOk returns a tuple with the StripReplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripReplies

`func (o *Route) SetStripReplies(v bool)`

SetStripReplies sets StripReplies field to given value.

### HasStripReplies

`func (o *Route) HasStripReplies() bool`

HasStripReplies returns a boolean if a field has been set.

### GetEnabled

`func (o *Route) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Route) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Route) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


