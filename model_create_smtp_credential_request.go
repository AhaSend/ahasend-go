package ahasend

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateSMTPCredentialRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSMTPCredentialRequest{}

// CreateSMTPCredentialRequest struct for CreateSMTPCredentialRequest
type CreateSMTPCredentialRequest struct {
	// Credential name
	Name string `json:"name"`
	// SMTP username
	Username string `json:"username"`
	// SMTP password
	Password string `json:"password"`
	// Whether this is a sandbox credential
	Sandbox *bool `json:"sandbox,omitempty"`
	// Credential scope - \"global\" or \"scoped\"
	Scope string `json:"scope"`
	// Required if scope is \"scoped\"
	Domains              []string `json:"domains,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateSMTPCredentialRequest CreateSMTPCredentialRequest

// NewCreateSMTPCredentialRequest instantiates a new CreateSMTPCredentialRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSMTPCredentialRequest(name string, username string, password string, scope string) *CreateSMTPCredentialRequest {
	this := CreateSMTPCredentialRequest{}
	this.Name = name
	this.Username = username
	this.Password = password
	var sandbox bool = false
	this.Sandbox = &sandbox
	this.Scope = scope
	return &this
}

// NewCreateSMTPCredentialRequestWithDefaults instantiates a new CreateSMTPCredentialRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSMTPCredentialRequestWithDefaults() *CreateSMTPCredentialRequest {
	this := CreateSMTPCredentialRequest{}
	var sandbox bool = false
	this.Sandbox = &sandbox
	return &this
}

// GetName returns the Name field value
func (o *CreateSMTPCredentialRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateSMTPCredentialRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateSMTPCredentialRequest) SetName(v string) {
	o.Name = v
}

// GetUsername returns the Username field value
func (o *CreateSMTPCredentialRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CreateSMTPCredentialRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CreateSMTPCredentialRequest) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *CreateSMTPCredentialRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CreateSMTPCredentialRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CreateSMTPCredentialRequest) SetPassword(v string) {
	o.Password = v
}

// GetSandbox returns the Sandbox field value if set, zero value otherwise.
func (o *CreateSMTPCredentialRequest) GetSandbox() bool {
	if o == nil || IsNil(o.Sandbox) {
		var ret bool
		return ret
	}
	return *o.Sandbox
}

// GetSandboxOk returns a tuple with the Sandbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSMTPCredentialRequest) GetSandboxOk() (*bool, bool) {
	if o == nil || IsNil(o.Sandbox) {
		return nil, false
	}
	return o.Sandbox, true
}

// HasSandbox returns a boolean if a field has been set.
func (o *CreateSMTPCredentialRequest) HasSandbox() bool {
	if o != nil && !IsNil(o.Sandbox) {
		return true
	}

	return false
}

// SetSandbox gets a reference to the given bool and assigns it to the Sandbox field.
func (o *CreateSMTPCredentialRequest) SetSandbox(v bool) {
	o.Sandbox = &v
}

// GetScope returns the Scope field value
func (o *CreateSMTPCredentialRequest) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *CreateSMTPCredentialRequest) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *CreateSMTPCredentialRequest) SetScope(v string) {
	o.Scope = v
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *CreateSMTPCredentialRequest) GetDomains() []string {
	if o == nil || IsNil(o.Domains) {
		var ret []string
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSMTPCredentialRequest) GetDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.Domains) {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *CreateSMTPCredentialRequest) HasDomains() bool {
	if o != nil && !IsNil(o.Domains) {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []string and assigns it to the Domains field.
func (o *CreateSMTPCredentialRequest) SetDomains(v []string) {
	o.Domains = v
}

func (o CreateSMTPCredentialRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSMTPCredentialRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password
	if !IsNil(o.Sandbox) {
		toSerialize["sandbox"] = o.Sandbox
	}
	toSerialize["scope"] = o.Scope
	if !IsNil(o.Domains) {
		toSerialize["domains"] = o.Domains
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateSMTPCredentialRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"username",
		"password",
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

	varCreateSMTPCredentialRequest := _CreateSMTPCredentialRequest{}

	err = json.Unmarshal(data, &varCreateSMTPCredentialRequest)

	if err != nil {
		return err
	}

	*o = CreateSMTPCredentialRequest(varCreateSMTPCredentialRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "username")
		delete(additionalProperties, "password")
		delete(additionalProperties, "sandbox")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "domains")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateSMTPCredentialRequest struct {
	value *CreateSMTPCredentialRequest
	isSet bool
}

func (v NullableCreateSMTPCredentialRequest) Get() *CreateSMTPCredentialRequest {
	return v.value
}

func (v *NullableCreateSMTPCredentialRequest) Set(val *CreateSMTPCredentialRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSMTPCredentialRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSMTPCredentialRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSMTPCredentialRequest(val *CreateSMTPCredentialRequest) *NullableCreateSMTPCredentialRequest {
	return &NullableCreateSMTPCredentialRequest{value: val, isSet: true}
}

func (v NullableCreateSMTPCredentialRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSMTPCredentialRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
