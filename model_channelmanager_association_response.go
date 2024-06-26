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

// checks if the ChannelmanagerAssociationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelmanagerAssociationResponse{}

// ChannelmanagerAssociationResponse struct for ChannelmanagerAssociationResponse
type ChannelmanagerAssociationResponse struct {
	Id *string `json:"id,omitempty"`
	Grn *string `json:"grn,omitempty"`
	AssociationEntities *ChannelmanagerAssociationResponseAssociation `json:"associationEntities,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewChannelmanagerAssociationResponse instantiates a new ChannelmanagerAssociationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelmanagerAssociationResponse() *ChannelmanagerAssociationResponse {
	this := ChannelmanagerAssociationResponse{}
	return &this
}

// NewChannelmanagerAssociationResponseWithDefaults instantiates a new ChannelmanagerAssociationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelmanagerAssociationResponseWithDefaults() *ChannelmanagerAssociationResponse {
	this := ChannelmanagerAssociationResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChannelmanagerAssociationResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerAssociationResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChannelmanagerAssociationResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChannelmanagerAssociationResponse) SetId(v string) {
	o.Id = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *ChannelmanagerAssociationResponse) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerAssociationResponse) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *ChannelmanagerAssociationResponse) HasGrn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *ChannelmanagerAssociationResponse) SetGrn(v string) {
	o.Grn = &v
}

// GetAssociationEntities returns the AssociationEntities field value if set, zero value otherwise.
func (o *ChannelmanagerAssociationResponse) GetAssociationEntities() ChannelmanagerAssociationResponseAssociation {
	if o == nil || IsNil(o.AssociationEntities) {
		var ret ChannelmanagerAssociationResponseAssociation
		return ret
	}
	return *o.AssociationEntities
}

// GetAssociationEntitiesOk returns a tuple with the AssociationEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerAssociationResponse) GetAssociationEntitiesOk() (*ChannelmanagerAssociationResponseAssociation, bool) {
	if o == nil || IsNil(o.AssociationEntities) {
		return nil, false
	}
	return o.AssociationEntities, true
}

// HasAssociationEntities returns a boolean if a field has been set.
func (o *ChannelmanagerAssociationResponse) HasAssociationEntities() bool {
	if o != nil && !IsNil(o.AssociationEntities) {
		return true
	}

	return false
}

// SetAssociationEntities gets a reference to the given ChannelmanagerAssociationResponseAssociation and assigns it to the AssociationEntities field.
func (o *ChannelmanagerAssociationResponse) SetAssociationEntities(v ChannelmanagerAssociationResponseAssociation) {
	o.AssociationEntities = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ChannelmanagerAssociationResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerAssociationResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ChannelmanagerAssociationResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ChannelmanagerAssociationResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ChannelmanagerAssociationResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelmanagerAssociationResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ChannelmanagerAssociationResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ChannelmanagerAssociationResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ChannelmanagerAssociationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelmanagerAssociationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !IsNil(o.AssociationEntities) {
		toSerialize["associationEntities"] = o.AssociationEntities
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableChannelmanagerAssociationResponse struct {
	value *ChannelmanagerAssociationResponse
	isSet bool
}

func (v NullableChannelmanagerAssociationResponse) Get() *ChannelmanagerAssociationResponse {
	return v.value
}

func (v *NullableChannelmanagerAssociationResponse) Set(val *ChannelmanagerAssociationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelmanagerAssociationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelmanagerAssociationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelmanagerAssociationResponse(val *ChannelmanagerAssociationResponse) *NullableChannelmanagerAssociationResponse {
	return &NullableChannelmanagerAssociationResponse{value: val, isSet: true}
}

func (v NullableChannelmanagerAssociationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelmanagerAssociationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


