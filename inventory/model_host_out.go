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

// HostOut A database entry representing a single host with its Inventory metadata.
type HostOut struct {
	// A set of facts belonging to the host.
	Facts *[]map[string]interface{} `json:"facts,omitempty"`
	// A UUID of the host machine BIOS.  This field is considered to be a canonical fact.
	BiosUuid NullableString `json:"bios_uuid,omitempty"`
	// Host’s reference in the external source e.g. AWS EC2, Azure, OpenStack, etc. This field is considered to be a canonical fact.
	ExternalId NullableString `json:"external_id,omitempty"`
	// A host’s Fully Qualified Domain Name.  This field is considered to be a canonical fact.
	Fqdn NullableString `json:"fqdn,omitempty"`
	// An ID defined in /etc/insights-client/machine-id. This field is considered a canonical fact.
	InsightsId NullableString `json:"insights_id,omitempty"`
	// Host’s network IP addresses.  This field is considered to be a canonical fact.
	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Host’s network interfaces MAC addresses.  This field is considered to be a canonical fact.
	MacAddresses []string `json:"mac_addresses,omitempty"`
	// A Machine ID of a RHEL host.  This field is considered to be a canonical fact.
	RhelMachineId NullableString `json:"rhel_machine_id,omitempty"`
	// A Red Hat Satellite ID of a RHEL host.  This field is considered to be a canonical fact.
	SatelliteId NullableString `json:"satellite_id,omitempty"`
	// A Red Hat Subcription Manager ID of a RHEL host.  This field is considered to be a canonical fact.
	SubscriptionManagerId NullableString `json:"subscription_manager_id,omitempty"`
	// A Red Hat Account number that owns the host.
	Account string `json:"account"`
	// The ansible host name for remediations
	AnsibleHost NullableString `json:"ansible_host,omitempty"`
	// A timestamp when the entry was created.
	Created *string `json:"created,omitempty"`
	// Timestamp from which the host is considered deleted.
	CulledTimestamp NullableTime `json:"culled_timestamp,omitempty"`
	// A host’s human-readable display name, e.g. in a form of a domain name.
	DisplayName NullableString `json:"display_name,omitempty"`
	// A durable and reliable platform-wide host identifier. Applications should use this identifier to reference hosts.
	Id *string `json:"id,omitempty"`
	// Reporting source of the host. Used when updating the stale_timestamp.
	Reporter NullableString `json:"reporter,omitempty"`
	// Timestamp from which the host is considered stale.
	StaleTimestamp NullableTime `json:"stale_timestamp,omitempty"`
	// Timestamp from which the host is considered too stale to be listed without an explicit toggle.
	StaleWarningTimestamp NullableTime `json:"stale_warning_timestamp,omitempty"`
	// A timestamp when the entry was last updated.
	Updated *string `json:"updated,omitempty"`
}

// NewHostOut instantiates a new HostOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostOut(account string, ) *HostOut {
	this := HostOut{}
	this.Account = account
	return &this
}

// NewHostOutWithDefaults instantiates a new HostOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostOutWithDefaults() *HostOut {
	this := HostOut{}
	return &this
}

// GetFacts returns the Facts field value if set, zero value otherwise.
func (o *HostOut) GetFacts() []map[string]interface{} {
	if o == nil || o.Facts == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Facts
}

// GetFactsOk returns a tuple with the Facts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostOut) GetFactsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Facts == nil {
		return nil, false
	}
	return o.Facts, true
}

// HasFacts returns a boolean if a field has been set.
func (o *HostOut) HasFacts() bool {
	if o != nil && o.Facts != nil {
		return true
	}

	return false
}

// SetFacts gets a reference to the given []map[string]interface{} and assigns it to the Facts field.
func (o *HostOut) SetFacts(v []map[string]interface{}) {
	o.Facts = &v
}

// GetBiosUuid returns the BiosUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HostOut) GetBiosUuid() string {
	if o == nil || o.BiosUuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.BiosUuid.Get()
}

// GetBiosUuidOk returns a tuple with the BiosUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostOut) GetBiosUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BiosUuid.Get(), o.BiosUuid.IsSet()
}

// HasBiosUuid returns a boolean if a field has been set.
func (o *HostOut) HasBiosUuid() bool {
	if o != nil && o.BiosUuid.IsSet() {
		return true
	}

	return false
}

