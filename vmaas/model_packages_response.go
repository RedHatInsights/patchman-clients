/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.27.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vmaas

import (
	"encoding/json"
)

// PackagesResponse struct for PackagesResponse
type PackagesResponse struct {
	PackageList *map[string]PackagesResponsePackageList `json:"package_list,omitempty"`
	LastChange *string `json:"last_change,omitempty"`
}

// NewPackagesResponse instantiates a new PackagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackagesResponse() *PackagesResponse {
	this := PackagesResponse{}
	return &this
}

// NewPackagesResponseWithDefaults instantiates a new PackagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackagesResponseWithDefaults() *PackagesResponse {
	this := PackagesResponse{}
	return &this
}

// GetPackageList returns the PackageList field value if set, zero value otherwise.
func (o *PackagesResponse) GetPackageList() map[string]PackagesResponsePackageList {
	if o == nil || o.PackageList == nil {
		var ret map[string]PackagesResponsePackageList
		return ret
	}
	return *o.PackageList
}

// GetPackageListOk returns a tuple with the PackageList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesResponse) GetPackageListOk() (*map[string]PackagesResponsePackageList, bool) {
	if o == nil || o.PackageList == nil {
		return nil, false
	}
	return o.PackageList, true
}

// HasPackageList returns a boolean if a field has been set.
func (o *PackagesResponse) HasPackageList() bool {
	if o != nil && o.PackageList != nil {
		return true
	}

	return false
}

// SetPackageList gets a reference to the given map[string]PackagesResponsePackageList and assigns it to the PackageList field.
func (o *PackagesResponse) SetPackageList(v map[string]PackagesResponsePackageList) {
	o.PackageList = &v
}

// GetLastChange returns the LastChange field value if set, zero value otherwise.
func (o *PackagesResponse) GetLastChange() string {
	if o == nil || o.LastChange == nil {
		var ret string
		return ret
	}
	return *o.LastChange
}

// GetLastChangeOk returns a tuple with the LastChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesResponse) GetLastChangeOk() (*string, bool) {
	if o == nil || o.LastChange == nil {
		return nil, false
	}
	return o.LastChange, true
}

// HasLastChange returns a boolean if a field has been set.
func (o *PackagesResponse) HasLastChange() bool {
	if o != nil && o.LastChange != nil {
		return true
	}

	return false
}

// SetLastChange gets a reference to the given string and assigns it to the LastChange field.
func (o *PackagesResponse) SetLastChange(v string) {
	o.LastChange = &v
}

func (o PackagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PackageList != nil {
		toSerialize["package_list"] = o.PackageList
	}
	if o.LastChange != nil {
		toSerialize["last_change"] = o.LastChange
	}
	return json.Marshal(toSerialize)
}

type NullablePackagesResponse struct {
	value *PackagesResponse
	isSet bool
}

func (v NullablePackagesResponse) Get() *PackagesResponse {
	return v.value
}

func (v *NullablePackagesResponse) Set(val *PackagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePackagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePackagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackagesResponse(val *PackagesResponse) *NullablePackagesResponse {
	return &NullablePackagesResponse{value: val, isSet: true}
}

func (v NullablePackagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


