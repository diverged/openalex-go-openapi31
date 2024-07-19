# ListWorks200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbstractInvertedIndex** | Pointer to **map[string]interface{}** | The abstract of the work, as an inverted index. This format allows for efficient searching and analysis of the abstract text. | [optional] 
**ApcList** | Pointer to [**ListWorks200ResponseResultsInnerApcList**](ListWorks200ResponseResultsInnerApcList.md) |  | [optional] 
**ApcPaid** | Pointer to [**ListWorks200ResponseResultsInnerApcPaid**](ListWorks200ResponseResultsInnerApcPaid.md) |  | [optional] 
**Authorships** | Pointer to [**[]ListWorks200ResponseResultsInnerAuthorshipsInner**](ListWorks200ResponseResultsInnerAuthorshipsInner.md) | List of authorship information for the work, including authors and their institutions. | [optional] 
**BestOaLocation** | Pointer to [**ListWorks200ResponseResultsInnerBestOaLocation**](ListWorks200ResponseResultsInnerBestOaLocation.md) |  | [optional] 
**Biblio** | Pointer to [**ListWorks200ResponseResultsInnerBiblio**](ListWorks200ResponseResultsInnerBiblio.md) |  | [optional] 
**CitedByApiUrl** | Pointer to **string** | A URL that displays a list of works that cite this work. | [optional] 
**CitedByCount** | Pointer to **int32** | The number of citations to this work. | [optional] 
**Concepts** | Pointer to **[]map[string]interface{}** | List of concepts (research topics) associated with the work. These are algorithmically inferred based on the work&#39;s content. | [optional] 
**CorrespondingAuthorIds** | Pointer to **[]string** | OpenAlex IDs of any authors for which authorships.is_corresponding is true. | [optional] 
**CorrespondingInstitutionIds** | Pointer to **[]string** | OpenAlex IDs of any institutions found within an authorship for which authorships.is_corresponding is true. | [optional] 
**CountriesDistinctCount** | Pointer to **int32** | Number of distinct country_codes among the authorships for this work. | [optional] 
**CountsByYear** | Pointer to [**[]ListWorks200ResponseResultsInnerCountsByYearInner**](ListWorks200ResponseResultsInnerCountsByYearInner.md) | Cited_by_count for each of the last ten years, binned by year. | [optional] 
**CreatedDate** | Pointer to **string** | The date this Work object was created in the OpenAlex dataset. | [optional] 
**DisplayName** | Pointer to **string** | The title of this work (same as title). | [optional] 
**Doi** | Pointer to **string** | The DOI for the work. | [optional] 
**FulltextOrigin** | Pointer to **string** | If a work&#39;s full text is searchable in OpenAlex, this tells you how we got the text. | [optional] 
**Grants** | Pointer to [**[]ListWorks200ResponseResultsInnerGrantsInner**](ListWorks200ResponseResultsInnerGrantsInner.md) | Information about grants that funded this work. | [optional] 
**HasFulltext** | Pointer to **bool** | Set to true if the work&#39;s full text is searchable in OpenAlex. | [optional] 
**Id** | Pointer to **string** | The OpenAlex ID for this work. | [optional] 
**Ids** | Pointer to [**ListWorks200ResponseResultsInnerIds**](ListWorks200ResponseResultsInnerIds.md) |  | [optional] 
**IndexedIn** | Pointer to **[]string** | The sources this work is indexed in. | [optional] 
**InstitutionsDistinctCount** | Pointer to **int32** | Number of distinct institutions among the authorships for this work. | [optional] 
**IsParatext** | Pointer to **bool** | True if we think this work is paratext. | [optional] 
**IsRetracted** | Pointer to **bool** | True if we know this work has been retracted. | [optional] 
**Keywords** | Pointer to [**[]ListKeywords200ResponseResultsInner**](ListKeywords200ResponseResultsInner.md) | A list of keywords associated with this work. | [optional] 
**Language** | Pointer to **string** | The language of the work in ISO 639-1 format. | [optional] 
**License** | Pointer to **string** | The license applied to this work at this host. Most toll-access works don&#39;t have an explicit license (they&#39;re under \&quot;all rights reserved\&quot; copyright), so this field generally has content only if is_oa is true. | [optional] 
**Locations** | Pointer to [**[]ListWorks200ResponseResultsInnerLocationsInner**](ListWorks200ResponseResultsInnerLocationsInner.md) | List of locations where the work can be found, including various versions and sources. | [optional] 
**LocationsCount** | Pointer to **int32** | Number of locations for this work. | [optional] 
**Mesh** | Pointer to [**[]ListWorks200ResponseResultsInnerMeshInner**](ListWorks200ResponseResultsInnerMeshInner.md) | List of Medical Subject Headings (MeSH) associated with the work, if applicable. MeSH terms are used to index and categorize biomedical literature. | [optional] 
**NgramsUrl** | Pointer to **string** | URL to retrieve n-grams for this work. | [optional] 
**OpenAccess** | Pointer to [**ListWorks200ResponseResultsInnerOpenAccess**](ListWorks200ResponseResultsInnerOpenAccess.md) |  | [optional] 
**PrimaryLocation** | Pointer to [**ListWorks200ResponseResultsInnerPrimaryLocation**](ListWorks200ResponseResultsInnerPrimaryLocation.md) |  | [optional] 
**PrimaryTopic** | Pointer to [**ListWorks200ResponseResultsInnerPrimaryTopic**](ListWorks200ResponseResultsInnerPrimaryTopic.md) |  | [optional] 
**PublicationDate** | Pointer to **string** | The day when this work was published, formatted as an ISO 8601 date. | [optional] 
**PublicationYear** | Pointer to **int32** | The year this work was published. | [optional] 
**ReferencedWorks** | Pointer to **[]string** | OpenAlex IDs for works that this work cites. | [optional] 
**RelatedWorks** | Pointer to **[]string** | OpenAlex IDs for works related to this work. These are computed algorithmically based on shared concepts. | [optional] 
**SustainableDevelopmentGoals** | Pointer to [**[]ListWorks200ResponseResultsInnerSustainableDevelopmentGoalsInner**](ListWorks200ResponseResultsInnerSustainableDevelopmentGoalsInner.md) | The UN Sustainable Development Goals relevant to this work. | [optional] 
**Title** | Pointer to **string** | The title of this work. | [optional] 
**Topics** | Pointer to [**[]ListWorks200ResponseResultsInnerTopicsInner**](ListWorks200ResponseResultsInnerTopicsInner.md) | The top ranked Topics for this work. | [optional] 
**Type** | Pointer to **string** | The type of the work. | [optional] 
**TypeCrossref** | Pointer to **string** | Legacy type information, using Crossref&#39;s \&quot;type\&quot; controlled vocabulary. | [optional] 
**UpdatedDate** | Pointer to **string** | The last time anything in this work object changed. Formatted as ISO 8601 extended format without time zone designator. | [optional] 

