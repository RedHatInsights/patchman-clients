# CreateHostOutAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | A Red Hat Account number that owns the host. | 
**AnsibleHost** | Pointer to **NullableString** | The ansible host name for remediations | [optional] 
**Created** | Pointer to **time.Time** | A timestamp when the entry was created. | [optional] 
**CulledTimestamp** | Pointer to **NullableTime** | Timestamp from which the host is considered deleted. | [optional] 
**DisplayName** | Pointer to **NullableString** | A host’s human-readable display name, e.g. in a form of a domain name. | [optional] 
**Facts** | Pointer to **[]map[string]interface{}** | A set of facts belonging to the host. | [optional] 
**Id** | Pointer to **string** | A durable and reliable platform-wide host identifier. Applications should use this identifier to reference hosts. | [optional] 
**Reporter** | Pointer to **NullableString** | Reporting source of the host. Used when updating the stale_timestamp. | [optional] 
**StaleTimestamp** | Pointer to **NullableTime** | Timestamp from which the host is considered stale. | [optional] 
**StaleWarningTimestamp** | Pointer to **NullableTime** | Timestamp from which the host is considered too stale to be listed without an explicit toggle. | [optional] 
**Updated** | Pointer to **time.Time** | A timestamp when the entry was last updated. | [optional] 

## Methods

### NewCreateHostOutAllOf

`func NewCreateHostOutAllOf(account string, ) *CreateHostOutAllOf`

NewCreateHostOutAllOf instantiates a new CreateHostOutAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHostOutAllOfWithDefaults

`func NewCreateHostOutAllOfWithDefaults() *CreateHostOutAllOf`

NewCreateHostOutAllOfWithDefaults instantiates a new CreateHostOutAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CreateHostOutAllOf) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CreateHostOutAllOf) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CreateHostOutAllOf) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetAnsibleHost

`func (o *CreateHostOutAllOf) GetAnsibleHost() string`

GetAnsibleHost returns the AnsibleHost field if non-nil, zero value otherwise.

### GetAnsibleHostOk

`func (o *CreateHostOutAllOf) GetAnsibleHostOk() (*string, bool)`

GetAnsibleHostOk returns a tuple with the AnsibleHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleHost

`func (o *CreateHostOutAllOf) SetAnsibleHost(v string)`

SetAnsibleHost sets AnsibleHost field to given value.

### HasAnsibleHost

`func (o *CreateHostOutAllOf) HasAnsibleHost() bool`

HasAnsibleHost returns a boolean if a field has been set.

### SetAnsibleHostNil

`func (o *CreateHostOutAllOf) SetAnsibleHostNil(b bool)`

 SetAnsibleHostNil sets the value for AnsibleHost to be an explicit nil

### UnsetAnsibleHost
`func (o *CreateHostOutAllOf) UnsetAnsibleHost()`

UnsetAnsibleHost ensures that no value is present for AnsibleHost, not even an explicit nil
### GetCreated

`func (o *CreateHostOutAllOf) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CreateHostOutAllOf) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CreateHostOutAllOf) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CreateHostOutAllOf) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCulledTimestamp

`func (o *CreateHostOutAllOf) GetCulledTimestamp() time.Time`

GetCulledTimestamp returns the CulledTimestamp field if non-nil, zero value otherwise.

### GetCulledTimestampOk

`func (o *CreateHostOutAllOf) GetCulledTimestampOk() (*time.Time, bool)`

GetCulledTimestampOk returns a tuple with the CulledTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCulledTimestamp

`func (o *CreateHostOutAllOf) SetCulledTimestamp(v time.Time)`

SetCulledTimestamp sets CulledTimestamp field to given value.

### HasCulledTimestamp

`func (o *CreateHostOutAllOf) HasCulledTimestamp() bool`

HasCulledTimestamp returns a boolean if a field has been set.

### SetCulledTimestampNil

`func (o *CreateHostOutAllOf) SetCulledTimestampNil(b bool)`

 SetCulledTimestampNil sets the value for CulledTimestamp to be an explicit nil

### UnsetCulledTimestamp
`func (o *CreateHostOutAllOf) UnsetCulledTimestamp()`

UnsetCulledTimestamp ensures that no value is present for CulledTimestamp, not even an explicit nil
### GetDisplayName

`func (o *CreateHostOutAllOf) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateHostOutAllOf) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateHostOutAllOf) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateHostOutAllOf) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CreateHostOutAllOf) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CreateHostOutAllOf) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetFacts

