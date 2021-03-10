# CrossAccountRequestDetailByAccount

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
**FirstName** | Pointer to **interface{}** |  | [optional] 
**LastName** | Pointer to **interface{}** |  | [optional] 
**Email** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewCrossAccountRequestDetailByAccount

`func NewCrossAccountRequestDetailByAccount() *CrossAccountRequestDetailByAccount`

NewCrossAccountRequestDetailByAccount instantiates a new CrossAccountRequestDetailByAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossAccountRequestDetailByAccountWithDefaults

`func NewCrossAccountRequestDetailByAccountWithDefaults() *CrossAccountRequestDetailByAccount`

NewCrossAccountRequestDetailByAccountWithDefaults instantiates a new CrossAccountRequestDetailByAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CrossAccountRequestDetailByAccount) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CrossAccountRequestDetailByAccount) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CrossAccountRequestDetailByAccount) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CrossAccountRequestDetailByAccount) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetTargetAccount

`func (o *CrossAccountRequestDetailByAccount) GetTargetAccount() string`

GetTargetAccount returns the TargetAccount field if non-nil, zero value otherwise.

### GetTargetAccountOk

`func (o *CrossAccountRequestDetailByAccount) GetTargetAccountOk() (*string, bool)`

GetTargetAccountOk returns a tuple with the TargetAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccount

`func (o *CrossAccountRequestDetailByAccount) SetTargetAccount(v string)`

SetTargetAccount sets TargetAccount field to given value.

### HasTargetAccount

`func (o *CrossAccountRequestDetailByAccount) HasTargetAccount() bool`

HasTargetAccount returns a boolean if a field has been set.

### GetStartDate

`func (o *CrossAccountRequestDetailByAccount) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CrossAccountRequestDetailByAccount) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CrossAccountRequestDetailByAccount) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CrossAccountRequestDetailByAccount) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *CrossAccountRequestDetailByAccount) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CrossAccountRequestDetailByAccount) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CrossAccountRequestDetailByAccount) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CrossAccountRequestDetailByAccount) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStatus

`func (o *CrossAccountRequestDetailByAccount) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CrossAccountRequestDetailByAccount) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CrossAccountRequestDetailByAccount) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CrossAccountRequestDetailByAccount) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreated

`func (o *CrossAccountRequestDetailByAccount) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CrossAccountRequestDetailByAccount) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CrossAccountRequestDetailByAccount) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CrossAccountRequestDetailByAccount) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetRoles

`func (o *CrossAccountRequestDetailByAccount) GetRoles() []CrossAccountRequestWithRolesRoles`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CrossAccountRequestDetailByAccount) GetRolesOk() (*[]CrossAccountRequestWithRolesRoles, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CrossAccountRequestDetailByAccount) SetRoles(v []CrossAccountRequestWithRolesRoles)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CrossAccountRequestDetailByAccount) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetFirstName

`func (o *CrossAccountRequestDetailByAccount) GetFirstName() interface{}`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CrossAccountRequestDetailByAccount) GetFirstNameOk() (*interface{}, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CrossAccountRequestDetailByAccount) SetFirstName(v interface{})`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CrossAccountRequestDetailByAccount) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *CrossAccountRequestDetailByAccount) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *CrossAccountRequestDetailByAccount) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *CrossAccountRequestDetailByAccount) GetLastName() interface{}`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CrossAccountRequestDetailByAccount) GetLastNameOk() (*interface{}, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CrossAccountRequestDetailByAccount) SetLastName(v interface{})`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CrossAccountRequestDetailByAccount) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *CrossAccountRequestDetailByAccount) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *CrossAccountRequestDetailByAccount) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetEmail

`func (o *CrossAccountRequestDetailByAccount) GetEmail() interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CrossAccountRequestDetailByAccount) GetEmailOk() (*interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CrossAccountRequestDetailByAccount) SetEmail(v interface{})`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CrossAccountRequestDetailByAccount) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CrossAccountRequestDetailByAccount) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CrossAccountRequestDetailByAccount) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


