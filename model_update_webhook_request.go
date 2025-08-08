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

// checks if the UpdateWebhookRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateWebhookRequest{}

// UpdateWebhookRequest struct for UpdateWebhookRequest
type UpdateWebhookRequest struct {
	// Webhook name
	Name *string `json:"name,omitempty"`
	// Webhook URL
	Url *string `json:"url,omitempty"`
	// Whether the webhook is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Trigger on message reception
	OnReception *bool `json:"on_reception,omitempty"`
	// Trigger on message delivery
	OnDelivered *bool `json:"on_delivered,omitempty"`
	// Trigger on transient errors
	OnTransientError *bool `json:"on_transient_error,omitempty"`
	// Trigger on permanent failures
	OnFailed *bool `json:"on_failed,omitempty"`
	// Trigger on bounces
	OnBounced *bool `json:"on_bounced,omitempty"`
	// Trigger on suppressions
	OnSuppressed *bool `json:"on_suppressed,omitempty"`
	// Trigger on opens
	OnOpened *bool `json:"on_opened,omitempty"`
	// Trigger on clicks
	OnClicked *bool `json:"on_clicked,omitempty"`
	// Trigger on suppression creation
	OnSuppressionCreated *bool `json:"on_suppression_created,omitempty"`
	// Trigger on DNS errors
	OnDnsError *bool `json:"on_dns_error,omitempty"`
	// Webhook scope
	Scope *string `json:"scope,omitempty"`
	// Domains this webhook applies to
	Domains              []string `json:"domains,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateWebhookRequest UpdateWebhookRequest

// NewUpdateWebhookRequest instantiates a new UpdateWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWebhookRequest() *UpdateWebhookRequest {
	this := UpdateWebhookRequest{}
	return &this
}

// NewUpdateWebhookRequestWithDefaults instantiates a new UpdateWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWebhookRequestWithDefaults() *UpdateWebhookRequest {
	this := UpdateWebhookRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateWebhookRequest) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UpdateWebhookRequest) SetUrl(v string) {
	o.Url = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateWebhookRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetOnReception returns the OnReception field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetOnReception() bool {
	if o == nil || IsNil(o.OnReception) {
		var ret bool
		return ret
	}
	return *o.OnReception
}

// GetOnReceptionOk returns a tuple with the OnReception field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetOnReceptionOk() (*bool, bool) {
	if o == nil || IsNil(o.OnReception) {
		return nil, false
	}
	return o.OnReception, true
}

// HasOnReception returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasOnReception() bool {
	if o != nil && !IsNil(o.OnReception) {
		return true
	}

	return false
}

// SetOnReception gets a reference to the given bool and assigns it to the OnReception field.
func (o *UpdateWebhookRequest) SetOnReception(v bool) {
	o.OnReception = &v
}

// GetOnDelivered returns the OnDelivered field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetOnDelivered() bool {
	if o == nil || IsNil(o.OnDelivered) {
		var ret bool
		return ret
	}
	return *o.OnDelivered
}

// GetOnDeliveredOk returns a tuple with the OnDelivered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetOnDeliveredOk() (*bool, bool) {
	if o == nil || IsNil(o.OnDelivered) {
		return nil, false
	}
	return o.OnDelivered, true
}

// HasOnDelivered returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasOnDelivered() bool {
	if o != nil && !IsNil(o.OnDelivered) {
		return true
	}

	return false
}

// SetOnDelivered gets a reference to the given bool and assigns it to the OnDelivered field.
func (o *UpdateWebhookRequest) SetOnDelivered(v bool) {
	o.OnDelivered = &v
}

// GetOnTransientError returns the OnTransientError field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetOnTransientError() bool {
	if o == nil || IsNil(o.OnTransientError) {
		var ret bool
		return ret
	}
	return *o.OnTransientError
}

// GetOnTransientErrorOk returns a tuple with the OnTransientError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetOnTransientErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.OnTransientError) {
		return nil, false
	}
	return o.OnTransientError, true
}

// HasOnTransientError returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasOnTransientError() bool {
	if o != nil && !IsNil(o.OnTransientError) {
		return true
	}

	return false
}

// SetOnTransientError gets a reference to the given bool and assigns it to the OnTransientError field.
func (o *UpdateWebhookRequest) SetOnTransientError(v bool) {
	o.OnTransientError = &v
}

// GetOnFailed returns the OnFailed field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetOnFailed() bool {
	if o == nil || IsNil(o.OnFailed) {
		var ret bool
		return ret
	}
	return *o.OnFailed
}

// GetOnFailedOk returns a tuple with the OnFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetOnFailedOk() (*bool, bool) {
	if o == nil || IsNil(o.OnFailed) {
		return nil, false
	}
	return o.OnFailed, true
}

// HasOnFailed returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasOnFailed() bool {
	if o != nil && !IsNil(o.OnFailed) {
		return true
	}

	return false
}

// SetOnFailed gets a reference to the given bool and assigns it to the OnFailed field.
func (o *UpdateWebhookRequest) SetOnFailed(v bool) {
	o.OnFailed = &v
}

// GetOnBounced returns the OnBounced field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetOnBounced() bool {
	if o == nil || IsNil(o.OnBounced) {
		var ret bool
		return ret
	}
	return *o.OnBounced
}

// GetOnBouncedOk returns a tuple with the OnBounced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetOnBouncedOk() (*bool, bool) {
	if o == nil || IsNil(o.OnBounced) {
		return nil, false
	}
	return o.OnBounced, true
}

// HasOnBounced returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasOnBounced() bool {
	if o != nil && !IsNil(o.OnBounced) {
		return true
	}

	return false
}

// SetOnBounced gets a reference to the given bool and assigns it to the OnBounced field.
func (o *UpdateWebhookRequest) SetOnBounced(v bool) {
	o.OnBounced = &v
}

// GetOnSuppressed returns the OnSuppressed field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetOnSuppressed() bool {
	if o == nil || IsNil(o.OnSuppressed) {
		var ret bool
		return ret
	}
	return *o.OnSuppressed
}

// GetOnSuppressedOk returns a tuple with the OnSuppressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetOnSuppressedOk() (*bool, bool) {
	if o == nil || IsNil(o.OnSuppressed) {
		return nil, false
	}
	return o.OnSuppressed, true
}

// HasOnSuppressed returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasOnSuppressed() bool {
	if o != nil && !IsNil(o.OnSuppressed) {
		return true
	}

	return false
}

// SetOnSuppressed gets a reference to the given bool and assigns it to the OnSuppressed field.
func (o *UpdateWebhookRequest) SetOnSuppressed(v bool) {
	o.OnSuppressed = &v
}

// GetOnOpened returns the OnOpened field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetOnOpened() bool {
	if o == nil || IsNil(o.OnOpened) {
		var ret bool
		return ret
	}
	return *o.OnOpened
}

// GetOnOpenedOk returns a tuple with the OnOpened field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetOnOpenedOk() (*bool, bool) {
	if o == nil || IsNil(o.OnOpened) {
		return nil, false
	}
	return o.OnOpened, true
}

// HasOnOpened returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasOnOpened() bool {
	if o != nil && !IsNil(o.OnOpened) {
		return true
	}

	return false
}

// SetOnOpened gets a reference to the given bool and assigns it to the OnOpened field.
func (o *UpdateWebhookRequest) SetOnOpened(v bool) {
	o.OnOpened = &v
}

// GetOnClicked returns the OnClicked field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetOnClicked() bool {
	if o == nil || IsNil(o.OnClicked) {
		var ret bool
		return ret
	}
	return *o.OnClicked
}

// GetOnClickedOk returns a tuple with the OnClicked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetOnClickedOk() (*bool, bool) {
	if o == nil || IsNil(o.OnClicked) {
		return nil, false
	}
	return o.OnClicked, true
}

// HasOnClicked returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasOnClicked() bool {
	if o != nil && !IsNil(o.OnClicked) {
		return true
	}

	return false
}

// SetOnClicked gets a reference to the given bool and assigns it to the OnClicked field.
func (o *UpdateWebhookRequest) SetOnClicked(v bool) {
	o.OnClicked = &v
}

// GetOnSuppressionCreated returns the OnSuppressionCreated field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetOnSuppressionCreated() bool {
	if o == nil || IsNil(o.OnSuppressionCreated) {
		var ret bool
		return ret
	}
	return *o.OnSuppressionCreated
}

// GetOnSuppressionCreatedOk returns a tuple with the OnSuppressionCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetOnSuppressionCreatedOk() (*bool, bool) {
	if o == nil || IsNil(o.OnSuppressionCreated) {
		return nil, false
	}
	return o.OnSuppressionCreated, true
}

// HasOnSuppressionCreated returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasOnSuppressionCreated() bool {
	if o != nil && !IsNil(o.OnSuppressionCreated) {
		return true
	}

	return false
}

// SetOnSuppressionCreated gets a reference to the given bool and assigns it to the OnSuppressionCreated field.
func (o *UpdateWebhookRequest) SetOnSuppressionCreated(v bool) {
	o.OnSuppressionCreated = &v
}

// GetOnDnsError returns the OnDnsError field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetOnDnsError() bool {
	if o == nil || IsNil(o.OnDnsError) {
		var ret bool
		return ret
	}
	return *o.OnDnsError
}

// GetOnDnsErrorOk returns a tuple with the OnDnsError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetOnDnsErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.OnDnsError) {
		return nil, false
	}
	return o.OnDnsError, true
}

// HasOnDnsError returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasOnDnsError() bool {
	if o != nil && !IsNil(o.OnDnsError) {
		return true
	}

	return false
}

// SetOnDnsError gets a reference to the given bool and assigns it to the OnDnsError field.
func (o *UpdateWebhookRequest) SetOnDnsError(v bool) {
	o.OnDnsError = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *UpdateWebhookRequest) SetScope(v string) {
	o.Scope = &v
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetDomains() []string {
	if o == nil || IsNil(o.Domains) {
		var ret []string
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.Domains) {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasDomains() bool {
	if o != nil && !IsNil(o.Domains) {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []string and assigns it to the Domains field.
func (o *UpdateWebhookRequest) SetDomains(v []string) {
	o.Domains = v
}

func (o UpdateWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateWebhookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.OnReception) {
		toSerialize["on_reception"] = o.OnReception
	}
	if !IsNil(o.OnDelivered) {
		toSerialize["on_delivered"] = o.OnDelivered
	}
	if !IsNil(o.OnTransientError) {
		toSerialize["on_transient_error"] = o.OnTransientError
	}
	if !IsNil(o.OnFailed) {
		toSerialize["on_failed"] = o.OnFailed
	}
	if !IsNil(o.OnBounced) {
		toSerialize["on_bounced"] = o.OnBounced
	}
	if !IsNil(o.OnSuppressed) {
		toSerialize["on_suppressed"] = o.OnSuppressed
	}
	if !IsNil(o.OnOpened) {
		toSerialize["on_opened"] = o.OnOpened
	}
	if !IsNil(o.OnClicked) {
		toSerialize["on_clicked"] = o.OnClicked
	}
	if !IsNil(o.OnSuppressionCreated) {
		toSerialize["on_suppression_created"] = o.OnSuppressionCreated
	}
	if !IsNil(o.OnDnsError) {
		toSerialize["on_dns_error"] = o.OnDnsError
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.Domains) {
		toSerialize["domains"] = o.Domains
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateWebhookRequest) UnmarshalJSON(data []byte) (err error) {
	varUpdateWebhookRequest := _UpdateWebhookRequest{}

	err = json.Unmarshal(data, &varUpdateWebhookRequest)

	if err != nil {
		return err
	}

	*o = UpdateWebhookRequest(varUpdateWebhookRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "url")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "on_reception")
		delete(additionalProperties, "on_delivered")
		delete(additionalProperties, "on_transient_error")
		delete(additionalProperties, "on_failed")
		delete(additionalProperties, "on_bounced")
		delete(additionalProperties, "on_suppressed")
		delete(additionalProperties, "on_opened")
		delete(additionalProperties, "on_clicked")
		delete(additionalProperties, "on_suppression_created")
		delete(additionalProperties, "on_dns_error")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "domains")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateWebhookRequest struct {
	value *UpdateWebhookRequest
	isSet bool
}

func (v NullableUpdateWebhookRequest) Get() *UpdateWebhookRequest {
	return v.value
}

func (v *NullableUpdateWebhookRequest) Set(val *UpdateWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWebhookRequest(val *UpdateWebhookRequest) *NullableUpdateWebhookRequest {
	return &NullableUpdateWebhookRequest{value: val, isSet: true}
}

func (v NullableUpdateWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
