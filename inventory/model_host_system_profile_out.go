/*
 * Insights Host Inventory REST Interface
 *
 * REST interface for the Insights Platform Host Inventory application.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package inventory

import (
	"encoding/json"
)

// HostSystemProfileOut Individual host record that contains only the host id and system profile
type HostSystemProfileOut struct {
	Id *string `json:"id,omitempty"`
	SystemProfile *SystemProfile `json:"system_profile,omitempty"`
}

// NewHostSystemProfileOut instantiates a new HostSystemProfileOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostSystemProfileOut() *HostSystemProfileOut {
	this := HostSystemProfileOut{}
	return &this
}

// NewHostSystemProfileOutWithDefaults instantiates a new HostSystemProfileOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostSystemProfileOutWithDefaults() *HostSystemProfileOut {
	this := HostSystemProfileOut{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HostSystemProfileOut) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostSystemProfileOut) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HostSystemProfileOut) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HostSystemProfileOut) SetId(v string) {
	o.Id = &v
}

// GetSystemProfile returns the SystemProfile field value if set, zero value otherwise.
func (o *HostSystemProfileOut) GetSystemProfile() SystemProfile {
	if o == nil || o.SystemProfile == nil {
		var ret SystemProfile
		return ret
	}
	return *o.SystemProfile
}

// GetSystemProfileOk returns a tuple with the SystemProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostSystemProfileOut) GetSystemProfileOk() (*SystemProfile, bool) {
	if o == nil || o.SystemProfile == nil {
		return nil, false
	}
	return o.SystemProfile, true
}

// HasSystemProfile returns a boolean if a field has been set.
func (o *HostSystemProfileOut) HasSystemProfile() bool {
	if o != nil && o.SystemProfile != nil {
		return true
	}

	return false
}

// SetSystemProfile gets a reference to the given SystemProfile and assigns it to the SystemProfile field.
func (o *HostSystemProfileOut) SetSystemProfile(v SystemProfile) {
	o.SystemProfile = &v
}

func (o HostSystemProfileOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.SystemProfile != nil {
		toSerialize["system_profile"] = o.SystemProfile
	}
	return json.Marshal(toSerialize)
}

type NullableHostSystemProfileOut struct {
	value *HostSystemProfileOut
	isSet bool
}

func (v NullableHostSystemProfileOut) Get() *HostSystemProfileOut {
	return v.value
}

func (v *NullableHostSystemProfileOut) Set(val *HostSystemProfileOut) {
	v.value = val
	v.isSet = true
}

func (v NullableHostSystemProfileOut) IsSet() bool {
	return v.isSet
}

func (v *NullableHostSystemProfileOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostSystemProfileOut(val *HostSystemProfileOut) *NullableHostSystemProfileOut {
	return &NullableHostSystemProfileOut{value: val, isSet: true}
}

func (v NullableHostSystemProfileOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostSystemProfileOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


