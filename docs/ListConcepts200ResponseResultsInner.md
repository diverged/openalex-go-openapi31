# ListConcepts200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The OpenAlex ID for this concept. | [optional] 
**Wikidata** | Pointer to **string** | The Wikidata ID for this concept. This is the Canonical External ID for concepts. | [optional] 
**DisplayName** | Pointer to **string** | The English-language label of the concept. | [optional] 
**Level** | Pointer to **int32** | The level in the concept tree where this concept lives. Lower-level concepts are more general, and higher-level concepts are more specific. | [optional] 
**Description** | Pointer to **string** | A brief description of this concept. | [optional] 
**WorksCount** | Pointer to **int32** | The number of works tagged with this concept. | [optional] 
**CitedByCount** | Pointer to **int32** | The number citations to works that have been tagged with this concept. | [optional] 
**Ids** | Pointer to [**ListConcepts200ResponseResultsInnerIds**](ListConcepts200ResponseResultsInnerIds.md) |  | [optional] 
**ImageUrl** | Pointer to **string** | URL where you can get an image representing this concept, where available. | [optional] 
**ImageThumbnailUrl** | Pointer to **string** | Same as image_url, but it&#39;s a smaller image. | [optional] 
**International** | Pointer to [**ListConcepts200ResponseResultsInnerInternational**](ListConcepts200ResponseResultsInnerInternational.md) |  | [optional] 
**Ancestors** | Pointer to [**[]ListConcepts200ResponseResultsInnerAncestorsInner**](ListConcepts200ResponseResultsInnerAncestorsInner.md) | List of concepts that this concept descends from, as dehydrated Concept objects. | [optional] 
**RelatedConcepts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CountsByYear** | Pointer to [**[]ListAuthors200ResponseResultsInnerCountsByYearInner**](ListAuthors200ResponseResultsInnerCountsByYearInner.md) | The values of works_count and cited_by_count for each of the last ten years, binned by year. | [optional] 
**WorksApiUrl** | Pointer to **string** | An URL that will get you a list of all the works tagged with this concept. | [optional] 
**UpdatedDate** | Pointer to **string** | The last time anything in this concept object changed. Formatted as ISO 8601 extended format without time zone designator. | [optional] 
**CreatedDate** | Pointer to **string** | The date this Concept object was created in the OpenAlex dataset, expressed as an ISO 8601 date string. | [optional] 
**SummaryStats** | Pointer to [**ListConcepts200ResponseResultsInnerSummaryStats**](ListConcepts200ResponseResultsInnerSummaryStats.md) |  | [optional] 

## Methods

### NewListConcepts200ResponseResultsInner

`func NewListConcepts200ResponseResultsInner() *ListConcepts200ResponseResultsInner`

NewListConcepts200ResponseResultsInner instantiates a new ListConcepts200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConcepts200ResponseResultsInnerWithDefaults

`func NewListConcepts200ResponseResultsInnerWithDefaults() *ListConcepts200ResponseResultsInner`

NewListConcepts200ResponseResultsInnerWithDefaults instantiates a new ListConcepts200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListConcepts200ResponseResultsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListConcepts200ResponseResultsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListConcepts200ResponseResultsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListConcepts200ResponseResultsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWikidata

`func (o *ListConcepts200ResponseResultsInner) GetWikidata() string`

GetWikidata returns the Wikidata field if non-nil, zero value otherwise.

### GetWikidataOk

`func (o *ListConcepts200ResponseResultsInner) GetWikidataOk() (*string, bool)`

GetWikidataOk returns a tuple with the Wikidata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikidata

`func (o *ListConcepts200ResponseResultsInner) SetWikidata(v string)`

SetWikidata sets Wikidata field to given value.

### HasWikidata

`func (o *ListConcepts200ResponseResultsInner) HasWikidata() bool`

HasWikidata returns a boolean if a field has been set.

### GetDisplayName

`func (o *ListConcepts200ResponseResultsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListConcepts200ResponseResultsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListConcepts200ResponseResultsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ListConcepts200ResponseResultsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLevel

`func (o *ListConcepts200ResponseResultsInner) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ListConcepts200ResponseResultsInner) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ListConcepts200ResponseResultsInner) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ListConcepts200ResponseResultsInner) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetDescription

