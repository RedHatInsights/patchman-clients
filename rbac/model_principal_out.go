/*
 * Role Based Access Control
 *
 * The API for Role Based Access Control.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rbac

import (
	"encoding/json"
)

// PrincipalOut struct for PrincipalOut
type PrincipalOut struct {
	Username string `json:"username"`
	Email string `json:"email"`
	FirstName *string `json:"first_name,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	IsOrgAdmin *bool `json:"is_org_admin,omitempty"`
	Uuid string `json:"uuid"`
}

// NewPrincipalOut instantiates a new PrincipalOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalOut(username string, email string, uuid string, ) *PrincipalOut {
	this := PrincipalOut{}
	this.Username = username
	this.Email = email
	this.Uuid = uuid
	return &this
}

// NewPrincipalOutWithDefaults instantiates a new PrincipalOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalOutWithDefaults() *PrincipalOut {
	this := PrincipalOut{}
	return &this
}

// GetUsername returns the Username field value
func (o *PrincipalOut) GetUsername() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *PrincipalOut) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *PrincipalOut) SetUsername(v string) {
	o.Username = v
}

// GetEmail returns the Email field value
func (o *PrincipalOut) GetEmail() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *PrincipalOut) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *PrincipalOut) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *PrincipalOut) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalOut) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *PrincipalOut) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *PrincipalOut) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *PrincipalOut) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalOut) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *PrincipalOut) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *PrincipalOut) SetLastName(v string) {
	o.LastName = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *PrincipalOut) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalOut) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *PrincipalOut) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *PrincipalOut) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsOrgAdmin returns the IsOrgAdmin field value if set, zero value otherwise.
func (o *PrincipalOut) GetIsOrgAdmin() bool {
	if o == nil || o.IsOrgAdmin == nil {
		var ret bool
		return ret
	}
	return *o.IsOrgAdmin
}

// GetIsOrgAdminOk returns a tuple with the IsOrgAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalOut) GetIsOrgAdminOk() (*bool, bool) {
	if o == nil || o.IsOrgAdmin == nil {
		return nil, false
	}
	return o.IsOrgAdmin, true
}

// HasIsOrgAdmin returns a boolean if a field has been set.
func (o *PrincipalOut) HasIsOrgAdmin() bool {
	if o != nil && o.IsOrgAdmin != nil {
		return true
	}

	return false
}

// SetIsOrgAdmin gets a reference to the given bool and assigns it to the IsOrgAdmin field.
func (o *PrincipalOut) SetIsOrgAdmin(v bool) {
	o.IsOrgAdmin = &v
}

// GetUuid returns the Uuid field value
func (o *PrincipalOut) GetUuid() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *PrincipalOut) GetUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *PrincipalOut) SetUuid(v string) {
	o.Uuid = v
}

func (o PrincipalOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if o.IsActive != nil {
		toSerialize["is_active"] = o.IsActive
	}
	if o.IsOrgAdmin != nil {
		toSerialize["is_org_admin"] = o.IsOrgAdmin
	}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullablePrincipalOut struct {
	value *PrincipalOut
	isSet bool
}

func (v NullablePrincipalOut) Get() *PrincipalOut {
	return v.value
}

func (v *NullablePrincipalOut) Set(val *PrincipalOut) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalOut) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalOut(val *PrincipalOut) *NullablePrincipalOut {
	return &NullablePrincipalOut{value: val, isSet: true}
}

func (v NullablePrincipalOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


