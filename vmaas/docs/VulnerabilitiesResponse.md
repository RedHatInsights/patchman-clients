# VulnerabilitiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CveList** | **[]string** |  | 
**UnpatchedCveList** | **[]string** |  | 

## Methods

### NewVulnerabilitiesResponse

`func NewVulnerabilitiesResponse(cveList []string, unpatchedCveList []string, ) *VulnerabilitiesResponse`

NewVulnerabilitiesResponse instantiates a new VulnerabilitiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVulnerabilitiesResponseWithDefaults

`func NewVulnerabilitiesResponseWithDefaults() *VulnerabilitiesResponse`

NewVulnerabilitiesResponseWithDefaults instantiates a new VulnerabilitiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCveList

`func (o *VulnerabilitiesResponse) GetCveList() []string`

GetCveList returns the CveList field if non-nil, zero value otherwise.

### GetCveListOk

`func (o *VulnerabilitiesResponse) GetCveListOk() (*[]string, bool)`

GetCveListOk returns a tuple with the CveList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveList

`func (o *VulnerabilitiesResponse) SetCveList(v []string)`

SetCveList sets CveList field to given value.


### GetUnpatchedCveList

`func (o *VulnerabilitiesResponse) GetUnpatchedCveList() []string`

GetUnpatchedCveList returns the UnpatchedCveList field if non-nil, zero value otherwise.

### GetUnpatchedCveListOk

`func (o *VulnerabilitiesResponse) GetUnpatchedCveListOk() (*[]string, bool)`

GetUnpatchedCveListOk returns a tuple with the UnpatchedCveList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpatchedCveList

`func (o *VulnerabilitiesResponse) SetUnpatchedCveList(v []string)`

SetUnpatchedCveList sets UnpatchedCveList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


