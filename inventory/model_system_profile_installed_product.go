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

// SystemProfileInstalledProduct Representation of one installed product
type SystemProfileInstalledProduct struct {
	// The product ID
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// Subscription status for product
	Status *string `json:"status,omitempty"`
}

// NewSystemProfileInstalledProduct instantiates a new SystemProfileInstalledProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemProfileInstalledProduct() *SystemProfileInstalledProduct {
	this := SystemProfileInstalledProduct{}
	return &this
}

// NewSystemProfileInstalledProductWithDefaults instantiates a new SystemProfileInstalledProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemProfileInstalledProductWithDefaults() *SystemProfileInstalledProduct {
	this := SystemProfileInstalledProduct{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SystemProfileInstalledProduct) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemProfileInstalledProduct) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SystemProfileInstalledProduct) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SystemProfileInstalledProduct) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SystemProfileInstalledProduct) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemProfileInstalledProduct) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SystemProfileInstalledProduct) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SystemProfileInstalledProduct) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SystemProfileInstalledProduct) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemProfileInstalledProduct) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SystemProfileInstalledProduct) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SystemProfileInstalledProduct) SetStatus(v string) {
	o.Status = &v
}

func (o SystemProfileInstalledProduct) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableSystemProfileInstalledProduct struct {
	value *SystemProfileInstalledProduct
	isSet bool
}

func (v NullableSystemProfileInstalledProduct) Get() *SystemProfileInstalledProduct {
	return v.value
}

func (v *NullableSystemProfileInstalledProduct) Set(val *SystemProfileInstalledProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemProfileInstalledProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemProfileInstalledProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemProfileInstalledProduct(val *SystemProfileInstalledProduct) *NullableSystemProfileInstalledProduct {
	return &NullableSystemProfileInstalledProduct{value: val, isSet: true}
}

func (v NullableSystemProfileInstalledProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemProfileInstalledProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

