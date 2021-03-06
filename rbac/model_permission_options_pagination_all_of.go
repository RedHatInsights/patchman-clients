/*
 * Role Based Access Control
 *
 * The API for Role Based Access Control.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rbac

import (
	"encoding/json"
)

// PermissionOptionsPaginationAllOf struct for PermissionOptionsPaginationAllOf
type PermissionOptionsPaginationAllOf struct {
	Data []string `json:"data"`
}

// NewPermissionOptionsPaginationAllOf instantiates a new PermissionOptionsPaginationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionOptionsPaginationAllOf(data []string, ) *PermissionOptionsPaginationAllOf {
	this := PermissionOptionsPaginationAllOf{}
	this.Data = data
	return &this
}

// NewPermissionOptionsPaginationAllOfWithDefaults instantiates a new PermissionOptionsPaginationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionOptionsPaginationAllOfWithDefaults() *PermissionOptionsPaginationAllOf {
	this := PermissionOptionsPaginationAllOf{}
	return &this
}

// GetData returns the Data field value
func (o *PermissionOptionsPaginationAllOf) GetData() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PermissionOptionsPaginationAllOf) GetDataOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PermissionOptionsPaginationAllOf) SetData(v []string) {
	o.Data = v
}

func (o PermissionOptionsPaginationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePermissionOptionsPaginationAllOf struct {
	value *PermissionOptionsPaginationAllOf
	isSet bool
}

func (v NullablePermissionOptionsPaginationAllOf) Get() *PermissionOptionsPaginationAllOf {
	return v.value
}

func (v *NullablePermissionOptionsPaginationAllOf) Set(val *PermissionOptionsPaginationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionOptionsPaginationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionOptionsPaginationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionOptionsPaginationAllOf(val *PermissionOptionsPaginationAllOf) *NullablePermissionOptionsPaginationAllOf {
	return &NullablePermissionOptionsPaginationAllOf{value: val, isSet: true}
}

func (v NullablePermissionOptionsPaginationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionOptionsPaginationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


