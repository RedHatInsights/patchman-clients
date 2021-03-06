# AHostSystemProfileQueryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | A number of entries on the current page. | 
**Page** | **int32** | A current page number. | 
**PerPage** | **int32** | A page size – a number of entries per single page. | 
**Results** | **[]map[string]interface{}** | Actual host search query result entries. | 
**Total** | **int32** | A total count of the found entries. | 

## Methods

### NewAHostSystemProfileQueryResult

`func NewAHostSystemProfileQueryResult(count int32, page int32, perPage int32, results []map[string]interface{}, total int32, ) *AHostSystemProfileQueryResult`

NewAHostSystemProfileQueryResult instantiates a new AHostSystemProfileQueryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAHostSystemProfileQueryResultWithDefaults

`func NewAHostSystemProfileQueryResultWithDefaults() *AHostSystemProfileQueryResult`

NewAHostSystemProfileQueryResultWithDefaults instantiates a new AHostSystemProfileQueryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *AHostSystemProfileQueryResult) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AHostSystemProfileQueryResult) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AHostSystemProfileQueryResult) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPage

`func (o *AHostSystemProfileQueryResult) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *AHostSystemProfileQueryResult) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *AHostSystemProfileQueryResult) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPerPage

`func (o *AHostSystemProfileQueryResult) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *AHostSystemProfileQueryResult) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *AHostSystemProfileQueryResult) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.


### GetResults

`func (o *AHostSystemProfileQueryResult) GetResults() []map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AHostSystemProfileQueryResult) GetResultsOk() (*[]map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AHostSystemProfileQueryResult) SetResults(v []map[string]interface{})`

SetResults sets Results field to given value.


### GetTotal

`func (o *AHostSystemProfileQueryResult) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AHostSystemProfileQueryResult) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AHostSystemProfileQueryResult) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


