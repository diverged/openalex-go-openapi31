# ListPublishers200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternateTitles** | Pointer to **[]string** | A list of alternate titles for this publisher. | [optional] 
**CitedByCount** | Pointer to **int32** | The number of citations to works that are linked to this publisher through journals or other sources. | [optional] 
**CountryCodes** | Pointer to **[]string** | The countries where the publisher is primarily located, as an ISO two-letter country code. | [optional] 
**CountsByYear** | Pointer to [**[]ListAuthors200ResponseResultsInnerCountsByYearInner**](ListAuthors200ResponseResultsInnerCountsByYearInner.md) | The values of works_count and cited_by_count for each of the last ten years, binned by year. | [optional] 
**CreatedDate** | Pointer to **string** | The date this Publisher object was created in the OpenAlex dataset, expressed as an ISO 8601 date string. | [optional] 
**DisplayName** | Pointer to **string** | The primary name of the publisher. | [optional] 
**HierarchyLevel** | Pointer to **int32** | The hierarchy level for this publisher. A publisher with hierarchy level 0 has no parent publishers. | [optional] 
**Id** | Pointer to **string** | The OpenAlex ID for this publisher. | [optional] 
**Ids** | Pointer to [**ListPublishers200ResponseResultsInnerIds**](ListPublishers200ResponseResultsInnerIds.md) |  | [optional] 
**ImageThumbnailUrl** | Pointer to **string** | URL for a smaller version of the image representing this publisher. | [optional] 
**ImageUrl** | Pointer to **string** | URL where you can get an image representing this publisher. | [optional] 
**Lineage** | Pointer to **[]string** | OpenAlex IDs of publishers, including this publisher&#39;s ID and any parent publishers. | [optional] 
**ParentPublisher** | Pointer to [**NullableListPublishers200ResponseResultsInnerParentPublisher**](ListPublishers200ResponseResultsInnerParentPublisher.md) |  | [optional] 
**Roles** | Pointer to [**[]ListFunders200ResponseResultsInnerRolesInner**](ListFunders200ResponseResultsInnerRolesInner.md) | List of role objects, which include the role, the id, and the works_count. | [optional] 
**SourcesApiUrl** | Pointer to **string** | An URL that will get you a list of all the sources published by this publisher. | [optional] 
**SummaryStats** | Pointer to [**ListPublishers200ResponseResultsInnerSummaryStats**](ListPublishers200ResponseResultsInnerSummaryStats.md) |  | [optional] 
**UpdatedDate** | Pointer to **string** | The last time anything in this publisher object changed. Formatted as ISO 8601 extended format without time zone designator. | [optional] 
**WorksCount** | Pointer to **int32** | The number of works published by this publisher. | [optional] 

## Methods

### NewListPublishers200ResponseResultsInner

`func NewListPublishers200ResponseResultsInner() *ListPublishers200ResponseResultsInner`

NewListPublishers200ResponseResultsInner instantiates a new ListPublishers200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPublishers200ResponseResultsInnerWithDefaults

`func NewListPublishers200ResponseResultsInnerWithDefaults() *ListPublishers200ResponseResultsInner`

NewListPublishers200ResponseResultsInnerWithDefaults instantiates a new ListPublishers200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateTitles

`func (o *ListPublishers200ResponseResultsInner) GetAlternateTitles() []string`

GetAlternateTitles returns the AlternateTitles field if non-nil, zero value otherwise.

### GetAlternateTitlesOk

`func (o *ListPublishers200ResponseResultsInner) GetAlternateTitlesOk() (*[]string, bool)`

GetAlternateTitlesOk returns a tuple with the AlternateTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateTitles

`func (o *ListPublishers200ResponseResultsInner) SetAlternateTitles(v []string)`

SetAlternateTitles sets AlternateTitles field to given value.

### HasAlternateTitles

`func (o *ListPublishers200ResponseResultsInner) HasAlternateTitles() bool`

HasAlternateTitles returns a boolean if a field has been set.

### GetCitedByCount

`func (o *ListPublishers200ResponseResultsInner) GetCitedByCount() int32`

GetCitedByCount returns the CitedByCount field if non-nil, zero value otherwise.

### GetCitedByCountOk

`func (o *ListPublishers200ResponseResultsInner) GetCitedByCountOk() (*int32, bool)`

GetCitedByCountOk returns a tuple with the CitedByCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitedByCount

`func (o *ListPublishers200ResponseResultsInner) SetCitedByCount(v int32)`

SetCitedByCount sets CitedByCount field to given value.

### HasCitedByCount

`func (o *ListPublishers200ResponseResultsInner) HasCitedByCount() bool`