`func (o *ListConcepts200ResponseResultsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListConcepts200ResponseResultsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListConcepts200ResponseResultsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListConcepts200ResponseResultsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWorksCount

`func (o *ListConcepts200ResponseResultsInner) GetWorksCount() int32`

GetWorksCount returns the WorksCount field if non-nil, zero value otherwise.

### GetWorksCountOk

`func (o *ListConcepts200ResponseResultsInner) GetWorksCountOk() (*int32, bool)`

GetWorksCountOk returns a tuple with the WorksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksCount

`func (o *ListConcepts200ResponseResultsInner) SetWorksCount(v int32)`

SetWorksCount sets WorksCount field to given value.

### HasWorksCount

`func (o *ListConcepts200ResponseResultsInner) HasWorksCount() bool`

HasWorksCount returns a boolean if a field has been set.

### GetCitedByCount

`func (o *ListConcepts200ResponseResultsInner) GetCitedByCount() int32`

GetCitedByCount returns the CitedByCount field if non-nil, zero value otherwise.

### GetCitedByCountOk

`func (o *ListConcepts200ResponseResultsInner) GetCitedByCountOk() (*int32, bool)`

GetCitedByCountOk returns a tuple with the CitedByCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitedByCount

`func (o *ListConcepts200ResponseResultsInner) SetCitedByCount(v int32)`

SetCitedByCount sets CitedByCount field to given value.

### HasCitedByCount

`func (o *ListConcepts200ResponseResultsInner) HasCitedByCount() bool`

HasCitedByCount returns a boolean if a field has been set.

### GetIds

`func (o *ListConcepts200ResponseResultsInner) GetIds() ListConcepts200ResponseResultsInnerIds`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListConcepts200ResponseResultsInner) GetIdsOk() (*ListConcepts200ResponseResultsInnerIds, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListConcepts200ResponseResultsInner) SetIds(v ListConcepts200ResponseResultsInnerIds)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ListConcepts200ResponseResultsInner) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetImageUrl

`func (o *ListConcepts200ResponseResultsInner) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *ListConcepts200ResponseResultsInner) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *ListConcepts200ResponseResultsInner) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *ListConcepts200ResponseResultsInner) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetImageThumbnailUrl

`func (o *ListConcepts200ResponseResultsInner) GetImageThumbnailUrl() string`

GetImageThumbnailUrl returns the ImageThumbnailUrl field if non-nil, zero value otherwise.

### GetImageThumbnailUrlOk

`func (o *ListConcepts200ResponseResultsInner) GetImageThumbnailUrlOk() (*string, bool)`

GetImageThumbnailUrlOk returns a tuple with the ImageThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageThumbnailUrl

`func (o *ListConcepts200ResponseResultsInner) SetImageThumbnailUrl(v string)`

SetImageThumbnailUrl sets ImageThumbnailUrl field to given value.

### HasImageThumbnailUrl

`func (o *ListConcepts200ResponseResultsInner) HasImageThumbnailUrl() bool`

HasImageThumbnailUrl returns a boolean if a field has been set.

### GetInternational

`func (o *ListConcepts200ResponseResultsInner) GetInternational() ListConcepts200ResponseResultsInnerInternational`

GetInternational returns the International field if non-nil, zero value otherwise.

### GetInternationalOk

`func (o *ListConcepts200ResponseResultsInner) GetInternationalOk() (*ListConcepts200ResponseResultsInnerInternational, bool)`

GetInternationalOk returns a tuple with the International field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternational

`func (o *ListConcepts200ResponseResultsInner) SetInternational(v ListConcepts200ResponseResultsInnerInternational)`

SetInternational sets International field to given value.

### HasInternational

`func (o *ListConcepts200ResponseResultsInner) HasInternational() bool`

HasInternational returns a boolean if a field has been set.

### GetAncestors

`func (o *ListConcepts200ResponseResultsInner) GetAncestors() []ListConcepts200ResponseResultsInnerAncestorsInner`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *ListConcepts200ResponseResultsInner) GetAncestorsOk() (*[]ListConcepts200ResponseResultsInnerAncestorsInner, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *ListConcepts200ResponseResultsInner) SetAncestors(v []ListConcepts200ResponseResultsInnerAncestorsInner)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *ListConcepts200ResponseResultsInner) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### GetRelatedConcepts

