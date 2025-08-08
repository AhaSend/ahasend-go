# AddMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email address of the user to add | 
**Name** | Pointer to **string** | Display name for the user | [optional] 
**Role** | **string** | Role to assign to the user | 

## Methods

### NewAddMemberRequest

`func NewAddMemberRequest(email string, role string, ) *AddMemberRequest`

NewAddMemberRequest instantiates a new AddMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddMemberRequestWithDefaults

`func NewAddMemberRequestWithDefaults() *AddMemberRequest`

NewAddMemberRequestWithDefaults instantiates a new AddMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AddMemberRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddMemberRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddMemberRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *AddMemberRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddMemberRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddMemberRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddMemberRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRole

`func (o *AddMemberRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AddMemberRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AddMemberRequest) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


