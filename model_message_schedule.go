package ahasend

import (
	"encoding/json"
	"time"
)

// checks if the MessageSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageSchedule{}

// MessageSchedule struct for MessageSchedule
type MessageSchedule struct {
	// The time to make the first attempt for delivering the message (RFC3339 format)
	FirstAttempt *time.Time `json:"first_attempt,omitempty"`
	// Expire and drop the message if not delivered by this time (RFC3339 format)
	Expires              *time.Time `json:"expires,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MessageSchedule MessageSchedule

// NewMessageSchedule instantiates a new MessageSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageSchedule() *MessageSchedule {
	this := MessageSchedule{}
	return &this
}

// NewMessageScheduleWithDefaults instantiates a new MessageSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageScheduleWithDefaults() *MessageSchedule {
	this := MessageSchedule{}
	return &this
}

// GetFirstAttempt returns the FirstAttempt field value if set, zero value otherwise.
func (o *MessageSchedule) GetFirstAttempt() time.Time {
	if o == nil || IsNil(o.FirstAttempt) {
		var ret time.Time
		return ret
	}
	return *o.FirstAttempt
}

// GetFirstAttemptOk returns a tuple with the FirstAttempt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSchedule) GetFirstAttemptOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FirstAttempt) {
		return nil, false
	}
	return o.FirstAttempt, true
}

// HasFirstAttempt returns a boolean if a field has been set.
func (o *MessageSchedule) HasFirstAttempt() bool {
	if o != nil && !IsNil(o.FirstAttempt) {
		return true
	}

	return false
}

// SetFirstAttempt gets a reference to the given time.Time and assigns it to the FirstAttempt field.
func (o *MessageSchedule) SetFirstAttempt(v time.Time) {
	o.FirstAttempt = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *MessageSchedule) GetExpires() time.Time {
	if o == nil || IsNil(o.Expires) {
		var ret time.Time
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSchedule) GetExpiresOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expires) {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *MessageSchedule) HasExpires() bool {
	if o != nil && !IsNil(o.Expires) {
		return true
	}

	return false
}

// SetExpires gets a reference to the given time.Time and assigns it to the Expires field.
func (o *MessageSchedule) SetExpires(v time.Time) {
	o.Expires = &v
}

func (o MessageSchedule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstAttempt) {
		toSerialize["first_attempt"] = o.FirstAttempt
	}
	if !IsNil(o.Expires) {
		toSerialize["expires"] = o.Expires
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MessageSchedule) UnmarshalJSON(data []byte) (err error) {
	varMessageSchedule := _MessageSchedule{}

	err = json.Unmarshal(data, &varMessageSchedule)

	if err != nil {
		return err
	}

	*o = MessageSchedule(varMessageSchedule)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "first_attempt")
		delete(additionalProperties, "expires")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMessageSchedule struct {
	value *MessageSchedule
	isSet bool
}

func (v NullableMessageSchedule) Get() *MessageSchedule {
	return v.value
}

func (v *NullableMessageSchedule) Set(val *MessageSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageSchedule(val *MessageSchedule) *NullableMessageSchedule {
	return &NullableMessageSchedule{value: val, isSet: true}
}

func (v NullableMessageSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
