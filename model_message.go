package ahasend

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// checks if the Message type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Message{}

// Message struct for Message
type Message struct {
	// Object type identifier
	Object string `json:"object"`
	// When the message was created
	CreatedAt time.Time `json:"created_at"`
	// When the message was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// When the message was sent
	SentAt *time.Time `json:"sent_at,omitempty"`
	// When the message was delivered
	DeliveredAt *time.Time `json:"delivered_at,omitempty"`
	// When the message data will be purged
	RetainUntil time.Time `json:"retain_until"`
	// Message direction
	Direction string `json:"direction"`
	// Whether this is a bounce notification
	IsBounceNotification bool `json:"is_bounce_notification"`
	// Classification of bounce if applicable
	BounceClassification *string `json:"bounce_classification,omitempty"`
	// List of delivery attempts for this message
	DeliveryAttempts []DeliveryEvent `json:"delivery_attempts"`
	// Message-ID header value
	MessageId string `json:"message_id"`
	// API-generated message ID
	ApiId uuid.UUID `json:"api_id"`
	// Internal AhaSend message ID
	AhasendId string `json:"ahasend_id"`
	// Message subject
	Subject string `json:"subject"`
	// Tags associated with the message
	Tags []string `json:"tags"`
	// Sender email address
	Sender string `json:"sender"`
	// Recipient email address
	Recipient string `json:"recipient"`
	// Current message status
	Status string `json:"status"`
	// Number of delivery attempts
	NumAttempts int32 `json:"num_attempts"`
	// Number of clicks tracked for this message
	ClickCount int32 `json:"click_count"`
	// Number of opens tracked for this message
	OpenCount int32 `json:"open_count"`
	// ID of the original message (for bounce messages)
	ReferenceMessageId *int64 `json:"reference_message_id,omitempty"`
	// Domain ID this message was sent from
	DomainId uuid.UUID `json:"domain_id"`
	// Account ID this message belongs to
	AccountId            uuid.UUID `json:"account_id"`
	AdditionalProperties map[string]interface{}
}

type _Message Message

// NewMessage instantiates a new Message object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessage(object string, createdAt time.Time, updatedAt time.Time, retainUntil time.Time, direction string, isBounceNotification bool, deliveryAttempts []DeliveryEvent, messageId string, apiId uuid.UUID, ahasendId string, subject string, tags []string, sender string, recipient string, status string, numAttempts int32, clickCount int32, openCount int32, domainId uuid.UUID, accountId uuid.UUID) *Message {
	this := Message{}
	this.Object = object
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.RetainUntil = retainUntil
	this.Direction = direction
	this.IsBounceNotification = isBounceNotification
	this.DeliveryAttempts = deliveryAttempts
	this.MessageId = messageId
	this.ApiId = apiId
	this.AhasendId = ahasendId
	this.Subject = subject
	this.Tags = tags
	this.Sender = sender
	this.Recipient = recipient
	this.Status = status
	this.NumAttempts = numAttempts
	this.ClickCount = clickCount
	this.OpenCount = openCount
	this.DomainId = domainId
	this.AccountId = accountId
	return &this
}

// NewMessageWithDefaults instantiates a new Message object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageWithDefaults() *Message {
	this := Message{}
	return &this
}

// GetObject returns the Object field value
func (o *Message) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Message) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Message) SetObject(v string) {
	o.Object = v
}

// GetDeliveryAttempts returns the DeliveryAttempts field value
func (o *Message) GetDeliveryAttempts() []DeliveryEvent {
	if o == nil {
		var ret []DeliveryEvent
		return ret
	}

	return o.DeliveryAttempts
}

// GetDeliveryAttemptsOk returns a tuple with the DeliveryAttempts field value
// and a boolean to check if the value has been set.
func (o *Message) GetDeliveryAttemptsOk() ([]DeliveryEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeliveryAttempts, true
}

// SetDeliveryAttempts sets field value
func (o *Message) SetDeliveryAttempts(v []DeliveryEvent) {
	o.DeliveryAttempts = v
}

// GetMessageId returns the MessageId field value
func (o *Message) GetMessageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *Message) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *Message) SetMessageId(v string) {
	o.MessageId = v
}

// GetApiId returns the ApiId field value
func (o *Message) GetApiId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}

	return o.ApiId
}

// GetApiIdOk returns a tuple with the ApiId field value
// and a boolean to check if the value has been set.
func (o *Message) GetApiIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiId, true
}

// SetApiId sets field value
func (o *Message) SetApiId(v uuid.UUID) {
	o.ApiId = v
}

// GetAhasendId returns the AhasendId field value
func (o *Message) GetAhasendId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AhasendId
}

// GetAhasendIdOk returns a tuple with the AhasendId field value
// and a boolean to check if the value has been set.
func (o *Message) GetAhasendIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AhasendId, true
}

// SetAhasendId sets field value
func (o *Message) SetAhasendId(v string) {
	o.AhasendId = v
}

