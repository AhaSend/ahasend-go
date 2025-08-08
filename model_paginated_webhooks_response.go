package ahasend

import (
	"encoding/json"
	"fmt"
)

// checks if the PaginatedWebhooksResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedWebhooksResponse{}

// PaginatedWebhooksResponse struct for PaginatedWebhooksResponse
type PaginatedWebhooksResponse struct {
	// Object type identifier
	Object string `json:"object"`
	// Array of webhooks
	Data                 []Webhook      `json:"data"`
	Pagination           PaginationInfo `json:"pagination"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedWebhooksResponse PaginatedWebhooksResponse

// NewPaginatedWebhooksResponse instantiates a new PaginatedWebhooksResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedWebhooksResponse(object string, data []Webhook, pagination PaginationInfo) *PaginatedWebhooksResponse {
	this := PaginatedWebhooksResponse{}
	this.Object = object
	this.Data = data
	this.Pagination = pagination
	return &this
}

// NewPaginatedWebhooksResponseWithDefaults instantiates a new PaginatedWebhooksResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedWebhooksResponseWithDefaults() *PaginatedWebhooksResponse {
	this := PaginatedWebhooksResponse{}
	return &this
}

// GetObject returns the Object field value
func (o *PaginatedWebhooksResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *PaginatedWebhooksResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *PaginatedWebhooksResponse) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *PaginatedWebhooksResponse) GetData() []Webhook {
	if o == nil {
		var ret []Webhook
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PaginatedWebhooksResponse) GetDataOk() ([]Webhook, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PaginatedWebhooksResponse) SetData(v []Webhook) {
	o.Data = v
}

// GetPagination returns the Pagination field value
func (o *PaginatedWebhooksResponse) GetPagination() PaginationInfo {
	if o == nil {
		var ret PaginationInfo
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedWebhooksResponse) GetPaginationOk() (*PaginationInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedWebhooksResponse) SetPagination(v PaginationInfo) {
	o.Pagination = v
}

func (o PaginatedWebhooksResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedWebhooksResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["data"] = o.Data
	toSerialize["pagination"] = o.Pagination

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedWebhooksResponse) UnmarshalJSON(data []byte) (err error) {
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

	varPaginatedWebhooksResponse := _PaginatedWebhooksResponse{}

	err = json.Unmarshal(data, &varPaginatedWebhooksResponse)

	if err != nil {
		return err
	}

	*o = PaginatedWebhooksResponse(varPaginatedWebhooksResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "data")
		delete(additionalProperties, "pagination")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedWebhooksResponse struct {
	value *PaginatedWebhooksResponse
	isSet bool
}

func (v NullablePaginatedWebhooksResponse) Get() *PaginatedWebhooksResponse {
	return v.value
}

func (v *NullablePaginatedWebhooksResponse) Set(val *PaginatedWebhooksResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedWebhooksResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedWebhooksResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedWebhooksResponse(val *PaginatedWebhooksResponse) *NullablePaginatedWebhooksResponse {
	return &NullablePaginatedWebhooksResponse{value: val, isSet: true}
}

func (v NullablePaginatedWebhooksResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedWebhooksResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
