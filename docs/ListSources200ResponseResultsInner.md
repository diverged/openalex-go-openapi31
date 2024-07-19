# ListSources200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbbreviatedTitle** | Pointer to **string** | An abbreviated title obtained from the ISSN Centre. | [optional] 
**AlternateTitles** | Pointer to **[]string** | Alternate titles for this source, as obtained from the ISSN Centre and individual work records, like Crossref DOIs, that carry the source name as a string. These are commonly abbreviations or translations of the source&#39;s canonical name. | [optional] 
**ApcPrices** | Pointer to [**[]ListSources200ResponseResultsInnerApcPricesInner**](ListSources200ResponseResultsInnerApcPricesInner.md) | Article processing charge information, taken directly from DOAJ. | [optional] 
**ApcUsd** | Pointer to **int32** | The source&#39;s article processing charge in US Dollars, if available from DOAJ. The apc_usd value is calculated by taking the APC price with a currency of USD if it is available. If it&#39;s not available, we convert the first available value from apc_prices into USD, using recent exchange rates. | [optional] 
**CitedByCount** | Pointer to **int32** | The total number of Works that cite a Work hosted in this source. | [optional] 
**CountryCode** | Pointer to **string** | The country that this source is associated with, represented as an ISO two-letter country code. | [optional] 
**CountsByYear** | Pointer to [**[]ListAuthors200ResponseResultsInnerCountsByYearInner**](ListAuthors200ResponseResultsInnerCountsByYearInner.md) | works_count and cited_by_count for each of the last ten years, binned by year. If the source was founded less than ten years ago, there will naturally be fewer than ten years in this list. Years with zero citations and zero works have been removed. | [optional] 
**CreatedDate** | Pointer to **string** | The date this Source object was created in the OpenAlex dataset, expressed as an ISO 8601 date string. | [optional] 
**DisplayName** | Pointer to **string** | The name of the source. | [optional] 
**HomepageUrl** | Pointer to **string** | The starting page for navigating the contents of this source; the homepage for this source&#39;s website. | [optional] 
**HostOrganization** | Pointer to **string** | The host organization for this source as an OpenAlex ID. This will be an Institution.id if the source is a repository, and a Publisher.id if the source is a journal, conference, or eBook platform (based on the type field). | [optional] 
**HostOrganizationLineage** | Pointer to **[]string** | OpenAlex IDs of the host organization&#39;s lineage. This will only be included if the host_organization is a publisher (and not if the host_organization is an institution). | [optional] 
**HostOrganizationName** | Pointer to **string** | The display_name from the host_organization, shown for convenience. | [optional] 
**Id** | Pointer to **string** | The OpenAlex ID for this source. | [optional] 
**Ids** | Pointer to [**ListSources200ResponseResultsInnerIds**](ListSources200ResponseResultsInnerIds.md) |  | [optional] 
**IsCore** | Pointer to **bool** | Whether this source is identified as a \&quot;core source\&quot; by [CWTS](https://www.cwts.nl/), used in the [Open Leiden Ranking](https://open.leidenranking.com/) of universities around the world. The list of core sources can be found [here](https://zenodo.org/records/10949671). | [optional] 
**IsInDoaj** | Pointer to **bool** | Whether this is a journal listed in the Directory of Open Access Journals (DOAJ). | [optional] 
**IsOa** | Pointer to **bool** | Whether this is currently fully-open-access source. This could be true for a preprint repository where everything uploaded is free to read, or for a Gold or Diamond open access journal, where all newly published Works are available for free under an open license. | [optional] 
**Issn** | Pointer to **[]string** | The ISSNs used by this source. Many publications have multiple ISSNs, so ISSN-L should be used when possible. | [optional] 
**IssnL** | Pointer to **string** | The ISSN-L identifying this source. This is the Canonical External ID for sources. ISSN-L or Linking ISSN solves the problem by designating a single canonical ISSN for all media versions of the title. It&#39;s usually the same as the print ISSN. | [optional] 
**Societies** | Pointer to [**[]ListSources200ResponseResultsInnerSocietiesInner**](ListSources200ResponseResultsInnerSocietiesInner.md) | Societies on whose behalf the source is published and maintained, obtained from our crowdsourced list. | [optional] 
**SummaryStats** | Pointer to [**ListSources200ResponseResultsInnerSummaryStats**](ListSources200ResponseResultsInnerSummaryStats.md) |  | [optional] 
**Type** | Pointer to **string** | The type of source. | [optional] 
**UpdatedDate** | Pointer to **string** | The last time anything in this source object changed. Formatted as ISO 8601 extended format without time zone designator. | [optional] 
**WorksApiUrl** | Pointer to **string** | A URL that will get you a list of all this source&#39;s Works. We express this as an API URL (instead of just listing the works themselves) because sometimes a source&#39;s publication list is too long to reasonably fit into a single Source object. | [optional] 
**WorksCount** | Pointer to **int32** | The number of Works this source hosts. | [optional] 
**XConcepts** | Pointer to **[]map[string]interface{}** | The Concepts most frequently applied to works hosted by this source. | [optional] 

## Methods

### NewListSources200ResponseResultsInner

`func NewListSources200ResponseResultsInner() *ListSources200ResponseResultsInner`

NewListSources200ResponseResultsInner instantiates a new ListSources200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSources200ResponseResultsInnerWithDefaults

`func NewListSources200ResponseResultsInnerWithDefaults() *ListSources200ResponseResultsInner`

NewListSources200ResponseResultsInnerWithDefaults instantiates a new ListSources200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbbreviatedTitle

`func (o *ListSources200ResponseResultsInner) GetAbbreviatedTitle() string`

GetAbbreviatedTitle returns the AbbreviatedTitle field if non-nil, zero value otherwise.

### GetAbbreviatedTitleOk

`func (o *ListSources200ResponseResultsInner) GetAbbreviatedTitleOk() (*string, bool)`

GetAbbreviatedTitleOk returns a tuple with the AbbreviatedTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbbreviatedTitle

`func (o *ListSources200ResponseResultsInner) SetAbbreviatedTitle(v string)`

SetAbbreviatedTitle sets AbbreviatedTitle field to given value.

### HasAbbreviatedTitle

`func (o *ListSources200ResponseResultsInner) HasAbbreviatedTitle() bool`

HasAbbreviatedTitle returns a boolean if a field has been set.

### GetAlternateTitles

`func (o *ListSources200ResponseResultsInner) GetAlternateTitles() []string`

GetAlternateTitles returns the AlternateTitles field if non-nil, zero value otherwise.

### GetAlternateTitlesOk

`func (o *ListSources200ResponseResultsInner) GetAlternateTitlesOk() (*[]string, bool)`

GetAlternateTitlesOk returns a tuple with the AlternateTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateTitles

`func (o *ListSources200ResponseResultsInner) SetAlternateTitles(v []string)`

SetAlternateTitles sets AlternateTitles field to given value.

### HasAlternateTitles

`func (o *ListSources200ResponseResultsInner) HasAlternateTitles() bool`

HasAlternateTitles returns a boolean if a field has been set.

### GetApcPrices

`func (o *ListSources200ResponseResultsInner) GetApcPrices() []ListSources200ResponseResultsInnerApcPricesInner`

GetApcPrices returns the ApcPrices field if non-nil, zero value otherwise.

### GetApcPricesOk

`func (o *ListSources200ResponseResultsInner) GetApcPricesOk() (*[]ListSources200ResponseResultsInnerApcPricesInner, bool)`

GetApcPricesOk returns a tuple with the ApcPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApcPrices

`func (o *ListSources200ResponseResultsInner) SetApcPrices(v []ListSources200ResponseResultsInnerApcPricesInner)`

SetApcPrices sets ApcPrices field to given value.

### HasApcPrices

`func (o *ListSources200ResponseResultsInner) HasApcPrices() bool`

HasApcPrices returns a boolean if a field has been set.

### GetApcUsd

`func (o *ListSources200ResponseResultsInner) GetApcUsd() int32`

GetApcUsd returns the ApcUsd field if non-nil, zero value otherwise.

### GetApcUsdOk

`func (o *ListSources200ResponseResultsInner) GetApcUsdOk() (*int32, bool)`

GetApcUsdOk returns a tuple with the ApcUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApcUsd

`func (o *ListSources200ResponseResultsInner) SetApcUsd(v int32)`

SetApcUsd sets ApcUsd field to given value.

### HasApcUsd

`func (o *ListSources200ResponseResultsInner) HasApcUsd() bool`

HasApcUsd returns a boolean if a field has been set.

### GetCitedByCount

`func (o *ListSources200ResponseResultsInner) GetCitedByCount() int32`

GetCitedByCount returns the CitedByCount field if non-nil, zero value otherwise.

### GetCitedByCountOk

`func (o *ListSources200ResponseResultsInner) GetCitedByCountOk() (*int32, bool)`

GetCitedByCountOk returns a tuple with the CitedByCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitedByCount

`func (o *ListSources200ResponseResultsInner) SetCitedByCount(v int32)`

SetCitedByCount sets CitedByCount field to given value.

### HasCitedByCount

`func (o *ListSources200ResponseResultsInner) HasCitedByCount() bool`

HasCitedByCount returns a boolean if a field has been set.

### GetCountryCode

`func (o *ListSources200ResponseResultsInner) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ListSources200ResponseResultsInner) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ListSources200ResponseResultsInner) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *ListSources200ResponseResultsInner) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCountsByYear

`func (o *ListSources200ResponseResultsInner) GetCountsByYear() []ListAuthors200ResponseResultsInnerCountsByYearInner`

GetCountsByYear returns the CountsByYear field if non-nil, zero value otherwise.

### GetCountsByYearOk

`func (o *ListSources200ResponseResultsInner) GetCountsByYearOk() (*[]ListAuthors200ResponseResultsInnerCountsByYearInner, bool)`

GetCountsByYearOk returns a tuple with the CountsByYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountsByYear

`func (o *ListSources200ResponseResultsInner) SetCountsByYear(v []ListAuthors200ResponseResultsInnerCountsByYearInner)`

SetCountsByYear sets CountsByYear field to given value.

### HasCountsByYear

`func (o *ListSources200ResponseResultsInner) HasCountsByYear() bool`

HasCountsByYear returns a boolean if a field has been set.

### GetCreatedDate

`func (o *ListSources200ResponseResultsInner) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ListSources200ResponseResultsInner) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ListSources200ResponseResultsInner) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ListSources200ResponseResultsInner) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetDisplayName

