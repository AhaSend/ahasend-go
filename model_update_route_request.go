package ahasend

import (
	"encoding/json"
)

// checks if the UpdateRouteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRouteRequest{}

// UpdateRouteRequest struct for UpdateRouteRequest
type UpdateRouteRequest struct {
	// Route name
	Name *string `json:"name,omitempty"`
	// Webhook URL for the route
	Url *string `json:"url,omitempty"`
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

type _UpdateRouteRequest UpdateRouteRequest

// NewUpdateRouteRequest instantiates a new UpdateRouteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRouteRequest() *UpdateRouteRequest {
	this := UpdateRouteRequest{}
	return &this
}

// NewUpdateRouteRequestWithDefaults instantiates a new UpdateRouteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRouteRequestWithDefaults() *UpdateRouteRequest {
	this := UpdateRouteRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateRouteRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRouteRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateRouteRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateRouteRequest) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UpdateRouteRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRouteRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UpdateRouteRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UpdateRouteRequest) SetUrl(v string) {
	o.Url = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *UpdateRouteRequest) GetRecipient() string {
	if o == nil || IsNil(o.Recipient) {
		var ret string
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRouteRequest) GetRecipientOk() (*string, bool) {
	if o == nil || IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *UpdateRouteRequest) HasRecipient() bool {
	if o != nil && !IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given string and assigns it to the Recipient field.
func (o *UpdateRouteRequest) SetRecipient(v string) {
	o.Recipient = &v
}

// GetIncludeAttachments returns the IncludeAttachments field value if set, zero value otherwise.
func (o *UpdateRouteRequest) GetIncludeAttachments() bool {
	if o == nil || IsNil(o.IncludeAttachments) {
		var ret bool
		return ret
	}
	return *o.IncludeAttachments
}

// GetIncludeAttachmentsOk returns a tuple with the IncludeAttachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRouteRequest) GetIncludeAttachmentsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeAttachments) {
		return nil, false
	}
	return o.IncludeAttachments, true
}

// HasIncludeAttachments returns a boolean if a field has been set.
func (o *UpdateRouteRequest) HasIncludeAttachments() bool {
	if o != nil && !IsNil(o.IncludeAttachments) {
		return true
	}

	return false
}

// SetIncludeAttachments gets a reference to the given bool and assigns it to the IncludeAttachments field.
func (o *UpdateRouteRequest) SetIncludeAttachments(v bool) {
	o.IncludeAttachments = &v
}

// GetIncludeHeaders returns the IncludeHeaders field value if set, zero value otherwise.
func (o *UpdateRouteRequest) GetIncludeHeaders() bool {
	if o == nil || IsNil(o.IncludeHeaders) {
		var ret bool
		return ret
	}
	return *o.IncludeHeaders
}

// GetIncludeHeadersOk returns a tuple with the IncludeHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRouteRequest) GetIncludeHeadersOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeHeaders) {
		return nil, false
	}
	return o.IncludeHeaders, true
}

// HasIncludeHeaders returns a boolean if a field has been set.
func (o *UpdateRouteRequest) HasIncludeHeaders() bool {
	if o != nil && !IsNil(o.IncludeHeaders) {
		return true
	}

	return false
}

// SetIncludeHeaders gets a reference to the given bool and assigns it to the IncludeHeaders field.
func (o *UpdateRouteRequest) SetIncludeHeaders(v bool) {
	o.IncludeHeaders = &v
}

// GetGroupByMessageId returns the GroupByMessageId field value if set, zero value otherwise.
func (o *UpdateRouteRequest) GetGroupByMessageId() bool {
	if o == nil || IsNil(o.GroupByMessageId) {
		var ret bool
		return ret
	}
	return *o.GroupByMessageId
}

// GetGroupByMessageIdOk returns a tuple with the GroupByMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRouteRequest) GetGroupByMessageIdOk() (*bool, bool) {
	if o == nil || IsNil(o.GroupByMessageId) {
		return nil, false
	}
	return o.GroupByMessageId, true
}

// HasGroupByMessageId returns a boolean if a field has been set.
func (o *UpdateRouteRequest) HasGroupByMessageId() bool {
	if o != nil && !IsNil(o.GroupByMessageId) {
		return true
	}

	return false
}

// SetGroupByMessageId gets a reference to the given bool and assigns it to the GroupByMessageId field.
func (o *UpdateRouteRequest) SetGroupByMessageId(v bool) {
	o.GroupByMessageId = &v
}

// GetStripReplies returns the StripReplies field value if set, zero value otherwise.
func (o *UpdateRouteRequest) GetStripReplies() bool {
	if o == nil || IsNil(o.StripReplies) {
		var ret bool
		return ret
	}
	return *o.StripReplies
}

// GetStripRepliesOk returns a tuple with the StripReplies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRouteRequest) GetStripRepliesOk() (*bool, bool) {
	if o == nil || IsNil(o.StripReplies) {
		return nil, false
	}
	return o.StripReplies, true
}

// HasStripReplies returns a boolean if a field has been set.
func (o *UpdateRouteRequest) HasStripReplies() bool {
	if o != nil && !IsNil(o.StripReplies) {
		return true
	}

	return false
}

// SetStripReplies gets a reference to the given bool and assigns it to the StripReplies field.
func (o *UpdateRouteRequest) SetStripReplies(v bool) {
	o.StripReplies = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateRouteRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRouteRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateRouteRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateRouteRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o UpdateRouteRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRouteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
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

func (o *UpdateRouteRequest) UnmarshalJSON(data []byte) (err error) {
	varUpdateRouteRequest := _UpdateRouteRequest{}

	err = json.Unmarshal(data, &varUpdateRouteRequest)

	if err != nil {
		return err
	}

	*o = UpdateRouteRequest(varUpdateRouteRequest)

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

type NullableUpdateRouteRequest struct {
	value *UpdateRouteRequest
	isSet bool
}

func (v NullableUpdateRouteRequest) Get() *UpdateRouteRequest {
	return v.value
}

func (v *NullableUpdateRouteRequest) Set(val *UpdateRouteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRouteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRouteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRouteRequest(val *UpdateRouteRequest) *NullableUpdateRouteRequest {
	return &NullableUpdateRouteRequest{value: val, isSet: true}
}

func (v NullableUpdateRouteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRouteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
