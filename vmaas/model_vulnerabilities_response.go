/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.31.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vmaas

import (
	"encoding/json"
)

// VulnerabilitiesResponse struct for VulnerabilitiesResponse
type VulnerabilitiesResponse struct {
	CveList []string `json:"cve_list"`
	ManuallyFixableCveList []string `json:"manually_fixable_cve_list"`
	UnpatchedCveList []string `json:"unpatched_cve_list"`
	LastChange *string `json:"last_change,omitempty"`
}

// NewVulnerabilitiesResponse instantiates a new VulnerabilitiesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVulnerabilitiesResponse(cveList []string, manuallyFixableCveList []string, unpatchedCveList []string, ) *VulnerabilitiesResponse {
	this := VulnerabilitiesResponse{}
	this.CveList = cveList
	this.ManuallyFixableCveList = manuallyFixableCveList
	this.UnpatchedCveList = unpatchedCveList
	return &this
}

// NewVulnerabilitiesResponseWithDefaults instantiates a new VulnerabilitiesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVulnerabilitiesResponseWithDefaults() *VulnerabilitiesResponse {
	this := VulnerabilitiesResponse{}
	return &this
}

// GetCveList returns the CveList field value
func (o *VulnerabilitiesResponse) GetCveList() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.CveList
}

// GetCveListOk returns a tuple with the CveList field value
// and a boolean to check if the value has been set.
func (o *VulnerabilitiesResponse) GetCveListOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CveList, true
}

// SetCveList sets field value
func (o *VulnerabilitiesResponse) SetCveList(v []string) {
	o.CveList = v
}

// GetManuallyFixableCveList returns the ManuallyFixableCveList field value
func (o *VulnerabilitiesResponse) GetManuallyFixableCveList() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.ManuallyFixableCveList
}

// GetManuallyFixableCveListOk returns a tuple with the ManuallyFixableCveList field value
// and a boolean to check if the value has been set.
func (o *VulnerabilitiesResponse) GetManuallyFixableCveListOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ManuallyFixableCveList, true
}

// SetManuallyFixableCveList sets field value
func (o *VulnerabilitiesResponse) SetManuallyFixableCveList(v []string) {
	o.ManuallyFixableCveList = v
}

// GetUnpatchedCveList returns the UnpatchedCveList field value
func (o *VulnerabilitiesResponse) GetUnpatchedCveList() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.UnpatchedCveList
}

// GetUnpatchedCveListOk returns a tuple with the UnpatchedCveList field value
// and a boolean to check if the value has been set.
func (o *VulnerabilitiesResponse) GetUnpatchedCveListOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UnpatchedCveList, true
}

// SetUnpatchedCveList sets field value
func (o *VulnerabilitiesResponse) SetUnpatchedCveList(v []string) {
	o.UnpatchedCveList = v
}

// GetLastChange returns the LastChange field value if set, zero value otherwise.
func (o *VulnerabilitiesResponse) GetLastChange() string {
	if o == nil || o.LastChange == nil {
		var ret string
		return ret
	}
	return *o.LastChange
}

// GetLastChangeOk returns a tuple with the LastChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnerabilitiesResponse) GetLastChangeOk() (*string, bool) {
	if o == nil || o.LastChange == nil {
		return nil, false
	}
	return o.LastChange, true
}

// HasLastChange returns a boolean if a field has been set.
func (o *VulnerabilitiesResponse) HasLastChange() bool {
	if o != nil && o.LastChange != nil {
		return true
	}

	return false
}

// SetLastChange gets a reference to the given string and assigns it to the LastChange field.
func (o *VulnerabilitiesResponse) SetLastChange(v string) {
	o.LastChange = &v
}

func (o VulnerabilitiesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cve_list"] = o.CveList
	}
	if true {
		toSerialize["manually_fixable_cve_list"] = o.ManuallyFixableCveList
	}
	if true {
		toSerialize["unpatched_cve_list"] = o.UnpatchedCveList
	}
	if o.LastChange != nil {
		toSerialize["last_change"] = o.LastChange
	}
	return json.Marshal(toSerialize)
}

type NullableVulnerabilitiesResponse struct {
	value *VulnerabilitiesResponse
	isSet bool
}

func (v NullableVulnerabilitiesResponse) Get() *VulnerabilitiesResponse {
	return v.value
}

func (v *NullableVulnerabilitiesResponse) Set(val *VulnerabilitiesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVulnerabilitiesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVulnerabilitiesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVulnerabilitiesResponse(val *VulnerabilitiesResponse) *NullableVulnerabilitiesResponse {
	return &NullableVulnerabilitiesResponse{value: val, isSet: true}
}

func (v NullableVulnerabilitiesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVulnerabilitiesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


