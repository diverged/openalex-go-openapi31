# ListConcepts200ResponseResultsInnerRelatedConceptsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The OpenAlex ID for this concept. | 
**Wikidata** | Pointer to **string** | The Wikidata ID for this concept. | [optional] 
**DisplayName** | **string** | The human-readable name of the concept. | 
**Level** | **int32** | The hierarchical level of the concept. Level 0 is the most general, and level 5 is the most specific. | 
**Score** | **float32** | The strength of association between this concept and the listed concept, on a scale of 0-100. | 

## Methods

### NewListConcepts200ResponseResultsInnerRelatedConceptsInner

`func NewListConcepts200ResponseResultsInnerRelatedConceptsInner(id string, displayName string, level int32, score float32, ) *ListConcepts200ResponseResultsInnerRelatedConceptsInner`

NewListConcepts200ResponseResultsInnerRelatedConceptsInner instantiates a new ListConcepts200ResponseResultsInnerRelatedConceptsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConcepts200ResponseResultsInnerRelatedConceptsInnerWithDefaults

`func NewListConcepts200ResponseResultsInnerRelatedConceptsInnerWithDefaults() *ListConcepts200ResponseResultsInnerRelatedConceptsInner`

NewListConcepts200ResponseResultsInnerRelatedConceptsInnerWithDefaults instantiates a new ListConcepts200ResponseResultsInnerRelatedConceptsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListConcepts200ResponseResultsInnerRelatedConceptsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListConcepts200ResponseResultsInnerRelatedConceptsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListConcepts200ResponseResultsInnerRelatedConceptsInner) SetId(v string)`

SetId sets Id field to given value.


### GetWikidata

`func (o *ListConcepts200ResponseResultsInnerRelatedConceptsInner) GetWikidata() string`

GetWikidata returns the Wikidata field if non-nil, zero value otherwise.

### GetWikidataOk

`func (o *ListConcepts200ResponseResultsInnerRelatedConceptsInner) GetWikidataOk() (*string, bool)`

GetWikidataOk returns a tuple with the Wikidata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikidata

`func (o *ListConcepts200ResponseResultsInnerRelatedConceptsInner) SetWikidata(v string)`

SetWikidata sets Wikidata field to given value.

### HasWikidata

`func (o *ListConcepts200ResponseResultsInnerRelatedConceptsInner) HasWikidata() bool`

HasWikidata returns a boolean if a field has been set.

### GetDisplayName

`func (o *ListConcepts200ResponseResultsInnerRelatedConceptsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListConcepts200ResponseResultsInnerRelatedConceptsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListConcepts200ResponseResultsInnerRelatedConceptsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetLevel

`func (o *ListConcepts200ResponseResultsInnerRelatedConceptsInner) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ListConcepts200ResponseResultsInnerRelatedConceptsInner) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ListConcepts200ResponseResultsInnerRelatedConceptsInner) SetLevel(v int32)`

SetLevel sets Level field to given value.


### GetScore

`func (o *ListConcepts200ResponseResultsInnerRelatedConceptsInner) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ListConcepts200ResponseResultsInnerRelatedConceptsInner) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ListConcepts200ResponseResultsInnerRelatedConceptsInner) SetScore(v float32)`

SetScore sets Score field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


