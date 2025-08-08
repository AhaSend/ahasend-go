# UpdateAPIKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Human-readable label for the API key | [optional] 
**Scopes** | Pointer to **[]string** | Array of scope strings to grant to this API key | [optional] 

## Methods

### NewUpdateAPIKeyRequest

`func NewUpdateAPIKeyRequest() *UpdateAPIKeyRequest`

NewUpdateAPIKeyRequest instantiates a new UpdateAPIKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAPIKeyRequestWithDefaults

`func NewUpdateAPIKeyRequestWithDefaults() *UpdateAPIKeyRequest`

NewUpdateAPIKeyRequestWithDefaults instantiates a new UpdateAPIKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *UpdateAPIKeyRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateAPIKeyRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateAPIKeyRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateAPIKeyRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetScopes

`func (o *UpdateAPIKeyRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *UpdateAPIKeyRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *UpdateAPIKeyRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *UpdateAPIKeyRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


