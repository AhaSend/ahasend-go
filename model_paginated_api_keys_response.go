package ahasend

import (
	"encoding/json"
	"fmt"
)

// checks if the PaginatedAPIKeysResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedAPIKeysResponse{}

// PaginatedAPIKeysResponse struct for PaginatedAPIKeysResponse
type PaginatedAPIKeysResponse struct {
	// Object type identifier
	Object string `json:"object"`
	// Array of API keys
	Data                 []ModelAPIKey  `json:"data"`
	Pagination           PaginationInfo `json:"pagination"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedAPIKeysResponse PaginatedAPIKeysResponse

// NewPaginatedAPIKeysResponse instantiates a new PaginatedAPIKeysResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedAPIKeysResponse(object string, data []ModelAPIKey, pagination PaginationInfo) *PaginatedAPIKeysResponse {
	this := PaginatedAPIKeysResponse{}
	this.Object = object
	this.Data = data
	this.Pagination = pagination
	return &this
}

// NewPaginatedAPIKeysResponseWithDefaults instantiates a new PaginatedAPIKeysResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedAPIKeysResponseWithDefaults() *PaginatedAPIKeysResponse {
	this := PaginatedAPIKeysResponse{}
	return &this
}

// GetObject returns the Object field value
func (o *PaginatedAPIKeysResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *PaginatedAPIKeysResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *PaginatedAPIKeysResponse) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *PaginatedAPIKeysResponse) GetData() []ModelAPIKey {
	if o == nil {
		var ret []ModelAPIKey
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PaginatedAPIKeysResponse) GetDataOk() ([]ModelAPIKey, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PaginatedAPIKeysResponse) SetData(v []ModelAPIKey) {
	o.Data = v
}

// GetPagination returns the Pagination field value
func (o *PaginatedAPIKeysResponse) GetPagination() PaginationInfo {
	if o == nil {
		var ret PaginationInfo
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedAPIKeysResponse) GetPaginationOk() (*PaginationInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedAPIKeysResponse) SetPagination(v PaginationInfo) {
	o.Pagination = v
}

func (o PaginatedAPIKeysResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedAPIKeysResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["data"] = o.Data
	toSerialize["pagination"] = o.Pagination

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedAPIKeysResponse) UnmarshalJSON(data []byte) (err error) {
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

	varPaginatedAPIKeysResponse := _PaginatedAPIKeysResponse{}

	err = json.Unmarshal(data, &varPaginatedAPIKeysResponse)

	if err != nil {
		return err
	}

	*o = PaginatedAPIKeysResponse(varPaginatedAPIKeysResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "data")
		delete(additionalProperties, "pagination")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedAPIKeysResponse struct {
	value *PaginatedAPIKeysResponse
	isSet bool
}

func (v NullablePaginatedAPIKeysResponse) Get() *PaginatedAPIKeysResponse {
	return v.value
}

func (v *NullablePaginatedAPIKeysResponse) Set(val *PaginatedAPIKeysResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedAPIKeysResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedAPIKeysResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedAPIKeysResponse(val *PaginatedAPIKeysResponse) *NullablePaginatedAPIKeysResponse {
	return &NullablePaginatedAPIKeysResponse{value: val, isSet: true}
}

func (v NullablePaginatedAPIKeysResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedAPIKeysResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
