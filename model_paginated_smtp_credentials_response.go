package ahasend

import (
	"encoding/json"
	"fmt"
)

// checks if the PaginatedSMTPCredentialsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedSMTPCredentialsResponse{}

// PaginatedSMTPCredentialsResponse struct for PaginatedSMTPCredentialsResponse
type PaginatedSMTPCredentialsResponse struct {
	// Object type identifier
	Object string `json:"object"`
	// Array of SMTP credentials
	Data                 []SMTPCredential `json:"data"`
	Pagination           PaginationInfo   `json:"pagination"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedSMTPCredentialsResponse PaginatedSMTPCredentialsResponse

// NewPaginatedSMTPCredentialsResponse instantiates a new PaginatedSMTPCredentialsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedSMTPCredentialsResponse(object string, data []SMTPCredential, pagination PaginationInfo) *PaginatedSMTPCredentialsResponse {
	this := PaginatedSMTPCredentialsResponse{}
	this.Object = object
	this.Data = data
	this.Pagination = pagination
	return &this
}

// NewPaginatedSMTPCredentialsResponseWithDefaults instantiates a new PaginatedSMTPCredentialsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedSMTPCredentialsResponseWithDefaults() *PaginatedSMTPCredentialsResponse {
	this := PaginatedSMTPCredentialsResponse{}
	return &this
}

// GetObject returns the Object field value
func (o *PaginatedSMTPCredentialsResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *PaginatedSMTPCredentialsResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *PaginatedSMTPCredentialsResponse) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *PaginatedSMTPCredentialsResponse) GetData() []SMTPCredential {
	if o == nil {
		var ret []SMTPCredential
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PaginatedSMTPCredentialsResponse) GetDataOk() ([]SMTPCredential, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PaginatedSMTPCredentialsResponse) SetData(v []SMTPCredential) {
	o.Data = v
}

// GetPagination returns the Pagination field value
func (o *PaginatedSMTPCredentialsResponse) GetPagination() PaginationInfo {
	if o == nil {
		var ret PaginationInfo
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedSMTPCredentialsResponse) GetPaginationOk() (*PaginationInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedSMTPCredentialsResponse) SetPagination(v PaginationInfo) {
	o.Pagination = v
}

func (o PaginatedSMTPCredentialsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedSMTPCredentialsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["data"] = o.Data
	toSerialize["pagination"] = o.Pagination

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedSMTPCredentialsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varPaginatedSMTPCredentialsResponse := _PaginatedSMTPCredentialsResponse{}

	err = json.Unmarshal(data, &varPaginatedSMTPCredentialsResponse)

	if err != nil {
		return err
	}

	*o = PaginatedSMTPCredentialsResponse(varPaginatedSMTPCredentialsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "data")
		delete(additionalProperties, "pagination")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedSMTPCredentialsResponse struct {
	value *PaginatedSMTPCredentialsResponse
	isSet bool
}

func (v NullablePaginatedSMTPCredentialsResponse) Get() *PaginatedSMTPCredentialsResponse {
	return v.value
}

func (v *NullablePaginatedSMTPCredentialsResponse) Set(val *PaginatedSMTPCredentialsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedSMTPCredentialsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedSMTPCredentialsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedSMTPCredentialsResponse(val *PaginatedSMTPCredentialsResponse) *NullablePaginatedSMTPCredentialsResponse {
	return &NullablePaginatedSMTPCredentialsResponse{value: val, isSet: true}
}

func (v NullablePaginatedSMTPCredentialsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedSMTPCredentialsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
