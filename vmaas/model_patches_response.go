/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.32.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vmaas

import (
	"encoding/json"
)

// PatchesResponse struct for PatchesResponse
type PatchesResponse struct {
	ErrataList *[]string `json:"errata_list,omitempty"`
	LastChange *string `json:"last_change,omitempty"`
}

// NewPatchesResponse instantiates a new PatchesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchesResponse() *PatchesResponse {
	this := PatchesResponse{}
	return &this
}

// NewPatchesResponseWithDefaults instantiates a new PatchesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchesResponseWithDefaults() *PatchesResponse {
	this := PatchesResponse{}
	return &this
}

// GetErrataList returns the ErrataList field value if set, zero value otherwise.
func (o *PatchesResponse) GetErrataList() []string {
	if o == nil || o.ErrataList == nil {
		var ret []string
		return ret
	}
	return *o.ErrataList
}

// GetErrataListOk returns a tuple with the ErrataList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchesResponse) GetErrataListOk() (*[]string, bool) {
	if o == nil || o.ErrataList == nil {
		return nil, false
	}
	return o.ErrataList, true
}

// HasErrataList returns a boolean if a field has been set.
func (o *PatchesResponse) HasErrataList() bool {
	if o != nil && o.ErrataList != nil {
		return true
	}

	return false
}

// SetErrataList gets a reference to the given []string and assigns it to the ErrataList field.
func (o *PatchesResponse) SetErrataList(v []string) {
	o.ErrataList = &v
}

// GetLastChange returns the LastChange field value if set, zero value otherwise.
func (o *PatchesResponse) GetLastChange() string {
	if o == nil || o.LastChange == nil {
		var ret string
		return ret
	}
	return *o.LastChange
}

// GetLastChangeOk returns a tuple with the LastChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchesResponse) GetLastChangeOk() (*string, bool) {
	if o == nil || o.LastChange == nil {
		return nil, false
	}
	return o.LastChange, true
}

// HasLastChange returns a boolean if a field has been set.
func (o *PatchesResponse) HasLastChange() bool {
	if o != nil && o.LastChange != nil {
		return true
	}

	return false
}

// SetLastChange gets a reference to the given string and assigns it to the LastChange field.
func (o *PatchesResponse) SetLastChange(v string) {
	o.LastChange = &v
}

func (o PatchesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrataList != nil {
		toSerialize["errata_list"] = o.ErrataList
	}
	if o.LastChange != nil {
		toSerialize["last_change"] = o.LastChange
	}
	return json.Marshal(toSerialize)
}

type NullablePatchesResponse struct {
	value *PatchesResponse
	isSet bool
}

func (v NullablePatchesResponse) Get() *PatchesResponse {
	return v.value
}

func (v *NullablePatchesResponse) Set(val *PatchesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchesResponse(val *PatchesResponse) *NullablePatchesResponse {
	return &NullablePatchesResponse{value: val, isSet: true}
}

func (v NullablePatchesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


