# CrossAccountRequestWithRoles

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

## Methods

### NewCrossAccountRequestWithRoles

`func NewCrossAccountRequestWithRoles() *CrossAccountRequestWithRoles`

NewCrossAccountRequestWithRoles instantiates a new CrossAccountRequestWithRoles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossAccountRequestWithRolesWithDefaults

`func NewCrossAccountRequestWithRolesWithDefaults() *CrossAccountRequestWithRoles`

NewCrossAccountRequestWithRolesWithDefaults instantiates a new CrossAccountRequestWithRoles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CrossAccountRequestWithRoles) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CrossAccountRequestWithRoles) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CrossAccountRequestWithRoles) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CrossAccountRequestWithRoles) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetTargetAccount

`func (o *CrossAccountRequestWithRoles) GetTargetAccount() string`

GetTargetAccount returns the TargetAccount field if non-nil, zero value otherwise.

### GetTargetAccountOk

`func (o *CrossAccountRequestWithRoles) GetTargetAccountOk() (*string, bool)`

GetTargetAccountOk returns a tuple with the TargetAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccount

`func (o *CrossAccountRequestWithRoles) SetTargetAccount(v string)`

SetTargetAccount sets TargetAccount field to given value.

### HasTargetAccount

`func (o *CrossAccountRequestWithRoles) HasTargetAccount() bool`

HasTargetAccount returns a boolean if a field has been set.

### GetStartDate

`func (o *CrossAccountRequestWithRoles) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CrossAccountRequestWithRoles) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CrossAccountRequestWithRoles) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CrossAccountRequestWithRoles) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *CrossAccountRequestWithRoles) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CrossAccountRequestWithRoles) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CrossAccountRequestWithRoles) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CrossAccountRequestWithRoles) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStatus

`func (o *CrossAccountRequestWithRoles) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CrossAccountRequestWithRoles) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CrossAccountRequestWithRoles) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CrossAccountRequestWithRoles) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreated

`func (o *CrossAccountRequestWithRoles) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CrossAccountRequestWithRoles) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CrossAccountRequestWithRoles) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CrossAccountRequestWithRoles) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetRoles

`func (o *CrossAccountRequestWithRoles) GetRoles() []CrossAccountRequestWithRolesRoles`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CrossAccountRequestWithRoles) GetRolesOk() (*[]CrossAccountRequestWithRolesRoles, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CrossAccountRequestWithRoles) SetRoles(v []CrossAccountRequestWithRolesRoles)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CrossAccountRequestWithRoles) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


