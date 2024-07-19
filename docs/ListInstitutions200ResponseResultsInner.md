# ListInstitutions200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedInstitutions** | Pointer to [**[]ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner**](ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner.md) | List of institutions related to this one, represented as dehydrated Institution objects with an additional &#39;relationship&#39; property. | [optional] 
**CitedByCount** | Pointer to **int32** | The total number of works that cite a work created by an author affiliated with this institution. | [optional] 
**CountryCode** | Pointer to **string** | The country where this institution is located, represented as an ISO two-letter country code. | [optional] 
**CountsByYear** | Pointer to [**[]ListAuthors200ResponseResultsInnerCountsByYearInner**](ListAuthors200ResponseResultsInnerCountsByYearInner.md) | Works count and cited by count for each of the last ten years, binned by year. | [optional] 
**CreatedDate** | Pointer to **string** | The date this Institution object was created in the OpenAlex dataset. | [optional] 
**DisplayName** | Pointer to **string** | The primary name of the institution. | [optional] 
**DisplayNameAcronyms** | Pointer to **[]string** | Acronyms or initialisms that people sometimes use instead of the full display name. | [optional] 
**DisplayNameAlternatives** | Pointer to **[]string** | Other names people may use for this institution. | [optional] 
**Geo** | Pointer to [**ListInstitutions200ResponseResultsInnerGeo**](ListInstitutions200ResponseResultsInnerGeo.md) |  | [optional] 
**HomepageUrl** | Pointer to **string** | The URL for institution&#39;s primary homepage. | [optional] 
**Id** | Pointer to **string** | The OpenAlex ID for this institution. | [optional] 
**Ids** | Pointer to [**ListInstitutions200ResponseResultsInnerIds**](ListInstitutions200ResponseResultsInnerIds.md) |  | [optional] 
**ImageThumbnailUrl** | Pointer to **string** | URL for a thumbnail image representing this institution. | [optional] 
**ImageUrl** | Pointer to **string** | URL for an image representing this institution. | [optional] 
**International** | Pointer to [**ListInstitutions200ResponseResultsInnerInternational**](ListInstitutions200ResponseResultsInnerInternational.md) |  | [optional] 
**Lineage** | Pointer to **[]string** | OpenAlex IDs of this institution and its parent institutions. | [optional] 
**Repositories** | Pointer to [**[]ListInstitutions200ResponseResultsInnerRepositoriesInner**](ListInstitutions200ResponseResultsInnerRepositoriesInner.md) | Repositories that have this institution as their host organization. | [optional] 
**Roles** | Pointer to [**[]ListFunders200ResponseResultsInnerRolesInner**](ListFunders200ResponseResultsInnerRolesInner.md) | List of roles this institution plays, including itself and potentially funder and publisher roles. | [optional] 
**Ror** | Pointer to **string** | The ROR ID for this institution. | [optional] 
**SummaryStats** | Pointer to [**ListInstitutions200ResponseResultsInnerSummaryStats**](ListInstitutions200ResponseResultsInnerSummaryStats.md) |  | [optional] 
**Type** | Pointer to **string** | The institution&#39;s primary type, using the ROR \&quot;type\&quot; controlled vocabulary. | [optional] 
**UpdatedDate** | Pointer to **string** | The last time anything in this author object changed. Formatted as ISO 8601 extended format without time zone designator. | [optional] 
**WorksApiUrl** | Pointer to **string** | A URL that will get you a list of all the Works affiliated with this institution. | [optional] 
**WorksCount** | Pointer to **int32** | The number of Works created by authors affiliated with this institution. | [optional] 
**XConcepts** | Pointer to **[]map[string]interface{}** | The Concepts most frequently applied to works affiliated with this institution. | [optional] 

## Methods

### NewListInstitutions200ResponseResultsInner

`func NewListInstitutions200ResponseResultsInner() *ListInstitutions200ResponseResultsInner`

NewListInstitutions200ResponseResultsInner instantiates a new ListInstitutions200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstitutions200ResponseResultsInnerWithDefaults

`func NewListInstitutions200ResponseResultsInnerWithDefaults() *ListInstitutions200ResponseResultsInner`

NewListInstitutions200ResponseResultsInnerWithDefaults instantiates a new ListInstitutions200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatedInstitutions

`func (o *ListInstitutions200ResponseResultsInner) GetAssociatedInstitutions() []ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner`

GetAssociatedInstitutions returns the AssociatedInstitutions field if non-nil, zero value otherwise.

### GetAssociatedInstitutionsOk

`func (o *ListInstitutions200ResponseResultsInner) GetAssociatedInstitutionsOk() (*[]ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner, bool)`

GetAssociatedInstitutionsOk returns a tuple with the AssociatedInstitutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedInstitutions

