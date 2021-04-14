# PkgTreeItemRepositories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | 
**Name** | **string** |  | 
**Basearch** | **string** |  | 
**Releasever** | **string** |  | 
**Revision** | **time.Time** |  | 
**ModuleName** | Pointer to **string** |  | [optional] 
**ModuleStream** | Pointer to **string** |  | [optional] 

## Methods

### NewPkgTreeItemRepositories

`func NewPkgTreeItemRepositories(label string, name string, basearch string, releasever string, revision time.Time, ) *PkgTreeItemRepositories`

NewPkgTreeItemRepositories instantiates a new PkgTreeItemRepositories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkgTreeItemRepositoriesWithDefaults

`func NewPkgTreeItemRepositoriesWithDefaults() *PkgTreeItemRepositories`

NewPkgTreeItemRepositoriesWithDefaults instantiates a new PkgTreeItemRepositories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *PkgTreeItemRepositories) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PkgTreeItemRepositories) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PkgTreeItemRepositories) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *PkgTreeItemRepositories) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PkgTreeItemRepositories) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PkgTreeItemRepositories) SetName(v string)`

SetName sets Name field to given value.


### GetBasearch

`func (o *PkgTreeItemRepositories) GetBasearch() string`

GetBasearch returns the Basearch field if non-nil, zero value otherwise.

### GetBasearchOk

`func (o *PkgTreeItemRepositories) GetBasearchOk() (*string, bool)`

GetBasearchOk returns a tuple with the Basearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasearch

`func (o *PkgTreeItemRepositories) SetBasearch(v string)`

SetBasearch sets Basearch field to given value.


### GetReleasever

`func (o *PkgTreeItemRepositories) GetReleasever() string`

GetReleasever returns the Releasever field if non-nil, zero value otherwise.

### GetReleaseverOk

`func (o *PkgTreeItemRepositories) GetReleaseverOk() (*string, bool)`

GetReleaseverOk returns a tuple with the Releasever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasever

`func (o *PkgTreeItemRepositories) SetReleasever(v string)`

SetReleasever sets Releasever field to given value.


### GetRevision

`func (o *PkgTreeItemRepositories) GetRevision() time.Time`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *PkgTreeItemRepositories) GetRevisionOk() (*time.Time, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *PkgTreeItemRepositories) SetRevision(v time.Time)`

SetRevision sets Revision field to given value.


### GetModuleName

`func (o *PkgTreeItemRepositories) GetModuleName() string`

GetModuleName returns the ModuleName field if non-nil, zero value otherwise.

### GetModuleNameOk

`func (o *PkgTreeItemRepositories) GetModuleNameOk() (*string, bool)`

GetModuleNameOk returns a tuple with the ModuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleName

`func (o *PkgTreeItemRepositories) SetModuleName(v string)`

SetModuleName sets ModuleName field to given value.

### HasModuleName

`func (o *PkgTreeItemRepositories) HasModuleName() bool`

HasModuleName returns a boolean if a field has been set.

### GetModuleStream

`func (o *PkgTreeItemRepositories) GetModuleStream() string`

GetModuleStream returns the ModuleStream field if non-nil, zero value otherwise.

### GetModuleStreamOk

`func (o *PkgTreeItemRepositories) GetModuleStreamOk() (*string, bool)`

GetModuleStreamOk returns a tuple with the ModuleStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleStream

`func (o *PkgTreeItemRepositories) SetModuleStream(v string)`

SetModuleStream sets ModuleStream field to given value.

### HasModuleStream

`func (o *PkgTreeItemRepositories) HasModuleStream() bool`

HasModuleStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


