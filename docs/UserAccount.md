# UserAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**uuid.UUID**](uuid.UUID.md) | Unique identifier for the user account relationship | 
**CreatedAt** | **time.Time** | When the relationship was created | 
**UpdatedAt** | **time.Time** | When the relationship was last updated | 
**UserId** | [**uuid.UUID**](uuid.UUID.md) | User ID | 
**AccountId** | [**uuid.UUID**](uuid.UUID.md) | Account ID | 
**Role** | **string** | User role in the account | 

## Methods

### NewUserAccount

`func NewUserAccount(id uuid.UUID, createdAt time.Time, updatedAt time.Time, userId uuid.UUID, accountId uuid.UUID, role string, ) *UserAccount`

NewUserAccount instantiates a new UserAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccountWithDefaults

`func NewUserAccountWithDefaults() *UserAccount`

NewUserAccountWithDefaults instantiates a new UserAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserAccount) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserAccount) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserAccount) SetId(v uuid.UUID)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *UserAccount) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserAccount) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserAccount) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UserAccount) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserAccount) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserAccount) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *UserAccount) GetUserId() uuid.UUID`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserAccount) GetUserIdOk() (*uuid.UUID, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserAccount) SetUserId(v uuid.UUID)`

SetUserId sets UserId field to given value.


### GetAccountId

`func (o *UserAccount) GetAccountId() uuid.UUID`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserAccount) GetAccountIdOk() (*uuid.UUID, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserAccount) SetAccountId(v uuid.UUID)`

SetAccountId sets AccountId field to given value.


### GetRole

`func (o *UserAccount) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserAccount) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserAccount) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