// SetBiosUuid gets a reference to the given NullableString and assigns it to the BiosUuid field.
func (o *HostOut) SetBiosUuid(v string) {
	o.BiosUuid.Set(&v)
}
// SetBiosUuidNil sets the value for BiosUuid to be an explicit nil
func (o *HostOut) SetBiosUuidNil() {
	o.BiosUuid.Set(nil)
}

// UnsetBiosUuid ensures that no value is present for BiosUuid, not even an explicit nil
func (o *HostOut) UnsetBiosUuid() {
	o.BiosUuid.Unset()
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HostOut) GetExternalId() string {
	if o == nil || o.ExternalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostOut) GetExternalIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *HostOut) HasExternalId() bool {
	if o != nil && o.ExternalId.IsSet() {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given NullableString and assigns it to the ExternalId field.
func (o *HostOut) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}
// SetExternalIdNil sets the value for ExternalId to be an explicit nil
func (o *HostOut) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
func (o *HostOut) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HostOut) GetFqdn() string {
	if o == nil || o.Fqdn.Get() == nil {
		var ret string
		return ret
	}
	return *o.Fqdn.Get()
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostOut) GetFqdnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Fqdn.Get(), o.Fqdn.IsSet()
}

// HasFqdn returns a boolean if a field has been set.
func (o *HostOut) HasFqdn() bool {
	if o != nil && o.Fqdn.IsSet() {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given NullableString and assigns it to the Fqdn field.
func (o *HostOut) SetFqdn(v string) {
	o.Fqdn.Set(&v)
}
// SetFqdnNil sets the value for Fqdn to be an explicit nil
func (o *HostOut) SetFqdnNil() {
	o.Fqdn.Set(nil)
}

// UnsetFqdn ensures that no value is present for Fqdn, not even an explicit nil
func (o *HostOut) UnsetFqdn() {
	o.Fqdn.Unset()
}

// GetInsightsId returns the InsightsId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HostOut) GetInsightsId() string {
	if o == nil || o.InsightsId.Get() == nil {
		var ret string
		return ret
	}
	return *o.InsightsId.Get()
}

// GetInsightsIdOk returns a tuple with the InsightsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostOut) GetInsightsIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InsightsId.Get(), o.InsightsId.IsSet()
}

// HasInsightsId returns a boolean if a field has been set.
func (o *HostOut) HasInsightsId() bool {
	if o != nil && o.InsightsId.IsSet() {
		return true
	}

	return false
}

// SetInsightsId gets a reference to the given NullableString and assigns it to the InsightsId field.
func (o *HostOut) SetInsightsId(v string) {
	o.InsightsId.Set(&v)
}
// SetInsightsIdNil sets the value for InsightsId to be an explicit nil
func (o *HostOut) SetInsightsIdNil() {
	o.InsightsId.Set(nil)
}

// UnsetInsightsId ensures that no value is present for InsightsId, not even an explicit nil
func (o *HostOut) UnsetInsightsId() {
	o.InsightsId.Unset()
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HostOut) GetIpAddresses() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostOut) GetIpAddressesOk() (*[]string, bool) {
	if o == nil || o.IpAddresses == nil {
		return nil, false
	}
	return &o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *HostOut) HasIpAddresses() bool {
	if o != nil && o.IpAddresses != nil {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *HostOut) SetIpAddresses(v []string) {
	o.IpAddresses = v
}

// GetMacAddresses returns the MacAddresses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HostOut) GetMacAddresses() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.MacAddresses
}

// GetMacAddressesOk returns a tuple with the MacAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostOut) GetMacAddressesOk() (*[]string, bool) {
	if o == nil || o.MacAddresses == nil {
		return nil, false
	}
	return &o.MacAddresses, true
}

// HasMacAddresses returns a boolean if a field has been set.
func (o *HostOut) HasMacAddresses() bool {
	if o != nil && o.MacAddresses != nil {
		return true
	}

	return false
}

// SetMacAddresses gets a reference to the given []string and assigns it to the MacAddresses field.
func (o *HostOut) SetMacAddresses(v []string) {
	o.MacAddresses = v
}

// GetRhelMachineId returns the RhelMachineId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HostOut) GetRhelMachineId() string {
	if o == nil || o.RhelMachineId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RhelMachineId.Get()
}

