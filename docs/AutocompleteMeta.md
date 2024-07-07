# AutocompleteMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of matching results. | [optional] 
**DbResponseTimeMs** | Pointer to **int32** | The time taken by the database to respond, in milliseconds. | [optional] 
**Page** | Pointer to **int32** | The current page number (always 1 for autocomplete). | [optional] 
**PerPage** | Pointer to **int32** | The number of results per page (always 10 for autocomplete). | [optional] 

## Methods

### NewAutocompleteMeta

`func NewAutocompleteMeta() *AutocompleteMeta`

NewAutocompleteMeta instantiates a new AutocompleteMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutocompleteMetaWithDefaults

`func NewAutocompleteMetaWithDefaults() *AutocompleteMeta`

NewAutocompleteMetaWithDefaults instantiates a new AutocompleteMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *AutocompleteMeta) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AutocompleteMeta) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AutocompleteMeta) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *AutocompleteMeta) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDbResponseTimeMs

`func (o *AutocompleteMeta) GetDbResponseTimeMs() int32`

GetDbResponseTimeMs returns the DbResponseTimeMs field if non-nil, zero value otherwise.

### GetDbResponseTimeMsOk

`func (o *AutocompleteMeta) GetDbResponseTimeMsOk() (*int32, bool)`

GetDbResponseTimeMsOk returns a tuple with the DbResponseTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbResponseTimeMs

`func (o *AutocompleteMeta) SetDbResponseTimeMs(v int32)`

SetDbResponseTimeMs sets DbResponseTimeMs field to given value.

### HasDbResponseTimeMs

`func (o *AutocompleteMeta) HasDbResponseTimeMs() bool`

HasDbResponseTimeMs returns a boolean if a field has been set.

### GetPage

`func (o *AutocompleteMeta) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *AutocompleteMeta) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *AutocompleteMeta) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *AutocompleteMeta) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPerPage

`func (o *AutocompleteMeta) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *AutocompleteMeta) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *AutocompleteMeta) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *AutocompleteMeta) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


