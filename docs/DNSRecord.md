# DNSRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | DNS record type (e.g., CNAME, TXT, MX) | 
**Host** | **string** | DNS record host/name | 
**Content** | **string** | DNS record content/value | 
**Required** | **bool** | Whether this DNS record is required for domain verification | 
**Propagated** | **bool** | Whether this DNS record has been propagated and verified | 

## Methods

### NewDNSRecord

`func NewDNSRecord(type_ string, host string, content string, required bool, propagated bool, ) *DNSRecord`

NewDNSRecord instantiates a new DNSRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSRecordWithDefaults

`func NewDNSRecordWithDefaults() *DNSRecord`

NewDNSRecordWithDefaults instantiates a new DNSRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DNSRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSRecord) SetType(v string)`

SetType sets Type field to given value.


### GetHost

`func (o *DNSRecord) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DNSRecord) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DNSRecord) SetHost(v string)`

SetHost sets Host field to given value.


### GetContent

`func (o *DNSRecord) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DNSRecord) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DNSRecord) SetContent(v string)`

SetContent sets Content field to given value.


### GetRequired

`func (o *DNSRecord) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *DNSRecord) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *DNSRecord) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetPropagated

`func (o *DNSRecord) GetPropagated() bool`

GetPropagated returns the Propagated field if non-nil, zero value otherwise.

### GetPropagatedOk

`func (o *DNSRecord) GetPropagatedOk() (*bool, bool)`

GetPropagatedOk returns a tuple with the Propagated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagated

`func (o *DNSRecord) SetPropagated(v bool)`

SetPropagated sets Propagated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


