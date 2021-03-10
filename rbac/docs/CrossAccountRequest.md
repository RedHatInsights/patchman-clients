# CrossAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** |  | [optional] 
**TargetAccount** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**StartDate** | Pointer to **interface{}** |  | [optional] 
**EndDate** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewCrossAccountRequest

`func NewCrossAccountRequest() *CrossAccountRequest`

NewCrossAccountRequest instantiates a new CrossAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossAccountRequestWithDefaults

`func NewCrossAccountRequestWithDefaults() *CrossAccountRequest`

NewCrossAccountRequestWithDefaults instantiates a new CrossAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CrossAccountRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CrossAccountRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CrossAccountRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CrossAccountRequest) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetTargetAccount

`func (o *CrossAccountRequest) GetTargetAccount() string`

GetTargetAccount returns the TargetAccount field if non-nil, zero value otherwise.

### GetTargetAccountOk

`func (o *CrossAccountRequest) GetTargetAccountOk() (*string, bool)`

GetTargetAccountOk returns a tuple with the TargetAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccount

`func (o *CrossAccountRequest) SetTargetAccount(v string)`

SetTargetAccount sets TargetAccount field to given value.

### HasTargetAccount

`func (o *CrossAccountRequest) HasTargetAccount() bool`

HasTargetAccount returns a boolean if a field has been set.

### GetStatus

`func (o *CrossAccountRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CrossAccountRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CrossAccountRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CrossAccountRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreated

`func (o *CrossAccountRequest) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CrossAccountRequest) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CrossAccountRequest) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CrossAccountRequest) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetStartDate

`func (o *CrossAccountRequest) GetStartDate() interface{}`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CrossAccountRequest) GetStartDateOk() (*interface{}, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CrossAccountRequest) SetStartDate(v interface{})`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CrossAccountRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *CrossAccountRequest) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *CrossAccountRequest) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *CrossAccountRequest) GetEndDate() interface{}`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CrossAccountRequest) GetEndDateOk() (*interface{}, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CrossAccountRequest) SetEndDate(v interface{})`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CrossAccountRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *CrossAccountRequest) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *CrossAccountRequest) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