`func (o *ListInstitutions200ResponseResultsInner) SetAssociatedInstitutions(v []ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner)`

SetAssociatedInstitutions sets AssociatedInstitutions field to given value.

### HasAssociatedInstitutions

`func (o *ListInstitutions200ResponseResultsInner) HasAssociatedInstitutions() bool`

HasAssociatedInstitutions returns a boolean if a field has been set.

### GetCitedByCount

`func (o *ListInstitutions200ResponseResultsInner) GetCitedByCount() int32`

GetCitedByCount returns the CitedByCount field if non-nil, zero value otherwise.

### GetCitedByCountOk

`func (o *ListInstitutions200ResponseResultsInner) GetCitedByCountOk() (*int32, bool)`

GetCitedByCountOk returns a tuple with the CitedByCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitedByCount

`func (o *ListInstitutions200ResponseResultsInner) SetCitedByCount(v int32)`

SetCitedByCount sets CitedByCount field to given value.

### HasCitedByCount

`func (o *ListInstitutions200ResponseResultsInner) HasCitedByCount() bool`

HasCitedByCount returns a boolean if a field has been set.

### GetCountryCode

`func (o *ListInstitutions200ResponseResultsInner) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ListInstitutions200ResponseResultsInner) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ListInstitutions200ResponseResultsInner) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *ListInstitutions200ResponseResultsInner) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCountsByYear

`func (o *ListInstitutions200ResponseResultsInner) GetCountsByYear() []ListAuthors200ResponseResultsInnerCountsByYearInner`

GetCountsByYear returns the CountsByYear field if non-nil, zero value otherwise.

### GetCountsByYearOk

`func (o *ListInstitutions200ResponseResultsInner) GetCountsByYearOk() (*[]ListAuthors200ResponseResultsInnerCountsByYearInner, bool)`

GetCountsByYearOk returns a tuple with the CountsByYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountsByYear

`func (o *ListInstitutions200ResponseResultsInner) SetCountsByYear(v []ListAuthors200ResponseResultsInnerCountsByYearInner)`

SetCountsByYear sets CountsByYear field to given value.

### HasCountsByYear

`func (o *ListInstitutions200ResponseResultsInner) HasCountsByYear() bool`

HasCountsByYear returns a boolean if a field has been set.

### GetCreatedDate

`func (o *ListInstitutions200ResponseResultsInner) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ListInstitutions200ResponseResultsInner) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ListInstitutions200ResponseResultsInner) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ListInstitutions200ResponseResultsInner) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetDisplayName

`func (o *ListInstitutions200ResponseResultsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListInstitutions200ResponseResultsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListInstitutions200ResponseResultsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ListInstitutions200ResponseResultsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDisplayNameAcronyms

`func (o *ListInstitutions200ResponseResultsInner) GetDisplayNameAcronyms() []string`

GetDisplayNameAcronyms returns the DisplayNameAcronyms field if non-nil, zero value otherwise.

### GetDisplayNameAcronymsOk

`func (o *ListInstitutions200ResponseResultsInner) GetDisplayNameAcronymsOk() (*[]string, bool)`

GetDisplayNameAcronymsOk returns a tuple with the DisplayNameAcronyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNameAcronyms

`func (o *ListInstitutions200ResponseResultsInner) SetDisplayNameAcronyms(v []string)`

SetDisplayNameAcronyms sets DisplayNameAcronyms field to given value.

### HasDisplayNameAcronyms

`func (o *ListInstitutions200ResponseResultsInner) HasDisplayNameAcronyms() bool`

HasDisplayNameAcronyms returns a boolean if a field has been set.

### GetDisplayNameAlternatives

`func (o *ListInstitutions200ResponseResultsInner) GetDisplayNameAlternatives() []string`

GetDisplayNameAlternatives returns the DisplayNameAlternatives field if non-nil, zero value otherwise.

### GetDisplayNameAlternativesOk

`func (o *ListInstitutions200ResponseResultsInner) GetDisplayNameAlternativesOk() (*[]string, bool)`

GetDisplayNameAlternativesOk returns a tuple with the DisplayNameAlternatives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNameAlternatives

`func (o *ListInstitutions200ResponseResultsInner) SetDisplayNameAlternatives(v []string)`

SetDisplayNameAlternatives sets DisplayNameAlternatives field to given value.

### HasDisplayNameAlternatives

`func (o *ListInstitutions200ResponseResultsInner) HasDisplayNameAlternatives() bool`

HasDisplayNameAlternatives returns a boolean if a field has been set.

### GetGeo

`func (o *ListInstitutions200ResponseResultsInner) GetGeo() ListInstitutions200ResponseResultsInnerGeo`

GetGeo returns the Geo field if non-nil, zero value otherwise.

### GetGeoOk

`func (o *ListInstitutions200ResponseResultsInner) GetGeoOk() (*ListInstitutions200ResponseResultsInnerGeo, bool)`

GetGeoOk returns a tuple with the Geo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeo

`func (o *ListInstitutions200ResponseResultsInner) SetGeo(v ListInstitutions200ResponseResultsInnerGeo)`

SetGeo sets Geo field to given value.

### HasGeo

`func (o *ListInstitutions200ResponseResultsInner) HasGeo() bool`

HasGeo returns a boolean if a field has been set.

### GetHomepageUrl

`func (o *ListInstitutions200ResponseResultsInner) GetHomepageUrl() string`

GetHomepageUrl returns the HomepageUrl field if non-nil, zero value otherwise.

### GetHomepageUrlOk

`func (o *ListInstitutions200ResponseResultsInner) GetHomepageUrlOk() (*string, bool)`

GetHomepageUrlOk returns a tuple with the HomepageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepageUrl

`func (o *ListInstitutions200ResponseResultsInner) SetHomepageUrl(v string)`

SetHomepageUrl sets HomepageUrl field to given value.

### HasHomepageUrl

`func (o *ListInstitutions200ResponseResultsInner) HasHomepageUrl() bool`

HasHomepageUrl returns a boolean if a field has been set.

### GetId

`func (o *ListInstitutions200ResponseResultsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListInstitutions200ResponseResultsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListInstitutions200ResponseResultsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListInstitutions200ResponseResultsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIds

`func (o *ListInstitutions200ResponseResultsInner) GetIds() ListInstitutions200ResponseResultsInnerIds`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListInstitutions200ResponseResultsInner) GetIdsOk() (*ListInstitutions200ResponseResultsInnerIds, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListInstitutions200ResponseResultsInner) SetIds(v ListInstitutions200ResponseResultsInnerIds)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ListInstitutions200ResponseResultsInner) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetImageThumbnailUrl

