package ahasend

import (
	"encoding/json"
	"fmt"
)

// checks if the AccountMembersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountMembersResponse{}

// AccountMembersResponse struct for AccountMembersResponse
type AccountMembersResponse struct {
	// Object type identifier
	Object string `json:"object"`
	// Array of account members
	Data                 []UserAccount `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _AccountMembersResponse AccountMembersResponse

// NewAccountMembersResponse instantiates a new AccountMembersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountMembersResponse(object string, data []UserAccount) *AccountMembersResponse {
	this := AccountMembersResponse{}
	this.Object = object
	this.Data = data
	return &this
}

// NewAccountMembersResponseWithDefaults instantiates a new AccountMembersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountMembersResponseWithDefaults() *AccountMembersResponse {
	this := AccountMembersResponse{}
	return &this
}

// GetObject returns the Object field value
func (o *AccountMembersResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *AccountMembersResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *AccountMembersResponse) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *AccountMembersResponse) GetData() []UserAccount {
	if o == nil {
		var ret []UserAccount
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AccountMembersResponse) GetDataOk() ([]UserAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AccountMembersResponse) SetData(v []UserAccount) {
	o.Data = v
}

func (o AccountMembersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountMembersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountMembersResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"data",
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

	varAccountMembersResponse := _AccountMembersResponse{}

	err = json.Unmarshal(data, &varAccountMembersResponse)

	if err != nil {
		return err
	}

	*o = AccountMembersResponse(varAccountMembersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountMembersResponse struct {
	value *AccountMembersResponse
	isSet bool
}

func (v NullableAccountMembersResponse) Get() *AccountMembersResponse {
	return v.value
}

func (v *NullableAccountMembersResponse) Set(val *AccountMembersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountMembersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountMembersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountMembersResponse(val *AccountMembersResponse) *NullableAccountMembersResponse {
	return &NullableAccountMembersResponse{value: val, isSet: true}
}

func (v NullableAccountMembersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountMembersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