## Methods

### NewListWorks200ResponseResultsInner

`func NewListWorks200ResponseResultsInner() *ListWorks200ResponseResultsInner`

NewListWorks200ResponseResultsInner instantiates a new ListWorks200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorks200ResponseResultsInnerWithDefaults

`func NewListWorks200ResponseResultsInnerWithDefaults() *ListWorks200ResponseResultsInner`

NewListWorks200ResponseResultsInnerWithDefaults instantiates a new ListWorks200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbstractInvertedIndex

`func (o *ListWorks200ResponseResultsInner) GetAbstractInvertedIndex() map[string]interface{}`

GetAbstractInvertedIndex returns the AbstractInvertedIndex field if non-nil, zero value otherwise.

### GetAbstractInvertedIndexOk

`func (o *ListWorks200ResponseResultsInner) GetAbstractInvertedIndexOk() (*map[string]interface{}, bool)`

GetAbstractInvertedIndexOk returns a tuple with the AbstractInvertedIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstractInvertedIndex

`func (o *ListWorks200ResponseResultsInner) SetAbstractInvertedIndex(v map[string]interface{})`

SetAbstractInvertedIndex sets AbstractInvertedIndex field to given value.

### HasAbstractInvertedIndex

`func (o *ListWorks200ResponseResultsInner) HasAbstractInvertedIndex() bool`

HasAbstractInvertedIndex returns a boolean if a field has been set.

### GetApcList

`func (o *ListWorks200ResponseResultsInner) GetApcList() ListWorks200ResponseResultsInnerApcList`

GetApcList returns the ApcList field if non-nil, zero value otherwise.

### GetApcListOk

`func (o *ListWorks200ResponseResultsInner) GetApcListOk() (*ListWorks200ResponseResultsInnerApcList, bool)`

GetApcListOk returns a tuple with the ApcList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApcList

`func (o *ListWorks200ResponseResultsInner) SetApcList(v ListWorks200ResponseResultsInnerApcList)`

SetApcList sets ApcList field to given value.

### HasApcList

`func (o *ListWorks200ResponseResultsInner) HasApcList() bool`

HasApcList returns a boolean if a field has been set.

### GetApcPaid

`func (o *ListWorks200ResponseResultsInner) GetApcPaid() ListWorks200ResponseResultsInnerApcPaid`