`func (o *ListInstitutions200ResponseResultsInner) GetImageThumbnailUrl() string`

GetImageThumbnailUrl returns the ImageThumbnailUrl field if non-nil, zero value otherwise.

### GetImageThumbnailUrlOk

`func (o *ListInstitutions200ResponseResultsInner) GetImageThumbnailUrlOk() (*string, bool)`

GetImageThumbnailUrlOk returns a tuple with the ImageThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageThumbnailUrl

`func (o *ListInstitutions200ResponseResultsInner) SetImageThumbnailUrl(v string)`

SetImageThumbnailUrl sets ImageThumbnailUrl field to given value.

### HasImageThumbnailUrl

`func (o *ListInstitutions200ResponseResultsInner) HasImageThumbnailUrl() bool`

HasImageThumbnailUrl returns a boolean if a field has been set.

### GetImageUrl

`func (o *ListInstitutions200ResponseResultsInner) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *ListInstitutions200ResponseResultsInner) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *ListInstitutions200ResponseResultsInner) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *ListInstitutions200ResponseResultsInner) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetInternational

`func (o *ListInstitutions200ResponseResultsInner) GetInternational() ListInstitutions200ResponseResultsInnerInternational`

GetInternational returns the International field if non-nil, zero value otherwise.

### GetInternationalOk

`func (o *ListInstitutions200ResponseResultsInner) GetInternationalOk() (*ListInstitutions200ResponseResultsInnerInternational, bool)`

GetInternationalOk returns a tuple with the International field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternational

`func (o *ListInstitutions200ResponseResultsInner) SetInternational(v ListInstitutions200ResponseResultsInnerInternational)`

SetInternational sets International field to given value.

### HasInternational

`func (o *ListInstitutions200ResponseResultsInner) HasInternational() bool`

HasInternational returns a boolean if a field has been set.

### GetLineage

`func (o *ListInstitutions200ResponseResultsInner) GetLineage() []string`

GetLineage returns the Lineage field if non-nil, zero value otherwise.

### GetLineageOk

`func (o *ListInstitutions200ResponseResultsInner) GetLineageOk() (*[]string, bool)`

GetLineageOk returns a tuple with the Lineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineage

`func (o *ListInstitutions200ResponseResultsInner) SetLineage(v []string)`

SetLineage sets Lineage field to given value.

### HasLineage

`func (o *ListInstitutions200ResponseResultsInner) HasLineage() bool`

HasLineage returns a boolean if a field has been set.

### GetRepositories

`func (o *ListInstitutions200ResponseResultsInner) GetRepositories() []ListInstitutions200ResponseResultsInnerRepositoriesInner`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *ListInstitutions200ResponseResultsInner) GetRepositoriesOk() (*[]ListInstitutions200ResponseResultsInnerRepositoriesInner, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *ListInstitutions200ResponseResultsInner) SetRepositories(v []ListInstitutions200ResponseResultsInnerRepositoriesInner)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *ListInstitutions200ResponseResultsInner) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetRoles

`func (o *ListInstitutions200ResponseResultsInner) GetRoles() []ListFunders200ResponseResultsInnerRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ListInstitutions200ResponseResultsInner) GetRolesOk() (*[]ListFunders200ResponseResultsInnerRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ListInstitutions200ResponseResultsInner) SetRoles(v []ListFunders200ResponseResultsInnerRolesInner)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ListInstitutions200ResponseResultsInner) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetRor

