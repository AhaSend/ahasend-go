# CreateMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**SenderAddress**](SenderAddress.md) |  | 
**Recipients** | [**[]Recipient**](Recipient.md) | This does not set the To header to multiple addresses, it sends a separate message for each recipient | 
**ReplyTo** | Pointer to [**SenderAddress**](SenderAddress.md) | If provided, the reply-to header in headers array must not be provided | [optional] 
**Subject** | **string** | Email subject line | 
**TextContent** | Pointer to **string** | Plain text content. Required if html_content is empty | [optional] 
**HtmlContent** | Pointer to **string** | HTML content. Required if text_content is empty | [optional] 
**AmpContent** | Pointer to **string** | AMP HTML content | [optional] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) | File attachments | [optional] 
**Headers** | Pointer to **map[string]string** | Custom email headers. reply-to header cannot be provided if reply_to is provided, message-id will be ignored and automatically generated | [optional] 
**Substitutions** | Pointer to **map[string]interface{}** | Global substitutions, recipient substitutions override global | [optional] 
**Tags** | Pointer to **[]string** | Tags for categorizing messages | [optional] 
**Sandbox** | Pointer to **bool** | If true, the message will be sent to the sandbox environment | [optional] [default to false]
**SandboxResult** | Pointer to **string** | The result of the sandbox send | [optional] 
**Tracking** | Pointer to [**Tracking**](Tracking.md) | Tracking settings for the message, overrides default account settings | [optional] 
**Retention** | Pointer to [**Retention**](Retention.md) | Retention settings for the message, overrides default account settings | [optional] 
**Schedule** | Pointer to [**MessageSchedule**](MessageSchedule.md) | Schedule for message delivery | [optional] 

## Methods

### NewCreateMessageRequest

`func NewCreateMessageRequest(from SenderAddress, recipients []Recipient, subject string, ) *CreateMessageRequest`

NewCreateMessageRequest instantiates a new CreateMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMessageRequestWithDefaults

`func NewCreateMessageRequestWithDefaults() *CreateMessageRequest`

NewCreateMessageRequestWithDefaults instantiates a new CreateMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *CreateMessageRequest) GetFrom() SenderAddress`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CreateMessageRequest) GetFromOk() (*SenderAddress, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CreateMessageRequest) SetFrom(v SenderAddress)`

SetFrom sets From field to given value.


### GetRecipients

`func (o *CreateMessageRequest) GetRecipients() []Recipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *CreateMessageRequest) GetRecipientsOk() (*[]Recipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *CreateMessageRequest) SetRecipients(v []Recipient)`

SetRecipients sets Recipients field to given value.


### GetReplyTo

`func (o *CreateMessageRequest) GetReplyTo() SenderAddress`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *CreateMessageRequest) GetReplyToOk() (*SenderAddress, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *CreateMessageRequest) SetReplyTo(v SenderAddress)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *CreateMessageRequest) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSubject

`func (o *CreateMessageRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CreateMessageRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CreateMessageRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetTextContent

`func (o *CreateMessageRequest) GetTextContent() string`

GetTextContent returns the TextContent field if non-nil, zero value otherwise.

### GetTextContentOk

`func (o *CreateMessageRequest) GetTextContentOk() (*string, bool)`

GetTextContentOk returns a tuple with the TextContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextContent

`func (o *CreateMessageRequest) SetTextContent(v string)`

SetTextContent sets TextContent field to given value.

### HasTextContent

`func (o *CreateMessageRequest) HasTextContent() bool`

HasTextContent returns a boolean if a field has been set.

### GetHtmlContent

`func (o *CreateMessageRequest) GetHtmlContent() string`

GetHtmlContent returns the HtmlContent field if non-nil, zero value otherwise.

### GetHtmlContentOk

`func (o *CreateMessageRequest) GetHtmlContentOk() (*string, bool)`

GetHtmlContentOk returns a tuple with the HtmlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlContent

`func (o *CreateMessageRequest) SetHtmlContent(v string)`

SetHtmlContent sets HtmlContent field to given value.

### HasHtmlContent

`func (o *CreateMessageRequest) HasHtmlContent() bool`

HasHtmlContent returns a boolean if a field has been set.

### GetAmpContent

`func (o *CreateMessageRequest) GetAmpContent() string`

GetAmpContent returns the AmpContent field if non-nil, zero value otherwise.

### GetAmpContentOk

`func (o *CreateMessageRequest) GetAmpContentOk() (*string, bool)`

GetAmpContentOk returns a tuple with the AmpContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmpContent

`func (o *CreateMessageRequest) SetAmpContent(v string)`

SetAmpContent sets AmpContent field to given value.

### HasAmpContent

`func (o *CreateMessageRequest) HasAmpContent() bool`

HasAmpContent returns a boolean if a field has been set.

### GetAttachments

`func (o *CreateMessageRequest) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *CreateMessageRequest) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *CreateMessageRequest) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *CreateMessageRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetHeaders

`func (o *CreateMessageRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *CreateMessageRequest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *CreateMessageRequest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *CreateMessageRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetSubstitutions

`func (o *CreateMessageRequest) GetSubstitutions() map[string]interface{}`

GetSubstitutions returns the Substitutions field if non-nil, zero value otherwise.

### GetSubstitutionsOk

`func (o *CreateMessageRequest) GetSubstitutionsOk() (*map[string]interface{}, bool)`

GetSubstitutionsOk returns a tuple with the Substitutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstitutions

`func (o *CreateMessageRequest) SetSubstitutions(v map[string]interface{})`

SetSubstitutions sets Substitutions field to given value.

### HasSubstitutions

`func (o *CreateMessageRequest) HasSubstitutions() bool`

HasSubstitutions returns a boolean if a field has been set.

### GetTags

`func (o *CreateMessageRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateMessageRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateMessageRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateMessageRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSandbox

`func (o *CreateMessageRequest) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *CreateMessageRequest) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *CreateMessageRequest) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.

### HasSandbox

`func (o *CreateMessageRequest) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.

### GetSandboxResult

`func (o *CreateMessageRequest) GetSandboxResult() string`

GetSandboxResult returns the SandboxResult field if non-nil, zero value otherwise.

### GetSandboxResultOk

`func (o *CreateMessageRequest) GetSandboxResultOk() (*string, bool)`

GetSandboxResultOk returns a tuple with the SandboxResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxResult

`func (o *CreateMessageRequest) SetSandboxResult(v string)`

SetSandboxResult sets SandboxResult field to given value.

### HasSandboxResult

`func (o *CreateMessageRequest) HasSandboxResult() bool`

HasSandboxResult returns a boolean if a field has been set.

### GetTracking

`func (o *CreateMessageRequest) GetTracking() Tracking`

GetTracking returns the Tracking field if non-nil, zero value otherwise.

### GetTrackingOk

`func (o *CreateMessageRequest) GetTrackingOk() (*Tracking, bool)`

GetTrackingOk returns a tuple with the Tracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracking

`func (o *CreateMessageRequest) SetTracking(v Tracking)`

SetTracking sets Tracking field to given value.

### HasTracking

`func (o *CreateMessageRequest) HasTracking() bool`

HasTracking returns a boolean if a field has been set.

### GetRetention

`func (o *CreateMessageRequest) GetRetention() Retention`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *CreateMessageRequest) GetRetentionOk() (*Retention, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *CreateMessageRequest) SetRetention(v Retention)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *CreateMessageRequest) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetSchedule

`func (o *CreateMessageRequest) GetSchedule() MessageSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CreateMessageRequest) GetScheduleOk() (*MessageSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CreateMessageRequest) SetSchedule(v MessageSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CreateMessageRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


