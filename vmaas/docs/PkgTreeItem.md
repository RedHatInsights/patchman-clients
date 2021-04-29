# PkgTreeItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nevra** | **string** |  | 
**Summary** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FirstPublished** | **string** |  | 
**Repositories** | Pointer to [**[]PkgTreeItemRepositories**](PkgTreeItemRepositories.md) |  | [optional] 
**Errata** | Pointer to [**[]PkgTreeItemErrata**](PkgTreeItemErrata.md) |  | [optional] 

## Methods

### NewPkgTreeItem

`func NewPkgTreeItem(nevra string, firstPublished string, ) *PkgTreeItem`

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


### GetSummary

`func (o *PkgTreeItem) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PkgTreeItem) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PkgTreeItem) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PkgTreeItem) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetDescription

`func (o *PkgTreeItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PkgTreeItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PkgTreeItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PkgTreeItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFirstPublished

`func (o *PkgTreeItem) GetFirstPublished() string`

GetFirstPublished returns the FirstPublished field if non-nil, zero value otherwise.

### GetFirstPublishedOk

`func (o *PkgTreeItem) GetFirstPublishedOk() (*string, bool)`

GetFirstPublishedOk returns a tuple with the FirstPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPublished

`func (o *PkgTreeItem) SetFirstPublished(v string)`

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


