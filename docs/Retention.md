# Retention

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **int32** | Number of days to retain metadata | [optional] 
**Data** | Pointer to **int32** | Number of days to retain data | [optional] 

## Methods

### NewRetention

`func NewRetention() *Retention`

NewRetention instantiates a new Retention object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetentionWithDefaults

`func NewRetentionWithDefaults() *Retention`

NewRetentionWithDefaults instantiates a new Retention object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *Retention) GetMetadata() int32`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Retention) GetMetadataOk() (*int32, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Retention) SetMetadata(v int32)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Retention) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetData

`func (o *Retention) GetData() int32`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Retention) GetDataOk() (*int32, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Retention) SetData(v int32)`

SetData sets Data field to given value.

### HasData

`func (o *Retention) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


