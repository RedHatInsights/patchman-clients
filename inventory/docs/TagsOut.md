# TagsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | The number of items on the current page | 
**Page** | **int32** | The page number | 
**PerPage** | **int32** | The number of items to return per page | 
**Total** | **int32** | Total number of items | 
**Results** | Pointer to [**map[string][]StructuredTag**](array.md) | The list of tags on the systems | [optional] 

## Methods

### NewTagsOut

`func NewTagsOut(count int32, page int32, perPage int32, total int32, ) *TagsOut`

NewTagsOut instantiates a new TagsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagsOutWithDefaults

`func NewTagsOutWithDefaults() *TagsOut`

NewTagsOutWithDefaults instantiates a new TagsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *TagsOut) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TagsOut) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TagsOut) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPage

`func (o *TagsOut) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *TagsOut) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *TagsOut) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPerPage

`func (o *TagsOut) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *TagsOut) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *TagsOut) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.


### GetTotal

`func (o *TagsOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TagsOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TagsOut) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetResults

`func (o *TagsOut) GetResults() map[string][]StructuredTag`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *TagsOut) GetResultsOk() (*map[string][]StructuredTag, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *TagsOut) SetResults(v map[string][]StructuredTag)`

SetResults sets Results field to given value.

### HasResults

`func (o *TagsOut) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


