package ahasend

import (
	"encoding/json"
	"fmt"
)

// checks if the PaginatedMessagesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedMessagesResponse{}

// PaginatedMessagesResponse struct for PaginatedMessagesResponse
type PaginatedMessagesResponse struct {
	// Object type identifier
	Object string `json:"object"`
	// Array of messages
	Data                 []Message      `json:"data"`
	Pagination           PaginationInfo `json:"pagination"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedMessagesResponse PaginatedMessagesResponse

// NewPaginatedMessagesResponse instantiates a new PaginatedMessagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedMessagesResponse(object string, data []Message, pagination PaginationInfo) *PaginatedMessagesResponse {
	this := PaginatedMessagesResponse{}
	this.Object = object
	this.Data = data
	this.Pagination = pagination
	return &this
}

// NewPaginatedMessagesResponseWithDefaults instantiates a new PaginatedMessagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedMessagesResponseWithDefaults() *PaginatedMessagesResponse {
	this := PaginatedMessagesResponse{}
	return &this
}

// GetObject returns the Object field value
func (o *PaginatedMessagesResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *PaginatedMessagesResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *PaginatedMessagesResponse) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *PaginatedMessagesResponse) GetData() []Message {
	if o == nil {
		var ret []Message
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PaginatedMessagesResponse) GetDataOk() ([]Message, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PaginatedMessagesResponse) SetData(v []Message) {
	o.Data = v
}

// GetPagination returns the Pagination field value
func (o *PaginatedMessagesResponse) GetPagination() PaginationInfo {
	if o == nil {
		var ret PaginationInfo
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedMessagesResponse) GetPaginationOk() (*PaginationInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedMessagesResponse) SetPagination(v PaginationInfo) {
	o.Pagination = v
}

func (o PaginatedMessagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedMessagesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["data"] = o.Data
	toSerialize["pagination"] = o.Pagination

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedMessagesResponse) UnmarshalJSON(data []byte) (err error) {
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

	varPaginatedMessagesResponse := _PaginatedMessagesResponse{}

	err = json.Unmarshal(data, &varPaginatedMessagesResponse)

	if err != nil {
		return err
	}

	*o = PaginatedMessagesResponse(varPaginatedMessagesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "data")
		delete(additionalProperties, "pagination")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedMessagesResponse struct {
	value *PaginatedMessagesResponse
	isSet bool
}

func (v NullablePaginatedMessagesResponse) Get() *PaginatedMessagesResponse {
	return v.value
}

func (v *NullablePaginatedMessagesResponse) Set(val *PaginatedMessagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedMessagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedMessagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedMessagesResponse(val *PaginatedMessagesResponse) *NullablePaginatedMessagesResponse {
	return &NullablePaginatedMessagesResponse{value: val, isSet: true}
}

func (v NullablePaginatedMessagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedMessagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
