# VulnerabilitiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CveList** | [**[]OneOfstringobject**](OneOfstringobject.md) |  | 
**ManuallyFixableCveList** | [**[]OneOfstringobject**](OneOfstringobject.md) |  | 
**UnpatchedCveList** | [**[]OneOfstringobject**](OneOfstringobject.md) |  | 
**LastChange** | Pointer to **string** |  | [optional] 

## Methods

### NewVulnerabilitiesResponse

`func NewVulnerabilitiesResponse(cveList []OneOfstringobject, manuallyFixableCveList []OneOfstringobject, unpatchedCveList []OneOfstringobject, ) *VulnerabilitiesResponse`

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

`func (o *VulnerabilitiesResponse) GetCveList() []OneOfstringobject`

GetCveList returns the CveList field if non-nil, zero value otherwise.

### GetCveListOk

`func (o *VulnerabilitiesResponse) GetCveListOk() (*[]OneOfstringobject, bool)`

GetCveListOk returns a tuple with the CveList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveList

`func (o *VulnerabilitiesResponse) SetCveList(v []OneOfstringobject)`

SetCveList sets CveList field to given value.


### GetManuallyFixableCveList

`func (o *VulnerabilitiesResponse) GetManuallyFixableCveList() []OneOfstringobject`

GetManuallyFixableCveList returns the ManuallyFixableCveList field if non-nil, zero value otherwise.

### GetManuallyFixableCveListOk

`func (o *VulnerabilitiesResponse) GetManuallyFixableCveListOk() (*[]OneOfstringobject, bool)`

GetManuallyFixableCveListOk returns a tuple with the ManuallyFixableCveList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyFixableCveList

`func (o *VulnerabilitiesResponse) SetManuallyFixableCveList(v []OneOfstringobject)`

SetManuallyFixableCveList sets ManuallyFixableCveList field to given value.


### GetUnpatchedCveList

`func (o *VulnerabilitiesResponse) GetUnpatchedCveList() []OneOfstringobject`

GetUnpatchedCveList returns the UnpatchedCveList field if non-nil, zero value otherwise.

### GetUnpatchedCveListOk

`func (o *VulnerabilitiesResponse) GetUnpatchedCveListOk() (*[]OneOfstringobject, bool)`

GetUnpatchedCveListOk returns a tuple with the UnpatchedCveList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpatchedCveList

`func (o *VulnerabilitiesResponse) SetUnpatchedCveList(v []OneOfstringobject)`

SetUnpatchedCveList sets UnpatchedCveList field to given value.


### GetLastChange

`func (o *VulnerabilitiesResponse) GetLastChange() string`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *VulnerabilitiesResponse) GetLastChangeOk() (*string, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *VulnerabilitiesResponse) SetLastChange(v string)`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *VulnerabilitiesResponse) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


