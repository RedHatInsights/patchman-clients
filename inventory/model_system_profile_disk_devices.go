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

// SystemProfileDiskDevices Representation of one mounted device
type SystemProfileDiskDevices struct {
	Device *string `json:"device,omitempty"`
	// User-defined mount label
	Label *string `json:"label,omitempty"`
	// The mount point
	MountPoint *string `json:"mount_point,omitempty"`
	// Mount options for nested object
	Options *map[string]interface{} `json:"options,omitempty"`
	// The mount type
	Type *string `json:"type,omitempty"`
}

// NewSystemProfileDiskDevices instantiates a new SystemProfileDiskDevices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemProfileDiskDevices() *SystemProfileDiskDevices {
	this := SystemProfileDiskDevices{}
	return &this
}

// NewSystemProfileDiskDevicesWithDefaults instantiates a new SystemProfileDiskDevices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemProfileDiskDevicesWithDefaults() *SystemProfileDiskDevices {
	this := SystemProfileDiskDevices{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *SystemProfileDiskDevices) GetDevice() string {
	if o == nil || o.Device == nil {
		var ret string
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemProfileDiskDevices) GetDeviceOk() (*string, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *SystemProfileDiskDevices) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given string and assigns it to the Device field.
func (o *SystemProfileDiskDevices) SetDevice(v string) {
	o.Device = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *SystemProfileDiskDevices) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemProfileDiskDevices) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *SystemProfileDiskDevices) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *SystemProfileDiskDevices) SetLabel(v string) {
	o.Label = &v
}

// GetMountPoint returns the MountPoint field value if set, zero value otherwise.
func (o *SystemProfileDiskDevices) GetMountPoint() string {
	if o == nil || o.MountPoint == nil {
		var ret string
		return ret
	}
	return *o.MountPoint
}

// GetMountPointOk returns a tuple with the MountPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemProfileDiskDevices) GetMountPointOk() (*string, bool) {
	if o == nil || o.MountPoint == nil {
		return nil, false
	}
	return o.MountPoint, true
}

// HasMountPoint returns a boolean if a field has been set.
func (o *SystemProfileDiskDevices) HasMountPoint() bool {
	if o != nil && o.MountPoint != nil {
		return true
	}

	return false
}

// SetMountPoint gets a reference to the given string and assigns it to the MountPoint field.
func (o *SystemProfileDiskDevices) SetMountPoint(v string) {
	o.MountPoint = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SystemProfileDiskDevices) GetOptions() map[string]interface{} {
	if o == nil || o.Options == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemProfileDiskDevices) GetOptionsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SystemProfileDiskDevices) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *SystemProfileDiskDevices) SetOptions(v map[string]interface{}) {
	o.Options = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SystemProfileDiskDevices) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemProfileDiskDevices) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SystemProfileDiskDevices) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SystemProfileDiskDevices) SetType(v string) {
	o.Type = &v
}

func (o SystemProfileDiskDevices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.MountPoint != nil {
		toSerialize["mount_point"] = o.MountPoint
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSystemProfileDiskDevices struct {
	value *SystemProfileDiskDevices
	isSet bool
}

func (v NullableSystemProfileDiskDevices) Get() *SystemProfileDiskDevices {
	return v.value
}

func (v *NullableSystemProfileDiskDevices) Set(val *SystemProfileDiskDevices) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemProfileDiskDevices) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemProfileDiskDevices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemProfileDiskDevices(val *SystemProfileDiskDevices) *NullableSystemProfileDiskDevices {
	return &NullableSystemProfileDiskDevices{value: val, isSet: true}
}

func (v NullableSystemProfileDiskDevices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemProfileDiskDevices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

