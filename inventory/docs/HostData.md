# HostData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnsibleHost** | Pointer to **string** | The ansible host name for remediations | [optional] 
**DisplayName** | Pointer to **string** | A host’s human-readable display name, e.g. in a form of a domain name. | [optional] 

## Methods

### NewHostData

`func NewHostData() *HostData`

NewHostData instantiates a new HostData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostDataWithDefaults

`func NewHostDataWithDefaults() *HostData`

NewHostDataWithDefaults instantiates a new HostData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnsibleHost

`func (o *HostData) GetAnsibleHost() string`

GetAnsibleHost returns the AnsibleHost field if non-nil, zero value otherwise.

### GetAnsibleHostOk

`func (o *HostData) GetAnsibleHostOk() (*string, bool)`

GetAnsibleHostOk returns a tuple with the AnsibleHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleHost

`func (o *HostData) SetAnsibleHost(v string)`

SetAnsibleHost sets AnsibleHost field to given value.

### HasAnsibleHost

`func (o *HostData) HasAnsibleHost() bool`

HasAnsibleHost returns a boolean if a field has been set.

### GetDisplayName

`func (o *HostData) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *HostData) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *HostData) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *HostData) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


