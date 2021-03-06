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

// CrossAccountRequestDetailByUseIdAllOf struct for CrossAccountRequestDetailByUseIdAllOf
type CrossAccountRequestDetailByUseIdAllOf struct {
	UserId interface{} `json:"user_id,omitempty"`
}

// NewCrossAccountRequestDetailByUseIdAllOf instantiates a new CrossAccountRequestDetailByUseIdAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrossAccountRequestDetailByUseIdAllOf() *CrossAccountRequestDetailByUseIdAllOf {
	this := CrossAccountRequestDetailByUseIdAllOf{}
	return &this
}

// NewCrossAccountRequestDetailByUseIdAllOfWithDefaults instantiates a new CrossAccountRequestDetailByUseIdAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrossAccountRequestDetailByUseIdAllOfWithDefaults() *CrossAccountRequestDetailByUseIdAllOf {
	this := CrossAccountRequestDetailByUseIdAllOf{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CrossAccountRequestDetailByUseIdAllOf) GetUserId() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CrossAccountRequestDetailByUseIdAllOf) GetUserIdOk() (*interface{}, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return &o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *CrossAccountRequestDetailByUseIdAllOf) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given interface{} and assigns it to the UserId field.
func (o *CrossAccountRequestDetailByUseIdAllOf) SetUserId(v interface{}) {
	o.UserId = v
}

func (o CrossAccountRequestDetailByUseIdAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableCrossAccountRequestDetailByUseIdAllOf struct {
	value *CrossAccountRequestDetailByUseIdAllOf
	isSet bool
}

func (v NullableCrossAccountRequestDetailByUseIdAllOf) Get() *CrossAccountRequestDetailByUseIdAllOf {
	return v.value
}

func (v *NullableCrossAccountRequestDetailByUseIdAllOf) Set(val *CrossAccountRequestDetailByUseIdAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCrossAccountRequestDetailByUseIdAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCrossAccountRequestDetailByUseIdAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCrossAccountRequestDetailByUseIdAllOf(val *CrossAccountRequestDetailByUseIdAllOf) *NullableCrossAccountRequestDetailByUseIdAllOf {
	return &NullableCrossAccountRequestDetailByUseIdAllOf{value: val, isSet: true}
}

func (v NullableCrossAccountRequestDetailByUseIdAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrossAccountRequestDetailByUseIdAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


