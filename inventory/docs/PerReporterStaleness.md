# PerReporterStaleness

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastCheckIn** | Pointer to **string** |  | [optional] 
**StaleTimestamp** | Pointer to **string** |  | [optional] 
**CheckInSucceeded** | Pointer to **bool** |  | [optional] 

## Methods

### NewPerReporterStaleness

`func NewPerReporterStaleness() *PerReporterStaleness`

NewPerReporterStaleness instantiates a new PerReporterStaleness object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerReporterStalenessWithDefaults

`func NewPerReporterStalenessWithDefaults() *PerReporterStaleness`

NewPerReporterStalenessWithDefaults instantiates a new PerReporterStaleness object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastCheckIn

`func (o *PerReporterStaleness) GetLastCheckIn() string`

GetLastCheckIn returns the LastCheckIn field if non-nil, zero value otherwise.

### GetLastCheckInOk

`func (o *PerReporterStaleness) GetLastCheckInOk() (*string, bool)`

GetLastCheckInOk returns a tuple with the LastCheckIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckIn

`func (o *PerReporterStaleness) SetLastCheckIn(v string)`

SetLastCheckIn sets LastCheckIn field to given value.

### HasLastCheckIn

`func (o *PerReporterStaleness) HasLastCheckIn() bool`

HasLastCheckIn returns a boolean if a field has been set.

### GetStaleTimestamp

`func (o *PerReporterStaleness) GetStaleTimestamp() string`

GetStaleTimestamp returns the StaleTimestamp field if non-nil, zero value otherwise.

### GetStaleTimestampOk

`func (o *PerReporterStaleness) GetStaleTimestampOk() (*string, bool)`

GetStaleTimestampOk returns a tuple with the StaleTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleTimestamp

`func (o *PerReporterStaleness) SetStaleTimestamp(v string)`

SetStaleTimestamp sets StaleTimestamp field to given value.

### HasStaleTimestamp

`func (o *PerReporterStaleness) HasStaleTimestamp() bool`

HasStaleTimestamp returns a boolean if a field has been set.

### GetCheckInSucceeded

`func (o *PerReporterStaleness) GetCheckInSucceeded() bool`

GetCheckInSucceeded returns the CheckInSucceeded field if non-nil, zero value otherwise.

### GetCheckInSucceededOk

`func (o *PerReporterStaleness) GetCheckInSucceededOk() (*bool, bool)`

GetCheckInSucceededOk returns a tuple with the CheckInSucceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInSucceeded

`func (o *PerReporterStaleness) SetCheckInSucceeded(v bool)`

SetCheckInSucceeded sets CheckInSucceeded field to given value.

### HasCheckInSucceeded

`func (o *PerReporterStaleness) HasCheckInSucceeded() bool`

HasCheckInSucceeded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


