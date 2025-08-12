package ahasend

import (
	"encoding/json"
)

// checks if the UpdateAPIKeyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAPIKeyRequest{}

// UpdateAPIKeyRequest struct for UpdateAPIKeyRequest
type UpdateAPIKeyRequest struct {
	// Human-readable label for the API key
	Label *string `json:"label,omitempty"`
	// Array of scope strings to grant to this API key
	Scopes               []string `json:"scopes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateAPIKeyRequest UpdateAPIKeyRequest

// NewUpdateAPIKeyRequest instantiates a new UpdateAPIKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAPIKeyRequest() *UpdateAPIKeyRequest {
	this := UpdateAPIKeyRequest{}
	return &this
}

// NewUpdateAPIKeyRequestWithDefaults instantiates a new UpdateAPIKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAPIKeyRequestWithDefaults() *UpdateAPIKeyRequest {
	this := UpdateAPIKeyRequest{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *UpdateAPIKeyRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAPIKeyRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *UpdateAPIKeyRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *UpdateAPIKeyRequest) SetLabel(v string) {
	o.Label = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *UpdateAPIKeyRequest) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAPIKeyRequest) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *UpdateAPIKeyRequest) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *UpdateAPIKeyRequest) SetScopes(v []string) {
	o.Scopes = v
}

func (o UpdateAPIKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAPIKeyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateAPIKeyRequest) UnmarshalJSON(data []byte) (err error) {
	varUpdateAPIKeyRequest := _UpdateAPIKeyRequest{}

	err = json.Unmarshal(data, &varUpdateAPIKeyRequest)

	if err != nil {
		return err
	}

	*o = UpdateAPIKeyRequest(varUpdateAPIKeyRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "scopes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateAPIKeyRequest struct {
	value *UpdateAPIKeyRequest
	isSet bool
}

func (v NullableUpdateAPIKeyRequest) Get() *UpdateAPIKeyRequest {
	return v.value
}

func (v *NullableUpdateAPIKeyRequest) Set(val *UpdateAPIKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAPIKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAPIKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAPIKeyRequest(val *UpdateAPIKeyRequest) *NullableUpdateAPIKeyRequest {
	return &NullableUpdateAPIKeyRequest{value: val, isSet: true}
}

func (v NullableUpdateAPIKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAPIKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