HasCitedByCount returns a boolean if a field has been set.

### GetCountryCodes

`func (o *ListPublishers200ResponseResultsInner) GetCountryCodes() []string`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *ListPublishers200ResponseResultsInner) GetCountryCodesOk() (*[]string, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *ListPublishers200ResponseResultsInner) SetCountryCodes(v []string)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *ListPublishers200ResponseResultsInner) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.

### GetCountsByYear

`func (o *ListPublishers200ResponseResultsInner) GetCountsByYear() []ListAuthors200ResponseResultsInnerCountsByYearInner`

GetCountsByYear returns the CountsByYear field if non-nil, zero value otherwise.

### GetCountsByYearOk

`func (o *ListPublishers200ResponseResultsInner) GetCountsByYearOk() (*[]ListAuthors200ResponseResultsInnerCountsByYearInner, bool)`

GetCountsByYearOk returns a tuple with the CountsByYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountsByYear

`func (o *ListPublishers200ResponseResultsInner) SetCountsByYear(v []ListAuthors200ResponseResultsInnerCountsByYearInner)`

SetCountsByYear sets CountsByYear field to given value.

### HasCountsByYear

`func (o *ListPublishers200ResponseResultsInner) HasCountsByYear() bool`

HasCountsByYear returns a boolean if a field has been set.

### GetCreatedDate

`func (o *ListPublishers200ResponseResultsInner) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ListPublishers200ResponseResultsInner) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ListPublishers200ResponseResultsInner) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ListPublishers200ResponseResultsInner) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetDisplayName

`func (o *ListPublishers200ResponseResultsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListPublishers200ResponseResultsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListPublishers200ResponseResultsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ListPublishers200ResponseResultsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHierarchyLevel

`func (o *ListPublishers200ResponseResultsInner) GetHierarchyLevel() int32`

GetHierarchyLevel returns the HierarchyLevel field if non-nil, zero value otherwise.

### GetHierarchyLevelOk

`func (o *ListPublishers200ResponseResultsInner) GetHierarchyLevelOk() (*int32, bool)`

GetHierarchyLevelOk returns a tuple with the HierarchyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyLevel

`func (o *ListPublishers200ResponseResultsInner) SetHierarchyLevel(v int32)`

SetHierarchyLevel sets HierarchyLevel field to given value.

### HasHierarchyLevel

`func (o *ListPublishers200ResponseResultsInner) HasHierarchyLevel() bool`

HasHierarchyLevel returns a boolean if a field has been set.

### GetId

`func (o *ListPublishers200ResponseResultsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListPublishers200ResponseResultsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListPublishers200ResponseResultsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListPublishers200ResponseResultsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIds

`func (o *ListPublishers200ResponseResultsInner) GetIds() ListPublishers200ResponseResultsInnerIds`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListPublishers200ResponseResultsInner) GetIdsOk() (*ListPublishers200ResponseResultsInnerIds, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListPublishers200ResponseResultsInner) SetIds(v ListPublishers200ResponseResultsInnerIds)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ListPublishers200ResponseResultsInner) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetImageThumbnailUrl

`func (o *ListPublishers200ResponseResultsInner) GetImageThumbnailUrl() string`

GetImageThumbnailUrl returns the ImageThumbnailUrl field if non-nil, zero value otherwise.

### GetImageThumbnailUrlOk

`func (o *ListPublishers200ResponseResultsInner) GetImageThumbnailUrlOk() (*string, bool)`

GetImageThumbnailUrlOk returns a tuple with the ImageThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageThumbnailUrl

`func (o *ListPublishers200ResponseResultsInner) SetImageThumbnailUrl(v string)`

SetImageThumbnailUrl sets ImageThumbnailUrl field to given value.

### HasImageThumbnailUrl

`func (o *ListPublishers200ResponseResultsInner) HasImageThumbnailUrl() bool`

HasImageThumbnailUrl returns a boolean if a field has been set.

### GetImageUrl

`func (o *ListPublishers200ResponseResultsInner) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *ListPublishers200ResponseResultsInner) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *ListPublishers200ResponseResultsInner) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *ListPublishers200ResponseResultsInner) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetLineage

`func (o *ListPublishers200ResponseResultsInner) GetLineage() []string`

GetLineage returns the Lineage field if non-nil, zero value otherwise.

### GetLineageOk

`func (o *ListPublishers200ResponseResultsInner) GetLineageOk() (*[]string, bool)`

GetLineageOk returns a tuple with the Lineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineage

`func (o *ListPublishers200ResponseResultsInner) SetLineage(v []string)`

