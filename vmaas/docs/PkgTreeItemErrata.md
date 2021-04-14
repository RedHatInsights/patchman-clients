# PkgTreeItemErrata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Issued** | **time.Time** |  | 
**CveList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPkgTreeItemErrata

`func NewPkgTreeItemErrata(name string, issued time.Time, ) *PkgTreeItemErrata`

NewPkgTreeItemErrata instantiates a new PkgTreeItemErrata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkgTreeItemErrataWithDefaults

`func NewPkgTreeItemErrataWithDefaults() *PkgTreeItemErrata`

NewPkgTreeItemErrataWithDefaults instantiates a new PkgTreeItemErrata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PkgTreeItemErrata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PkgTreeItemErrata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PkgTreeItemErrata) SetName(v string)`

SetName sets Name field to given value.


### GetIssued

`func (o *PkgTreeItemErrata) GetIssued() time.Time`

GetIssued returns the Issued field if non-nil, zero value otherwise.

### GetIssuedOk

`func (o *PkgTreeItemErrata) GetIssuedOk() (*time.Time, bool)`

GetIssuedOk returns a tuple with the Issued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssued

`func (o *PkgTreeItemErrata) SetIssued(v time.Time)`

SetIssued sets Issued field to given value.


### GetCveList

`func (o *PkgTreeItemErrata) GetCveList() []string`

GetCveList returns the CveList field if non-nil, zero value otherwise.

### GetCveListOk

`func (o *PkgTreeItemErrata) GetCveListOk() (*[]string, bool)`

GetCveListOk returns a tuple with the CveList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveList

`func (o *PkgTreeItemErrata) SetCveList(v []string)`

SetCveList sets CveList field to given value.

### HasCveList

`func (o *PkgTreeItemErrata) HasCveList() bool`

HasCveList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


