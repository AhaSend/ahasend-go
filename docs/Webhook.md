# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Object type identifier | 
**Id** | [**uuid.UUID**](uuid.UUID.md) | Unique identifier for the webhook | 
**CreatedAt** | **time.Time** | When the webhook was created | 
**UpdatedAt** | **time.Time** | When the webhook was last updated | 
**Name** | **string** | Webhook name | 
**Url** | **string** | Webhook URL | 
**Enabled** | **bool** | Whether the webhook is enabled | 
**OnReception** | Pointer to **bool** | Trigger on message reception | [optional] 
**OnDelivered** | Pointer to **bool** | Trigger on message delivery | [optional] 
**OnTransientError** | Pointer to **bool** | Trigger on transient errors | [optional] 
**OnFailed** | Pointer to **bool** | Trigger on permanent failures | [optional] 
**OnBounced** | Pointer to **bool** | Trigger on bounces | [optional] 
**OnSuppressed** | Pointer to **bool** | Trigger on suppressions | [optional] 
**OnOpened** | Pointer to **bool** | Trigger on opens | [optional] 
**OnClicked** | Pointer to **bool** | Trigger on clicks | [optional] 
**OnSuppressionCreated** | Pointer to **bool** | Trigger on suppression creation | [optional] 
**OnDnsError** | Pointer to **bool** | Trigger on DNS errors | [optional] 
**Scope** | Pointer to **string** | Webhook scope | [optional] 
**Domains** | Pointer to **[]string** | Domains this webhook applies to | [optional] 

## Methods

### NewWebhook

`func NewWebhook(object string, id uuid.UUID, createdAt time.Time, updatedAt time.Time, name string, url string, enabled bool, ) *Webhook`

NewWebhook instantiates a new Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithDefaults

`func NewWebhookWithDefaults() *Webhook`

NewWebhookWithDefaults instantiates a new Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Webhook) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Webhook) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Webhook) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *Webhook) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Webhook) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Webhook) SetId(v uuid.UUID)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Webhook) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Webhook) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Webhook) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Webhook) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Webhook) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Webhook) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetName

