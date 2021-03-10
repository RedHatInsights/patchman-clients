# PackagesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageList** | **[]string** |  | 
**ThirdParty** | Pointer to **bool** | Include content from \&quot;third party\&quot; repositories into the response, disabled by default. | [optional] [default to false]

## Methods

### NewPackagesRequest

`func NewPackagesRequest(packageList []string, ) *PackagesRequest`

NewPackagesRequest instantiates a new PackagesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackagesRequestWithDefaults

`func NewPackagesRequestWithDefaults() *PackagesRequest`

NewPackagesRequestWithDefaults instantiates a new PackagesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageList

`func (o *PackagesRequest) GetPackageList() []string`

GetPackageList returns the PackageList field if non-nil, zero value otherwise.

### GetPackageListOk

`func (o *PackagesRequest) GetPackageListOk() (*[]string, bool)`

GetPackageListOk returns a tuple with the PackageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageList

`func (o *PackagesRequest) SetPackageList(v []string)`

SetPackageList sets PackageList field to given value.


### GetThirdParty

`func (o *PackagesRequest) GetThirdParty() bool`

GetThirdParty returns the ThirdParty field if non-nil, zero value otherwise.

### GetThirdPartyOk

`func (o *PackagesRequest) GetThirdPartyOk() (*bool, bool)`

GetThirdPartyOk returns a tuple with the ThirdParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdParty

`func (o *PackagesRequest) SetThirdParty(v bool)`

SetThirdParty sets ThirdParty field to given value.

### HasThirdParty

`func (o *PackagesRequest) HasThirdParty() bool`

HasThirdParty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


