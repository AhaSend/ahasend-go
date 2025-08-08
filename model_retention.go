package ahasend

import (
	"encoding/json"
)

// checks if the Retention type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Retention{}

// Retention struct for Retention
type Retention struct {
	// Number of days to retain metadata
	Metadata *int32 `json:"metadata,omitempty"`
	// Number of days to retain data
	Data                 *int32 `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Retention Retention

// NewRetention instantiates a new Retention object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetention() *Retention {
	this := Retention{}
	return &this
}

// NewRetentionWithDefaults instantiates a new Retention object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetentionWithDefaults() *Retention {
	this := Retention{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Retention) GetMetadata() int32 {
	if o == nil || IsNil(o.Metadata) {
		var ret int32
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Retention) GetMetadataOk() (*int32, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Retention) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given int32 and assigns it to the Metadata field.
func (o *Retention) SetMetadata(v int32) {
	o.Metadata = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Retention) GetData() int32 {
	if o == nil || IsNil(o.Data) {
		var ret int32
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Retention) GetDataOk() (*int32, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Retention) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given int32 and assigns it to the Data field.
func (o *Retention) SetData(v int32) {
	o.Data = &v
}

func (o Retention) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Retention) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Retention) UnmarshalJSON(data []byte) (err error) {
	varRetention := _Retention{}

	err = json.Unmarshal(data, &varRetention)

	if err != nil {
		return err
	}

	*o = Retention(varRetention)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRetention struct {
	value *Retention
	isSet bool
}

func (v NullableRetention) Get() *Retention {
	return v.value
}

func (v *NullableRetention) Set(val *Retention) {
	v.value = val
	v.isSet = true
}

func (v NullableRetention) IsSet() bool {
	return v.isSet
}

func (v *NullableRetention) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetention(val *Retention) *NullableRetention {
	return &NullableRetention{value: val, isSet: true}
}

func (v NullableRetention) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetention) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