// GetRhelMachineIdOk returns a tuple with the RhelMachineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostOut) GetRhelMachineIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RhelMachineId.Get(), o.RhelMachineId.IsSet()
}

// HasRhelMachineId returns a boolean if a field has been set.
func (o *HostOut) HasRhelMachineId() bool {
	if o != nil && o.RhelMachineId.IsSet() {
		return true
	}

	return false
}

// SetRhelMachineId gets a reference to the given NullableString and assigns it to the RhelMachineId field.
func (o *HostOut) SetRhelMachineId(v string) {
	o.RhelMachineId.Set(&v)
}
// SetRhelMachineIdNil sets the value for RhelMachineId to be an explicit nil
func (o *HostOut) SetRhelMachineIdNil() {
	o.RhelMachineId.Set(nil)
}

// UnsetRhelMachineId ensures that no value is present for RhelMachineId, not even an explicit nil
func (o *HostOut) UnsetRhelMachineId() {
	o.RhelMachineId.Unset()
}

// GetSatelliteId returns the SatelliteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HostOut) GetSatelliteId() string {
	if o == nil || o.SatelliteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SatelliteId.Get()
}

// GetSatelliteIdOk returns a tuple with the SatelliteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostOut) GetSatelliteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SatelliteId.Get(), o.SatelliteId.IsSet()
}

// HasSatelliteId returns a boolean if a field has been set.
func (o *HostOut) HasSatelliteId() bool {
	if o != nil && o.SatelliteId.IsSet() {
		return true
	}

	return false
}

// SetSatelliteId gets a reference to the given NullableString and assigns it to the SatelliteId field.
func (o *HostOut) SetSatelliteId(v string) {
	o.SatelliteId.Set(&v)
}
// SetSatelliteIdNil sets the value for SatelliteId to be an explicit nil
func (o *HostOut) SetSatelliteIdNil() {
	o.SatelliteId.Set(nil)
}

// UnsetSatelliteId ensures that no value is present for SatelliteId, not even an explicit nil
func (o *HostOut) UnsetSatelliteId() {
	o.SatelliteId.Unset()
}

// GetSubscriptionManagerId returns the SubscriptionManagerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HostOut) GetSubscriptionManagerId() string {
	if o == nil || o.SubscriptionManagerId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionManagerId.Get()
}

// GetSubscriptionManagerIdOk returns a tuple with the SubscriptionManagerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostOut) GetSubscriptionManagerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubscriptionManagerId.Get(), o.SubscriptionManagerId.IsSet()
}

// HasSubscriptionManagerId returns a boolean if a field has been set.
func (o *HostOut) HasSubscriptionManagerId() bool {
	if o != nil && o.SubscriptionManagerId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionManagerId gets a reference to the given NullableString and assigns it to the SubscriptionManagerId field.
func (o *HostOut) SetSubscriptionManagerId(v string) {
	o.SubscriptionManagerId.Set(&v)
}
// SetSubscriptionManagerIdNil sets the value for SubscriptionManagerId to be an explicit nil
func (o *HostOut) SetSubscriptionManagerIdNil() {
	o.SubscriptionManagerId.Set(nil)
}

// UnsetSubscriptionManagerId ensures that no value is present for SubscriptionManagerId, not even an explicit nil
func (o *HostOut) UnsetSubscriptionManagerId() {
	o.SubscriptionManagerId.Unset()
}

// GetAccount returns the Account field value
func (o *HostOut) GetAccount() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *HostOut) GetAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *HostOut) SetAccount(v string) {
	o.Account = v
}

// GetAnsibleHost returns the AnsibleHost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HostOut) GetAnsibleHost() string {
	if o == nil || o.AnsibleHost.Get() == nil {
		var ret string
		return ret
	}
	return *o.AnsibleHost.Get()
}

// GetAnsibleHostOk returns a tuple with the AnsibleHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostOut) GetAnsibleHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AnsibleHost.Get(), o.AnsibleHost.IsSet()
}

// HasAnsibleHost returns a boolean if a field has been set.
func (o *HostOut) HasAnsibleHost() bool {
	if o != nil && o.AnsibleHost.IsSet() {
		return true
	}

	return false
}

