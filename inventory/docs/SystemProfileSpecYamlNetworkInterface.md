# SystemProfileSpecYamlNetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4Addresses** | Pointer to **[]string** |  | [optional] 
**Ipv6Addresses** | Pointer to **[]string** |  | [optional] 
**MacAddress** | Pointer to **string** | MAC address (with or without colons) | [optional] 
**Mtu** | Pointer to **int32** | MTU (Maximum transmission unit) | [optional] 
**Name** | Pointer to **string** | Name of interface | [optional] 
**State** | Pointer to **string** | Interface state (UP, DOWN, UNKNOWN) | [optional] 
**Type** | Pointer to **string** | Interface type (ether, loopback) | [optional] 

## Methods

### NewSystemProfileSpecYamlNetworkInterface

`func NewSystemProfileSpecYamlNetworkInterface() *SystemProfileSpecYamlNetworkInterface`

NewSystemProfileSpecYamlNetworkInterface instantiates a new SystemProfileSpecYamlNetworkInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemProfileSpecYamlNetworkInterfaceWithDefaults

`func NewSystemProfileSpecYamlNetworkInterfaceWithDefaults() *SystemProfileSpecYamlNetworkInterface`

NewSystemProfileSpecYamlNetworkInterfaceWithDefaults instantiates a new SystemProfileSpecYamlNetworkInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4Addresses

`func (o *SystemProfileSpecYamlNetworkInterface) GetIpv4Addresses() []string`

GetIpv4Addresses returns the Ipv4Addresses field if non-nil, zero value otherwise.

### GetIpv4AddressesOk

`func (o *SystemProfileSpecYamlNetworkInterface) GetIpv4AddressesOk() (*[]string, bool)`

GetIpv4AddressesOk returns a tuple with the Ipv4Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addresses

`func (o *SystemProfileSpecYamlNetworkInterface) SetIpv4Addresses(v []string)`

SetIpv4Addresses sets Ipv4Addresses field to given value.

### HasIpv4Addresses

`func (o *SystemProfileSpecYamlNetworkInterface) HasIpv4Addresses() bool`

HasIpv4Addresses returns a boolean if a field has been set.

### GetIpv6Addresses

`func (o *SystemProfileSpecYamlNetworkInterface) GetIpv6Addresses() []string`

GetIpv6Addresses returns the Ipv6Addresses field if non-nil, zero value otherwise.

### GetIpv6AddressesOk

`func (o *SystemProfileSpecYamlNetworkInterface) GetIpv6AddressesOk() (*[]string, bool)`

GetIpv6AddressesOk returns a tuple with the Ipv6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addresses

`func (o *SystemProfileSpecYamlNetworkInterface) SetIpv6Addresses(v []string)`

SetIpv6Addresses sets Ipv6Addresses field to given value.

### HasIpv6Addresses

`func (o *SystemProfileSpecYamlNetworkInterface) HasIpv6Addresses() bool`

HasIpv6Addresses returns a boolean if a field has been set.

### GetMacAddress

`func (o *SystemProfileSpecYamlNetworkInterface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *SystemProfileSpecYamlNetworkInterface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *SystemProfileSpecYamlNetworkInterface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *SystemProfileSpecYamlNetworkInterface) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMtu

`func (o *SystemProfileSpecYamlNetworkInterface) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *SystemProfileSpecYamlNetworkInterface) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *SystemProfileSpecYamlNetworkInterface) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *SystemProfileSpecYamlNetworkInterface) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *SystemProfileSpecYamlNetworkInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemProfileSpecYamlNetworkInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemProfileSpecYamlNetworkInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SystemProfileSpecYamlNetworkInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *SystemProfileSpecYamlNetworkInterface) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SystemProfileSpecYamlNetworkInterface) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SystemProfileSpecYamlNetworkInterface) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SystemProfileSpecYamlNetworkInterface) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *SystemProfileSpecYamlNetworkInterface) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemProfileSpecYamlNetworkInterface) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemProfileSpecYamlNetworkInterface) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SystemProfileSpecYamlNetworkInterface) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


