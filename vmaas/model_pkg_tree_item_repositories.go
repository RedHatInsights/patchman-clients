/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.24.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vmaas

import (
	"encoding/json"
)

// PkgTreeItemRepositories struct for PkgTreeItemRepositories
type PkgTreeItemRepositories struct {
	Label string `json:"label"`
	Name string `json:"name"`
	Basearch string `json:"basearch"`
	Releasever string `json:"releasever"`
	Revision string `json:"revision"`
	ModuleName *string `json:"module_name,omitempty"`
	ModuleStream *string `json:"module_stream,omitempty"`
}

// NewPkgTreeItemRepositories instantiates a new PkgTreeItemRepositories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkgTreeItemRepositories(label string, name string, basearch string, releasever string, revision string, ) *PkgTreeItemRepositories {
	this := PkgTreeItemRepositories{}
	this.Label = label
	this.Name = name
	this.Basearch = basearch
	this.Releasever = releasever
	this.Revision = revision
	return &this
}

// NewPkgTreeItemRepositoriesWithDefaults instantiates a new PkgTreeItemRepositories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkgTreeItemRepositoriesWithDefaults() *PkgTreeItemRepositories {
	this := PkgTreeItemRepositories{}
	return &this
}

// GetLabel returns the Label field value
func (o *PkgTreeItemRepositories) GetLabel() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *PkgTreeItemRepositories) GetLabelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *PkgTreeItemRepositories) SetLabel(v string) {
	o.Label = v
}

// GetName returns the Name field value
func (o *PkgTreeItemRepositories) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PkgTreeItemRepositories) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PkgTreeItemRepositories) SetName(v string) {
	o.Name = v
}

// GetBasearch returns the Basearch field value
func (o *PkgTreeItemRepositories) GetBasearch() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Basearch
}

// GetBasearchOk returns a tuple with the Basearch field value
// and a boolean to check if the value has been set.
func (o *PkgTreeItemRepositories) GetBasearchOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Basearch, true
}

// SetBasearch sets field value
func (o *PkgTreeItemRepositories) SetBasearch(v string) {
	o.Basearch = v
}

// GetReleasever returns the Releasever field value
func (o *PkgTreeItemRepositories) GetReleasever() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Releasever
}

// GetReleaseverOk returns a tuple with the Releasever field value
// and a boolean to check if the value has been set.
func (o *PkgTreeItemRepositories) GetReleaseverOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Releasever, true
}

// SetReleasever sets field value
func (o *PkgTreeItemRepositories) SetReleasever(v string) {
	o.Releasever = v
}

// GetRevision returns the Revision field value
func (o *PkgTreeItemRepositories) GetRevision() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *PkgTreeItemRepositories) GetRevisionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *PkgTreeItemRepositories) SetRevision(v string) {
	o.Revision = v
}

// GetModuleName returns the ModuleName field value if set, zero value otherwise.
func (o *PkgTreeItemRepositories) GetModuleName() string {
	if o == nil || o.ModuleName == nil {
		var ret string
		return ret
	}
	return *o.ModuleName
}

// GetModuleNameOk returns a tuple with the ModuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgTreeItemRepositories) GetModuleNameOk() (*string, bool) {
	if o == nil || o.ModuleName == nil {
		return nil, false
	}
	return o.ModuleName, true
}

// HasModuleName returns a boolean if a field has been set.
func (o *PkgTreeItemRepositories) HasModuleName() bool {
	if o != nil && o.ModuleName != nil {
		return true
	}

	return false
}

// SetModuleName gets a reference to the given string and assigns it to the ModuleName field.
func (o *PkgTreeItemRepositories) SetModuleName(v string) {
	o.ModuleName = &v
}

// GetModuleStream returns the ModuleStream field value if set, zero value otherwise.
func (o *PkgTreeItemRepositories) GetModuleStream() string {
	if o == nil || o.ModuleStream == nil {
		var ret string
		return ret
	}
	return *o.ModuleStream
}

// GetModuleStreamOk returns a tuple with the ModuleStream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgTreeItemRepositories) GetModuleStreamOk() (*string, bool) {
	if o == nil || o.ModuleStream == nil {
		return nil, false
	}
	return o.ModuleStream, true
}

// HasModuleStream returns a boolean if a field has been set.
func (o *PkgTreeItemRepositories) HasModuleStream() bool {
	if o != nil && o.ModuleStream != nil {
		return true
	}

	return false
}

// SetModuleStream gets a reference to the given string and assigns it to the ModuleStream field.
func (o *PkgTreeItemRepositories) SetModuleStream(v string) {
	o.ModuleStream = &v
}

func (o PkgTreeItemRepositories) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["basearch"] = o.Basearch
	}
	if true {
		toSerialize["releasever"] = o.Releasever
	}
	if true {
		toSerialize["revision"] = o.Revision
	}
	if o.ModuleName != nil {
		toSerialize["module_name"] = o.ModuleName
	}
	if o.ModuleStream != nil {
		toSerialize["module_stream"] = o.ModuleStream
	}
	return json.Marshal(toSerialize)
}

type NullablePkgTreeItemRepositories struct {
	value *PkgTreeItemRepositories
	isSet bool
}

func (v NullablePkgTreeItemRepositories) Get() *PkgTreeItemRepositories {
	return v.value
}

func (v *NullablePkgTreeItemRepositories) Set(val *PkgTreeItemRepositories) {
	v.value = val
	v.isSet = true
}

func (v NullablePkgTreeItemRepositories) IsSet() bool {
	return v.isSet
}

func (v *NullablePkgTreeItemRepositories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkgTreeItemRepositories(val *PkgTreeItemRepositories) *NullablePkgTreeItemRepositories {
	return &NullablePkgTreeItemRepositories{value: val, isSet: true}
}

func (v NullablePkgTreeItemRepositories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkgTreeItemRepositories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


