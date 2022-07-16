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

// CreateWebhookResponseData The webhook that was created. 
type CreateWebhookResponseData struct {
	// The type of this resource: `webhooks`
	Type string `json:"type"`
	// The unique identifier for this webhook. 
	Id string `json:"id"`
	Attributes WebhookResourceAttributes `json:"attributes"`
	Relationships WebhookResourceRelationships `json:"relationships"`
	Links *AccountResourceLinks `json:"links,omitempty"`
}

// NewCreateWebhookResponseData instantiates a new CreateWebhookResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWebhookResponseData(type_ string, id string, attributes WebhookResourceAttributes, relationships WebhookResourceRelationships) *CreateWebhookResponseData {
	this := CreateWebhookResponseData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewCreateWebhookResponseDataWithDefaults instantiates a new CreateWebhookResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWebhookResponseDataWithDefaults() *CreateWebhookResponseData {
	this := CreateWebhookResponseData{}
	return &this
}

// GetType returns the Type field value
func (o *CreateWebhookResponseData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateWebhookResponseData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateWebhookResponseData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CreateWebhookResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateWebhookResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateWebhookResponseData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *CreateWebhookResponseData) GetAttributes() WebhookResourceAttributes {
	if o == nil {
		var ret WebhookResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CreateWebhookResponseData) GetAttributesOk() (*WebhookResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CreateWebhookResponseData) SetAttributes(v WebhookResourceAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *CreateWebhookResponseData) GetRelationships() WebhookResourceRelationships {
	if o == nil {
		var ret WebhookResourceRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *CreateWebhookResponseData) GetRelationshipsOk() (*WebhookResourceRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *CreateWebhookResponseData) SetRelationships(v WebhookResourceRelationships) {
	o.Relationships = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CreateWebhookResponseData) GetLinks() AccountResourceLinks {
	if o == nil || o.Links == nil {
		var ret AccountResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebhookResponseData) GetLinksOk() (*AccountResourceLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CreateWebhookResponseData) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AccountResourceLinks and assigns it to the Links field.
func (o *CreateWebhookResponseData) SetLinks(v AccountResourceLinks) {
	o.Links = &v
}

func (o CreateWebhookResponseData) MarshalJSON() ([]byte, error) {
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

type NullableCreateWebhookResponseData struct {
	value *CreateWebhookResponseData
	isSet bool
}

func (v NullableCreateWebhookResponseData) Get() *CreateWebhookResponseData {
	return v.value
}

func (v *NullableCreateWebhookResponseData) Set(val *CreateWebhookResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWebhookResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWebhookResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWebhookResponseData(val *CreateWebhookResponseData) *NullableCreateWebhookResponseData {
	return &NullableCreateWebhookResponseData{value: val, isSet: true}
}

func (v NullableCreateWebhookResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWebhookResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

