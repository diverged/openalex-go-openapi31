# ListAuthors200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affiliations** | Pointer to [**[]ListAuthors200ResponseResultsInnerAffiliationsInner**](ListAuthors200ResponseResultsInnerAffiliationsInner.md) | List of objects representing the affiliations this author has claimed in their publications. | [optional] 
**CitedByCount** | Pointer to **int32** | The total number of Works that cite a work this author has created. | [optional] 
**CountsByYear** | Pointer to [**[]ListAuthors200ResponseResultsInnerCountsByYearInner**](ListAuthors200ResponseResultsInnerCountsByYearInner.md) | Works_count and cited_by_count for each of the last ten years, binned by year. | [optional] 
**CreatedDate** | Pointer to **string** | The date this Author object was created in the OpenAlex dataset. | [optional] 
**DisplayName** | Pointer to **string** | The name of the author as a single string. | [optional] 
**DisplayNameAlternatives** | Pointer to **[]string** | Other ways that we&#39;ve found this author&#39;s name displayed. | [optional] 
**Id** | Pointer to **string** | The OpenAlex ID for this author. | [optional] 
**Ids** | Pointer to [**ListAuthors200ResponseResultsInnerIds**](ListAuthors200ResponseResultsInnerIds.md) |  | [optional] 
**LastKnownInstitution** | Pointer to [**ListAuthors200ResponseResultsInnerLastKnownInstitution**](ListAuthors200ResponseResultsInnerLastKnownInstitution.md) |  | [optional] 
**LastKnownInstitutions** | Pointer to [**[]ListAuthors200ResponseResultsInnerAffiliationsInnerInstitution**](ListAuthors200ResponseResultsInnerAffiliationsInnerInstitution.md) | This author&#39;s last known institutional affiliations. | [optional] 
**Orcid** | Pointer to **string** | The ORCID ID for this author. | [optional] 
**SummaryStats** | Pointer to [**ListAuthors200ResponseResultsInnerSummaryStats**](ListAuthors200ResponseResultsInnerSummaryStats.md) |  | [optional] 
**Topics** | Pointer to [**[]ListAuthors200ResponseResultsInnerTopicsInner**](ListAuthors200ResponseResultsInnerTopicsInner.md) | Topics that are frequently associated with works affiliated with this source, in descending order of count. | [optional] 
**TopicShare** | Pointer to [**[]ListAuthors200ResponseResultsInnerTopicsInner**](ListAuthors200ResponseResultsInnerTopicsInner.md) | Topics that are frequently associated with works affiliated with this source, in descending order of value. | [optional] 
**UpdatedDate** | Pointer to **string** | The last time anything in this author object changed. Formatted as ISO 8601 extended format without time zone designator. | [optional] 
**WorksApiUrl** | Pointer to **string** | A URL that will get you a list of all this author&#39;s works. | [optional] 
**WorksCount** | Pointer to **int32** | The number of Works this author has created. | [optional] 
**XConcepts** | Pointer to **[]map[string]interface{}** | The concepts most frequently applied to works created by this author (deprecated). | [optional] 

## Methods

### NewListAuthors200ResponseResultsInner

`func NewListAuthors200ResponseResultsInner() *ListAuthors200ResponseResultsInner`

NewListAuthors200ResponseResultsInner instantiates a new ListAuthors200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAuthors200ResponseResultsInnerWithDefaults

`func NewListAuthors200ResponseResultsInnerWithDefaults() *ListAuthors200ResponseResultsInner`

NewListAuthors200ResponseResultsInnerWithDefaults instantiates a new ListAuthors200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffiliations

`func (o *ListAuthors200ResponseResultsInner) GetAffiliations() []ListAuthors200ResponseResultsInnerAffiliationsInner`

GetAffiliations returns the Affiliations field if non-nil, zero value otherwise.

### GetAffiliationsOk

