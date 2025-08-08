/*
AhaSend API v2

The AhaSend API v2 allows you to send transactional emails, manage domains, webhooks, routes, API keys, and view statistics.  ## Authentication All API requests must be authenticated using a Bearer token in the Authorization header: ``` Authorization: Bearer aha-sk-64-CHARACTER-RANDOM-STRING ```  ## Scopes API keys have specific scopes that control access to different resources and actions:  ### Message Scopes - `messages:send:all` - Send messages from any domain in the account - `messages:send:{domain}` - Send messages from a specific domain - `messages:cancel:all` - Cancel messages from any domain - `messages:cancel:{domain}` - Cancel messages from a specific domain - `messages:read:all` - Read messages from any domain - `messages:read:{domain}` - Read messages from a specific domain  ### Domain Scopes - `domains:read` - Read all domains - `domains:write` - Create and update domains - `domains:delete:all` - Delete any domain - `domains:delete:{domain}` - Delete a specific domain  ### Account Scopes - `accounts:read` - Read account information - `accounts:write` - Update account settings - `accounts:billing` - Access billing information - `accounts:members:read` - Read account members - `accounts:members:add` - Add account members - `accounts:members:update` - Update account members - `accounts:members:remove` - Remove account members  ### Webhook Scopes - `webhooks:read:all` - Read all webhooks - `webhooks:read:{domain}` - Read webhooks for a specific domain - `webhooks:write:all` - Create and update webhooks - `webhooks:write:{domain}` - Create and update webhooks for a specific domain - `webhooks:delete:all` - Delete any webhook - `webhooks:delete:{domain}` - Delete webhooks for a specific domain  ### Route Scopes - `routes:read:all` - Read all routes - `routes:read:{domain}` - Read routes for a specific domain - `routes:write:all` - Create and update routes - `routes:write:{domain}` - Create and update routes for a specific domain - `routes:delete:all` - Delete any route - `routes:delete:{domain}` - Delete routes for a specific domain  ### Suppression Scopes - `suppressions:read` - Read suppressions - `suppressions:write` - Create suppressions - `suppressions:delete` - Delete suppressions - `suppressions:wipe` - Delete all suppressions (dangerous)  ### SMTP Credentials Scopes - `smtp-credentials:read:all` - Read all SMTP credentials - `smtp-credentials:read:{domain}` - Read SMTP credentials for a specific domain - `smtp-credentials:write:all` - Create SMTP credentials - `smtp-credentials:write:{domain}` - Create SMTP credentials for a specific domain - `smtp-credentials:delete:all` - Delete any SMTP credentials - `smtp-credentials:delete:{domain}` - Delete SMTP credentials for a specific domain  ### Statistics Scopes - `statistics-transactional:read:all` - Read all transactional statistics - `statistics-transactional:read:{domain}` - Read transactional statistics for a specific domain  ### API Key Scopes - `api-keys:read` - Read API keys - `api-keys:write` - Create and update API keys - `api-keys:delete` - Delete API keys  ## Rate Limiting - General API endpoints: 100 requests per second, 200 burst - Statistics endpoints: 1 request per second, 1 burst  ## Pagination List endpoints use cursor-based pagination with the following parameters: - `limit`: Maximum number of items to return (default: 100, max: 100) - `cursor`: Pagination cursor for the next page  ## Time Formats All timestamps must be in RFC3339 format, e.g., `2023-12-25T10:30:00Z`  ## Idempotency POST requests support idempotency through the optional `Idempotency-Key` header. When provided: - The same request can be safely retried multiple times - Duplicate requests return the same response with `Idempotent-Replayed: true` - In-progress requests return HTTP 409 with `Idempotent-Replayed: false` - Failed requests return HTTP 412 with `Idempotent-Replayed: false` - Idempotency keys expire after 24 hours

API version: 2.0.0
Contact: support@ahasend.com
*/

package ahasend

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"time"
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
