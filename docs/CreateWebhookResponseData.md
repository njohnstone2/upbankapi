# CreateWebhookResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of this resource: &#x60;webhooks&#x60; | 
**Id** | **string** | The unique identifier for this webhook.  | 
**Attributes** | [**WebhookResourceAttributes**](WebhookResourceAttributes.md) |  | 
**Relationships** | [**WebhookResourceRelationships**](WebhookResourceRelationships.md) |  | 
**Links** | Pointer to [**AccountResourceLinks**](AccountResourceLinks.md) |  | [optional] 

## Methods

### NewCreateWebhookResponseData

`func NewCreateWebhookResponseData(type_ string, id string, attributes WebhookResourceAttributes, relationships WebhookResourceRelationships, ) *CreateWebhookResponseData`

NewCreateWebhookResponseData instantiates a new CreateWebhookResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebhookResponseDataWithDefaults

`func NewCreateWebhookResponseDataWithDefaults() *CreateWebhookResponseData`

NewCreateWebhookResponseDataWithDefaults instantiates a new CreateWebhookResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateWebhookResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateWebhookResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateWebhookResponseData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CreateWebhookResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateWebhookResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateWebhookResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CreateWebhookResponseData) GetAttributes() WebhookResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CreateWebhookResponseData) GetAttributesOk() (*WebhookResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CreateWebhookResponseData) SetAttributes(v WebhookResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CreateWebhookResponseData) GetRelationships() WebhookResourceRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CreateWebhookResponseData) GetRelationshipsOk() (*WebhookResourceRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CreateWebhookResponseData) SetRelationships(v WebhookResourceRelationships)`

SetRelationships sets Relationships field to given value.


### GetLinks

`func (o *CreateWebhookResponseData) GetLinks() AccountResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateWebhookResponseData) GetLinksOk() (*AccountResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateWebhookResponseData) SetLinks(v AccountResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateWebhookResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


