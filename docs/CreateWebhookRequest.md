# CreateWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Webhook name | 
**Url** | **string** | Webhook URL | 
**Enabled** | Pointer to **bool** | Whether the webhook is enabled | [optional] [default to true]
**OnReception** | Pointer to **bool** | Trigger on message reception | [optional] [default to false]
**OnDelivered** | Pointer to **bool** | Trigger on message delivery | [optional] [default to false]
**OnTransientError** | Pointer to **bool** | Trigger on transient errors | [optional] [default to false]
**OnFailed** | Pointer to **bool** | Trigger on permanent failures | [optional] [default to false]
**OnBounced** | Pointer to **bool** | Trigger on bounces | [optional] [default to false]
**OnSuppressed** | Pointer to **bool** | Trigger on suppressions | [optional] [default to false]
**OnOpened** | Pointer to **bool** | Trigger on opens | [optional] [default to false]
**OnClicked** | Pointer to **bool** | Trigger on clicks | [optional] [default to false]
**OnSuppressionCreated** | Pointer to **bool** | Trigger on suppression creation | [optional] [default to false]
**OnDnsError** | Pointer to **bool** | Trigger on DNS errors | [optional] [default to false]
**Scope** | Pointer to **string** | Webhook scope | [optional] 
**Domains** | Pointer to **[]string** | Domains this webhook applies to | [optional] 

## Methods

### NewCreateWebhookRequest

`func NewCreateWebhookRequest(name string, url string, ) *CreateWebhookRequest`

NewCreateWebhookRequest instantiates a new CreateWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebhookRequestWithDefaults

`func NewCreateWebhookRequestWithDefaults() *CreateWebhookRequest`

NewCreateWebhookRequestWithDefaults instantiates a new CreateWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateWebhookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWebhookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWebhookRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *CreateWebhookRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateWebhookRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateWebhookRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEnabled

