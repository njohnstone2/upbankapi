# CashbackObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A brief description of why this cashback was paid.  | 
**Amount** | [**CashbackObjectAmount**](CashbackObjectAmount.md) |  | 

## Methods

### NewCashbackObject

`func NewCashbackObject(description string, amount CashbackObjectAmount, ) *CashbackObject`

NewCashbackObject instantiates a new CashbackObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashbackObjectWithDefaults

`func NewCashbackObjectWithDefaults() *CashbackObject`

NewCashbackObjectWithDefaults instantiates a new CashbackObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CashbackObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CashbackObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CashbackObject) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAmount

`func (o *CashbackObject) GetAmount() CashbackObjectAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CashbackObject) GetAmountOk() (*CashbackObjectAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CashbackObject) SetAmount(v CashbackObjectAmount)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


