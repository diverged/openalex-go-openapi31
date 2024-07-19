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

// checks if the ListConcepts200ResponseResultsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListConcepts200ResponseResultsInner{}

// ListConcepts200ResponseResultsInner struct for ListConcepts200ResponseResultsInner
type ListConcepts200ResponseResultsInner struct {
	// The OpenAlex ID for this concept.
	Id *string `json:"id,omitempty"`
	// The Wikidata ID for this concept. This is the Canonical External ID for concepts.
	Wikidata *string `json:"wikidata,omitempty"`
	// The English-language label of the concept.
	DisplayName *string `json:"display_name,omitempty"`
	// The level in the concept tree where this concept lives. Lower-level concepts are more general, and higher-level concepts are more specific.
	Level *int32 `json:"level,omitempty"`
	// A brief description of this concept.
	Description *string `json:"description,omitempty"`
	// The number of works tagged with this concept.
	WorksCount *int32 `json:"works_count,omitempty"`
	// The number citations to works that have been tagged with this concept.
	CitedByCount *int32 `json:"cited_by_count,omitempty"`
	Ids *ListConcepts200ResponseResultsInnerIds `json:"ids,omitempty"`
	// URL where you can get an image representing this concept, where available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Same as image_url, but it's a smaller image.
	ImageThumbnailUrl *string `json:"image_thumbnail_url,omitempty"`
	International *ListConcepts200ResponseResultsInnerInternational `json:"international,omitempty"`
	// List of concepts that this concept descends from, as dehydrated Concept objects.
	Ancestors []ListConcepts200ResponseResultsInnerAncestorsInner `json:"ancestors,omitempty"`
	RelatedConcepts []map[string]interface{} `json:"related_concepts,omitempty"`
	// The values of works_count and cited_by_count for each of the last ten years, binned by year.
	CountsByYear []ListAuthors200ResponseResultsInnerCountsByYearInner `json:"counts_by_year,omitempty"`
	// An URL that will get you a list of all the works tagged with this concept.
	WorksApiUrl *string `json:"works_api_url,omitempty"`
	// The last time anything in this concept object changed. Formatted as ISO 8601 extended format without time zone designator.
	UpdatedDate *string `json:"updated_date,omitempty"`
	// The date this Concept object was created in the OpenAlex dataset, expressed as an ISO 8601 date string.
	CreatedDate *string `json:"created_date,omitempty"`
	SummaryStats *ListConcepts200ResponseResultsInnerSummaryStats `json:"summary_stats,omitempty"`
}

// NewListConcepts200ResponseResultsInner instantiates a new ListConcepts200ResponseResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConcepts200ResponseResultsInner() *ListConcepts200ResponseResultsInner {
	this := ListConcepts200ResponseResultsInner{}
	return &this
}

// NewListConcepts200ResponseResultsInnerWithDefaults instantiates a new ListConcepts200ResponseResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConcepts200ResponseResultsInnerWithDefaults() *ListConcepts200ResponseResultsInner {
	this := ListConcepts200ResponseResultsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListConcepts200ResponseResultsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConcepts200ResponseResultsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListConcepts200ResponseResultsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListConcepts200ResponseResultsInner) SetId(v string) {
	o.Id = &v
}

// GetWikidata returns the Wikidata field value if set, zero value otherwise.
func (o *ListConcepts200ResponseResultsInner) GetWikidata() string {
	if o == nil || IsNil(o.Wikidata) {
		var ret string
		return ret
	}
	return *o.Wikidata
}

// GetWikidataOk returns a tuple with the Wikidata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConcepts200ResponseResultsInner) GetWikidataOk() (*string, bool) {
	if o == nil || IsNil(o.Wikidata) {
		return nil, false
	}
	return o.Wikidata, true
}

// HasWikidata returns a boolean if a field has been set.
func (o *ListConcepts200ResponseResultsInner) HasWikidata() bool {
	if o != nil && !IsNil(o.Wikidata) {
		return true
	}

	return false
}

// SetWikidata gets a reference to the given string and assigns it to the Wikidata field.
func (o *ListConcepts200ResponseResultsInner) SetWikidata(v string) {
	o.Wikidata = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ListConcepts200ResponseResultsInner) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConcepts200ResponseResultsInner) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ListConcepts200ResponseResultsInner) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ListConcepts200ResponseResultsInner) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *ListConcepts200ResponseResultsInner) GetLevel() int32 {
	if o == nil || IsNil(o.Level) {
		var ret int32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConcepts200ResponseResultsInner) GetLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *ListConcepts200ResponseResultsInner) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given int32 and assigns it to the Level field.
func (o *ListConcepts200ResponseResultsInner) SetLevel(v int32) {
	o.Level = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ListConcepts200ResponseResultsInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConcepts200ResponseResultsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ListConcepts200ResponseResultsInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ListConcepts200ResponseResultsInner) SetDescription(v string) {
	o.Description = &v
}

