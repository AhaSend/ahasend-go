package ahasend

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the DeliverabilityStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliverabilityStatistics{}

// DeliverabilityStatistics struct for DeliverabilityStatistics
type DeliverabilityStatistics struct {
	// Start time of the statistics bucket
	FromTimestamp time.Time `json:"from_timestamp"`
	// End time of the statistics bucket
	ToTimestamp time.Time `json:"to_timestamp"`
	// Number of messages accepted for delivery
	ReceptionCount *int32 `json:"reception_count,omitempty"`
	// Number of messages delivered
	DeliveredCount *int32 `json:"delivered_count,omitempty"`
	// Number of messages deferred
	DeferredCount *int32 `json:"deferred_count,omitempty"`
	// Number of messages bounced
	BouncedCount *int32 `json:"bounced_count,omitempty"`
	// Number of messages failed
	FailedCount *int32 `json:"failed_count,omitempty"`
	// Number of messages suppressed
	SuppressedCount *int32 `json:"suppressed_count,omitempty"`
	// Number of messages opened at least once
	OpenedCount *int32 `json:"opened_count,omitempty"`
	// Number of messages that have at least one link in them clicked.
	ClickedCount         *int32 `json:"clicked_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeliverabilityStatistics DeliverabilityStatistics

// NewDeliverabilityStatistics instantiates a new DeliverabilityStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliverabilityStatistics(fromTimestamp time.Time, toTimestamp time.Time) *DeliverabilityStatistics {
	this := DeliverabilityStatistics{}
	this.FromTimestamp = fromTimestamp
	this.ToTimestamp = toTimestamp
	return &this
}

// NewDeliverabilityStatisticsWithDefaults instantiates a new DeliverabilityStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliverabilityStatisticsWithDefaults() *DeliverabilityStatistics {
	this := DeliverabilityStatistics{}
	return &this
}

// GetFromTimestamp returns the FromTimestamp field value
func (o *DeliverabilityStatistics) GetFromTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.FromTimestamp
}

// GetFromTimestampOk returns a tuple with the FromTimestamp field value
// and a boolean to check if the value has been set.
func (o *DeliverabilityStatistics) GetFromTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromTimestamp, true
}

// SetFromTimestamp sets field value
func (o *DeliverabilityStatistics) SetFromTimestamp(v time.Time) {
	o.FromTimestamp = v
}

// GetToTimestamp returns the ToTimestamp field value
func (o *DeliverabilityStatistics) GetToTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ToTimestamp
}

// GetToTimestampOk returns a tuple with the ToTimestamp field value
// and a boolean to check if the value has been set.
func (o *DeliverabilityStatistics) GetToTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToTimestamp, true
}

// SetToTimestamp sets field value
func (o *DeliverabilityStatistics) SetToTimestamp(v time.Time) {
	o.ToTimestamp = v
}

// GetReceptionCount returns the ReceptionCount field value if set, zero value otherwise.
func (o *DeliverabilityStatistics) GetReceptionCount() int32 {
	if o == nil || IsNil(o.ReceptionCount) {
		var ret int32
		return ret
	}
	return *o.ReceptionCount
}

// GetReceptionCountOk returns a tuple with the ReceptionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverabilityStatistics) GetReceptionCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ReceptionCount) {
		return nil, false
	}
	return o.ReceptionCount, true
}

// HasReceptionCount returns a boolean if a field has been set.
func (o *DeliverabilityStatistics) HasReceptionCount() bool {
	if o != nil && !IsNil(o.ReceptionCount) {
		return true
	}

	return false
}

// SetReceptionCount gets a reference to the given int32 and assigns it to the ReceptionCount field.
func (o *DeliverabilityStatistics) SetReceptionCount(v int32) {
	o.ReceptionCount = &v
}

// GetDeliveredCount returns the DeliveredCount field value if set, zero value otherwise.
func (o *DeliverabilityStatistics) GetDeliveredCount() int32 {
	if o == nil || IsNil(o.DeliveredCount) {
		var ret int32
		return ret
	}
	return *o.DeliveredCount
}

// GetDeliveredCountOk returns a tuple with the DeliveredCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverabilityStatistics) GetDeliveredCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DeliveredCount) {
		return nil, false
	}
	return o.DeliveredCount, true
}

// HasDeliveredCount returns a boolean if a field has been set.
func (o *DeliverabilityStatistics) HasDeliveredCount() bool {
	if o != nil && !IsNil(o.DeliveredCount) {
		return true
	}

	return false
}

// SetDeliveredCount gets a reference to the given int32 and assigns it to the DeliveredCount field.
func (o *DeliverabilityStatistics) SetDeliveredCount(v int32) {
	o.DeliveredCount = &v
}

// GetDeferredCount returns the DeferredCount field value if set, zero value otherwise.
func (o *DeliverabilityStatistics) GetDeferredCount() int32 {
	if o == nil || IsNil(o.DeferredCount) {
		var ret int32
		return ret
	}
	return *o.DeferredCount
}

// GetDeferredCountOk returns a tuple with the DeferredCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverabilityStatistics) GetDeferredCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DeferredCount) {
		return nil, false
	}
	return o.DeferredCount, true
}

// HasDeferredCount returns a boolean if a field has been set.
func (o *DeliverabilityStatistics) HasDeferredCount() bool {
	if o != nil && !IsNil(o.DeferredCount) {
		return true
	}

	return false
}

// SetDeferredCount gets a reference to the given int32 and assigns it to the DeferredCount field.
func (o *DeliverabilityStatistics) SetDeferredCount(v int32) {
	o.DeferredCount = &v
}

// GetBouncedCount returns the BouncedCount field value if set, zero value otherwise.
func (o *DeliverabilityStatistics) GetBouncedCount() int32 {
	if o == nil || IsNil(o.BouncedCount) {
		var ret int32
		return ret
	}
	return *o.BouncedCount
}

// GetBouncedCountOk returns a tuple with the BouncedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverabilityStatistics) GetBouncedCountOk() (*int32, bool) {
	if o == nil || IsNil(o.BouncedCount) {
		return nil, false
	}
	return o.BouncedCount, true
}

// HasBouncedCount returns a boolean if a field has been set.
func (o *DeliverabilityStatistics) HasBouncedCount() bool {
	if o != nil && !IsNil(o.BouncedCount) {
		return true
	}

	return false
}

// SetBouncedCount gets a reference to the given int32 and assigns it to the BouncedCount field.
func (o *DeliverabilityStatistics) SetBouncedCount(v int32) {
	o.BouncedCount = &v
}

// GetFailedCount returns the FailedCount field value if set, zero value otherwise.
func (o *DeliverabilityStatistics) GetFailedCount() int32 {
	if o == nil || IsNil(o.FailedCount) {
		var ret int32
		return ret
	}
	return *o.FailedCount
}

// GetFailedCountOk returns a tuple with the FailedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverabilityStatistics) GetFailedCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FailedCount) {
		return nil, false
	}
	return o.FailedCount, true
}

// HasFailedCount returns a boolean if a field has been set.
func (o *DeliverabilityStatistics) HasFailedCount() bool {
	if o != nil && !IsNil(o.FailedCount) {
		return true
	}

	return false
}

// SetFailedCount gets a reference to the given int32 and assigns it to the FailedCount field.
func (o *DeliverabilityStatistics) SetFailedCount(v int32) {
	o.FailedCount = &v
}

// GetSuppressedCount returns the SuppressedCount field value if set, zero value otherwise.
func (o *DeliverabilityStatistics) GetSuppressedCount() int32 {
	if o == nil || IsNil(o.SuppressedCount) {
		var ret int32
		return ret
	}
	return *o.SuppressedCount
}

// GetSuppressedCountOk returns a tuple with the SuppressedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverabilityStatistics) GetSuppressedCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SuppressedCount) {
		return nil, false
	}
	return o.SuppressedCount, true
}

// HasSuppressedCount returns a boolean if a field has been set.
func (o *DeliverabilityStatistics) HasSuppressedCount() bool {
	if o != nil && !IsNil(o.SuppressedCount) {
		return true
	}

	return false
}

// SetSuppressedCount gets a reference to the given int32 and assigns it to the SuppressedCount field.
func (o *DeliverabilityStatistics) SetSuppressedCount(v int32) {
	o.SuppressedCount = &v
}

// GetOpenedCount returns the OpenedCount field value if set, zero value otherwise.
func (o *DeliverabilityStatistics) GetOpenedCount() int32 {
	if o == nil || IsNil(o.OpenedCount) {
		var ret int32
		return ret
	}
	return *o.OpenedCount
}

// GetOpenedCountOk returns a tuple with the OpenedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverabilityStatistics) GetOpenedCountOk() (*int32, bool) {
	if o == nil || IsNil(o.OpenedCount) {
		return nil, false
	}
	return o.OpenedCount, true
}

// HasOpenedCount returns a boolean if a field has been set.
func (o *DeliverabilityStatistics) HasOpenedCount() bool {
	if o != nil && !IsNil(o.OpenedCount) {
		return true
	}

	return false
}

// SetOpenedCount gets a reference to the given int32 and assigns it to the OpenedCount field.
func (o *DeliverabilityStatistics) SetOpenedCount(v int32) {
	o.OpenedCount = &v
}

// GetClickedCount returns the ClickedCount field value if set, zero value otherwise.
func (o *DeliverabilityStatistics) GetClickedCount() int32 {
	if o == nil || IsNil(o.ClickedCount) {
		var ret int32
		return ret
	}
	return *o.ClickedCount
}

// GetClickedCountOk returns a tuple with the ClickedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverabilityStatistics) GetClickedCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ClickedCount) {
		return nil, false
	}
	return o.ClickedCount, true
}

// HasClickedCount returns a boolean if a field has been set.
func (o *DeliverabilityStatistics) HasClickedCount() bool {
	if o != nil && !IsNil(o.ClickedCount) {
		return true
	}

	return false
}

// SetClickedCount gets a reference to the given int32 and assigns it to the ClickedCount field.
func (o *DeliverabilityStatistics) SetClickedCount(v int32) {
	o.ClickedCount = &v
}

func (o DeliverabilityStatistics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliverabilityStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["from_timestamp"] = o.FromTimestamp
	toSerialize["to_timestamp"] = o.ToTimestamp
	if !IsNil(o.ReceptionCount) {
		toSerialize["reception_count"] = o.ReceptionCount
	}
	if !IsNil(o.DeliveredCount) {
		toSerialize["delivered_count"] = o.DeliveredCount
	}
	if !IsNil(o.DeferredCount) {
		toSerialize["deferred_count"] = o.DeferredCount
	}
	if !IsNil(o.BouncedCount) {
		toSerialize["bounced_count"] = o.BouncedCount
	}
	if !IsNil(o.FailedCount) {
		toSerialize["failed_count"] = o.FailedCount
	}
	if !IsNil(o.SuppressedCount) {
		toSerialize["suppressed_count"] = o.SuppressedCount
	}
	if !IsNil(o.OpenedCount) {
		toSerialize["opened_count"] = o.OpenedCount
	}
	if !IsNil(o.ClickedCount) {
		toSerialize["clicked_count"] = o.ClickedCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeliverabilityStatistics) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"from_timestamp",
		"to_timestamp",
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

	varDeliverabilityStatistics := _DeliverabilityStatistics{}

	err = json.Unmarshal(data, &varDeliverabilityStatistics)

	if err != nil {
		return err
	}

	*o = DeliverabilityStatistics(varDeliverabilityStatistics)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "from_timestamp")
		delete(additionalProperties, "to_timestamp")
		delete(additionalProperties, "reception_count")
		delete(additionalProperties, "delivered_count")
		delete(additionalProperties, "deferred_count")
		delete(additionalProperties, "bounced_count")
		delete(additionalProperties, "failed_count")
		delete(additionalProperties, "suppressed_count")
		delete(additionalProperties, "opened_count")
		delete(additionalProperties, "clicked_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeliverabilityStatistics struct {
	value *DeliverabilityStatistics
	isSet bool
}

func (v NullableDeliverabilityStatistics) Get() *DeliverabilityStatistics {
	return v.value
}

func (v *NullableDeliverabilityStatistics) Set(val *DeliverabilityStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliverabilityStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliverabilityStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliverabilityStatistics(val *DeliverabilityStatistics) *NullableDeliverabilityStatistics {
	return &NullableDeliverabilityStatistics{value: val, isSet: true}
}

func (v NullableDeliverabilityStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliverabilityStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
