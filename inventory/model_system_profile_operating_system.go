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

// SystemProfileOperatingSystem Object for OS details. Supports range operations
type SystemProfileOperatingSystem struct {
	// Major release of OS (aka the x version)
	Major *int32 `json:"major,omitempty"`
	// Minor release of OS (aka the y version)
	Minor *int32 `json:"minor,omitempty"`
	// Name of the distro/os
	Name *string `json:"name,omitempty"`
}

// NewSystemProfileOperatingSystem instantiates a new SystemProfileOperatingSystem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemProfileOperatingSystem() *SystemProfileOperatingSystem {
	this := SystemProfileOperatingSystem{}
	return &this
}

// NewSystemProfileOperatingSystemWithDefaults instantiates a new SystemProfileOperatingSystem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemProfileOperatingSystemWithDefaults() *SystemProfileOperatingSystem {
	this := SystemProfileOperatingSystem{}
	return &this
}

// GetMajor returns the Major field value if set, zero value otherwise.
func (o *SystemProfileOperatingSystem) GetMajor() int32 {
	if o == nil || o.Major == nil {
		var ret int32
		return ret
	}
	return *o.Major
}

// GetMajorOk returns a tuple with the Major field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemProfileOperatingSystem) GetMajorOk() (*int32, bool) {
	if o == nil || o.Major == nil {
		return nil, false
	}
	return o.Major, true
}

// HasMajor returns a boolean if a field has been set.
func (o *SystemProfileOperatingSystem) HasMajor() bool {
	if o != nil && o.Major != nil {
		return true
	}

	return false
}

// SetMajor gets a reference to the given int32 and assigns it to the Major field.
func (o *SystemProfileOperatingSystem) SetMajor(v int32) {
	o.Major = &v
}

// GetMinor returns the Minor field value if set, zero value otherwise.
func (o *SystemProfileOperatingSystem) GetMinor() int32 {
	if o == nil || o.Minor == nil {
		var ret int32
		return ret
	}
	return *o.Minor
}

// GetMinorOk returns a tuple with the Minor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemProfileOperatingSystem) GetMinorOk() (*int32, bool) {
	if o == nil || o.Minor == nil {
		return nil, false
	}
	return o.Minor, true
}

// HasMinor returns a boolean if a field has been set.
func (o *SystemProfileOperatingSystem) HasMinor() bool {
	if o != nil && o.Minor != nil {
		return true
	}

	return false
}

// SetMinor gets a reference to the given int32 and assigns it to the Minor field.
func (o *SystemProfileOperatingSystem) SetMinor(v int32) {
	o.Minor = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SystemProfileOperatingSystem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemProfileOperatingSystem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SystemProfileOperatingSystem) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SystemProfileOperatingSystem) SetName(v string) {
	o.Name = &v
}

func (o SystemProfileOperatingSystem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Major != nil {
		toSerialize["major"] = o.Major
	}
	if o.Minor != nil {
		toSerialize["minor"] = o.Minor
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableSystemProfileOperatingSystem struct {
	value *SystemProfileOperatingSystem
	isSet bool
}

func (v NullableSystemProfileOperatingSystem) Get() *SystemProfileOperatingSystem {
	return v.value
}

func (v *NullableSystemProfileOperatingSystem) Set(val *SystemProfileOperatingSystem) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemProfileOperatingSystem) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemProfileOperatingSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemProfileOperatingSystem(val *SystemProfileOperatingSystem) *NullableSystemProfileOperatingSystem {
	return &NullableSystemProfileOperatingSystem{value: val, isSet: true}
}

func (v NullableSystemProfileOperatingSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemProfileOperatingSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

