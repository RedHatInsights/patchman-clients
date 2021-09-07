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

// UpdatesV3Request struct for UpdatesV3Request
type UpdatesV3Request struct {
	PackageList []string `json:"package_list"`
	RepositoryList *[]string `json:"repository_list,omitempty"`
	ModulesList *[]UpdatesV3RequestModulesList `json:"modules_list,omitempty"`
	Releasever *string `json:"releasever,omitempty"`
	Basearch *string `json:"basearch,omitempty"`
	SecurityOnly *bool `json:"security_only,omitempty"`
	LatestOnly *bool `json:"latest_only,omitempty"`
	// Include content from \"third party\" repositories into the response, disabled by default.
	ThirdParty *bool `json:"third_party,omitempty"`
	// Search for updates of unknown package EVRAs.
	OptimisticUpdates *bool `json:"optimistic_updates,omitempty"`
}

// NewUpdatesV3Request instantiates a new UpdatesV3Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatesV3Request(packageList []string, ) *UpdatesV3Request {
	this := UpdatesV3Request{}
	this.PackageList = packageList
	var thirdParty bool = false
	this.ThirdParty = &thirdParty
	var optimisticUpdates bool = false
	this.OptimisticUpdates = &optimisticUpdates
	return &this
}

// NewUpdatesV3RequestWithDefaults instantiates a new UpdatesV3Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatesV3RequestWithDefaults() *UpdatesV3Request {
	this := UpdatesV3Request{}
	var thirdParty bool = false
	this.ThirdParty = &thirdParty
	var optimisticUpdates bool = false
	this.OptimisticUpdates = &optimisticUpdates
	return &this
}

// GetPackageList returns the PackageList field value
func (o *UpdatesV3Request) GetPackageList() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.PackageList
}

// GetPackageListOk returns a tuple with the PackageList field value
// and a boolean to check if the value has been set.
func (o *UpdatesV3Request) GetPackageListOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PackageList, true
}

// SetPackageList sets field value
func (o *UpdatesV3Request) SetPackageList(v []string) {
	o.PackageList = v
}

// GetRepositoryList returns the RepositoryList field value if set, zero value otherwise.
func (o *UpdatesV3Request) GetRepositoryList() []string {
	if o == nil || o.RepositoryList == nil {
		var ret []string
		return ret
	}
	return *o.RepositoryList
}

// GetRepositoryListOk returns a tuple with the RepositoryList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesV3Request) GetRepositoryListOk() (*[]string, bool) {
	if o == nil || o.RepositoryList == nil {
		return nil, false
	}
	return o.RepositoryList, true
}

// HasRepositoryList returns a boolean if a field has been set.
func (o *UpdatesV3Request) HasRepositoryList() bool {
	if o != nil && o.RepositoryList != nil {
		return true
	}

	return false
}

// SetRepositoryList gets a reference to the given []string and assigns it to the RepositoryList field.
func (o *UpdatesV3Request) SetRepositoryList(v []string) {
	o.RepositoryList = &v
}

// GetModulesList returns the ModulesList field value if set, zero value otherwise.
func (o *UpdatesV3Request) GetModulesList() []UpdatesV3RequestModulesList {
	if o == nil || o.ModulesList == nil {
		var ret []UpdatesV3RequestModulesList
		return ret
	}
	return *o.ModulesList
}

// GetModulesListOk returns a tuple with the ModulesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesV3Request) GetModulesListOk() (*[]UpdatesV3RequestModulesList, bool) {
	if o == nil || o.ModulesList == nil {
		return nil, false
	}
	return o.ModulesList, true
}

// HasModulesList returns a boolean if a field has been set.
func (o *UpdatesV3Request) HasModulesList() bool {
	if o != nil && o.ModulesList != nil {
		return true
	}

	return false
}

// SetModulesList gets a reference to the given []UpdatesV3RequestModulesList and assigns it to the ModulesList field.
func (o *UpdatesV3Request) SetModulesList(v []UpdatesV3RequestModulesList) {
	o.ModulesList = &v
}

// GetReleasever returns the Releasever field value if set, zero value otherwise.
func (o *UpdatesV3Request) GetReleasever() string {
	if o == nil || o.Releasever == nil {
		var ret string
		return ret
	}
	return *o.Releasever
}

// GetReleaseverOk returns a tuple with the Releasever field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesV3Request) GetReleaseverOk() (*string, bool) {
	if o == nil || o.Releasever == nil {
		return nil, false
	}
	return o.Releasever, true
}

// HasReleasever returns a boolean if a field has been set.
func (o *UpdatesV3Request) HasReleasever() bool {
	if o != nil && o.Releasever != nil {
		return true
	}

	return false
}

// SetReleasever gets a reference to the given string and assigns it to the Releasever field.
func (o *UpdatesV3Request) SetReleasever(v string) {
	o.Releasever = &v
}

// GetBasearch returns the Basearch field value if set, zero value otherwise.
func (o *UpdatesV3Request) GetBasearch() string {
	if o == nil || o.Basearch == nil {
		var ret string
		return ret
	}
	return *o.Basearch
}

// GetBasearchOk returns a tuple with the Basearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesV3Request) GetBasearchOk() (*string, bool) {
	if o == nil || o.Basearch == nil {
		return nil, false
	}
	return o.Basearch, true
}

