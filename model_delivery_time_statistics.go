package ahasend

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the DeliveryTimeStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryTimeStatistics{}

// DeliveryTimeStatistics struct for DeliveryTimeStatistics
type DeliveryTimeStatistics struct {
	// Start time of the statistics bucket
	FromTimestamp time.Time `json:"from_timestamp"`
	// End time of the statistics bucket
	ToTimestamp time.Time `json:"to_timestamp"`
	// Average delivery time in seconds
	AvgDeliveryTime float64 `json:"avg_delivery_time"`
	// Number of messages
	DeliveredCount       int32          `json:"delivered_count"`
	DeliveryTimes        []DeliveryTime `json:"delivery_times,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeliveryTimeStatistics DeliveryTimeStatistics

// NewDeliveryTimeStatistics instantiates a new DeliveryTimeStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryTimeStatistics(fromTimestamp time.Time, toTimestamp time.Time, avgDeliveryTime float64, deliveredCount int32) *DeliveryTimeStatistics {
	this := DeliveryTimeStatistics{}
	this.FromTimestamp = fromTimestamp
	this.ToTimestamp = toTimestamp
	this.AvgDeliveryTime = avgDeliveryTime
	this.DeliveredCount = deliveredCount
	return &this
}

// NewDeliveryTimeStatisticsWithDefaults instantiates a new DeliveryTimeStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryTimeStatisticsWithDefaults() *DeliveryTimeStatistics {
	this := DeliveryTimeStatistics{}
	return &this
}

// GetFromTimestamp returns the FromTimestamp field value
func (o *DeliveryTimeStatistics) GetFromTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.FromTimestamp
}

// GetFromTimestampOk returns a tuple with the FromTimestamp field value
// and a boolean to check if the value has been set.
func (o *DeliveryTimeStatistics) GetFromTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromTimestamp, true
}

// SetFromTimestamp sets field value
func (o *DeliveryTimeStatistics) SetFromTimestamp(v time.Time) {
	o.FromTimestamp = v
}

// GetToTimestamp returns the ToTimestamp field value
func (o *DeliveryTimeStatistics) GetToTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ToTimestamp
}

// GetToTimestampOk returns a tuple with the ToTimestamp field value
// and a boolean to check if the value has been set.
func (o *DeliveryTimeStatistics) GetToTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToTimestamp, true
}

// SetToTimestamp sets field value
func (o *DeliveryTimeStatistics) SetToTimestamp(v time.Time) {
	o.ToTimestamp = v
}

// GetAvgDeliveryTime returns the AvgDeliveryTime field value
func (o *DeliveryTimeStatistics) GetAvgDeliveryTime() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AvgDeliveryTime
}

// GetAvgDeliveryTimeOk returns a tuple with the AvgDeliveryTime field value
// and a boolean to check if the value has been set.
func (o *DeliveryTimeStatistics) GetAvgDeliveryTimeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgDeliveryTime, true
}

// SetAvgDeliveryTime sets field value
func (o *DeliveryTimeStatistics) SetAvgDeliveryTime(v float64) {
	o.AvgDeliveryTime = v
}

// GetDeliveredCount returns the DeliveredCount field value
func (o *DeliveryTimeStatistics) GetDeliveredCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DeliveredCount
}

// GetDeliveredCountOk returns a tuple with the DeliveredCount field value
// and a boolean to check if the value has been set.
func (o *DeliveryTimeStatistics) GetDeliveredCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliveredCount, true
}

// SetDeliveredCount sets field value
func (o *DeliveryTimeStatistics) SetDeliveredCount(v int32) {
	o.DeliveredCount = v
}

// GetDeliveryTimes returns the DeliveryTimes field value if set, zero value otherwise.
func (o *DeliveryTimeStatistics) GetDeliveryTimes() []DeliveryTime {
	if o == nil || IsNil(o.DeliveryTimes) {
		var ret []DeliveryTime
		return ret
	}
	return o.DeliveryTimes
}

// GetDeliveryTimesOk returns a tuple with the DeliveryTimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryTimeStatistics) GetDeliveryTimesOk() ([]DeliveryTime, bool) {
	if o == nil || IsNil(o.DeliveryTimes) {
		return nil, false
	}
	return o.DeliveryTimes, true
}

// HasDeliveryTimes returns a boolean if a field has been set.
func (o *DeliveryTimeStatistics) HasDeliveryTimes() bool {
	if o != nil && !IsNil(o.DeliveryTimes) {
		return true
	}

	return false
}

// SetDeliveryTimes gets a reference to the given []DeliveryTime and assigns it to the DeliveryTimes field.
func (o *DeliveryTimeStatistics) SetDeliveryTimes(v []DeliveryTime) {
	o.DeliveryTimes = v
}

func (o DeliveryTimeStatistics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryTimeStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["from_timestamp"] = o.FromTimestamp
	toSerialize["to_timestamp"] = o.ToTimestamp
	toSerialize["avg_delivery_time"] = o.AvgDeliveryTime
	toSerialize["delivered_count"] = o.DeliveredCount
	if !IsNil(o.DeliveryTimes) {
		toSerialize["delivery_times"] = o.DeliveryTimes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeliveryTimeStatistics) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"from_timestamp",
		"to_timestamp",
		"avg_delivery_time",
		"delivered_count",
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

	varDeliveryTimeStatistics := _DeliveryTimeStatistics{}

	err = json.Unmarshal(data, &varDeliveryTimeStatistics)

	if err != nil {
		return err
	}

	*o = DeliveryTimeStatistics(varDeliveryTimeStatistics)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "from_timestamp")
		delete(additionalProperties, "to_timestamp")
		delete(additionalProperties, "avg_delivery_time")
		delete(additionalProperties, "delivered_count")
		delete(additionalProperties, "delivery_times")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeliveryTimeStatistics struct {
	value *DeliveryTimeStatistics
	isSet bool
}

func (v NullableDeliveryTimeStatistics) Get() *DeliveryTimeStatistics {
	return v.value
}

func (v *NullableDeliveryTimeStatistics) Set(val *DeliveryTimeStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryTimeStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryTimeStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryTimeStatistics(val *DeliveryTimeStatistics) *NullableDeliveryTimeStatistics {
	return &NullableDeliveryTimeStatistics{value: val, isSet: true}
}

func (v NullableDeliveryTimeStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryTimeStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