`func (o *ListAuthors200ResponseResultsInner) GetAffiliationsOk() (*[]ListAuthors200ResponseResultsInnerAffiliationsInner, bool)`

GetAffiliationsOk returns a tuple with the Affiliations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliations

`func (o *ListAuthors200ResponseResultsInner) SetAffiliations(v []ListAuthors200ResponseResultsInnerAffiliationsInner)`

SetAffiliations sets Affiliations field to given value.

### HasAffiliations

`func (o *ListAuthors200ResponseResultsInner) HasAffiliations() bool`

HasAffiliations returns a boolean if a field has been set.

### GetCitedByCount

`func (o *ListAuthors200ResponseResultsInner) GetCitedByCount() int32`

GetCitedByCount returns the CitedByCount field if non-nil, zero value otherwise.

### GetCitedByCountOk

`func (o *ListAuthors200ResponseResultsInner) GetCitedByCountOk() (*int32, bool)`

GetCitedByCountOk returns a tuple with the CitedByCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitedByCount

`func (o *ListAuthors200ResponseResultsInner) SetCitedByCount(v int32)`

SetCitedByCount sets CitedByCount field to given value.

### HasCitedByCount

`func (o *ListAuthors200ResponseResultsInner) HasCitedByCount() bool`

HasCitedByCount returns a boolean if a field has been set.

### GetCountsByYear

`func (o *ListAuthors200ResponseResultsInner) GetCountsByYear() []ListAuthors200ResponseResultsInnerCountsByYearInner`

GetCountsByYear returns the CountsByYear field if non-nil, zero value otherwise.

### GetCountsByYearOk

`func (o *ListAuthors200ResponseResultsInner) GetCountsByYearOk() (*[]ListAuthors200ResponseResultsInnerCountsByYearInner, bool)`

GetCountsByYearOk returns a tuple with the CountsByYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountsByYear

`func (o *ListAuthors200ResponseResultsInner) SetCountsByYear(v []ListAuthors200ResponseResultsInnerCountsByYearInner)`

SetCountsByYear sets CountsByYear field to given value.

### HasCountsByYear

`func (o *ListAuthors200ResponseResultsInner) HasCountsByYear() bool`

HasCountsByYear returns a boolean if a field has been set.

### GetCreatedDate

