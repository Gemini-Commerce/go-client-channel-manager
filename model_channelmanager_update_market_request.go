/*
Channel Manager Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package channelmanager

import (
	"encoding/json"
)

// checks if the ChannelmanagerUpdateMarketRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelmanagerUpdateMarketRequest{}

// ChannelmanagerUpdateMarketRequest struct for ChannelmanagerUpdateMarketRequest
type ChannelmanagerUpdateMarketRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Id *string `json:"id,omitempty"`
	Payload *ChannelmanagerUpdateMarketRequestPayload `json:"payload,omitempty"`
	PayloadMask []string `json:"payloadMask,omitempty"`
}

// NewChannelmanagerUpdateMarketRequest instantiates a new ChannelmanagerUpdateMarketRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelmanagerUpdateMarketRequest() *ChannelmanagerUpdateMarketRequest {
	this := ChannelmanagerUpdateMarketRequest{}
	return &this
}

// NewChannelmanagerUpdateMarketRequestWithDefaults instantiates a new ChannelmanagerUpdateMarketRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelmanagerUpdateMarketRequestWithDefaults() *ChannelmanagerUpdateMarketRequest {
	this := ChannelmanagerUpdateMarketRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ChannelmanagerUpdateMarketRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerUpdateMarketRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ChannelmanagerUpdateMarketRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ChannelmanagerUpdateMarketRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChannelmanagerUpdateMarketRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerUpdateMarketRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChannelmanagerUpdateMarketRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChannelmanagerUpdateMarketRequest) SetId(v string) {
	o.Id = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *ChannelmanagerUpdateMarketRequest) GetPayload() ChannelmanagerUpdateMarketRequestPayload {
	if o == nil || IsNil(o.Payload) {
		var ret ChannelmanagerUpdateMarketRequestPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerUpdateMarketRequest) GetPayloadOk() (*ChannelmanagerUpdateMarketRequestPayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *ChannelmanagerUpdateMarketRequest) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given ChannelmanagerUpdateMarketRequestPayload and assigns it to the Payload field.
func (o *ChannelmanagerUpdateMarketRequest) SetPayload(v ChannelmanagerUpdateMarketRequestPayload) {
	o.Payload = &v
}

// GetPayloadMask returns the PayloadMask field value if set, zero value otherwise.
func (o *ChannelmanagerUpdateMarketRequest) GetPayloadMask() []string {
	if o == nil || IsNil(o.PayloadMask) {
		var ret []string
		return ret
	}
	return o.PayloadMask
}

// GetPayloadMaskOk returns a tuple with the PayloadMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerUpdateMarketRequest) GetPayloadMaskOk() ([]string, bool) {
	if o == nil || IsNil(o.PayloadMask) {
		return nil, false
	}
	return o.PayloadMask, true
}

// HasPayloadMask returns a boolean if a field has been set.
func (o *ChannelmanagerUpdateMarketRequest) HasPayloadMask() bool {
	if o != nil && !IsNil(o.PayloadMask) {
		return true
	}

	return false
}

// SetPayloadMask gets a reference to the given []string and assigns it to the PayloadMask field.
func (o *ChannelmanagerUpdateMarketRequest) SetPayloadMask(v []string) {
	o.PayloadMask = v
}

func (o ChannelmanagerUpdateMarketRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelmanagerUpdateMarketRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.PayloadMask) {
		toSerialize["payloadMask"] = o.PayloadMask
	}
	return toSerialize, nil
}

type NullableChannelmanagerUpdateMarketRequest struct {
	value *ChannelmanagerUpdateMarketRequest
	isSet bool
}

func (v NullableChannelmanagerUpdateMarketRequest) Get() *ChannelmanagerUpdateMarketRequest {
	return v.value
}

func (v *NullableChannelmanagerUpdateMarketRequest) Set(val *ChannelmanagerUpdateMarketRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelmanagerUpdateMarketRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelmanagerUpdateMarketRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelmanagerUpdateMarketRequest(val *ChannelmanagerUpdateMarketRequest) *NullableChannelmanagerUpdateMarketRequest {
	return &NullableChannelmanagerUpdateMarketRequest{value: val, isSet: true}
}

func (v NullableChannelmanagerUpdateMarketRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelmanagerUpdateMarketRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

