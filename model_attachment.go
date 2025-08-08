package ahasend

import (
	"encoding/json"
	"fmt"
)

// checks if the Attachment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Attachment{}

// Attachment struct for Attachment
type Attachment struct {
	// If true, data must be encoded using base64. Otherwise, data will be interpreted as UTF-8
	Base64 *bool `json:"base64,omitempty"`
	// Either plaintext or base64 encoded attachment data (depending on base64 field)
	Data string `json:"data"`
	// The MIME type of the attachment
	ContentType string `json:"content_type"`
	// The Content-ID of the attachment for inline images
	ContentId *string `json:"content_id,omitempty"`
	// The filename of the attachment
	FileName             string `json:"file_name"`
	AdditionalProperties map[string]interface{}
}

type _Attachment Attachment

// NewAttachment instantiates a new Attachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachment(data string, contentType string, fileName string) *Attachment {
	this := Attachment{}
	var base64 = false
	this.Base64 = &base64
	this.Data = data
	this.ContentType = contentType
	this.FileName = fileName
	return &this
}

// NewAttachmentWithDefaults instantiates a new Attachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentWithDefaults() *Attachment {
	this := Attachment{}
	var base64 = false
	this.Base64 = &base64
	return &this
}

// GetBase64 returns the Base64 field value if set, zero value otherwise.
func (o *Attachment) GetBase64() bool {
	if o == nil || IsNil(o.Base64) {
		var ret bool
		return ret
	}
	return *o.Base64
}

// GetBase64Ok returns a tuple with the Base64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attachment) GetBase64Ok() (*bool, bool) {
	if o == nil || IsNil(o.Base64) {
		return nil, false
	}
	return o.Base64, true
}

// HasBase64 returns a boolean if a field has been set.
func (o *Attachment) HasBase64() bool {
	if o != nil && !IsNil(o.Base64) {
		return true
	}

	return false
}

// SetBase64 gets a reference to the given bool and assigns it to the Base64 field.
func (o *Attachment) SetBase64(v bool) {
	o.Base64 = &v
}

// GetData returns the Data field value
func (o *Attachment) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Attachment) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Attachment) SetData(v string) {
	o.Data = v
}

// GetContentType returns the ContentType field value
func (o *Attachment) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *Attachment) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *Attachment) SetContentType(v string) {
	o.ContentType = v
}

// GetContentId returns the ContentId field value if set, zero value otherwise.
func (o *Attachment) GetContentId() string {
	if o == nil || IsNil(o.ContentId) {
		var ret string
		return ret
	}
	return *o.ContentId
}

// GetContentIdOk returns a tuple with the ContentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attachment) GetContentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContentId) {
		return nil, false
	}
	return o.ContentId, true
}

// HasContentId returns a boolean if a field has been set.
func (o *Attachment) HasContentId() bool {
	if o != nil && !IsNil(o.ContentId) {
		return true
	}

	return false
}

// SetContentId gets a reference to the given string and assigns it to the ContentId field.
func (o *Attachment) SetContentId(v string) {
	o.ContentId = &v
}

// GetFileName returns the FileName field value
func (o *Attachment) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *Attachment) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *Attachment) SetFileName(v string) {
	o.FileName = v
}

func (o Attachment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Attachment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Base64) {
		toSerialize["base64"] = o.Base64
	}
	toSerialize["data"] = o.Data
	toSerialize["content_type"] = o.ContentType
	if !IsNil(o.ContentId) {
		toSerialize["content_id"] = o.ContentId
	}
	toSerialize["file_name"] = o.FileName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Attachment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"content_type",
		"file_name",
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

	varAttachment := _Attachment{}

	err = json.Unmarshal(data, &varAttachment)

	if err != nil {
		return err
	}

	*o = Attachment(varAttachment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "base64")
		delete(additionalProperties, "data")
		delete(additionalProperties, "content_type")
		delete(additionalProperties, "content_id")
		delete(additionalProperties, "file_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAttachment struct {
	value *Attachment
	isSet bool
}

func (v NullableAttachment) Get() *Attachment {
	return v.value
}

func (v *NullableAttachment) Set(val *Attachment) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachment(val *Attachment) *NullableAttachment {
	return &NullableAttachment{value: val, isSet: true}
}

func (v NullableAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