// GetWorksCount returns the WorksCount field value if set, zero value otherwise.
func (o *ListConcepts200ResponseResultsInner) GetWorksCount() int32 {
	if o == nil || IsNil(o.WorksCount) {
		var ret int32
		return ret
	}
	return *o.WorksCount
}

// GetWorksCountOk returns a tuple with the WorksCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConcepts200ResponseResultsInner) GetWorksCountOk() (*int32, bool) {
	if o == nil || IsNil(o.WorksCount) {
		return nil, false
	}
	return o.WorksCount, true
}

// HasWorksCount returns a boolean if a field has been set.
func (o *ListConcepts200ResponseResultsInner) HasWorksCount() bool {
	if o != nil && !IsNil(o.WorksCount) {
		return true
	}

	return false
}

// SetWorksCount gets a reference to the given int32 and assigns it to the WorksCount field.
func (o *ListConcepts200ResponseResultsInner) SetWorksCount(v int32) {
	o.WorksCount = &v
}

// GetCitedByCount returns the CitedByCount field value if set, zero value otherwise.
func (o *ListConcepts200ResponseResultsInner) GetCitedByCount() int32 {
	if o == nil || IsNil(o.CitedByCount) {
		var ret int32
		return ret
	}
	return *o.CitedByCount
}

// GetCitedByCountOk returns a tuple with the CitedByCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConcepts200ResponseResultsInner) GetCitedByCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CitedByCount) {
		return nil, false
	}
	return o.CitedByCount, true
}

// HasCitedByCount returns a boolean if a field has been set.
func (o *ListConcepts200ResponseResultsInner) HasCitedByCount() bool {
	if o != nil && !IsNil(o.CitedByCount) {
		return true
	}

	return false
}

// SetCitedByCount gets a reference to the given int32 and assigns it to the CitedByCount field.
func (o *ListConcepts200ResponseResultsInner) SetCitedByCount(v int32) {
	o.CitedByCount = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *ListConcepts200ResponseResultsInner) GetIds() ListConcepts200ResponseResultsInnerIds {
	if o == nil || IsNil(o.Ids) {
		var ret ListConcepts200ResponseResultsInnerIds
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConcepts200ResponseResultsInner) GetIdsOk() (*ListConcepts200ResponseResultsInnerIds, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *ListConcepts200ResponseResultsInner) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given ListConcepts200ResponseResultsInnerIds and assigns it to the Ids field.
func (o *ListConcepts200ResponseResultsInner) SetIds(v ListConcepts200ResponseResultsInnerIds) {
	o.Ids = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *ListConcepts200ResponseResultsInner) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConcepts200ResponseResultsInner) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *ListConcepts200ResponseResultsInner) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *ListConcepts200ResponseResultsInner) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetImageThumbnailUrl returns the ImageThumbnailUrl field value if set, zero value otherwise.
func (o *ListConcepts200ResponseResultsInner) GetImageThumbnailUrl() string {
	if o == nil || IsNil(o.ImageThumbnailUrl) {
		var ret string
		return ret
	}
	return *o.ImageThumbnailUrl
}

// GetImageThumbnailUrlOk returns a tuple with the ImageThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConcepts200ResponseResultsInner) GetImageThumbnailUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageThumbnailUrl) {
		return nil, false
	}
	return o.ImageThumbnailUrl, true
}

// HasImageThumbnailUrl returns a boolean if a field has been set.
func (o *ListConcepts200ResponseResultsInner) HasImageThumbnailUrl() bool {
	if o != nil && !IsNil(o.ImageThumbnailUrl) {
		return true
	}

	return false
}

// SetImageThumbnailUrl gets a reference to the given string and assigns it to the ImageThumbnailUrl field.
func (o *ListConcepts200ResponseResultsInner) SetImageThumbnailUrl(v string) {
	o.ImageThumbnailUrl = &v
}

// GetInternational returns the International field value if set, zero value otherwise.
func (o *ListConcepts200ResponseResultsInner) GetInternational() ListConcepts200ResponseResultsInnerInternational {
	if o == nil || IsNil(o.International) {
		var ret ListConcepts200ResponseResultsInnerInternational
		return ret
	}
	return *o.International
}

// GetInternationalOk returns a tuple with the International field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConcepts200ResponseResultsInner) GetInternationalOk() (*ListConcepts200ResponseResultsInnerInternational, bool) {
	if o == nil || IsNil(o.International) {
		return nil, false
	}
	return o.International, true
}

// HasInternational returns a boolean if a field has been set.
func (o *ListConcepts200ResponseResultsInner) HasInternational() bool {
	if o != nil && !IsNil(o.International) {
		return true
	}

	return false
}

