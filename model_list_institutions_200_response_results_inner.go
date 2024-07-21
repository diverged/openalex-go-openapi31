/*
OpenAlex API

An OpenAPI specification for the OpenAlex API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ListInstitutions200ResponseResultsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListInstitutions200ResponseResultsInner{}

// ListInstitutions200ResponseResultsInner struct for ListInstitutions200ResponseResultsInner
type ListInstitutions200ResponseResultsInner struct {
	// List of institutions related to this one, represented as dehydrated Institution objects with an additional 'relationship' property.
	AssociatedInstitutions []ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner `json:"associated_institutions,omitempty"`
	// The total number of works that cite a work created by an author affiliated with this institution.
	CitedByCount *int32 `json:"cited_by_count,omitempty"`
	// The country where this institution is located, represented as an ISO two-letter country code.
	CountryCode *string `json:"country_code,omitempty"`
	// Works count and cited by count for each of the last ten years, binned by year.
	CountsByYear []ListAuthors200ResponseResultsInnerCountsByYearInner `json:"counts_by_year,omitempty"`
	// The date this Institution object was created in the OpenAlex dataset.
	CreatedDate *string `json:"created_date,omitempty"`
	// The primary name of the institution.
	DisplayName *string `json:"display_name,omitempty"`
	// Acronyms or initialisms that people sometimes use instead of the full display name.
	DisplayNameAcronyms []string `json:"display_name_acronyms,omitempty"`
	// Other names people may use for this institution.
	DisplayNameAlternatives []string `json:"display_name_alternatives,omitempty"`
	Geo *ListInstitutions200ResponseResultsInnerGeo `json:"geo,omitempty"`
	// The URL for institution's primary homepage.
	HomepageUrl *string `json:"homepage_url,omitempty"`
	// The OpenAlex ID for this institution.
	Id *string `json:"id,omitempty"`
	Ids *ListInstitutions200ResponseResultsInnerIds `json:"ids,omitempty"`
	// URL for a thumbnail image representing this institution.
	ImageThumbnailUrl *string `json:"image_thumbnail_url,omitempty"`
	// URL for an image representing this institution.
	ImageUrl *string `json:"image_url,omitempty"`
	International *ListInstitutions200ResponseResultsInnerInternational `json:"international,omitempty"`
	// OpenAlex IDs of this institution and its parent institutions.
	Lineage []string `json:"lineage,omitempty"`
	// Repositories that have this institution as their host organization.
	Repositories []ListInstitutions200ResponseResultsInnerRepositoriesInner `json:"repositories,omitempty"`
	// List of roles this institution plays, including itself and potentially funder and publisher roles.
	Roles []ListFunders200ResponseResultsInnerRolesInner `json:"roles,omitempty"`
	// The ROR ID for this institution.
	Ror *string `json:"ror,omitempty"`
	SummaryStats *ListInstitutions200ResponseResultsInnerSummaryStats `json:"summary_stats,omitempty"`
	// Topics that are frequently associated with works affiliated with this institution, in descending order of count.
	Topics []ListInstitutions200ResponseResultsInnerTopicsInner `json:"topics,omitempty"`
	// Topics that are frequently associated with works affiliated with this institution, in descending order of value.
	TopicShare []ListInstitutions200ResponseResultsInnerTopicsInner `json:"topic_share,omitempty"`
	// The institution's primary type, using the ROR \"type\" controlled vocabulary.
	Type *string `json:"type,omitempty"`
	// The last time anything in this institution object changed. Formatted as ISO 8601 extended format without time zone designator.
	UpdatedDate *string `json:"updated_date,omitempty"`
	// A URL that will get you a list of all the Works affiliated with this institution.
	WorksApiUrl *string `json:"works_api_url,omitempty"`
	// The number of Works created by authors affiliated with this institution.
	WorksCount *int32 `json:"works_count,omitempty"`
	// The Concepts most frequently applied to works affiliated with this institution.
	XConcepts []map[string]interface{} `json:"x_concepts,omitempty"`
}

// NewListInstitutions200ResponseResultsInner instantiates a new ListInstitutions200ResponseResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInstitutions200ResponseResultsInner() *ListInstitutions200ResponseResultsInner {
	this := ListInstitutions200ResponseResultsInner{}
	return &this
}

// NewListInstitutions200ResponseResultsInnerWithDefaults instantiates a new ListInstitutions200ResponseResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInstitutions200ResponseResultsInnerWithDefaults() *ListInstitutions200ResponseResultsInner {
	this := ListInstitutions200ResponseResultsInner{}
	return &this
}

// GetAssociatedInstitutions returns the AssociatedInstitutions field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetAssociatedInstitutions() []ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner {
	if o == nil || IsNil(o.AssociatedInstitutions) {
		var ret []ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner
		return ret
	}
	return o.AssociatedInstitutions
}

// GetAssociatedInstitutionsOk returns a tuple with the AssociatedInstitutions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetAssociatedInstitutionsOk() ([]ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner, bool) {
	if o == nil || IsNil(o.AssociatedInstitutions) {
		return nil, false
	}
	return o.AssociatedInstitutions, true
}

// HasAssociatedInstitutions returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasAssociatedInstitutions() bool {
	if o != nil && !IsNil(o.AssociatedInstitutions) {
		return true
	}

	return false
}

// SetAssociatedInstitutions gets a reference to the given []ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner and assigns it to the AssociatedInstitutions field.
func (o *ListInstitutions200ResponseResultsInner) SetAssociatedInstitutions(v []ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) {
	o.AssociatedInstitutions = v
}

// GetCitedByCount returns the CitedByCount field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetCitedByCount() int32 {
	if o == nil || IsNil(o.CitedByCount) {
		var ret int32
		return ret
	}
	return *o.CitedByCount
}

// GetCitedByCountOk returns a tuple with the CitedByCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetCitedByCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CitedByCount) {
		return nil, false
	}
	return o.CitedByCount, true
}

// HasCitedByCount returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasCitedByCount() bool {
	if o != nil && !IsNil(o.CitedByCount) {
		return true
	}

	return false
}

// SetCitedByCount gets a reference to the given int32 and assigns it to the CitedByCount field.
func (o *ListInstitutions200ResponseResultsInner) SetCitedByCount(v int32) {
	o.CitedByCount = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *ListInstitutions200ResponseResultsInner) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetCountsByYear returns the CountsByYear field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetCountsByYear() []ListAuthors200ResponseResultsInnerCountsByYearInner {
	if o == nil || IsNil(o.CountsByYear) {
		var ret []ListAuthors200ResponseResultsInnerCountsByYearInner
		return ret
	}
	return o.CountsByYear
}

// GetCountsByYearOk returns a tuple with the CountsByYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetCountsByYearOk() ([]ListAuthors200ResponseResultsInnerCountsByYearInner, bool) {
	if o == nil || IsNil(o.CountsByYear) {
		return nil, false
	}
	return o.CountsByYear, true
}

// HasCountsByYear returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasCountsByYear() bool {
	if o != nil && !IsNil(o.CountsByYear) {
		return true
	}

	return false
}

// SetCountsByYear gets a reference to the given []ListAuthors200ResponseResultsInnerCountsByYearInner and assigns it to the CountsByYear field.
func (o *ListInstitutions200ResponseResultsInner) SetCountsByYear(v []ListAuthors200ResponseResultsInnerCountsByYearInner) {
	o.CountsByYear = v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetCreatedDate() string {
	if o == nil || IsNil(o.CreatedDate) {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetCreatedDateOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *ListInstitutions200ResponseResultsInner) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ListInstitutions200ResponseResultsInner) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDisplayNameAcronyms returns the DisplayNameAcronyms field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetDisplayNameAcronyms() []string {
	if o == nil || IsNil(o.DisplayNameAcronyms) {
		var ret []string
		return ret
	}
	return o.DisplayNameAcronyms
}

// GetDisplayNameAcronymsOk returns a tuple with the DisplayNameAcronyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetDisplayNameAcronymsOk() ([]string, bool) {
	if o == nil || IsNil(o.DisplayNameAcronyms) {
		return nil, false
	}
	return o.DisplayNameAcronyms, true
}

// HasDisplayNameAcronyms returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasDisplayNameAcronyms() bool {
	if o != nil && !IsNil(o.DisplayNameAcronyms) {
		return true
	}

	return false
}

// SetDisplayNameAcronyms gets a reference to the given []string and assigns it to the DisplayNameAcronyms field.
func (o *ListInstitutions200ResponseResultsInner) SetDisplayNameAcronyms(v []string) {
	o.DisplayNameAcronyms = v
}

// GetDisplayNameAlternatives returns the DisplayNameAlternatives field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetDisplayNameAlternatives() []string {
	if o == nil || IsNil(o.DisplayNameAlternatives) {
		var ret []string
		return ret
	}
	return o.DisplayNameAlternatives
}

// GetDisplayNameAlternativesOk returns a tuple with the DisplayNameAlternatives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetDisplayNameAlternativesOk() ([]string, bool) {
	if o == nil || IsNil(o.DisplayNameAlternatives) {
		return nil, false
	}
	return o.DisplayNameAlternatives, true
}

// HasDisplayNameAlternatives returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasDisplayNameAlternatives() bool {
	if o != nil && !IsNil(o.DisplayNameAlternatives) {
		return true
	}

	return false
}

// SetDisplayNameAlternatives gets a reference to the given []string and assigns it to the DisplayNameAlternatives field.
func (o *ListInstitutions200ResponseResultsInner) SetDisplayNameAlternatives(v []string) {
	o.DisplayNameAlternatives = v
}

// GetGeo returns the Geo field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetGeo() ListInstitutions200ResponseResultsInnerGeo {
	if o == nil || IsNil(o.Geo) {
		var ret ListInstitutions200ResponseResultsInnerGeo
		return ret
	}
	return *o.Geo
}

// GetGeoOk returns a tuple with the Geo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetGeoOk() (*ListInstitutions200ResponseResultsInnerGeo, bool) {
	if o == nil || IsNil(o.Geo) {
		return nil, false
	}
	return o.Geo, true
}

// HasGeo returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasGeo() bool {
	if o != nil && !IsNil(o.Geo) {
		return true
	}

	return false
}

// SetGeo gets a reference to the given ListInstitutions200ResponseResultsInnerGeo and assigns it to the Geo field.
func (o *ListInstitutions200ResponseResultsInner) SetGeo(v ListInstitutions200ResponseResultsInnerGeo) {
	o.Geo = &v
}

// GetHomepageUrl returns the HomepageUrl field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetHomepageUrl() string {
	if o == nil || IsNil(o.HomepageUrl) {
		var ret string
		return ret
	}
	return *o.HomepageUrl
}

// GetHomepageUrlOk returns a tuple with the HomepageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetHomepageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HomepageUrl) {
		return nil, false
	}
	return o.HomepageUrl, true
}

// HasHomepageUrl returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasHomepageUrl() bool {
	if o != nil && !IsNil(o.HomepageUrl) {
		return true
	}

	return false
}

// SetHomepageUrl gets a reference to the given string and assigns it to the HomepageUrl field.
func (o *ListInstitutions200ResponseResultsInner) SetHomepageUrl(v string) {
	o.HomepageUrl = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListInstitutions200ResponseResultsInner) SetId(v string) {
	o.Id = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetIds() ListInstitutions200ResponseResultsInnerIds {
	if o == nil || IsNil(o.Ids) {
		var ret ListInstitutions200ResponseResultsInnerIds
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetIdsOk() (*ListInstitutions200ResponseResultsInnerIds, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given ListInstitutions200ResponseResultsInnerIds and assigns it to the Ids field.
func (o *ListInstitutions200ResponseResultsInner) SetIds(v ListInstitutions200ResponseResultsInnerIds) {
	o.Ids = &v
}

// GetImageThumbnailUrl returns the ImageThumbnailUrl field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetImageThumbnailUrl() string {
	if o == nil || IsNil(o.ImageThumbnailUrl) {
		var ret string
		return ret
	}
	return *o.ImageThumbnailUrl
}

// GetImageThumbnailUrlOk returns a tuple with the ImageThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetImageThumbnailUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageThumbnailUrl) {
		return nil, false
	}
	return o.ImageThumbnailUrl, true
}

// HasImageThumbnailUrl returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasImageThumbnailUrl() bool {
	if o != nil && !IsNil(o.ImageThumbnailUrl) {
		return true
	}

	return false
}

// SetImageThumbnailUrl gets a reference to the given string and assigns it to the ImageThumbnailUrl field.
func (o *ListInstitutions200ResponseResultsInner) SetImageThumbnailUrl(v string) {
	o.ImageThumbnailUrl = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *ListInstitutions200ResponseResultsInner) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetInternational returns the International field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetInternational() ListInstitutions200ResponseResultsInnerInternational {
	if o == nil || IsNil(o.International) {
		var ret ListInstitutions200ResponseResultsInnerInternational
		return ret
	}
	return *o.International
}

// GetInternationalOk returns a tuple with the International field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetInternationalOk() (*ListInstitutions200ResponseResultsInnerInternational, bool) {
	if o == nil || IsNil(o.International) {
		return nil, false
	}
	return o.International, true
}

// HasInternational returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasInternational() bool {
	if o != nil && !IsNil(o.International) {
		return true
	}

	return false
}

// SetInternational gets a reference to the given ListInstitutions200ResponseResultsInnerInternational and assigns it to the International field.
func (o *ListInstitutions200ResponseResultsInner) SetInternational(v ListInstitutions200ResponseResultsInnerInternational) {
	o.International = &v
}

// GetLineage returns the Lineage field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetLineage() []string {
	if o == nil || IsNil(o.Lineage) {
		var ret []string
		return ret
	}
	return o.Lineage
}

// GetLineageOk returns a tuple with the Lineage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetLineageOk() ([]string, bool) {
	if o == nil || IsNil(o.Lineage) {
		return nil, false
	}
	return o.Lineage, true
}

// HasLineage returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasLineage() bool {
	if o != nil && !IsNil(o.Lineage) {
		return true
	}

	return false
}

// SetLineage gets a reference to the given []string and assigns it to the Lineage field.
func (o *ListInstitutions200ResponseResultsInner) SetLineage(v []string) {
	o.Lineage = v
}

// GetRepositories returns the Repositories field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetRepositories() []ListInstitutions200ResponseResultsInnerRepositoriesInner {
	if o == nil || IsNil(o.Repositories) {
		var ret []ListInstitutions200ResponseResultsInnerRepositoriesInner
		return ret
	}
	return o.Repositories
}

// GetRepositoriesOk returns a tuple with the Repositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetRepositoriesOk() ([]ListInstitutions200ResponseResultsInnerRepositoriesInner, bool) {
	if o == nil || IsNil(o.Repositories) {
		return nil, false
	}
	return o.Repositories, true
}

// HasRepositories returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasRepositories() bool {
	if o != nil && !IsNil(o.Repositories) {
		return true
	}

	return false
}

// SetRepositories gets a reference to the given []ListInstitutions200ResponseResultsInnerRepositoriesInner and assigns it to the Repositories field.
func (o *ListInstitutions200ResponseResultsInner) SetRepositories(v []ListInstitutions200ResponseResultsInnerRepositoriesInner) {
	o.Repositories = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetRoles() []ListFunders200ResponseResultsInnerRolesInner {
	if o == nil || IsNil(o.Roles) {
		var ret []ListFunders200ResponseResultsInnerRolesInner
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetRolesOk() ([]ListFunders200ResponseResultsInnerRolesInner, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []ListFunders200ResponseResultsInnerRolesInner and assigns it to the Roles field.
func (o *ListInstitutions200ResponseResultsInner) SetRoles(v []ListFunders200ResponseResultsInnerRolesInner) {
	o.Roles = v
}

// GetRor returns the Ror field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetRor() string {
	if o == nil || IsNil(o.Ror) {
		var ret string
		return ret
	}
	return *o.Ror
}

// GetRorOk returns a tuple with the Ror field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetRorOk() (*string, bool) {
	if o == nil || IsNil(o.Ror) {
		return nil, false
	}
	return o.Ror, true
}

// HasRor returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasRor() bool {
	if o != nil && !IsNil(o.Ror) {
		return true
	}

	return false
}

// SetRor gets a reference to the given string and assigns it to the Ror field.
func (o *ListInstitutions200ResponseResultsInner) SetRor(v string) {
	o.Ror = &v
}

// GetSummaryStats returns the SummaryStats field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetSummaryStats() ListInstitutions200ResponseResultsInnerSummaryStats {
	if o == nil || IsNil(o.SummaryStats) {
		var ret ListInstitutions200ResponseResultsInnerSummaryStats
		return ret
	}
	return *o.SummaryStats
}

// GetSummaryStatsOk returns a tuple with the SummaryStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetSummaryStatsOk() (*ListInstitutions200ResponseResultsInnerSummaryStats, bool) {
	if o == nil || IsNil(o.SummaryStats) {
		return nil, false
	}
	return o.SummaryStats, true
}

// HasSummaryStats returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasSummaryStats() bool {
	if o != nil && !IsNil(o.SummaryStats) {
		return true
	}

	return false
}

// SetSummaryStats gets a reference to the given ListInstitutions200ResponseResultsInnerSummaryStats and assigns it to the SummaryStats field.
func (o *ListInstitutions200ResponseResultsInner) SetSummaryStats(v ListInstitutions200ResponseResultsInnerSummaryStats) {
	o.SummaryStats = &v
}

// GetTopics returns the Topics field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetTopics() []ListInstitutions200ResponseResultsInnerTopicsInner {
	if o == nil || IsNil(o.Topics) {
		var ret []ListInstitutions200ResponseResultsInnerTopicsInner
		return ret
	}
	return o.Topics
}

// GetTopicsOk returns a tuple with the Topics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetTopicsOk() ([]ListInstitutions200ResponseResultsInnerTopicsInner, bool) {
	if o == nil || IsNil(o.Topics) {
		return nil, false
	}
	return o.Topics, true
}

// HasTopics returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasTopics() bool {
	if o != nil && !IsNil(o.Topics) {
		return true
	}

	return false
}

// SetTopics gets a reference to the given []ListInstitutions200ResponseResultsInnerTopicsInner and assigns it to the Topics field.
func (o *ListInstitutions200ResponseResultsInner) SetTopics(v []ListInstitutions200ResponseResultsInnerTopicsInner) {
	o.Topics = v
}

// GetTopicShare returns the TopicShare field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetTopicShare() []ListInstitutions200ResponseResultsInnerTopicsInner {
	if o == nil || IsNil(o.TopicShare) {
		var ret []ListInstitutions200ResponseResultsInnerTopicsInner
		return ret
	}
	return o.TopicShare
}

// GetTopicShareOk returns a tuple with the TopicShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetTopicShareOk() ([]ListInstitutions200ResponseResultsInnerTopicsInner, bool) {
	if o == nil || IsNil(o.TopicShare) {
		return nil, false
	}
	return o.TopicShare, true
}

// HasTopicShare returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasTopicShare() bool {
	if o != nil && !IsNil(o.TopicShare) {
		return true
	}

	return false
}

// SetTopicShare gets a reference to the given []ListInstitutions200ResponseResultsInnerTopicsInner and assigns it to the TopicShare field.
func (o *ListInstitutions200ResponseResultsInner) SetTopicShare(v []ListInstitutions200ResponseResultsInnerTopicsInner) {
	o.TopicShare = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ListInstitutions200ResponseResultsInner) SetType(v string) {
	o.Type = &v
}

// GetUpdatedDate returns the UpdatedDate field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetUpdatedDate() string {
	if o == nil || IsNil(o.UpdatedDate) {
		var ret string
		return ret
	}
	return *o.UpdatedDate
}

// GetUpdatedDateOk returns a tuple with the UpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetUpdatedDateOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedDate) {
		return nil, false
	}
	return o.UpdatedDate, true
}

// HasUpdatedDate returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasUpdatedDate() bool {
	if o != nil && !IsNil(o.UpdatedDate) {
		return true
	}

	return false
}

// SetUpdatedDate gets a reference to the given string and assigns it to the UpdatedDate field.
func (o *ListInstitutions200ResponseResultsInner) SetUpdatedDate(v string) {
	o.UpdatedDate = &v
}

// GetWorksApiUrl returns the WorksApiUrl field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetWorksApiUrl() string {
	if o == nil || IsNil(o.WorksApiUrl) {
		var ret string
		return ret
	}
	return *o.WorksApiUrl
}

// GetWorksApiUrlOk returns a tuple with the WorksApiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetWorksApiUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WorksApiUrl) {
		return nil, false
	}
	return o.WorksApiUrl, true
}

// HasWorksApiUrl returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasWorksApiUrl() bool {
	if o != nil && !IsNil(o.WorksApiUrl) {
		return true
	}

	return false
}

// SetWorksApiUrl gets a reference to the given string and assigns it to the WorksApiUrl field.
func (o *ListInstitutions200ResponseResultsInner) SetWorksApiUrl(v string) {
	o.WorksApiUrl = &v
}

// GetWorksCount returns the WorksCount field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetWorksCount() int32 {
	if o == nil || IsNil(o.WorksCount) {
		var ret int32
		return ret
	}
	return *o.WorksCount
}

// GetWorksCountOk returns a tuple with the WorksCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetWorksCountOk() (*int32, bool) {
	if o == nil || IsNil(o.WorksCount) {
		return nil, false
	}
	return o.WorksCount, true
}

// HasWorksCount returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasWorksCount() bool {
	if o != nil && !IsNil(o.WorksCount) {
		return true
	}

	return false
}

// SetWorksCount gets a reference to the given int32 and assigns it to the WorksCount field.
func (o *ListInstitutions200ResponseResultsInner) SetWorksCount(v int32) {
	o.WorksCount = &v
}

// GetXConcepts returns the XConcepts field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInner) GetXConcepts() []map[string]interface{} {
	if o == nil || IsNil(o.XConcepts) {
		var ret []map[string]interface{}
		return ret
	}
	return o.XConcepts
}

// GetXConceptsOk returns a tuple with the XConcepts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInner) GetXConceptsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.XConcepts) {
		return nil, false
	}
	return o.XConcepts, true
}

// HasXConcepts returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInner) HasXConcepts() bool {
	if o != nil && !IsNil(o.XConcepts) {
		return true
	}

	return false
}

// SetXConcepts gets a reference to the given []map[string]interface{} and assigns it to the XConcepts field.
func (o *ListInstitutions200ResponseResultsInner) SetXConcepts(v []map[string]interface{}) {
	o.XConcepts = v
}

func (o ListInstitutions200ResponseResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListInstitutions200ResponseResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssociatedInstitutions) {
		toSerialize["associated_institutions"] = o.AssociatedInstitutions
	}
	if !IsNil(o.CitedByCount) {
		toSerialize["cited_by_count"] = o.CitedByCount
	}
	if !IsNil(o.CountryCode) {
		toSerialize["country_code"] = o.CountryCode
	}
	if !IsNil(o.CountsByYear) {
		toSerialize["counts_by_year"] = o.CountsByYear
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["created_date"] = o.CreatedDate
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.DisplayNameAcronyms) {
		toSerialize["display_name_acronyms"] = o.DisplayNameAcronyms
	}
	if !IsNil(o.DisplayNameAlternatives) {
		toSerialize["display_name_alternatives"] = o.DisplayNameAlternatives
	}
	if !IsNil(o.Geo) {
		toSerialize["geo"] = o.Geo
	}
	if !IsNil(o.HomepageUrl) {
		toSerialize["homepage_url"] = o.HomepageUrl
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !IsNil(o.ImageThumbnailUrl) {
		toSerialize["image_thumbnail_url"] = o.ImageThumbnailUrl
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["image_url"] = o.ImageUrl
	}
	if !IsNil(o.International) {
		toSerialize["international"] = o.International
	}
	if !IsNil(o.Lineage) {
		toSerialize["lineage"] = o.Lineage
	}
	if !IsNil(o.Repositories) {
		toSerialize["repositories"] = o.Repositories
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Ror) {
		toSerialize["ror"] = o.Ror
	}
	if !IsNil(o.SummaryStats) {
		toSerialize["summary_stats"] = o.SummaryStats
	}
	if !IsNil(o.Topics) {
		toSerialize["topics"] = o.Topics
	}
	if !IsNil(o.TopicShare) {
		toSerialize["topic_share"] = o.TopicShare
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UpdatedDate) {
		toSerialize["updated_date"] = o.UpdatedDate
	}
	if !IsNil(o.WorksApiUrl) {
		toSerialize["works_api_url"] = o.WorksApiUrl
	}
	if !IsNil(o.WorksCount) {
		toSerialize["works_count"] = o.WorksCount
	}
	if !IsNil(o.XConcepts) {
		toSerialize["x_concepts"] = o.XConcepts
	}
	return toSerialize, nil
}

type NullableListInstitutions200ResponseResultsInner struct {
	value *ListInstitutions200ResponseResultsInner
	isSet bool
}

func (v NullableListInstitutions200ResponseResultsInner) Get() *ListInstitutions200ResponseResultsInner {
	return v.value
}

func (v *NullableListInstitutions200ResponseResultsInner) Set(val *ListInstitutions200ResponseResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListInstitutions200ResponseResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListInstitutions200ResponseResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInstitutions200ResponseResultsInner(val *ListInstitutions200ResponseResultsInner) *NullableListInstitutions200ResponseResultsInner {
	return &NullableListInstitutions200ResponseResultsInner{value: val, isSet: true}
}

func (v NullableListInstitutions200ResponseResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInstitutions200ResponseResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


