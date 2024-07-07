# ListFunders200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternateTitles** | Pointer to **[]string** | A list of alternate titles for this funder. | [optional] 
**CitedByCount** | Pointer to **int32** | The total number of Works that cite a work linked to this funder. | [optional] 
**CountryCode** | Pointer to **string** | The country where this funder is located, represented as an ISO two-letter country code. | [optional] 
**CountsByYear** | Pointer to [**[]ListAuthors200ResponseResultsInnerCountsByYearInner**](ListAuthors200ResponseResultsInnerCountsByYearInner.md) | The values of works_count and cited_by_count for each of the last ten years, binned by year. | [optional] 
**CreatedDate** | Pointer to **string** | The date this Funder object was created in the OpenAlex dataset, expressed as an ISO 8601 date string. | [optional] 
**Description** | Pointer to **string** | A short description of this funder, taken from Wikidata. | [optional] 
**DisplayName** | Pointer to **string** | The primary name of the funder. | [optional] 
**GrantsCount** | Pointer to **int32** | The number of grants linked to this funder. | [optional] 
**HomepageUrl** | Pointer to **string** | The URL for this funder&#39;s primary homepage. | [optional] 
**Id** | Pointer to **string** | The OpenAlex ID for this funder. | [optional] 
**Ids** | Pointer to [**ListFunders200ResponseResultsInnerIds**](ListFunders200ResponseResultsInnerIds.md) |  | [optional] 
**ImageThumbnailUrl** | Pointer to **string** | Same as image_url, but it&#39;s a smaller image. | [optional] 
**ImageUrl** | Pointer to **string** | URL where you can get an image representing this funder. | [optional] 
**Roles** | Pointer to [**[]ListFunders200ResponseResultsInnerRolesInner**](ListFunders200ResponseResultsInnerRolesInner.md) | List of role objects, which include the role, the id, and the works_count. | [optional] 
**SummaryStats** | Pointer to [**ListFunders200ResponseResultsInnerSummaryStats**](ListFunders200ResponseResultsInnerSummaryStats.md) |  | [optional] 
**UpdatedDate** | Pointer to **time.Time** | The last time anything in this funder object changed, expressed as an ISO 8601 date string. | [optional] 
**WorksCount** | Pointer to **int32** | The number of works linked to this funder. | [optional] 

## Methods

### NewListFunders200ResponseResultsInner

`func NewListFunders200ResponseResultsInner() *ListFunders200ResponseResultsInner`

NewListFunders200ResponseResultsInner instantiates a new ListFunders200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFunders200ResponseResultsInnerWithDefaults

`func NewListFunders200ResponseResultsInnerWithDefaults() *ListFunders200ResponseResultsInner`

NewListFunders200ResponseResultsInnerWithDefaults instantiates a new ListFunders200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateTitles

`func (o *ListFunders200ResponseResultsInner) GetAlternateTitles() []string`

GetAlternateTitles returns the AlternateTitles field if non-nil, zero value otherwise.

### GetAlternateTitlesOk

`func (o *ListFunders200ResponseResultsInner) GetAlternateTitlesOk() (*[]string, bool)`

GetAlternateTitlesOk returns a tuple with the AlternateTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateTitles

`func (o *ListFunders200ResponseResultsInner) SetAlternateTitles(v []string)`

SetAlternateTitles sets AlternateTitles field to given value.

### HasAlternateTitles

`func (o *ListFunders200ResponseResultsInner) HasAlternateTitles() bool`

HasAlternateTitles returns a boolean if a field has been set.

### GetCitedByCount

`func (o *ListFunders200ResponseResultsInner) GetCitedByCount() int32`

GetCitedByCount returns the CitedByCount field if non-nil, zero value otherwise.

### GetCitedByCountOk

`func (o *ListFunders200ResponseResultsInner) GetCitedByCountOk() (*int32, bool)`

GetCitedByCountOk returns a tuple with the CitedByCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitedByCount

`func (o *ListFunders200ResponseResultsInner) SetCitedByCount(v int32)`

SetCitedByCount sets CitedByCount field to given value.

### HasCitedByCount

`func (o *ListFunders200ResponseResultsInner) HasCitedByCount() bool`

HasCitedByCount returns a boolean if a field has been set.

### GetCountryCode

`func (o *ListFunders200ResponseResultsInner) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ListFunders200ResponseResultsInner) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ListFunders200ResponseResultsInner) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *ListFunders200ResponseResultsInner) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCountsByYear

`func (o *ListFunders200ResponseResultsInner) GetCountsByYear() []ListAuthors200ResponseResultsInnerCountsByYearInner`

GetCountsByYear returns the CountsByYear field if non-nil, zero value otherwise.

### GetCountsByYearOk

`func (o *ListFunders200ResponseResultsInner) GetCountsByYearOk() (*[]ListAuthors200ResponseResultsInnerCountsByYearInner, bool)`

GetCountsByYearOk returns a tuple with the CountsByYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountsByYear

`func (o *ListFunders200ResponseResultsInner) SetCountsByYear(v []ListAuthors200ResponseResultsInnerCountsByYearInner)`

SetCountsByYear sets CountsByYear field to given value.

### HasCountsByYear

`func (o *ListFunders200ResponseResultsInner) HasCountsByYear() bool`

HasCountsByYear returns a boolean if a field has been set.

### GetCreatedDate

`func (o *ListFunders200ResponseResultsInner) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ListFunders200ResponseResultsInner) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ListFunders200ResponseResultsInner) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ListFunders200ResponseResultsInner) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetDescription

