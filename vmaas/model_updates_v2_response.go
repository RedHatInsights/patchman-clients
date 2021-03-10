/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vmaas

import (
	"encoding/json"
)

// UpdatesV2Response struct for UpdatesV2Response
type UpdatesV2Response struct {
	UpdateList *map[string]UpdatesV2ResponseUpdateList `json:"update_list,omitempty"`
	RepositoryList *[]string `json:"repository_list,omitempty"`
	ModulesList *[]UpdatesV3RequestModulesList `json:"modules_list,omitempty"`
	Releasever *string `json:"releasever,omitempty"`
	Basearch *string `json:"basearch,omitempty"`
}

// NewUpdatesV2Response instantiates a new UpdatesV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatesV2Response() *UpdatesV2Response {
	this := UpdatesV2Response{}
	return &this
}

// NewUpdatesV2ResponseWithDefaults instantiates a new UpdatesV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatesV2ResponseWithDefaults() *UpdatesV2Response {
	this := UpdatesV2Response{}
	return &this
}

// GetUpdateList returns the UpdateList field value if set, zero value otherwise.
func (o *UpdatesV2Response) GetUpdateList() map[string]UpdatesV2ResponseUpdateList {
	if o == nil || o.UpdateList == nil {
		var ret map[string]UpdatesV2ResponseUpdateList
		return ret
	}
	return *o.UpdateList
}

// GetUpdateListOk returns a tuple with the UpdateList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesV2Response) GetUpdateListOk() (*map[string]UpdatesV2ResponseUpdateList, bool) {
	if o == nil || o.UpdateList == nil {
		return nil, false
	}
	return o.UpdateList, true
}

// HasUpdateList returns a boolean if a field has been set.
func (o *UpdatesV2Response) HasUpdateList() bool {
	if o != nil && o.UpdateList != nil {
		return true
	}

	return false
}

// SetUpdateList gets a reference to the given map[string]UpdatesV2ResponseUpdateList and assigns it to the UpdateList field.
func (o *UpdatesV2Response) SetUpdateList(v map[string]UpdatesV2ResponseUpdateList) {
	o.UpdateList = &v
}

// GetRepositoryList returns the RepositoryList field value if set, zero value otherwise.
func (o *UpdatesV2Response) GetRepositoryList() []string {
	if o == nil || o.RepositoryList == nil {
		var ret []string
		return ret
	}
	return *o.RepositoryList
}

// GetRepositoryListOk returns a tuple with the RepositoryList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesV2Response) GetRepositoryListOk() (*[]string, bool) {
	if o == nil || o.RepositoryList == nil {
		return nil, false
	}
	return o.RepositoryList, true
}

// HasRepositoryList returns a boolean if a field has been set.
func (o *UpdatesV2Response) HasRepositoryList() bool {
	if o != nil && o.RepositoryList != nil {
		return true
	}

	return false
}

// SetRepositoryList gets a reference to the given []string and assigns it to the RepositoryList field.
func (o *UpdatesV2Response) SetRepositoryList(v []string) {
	o.RepositoryList = &v
}

// GetModulesList returns the ModulesList field value if set, zero value otherwise.
func (o *UpdatesV2Response) GetModulesList() []UpdatesV3RequestModulesList {
	if o == nil || o.ModulesList == nil {
		var ret []UpdatesV3RequestModulesList
		return ret
	}
	return *o.ModulesList
}

// GetModulesListOk returns a tuple with the ModulesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesV2Response) GetModulesListOk() (*[]UpdatesV3RequestModulesList, bool) {
	if o == nil || o.ModulesList == nil {
		return nil, false
	}
	return o.ModulesList, true
}

// HasModulesList returns a boolean if a field has been set.
func (o *UpdatesV2Response) HasModulesList() bool {
	if o != nil && o.ModulesList != nil {
		return true
	}

	return false
}

// SetModulesList gets a reference to the given []UpdatesV3RequestModulesList and assigns it to the ModulesList field.
func (o *UpdatesV2Response) SetModulesList(v []UpdatesV3RequestModulesList) {
	o.ModulesList = &v
}

// GetReleasever returns the Releasever field value if set, zero value otherwise.
func (o *UpdatesV2Response) GetReleasever() string {
	if o == nil || o.Releasever == nil {
		var ret string
		return ret
	}
	return *o.Releasever
}

// GetReleaseverOk returns a tuple with the Releasever field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesV2Response) GetReleaseverOk() (*string, bool) {
	if o == nil || o.Releasever == nil {
		return nil, false
	}
	return o.Releasever, true
}

// HasReleasever returns a boolean if a field has been set.
func (o *UpdatesV2Response) HasReleasever() bool {
	if o != nil && o.Releasever != nil {
		return true
	}

	return false
}

// SetReleasever gets a reference to the given string and assigns it to the Releasever field.
func (o *UpdatesV2Response) SetReleasever(v string) {
	o.Releasever = &v
}

// GetBasearch returns the Basearch field value if set, zero value otherwise.
func (o *UpdatesV2Response) GetBasearch() string {
	if o == nil || o.Basearch == nil {
		var ret string
		return ret
	}
	return *o.Basearch
}

// GetBasearchOk returns a tuple with the Basearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesV2Response) GetBasearchOk() (*string, bool) {
	if o == nil || o.Basearch == nil {
		return nil, false
	}
	return o.Basearch, true
}

// HasBasearch returns a boolean if a field has been set.
func (o *UpdatesV2Response) HasBasearch() bool {
	if o != nil && o.Basearch != nil {
		return true
	}

	return false
}

// SetBasearch gets a reference to the given string and assigns it to the Basearch field.
func (o *UpdatesV2Response) SetBasearch(v string) {
	o.Basearch = &v
}

func (o UpdatesV2Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UpdateList != nil {
		toSerialize["update_list"] = o.UpdateList
	}
	if o.RepositoryList != nil {
		toSerialize["repository_list"] = o.RepositoryList
	}
	if o.ModulesList != nil {
		toSerialize["modules_list"] = o.ModulesList
	}
	if o.Releasever != nil {
		toSerialize["releasever"] = o.Releasever
	}
	if o.Basearch != nil {
		toSerialize["basearch"] = o.Basearch
	}
	return json.Marshal(toSerialize)
}

type NullableUpdatesV2Response struct {
	value *UpdatesV2Response
	isSet bool
}

func (v NullableUpdatesV2Response) Get() *UpdatesV2Response {
	return v.value
}

func (v *NullableUpdatesV2Response) Set(val *UpdatesV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatesV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatesV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatesV2Response(val *UpdatesV2Response) *NullableUpdatesV2Response {
	return &NullableUpdatesV2Response{value: val, isSet: true}
}

func (v NullableUpdatesV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatesV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


