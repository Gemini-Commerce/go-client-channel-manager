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

// checks if the ChannelmanagerCreateChannelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelmanagerCreateChannelRequest{}

// ChannelmanagerCreateChannelRequest struct for ChannelmanagerCreateChannelRequest
type ChannelmanagerCreateChannelRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Code *string `json:"code,omitempty"`
	Image *string `json:"image,omitempty"`
	Name *string `json:"name,omitempty"`
	DefaultLanguage *ChannelmanagerLanguageCode `json:"defaultLanguage,omitempty"`
	Settings *ChannelmanagerChannelSettings `json:"settings,omitempty"`
	Website *ChannelmanagerChannelTypeWebsite `json:"website,omitempty"`
	Status *ChannelmanagerChannelStatus `json:"status,omitempty"`
}

// NewChannelmanagerCreateChannelRequest instantiates a new ChannelmanagerCreateChannelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelmanagerCreateChannelRequest() *ChannelmanagerCreateChannelRequest {
	this := ChannelmanagerCreateChannelRequest{}
	var defaultLanguage ChannelmanagerLanguageCode = CHANNELMANAGERLANGUAGECODE_UNKNOWN
	this.DefaultLanguage = &defaultLanguage
	var status ChannelmanagerChannelStatus = CHANNELMANAGERCHANNELSTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// NewChannelmanagerCreateChannelRequestWithDefaults instantiates a new ChannelmanagerCreateChannelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelmanagerCreateChannelRequestWithDefaults() *ChannelmanagerCreateChannelRequest {
	this := ChannelmanagerCreateChannelRequest{}
	var defaultLanguage ChannelmanagerLanguageCode = CHANNELMANAGERLANGUAGECODE_UNKNOWN
	this.DefaultLanguage = &defaultLanguage
	var status ChannelmanagerChannelStatus = CHANNELMANAGERCHANNELSTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ChannelmanagerCreateChannelRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerCreateChannelRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ChannelmanagerCreateChannelRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ChannelmanagerCreateChannelRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ChannelmanagerCreateChannelRequest) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerCreateChannelRequest) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ChannelmanagerCreateChannelRequest) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ChannelmanagerCreateChannelRequest) SetCode(v string) {
	o.Code = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *ChannelmanagerCreateChannelRequest) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerCreateChannelRequest) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *ChannelmanagerCreateChannelRequest) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *ChannelmanagerCreateChannelRequest) SetImage(v string) {
	o.Image = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ChannelmanagerCreateChannelRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerCreateChannelRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ChannelmanagerCreateChannelRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ChannelmanagerCreateChannelRequest) SetName(v string) {
	o.Name = &v
}

// GetDefaultLanguage returns the DefaultLanguage field value if set, zero value otherwise.
func (o *ChannelmanagerCreateChannelRequest) GetDefaultLanguage() ChannelmanagerLanguageCode {
	if o == nil || IsNil(o.DefaultLanguage) {
		var ret ChannelmanagerLanguageCode
		return ret
	}
	return *o.DefaultLanguage
}

// GetDefaultLanguageOk returns a tuple with the DefaultLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerCreateChannelRequest) GetDefaultLanguageOk() (*ChannelmanagerLanguageCode, bool) {
	if o == nil || IsNil(o.DefaultLanguage) {
		return nil, false
	}
	return o.DefaultLanguage, true
}

// HasDefaultLanguage returns a boolean if a field has been set.
func (o *ChannelmanagerCreateChannelRequest) HasDefaultLanguage() bool {
	if o != nil && !IsNil(o.DefaultLanguage) {
		return true
	}

	return false
}

// SetDefaultLanguage gets a reference to the given ChannelmanagerLanguageCode and assigns it to the DefaultLanguage field.
func (o *ChannelmanagerCreateChannelRequest) SetDefaultLanguage(v ChannelmanagerLanguageCode) {
	o.DefaultLanguage = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *ChannelmanagerCreateChannelRequest) GetSettings() ChannelmanagerChannelSettings {
	if o == nil || IsNil(o.Settings) {
		var ret ChannelmanagerChannelSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerCreateChannelRequest) GetSettingsOk() (*ChannelmanagerChannelSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *ChannelmanagerCreateChannelRequest) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given ChannelmanagerChannelSettings and assigns it to the Settings field.
func (o *ChannelmanagerCreateChannelRequest) SetSettings(v ChannelmanagerChannelSettings) {
	o.Settings = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *ChannelmanagerCreateChannelRequest) GetWebsite() ChannelmanagerChannelTypeWebsite {
	if o == nil || IsNil(o.Website) {
		var ret ChannelmanagerChannelTypeWebsite
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerCreateChannelRequest) GetWebsiteOk() (*ChannelmanagerChannelTypeWebsite, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *ChannelmanagerCreateChannelRequest) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given ChannelmanagerChannelTypeWebsite and assigns it to the Website field.
func (o *ChannelmanagerCreateChannelRequest) SetWebsite(v ChannelmanagerChannelTypeWebsite) {
	o.Website = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ChannelmanagerCreateChannelRequest) GetStatus() ChannelmanagerChannelStatus {
	if o == nil || IsNil(o.Status) {
		var ret ChannelmanagerChannelStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerCreateChannelRequest) GetStatusOk() (*ChannelmanagerChannelStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ChannelmanagerCreateChannelRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ChannelmanagerChannelStatus and assigns it to the Status field.
func (o *ChannelmanagerCreateChannelRequest) SetStatus(v ChannelmanagerChannelStatus) {
	o.Status = &v
}

func (o ChannelmanagerCreateChannelRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelmanagerCreateChannelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
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

type NullableChannelmanagerCreateChannelRequest struct {
	value *ChannelmanagerCreateChannelRequest
	isSet bool
}

func (v NullableChannelmanagerCreateChannelRequest) Get() *ChannelmanagerCreateChannelRequest {
	return v.value
}

func (v *NullableChannelmanagerCreateChannelRequest) Set(val *ChannelmanagerCreateChannelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelmanagerCreateChannelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelmanagerCreateChannelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelmanagerCreateChannelRequest(val *ChannelmanagerCreateChannelRequest) *NullableChannelmanagerCreateChannelRequest {
	return &NullableChannelmanagerCreateChannelRequest{value: val, isSet: true}
}

func (v NullableChannelmanagerCreateChannelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelmanagerCreateChannelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


