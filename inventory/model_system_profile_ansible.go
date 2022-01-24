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

// SystemProfileAnsible Object containing data specific to Ansible Automation Platform
type SystemProfileAnsible struct {
	// The catalog-worker version on the host
	CatalogWorkerVersion *string `json:"catalog_worker_version,omitempty"`
	// The ansible-tower or automation-controller version on the host
	ControllerVersion *string `json:"controller_version,omitempty"`
	// The automation-hub version on the host
	HubVersion *string `json:"hub_version,omitempty"`
	// The SSO version on the host
	SsoVersion *string `json:"sso_version,omitempty"`
}

// NewSystemProfileAnsible instantiates a new SystemProfileAnsible object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemProfileAnsible() *SystemProfileAnsible {
	this := SystemProfileAnsible{}
	return &this
}

// NewSystemProfileAnsibleWithDefaults instantiates a new SystemProfileAnsible object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemProfileAnsibleWithDefaults() *SystemProfileAnsible {
	this := SystemProfileAnsible{}
	return &this
}

// GetCatalogWorkerVersion returns the CatalogWorkerVersion field value if set, zero value otherwise.
func (o *SystemProfileAnsible) GetCatalogWorkerVersion() string {
	if o == nil || o.CatalogWorkerVersion == nil {
		var ret string
		return ret
	}
	return *o.CatalogWorkerVersion
}

// GetCatalogWorkerVersionOk returns a tuple with the CatalogWorkerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemProfileAnsible) GetCatalogWorkerVersionOk() (*string, bool) {
	if o == nil || o.CatalogWorkerVersion == nil {
		return nil, false
	}
	return o.CatalogWorkerVersion, true
}

// HasCatalogWorkerVersion returns a boolean if a field has been set.
func (o *SystemProfileAnsible) HasCatalogWorkerVersion() bool {
	if o != nil && o.CatalogWorkerVersion != nil {
		return true
	}

	return false
}

// SetCatalogWorkerVersion gets a reference to the given string and assigns it to the CatalogWorkerVersion field.
func (o *SystemProfileAnsible) SetCatalogWorkerVersion(v string) {
	o.CatalogWorkerVersion = &v
}

// GetControllerVersion returns the ControllerVersion field value if set, zero value otherwise.
func (o *SystemProfileAnsible) GetControllerVersion() string {
	if o == nil || o.ControllerVersion == nil {
		var ret string
		return ret
	}
	return *o.ControllerVersion
}

// GetControllerVersionOk returns a tuple with the ControllerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemProfileAnsible) GetControllerVersionOk() (*string, bool) {
	if o == nil || o.ControllerVersion == nil {
		return nil, false
	}
	return o.ControllerVersion, true
}

// HasControllerVersion returns a boolean if a field has been set.
func (o *SystemProfileAnsible) HasControllerVersion() bool {
	if o != nil && o.ControllerVersion != nil {
		return true
	}

	return false
}

// SetControllerVersion gets a reference to the given string and assigns it to the ControllerVersion field.
func (o *SystemProfileAnsible) SetControllerVersion(v string) {
	o.ControllerVersion = &v
}

// GetHubVersion returns the HubVersion field value if set, zero value otherwise.
func (o *SystemProfileAnsible) GetHubVersion() string {
	if o == nil || o.HubVersion == nil {
		var ret string
		return ret
	}
	return *o.HubVersion
}

// GetHubVersionOk returns a tuple with the HubVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemProfileAnsible) GetHubVersionOk() (*string, bool) {
	if o == nil || o.HubVersion == nil {
		return nil, false
	}
	return o.HubVersion, true
}

// HasHubVersion returns a boolean if a field has been set.
func (o *SystemProfileAnsible) HasHubVersion() bool {
	if o != nil && o.HubVersion != nil {
		return true
	}

	return false
}

// SetHubVersion gets a reference to the given string and assigns it to the HubVersion field.
func (o *SystemProfileAnsible) SetHubVersion(v string) {
	o.HubVersion = &v
}

// GetSsoVersion returns the SsoVersion field value if set, zero value otherwise.
func (o *SystemProfileAnsible) GetSsoVersion() string {
	if o == nil || o.SsoVersion == nil {
		var ret string
		return ret
	}
	return *o.SsoVersion
}

// GetSsoVersionOk returns a tuple with the SsoVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemProfileAnsible) GetSsoVersionOk() (*string, bool) {
	if o == nil || o.SsoVersion == nil {
		return nil, false
	}
	return o.SsoVersion, true
}

// HasSsoVersion returns a boolean if a field has been set.
func (o *SystemProfileAnsible) HasSsoVersion() bool {
	if o != nil && o.SsoVersion != nil {
		return true
	}

	return false
}

// SetSsoVersion gets a reference to the given string and assigns it to the SsoVersion field.
func (o *SystemProfileAnsible) SetSsoVersion(v string) {
	o.SsoVersion = &v
}

func (o SystemProfileAnsible) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CatalogWorkerVersion != nil {
		toSerialize["catalog_worker_version"] = o.CatalogWorkerVersion
	}
	if o.ControllerVersion != nil {
		toSerialize["controller_version"] = o.ControllerVersion
	}
	if o.HubVersion != nil {
		toSerialize["hub_version"] = o.HubVersion
	}
	if o.SsoVersion != nil {
		toSerialize["sso_version"] = o.SsoVersion
	}
	return json.Marshal(toSerialize)
}

type NullableSystemProfileAnsible struct {
	value *SystemProfileAnsible
	isSet bool
}

func (v NullableSystemProfileAnsible) Get() *SystemProfileAnsible {
	return v.value
}

func (v *NullableSystemProfileAnsible) Set(val *SystemProfileAnsible) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemProfileAnsible) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemProfileAnsible) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemProfileAnsible(val *SystemProfileAnsible) *NullableSystemProfileAnsible {
	return &NullableSystemProfileAnsible{value: val, isSet: true}
}

func (v NullableSystemProfileAnsible) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemProfileAnsible) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