GetApcPaid returns the ApcPaid field if non-nil, zero value otherwise.

### GetApcPaidOk

`func (o *ListWorks200ResponseResultsInner) GetApcPaidOk() (*ListWorks200ResponseResultsInnerApcPaid, bool)`

GetApcPaidOk returns a tuple with the ApcPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApcPaid

`func (o *ListWorks200ResponseResultsInner) SetApcPaid(v ListWorks200ResponseResultsInnerApcPaid)`

SetApcPaid sets ApcPaid field to given value.

### HasApcPaid

`func (o *ListWorks200ResponseResultsInner) HasApcPaid() bool`

HasApcPaid returns a boolean if a field has been set.

### GetAuthorships

`func (o *ListWorks200ResponseResultsInner) GetAuthorships() []ListWorks200ResponseResultsInnerAuthorshipsInner`

GetAuthorships returns the Authorships field if non-nil, zero value otherwise.

### GetAuthorshipsOk

`func (o *ListWorks200ResponseResultsInner) GetAuthorshipsOk() (*[]ListWorks200ResponseResultsInnerAuthorshipsInner, bool)`

GetAuthorshipsOk returns a tuple with the Authorships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorships

`func (o *ListWorks200ResponseResultsInner) SetAuthorships(v []ListWorks200ResponseResultsInnerAuthorshipsInner)`

SetAuthorships sets Authorships field to given value.

### HasAuthorships

`func (o *ListWorks200ResponseResultsInner) HasAuthorships() bool`

HasAuthorships returns a boolean if a field has been set.

### GetBestOaLocation

`func (o *ListWorks200ResponseResultsInner) GetBestOaLocation() ListWorks200ResponseResultsInnerBestOaLocation`

GetBestOaLocation returns the BestOaLocation field if non-nil, zero value otherwise.

### GetBestOaLocationOk

`func (o *ListWorks200ResponseResultsInner) GetBestOaLocationOk() (*ListWorks200ResponseResultsInnerBestOaLocation, bool)`

GetBestOaLocationOk returns a tuple with the BestOaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestOaLocation

`func (o *ListWorks200ResponseResultsInner) SetBestOaLocation(v ListWorks200ResponseResultsInnerBestOaLocation)`

SetBestOaLocation sets BestOaLocation field to given value.

### HasBestOaLocation

`func (o *ListWorks200ResponseResultsInner) HasBestOaLocation() bool`

HasBestOaLocation returns a boolean if a field has been set.

### GetBiblio

`func (o *ListWorks200ResponseResultsInner) GetBiblio() ListWorks200ResponseResultsInnerBiblio`

GetBiblio returns the Biblio field if non-nil, zero value otherwise.

### GetBiblioOk

`func (o *ListWorks200ResponseResultsInner) GetBiblioOk() (*ListWorks200ResponseResultsInnerBiblio, bool)`

GetBiblioOk returns a tuple with the Biblio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiblio

`func (o *ListWorks200ResponseResultsInner) SetBiblio(v ListWorks200ResponseResultsInnerBiblio)`

SetBiblio sets Biblio field to given value.

### HasBiblio

`func (o *ListWorks200ResponseResultsInner) HasBiblio() bool`

HasBiblio returns a boolean if a field has been set.

### GetCitedByApiUrl

`func (o *ListWorks200ResponseResultsInner) GetCitedByApiUrl() string`

GetCitedByApiUrl returns the CitedByApiUrl field if non-nil, zero value otherwise.

### GetCitedByApiUrlOk

`func (o *ListWorks200ResponseResultsInner) GetCitedByApiUrlOk() (*string, bool)`

GetCitedByApiUrlOk returns a tuple with the CitedByApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitedByApiUrl

`func (o *ListWorks200ResponseResultsInner) SetCitedByApiUrl(v string)`

SetCitedByApiUrl sets CitedByApiUrl field to given value.

### HasCitedByApiUrl

`func (o *ListWorks200ResponseResultsInner) HasCitedByApiUrl() bool`

HasCitedByApiUrl returns a boolean if a field has been set.

### GetCitedByCount

`func (o *ListWorks200ResponseResultsInner) GetCitedByCount() int32`

GetCitedByCount returns the CitedByCount field if non-nil, zero value otherwise.

### GetCitedByCountOk

`func (o *ListWorks200ResponseResultsInner) GetCitedByCountOk() (*int32, bool)`

GetCitedByCountOk returns a tuple with the CitedByCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitedByCount

`func (o *ListWorks200ResponseResultsInner) SetCitedByCount(v int32)`

SetCitedByCount sets CitedByCount field to given value.

### HasCitedByCount

`func (o *ListWorks200ResponseResultsInner) HasCitedByCount() bool`

