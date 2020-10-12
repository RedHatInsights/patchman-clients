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

// PermissionPagination struct for PermissionPagination
type PermissionPagination struct {
	Meta *PaginationMeta `json:"meta,omitempty"`
	Links *PaginationLinks `json:"links,omitempty"`
	Data []Permission `json:"data"`
}

// NewPermissionPagination instantiates a new PermissionPagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionPagination(data []Permission, ) *PermissionPagination {
	this := PermissionPagination{}
	this.Data = data
	return &this
}

// NewPermissionPaginationWithDefaults instantiates a new PermissionPagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionPaginationWithDefaults() *PermissionPagination {
	this := PermissionPagination{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *PermissionPagination) GetMeta() PaginationMeta {
	if o == nil || o.Meta == nil {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionPagination) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *PermissionPagination) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *PermissionPagination) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PermissionPagination) GetLinks() PaginationLinks {
	if o == nil || o.Links == nil {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionPagination) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PermissionPagination) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *PermissionPagination) SetLinks(v PaginationLinks) {
	o.Links = &v
}

// GetData returns the Data field value
func (o *PermissionPagination) GetData() []Permission {
	if o == nil  {
		var ret []Permission
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PermissionPagination) GetDataOk() (*[]Permission, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PermissionPagination) SetData(v []Permission) {
	o.Data = v
}

func (o PermissionPagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePermissionPagination struct {
	value *PermissionPagination
	isSet bool
}

func (v NullablePermissionPagination) Get() *PermissionPagination {
	return v.value
}

func (v *NullablePermissionPagination) Set(val *PermissionPagination) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionPagination) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionPagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionPagination(val *PermissionPagination) *NullablePermissionPagination {
	return &NullablePermissionPagination{value: val, isSet: true}
}

func (v NullablePermissionPagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionPagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

