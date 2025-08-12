package ahasend

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the BounceStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BounceStatistics{}

// BounceStatistics struct for BounceStatistics
type BounceStatistics struct {
	// Start time of the statistics bucket
	FromTimestamp time.Time `json:"from_timestamp"`
	// End time of the statistics bucket
	ToTimestamp time.Time `json:"to_timestamp"`
	// Bounce count per bounce classification
	Bounces              []Bounce `json:"bounces"`
	AdditionalProperties map[string]interface{}
}

type _BounceStatistics BounceStatistics

// NewBounceStatistics instantiates a new BounceStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBounceStatistics(fromTimestamp time.Time, toTimestamp time.Time, bounces []Bounce) *BounceStatistics {
	this := BounceStatistics{}
	this.FromTimestamp = fromTimestamp
	this.ToTimestamp = toTimestamp
	this.Bounces = bounces
	return &this
}

// NewBounceStatisticsWithDefaults instantiates a new BounceStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBounceStatisticsWithDefaults() *BounceStatistics {
	this := BounceStatistics{}
	return &this
}

// GetFromTimestamp returns the FromTimestamp field value
func (o *BounceStatistics) GetFromTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.FromTimestamp
}

// GetFromTimestampOk returns a tuple with the FromTimestamp field value
// and a boolean to check if the value has been set.
func (o *BounceStatistics) GetFromTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromTimestamp, true
}

// SetFromTimestamp sets field value
func (o *BounceStatistics) SetFromTimestamp(v time.Time) {
	o.FromTimestamp = v
}

// GetToTimestamp returns the ToTimestamp field value
func (o *BounceStatistics) GetToTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ToTimestamp
}

// GetToTimestampOk returns a tuple with the ToTimestamp field value
// and a boolean to check if the value has been set.
func (o *BounceStatistics) GetToTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToTimestamp, true
}

// SetToTimestamp sets field value
func (o *BounceStatistics) SetToTimestamp(v time.Time) {
	o.ToTimestamp = v
}

// GetBounces returns the Bounces field value
func (o *BounceStatistics) GetBounces() []Bounce {
	if o == nil {
		var ret []Bounce
		return ret
	}

	return o.Bounces
}

// GetBouncesOk returns a tuple with the Bounces field value
// and a boolean to check if the value has been set.
func (o *BounceStatistics) GetBouncesOk() ([]Bounce, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bounces, true
}

// SetBounces sets field value
func (o *BounceStatistics) SetBounces(v []Bounce) {
	o.Bounces = v
}

func (o BounceStatistics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BounceStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["from_timestamp"] = o.FromTimestamp
	toSerialize["to_timestamp"] = o.ToTimestamp
	toSerialize["bounces"] = o.Bounces

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BounceStatistics) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"from_timestamp",
		"to_timestamp",
		"bounces",
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

	varBounceStatistics := _BounceStatistics{}

	err = json.Unmarshal(data, &varBounceStatistics)

	if err != nil {
		return err
	}

	*o = BounceStatistics(varBounceStatistics)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "from_timestamp")
		delete(additionalProperties, "to_timestamp")
		delete(additionalProperties, "bounces")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBounceStatistics struct {
	value *BounceStatistics
	isSet bool
}

func (v NullableBounceStatistics) Get() *BounceStatistics {
	return v.value
}

func (v *NullableBounceStatistics) Set(val *BounceStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableBounceStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableBounceStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBounceStatistics(val *BounceStatistics) *NullableBounceStatistics {
	return &NullableBounceStatistics{value: val, isSet: true}
}

func (v NullableBounceStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBounceStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
