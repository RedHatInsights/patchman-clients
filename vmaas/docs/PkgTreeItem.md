# PkgTreeItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nevra** | **string** |  | 
**FirstPublished** | **time.Time** |  | 
**Repositories** | Pointer to [**[]PkgTreeItemRepositories**](PkgTreeItemRepositories.md) |  | [optional] 
**Errata** | Pointer to [**[]PkgTreeItemErrata**](PkgTreeItemErrata.md) |  | [optional] 

## Methods

### NewPkgTreeItem

`func NewPkgTreeItem(nevra string, firstPublished time.Time, ) *PkgTreeItem`

NewPkgTreeItem instantiates a new PkgTreeItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkgTreeItemWithDefaults

`func NewPkgTreeItemWithDefaults() *PkgTreeItem`

NewPkgTreeItemWithDefaults instantiates a new PkgTreeItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNevra

`func (o *PkgTreeItem) GetNevra() string`

GetNevra returns the Nevra field if non-nil, zero value otherwise.

### GetNevraOk

`func (o *PkgTreeItem) GetNevraOk() (*string, bool)`

GetNevraOk returns a tuple with the Nevra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNevra

`func (o *PkgTreeItem) SetNevra(v string)`

SetNevra sets Nevra field to given value.


### GetFirstPublished

`func (o *PkgTreeItem) GetFirstPublished() time.Time`

GetFirstPublished returns the FirstPublished field if non-nil, zero value otherwise.

### GetFirstPublishedOk

`func (o *PkgTreeItem) GetFirstPublishedOk() (*time.Time, bool)`

GetFirstPublishedOk returns a tuple with the FirstPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPublished

`func (o *PkgTreeItem) SetFirstPublished(v time.Time)`

SetFirstPublished sets FirstPublished field to given value.


### GetRepositories

`func (o *PkgTreeItem) GetRepositories() []PkgTreeItemRepositories`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *PkgTreeItem) GetRepositoriesOk() (*[]PkgTreeItemRepositories, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *PkgTreeItem) SetRepositories(v []PkgTreeItemRepositories)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *PkgTreeItem) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetErrata

`func (o *PkgTreeItem) GetErrata() []PkgTreeItemErrata`

GetErrata returns the Errata field if non-nil, zero value otherwise.

### GetErrataOk

`func (o *PkgTreeItem) GetErrataOk() (*[]PkgTreeItemErrata, bool)`

GetErrataOk returns a tuple with the Errata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrata

`func (o *PkgTreeItem) SetErrata(v []PkgTreeItemErrata)`

SetErrata sets Errata field to given value.

### HasErrata

`func (o *PkgTreeItem) HasErrata() bool`

HasErrata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

