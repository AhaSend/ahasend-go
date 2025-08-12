package ahasend

import (
	"encoding/json"
	"fmt"
)

// checks if the SenderAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SenderAddress{}

// SenderAddress struct for SenderAddress
type SenderAddress struct {
	// Valid email address from a domain defined in your account with valid DNS records
	Email string `json:"email"`
	// Display name for the sender
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SenderAddress SenderAddress

// NewSenderAddress instantiates a new SenderAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSenderAddress(email string) *SenderAddress {
	this := SenderAddress{}
	this.Email = email
	return &this
}

// NewSenderAddressWithDefaults instantiates a new SenderAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSenderAddressWithDefaults() *SenderAddress {
	this := SenderAddress{}
	return &this
}

// GetEmail returns the Email field value
func (o *SenderAddress) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SenderAddress) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SenderAddress) SetEmail(v string) {
	o.Email = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SenderAddress) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SenderAddress) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SenderAddress) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SenderAddress) SetName(v string) {
	o.Name = &v
}

func (o SenderAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SenderAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SenderAddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varSenderAddress := _SenderAddress{}

	err = json.Unmarshal(data, &varSenderAddress)

	if err != nil {
		return err
	}

	*o = SenderAddress(varSenderAddress)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSenderAddress struct {
	value *SenderAddress
	isSet bool
}

func (v NullableSenderAddress) Get() *SenderAddress {
	return v.value
}

func (v *NullableSenderAddress) Set(val *SenderAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableSenderAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableSenderAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSenderAddress(val *SenderAddress) *NullableSenderAddress {
	return &NullableSenderAddress{value: val, isSet: true}
}

func (v NullableSenderAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSenderAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