// HasBasearch returns a boolean if a field has been set.
func (o *UpdatesV3Request) HasBasearch() bool {
	if o != nil && o.Basearch != nil {
		return true
	}

	return false
}

// SetBasearch gets a reference to the given string and assigns it to the Basearch field.
func (o *UpdatesV3Request) SetBasearch(v string) {
	o.Basearch = &v
}

// GetSecurityOnly returns the SecurityOnly field value if set, zero value otherwise.
func (o *UpdatesV3Request) GetSecurityOnly() bool {
	if o == nil || o.SecurityOnly == nil {
		var ret bool
		return ret
	}
	return *o.SecurityOnly
}

// GetSecurityOnlyOk returns a tuple with the SecurityOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesV3Request) GetSecurityOnlyOk() (*bool, bool) {
	if o == nil || o.SecurityOnly == nil {
		return nil, false
	}
	return o.SecurityOnly, true
}

// HasSecurityOnly returns a boolean if a field has been set.
func (o *UpdatesV3Request) HasSecurityOnly() bool {
	if o != nil && o.SecurityOnly != nil {
		return true
	}

	return false
}

// SetSecurityOnly gets a reference to the given bool and assigns it to the SecurityOnly field.
func (o *UpdatesV3Request) SetSecurityOnly(v bool) {
	o.SecurityOnly = &v
}

// GetLatestOnly returns the LatestOnly field value if set, zero value otherwise.
func (o *UpdatesV3Request) GetLatestOnly() bool {
	if o == nil || o.LatestOnly == nil {
		var ret bool
		return ret
	}
	return *o.LatestOnly
}

// GetLatestOnlyOk returns a tuple with the LatestOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesV3Request) GetLatestOnlyOk() (*bool, bool) {
	if o == nil || o.LatestOnly == nil {
		return nil, false
	}
	return o.LatestOnly, true
}

// HasLatestOnly returns a boolean if a field has been set.
func (o *UpdatesV3Request) HasLatestOnly() bool {
	if o != nil && o.LatestOnly != nil {
		return true
	}

	return false
}

// SetLatestOnly gets a reference to the given bool and assigns it to the LatestOnly field.
func (o *UpdatesV3Request) SetLatestOnly(v bool) {
	o.LatestOnly = &v
}

// GetThirdParty returns the ThirdParty field value if set, zero value otherwise.
func (o *UpdatesV3Request) GetThirdParty() bool {
	if o == nil || o.ThirdParty == nil {
		var ret bool
		return ret
	}
	return *o.ThirdParty
}

// GetThirdPartyOk returns a tuple with the ThirdParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesV3Request) GetThirdPartyOk() (*bool, bool) {
	if o == nil || o.ThirdParty == nil {
		return nil, false
	}
	return o.ThirdParty, true
}

// HasThirdParty returns a boolean if a field has been set.
func (o *UpdatesV3Request) HasThirdParty() bool {
	if o != nil && o.ThirdParty != nil {
		return true
	}

	return false
}

// SetThirdParty gets a reference to the given bool and assigns it to the ThirdParty field.
func (o *UpdatesV3Request) SetThirdParty(v bool) {
	o.ThirdParty = &v
}

// GetOptimisticUpdates returns the OptimisticUpdates field value if set, zero value otherwise.
func (o *UpdatesV3Request) GetOptimisticUpdates() bool {
	if o == nil || o.OptimisticUpdates == nil {
		var ret bool
		return ret
	}
	return *o.OptimisticUpdates
}

// GetOptimisticUpdatesOk returns a tuple with the OptimisticUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesV3Request) GetOptimisticUpdatesOk() (*bool, bool) {
	if o == nil || o.OptimisticUpdates == nil {
		return nil, false
	}
	return o.OptimisticUpdates, true
}

// HasOptimisticUpdates returns a boolean if a field has been set.
func (o *UpdatesV3Request) HasOptimisticUpdates() bool {
	if o != nil && o.OptimisticUpdates != nil {
		return true
	}

	return false
}

// SetOptimisticUpdates gets a reference to the given bool and assigns it to the OptimisticUpdates field.
func (o *UpdatesV3Request) SetOptimisticUpdates(v bool) {
	o.OptimisticUpdates = &v
}

func (o UpdatesV3Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["package_list"] = o.PackageList
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
	if o.SecurityOnly != nil {
		toSerialize["security_only"] = o.SecurityOnly
	}
	if o.LatestOnly != nil {
		toSerialize["latest_only"] = o.LatestOnly
	}
	if o.ThirdParty != nil {
		toSerialize["third_party"] = o.ThirdParty
	}
	if o.OptimisticUpdates != nil {
		toSerialize["optimistic_updates"] = o.OptimisticUpdates
	}
	return json.Marshal(toSerialize)
}

type NullableUpdatesV3Request struct {
	value *UpdatesV3Request
	isSet bool
}

func (v NullableUpdatesV3Request) Get() *UpdatesV3Request {
	return v.value
}

func (v *NullableUpdatesV3Request) Set(val *UpdatesV3Request) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatesV3Request) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatesV3Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatesV3Request(val *UpdatesV3Request) *NullableUpdatesV3Request {
	return &NullableUpdatesV3Request{value: val, isSet: true}
}

func (v NullableUpdatesV3Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatesV3Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