`func (o *ListConcepts200ResponseResultsInner) GetRelatedConcepts() []map[string]interface{}`

GetRelatedConcepts returns the RelatedConcepts field if non-nil, zero value otherwise.

### GetRelatedConceptsOk

`func (o *ListConcepts200ResponseResultsInner) GetRelatedConceptsOk() (*[]map[string]interface{}, bool)`

GetRelatedConceptsOk returns a tuple with the RelatedConcepts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedConcepts

`func (o *ListConcepts200ResponseResultsInner) SetRelatedConcepts(v []map[string]interface{})`

SetRelatedConcepts sets RelatedConcepts field to given value.

### HasRelatedConcepts

`func (o *ListConcepts200ResponseResultsInner) HasRelatedConcepts() bool`

HasRelatedConcepts returns a boolean if a field has been set.

### GetCountsByYear

`func (o *ListConcepts200ResponseResultsInner) GetCountsByYear() []ListAuthors200ResponseResultsInnerCountsByYearInner`

GetCountsByYear returns the CountsByYear field if non-nil, zero value otherwise.

### GetCountsByYearOk

`func (o *ListConcepts200ResponseResultsInner) GetCountsByYearOk() (*[]ListAuthors200ResponseResultsInnerCountsByYearInner, bool)`

GetCountsByYearOk returns a tuple with the CountsByYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountsByYear

`func (o *ListConcepts200ResponseResultsInner) SetCountsByYear(v []ListAuthors200ResponseResultsInnerCountsByYearInner)`

SetCountsByYear sets CountsByYear field to given value.

### HasCountsByYear

`func (o *ListConcepts200ResponseResultsInner) HasCountsByYear() bool`

HasCountsByYear returns a boolean if a field has been set.

### GetWorksApiUrl

`func (o *ListConcepts200ResponseResultsInner) GetWorksApiUrl() string`

GetWorksApiUrl returns the WorksApiUrl field if non-nil, zero value otherwise.

### GetWorksApiUrlOk

`func (o *ListConcepts200ResponseResultsInner) GetWorksApiUrlOk() (*string, bool)`

GetWorksApiUrlOk returns a tuple with the WorksApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksApiUrl

`func (o *ListConcepts200ResponseResultsInner) SetWorksApiUrl(v string)`

SetWorksApiUrl sets WorksApiUrl field to given value.

### HasWorksApiUrl

`func (o *ListConcepts200ResponseResultsInner) HasWorksApiUrl() bool`

HasWorksApiUrl returns a boolean if a field has been set.

### GetUpdatedDate

`func (o *ListConcepts200ResponseResultsInner) GetUpdatedDate() string`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *ListConcepts200ResponseResultsInner) GetUpdatedDateOk() (*string, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *ListConcepts200ResponseResultsInner) SetUpdatedDate(v string)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *ListConcepts200ResponseResultsInner) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.

### GetCreatedDate

`func (o *ListConcepts200ResponseResultsInner) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ListConcepts200ResponseResultsInner) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ListConcepts200ResponseResultsInner) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ListConcepts200ResponseResultsInner) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetSummaryStats

`func (o *ListConcepts200ResponseResultsInner) GetSummaryStats() ListConcepts200ResponseResultsInnerSummaryStats`

GetSummaryStats returns the SummaryStats field if non-nil, zero value otherwise.

### GetSummaryStatsOk

`func (o *ListConcepts200ResponseResultsInner) GetSummaryStatsOk() (*ListConcepts200ResponseResultsInnerSummaryStats, bool)`

GetSummaryStatsOk returns a tuple with the SummaryStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryStats

`func (o *ListConcepts200ResponseResultsInner) SetSummaryStats(v ListConcepts200ResponseResultsInnerSummaryStats)`

SetSummaryStats sets SummaryStats field to given value.

### HasSummaryStats

`func (o *ListConcepts200ResponseResultsInner) HasSummaryStats() bool`

HasSummaryStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


