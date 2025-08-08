/*
AhaSend API v2

The AhaSend API v2 allows you to send transactional emails, manage domains, webhooks, routes, API keys, and view statistics.  ## Authentication All API requests must be authenticated using a Bearer token in the Authorization header: ``` Authorization: Bearer aha-sk-64-CHARACTER-RANDOM-STRING ```  ## Scopes API keys have specific scopes that control access to different resources and actions:  ### Message Scopes - `messages:send:all` - Send messages from any domain in the account - `messages:send:{domain}` - Send messages from a specific domain - `messages:cancel:all` - Cancel messages from any domain - `messages:cancel:{domain}` - Cancel messages from a specific domain - `messages:read:all` - Read messages from any domain - `messages:read:{domain}` - Read messages from a specific domain  ### Domain Scopes - `domains:read` - Read all domains - `domains:write` - Create and update domains - `domains:delete:all` - Delete any domain - `domains:delete:{domain}` - Delete a specific domain  ### Account Scopes - `accounts:read` - Read account information - `accounts:write` - Update account settings - `accounts:billing` - Access billing information - `accounts:members:read` - Read account members - `accounts:members:add` - Add account members - `accounts:members:update` - Update account members - `accounts:members:remove` - Remove account members  ### Webhook Scopes - `webhooks:read:all` - Read all webhooks - `webhooks:read:{domain}` - Read webhooks for a specific domain - `webhooks:write:all` - Create and update webhooks - `webhooks:write:{domain}` - Create and update webhooks for a specific domain - `webhooks:delete:all` - Delete any webhook - `webhooks:delete:{domain}` - Delete webhooks for a specific domain  ### Route Scopes - `routes:read:all` - Read all routes - `routes:read:{domain}` - Read routes for a specific domain - `routes:write:all` - Create and update routes - `routes:write:{domain}` - Create and update routes for a specific domain - `routes:delete:all` - Delete any route - `routes:delete:{domain}` - Delete routes for a specific domain  ### Suppression Scopes - `suppressions:read` - Read suppressions - `suppressions:write` - Create suppressions - `suppressions:delete` - Delete suppressions - `suppressions:wipe` - Delete all suppressions (dangerous)  ### SMTP Credentials Scopes - `smtp-credentials:read:all` - Read all SMTP credentials - `smtp-credentials:read:{domain}` - Read SMTP credentials for a specific domain - `smtp-credentials:write:all` - Create SMTP credentials - `smtp-credentials:write:{domain}` - Create SMTP credentials for a specific domain - `smtp-credentials:delete:all` - Delete any SMTP credentials - `smtp-credentials:delete:{domain}` - Delete SMTP credentials for a specific domain  ### Statistics Scopes - `statistics-transactional:read:all` - Read all transactional statistics - `statistics-transactional:read:{domain}` - Read transactional statistics for a specific domain  ### API Key Scopes - `api-keys:read` - Read API keys - `api-keys:write` - Create and update API keys - `api-keys:delete` - Delete API keys  ## Rate Limiting - General API endpoints: 100 requests per second, 200 burst - Statistics endpoints: 1 request per second, 1 burst  ## Pagination List endpoints use cursor-based pagination with the following parameters: - `limit`: Maximum number of items to return (default: 100, max: 100) - `cursor`: Pagination cursor for the next page  ## Time Formats All timestamps must be in RFC3339 format, e.g., `2023-12-25T10:30:00Z`  ## Idempotency POST requests support idempotency through the optional `Idempotency-Key` header. When provided: - The same request can be safely retried multiple times - Duplicate requests return the same response with `Idempotent-Replayed: true` - In-progress requests return HTTP 409 with `Idempotent-Replayed: false` - Failed requests return HTTP 412 with `Idempotent-Replayed: false` - Idempotency keys expire after 24 hours

API version: 2.0.0
Contact: support@ahasend.com
*/

package ahasend

import (
	"encoding/json"
)

// checks if the Tracking type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tracking{}

// Tracking struct for Tracking
type Tracking struct {
	// Whether to track opens
	Open *bool `json:"open,omitempty"`
	// Whether to track clicks
	Click                *bool `json:"click,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Tracking Tracking

// NewTracking instantiates a new Tracking object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTracking() *Tracking {
	this := Tracking{}
	return &this
}

// NewTrackingWithDefaults instantiates a new Tracking object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackingWithDefaults() *Tracking {
	this := Tracking{}
	return &this
}

// GetOpen returns the Open field value if set, zero value otherwise.
func (o *Tracking) GetOpen() bool {
	if o == nil || IsNil(o.Open) {
		var ret bool
		return ret
	}
	return *o.Open
}

// GetOpenOk returns a tuple with the Open field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tracking) GetOpenOk() (*bool, bool) {
	if o == nil || IsNil(o.Open) {
		return nil, false
	}
	return o.Open, true
}

// HasOpen returns a boolean if a field has been set.
func (o *Tracking) HasOpen() bool {
	if o != nil && !IsNil(o.Open) {
		return true
	}

	return false
}

// SetOpen gets a reference to the given bool and assigns it to the Open field.
func (o *Tracking) SetOpen(v bool) {
	o.Open = &v
}

// GetClick returns the Click field value if set, zero value otherwise.
func (o *Tracking) GetClick() bool {
	if o == nil || IsNil(o.Click) {
		var ret bool
		return ret
	}
	return *o.Click
}

// GetClickOk returns a tuple with the Click field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tracking) GetClickOk() (*bool, bool) {
	if o == nil || IsNil(o.Click) {
		return nil, false
	}
	return o.Click, true
}

// HasClick returns a boolean if a field has been set.
func (o *Tracking) HasClick() bool {
	if o != nil && !IsNil(o.Click) {
		return true
	}

	return false
}

// SetClick gets a reference to the given bool and assigns it to the Click field.
func (o *Tracking) SetClick(v bool) {
	o.Click = &v
}

func (o Tracking) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tracking) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Open) {
		toSerialize["open"] = o.Open
	}
	if !IsNil(o.Click) {
		toSerialize["click"] = o.Click
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Tracking) UnmarshalJSON(data []byte) (err error) {
	varTracking := _Tracking{}

	err = json.Unmarshal(data, &varTracking)

	if err != nil {
		return err
	}

	*o = Tracking(varTracking)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "open")
		delete(additionalProperties, "click")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTracking struct {
	value *Tracking
	isSet bool
}

func (v NullableTracking) Get() *Tracking {
	return v.value
}

func (v *NullableTracking) Set(val *Tracking) {
	v.value = val
	v.isSet = true
}

func (v NullableTracking) IsSet() bool {
	return v.isSet
}

func (v *NullableTracking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTracking(val *Tracking) *NullableTracking {
	return &NullableTracking{value: val, isSet: true}
}

func (v NullableTracking) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTracking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
