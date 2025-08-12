package ahasend

import (
	"encoding/json"
)

// checks if the DeliveryTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryTime{}

// DeliveryTime struct for DeliveryTime
type DeliveryTime struct {
	// The recipient domain
	RecipientDomain *string `json:"recipient_domain,omitempty"`
	// The average time from reception to delivery in seconds
	DeliveryTime         *float64 `json:"delivery_time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeliveryTime DeliveryTime

// NewDeliveryTime instantiates a new DeliveryTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryTime() *DeliveryTime {
	this := DeliveryTime{}
	return &this
}

// NewDeliveryTimeWithDefaults instantiates a new DeliveryTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryTimeWithDefaults() *DeliveryTime {
	this := DeliveryTime{}
	return &this
}

// GetRecipientDomain returns the RecipientDomain field value if set, zero value otherwise.
func (o *DeliveryTime) GetRecipientDomain() string {
	if o == nil || IsNil(o.RecipientDomain) {
		var ret string
		return ret
	}
	return *o.RecipientDomain
}

// GetRecipientDomainOk returns a tuple with the RecipientDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryTime) GetRecipientDomainOk() (*string, bool) {
	if o == nil || IsNil(o.RecipientDomain) {
		return nil, false
	}
	return o.RecipientDomain, true
}

// HasRecipientDomain returns a boolean if a field has been set.
func (o *DeliveryTime) HasRecipientDomain() bool {
	if o != nil && !IsNil(o.RecipientDomain) {
		return true
	}

	return false
}

// SetRecipientDomain gets a reference to the given string and assigns it to the RecipientDomain field.
func (o *DeliveryTime) SetRecipientDomain(v string) {
	o.RecipientDomain = &v
}

// GetDeliveryTime returns the DeliveryTime field value if set, zero value otherwise.
func (o *DeliveryTime) GetDeliveryTime() float64 {
	if o == nil || IsNil(o.DeliveryTime) {
		var ret float64
		return ret
	}
	return *o.DeliveryTime
}

// GetDeliveryTimeOk returns a tuple with the DeliveryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryTime) GetDeliveryTimeOk() (*float64, bool) {
	if o == nil || IsNil(o.DeliveryTime) {
		return nil, false
	}
	return o.DeliveryTime, true
}

// HasDeliveryTime returns a boolean if a field has been set.
func (o *DeliveryTime) HasDeliveryTime() bool {
	if o != nil && !IsNil(o.DeliveryTime) {
		return true
	}

	return false
}

// SetDeliveryTime gets a reference to the given float64 and assigns it to the DeliveryTime field.
func (o *DeliveryTime) SetDeliveryTime(v float64) {
	o.DeliveryTime = &v
}

func (o DeliveryTime) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecipientDomain) {
		toSerialize["recipient_domain"] = o.RecipientDomain
	}
	if !IsNil(o.DeliveryTime) {
		toSerialize["delivery_time"] = o.DeliveryTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeliveryTime) UnmarshalJSON(data []byte) (err error) {
	varDeliveryTime := _DeliveryTime{}

	err = json.Unmarshal(data, &varDeliveryTime)

	if err != nil {
		return err
	}

	*o = DeliveryTime(varDeliveryTime)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "recipient_domain")
		delete(additionalProperties, "delivery_time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeliveryTime struct {
	value *DeliveryTime
	isSet bool
}

func (v NullableDeliveryTime) Get() *DeliveryTime {
	return v.value
}

func (v *NullableDeliveryTime) Set(val *DeliveryTime) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryTime) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryTime(val *DeliveryTime) *NullableDeliveryTime {
	return &NullableDeliveryTime{value: val, isSet: true}
}

func (v NullableDeliveryTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
