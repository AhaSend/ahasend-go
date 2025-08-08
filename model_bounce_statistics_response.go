package ahasend

import (
	"encoding/json"
	"fmt"
)

// checks if the BounceStatisticsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BounceStatisticsResponse{}

// BounceStatisticsResponse struct for BounceStatisticsResponse
type BounceStatisticsResponse struct {
	// Object type identifier
	Object string `json:"object"`
	// Array of bounce statistics
	Data                 []BounceStatistics `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _BounceStatisticsResponse BounceStatisticsResponse

// NewBounceStatisticsResponse instantiates a new BounceStatisticsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBounceStatisticsResponse(object string, data []BounceStatistics) *BounceStatisticsResponse {
	this := BounceStatisticsResponse{}
	this.Object = object
	this.Data = data
	return &this
}

// NewBounceStatisticsResponseWithDefaults instantiates a new BounceStatisticsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBounceStatisticsResponseWithDefaults() *BounceStatisticsResponse {
	this := BounceStatisticsResponse{}
	return &this
}

// GetObject returns the Object field value
func (o *BounceStatisticsResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *BounceStatisticsResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *BounceStatisticsResponse) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *BounceStatisticsResponse) GetData() []BounceStatistics {
	if o == nil {
		var ret []BounceStatistics
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BounceStatisticsResponse) GetDataOk() ([]BounceStatistics, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BounceStatisticsResponse) SetData(v []BounceStatistics) {
	o.Data = v
}

func (o BounceStatisticsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BounceStatisticsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BounceStatisticsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varBounceStatisticsResponse := _BounceStatisticsResponse{}

	err = json.Unmarshal(data, &varBounceStatisticsResponse)

	if err != nil {
		return err
	}

	*o = BounceStatisticsResponse(varBounceStatisticsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBounceStatisticsResponse struct {
	value *BounceStatisticsResponse
	isSet bool
}

func (v NullableBounceStatisticsResponse) Get() *BounceStatisticsResponse {
	return v.value
}

func (v *NullableBounceStatisticsResponse) Set(val *BounceStatisticsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBounceStatisticsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBounceStatisticsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBounceStatisticsResponse(val *BounceStatisticsResponse) *NullableBounceStatisticsResponse {
	return &NullableBounceStatisticsResponse{value: val, isSet: true}
}

func (v NullableBounceStatisticsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBounceStatisticsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
