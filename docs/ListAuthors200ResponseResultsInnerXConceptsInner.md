# ListAuthors200ResponseResultsInnerXConceptsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The OpenAlex ID for this concept. | 
**Wikidata** | **string** | The Wikidata ID for this concept. | 
**DisplayName** | **string** | The human-readable name of the concept. | 
**Level** | **int32** | The hierarchical level of the concept. Level 0 is the most general, and level 5 is the most specific. | 
**Score** | Pointer to **float32** |  | [optional] 

## Methods

### NewListAuthors200ResponseResultsInnerXConceptsInner

`func NewListAuthors200ResponseResultsInnerXConceptsInner(id string, wikidata string, displayName string, level int32, ) *ListAuthors200ResponseResultsInnerXConceptsInner`

NewListAuthors200ResponseResultsInnerXConceptsInner instantiates a new ListAuthors200ResponseResultsInnerXConceptsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAuthors200ResponseResultsInnerXConceptsInnerWithDefaults

`func NewListAuthors200ResponseResultsInnerXConceptsInnerWithDefaults() *ListAuthors200ResponseResultsInnerXConceptsInner`

NewListAuthors200ResponseResultsInnerXConceptsInnerWithDefaults instantiates a new ListAuthors200ResponseResultsInnerXConceptsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListAuthors200ResponseResultsInnerXConceptsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListAuthors200ResponseResultsInnerXConceptsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListAuthors200ResponseResultsInnerXConceptsInner) SetId(v string)`

SetId sets Id field to given value.


### GetWikidata

`func (o *ListAuthors200ResponseResultsInnerXConceptsInner) GetWikidata() string`

GetWikidata returns the Wikidata field if non-nil, zero value otherwise.

### GetWikidataOk

`func (o *ListAuthors200ResponseResultsInnerXConceptsInner) GetWikidataOk() (*string, bool)`

GetWikidataOk returns a tuple with the Wikidata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikidata

`func (o *ListAuthors200ResponseResultsInnerXConceptsInner) SetWikidata(v string)`

SetWikidata sets Wikidata field to given value.


### GetDisplayName

`func (o *ListAuthors200ResponseResultsInnerXConceptsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListAuthors200ResponseResultsInnerXConceptsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListAuthors200ResponseResultsInnerXConceptsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetLevel

`func (o *ListAuthors200ResponseResultsInnerXConceptsInner) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ListAuthors200ResponseResultsInnerXConceptsInner) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ListAuthors200ResponseResultsInnerXConceptsInner) SetLevel(v int32)`

SetLevel sets Level field to given value.


### GetScore

`func (o *ListAuthors200ResponseResultsInnerXConceptsInner) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ListAuthors200ResponseResultsInnerXConceptsInner) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ListAuthors200ResponseResultsInnerXConceptsInner) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ListAuthors200ResponseResultsInnerXConceptsInner) HasScore() bool`

HasScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


