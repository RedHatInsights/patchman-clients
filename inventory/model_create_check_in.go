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

// CreateCheckIn Data required to create a check-in record for a host.
type CreateCheckIn struct {
	BiosUuid string `json:"bios_uuid"`
	Fqdn string `json:"fqdn"`
	InsightsId string `json:"insights_id"`
	IpAddresses []string `json:"ip_addresses"`
	MacAddresses []string `json:"mac_addresses"`
	ProviderId string `json:"provider_id"`
	ProviderType string `json:"provider_type"`
	SatelliteId string `json:"satellite_id"`
	SubscriptionManagerId string `json:"subscription_manager_id"`
	// How long from now to expect another check-in (in minutes).
	CheckinFrequency *int32 `json:"checkin_frequency,omitempty"`
}

// NewCreateCheckIn instantiates a new CreateCheckIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCheckIn(biosUuid string, fqdn string, insightsId string, ipAddresses []string, macAddresses []string, providerId string, providerType string, satelliteId string, subscriptionManagerId string, ) *CreateCheckIn {
	this := CreateCheckIn{}
	this.BiosUuid = biosUuid
	this.Fqdn = fqdn
	this.InsightsId = insightsId
	this.IpAddresses = ipAddresses
	this.MacAddresses = macAddresses
	this.ProviderId = providerId
	this.ProviderType = providerType
	this.SatelliteId = satelliteId
	this.SubscriptionManagerId = subscriptionManagerId
	return &this
}

// NewCreateCheckInWithDefaults instantiates a new CreateCheckIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCheckInWithDefaults() *CreateCheckIn {
	this := CreateCheckIn{}
	return &this
}

// GetBiosUuid returns the BiosUuid field value
func (o *CreateCheckIn) GetBiosUuid() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.BiosUuid
}

// GetBiosUuidOk returns a tuple with the BiosUuid field value
// and a boolean to check if the value has been set.
func (o *CreateCheckIn) GetBiosUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BiosUuid, true
}

// SetBiosUuid sets field value
func (o *CreateCheckIn) SetBiosUuid(v string) {
	o.BiosUuid = v
}

// GetFqdn returns the Fqdn field value
func (o *CreateCheckIn) GetFqdn() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value
// and a boolean to check if the value has been set.
func (o *CreateCheckIn) GetFqdnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Fqdn, true
}

// SetFqdn sets field value
func (o *CreateCheckIn) SetFqdn(v string) {
	o.Fqdn = v
}

// GetInsightsId returns the InsightsId field value
func (o *CreateCheckIn) GetInsightsId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.InsightsId
}

// GetInsightsIdOk returns a tuple with the InsightsId field value
// and a boolean to check if the value has been set.
func (o *CreateCheckIn) GetInsightsIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InsightsId, true
}

// SetInsightsId sets field value
func (o *CreateCheckIn) SetInsightsId(v string) {
	o.InsightsId = v
}

// GetIpAddresses returns the IpAddresses field value
func (o *CreateCheckIn) GetIpAddresses() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value
// and a boolean to check if the value has been set.
func (o *CreateCheckIn) GetIpAddressesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IpAddresses, true
}

// SetIpAddresses sets field value
func (o *CreateCheckIn) SetIpAddresses(v []string) {
	o.IpAddresses = v
}

// GetMacAddresses returns the MacAddresses field value
func (o *CreateCheckIn) GetMacAddresses() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.MacAddresses
}

// GetMacAddressesOk returns a tuple with the MacAddresses field value
// and a boolean to check if the value has been set.
func (o *CreateCheckIn) GetMacAddressesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MacAddresses, true
}

// SetMacAddresses sets field value
func (o *CreateCheckIn) SetMacAddresses(v []string) {
	o.MacAddresses = v
}

// GetProviderId returns the ProviderId field value
func (o *CreateCheckIn) GetProviderId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value
// and a boolean to check if the value has been set.
func (o *CreateCheckIn) GetProviderIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProviderId, true
}