// SetInternational gets a reference to the given ListConcepts200ResponseResultsInnerInternational and assigns it to the International field.
func (o *ListConcepts200ResponseResultsInner) SetInternational(v ListConcepts200ResponseResultsInnerInternational) {
	o.International = &v
}

// GetAncestors returns the Ancestors field value if set, zero value otherwise.
func (o *ListConcepts200ResponseResultsInner) GetAncestors() []ListConcepts200ResponseResultsInnerAncestorsInner {
	if o == nil || IsNil(o.Ancestors) {
		var ret []ListConcepts200ResponseResultsInnerAncestorsInner
		return ret
	}
	return o.Ancestors
}

// GetAncestorsOk returns a tuple with the Ancestors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConcepts200ResponseResultsInner) GetAncestorsOk() ([]ListConcepts200ResponseResultsInnerAncestorsInner, bool) {
	if o == nil || IsNil(o.Ancestors) {
		return nil, false
	}
	return o.Ancestors, true
}

// HasAncestors returns a boolean if a field has been set.
func (o *ListConcepts200ResponseResultsInner) HasAncestors() bool {
	if o != nil && !IsNil(o.Ancestors) {
		return true
	}

	return false
}

// SetAncestors gets a reference to the given []ListConcepts200ResponseResultsInnerAncestorsInner and assigns it to the Ancestors field.
func (o *ListConcepts200ResponseResultsInner) SetAncestors(v []ListConcepts200ResponseResultsInnerAncestorsInner) {
	o.Ancestors = v
}

// GetRelatedConcepts returns the RelatedConcepts field value if set, zero value otherwise.
func (o *ListConcepts200ResponseResultsInner) GetRelatedConcepts() []map[string]interface{} {
	if o == nil || IsNil(o.RelatedConcepts) {
		var ret []map[string]interface{}
		return ret
	}
	return o.RelatedConcepts
}

// GetRelatedConceptsOk returns a tuple with the RelatedConcepts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConcepts200ResponseResultsInner) GetRelatedConceptsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.RelatedConcepts) {
		return nil, false
	}
	return o.RelatedConcepts, true
}

// HasRelatedConcepts returns a boolean if a field has been set.
func (o *ListConcepts200ResponseResultsInner) HasRelatedConcepts() bool {
	if o != nil && !IsNil(o.RelatedConcepts) {
		return true
	}

	return false
}

// SetRelatedConcepts gets a reference to the given []map[string]interface{} and assigns it to the RelatedConcepts field.
func (o *ListConcepts200ResponseResultsInner) SetRelatedConcepts(v []map[string]interface{}) {
	o.RelatedConcepts = v
}

// GetCountsByYear returns the CountsByYear field value if set, zero value otherwise.
func (o *ListConcepts200ResponseResultsInner) GetCountsByYear() []ListAuthors200ResponseResultsInnerCountsByYearInner {
	if o == nil || IsNil(o.CountsByYear) {
		var ret []ListAuthors200ResponseResultsInnerCountsByYearInner
		return ret
	}
	return o.CountsByYear
}

// GetCountsByYearOk returns a tuple with the CountsByYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConcepts200ResponseResultsInner) GetCountsByYearOk() ([]ListAuthors200ResponseResultsInnerCountsByYearInner, bool) {
	if o == nil || IsNil(o.CountsByYear) {
		return nil, false
	}
	return o.CountsByYear, true
}

// HasCountsByYear returns a boolean if a field has been set.
func (o *ListConcepts200ResponseResultsInner) HasCountsByYear() bool {
	if o != nil && !IsNil(o.CountsByYear) {
		return true
	}

	return false
}

// SetCountsByYear gets a reference to the given []ListAuthors200ResponseResultsInnerCountsByYearInner and assigns it to the CountsByYear field.
func (o *ListConcepts200ResponseResultsInner) SetCountsByYear(v []ListAuthors200ResponseResultsInnerCountsByYearInner) {
	o.CountsByYear = v
}

// GetWorksApiUrl returns the WorksApiUrl field value if set, zero value otherwise.
func (o *ListConcepts200ResponseResultsInner) GetWorksApiUrl() string {
	if o == nil || IsNil(o.WorksApiUrl) {
		var ret string
		return ret
	}
	return *o.WorksApiUrl
}

// GetWorksApiUrlOk returns a tuple with the WorksApiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConcepts200ResponseResultsInner) GetWorksApiUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WorksApiUrl) {
		return nil, false
	}
	return o.WorksApiUrl, true
}

// HasWorksApiUrl returns a boolean if a field has been set.
func (o *ListConcepts200ResponseResultsInner) HasWorksApiUrl() bool {
	if o != nil && !IsNil(o.WorksApiUrl) {
		return true
	}

	return false
}

// SetWorksApiUrl gets a reference to the given string and assigns it to the WorksApiUrl field.
func (o *ListConcepts200ResponseResultsInner) SetWorksApiUrl(v string) {
	o.WorksApiUrl = &v
}