// GetSubject returns the Subject field value
func (o *Message) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *Message) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *Message) SetSubject(v string) {
	o.Subject = v
}

// GetTags returns the Tags field value
func (o *Message) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *Message) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *Message) SetTags(v []string) {
	o.Tags = v
}

// GetSender returns the Sender field value
func (o *Message) GetSender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *Message) GetSenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *Message) SetSender(v string) {
	o.Sender = v
}

// GetRecipient returns the Recipient field value
func (o *Message) GetRecipient() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value
// and a boolean to check if the value has been set.
func (o *Message) GetRecipientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipient, true
}

// SetRecipient sets field value
func (o *Message) SetRecipient(v string) {
	o.Recipient = v
}

// GetStatus returns the Status field value
func (o *Message) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Message) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Message) SetStatus(v string) {
	o.Status = v
}

// GetNumAttempts returns the NumAttempts field value
func (o *Message) GetNumAttempts() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumAttempts
}

// GetNumAttemptsOk returns a tuple with the NumAttempts field value
// and a boolean to check if the value has been set.
func (o *Message) GetNumAttemptsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumAttempts, true
}

// SetNumAttempts sets field value
func (o *Message) SetNumAttempts(v int32) {
	o.NumAttempts = v
}

// GetClickCount returns the ClickCount field value
func (o *Message) GetClickCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClickCount
}

// GetClickCountOk returns a tuple with the ClickCount field value
// and a boolean to check if the value has been set.
func (o *Message) GetClickCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClickCount, true
}

// SetClickCount sets field value
func (o *Message) SetClickCount(v int32) {
	o.ClickCount = v
}

// GetOpenCount returns the OpenCount field value
func (o *Message) GetOpenCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OpenCount
}

// GetOpenCountOk returns a tuple with the OpenCount field value
// and a boolean to check if the value has been set.
func (o *Message) GetOpenCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OpenCount, true
}

// SetOpenCount sets field value
func (o *Message) SetOpenCount(v int32) {
	o.OpenCount = v
}

// GetReferenceMessageId returns the ReferenceMessageId field value if set, zero value otherwise.
func (o *Message) GetReferenceMessageId() int64 {
	if o == nil || IsNil(o.ReferenceMessageId) {
		var ret int64
		return ret
	}
	return *o.ReferenceMessageId
}

// GetReferenceMessageIdOk returns a tuple with the ReferenceMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetReferenceMessageIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ReferenceMessageId) {
		return nil, false
	}
	return o.ReferenceMessageId, true
}

// HasReferenceMessageId returns a boolean if a field has been set.
func (o *Message) HasReferenceMessageId() bool {
	if o != nil && !IsNil(o.ReferenceMessageId) {
		return true
	}

	return false
}

// SetReferenceMessageId gets a reference to the given int64 and assigns it to the ReferenceMessageId field.
func (o *Message) SetReferenceMessageId(v int64) {
	o.ReferenceMessageId = &v
}

// GetDomainId returns the DomainId field value
func (o *Message) GetDomainId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}

	return o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value
// and a boolean to check if the value has been set.
func (o *Message) GetDomainIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DomainId, true
}

// SetDomainId sets field value
func (o *Message) SetDomainId(v uuid.UUID) {
	o.DomainId = v
}

// GetAccountId returns the AccountId field value
func (o *Message) GetAccountId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *Message) GetAccountIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *Message) SetAccountId(v uuid.UUID) {
	o.AccountId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Message) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Message) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Message) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Message) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Message) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Message) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetSentAt returns the SentAt field value if set, zero value otherwise.
func (o *Message) GetSentAt() time.Time {
	if o == nil || IsNil(o.SentAt) {
		var ret time.Time
		return ret
	}
	return *o.SentAt
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetSentAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SentAt) {
		return nil, false
	}
	return o.SentAt, true
}

// HasSentAt returns a boolean if a field has been set.
func (o *Message) HasSentAt() bool {
	if o != nil && !IsNil(o.SentAt) {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given time.Time and assigns it to the SentAt field.
func (o *Message) SetSentAt(v time.Time) {
	o.SentAt = &v
}

// GetDeliveredAt returns the DeliveredAt field value if set, zero value otherwise.
func (o *Message) GetDeliveredAt() time.Time {
	if o == nil || IsNil(o.DeliveredAt) {
		var ret time.Time
		return ret
	}
	return *o.DeliveredAt
}

// GetDeliveredAtOk returns a tuple with the DeliveredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetDeliveredAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeliveredAt) {
		return nil, false
	}
	return o.DeliveredAt, true
}

// HasDeliveredAt returns a boolean if a field has been set.
func (o *Message) HasDeliveredAt() bool {
	if o != nil && !IsNil(o.DeliveredAt) {
		return true
	}

	return false
}

// SetDeliveredAt gets a reference to the given time.Time and assigns it to the DeliveredAt field.
func (o *Message) SetDeliveredAt(v time.Time) {
	o.DeliveredAt = &v
}

// GetRetainUntil returns the RetainUntil field value
func (o *Message) GetRetainUntil() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.RetainUntil
}

