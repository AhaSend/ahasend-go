# Recipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Recipient email address | 
**Name** | Pointer to **string** | Display name for the recipient | [optional] 
**SubstitutionData** | Pointer to **map[string]interface{}** | Substitution data for the recipient. Used with jinja2 templating language for dynamic content | [optional] 

## Methods

### NewRecipient

`func NewRecipient(email string, ) *Recipient`

NewRecipient instantiates a new Recipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipientWithDefaults

`func NewRecipientWithDefaults() *Recipient`

NewRecipientWithDefaults instantiates a new Recipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *Recipient) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Recipient) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Recipient) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *Recipient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Recipient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Recipient) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Recipient) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubstitutionData

`func (o *Recipient) GetSubstitutionData() map[string]interface{}`

GetSubstitutionData returns the SubstitutionData field if non-nil, zero value otherwise.

### GetSubstitutionDataOk

`func (o *Recipient) GetSubstitutionDataOk() (*map[string]interface{}, bool)`

GetSubstitutionDataOk returns a tuple with the SubstitutionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstitutionData

`func (o *Recipient) SetSubstitutionData(v map[string]interface{})`

SetSubstitutionData sets SubstitutionData field to given value.

### HasSubstitutionData

`func (o *Recipient) HasSubstitutionData() bool`

HasSubstitutionData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


