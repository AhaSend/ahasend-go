# UpdateWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Webhook name | [optional] 
**Url** | Pointer to **string** | Webhook URL | [optional] 
**Enabled** | Pointer to **bool** | Whether the webhook is enabled | [optional] 
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

### NewUpdateWebhookRequest

`func NewUpdateWebhookRequest() *UpdateWebhookRequest`

NewUpdateWebhookRequest instantiates a new UpdateWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWebhookRequestWithDefaults

`func NewUpdateWebhookRequestWithDefaults() *UpdateWebhookRequest`

NewUpdateWebhookRequestWithDefaults instantiates a new UpdateWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateWebhookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateWebhookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateWebhookRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateWebhookRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateWebhookRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateWebhookRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateWebhookRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateWebhookRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateWebhookRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateWebhookRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateWebhookRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateWebhookRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOnReception

`func (o *UpdateWebhookRequest) GetOnReception() bool`

GetOnReception returns the OnReception field if non-nil, zero value otherwise.

### GetOnReceptionOk

`func (o *UpdateWebhookRequest) GetOnReceptionOk() (*bool, bool)`

GetOnReceptionOk returns a tuple with the OnReception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnReception

`func (o *UpdateWebhookRequest) SetOnReception(v bool)`

SetOnReception sets OnReception field to given value.

### HasOnReception

`func (o *UpdateWebhookRequest) HasOnReception() bool`

HasOnReception returns a boolean if a field has been set.

### GetOnDelivered

`func (o *UpdateWebhookRequest) GetOnDelivered() bool`

GetOnDelivered returns the OnDelivered field if non-nil, zero value otherwise.

### GetOnDeliveredOk

`func (o *UpdateWebhookRequest) GetOnDeliveredOk() (*bool, bool)`

GetOnDeliveredOk returns a tuple with the OnDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDelivered

`func (o *UpdateWebhookRequest) SetOnDelivered(v bool)`

SetOnDelivered sets OnDelivered field to given value.

### HasOnDelivered

`func (o *UpdateWebhookRequest) HasOnDelivered() bool`

HasOnDelivered returns a boolean if a field has been set.

### GetOnTransientError

`func (o *UpdateWebhookRequest) GetOnTransientError() bool`

GetOnTransientError returns the OnTransientError field if non-nil, zero value otherwise.

### GetOnTransientErrorOk

`func (o *UpdateWebhookRequest) GetOnTransientErrorOk() (*bool, bool)`

GetOnTransientErrorOk returns a tuple with the OnTransientError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnTransientError

`func (o *UpdateWebhookRequest) SetOnTransientError(v bool)`

SetOnTransientError sets OnTransientError field to given value.

### HasOnTransientError

`func (o *UpdateWebhookRequest) HasOnTransientError() bool`

HasOnTransientError returns a boolean if a field has been set.

### GetOnFailed

`func (o *UpdateWebhookRequest) GetOnFailed() bool`

GetOnFailed returns the OnFailed field if non-nil, zero value otherwise.

### GetOnFailedOk

`func (o *UpdateWebhookRequest) GetOnFailedOk() (*bool, bool)`

GetOnFailedOk returns a tuple with the OnFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailed

`func (o *UpdateWebhookRequest) SetOnFailed(v bool)`

SetOnFailed sets OnFailed field to given value.

### HasOnFailed

`func (o *UpdateWebhookRequest) HasOnFailed() bool`

HasOnFailed returns a boolean if a field has been set.

### GetOnBounced

`func (o *UpdateWebhookRequest) GetOnBounced() bool`

GetOnBounced returns the OnBounced field if non-nil, zero value otherwise.

### GetOnBouncedOk

`func (o *UpdateWebhookRequest) GetOnBouncedOk() (*bool, bool)`

GetOnBouncedOk returns a tuple with the OnBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnBounced

`func (o *UpdateWebhookRequest) SetOnBounced(v bool)`

SetOnBounced sets OnBounced field to given value.

