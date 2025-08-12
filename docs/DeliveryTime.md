# DeliveryTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecipientDomain** | Pointer to **string** | The recipient domain | [optional] 
**DeliveryTime** | Pointer to **float64** | The average time from reception to delivery in seconds | [optional] 

## Methods

### NewDeliveryTime

`func NewDeliveryTime() *DeliveryTime`

NewDeliveryTime instantiates a new DeliveryTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryTimeWithDefaults

`func NewDeliveryTimeWithDefaults() *DeliveryTime`

NewDeliveryTimeWithDefaults instantiates a new DeliveryTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipientDomain

`func (o *DeliveryTime) GetRecipientDomain() string`

GetRecipientDomain returns the RecipientDomain field if non-nil, zero value otherwise.

### GetRecipientDomainOk

`func (o *DeliveryTime) GetRecipientDomainOk() (*string, bool)`

GetRecipientDomainOk returns a tuple with the RecipientDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientDomain

`func (o *DeliveryTime) SetRecipientDomain(v string)`

SetRecipientDomain sets RecipientDomain field to given value.

### HasRecipientDomain

`func (o *DeliveryTime) HasRecipientDomain() bool`

HasRecipientDomain returns a boolean if a field has been set.

### GetDeliveryTime

`func (o *DeliveryTime) GetDeliveryTime() float64`

GetDeliveryTime returns the DeliveryTime field if non-nil, zero value otherwise.

### GetDeliveryTimeOk

`func (o *DeliveryTime) GetDeliveryTimeOk() (*float64, bool)`

GetDeliveryTimeOk returns a tuple with the DeliveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTime

`func (o *DeliveryTime) SetDeliveryTime(v float64)`

SetDeliveryTime sets DeliveryTime field to given value.

### HasDeliveryTime

`func (o *DeliveryTime) HasDeliveryTime() bool`

HasDeliveryTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


