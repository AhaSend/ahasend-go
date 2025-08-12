package ahasend

import (
	"encoding/json"
)

// checks if the Tracking type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tracking{}

// Tracking struct for Tracking
type Tracking struct {
	// Whether to track opens
	Open *bool `json:"open,omitempty"`
	// Whether to track clicks
	Click                *bool `json:"click,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Tracking Tracking

// NewTracking instantiates a new Tracking object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTracking() *Tracking {
	this := Tracking{}
	return &this
}

// NewTrackingWithDefaults instantiates a new Tracking object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackingWithDefaults() *Tracking {
	this := Tracking{}
	return &this
}

// GetOpen returns the Open field value if set, zero value otherwise.
func (o *Tracking) GetOpen() bool {
	if o == nil || IsNil(o.Open) {
		var ret bool
		return ret
	}
	return *o.Open
}

// GetOpenOk returns a tuple with the Open field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tracking) GetOpenOk() (*bool, bool) {
	if o == nil || IsNil(o.Open) {
		return nil, false
	}
	return o.Open, true
}

// HasOpen returns a boolean if a field has been set.
func (o *Tracking) HasOpen() bool {
	if o != nil && !IsNil(o.Open) {
		return true
	}

	return false
}

// SetOpen gets a reference to the given bool and assigns it to the Open field.
func (o *Tracking) SetOpen(v bool) {
	o.Open = &v
}

// GetClick returns the Click field value if set, zero value otherwise.
func (o *Tracking) GetClick() bool {
	if o == nil || IsNil(o.Click) {
		var ret bool
		return ret
	}
	return *o.Click
}

// GetClickOk returns a tuple with the Click field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tracking) GetClickOk() (*bool, bool) {
	if o == nil || IsNil(o.Click) {
		return nil, false
	}
	return o.Click, true
}

// HasClick returns a boolean if a field has been set.
func (o *Tracking) HasClick() bool {
	if o != nil && !IsNil(o.Click) {
		return true
	}

	return false
}

// SetClick gets a reference to the given bool and assigns it to the Click field.
func (o *Tracking) SetClick(v bool) {
	o.Click = &v
}

func (o Tracking) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tracking) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Open) {
		toSerialize["open"] = o.Open
	}
	if !IsNil(o.Click) {
		toSerialize["click"] = o.Click
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Tracking) UnmarshalJSON(data []byte) (err error) {
	varTracking := _Tracking{}

	err = json.Unmarshal(data, &varTracking)

	if err != nil {
		return err
	}

	*o = Tracking(varTracking)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "open")
		delete(additionalProperties, "click")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTracking struct {
	value *Tracking
	isSet bool
}

func (v NullableTracking) Get() *Tracking {
	return v.value
}

func (v *NullableTracking) Set(val *Tracking) {
	v.value = val
	v.isSet = true
}

func (v NullableTracking) IsSet() bool {
	return v.isSet
}

func (v *NullableTracking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTracking(val *Tracking) *NullableTracking {
	return &NullableTracking{value: val, isSet: true}
}

func (v NullableTracking) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTracking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
