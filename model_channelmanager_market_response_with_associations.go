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

// checks if the ChannelmanagerMarketResponseWithAssociations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelmanagerMarketResponseWithAssociations{}

// ChannelmanagerMarketResponseWithAssociations struct for ChannelmanagerMarketResponseWithAssociations
type ChannelmanagerMarketResponseWithAssociations struct {
	Id *string `json:"id,omitempty"`
	Grn *string `json:"grn,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Countries []ChannelmanagerCountryCode `json:"countries,omitempty"`
	Associations []ChannelmanagerMarketResponseWithAssociationsAssociation `json:"associations,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewChannelmanagerMarketResponseWithAssociations instantiates a new ChannelmanagerMarketResponseWithAssociations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelmanagerMarketResponseWithAssociations() *ChannelmanagerMarketResponseWithAssociations {
	this := ChannelmanagerMarketResponseWithAssociations{}
	return &this
}

// NewChannelmanagerMarketResponseWithAssociationsWithDefaults instantiates a new ChannelmanagerMarketResponseWithAssociations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelmanagerMarketResponseWithAssociationsWithDefaults() *ChannelmanagerMarketResponseWithAssociations {
	this := ChannelmanagerMarketResponseWithAssociations{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChannelmanagerMarketResponseWithAssociations) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerMarketResponseWithAssociations) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChannelmanagerMarketResponseWithAssociations) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChannelmanagerMarketResponseWithAssociations) SetId(v string) {
	o.Id = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *ChannelmanagerMarketResponseWithAssociations) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerMarketResponseWithAssociations) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *ChannelmanagerMarketResponseWithAssociations) HasGrn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *ChannelmanagerMarketResponseWithAssociations) SetGrn(v string) {
	o.Grn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ChannelmanagerMarketResponseWithAssociations) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerMarketResponseWithAssociations) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ChannelmanagerMarketResponseWithAssociations) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ChannelmanagerMarketResponseWithAssociations) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ChannelmanagerMarketResponseWithAssociations) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerMarketResponseWithAssociations) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ChannelmanagerMarketResponseWithAssociations) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ChannelmanagerMarketResponseWithAssociations) SetDescription(v string) {
	o.Description = &v
}

// GetCountries returns the Countries field value if set, zero value otherwise.
func (o *ChannelmanagerMarketResponseWithAssociations) GetCountries() []ChannelmanagerCountryCode {
	if o == nil || IsNil(o.Countries) {
		var ret []ChannelmanagerCountryCode
		return ret
	}
	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerMarketResponseWithAssociations) GetCountriesOk() ([]ChannelmanagerCountryCode, bool) {
	if o == nil || IsNil(o.Countries) {
		return nil, false
	}
	return o.Countries, true
}

// HasCountries returns a boolean if a field has been set.
func (o *ChannelmanagerMarketResponseWithAssociations) HasCountries() bool {
	if o != nil && !IsNil(o.Countries) {
		return true
	}

	return false
}

// SetCountries gets a reference to the given []ChannelmanagerCountryCode and assigns it to the Countries field.
func (o *ChannelmanagerMarketResponseWithAssociations) SetCountries(v []ChannelmanagerCountryCode) {
	o.Countries = v
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *ChannelmanagerMarketResponseWithAssociations) GetAssociations() []ChannelmanagerMarketResponseWithAssociationsAssociation {
	if o == nil || IsNil(o.Associations) {
		var ret []ChannelmanagerMarketResponseWithAssociationsAssociation
		return ret
	}
	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerMarketResponseWithAssociations) GetAssociationsOk() ([]ChannelmanagerMarketResponseWithAssociationsAssociation, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *ChannelmanagerMarketResponseWithAssociations) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given []ChannelmanagerMarketResponseWithAssociationsAssociation and assigns it to the Associations field.
func (o *ChannelmanagerMarketResponseWithAssociations) SetAssociations(v []ChannelmanagerMarketResponseWithAssociationsAssociation) {
	o.Associations = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ChannelmanagerMarketResponseWithAssociations) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerMarketResponseWithAssociations) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ChannelmanagerMarketResponseWithAssociations) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ChannelmanagerMarketResponseWithAssociations) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ChannelmanagerMarketResponseWithAssociations) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerMarketResponseWithAssociations) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ChannelmanagerMarketResponseWithAssociations) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ChannelmanagerMarketResponseWithAssociations) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ChannelmanagerMarketResponseWithAssociations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelmanagerMarketResponseWithAssociations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Countries) {
		toSerialize["countries"] = o.Countries
	}
	if !IsNil(o.Associations) {
		toSerialize["associations"] = o.Associations
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableChannelmanagerMarketResponseWithAssociations struct {
	value *ChannelmanagerMarketResponseWithAssociations
	isSet bool
}

func (v NullableChannelmanagerMarketResponseWithAssociations) Get() *ChannelmanagerMarketResponseWithAssociations {
	return v.value
}

func (v *NullableChannelmanagerMarketResponseWithAssociations) Set(val *ChannelmanagerMarketResponseWithAssociations) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelmanagerMarketResponseWithAssociations) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelmanagerMarketResponseWithAssociations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelmanagerMarketResponseWithAssociations(val *ChannelmanagerMarketResponseWithAssociations) *NullableChannelmanagerMarketResponseWithAssociations {
	return &NullableChannelmanagerMarketResponseWithAssociations{value: val, isSet: true}
}

func (v NullableChannelmanagerMarketResponseWithAssociations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelmanagerMarketResponseWithAssociations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

