# SystemProfileDiskDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | User-defined mount label | [optional] 
**MountPoint** | Pointer to **string** | The mount point | [optional] 
**Options** | Pointer to **map[string]map[string]interface{}** | An arbitrary object that does not allow empty string keys. | [optional] 
**Type** | Pointer to **string** | The mount type | [optional] 

## Methods

### NewSystemProfileDiskDevice

`func NewSystemProfileDiskDevice() *SystemProfileDiskDevice`

NewSystemProfileDiskDevice instantiates a new SystemProfileDiskDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemProfileDiskDeviceWithDefaults

`func NewSystemProfileDiskDeviceWithDefaults() *SystemProfileDiskDevice`

NewSystemProfileDiskDeviceWithDefaults instantiates a new SystemProfileDiskDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *SystemProfileDiskDevice) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *SystemProfileDiskDevice) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *SystemProfileDiskDevice) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *SystemProfileDiskDevice) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetLabel

`func (o *SystemProfileDiskDevice) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SystemProfileDiskDevice) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SystemProfileDiskDevice) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SystemProfileDiskDevice) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMountPoint

`func (o *SystemProfileDiskDevice) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *SystemProfileDiskDevice) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *SystemProfileDiskDevice) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.

### HasMountPoint

`func (o *SystemProfileDiskDevice) HasMountPoint() bool`

HasMountPoint returns a boolean if a field has been set.

### GetOptions

`func (o *SystemProfileDiskDevice) GetOptions() map[string]map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SystemProfileDiskDevice) GetOptionsOk() (*map[string]map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SystemProfileDiskDevice) SetOptions(v map[string]map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SystemProfileDiskDevice) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetType

`func (o *SystemProfileDiskDevice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemProfileDiskDevice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemProfileDiskDevice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SystemProfileDiskDevice) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


