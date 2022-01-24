# SystemProfileOperatingSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Major** | Pointer to **int32** | Major release of OS (aka the x version) | [optional] 
**Minor** | Pointer to **int32** | Minor release of OS (aka the y version) | [optional] 
**Name** | Pointer to **string** | Name of the distro/os | [optional] 

## Methods

### NewSystemProfileOperatingSystem

`func NewSystemProfileOperatingSystem() *SystemProfileOperatingSystem`

NewSystemProfileOperatingSystem instantiates a new SystemProfileOperatingSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemProfileOperatingSystemWithDefaults

`func NewSystemProfileOperatingSystemWithDefaults() *SystemProfileOperatingSystem`

NewSystemProfileOperatingSystemWithDefaults instantiates a new SystemProfileOperatingSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMajor

`func (o *SystemProfileOperatingSystem) GetMajor() int32`

GetMajor returns the Major field if non-nil, zero value otherwise.

### GetMajorOk

`func (o *SystemProfileOperatingSystem) GetMajorOk() (*int32, bool)`

GetMajorOk returns a tuple with the Major field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajor

`func (o *SystemProfileOperatingSystem) SetMajor(v int32)`

SetMajor sets Major field to given value.

### HasMajor

`func (o *SystemProfileOperatingSystem) HasMajor() bool`

HasMajor returns a boolean if a field has been set.

### GetMinor

`func (o *SystemProfileOperatingSystem) GetMinor() int32`

GetMinor returns the Minor field if non-nil, zero value otherwise.

### GetMinorOk

`func (o *SystemProfileOperatingSystem) GetMinorOk() (*int32, bool)`

GetMinorOk returns a tuple with the Minor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinor

`func (o *SystemProfileOperatingSystem) SetMinor(v int32)`

SetMinor sets Minor field to given value.

### HasMinor

`func (o *SystemProfileOperatingSystem) HasMinor() bool`

HasMinor returns a boolean if a field has been set.

### GetName

`func (o *SystemProfileOperatingSystem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemProfileOperatingSystem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemProfileOperatingSystem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SystemProfileOperatingSystem) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


