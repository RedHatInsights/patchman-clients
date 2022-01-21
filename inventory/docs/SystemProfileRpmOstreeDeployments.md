# SystemProfileRpmOstreeDeployments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the deployment | 
**Checksum** | **string** | The checksum / commit of the deployment | 
**Origin** | **string** | The origin repo from which the commit was installed | 
**Osname** | **string** | The operating system name | 
**Version** | Pointer to **string** | The version of the deployment | [optional] 
**Booted** | **bool** | Whether the deployment is currently booted | 
**Pinned** | **bool** | Whether the deployment is currently pinned | 

## Methods

### NewSystemProfileRpmOstreeDeployments

`func NewSystemProfileRpmOstreeDeployments(id string, checksum string, origin string, osname string, booted bool, pinned bool, ) *SystemProfileRpmOstreeDeployments`

NewSystemProfileRpmOstreeDeployments instantiates a new SystemProfileRpmOstreeDeployments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemProfileRpmOstreeDeploymentsWithDefaults

`func NewSystemProfileRpmOstreeDeploymentsWithDefaults() *SystemProfileRpmOstreeDeployments`

NewSystemProfileRpmOstreeDeploymentsWithDefaults instantiates a new SystemProfileRpmOstreeDeployments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SystemProfileRpmOstreeDeployments) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SystemProfileRpmOstreeDeployments) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SystemProfileRpmOstreeDeployments) SetId(v string)`

SetId sets Id field to given value.


### GetChecksum

`func (o *SystemProfileRpmOstreeDeployments) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *SystemProfileRpmOstreeDeployments) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *SystemProfileRpmOstreeDeployments) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.


### GetOrigin

`func (o *SystemProfileRpmOstreeDeployments) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *SystemProfileRpmOstreeDeployments) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *SystemProfileRpmOstreeDeployments) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetOsname

`func (o *SystemProfileRpmOstreeDeployments) GetOsname() string`

GetOsname returns the Osname field if non-nil, zero value otherwise.

### GetOsnameOk

`func (o *SystemProfileRpmOstreeDeployments) GetOsnameOk() (*string, bool)`

GetOsnameOk returns a tuple with the Osname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsname

`func (o *SystemProfileRpmOstreeDeployments) SetOsname(v string)`

SetOsname sets Osname field to given value.


### GetVersion

`func (o *SystemProfileRpmOstreeDeployments) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SystemProfileRpmOstreeDeployments) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SystemProfileRpmOstreeDeployments) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SystemProfileRpmOstreeDeployments) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBooted

`func (o *SystemProfileRpmOstreeDeployments) GetBooted() bool`

GetBooted returns the Booted field if non-nil, zero value otherwise.

### GetBootedOk

`func (o *SystemProfileRpmOstreeDeployments) GetBootedOk() (*bool, bool)`

GetBootedOk returns a tuple with the Booted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooted

`func (o *SystemProfileRpmOstreeDeployments) SetBooted(v bool)`

SetBooted sets Booted field to given value.


### GetPinned

`func (o *SystemProfileRpmOstreeDeployments) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *SystemProfileRpmOstreeDeployments) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *SystemProfileRpmOstreeDeployments) SetPinned(v bool)`

SetPinned sets Pinned field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


