# Attachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base64** | Pointer to **bool** | If true, data must be encoded using base64. Otherwise, data will be interpreted as UTF-8 | [optional] [default to false]
**Data** | **string** | Either plaintext or base64 encoded attachment data (depending on base64 field) | 
**ContentType** | **string** | The MIME type of the attachment | 
**ContentId** | Pointer to **string** | The Content-ID of the attachment for inline images | [optional] 
**FileName** | **string** | The filename of the attachment | 

## Methods

### NewAttachment

`func NewAttachment(data string, contentType string, fileName string, ) *Attachment`

NewAttachment instantiates a new Attachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentWithDefaults

`func NewAttachmentWithDefaults() *Attachment`

NewAttachmentWithDefaults instantiates a new Attachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase64

`func (o *Attachment) GetBase64() bool`

GetBase64 returns the Base64 field if non-nil, zero value otherwise.

### GetBase64Ok

`func (o *Attachment) GetBase64Ok() (*bool, bool)`

GetBase64Ok returns a tuple with the Base64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase64

`func (o *Attachment) SetBase64(v bool)`

SetBase64 sets Base64 field to given value.

### HasBase64

`func (o *Attachment) HasBase64() bool`

HasBase64 returns a boolean if a field has been set.

### GetData

`func (o *Attachment) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Attachment) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Attachment) SetData(v string)`

SetData sets Data field to given value.


### GetContentType

`func (o *Attachment) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *Attachment) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *Attachment) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetContentId

`func (o *Attachment) GetContentId() string`

GetContentId returns the ContentId field if non-nil, zero value otherwise.

### GetContentIdOk

`func (o *Attachment) GetContentIdOk() (*string, bool)`

GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentId

`func (o *Attachment) SetContentId(v string)`

SetContentId sets ContentId field to given value.

### HasContentId

`func (o *Attachment) HasContentId() bool`

HasContentId returns a boolean if a field has been set.

### GetFileName

`func (o *Attachment) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *Attachment) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *Attachment) SetFileName(v string)`

SetFileName sets FileName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


