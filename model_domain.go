package ahasend

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"time"
)

// checks if the Domain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Domain{}

// Domain struct for Domain
type Domain struct {
	// Object type identifier
	Object string `json:"object"`
	// Unique identifier for the domain
	Id uuid.UUID `json:"id"`
	// When the domain was created
	CreatedAt time.Time `json:"created_at"`
	// When the domain was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// The domain name
	Domain string `json:"domain"`
	// Account ID this domain belongs to
	AccountId uuid.UUID `json:"account_id"`
	// DNS records required for domain verification
	DnsRecords []DNSRecord `json:"dns_records"`
	// When DNS records were last checked
	LastDnsCheckAt *time.Time `json:"last_dns_check_at,omitempty"`
	// Whether all required DNS records are properly configured
	DnsValid             bool `json:"dns_valid"`
	AdditionalProperties map[string]interface{}
}

type _Domain Domain

// NewDomain instantiates a new Domain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomain(object string, id uuid.UUID, createdAt time.Time, updatedAt time.Time, domain string, accountId uuid.UUID, dnsRecords []DNSRecord, dnsValid bool) *Domain {
	this := Domain{}
	this.Object = object
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Domain = domain
	this.AccountId = accountId
	this.DnsRecords = dnsRecords
	this.DnsValid = dnsValid
	return &this
}

// NewDomainWithDefaults instantiates a new Domain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainWithDefaults() *Domain {
	this := Domain{}
	return &this
}

// GetObject returns the Object field value
func (o *Domain) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Domain) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Domain) SetObject(v string) {
	o.Object = v
}

// GetId returns the Id field value
func (o *Domain) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Domain) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Domain) SetId(v uuid.UUID) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Domain) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Domain) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Domain) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Domain) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Domain) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Domain) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetDomain returns the Domain field value
func (o *Domain) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *Domain) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *Domain) SetDomain(v string) {
	o.Domain = v
}

// GetAccountId returns the AccountId field value
func (o *Domain) GetAccountId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *Domain) GetAccountIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *Domain) SetAccountId(v uuid.UUID) {
	o.AccountId = v
}

// GetDnsRecords returns the DnsRecords field value
func (o *Domain) GetDnsRecords() []DNSRecord {
	if o == nil {
		var ret []DNSRecord
		return ret
	}

	return o.DnsRecords
}

// GetDnsRecordsOk returns a tuple with the DnsRecords field value
// and a boolean to check if the value has been set.
func (o *Domain) GetDnsRecordsOk() ([]DNSRecord, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnsRecords, true
}

// SetDnsRecords sets field value
func (o *Domain) SetDnsRecords(v []DNSRecord) {
	o.DnsRecords = v
}

// GetLastDnsCheckAt returns the LastDnsCheckAt field value if set, zero value otherwise.
func (o *Domain) GetLastDnsCheckAt() time.Time {
	if o == nil || IsNil(o.LastDnsCheckAt) {
		var ret time.Time
		return ret
	}
	return *o.LastDnsCheckAt
}

// GetLastDnsCheckAtOk returns a tuple with the LastDnsCheckAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetLastDnsCheckAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastDnsCheckAt) {
		return nil, false
	}
	return o.LastDnsCheckAt, true
}

// HasLastDnsCheckAt returns a boolean if a field has been set.
func (o *Domain) HasLastDnsCheckAt() bool {
	if o != nil && !IsNil(o.LastDnsCheckAt) {
		return true
	}

	return false
}

// SetLastDnsCheckAt gets a reference to the given time.Time and assigns it to the LastDnsCheckAt field.
func (o *Domain) SetLastDnsCheckAt(v time.Time) {
	o.LastDnsCheckAt = &v
}

// GetDnsValid returns the DnsValid field value
func (o *Domain) GetDnsValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DnsValid
}

// GetDnsValidOk returns a tuple with the DnsValid field value
// and a boolean to check if the value has been set.
func (o *Domain) GetDnsValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DnsValid, true
}

// SetDnsValid sets field value
func (o *Domain) SetDnsValid(v bool) {
	o.DnsValid = v
}

func (o Domain) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Domain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["domain"] = o.Domain
	toSerialize["account_id"] = o.AccountId
	toSerialize["dns_records"] = o.DnsRecords
	if !IsNil(o.LastDnsCheckAt) {
		toSerialize["last_dns_check_at"] = o.LastDnsCheckAt
	}
	toSerialize["dns_valid"] = o.DnsValid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Domain) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"id",
		"created_at",
		"updated_at",
		"domain",
		"account_id",
		"dns_records",
		"dns_valid",
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

	varDomain := _Domain{}

	err = json.Unmarshal(data, &varDomain)

	if err != nil {
		return err
	}

	*o = Domain(varDomain)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "dns_records")
		delete(additionalProperties, "last_dns_check_at")
		delete(additionalProperties, "dns_valid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDomain struct {
	value *Domain
	isSet bool
}

func (v NullableDomain) Get() *Domain {
	return v.value
}

func (v *NullableDomain) Set(val *Domain) {
	v.value = val
	v.isSet = true
}

func (v NullableDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomain(val *Domain) *NullableDomain {
	return &NullableDomain{value: val, isSet: true}
}

func (v NullableDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
