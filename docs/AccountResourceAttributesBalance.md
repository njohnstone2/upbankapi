# AccountResourceAttributesBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | **string** | The ISO 4217 currency code.  | 
**Value** | **string** | The amount of money, formatted as a string in the relevant currency. For example, for an Australian dollar value of $10.56, this field will be &#x60;\&quot;10.56\&quot;&#x60;. The currency symbol is not included in the string.  | 
**ValueInBaseUnits** | **int32** | The amount of money in the smallest denomination for the currency, as a 64-bit integer.  For example, for an Australian dollar value of $10.56, this field will be &#x60;1056&#x60;.  | 

## Methods

### NewAccountResourceAttributesBalance

`func NewAccountResourceAttributesBalance(currencyCode string, value string, valueInBaseUnits int32, ) *AccountResourceAttributesBalance`

NewAccountResourceAttributesBalance instantiates a new AccountResourceAttributesBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountResourceAttributesBalanceWithDefaults

`func NewAccountResourceAttributesBalanceWithDefaults() *AccountResourceAttributesBalance`

NewAccountResourceAttributesBalanceWithDefaults instantiates a new AccountResourceAttributesBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyCode

`func (o *AccountResourceAttributesBalance) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AccountResourceAttributesBalance) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AccountResourceAttributesBalance) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetValue

`func (o *AccountResourceAttributesBalance) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AccountResourceAttributesBalance) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AccountResourceAttributesBalance) SetValue(v string)`

SetValue sets Value field to given value.


### GetValueInBaseUnits

`func (o *AccountResourceAttributesBalance) GetValueInBaseUnits() int32`

GetValueInBaseUnits returns the ValueInBaseUnits field if non-nil, zero value otherwise.

### GetValueInBaseUnitsOk

`func (o *AccountResourceAttributesBalance) GetValueInBaseUnitsOk() (*int32, bool)`

GetValueInBaseUnitsOk returns a tuple with the ValueInBaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueInBaseUnits

`func (o *AccountResourceAttributesBalance) SetValueInBaseUnits(v int32)`

SetValueInBaseUnits sets ValueInBaseUnits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


