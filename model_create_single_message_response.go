package ahasend

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateSingleMessageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSingleMessageResponse{}

// CreateSingleMessageResponse struct for CreateSingleMessageResponse
type CreateSingleMessageResponse struct {
	// Object type identifier
	Object string `json:"object"`
	// Message ID (null if the message was not sent)
	Id        *string   `json:"id,omitempty"`
	Recipient Recipient `json:"recipient"`
	// Status of the message
	Status string `json:"status"`
	// Error message if the message was not sent
	Error *string `json:"error,omitempty"`
	// Provided if the request contained a schedule
	Schedule             *MessageSchedule `json:"schedule,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateSingleMessageResponse CreateSingleMessageResponse

// NewCreateSingleMessageResponse instantiates a new CreateSingleMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSingleMessageResponse(object string, recipient Recipient, status string) *CreateSingleMessageResponse {
	this := CreateSingleMessageResponse{}
	this.Object = object
	this.Recipient = recipient
	this.Status = status
	return &this
}

// NewCreateSingleMessageResponseWithDefaults instantiates a new CreateSingleMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSingleMessageResponseWithDefaults() *CreateSingleMessageResponse {
	this := CreateSingleMessageResponse{}
	return &this
}

// GetObject returns the Object field value
func (o *CreateSingleMessageResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *CreateSingleMessageResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *CreateSingleMessageResponse) SetObject(v string) {
	o.Object = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateSingleMessageResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSingleMessageResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateSingleMessageResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateSingleMessageResponse) SetId(v string) {
	o.Id = &v
}

// GetRecipient returns the Recipient field value
func (o *CreateSingleMessageResponse) GetRecipient() Recipient {
	if o == nil {
		var ret Recipient
		return ret
	}

	return o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value
// and a boolean to check if the value has been set.
func (o *CreateSingleMessageResponse) GetRecipientOk() (*Recipient, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipient, true
}

// SetRecipient sets field value
func (o *CreateSingleMessageResponse) SetRecipient(v Recipient) {
	o.Recipient = v
}

// GetStatus returns the Status field value
func (o *CreateSingleMessageResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CreateSingleMessageResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CreateSingleMessageResponse) SetStatus(v string) {
	o.Status = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *CreateSingleMessageResponse) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSingleMessageResponse) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *CreateSingleMessageResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *CreateSingleMessageResponse) SetError(v string) {
	o.Error = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *CreateSingleMessageResponse) GetSchedule() MessageSchedule {
	if o == nil || IsNil(o.Schedule) {
		var ret MessageSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSingleMessageResponse) GetScheduleOk() (*MessageSchedule, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *CreateSingleMessageResponse) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given MessageSchedule and assigns it to the Schedule field.
func (o *CreateSingleMessageResponse) SetSchedule(v MessageSchedule) {
	o.Schedule = &v
}

func (o CreateSingleMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSingleMessageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["recipient"] = o.Recipient
	toSerialize["status"] = o.Status
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateSingleMessageResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"recipient",
		"status",
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

	varCreateSingleMessageResponse := _CreateSingleMessageResponse{}

	err = json.Unmarshal(data, &varCreateSingleMessageResponse)

	if err != nil {
		return err
	}

	*o = CreateSingleMessageResponse(varCreateSingleMessageResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "id")
		delete(additionalProperties, "recipient")
		delete(additionalProperties, "status")
		delete(additionalProperties, "error")
		delete(additionalProperties, "schedule")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateSingleMessageResponse struct {
	value *CreateSingleMessageResponse
	isSet bool
}

func (v NullableCreateSingleMessageResponse) Get() *CreateSingleMessageResponse {
	return v.value
}

func (v *NullableCreateSingleMessageResponse) Set(val *CreateSingleMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSingleMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSingleMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSingleMessageResponse(val *CreateSingleMessageResponse) *NullableCreateSingleMessageResponse {
	return &NullableCreateSingleMessageResponse{value: val, isSet: true}
}

func (v NullableCreateSingleMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSingleMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