// SetAnsibleHost gets a reference to the given NullableString and assigns it to the AnsibleHost field.
func (o *HostOut) SetAnsibleHost(v string) {
	o.AnsibleHost.Set(&v)
}
// SetAnsibleHostNil sets the value for AnsibleHost to be an explicit nil
func (o *HostOut) SetAnsibleHostNil() {
	o.AnsibleHost.Set(nil)
}

// UnsetAnsibleHost ensures that no value is present for AnsibleHost, not even an explicit nil
func (o *HostOut) UnsetAnsibleHost() {
	o.AnsibleHost.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *HostOut) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostOut) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *HostOut) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *HostOut) SetCreated(v string) {
	o.Created = &v
}

// GetCulledTimestamp returns the CulledTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HostOut) GetCulledTimestamp() string {
	if o == nil || o.CulledTimestamp.Get() == nil {
		var ret string
		return ret
	}
	return *o.CulledTimestamp.Get()
}

// GetCulledTimestampOk returns a tuple with the CulledTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostOut) GetCulledTimestampOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CulledTimestamp.Get(), o.CulledTimestamp.IsSet()
}

// HasCulledTimestamp returns a boolean if a field has been set.
func (o *HostOut) HasCulledTimestamp() bool {
	if o != nil && o.CulledTimestamp.IsSet() {
		return true
	}

	return false
}

// SetCulledTimestamp gets a reference to the given NullableTime and assigns it to the CulledTimestamp field.
func (o *HostOut) SetCulledTimestamp(v string) {
	o.CulledTimestamp.Set(&v)
}
// SetCulledTimestampNil sets the value for CulledTimestamp to be an explicit nil
func (o *HostOut) SetCulledTimestampNil() {
	o.CulledTimestamp.Set(nil)
}

// UnsetCulledTimestamp ensures that no value is present for CulledTimestamp, not even an explicit nil
func (o *HostOut) UnsetCulledTimestamp() {
	o.CulledTimestamp.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HostOut) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostOut) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *HostOut) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *HostOut) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *HostOut) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *HostOut) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HostOut) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostOut) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HostOut) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HostOut) SetId(v string) {
	o.Id = &v
}

// GetReporter returns the Reporter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HostOut) GetReporter() string {
	if o == nil || o.Reporter.Get() == nil {
		var ret string
		return ret
	}
	return *o.Reporter.Get()
}

// GetReporterOk returns a tuple with the Reporter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostOut) GetReporterOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Reporter.Get(), o.Reporter.IsSet()
}

// HasReporter returns a boolean if a field has been set.
func (o *HostOut) HasReporter() bool {
	if o != nil && o.Reporter.IsSet() {
		return true
	}

	return false
}

// SetReporter gets a reference to the given NullableString and assigns it to the Reporter field.
func (o *HostOut) SetReporter(v string) {
	o.Reporter.Set(&v)
}
// SetReporterNil sets the value for Reporter to be an explicit nil
func (o *HostOut) SetReporterNil() {
	o.Reporter.Set(nil)
}

// UnsetReporter ensures that no value is present for Reporter, not even an explicit nil
func (o *HostOut) UnsetReporter() {
	o.Reporter.Unset()
}

// GetStaleTimestamp returns the StaleTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HostOut) GetStaleTimestamp() string {
	if o == nil || o.StaleTimestamp.Get() == nil {
		var ret string
		return ret
	}
	return *o.StaleTimestamp.Get()
}

// GetStaleTimestampOk returns a tuple with the StaleTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostOut) GetStaleTimestampOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StaleTimestamp.Get(), o.StaleTimestamp.IsSet()
}

// HasStaleTimestamp returns a boolean if a field has been set.
func (o *HostOut) HasStaleTimestamp() bool {
	if o != nil && o.StaleTimestamp.IsSet() {
		return true
	}

	return false
}

// SetStaleTimestamp gets a reference to the given NullableTime and assigns it to the StaleTimestamp field.
func (o *HostOut) SetStaleTimestamp(v string) {
	o.StaleTimestamp.Set(&v)
}
// SetStaleTimestampNil sets the value for StaleTimestamp to be an explicit nil
func (o *HostOut) SetStaleTimestampNil() {
	o.StaleTimestamp.Set(nil)
}

