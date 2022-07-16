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

// CreateWebhookRequestData The webhook resource to create. 
type CreateWebhookRequestData struct {
	Attributes WebhookInputResourceAttributes `json:"attributes"`
}

// NewCreateWebhookRequestData instantiates a new CreateWebhookRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWebhookRequestData(attributes WebhookInputResourceAttributes) *CreateWebhookRequestData {
	this := CreateWebhookRequestData{}
	this.Attributes = attributes
	return &this
}

// NewCreateWebhookRequestDataWithDefaults instantiates a new CreateWebhookRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWebhookRequestDataWithDefaults() *CreateWebhookRequestData {
	this := CreateWebhookRequestData{}
	return &this
}

// GetAttributes returns the Attributes field value
func (o *CreateWebhookRequestData) GetAttributes() WebhookInputResourceAttributes {
	if o == nil {
		var ret WebhookInputResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CreateWebhookRequestData) GetAttributesOk() (*WebhookInputResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CreateWebhookRequestData) SetAttributes(v WebhookInputResourceAttributes) {
	o.Attributes = v
}

func (o CreateWebhookRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableCreateWebhookRequestData struct {
	value *CreateWebhookRequestData
	isSet bool
}

func (v NullableCreateWebhookRequestData) Get() *CreateWebhookRequestData {
	return v.value
}

func (v *NullableCreateWebhookRequestData) Set(val *CreateWebhookRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWebhookRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWebhookRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWebhookRequestData(val *CreateWebhookRequestData) *NullableCreateWebhookRequestData {
	return &NullableCreateWebhookRequestData{value: val, isSet: true}
}

func (v NullableCreateWebhookRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWebhookRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

