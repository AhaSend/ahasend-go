# ModelAPIKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Object type identifier | 
**Id** | [**uuid.UUID**](uuid.UUID.md) | Unique identifier for the API key | 
**CreatedAt** | **time.Time** | When the API key was created | 
**UpdatedAt** | **time.Time** | When the API key was last updated | 
**LastUsedAt** | Pointer to **time.Time** | When the API key was last used (updates every 5-10 minutes) | [optional] 
**AccountId** | [**uuid.UUID**](uuid.UUID.md) | Account ID this API key belongs to | 
**Label** | **string** | Human-readable label for the API key | 
**PublicKey** | **string** | Public portion of the API key | 
**SecretKey** | Pointer to **string** | Secret key (only returned on creation) | [optional] 
**Scopes** | [**[]APIKeyScope**](APIKeyScope.md) | Scopes granted to this API key | 

## Methods

### NewModelAPIKey

`func NewModelAPIKey(object string, id uuid.UUID, createdAt time.Time, updatedAt time.Time, accountId uuid.UUID, label string, publicKey string, scopes []APIKeyScope, ) *ModelAPIKey`

NewModelAPIKey instantiates a new ModelAPIKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIKeyWithDefaults

`func NewModelAPIKeyWithDefaults() *ModelAPIKey`

NewModelAPIKeyWithDefaults instantiates a new ModelAPIKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ModelAPIKey) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ModelAPIKey) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ModelAPIKey) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *ModelAPIKey) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelAPIKey) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelAPIKey) SetId(v uuid.UUID)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ModelAPIKey) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelAPIKey) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelAPIKey) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ModelAPIKey) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelAPIKey) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelAPIKey) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLastUsedAt

`func (o *ModelAPIKey) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *ModelAPIKey) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *ModelAPIKey) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *ModelAPIKey) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetAccountId

`func (o *ModelAPIKey) GetAccountId() uuid.UUID`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ModelAPIKey) GetAccountIdOk() (*uuid.UUID, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ModelAPIKey) SetAccountId(v uuid.UUID)`

SetAccountId sets AccountId field to given value.


### GetLabel

`func (o *ModelAPIKey) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ModelAPIKey) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ModelAPIKey) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetPublicKey

`func (o *ModelAPIKey) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *ModelAPIKey) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *ModelAPIKey) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetSecretKey

`func (o *ModelAPIKey) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *ModelAPIKey) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *ModelAPIKey) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *ModelAPIKey) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetScopes

`func (o *ModelAPIKey) GetScopes() []APIKeyScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ModelAPIKey) GetScopesOk() (*[]APIKeyScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ModelAPIKey) SetScopes(v []APIKeyScope)`

SetScopes sets Scopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


