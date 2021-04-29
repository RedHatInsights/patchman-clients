# PkgtreeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **float32** |  | [optional] 
**PageSize** | Pointer to **float32** |  | [optional] 
**PackageNameList** | **[]string** |  | 
**ModifiedSince** | Pointer to **string** |  | [optional] 
**ThirdParty** | Pointer to **bool** | Include content from \&quot;third party\&quot; repositories into the response, disabled by default. | [optional] [default to false]
**ReturnRepositories** | Pointer to **bool** | Include nevra repositories info into the response. | [optional] [default to true]
**ReturnErrata** | Pointer to **bool** | Include nevra errata info into the response. | [optional] [default to true]
**ReturnSummary** | Pointer to **bool** | Include nevra summary info into the response. | [optional] [default to false]
**ReturnDescription** | Pointer to **bool** | Include nevra description info into the response. | [optional] [default to false]

## Methods

### NewPkgtreeRequest

`func NewPkgtreeRequest(packageNameList []string, ) *PkgtreeRequest`

NewPkgtreeRequest instantiates a new PkgtreeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkgtreeRequestWithDefaults

`func NewPkgtreeRequestWithDefaults() *PkgtreeRequest`

NewPkgtreeRequestWithDefaults instantiates a new PkgtreeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PkgtreeRequest) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PkgtreeRequest) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PkgtreeRequest) SetPage(v float32)`

SetPage sets Page field to given value.

### HasPage

`func (o *PkgtreeRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *PkgtreeRequest) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PkgtreeRequest) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PkgtreeRequest) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *PkgtreeRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPackageNameList

`func (o *PkgtreeRequest) GetPackageNameList() []string`

GetPackageNameList returns the PackageNameList field if non-nil, zero value otherwise.

### GetPackageNameListOk

`func (o *PkgtreeRequest) GetPackageNameListOk() (*[]string, bool)`

GetPackageNameListOk returns a tuple with the PackageNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageNameList

`func (o *PkgtreeRequest) SetPackageNameList(v []string)`

SetPackageNameList sets PackageNameList field to given value.


### GetModifiedSince

`func (o *PkgtreeRequest) GetModifiedSince() string`

GetModifiedSince returns the ModifiedSince field if non-nil, zero value otherwise.

### GetModifiedSinceOk

`func (o *PkgtreeRequest) GetModifiedSinceOk() (*string, bool)`

GetModifiedSinceOk returns a tuple with the ModifiedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedSince

`func (o *PkgtreeRequest) SetModifiedSince(v string)`

SetModifiedSince sets ModifiedSince field to given value.

### HasModifiedSince

`func (o *PkgtreeRequest) HasModifiedSince() bool`

HasModifiedSince returns a boolean if a field has been set.

### GetThirdParty

`func (o *PkgtreeRequest) GetThirdParty() bool`

GetThirdParty returns the ThirdParty field if non-nil, zero value otherwise.

### GetThirdPartyOk

`func (o *PkgtreeRequest) GetThirdPartyOk() (*bool, bool)`

GetThirdPartyOk returns a tuple with the ThirdParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdParty

`func (o *PkgtreeRequest) SetThirdParty(v bool)`

SetThirdParty sets ThirdParty field to given value.

### HasThirdParty

`func (o *PkgtreeRequest) HasThirdParty() bool`

HasThirdParty returns a boolean if a field has been set.

### GetReturnRepositories

`func (o *PkgtreeRequest) GetReturnRepositories() bool`

GetReturnRepositories returns the ReturnRepositories field if non-nil, zero value otherwise.

### GetReturnRepositoriesOk

`func (o *PkgtreeRequest) GetReturnRepositoriesOk() (*bool, bool)`

GetReturnRepositoriesOk returns a tuple with the ReturnRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnRepositories

`func (o *PkgtreeRequest) SetReturnRepositories(v bool)`

SetReturnRepositories sets ReturnRepositories field to given value.

### HasReturnRepositories

`func (o *PkgtreeRequest) HasReturnRepositories() bool`

HasReturnRepositories returns a boolean if a field has been set.

### GetReturnErrata

`func (o *PkgtreeRequest) GetReturnErrata() bool`

GetReturnErrata returns the ReturnErrata field if non-nil, zero value otherwise.

### GetReturnErrataOk

`func (o *PkgtreeRequest) GetReturnErrataOk() (*bool, bool)`

GetReturnErrataOk returns a tuple with the ReturnErrata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnErrata

`func (o *PkgtreeRequest) SetReturnErrata(v bool)`

SetReturnErrata sets ReturnErrata field to given value.

### HasReturnErrata

`func (o *PkgtreeRequest) HasReturnErrata() bool`

HasReturnErrata returns a boolean if a field has been set.

### GetReturnSummary

`func (o *PkgtreeRequest) GetReturnSummary() bool`

GetReturnSummary returns the ReturnSummary field if non-nil, zero value otherwise.

### GetReturnSummaryOk

`func (o *PkgtreeRequest) GetReturnSummaryOk() (*bool, bool)`

GetReturnSummaryOk returns a tuple with the ReturnSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnSummary

`func (o *PkgtreeRequest) SetReturnSummary(v bool)`

SetReturnSummary sets ReturnSummary field to given value.

### HasReturnSummary

`func (o *PkgtreeRequest) HasReturnSummary() bool`

HasReturnSummary returns a boolean if a field has been set.

### GetReturnDescription

`func (o *PkgtreeRequest) GetReturnDescription() bool`

GetReturnDescription returns the ReturnDescription field if non-nil, zero value otherwise.

### GetReturnDescriptionOk

`func (o *PkgtreeRequest) GetReturnDescriptionOk() (*bool, bool)`

GetReturnDescriptionOk returns a tuple with the ReturnDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnDescription

`func (o *PkgtreeRequest) SetReturnDescription(v bool)`

SetReturnDescription sets ReturnDescription field to given value.

### HasReturnDescription

`func (o *PkgtreeRequest) HasReturnDescription() bool`

HasReturnDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


