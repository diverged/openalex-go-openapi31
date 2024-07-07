# Autocomplete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**AutocompleteMeta**](AutocompleteMeta.md) |  | [optional] 
**Results** | Pointer to [**[]AutocompleteResultsInner**](AutocompleteResultsInner.md) | A list of up to ten autocomplete results for the query, sorted by citation count. | [optional] 

## Methods

### NewAutocomplete

`func NewAutocomplete() *Autocomplete`

NewAutocomplete instantiates a new Autocomplete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutocompleteWithDefaults

`func NewAutocompleteWithDefaults() *Autocomplete`

NewAutocompleteWithDefaults instantiates a new Autocomplete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *Autocomplete) GetMeta() AutocompleteMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Autocomplete) GetMetaOk() (*AutocompleteMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Autocomplete) SetMeta(v AutocompleteMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Autocomplete) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetResults

`func (o *Autocomplete) GetResults() []AutocompleteResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *Autocomplete) GetResultsOk() (*[]AutocompleteResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *Autocomplete) SetResults(v []AutocompleteResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *Autocomplete) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


