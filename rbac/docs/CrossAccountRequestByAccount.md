# CrossAccountRequestByAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** |  | [optional] 
**TargetAccount** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **interface{}** |  | [optional] 
**EndDate** | Pointer to **interface{}** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 

## Methods

### NewCrossAccountRequestByAccount

`func NewCrossAccountRequestByAccount() *CrossAccountRequestByAccount`

NewCrossAccountRequestByAccount instantiates a new CrossAccountRequestByAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossAccountRequestByAccountWithDefaults

`func NewCrossAccountRequestByAccountWithDefaults() *CrossAccountRequestByAccount`

NewCrossAccountRequestByAccountWithDefaults instantiates a new CrossAccountRequestByAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CrossAccountRequestByAccount) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CrossAccountRequestByAccount) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CrossAccountRequestByAccount) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CrossAccountRequestByAccount) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetTargetAccount

`func (o *CrossAccountRequestByAccount) GetTargetAccount() string`

GetTargetAccount returns the TargetAccount field if non-nil, zero value otherwise.

### GetTargetAccountOk

`func (o *CrossAccountRequestByAccount) GetTargetAccountOk() (*string, bool)`

GetTargetAccountOk returns a tuple with the TargetAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccount

`func (o *CrossAccountRequestByAccount) SetTargetAccount(v string)`

SetTargetAccount sets TargetAccount field to given value.

### HasTargetAccount

`func (o *CrossAccountRequestByAccount) HasTargetAccount() bool`

HasTargetAccount returns a boolean if a field has been set.

### GetStatus

`func (o *CrossAccountRequestByAccount) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CrossAccountRequestByAccount) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CrossAccountRequestByAccount) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CrossAccountRequestByAccount) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreated

`func (o *CrossAccountRequestByAccount) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CrossAccountRequestByAccount) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CrossAccountRequestByAccount) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CrossAccountRequestByAccount) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetStartDate

`func (o *CrossAccountRequestByAccount) GetStartDate() interface{}`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CrossAccountRequestByAccount) GetStartDateOk() (*interface{}, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CrossAccountRequestByAccount) SetStartDate(v interface{})`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CrossAccountRequestByAccount) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *CrossAccountRequestByAccount) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *CrossAccountRequestByAccount) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *CrossAccountRequestByAccount) GetEndDate() interface{}`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CrossAccountRequestByAccount) GetEndDateOk() (*interface{}, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CrossAccountRequestByAccount) SetEndDate(v interface{})`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CrossAccountRequestByAccount) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *CrossAccountRequestByAccount) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *CrossAccountRequestByAccount) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetFirstName

`func (o *CrossAccountRequestByAccount) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CrossAccountRequestByAccount) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CrossAccountRequestByAccount) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CrossAccountRequestByAccount) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *CrossAccountRequestByAccount) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CrossAccountRequestByAccount) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CrossAccountRequestByAccount) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CrossAccountRequestByAccount) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *CrossAccountRequestByAccount) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CrossAccountRequestByAccount) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CrossAccountRequestByAccount) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CrossAccountRequestByAccount) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