`func (o *ListSources200ResponseResultsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListSources200ResponseResultsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListSources200ResponseResultsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ListSources200ResponseResultsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHomepageUrl

`func (o *ListSources200ResponseResultsInner) GetHomepageUrl() string`

GetHomepageUrl returns the HomepageUrl field if non-nil, zero value otherwise.

### GetHomepageUrlOk

`func (o *ListSources200ResponseResultsInner) GetHomepageUrlOk() (*string, bool)`

GetHomepageUrlOk returns a tuple with the HomepageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepageUrl

`func (o *ListSources200ResponseResultsInner) SetHomepageUrl(v string)`

SetHomepageUrl sets HomepageUrl field to given value.

### HasHomepageUrl

`func (o *ListSources200ResponseResultsInner) HasHomepageUrl() bool`

HasHomepageUrl returns a boolean if a field has been set.

### GetHostOrganization

`func (o *ListSources200ResponseResultsInner) GetHostOrganization() string`

GetHostOrganization returns the HostOrganization field if non-nil, zero value otherwise.

### GetHostOrganizationOk

`func (o *ListSources200ResponseResultsInner) GetHostOrganizationOk() (*string, bool)`

GetHostOrganizationOk returns a tuple with the HostOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOrganization

`func (o *ListSources200ResponseResultsInner) SetHostOrganization(v string)`

SetHostOrganization sets HostOrganization field to given value.

### HasHostOrganization

`func (o *ListSources200ResponseResultsInner) HasHostOrganization() bool`

HasHostOrganization returns a boolean if a field has been set.

### GetHostOrganizationLineage

`func (o *ListSources200ResponseResultsInner) GetHostOrganizationLineage() []string`

GetHostOrganizationLineage returns the HostOrganizationLineage field if non-nil, zero value otherwise.

### GetHostOrganizationLineageOk

`func (o *ListSources200ResponseResultsInner) GetHostOrganizationLineageOk() (*[]string, bool)`

GetHostOrganizationLineageOk returns a tuple with the HostOrganizationLineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOrganizationLineage

`func (o *ListSources200ResponseResultsInner) SetHostOrganizationLineage(v []string)`

SetHostOrganizationLineage sets HostOrganizationLineage field to given value.

### HasHostOrganizationLineage

`func (o *ListSources200ResponseResultsInner) HasHostOrganizationLineage() bool`

HasHostOrganizationLineage returns a boolean if a field has been set.

### GetHostOrganizationName

`func (o *ListSources200ResponseResultsInner) GetHostOrganizationName() string`

GetHostOrganizationName returns the HostOrganizationName field if non-nil, zero value otherwise.

### GetHostOrganizationNameOk

`func (o *ListSources200ResponseResultsInner) GetHostOrganizationNameOk() (*string, bool)`

GetHostOrganizationNameOk returns a tuple with the HostOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOrganizationName

`func (o *ListSources200ResponseResultsInner) SetHostOrganizationName(v string)`

SetHostOrganizationName sets HostOrganizationName field to given value.

### HasHostOrganizationName

`func (o *ListSources200ResponseResultsInner) HasHostOrganizationName() bool`

HasHostOrganizationName returns a boolean if a field has been set.

### GetId

`func (o *ListSources200ResponseResultsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListSources200ResponseResultsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListSources200ResponseResultsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListSources200ResponseResultsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIds

`func (o *ListSources200ResponseResultsInner) GetIds() ListSources200ResponseResultsInnerIds`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListSources200ResponseResultsInner) GetIdsOk() (*ListSources200ResponseResultsInnerIds, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListSources200ResponseResultsInner) SetIds(v ListSources200ResponseResultsInnerIds)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ListSources200ResponseResultsInner) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetIsCore

