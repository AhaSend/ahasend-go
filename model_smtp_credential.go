package ahasend

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// checks if the SMTPCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SMTPCredential{}

// SMTPCredential struct for SMTPCredential
type SMTPCredential struct {
	// Object type identifier
	Object string `json:"object"`
	// Unique identifier for the SMTP credential
	Id uuid.UUID `json:"id"`
	// When the credential was created
	CreatedAt time.Time `json:"created_at"`
	// When the credential was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// Credential name
	Name string `json:"name"`
	// SMTP username
	Username string `json:"username"`
	// Whether this is a sandbox credential
	Sandbox bool `json:"sandbox"`
	// Credential scope
	Scope string `json:"scope"`
	// Domains this credential can send from
	Domains              []string `json:"domains,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SMTPCredential SMTPCredential

// NewSMTPCredential instantiates a new SMTPCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSMTPCredential(object string, id uuid.UUID, createdAt time.Time, updatedAt time.Time, name string, username string, sandbox bool, scope string) *SMTPCredential {
	this := SMTPCredential{}
	this.Object = object
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Name = name
	this.Username = username
	this.Sandbox = sandbox
	this.Scope = scope
	return &this
}

// NewSMTPCredentialWithDefaults instantiates a new SMTPCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSMTPCredentialWithDefaults() *SMTPCredential {
	this := SMTPCredential{}
	return &this
}

// GetObject returns the Object field value
func (o *SMTPCredential) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *SMTPCredential) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *SMTPCredential) SetObject(v string) {
	o.Object = v
}

// GetId returns the Id field value
func (o *SMTPCredential) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SMTPCredential) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SMTPCredential) SetId(v uuid.UUID) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SMTPCredential) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SMTPCredential) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SMTPCredential) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *SMTPCredential) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *SMTPCredential) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *SMTPCredential) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetName returns the Name field value
func (o *SMTPCredential) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SMTPCredential) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SMTPCredential) SetName(v string) {
	o.Name = v
}

// GetUsername returns the Username field value
func (o *SMTPCredential) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *SMTPCredential) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *SMTPCredential) SetUsername(v string) {
	o.Username = v
}

// GetSandbox returns the Sandbox field value
func (o *SMTPCredential) GetSandbox() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Sandbox
}

// GetSandboxOk returns a tuple with the Sandbox field value
// and a boolean to check if the value has been set.
func (o *SMTPCredential) GetSandboxOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sandbox, true
}

// SetSandbox sets field value
func (o *SMTPCredential) SetSandbox(v bool) {
	o.Sandbox = v
}

// GetScope returns the Scope field value
func (o *SMTPCredential) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *SMTPCredential) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *SMTPCredential) SetScope(v string) {
	o.Scope = v
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *SMTPCredential) GetDomains() []string {
	if o == nil || IsNil(o.Domains) {
		var ret []string
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMTPCredential) GetDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.Domains) {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *SMTPCredential) HasDomains() bool {
	if o != nil && !IsNil(o.Domains) {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []string and assigns it to the Domains field.
func (o *SMTPCredential) SetDomains(v []string) {
	o.Domains = v
}

func (o SMTPCredential) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SMTPCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["name"] = o.Name
	toSerialize["username"] = o.Username
	toSerialize["sandbox"] = o.Sandbox
	toSerialize["scope"] = o.Scope
	if !IsNil(o.Domains) {
		toSerialize["domains"] = o.Domains
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SMTPCredential) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"id",
		"created_at",
		"updated_at",
		"name",
		"username",
		"sandbox",
		"scope",
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

	varSMTPCredential := _SMTPCredential{}

	err = json.Unmarshal(data, &varSMTPCredential)

	if err != nil {
		return err
	}

	*o = SMTPCredential(varSMTPCredential)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "name")
		delete(additionalProperties, "username")
		delete(additionalProperties, "sandbox")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "domains")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSMTPCredential struct {
	value *SMTPCredential
	isSet bool
}

func (v NullableSMTPCredential) Get() *SMTPCredential {
	return v.value
}

func (v *NullableSMTPCredential) Set(val *SMTPCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableSMTPCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableSMTPCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSMTPCredential(val *SMTPCredential) *NullableSMTPCredential {
	return &NullableSMTPCredential{value: val, isSet: true}
}

func (v NullableSMTPCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSMTPCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
