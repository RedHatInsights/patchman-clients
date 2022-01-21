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

// FactSet A set of string facts belonging to a single namespace.
type FactSet struct {
	// A namespace the facts belong to.
	Namespace string `json:"namespace"`
	// The facts themselves.
	Facts map[string]interface{} `json:"facts"`
}

// NewFactSet instantiates a new FactSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFactSet(namespace string, facts map[string]interface{}, ) *FactSet {
	this := FactSet{}
	this.Namespace = namespace
	this.Facts = facts
	return &this
}

// NewFactSetWithDefaults instantiates a new FactSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFactSetWithDefaults() *FactSet {
	this := FactSet{}
	return &this
}

// GetNamespace returns the Namespace field value
func (o *FactSet) GetNamespace() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *FactSet) GetNamespaceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *FactSet) SetNamespace(v string) {
	o.Namespace = v
}

// GetFacts returns the Facts field value
func (o *FactSet) GetFacts() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}

	return o.Facts
}

// GetFactsOk returns a tuple with the Facts field value
// and a boolean to check if the value has been set.
func (o *FactSet) GetFactsOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Facts, true
}

// SetFacts sets field value
func (o *FactSet) SetFacts(v map[string]interface{}) {
	o.Facts = v
}

func (o FactSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["namespace"] = o.Namespace
	}
	if true {
		toSerialize["facts"] = o.Facts
	}
	return json.Marshal(toSerialize)
}

type NullableFactSet struct {
	value *FactSet
	isSet bool
}

func (v NullableFactSet) Get() *FactSet {
	return v.value
}

func (v *NullableFactSet) Set(val *FactSet) {
	v.value = val
	v.isSet = true
}

func (v NullableFactSet) IsSet() bool {
	return v.isSet
}

func (v *NullableFactSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFactSet(val *FactSet) *NullableFactSet {
	return &NullableFactSet{value: val, isSet: true}
}

func (v NullableFactSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFactSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


