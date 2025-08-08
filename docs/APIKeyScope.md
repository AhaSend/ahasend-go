# APIKeyScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**uuid.UUID**](uuid.UUID.md) | Unique identifier for the scope | 
**CreatedAt** | **time.Time** | When the scope was created | 
**UpdatedAt** | **time.Time** | When the scope was last updated | 
**ApiKeyId** | [**uuid.UUID**](uuid.UUID.md) | ID of the API key this scope belongs to | 
**Scope** | **string** | The scope string | 
**DomainId** | Pointer to [**uuid.UUID**](uuid.UUID.md) | Domain ID for domain-specific scopes | [optional] 

## Methods

### NewAPIKeyScope

`func NewAPIKeyScope(id uuid.UUID, createdAt time.Time, updatedAt time.Time, apiKeyId uuid.UUID, scope string, ) *APIKeyScope`

NewAPIKeyScope instantiates a new APIKeyScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIKeyScopeWithDefaults

`func NewAPIKeyScopeWithDefaults() *APIKeyScope`

NewAPIKeyScopeWithDefaults instantiates a new APIKeyScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *APIKeyScope) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIKeyScope) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIKeyScope) SetId(v uuid.UUID)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *APIKeyScope) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *APIKeyScope) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *APIKeyScope) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *APIKeyScope) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *APIKeyScope) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *APIKeyScope) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetApiKeyId

`func (o *APIKeyScope) GetApiKeyId() uuid.UUID`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *APIKeyScope) GetApiKeyIdOk() (*uuid.UUID, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *APIKeyScope) SetApiKeyId(v uuid.UUID)`

SetApiKeyId sets ApiKeyId field to given value.


### GetScope

`func (o *APIKeyScope) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *APIKeyScope) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *APIKeyScope) SetScope(v string)`

SetScope sets Scope field to given value.


### GetDomainId

`func (o *APIKeyScope) GetDomainId() uuid.UUID`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *APIKeyScope) GetDomainIdOk() (*uuid.UUID, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *APIKeyScope) SetDomainId(v uuid.UUID)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *APIKeyScope) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