`func (o *ListFunders200ResponseResultsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListFunders200ResponseResultsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListFunders200ResponseResultsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListFunders200ResponseResultsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *ListFunders200ResponseResultsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListFunders200ResponseResultsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListFunders200ResponseResultsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ListFunders200ResponseResultsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetGrantsCount

`func (o *ListFunders200ResponseResultsInner) GetGrantsCount() int32`

GetGrantsCount returns the GrantsCount field if non-nil, zero value otherwise.

### GetGrantsCountOk

`func (o *ListFunders200ResponseResultsInner) GetGrantsCountOk() (*int32, bool)`

GetGrantsCountOk returns a tuple with the GrantsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantsCount

`func (o *ListFunders200ResponseResultsInner) SetGrantsCount(v int32)`

SetGrantsCount sets GrantsCount field to given value.

### HasGrantsCount

`func (o *ListFunders200ResponseResultsInner) HasGrantsCount() bool`

HasGrantsCount returns a boolean if a field has been set.

### GetHomepageUrl

`func (o *ListFunders200ResponseResultsInner) GetHomepageUrl() string`

GetHomepageUrl returns the HomepageUrl field if non-nil, zero value otherwise.

### GetHomepageUrlOk

`func (o *ListFunders200ResponseResultsInner) GetHomepageUrlOk() (*string, bool)`

GetHomepageUrlOk returns a tuple with the HomepageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepageUrl

`func (o *ListFunders200ResponseResultsInner) SetHomepageUrl(v string)`

SetHomepageUrl sets HomepageUrl field to given value.

### HasHomepageUrl

`func (o *ListFunders200ResponseResultsInner) HasHomepageUrl() bool`

HasHomepageUrl returns a boolean if a field has been set.

### GetId

`func (o *ListFunders200ResponseResultsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListFunders200ResponseResultsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListFunders200ResponseResultsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListFunders200ResponseResultsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIds

`func (o *ListFunders200ResponseResultsInner) GetIds() ListFunders200ResponseResultsInnerIds`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListFunders200ResponseResultsInner) GetIdsOk() (*ListFunders200ResponseResultsInnerIds, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListFunders200ResponseResultsInner) SetIds(v ListFunders200ResponseResultsInnerIds)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ListFunders200ResponseResultsInner) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetImageThumbnailUrl

`func (o *ListFunders200ResponseResultsInner) GetImageThumbnailUrl() string`

GetImageThumbnailUrl returns the ImageThumbnailUrl field if non-nil, zero value otherwise.

### GetImageThumbnailUrlOk

`func (o *ListFunders200ResponseResultsInner) GetImageThumbnailUrlOk() (*string, bool)`

GetImageThumbnailUrlOk returns a tuple with the ImageThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageThumbnailUrl

`func (o *ListFunders200ResponseResultsInner) SetImageThumbnailUrl(v string)`

SetImageThumbnailUrl sets ImageThumbnailUrl field to given value.

### HasImageThumbnailUrl

`func (o *ListFunders200ResponseResultsInner) HasImageThumbnailUrl() bool`

HasImageThumbnailUrl returns a boolean if a field has been set.

### GetImageUrl

`func (o *ListFunders200ResponseResultsInner) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *ListFunders200ResponseResultsInner) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *ListFunders200ResponseResultsInner) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *ListFunders200ResponseResultsInner) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetRoles

`func (o *ListFunders200ResponseResultsInner) GetRoles() []ListFunders200ResponseResultsInnerRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ListFunders200ResponseResultsInner) GetRolesOk() (*[]ListFunders200ResponseResultsInnerRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ListFunders200ResponseResultsInner) SetRoles(v []ListFunders200ResponseResultsInnerRolesInner)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ListFunders200ResponseResultsInner) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSummaryStats

`func (o *ListFunders200ResponseResultsInner) GetSummaryStats() ListFunders200ResponseResultsInnerSummaryStats`

GetSummaryStats returns the SummaryStats field if non-nil, zero value otherwise.

### GetSummaryStatsOk

`func (o *ListFunders200ResponseResultsInner) GetSummaryStatsOk() (*ListFunders200ResponseResultsInnerSummaryStats, bool)`

GetSummaryStatsOk returns a tuple with the SummaryStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryStats

`func (o *ListFunders200ResponseResultsInner) SetSummaryStats(v ListFunders200ResponseResultsInnerSummaryStats)`

SetSummaryStats sets SummaryStats field to given value.

### HasSummaryStats

`func (o *ListFunders200ResponseResultsInner) HasSummaryStats() bool`

HasSummaryStats returns a boolean if a field has been set.

### GetUpdatedDate

`func (o *ListFunders200ResponseResultsInner) GetUpdatedDate() time.Time`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *ListFunders200ResponseResultsInner) GetUpdatedDateOk() (*time.Time, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *ListFunders200ResponseResultsInner) SetUpdatedDate(v time.Time)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *ListFunders200ResponseResultsInner) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.

### GetWorksCount

`func (o *ListFunders200ResponseResultsInner) GetWorksCount() int32`

GetWorksCount returns the WorksCount field if non-nil, zero value otherwise.

### GetWorksCountOk

`func (o *ListFunders200ResponseResultsInner) GetWorksCountOk() (*int32, bool)`

GetWorksCountOk returns a tuple with the WorksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksCount

`func (o *ListFunders200ResponseResultsInner) SetWorksCount(v int32)`

SetWorksCount sets WorksCount field to given value.

### HasWorksCount

`func (o *ListFunders200ResponseResultsInner) HasWorksCount() bool`

HasWorksCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


