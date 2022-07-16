# GetWebhookResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of this resource: &#x60;webhooks&#x60; | 
**Id** | **string** | The unique identifier for this webhook.  | 
**Attributes** | [**WebhookResourceAttributes**](WebhookResourceAttributes.md) |  | 
**Relationships** | [**WebhookResourceRelationships**](WebhookResourceRelationships.md) |  | 
**Links** | Pointer to [**AccountResourceLinks**](AccountResourceLinks.md) |  | [optional] 

## Methods

### NewGetWebhookResponseData

`func NewGetWebhookResponseData(type_ string, id string, attributes WebhookResourceAttributes, relationships WebhookResourceRelationships, ) *GetWebhookResponseData`

NewGetWebhookResponseData instantiates a new GetWebhookResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWebhookResponseDataWithDefaults

`func NewGetWebhookResponseDataWithDefaults() *GetWebhookResponseData`

NewGetWebhookResponseDataWithDefaults instantiates a new GetWebhookResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetWebhookResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetWebhookResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetWebhookResponseData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GetWebhookResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetWebhookResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetWebhookResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GetWebhookResponseData) GetAttributes() WebhookResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetWebhookResponseData) GetAttributesOk() (*WebhookResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetWebhookResponseData) SetAttributes(v WebhookResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *GetWebhookResponseData) GetRelationships() WebhookResourceRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GetWebhookResponseData) GetRelationshipsOk() (*WebhookResourceRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GetWebhookResponseData) SetRelationships(v WebhookResourceRelationships)`

SetRelationships sets Relationships field to given value.


### GetLinks

`func (o *GetWebhookResponseData) GetLinks() AccountResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetWebhookResponseData) GetLinksOk() (*AccountResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetWebhookResponseData) SetLinks(v AccountResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetWebhookResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


