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

// WebhookEventCallbackData The webhook event data sent to the subscribed webhook. 
type WebhookEventCallbackData struct {
	// The type of this resource: `webhook-events`
	Type string `json:"type"`
	// The unique identifier for this event. This will remain constant across delivery retries. 
	Id string `json:"id"`
	Attributes WebhookEventResourceAttributes `json:"attributes"`
	Relationships WebhookEventResourceRelationships `json:"relationships"`
}

// NewWebhookEventCallbackData instantiates a new WebhookEventCallbackData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookEventCallbackData(type_ string, id string, attributes WebhookEventResourceAttributes, relationships WebhookEventResourceRelationships) *WebhookEventCallbackData {
	this := WebhookEventCallbackData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewWebhookEventCallbackDataWithDefaults instantiates a new WebhookEventCallbackData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookEventCallbackDataWithDefaults() *WebhookEventCallbackData {
	this := WebhookEventCallbackData{}
	return &this
}

// GetType returns the Type field value
func (o *WebhookEventCallbackData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WebhookEventCallbackData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WebhookEventCallbackData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *WebhookEventCallbackData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebhookEventCallbackData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebhookEventCallbackData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *WebhookEventCallbackData) GetAttributes() WebhookEventResourceAttributes {
	if o == nil {
		var ret WebhookEventResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *WebhookEventCallbackData) GetAttributesOk() (*WebhookEventResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *WebhookEventCallbackData) SetAttributes(v WebhookEventResourceAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *WebhookEventCallbackData) GetRelationships() WebhookEventResourceRelationships {
	if o == nil {
		var ret WebhookEventResourceRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *WebhookEventCallbackData) GetRelationshipsOk() (*WebhookEventResourceRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *WebhookEventCallbackData) SetRelationships(v WebhookEventResourceRelationships) {
	o.Relationships = v
}

func (o WebhookEventCallbackData) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableWebhookEventCallbackData struct {
	value *WebhookEventCallbackData
	isSet bool
}

func (v NullableWebhookEventCallbackData) Get() *WebhookEventCallbackData {
	return v.value
}

func (v *NullableWebhookEventCallbackData) Set(val *WebhookEventCallbackData) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEventCallbackData) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEventCallbackData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEventCallbackData(val *WebhookEventCallbackData) *NullableWebhookEventCallbackData {
	return &NullableWebhookEventCallbackData{value: val, isSet: true}
}

func (v NullableWebhookEventCallbackData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEventCallbackData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


