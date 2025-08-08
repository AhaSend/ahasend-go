# Suppression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Object type identifier | 
**Id** | [**uuid.UUID**](uuid.UUID.md) | Unique identifier for the suppression | 
**CreatedAt** | **time.Time** | When the suppression was created | 
**UpdatedAt** | **time.Time** | When the suppression was last updated | 
**Email** | **string** | Suppressed email address | 
**Domain** | Pointer to **string** | Domain for which the email is suppressed | [optional] 
**Reason** | Pointer to **string** | Reason for suppression | [optional] 
**ExpiresAt** | **time.Time** | When the suppression expires | 

## Methods

### NewSuppression

`func NewSuppression(object string, id uuid.UUID, createdAt time.Time, updatedAt time.Time, email string, expiresAt time.Time, ) *Suppression`

NewSuppression instantiates a new Suppression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuppressionWithDefaults

`func NewSuppressionWithDefaults() *Suppression`

NewSuppressionWithDefaults instantiates a new Suppression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Suppression) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Suppression) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Suppression) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *Suppression) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Suppression) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Suppression) SetId(v uuid.UUID)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Suppression) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Suppression) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Suppression) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Suppression) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Suppression) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Suppression) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetEmail

`func (o *Suppression) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Suppression) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Suppression) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetDomain

`func (o *Suppression) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Suppression) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Suppression) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Suppression) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetReason

`func (o *Suppression) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Suppression) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Suppression) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Suppression) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetExpiresAt

`func (o *Suppression) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Suppression) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Suppression) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


