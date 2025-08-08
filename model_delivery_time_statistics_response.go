package ahasend

import (
	"encoding/json"
	"fmt"
)

// checks if the DeliveryTimeStatisticsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryTimeStatisticsResponse{}

// DeliveryTimeStatisticsResponse struct for DeliveryTimeStatisticsResponse
type DeliveryTimeStatisticsResponse struct {
	// Object type identifier
	Object string `json:"object"`
	// Array of delivery time statistics
	Data                 []DeliveryTimeStatistics `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _DeliveryTimeStatisticsResponse DeliveryTimeStatisticsResponse

// NewDeliveryTimeStatisticsResponse instantiates a new DeliveryTimeStatisticsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryTimeStatisticsResponse(object string, data []DeliveryTimeStatistics) *DeliveryTimeStatisticsResponse {
	this := DeliveryTimeStatisticsResponse{}
	this.Object = object
	this.Data = data
	return &this
}

// NewDeliveryTimeStatisticsResponseWithDefaults instantiates a new DeliveryTimeStatisticsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryTimeStatisticsResponseWithDefaults() *DeliveryTimeStatisticsResponse {
	this := DeliveryTimeStatisticsResponse{}
	return &this
}

// GetObject returns the Object field value
func (o *DeliveryTimeStatisticsResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *DeliveryTimeStatisticsResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *DeliveryTimeStatisticsResponse) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *DeliveryTimeStatisticsResponse) GetData() []DeliveryTimeStatistics {
	if o == nil {
		var ret []DeliveryTimeStatistics
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeliveryTimeStatisticsResponse) GetDataOk() ([]DeliveryTimeStatistics, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *DeliveryTimeStatisticsResponse) SetData(v []DeliveryTimeStatistics) {
	o.Data = v
}

func (o DeliveryTimeStatisticsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryTimeStatisticsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeliveryTimeStatisticsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varDeliveryTimeStatisticsResponse := _DeliveryTimeStatisticsResponse{}

	err = json.Unmarshal(data, &varDeliveryTimeStatisticsResponse)

	if err != nil {
		return err
	}

	*o = DeliveryTimeStatisticsResponse(varDeliveryTimeStatisticsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeliveryTimeStatisticsResponse struct {
	value *DeliveryTimeStatisticsResponse
	isSet bool
}

func (v NullableDeliveryTimeStatisticsResponse) Get() *DeliveryTimeStatisticsResponse {
	return v.value
}

func (v *NullableDeliveryTimeStatisticsResponse) Set(val *DeliveryTimeStatisticsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryTimeStatisticsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryTimeStatisticsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryTimeStatisticsResponse(val *DeliveryTimeStatisticsResponse) *NullableDeliveryTimeStatisticsResponse {
	return &NullableDeliveryTimeStatisticsResponse{value: val, isSet: true}
}

func (v NullableDeliveryTimeStatisticsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryTimeStatisticsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
