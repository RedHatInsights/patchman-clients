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

// HostQueryOutputAllOf struct for HostQueryOutputAllOf
type HostQueryOutputAllOf struct {
	// Actual host search query result entries.
	Results []HostOut `json:"results"`
}

// NewHostQueryOutputAllOf instantiates a new HostQueryOutputAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostQueryOutputAllOf(results []HostOut, ) *HostQueryOutputAllOf {
	this := HostQueryOutputAllOf{}
	this.Results = results
	return &this
}

// NewHostQueryOutputAllOfWithDefaults instantiates a new HostQueryOutputAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostQueryOutputAllOfWithDefaults() *HostQueryOutputAllOf {
	this := HostQueryOutputAllOf{}
	return &this
}

// GetResults returns the Results field value
func (o *HostQueryOutputAllOf) GetResults() []HostOut {
	if o == nil  {
		var ret []HostOut
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *HostQueryOutputAllOf) GetResultsOk() (*[]HostOut, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *HostQueryOutputAllOf) SetResults(v []HostOut) {
	o.Results = v
}

func (o HostQueryOutputAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableHostQueryOutputAllOf struct {
	value *HostQueryOutputAllOf
	isSet bool
}

func (v NullableHostQueryOutputAllOf) Get() *HostQueryOutputAllOf {
	return v.value
}

func (v *NullableHostQueryOutputAllOf) Set(val *HostQueryOutputAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHostQueryOutputAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHostQueryOutputAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostQueryOutputAllOf(val *HostQueryOutputAllOf) *NullableHostQueryOutputAllOf {
	return &NullableHostQueryOutputAllOf{value: val, isSet: true}
}

func (v NullableHostQueryOutputAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostQueryOutputAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


