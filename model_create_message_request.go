package ahasend

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMessageRequest{}

// CreateMessageRequest struct for CreateMessageRequest
type CreateMessageRequest struct {
	From SenderAddress `json:"from"`
	// This does not set the To header to multiple addresses, it sends a separate message for each recipient
	Recipients []Recipient `json:"recipients"`
	// If provided, the reply-to header in headers array must not be provided
	ReplyTo *SenderAddress `json:"reply_to,omitempty"`
	// Email subject line
	Subject string `json:"subject"`
	// Plain text content. Required if html_content is empty
	TextContent *string `json:"text_content,omitempty"`
	// HTML content. Required if text_content is empty
	HtmlContent *string `json:"html_content,omitempty"`
	// AMP HTML content
	AmpContent *string `json:"amp_content,omitempty"`
	// File attachments
	Attachments []Attachment `json:"attachments,omitempty"`
	// Custom email headers. reply-to header cannot be provided if reply_to is provided, message-id will be ignored and automatically generated
	Headers map[string]string `json:"headers,omitempty"`
	// Global substitutions, recipient substitutions override global
	Substitutions map[string]interface{} `json:"substitutions,omitempty"`
	// Tags for categorizing messages
	Tags []string `json:"tags,omitempty"`
	// If true, the message will be sent to the sandbox environment
	Sandbox *bool `json:"sandbox,omitempty"`
	// The result of the sandbox send
	SandboxResult *string `json:"sandbox_result,omitempty"`
	// Tracking settings for the message, overrides default account settings
	Tracking *Tracking `json:"tracking,omitempty"`
	// Retention settings for the message, overrides default account settings
	Retention *Retention `json:"retention,omitempty"`
	// Schedule for message delivery
	Schedule             *MessageSchedule `json:"schedule,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateMessageRequest CreateMessageRequest

// NewCreateMessageRequest instantiates a new CreateMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMessageRequest(from SenderAddress, recipients []Recipient, subject string) *CreateMessageRequest {
	this := CreateMessageRequest{}
	this.From = from
	this.Recipients = recipients
	this.Subject = subject
	var sandbox bool = false
	this.Sandbox = &sandbox
	return &this
}

// NewCreateMessageRequestWithDefaults instantiates a new CreateMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMessageRequestWithDefaults() *CreateMessageRequest {
	this := CreateMessageRequest{}
	var sandbox bool = false
	this.Sandbox = &sandbox
	return &this
}

// GetFrom returns the From field value
func (o *CreateMessageRequest) GetFrom() SenderAddress {
	if o == nil {
		var ret SenderAddress
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *CreateMessageRequest) GetFromOk() (*SenderAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *CreateMessageRequest) SetFrom(v SenderAddress) {
	o.From = v
}

// GetRecipients returns the Recipients field value
func (o *CreateMessageRequest) GetRecipients() []Recipient {
	if o == nil {
		var ret []Recipient
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *CreateMessageRequest) GetRecipientsOk() ([]Recipient, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipients, true
}

// SetRecipients sets field value
func (o *CreateMessageRequest) SetRecipients(v []Recipient) {
	o.Recipients = v
}

// GetReplyTo returns the ReplyTo field value if set, zero value otherwise.
func (o *CreateMessageRequest) GetReplyTo() SenderAddress {
	if o == nil || IsNil(o.ReplyTo) {
		var ret SenderAddress
		return ret
	}
	return *o.ReplyTo
}

// GetReplyToOk returns a tuple with the ReplyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessageRequest) GetReplyToOk() (*SenderAddress, bool) {
	if o == nil || IsNil(o.ReplyTo) {
		return nil, false
	}
	return o.ReplyTo, true
}

// HasReplyTo returns a boolean if a field has been set.
func (o *CreateMessageRequest) HasReplyTo() bool {
	if o != nil && !IsNil(o.ReplyTo) {
		return true
	}

	return false
}

// SetReplyTo gets a reference to the given SenderAddress and assigns it to the ReplyTo field.
func (o *CreateMessageRequest) SetReplyTo(v SenderAddress) {
	o.ReplyTo = &v
}

// GetSubject returns the Subject field value
func (o *CreateMessageRequest) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *CreateMessageRequest) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *CreateMessageRequest) SetSubject(v string) {
	o.Subject = v
}

// GetTextContent returns the TextContent field value if set, zero value otherwise.
func (o *CreateMessageRequest) GetTextContent() string {
	if o == nil || IsNil(o.TextContent) {
		var ret string
		return ret
	}
	return *o.TextContent
}

// GetTextContentOk returns a tuple with the TextContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessageRequest) GetTextContentOk() (*string, bool) {
	if o == nil || IsNil(o.TextContent) {
		return nil, false
	}
	return o.TextContent, true
}

// HasTextContent returns a boolean if a field has been set.
func (o *CreateMessageRequest) HasTextContent() bool {
	if o != nil && !IsNil(o.TextContent) {
		return true
	}

	return false
}

// SetTextContent gets a reference to the given string and assigns it to the TextContent field.
func (o *CreateMessageRequest) SetTextContent(v string) {
	o.TextContent = &v
}

// GetHtmlContent returns the HtmlContent field value if set, zero value otherwise.
func (o *CreateMessageRequest) GetHtmlContent() string {
	if o == nil || IsNil(o.HtmlContent) {
		var ret string
		return ret
	}
	return *o.HtmlContent
}

// GetHtmlContentOk returns a tuple with the HtmlContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessageRequest) GetHtmlContentOk() (*string, bool) {
	if o == nil || IsNil(o.HtmlContent) {
		return nil, false
	}
	return o.HtmlContent, true
}

// HasHtmlContent returns a boolean if a field has been set.
func (o *CreateMessageRequest) HasHtmlContent() bool {
	if o != nil && !IsNil(o.HtmlContent) {
		return true
	}

	return false
}

// SetHtmlContent gets a reference to the given string and assigns it to the HtmlContent field.
func (o *CreateMessageRequest) SetHtmlContent(v string) {
	o.HtmlContent = &v
}

// GetAmpContent returns the AmpContent field value if set, zero value otherwise.
func (o *CreateMessageRequest) GetAmpContent() string {
	if o == nil || IsNil(o.AmpContent) {
		var ret string
		return ret
	}
	return *o.AmpContent
}

// GetAmpContentOk returns a tuple with the AmpContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessageRequest) GetAmpContentOk() (*string, bool) {
	if o == nil || IsNil(o.AmpContent) {
		return nil, false
	}
	return o.AmpContent, true
}

// HasAmpContent returns a boolean if a field has been set.
func (o *CreateMessageRequest) HasAmpContent() bool {
	if o != nil && !IsNil(o.AmpContent) {
		return true
	}

	return false
}

// SetAmpContent gets a reference to the given string and assigns it to the AmpContent field.
func (o *CreateMessageRequest) SetAmpContent(v string) {
	o.AmpContent = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *CreateMessageRequest) GetAttachments() []Attachment {
	if o == nil || IsNil(o.Attachments) {
		var ret []Attachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessageRequest) GetAttachmentsOk() ([]Attachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *CreateMessageRequest) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []Attachment and assigns it to the Attachments field.
func (o *CreateMessageRequest) SetAttachments(v []Attachment) {
	o.Attachments = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *CreateMessageRequest) GetHeaders() map[string]string {
	if o == nil || IsNil(o.Headers) {
		var ret map[string]string
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessageRequest) GetHeadersOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return map[string]string{}, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *CreateMessageRequest) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *CreateMessageRequest) SetHeaders(v map[string]string) {
	o.Headers = v
}

// GetSubstitutions returns the Substitutions field value if set, zero value otherwise.
func (o *CreateMessageRequest) GetSubstitutions() map[string]interface{} {
	if o == nil || IsNil(o.Substitutions) {
		var ret map[string]interface{}
		return ret
	}
	return o.Substitutions
}

// GetSubstitutionsOk returns a tuple with the Substitutions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessageRequest) GetSubstitutionsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Substitutions) {
		return map[string]interface{}{}, false
	}
	return o.Substitutions, true
}

// HasSubstitutions returns a boolean if a field has been set.
func (o *CreateMessageRequest) HasSubstitutions() bool {
	if o != nil && !IsNil(o.Substitutions) {
		return true
	}

	return false
}

// SetSubstitutions gets a reference to the given map[string]interface{} and assigns it to the Substitutions field.
func (o *CreateMessageRequest) SetSubstitutions(v map[string]interface{}) {
	o.Substitutions = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateMessageRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessageRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateMessageRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateMessageRequest) SetTags(v []string) {
	o.Tags = v
}

// GetSandbox returns the Sandbox field value if set, zero value otherwise.
func (o *CreateMessageRequest) GetSandbox() bool {
	if o == nil || IsNil(o.Sandbox) {
		var ret bool
		return ret
	}
	return *o.Sandbox
}

// GetSandboxOk returns a tuple with the Sandbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessageRequest) GetSandboxOk() (*bool, bool) {
	if o == nil || IsNil(o.Sandbox) {
		return nil, false
	}
	return o.Sandbox, true
}

// HasSandbox returns a boolean if a field has been set.
func (o *CreateMessageRequest) HasSandbox() bool {
	if o != nil && !IsNil(o.Sandbox) {
		return true
	}

	return false
}

// SetSandbox gets a reference to the given bool and assigns it to the Sandbox field.
func (o *CreateMessageRequest) SetSandbox(v bool) {
	o.Sandbox = &v
}

// GetSandboxResult returns the SandboxResult field value if set, zero value otherwise.
func (o *CreateMessageRequest) GetSandboxResult() string {
	if o == nil || IsNil(o.SandboxResult) {
		var ret string
		return ret
	}
	return *o.SandboxResult
}

// GetSandboxResultOk returns a tuple with the SandboxResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessageRequest) GetSandboxResultOk() (*string, bool) {
	if o == nil || IsNil(o.SandboxResult) {
		return nil, false
	}
	return o.SandboxResult, true
}

// HasSandboxResult returns a boolean if a field has been set.
func (o *CreateMessageRequest) HasSandboxResult() bool {
	if o != nil && !IsNil(o.SandboxResult) {
		return true
	}

	return false
}

// SetSandboxResult gets a reference to the given string and assigns it to the SandboxResult field.
func (o *CreateMessageRequest) SetSandboxResult(v string) {
	o.SandboxResult = &v
}

// GetTracking returns the Tracking field value if set, zero value otherwise.
func (o *CreateMessageRequest) GetTracking() Tracking {
	if o == nil || IsNil(o.Tracking) {
		var ret Tracking
		return ret
	}
	return *o.Tracking
}

// GetTrackingOk returns a tuple with the Tracking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessageRequest) GetTrackingOk() (*Tracking, bool) {
	if o == nil || IsNil(o.Tracking) {
		return nil, false
	}
	return o.Tracking, true
}

// HasTracking returns a boolean if a field has been set.
func (o *CreateMessageRequest) HasTracking() bool {
	if o != nil && !IsNil(o.Tracking) {
		return true
	}

	return false
}

// SetTracking gets a reference to the given Tracking and assigns it to the Tracking field.
func (o *CreateMessageRequest) SetTracking(v Tracking) {
	o.Tracking = &v
}

// GetRetention returns the Retention field value if set, zero value otherwise.
func (o *CreateMessageRequest) GetRetention() Retention {
	if o == nil || IsNil(o.Retention) {
		var ret Retention
		return ret
	}
	return *o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessageRequest) GetRetentionOk() (*Retention, bool) {
	if o == nil || IsNil(o.Retention) {
		return nil, false
	}
	return o.Retention, true
}

// HasRetention returns a boolean if a field has been set.
func (o *CreateMessageRequest) HasRetention() bool {
	if o != nil && !IsNil(o.Retention) {
		return true
	}

	return false
}

// SetRetention gets a reference to the given Retention and assigns it to the Retention field.
func (o *CreateMessageRequest) SetRetention(v Retention) {
	o.Retention = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *CreateMessageRequest) GetSchedule() MessageSchedule {
	if o == nil || IsNil(o.Schedule) {
		var ret MessageSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessageRequest) GetScheduleOk() (*MessageSchedule, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *CreateMessageRequest) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given MessageSchedule and assigns it to the Schedule field.
func (o *CreateMessageRequest) SetSchedule(v MessageSchedule) {
	o.Schedule = &v
}

func (o CreateMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["from"] = o.From
	toSerialize["recipients"] = o.Recipients
	if !IsNil(o.ReplyTo) {
		toSerialize["reply_to"] = o.ReplyTo
	}
	toSerialize["subject"] = o.Subject
	if !IsNil(o.TextContent) {
		toSerialize["text_content"] = o.TextContent
	}
	if !IsNil(o.HtmlContent) {
		toSerialize["html_content"] = o.HtmlContent
	}
	if !IsNil(o.AmpContent) {
		toSerialize["amp_content"] = o.AmpContent
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.Substitutions) {
		toSerialize["substitutions"] = o.Substitutions
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Sandbox) {
		toSerialize["sandbox"] = o.Sandbox
	}
	if !IsNil(o.SandboxResult) {
		toSerialize["sandbox_result"] = o.SandboxResult
	}
	if !IsNil(o.Tracking) {
		toSerialize["tracking"] = o.Tracking
	}
	if !IsNil(o.Retention) {
		toSerialize["retention"] = o.Retention
	}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateMessageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"from",
		"recipients",
		"subject",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateMessageRequest := _CreateMessageRequest{}

	err = json.Unmarshal(data, &varCreateMessageRequest)

	if err != nil {
		return err
	}

	*o = CreateMessageRequest(varCreateMessageRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "from")
		delete(additionalProperties, "recipients")
		delete(additionalProperties, "reply_to")
		delete(additionalProperties, "subject")
		delete(additionalProperties, "text_content")
		delete(additionalProperties, "html_content")
		delete(additionalProperties, "amp_content")
		delete(additionalProperties, "attachments")
		delete(additionalProperties, "headers")
		delete(additionalProperties, "substitutions")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "sandbox")
		delete(additionalProperties, "sandbox_result")
		delete(additionalProperties, "tracking")
		delete(additionalProperties, "retention")
		delete(additionalProperties, "schedule")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateMessageRequest struct {
	value *CreateMessageRequest
	isSet bool
}

func (v NullableCreateMessageRequest) Get() *CreateMessageRequest {
	return v.value
}

func (v *NullableCreateMessageRequest) Set(val *CreateMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMessageRequest(val *CreateMessageRequest) *NullableCreateMessageRequest {
	return &NullableCreateMessageRequest{value: val, isSet: true}
}

func (v NullableCreateMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
