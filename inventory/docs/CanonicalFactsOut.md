# CanonicalFactsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BiosUuid** | Pointer to **NullableString** | A UUID of the host machine BIOS.  This field is considered to be a canonical fact. | [optional] 
**ExternalId** | Pointer to **NullableString** | Host’s reference in the external source e.g. AWS EC2, Azure, OpenStack, etc. This field is considered to be a canonical fact. | [optional] 
**Fqdn** | Pointer to **NullableString** | A host’s Fully Qualified Domain Name.  This field is considered to be a canonical fact. | [optional] 
**InsightsId** | Pointer to **NullableString** | An ID defined in /etc/insights-client/machine-id. This field is considered a canonical fact. | [optional] 
**IpAddresses** | Pointer to **[]string** | Host’s network IP addresses.  This field is considered to be a canonical fact. | [optional] 
**MacAddresses** | Pointer to **[]string** | Host’s network interfaces MAC addresses.  This field is considered to be a canonical fact. | [optional] 
**RhelMachineId** | Pointer to **NullableString** | A Machine ID of a RHEL host.  This field is considered to be a canonical fact. | [optional] 
**SatelliteId** | Pointer to **NullableString** | A Red Hat Satellite ID of a RHEL host.  This field is considered to be a canonical fact. | [optional] 
**SubscriptionManagerId** | Pointer to **NullableString** | A Red Hat Subcription Manager ID of a RHEL host.  This field is considered to be a canonical fact. | [optional] 

## Methods

### NewCanonicalFactsOut

`func NewCanonicalFactsOut() *CanonicalFactsOut`

NewCanonicalFactsOut instantiates a new CanonicalFactsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCanonicalFactsOutWithDefaults

`func NewCanonicalFactsOutWithDefaults() *CanonicalFactsOut`

NewCanonicalFactsOutWithDefaults instantiates a new CanonicalFactsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBiosUuid

`func (o *CanonicalFactsOut) GetBiosUuid() string`

GetBiosUuid returns the BiosUuid field if non-nil, zero value otherwise.

### GetBiosUuidOk

`func (o *CanonicalFactsOut) GetBiosUuidOk() (*string, bool)`

GetBiosUuidOk returns a tuple with the BiosUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosUuid

`func (o *CanonicalFactsOut) SetBiosUuid(v string)`

SetBiosUuid sets BiosUuid field to given value.

### HasBiosUuid

`func (o *CanonicalFactsOut) HasBiosUuid() bool`

HasBiosUuid returns a boolean if a field has been set.

### SetBiosUuidNil

`func (o *CanonicalFactsOut) SetBiosUuidNil(b bool)`

 SetBiosUuidNil sets the value for BiosUuid to be an explicit nil

### UnsetBiosUuid
`func (o *CanonicalFactsOut) UnsetBiosUuid()`

UnsetBiosUuid ensures that no value is present for BiosUuid, not even an explicit nil
### GetExternalId

`func (o *CanonicalFactsOut) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CanonicalFactsOut) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CanonicalFactsOut) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CanonicalFactsOut) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *CanonicalFactsOut) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *CanonicalFactsOut) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetFqdn

`func (o *CanonicalFactsOut) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *CanonicalFactsOut) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *CanonicalFactsOut) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *CanonicalFactsOut) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### SetFqdnNil

`func (o *CanonicalFactsOut) SetFqdnNil(b bool)`

 SetFqdnNil sets the value for Fqdn to be an explicit nil

### UnsetFqdn
`func (o *CanonicalFactsOut) UnsetFqdn()`

UnsetFqdn ensures that no value is present for Fqdn, not even an explicit nil
### GetInsightsId

`func (o *CanonicalFactsOut) GetInsightsId() string`

GetInsightsId returns the InsightsId field if non-nil, zero value otherwise.

### GetInsightsIdOk

`func (o *CanonicalFactsOut) GetInsightsIdOk() (*string, bool)`

GetInsightsIdOk returns a tuple with the InsightsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightsId

`func (o *CanonicalFactsOut) SetInsightsId(v string)`

SetInsightsId sets InsightsId field to given value.

### HasInsightsId

`func (o *CanonicalFactsOut) HasInsightsId() bool`

HasInsightsId returns a boolean if a field has been set.

### SetInsightsIdNil

`func (o *CanonicalFactsOut) SetInsightsIdNil(b bool)`

 SetInsightsIdNil sets the value for InsightsId to be an explicit nil

### UnsetInsightsId
`func (o *CanonicalFactsOut) UnsetInsightsId()`

UnsetInsightsId ensures that no value is present for InsightsId, not even an explicit nil
### GetIpAddresses

