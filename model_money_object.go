/*
Up API

The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package upbankapi

import (
	"encoding/json"
)

// MoneyObject Provides information about a value of money. 
type MoneyObject struct {
	// The ISO 4217 currency code. 
	CurrencyCode string `json:"currencyCode"`
	// The amount of money, formatted as a string in the relevant currency. For example, for an Australian dollar value of $10.56, this field will be `\"10.56\"`. The currency symbol is not included in the string. 
	Value string `json:"value"`
	// The amount of money in the smallest denomination for the currency, as a 64-bit integer.  For example, for an Australian dollar value of $10.56, this field will be `1056`. 
	ValueInBaseUnits int32 `json:"valueInBaseUnits"`
}

// NewMoneyObject instantiates a new MoneyObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoneyObject(currencyCode string, value string, valueInBaseUnits int32) *MoneyObject {
	this := MoneyObject{}
	this.CurrencyCode = currencyCode
	this.Value = value
	this.ValueInBaseUnits = valueInBaseUnits
	return &this
}

// NewMoneyObjectWithDefaults instantiates a new MoneyObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoneyObjectWithDefaults() *MoneyObject {
	this := MoneyObject{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *MoneyObject) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *MoneyObject) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *MoneyObject) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetValue returns the Value field value
func (o *MoneyObject) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *MoneyObject) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *MoneyObject) SetValue(v string) {
	o.Value = v
}

// GetValueInBaseUnits returns the ValueInBaseUnits field value
func (o *MoneyObject) GetValueInBaseUnits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ValueInBaseUnits
}

// GetValueInBaseUnitsOk returns a tuple with the ValueInBaseUnits field value
// and a boolean to check if the value has been set.
func (o *MoneyObject) GetValueInBaseUnitsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueInBaseUnits, true
}

// SetValueInBaseUnits sets field value
func (o *MoneyObject) SetValueInBaseUnits(v int32) {
	o.ValueInBaseUnits = v
}

func (o MoneyObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["valueInBaseUnits"] = o.ValueInBaseUnits
	}
	return json.Marshal(toSerialize)
}

type NullableMoneyObject struct {
	value *MoneyObject
	isSet bool
}

func (v NullableMoneyObject) Get() *MoneyObject {
	return v.value
}

func (v *NullableMoneyObject) Set(val *MoneyObject) {
	v.value = val
	v.isSet = true
}

func (v NullableMoneyObject) IsSet() bool {
	return v.isSet
}

func (v *NullableMoneyObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoneyObject(val *MoneyObject) *NullableMoneyObject {
	return &NullableMoneyObject{value: val, isSet: true}
}

func (v NullableMoneyObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoneyObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


