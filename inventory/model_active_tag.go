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

// ActiveTag Information about a host tag
type ActiveTag struct {
	// The number of hosts with the given tag. If the value is null this indicates that the count is unknown.
	Count NullableInt32 `json:"count"`
	Tag StructuredTag `json:"tag"`
}

// NewActiveTag instantiates a new ActiveTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveTag(count NullableInt32, tag StructuredTag, ) *ActiveTag {
	this := ActiveTag{}
	this.Count = count
	this.Tag = tag
	return &this
}

// NewActiveTagWithDefaults instantiates a new ActiveTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveTagWithDefaults() *ActiveTag {
	this := ActiveTag{}
	return &this
}

// GetCount returns the Count field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ActiveTag) GetCount() int32 {
	if o == nil || o.Count.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Count.Get()
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveTag) GetCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Count.Get(), o.Count.IsSet()
}

// SetCount sets field value
func (o *ActiveTag) SetCount(v int32) {
	o.Count.Set(&v)
}

// GetTag returns the Tag field value
func (o *ActiveTag) GetTag() StructuredTag {
	if o == nil  {
		var ret StructuredTag
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *ActiveTag) GetTagOk() (*StructuredTag, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *ActiveTag) SetTag(v StructuredTag) {
	o.Tag = v
}

func (o ActiveTag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count.Get()
	}
	if true {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableActiveTag struct {
	value *ActiveTag
	isSet bool
}

func (v NullableActiveTag) Get() *ActiveTag {
	return v.value
}

func (v *NullableActiveTag) Set(val *ActiveTag) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveTag) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveTag(val *ActiveTag) *NullableActiveTag {
	return &NullableActiveTag{value: val, isSet: true}
}

func (v NullableActiveTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


