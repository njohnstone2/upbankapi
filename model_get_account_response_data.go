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

// GetAccountResponseData The account returned in this response. 
type GetAccountResponseData struct {
	// The type of this resource: `accounts`
	Type string `json:"type"`
	// The unique identifier for this account. 
	Id string `json:"id"`
	Attributes AccountResourceAttributes `json:"attributes"`
	Relationships AccountResourceRelationships `json:"relationships"`
	Links *AccountResourceLinks `json:"links,omitempty"`
}

// NewGetAccountResponseData instantiates a new GetAccountResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountResponseData(type_ string, id string, attributes AccountResourceAttributes, relationships AccountResourceRelationships) *GetAccountResponseData {
	this := GetAccountResponseData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewGetAccountResponseDataWithDefaults instantiates a new GetAccountResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountResponseDataWithDefaults() *GetAccountResponseData {
	this := GetAccountResponseData{}
	return &this
}

// GetType returns the Type field value
func (o *GetAccountResponseData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetAccountResponseData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetAccountResponseData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GetAccountResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetAccountResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetAccountResponseData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *GetAccountResponseData) GetAttributes() AccountResourceAttributes {
	if o == nil {
		var ret AccountResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GetAccountResponseData) GetAttributesOk() (*AccountResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GetAccountResponseData) SetAttributes(v AccountResourceAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *GetAccountResponseData) GetRelationships() AccountResourceRelationships {
	if o == nil {
		var ret AccountResourceRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *GetAccountResponseData) GetRelationshipsOk() (*AccountResourceRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *GetAccountResponseData) SetRelationships(v AccountResourceRelationships) {
	o.Relationships = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetAccountResponseData) GetLinks() AccountResourceLinks {
	if o == nil || o.Links == nil {
		var ret AccountResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountResponseData) GetLinksOk() (*AccountResourceLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetAccountResponseData) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AccountResourceLinks and assigns it to the Links field.
func (o *GetAccountResponseData) SetLinks(v AccountResourceLinks) {
	o.Links = &v
}

func (o GetAccountResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["relationships"] = o.Relationships
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableGetAccountResponseData struct {
	value *GetAccountResponseData
	isSet bool
}

func (v NullableGetAccountResponseData) Get() *GetAccountResponseData {
	return v.value
}

func (v *NullableGetAccountResponseData) Set(val *GetAccountResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountResponseData(val *GetAccountResponseData) *NullableGetAccountResponseData {
	return &NullableGetAccountResponseData{value: val, isSet: true}
}

func (v NullableGetAccountResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