SetLineage sets Lineage field to given value.

### HasLineage

`func (o *ListPublishers200ResponseResultsInner) HasLineage() bool`

HasLineage returns a boolean if a field has been set.

### GetParentPublisher

`func (o *ListPublishers200ResponseResultsInner) GetParentPublisher() ListPublishers200ResponseResultsInnerParentPublisher`

GetParentPublisher returns the ParentPublisher field if non-nil, zero value otherwise.

### GetParentPublisherOk

`func (o *ListPublishers200ResponseResultsInner) GetParentPublisherOk() (*ListPublishers200ResponseResultsInnerParentPublisher, bool)`

GetParentPublisherOk returns a tuple with the ParentPublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPublisher

`func (o *ListPublishers200ResponseResultsInner) SetParentPublisher(v ListPublishers200ResponseResultsInnerParentPublisher)`

SetParentPublisher sets ParentPublisher field to given value.

### HasParentPublisher

`func (o *ListPublishers200ResponseResultsInner) HasParentPublisher() bool`

HasParentPublisher returns a boolean if a field has been set.

### SetParentPublisherNil

`func (o *ListPublishers200ResponseResultsInner) SetParentPublisherNil(b bool)`

 SetParentPublisherNil sets the value for ParentPublisher to be an explicit nil

### UnsetParentPublisher
`func (o *ListPublishers200ResponseResultsInner) UnsetParentPublisher()`

UnsetParentPublisher ensures that no value is present for ParentPublisher, not even an explicit nil
### GetRoles

`func (o *ListPublishers200ResponseResultsInner) GetRoles() []ListFunders200ResponseResultsInnerRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ListPublishers200ResponseResultsInner) GetRolesOk() (*[]ListFunders200ResponseResultsInnerRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ListPublishers200ResponseResultsInner) SetRoles(v []ListFunders200ResponseResultsInnerRolesInner)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ListPublishers200ResponseResultsInner) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSourcesApiUrl

`func (o *ListPublishers200ResponseResultsInner) GetSourcesApiUrl() string`

GetSourcesApiUrl returns the SourcesApiUrl field if non-nil, zero value otherwise.

### GetSourcesApiUrlOk

`func (o *ListPublishers200ResponseResultsInner) GetSourcesApiUrlOk() (*string, bool)`

GetSourcesApiUrlOk returns a tuple with the SourcesApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcesApiUrl

`func (o *ListPublishers200ResponseResultsInner) SetSourcesApiUrl(v string)`

SetSourcesApiUrl sets SourcesApiUrl field to given value.

### HasSourcesApiUrl

`func (o *ListPublishers200ResponseResultsInner) HasSourcesApiUrl() bool`

HasSourcesApiUrl returns a boolean if a field has been set.

### GetSummaryStats

`func (o *ListPublishers200ResponseResultsInner) GetSummaryStats() ListPublishers200ResponseResultsInnerSummaryStats`

GetSummaryStats returns the SummaryStats field if non-nil, zero value otherwise.

### GetSummaryStatsOk

`func (o *ListPublishers200ResponseResultsInner) GetSummaryStatsOk() (*ListPublishers200ResponseResultsInnerSummaryStats, bool)`

GetSummaryStatsOk returns a tuple with the SummaryStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryStats

`func (o *ListPublishers200ResponseResultsInner) SetSummaryStats(v ListPublishers200ResponseResultsInnerSummaryStats)`

SetSummaryStats sets SummaryStats field to given value.

### HasSummaryStats

`func (o *ListPublishers200ResponseResultsInner) HasSummaryStats() bool`

HasSummaryStats returns a boolean if a field has been set.

### GetUpdatedDate

`func (o *ListPublishers200ResponseResultsInner) GetUpdatedDate() string`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *ListPublishers200ResponseResultsInner) GetUpdatedDateOk() (*string, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *ListPublishers200ResponseResultsInner) SetUpdatedDate(v string)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *ListPublishers200ResponseResultsInner) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.

### GetWorksCount

`func (o *ListPublishers200ResponseResultsInner) GetWorksCount() int32`

GetWorksCount returns the WorksCount field if non-nil, zero value otherwise.

### GetWorksCountOk

`func (o *ListPublishers200ResponseResultsInner) GetWorksCountOk() (*int32, bool)`

GetWorksCountOk returns a tuple with the WorksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksCount

`func (o *ListPublishers200ResponseResultsInner) SetWorksCount(v int32)`

SetWorksCount sets WorksCount field to given value.

### HasWorksCount

`func (o *ListPublishers200ResponseResultsInner) HasWorksCount() bool`

HasWorksCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


