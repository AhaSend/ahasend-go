# Bounce

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classification** | **string** | Bounce classification | 
**Count** | **int32** | Number of bounces | 

## Methods

### NewBounce

`func NewBounce(classification string, count int32, ) *Bounce`

NewBounce instantiates a new Bounce object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBounceWithDefaults

`func NewBounceWithDefaults() *Bounce`

NewBounceWithDefaults instantiates a new Bounce object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassification

`func (o *Bounce) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *Bounce) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *Bounce) SetClassification(v string)`

SetClassification sets Classification field to given value.


### GetCount

`func (o *Bounce) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Bounce) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Bounce) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