`func (o *ListSources200ResponseResultsInner) GetIsCore() bool`

GetIsCore returns the IsCore field if non-nil, zero value otherwise.

### GetIsCoreOk

`func (o *ListSources200ResponseResultsInner) GetIsCoreOk() (*bool, bool)`

GetIsCoreOk returns a tuple with the IsCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCore

`func (o *ListSources200ResponseResultsInner) SetIsCore(v bool)`

SetIsCore sets IsCore field to given value.

### HasIsCore

`func (o *ListSources200ResponseResultsInner) HasIsCore() bool`

HasIsCore returns a boolean if a field has been set.

### GetIsInDoaj

`func (o *ListSources200ResponseResultsInner) GetIsInDoaj() bool`

GetIsInDoaj returns the IsInDoaj field if non-nil, zero value otherwise.

### GetIsInDoajOk

`func (o *ListSources200ResponseResultsInner) GetIsInDoajOk() (*bool, bool)`

GetIsInDoajOk returns a tuple with the IsInDoaj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInDoaj

`func (o *ListSources200ResponseResultsInner) SetIsInDoaj(v bool)`

SetIsInDoaj sets IsInDoaj field to given value.

### HasIsInDoaj

`func (o *ListSources200ResponseResultsInner) HasIsInDoaj() bool`

