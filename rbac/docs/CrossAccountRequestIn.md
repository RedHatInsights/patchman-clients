# CrossAccountRequestIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetAccount** | **string** |  | 
**StartDate** | **string** |  | 
**EndDate** | **string** |  | 
**Roles** | **[]string** |  | 

## Methods

### NewCrossAccountRequestIn

`func NewCrossAccountRequestIn(targetAccount string, startDate string, endDate string, roles []string, ) *CrossAccountRequestIn`

NewCrossAccountRequestIn instantiates a new CrossAccountRequestIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossAccountRequestInWithDefaults

`func NewCrossAccountRequestInWithDefaults() *CrossAccountRequestIn`

NewCrossAccountRequestInWithDefaults instantiates a new CrossAccountRequestIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetAccount

`func (o *CrossAccountRequestIn) GetTargetAccount() string`

GetTargetAccount returns the TargetAccount field if non-nil, zero value otherwise.

### GetTargetAccountOk

`func (o *CrossAccountRequestIn) GetTargetAccountOk() (*string, bool)`

GetTargetAccountOk returns a tuple with the TargetAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccount

`func (o *CrossAccountRequestIn) SetTargetAccount(v string)`

SetTargetAccount sets TargetAccount field to given value.


### GetStartDate

`func (o *CrossAccountRequestIn) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CrossAccountRequestIn) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CrossAccountRequestIn) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *CrossAccountRequestIn) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CrossAccountRequestIn) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CrossAccountRequestIn) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetRoles

`func (o *CrossAccountRequestIn) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CrossAccountRequestIn) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CrossAccountRequestIn) SetRoles(v []string)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


