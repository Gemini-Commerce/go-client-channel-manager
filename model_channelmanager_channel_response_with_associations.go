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
	"time"
)

// checks if the ChannelmanagerChannelResponseWithAssociations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelmanagerChannelResponseWithAssociations{}

// ChannelmanagerChannelResponseWithAssociations struct for ChannelmanagerChannelResponseWithAssociations
type ChannelmanagerChannelResponseWithAssociations struct {
	Id *string `json:"id,omitempty"`
	Grn *string `json:"grn,omitempty"`
	Code *string `json:"code,omitempty"`
	Image *string `json:"image,omitempty"`
	Name *string `json:"name,omitempty"`
	DefaultLanguage *ChannelmanagerLanguageCode `json:"defaultLanguage,omitempty"`
	Settings *ChannelmanagerChannelSettings `json:"settings,omitempty"`
	Website *ChannelmanagerChannelTypeWebsite `json:"website,omitempty"`
	Associations []ChannelmanagerChannelResponseWithAssociationsAssociation `json:"associations,omitempty"`
	Status *ChannelmanagerChannelStatus `json:"status,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewChannelmanagerChannelResponseWithAssociations instantiates a new ChannelmanagerChannelResponseWithAssociations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelmanagerChannelResponseWithAssociations() *ChannelmanagerChannelResponseWithAssociations {
	this := ChannelmanagerChannelResponseWithAssociations{}
	var defaultLanguage ChannelmanagerLanguageCode = CHANNELMANAGERLANGUAGECODE_UNKNOWN
	this.DefaultLanguage = &defaultLanguage
	var status ChannelmanagerChannelStatus = CHANNELMANAGERCHANNELSTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// NewChannelmanagerChannelResponseWithAssociationsWithDefaults instantiates a new ChannelmanagerChannelResponseWithAssociations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelmanagerChannelResponseWithAssociationsWithDefaults() *ChannelmanagerChannelResponseWithAssociations {
	this := ChannelmanagerChannelResponseWithAssociations{}
	var defaultLanguage ChannelmanagerLanguageCode = CHANNELMANAGERLANGUAGECODE_UNKNOWN
	this.DefaultLanguage = &defaultLanguage
	var status ChannelmanagerChannelStatus = CHANNELMANAGERCHANNELSTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChannelmanagerChannelResponseWithAssociations) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerChannelResponseWithAssociations) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChannelmanagerChannelResponseWithAssociations) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChannelmanagerChannelResponseWithAssociations) SetId(v string) {
	o.Id = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *ChannelmanagerChannelResponseWithAssociations) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerChannelResponseWithAssociations) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *ChannelmanagerChannelResponseWithAssociations) HasGrn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *ChannelmanagerChannelResponseWithAssociations) SetGrn(v string) {
	o.Grn = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ChannelmanagerChannelResponseWithAssociations) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerChannelResponseWithAssociations) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ChannelmanagerChannelResponseWithAssociations) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ChannelmanagerChannelResponseWithAssociations) SetCode(v string) {
	o.Code = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *ChannelmanagerChannelResponseWithAssociations) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerChannelResponseWithAssociations) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *ChannelmanagerChannelResponseWithAssociations) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *ChannelmanagerChannelResponseWithAssociations) SetImage(v string) {
	o.Image = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ChannelmanagerChannelResponseWithAssociations) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerChannelResponseWithAssociations) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ChannelmanagerChannelResponseWithAssociations) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ChannelmanagerChannelResponseWithAssociations) SetName(v string) {
	o.Name = &v
}

// GetDefaultLanguage returns the DefaultLanguage field value if set, zero value otherwise.
func (o *ChannelmanagerChannelResponseWithAssociations) GetDefaultLanguage() ChannelmanagerLanguageCode {
	if o == nil || IsNil(o.DefaultLanguage) {
		var ret ChannelmanagerLanguageCode
		return ret
	}
	return *o.DefaultLanguage
}

// GetDefaultLanguageOk returns a tuple with the DefaultLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerChannelResponseWithAssociations) GetDefaultLanguageOk() (*ChannelmanagerLanguageCode, bool) {
	if o == nil || IsNil(o.DefaultLanguage) {
		return nil, false
	}
	return o.DefaultLanguage, true
}

// HasDefaultLanguage returns a boolean if a field has been set.
func (o *ChannelmanagerChannelResponseWithAssociations) HasDefaultLanguage() bool {
	if o != nil && !IsNil(o.DefaultLanguage) {
		return true
	}

	return false
}

// SetDefaultLanguage gets a reference to the given ChannelmanagerLanguageCode and assigns it to the DefaultLanguage field.
func (o *ChannelmanagerChannelResponseWithAssociations) SetDefaultLanguage(v ChannelmanagerLanguageCode) {
	o.DefaultLanguage = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *ChannelmanagerChannelResponseWithAssociations) GetSettings() ChannelmanagerChannelSettings {
	if o == nil || IsNil(o.Settings) {
		var ret ChannelmanagerChannelSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerChannelResponseWithAssociations) GetSettingsOk() (*ChannelmanagerChannelSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *ChannelmanagerChannelResponseWithAssociations) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given ChannelmanagerChannelSettings and assigns it to the Settings field.
func (o *ChannelmanagerChannelResponseWithAssociations) SetSettings(v ChannelmanagerChannelSettings) {
	o.Settings = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *ChannelmanagerChannelResponseWithAssociations) GetWebsite() ChannelmanagerChannelTypeWebsite {
	if o == nil || IsNil(o.Website) {
		var ret ChannelmanagerChannelTypeWebsite
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerChannelResponseWithAssociations) GetWebsiteOk() (*ChannelmanagerChannelTypeWebsite, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *ChannelmanagerChannelResponseWithAssociations) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given ChannelmanagerChannelTypeWebsite and assigns it to the Website field.
func (o *ChannelmanagerChannelResponseWithAssociations) SetWebsite(v ChannelmanagerChannelTypeWebsite) {
	o.Website = &v
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *ChannelmanagerChannelResponseWithAssociations) GetAssociations() []ChannelmanagerChannelResponseWithAssociationsAssociation {
	if o == nil || IsNil(o.Associations) {
		var ret []ChannelmanagerChannelResponseWithAssociationsAssociation
		return ret
	}
	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerChannelResponseWithAssociations) GetAssociationsOk() ([]ChannelmanagerChannelResponseWithAssociationsAssociation, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *ChannelmanagerChannelResponseWithAssociations) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given []ChannelmanagerChannelResponseWithAssociationsAssociation and assigns it to the Associations field.
func (o *ChannelmanagerChannelResponseWithAssociations) SetAssociations(v []ChannelmanagerChannelResponseWithAssociationsAssociation) {
	o.Associations = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ChannelmanagerChannelResponseWithAssociations) GetStatus() ChannelmanagerChannelStatus {
	if o == nil || IsNil(o.Status) {
		var ret ChannelmanagerChannelStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerChannelResponseWithAssociations) GetStatusOk() (*ChannelmanagerChannelStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ChannelmanagerChannelResponseWithAssociations) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ChannelmanagerChannelStatus and assigns it to the Status field.
func (o *ChannelmanagerChannelResponseWithAssociations) SetStatus(v ChannelmanagerChannelStatus) {
	o.Status = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ChannelmanagerChannelResponseWithAssociations) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerChannelResponseWithAssociations) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ChannelmanagerChannelResponseWithAssociations) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ChannelmanagerChannelResponseWithAssociations) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ChannelmanagerChannelResponseWithAssociations) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerChannelResponseWithAssociations) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ChannelmanagerChannelResponseWithAssociations) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ChannelmanagerChannelResponseWithAssociations) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ChannelmanagerChannelResponseWithAssociations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelmanagerChannelResponseWithAssociations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Grn) {
		toSerialize["grn"] = o.Grn
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
	if !IsNil(o.Associations) {
		toSerialize["associations"] = o.Associations
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableChannelmanagerChannelResponseWithAssociations struct {
	value *ChannelmanagerChannelResponseWithAssociations
	isSet bool
}

func (v NullableChannelmanagerChannelResponseWithAssociations) Get() *ChannelmanagerChannelResponseWithAssociations {
	return v.value
}

func (v *NullableChannelmanagerChannelResponseWithAssociations) Set(val *ChannelmanagerChannelResponseWithAssociations) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelmanagerChannelResponseWithAssociations) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelmanagerChannelResponseWithAssociations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelmanagerChannelResponseWithAssociations(val *ChannelmanagerChannelResponseWithAssociations) *NullableChannelmanagerChannelResponseWithAssociations {
	return &NullableChannelmanagerChannelResponseWithAssociations{value: val, isSet: true}
}

func (v NullableChannelmanagerChannelResponseWithAssociations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelmanagerChannelResponseWithAssociations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