// UnsetStaleTimestamp ensures that no value is present for StaleTimestamp, not even an explicit nil
func (o *HostOut) UnsetStaleTimestamp() {
	o.StaleTimestamp.Unset()
}

// GetStaleWarningTimestamp returns the StaleWarningTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HostOut) GetStaleWarningTimestamp() string {
	if o == nil || o.StaleWarningTimestamp.Get() == nil {
		var ret string
		return ret
	}
	return *o.StaleWarningTimestamp.Get()
}

// GetStaleWarningTimestampOk returns a tuple with the StaleWarningTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostOut) GetStaleWarningTimestampOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StaleWarningTimestamp.Get(), o.StaleWarningTimestamp.IsSet()
}

// HasStaleWarningTimestamp returns a boolean if a field has been set.
func (o *HostOut) HasStaleWarningTimestamp() bool {
	if o != nil && o.StaleWarningTimestamp.IsSet() {
		return true
	}

	return false
}

// SetStaleWarningTimestamp gets a reference to the given NullableTime and assigns it to the StaleWarningTimestamp field.
func (o *HostOut) SetStaleWarningTimestamp(v string) {
	o.StaleWarningTimestamp.Set(&v)
}
// SetStaleWarningTimestampNil sets the value for StaleWarningTimestamp to be an explicit nil
func (o *HostOut) SetStaleWarningTimestampNil() {
	o.StaleWarningTimestamp.Set(nil)
}

// UnsetStaleWarningTimestamp ensures that no value is present for StaleWarningTimestamp, not even an explicit nil
func (o *HostOut) UnsetStaleWarningTimestamp() {
	o.StaleWarningTimestamp.Unset()
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *HostOut) GetUpdated() string {
	if o == nil || o.Updated == nil {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostOut) GetUpdatedOk() (*string, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *HostOut) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *HostOut) SetUpdated(v string) {
	o.Updated = &v
}

func (o HostOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Facts != nil {
		toSerialize["facts"] = o.Facts
	}
	if o.BiosUuid.IsSet() {
		toSerialize["bios_uuid"] = o.BiosUuid.Get()
	}
	if o.ExternalId.IsSet() {
		toSerialize["external_id"] = o.ExternalId.Get()
	}
	if o.Fqdn.IsSet() {
		toSerialize["fqdn"] = o.Fqdn.Get()
	}
	if o.InsightsId.IsSet() {
		toSerialize["insights_id"] = o.InsightsId.Get()
	}
	if o.IpAddresses != nil {
		toSerialize["ip_addresses"] = o.IpAddresses
	}
	if o.MacAddresses != nil {
		toSerialize["mac_addresses"] = o.MacAddresses
	}
	if o.RhelMachineId.IsSet() {
		toSerialize["rhel_machine_id"] = o.RhelMachineId.Get()
	}
	if o.SatelliteId.IsSet() {
		toSerialize["satellite_id"] = o.SatelliteId.Get()
	}
	if o.SubscriptionManagerId.IsSet() {
		toSerialize["subscription_manager_id"] = o.SubscriptionManagerId.Get()
	}
	if true {
		toSerialize["account"] = o.Account
	}
	if o.AnsibleHost.IsSet() {
		toSerialize["ansible_host"] = o.AnsibleHost.Get()
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.CulledTimestamp.IsSet() {
		toSerialize["culled_timestamp"] = o.CulledTimestamp.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["display_name"] = o.DisplayName.Get()
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Reporter.IsSet() {
		toSerialize["reporter"] = o.Reporter.Get()
	}
	if o.StaleTimestamp.IsSet() {
		toSerialize["stale_timestamp"] = o.StaleTimestamp.Get()
	}
	if o.StaleWarningTimestamp.IsSet() {
		toSerialize["stale_warning_timestamp"] = o.StaleWarningTimestamp.Get()
	}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	return json.Marshal(toSerialize)
}

type NullableHostOut struct {
	value *HostOut
	isSet bool
}

func (v NullableHostOut) Get() *HostOut {
	return v.value
}

func (v *NullableHostOut) Set(val *HostOut) {
	v.value = val
	v.isSet = true
}

func (v NullableHostOut) IsSet() bool {
	return v.isSet
}

func (v *NullableHostOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostOut(val *HostOut) *NullableHostOut {
	return &NullableHostOut{value: val, isSet: true}
}

func (v NullableHostOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


