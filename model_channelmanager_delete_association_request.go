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

// checks if the ChannelmanagerDeleteAssociationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelmanagerDeleteAssociationRequest{}

// ChannelmanagerDeleteAssociationRequest struct for ChannelmanagerDeleteAssociationRequest
type ChannelmanagerDeleteAssociationRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Id *string `json:"id,omitempty"`
}

// NewChannelmanagerDeleteAssociationRequest instantiates a new ChannelmanagerDeleteAssociationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelmanagerDeleteAssociationRequest() *ChannelmanagerDeleteAssociationRequest {
	this := ChannelmanagerDeleteAssociationRequest{}
	return &this
}

// NewChannelmanagerDeleteAssociationRequestWithDefaults instantiates a new ChannelmanagerDeleteAssociationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelmanagerDeleteAssociationRequestWithDefaults() *ChannelmanagerDeleteAssociationRequest {
	this := ChannelmanagerDeleteAssociationRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ChannelmanagerDeleteAssociationRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerDeleteAssociationRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ChannelmanagerDeleteAssociationRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ChannelmanagerDeleteAssociationRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChannelmanagerDeleteAssociationRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerDeleteAssociationRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChannelmanagerDeleteAssociationRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChannelmanagerDeleteAssociationRequest) SetId(v string) {
	o.Id = &v
}

func (o ChannelmanagerDeleteAssociationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelmanagerDeleteAssociationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableChannelmanagerDeleteAssociationRequest struct {
	value *ChannelmanagerDeleteAssociationRequest
	isSet bool
}

func (v NullableChannelmanagerDeleteAssociationRequest) Get() *ChannelmanagerDeleteAssociationRequest {
	return v.value
}

func (v *NullableChannelmanagerDeleteAssociationRequest) Set(val *ChannelmanagerDeleteAssociationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelmanagerDeleteAssociationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelmanagerDeleteAssociationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelmanagerDeleteAssociationRequest(val *ChannelmanagerDeleteAssociationRequest) *NullableChannelmanagerDeleteAssociationRequest {
	return &NullableChannelmanagerDeleteAssociationRequest{value: val, isSet: true}
}

func (v NullableChannelmanagerDeleteAssociationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelmanagerDeleteAssociationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


