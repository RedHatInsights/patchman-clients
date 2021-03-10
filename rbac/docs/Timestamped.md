# Timestamped

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** |  | 
**Modified** | **time.Time** |  | 

## Methods

### NewTimestamped

`func NewTimestamped(created time.Time, modified time.Time, ) *Timestamped`

NewTimestamped instantiates a new Timestamped object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimestampedWithDefaults

`func NewTimestampedWithDefaults() *Timestamped`

NewTimestampedWithDefaults instantiates a new Timestamped object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *Timestamped) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Timestamped) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Timestamped) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *Timestamped) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Timestamped) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Timestamped) SetModified(v time.Time)`

SetModified sets Modified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


