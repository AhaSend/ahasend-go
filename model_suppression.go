package ahasend

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// checks if the Suppression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Suppression{}

// Suppression struct for Suppression
type Suppression struct {
	// Object type identifier
	Object string `json:"object"`
	// Unique identifier for the suppression
	Id uuid.UUID `json:"id"`
	// When the suppression was created
	CreatedAt time.Time `json:"created_at"`
	// When the suppression was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// Suppressed email address
	Email string `json:"email"`
	// Domain for which the email is suppressed
	Domain *string `json:"domain,omitempty"`
	// Reason for suppression
	Reason *string `json:"reason,omitempty"`
	// When the suppression expires
	ExpiresAt            time.Time `json:"expires_at"`
	AdditionalProperties map[string]interface{}
}

type _Suppression Suppression

// NewSuppression instantiates a new Suppression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuppression(object string, id uuid.UUID, createdAt time.Time, updatedAt time.Time, email string, expiresAt time.Time) *Suppression {
	this := Suppression{}
	this.Object = object
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Email = email
	this.ExpiresAt = expiresAt
	return &this
}

// NewSuppressionWithDefaults instantiates a new Suppression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuppressionWithDefaults() *Suppression {
	this := Suppression{}
	return &this
}

// GetObject returns the Object field value
func (o *Suppression) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Suppression) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Suppression) SetObject(v string) {
	o.Object = v
}

// GetId returns the Id field value
func (o *Suppression) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Suppression) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Suppression) SetId(v uuid.UUID) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Suppression) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Suppression) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Suppression) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Suppression) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Suppression) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Suppression) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetEmail returns the Email field value
func (o *Suppression) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *Suppression) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *Suppression) SetEmail(v string) {
	o.Email = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *Suppression) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Suppression) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *Suppression) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *Suppression) SetDomain(v string) {
	o.Domain = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *Suppression) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Suppression) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *Suppression) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *Suppression) SetReason(v string) {
	o.Reason = &v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *Suppression) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *Suppression) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *Suppression) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

func (o Suppression) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Suppression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
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

func (o *Suppression) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"id",
		"created_at",
		"updated_at",
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

	varSuppression := _Suppression{}

	err = json.Unmarshal(data, &varSuppression)

	if err != nil {
		return err
	}

	*o = Suppression(varSuppression)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "email")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "expires_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSuppression struct {
	value *Suppression
	isSet bool
}

func (v NullableSuppression) Get() *Suppression {
	return v.value
}

func (v *NullableSuppression) Set(val *Suppression) {
	v.value = val
	v.isSet = true
}

func (v NullableSuppression) IsSet() bool {
	return v.isSet
}

func (v *NullableSuppression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuppression(val *Suppression) *NullableSuppression {
	return &NullableSuppression{value: val, isSet: true}
}

func (v NullableSuppression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuppression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
