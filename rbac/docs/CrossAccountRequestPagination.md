# CrossAccountRequestPagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 
**Data** | [**[]CrossAccountRequestByAccount**](CrossAccountRequestByAccount.md) |  | 

## Methods

### NewCrossAccountRequestPagination

`func NewCrossAccountRequestPagination(data []CrossAccountRequestByAccount, ) *CrossAccountRequestPagination`

NewCrossAccountRequestPagination instantiates a new CrossAccountRequestPagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossAccountRequestPaginationWithDefaults

`func NewCrossAccountRequestPaginationWithDefaults() *CrossAccountRequestPagination`

NewCrossAccountRequestPaginationWithDefaults instantiates a new CrossAccountRequestPagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *CrossAccountRequestPagination) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CrossAccountRequestPagination) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CrossAccountRequestPagination) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CrossAccountRequestPagination) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *CrossAccountRequestPagination) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CrossAccountRequestPagination) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CrossAccountRequestPagination) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CrossAccountRequestPagination) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetData

`func (o *CrossAccountRequestPagination) GetData() []CrossAccountRequestByAccount`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CrossAccountRequestPagination) GetDataOk() (*[]CrossAccountRequestByAccount, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CrossAccountRequestPagination) SetData(v []CrossAccountRequestByAccount)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


