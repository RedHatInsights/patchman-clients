# SystemProfileSpecYamlDiskDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | User-defined mount label | [optional] 
**MountPoint** | Pointer to **string** | The mount point | [optional] 
**Options** | Pointer to **map[string]map[string]interface{}** | An arbitrary object that does not allow empty string keys. | [optional] 
**Type** | Pointer to **string** | The mount type | [optional] 

## Methods

### NewSystemProfileSpecYamlDiskDevice

`func NewSystemProfileSpecYamlDiskDevice() *SystemProfileSpecYamlDiskDevice`

NewSystemProfileSpecYamlDiskDevice instantiates a new SystemProfileSpecYamlDiskDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemProfileSpecYamlDiskDeviceWithDefaults

`func NewSystemProfileSpecYamlDiskDeviceWithDefaults() *SystemProfileSpecYamlDiskDevice`

NewSystemProfileSpecYamlDiskDeviceWithDefaults instantiates a new SystemProfileSpecYamlDiskDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *SystemProfileSpecYamlDiskDevice) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *SystemProfileSpecYamlDiskDevice) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *SystemProfileSpecYamlDiskDevice) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *SystemProfileSpecYamlDiskDevice) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetLabel

`func (o *SystemProfileSpecYamlDiskDevice) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SystemProfileSpecYamlDiskDevice) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SystemProfileSpecYamlDiskDevice) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SystemProfileSpecYamlDiskDevice) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMountPoint

`func (o *SystemProfileSpecYamlDiskDevice) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *SystemProfileSpecYamlDiskDevice) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *SystemProfileSpecYamlDiskDevice) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.

### HasMountPoint

`func (o *SystemProfileSpecYamlDiskDevice) HasMountPoint() bool`

HasMountPoint returns a boolean if a field has been set.

### GetOptions

`func (o *SystemProfileSpecYamlDiskDevice) GetOptions() map[string]map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SystemProfileSpecYamlDiskDevice) GetOptionsOk() (*map[string]map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SystemProfileSpecYamlDiskDevice) SetOptions(v map[string]map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SystemProfileSpecYamlDiskDevice) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetType

`func (o *SystemProfileSpecYamlDiskDevice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemProfileSpecYamlDiskDevice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemProfileSpecYamlDiskDevice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SystemProfileSpecYamlDiskDevice) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


