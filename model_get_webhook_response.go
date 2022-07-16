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

// GetWebhookResponse Successful response to get a single webhook. 
type GetWebhookResponse struct {
	Data GetWebhookResponseData `json:"data"`
}

// NewGetWebhookResponse instantiates a new GetWebhookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWebhookResponse(data GetWebhookResponseData) *GetWebhookResponse {
	this := GetWebhookResponse{}
	this.Data = data
	return &this
}

// NewGetWebhookResponseWithDefaults instantiates a new GetWebhookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWebhookResponseWithDefaults() *GetWebhookResponse {
	this := GetWebhookResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetWebhookResponse) GetData() GetWebhookResponseData {
	if o == nil {
		var ret GetWebhookResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetWebhookResponse) GetDataOk() (*GetWebhookResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetWebhookResponse) SetData(v GetWebhookResponseData) {
	o.Data = v
}

func (o GetWebhookResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetWebhookResponse struct {
	value *GetWebhookResponse
	isSet bool
}

func (v NullableGetWebhookResponse) Get() *GetWebhookResponse {
	return v.value
}

func (v *NullableGetWebhookResponse) Set(val *GetWebhookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWebhookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWebhookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWebhookResponse(val *GetWebhookResponse) *NullableGetWebhookResponse {
	return &NullableGetWebhookResponse{value: val, isSet: true}
}

func (v NullableGetWebhookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWebhookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

