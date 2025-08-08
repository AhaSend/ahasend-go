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
	// Time bucket for the statistics
	TimeBucket time.Time `json:"time_bucket"`
	// Bounce classification
	Classification string `json:"classification"`
	// Number of bounces
	Count                int32 `json:"count"`
	AdditionalProperties map[string]interface{}
}

type _BounceStatistics BounceStatistics

// NewBounceStatistics instantiates a new BounceStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBounceStatistics(timeBucket time.Time, classification string, count int32) *BounceStatistics {
	this := BounceStatistics{}
	this.TimeBucket = timeBucket
	this.Classification = classification
	this.Count = count
	return &this
}

// NewBounceStatisticsWithDefaults instantiates a new BounceStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBounceStatisticsWithDefaults() *BounceStatistics {
	this := BounceStatistics{}
	return &this
}

// GetTimeBucket returns the TimeBucket field value
func (o *BounceStatistics) GetTimeBucket() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TimeBucket
}

// GetTimeBucketOk returns a tuple with the TimeBucket field value
// and a boolean to check if the value has been set.
func (o *BounceStatistics) GetTimeBucketOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeBucket, true
}

// SetTimeBucket sets field value
func (o *BounceStatistics) SetTimeBucket(v time.Time) {
	o.TimeBucket = v
}

// GetClassification returns the Classification field value
func (o *BounceStatistics) GetClassification() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Classification
}

// GetClassificationOk returns a tuple with the Classification field value
// and a boolean to check if the value has been set.
func (o *BounceStatistics) GetClassificationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Classification, true
}

// SetClassification sets field value
func (o *BounceStatistics) SetClassification(v string) {
	o.Classification = v
}

// GetCount returns the Count field value
func (o *BounceStatistics) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *BounceStatistics) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *BounceStatistics) SetCount(v int32) {
	o.Count = v
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
	toSerialize["time_bucket"] = o.TimeBucket
	toSerialize["classification"] = o.Classification
	toSerialize["count"] = o.Count

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
		"time_bucket",
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

	varBounceStatistics := _BounceStatistics{}

	err = json.Unmarshal(data, &varBounceStatistics)

	if err != nil {
		return err
	}

	*o = BounceStatistics(varBounceStatistics)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "time_bucket")
		delete(additionalProperties, "classification")
		delete(additionalProperties, "count")
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
