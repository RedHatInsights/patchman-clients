# CreateCheckInAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckinFrequency** | Pointer to **int32** | How long from now to expect another check-in (in minutes). | [optional] 

## Methods

### NewCreateCheckInAllOf

`func NewCreateCheckInAllOf() *CreateCheckInAllOf`

NewCreateCheckInAllOf instantiates a new CreateCheckInAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCheckInAllOfWithDefaults

`func NewCreateCheckInAllOfWithDefaults() *CreateCheckInAllOf`

NewCreateCheckInAllOfWithDefaults instantiates a new CreateCheckInAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckinFrequency

`func (o *CreateCheckInAllOf) GetCheckinFrequency() int32`

GetCheckinFrequency returns the CheckinFrequency field if non-nil, zero value otherwise.

### GetCheckinFrequencyOk

`func (o *CreateCheckInAllOf) GetCheckinFrequencyOk() (*int32, bool)`

GetCheckinFrequencyOk returns a tuple with the CheckinFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckinFrequency

`func (o *CreateCheckInAllOf) SetCheckinFrequency(v int32)`

SetCheckinFrequency sets CheckinFrequency field to given value.

### HasCheckinFrequency

`func (o *CreateCheckInAllOf) HasCheckinFrequency() bool`

HasCheckinFrequency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