`func (o *ListAuthors200ResponseResultsInner) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ListAuthors200ResponseResultsInner) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ListAuthors200ResponseResultsInner) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ListAuthors200ResponseResultsInner) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetDisplayName

`func (o *ListAuthors200ResponseResultsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListAuthors200ResponseResultsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListAuthors200ResponseResultsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ListAuthors200ResponseResultsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDisplayNameAlternatives

`func (o *ListAuthors200ResponseResultsInner) GetDisplayNameAlternatives() []string`

GetDisplayNameAlternatives returns the DisplayNameAlternatives field if non-nil, zero value otherwise.

### GetDisplayNameAlternativesOk

`func (o *ListAuthors200ResponseResultsInner) GetDisplayNameAlternativesOk() (*[]string, bool)`

GetDisplayNameAlternativesOk returns a tuple with the DisplayNameAlternatives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNameAlternatives

`func (o *ListAuthors200ResponseResultsInner) SetDisplayNameAlternatives(v []string)`

SetDisplayNameAlternatives sets DisplayNameAlternatives field to given value.

### HasDisplayNameAlternatives

`func (o *ListAuthors200ResponseResultsInner) HasDisplayNameAlternatives() bool`

HasDisplayNameAlternatives returns a boolean if a field has been set.

### GetId

`func (o *ListAuthors200ResponseResultsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListAuthors200ResponseResultsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListAuthors200ResponseResultsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListAuthors200ResponseResultsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIds

`func (o *ListAuthors200ResponseResultsInner) GetIds() ListAuthors200ResponseResultsInnerIds`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListAuthors200ResponseResultsInner) GetIdsOk() (*ListAuthors200ResponseResultsInnerIds, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListAuthors200ResponseResultsInner) SetIds(v ListAuthors200ResponseResultsInnerIds)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ListAuthors200ResponseResultsInner) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetLastKnownInstitution

`func (o *ListAuthors200ResponseResultsInner) GetLastKnownInstitution() ListAuthors200ResponseResultsInnerLastKnownInstitution`

GetLastKnownInstitution returns the LastKnownInstitution field if non-nil, zero value otherwise.

### GetLastKnownInstitutionOk

`func (o *ListAuthors200ResponseResultsInner) GetLastKnownInstitutionOk() (*ListAuthors200ResponseResultsInnerLastKnownInstitution, bool)`

GetLastKnownInstitutionOk returns a tuple with the LastKnownInstitution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastKnownInstitution

`func (o *ListAuthors200ResponseResultsInner) SetLastKnownInstitution(v ListAuthors200ResponseResultsInnerLastKnownInstitution)`

SetLastKnownInstitution sets LastKnownInstitution field to given value.

### HasLastKnownInstitution

`func (o *ListAuthors200ResponseResultsInner) HasLastKnownInstitution() bool`

HasLastKnownInstitution returns a boolean if a field has been set.

### GetLastKnownInstitutions

`func (o *ListAuthors200ResponseResultsInner) GetLastKnownInstitutions() []ListAuthors200ResponseResultsInnerAffiliationsInnerInstitution`

GetLastKnownInstitutions returns the LastKnownInstitutions field if non-nil, zero value otherwise.

### GetLastKnownInstitutionsOk

`func (o *ListAuthors200ResponseResultsInner) GetLastKnownInstitutionsOk() (*[]ListAuthors200ResponseResultsInnerAffiliationsInnerInstitution, bool)`

GetLastKnownInstitutionsOk returns a tuple with the LastKnownInstitutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastKnownInstitutions

`func (o *ListAuthors200ResponseResultsInner) SetLastKnownInstitutions(v []ListAuthors200ResponseResultsInnerAffiliationsInnerInstitution)`

SetLastKnownInstitutions sets LastKnownInstitutions field to given value.

### HasLastKnownInstitutions

`func (o *ListAuthors200ResponseResultsInner) HasLastKnownInstitutions() bool`

HasLastKnownInstitutions returns a boolean if a field has been set.

### GetOrcid

`func (o *ListAuthors200ResponseResultsInner) GetOrcid() string`

GetOrcid returns the Orcid field if non-nil, zero value otherwise.

### GetOrcidOk

`func (o *ListAuthors200ResponseResultsInner) GetOrcidOk() (*string, bool)`

GetOrcidOk returns a tuple with the Orcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrcid

`func (o *ListAuthors200ResponseResultsInner) SetOrcid(v string)`

SetOrcid sets Orcid field to given value.

### HasOrcid

`func (o *ListAuthors200ResponseResultsInner) HasOrcid() bool`

HasOrcid returns a boolean if a field has been set.

### GetSummaryStats

`func (o *ListAuthors200ResponseResultsInner) GetSummaryStats() ListAuthors200ResponseResultsInnerSummaryStats`

GetSummaryStats returns the SummaryStats field if non-nil, zero value otherwise.

### GetSummaryStatsOk

`func (o *ListAuthors200ResponseResultsInner) GetSummaryStatsOk() (*ListAuthors200ResponseResultsInnerSummaryStats, bool)`

GetSummaryStatsOk returns a tuple with the SummaryStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryStats

`func (o *ListAuthors200ResponseResultsInner) SetSummaryStats(v ListAuthors200ResponseResultsInnerSummaryStats)`

SetSummaryStats sets SummaryStats field to given value.

### HasSummaryStats

`func (o *ListAuthors200ResponseResultsInner) HasSummaryStats() bool`

HasSummaryStats returns a boolean if a field has been set.

### GetTopics

`func (o *ListAuthors200ResponseResultsInner) GetTopics() []ListAuthors200ResponseResultsInnerTopicsInner`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *ListAuthors200ResponseResultsInner) GetTopicsOk() (*[]ListAuthors200ResponseResultsInnerTopicsInner, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *ListAuthors200ResponseResultsInner) SetTopics(v []ListAuthors200ResponseResultsInnerTopicsInner)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *ListAuthors200ResponseResultsInner) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### GetTopicShare

`func (o *ListAuthors200ResponseResultsInner) GetTopicShare() []ListAuthors200ResponseResultsInnerTopicsInner`

GetTopicShare returns the TopicShare field if non-nil, zero value otherwise.

### GetTopicShareOk

`func (o *ListAuthors200ResponseResultsInner) GetTopicShareOk() (*[]ListAuthors200ResponseResultsInnerTopicsInner, bool)`

GetTopicShareOk returns a tuple with the TopicShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicShare

`func (o *ListAuthors200ResponseResultsInner) SetTopicShare(v []ListAuthors200ResponseResultsInnerTopicsInner)`

SetTopicShare sets TopicShare field to given value.

### HasTopicShare

`func (o *ListAuthors200ResponseResultsInner) HasTopicShare() bool`

HasTopicShare returns a boolean if a field has been set.

### GetUpdatedDate

`func (o *ListAuthors200ResponseResultsInner) GetUpdatedDate() string`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *ListAuthors200ResponseResultsInner) GetUpdatedDateOk() (*string, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *ListAuthors200ResponseResultsInner) SetUpdatedDate(v string)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *ListAuthors200ResponseResultsInner) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.

### GetWorksApiUrl

`func (o *ListAuthors200ResponseResultsInner) GetWorksApiUrl() string`

GetWorksApiUrl returns the WorksApiUrl field if non-nil, zero value otherwise.

### GetWorksApiUrlOk

`func (o *ListAuthors200ResponseResultsInner) GetWorksApiUrlOk() (*string, bool)`

GetWorksApiUrlOk returns a tuple with the WorksApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksApiUrl

`func (o *ListAuthors200ResponseResultsInner) SetWorksApiUrl(v string)`

SetWorksApiUrl sets WorksApiUrl field to given value.

### HasWorksApiUrl

`func (o *ListAuthors200ResponseResultsInner) HasWorksApiUrl() bool`

HasWorksApiUrl returns a boolean if a field has been set.

### GetWorksCount

`func (o *ListAuthors200ResponseResultsInner) GetWorksCount() int32`

GetWorksCount returns the WorksCount field if non-nil, zero value otherwise.

### GetWorksCountOk

`func (o *ListAuthors200ResponseResultsInner) GetWorksCountOk() (*int32, bool)`

GetWorksCountOk returns a tuple with the WorksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksCount

`func (o *ListAuthors200ResponseResultsInner) SetWorksCount(v int32)`

SetWorksCount sets WorksCount field to given value.

### HasWorksCount

`func (o *ListAuthors200ResponseResultsInner) HasWorksCount() bool`

HasWorksCount returns a boolean if a field has been set.

### GetXConcepts

`func (o *ListAuthors200ResponseResultsInner) GetXConcepts() []map[string]interface{}`

GetXConcepts returns the XConcepts field if non-nil, zero value otherwise.

### GetXConceptsOk

`func (o *ListAuthors200ResponseResultsInner) GetXConceptsOk() (*[]map[string]interface{}, bool)`

GetXConceptsOk returns a tuple with the XConcepts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXConcepts

`func (o *ListAuthors200ResponseResultsInner) SetXConcepts(v []map[string]interface{})`

SetXConcepts sets XConcepts field to given value.

### HasXConcepts

`func (o *ListAuthors200ResponseResultsInner) HasXConcepts() bool`

HasXConcepts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


