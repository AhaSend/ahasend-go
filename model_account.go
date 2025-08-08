package ahasend

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// checks if the Account type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Account{}

// Account struct for Account
type Account struct {
	// Object type identifier
	Object string `json:"object"`
	// Unique identifier for the account
	Id uuid.UUID `json:"id"`
	// When the account was created
	CreatedAt time.Time `json:"created_at"`
	// When the account was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// Account name
	Name string `json:"name"`
	// Account website URL
	Website *string `json:"website,omitempty"`
	// Account description
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

type _Account Account

// NewAccount instantiates a new Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccount(object string, id uuid.UUID, createdAt time.Time, updatedAt time.Time, name string) *Account {
	this := Account{}
	this.Object = object
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Name = name
	return &this
}

// NewAccountWithDefaults instantiates a new Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountWithDefaults() *Account {
	this := Account{}
	return &this
}

// GetObject returns the Object field value
func (o *Account) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Account) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Account) SetObject(v string) {
	o.Object = v
}

// GetId returns the Id field value
func (o *Account) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Account) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Account) SetId(v uuid.UUID) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Account) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Account) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Account) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Account) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Account) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Account) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetName returns the Name field value
func (o *Account) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Account) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Account) SetName(v string) {
	o.Name = v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *Account) GetWebsite() string {
	if o == nil || IsNil(o.Website) {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *Account) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *Account) SetWebsite(v string) {
	o.Website = &v
}

// GetAbout returns the About field value if set, zero value otherwise.
func (o *Account) GetAbout() string {
	if o == nil || IsNil(o.About) {
		var ret string
		return ret
	}
	return *o.About
}

// GetAboutOk returns a tuple with the About field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetAboutOk() (*string, bool) {
	if o == nil || IsNil(o.About) {
		return nil, false
	}
	return o.About, true
}

// HasAbout returns a boolean if a field has been set.
func (o *Account) HasAbout() bool {
	if o != nil && !IsNil(o.About) {
		return true
	}

	return false
}

// SetAbout gets a reference to the given string and assigns it to the About field.
func (o *Account) SetAbout(v string) {
	o.About = &v
}

// GetTrackOpens returns the TrackOpens field value if set, zero value otherwise.
func (o *Account) GetTrackOpens() bool {
	if o == nil || IsNil(o.TrackOpens) {
		var ret bool
		return ret
	}
	return *o.TrackOpens
}

// GetTrackOpensOk returns a tuple with the TrackOpens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetTrackOpensOk() (*bool, bool) {
	if o == nil || IsNil(o.TrackOpens) {
		return nil, false
	}
	return o.TrackOpens, true
}

// HasTrackOpens returns a boolean if a field has been set.
func (o *Account) HasTrackOpens() bool {
	if o != nil && !IsNil(o.TrackOpens) {
		return true
	}

	return false
}

// SetTrackOpens gets a reference to the given bool and assigns it to the TrackOpens field.
func (o *Account) SetTrackOpens(v bool) {
	o.TrackOpens = &v
}

// GetTrackClicks returns the TrackClicks field value if set, zero value otherwise.
func (o *Account) GetTrackClicks() bool {
	if o == nil || IsNil(o.TrackClicks) {
		var ret bool
		return ret
	}
	return *o.TrackClicks
}

// GetTrackClicksOk returns a tuple with the TrackClicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetTrackClicksOk() (*bool, bool) {
	if o == nil || IsNil(o.TrackClicks) {
		return nil, false
	}
	return o.TrackClicks, true
}

// HasTrackClicks returns a boolean if a field has been set.
func (o *Account) HasTrackClicks() bool {
	if o != nil && !IsNil(o.TrackClicks) {
		return true
	}

	return false
}

// SetTrackClicks gets a reference to the given bool and assigns it to the TrackClicks field.
func (o *Account) SetTrackClicks(v bool) {
	o.TrackClicks = &v
}

// GetRejectBadRecipients returns the RejectBadRecipients field value if set, zero value otherwise.
func (o *Account) GetRejectBadRecipients() bool {
	if o == nil || IsNil(o.RejectBadRecipients) {
		var ret bool
		return ret
	}
	return *o.RejectBadRecipients
}

