# GetTransactionResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of this resource: &#x60;transactions&#x60; | 
**Id** | **string** | The unique identifier for this transaction.  | 
**Attributes** | [**TransactionResourceAttributes**](TransactionResourceAttributes.md) |  | 
**Relationships** | [**TransactionResourceRelationships**](TransactionResourceRelationships.md) |  | 
**Links** | Pointer to [**AccountResourceLinks**](AccountResourceLinks.md) |  | [optional] 

## Methods

### NewGetTransactionResponseData

`func NewGetTransactionResponseData(type_ string, id string, attributes TransactionResourceAttributes, relationships TransactionResourceRelationships, ) *GetTransactionResponseData`

NewGetTransactionResponseData instantiates a new GetTransactionResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionResponseDataWithDefaults

`func NewGetTransactionResponseDataWithDefaults() *GetTransactionResponseData`

NewGetTransactionResponseDataWithDefaults instantiates a new GetTransactionResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetTransactionResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTransactionResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTransactionResponseData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GetTransactionResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTransactionResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTransactionResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GetTransactionResponseData) GetAttributes() TransactionResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetTransactionResponseData) GetAttributesOk() (*TransactionResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetTransactionResponseData) SetAttributes(v TransactionResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *GetTransactionResponseData) GetRelationships() TransactionResourceRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GetTransactionResponseData) GetRelationshipsOk() (*TransactionResourceRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GetTransactionResponseData) SetRelationships(v TransactionResourceRelationships)`

SetRelationships sets Relationships field to given value.


### GetLinks

`func (o *GetTransactionResponseData) GetLinks() AccountResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetTransactionResponseData) GetLinksOk() (*AccountResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetTransactionResponseData) SetLinks(v AccountResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetTransactionResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


