package ahasend

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateAPIKeyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAPIKeyRequest{}

// CreateAPIKeyRequest struct for CreateAPIKeyRequest
type CreateAPIKeyRequest struct {
	// Human-readable label for the API key
	Label string `json:"label"`
	// Array of scope strings to grant to this API key
	Scopes               []string `json:"scopes"`
	AdditionalProperties map[string]interface{}
}

type _CreateAPIKeyRequest CreateAPIKeyRequest

// NewCreateAPIKeyRequest instantiates a new CreateAPIKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAPIKeyRequest(label string, scopes []string) *CreateAPIKeyRequest {
	this := CreateAPIKeyRequest{}
	this.Label = label
	this.Scopes = scopes
	return &this
}

// NewCreateAPIKeyRequestWithDefaults instantiates a new CreateAPIKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAPIKeyRequestWithDefaults() *CreateAPIKeyRequest {
	this := CreateAPIKeyRequest{}
	return &this
}

// GetLabel returns the Label field value
func (o *CreateAPIKeyRequest) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *CreateAPIKeyRequest) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *CreateAPIKeyRequest) SetLabel(v string) {
	o.Label = v
}

// GetScopes returns the Scopes field value
func (o *CreateAPIKeyRequest) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *CreateAPIKeyRequest) GetScopesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *CreateAPIKeyRequest) SetScopes(v []string) {
	o.Scopes = v
}

func (o CreateAPIKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAPIKeyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["scopes"] = o.Scopes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateAPIKeyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
		"scopes",
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

	varCreateAPIKeyRequest := _CreateAPIKeyRequest{}

	err = json.Unmarshal(data, &varCreateAPIKeyRequest)

	if err != nil {
		return err
	}

	*o = CreateAPIKeyRequest(varCreateAPIKeyRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "scopes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateAPIKeyRequest struct {
	value *CreateAPIKeyRequest
	isSet bool
}

func (v NullableCreateAPIKeyRequest) Get() *CreateAPIKeyRequest {
	return v.value
}

func (v *NullableCreateAPIKeyRequest) Set(val *CreateAPIKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAPIKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAPIKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAPIKeyRequest(val *CreateAPIKeyRequest) *NullableCreateAPIKeyRequest {
	return &NullableCreateAPIKeyRequest{value: val, isSet: true}
}

func (v NullableCreateAPIKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAPIKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
