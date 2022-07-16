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

// CategoryResourceRelationshipsParentData struct for CategoryResourceRelationshipsParentData
type CategoryResourceRelationshipsParentData struct {
	// The type of this resource: `categories`
	Type string `json:"type"`
	// The unique identifier of the resource within its type. 
	Id string `json:"id"`
}

// NewCategoryResourceRelationshipsParentData instantiates a new CategoryResourceRelationshipsParentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryResourceRelationshipsParentData(type_ string, id string) *CategoryResourceRelationshipsParentData {
	this := CategoryResourceRelationshipsParentData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewCategoryResourceRelationshipsParentDataWithDefaults instantiates a new CategoryResourceRelationshipsParentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryResourceRelationshipsParentDataWithDefaults() *CategoryResourceRelationshipsParentData {
	this := CategoryResourceRelationshipsParentData{}
	return &this
}

// GetType returns the Type field value
func (o *CategoryResourceRelationshipsParentData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CategoryResourceRelationshipsParentData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CategoryResourceRelationshipsParentData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CategoryResourceRelationshipsParentData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CategoryResourceRelationshipsParentData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CategoryResourceRelationshipsParentData) SetId(v string) {
	o.Id = v
}

func (o CategoryResourceRelationshipsParentData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableCategoryResourceRelationshipsParentData struct {
	value *CategoryResourceRelationshipsParentData
	isSet bool
}

func (v NullableCategoryResourceRelationshipsParentData) Get() *CategoryResourceRelationshipsParentData {
	return v.value
}

func (v *NullableCategoryResourceRelationshipsParentData) Set(val *CategoryResourceRelationshipsParentData) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryResourceRelationshipsParentData) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryResourceRelationshipsParentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryResourceRelationshipsParentData(val *CategoryResourceRelationshipsParentData) *NullableCategoryResourceRelationshipsParentData {
	return &NullableCategoryResourceRelationshipsParentData{value: val, isSet: true}
}

func (v NullableCategoryResourceRelationshipsParentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryResourceRelationshipsParentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


