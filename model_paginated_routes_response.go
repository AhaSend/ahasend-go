package ahasend

import (
	"encoding/json"
	"fmt"
)

// checks if the PaginatedRoutesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedRoutesResponse{}

// PaginatedRoutesResponse struct for PaginatedRoutesResponse
type PaginatedRoutesResponse struct {
	// Object type identifier
	Object string `json:"object"`
	// Array of routes
	Data                 []Route        `json:"data"`
	Pagination           PaginationInfo `json:"pagination"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedRoutesResponse PaginatedRoutesResponse

// NewPaginatedRoutesResponse instantiates a new PaginatedRoutesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedRoutesResponse(object string, data []Route, pagination PaginationInfo) *PaginatedRoutesResponse {
	this := PaginatedRoutesResponse{}
	this.Object = object
	this.Data = data
	this.Pagination = pagination
	return &this
}

// NewPaginatedRoutesResponseWithDefaults instantiates a new PaginatedRoutesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedRoutesResponseWithDefaults() *PaginatedRoutesResponse {
	this := PaginatedRoutesResponse{}
	return &this
}

// GetObject returns the Object field value
func (o *PaginatedRoutesResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *PaginatedRoutesResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *PaginatedRoutesResponse) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *PaginatedRoutesResponse) GetData() []Route {
	if o == nil {
		var ret []Route
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PaginatedRoutesResponse) GetDataOk() ([]Route, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PaginatedRoutesResponse) SetData(v []Route) {
	o.Data = v
}

// GetPagination returns the Pagination field value
func (o *PaginatedRoutesResponse) GetPagination() PaginationInfo {
	if o == nil {
		var ret PaginationInfo
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedRoutesResponse) GetPaginationOk() (*PaginationInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedRoutesResponse) SetPagination(v PaginationInfo) {
	o.Pagination = v
}

func (o PaginatedRoutesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedRoutesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["data"] = o.Data
	toSerialize["pagination"] = o.Pagination

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedRoutesResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"data",
		"pagination",
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

	varPaginatedRoutesResponse := _PaginatedRoutesResponse{}

	err = json.Unmarshal(data, &varPaginatedRoutesResponse)

	if err != nil {
		return err
	}

	*o = PaginatedRoutesResponse(varPaginatedRoutesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "data")
		delete(additionalProperties, "pagination")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedRoutesResponse struct {
	value *PaginatedRoutesResponse
	isSet bool
}

func (v NullablePaginatedRoutesResponse) Get() *PaginatedRoutesResponse {
	return v.value
}

func (v *NullablePaginatedRoutesResponse) Set(val *PaginatedRoutesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedRoutesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedRoutesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedRoutesResponse(val *PaginatedRoutesResponse) *NullablePaginatedRoutesResponse {
	return &NullablePaginatedRoutesResponse{value: val, isSet: true}
}

func (v NullablePaginatedRoutesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedRoutesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
