package ahasend

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// checks if the APIKeyScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIKeyScope{}

// APIKeyScope struct for APIKeyScope
type APIKeyScope struct {
	// Unique identifier for the scope
	Id uuid.UUID `json:"id"`
	// When the scope was created
	CreatedAt time.Time `json:"created_at"`
	// When the scope was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// ID of the API key this scope belongs to
	ApiKeyId uuid.UUID `json:"api_key_id"`
	// The scope string
	Scope string `json:"scope"`
	// Domain ID for domain-specific scopes
	DomainId             *uuid.UUID `json:"domain_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _APIKeyScope APIKeyScope

// NewAPIKeyScope instantiates a new APIKeyScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIKeyScope(id uuid.UUID, createdAt time.Time, updatedAt time.Time, apiKeyId uuid.UUID, scope string) *APIKeyScope {
	this := APIKeyScope{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.ApiKeyId = apiKeyId
	this.Scope = scope
	return &this
}

// NewAPIKeyScopeWithDefaults instantiates a new APIKeyScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIKeyScopeWithDefaults() *APIKeyScope {
	this := APIKeyScope{}
	return &this
}

// GetId returns the Id field value
func (o *APIKeyScope) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *APIKeyScope) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *APIKeyScope) SetId(v uuid.UUID) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *APIKeyScope) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *APIKeyScope) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *APIKeyScope) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *APIKeyScope) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *APIKeyScope) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *APIKeyScope) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetApiKeyId returns the ApiKeyId field value
func (o *APIKeyScope) GetApiKeyId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}

	return o.ApiKeyId
}

// GetApiKeyIdOk returns a tuple with the ApiKeyId field value
// and a boolean to check if the value has been set.
func (o *APIKeyScope) GetApiKeyIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKeyId, true
}

// SetApiKeyId sets field value
func (o *APIKeyScope) SetApiKeyId(v uuid.UUID) {
	o.ApiKeyId = v
}

// GetScope returns the Scope field value
func (o *APIKeyScope) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *APIKeyScope) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *APIKeyScope) SetScope(v string) {
	o.Scope = v
}

// GetDomainId returns the DomainId field value if set, zero value otherwise.
func (o *APIKeyScope) GetDomainId() uuid.UUID {
	if o == nil || IsNil(o.DomainId) {
		var ret uuid.UUID
		return ret
	}
	return *o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIKeyScope) GetDomainIdOk() (*uuid.UUID, bool) {
	if o == nil || IsNil(o.DomainId) {
		return nil, false
	}
	return o.DomainId, true
}

// HasDomainId returns a boolean if a field has been set.
func (o *APIKeyScope) HasDomainId() bool {
	if o != nil && !IsNil(o.DomainId) {
		return true
	}

	return false
}

// SetDomainId gets a reference to the given uuid.UUID and assigns it to the DomainId field.
func (o *APIKeyScope) SetDomainId(v uuid.UUID) {
	o.DomainId = &v
}

func (o APIKeyScope) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIKeyScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["api_key_id"] = o.ApiKeyId
	toSerialize["scope"] = o.Scope
	if !IsNil(o.DomainId) {
		toSerialize["domain_id"] = o.DomainId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *APIKeyScope) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
		"api_key_id",
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

	varAPIKeyScope := _APIKeyScope{}

	err = json.Unmarshal(data, &varAPIKeyScope)

	if err != nil {
		return err
	}

	*o = APIKeyScope(varAPIKeyScope)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "api_key_id")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "domain_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAPIKeyScope struct {
	value *APIKeyScope
	isSet bool
}

func (v NullableAPIKeyScope) Get() *APIKeyScope {
	return v.value
}

func (v *NullableAPIKeyScope) Set(val *APIKeyScope) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIKeyScope) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIKeyScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIKeyScope(val *APIKeyScope) *NullableAPIKeyScope {
	return &NullableAPIKeyScope{value: val, isSet: true}
}

func (v NullableAPIKeyScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIKeyScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