`func (o *Webhook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Webhook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Webhook) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *Webhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Webhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Webhook) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEnabled

`func (o *Webhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Webhook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Webhook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetOnReception

`func (o *Webhook) GetOnReception() bool`

GetOnReception returns the OnReception field if non-nil, zero value otherwise.

### GetOnReceptionOk

`func (o *Webhook) GetOnReceptionOk() (*bool, bool)`

GetOnReceptionOk returns a tuple with the OnReception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnReception

`func (o *Webhook) SetOnReception(v bool)`

SetOnReception sets OnReception field to given value.

### HasOnReception

`func (o *Webhook) HasOnReception() bool`

HasOnReception returns a boolean if a field has been set.

### GetOnDelivered

`func (o *Webhook) GetOnDelivered() bool`

GetOnDelivered returns the OnDelivered field if non-nil, zero value otherwise.

### GetOnDeliveredOk

`func (o *Webhook) GetOnDeliveredOk() (*bool, bool)`

GetOnDeliveredOk returns a tuple with the OnDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDelivered

`func (o *Webhook) SetOnDelivered(v bool)`

SetOnDelivered sets OnDelivered field to given value.

### HasOnDelivered

`func (o *Webhook) HasOnDelivered() bool`

HasOnDelivered returns a boolean if a field has been set.

### GetOnTransientError

`func (o *Webhook) GetOnTransientError() bool`

GetOnTransientError returns the OnTransientError field if non-nil, zero value otherwise.

### GetOnTransientErrorOk

`func (o *Webhook) GetOnTransientErrorOk() (*bool, bool)`

GetOnTransientErrorOk returns a tuple with the OnTransientError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnTransientError

`func (o *Webhook) SetOnTransientError(v bool)`

SetOnTransientError sets OnTransientError field to given value.

### HasOnTransientError

`func (o *Webhook) HasOnTransientError() bool`

HasOnTransientError returns a boolean if a field has been set.

### GetOnFailed

`func (o *Webhook) GetOnFailed() bool`

GetOnFailed returns the OnFailed field if non-nil, zero value otherwise.

### GetOnFailedOk

`func (o *Webhook) GetOnFailedOk() (*bool, bool)`

GetOnFailedOk returns a tuple with the OnFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailed

`func (o *Webhook) SetOnFailed(v bool)`

SetOnFailed sets OnFailed field to given value.

### HasOnFailed

`func (o *Webhook) HasOnFailed() bool`

HasOnFailed returns a boolean if a field has been set.

### GetOnBounced

`func (o *Webhook) GetOnBounced() bool`

GetOnBounced returns the OnBounced field if non-nil, zero value otherwise.

### GetOnBouncedOk

`func (o *Webhook) GetOnBouncedOk() (*bool, bool)`

GetOnBouncedOk returns a tuple with the OnBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnBounced

`func (o *Webhook) SetOnBounced(v bool)`

SetOnBounced sets OnBounced field to given value.

### HasOnBounced

`func (o *Webhook) HasOnBounced() bool`

HasOnBounced returns a boolean if a field has been set.

### GetOnSuppressed

`func (o *Webhook) GetOnSuppressed() bool`

GetOnSuppressed returns the OnSuppressed field if non-nil, zero value otherwise.

### GetOnSuppressedOk

`func (o *Webhook) GetOnSuppressedOk() (*bool, bool)`

GetOnSuppressedOk returns a tuple with the OnSuppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSuppressed

`func (o *Webhook) SetOnSuppressed(v bool)`

SetOnSuppressed sets OnSuppressed field to given value.

### HasOnSuppressed

`func (o *Webhook) HasOnSuppressed() bool`

HasOnSuppressed returns a boolean if a field has been set.

### GetOnOpened

`func (o *Webhook) GetOnOpened() bool`

GetOnOpened returns the OnOpened field if non-nil, zero value otherwise.

### GetOnOpenedOk

`func (o *Webhook) GetOnOpenedOk() (*bool, bool)`

GetOnOpenedOk returns a tuple with the OnOpened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnOpened

`func (o *Webhook) SetOnOpened(v bool)`

SetOnOpened sets OnOpened field to given value.

### HasOnOpened

`func (o *Webhook) HasOnOpened() bool`

HasOnOpened returns a boolean if a field has been set.

### GetOnClicked

`func (o *Webhook) GetOnClicked() bool`

GetOnClicked returns the OnClicked field if non-nil, zero value otherwise.

### GetOnClickedOk

`func (o *Webhook) GetOnClickedOk() (*bool, bool)`

GetOnClickedOk returns a tuple with the OnClicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnClicked

`func (o *Webhook) SetOnClicked(v bool)`

SetOnClicked sets OnClicked field to given value.

### HasOnClicked

`func (o *Webhook) HasOnClicked() bool`

HasOnClicked returns a boolean if a field has been set.

### GetOnSuppressionCreated

`func (o *Webhook) GetOnSuppressionCreated() bool`

GetOnSuppressionCreated returns the OnSuppressionCreated field if non-nil, zero value otherwise.

### GetOnSuppressionCreatedOk

`func (o *Webhook) GetOnSuppressionCreatedOk() (*bool, bool)`

GetOnSuppressionCreatedOk returns a tuple with the OnSuppressionCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSuppressionCreated

`func (o *Webhook) SetOnSuppressionCreated(v bool)`

SetOnSuppressionCreated sets OnSuppressionCreated field to given value.

### HasOnSuppressionCreated

`func (o *Webhook) HasOnSuppressionCreated() bool`

HasOnSuppressionCreated returns a boolean if a field has been set.

### GetOnDnsError

`func (o *Webhook) GetOnDnsError() bool`

GetOnDnsError returns the OnDnsError field if non-nil, zero value otherwise.

### GetOnDnsErrorOk

`func (o *Webhook) GetOnDnsErrorOk() (*bool, bool)`

GetOnDnsErrorOk returns a tuple with the OnDnsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDnsError

`func (o *Webhook) SetOnDnsError(v bool)`

SetOnDnsError sets OnDnsError field to given value.

### HasOnDnsError

`func (o *Webhook) HasOnDnsError() bool`

HasOnDnsError returns a boolean if a field has been set.

### GetScope

`func (o *Webhook) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Webhook) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Webhook) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Webhook) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetDomains

`func (o *Webhook) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *Webhook) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *Webhook) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *Webhook) HasDomains() bool`

HasDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


