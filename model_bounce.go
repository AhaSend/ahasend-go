package ahasend

import (
	"encoding/json"
	"fmt"
)

// checks if the Bounce type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Bounce{}

// Bounce struct for Bounce
type Bounce struct {
	// Bounce classification
	Classification string `json:"classification"`
	// Number of bounces
	Count                int32 `json:"count"`
	AdditionalProperties map[string]interface{}
}

type _Bounce Bounce

// NewBounce instantiates a new Bounce object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBounce(classification string, count int32) *Bounce {
	this := Bounce{}
	this.Classification = classification
	this.Count = count
	return &this
}

// NewBounceWithDefaults instantiates a new Bounce object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBounceWithDefaults() *Bounce {
	this := Bounce{}
	return &this
}

// GetClassification returns the Classification field value
func (o *Bounce) GetClassification() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Classification
}

// GetClassificationOk returns a tuple with the Classification field value
// and a boolean to check if the value has been set.
func (o *Bounce) GetClassificationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Classification, true
}

// SetClassification sets field value
func (o *Bounce) SetClassification(v string) {
	o.Classification = v
}

// GetCount returns the Count field value
func (o *Bounce) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *Bounce) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *Bounce) SetCount(v int32) {
	o.Count = v
}

func (o Bounce) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Bounce) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["classification"] = o.Classification
	toSerialize["count"] = o.Count

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Bounce) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"classification",
		"count",
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

	varBounce := _Bounce{}

	err = json.Unmarshal(data, &varBounce)

	if err != nil {
		return err
	}

	*o = Bounce(varBounce)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "classification")
		delete(additionalProperties, "count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBounce struct {
	value *Bounce
	isSet bool
}

func (v NullableBounce) Get() *Bounce {
	return v.value
}

func (v *NullableBounce) Set(val *Bounce) {
	v.value = val
	v.isSet = true
}

func (v NullableBounce) IsSet() bool {
	return v.isSet
}

func (v *NullableBounce) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBounce(val *Bounce) *NullableBounce {
	return &NullableBounce{value: val, isSet: true}
}

func (v NullableBounce) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBounce) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
