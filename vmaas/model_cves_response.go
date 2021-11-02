/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.32.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vmaas

import (
	"encoding/json"
)

// CvesResponse struct for CvesResponse
type CvesResponse struct {
	Page *float32 `json:"page,omitempty"`
	PageSize *float32 `json:"page_size,omitempty"`
	Pages *float32 `json:"pages,omitempty"`
	CveList *map[string]CvesResponseCveList `json:"cve_list,omitempty"`
	LastChange *string `json:"last_change,omitempty"`
}

// NewCvesResponse instantiates a new CvesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCvesResponse() *CvesResponse {
	this := CvesResponse{}
	return &this
}

// NewCvesResponseWithDefaults instantiates a new CvesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCvesResponseWithDefaults() *CvesResponse {
	this := CvesResponse{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *CvesResponse) GetPage() float32 {
	if o == nil || o.Page == nil {
		var ret float32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesResponse) GetPageOk() (*float32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *CvesResponse) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given float32 and assigns it to the Page field.
func (o *CvesResponse) SetPage(v float32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *CvesResponse) GetPageSize() float32 {
	if o == nil || o.PageSize == nil {
		var ret float32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesResponse) GetPageSizeOk() (*float32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *CvesResponse) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given float32 and assigns it to the PageSize field.
func (o *CvesResponse) SetPageSize(v float32) {
	o.PageSize = &v
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *CvesResponse) GetPages() float32 {
	if o == nil || o.Pages == nil {
		var ret float32
		return ret
	}
	return *o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesResponse) GetPagesOk() (*float32, bool) {
	if o == nil || o.Pages == nil {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *CvesResponse) HasPages() bool {
	if o != nil && o.Pages != nil {
		return true
	}

	return false
}

// SetPages gets a reference to the given float32 and assigns it to the Pages field.
func (o *CvesResponse) SetPages(v float32) {
	o.Pages = &v
}

// GetCveList returns the CveList field value if set, zero value otherwise.
func (o *CvesResponse) GetCveList() map[string]CvesResponseCveList {
	if o == nil || o.CveList == nil {
		var ret map[string]CvesResponseCveList
		return ret
	}
	return *o.CveList
}

// GetCveListOk returns a tuple with the CveList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesResponse) GetCveListOk() (*map[string]CvesResponseCveList, bool) {
	if o == nil || o.CveList == nil {
		return nil, false
	}
	return o.CveList, true
}

// HasCveList returns a boolean if a field has been set.
func (o *CvesResponse) HasCveList() bool {
	if o != nil && o.CveList != nil {
		return true
	}

	return false
}

// SetCveList gets a reference to the given map[string]CvesResponseCveList and assigns it to the CveList field.
func (o *CvesResponse) SetCveList(v map[string]CvesResponseCveList) {
	o.CveList = &v
}

// GetLastChange returns the LastChange field value if set, zero value otherwise.
func (o *CvesResponse) GetLastChange() string {
	if o == nil || o.LastChange == nil {
		var ret string
		return ret
	}
	return *o.LastChange
}

// GetLastChangeOk returns a tuple with the LastChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesResponse) GetLastChangeOk() (*string, bool) {
	if o == nil || o.LastChange == nil {
		return nil, false
	}
	return o.LastChange, true
}

// HasLastChange returns a boolean if a field has been set.
func (o *CvesResponse) HasLastChange() bool {
	if o != nil && o.LastChange != nil {
		return true
	}

	return false
}

// SetLastChange gets a reference to the given string and assigns it to the LastChange field.
func (o *CvesResponse) SetLastChange(v string) {
	o.LastChange = &v
}

func (o CvesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.PageSize != nil {
		toSerialize["page_size"] = o.PageSize
	}
	if o.Pages != nil {
		toSerialize["pages"] = o.Pages
	}
	if o.CveList != nil {
		toSerialize["cve_list"] = o.CveList
	}
	if o.LastChange != nil {
		toSerialize["last_change"] = o.LastChange
	}
	return json.Marshal(toSerialize)
}

type NullableCvesResponse struct {
	value *CvesResponse
	isSet bool
}

func (v NullableCvesResponse) Get() *CvesResponse {
	return v.value
}

func (v *NullableCvesResponse) Set(val *CvesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCvesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCvesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCvesResponse(val *CvesResponse) *NullableCvesResponse {
	return &NullableCvesResponse{value: val, isSet: true}
}

func (v NullableCvesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCvesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


