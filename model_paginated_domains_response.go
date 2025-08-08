package ahasend

import (
	"encoding/json"
	"fmt"
)

// checks if the PaginatedDomainsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedDomainsResponse{}

// PaginatedDomainsResponse struct for PaginatedDomainsResponse
type PaginatedDomainsResponse struct {
	// Object type identifier
	Object string `json:"object"`
	// Array of domains
	Data                 []Domain       `json:"data"`
	Pagination           PaginationInfo `json:"pagination"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedDomainsResponse PaginatedDomainsResponse

// NewPaginatedDomainsResponse instantiates a new PaginatedDomainsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedDomainsResponse(object string, data []Domain, pagination PaginationInfo) *PaginatedDomainsResponse {
	this := PaginatedDomainsResponse{}
	this.Object = object
	this.Data = data
	this.Pagination = pagination
	return &this
}

// NewPaginatedDomainsResponseWithDefaults instantiates a new PaginatedDomainsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedDomainsResponseWithDefaults() *PaginatedDomainsResponse {
	this := PaginatedDomainsResponse{}
	return &this
}

// GetObject returns the Object field value
func (o *PaginatedDomainsResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *PaginatedDomainsResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *PaginatedDomainsResponse) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *PaginatedDomainsResponse) GetData() []Domain {
	if o == nil {
		var ret []Domain
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PaginatedDomainsResponse) GetDataOk() ([]Domain, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PaginatedDomainsResponse) SetData(v []Domain) {
	o.Data = v
}

// GetPagination returns the Pagination field value
func (o *PaginatedDomainsResponse) GetPagination() PaginationInfo {
	if o == nil {
		var ret PaginationInfo
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedDomainsResponse) GetPaginationOk() (*PaginationInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedDomainsResponse) SetPagination(v PaginationInfo) {
	o.Pagination = v
}

func (o PaginatedDomainsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedDomainsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["data"] = o.Data
	toSerialize["pagination"] = o.Pagination

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedDomainsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varPaginatedDomainsResponse := _PaginatedDomainsResponse{}

	err = json.Unmarshal(data, &varPaginatedDomainsResponse)

	if err != nil {
		return err
	}

	*o = PaginatedDomainsResponse(varPaginatedDomainsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "data")
		delete(additionalProperties, "pagination")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedDomainsResponse struct {
	value *PaginatedDomainsResponse
	isSet bool
}

func (v NullablePaginatedDomainsResponse) Get() *PaginatedDomainsResponse {
	return v.value
}

func (v *NullablePaginatedDomainsResponse) Set(val *PaginatedDomainsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedDomainsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedDomainsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedDomainsResponse(val *PaginatedDomainsResponse) *NullablePaginatedDomainsResponse {
	return &NullablePaginatedDomainsResponse{value: val, isSet: true}
}

func (v NullablePaginatedDomainsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedDomainsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
