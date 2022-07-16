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

// ListTagsResponse Successful response to get all tags. This returns a paginated list of tags, which can be scrolled by following the `prev` and `next` links if present. 
type ListTagsResponse struct {
	// The list of tags returned in this response. 
	Data []TagResource `json:"data"`
	Links ListAccountsResponseLinks `json:"links"`
}

// NewListTagsResponse instantiates a new ListTagsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTagsResponse(data []TagResource, links ListAccountsResponseLinks) *ListTagsResponse {
	this := ListTagsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewListTagsResponseWithDefaults instantiates a new ListTagsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTagsResponseWithDefaults() *ListTagsResponse {
	this := ListTagsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListTagsResponse) GetData() []TagResource {
	if o == nil {
		var ret []TagResource
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListTagsResponse) GetDataOk() ([]TagResource, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListTagsResponse) SetData(v []TagResource) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *ListTagsResponse) GetLinks() ListAccountsResponseLinks {
	if o == nil {
		var ret ListAccountsResponseLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ListTagsResponse) GetLinksOk() (*ListAccountsResponseLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ListTagsResponse) SetLinks(v ListAccountsResponseLinks) {
	o.Links = v
}

func (o ListTagsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableListTagsResponse struct {
	value *ListTagsResponse
	isSet bool
}

func (v NullableListTagsResponse) Get() *ListTagsResponse {
	return v.value
}

func (v *NullableListTagsResponse) Set(val *ListTagsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListTagsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListTagsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTagsResponse(val *ListTagsResponse) *NullableListTagsResponse {
	return &NullableListTagsResponse{value: val, isSet: true}
}

func (v NullableListTagsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTagsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


