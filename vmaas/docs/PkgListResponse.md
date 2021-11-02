# PkgListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **float32** |  | [optional] 
**PageSize** | Pointer to **float32** |  | [optional] 
**Pages** | Pointer to **float32** |  | [optional] 
**LastChange** | Pointer to **string** |  | [optional] 
**PackageList** | Pointer to [**[]PkgListItem**](PkgListItem.md) |  | [optional] 
**Total** | Pointer to **float32** | Total number of packages to return. | [optional] 

## Methods

### NewPkgListResponse

`func NewPkgListResponse() *PkgListResponse`

NewPkgListResponse instantiates a new PkgListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkgListResponseWithDefaults

`func NewPkgListResponseWithDefaults() *PkgListResponse`

NewPkgListResponseWithDefaults instantiates a new PkgListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PkgListResponse) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PkgListResponse) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PkgListResponse) SetPage(v float32)`

SetPage sets Page field to given value.

### HasPage

`func (o *PkgListResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *PkgListResponse) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PkgListResponse) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PkgListResponse) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *PkgListResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPages

`func (o *PkgListResponse) GetPages() float32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *PkgListResponse) GetPagesOk() (*float32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *PkgListResponse) SetPages(v float32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *PkgListResponse) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetLastChange

`func (o *PkgListResponse) GetLastChange() string`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *PkgListResponse) GetLastChangeOk() (*string, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *PkgListResponse) SetLastChange(v string)`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *PkgListResponse) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.

### GetPackageList

`func (o *PkgListResponse) GetPackageList() []PkgListItem`

GetPackageList returns the PackageList field if non-nil, zero value otherwise.

### GetPackageListOk

`func (o *PkgListResponse) GetPackageListOk() (*[]PkgListItem, bool)`

GetPackageListOk returns a tuple with the PackageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageList

`func (o *PkgListResponse) SetPackageList(v []PkgListItem)`

SetPackageList sets PackageList field to given value.

### HasPackageList

`func (o *PkgListResponse) HasPackageList() bool`

HasPackageList returns a boolean if a field has been set.

### GetTotal

`func (o *PkgListResponse) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PkgListResponse) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PkgListResponse) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PkgListResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


