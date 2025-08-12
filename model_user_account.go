package ahasend

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// checks if the UserAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserAccount{}

// UserAccount struct for UserAccount
type UserAccount struct {
	// Unique identifier for the user account relationship
	Id uuid.UUID `json:"id"`
	// When the relationship was created
	CreatedAt time.Time `json:"created_at"`
	// When the relationship was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// User ID
	UserId uuid.UUID `json:"user_id"`
	// Account ID
	AccountId uuid.UUID `json:"account_id"`
	// User role in the account
	Role                 string `json:"role"`
	AdditionalProperties map[string]interface{}
}

type _UserAccount UserAccount

// NewUserAccount instantiates a new UserAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAccount(id uuid.UUID, createdAt time.Time, updatedAt time.Time, userId uuid.UUID, accountId uuid.UUID, role string) *UserAccount {
	this := UserAccount{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.UserId = userId
	this.AccountId = accountId
	this.Role = role
	return &this
}

// NewUserAccountWithDefaults instantiates a new UserAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAccountWithDefaults() *UserAccount {
	this := UserAccount{}
	return &this
}

// GetId returns the Id field value
func (o *UserAccount) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserAccount) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserAccount) SetId(v uuid.UUID) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *UserAccount) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *UserAccount) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *UserAccount) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *UserAccount) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *UserAccount) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *UserAccount) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetUserId returns the UserId field value
func (o *UserAccount) GetUserId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UserAccount) GetUserIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UserAccount) SetUserId(v uuid.UUID) {
	o.UserId = v
}

// GetAccountId returns the AccountId field value
func (o *UserAccount) GetAccountId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *UserAccount) GetAccountIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *UserAccount) SetAccountId(v uuid.UUID) {
	o.AccountId = v
}

// GetRole returns the Role field value
func (o *UserAccount) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *UserAccount) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *UserAccount) SetRole(v string) {
	o.Role = v
}

func (o UserAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["user_id"] = o.UserId
	toSerialize["account_id"] = o.AccountId
	toSerialize["role"] = o.Role

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
		"user_id",
		"account_id",
		"role",
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

	varUserAccount := _UserAccount{}

	err = json.Unmarshal(data, &varUserAccount)

	if err != nil {
		return err
	}

	*o = UserAccount(varUserAccount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserAccount struct {
	value *UserAccount
	isSet bool
}

func (v NullableUserAccount) Get() *UserAccount {
	return v.value
}

func (v *NullableUserAccount) Set(val *UserAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAccount(val *UserAccount) *NullableUserAccount {
	return &NullableUserAccount{value: val, isSet: true}
}

func (v NullableUserAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