`func (o *CreateHostOutAllOf) GetFacts() []map[string]interface{}`

GetFacts returns the Facts field if non-nil, zero value otherwise.

### GetFactsOk

`func (o *CreateHostOutAllOf) GetFactsOk() (*[]map[string]interface{}, bool)`

GetFactsOk returns a tuple with the Facts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacts

`func (o *CreateHostOutAllOf) SetFacts(v []map[string]interface{})`

SetFacts sets Facts field to given value.

### HasFacts

`func (o *CreateHostOutAllOf) HasFacts() bool`

HasFacts returns a boolean if a field has been set.

### GetId

`func (o *CreateHostOutAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateHostOutAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateHostOutAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateHostOutAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReporter

`func (o *CreateHostOutAllOf) GetReporter() string`

GetReporter returns the Reporter field if non-nil, zero value otherwise.

### GetReporterOk

`func (o *CreateHostOutAllOf) GetReporterOk() (*string, bool)`

GetReporterOk returns a tuple with the Reporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporter

`func (o *CreateHostOutAllOf) SetReporter(v string)`

SetReporter sets Reporter field to given value.

### HasReporter

`func (o *CreateHostOutAllOf) HasReporter() bool`

HasReporter returns a boolean if a field has been set.

### SetReporterNil

`func (o *CreateHostOutAllOf) SetReporterNil(b bool)`

 SetReporterNil sets the value for Reporter to be an explicit nil

### UnsetReporter
`func (o *CreateHostOutAllOf) UnsetReporter()`

UnsetReporter ensures that no value is present for Reporter, not even an explicit nil
### GetStaleTimestamp

`func (o *CreateHostOutAllOf) GetStaleTimestamp() time.Time`

GetStaleTimestamp returns the StaleTimestamp field if non-nil, zero value otherwise.

### GetStaleTimestampOk

`func (o *CreateHostOutAllOf) GetStaleTimestampOk() (*time.Time, bool)`

GetStaleTimestampOk returns a tuple with the StaleTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleTimestamp

`func (o *CreateHostOutAllOf) SetStaleTimestamp(v time.Time)`

SetStaleTimestamp sets StaleTimestamp field to given value.

### HasStaleTimestamp

`func (o *CreateHostOutAllOf) HasStaleTimestamp() bool`

HasStaleTimestamp returns a boolean if a field has been set.

### SetStaleTimestampNil

`func (o *CreateHostOutAllOf) SetStaleTimestampNil(b bool)`

 SetStaleTimestampNil sets the value for StaleTimestamp to be an explicit nil

### UnsetStaleTimestamp
`func (o *CreateHostOutAllOf) UnsetStaleTimestamp()`

UnsetStaleTimestamp ensures that no value is present for StaleTimestamp, not even an explicit nil
### GetStaleWarningTimestamp

`func (o *CreateHostOutAllOf) GetStaleWarningTimestamp() time.Time`

GetStaleWarningTimestamp returns the StaleWarningTimestamp field if non-nil, zero value otherwise.

### GetStaleWarningTimestampOk

`func (o *CreateHostOutAllOf) GetStaleWarningTimestampOk() (*time.Time, bool)`

GetStaleWarningTimestampOk returns a tuple with the StaleWarningTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleWarningTimestamp

`func (o *CreateHostOutAllOf) SetStaleWarningTimestamp(v time.Time)`

SetStaleWarningTimestamp sets StaleWarningTimestamp field to given value.

### HasStaleWarningTimestamp

`func (o *CreateHostOutAllOf) HasStaleWarningTimestamp() bool`

HasStaleWarningTimestamp returns a boolean if a field has been set.

### SetStaleWarningTimestampNil

`func (o *CreateHostOutAllOf) SetStaleWarningTimestampNil(b bool)`

 SetStaleWarningTimestampNil sets the value for StaleWarningTimestamp to be an explicit nil

### UnsetStaleWarningTimestamp
`func (o *CreateHostOutAllOf) UnsetStaleWarningTimestamp()`

UnsetStaleWarningTimestamp ensures that no value is present for StaleWarningTimestamp, not even an explicit nil
### GetUpdated

`func (o *CreateHostOutAllOf) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CreateHostOutAllOf) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CreateHostOutAllOf) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CreateHostOutAllOf) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


