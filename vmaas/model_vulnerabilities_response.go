/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.14.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vmaas

import (
	"encoding/json"
)

// VulnerabilitiesResponse struct for VulnerabilitiesResponse
type VulnerabilitiesResponse struct {
	CveList []string `json:"cve_list"`
	UnpatchedCveList []string `json:"unpatched_cve_list"`
}

// NewVulnerabilitiesResponse instantiates a new VulnerabilitiesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVulnerabilitiesResponse(cveList []string, unpatchedCveList []string, ) *VulnerabilitiesResponse {
	this := VulnerabilitiesResponse{}
	this.CveList = cveList
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

func (o VulnerabilitiesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cve_list"] = o.CveList
	}
	if true {
		toSerialize["unpatched_cve_list"] = o.UnpatchedCveList
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