HasIsInDoaj returns a boolean if a field has been set.

### GetIsOa

`func (o *ListSources200ResponseResultsInner) GetIsOa() bool`

GetIsOa returns the IsOa field if non-nil, zero value otherwise.

### GetIsOaOk

`func (o *ListSources200ResponseResultsInner) GetIsOaOk() (*bool, bool)`

GetIsOaOk returns a tuple with the IsOa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOa

`func (o *ListSources200ResponseResultsInner) SetIsOa(v bool)`

SetIsOa sets IsOa field to given value.

### HasIsOa

`func (o *ListSources200ResponseResultsInner) HasIsOa() bool`

HasIsOa returns a boolean if a field has been set.

### GetIssn

`func (o *ListSources200ResponseResultsInner) GetIssn() []string`

GetIssn returns the Issn field if non-nil, zero value otherwise.

### GetIssnOk

`func (o *ListSources200ResponseResultsInner) GetIssnOk() (*[]string, bool)`

GetIssnOk returns a tuple with the Issn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssn

`func (o *ListSources200ResponseResultsInner) SetIssn(v []string)`

SetIssn sets Issn field to given value.

### HasIssn

`func (o *ListSources200ResponseResultsInner) HasIssn() bool`

HasIssn returns a boolean if a field has been set.

### GetIssnL

`func (o *ListSources200ResponseResultsInner) GetIssnL() string`

GetIssnL returns the IssnL field if non-nil, zero value otherwise.

### GetIssnLOk

`func (o *ListSources200ResponseResultsInner) GetIssnLOk() (*string, bool)`

GetIssnLOk returns a tuple with the IssnL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssnL

`func (o *ListSources200ResponseResultsInner) SetIssnL(v string)`

SetIssnL sets IssnL field to given value.

### HasIssnL

