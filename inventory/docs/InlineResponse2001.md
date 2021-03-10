# InlineResponse2001

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | A number of entries on the current page. | [optional] 
**Page** | Pointer to **int32** | A current page number. | [optional] 
**PerPage** | Pointer to **int32** | A page size – a number of entries per single page. | [optional] 
**Results** | Pointer to **map[string]int32** | The list of tags on the systems | [optional] 
**Total** | Pointer to **int32** | Total number of items in the \&quot;data\&quot; list. | [optional] 

## Methods

### NewInlineResponse2001

`func NewInlineResponse2001() *InlineResponse2001`

NewInlineResponse2001 instantiates a new InlineResponse2001 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2001WithDefaults

`func NewInlineResponse2001WithDefaults() *InlineResponse2001`

NewInlineResponse2001WithDefaults instantiates a new InlineResponse2001 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *InlineResponse2001) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *InlineResponse2001) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *InlineResponse2001) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *InlineResponse2001) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPage

`func (o *InlineResponse2001) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *InlineResponse2001) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *InlineResponse2001) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *InlineResponse2001) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPerPage

`func (o *InlineResponse2001) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *InlineResponse2001) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *InlineResponse2001) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *InlineResponse2001) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetResults

`func (o *InlineResponse2001) GetResults() map[string]int32`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *InlineResponse2001) GetResultsOk() (*map[string]int32, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *InlineResponse2001) SetResults(v map[string]int32)`

SetResults sets Results field to given value.

### HasResults

`func (o *InlineResponse2001) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotal

`func (o *InlineResponse2001) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse2001) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse2001) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse2001) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


