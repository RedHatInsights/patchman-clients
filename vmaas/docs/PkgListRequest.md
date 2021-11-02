# PkgListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **float32** |  | [optional] 
**PageSize** | Pointer to **float32** |  | [optional] 
**ModifiedSince** | Pointer to **string** |  | [optional] 
**ReturnModified** | Pointer to **bool** | Include &#39;modified&#39; package attribute into the response | [optional] [default to false]

## Methods

### NewPkgListRequest

`func NewPkgListRequest() *PkgListRequest`

NewPkgListRequest instantiates a new PkgListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkgListRequestWithDefaults

`func NewPkgListRequestWithDefaults() *PkgListRequest`

NewPkgListRequestWithDefaults instantiates a new PkgListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PkgListRequest) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PkgListRequest) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PkgListRequest) SetPage(v float32)`

SetPage sets Page field to given value.

### HasPage

`func (o *PkgListRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *PkgListRequest) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PkgListRequest) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PkgListRequest) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *PkgListRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetModifiedSince

`func (o *PkgListRequest) GetModifiedSince() string`

GetModifiedSince returns the ModifiedSince field if non-nil, zero value otherwise.

### GetModifiedSinceOk

`func (o *PkgListRequest) GetModifiedSinceOk() (*string, bool)`

GetModifiedSinceOk returns a tuple with the ModifiedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedSince

`func (o *PkgListRequest) SetModifiedSince(v string)`

SetModifiedSince sets ModifiedSince field to given value.

### HasModifiedSince

`func (o *PkgListRequest) HasModifiedSince() bool`

HasModifiedSince returns a boolean if a field has been set.

### GetReturnModified

`func (o *PkgListRequest) GetReturnModified() bool`

GetReturnModified returns the ReturnModified field if non-nil, zero value otherwise.

### GetReturnModifiedOk

`func (o *PkgListRequest) GetReturnModifiedOk() (*bool, bool)`

GetReturnModifiedOk returns a tuple with the ReturnModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnModified

`func (o *PkgListRequest) SetReturnModified(v bool)`

SetReturnModified sets ReturnModified field to given value.

### HasReturnModified

`func (o *PkgListRequest) HasReturnModified() bool`

HasReturnModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