`func (o *ListSources200ResponseResultsInner) HasIssnL() bool`

HasIssnL returns a boolean if a field has been set.

### GetSocieties

`func (o *ListSources200ResponseResultsInner) GetSocieties() []ListSources200ResponseResultsInnerSocietiesInner`

GetSocieties returns the Societies field if non-nil, zero value otherwise.

### GetSocietiesOk

`func (o *ListSources200ResponseResultsInner) GetSocietiesOk() (*[]ListSources200ResponseResultsInnerSocietiesInner, bool)`

GetSocietiesOk returns a tuple with the Societies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocieties

`func (o *ListSources200ResponseResultsInner) SetSocieties(v []ListSources200ResponseResultsInnerSocietiesInner)`

SetSocieties sets Societies field to given value.

### HasSocieties

`func (o *ListSources200ResponseResultsInner) HasSocieties() bool`

HasSocieties returns a boolean if a field has been set.

### GetSummaryStats

`func (o *ListSources200ResponseResultsInner) GetSummaryStats() ListSources200ResponseResultsInnerSummaryStats`

GetSummaryStats returns the SummaryStats field if non-nil, zero value otherwise.

### GetSummaryStatsOk

`func (o *ListSources200ResponseResultsInner) GetSummaryStatsOk() (*ListSources200ResponseResultsInnerSummaryStats, bool)`

GetSummaryStatsOk returns a tuple with the SummaryStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryStats

`func (o *ListSources200ResponseResultsInner) SetSummaryStats(v ListSources200ResponseResultsInnerSummaryStats)`

SetSummaryStats sets SummaryStats field to given value.

### HasSummaryStats

`func (o *ListSources200ResponseResultsInner) HasSummaryStats() bool`

HasSummaryStats returns a boolean if a field has been set.

### GetType

`func (o *ListSources200ResponseResultsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListSources200ResponseResultsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListSources200ResponseResultsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListSources200ResponseResultsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedDate

`func (o *ListSources200ResponseResultsInner) GetUpdatedDate() string`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *ListSources200ResponseResultsInner) GetUpdatedDateOk() (*string, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *ListSources200ResponseResultsInner) SetUpdatedDate(v string)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *ListSources200ResponseResultsInner) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.

### GetWorksApiUrl

`func (o *ListSources200ResponseResultsInner) GetWorksApiUrl() string`

GetWorksApiUrl returns the WorksApiUrl field if non-nil, zero value otherwise.

### GetWorksApiUrlOk

`func (o *ListSources200ResponseResultsInner) GetWorksApiUrlOk() (*string, bool)`

GetWorksApiUrlOk returns a tuple with the WorksApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksApiUrl

`func (o *ListSources200ResponseResultsInner) SetWorksApiUrl(v string)`

SetWorksApiUrl sets WorksApiUrl field to given value.

### HasWorksApiUrl

`func (o *ListSources200ResponseResultsInner) HasWorksApiUrl() bool`

HasWorksApiUrl returns a boolean if a field has been set.

### GetWorksCount

`func (o *ListSources200ResponseResultsInner) GetWorksCount() int32`

GetWorksCount returns the WorksCount field if non-nil, zero value otherwise.

### GetWorksCountOk

`func (o *ListSources200ResponseResultsInner) GetWorksCountOk() (*int32, bool)`

GetWorksCountOk returns a tuple with the WorksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksCount

`func (o *ListSources200ResponseResultsInner) SetWorksCount(v int32)`

SetWorksCount sets WorksCount field to given value.

### HasWorksCount

`func (o *ListSources200ResponseResultsInner) HasWorksCount() bool`

HasWorksCount returns a boolean if a field has been set.

### GetXConcepts

`func (o *ListSources200ResponseResultsInner) GetXConcepts() []map[string]interface{}`

GetXConcepts returns the XConcepts field if non-nil, zero value otherwise.

### GetXConceptsOk

`func (o *ListSources200ResponseResultsInner) GetXConceptsOk() (*[]map[string]interface{}, bool)`

GetXConceptsOk returns a tuple with the XConcepts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXConcepts

`func (o *ListSources200ResponseResultsInner) SetXConcepts(v []map[string]interface{})`

SetXConcepts sets XConcepts field to given value.

### HasXConcepts

`func (o *ListSources200ResponseResultsInner) HasXConcepts() bool`

HasXConcepts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


