# CrossAccountRequestDetailByUseId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** |  | [optional] 
**TargetAccount** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Roles** | Pointer to [**[]CrossAccountRequestWithRolesRoles**](CrossAccountRequestWithRolesRoles.md) |  | [optional] 
**UserId** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewCrossAccountRequestDetailByUseId

`func NewCrossAccountRequestDetailByUseId() *CrossAccountRequestDetailByUseId`

NewCrossAccountRequestDetailByUseId instantiates a new CrossAccountRequestDetailByUseId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossAccountRequestDetailByUseIdWithDefaults

`func NewCrossAccountRequestDetailByUseIdWithDefaults() *CrossAccountRequestDetailByUseId`

NewCrossAccountRequestDetailByUseIdWithDefaults instantiates a new CrossAccountRequestDetailByUseId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CrossAccountRequestDetailByUseId) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CrossAccountRequestDetailByUseId) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CrossAccountRequestDetailByUseId) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CrossAccountRequestDetailByUseId) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetTargetAccount

`func (o *CrossAccountRequestDetailByUseId) GetTargetAccount() string`

GetTargetAccount returns the TargetAccount field if non-nil, zero value otherwise.

### GetTargetAccountOk

`func (o *CrossAccountRequestDetailByUseId) GetTargetAccountOk() (*string, bool)`

GetTargetAccountOk returns a tuple with the TargetAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccount

`func (o *CrossAccountRequestDetailByUseId) SetTargetAccount(v string)`

SetTargetAccount sets TargetAccount field to given value.

### HasTargetAccount

`func (o *CrossAccountRequestDetailByUseId) HasTargetAccount() bool`

HasTargetAccount returns a boolean if a field has been set.

### GetStartDate

`func (o *CrossAccountRequestDetailByUseId) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CrossAccountRequestDetailByUseId) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CrossAccountRequestDetailByUseId) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CrossAccountRequestDetailByUseId) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *CrossAccountRequestDetailByUseId) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CrossAccountRequestDetailByUseId) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CrossAccountRequestDetailByUseId) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CrossAccountRequestDetailByUseId) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStatus

`func (o *CrossAccountRequestDetailByUseId) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CrossAccountRequestDetailByUseId) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CrossAccountRequestDetailByUseId) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CrossAccountRequestDetailByUseId) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreated

`func (o *CrossAccountRequestDetailByUseId) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CrossAccountRequestDetailByUseId) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CrossAccountRequestDetailByUseId) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CrossAccountRequestDetailByUseId) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetRoles

`func (o *CrossAccountRequestDetailByUseId) GetRoles() []CrossAccountRequestWithRolesRoles`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CrossAccountRequestDetailByUseId) GetRolesOk() (*[]CrossAccountRequestWithRolesRoles, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CrossAccountRequestDetailByUseId) SetRoles(v []CrossAccountRequestWithRolesRoles)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CrossAccountRequestDetailByUseId) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetUserId

`func (o *CrossAccountRequestDetailByUseId) GetUserId() interface{}`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CrossAccountRequestDetailByUseId) GetUserIdOk() (*interface{}, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CrossAccountRequestDetailByUseId) SetUserId(v interface{})`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CrossAccountRequestDetailByUseId) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *CrossAccountRequestDetailByUseId) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *CrossAccountRequestDetailByUseId) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


