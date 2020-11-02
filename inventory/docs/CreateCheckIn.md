# CreateCheckIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanonicalFacts** | **map[string]interface{}** | A set of string facts about a host. | 
**CheckinFrequency** | Pointer to **int32** | Defines how far in the future the host becomes stale (in minutes). | [optional] [default to 1440]

## Methods

### NewCreateCheckIn

`func NewCreateCheckIn(canonicalFacts map[string]interface{}, ) *CreateCheckIn`

NewCreateCheckIn instantiates a new CreateCheckIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCheckInWithDefaults

`func NewCreateCheckInWithDefaults() *CreateCheckIn`

NewCreateCheckInWithDefaults instantiates a new CreateCheckIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanonicalFacts

`func (o *CreateCheckIn) GetCanonicalFacts() map[string]interface{}`

GetCanonicalFacts returns the CanonicalFacts field if non-nil, zero value otherwise.

### GetCanonicalFactsOk

`func (o *CreateCheckIn) GetCanonicalFactsOk() (*map[string]interface{}, bool)`

GetCanonicalFactsOk returns a tuple with the CanonicalFacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalFacts

`func (o *CreateCheckIn) SetCanonicalFacts(v map[string]interface{})`

SetCanonicalFacts sets CanonicalFacts field to given value.


### GetCheckinFrequency

`func (o *CreateCheckIn) GetCheckinFrequency() int32`

GetCheckinFrequency returns the CheckinFrequency field if non-nil, zero value otherwise.

### GetCheckinFrequencyOk

`func (o *CreateCheckIn) GetCheckinFrequencyOk() (*int32, bool)`

GetCheckinFrequencyOk returns a tuple with the CheckinFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckinFrequency

`func (o *CreateCheckIn) SetCheckinFrequency(v int32)`

SetCheckinFrequency sets CheckinFrequency field to given value.

### HasCheckinFrequency

`func (o *CreateCheckIn) HasCheckinFrequency() bool`

HasCheckinFrequency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