// SetProviderId sets field value
func (o *CreateCheckIn) SetProviderId(v string) {
	o.ProviderId = v
}

// GetProviderType returns the ProviderType field value
func (o *CreateCheckIn) GetProviderType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ProviderType
}

// GetProviderTypeOk returns a tuple with the ProviderType field value
// and a boolean to check if the value has been set.
func (o *CreateCheckIn) GetProviderTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProviderType, true
}

// SetProviderType sets field value
func (o *CreateCheckIn) SetProviderType(v string) {
	o.ProviderType = v
}

// GetSatelliteId returns the SatelliteId field value
func (o *CreateCheckIn) GetSatelliteId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SatelliteId
}

// GetSatelliteIdOk returns a tuple with the SatelliteId field value
// and a boolean to check if the value has been set.
func (o *CreateCheckIn) GetSatelliteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SatelliteId, true
}

// SetSatelliteId sets field value
func (o *CreateCheckIn) SetSatelliteId(v string) {
	o.SatelliteId = v
}

// GetSubscriptionManagerId returns the SubscriptionManagerId field value
func (o *CreateCheckIn) GetSubscriptionManagerId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SubscriptionManagerId
}

// GetSubscriptionManagerIdOk returns a tuple with the SubscriptionManagerId field value
// and a boolean to check if the value has been set.
func (o *CreateCheckIn) GetSubscriptionManagerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SubscriptionManagerId, true
}

// SetSubscriptionManagerId sets field value
func (o *CreateCheckIn) SetSubscriptionManagerId(v string) {
	o.SubscriptionManagerId = v
}

// GetCheckinFrequency returns the CheckinFrequency field value if set, zero value otherwise.
func (o *CreateCheckIn) GetCheckinFrequency() int32 {
	if o == nil || o.CheckinFrequency == nil {
		var ret int32
		return ret
	}
	return *o.CheckinFrequency
}

// GetCheckinFrequencyOk returns a tuple with the CheckinFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCheckIn) GetCheckinFrequencyOk() (*int32, bool) {
	if o == nil || o.CheckinFrequency == nil {
		return nil, false
	}
	return o.CheckinFrequency, true
}

// HasCheckinFrequency returns a boolean if a field has been set.
func (o *CreateCheckIn) HasCheckinFrequency() bool {
	if o != nil && o.CheckinFrequency != nil {
		return true
	}

	return false
}

// SetCheckinFrequency gets a reference to the given int32 and assigns it to the CheckinFrequency field.
func (o *CreateCheckIn) SetCheckinFrequency(v int32) {
	o.CheckinFrequency = &v
}

func (o CreateCheckIn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bios_uuid"] = o.BiosUuid
	}
	if true {
		toSerialize["fqdn"] = o.Fqdn
	}
	if true {
		toSerialize["insights_id"] = o.InsightsId
	}
	if true {
		toSerialize["ip_addresses"] = o.IpAddresses
	}
	if true {
		toSerialize["mac_addresses"] = o.MacAddresses
	}
	if true {
		toSerialize["provider_id"] = o.ProviderId
	}
	if true {
		toSerialize["provider_type"] = o.ProviderType
	}
	if true {
		toSerialize["satellite_id"] = o.SatelliteId
	}
	if true {
		toSerialize["subscription_manager_id"] = o.SubscriptionManagerId
	}
	if o.CheckinFrequency != nil {
		toSerialize["checkin_frequency"] = o.CheckinFrequency
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCheckIn struct {
	value *CreateCheckIn
	isSet bool
}

func (v NullableCreateCheckIn) Get() *CreateCheckIn {
	return v.value
}

func (v *NullableCreateCheckIn) Set(val *CreateCheckIn) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCheckIn) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCheckIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCheckIn(val *CreateCheckIn) *NullableCreateCheckIn {
	return &NullableCreateCheckIn{value: val, isSet: true}
}

func (v NullableCreateCheckIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCheckIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