`func (o *CreateWebhookRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateWebhookRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateWebhookRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateWebhookRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOnReception

`func (o *CreateWebhookRequest) GetOnReception() bool`

GetOnReception returns the OnReception field if non-nil, zero value otherwise.

### GetOnReceptionOk

`func (o *CreateWebhookRequest) GetOnReceptionOk() (*bool, bool)`

GetOnReceptionOk returns a tuple with the OnReception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnReception

`func (o *CreateWebhookRequest) SetOnReception(v bool)`

SetOnReception sets OnReception field to given value.

### HasOnReception

`func (o *CreateWebhookRequest) HasOnReception() bool`

HasOnReception returns a boolean if a field has been set.

### GetOnDelivered

`func (o *CreateWebhookRequest) GetOnDelivered() bool`

GetOnDelivered returns the OnDelivered field if non-nil, zero value otherwise.

### GetOnDeliveredOk

`func (o *CreateWebhookRequest) GetOnDeliveredOk() (*bool, bool)`

GetOnDeliveredOk returns a tuple with the OnDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDelivered

`func (o *CreateWebhookRequest) SetOnDelivered(v bool)`

SetOnDelivered sets OnDelivered field to given value.

### HasOnDelivered

`func (o *CreateWebhookRequest) HasOnDelivered() bool`

HasOnDelivered returns a boolean if a field has been set.

### GetOnTransientError

`func (o *CreateWebhookRequest) GetOnTransientError() bool`

GetOnTransientError returns the OnTransientError field if non-nil, zero value otherwise.

### GetOnTransientErrorOk

`func (o *CreateWebhookRequest) GetOnTransientErrorOk() (*bool, bool)`

GetOnTransientErrorOk returns a tuple with the OnTransientError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnTransientError

`func (o *CreateWebhookRequest) SetOnTransientError(v bool)`

SetOnTransientError sets OnTransientError field to given value.

### HasOnTransientError

`func (o *CreateWebhookRequest) HasOnTransientError() bool`

HasOnTransientError returns a boolean if a field has been set.

### GetOnFailed

`func (o *CreateWebhookRequest) GetOnFailed() bool`

GetOnFailed returns the OnFailed field if non-nil, zero value otherwise.

### GetOnFailedOk

`func (o *CreateWebhookRequest) GetOnFailedOk() (*bool, bool)`

GetOnFailedOk returns a tuple with the OnFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailed

`func (o *CreateWebhookRequest) SetOnFailed(v bool)`

SetOnFailed sets OnFailed field to given value.

### HasOnFailed

`func (o *CreateWebhookRequest) HasOnFailed() bool`

HasOnFailed returns a boolean if a field has been set.

### GetOnBounced

`func (o *CreateWebhookRequest) GetOnBounced() bool`

GetOnBounced returns the OnBounced field if non-nil, zero value otherwise.

### GetOnBouncedOk

`func (o *CreateWebhookRequest) GetOnBouncedOk() (*bool, bool)`

GetOnBouncedOk returns a tuple with the OnBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnBounced

`func (o *CreateWebhookRequest) SetOnBounced(v bool)`

SetOnBounced sets OnBounced field to given value.

### HasOnBounced

`func (o *CreateWebhookRequest) HasOnBounced() bool`

HasOnBounced returns a boolean if a field has been set.

### GetOnSuppressed

`func (o *CreateWebhookRequest) GetOnSuppressed() bool`

GetOnSuppressed returns the OnSuppressed field if non-nil, zero value otherwise.

### GetOnSuppressedOk

`func (o *CreateWebhookRequest) GetOnSuppressedOk() (*bool, bool)`

GetOnSuppressedOk returns a tuple with the OnSuppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSuppressed

`func (o *CreateWebhookRequest) SetOnSuppressed(v bool)`

SetOnSuppressed sets OnSuppressed field to given value.

### HasOnSuppressed

`func (o *CreateWebhookRequest) HasOnSuppressed() bool`

HasOnSuppressed returns a boolean if a field has been set.

### GetOnOpened

`func (o *CreateWebhookRequest) GetOnOpened() bool`

GetOnOpened returns the OnOpened field if non-nil, zero value otherwise.

### GetOnOpenedOk

`func (o *CreateWebhookRequest) GetOnOpenedOk() (*bool, bool)`

GetOnOpenedOk returns a tuple with the OnOpened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnOpened

`func (o *CreateWebhookRequest) SetOnOpened(v bool)`

SetOnOpened sets OnOpened field to given value.

### HasOnOpened

`func (o *CreateWebhookRequest) HasOnOpened() bool`

HasOnOpened returns a boolean if a field has been set.

### GetOnClicked

`func (o *CreateWebhookRequest) GetOnClicked() bool`

GetOnClicked returns the OnClicked field if non-nil, zero value otherwise.

### GetOnClickedOk

`func (o *CreateWebhookRequest) GetOnClickedOk() (*bool, bool)`

GetOnClickedOk returns a tuple with the OnClicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnClicked

`func (o *CreateWebhookRequest) SetOnClicked(v bool)`

SetOnClicked sets OnClicked field to given value.

### HasOnClicked

`func (o *CreateWebhookRequest) HasOnClicked() bool`

HasOnClicked returns a boolean if a field has been set.

### GetOnSuppressionCreated

`func (o *CreateWebhookRequest) GetOnSuppressionCreated() bool`

GetOnSuppressionCreated returns the OnSuppressionCreated field if non-nil, zero value otherwise.

### GetOnSuppressionCreatedOk

`func (o *CreateWebhookRequest) GetOnSuppressionCreatedOk() (*bool, bool)`

GetOnSuppressionCreatedOk returns a tuple with the OnSuppressionCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSuppressionCreated

`func (o *CreateWebhookRequest) SetOnSuppressionCreated(v bool)`

SetOnSuppressionCreated sets OnSuppressionCreated field to given value.

### HasOnSuppressionCreated

`func (o *CreateWebhookRequest) HasOnSuppressionCreated() bool`

HasOnSuppressionCreated returns a boolean if a field has been set.

### GetOnDnsError

`func (o *CreateWebhookRequest) GetOnDnsError() bool`

GetOnDnsError returns the OnDnsError field if non-nil, zero value otherwise.

### GetOnDnsErrorOk

`func (o *CreateWebhookRequest) GetOnDnsErrorOk() (*bool, bool)`

GetOnDnsErrorOk returns a tuple with the OnDnsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDnsError

`func (o *CreateWebhookRequest) SetOnDnsError(v bool)`

SetOnDnsError sets OnDnsError field to given value.

### HasOnDnsError

`func (o *CreateWebhookRequest) HasOnDnsError() bool`

HasOnDnsError returns a boolean if a field has been set.

### GetScope

`func (o *CreateWebhookRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateWebhookRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateWebhookRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CreateWebhookRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetDomains

`func (o *CreateWebhookRequest) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *CreateWebhookRequest) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *CreateWebhookRequest) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *CreateWebhookRequest) HasDomains() bool`

HasDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


