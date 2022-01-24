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

// SystemProfileSystemPurpose Object for system purpose information
type SystemProfileSystemPurpose struct {
	// The intended role of the system
	Role *string `json:"role,omitempty"`
	// The intended SLA of the system
	Sla *string `json:"sla,omitempty"`
	// The intended usage of the system
	Usage *string `json:"usage,omitempty"`
}

// NewSystemProfileSystemPurpose instantiates a new SystemProfileSystemPurpose object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemProfileSystemPurpose() *SystemProfileSystemPurpose {
	this := SystemProfileSystemPurpose{}
	return &this
}

// NewSystemProfileSystemPurposeWithDefaults instantiates a new SystemProfileSystemPurpose object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemProfileSystemPurposeWithDefaults() *SystemProfileSystemPurpose {
	this := SystemProfileSystemPurpose{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *SystemProfileSystemPurpose) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemProfileSystemPurpose) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *SystemProfileSystemPurpose) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *SystemProfileSystemPurpose) SetRole(v string) {
	o.Role = &v
}

// GetSla returns the Sla field value if set, zero value otherwise.
func (o *SystemProfileSystemPurpose) GetSla() string {
	if o == nil || o.Sla == nil {
		var ret string
		return ret
	}
	return *o.Sla
}

// GetSlaOk returns a tuple with the Sla field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemProfileSystemPurpose) GetSlaOk() (*string, bool) {
	if o == nil || o.Sla == nil {
		return nil, false
	}
	return o.Sla, true
}

// HasSla returns a boolean if a field has been set.
func (o *SystemProfileSystemPurpose) HasSla() bool {
	if o != nil && o.Sla != nil {
		return true
	}

	return false
}

// SetSla gets a reference to the given string and assigns it to the Sla field.
func (o *SystemProfileSystemPurpose) SetSla(v string) {
	o.Sla = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *SystemProfileSystemPurpose) GetUsage() string {
	if o == nil || o.Usage == nil {
		var ret string
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemProfileSystemPurpose) GetUsageOk() (*string, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *SystemProfileSystemPurpose) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given string and assigns it to the Usage field.
func (o *SystemProfileSystemPurpose) SetUsage(v string) {
	o.Usage = &v
}

func (o SystemProfileSystemPurpose) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Sla != nil {
		toSerialize["sla"] = o.Sla
	}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableSystemProfileSystemPurpose struct {
	value *SystemProfileSystemPurpose
	isSet bool
}

func (v NullableSystemProfileSystemPurpose) Get() *SystemProfileSystemPurpose {
	return v.value
}

func (v *NullableSystemProfileSystemPurpose) Set(val *SystemProfileSystemPurpose) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemProfileSystemPurpose) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemProfileSystemPurpose) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemProfileSystemPurpose(val *SystemProfileSystemPurpose) *NullableSystemProfileSystemPurpose {
	return &NullableSystemProfileSystemPurpose{value: val, isSet: true}
}

func (v NullableSystemProfileSystemPurpose) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemProfileSystemPurpose) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


