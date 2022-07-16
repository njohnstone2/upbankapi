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

// AccountResourceRelationships struct for AccountResourceRelationships
type AccountResourceRelationships struct {
	Transactions AccountResourceRelationshipsTransactions `json:"transactions"`
}

// NewAccountResourceRelationships instantiates a new AccountResourceRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountResourceRelationships(transactions AccountResourceRelationshipsTransactions) *AccountResourceRelationships {
	this := AccountResourceRelationships{}
	this.Transactions = transactions
	return &this
}

// NewAccountResourceRelationshipsWithDefaults instantiates a new AccountResourceRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountResourceRelationshipsWithDefaults() *AccountResourceRelationships {
	this := AccountResourceRelationships{}
	return &this
}

// GetTransactions returns the Transactions field value
func (o *AccountResourceRelationships) GetTransactions() AccountResourceRelationshipsTransactions {
	if o == nil {
		var ret AccountResourceRelationshipsTransactions
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *AccountResourceRelationships) GetTransactionsOk() (*AccountResourceRelationshipsTransactions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transactions, true
}

// SetTransactions sets field value
func (o *AccountResourceRelationships) SetTransactions(v AccountResourceRelationshipsTransactions) {
	o.Transactions = v
}

func (o AccountResourceRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transactions"] = o.Transactions
	}
	return json.Marshal(toSerialize)
}

type NullableAccountResourceRelationships struct {
	value *AccountResourceRelationships
	isSet bool
}

func (v NullableAccountResourceRelationships) Get() *AccountResourceRelationships {
	return v.value
}

func (v *NullableAccountResourceRelationships) Set(val *AccountResourceRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountResourceRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountResourceRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountResourceRelationships(val *AccountResourceRelationships) *NullableAccountResourceRelationships {
	return &NullableAccountResourceRelationships{value: val, isSet: true}
}

func (v NullableAccountResourceRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountResourceRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