HasCitedByCount returns a boolean if a field has been set.

### GetConcepts

`func (o *ListWorks200ResponseResultsInner) GetConcepts() []map[string]interface{}`

GetConcepts returns the Concepts field if non-nil, zero value otherwise.

### GetConceptsOk

`func (o *ListWorks200ResponseResultsInner) GetConceptsOk() (*[]map[string]interface{}, bool)`

GetConceptsOk returns a tuple with the Concepts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcepts

`func (o *ListWorks200ResponseResultsInner) SetConcepts(v []map[string]interface{})`

SetConcepts sets Concepts field to given value.

### HasConcepts

`func (o *ListWorks200ResponseResultsInner) HasConcepts() bool`

HasConcepts returns a boolean if a field has been set.

### GetCorrespondingAuthorIds

`func (o *ListWorks200ResponseResultsInner) GetCorrespondingAuthorIds() []string`

GetCorrespondingAuthorIds returns the CorrespondingAuthorIds field if non-nil, zero value otherwise.

### GetCorrespondingAuthorIdsOk

`func (o *ListWorks200ResponseResultsInner) GetCorrespondingAuthorIdsOk() (*[]string, bool)`

GetCorrespondingAuthorIdsOk returns a tuple with the CorrespondingAuthorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondingAuthorIds

`func (o *ListWorks200ResponseResultsInner) SetCorrespondingAuthorIds(v []string)`

SetCorrespondingAuthorIds sets CorrespondingAuthorIds field to given value.

### HasCorrespondingAuthorIds

`func (o *ListWorks200ResponseResultsInner) HasCorrespondingAuthorIds() bool`

HasCorrespondingAuthorIds returns a boolean if a field has been set.

### GetCorrespondingInstitutionIds

`func (o *ListWorks200ResponseResultsInner) GetCorrespondingInstitutionIds() []string`

GetCorrespondingInstitutionIds returns the CorrespondingInstitutionIds field if non-nil, zero value otherwise.

### GetCorrespondingInstitutionIdsOk

`func (o *ListWorks200ResponseResultsInner) GetCorrespondingInstitutionIdsOk() (*[]string, bool)`

GetCorrespondingInstitutionIdsOk returns a tuple with the CorrespondingInstitutionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondingInstitutionIds

`func (o *ListWorks200ResponseResultsInner) SetCorrespondingInstitutionIds(v []string)`

SetCorrespondingInstitutionIds sets CorrespondingInstitutionIds field to given value.

### HasCorrespondingInstitutionIds

`func (o *ListWorks200ResponseResultsInner) HasCorrespondingInstitutionIds() bool`

HasCorrespondingInstitutionIds returns a boolean if a field has been set.

### GetCountriesDistinctCount

`func (o *ListWorks200ResponseResultsInner) GetCountriesDistinctCount() int32`

GetCountriesDistinctCount returns the CountriesDistinctCount field if non-nil, zero value otherwise.

### GetCountriesDistinctCountOk

`func (o *ListWorks200ResponseResultsInner) GetCountriesDistinctCountOk() (*int32, bool)`

GetCountriesDistinctCountOk returns a tuple with the CountriesDistinctCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountriesDistinctCount

`func (o *ListWorks200ResponseResultsInner) SetCountriesDistinctCount(v int32)`

SetCountriesDistinctCount sets CountriesDistinctCount field to given value.

### HasCountriesDistinctCount

`func (o *ListWorks200ResponseResultsInner) HasCountriesDistinctCount() bool`

HasCountriesDistinctCount returns a boolean if a field has been set.

### GetCountsByYear

`func (o *ListWorks200ResponseResultsInner) GetCountsByYear() []ListWorks200ResponseResultsInnerCountsByYearInner`

GetCountsByYear returns the CountsByYear field if non-nil, zero value otherwise.

### GetCountsByYearOk

`func (o *ListWorks200ResponseResultsInner) GetCountsByYearOk() (*[]ListWorks200ResponseResultsInnerCountsByYearInner, bool)`

GetCountsByYearOk returns a tuple with the CountsByYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountsByYear

`func (o *ListWorks200ResponseResultsInner) SetCountsByYear(v []ListWorks200ResponseResultsInnerCountsByYearInner)`

SetCountsByYear sets CountsByYear field to given value.

### HasCountsByYear

`func (o *ListWorks200ResponseResultsInner) HasCountsByYear() bool`

HasCountsByYear returns a boolean if a field has been set.

### GetCreatedDate

