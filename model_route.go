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

// checks if the Route type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Route{}

// Route struct for Route
type Route struct {
	// Object type identifier
	Object string `json:"object"`
	// Unique identifier for the route
	Id uuid.UUID `json:"id"`
	// When the route was created
	CreatedAt time.Time `json:"created_at"`
	// When the route was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// Route name
	Name string `json:"name"`
	// Webhook URL for the route
	Url string `json:"url"`
	// Recipient filter
	Recipient *string `json:"recipient,omitempty"`
	// Whether to include attachments in webhooks
	IncludeAttachments *bool `json:"include_attachments,omitempty"`
	// Whether to include headers in webhooks
	IncludeHeaders *bool `json:"include_headers,omitempty"`
	// Whether to group by message ID
	GroupByMessageId *bool `json:"group_by_message_id,omitempty"`
	// Whether to strip reply content
	StripReplies *bool `json:"strip_replies,omitempty"`
	// Whether the route is enabled
	Enabled              bool `json:"enabled"`
	AdditionalProperties map[string]interface{}
}

type _Route Route

// NewRoute instantiates a new Route object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoute(object string, id uuid.UUID, createdAt time.Time, updatedAt time.Time, name string, url string, enabled bool) *Route {
	this := Route{}
	this.Object = object
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Name = name
	this.Url = url
	this.Enabled = enabled
	return &this
}

// NewRouteWithDefaults instantiates a new Route object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteWithDefaults() *Route {
	this := Route{}
	return &this
}

// GetObject returns the Object field value
func (o *Route) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Route) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Route) SetObject(v string) {
	o.Object = v
}

// GetId returns the Id field value
func (o *Route) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Route) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Route) SetId(v uuid.UUID) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Route) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Route) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Route) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Route) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Route) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Route) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetName returns the Name field value
func (o *Route) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Route) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Route) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *Route) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Route) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Route) SetUrl(v string) {
	o.Url = v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *Route) GetRecipient() string {
	if o == nil || IsNil(o.Recipient) {
		var ret string
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetRecipientOk() (*string, bool) {
	if o == nil || IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *Route) HasRecipient() bool {
	if o != nil && !IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given string and assigns it to the Recipient field.
func (o *Route) SetRecipient(v string) {
	o.Recipient = &v
}

// GetIncludeAttachments returns the IncludeAttachments field value if set, zero value otherwise.
func (o *Route) GetIncludeAttachments() bool {
	if o == nil || IsNil(o.IncludeAttachments) {
		var ret bool
		return ret
	}
	return *o.IncludeAttachments
}

// GetIncludeAttachmentsOk returns a tuple with the IncludeAttachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetIncludeAttachmentsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeAttachments) {
		return nil, false
	}
	return o.IncludeAttachments, true
}

// HasIncludeAttachments returns a boolean if a field has been set.
func (o *Route) HasIncludeAttachments() bool {
	if o != nil && !IsNil(o.IncludeAttachments) {
		return true
	}

	return false
}

// SetIncludeAttachments gets a reference to the given bool and assigns it to the IncludeAttachments field.
func (o *Route) SetIncludeAttachments(v bool) {
	o.IncludeAttachments = &v
}

// GetIncludeHeaders returns the IncludeHeaders field value if set, zero value otherwise.
func (o *Route) GetIncludeHeaders() bool {
	if o == nil || IsNil(o.IncludeHeaders) {
		var ret bool
		return ret
	}
	return *o.IncludeHeaders
}

// GetIncludeHeadersOk returns a tuple with the IncludeHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetIncludeHeadersOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeHeaders) {
		return nil, false
	}
	return o.IncludeHeaders, true
}

// HasIncludeHeaders returns a boolean if a field has been set.
func (o *Route) HasIncludeHeaders() bool {
	if o != nil && !IsNil(o.IncludeHeaders) {
		return true
	}

	return false
}

// SetIncludeHeaders gets a reference to the given bool and assigns it to the IncludeHeaders field.
func (o *Route) SetIncludeHeaders(v bool) {
	o.IncludeHeaders = &v
}

// GetGroupByMessageId returns the GroupByMessageId field value if set, zero value otherwise.
func (o *Route) GetGroupByMessageId() bool {
	if o == nil || IsNil(o.GroupByMessageId) {
		var ret bool
		return ret
	}
	return *o.GroupByMessageId
}

// GetGroupByMessageIdOk returns a tuple with the GroupByMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetGroupByMessageIdOk() (*bool, bool) {
	if o == nil || IsNil(o.GroupByMessageId) {
		return nil, false
	}
	return o.GroupByMessageId, true
}

// HasGroupByMessageId returns a boolean if a field has been set.
func (o *Route) HasGroupByMessageId() bool {
	if o != nil && !IsNil(o.GroupByMessageId) {
		return true
	}

	return false
}

// SetGroupByMessageId gets a reference to the given bool and assigns it to the GroupByMessageId field.
func (o *Route) SetGroupByMessageId(v bool) {
	o.GroupByMessageId = &v
}

// GetStripReplies returns the StripReplies field value if set, zero value otherwise.
func (o *Route) GetStripReplies() bool {
	if o == nil || IsNil(o.StripReplies) {
		var ret bool
		return ret
	}
	return *o.StripReplies
}

// GetStripRepliesOk returns a tuple with the StripReplies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetStripRepliesOk() (*bool, bool) {
	if o == nil || IsNil(o.StripReplies) {
		return nil, false
	}
	return o.StripReplies, true
}

// HasStripReplies returns a boolean if a field has been set.
func (o *Route) HasStripReplies() bool {
	if o != nil && !IsNil(o.StripReplies) {
		return true
	}

	return false
}

// SetStripReplies gets a reference to the given bool and assigns it to the StripReplies field.
func (o *Route) SetStripReplies(v bool) {
	o.StripReplies = &v
}

// GetEnabled returns the Enabled field value
func (o *Route) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Route) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Route) SetEnabled(v bool) {
	o.Enabled = v
}

func (o Route) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Route) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["name"] = o.Name
	toSerialize["url"] = o.Url
	if !IsNil(o.Recipient) {
		toSerialize["recipient"] = o.Recipient
	}
	if !IsNil(o.IncludeAttachments) {
		toSerialize["include_attachments"] = o.IncludeAttachments
	}
	if !IsNil(o.IncludeHeaders) {
		toSerialize["include_headers"] = o.IncludeHeaders
	}
	if !IsNil(o.GroupByMessageId) {
		toSerialize["group_by_message_id"] = o.GroupByMessageId
	}
	if !IsNil(o.StripReplies) {
		toSerialize["strip_replies"] = o.StripReplies
	}
	toSerialize["enabled"] = o.Enabled

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Route) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"id",
		"created_at",
		"updated_at",
		"name",
		"url",
		"enabled",
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

	varRoute := _Route{}

	err = json.Unmarshal(data, &varRoute)

	if err != nil {
		return err
	}

	*o = Route(varRoute)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "name")
		delete(additionalProperties, "url")
		delete(additionalProperties, "recipient")
		delete(additionalProperties, "include_attachments")
		delete(additionalProperties, "include_headers")
		delete(additionalProperties, "group_by_message_id")
		delete(additionalProperties, "strip_replies")
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoute struct {
	value *Route
	isSet bool
}

func (v NullableRoute) Get() *Route {
	return v.value
}

func (v *NullableRoute) Set(val *Route) {
	v.value = val
	v.isSet = true
}

func (v NullableRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoute(val *Route) *NullableRoute {
	return &NullableRoute{value: val, isSet: true}
}

func (v NullableRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
