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

// DBChangeResponse struct for DBChangeResponse
type DBChangeResponse struct {
	Dbchange *DBChangeResponseDbchange `json:"dbchange,omitempty"`
}

// NewDBChangeResponse instantiates a new DBChangeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDBChangeResponse() *DBChangeResponse {
	this := DBChangeResponse{}
	return &this
}

// NewDBChangeResponseWithDefaults instantiates a new DBChangeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDBChangeResponseWithDefaults() *DBChangeResponse {
	this := DBChangeResponse{}
	return &this
}

// GetDbchange returns the Dbchange field value if set, zero value otherwise.
func (o *DBChangeResponse) GetDbchange() DBChangeResponseDbchange {
	if o == nil || o.Dbchange == nil {
		var ret DBChangeResponseDbchange
		return ret
	}
	return *o.Dbchange
}

// GetDbchangeOk returns a tuple with the Dbchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBChangeResponse) GetDbchangeOk() (*DBChangeResponseDbchange, bool) {
	if o == nil || o.Dbchange == nil {
		return nil, false
	}
	return o.Dbchange, true
}

// HasDbchange returns a boolean if a field has been set.
func (o *DBChangeResponse) HasDbchange() bool {
	if o != nil && o.Dbchange != nil {
		return true
	}

	return false
}

// SetDbchange gets a reference to the given DBChangeResponseDbchange and assigns it to the Dbchange field.
func (o *DBChangeResponse) SetDbchange(v DBChangeResponseDbchange) {
	o.Dbchange = &v
}

func (o DBChangeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dbchange != nil {
		toSerialize["dbchange"] = o.Dbchange
	}
	return json.Marshal(toSerialize)
}

type NullableDBChangeResponse struct {
	value *DBChangeResponse
	isSet bool
}

func (v NullableDBChangeResponse) Get() *DBChangeResponse {
	return v.value
}

func (v *NullableDBChangeResponse) Set(val *DBChangeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDBChangeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDBChangeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDBChangeResponse(val *DBChangeResponse) *NullableDBChangeResponse {
	return &NullableDBChangeResponse{value: val, isSet: true}
}

func (v NullableDBChangeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDBChangeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