`func (o *ListWorks200ResponseResultsInner) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ListWorks200ResponseResultsInner) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ListWorks200ResponseResultsInner) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ListWorks200ResponseResultsInner) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetDisplayName

`func (o *ListWorks200ResponseResultsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListWorks200ResponseResultsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListWorks200ResponseResultsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ListWorks200ResponseResultsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDoi

`func (o *ListWorks200ResponseResultsInner) GetDoi() string`

GetDoi returns the Doi field if non-nil, zero value otherwise.

### GetDoiOk

`func (o *ListWorks200ResponseResultsInner) GetDoiOk() (*string, bool)`

GetDoiOk returns a tuple with the Doi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoi

`func (o *ListWorks200ResponseResultsInner) SetDoi(v string)`

SetDoi sets Doi field to given value.

### HasDoi

`func (o *ListWorks200ResponseResultsInner) HasDoi() bool`

HasDoi returns a boolean if a field has been set.

### GetFulltextOrigin

`func (o *ListWorks200ResponseResultsInner) GetFulltextOrigin() string`

GetFulltextOrigin returns the FulltextOrigin field if non-nil, zero value otherwise.

### GetFulltextOriginOk

`func (o *ListWorks200ResponseResultsInner) GetFulltextOriginOk() (*string, bool)`

GetFulltextOriginOk returns a tuple with the FulltextOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulltextOrigin

`func (o *ListWorks200ResponseResultsInner) SetFulltextOrigin(v string)`

SetFulltextOrigin sets FulltextOrigin field to given value.

### HasFulltextOrigin

`func (o *ListWorks200ResponseResultsInner) HasFulltextOrigin() bool`

HasFulltextOrigin returns a boolean if a field has been set.

### GetGrants

`func (o *ListWorks200ResponseResultsInner) GetGrants() []ListWorks200ResponseResultsInnerGrantsInner`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *ListWorks200ResponseResultsInner) GetGrantsOk() (*[]ListWorks200ResponseResultsInnerGrantsInner, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *ListWorks200ResponseResultsInner) SetGrants(v []ListWorks200ResponseResultsInnerGrantsInner)`

SetGrants sets Grants field to given value.

### HasGrants

`func (o *ListWorks200ResponseResultsInner) HasGrants() bool`

HasGrants returns a boolean if a field has been set.

### GetHasFulltext

`func (o *ListWorks200ResponseResultsInner) GetHasFulltext() bool`

GetHasFulltext returns the HasFulltext field if non-nil, zero value otherwise.

### GetHasFulltextOk

`func (o *ListWorks200ResponseResultsInner) GetHasFulltextOk() (*bool, bool)`

GetHasFulltextOk returns a tuple with the HasFulltext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFulltext

`func (o *ListWorks200ResponseResultsInner) SetHasFulltext(v bool)`

SetHasFulltext sets HasFulltext field to given value.

### HasHasFulltext

`func (o *ListWorks200ResponseResultsInner) HasHasFulltext() bool`

HasHasFulltext returns a boolean if a field has been set.

### GetId

`func (o *ListWorks200ResponseResultsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListWorks200ResponseResultsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListWorks200ResponseResultsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListWorks200ResponseResultsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIds

`func (o *ListWorks200ResponseResultsInner) GetIds() ListWorks200ResponseResultsInnerIds`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListWorks200ResponseResultsInner) GetIdsOk() (*ListWorks200ResponseResultsInnerIds, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListWorks200ResponseResultsInner) SetIds(v ListWorks200ResponseResultsInnerIds)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ListWorks200ResponseResultsInner) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetIndexedIn

`func (o *ListWorks200ResponseResultsInner) GetIndexedIn() []string`

GetIndexedIn returns the IndexedIn field if non-nil, zero value otherwise.

### GetIndexedInOk

`func (o *ListWorks200ResponseResultsInner) GetIndexedInOk() (*[]string, bool)`

GetIndexedInOk returns a tuple with the IndexedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedIn

`func (o *ListWorks200ResponseResultsInner) SetIndexedIn(v []string)`

SetIndexedIn sets IndexedIn field to given value.

### HasIndexedIn

`func (o *ListWorks200ResponseResultsInner) HasIndexedIn() bool`

HasIndexedIn returns a boolean if a field has been set.

### GetInstitutionsDistinctCount

`func (o *ListWorks200ResponseResultsInner) GetInstitutionsDistinctCount() int32`

GetInstitutionsDistinctCount returns the InstitutionsDistinctCount field if non-nil, zero value otherwise.

### GetInstitutionsDistinctCountOk

`func (o *ListWorks200ResponseResultsInner) GetInstitutionsDistinctCountOk() (*int32, bool)`

GetInstitutionsDistinctCountOk returns a tuple with the InstitutionsDistinctCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionsDistinctCount

