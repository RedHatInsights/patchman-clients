# CreateCheckIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckinFrequency** | Pointer to **int32** | How long from now to expect another check-in (in minutes). | [optional] 
**BiosUuid** | **string** |  | 
**ExternalId** | **string** |  | 
**Fqdn** | **string** |  | 
**InsightsId** | **string** |  | 
**IpAddresses** | **[]string** |  | 
**MacAddresses** | **[]string** |  | 
**RhelMachineId** | **string** |  | 
**SatelliteId** | **string** |  | 
**SubscriptionManagerId** | **string** |  | 

## Methods

### NewCreateCheckIn

`func NewCreateCheckIn(biosUuid string, externalId string, fqdn string, insightsId string, ipAddresses []string, macAddresses []string, rhelMachineId string, satelliteId string, subscriptionManagerId string, ) *CreateCheckIn`

NewCreateCheckIn instantiates a new CreateCheckIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCheckInWithDefaults

`func NewCreateCheckInWithDefaults() *CreateCheckIn`

NewCreateCheckInWithDefaults instantiates a new CreateCheckIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckinFrequency

`func (o *CreateCheckIn) GetCheckinFrequency() int32`

GetCheckinFrequency returns the CheckinFrequency field if non-nil, zero value otherwise.

### GetCheckinFrequencyOk

`func (o *CreateCheckIn) GetCheckinFrequencyOk() (*int32, bool)`

GetCheckinFrequencyOk returns a tuple with the CheckinFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckinFrequency

`func (o *CreateCheckIn) SetCheckinFrequency(v int32)`

SetCheckinFrequency sets CheckinFrequency field to given value.

### HasCheckinFrequency

`func (o *CreateCheckIn) HasCheckinFrequency() bool`

HasCheckinFrequency returns a boolean if a field has been set.

### GetBiosUuid

`func (o *CreateCheckIn) GetBiosUuid() string`

GetBiosUuid returns the BiosUuid field if non-nil, zero value otherwise.

### GetBiosUuidOk

`func (o *CreateCheckIn) GetBiosUuidOk() (*string, bool)`

GetBiosUuidOk returns a tuple with the BiosUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosUuid

`func (o *CreateCheckIn) SetBiosUuid(v string)`

SetBiosUuid sets BiosUuid field to given value.


### GetExternalId

`func (o *CreateCheckIn) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateCheckIn) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateCheckIn) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetFqdn

`func (o *CreateCheckIn) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *CreateCheckIn) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *CreateCheckIn) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetInsightsId

`func (o *CreateCheckIn) GetInsightsId() string`

GetInsightsId returns the InsightsId field if non-nil, zero value otherwise.

### GetInsightsIdOk

`func (o *CreateCheckIn) GetInsightsIdOk() (*string, bool)`

GetInsightsIdOk returns a tuple with the InsightsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightsId

`func (o *CreateCheckIn) SetInsightsId(v string)`

SetInsightsId sets InsightsId field to given value.


### GetIpAddresses

`func (o *CreateCheckIn) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *CreateCheckIn) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *CreateCheckIn) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.


### GetMacAddresses

`func (o *CreateCheckIn) GetMacAddresses() []string`

GetMacAddresses returns the MacAddresses field if non-nil, zero value otherwise.

### GetMacAddressesOk

`func (o *CreateCheckIn) GetMacAddressesOk() (*[]string, bool)`

GetMacAddressesOk returns a tuple with the MacAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddresses

`func (o *CreateCheckIn) SetMacAddresses(v []string)`

SetMacAddresses sets MacAddresses field to given value.


### GetRhelMachineId

`func (o *CreateCheckIn) GetRhelMachineId() string`

GetRhelMachineId returns the RhelMachineId field if non-nil, zero value otherwise.

### GetRhelMachineIdOk

`func (o *CreateCheckIn) GetRhelMachineIdOk() (*string, bool)`

GetRhelMachineIdOk returns a tuple with the RhelMachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRhelMachineId

`func (o *CreateCheckIn) SetRhelMachineId(v string)`

SetRhelMachineId sets RhelMachineId field to given value.


### GetSatelliteId

`func (o *CreateCheckIn) GetSatelliteId() string`

GetSatelliteId returns the SatelliteId field if non-nil, zero value otherwise.

### GetSatelliteIdOk

`func (o *CreateCheckIn) GetSatelliteIdOk() (*string, bool)`

GetSatelliteIdOk returns a tuple with the SatelliteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatelliteId

`func (o *CreateCheckIn) SetSatelliteId(v string)`

SetSatelliteId sets SatelliteId field to given value.


### GetSubscriptionManagerId

`func (o *CreateCheckIn) GetSubscriptionManagerId() string`

GetSubscriptionManagerId returns the SubscriptionManagerId field if non-nil, zero value otherwise.

### GetSubscriptionManagerIdOk

`func (o *CreateCheckIn) GetSubscriptionManagerIdOk() (*string, bool)`

GetSubscriptionManagerIdOk returns a tuple with the SubscriptionManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionManagerId

`func (o *CreateCheckIn) SetSubscriptionManagerId(v string)`

SetSubscriptionManagerId sets SubscriptionManagerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


