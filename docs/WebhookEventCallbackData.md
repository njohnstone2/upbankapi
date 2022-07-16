# WebhookEventCallbackData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of this resource: &#x60;webhook-events&#x60; | 
**Id** | **string** | The unique identifier for this event. This will remain constant across delivery retries.  | 
**Attributes** | [**WebhookEventResourceAttributes**](WebhookEventResourceAttributes.md) |  | 
**Relationships** | [**WebhookEventResourceRelationships**](WebhookEventResourceRelationships.md) |  | 

## Methods

### NewWebhookEventCallbackData

`func NewWebhookEventCallbackData(type_ string, id string, attributes WebhookEventResourceAttributes, relationships WebhookEventResourceRelationships, ) *WebhookEventCallbackData`

NewWebhookEventCallbackData instantiates a new WebhookEventCallbackData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEventCallbackDataWithDefaults

`func NewWebhookEventCallbackDataWithDefaults() *WebhookEventCallbackData`

NewWebhookEventCallbackDataWithDefaults instantiates a new WebhookEventCallbackData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WebhookEventCallbackData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookEventCallbackData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookEventCallbackData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *WebhookEventCallbackData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookEventCallbackData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookEventCallbackData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *WebhookEventCallbackData) GetAttributes() WebhookEventResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WebhookEventCallbackData) GetAttributesOk() (*WebhookEventResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WebhookEventCallbackData) SetAttributes(v WebhookEventResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *WebhookEventCallbackData) GetRelationships() WebhookEventResourceRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WebhookEventCallbackData) GetRelationshipsOk() (*WebhookEventResourceRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WebhookEventCallbackData) SetRelationships(v WebhookEventResourceRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


