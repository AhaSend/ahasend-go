package ahasend

import (
	"encoding/json"
	"fmt"
)

// checks if the AddMemberRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddMemberRequest{}

// AddMemberRequest struct for AddMemberRequest
type AddMemberRequest struct {
	// Email address of the user to add
	Email string `json:"email"`
	// Display name for the user
	Name *string `json:"name,omitempty"`
	// Role to assign to the user
	Role                 string `json:"role"`
	AdditionalProperties map[string]interface{}
}

type _AddMemberRequest AddMemberRequest

// NewAddMemberRequest instantiates a new AddMemberRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddMemberRequest(email string, role string) *AddMemberRequest {
	this := AddMemberRequest{}
	this.Email = email
	this.Role = role
	return &this
}

// NewAddMemberRequestWithDefaults instantiates a new AddMemberRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddMemberRequestWithDefaults() *AddMemberRequest {
	this := AddMemberRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *AddMemberRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *AddMemberRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *AddMemberRequest) SetEmail(v string) {
	o.Email = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AddMemberRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddMemberRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AddMemberRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AddMemberRequest) SetName(v string) {
	o.Name = &v
}

// GetRole returns the Role field value
func (o *AddMemberRequest) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *AddMemberRequest) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *AddMemberRequest) SetRole(v string) {
	o.Role = v
}

func (o AddMemberRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddMemberRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["role"] = o.Role

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AddMemberRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"role",
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

	varAddMemberRequest := _AddMemberRequest{}

	err = json.Unmarshal(data, &varAddMemberRequest)

	if err != nil {
		return err
	}

	*o = AddMemberRequest(varAddMemberRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "name")
		delete(additionalProperties, "role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAddMemberRequest struct {
	value *AddMemberRequest
	isSet bool
}

func (v NullableAddMemberRequest) Get() *AddMemberRequest {
	return v.value
}

func (v *NullableAddMemberRequest) Set(val *AddMemberRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddMemberRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddMemberRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddMemberRequest(val *AddMemberRequest) *NullableAddMemberRequest {
	return &NullableAddMemberRequest{value: val, isSet: true}
}

func (v NullableAddMemberRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddMemberRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
