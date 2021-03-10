# HostData1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | The number of items on the current page | 
**Page** | **int32** | The page number | 
**PerPage** | **int32** | The number of items to return per page | 
**Results** | [**[]ActiveTag**](ActiveTag.md) |  | 
**Total** | **int32** | Total number of items | 

## Methods

### NewHostData1

`func NewHostData1(count int32, page int32, perPage int32, results []ActiveTag, total int32, ) *HostData1`

NewHostData1 instantiates a new HostData1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostData1WithDefaults

`func NewHostData1WithDefaults() *HostData1`

NewHostData1WithDefaults instantiates a new HostData1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *HostData1) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *HostData1) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *HostData1) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPage

`func (o *HostData1) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *HostData1) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *HostData1) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPerPage

`func (o *HostData1) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *HostData1) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *HostData1) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.


### GetResults

`func (o *HostData1) GetResults() []ActiveTag`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *HostData1) GetResultsOk() (*[]ActiveTag, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *HostData1) SetResults(v []ActiveTag)`

SetResults sets Results field to given value.


### GetTotal

`func (o *HostData1) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *HostData1) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *HostData1) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


