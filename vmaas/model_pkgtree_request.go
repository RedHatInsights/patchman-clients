/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.14.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vmaas

import (
	"encoding/json"
)

// PkgtreeRequest struct for PkgtreeRequest
type PkgtreeRequest struct {
	Page *float32 `json:"page,omitempty"`
	PageSize *float32 `json:"page_size,omitempty"`
	PackageNameList []string `json:"package_name_list"`
	ModifiedSince *string `json:"modified_since,omitempty"`
	// Include content from \"third party\" repositories into the response, disabled by default.
	ThirdParty *bool `json:"third_party,omitempty"`
	// Include nevra repositories info into the response.
	ReturnRepositories *bool `json:"return_repositories,omitempty"`
	// Include nevra errata info into the response.
	ReturnErrata *bool `json:"return_errata,omitempty"`
	// Include nevra summary info into the response.
	ReturnSummary *bool `json:"return_summary,omitempty"`
	// Include nevra description info into the response.
	ReturnDescription *bool `json:"return_description,omitempty"`
}

// NewPkgtreeRequest instantiates a new PkgtreeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkgtreeRequest(packageNameList []string, ) *PkgtreeRequest {
	this := PkgtreeRequest{}
	this.PackageNameList = packageNameList
	var thirdParty bool = false
	this.ThirdParty = &thirdParty
	var returnRepositories bool = true
	this.ReturnRepositories = &returnRepositories
	var returnErrata bool = true
	this.ReturnErrata = &returnErrata
	var returnSummary bool = false
	this.ReturnSummary = &returnSummary
	var returnDescription bool = false
	this.ReturnDescription = &returnDescription
	return &this
}

// NewPkgtreeRequestWithDefaults instantiates a new PkgtreeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkgtreeRequestWithDefaults() *PkgtreeRequest {
	this := PkgtreeRequest{}
	var thirdParty bool = false
	this.ThirdParty = &thirdParty
	var returnRepositories bool = true
	this.ReturnRepositories = &returnRepositories
	var returnErrata bool = true
	this.ReturnErrata = &returnErrata
	var returnSummary bool = false
	this.ReturnSummary = &returnSummary
	var returnDescription bool = false
	this.ReturnDescription = &returnDescription
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *PkgtreeRequest) GetPage() float32 {
	if o == nil || o.Page == nil {
		var ret float32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgtreeRequest) GetPageOk() (*float32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *PkgtreeRequest) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given float32 and assigns it to the Page field.
func (o *PkgtreeRequest) SetPage(v float32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *PkgtreeRequest) GetPageSize() float32 {
	if o == nil || o.PageSize == nil {
		var ret float32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgtreeRequest) GetPageSizeOk() (*float32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *PkgtreeRequest) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given float32 and assigns it to the PageSize field.
func (o *PkgtreeRequest) SetPageSize(v float32) {
	o.PageSize = &v
}

// GetPackageNameList returns the PackageNameList field value
func (o *PkgtreeRequest) GetPackageNameList() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.PackageNameList
}

// GetPackageNameListOk returns a tuple with the PackageNameList field value
// and a boolean to check if the value has been set.
func (o *PkgtreeRequest) GetPackageNameListOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PackageNameList, true
}

// SetPackageNameList sets field value
func (o *PkgtreeRequest) SetPackageNameList(v []string) {
	o.PackageNameList = v
}

// GetModifiedSince returns the ModifiedSince field value if set, zero value otherwise.
func (o *PkgtreeRequest) GetModifiedSince() string {
	if o == nil || o.ModifiedSince == nil {
		var ret string
		return ret
	}
	return *o.ModifiedSince
}

// GetModifiedSinceOk returns a tuple with the ModifiedSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgtreeRequest) GetModifiedSinceOk() (*string, bool) {
	if o == nil || o.ModifiedSince == nil {
		return nil, false
	}
	return o.ModifiedSince, true
}

// HasModifiedSince returns a boolean if a field has been set.
func (o *PkgtreeRequest) HasModifiedSince() bool {
	if o != nil && o.ModifiedSince != nil {
		return true
	}

	return false
}

// SetModifiedSince gets a reference to the given string and assigns it to the ModifiedSince field.
func (o *PkgtreeRequest) SetModifiedSince(v string) {
	o.ModifiedSince = &v
}

// GetThirdParty returns the ThirdParty field value if set, zero value otherwise.
func (o *PkgtreeRequest) GetThirdParty() bool {
	if o == nil || o.ThirdParty == nil {
		var ret bool
		return ret
	}
	return *o.ThirdParty
}

// GetThirdPartyOk returns a tuple with the ThirdParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgtreeRequest) GetThirdPartyOk() (*bool, bool) {
	if o == nil || o.ThirdParty == nil {
		return nil, false
	}
	return o.ThirdParty, true
}

// HasThirdParty returns a boolean if a field has been set.
func (o *PkgtreeRequest) HasThirdParty() bool {
	if o != nil && o.ThirdParty != nil {
		return true
	}

	return false
}

// SetThirdParty gets a reference to the given bool and assigns it to the ThirdParty field.
func (o *PkgtreeRequest) SetThirdParty(v bool) {
	o.ThirdParty = &v
}

