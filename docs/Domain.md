# Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Object type identifier | 
**Id** | [**uuid.UUID**](uuid.UUID.md) | Unique identifier for the domain | 
**CreatedAt** | **time.Time** | When the domain was created | 
**UpdatedAt** | **time.Time** | When the domain was last updated | 
**Domain** | **string** | The domain name | 
**AccountId** | [**uuid.UUID**](uuid.UUID.md) | Account ID this domain belongs to | 
**DnsRecords** | [**[]DNSRecord**](DNSRecord.md) | DNS records required for domain verification | 
**LastDnsCheckAt** | Pointer to **time.Time** | When DNS records were last checked | [optional] 
**DnsValid** | **bool** | Whether all required DNS records are properly configured | 

## Methods

### NewDomain

`func NewDomain(object string, id uuid.UUID, createdAt time.Time, updatedAt time.Time, domain string, accountId uuid.UUID, dnsRecords []DNSRecord, dnsValid bool, ) *Domain`

NewDomain instantiates a new Domain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainWithDefaults

`func NewDomainWithDefaults() *Domain`

NewDomainWithDefaults instantiates a new Domain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Domain) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Domain) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Domain) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *Domain) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Domain) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Domain) SetId(v uuid.UUID)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Domain) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Domain) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Domain) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Domain) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Domain) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Domain) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDomain

`func (o *Domain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Domain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Domain) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetAccountId

`func (o *Domain) GetAccountId() uuid.UUID`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Domain) GetAccountIdOk() (*uuid.UUID, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Domain) SetAccountId(v uuid.UUID)`

SetAccountId sets AccountId field to given value.


### GetDnsRecords

`func (o *Domain) GetDnsRecords() []DNSRecord`

GetDnsRecords returns the DnsRecords field if non-nil, zero value otherwise.

### GetDnsRecordsOk

`func (o *Domain) GetDnsRecordsOk() (*[]DNSRecord, bool)`

GetDnsRecordsOk returns a tuple with the DnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecords

`func (o *Domain) SetDnsRecords(v []DNSRecord)`

SetDnsRecords sets DnsRecords field to given value.


### GetLastDnsCheckAt

`func (o *Domain) GetLastDnsCheckAt() time.Time`

GetLastDnsCheckAt returns the LastDnsCheckAt field if non-nil, zero value otherwise.

### GetLastDnsCheckAtOk

`func (o *Domain) GetLastDnsCheckAtOk() (*time.Time, bool)`

GetLastDnsCheckAtOk returns a tuple with the LastDnsCheckAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDnsCheckAt

`func (o *Domain) SetLastDnsCheckAt(v time.Time)`

SetLastDnsCheckAt sets LastDnsCheckAt field to given value.

### HasLastDnsCheckAt

`func (o *Domain) HasLastDnsCheckAt() bool`

HasLastDnsCheckAt returns a boolean if a field has been set.

### GetDnsValid

`func (o *Domain) GetDnsValid() bool`

GetDnsValid returns the DnsValid field if non-nil, zero value otherwise.

### GetDnsValidOk

`func (o *Domain) GetDnsValidOk() (*bool, bool)`

GetDnsValidOk returns a tuple with the DnsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsValid

`func (o *Domain) SetDnsValid(v bool)`

SetDnsValid sets DnsValid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


