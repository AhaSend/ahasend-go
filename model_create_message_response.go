package ahasend

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateMessageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMessageResponse{}

// CreateMessageResponse struct for CreateMessageResponse
type CreateMessageResponse struct {
	// Object type identifier
	Object string `json:"object"`
	// List of messages and their statuses
	Data                 []CreateSingleMessageResponse `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _CreateMessageResponse CreateMessageResponse

// NewCreateMessageResponse instantiates a new CreateMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMessageResponse(object string, data []CreateSingleMessageResponse) *CreateMessageResponse {
	this := CreateMessageResponse{}
	this.Object = object
	this.Data = data
	return &this
}

// NewCreateMessageResponseWithDefaults instantiates a new CreateMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMessageResponseWithDefaults() *CreateMessageResponse {
	this := CreateMessageResponse{}
	return &this
}

// GetObject returns the Object field value
func (o *CreateMessageResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *CreateMessageResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *CreateMessageResponse) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *CreateMessageResponse) GetData() []CreateSingleMessageResponse {
	if o == nil {
		var ret []CreateSingleMessageResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateMessageResponse) GetDataOk() ([]CreateSingleMessageResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *CreateMessageResponse) SetData(v []CreateSingleMessageResponse) {
	o.Data = v
}

func (o CreateMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMessageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateMessageResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCreateMessageResponse := _CreateMessageResponse{}

	err = json.Unmarshal(data, &varCreateMessageResponse)

	if err != nil {
		return err
	}

	*o = CreateMessageResponse(varCreateMessageResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateMessageResponse struct {
	value *CreateMessageResponse
	isSet bool
}

func (v NullableCreateMessageResponse) Get() *CreateMessageResponse {
	return v.value
}

func (v *NullableCreateMessageResponse) Set(val *CreateMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMessageResponse(val *CreateMessageResponse) *NullableCreateMessageResponse {
	return &NullableCreateMessageResponse{value: val, isSet: true}
}

func (v NullableCreateMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
