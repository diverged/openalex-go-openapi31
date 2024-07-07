# AutocompleteResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The OpenAlex ID for this result entity. | [optional] 
**ExternalId** | Pointer to **string** | The Canonical External ID for this result entity. | [optional] 
**DisplayName** | Pointer to **string** | The entity&#39;s display_name property. | [optional] 
**EntityType** | Pointer to **string** | The entity&#39;s type (author, concept, institution, source, publisher, funder, or work). | [optional] 
**CitedByCount** | Pointer to **int32** | The entity&#39;s cited_by_count property. For works, this is the number of incoming citations. For other entities, it&#39;s the sum of incoming citations for all linked works. | [optional] 
**WorksCount** | Pointer to **int32** | The number of works associated with the entity. For entity type &#39;work&#39; it&#39;s always null. | [optional] 
**Hint** | Pointer to **string** | Extra information to help identify the right item. Content varies by entity type. | [optional] 

## Methods

### NewAutocompleteResultsInner

`func NewAutocompleteResultsInner() *AutocompleteResultsInner`

NewAutocompleteResultsInner instantiates a new AutocompleteResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutocompleteResultsInnerWithDefaults

`func NewAutocompleteResultsInnerWithDefaults() *AutocompleteResultsInner`

NewAutocompleteResultsInnerWithDefaults instantiates a new AutocompleteResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutocompleteResultsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutocompleteResultsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutocompleteResultsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutocompleteResultsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalId

`func (o *AutocompleteResultsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AutocompleteResultsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AutocompleteResultsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AutocompleteResultsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDisplayName

`func (o *AutocompleteResultsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AutocompleteResultsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AutocompleteResultsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AutocompleteResultsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEntityType

`func (o *AutocompleteResultsInner) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *AutocompleteResultsInner) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *AutocompleteResultsInner) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *AutocompleteResultsInner) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetCitedByCount

`func (o *AutocompleteResultsInner) GetCitedByCount() int32`

GetCitedByCount returns the CitedByCount field if non-nil, zero value otherwise.

### GetCitedByCountOk

`func (o *AutocompleteResultsInner) GetCitedByCountOk() (*int32, bool)`

GetCitedByCountOk returns a tuple with the CitedByCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitedByCount

`func (o *AutocompleteResultsInner) SetCitedByCount(v int32)`

SetCitedByCount sets CitedByCount field to given value.

### HasCitedByCount

`func (o *AutocompleteResultsInner) HasCitedByCount() bool`

HasCitedByCount returns a boolean if a field has been set.

### GetWorksCount

`func (o *AutocompleteResultsInner) GetWorksCount() int32`

GetWorksCount returns the WorksCount field if non-nil, zero value otherwise.

### GetWorksCountOk

`func (o *AutocompleteResultsInner) GetWorksCountOk() (*int32, bool)`

GetWorksCountOk returns a tuple with the WorksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksCount

`func (o *AutocompleteResultsInner) SetWorksCount(v int32)`

SetWorksCount sets WorksCount field to given value.

### HasWorksCount

`func (o *AutocompleteResultsInner) HasWorksCount() bool`

HasWorksCount returns a boolean if a field has been set.

### GetHint

`func (o *AutocompleteResultsInner) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *AutocompleteResultsInner) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *AutocompleteResultsInner) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *AutocompleteResultsInner) HasHint() bool`

HasHint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


