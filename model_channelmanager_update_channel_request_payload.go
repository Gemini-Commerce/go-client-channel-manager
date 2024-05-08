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

// checks if the ChannelmanagerUpdateChannelRequestPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelmanagerUpdateChannelRequestPayload{}

// ChannelmanagerUpdateChannelRequestPayload struct for ChannelmanagerUpdateChannelRequestPayload
type ChannelmanagerUpdateChannelRequestPayload struct {
	Image *string `json:"image,omitempty"`
	Name *string `json:"name,omitempty"`
	DefaultLanguage *ChannelmanagerLanguageCode `json:"defaultLanguage,omitempty"`
	Settings *ChannelmanagerChannelSettings `json:"settings,omitempty"`
	Website *ChannelmanagerChannelTypeWebsite `json:"website,omitempty"`
	Status *ChannelmanagerChannelStatus `json:"status,omitempty"`
}

// NewChannelmanagerUpdateChannelRequestPayload instantiates a new ChannelmanagerUpdateChannelRequestPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelmanagerUpdateChannelRequestPayload() *ChannelmanagerUpdateChannelRequestPayload {
	this := ChannelmanagerUpdateChannelRequestPayload{}
	var defaultLanguage ChannelmanagerLanguageCode = CHANNELMANAGERLANGUAGECODE_UNKNOWN
	this.DefaultLanguage = &defaultLanguage
	var status ChannelmanagerChannelStatus = CHANNELMANAGERCHANNELSTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// NewChannelmanagerUpdateChannelRequestPayloadWithDefaults instantiates a new ChannelmanagerUpdateChannelRequestPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelmanagerUpdateChannelRequestPayloadWithDefaults() *ChannelmanagerUpdateChannelRequestPayload {
	this := ChannelmanagerUpdateChannelRequestPayload{}
	var defaultLanguage ChannelmanagerLanguageCode = CHANNELMANAGERLANGUAGECODE_UNKNOWN
	this.DefaultLanguage = &defaultLanguage
	var status ChannelmanagerChannelStatus = CHANNELMANAGERCHANNELSTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *ChannelmanagerUpdateChannelRequestPayload) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerUpdateChannelRequestPayload) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *ChannelmanagerUpdateChannelRequestPayload) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *ChannelmanagerUpdateChannelRequestPayload) SetImage(v string) {
	o.Image = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ChannelmanagerUpdateChannelRequestPayload) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerUpdateChannelRequestPayload) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ChannelmanagerUpdateChannelRequestPayload) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ChannelmanagerUpdateChannelRequestPayload) SetName(v string) {
	o.Name = &v
}

// GetDefaultLanguage returns the DefaultLanguage field value if set, zero value otherwise.
func (o *ChannelmanagerUpdateChannelRequestPayload) GetDefaultLanguage() ChannelmanagerLanguageCode {
	if o == nil || IsNil(o.DefaultLanguage) {
		var ret ChannelmanagerLanguageCode
		return ret
	}
	return *o.DefaultLanguage
}

// GetDefaultLanguageOk returns a tuple with the DefaultLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerUpdateChannelRequestPayload) GetDefaultLanguageOk() (*ChannelmanagerLanguageCode, bool) {
	if o == nil || IsNil(o.DefaultLanguage) {
		return nil, false
	}
	return o.DefaultLanguage, true
}

// HasDefaultLanguage returns a boolean if a field has been set.
func (o *ChannelmanagerUpdateChannelRequestPayload) HasDefaultLanguage() bool {
	if o != nil && !IsNil(o.DefaultLanguage) {
		return true
	}

	return false
}

// SetDefaultLanguage gets a reference to the given ChannelmanagerLanguageCode and assigns it to the DefaultLanguage field.
func (o *ChannelmanagerUpdateChannelRequestPayload) SetDefaultLanguage(v ChannelmanagerLanguageCode) {
	o.DefaultLanguage = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *ChannelmanagerUpdateChannelRequestPayload) GetSettings() ChannelmanagerChannelSettings {
	if o == nil || IsNil(o.Settings) {
		var ret ChannelmanagerChannelSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerUpdateChannelRequestPayload) GetSettingsOk() (*ChannelmanagerChannelSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *ChannelmanagerUpdateChannelRequestPayload) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given ChannelmanagerChannelSettings and assigns it to the Settings field.
func (o *ChannelmanagerUpdateChannelRequestPayload) SetSettings(v ChannelmanagerChannelSettings) {
	o.Settings = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *ChannelmanagerUpdateChannelRequestPayload) GetWebsite() ChannelmanagerChannelTypeWebsite {
	if o == nil || IsNil(o.Website) {
		var ret ChannelmanagerChannelTypeWebsite
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerUpdateChannelRequestPayload) GetWebsiteOk() (*ChannelmanagerChannelTypeWebsite, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *ChannelmanagerUpdateChannelRequestPayload) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given ChannelmanagerChannelTypeWebsite and assigns it to the Website field.
func (o *ChannelmanagerUpdateChannelRequestPayload) SetWebsite(v ChannelmanagerChannelTypeWebsite) {
	o.Website = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ChannelmanagerUpdateChannelRequestPayload) GetStatus() ChannelmanagerChannelStatus {
	if o == nil || IsNil(o.Status) {
		var ret ChannelmanagerChannelStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerUpdateChannelRequestPayload) GetStatusOk() (*ChannelmanagerChannelStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ChannelmanagerUpdateChannelRequestPayload) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ChannelmanagerChannelStatus and assigns it to the Status field.
func (o *ChannelmanagerUpdateChannelRequestPayload) SetStatus(v ChannelmanagerChannelStatus) {
	o.Status = &v
}

func (o ChannelmanagerUpdateChannelRequestPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelmanagerUpdateChannelRequestPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DefaultLanguage) {
		toSerialize["defaultLanguage"] = o.DefaultLanguage
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.Website) {
		toSerialize["website"] = o.Website
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableChannelmanagerUpdateChannelRequestPayload struct {
	value *ChannelmanagerUpdateChannelRequestPayload
	isSet bool
}

func (v NullableChannelmanagerUpdateChannelRequestPayload) Get() *ChannelmanagerUpdateChannelRequestPayload {
	return v.value
}

func (v *NullableChannelmanagerUpdateChannelRequestPayload) Set(val *ChannelmanagerUpdateChannelRequestPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelmanagerUpdateChannelRequestPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelmanagerUpdateChannelRequestPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelmanagerUpdateChannelRequestPayload(val *ChannelmanagerUpdateChannelRequestPayload) *NullableChannelmanagerUpdateChannelRequestPayload {
	return &NullableChannelmanagerUpdateChannelRequestPayload{value: val, isSet: true}
}

func (v NullableChannelmanagerUpdateChannelRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelmanagerUpdateChannelRequestPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

