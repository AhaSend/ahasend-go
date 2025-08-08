package ahasend

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// checks if the ModelAPIKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelAPIKey{}

// ModelAPIKey struct for ModelAPIKey
type ModelAPIKey struct {
	// Object type identifier
	Object string `json:"object"`
	// Unique identifier for the API key
	Id uuid.UUID `json:"id"`
	// When the API key was created
	CreatedAt time.Time `json:"created_at"`
	// When the API key was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// When the API key was last used (updates every 5-10 minutes)
	LastUsedAt *time.Time `json:"last_used_at,omitempty"`
	// Account ID this API key belongs to
	AccountId uuid.UUID `json:"account_id"`
	// Human-readable label for the API key
	Label string `json:"label"`
	// Public portion of the API key
	PublicKey string `json:"public_key"`
	// Secret key (only returned on creation)
	SecretKey *string `json:"secret_key,omitempty"`
	// Scopes granted to this API key
	Scopes               []APIKeyScope `json:"scopes"`
	AdditionalProperties map[string]interface{}
}

type _ModelAPIKey ModelAPIKey

// NewModelAPIKey instantiates a new ModelAPIKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelAPIKey(object string, id uuid.UUID, createdAt time.Time, updatedAt time.Time, accountId uuid.UUID, label string, publicKey string, scopes []APIKeyScope) *ModelAPIKey {
	this := ModelAPIKey{}
	this.Object = object
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.AccountId = accountId
	this.Label = label
	this.PublicKey = publicKey
	this.Scopes = scopes
	return &this
}

// NewModelAPIKeyWithDefaults instantiates a new ModelAPIKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelAPIKeyWithDefaults() *ModelAPIKey {
	this := ModelAPIKey{}
	return &this
}

// GetObject returns the Object field value
func (o *ModelAPIKey) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ModelAPIKey) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ModelAPIKey) SetObject(v string) {
	o.Object = v
}

// GetId returns the Id field value
func (o *ModelAPIKey) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModelAPIKey) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModelAPIKey) SetId(v uuid.UUID) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ModelAPIKey) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelAPIKey) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ModelAPIKey) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ModelAPIKey) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelAPIKey) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ModelAPIKey) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise.
func (o *ModelAPIKey) GetLastUsedAt() time.Time {
	if o == nil || IsNil(o.LastUsedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIKey) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUsedAt) {
		return nil, false
	}
	return o.LastUsedAt, true
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *ModelAPIKey) HasLastUsedAt() bool {
	if o != nil && !IsNil(o.LastUsedAt) {
		return true
	}

	return false
}

// SetLastUsedAt gets a reference to the given time.Time and assigns it to the LastUsedAt field.
func (o *ModelAPIKey) SetLastUsedAt(v time.Time) {
	o.LastUsedAt = &v
}

// GetAccountId returns the AccountId field value
func (o *ModelAPIKey) GetAccountId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *ModelAPIKey) GetAccountIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *ModelAPIKey) SetAccountId(v uuid.UUID) {
	o.AccountId = v
}

// GetLabel returns the Label field value
func (o *ModelAPIKey) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ModelAPIKey) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ModelAPIKey) SetLabel(v string) {
	o.Label = v
}

// GetPublicKey returns the PublicKey field value
func (o *ModelAPIKey) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *ModelAPIKey) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *ModelAPIKey) SetPublicKey(v string) {
	o.PublicKey = v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *ModelAPIKey) GetSecretKey() string {
	if o == nil || IsNil(o.SecretKey) {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIKey) GetSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SecretKey) {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *ModelAPIKey) HasSecretKey() bool {
	if o != nil && !IsNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *ModelAPIKey) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetScopes returns the Scopes field value
func (o *ModelAPIKey) GetScopes() []APIKeyScope {
	if o == nil {
		var ret []APIKeyScope
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *ModelAPIKey) GetScopesOk() ([]APIKeyScope, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *ModelAPIKey) SetScopes(v []APIKeyScope) {
	o.Scopes = v
}

func (o ModelAPIKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelAPIKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	if !IsNil(o.LastUsedAt) {
		toSerialize["last_used_at"] = o.LastUsedAt
	}
	toSerialize["account_id"] = o.AccountId
	toSerialize["label"] = o.Label
	toSerialize["public_key"] = o.PublicKey
	if !IsNil(o.SecretKey) {
		toSerialize["secret_key"] = o.SecretKey
	}
	toSerialize["scopes"] = o.Scopes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelAPIKey) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"id",
		"created_at",
		"updated_at",
		"account_id",
		"label",
		"public_key",
		"scopes",
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

	varModelAPIKey := _ModelAPIKey{}

	err = json.Unmarshal(data, &varModelAPIKey)

	if err != nil {
		return err
	}

	*o = ModelAPIKey(varModelAPIKey)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "last_used_at")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "label")
		delete(additionalProperties, "public_key")
		delete(additionalProperties, "secret_key")
		delete(additionalProperties, "scopes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelAPIKey struct {
	value *ModelAPIKey
	isSet bool
}

func (v NullableModelAPIKey) Get() *ModelAPIKey {
	return v.value
}

func (v *NullableModelAPIKey) Set(val *ModelAPIKey) {
	v.value = val
	v.isSet = true
}

func (v NullableModelAPIKey) IsSet() bool {
	return v.isSet
}

func (v *NullableModelAPIKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelAPIKey(val *ModelAPIKey) *NullableModelAPIKey {
	return &NullableModelAPIKey{value: val, isSet: true}
}

func (v NullableModelAPIKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelAPIKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
