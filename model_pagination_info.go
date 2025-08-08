package ahasend

import (
	"encoding/json"
	"fmt"
)

// checks if the PaginationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginationInfo{}

// PaginationInfo struct for PaginationInfo
type PaginationInfo struct {
	// Whether there are more items available
	HasMore bool `json:"has_more"`
	// Cursor for the next page of results
	NextCursor           *string `json:"next_cursor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaginationInfo PaginationInfo

// NewPaginationInfo instantiates a new PaginationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationInfo(hasMore bool) *PaginationInfo {
	this := PaginationInfo{}
	this.HasMore = hasMore
	return &this
}

// NewPaginationInfoWithDefaults instantiates a new PaginationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationInfoWithDefaults() *PaginationInfo {
	this := PaginationInfo{}
	return &this
}

// GetHasMore returns the HasMore field value
func (o *PaginationInfo) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *PaginationInfo) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *PaginationInfo) SetHasMore(v bool) {
	o.HasMore = v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *PaginationInfo) GetNextCursor() string {
	if o == nil || IsNil(o.NextCursor) {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationInfo) GetNextCursorOk() (*string, bool) {
	if o == nil || IsNil(o.NextCursor) {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *PaginationInfo) HasNextCursor() bool {
	if o != nil && !IsNil(o.NextCursor) {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *PaginationInfo) SetNextCursor(v string) {
	o.NextCursor = &v
}

func (o PaginationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["has_more"] = o.HasMore
	if !IsNil(o.NextCursor) {
		toSerialize["next_cursor"] = o.NextCursor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginationInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"has_more",
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

	varPaginationInfo := _PaginationInfo{}

	err = json.Unmarshal(data, &varPaginationInfo)

	if err != nil {
		return err
	}

	*o = PaginationInfo(varPaginationInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "has_more")
		delete(additionalProperties, "next_cursor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginationInfo struct {
	value *PaginationInfo
	isSet bool
}

func (v NullablePaginationInfo) Get() *PaginationInfo {
	return v.value
}

func (v *NullablePaginationInfo) Set(val *PaginationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationInfo(val *PaginationInfo) *NullablePaginationInfo {
	return &NullablePaginationInfo{value: val, isSet: true}
}

func (v NullablePaginationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
