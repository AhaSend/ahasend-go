# SMTPCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Object type identifier | 
**Id** | [**uuid.UUID**](uuid.UUID.md) | Unique identifier for the SMTP credential | 
**CreatedAt** | **time.Time** | When the credential was created | 
**UpdatedAt** | **time.Time** | When the credential was last updated | 
**Name** | **string** | Credential name | 
**Username** | **string** | SMTP username | 
**Sandbox** | **bool** | Whether this is a sandbox credential | 
**Scope** | **string** | Credential scope | 
**Domains** | Pointer to **[]string** | Domains this credential can send from | [optional] 

## Methods

### NewSMTPCredential

`func NewSMTPCredential(object string, id uuid.UUID, createdAt time.Time, updatedAt time.Time, name string, username string, sandbox bool, scope string, ) *SMTPCredential`

NewSMTPCredential instantiates a new SMTPCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMTPCredentialWithDefaults

`func NewSMTPCredentialWithDefaults() *SMTPCredential`

NewSMTPCredentialWithDefaults instantiates a new SMTPCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *SMTPCredential) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SMTPCredential) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SMTPCredential) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *SMTPCredential) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SMTPCredential) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SMTPCredential) SetId(v uuid.UUID)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *SMTPCredential) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SMTPCredential) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SMTPCredential) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SMTPCredential) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SMTPCredential) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SMTPCredential) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetName

`func (o *SMTPCredential) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SMTPCredential) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SMTPCredential) SetName(v string)`

SetName sets Name field to given value.


### GetUsername

`func (o *SMTPCredential) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SMTPCredential) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SMTPCredential) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetSandbox

`func (o *SMTPCredential) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *SMTPCredential) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *SMTPCredential) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.


### GetScope

`func (o *SMTPCredential) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SMTPCredential) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SMTPCredential) SetScope(v string)`

SetScope sets Scope field to given value.


### GetDomains

`func (o *SMTPCredential) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *SMTPCredential) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *SMTPCredential) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *SMTPCredential) HasDomains() bool`

HasDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


