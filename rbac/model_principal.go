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

// Principal struct for Principal
type Principal struct {
	Username string `json:"username"`
	Email string `json:"email"`
	FirstName *string `json:"first_name,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	IsOrgAdmin *bool `json:"is_org_admin,omitempty"`
}

// NewPrincipal instantiates a new Principal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipal(username string, email string, ) *Principal {
	this := Principal{}
	this.Username = username
	this.Email = email
	return &this
}

// NewPrincipalWithDefaults instantiates a new Principal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalWithDefaults() *Principal {
	this := Principal{}
	return &this
}

// GetUsername returns the Username field value
func (o *Principal) GetUsername() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *Principal) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *Principal) SetUsername(v string) {
	o.Username = v
}

// GetEmail returns the Email field value
func (o *Principal) GetEmail() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *Principal) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *Principal) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *Principal) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Principal) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *Principal) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *Principal) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *Principal) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Principal) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *Principal) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *Principal) SetLastName(v string) {
	o.LastName = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *Principal) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Principal) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *Principal) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *Principal) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsOrgAdmin returns the IsOrgAdmin field value if set, zero value otherwise.
func (o *Principal) GetIsOrgAdmin() bool {
	if o == nil || o.IsOrgAdmin == nil {
		var ret bool
		return ret
	}
	return *o.IsOrgAdmin
}

// GetIsOrgAdminOk returns a tuple with the IsOrgAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Principal) GetIsOrgAdminOk() (*bool, bool) {
	if o == nil || o.IsOrgAdmin == nil {
		return nil, false
	}
	return o.IsOrgAdmin, true
}

// HasIsOrgAdmin returns a boolean if a field has been set.
func (o *Principal) HasIsOrgAdmin() bool {
	if o != nil && o.IsOrgAdmin != nil {
		return true
	}

	return false
}

// SetIsOrgAdmin gets a reference to the given bool and assigns it to the IsOrgAdmin field.
func (o *Principal) SetIsOrgAdmin(v bool) {
	o.IsOrgAdmin = &v
}

func (o Principal) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullablePrincipal struct {
	value *Principal
	isSet bool
}

func (v NullablePrincipal) Get() *Principal {
	return v.value
}

func (v *NullablePrincipal) Set(val *Principal) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipal) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipal(val *Principal) *NullablePrincipal {
	return &NullablePrincipal{value: val, isSet: true}
}

func (v NullablePrincipal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