// GetRetainUntilOk returns a tuple with the RetainUntil field value
// and a boolean to check if the value has been set.
func (o *Message) GetRetainUntilOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetainUntil, true
}

// SetRetainUntil sets field value
func (o *Message) SetRetainUntil(v time.Time) {
	o.RetainUntil = v
}

// GetDirection returns the Direction field value
func (o *Message) GetDirection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *Message) GetDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *Message) SetDirection(v string) {
	o.Direction = v
}

// GetIsBounceNotification returns the IsBounceNotification field value
func (o *Message) GetIsBounceNotification() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBounceNotification
}

// GetIsBounceNotificationOk returns a tuple with the IsBounceNotification field value
// and a boolean to check if the value has been set.
func (o *Message) GetIsBounceNotificationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBounceNotification, true
}

// SetIsBounceNotification sets field value
func (o *Message) SetIsBounceNotification(v bool) {
	o.IsBounceNotification = v
}

// GetBounceClassification returns the BounceClassification field value if set, zero value otherwise.
func (o *Message) GetBounceClassification() string {
	if o == nil || IsNil(o.BounceClassification) {
		var ret string
		return ret
	}
	return *o.BounceClassification
}

// GetBounceClassificationOk returns a tuple with the BounceClassification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetBounceClassificationOk() (*string, bool) {
	if o == nil || IsNil(o.BounceClassification) {
		return nil, false
	}
	return o.BounceClassification, true
}

// HasBounceClassification returns a boolean if a field has been set.
func (o *Message) HasBounceClassification() bool {
	if o != nil && !IsNil(o.BounceClassification) {
		return true
	}

	return false
}

// SetBounceClassification gets a reference to the given string and assigns it to the BounceClassification field.
func (o *Message) SetBounceClassification(v string) {
	o.BounceClassification = &v
}

func (o Message) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Message) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	if !IsNil(o.SentAt) {
		toSerialize["sent_at"] = o.SentAt
	}
	if !IsNil(o.DeliveredAt) {
		toSerialize["delivered_at"] = o.DeliveredAt
	}
	toSerialize["retain_until"] = o.RetainUntil
	toSerialize["direction"] = o.Direction
	toSerialize["is_bounce_notification"] = o.IsBounceNotification
	if !IsNil(o.BounceClassification) {
		toSerialize["bounce_classification"] = o.BounceClassification
	}
	toSerialize["delivery_attempts"] = o.DeliveryAttempts
	toSerialize["message_id"] = o.MessageId
	toSerialize["api_id"] = o.ApiId
	toSerialize["ahasend_id"] = o.AhasendId
	toSerialize["subject"] = o.Subject
	toSerialize["tags"] = o.Tags
	toSerialize["sender"] = o.Sender
	toSerialize["recipient"] = o.Recipient
	toSerialize["status"] = o.Status
	toSerialize["num_attempts"] = o.NumAttempts
	toSerialize["click_count"] = o.ClickCount
	toSerialize["open_count"] = o.OpenCount
	if !IsNil(o.ReferenceMessageId) {
		toSerialize["reference_message_id"] = o.ReferenceMessageId
	}
	toSerialize["domain_id"] = o.DomainId
	toSerialize["account_id"] = o.AccountId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Message) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"created_at",
		"updated_at",
		"retain_until",
		"direction",
		"is_bounce_notification",
		"delivery_attempts",
		"message_id",
		"api_id",
		"ahasend_id",
		"subject",
		"tags",
		"sender",
		"recipient",
		"status",
		"num_attempts",
		"click_count",
		"open_count",
		"domain_id",
		"account_id",
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

	varMessage := _Message{}

	err = json.Unmarshal(data, &varMessage)

	if err != nil {
		return err
	}

	*o = Message(varMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "sent_at")
		delete(additionalProperties, "delivered_at")
		delete(additionalProperties, "retain_until")
		delete(additionalProperties, "direction")
		delete(additionalProperties, "is_bounce_notification")
		delete(additionalProperties, "bounce_classification")
		delete(additionalProperties, "delivery_attempts")
		delete(additionalProperties, "message_id")
		delete(additionalProperties, "api_id")
		delete(additionalProperties, "ahasend_id")
		delete(additionalProperties, "subject")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "sender")
		delete(additionalProperties, "recipient")
		delete(additionalProperties, "status")
		delete(additionalProperties, "num_attempts")
		delete(additionalProperties, "click_count")
		delete(additionalProperties, "open_count")
		delete(additionalProperties, "reference_message_id")
		delete(additionalProperties, "domain_id")
		delete(additionalProperties, "account_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMessage struct {
	value *Message
	isSet bool
}

func (v NullableMessage) Get() *Message {
	return v.value
}

func (v *NullableMessage) Set(val *Message) {
	v.value = val
	v.isSet = true
}

func (v NullableMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessage(val *Message) *NullableMessage {
	return &NullableMessage{value: val, isSet: true}
}

func (v NullableMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
