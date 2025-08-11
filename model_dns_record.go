package ahasend

import (
	"encoding/json"
	"fmt"
)

// checks if the DNSRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DNSRecord{}

// DNSRecord struct for DNSRecord
type DNSRecord struct {
	// DNS record type (e.g., CNAME, TXT, MX)
	Type string `json:"type"`
	// DNS record host/name
	Host string `json:"host"`
	// DNS record content/value
	Content string `json:"content"`
	// Whether this DNS record is required for domain verification
	Required bool `json:"required"`
	// Whether this DNS record has been propagated and verified
	Propagated           bool `json:"propagated"`
	AdditionalProperties map[string]interface{}
}

type _DNSRecord DNSRecord

// NewDNSRecord instantiates a new DNSRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDNSRecord(recordType string, host string, content string, required bool, propagated bool) *DNSRecord {
	this := DNSRecord{}
	this.Type = recordType
	this.Host = host
	this.Content = content
	this.Required = required
	this.Propagated = propagated
	return &this
}

// NewDNSRecordWithDefaults instantiates a new DNSRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDNSRecordWithDefaults() *DNSRecord {
	this := DNSRecord{}
	return &this
}

// GetType returns the Type field value
func (o *DNSRecord) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DNSRecord) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DNSRecord) SetType(v string) {
	o.Type = v
}

// GetHost returns the Host field value
func (o *DNSRecord) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *DNSRecord) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *DNSRecord) SetHost(v string) {
	o.Host = v
}

// GetContent returns the Content field value
func (o *DNSRecord) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *DNSRecord) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *DNSRecord) SetContent(v string) {
	o.Content = v
}

// GetRequired returns the Required field value
func (o *DNSRecord) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *DNSRecord) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *DNSRecord) SetRequired(v bool) {
	o.Required = v
}

// GetPropagated returns the Propagated field value
func (o *DNSRecord) GetPropagated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Propagated
}

// GetPropagatedOk returns a tuple with the Propagated field value
// and a boolean to check if the value has been set.
func (o *DNSRecord) GetPropagatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Propagated, true
}

// SetPropagated sets field value
func (o *DNSRecord) SetPropagated(v bool) {
	o.Propagated = v
}

func (o DNSRecord) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DNSRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["host"] = o.Host
	toSerialize["content"] = o.Content
	toSerialize["required"] = o.Required
	toSerialize["propagated"] = o.Propagated

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DNSRecord) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"host",
		"content",
		"required",
		"propagated",
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

	varDNSRecord := _DNSRecord{}

	err = json.Unmarshal(data, &varDNSRecord)

	if err != nil {
		return err
	}

	*o = DNSRecord(varDNSRecord)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "host")
		delete(additionalProperties, "content")
		delete(additionalProperties, "required")
		delete(additionalProperties, "propagated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

func (o *DNSRecord) String() string {
	statusInfo := ""
	if o.Required {
		if o.Propagated {
			statusInfo = " ✓ (required, propagated)"
		} else {
			statusInfo = " ✗ (not propagated)"
		}
	} else {
		if o.Propagated {
			statusInfo = " ✓ (optional, propagated)"
		} else {
			statusInfo = " ✗ (optional, not propagated)"
		}
	}

	return fmt.Sprintf("\n    %s %s%s\n        %s", o.Type, o.Host, statusInfo, o.Content)
}

type NullableDNSRecord struct {
	value *DNSRecord
	isSet bool
}

func (v NullableDNSRecord) Get() *DNSRecord {
	return v.value
}

func (v *NullableDNSRecord) Set(val *DNSRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableDNSRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableDNSRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDNSRecord(val *DNSRecord) *NullableDNSRecord {
	return &NullableDNSRecord{value: val, isSet: true}
}

func (v NullableDNSRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDNSRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
