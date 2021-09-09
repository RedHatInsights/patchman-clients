# PatchesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrataList** | Pointer to **[]string** |  | [optional] 
**LastChange** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchesResponse

`func NewPatchesResponse() *PatchesResponse`

NewPatchesResponse instantiates a new PatchesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchesResponseWithDefaults

`func NewPatchesResponseWithDefaults() *PatchesResponse`

NewPatchesResponseWithDefaults instantiates a new PatchesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrataList

`func (o *PatchesResponse) GetErrataList() []string`

GetErrataList returns the ErrataList field if non-nil, zero value otherwise.

### GetErrataListOk

`func (o *PatchesResponse) GetErrataListOk() (*[]string, bool)`

GetErrataListOk returns a tuple with the ErrataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrataList

`func (o *PatchesResponse) SetErrataList(v []string)`

SetErrataList sets ErrataList field to given value.

### HasErrataList

`func (o *PatchesResponse) HasErrataList() bool`

HasErrataList returns a boolean if a field has been set.

### GetLastChange

`func (o *PatchesResponse) GetLastChange() string`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *PatchesResponse) GetLastChangeOk() (*string, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *PatchesResponse) SetLastChange(v string)`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *PatchesResponse) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


