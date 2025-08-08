# CreateAPIKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Human-readable label for the API key | 
**Scopes** | **[]string** | Array of scope strings to grant to this API key | 

## Methods

### NewCreateAPIKeyRequest

`func NewCreateAPIKeyRequest(label string, scopes []string, ) *CreateAPIKeyRequest`

NewCreateAPIKeyRequest instantiates a new CreateAPIKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAPIKeyRequestWithDefaults

`func NewCreateAPIKeyRequestWithDefaults() *CreateAPIKeyRequest`

NewCreateAPIKeyRequestWithDefaults instantiates a new CreateAPIKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateAPIKeyRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateAPIKeyRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateAPIKeyRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetScopes

`func (o *CreateAPIKeyRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CreateAPIKeyRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CreateAPIKeyRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


