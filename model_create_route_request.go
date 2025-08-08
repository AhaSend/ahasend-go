package ahasend

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateRouteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRouteRequest{}

// CreateRouteRequest struct for CreateRouteRequest
type CreateRouteRequest struct {
	// Route name
	Name string `json:"name"`
	// Webhook URL for the route
	Url string `json:"url"`
	// Recipient filter
	Recipient *string `json:"recipient,omitempty"`
	// Whether to include attachments in webhooks
	IncludeAttachments *bool `json:"include_attachments,omitempty"`
	// Whether to include headers in webhooks
	IncludeHeaders *bool `json:"include_headers,omitempty"`
	// Whether to group by message ID
	GroupByMessageId *bool `json:"group_by_message_id,omitempty"`
	// Whether to strip reply content
	StripReplies *bool `json:"strip_replies,omitempty"`
	// Whether the route is enabled
	Enabled              *bool `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateRouteRequest CreateRouteRequest

// NewCreateRouteRequest instantiates a new CreateRouteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRouteRequest(name string, url string) *CreateRouteRequest {
	this := CreateRouteRequest{}
	this.Name = name
	this.Url = url
	var includeAttachments bool = false
	this.IncludeAttachments = &includeAttachments
	var includeHeaders bool = false
	this.IncludeHeaders = &includeHeaders
	var groupByMessageId bool = false
	this.GroupByMessageId = &groupByMessageId
	var stripReplies bool = false
	this.StripReplies = &stripReplies
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// NewCreateRouteRequestWithDefaults instantiates a new CreateRouteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRouteRequestWithDefaults() *CreateRouteRequest {
	this := CreateRouteRequest{}
	var includeAttachments bool = false
	this.IncludeAttachments = &includeAttachments
	var includeHeaders bool = false
	this.IncludeHeaders = &includeHeaders
	var groupByMessageId bool = false
	this.GroupByMessageId = &groupByMessageId
	var stripReplies bool = false
	this.StripReplies = &stripReplies
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetName returns the Name field value
func (o *CreateRouteRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateRouteRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateRouteRequest) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *CreateRouteRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CreateRouteRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CreateRouteRequest) SetUrl(v string) {
	o.Url = v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *CreateRouteRequest) GetRecipient() string {
	if o == nil || IsNil(o.Recipient) {
		var ret string
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteRequest) GetRecipientOk() (*string, bool) {
	if o == nil || IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *CreateRouteRequest) HasRecipient() bool {
	if o != nil && !IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given string and assigns it to the Recipient field.
func (o *CreateRouteRequest) SetRecipient(v string) {
	o.Recipient = &v
}

// GetIncludeAttachments returns the IncludeAttachments field value if set, zero value otherwise.
func (o *CreateRouteRequest) GetIncludeAttachments() bool {
	if o == nil || IsNil(o.IncludeAttachments) {
		var ret bool
		return ret
	}
	return *o.IncludeAttachments
}

// GetIncludeAttachmentsOk returns a tuple with the IncludeAttachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteRequest) GetIncludeAttachmentsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeAttachments) {
		return nil, false
	}
	return o.IncludeAttachments, true
}

// HasIncludeAttachments returns a boolean if a field has been set.
func (o *CreateRouteRequest) HasIncludeAttachments() bool {
	if o != nil && !IsNil(o.IncludeAttachments) {
		return true
	}

	return false
}

// SetIncludeAttachments gets a reference to the given bool and assigns it to the IncludeAttachments field.
func (o *CreateRouteRequest) SetIncludeAttachments(v bool) {
	o.IncludeAttachments = &v
}

// GetIncludeHeaders returns the IncludeHeaders field value if set, zero value otherwise.
func (o *CreateRouteRequest) GetIncludeHeaders() bool {
	if o == nil || IsNil(o.IncludeHeaders) {
		var ret bool
		return ret
	}
	return *o.IncludeHeaders
}

// GetIncludeHeadersOk returns a tuple with the IncludeHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteRequest) GetIncludeHeadersOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeHeaders) {
		return nil, false
	}
	return o.IncludeHeaders, true
}

// HasIncludeHeaders returns a boolean if a field has been set.
func (o *CreateRouteRequest) HasIncludeHeaders() bool {
	if o != nil && !IsNil(o.IncludeHeaders) {
		return true
	}

	return false
}

// SetIncludeHeaders gets a reference to the given bool and assigns it to the IncludeHeaders field.
func (o *CreateRouteRequest) SetIncludeHeaders(v bool) {
	o.IncludeHeaders = &v
}

// GetGroupByMessageId returns the GroupByMessageId field value if set, zero value otherwise.
func (o *CreateRouteRequest) GetGroupByMessageId() bool {
	if o == nil || IsNil(o.GroupByMessageId) {
		var ret bool
		return ret
	}
	return *o.GroupByMessageId
}

// GetGroupByMessageIdOk returns a tuple with the GroupByMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteRequest) GetGroupByMessageIdOk() (*bool, bool) {
	if o == nil || IsNil(o.GroupByMessageId) {
		return nil, false
	}
	return o.GroupByMessageId, true
}

// HasGroupByMessageId returns a boolean if a field has been set.
func (o *CreateRouteRequest) HasGroupByMessageId() bool {
	if o != nil && !IsNil(o.GroupByMessageId) {
		return true
	}

	return false
}

// SetGroupByMessageId gets a reference to the given bool and assigns it to the GroupByMessageId field.
func (o *CreateRouteRequest) SetGroupByMessageId(v bool) {
	o.GroupByMessageId = &v
}

// GetStripReplies returns the StripReplies field value if set, zero value otherwise.
func (o *CreateRouteRequest) GetStripReplies() bool {
	if o == nil || IsNil(o.StripReplies) {
		var ret bool
		return ret
	}
	return *o.StripReplies
}

// GetStripRepliesOk returns a tuple with the StripReplies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteRequest) GetStripRepliesOk() (*bool, bool) {
	if o == nil || IsNil(o.StripReplies) {
		return nil, false
	}
	return o.StripReplies, true
}

// HasStripReplies returns a boolean if a field has been set.
func (o *CreateRouteRequest) HasStripReplies() bool {
	if o != nil && !IsNil(o.StripReplies) {
		return true
	}

	return false
}

// SetStripReplies gets a reference to the given bool and assigns it to the StripReplies field.
func (o *CreateRouteRequest) SetStripReplies(v bool) {
	o.StripReplies = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateRouteRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateRouteRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateRouteRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o CreateRouteRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRouteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["url"] = o.Url
	if !IsNil(o.Recipient) {
		toSerialize["recipient"] = o.Recipient
	}
	if !IsNil(o.IncludeAttachments) {
		toSerialize["include_attachments"] = o.IncludeAttachments
	}
	if !IsNil(o.IncludeHeaders) {
		toSerialize["include_headers"] = o.IncludeHeaders
	}
	if !IsNil(o.GroupByMessageId) {
		toSerialize["group_by_message_id"] = o.GroupByMessageId
	}
	if !IsNil(o.StripReplies) {
		toSerialize["strip_replies"] = o.StripReplies
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateRouteRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"url",
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

	varCreateRouteRequest := _CreateRouteRequest{}

	err = json.Unmarshal(data, &varCreateRouteRequest)

	if err != nil {
		return err
	}

	*o = CreateRouteRequest(varCreateRouteRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "url")
		delete(additionalProperties, "recipient")
		delete(additionalProperties, "include_attachments")
		delete(additionalProperties, "include_headers")
		delete(additionalProperties, "group_by_message_id")
		delete(additionalProperties, "strip_replies")
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateRouteRequest struct {
	value *CreateRouteRequest
	isSet bool
}

func (v NullableCreateRouteRequest) Get() *CreateRouteRequest {
	return v.value
}

func (v *NullableCreateRouteRequest) Set(val *CreateRouteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRouteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRouteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRouteRequest(val *CreateRouteRequest) *NullableCreateRouteRequest {
	return &NullableCreateRouteRequest{value: val, isSet: true}
}

func (v NullableCreateRouteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRouteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
