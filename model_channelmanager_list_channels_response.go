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

// checks if the ChannelmanagerListChannelsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelmanagerListChannelsResponse{}

// ChannelmanagerListChannelsResponse struct for ChannelmanagerListChannelsResponse
type ChannelmanagerListChannelsResponse struct {
	Channels []ChannelmanagerChannelResponse `json:"channels,omitempty"`
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

// NewChannelmanagerListChannelsResponse instantiates a new ChannelmanagerListChannelsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelmanagerListChannelsResponse() *ChannelmanagerListChannelsResponse {
	this := ChannelmanagerListChannelsResponse{}
	return &this
}

// NewChannelmanagerListChannelsResponseWithDefaults instantiates a new ChannelmanagerListChannelsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelmanagerListChannelsResponseWithDefaults() *ChannelmanagerListChannelsResponse {
	this := ChannelmanagerListChannelsResponse{}
	return &this
}

// GetChannels returns the Channels field value if set, zero value otherwise.
func (o *ChannelmanagerListChannelsResponse) GetChannels() []ChannelmanagerChannelResponse {
	if o == nil || IsNil(o.Channels) {
		var ret []ChannelmanagerChannelResponse
		return ret
	}
	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerListChannelsResponse) GetChannelsOk() ([]ChannelmanagerChannelResponse, bool) {
	if o == nil || IsNil(o.Channels) {
		return nil, false
	}
	return o.Channels, true
}

// HasChannels returns a boolean if a field has been set.
func (o *ChannelmanagerListChannelsResponse) HasChannels() bool {
	if o != nil && !IsNil(o.Channels) {
		return true
	}

	return false
}

// SetChannels gets a reference to the given []ChannelmanagerChannelResponse and assigns it to the Channels field.
func (o *ChannelmanagerListChannelsResponse) SetChannels(v []ChannelmanagerChannelResponse) {
	o.Channels = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ChannelmanagerListChannelsResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerListChannelsResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ChannelmanagerListChannelsResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ChannelmanagerListChannelsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o ChannelmanagerListChannelsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelmanagerListChannelsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Channels) {
		toSerialize["channels"] = o.Channels
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	return toSerialize, nil
}

type NullableChannelmanagerListChannelsResponse struct {
	value *ChannelmanagerListChannelsResponse
	isSet bool
}

func (v NullableChannelmanagerListChannelsResponse) Get() *ChannelmanagerListChannelsResponse {
	return v.value
}

func (v *NullableChannelmanagerListChannelsResponse) Set(val *ChannelmanagerListChannelsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelmanagerListChannelsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelmanagerListChannelsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelmanagerListChannelsResponse(val *ChannelmanagerListChannelsResponse) *NullableChannelmanagerListChannelsResponse {
	return &NullableChannelmanagerListChannelsResponse{value: val, isSet: true}
}

func (v NullableChannelmanagerListChannelsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelmanagerListChannelsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

