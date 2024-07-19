# ListAuthors200ResponseResultsInnerXConceptsInnerInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The OpenAlex ID for this concept. | 
**Wikidata** | **string** | The Wikidata ID for this concept. | 
**DisplayName** | **string** | The human-readable name of the concept. | 
**Level** | **int32** | The hierarchical level of the concept. Level 0 is the most general, and level 5 is the most specific. | 
**Score** | Pointer to **NullableFloat32** | The strength of association between this source and the listed concept, from 0-100. | [optional] 

## Methods

### NewListAuthors200ResponseResultsInnerXConceptsInnerInner

`func NewListAuthors200ResponseResultsInnerXConceptsInnerInner(id string, wikidata string, displayName string, level int32, ) *ListAuthors200ResponseResultsInnerXConceptsInnerInner`

NewListAuthors200ResponseResultsInnerXConceptsInnerInner instantiates a new ListAuthors200ResponseResultsInnerXConceptsInnerInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAuthors200ResponseResultsInnerXConceptsInnerInnerWithDefaults

`func NewListAuthors200ResponseResultsInnerXConceptsInnerInnerWithDefaults() *ListAuthors200ResponseResultsInnerXConceptsInnerInner`

NewListAuthors200ResponseResultsInnerXConceptsInnerInnerWithDefaults instantiates a new ListAuthors200ResponseResultsInnerXConceptsInnerInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListAuthors200ResponseResultsInnerXConceptsInnerInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListAuthors200ResponseResultsInnerXConceptsInnerInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListAuthors200ResponseResultsInnerXConceptsInnerInner) SetId(v string)`

SetId sets Id field to given value.


### GetWikidata

`func (o *ListAuthors200ResponseResultsInnerXConceptsInnerInner) GetWikidata() string`

GetWikidata returns the Wikidata field if non-nil, zero value otherwise.

### GetWikidataOk

`func (o *ListAuthors200ResponseResultsInnerXConceptsInnerInner) GetWikidataOk() (*string, bool)`

GetWikidataOk returns a tuple with the Wikidata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikidata

`func (o *ListAuthors200ResponseResultsInnerXConceptsInnerInner) SetWikidata(v string)`

SetWikidata sets Wikidata field to given value.


### GetDisplayName

`func (o *ListAuthors200ResponseResultsInnerXConceptsInnerInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListAuthors200ResponseResultsInnerXConceptsInnerInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListAuthors200ResponseResultsInnerXConceptsInnerInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetLevel

`func (o *ListAuthors200ResponseResultsInnerXConceptsInnerInner) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ListAuthors200ResponseResultsInnerXConceptsInnerInner) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ListAuthors200ResponseResultsInnerXConceptsInnerInner) SetLevel(v int32)`

SetLevel sets Level field to given value.


### GetScore

`func (o *ListAuthors200ResponseResultsInnerXConceptsInnerInner) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ListAuthors200ResponseResultsInnerXConceptsInnerInner) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ListAuthors200ResponseResultsInnerXConceptsInnerInner) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ListAuthors200ResponseResultsInnerXConceptsInnerInner) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *ListAuthors200ResponseResultsInnerXConceptsInnerInner) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *ListAuthors200ResponseResultsInnerXConceptsInnerInner) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


