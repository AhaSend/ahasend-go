package ahasend

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the DeliveryEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryEvent{}

// DeliveryEvent struct for DeliveryEvent
type DeliveryEvent struct {
	// Timestamp of the delivery event
	Time time.Time `json:"time"`
	// Log message for the delivery event
	Log string `json:"log"`
	// Status of the delivery event
	Status               string `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _DeliveryEvent DeliveryEvent

// NewDeliveryEvent instantiates a new DeliveryEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryEvent(time time.Time, log string, status string) *DeliveryEvent {
	this := DeliveryEvent{}
	this.Time = time
	this.Log = log
	this.Status = status
	return &this
}

// NewDeliveryEventWithDefaults instantiates a new DeliveryEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryEventWithDefaults() *DeliveryEvent {
	this := DeliveryEvent{}
	return &this
}

// GetTime returns the Time field value
func (o *DeliveryEvent) GetTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *DeliveryEvent) GetTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *DeliveryEvent) SetTime(v time.Time) {
	o.Time = v
}

// GetLog returns the Log field value
func (o *DeliveryEvent) GetLog() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Log
}

// GetLogOk returns a tuple with the Log field value
// and a boolean to check if the value has been set.
func (o *DeliveryEvent) GetLogOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Log, true
}

// SetLog sets field value
func (o *DeliveryEvent) SetLog(v string) {
	o.Log = v
}

// GetStatus returns the Status field value
func (o *DeliveryEvent) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DeliveryEvent) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DeliveryEvent) SetStatus(v string) {
	o.Status = v
}

func (o DeliveryEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["time"] = o.Time
	toSerialize["log"] = o.Log
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeliveryEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"time",
		"log",
		"status",
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

	varDeliveryEvent := _DeliveryEvent{}

	err = json.Unmarshal(data, &varDeliveryEvent)

	if err != nil {
		return err
	}

	*o = DeliveryEvent(varDeliveryEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "time")
		delete(additionalProperties, "log")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeliveryEvent struct {
	value *DeliveryEvent
	isSet bool
}

func (v NullableDeliveryEvent) Get() *DeliveryEvent {
	return v.value
}

func (v *NullableDeliveryEvent) Set(val *DeliveryEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryEvent(val *DeliveryEvent) *NullableDeliveryEvent {
	return &NullableDeliveryEvent{value: val, isSet: true}
}

func (v NullableDeliveryEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
