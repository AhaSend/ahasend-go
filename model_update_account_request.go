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

// checks if the UpdateAccountRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAccountRequest{}

// UpdateAccountRequest struct for UpdateAccountRequest
type UpdateAccountRequest struct {
	// Account name
	Name *string `json:"name,omitempty"`
	// Account website URL
	Website *string `json:"website,omitempty"`
	// Account description (used for account verification)
	About *string `json:"about,omitempty"`
	// Default open tracking setting
	TrackOpens *bool `json:"track_opens,omitempty"`
	// Default click tracking setting
	TrackClicks *bool `json:"track_clicks,omitempty"`
	// Whether to reject bad recipients
	RejectBadRecipients *bool `json:"reject_bad_recipients,omitempty"`
	// Whether to reject mistyped recipients
	RejectMistypedRecipients *bool `json:"reject_mistyped_recipients,omitempty"`
	// Default message metadata retention in days
	MessageMetadataRetention *int32 `json:"message_metadata_retention,omitempty"`
	// Default message data retention in days
	MessageDataRetention *int32 `json:"message_data_retention,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateAccountRequest UpdateAccountRequest

// NewUpdateAccountRequest instantiates a new UpdateAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAccountRequest() *UpdateAccountRequest {
	this := UpdateAccountRequest{}
	return &this
}

// NewUpdateAccountRequestWithDefaults instantiates a new UpdateAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAccountRequestWithDefaults() *UpdateAccountRequest {
	this := UpdateAccountRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateAccountRequest) SetName(v string) {
	o.Name = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetWebsite() string {
	if o == nil || IsNil(o.Website) {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *UpdateAccountRequest) SetWebsite(v string) {
	o.Website = &v
}

// GetAbout returns the About field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetAbout() string {
	if o == nil || IsNil(o.About) {
		var ret string
		return ret
	}
	return *o.About
}

// GetAboutOk returns a tuple with the About field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetAboutOk() (*string, bool) {
	if o == nil || IsNil(o.About) {
		return nil, false
	}
	return o.About, true
}

// HasAbout returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasAbout() bool {
	if o != nil && !IsNil(o.About) {
		return true
	}

	return false
}

// SetAbout gets a reference to the given string and assigns it to the About field.
func (o *UpdateAccountRequest) SetAbout(v string) {
	o.About = &v
}

// GetTrackOpens returns the TrackOpens field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetTrackOpens() bool {
	if o == nil || IsNil(o.TrackOpens) {
		var ret bool
		return ret
	}
	return *o.TrackOpens
}

// GetTrackOpensOk returns a tuple with the TrackOpens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetTrackOpensOk() (*bool, bool) {
	if o == nil || IsNil(o.TrackOpens) {
		return nil, false
	}
	return o.TrackOpens, true
}

// HasTrackOpens returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasTrackOpens() bool {
	if o != nil && !IsNil(o.TrackOpens) {
		return true
	}

	return false
}

// SetTrackOpens gets a reference to the given bool and assigns it to the TrackOpens field.
func (o *UpdateAccountRequest) SetTrackOpens(v bool) {
	o.TrackOpens = &v
}

// GetTrackClicks returns the TrackClicks field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetTrackClicks() bool {
	if o == nil || IsNil(o.TrackClicks) {
		var ret bool
		return ret
	}
	return *o.TrackClicks
}

// GetTrackClicksOk returns a tuple with the TrackClicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetTrackClicksOk() (*bool, bool) {
	if o == nil || IsNil(o.TrackClicks) {
		return nil, false
	}
	return o.TrackClicks, true
}

// HasTrackClicks returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasTrackClicks() bool {
	if o != nil && !IsNil(o.TrackClicks) {
		return true
	}

	return false
}

// SetTrackClicks gets a reference to the given bool and assigns it to the TrackClicks field.
func (o *UpdateAccountRequest) SetTrackClicks(v bool) {
	o.TrackClicks = &v
}

// GetRejectBadRecipients returns the RejectBadRecipients field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetRejectBadRecipients() bool {
	if o == nil || IsNil(o.RejectBadRecipients) {
		var ret bool
		return ret
	}
	return *o.RejectBadRecipients
}

// GetRejectBadRecipientsOk returns a tuple with the RejectBadRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetRejectBadRecipientsOk() (*bool, bool) {
	if o == nil || IsNil(o.RejectBadRecipients) {
		return nil, false
	}
	return o.RejectBadRecipients, true
}

// HasRejectBadRecipients returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasRejectBadRecipients() bool {
	if o != nil && !IsNil(o.RejectBadRecipients) {
		return true
	}

	return false
}

// SetRejectBadRecipients gets a reference to the given bool and assigns it to the RejectBadRecipients field.
func (o *UpdateAccountRequest) SetRejectBadRecipients(v bool) {
	o.RejectBadRecipients = &v
}

// GetRejectMistypedRecipients returns the RejectMistypedRecipients field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetRejectMistypedRecipients() bool {
	if o == nil || IsNil(o.RejectMistypedRecipients) {
		var ret bool
		return ret
	}
	return *o.RejectMistypedRecipients
}

// GetRejectMistypedRecipientsOk returns a tuple with the RejectMistypedRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetRejectMistypedRecipientsOk() (*bool, bool) {
	if o == nil || IsNil(o.RejectMistypedRecipients) {
		return nil, false
	}
	return o.RejectMistypedRecipients, true
}

// HasRejectMistypedRecipients returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasRejectMistypedRecipients() bool {
	if o != nil && !IsNil(o.RejectMistypedRecipients) {
		return true
	}

	return false
}

// SetRejectMistypedRecipients gets a reference to the given bool and assigns it to the RejectMistypedRecipients field.
func (o *UpdateAccountRequest) SetRejectMistypedRecipients(v bool) {
	o.RejectMistypedRecipients = &v
}

// GetMessageMetadataRetention returns the MessageMetadataRetention field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetMessageMetadataRetention() int32 {
	if o == nil || IsNil(o.MessageMetadataRetention) {
		var ret int32
		return ret
	}
	return *o.MessageMetadataRetention
}

// GetMessageMetadataRetentionOk returns a tuple with the MessageMetadataRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetMessageMetadataRetentionOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageMetadataRetention) {
		return nil, false
	}
	return o.MessageMetadataRetention, true
}

// HasMessageMetadataRetention returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasMessageMetadataRetention() bool {
	if o != nil && !IsNil(o.MessageMetadataRetention) {
		return true
	}

	return false
}

// SetMessageMetadataRetention gets a reference to the given int32 and assigns it to the MessageMetadataRetention field.
func (o *UpdateAccountRequest) SetMessageMetadataRetention(v int32) {
	o.MessageMetadataRetention = &v
}

// GetMessageDataRetention returns the MessageDataRetention field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetMessageDataRetention() int32 {
	if o == nil || IsNil(o.MessageDataRetention) {
		var ret int32
		return ret
	}
	return *o.MessageDataRetention
}

// GetMessageDataRetentionOk returns a tuple with the MessageDataRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetMessageDataRetentionOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageDataRetention) {
		return nil, false
	}
	return o.MessageDataRetention, true
}

// HasMessageDataRetention returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasMessageDataRetention() bool {
	if o != nil && !IsNil(o.MessageDataRetention) {
		return true
	}

	return false
}

// SetMessageDataRetention gets a reference to the given int32 and assigns it to the MessageDataRetention field.
func (o *UpdateAccountRequest) SetMessageDataRetention(v int32) {
	o.MessageDataRetention = &v
}

func (o UpdateAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAccountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Website) {
		toSerialize["website"] = o.Website
	}
	if !IsNil(o.About) {
		toSerialize["about"] = o.About
	}
	if !IsNil(o.TrackOpens) {
		toSerialize["track_opens"] = o.TrackOpens
	}
	if !IsNil(o.TrackClicks) {
		toSerialize["track_clicks"] = o.TrackClicks
	}
	if !IsNil(o.RejectBadRecipients) {
		toSerialize["reject_bad_recipients"] = o.RejectBadRecipients
	}
	if !IsNil(o.RejectMistypedRecipients) {
		toSerialize["reject_mistyped_recipients"] = o.RejectMistypedRecipients
	}
	if !IsNil(o.MessageMetadataRetention) {
		toSerialize["message_metadata_retention"] = o.MessageMetadataRetention
	}
	if !IsNil(o.MessageDataRetention) {
		toSerialize["message_data_retention"] = o.MessageDataRetention
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateAccountRequest) UnmarshalJSON(data []byte) (err error) {
	varUpdateAccountRequest := _UpdateAccountRequest{}

	err = json.Unmarshal(data, &varUpdateAccountRequest)

	if err != nil {
		return err
	}

	*o = UpdateAccountRequest(varUpdateAccountRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "website")
		delete(additionalProperties, "about")
		delete(additionalProperties, "track_opens")
		delete(additionalProperties, "track_clicks")
		delete(additionalProperties, "reject_bad_recipients")
		delete(additionalProperties, "reject_mistyped_recipients")
		delete(additionalProperties, "message_metadata_retention")
		delete(additionalProperties, "message_data_retention")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateAccountRequest struct {
	value *UpdateAccountRequest
	isSet bool
}

func (v NullableUpdateAccountRequest) Get() *UpdateAccountRequest {
	return v.value
}

func (v *NullableUpdateAccountRequest) Set(val *UpdateAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAccountRequest(val *UpdateAccountRequest) *NullableUpdateAccountRequest {
	return &NullableUpdateAccountRequest{value: val, isSet: true}
}

func (v NullableUpdateAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