`func (o *ListWorks200ResponseResultsInner) SetInstitutionsDistinctCount(v int32)`

SetInstitutionsDistinctCount sets InstitutionsDistinctCount field to given value.

### HasInstitutionsDistinctCount

`func (o *ListWorks200ResponseResultsInner) HasInstitutionsDistinctCount() bool`

HasInstitutionsDistinctCount returns a boolean if a field has been set.

### GetIsParatext

`func (o *ListWorks200ResponseResultsInner) GetIsParatext() bool`

GetIsParatext returns the IsParatext field if non-nil, zero value otherwise.

### GetIsParatextOk

`func (o *ListWorks200ResponseResultsInner) GetIsParatextOk() (*bool, bool)`

GetIsParatextOk returns a tuple with the IsParatext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsParatext

`func (o *ListWorks200ResponseResultsInner) SetIsParatext(v bool)`

SetIsParatext sets IsParatext field to given value.

### HasIsParatext

`func (o *ListWorks200ResponseResultsInner) HasIsParatext() bool`

HasIsParatext returns a boolean if a field has been set.

### GetIsRetracted

`func (o *ListWorks200ResponseResultsInner) GetIsRetracted() bool`

GetIsRetracted returns the IsRetracted field if non-nil, zero value otherwise.

### GetIsRetractedOk

`func (o *ListWorks200ResponseResultsInner) GetIsRetractedOk() (*bool, bool)`

GetIsRetractedOk returns a tuple with the IsRetracted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRetracted

`func (o *ListWorks200ResponseResultsInner) SetIsRetracted(v bool)`

SetIsRetracted sets IsRetracted field to given value.

### HasIsRetracted

`func (o *ListWorks200ResponseResultsInner) HasIsRetracted() bool`

HasIsRetracted returns a boolean if a field has been set.

### GetKeywords

`func (o *ListWorks200ResponseResultsInner) GetKeywords() []ListKeywords200ResponseResultsInner`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *ListWorks200ResponseResultsInner) GetKeywordsOk() (*[]ListKeywords200ResponseResultsInner, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *ListWorks200ResponseResultsInner) SetKeywords(v []ListKeywords200ResponseResultsInner)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *ListWorks200ResponseResultsInner) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetLanguage

`func (o *ListWorks200ResponseResultsInner) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ListWorks200ResponseResultsInner) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ListWorks200ResponseResultsInner) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ListWorks200ResponseResultsInner) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLicense

`func (o *ListWorks200ResponseResultsInner) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *ListWorks200ResponseResultsInner) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *ListWorks200ResponseResultsInner) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *ListWorks200ResponseResultsInner) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetLocations

`func (o *ListWorks200ResponseResultsInner) GetLocations() []ListWorks200ResponseResultsInnerLocationsInner`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *ListWorks200ResponseResultsInner) GetLocationsOk() (*[]ListWorks200ResponseResultsInnerLocationsInner, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *ListWorks200ResponseResultsInner) SetLocations(v []ListWorks200ResponseResultsInnerLocationsInner)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *ListWorks200ResponseResultsInner) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetLocationsCount

`func (o *ListWorks200ResponseResultsInner) GetLocationsCount() int32`

GetLocationsCount returns the LocationsCount field if non-nil, zero value otherwise.

### GetLocationsCountOk

`func (o *ListWorks200ResponseResultsInner) GetLocationsCountOk() (*int32, bool)`

GetLocationsCountOk returns a tuple with the LocationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationsCount

`func (o *ListWorks200ResponseResultsInner) SetLocationsCount(v int32)`

SetLocationsCount sets LocationsCount field to given value.

### HasLocationsCount

`func (o *ListWorks200ResponseResultsInner) HasLocationsCount() bool`

HasLocationsCount returns a boolean if a field has been set.

### GetMesh

`func (o *ListWorks200ResponseResultsInner) GetMesh() []ListWorks200ResponseResultsInnerMeshInner`

GetMesh returns the Mesh field if non-nil, zero value otherwise.

### GetMeshOk

`func (o *ListWorks200ResponseResultsInner) GetMeshOk() (*[]ListWorks200ResponseResultsInnerMeshInner, bool)`

GetMeshOk returns a tuple with the Mesh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMesh

`func (o *ListWorks200ResponseResultsInner) SetMesh(v []ListWorks200ResponseResultsInnerMeshInner)`

SetMesh sets Mesh field to given value.

### HasMesh

`func (o *ListWorks200ResponseResultsInner) HasMesh() bool`

HasMesh returns a boolean if a field has been set.

### GetNgramsUrl

`func (o *ListWorks200ResponseResultsInner) GetNgramsUrl() string`