`func (o *CanonicalFactsOut) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *CanonicalFactsOut) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *CanonicalFactsOut) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *CanonicalFactsOut) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### SetIpAddressesNil

`func (o *CanonicalFactsOut) SetIpAddressesNil(b bool)`

 SetIpAddressesNil sets the value for IpAddresses to be an explicit nil

### UnsetIpAddresses
`func (o *CanonicalFactsOut) UnsetIpAddresses()`

UnsetIpAddresses ensures that no value is present for IpAddresses, not even an explicit nil
### GetMacAddresses

`func (o *CanonicalFactsOut) GetMacAddresses() []string`

GetMacAddresses returns the MacAddresses field if non-nil, zero value otherwise.

### GetMacAddressesOk

`func (o *CanonicalFactsOut) GetMacAddressesOk() (*[]string, bool)`

GetMacAddressesOk returns a tuple with the MacAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddresses

`func (o *CanonicalFactsOut) SetMacAddresses(v []string)`

SetMacAddresses sets MacAddresses field to given value.

### HasMacAddresses

`func (o *CanonicalFactsOut) HasMacAddresses() bool`

HasMacAddresses returns a boolean if a field has been set.

### SetMacAddressesNil

`func (o *CanonicalFactsOut) SetMacAddressesNil(b bool)`

 SetMacAddressesNil sets the value for MacAddresses to be an explicit nil

### UnsetMacAddresses
`func (o *CanonicalFactsOut) UnsetMacAddresses()`

UnsetMacAddresses ensures that no value is present for MacAddresses, not even an explicit nil
### GetRhelMachineId

`func (o *CanonicalFactsOut) GetRhelMachineId() string`

GetRhelMachineId returns the RhelMachineId field if non-nil, zero value otherwise.

### GetRhelMachineIdOk

`func (o *CanonicalFactsOut) GetRhelMachineIdOk() (*string, bool)`

GetRhelMachineIdOk returns a tuple with the RhelMachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRhelMachineId

`func (o *CanonicalFactsOut) SetRhelMachineId(v string)`

SetRhelMachineId sets RhelMachineId field to given value.

### HasRhelMachineId

`func (o *CanonicalFactsOut) HasRhelMachineId() bool`

HasRhelMachineId returns a boolean if a field has been set.

### SetRhelMachineIdNil

`func (o *CanonicalFactsOut) SetRhelMachineIdNil(b bool)`

 SetRhelMachineIdNil sets the value for RhelMachineId to be an explicit nil

### UnsetRhelMachineId
`func (o *CanonicalFactsOut) UnsetRhelMachineId()`

UnsetRhelMachineId ensures that no value is present for RhelMachineId, not even an explicit nil
### GetSatelliteId

`func (o *CanonicalFactsOut) GetSatelliteId() string`

GetSatelliteId returns the SatelliteId field if non-nil, zero value otherwise.

### GetSatelliteIdOk

`func (o *CanonicalFactsOut) GetSatelliteIdOk() (*string, bool)`

GetSatelliteIdOk returns a tuple with the SatelliteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatelliteId

`func (o *CanonicalFactsOut) SetSatelliteId(v string)`

SetSatelliteId sets SatelliteId field to given value.

### HasSatelliteId

`func (o *CanonicalFactsOut) HasSatelliteId() bool`

HasSatelliteId returns a boolean if a field has been set.

### SetSatelliteIdNil

`func (o *CanonicalFactsOut) SetSatelliteIdNil(b bool)`

 SetSatelliteIdNil sets the value for SatelliteId to be an explicit nil

### UnsetSatelliteId
`func (o *CanonicalFactsOut) UnsetSatelliteId()`

UnsetSatelliteId ensures that no value is present for SatelliteId, not even an explicit nil
### GetSubscriptionManagerId

`func (o *CanonicalFactsOut) GetSubscriptionManagerId() string`

GetSubscriptionManagerId returns the SubscriptionManagerId field if non-nil, zero value otherwise.

### GetSubscriptionManagerIdOk

`func (o *CanonicalFactsOut) GetSubscriptionManagerIdOk() (*string, bool)`

GetSubscriptionManagerIdOk returns a tuple with the SubscriptionManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionManagerId

`func (o *CanonicalFactsOut) SetSubscriptionManagerId(v string)`

SetSubscriptionManagerId sets SubscriptionManagerId field to given value.

### HasSubscriptionManagerId

`func (o *CanonicalFactsOut) HasSubscriptionManagerId() bool`

HasSubscriptionManagerId returns a boolean if a field has been set.

### SetSubscriptionManagerIdNil

`func (o *CanonicalFactsOut) SetSubscriptionManagerIdNil(b bool)`

 SetSubscriptionManagerIdNil sets the value for SubscriptionManagerId to be an explicit nil

### UnsetSubscriptionManagerId
`func (o *CanonicalFactsOut) UnsetSubscriptionManagerId()`

UnsetSubscriptionManagerId ensures that no value is present for SubscriptionManagerId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


