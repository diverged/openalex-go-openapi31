# ListWorks200ResponseResultsInnerGrantsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Funder** | **string** | The OpenAlex ID for the funder of this grant. | 
**FunderDisplayName** | **string** | The display name of the funder. | 
**AwardId** | Pointer to **string** | The identifier for this specific grant, as given by the funder. | [optional] 

## Methods

### NewListWorks200ResponseResultsInnerGrantsInner

`func NewListWorks200ResponseResultsInnerGrantsInner(funder string, funderDisplayName string, ) *ListWorks200ResponseResultsInnerGrantsInner`

NewListWorks200ResponseResultsInnerGrantsInner instantiates a new ListWorks200ResponseResultsInnerGrantsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorks200ResponseResultsInnerGrantsInnerWithDefaults

`func NewListWorks200ResponseResultsInnerGrantsInnerWithDefaults() *ListWorks200ResponseResultsInnerGrantsInner`

NewListWorks200ResponseResultsInnerGrantsInnerWithDefaults instantiates a new ListWorks200ResponseResultsInnerGrantsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunder

`func (o *ListWorks200ResponseResultsInnerGrantsInner) GetFunder() string`

GetFunder returns the Funder field if non-nil, zero value otherwise.

### GetFunderOk

`func (o *ListWorks200ResponseResultsInnerGrantsInner) GetFunderOk() (*string, bool)`

GetFunderOk returns a tuple with the Funder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunder

`func (o *ListWorks200ResponseResultsInnerGrantsInner) SetFunder(v string)`

SetFunder sets Funder field to given value.


### GetFunderDisplayName

`func (o *ListWorks200ResponseResultsInnerGrantsInner) GetFunderDisplayName() string`

GetFunderDisplayName returns the FunderDisplayName field if non-nil, zero value otherwise.

### GetFunderDisplayNameOk

`func (o *ListWorks200ResponseResultsInnerGrantsInner) GetFunderDisplayNameOk() (*string, bool)`

GetFunderDisplayNameOk returns a tuple with the FunderDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunderDisplayName

`func (o *ListWorks200ResponseResultsInnerGrantsInner) SetFunderDisplayName(v string)`

SetFunderDisplayName sets FunderDisplayName field to given value.


### GetAwardId

`func (o *ListWorks200ResponseResultsInnerGrantsInner) GetAwardId() string`

GetAwardId returns the AwardId field if non-nil, zero value otherwise.

### GetAwardIdOk

`func (o *ListWorks200ResponseResultsInnerGrantsInner) GetAwardIdOk() (*string, bool)`

GetAwardIdOk returns a tuple with the AwardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwardId

`func (o *ListWorks200ResponseResultsInnerGrantsInner) SetAwardId(v string)`

SetAwardId sets AwardId field to given value.

### HasAwardId

`func (o *ListWorks200ResponseResultsInnerGrantsInner) HasAwardId() bool`

HasAwardId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


