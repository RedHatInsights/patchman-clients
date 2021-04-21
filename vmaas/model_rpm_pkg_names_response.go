/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.13.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vmaas

import (
	"encoding/json"
)

// RPMPkgNamesResponse struct for RPMPkgNamesResponse
type RPMPkgNamesResponse struct {
	LastChange interface{} `json:"last_change,omitempty"`
	RpmNameList *map[string][]string `json:"rpm_name_list,omitempty"`
}

// NewRPMPkgNamesResponse instantiates a new RPMPkgNamesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRPMPkgNamesResponse() *RPMPkgNamesResponse {
	this := RPMPkgNamesResponse{}
	return &this
}

// NewRPMPkgNamesResponseWithDefaults instantiates a new RPMPkgNamesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRPMPkgNamesResponseWithDefaults() *RPMPkgNamesResponse {
	this := RPMPkgNamesResponse{}
	return &this
}

// GetLastChange returns the LastChange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RPMPkgNamesResponse) GetLastChange() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.LastChange
}

// GetLastChangeOk returns a tuple with the LastChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RPMPkgNamesResponse) GetLastChangeOk() (*interface{}, bool) {
	if o == nil || o.LastChange == nil {
		return nil, false
	}
	return &o.LastChange, true
}

// HasLastChange returns a boolean if a field has been set.
func (o *RPMPkgNamesResponse) HasLastChange() bool {
	if o != nil && o.LastChange != nil {
		return true
	}

	return false
}

// SetLastChange gets a reference to the given interface{} and assigns it to the LastChange field.
func (o *RPMPkgNamesResponse) SetLastChange(v interface{}) {
	o.LastChange = v
}

// GetRpmNameList returns the RpmNameList field value if set, zero value otherwise.
func (o *RPMPkgNamesResponse) GetRpmNameList() map[string][]string {
	if o == nil || o.RpmNameList == nil {
		var ret map[string][]string
		return ret
	}
	return *o.RpmNameList
}

// GetRpmNameListOk returns a tuple with the RpmNameList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RPMPkgNamesResponse) GetRpmNameListOk() (*map[string][]string, bool) {
	if o == nil || o.RpmNameList == nil {
		return nil, false
	}
	return o.RpmNameList, true
}

// HasRpmNameList returns a boolean if a field has been set.
func (o *RPMPkgNamesResponse) HasRpmNameList() bool {
	if o != nil && o.RpmNameList != nil {
		return true
	}

	return false
}

// SetRpmNameList gets a reference to the given map[string][]string and assigns it to the RpmNameList field.
func (o *RPMPkgNamesResponse) SetRpmNameList(v map[string][]string) {
	o.RpmNameList = &v
}

func (o RPMPkgNamesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastChange != nil {
		toSerialize["last_change"] = o.LastChange
	}
	if o.RpmNameList != nil {
		toSerialize["rpm_name_list"] = o.RpmNameList
	}
	return json.Marshal(toSerialize)
}

type NullableRPMPkgNamesResponse struct {
	value *RPMPkgNamesResponse
	isSet bool
}

func (v NullableRPMPkgNamesResponse) Get() *RPMPkgNamesResponse {
	return v.value
}

func (v *NullableRPMPkgNamesResponse) Set(val *RPMPkgNamesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRPMPkgNamesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRPMPkgNamesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRPMPkgNamesResponse(val *RPMPkgNamesResponse) *NullableRPMPkgNamesResponse {
	return &NullableRPMPkgNamesResponse{value: val, isSet: true}
}

func (v NullableRPMPkgNamesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRPMPkgNamesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


