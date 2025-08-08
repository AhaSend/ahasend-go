# CreateSMTPCredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Credential name | 
**Username** | **string** | SMTP username | 
**Password** | **string** | SMTP password | 
**Sandbox** | Pointer to **bool** | Whether this is a sandbox credential | [optional] [default to false]
**Scope** | **string** | Credential scope - \&quot;global\&quot; or \&quot;scoped\&quot; | 
**Domains** | Pointer to **[]string** | Required if scope is \&quot;scoped\&quot; | [optional] 

## Methods

### NewCreateSMTPCredentialRequest

`func NewCreateSMTPCredentialRequest(name string, username string, password string, scope string, ) *CreateSMTPCredentialRequest`

NewCreateSMTPCredentialRequest instantiates a new CreateSMTPCredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSMTPCredentialRequestWithDefaults

`func NewCreateSMTPCredentialRequestWithDefaults() *CreateSMTPCredentialRequest`

NewCreateSMTPCredentialRequestWithDefaults instantiates a new CreateSMTPCredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSMTPCredentialRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSMTPCredentialRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSMTPCredentialRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUsername

`func (o *CreateSMTPCredentialRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateSMTPCredentialRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateSMTPCredentialRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *CreateSMTPCredentialRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateSMTPCredentialRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateSMTPCredentialRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetSandbox

`func (o *CreateSMTPCredentialRequest) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *CreateSMTPCredentialRequest) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *CreateSMTPCredentialRequest) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.

### HasSandbox

`func (o *CreateSMTPCredentialRequest) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.

### GetScope

`func (o *CreateSMTPCredentialRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateSMTPCredentialRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateSMTPCredentialRequest) SetScope(v string)`

SetScope sets Scope field to given value.


### GetDomains

`func (o *CreateSMTPCredentialRequest) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *CreateSMTPCredentialRequest) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *CreateSMTPCredentialRequest) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *CreateSMTPCredentialRequest) HasDomains() bool`

HasDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