// GetReturnRepositories returns the ReturnRepositories field value if set, zero value otherwise.
func (o *PkgtreeRequest) GetReturnRepositories() bool {
	if o == nil || o.ReturnRepositories == nil {
		var ret bool
		return ret
	}
	return *o.ReturnRepositories
}

// GetReturnRepositoriesOk returns a tuple with the ReturnRepositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgtreeRequest) GetReturnRepositoriesOk() (*bool, bool) {
	if o == nil || o.ReturnRepositories == nil {
		return nil, false
	}
	return o.ReturnRepositories, true
}

// HasReturnRepositories returns a boolean if a field has been set.
func (o *PkgtreeRequest) HasReturnRepositories() bool {
	if o != nil && o.ReturnRepositories != nil {
		return true
	}

	return false
}

// SetReturnRepositories gets a reference to the given bool and assigns it to the ReturnRepositories field.
func (o *PkgtreeRequest) SetReturnRepositories(v bool) {
	o.ReturnRepositories = &v
}

// GetReturnErrata returns the ReturnErrata field value if set, zero value otherwise.
func (o *PkgtreeRequest) GetReturnErrata() bool {
	if o == nil || o.ReturnErrata == nil {
		var ret bool
		return ret
	}
	return *o.ReturnErrata
}

// GetReturnErrataOk returns a tuple with the ReturnErrata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgtreeRequest) GetReturnErrataOk() (*bool, bool) {
	if o == nil || o.ReturnErrata == nil {
		return nil, false
	}
	return o.ReturnErrata, true
}

// HasReturnErrata returns a boolean if a field has been set.
func (o *PkgtreeRequest) HasReturnErrata() bool {
	if o != nil && o.ReturnErrata != nil {
		return true
	}

	return false
}

// SetReturnErrata gets a reference to the given bool and assigns it to the ReturnErrata field.
func (o *PkgtreeRequest) SetReturnErrata(v bool) {
	o.ReturnErrata = &v
}

// GetReturnSummary returns the ReturnSummary field value if set, zero value otherwise.
func (o *PkgtreeRequest) GetReturnSummary() bool {
	if o == nil || o.ReturnSummary == nil {
		var ret bool
		return ret
	}
	return *o.ReturnSummary
}

// GetReturnSummaryOk returns a tuple with the ReturnSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgtreeRequest) GetReturnSummaryOk() (*bool, bool) {
	if o == nil || o.ReturnSummary == nil {
		return nil, false
	}
	return o.ReturnSummary, true
}

// HasReturnSummary returns a boolean if a field has been set.
func (o *PkgtreeRequest) HasReturnSummary() bool {
	if o != nil && o.ReturnSummary != nil {
		return true
	}

	return false
}

// SetReturnSummary gets a reference to the given bool and assigns it to the ReturnSummary field.
func (o *PkgtreeRequest) SetReturnSummary(v bool) {
	o.ReturnSummary = &v
}

// GetReturnDescription returns the ReturnDescription field value if set, zero value otherwise.
func (o *PkgtreeRequest) GetReturnDescription() bool {
	if o == nil || o.ReturnDescription == nil {
		var ret bool
		return ret
	}
	return *o.ReturnDescription
}

// GetReturnDescriptionOk returns a tuple with the ReturnDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgtreeRequest) GetReturnDescriptionOk() (*bool, bool) {
	if o == nil || o.ReturnDescription == nil {
		return nil, false
	}
	return o.ReturnDescription, true
}

// HasReturnDescription returns a boolean if a field has been set.
func (o *PkgtreeRequest) HasReturnDescription() bool {
	if o != nil && o.ReturnDescription != nil {
		return true
	}

	return false
}

// SetReturnDescription gets a reference to the given bool and assigns it to the ReturnDescription field.
func (o *PkgtreeRequest) SetReturnDescription(v bool) {
	o.ReturnDescription = &v
}

func (o PkgtreeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.PageSize != nil {
		toSerialize["page_size"] = o.PageSize
	}
	if true {
		toSerialize["package_name_list"] = o.PackageNameList
	}
	if o.ModifiedSince != nil {
		toSerialize["modified_since"] = o.ModifiedSince
	}
	if o.ThirdParty != nil {
		toSerialize["third_party"] = o.ThirdParty
	}
	if o.ReturnRepositories != nil {
		toSerialize["return_repositories"] = o.ReturnRepositories
	}
	if o.ReturnErrata != nil {
		toSerialize["return_errata"] = o.ReturnErrata
	}
	if o.ReturnSummary != nil {
		toSerialize["return_summary"] = o.ReturnSummary
	}
	if o.ReturnDescription != nil {
		toSerialize["return_description"] = o.ReturnDescription
	}
	return json.Marshal(toSerialize)
}

type NullablePkgtreeRequest struct {
	value *PkgtreeRequest
	isSet bool
}

func (v NullablePkgtreeRequest) Get() *PkgtreeRequest {
	return v.value
}

func (v *NullablePkgtreeRequest) Set(val *PkgtreeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePkgtreeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePkgtreeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkgtreeRequest(val *PkgtreeRequest) *NullablePkgtreeRequest {
	return &NullablePkgtreeRequest{value: val, isSet: true}
}

func (v NullablePkgtreeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkgtreeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