// GetUpdatedDate returns the UpdatedDate field value if set, zero value otherwise.
func (o *ListConcepts200ResponseResultsInner) GetUpdatedDate() string {
	if o == nil || IsNil(o.UpdatedDate) {
		var ret string
		return ret
	}
	return *o.UpdatedDate
}

// GetUpdatedDateOk returns a tuple with the UpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConcepts200ResponseResultsInner) GetUpdatedDateOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedDate) {
		return nil, false
	}
	return o.UpdatedDate, true
}

// HasUpdatedDate returns a boolean if a field has been set.
func (o *ListConcepts200ResponseResultsInner) HasUpdatedDate() bool {
	if o != nil && !IsNil(o.UpdatedDate) {
		return true
	}

	return false
}

// SetUpdatedDate gets a reference to the given string and assigns it to the UpdatedDate field.
func (o *ListConcepts200ResponseResultsInner) SetUpdatedDate(v string) {
	o.UpdatedDate = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *ListConcepts200ResponseResultsInner) GetCreatedDate() string {
	if o == nil || IsNil(o.CreatedDate) {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConcepts200ResponseResultsInner) GetCreatedDateOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *ListConcepts200ResponseResultsInner) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *ListConcepts200ResponseResultsInner) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

// GetSummaryStats returns the SummaryStats field value if set, zero value otherwise.
func (o *ListConcepts200ResponseResultsInner) GetSummaryStats() ListConcepts200ResponseResultsInnerSummaryStats {
	if o == nil || IsNil(o.SummaryStats) {
		var ret ListConcepts200ResponseResultsInnerSummaryStats
		return ret
	}
	return *o.SummaryStats
}

// GetSummaryStatsOk returns a tuple with the SummaryStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConcepts200ResponseResultsInner) GetSummaryStatsOk() (*ListConcepts200ResponseResultsInnerSummaryStats, bool) {
	if o == nil || IsNil(o.SummaryStats) {
		return nil, false
	}
	return o.SummaryStats, true
}

// HasSummaryStats returns a boolean if a field has been set.
func (o *ListConcepts200ResponseResultsInner) HasSummaryStats() bool {
	if o != nil && !IsNil(o.SummaryStats) {
		return true
	}

	return false
}

// SetSummaryStats gets a reference to the given ListConcepts200ResponseResultsInnerSummaryStats and assigns it to the SummaryStats field.
func (o *ListConcepts200ResponseResultsInner) SetSummaryStats(v ListConcepts200ResponseResultsInnerSummaryStats) {
	o.SummaryStats = &v
}

func (o ListConcepts200ResponseResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListConcepts200ResponseResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Wikidata) {
		toSerialize["wikidata"] = o.Wikidata
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.WorksCount) {
		toSerialize["works_count"] = o.WorksCount
	}
	if !IsNil(o.CitedByCount) {
		toSerialize["cited_by_count"] = o.CitedByCount
	}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["image_url"] = o.ImageUrl
	}
	if !IsNil(o.ImageThumbnailUrl) {
		toSerialize["image_thumbnail_url"] = o.ImageThumbnailUrl
	}
	if !IsNil(o.International) {
		toSerialize["international"] = o.International
	}
	if !IsNil(o.Ancestors) {
		toSerialize["ancestors"] = o.Ancestors
	}
	if !IsNil(o.RelatedConcepts) {
		toSerialize["related_concepts"] = o.RelatedConcepts
	}
	if !IsNil(o.CountsByYear) {
		toSerialize["counts_by_year"] = o.CountsByYear
	}
	if !IsNil(o.WorksApiUrl) {
		toSerialize["works_api_url"] = o.WorksApiUrl
	}
	if !IsNil(o.UpdatedDate) {
		toSerialize["updated_date"] = o.UpdatedDate
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["created_date"] = o.CreatedDate
	}
	if !IsNil(o.SummaryStats) {
		toSerialize["summary_stats"] = o.SummaryStats
	}
	return toSerialize, nil
}

type NullableListConcepts200ResponseResultsInner struct {
	value *ListConcepts200ResponseResultsInner
	isSet bool
}

func (v NullableListConcepts200ResponseResultsInner) Get() *ListConcepts200ResponseResultsInner {
	return v.value
}

func (v *NullableListConcepts200ResponseResultsInner) Set(val *ListConcepts200ResponseResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListConcepts200ResponseResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListConcepts200ResponseResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConcepts200ResponseResultsInner(val *ListConcepts200ResponseResultsInner) *NullableListConcepts200ResponseResultsInner {
	return &NullableListConcepts200ResponseResultsInner{value: val, isSet: true}
}

func (v NullableListConcepts200ResponseResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConcepts200ResponseResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


