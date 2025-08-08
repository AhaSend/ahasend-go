package ahasend

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the CreateSuppressionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSuppressionRequest{}

// CreateSuppressionRequest struct for CreateSuppressionRequest
type CreateSuppressionRequest struct {
	// Email address to suppress
	Email string `json:"email"`
	// Domain for which to suppress the email
	Domain *string `json:"domain,omitempty"`
	// Reason for suppression
	Reason *string `json:"reason,omitempty"`
	// When the suppression expires (RFC3339 format)
	ExpiresAt            time.Time `json:"expires_at"`
	AdditionalProperties map[string]interface{}
}

type _CreateSuppressionRequest CreateSuppressionRequest

// NewCreateSuppressionRequest instantiates a new CreateSuppressionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSuppressionRequest(email string, expiresAt time.Time) *CreateSuppressionRequest {
	this := CreateSuppressionRequest{}
	this.Email = email
	this.ExpiresAt = expiresAt
	return &this
}

// NewCreateSuppressionRequestWithDefaults instantiates a new CreateSuppressionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSuppressionRequestWithDefaults() *CreateSuppressionRequest {
	this := CreateSuppressionRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *CreateSuppressionRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CreateSuppressionRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CreateSuppressionRequest) SetEmail(v string) {
	o.Email = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *CreateSuppressionRequest) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSuppressionRequest) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *CreateSuppressionRequest) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *CreateSuppressionRequest) SetDomain(v string) {
	o.Domain = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *CreateSuppressionRequest) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSuppressionRequest) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *CreateSuppressionRequest) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *CreateSuppressionRequest) SetReason(v string) {
	o.Reason = &v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *CreateSuppressionRequest) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *CreateSuppressionRequest) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *CreateSuppressionRequest) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

func (o CreateSuppressionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSuppressionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	toSerialize["expires_at"] = o.ExpiresAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateSuppressionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"expires_at",
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

	varCreateSuppressionRequest := _CreateSuppressionRequest{}

	err = json.Unmarshal(data, &varCreateSuppressionRequest)

	if err != nil {
		return err
	}

	*o = CreateSuppressionRequest(varCreateSuppressionRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "expires_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateSuppressionRequest struct {
	value *CreateSuppressionRequest
	isSet bool
}

func (v NullableCreateSuppressionRequest) Get() *CreateSuppressionRequest {
	return v.value
}

func (v *NullableCreateSuppressionRequest) Set(val *CreateSuppressionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSuppressionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSuppressionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSuppressionRequest(val *CreateSuppressionRequest) *NullableCreateSuppressionRequest {
	return &NullableCreateSuppressionRequest{value: val, isSet: true}
}

func (v NullableCreateSuppressionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSuppressionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