### HasOnBounced

`func (o *UpdateWebhookRequest) HasOnBounced() bool`

HasOnBounced returns a boolean if a field has been set.

### GetOnSuppressed

`func (o *UpdateWebhookRequest) GetOnSuppressed() bool`

GetOnSuppressed returns the OnSuppressed field if non-nil, zero value otherwise.

### GetOnSuppressedOk

`func (o *UpdateWebhookRequest) GetOnSuppressedOk() (*bool, bool)`

GetOnSuppressedOk returns a tuple with the OnSuppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSuppressed

`func (o *UpdateWebhookRequest) SetOnSuppressed(v bool)`

SetOnSuppressed sets OnSuppressed field to given value.

### HasOnSuppressed

`func (o *UpdateWebhookRequest) HasOnSuppressed() bool`

HasOnSuppressed returns a boolean if a field has been set.

### GetOnOpened

`func (o *UpdateWebhookRequest) GetOnOpened() bool`

GetOnOpened returns the OnOpened field if non-nil, zero value otherwise.

### GetOnOpenedOk

`func (o *UpdateWebhookRequest) GetOnOpenedOk() (*bool, bool)`

GetOnOpenedOk returns a tuple with the OnOpened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnOpened

`func (o *UpdateWebhookRequest) SetOnOpened(v bool)`

SetOnOpened sets OnOpened field to given value.

### HasOnOpened

`func (o *UpdateWebhookRequest) HasOnOpened() bool`

HasOnOpened returns a boolean if a field has been set.

### GetOnClicked

`func (o *UpdateWebhookRequest) GetOnClicked() bool`

GetOnClicked returns the OnClicked field if non-nil, zero value otherwise.

### GetOnClickedOk

`func (o *UpdateWebhookRequest) GetOnClickedOk() (*bool, bool)`

GetOnClickedOk returns a tuple with the OnClicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnClicked

`func (o *UpdateWebhookRequest) SetOnClicked(v bool)`

SetOnClicked sets OnClicked field to given value.

### HasOnClicked

`func (o *UpdateWebhookRequest) HasOnClicked() bool`

HasOnClicked returns a boolean if a field has been set.

### GetOnSuppressionCreated

`func (o *UpdateWebhookRequest) GetOnSuppressionCreated() bool`

GetOnSuppressionCreated returns the OnSuppressionCreated field if non-nil, zero value otherwise.

### GetOnSuppressionCreatedOk

`func (o *UpdateWebhookRequest) GetOnSuppressionCreatedOk() (*bool, bool)`

GetOnSuppressionCreatedOk returns a tuple with the OnSuppressionCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSuppressionCreated

`func (o *UpdateWebhookRequest) SetOnSuppressionCreated(v bool)`

SetOnSuppressionCreated sets OnSuppressionCreated field to given value.

### HasOnSuppressionCreated

`func (o *UpdateWebhookRequest) HasOnSuppressionCreated() bool`

HasOnSuppressionCreated returns a boolean if a field has been set.

### GetOnDnsError

`func (o *UpdateWebhookRequest) GetOnDnsError() bool`

GetOnDnsError returns the OnDnsError field if non-nil, zero value otherwise.

### GetOnDnsErrorOk

`func (o *UpdateWebhookRequest) GetOnDnsErrorOk() (*bool, bool)`

GetOnDnsErrorOk returns a tuple with the OnDnsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDnsError

`func (o *UpdateWebhookRequest) SetOnDnsError(v bool)`

SetOnDnsError sets OnDnsError field to given value.

### HasOnDnsError

`func (o *UpdateWebhookRequest) HasOnDnsError() bool`

HasOnDnsError returns a boolean if a field has been set.

### GetScope

`func (o *UpdateWebhookRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UpdateWebhookRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UpdateWebhookRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *UpdateWebhookRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetDomains

`func (o *UpdateWebhookRequest) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *UpdateWebhookRequest) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *UpdateWebhookRequest) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *UpdateWebhookRequest) HasDomains() bool`

HasDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


