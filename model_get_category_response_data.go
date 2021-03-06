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

// GetCategoryResponseData The category returned in this response. 
type GetCategoryResponseData struct {
	// The type of this resource: `categories`
	Type string `json:"type"`
	// The unique identifier for this category. This is a human-readable but URL-safe value. 
	Id string `json:"id"`
	Attributes CategoryResourceAttributes `json:"attributes"`
	Relationships CategoryResourceRelationships `json:"relationships"`
	Links *AccountResourceLinks `json:"links,omitempty"`
}

// NewGetCategoryResponseData instantiates a new GetCategoryResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCategoryResponseData(type_ string, id string, attributes CategoryResourceAttributes, relationships CategoryResourceRelationships) *GetCategoryResponseData {
	this := GetCategoryResponseData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewGetCategoryResponseDataWithDefaults instantiates a new GetCategoryResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCategoryResponseDataWithDefaults() *GetCategoryResponseData {
	this := GetCategoryResponseData{}
	return &this
}

// GetType returns the Type field value
func (o *GetCategoryResponseData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetCategoryResponseData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetCategoryResponseData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GetCategoryResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetCategoryResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetCategoryResponseData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *GetCategoryResponseData) GetAttributes() CategoryResourceAttributes {
	if o == nil {
		var ret CategoryResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GetCategoryResponseData) GetAttributesOk() (*CategoryResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GetCategoryResponseData) SetAttributes(v CategoryResourceAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *GetCategoryResponseData) GetRelationships() CategoryResourceRelationships {
	if o == nil {
		var ret CategoryResourceRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *GetCategoryResponseData) GetRelationshipsOk() (*CategoryResourceRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *GetCategoryResponseData) SetRelationships(v CategoryResourceRelationships) {
	o.Relationships = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetCategoryResponseData) GetLinks() AccountResourceLinks {
	if o == nil || o.Links == nil {
		var ret AccountResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCategoryResponseData) GetLinksOk() (*AccountResourceLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetCategoryResponseData) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AccountResourceLinks and assigns it to the Links field.
func (o *GetCategoryResponseData) SetLinks(v AccountResourceLinks) {
	o.Links = &v
}

func (o GetCategoryResponseData) MarshalJSON() ([]byte, error) {
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

type NullableGetCategoryResponseData struct {
	value *GetCategoryResponseData
	isSet bool
}

func (v NullableGetCategoryResponseData) Get() *GetCategoryResponseData {
	return v.value
}

func (v *NullableGetCategoryResponseData) Set(val *GetCategoryResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCategoryResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCategoryResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCategoryResponseData(val *GetCategoryResponseData) *NullableGetCategoryResponseData {
	return &NullableGetCategoryResponseData{value: val, isSet: true}
}

func (v NullableGetCategoryResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCategoryResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


