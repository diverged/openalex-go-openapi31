# ListConcepts200ResponseResultsInnerAncestorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The OpenAlex ID for this concept. | 
**Wikidata** | Pointer to **string** | The Wikidata ID for this concept. | [optional] 
**DisplayName** | **string** | The human-readable name of the concept. | 
**Level** | **int32** | The hierarchical level of the concept. Level 0 is the most general, and level 5 is the most specific. | 
**Score** | **float32** | The relevance score of this concept to the work. Higher scores indicate greater relevance. | 

## Methods

### NewListConcepts200ResponseResultsInnerAncestorsInner

`func NewListConcepts200ResponseResultsInnerAncestorsInner(id string, displayName string, level int32, score float32, ) *ListConcepts200ResponseResultsInnerAncestorsInner`

NewListConcepts200ResponseResultsInnerAncestorsInner instantiates a new ListConcepts200ResponseResultsInnerAncestorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConcepts200ResponseResultsInnerAncestorsInnerWithDefaults

`func NewListConcepts200ResponseResultsInnerAncestorsInnerWithDefaults() *ListConcepts200ResponseResultsInnerAncestorsInner`

NewListConcepts200ResponseResultsInnerAncestorsInnerWithDefaults instantiates a new ListConcepts200ResponseResultsInnerAncestorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListConcepts200ResponseResultsInnerAncestorsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListConcepts200ResponseResultsInnerAncestorsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListConcepts200ResponseResultsInnerAncestorsInner) SetId(v string)`

SetId sets Id field to given value.


### GetWikidata

`func (o *ListConcepts200ResponseResultsInnerAncestorsInner) GetWikidata() string`

GetWikidata returns the Wikidata field if non-nil, zero value otherwise.

### GetWikidataOk

`func (o *ListConcepts200ResponseResultsInnerAncestorsInner) GetWikidataOk() (*string, bool)`

GetWikidataOk returns a tuple with the Wikidata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikidata

`func (o *ListConcepts200ResponseResultsInnerAncestorsInner) SetWikidata(v string)`

SetWikidata sets Wikidata field to given value.

### HasWikidata

`func (o *ListConcepts200ResponseResultsInnerAncestorsInner) HasWikidata() bool`

HasWikidata returns a boolean if a field has been set.

### GetDisplayName

`func (o *ListConcepts200ResponseResultsInnerAncestorsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListConcepts200ResponseResultsInnerAncestorsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListConcepts200ResponseResultsInnerAncestorsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetLevel

`func (o *ListConcepts200ResponseResultsInnerAncestorsInner) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ListConcepts200ResponseResultsInnerAncestorsInner) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ListConcepts200ResponseResultsInnerAncestorsInner) SetLevel(v int32)`

SetLevel sets Level field to given value.


### GetScore

`func (o *ListConcepts200ResponseResultsInnerAncestorsInner) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ListConcepts200ResponseResultsInnerAncestorsInner) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ListConcepts200ResponseResultsInnerAncestorsInner) SetScore(v float32)`

SetScore sets Score field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


