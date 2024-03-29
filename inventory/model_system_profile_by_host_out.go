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

// SystemProfileByHostOut Structure of the output of the host system profile query
type SystemProfileByHostOut struct {
	// The number of items on the current page
	Count int32 `json:"count"`
	// The page number
	Page int32 `json:"page"`
	// The number of items to return per page
	PerPage int32 `json:"per_page"`
	// Total number of items
	Total int32 `json:"total"`
	// Actual host search query result entries.
	Results []HostSystemProfileOut `json:"results"`
}

// NewSystemProfileByHostOut instantiates a new SystemProfileByHostOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemProfileByHostOut(count int32, page int32, perPage int32, total int32, results []HostSystemProfileOut, ) *SystemProfileByHostOut {
	this := SystemProfileByHostOut{}
	this.Count = count
	this.Page = page
	this.PerPage = perPage
	this.Total = total
	this.Results = results
	return &this
}

// NewSystemProfileByHostOutWithDefaults instantiates a new SystemProfileByHostOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemProfileByHostOutWithDefaults() *SystemProfileByHostOut {
	this := SystemProfileByHostOut{}
	return &this
}

// GetCount returns the Count field value
func (o *SystemProfileByHostOut) GetCount() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *SystemProfileByHostOut) GetCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *SystemProfileByHostOut) SetCount(v int32) {
	o.Count = v
}

// GetPage returns the Page field value
func (o *SystemProfileByHostOut) GetPage() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *SystemProfileByHostOut) GetPageOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *SystemProfileByHostOut) SetPage(v int32) {
	o.Page = v
}

// GetPerPage returns the PerPage field value
func (o *SystemProfileByHostOut) GetPerPage() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.PerPage
}

// GetPerPageOk returns a tuple with the PerPage field value
// and a boolean to check if the value has been set.
func (o *SystemProfileByHostOut) GetPerPageOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PerPage, true
}

// SetPerPage sets field value
func (o *SystemProfileByHostOut) SetPerPage(v int32) {
	o.PerPage = v
}

// GetTotal returns the Total field value
func (o *SystemProfileByHostOut) GetTotal() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *SystemProfileByHostOut) GetTotalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *SystemProfileByHostOut) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *SystemProfileByHostOut) GetResults() []HostSystemProfileOut {
	if o == nil  {
		var ret []HostSystemProfileOut
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *SystemProfileByHostOut) GetResultsOk() (*[]HostSystemProfileOut, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *SystemProfileByHostOut) SetResults(v []HostSystemProfileOut) {
	o.Results = v
}

func (o SystemProfileByHostOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["page"] = o.Page
	}
	if true {
		toSerialize["per_page"] = o.PerPage
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableSystemProfileByHostOut struct {
	value *SystemProfileByHostOut
	isSet bool
}

func (v NullableSystemProfileByHostOut) Get() *SystemProfileByHostOut {
	return v.value
}

func (v *NullableSystemProfileByHostOut) Set(val *SystemProfileByHostOut) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemProfileByHostOut) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemProfileByHostOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemProfileByHostOut(val *SystemProfileByHostOut) *NullableSystemProfileByHostOut {
	return &NullableSystemProfileByHostOut{value: val, isSet: true}
}

func (v NullableSystemProfileByHostOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemProfileByHostOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


