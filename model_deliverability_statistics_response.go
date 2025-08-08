package ahasend

import (
	"encoding/json"
	"fmt"
)

// checks if the DeliverabilityStatisticsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliverabilityStatisticsResponse{}

// DeliverabilityStatisticsResponse struct for DeliverabilityStatisticsResponse
type DeliverabilityStatisticsResponse struct {
	// Object type identifier
	Object string `json:"object"`
	// Array of deliverability statistics
	Data                 []DeliverabilityStatistics `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _DeliverabilityStatisticsResponse DeliverabilityStatisticsResponse

// NewDeliverabilityStatisticsResponse instantiates a new DeliverabilityStatisticsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliverabilityStatisticsResponse(object string, data []DeliverabilityStatistics) *DeliverabilityStatisticsResponse {
	this := DeliverabilityStatisticsResponse{}
	this.Object = object
	this.Data = data
	return &this
}

// NewDeliverabilityStatisticsResponseWithDefaults instantiates a new DeliverabilityStatisticsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliverabilityStatisticsResponseWithDefaults() *DeliverabilityStatisticsResponse {
	this := DeliverabilityStatisticsResponse{}
	return &this
}

// GetObject returns the Object field value
func (o *DeliverabilityStatisticsResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *DeliverabilityStatisticsResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *DeliverabilityStatisticsResponse) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *DeliverabilityStatisticsResponse) GetData() []DeliverabilityStatistics {
	if o == nil {
		var ret []DeliverabilityStatistics
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeliverabilityStatisticsResponse) GetDataOk() ([]DeliverabilityStatistics, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *DeliverabilityStatisticsResponse) SetData(v []DeliverabilityStatistics) {
	o.Data = v
}

func (o DeliverabilityStatisticsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliverabilityStatisticsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeliverabilityStatisticsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varDeliverabilityStatisticsResponse := _DeliverabilityStatisticsResponse{}

	err = json.Unmarshal(data, &varDeliverabilityStatisticsResponse)

	if err != nil {
		return err
	}

	*o = DeliverabilityStatisticsResponse(varDeliverabilityStatisticsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeliverabilityStatisticsResponse struct {
	value *DeliverabilityStatisticsResponse
	isSet bool
}

func (v NullableDeliverabilityStatisticsResponse) Get() *DeliverabilityStatisticsResponse {
	return v.value
}

func (v *NullableDeliverabilityStatisticsResponse) Set(val *DeliverabilityStatisticsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliverabilityStatisticsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliverabilityStatisticsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliverabilityStatisticsResponse(val *DeliverabilityStatisticsResponse) *NullableDeliverabilityStatisticsResponse {
	return &NullableDeliverabilityStatisticsResponse{value: val, isSet: true}
}

func (v NullableDeliverabilityStatisticsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliverabilityStatisticsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
