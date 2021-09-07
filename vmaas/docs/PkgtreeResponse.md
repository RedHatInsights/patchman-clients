# PkgtreeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **float32** |  | [optional] 
**PageSize** | Pointer to **float32** |  | [optional] 
**Pages** | Pointer to **float32** |  | [optional] 
**LastChange** | Pointer to **string** |  | [optional] 
**PackageNameList** | Pointer to [**map[string][]PkgTreeItem**](array.md) |  | [optional] 

## Methods

### NewPkgtreeResponse

`func NewPkgtreeResponse() *PkgtreeResponse`

NewPkgtreeResponse instantiates a new PkgtreeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkgtreeResponseWithDefaults

`func NewPkgtreeResponseWithDefaults() *PkgtreeResponse`

NewPkgtreeResponseWithDefaults instantiates a new PkgtreeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PkgtreeResponse) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PkgtreeResponse) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PkgtreeResponse) SetPage(v float32)`

SetPage sets Page field to given value.

### HasPage

`func (o *PkgtreeResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *PkgtreeResponse) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PkgtreeResponse) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PkgtreeResponse) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *PkgtreeResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPages

`func (o *PkgtreeResponse) GetPages() float32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *PkgtreeResponse) GetPagesOk() (*float32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *PkgtreeResponse) SetPages(v float32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *PkgtreeResponse) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetLastChange

`func (o *PkgtreeResponse) GetLastChange() string`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *PkgtreeResponse) GetLastChangeOk() (*string, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *PkgtreeResponse) SetLastChange(v string)`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *PkgtreeResponse) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.

### GetPackageNameList

`func (o *PkgtreeResponse) GetPackageNameList() map[string][]PkgTreeItem`

GetPackageNameList returns the PackageNameList field if non-nil, zero value otherwise.

### GetPackageNameListOk

`func (o *PkgtreeResponse) GetPackageNameListOk() (*map[string][]PkgTreeItem, bool)`

GetPackageNameListOk returns a tuple with the PackageNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageNameList

`func (o *PkgtreeResponse) SetPackageNameList(v map[string][]PkgTreeItem)`

SetPackageNameList sets PackageNameList field to given value.

### HasPackageNameList

`func (o *PkgtreeResponse) HasPackageNameList() bool`

HasPackageNameList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


