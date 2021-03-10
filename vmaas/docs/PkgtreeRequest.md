# PkgtreeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageNameList** | **[]string** |  | 
**ThirdParty** | Pointer to **bool** | Include content from \&quot;third party\&quot; repositories into the response, disabled by default. | [optional] [default to false]

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