// GetRejectBadRecipientsOk returns a tuple with the RejectBadRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetRejectBadRecipientsOk() (*bool, bool) {
	if o == nil || IsNil(o.RejectBadRecipients) {
		return nil, false
	}
	return o.RejectBadRecipients, true
}

// HasRejectBadRecipients returns a boolean if a field has been set.
func (o *Account) HasRejectBadRecipients() bool {
	if o != nil && !IsNil(o.RejectBadRecipients) {
		return true
	}

	return false
}

// SetRejectBadRecipients gets a reference to the given bool and assigns it to the RejectBadRecipients field.
func (o *Account) SetRejectBadRecipients(v bool) {
	o.RejectBadRecipients = &v
}

// GetRejectMistypedRecipients returns the RejectMistypedRecipients field value if set, zero value otherwise.
func (o *Account) GetRejectMistypedRecipients() bool {
	if o == nil || IsNil(o.RejectMistypedRecipients) {
		var ret bool
		return ret
	}
	return *o.RejectMistypedRecipients
}

// GetRejectMistypedRecipientsOk returns a tuple with the RejectMistypedRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetRejectMistypedRecipientsOk() (*bool, bool) {
	if o == nil || IsNil(o.RejectMistypedRecipients) {
		return nil, false
	}
	return o.RejectMistypedRecipients, true
}

// HasRejectMistypedRecipients returns a boolean if a field has been set.
func (o *Account) HasRejectMistypedRecipients() bool {
	if o != nil && !IsNil(o.RejectMistypedRecipients) {
		return true
	}

	return false
}

// SetRejectMistypedRecipients gets a reference to the given bool and assigns it to the RejectMistypedRecipients field.
func (o *Account) SetRejectMistypedRecipients(v bool) {
	o.RejectMistypedRecipients = &v
}

// GetMessageMetadataRetention returns the MessageMetadataRetention field value if set, zero value otherwise.
func (o *Account) GetMessageMetadataRetention() int32 {
	if o == nil || IsNil(o.MessageMetadataRetention) {
		var ret int32
		return ret
	}
	return *o.MessageMetadataRetention
}

// GetMessageMetadataRetentionOk returns a tuple with the MessageMetadataRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetMessageMetadataRetentionOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageMetadataRetention) {
		return nil, false
	}
	return o.MessageMetadataRetention, true
}

// HasMessageMetadataRetention returns a boolean if a field has been set.
func (o *Account) HasMessageMetadataRetention() bool {
	if o != nil && !IsNil(o.MessageMetadataRetention) {
		return true
	}

	return false
}

// SetMessageMetadataRetention gets a reference to the given int32 and assigns it to the MessageMetadataRetention field.
func (o *Account) SetMessageMetadataRetention(v int32) {
	o.MessageMetadataRetention = &v
}

// GetMessageDataRetention returns the MessageDataRetention field value if set, zero value otherwise.
func (o *Account) GetMessageDataRetention() int32 {
	if o == nil || IsNil(o.MessageDataRetention) {
		var ret int32
		return ret
	}
	return *o.MessageDataRetention
}

// GetMessageDataRetentionOk returns a tuple with the MessageDataRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetMessageDataRetentionOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageDataRetention) {
		return nil, false
	}
	return o.MessageDataRetention, true
}

// HasMessageDataRetention returns a boolean if a field has been set.
func (o *Account) HasMessageDataRetention() bool {
	if o != nil && !IsNil(o.MessageDataRetention) {
		return true
	}

	return false
}

// SetMessageDataRetention gets a reference to the given int32 and assigns it to the MessageDataRetention field.
func (o *Account) SetMessageDataRetention(v int32) {
	o.MessageDataRetention = &v
}

func (o Account) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Account) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["name"] = o.Name
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

func (o *Account) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"id",
		"created_at",
		"updated_at",
		"name",
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

	varAccount := _Account{}

	err = json.Unmarshal(data, &varAccount)

	if err != nil {
		return err
	}

	*o = Account(varAccount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
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

type NullableAccount struct {
	value *Account
	isSet bool
}

func (v NullableAccount) Get() *Account {
	return v.value
}

func (v *NullableAccount) Set(val *Account) {
	v.value = val
	v.isSet = true
}

func (v NullableAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccount(val *Account) *NullableAccount {
	return &NullableAccount{value: val, isSet: true}
}

func (v NullableAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