GetNgramsUrl returns the NgramsUrl field if non-nil, zero value otherwise.

### GetNgramsUrlOk

`func (o *ListWorks200ResponseResultsInner) GetNgramsUrlOk() (*string, bool)`

GetNgramsUrlOk returns a tuple with the NgramsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgramsUrl

`func (o *ListWorks200ResponseResultsInner) SetNgramsUrl(v string)`

SetNgramsUrl sets NgramsUrl field to given value.

### HasNgramsUrl

`func (o *ListWorks200ResponseResultsInner) HasNgramsUrl() bool`

HasNgramsUrl returns a boolean if a field has been set.

### GetOpenAccess

`func (o *ListWorks200ResponseResultsInner) GetOpenAccess() ListWorks200ResponseResultsInnerOpenAccess`

GetOpenAccess returns the OpenAccess field if non-nil, zero value otherwise.

### GetOpenAccessOk

`func (o *ListWorks200ResponseResultsInner) GetOpenAccessOk() (*ListWorks200ResponseResultsInnerOpenAccess, bool)`

GetOpenAccessOk returns a tuple with the OpenAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAccess

`func (o *ListWorks200ResponseResultsInner) SetOpenAccess(v ListWorks200ResponseResultsInnerOpenAccess)`

SetOpenAccess sets OpenAccess field to given value.

### HasOpenAccess

`func (o *ListWorks200ResponseResultsInner) HasOpenAccess() bool`

HasOpenAccess returns a boolean if a field has been set.

### GetPrimaryLocation

`func (o *ListWorks200ResponseResultsInner) GetPrimaryLocation() ListWorks200ResponseResultsInnerPrimaryLocation`

GetPrimaryLocation returns the PrimaryLocation field if non-nil, zero value otherwise.

### GetPrimaryLocationOk

`func (o *ListWorks200ResponseResultsInner) GetPrimaryLocationOk() (*ListWorks200ResponseResultsInnerPrimaryLocation, bool)`

GetPrimaryLocationOk returns a tuple with the PrimaryLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLocation

`func (o *ListWorks200ResponseResultsInner) SetPrimaryLocation(v ListWorks200ResponseResultsInnerPrimaryLocation)`

SetPrimaryLocation sets PrimaryLocation field to given value.

### HasPrimaryLocation

`func (o *ListWorks200ResponseResultsInner) HasPrimaryLocation() bool`

HasPrimaryLocation returns a boolean if a field has been set.

### GetPrimaryTopic

`func (o *ListWorks200ResponseResultsInner) GetPrimaryTopic() ListWorks200ResponseResultsInnerPrimaryTopic`

GetPrimaryTopic returns the PrimaryTopic field if non-nil, zero value otherwise.

### GetPrimaryTopicOk

`func (o *ListWorks200ResponseResultsInner) GetPrimaryTopicOk() (*ListWorks200ResponseResultsInnerPrimaryTopic, bool)`

GetPrimaryTopicOk returns a tuple with the PrimaryTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTopic

`func (o *ListWorks200ResponseResultsInner) SetPrimaryTopic(v ListWorks200ResponseResultsInnerPrimaryTopic)`

SetPrimaryTopic sets PrimaryTopic field to given value.

### HasPrimaryTopic

`func (o *ListWorks200ResponseResultsInner) HasPrimaryTopic() bool`

HasPrimaryTopic returns a boolean if a field has been set.

### GetPublicationDate

`func (o *ListWorks200ResponseResultsInner) GetPublicationDate() string`

GetPublicationDate returns the PublicationDate field if non-nil, zero value otherwise.

### GetPublicationDateOk

`func (o *ListWorks200ResponseResultsInner) GetPublicationDateOk() (*string, bool)`

GetPublicationDateOk returns a tuple with the PublicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationDate

`func (o *ListWorks200ResponseResultsInner) SetPublicationDate(v string)`

SetPublicationDate sets PublicationDate field to given value.

### HasPublicationDate

`func (o *ListWorks200ResponseResultsInner) HasPublicationDate() bool`

HasPublicationDate returns a boolean if a field has been set.

### GetPublicationYear

`func (o *ListWorks200ResponseResultsInner) GetPublicationYear() int32`

GetPublicationYear returns the PublicationYear field if non-nil, zero value otherwise.

### GetPublicationYearOk

`func (o *ListWorks200ResponseResultsInner) GetPublicationYearOk() (*int32, bool)`

GetPublicationYearOk returns a tuple with the PublicationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationYear

`func (o *ListWorks200ResponseResultsInner) SetPublicationYear(v int32)`

SetPublicationYear sets PublicationYear field to given value.

### HasPublicationYear

`func (o *ListWorks200ResponseResultsInner) HasPublicationYear() bool`

