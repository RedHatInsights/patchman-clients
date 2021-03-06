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

// GroupPagination struct for GroupPagination
type GroupPagination struct {
	Meta *PaginationMeta `json:"meta,omitempty"`
	Links *PaginationLinks `json:"links,omitempty"`
	Data []GroupOut `json:"data"`
}

// NewGroupPagination instantiates a new GroupPagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupPagination(data []GroupOut, ) *GroupPagination {
	this := GroupPagination{}
	this.Data = data
	return &this
}

// NewGroupPaginationWithDefaults instantiates a new GroupPagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupPaginationWithDefaults() *GroupPagination {
	this := GroupPagination{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GroupPagination) GetMeta() PaginationMeta {
	if o == nil || o.Meta == nil {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPagination) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GroupPagination) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *GroupPagination) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GroupPagination) GetLinks() PaginationLinks {
	if o == nil || o.Links == nil {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPagination) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GroupPagination) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *GroupPagination) SetLinks(v PaginationLinks) {
	o.Links = &v
}

// GetData returns the Data field value
func (o *GroupPagination) GetData() []GroupOut {
	if o == nil  {
		var ret []GroupOut
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GroupPagination) GetDataOk() (*[]GroupOut, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GroupPagination) SetData(v []GroupOut) {
	o.Data = v
}

func (o GroupPagination) MarshalJSON() ([]byte, error) {
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

type NullableGroupPagination struct {
	value *GroupPagination
	isSet bool
}

func (v NullableGroupPagination) Get() *GroupPagination {
	return v.value
}

func (v *NullableGroupPagination) Set(val *GroupPagination) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupPagination) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupPagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupPagination(val *GroupPagination) *NullableGroupPagination {
	return &NullableGroupPagination{value: val, isSet: true}
}

func (v NullableGroupPagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupPagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