`func (o *ListInstitutions200ResponseResultsInner) GetRor() string`

GetRor returns the Ror field if non-nil, zero value otherwise.

### GetRorOk

`func (o *ListInstitutions200ResponseResultsInner) GetRorOk() (*string, bool)`

GetRorOk returns a tuple with the Ror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRor

`func (o *ListInstitutions200ResponseResultsInner) SetRor(v string)`

SetRor sets Ror field to given value.

### HasRor

`func (o *ListInstitutions200ResponseResultsInner) HasRor() bool`

HasRor returns a boolean if a field has been set.

### GetSummaryStats

`func (o *ListInstitutions200ResponseResultsInner) GetSummaryStats() ListInstitutions200ResponseResultsInnerSummaryStats`

GetSummaryStats returns the SummaryStats field if non-nil, zero value otherwise.

### GetSummaryStatsOk

`func (o *ListInstitutions200ResponseResultsInner) GetSummaryStatsOk() (*ListInstitutions200ResponseResultsInnerSummaryStats, bool)`

GetSummaryStatsOk returns a tuple with the SummaryStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryStats

`func (o *ListInstitutions200ResponseResultsInner) SetSummaryStats(v ListInstitutions200ResponseResultsInnerSummaryStats)`

SetSummaryStats sets SummaryStats field to given value.

### HasSummaryStats

`func (o *ListInstitutions200ResponseResultsInner) HasSummaryStats() bool`

HasSummaryStats returns a boolean if a field has been set.

### GetType

`func (o *ListInstitutions200ResponseResultsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListInstitutions200ResponseResultsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListInstitutions200ResponseResultsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListInstitutions200ResponseResultsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedDate

`func (o *ListInstitutions200ResponseResultsInner) GetUpdatedDate() string`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *ListInstitutions200ResponseResultsInner) GetUpdatedDateOk() (*string, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *ListInstitutions200ResponseResultsInner) SetUpdatedDate(v string)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *ListInstitutions200ResponseResultsInner) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.

### GetWorksApiUrl

`func (o *ListInstitutions200ResponseResultsInner) GetWorksApiUrl() string`

GetWorksApiUrl returns the WorksApiUrl field if non-nil, zero value otherwise.

### GetWorksApiUrlOk

`func (o *ListInstitutions200ResponseResultsInner) GetWorksApiUrlOk() (*string, bool)`

GetWorksApiUrlOk returns a tuple with the WorksApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksApiUrl

`func (o *ListInstitutions200ResponseResultsInner) SetWorksApiUrl(v string)`

SetWorksApiUrl sets WorksApiUrl field to given value.

### HasWorksApiUrl

`func (o *ListInstitutions200ResponseResultsInner) HasWorksApiUrl() bool`

HasWorksApiUrl returns a boolean if a field has been set.

### GetWorksCount

`func (o *ListInstitutions200ResponseResultsInner) GetWorksCount() int32`

GetWorksCount returns the WorksCount field if non-nil, zero value otherwise.

### GetWorksCountOk

`func (o *ListInstitutions200ResponseResultsInner) GetWorksCountOk() (*int32, bool)`

GetWorksCountOk returns a tuple with the WorksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksCount

`func (o *ListInstitutions200ResponseResultsInner) SetWorksCount(v int32)`

SetWorksCount sets WorksCount field to given value.

### HasWorksCount

`func (o *ListInstitutions200ResponseResultsInner) HasWorksCount() bool`

HasWorksCount returns a boolean if a field has been set.

### GetXConcepts

`func (o *ListInstitutions200ResponseResultsInner) GetXConcepts() []map[string]interface{}`

GetXConcepts returns the XConcepts field if non-nil, zero value otherwise.

### GetXConceptsOk

`func (o *ListInstitutions200ResponseResultsInner) GetXConceptsOk() (*[]map[string]interface{}, bool)`

GetXConceptsOk returns a tuple with the XConcepts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXConcepts

`func (o *ListInstitutions200ResponseResultsInner) SetXConcepts(v []map[string]interface{})`

SetXConcepts sets XConcepts field to given value.

### HasXConcepts

`func (o *ListInstitutions200ResponseResultsInner) HasXConcepts() bool`

HasXConcepts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


