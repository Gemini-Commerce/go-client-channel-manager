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

// checks if the ChannelmanagerListChannelsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelmanagerListChannelsRequest{}

// ChannelmanagerListChannelsRequest struct for ChannelmanagerListChannelsRequest
type ChannelmanagerListChannelsRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	// The number of items to return per page. If not specified, it will returns all items.
	PageSize *int64 `json:"pageSize,omitempty"`
	PageToken *string `json:"pageToken,omitempty"`
}

// NewChannelmanagerListChannelsRequest instantiates a new ChannelmanagerListChannelsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelmanagerListChannelsRequest() *ChannelmanagerListChannelsRequest {
	this := ChannelmanagerListChannelsRequest{}
	return &this
}

// NewChannelmanagerListChannelsRequestWithDefaults instantiates a new ChannelmanagerListChannelsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelmanagerListChannelsRequestWithDefaults() *ChannelmanagerListChannelsRequest {
	this := ChannelmanagerListChannelsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ChannelmanagerListChannelsRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerListChannelsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ChannelmanagerListChannelsRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ChannelmanagerListChannelsRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ChannelmanagerListChannelsRequest) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerListChannelsRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ChannelmanagerListChannelsRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *ChannelmanagerListChannelsRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *ChannelmanagerListChannelsRequest) GetPageToken() string {
	if o == nil || IsNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerListChannelsRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PageToken) {
		return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *ChannelmanagerListChannelsRequest) HasPageToken() bool {
	if o != nil && !IsNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *ChannelmanagerListChannelsRequest) SetPageToken(v string) {
	o.PageToken = &v
}

func (o ChannelmanagerListChannelsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelmanagerListChannelsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.PageToken) {
		toSerialize["pageToken"] = o.PageToken
	}
	return toSerialize, nil
}

type NullableChannelmanagerListChannelsRequest struct {
	value *ChannelmanagerListChannelsRequest
	isSet bool
}

func (v NullableChannelmanagerListChannelsRequest) Get() *ChannelmanagerListChannelsRequest {
	return v.value
}

func (v *NullableChannelmanagerListChannelsRequest) Set(val *ChannelmanagerListChannelsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelmanagerListChannelsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelmanagerListChannelsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelmanagerListChannelsRequest(val *ChannelmanagerListChannelsRequest) *NullableChannelmanagerListChannelsRequest {
	return &NullableChannelmanagerListChannelsRequest{value: val, isSet: true}
}

func (v NullableChannelmanagerListChannelsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelmanagerListChannelsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


