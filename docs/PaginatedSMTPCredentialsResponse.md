# PaginatedSMTPCredentialsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Object type identifier | 
**Data** | [**[]SMTPCredential**](SMTPCredential.md) | Array of SMTP credentials | 
**Pagination** | [**PaginationInfo**](PaginationInfo.md) |  | 

## Methods

### NewPaginatedSMTPCredentialsResponse

`func NewPaginatedSMTPCredentialsResponse(object string, data []SMTPCredential, pagination PaginationInfo, ) *PaginatedSMTPCredentialsResponse`

NewPaginatedSMTPCredentialsResponse instantiates a new PaginatedSMTPCredentialsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedSMTPCredentialsResponseWithDefaults

`func NewPaginatedSMTPCredentialsResponseWithDefaults() *PaginatedSMTPCredentialsResponse`

NewPaginatedSMTPCredentialsResponseWithDefaults instantiates a new PaginatedSMTPCredentialsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *PaginatedSMTPCredentialsResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaginatedSMTPCredentialsResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaginatedSMTPCredentialsResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetData

`func (o *PaginatedSMTPCredentialsResponse) GetData() []SMTPCredential`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaginatedSMTPCredentialsResponse) GetDataOk() (*[]SMTPCredential, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaginatedSMTPCredentialsResponse) SetData(v []SMTPCredential)`

SetData sets Data field to given value.


### GetPagination

`func (o *PaginatedSMTPCredentialsResponse) GetPagination() PaginationInfo`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedSMTPCredentialsResponse) GetPaginationOk() (*PaginationInfo, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedSMTPCredentialsResponse) SetPagination(v PaginationInfo)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