HasPublicationYear returns a boolean if a field has been set.

### GetReferencedWorks

`func (o *ListWorks200ResponseResultsInner) GetReferencedWorks() []string`

GetReferencedWorks returns the ReferencedWorks field if non-nil, zero value otherwise.

### GetReferencedWorksOk

`func (o *ListWorks200ResponseResultsInner) GetReferencedWorksOk() (*[]string, bool)`

GetReferencedWorksOk returns a tuple with the ReferencedWorks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedWorks

`func (o *ListWorks200ResponseResultsInner) SetReferencedWorks(v []string)`

SetReferencedWorks sets ReferencedWorks field to given value.

### HasReferencedWorks

`func (o *ListWorks200ResponseResultsInner) HasReferencedWorks() bool`

HasReferencedWorks returns a boolean if a field has been set.

### GetRelatedWorks

`func (o *ListWorks200ResponseResultsInner) GetRelatedWorks() []string`

GetRelatedWorks returns the RelatedWorks field if non-nil, zero value otherwise.

### GetRelatedWorksOk

`func (o *ListWorks200ResponseResultsInner) GetRelatedWorksOk() (*[]string, bool)`

GetRelatedWorksOk returns a tuple with the RelatedWorks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedWorks

`func (o *ListWorks200ResponseResultsInner) SetRelatedWorks(v []string)`

SetRelatedWorks sets RelatedWorks field to given value.

### HasRelatedWorks

`func (o *ListWorks200ResponseResultsInner) HasRelatedWorks() bool`

HasRelatedWorks returns a boolean if a field has been set.

### GetSustainableDevelopmentGoals

`func (o *ListWorks200ResponseResultsInner) GetSustainableDevelopmentGoals() []ListWorks200ResponseResultsInnerSustainableDevelopmentGoalsInner`

GetSustainableDevelopmentGoals returns the SustainableDevelopmentGoals field if non-nil, zero value otherwise.

### GetSustainableDevelopmentGoalsOk

`func (o *ListWorks200ResponseResultsInner) GetSustainableDevelopmentGoalsOk() (*[]ListWorks200ResponseResultsInnerSustainableDevelopmentGoalsInner, bool)`

GetSustainableDevelopmentGoalsOk returns a tuple with the SustainableDevelopmentGoals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSustainableDevelopmentGoals

`func (o *ListWorks200ResponseResultsInner) SetSustainableDevelopmentGoals(v []ListWorks200ResponseResultsInnerSustainableDevelopmentGoalsInner)`

SetSustainableDevelopmentGoals sets SustainableDevelopmentGoals field to given value.

### HasSustainableDevelopmentGoals

`func (o *ListWorks200ResponseResultsInner) HasSustainableDevelopmentGoals() bool`

HasSustainableDevelopmentGoals returns a boolean if a field has been set.

### GetTitle

`func (o *ListWorks200ResponseResultsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListWorks200ResponseResultsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListWorks200ResponseResultsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ListWorks200ResponseResultsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTopics

`func (o *ListWorks200ResponseResultsInner) GetTopics() []ListWorks200ResponseResultsInnerTopicsInner`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *ListWorks200ResponseResultsInner) GetTopicsOk() (*[]ListWorks200ResponseResultsInnerTopicsInner, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *ListWorks200ResponseResultsInner) SetTopics(v []ListWorks200ResponseResultsInnerTopicsInner)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *ListWorks200ResponseResultsInner) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### GetType

`func (o *ListWorks200ResponseResultsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListWorks200ResponseResultsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListWorks200ResponseResultsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListWorks200ResponseResultsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeCrossref

`func (o *ListWorks200ResponseResultsInner) GetTypeCrossref() string`

GetTypeCrossref returns the TypeCrossref field if non-nil, zero value otherwise.

### GetTypeCrossrefOk

`func (o *ListWorks200ResponseResultsInner) GetTypeCrossrefOk() (*string, bool)`

GetTypeCrossrefOk returns a tuple with the TypeCrossref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCrossref

`func (o *ListWorks200ResponseResultsInner) SetTypeCrossref(v string)`

SetTypeCrossref sets TypeCrossref field to given value.

### HasTypeCrossref

`func (o *ListWorks200ResponseResultsInner) HasTypeCrossref() bool`

HasTypeCrossref returns a boolean if a field has been set.

### GetUpdatedDate

`func (o *ListWorks200ResponseResultsInner) GetUpdatedDate() string`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *ListWorks200ResponseResultsInner) GetUpdatedDateOk() (*string, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *ListWorks200ResponseResultsInner) SetUpdatedDate(v string)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *ListWorks200ResponseResultsInner) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


