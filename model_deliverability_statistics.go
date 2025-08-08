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
	// Time bucket for the statistics
	TimeBucket time.Time `json:"time_bucket"`
	// Message direction
	Direction string `json:"direction"`
	// Number of messages sent
	Sent *int32 `json:"sent,omitempty"`
	// Number of messages delivered
	Delivered *int32 `json:"delivered,omitempty"`
	// Number of messages bounced
	Bounced *int32 `json:"bounced,omitempty"`
	// Number of messages failed
	Failed *int32 `json:"failed,omitempty"`
	// Number of messages suppressed
	Suppressed           *int32 `json:"suppressed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeliverabilityStatistics DeliverabilityStatistics

// NewDeliverabilityStatistics instantiates a new DeliverabilityStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliverabilityStatistics(timeBucket time.Time, direction string) *DeliverabilityStatistics {
	this := DeliverabilityStatistics{}
	this.TimeBucket = timeBucket
	this.Direction = direction
	return &this
}

// NewDeliverabilityStatisticsWithDefaults instantiates a new DeliverabilityStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliverabilityStatisticsWithDefaults() *DeliverabilityStatistics {
	this := DeliverabilityStatistics{}
	return &this
}

// GetTimeBucket returns the TimeBucket field value
func (o *DeliverabilityStatistics) GetTimeBucket() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TimeBucket
}

// GetTimeBucketOk returns a tuple with the TimeBucket field value
// and a boolean to check if the value has been set.
func (o *DeliverabilityStatistics) GetTimeBucketOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeBucket, true
}

// SetTimeBucket sets field value
func (o *DeliverabilityStatistics) SetTimeBucket(v time.Time) {
	o.TimeBucket = v
}

// GetDirection returns the Direction field value
func (o *DeliverabilityStatistics) GetDirection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *DeliverabilityStatistics) GetDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *DeliverabilityStatistics) SetDirection(v string) {
	o.Direction = v
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *DeliverabilityStatistics) GetSent() int32 {
	if o == nil || IsNil(o.Sent) {
		var ret int32
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverabilityStatistics) GetSentOk() (*int32, bool) {
	if o == nil || IsNil(o.Sent) {
		return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *DeliverabilityStatistics) HasSent() bool {
	if o != nil && !IsNil(o.Sent) {
		return true
	}

	return false
}

// SetSent gets a reference to the given int32 and assigns it to the Sent field.
func (o *DeliverabilityStatistics) SetSent(v int32) {
	o.Sent = &v
}

// GetDelivered returns the Delivered field value if set, zero value otherwise.
func (o *DeliverabilityStatistics) GetDelivered() int32 {
	if o == nil || IsNil(o.Delivered) {
		var ret int32
		return ret
	}
	return *o.Delivered
}

// GetDeliveredOk returns a tuple with the Delivered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverabilityStatistics) GetDeliveredOk() (*int32, bool) {
	if o == nil || IsNil(o.Delivered) {
		return nil, false
	}
	return o.Delivered, true
}

// HasDelivered returns a boolean if a field has been set.
func (o *DeliverabilityStatistics) HasDelivered() bool {
	if o != nil && !IsNil(o.Delivered) {
		return true
	}

	return false
}

// SetDelivered gets a reference to the given int32 and assigns it to the Delivered field.
func (o *DeliverabilityStatistics) SetDelivered(v int32) {
	o.Delivered = &v
}

// GetBounced returns the Bounced field value if set, zero value otherwise.
func (o *DeliverabilityStatistics) GetBounced() int32 {
	if o == nil || IsNil(o.Bounced) {
		var ret int32
		return ret
	}
	return *o.Bounced
}

// GetBouncedOk returns a tuple with the Bounced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverabilityStatistics) GetBouncedOk() (*int32, bool) {
	if o == nil || IsNil(o.Bounced) {
		return nil, false
	}
	return o.Bounced, true
}

// HasBounced returns a boolean if a field has been set.
func (o *DeliverabilityStatistics) HasBounced() bool {
	if o != nil && !IsNil(o.Bounced) {
		return true
	}

	return false
}

// SetBounced gets a reference to the given int32 and assigns it to the Bounced field.
func (o *DeliverabilityStatistics) SetBounced(v int32) {
	o.Bounced = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *DeliverabilityStatistics) GetFailed() int32 {
	if o == nil || IsNil(o.Failed) {
		var ret int32
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverabilityStatistics) GetFailedOk() (*int32, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *DeliverabilityStatistics) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given int32 and assigns it to the Failed field.
func (o *DeliverabilityStatistics) SetFailed(v int32) {
	o.Failed = &v
}

// GetSuppressed returns the Suppressed field value if set, zero value otherwise.
func (o *DeliverabilityStatistics) GetSuppressed() int32 {
	if o == nil || IsNil(o.Suppressed) {
		var ret int32
		return ret
	}
	return *o.Suppressed
}

// GetSuppressedOk returns a tuple with the Suppressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverabilityStatistics) GetSuppressedOk() (*int32, bool) {
	if o == nil || IsNil(o.Suppressed) {
		return nil, false
	}
	return o.Suppressed, true
}

// HasSuppressed returns a boolean if a field has been set.
func (o *DeliverabilityStatistics) HasSuppressed() bool {
	if o != nil && !IsNil(o.Suppressed) {
		return true
	}

	return false
}

// SetSuppressed gets a reference to the given int32 and assigns it to the Suppressed field.
func (o *DeliverabilityStatistics) SetSuppressed(v int32) {
	o.Suppressed = &v
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
	toSerialize["time_bucket"] = o.TimeBucket
	toSerialize["direction"] = o.Direction
	if !IsNil(o.Sent) {
		toSerialize["sent"] = o.Sent
	}
	if !IsNil(o.Delivered) {
		toSerialize["delivered"] = o.Delivered
	}
	if !IsNil(o.Bounced) {
		toSerialize["bounced"] = o.Bounced
	}
	if !IsNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	if !IsNil(o.Suppressed) {
		toSerialize["suppressed"] = o.Suppressed
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
		"time_bucket",
		"direction",
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
		delete(additionalProperties, "time_bucket")
		delete(additionalProperties, "direction")
		delete(additionalProperties, "sent")
		delete(additionalProperties, "delivered")
		delete(additionalProperties, "bounced")
		delete(additionalProperties, "failed")
		delete(additionalProperties, "suppressed")
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
