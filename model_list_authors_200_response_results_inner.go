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

// checks if the ListAuthors200ResponseResultsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAuthors200ResponseResultsInner{}

// ListAuthors200ResponseResultsInner struct for ListAuthors200ResponseResultsInner
type ListAuthors200ResponseResultsInner struct {
	// List of objects representing the affiliations this author has claimed in their publications.
	Affiliations []ListAuthors200ResponseResultsInnerAffiliationsInner `json:"affiliations,omitempty"`
	// The total number of Works that cite a work this author has created.
	CitedByCount *int32 `json:"cited_by_count,omitempty"`
	// Works_count and cited_by_count for each of the last ten years, binned by year.
	CountsByYear []ListAuthors200ResponseResultsInnerCountsByYearInner `json:"counts_by_year,omitempty"`
	// The date this Author object was created in the OpenAlex dataset.
	CreatedDate *string `json:"created_date,omitempty"`
	// The name of the author as a single string.
	DisplayName *string `json:"display_name,omitempty"`
	// Other ways that we've found this author's name displayed.
	DisplayNameAlternatives []string `json:"display_name_alternatives,omitempty"`
	// The OpenAlex ID for this author.
	Id *string `json:"id,omitempty"`
	Ids *ListAuthors200ResponseResultsInnerIds `json:"ids,omitempty"`
	LastKnownInstitution *ListAuthors200ResponseResultsInnerLastKnownInstitution `json:"last_known_institution,omitempty"`
	// This author's last known institutional affiliations.
	LastKnownInstitutions []ListAuthors200ResponseResultsInnerAffiliationsInnerInstitution `json:"last_known_institutions,omitempty"`
	// The ORCID ID for this author.
	Orcid *string `json:"orcid,omitempty"`
	SummaryStats *ListAuthors200ResponseResultsInnerSummaryStats `json:"summary_stats,omitempty"`
	// The last time anything in this author object changed. Formatted as ISO 8601 extended format without time zone designator.
	UpdatedDate *string `json:"updated_date,omitempty"`
	// A URL that will get you a list of all this author's works.
	WorksApiUrl *string `json:"works_api_url,omitempty"`
	// The number of Works this author has created.
	WorksCount *int32 `json:"works_count,omitempty"`
	// The concepts most frequently applied to works created by this author (deprecated).
	XConcepts []ListAuthors200ResponseResultsInnerXConceptsInner `json:"x_concepts,omitempty"`
}

// NewListAuthors200ResponseResultsInner instantiates a new ListAuthors200ResponseResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAuthors200ResponseResultsInner() *ListAuthors200ResponseResultsInner {
	this := ListAuthors200ResponseResultsInner{}
	return &this
}

// NewListAuthors200ResponseResultsInnerWithDefaults instantiates a new ListAuthors200ResponseResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAuthors200ResponseResultsInnerWithDefaults() *ListAuthors200ResponseResultsInner {
	this := ListAuthors200ResponseResultsInner{}
	return &this
}

// GetAffiliations returns the Affiliations field value if set, zero value otherwise.
func (o *ListAuthors200ResponseResultsInner) GetAffiliations() []ListAuthors200ResponseResultsInnerAffiliationsInner {
	if o == nil || IsNil(o.Affiliations) {
		var ret []ListAuthors200ResponseResultsInnerAffiliationsInner
		return ret
	}
	return o.Affiliations
}

// GetAffiliationsOk returns a tuple with the Affiliations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAuthors200ResponseResultsInner) GetAffiliationsOk() ([]ListAuthors200ResponseResultsInnerAffiliationsInner, bool) {
	if o == nil || IsNil(o.Affiliations) {
		return nil, false
	}
	return o.Affiliations, true
}

// HasAffiliations returns a boolean if a field has been set.
func (o *ListAuthors200ResponseResultsInner) HasAffiliations() bool {
	if o != nil && !IsNil(o.Affiliations) {
		return true
	}

	return false
}

// SetAffiliations gets a reference to the given []ListAuthors200ResponseResultsInnerAffiliationsInner and assigns it to the Affiliations field.
func (o *ListAuthors200ResponseResultsInner) SetAffiliations(v []ListAuthors200ResponseResultsInnerAffiliationsInner) {
	o.Affiliations = v
}

// GetCitedByCount returns the CitedByCount field value if set, zero value otherwise.
func (o *ListAuthors200ResponseResultsInner) GetCitedByCount() int32 {
	if o == nil || IsNil(o.CitedByCount) {
		var ret int32
		return ret
	}
	return *o.CitedByCount
}

// GetCitedByCountOk returns a tuple with the CitedByCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAuthors200ResponseResultsInner) GetCitedByCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CitedByCount) {
		return nil, false
	}
	return o.CitedByCount, true
}

// HasCitedByCount returns a boolean if a field has been set.
func (o *ListAuthors200ResponseResultsInner) HasCitedByCount() bool {
	if o != nil && !IsNil(o.CitedByCount) {
		return true
	}

	return false
}

// SetCitedByCount gets a reference to the given int32 and assigns it to the CitedByCount field.
func (o *ListAuthors200ResponseResultsInner) SetCitedByCount(v int32) {
	o.CitedByCount = &v
}

// GetCountsByYear returns the CountsByYear field value if set, zero value otherwise.
func (o *ListAuthors200ResponseResultsInner) GetCountsByYear() []ListAuthors200ResponseResultsInnerCountsByYearInner {
	if o == nil || IsNil(o.CountsByYear) {
		var ret []ListAuthors200ResponseResultsInnerCountsByYearInner
		return ret
	}
	return o.CountsByYear
}

// GetCountsByYearOk returns a tuple with the CountsByYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAuthors200ResponseResultsInner) GetCountsByYearOk() ([]ListAuthors200ResponseResultsInnerCountsByYearInner, bool) {
	if o == nil || IsNil(o.CountsByYear) {
		return nil, false
	}
	return o.CountsByYear, true
}

// HasCountsByYear returns a boolean if a field has been set.
func (o *ListAuthors200ResponseResultsInner) HasCountsByYear() bool {
	if o != nil && !IsNil(o.CountsByYear) {
		return true
	}

	return false
}

// SetCountsByYear gets a reference to the given []ListAuthors200ResponseResultsInnerCountsByYearInner and assigns it to the CountsByYear field.
func (o *ListAuthors200ResponseResultsInner) SetCountsByYear(v []ListAuthors200ResponseResultsInnerCountsByYearInner) {
	o.CountsByYear = v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *ListAuthors200ResponseResultsInner) GetCreatedDate() string {
	if o == nil || IsNil(o.CreatedDate) {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAuthors200ResponseResultsInner) GetCreatedDateOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *ListAuthors200ResponseResultsInner) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *ListAuthors200ResponseResultsInner) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ListAuthors200ResponseResultsInner) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAuthors200ResponseResultsInner) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ListAuthors200ResponseResultsInner) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ListAuthors200ResponseResultsInner) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDisplayNameAlternatives returns the DisplayNameAlternatives field value if set, zero value otherwise.
func (o *ListAuthors200ResponseResultsInner) GetDisplayNameAlternatives() []string {
	if o == nil || IsNil(o.DisplayNameAlternatives) {
		var ret []string
		return ret
	}
	return o.DisplayNameAlternatives
}

// GetDisplayNameAlternativesOk returns a tuple with the DisplayNameAlternatives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAuthors200ResponseResultsInner) GetDisplayNameAlternativesOk() ([]string, bool) {
	if o == nil || IsNil(o.DisplayNameAlternatives) {
		return nil, false
	}
	return o.DisplayNameAlternatives, true
}

// HasDisplayNameAlternatives returns a boolean if a field has been set.
func (o *ListAuthors200ResponseResultsInner) HasDisplayNameAlternatives() bool {
	if o != nil && !IsNil(o.DisplayNameAlternatives) {
		return true
	}

	return false
}

// SetDisplayNameAlternatives gets a reference to the given []string and assigns it to the DisplayNameAlternatives field.
func (o *ListAuthors200ResponseResultsInner) SetDisplayNameAlternatives(v []string) {
	o.DisplayNameAlternatives = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListAuthors200ResponseResultsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAuthors200ResponseResultsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListAuthors200ResponseResultsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListAuthors200ResponseResultsInner) SetId(v string) {
	o.Id = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *ListAuthors200ResponseResultsInner) GetIds() ListAuthors200ResponseResultsInnerIds {
	if o == nil || IsNil(o.Ids) {
		var ret ListAuthors200ResponseResultsInnerIds
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAuthors200ResponseResultsInner) GetIdsOk() (*ListAuthors200ResponseResultsInnerIds, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *ListAuthors200ResponseResultsInner) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given ListAuthors200ResponseResultsInnerIds and assigns it to the Ids field.
func (o *ListAuthors200ResponseResultsInner) SetIds(v ListAuthors200ResponseResultsInnerIds) {
	o.Ids = &v
}

// GetLastKnownInstitution returns the LastKnownInstitution field value if set, zero value otherwise.
func (o *ListAuthors200ResponseResultsInner) GetLastKnownInstitution() ListAuthors200ResponseResultsInnerLastKnownInstitution {
	if o == nil || IsNil(o.LastKnownInstitution) {
		var ret ListAuthors200ResponseResultsInnerLastKnownInstitution
		return ret
	}
	return *o.LastKnownInstitution
}

// GetLastKnownInstitutionOk returns a tuple with the LastKnownInstitution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAuthors200ResponseResultsInner) GetLastKnownInstitutionOk() (*ListAuthors200ResponseResultsInnerLastKnownInstitution, bool) {
	if o == nil || IsNil(o.LastKnownInstitution) {
		return nil, false
	}
	return o.LastKnownInstitution, true
}

// HasLastKnownInstitution returns a boolean if a field has been set.
func (o *ListAuthors200ResponseResultsInner) HasLastKnownInstitution() bool {
	if o != nil && !IsNil(o.LastKnownInstitution) {
		return true
	}

	return false
}

// SetLastKnownInstitution gets a reference to the given ListAuthors200ResponseResultsInnerLastKnownInstitution and assigns it to the LastKnownInstitution field.
func (o *ListAuthors200ResponseResultsInner) SetLastKnownInstitution(v ListAuthors200ResponseResultsInnerLastKnownInstitution) {
	o.LastKnownInstitution = &v
}

// GetLastKnownInstitutions returns the LastKnownInstitutions field value if set, zero value otherwise.
func (o *ListAuthors200ResponseResultsInner) GetLastKnownInstitutions() []ListAuthors200ResponseResultsInnerAffiliationsInnerInstitution {
	if o == nil || IsNil(o.LastKnownInstitutions) {
		var ret []ListAuthors200ResponseResultsInnerAffiliationsInnerInstitution
		return ret
	}
	return o.LastKnownInstitutions
}

// GetLastKnownInstitutionsOk returns a tuple with the LastKnownInstitutions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAuthors200ResponseResultsInner) GetLastKnownInstitutionsOk() ([]ListAuthors200ResponseResultsInnerAffiliationsInnerInstitution, bool) {
	if o == nil || IsNil(o.LastKnownInstitutions) {
		return nil, false
	}
	return o.LastKnownInstitutions, true
}

// HasLastKnownInstitutions returns a boolean if a field has been set.
func (o *ListAuthors200ResponseResultsInner) HasLastKnownInstitutions() bool {
	if o != nil && !IsNil(o.LastKnownInstitutions) {
		return true
	}

	return false
}

// SetLastKnownInstitutions gets a reference to the given []ListAuthors200ResponseResultsInnerAffiliationsInnerInstitution and assigns it to the LastKnownInstitutions field.
func (o *ListAuthors200ResponseResultsInner) SetLastKnownInstitutions(v []ListAuthors200ResponseResultsInnerAffiliationsInnerInstitution) {
	o.LastKnownInstitutions = v
}

// GetOrcid returns the Orcid field value if set, zero value otherwise.
func (o *ListAuthors200ResponseResultsInner) GetOrcid() string {
	if o == nil || IsNil(o.Orcid) {
		var ret string
		return ret
	}
	return *o.Orcid
}

// GetOrcidOk returns a tuple with the Orcid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAuthors200ResponseResultsInner) GetOrcidOk() (*string, bool) {
	if o == nil || IsNil(o.Orcid) {
		return nil, false
	}
	return o.Orcid, true
}

// HasOrcid returns a boolean if a field has been set.
func (o *ListAuthors200ResponseResultsInner) HasOrcid() bool {
	if o != nil && !IsNil(o.Orcid) {
		return true
	}

	return false
}

// SetOrcid gets a reference to the given string and assigns it to the Orcid field.
func (o *ListAuthors200ResponseResultsInner) SetOrcid(v string) {
	o.Orcid = &v
}

// GetSummaryStats returns the SummaryStats field value if set, zero value otherwise.
func (o *ListAuthors200ResponseResultsInner) GetSummaryStats() ListAuthors200ResponseResultsInnerSummaryStats {
	if o == nil || IsNil(o.SummaryStats) {
		var ret ListAuthors200ResponseResultsInnerSummaryStats
		return ret
	}
	return *o.SummaryStats
}

// GetSummaryStatsOk returns a tuple with the SummaryStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAuthors200ResponseResultsInner) GetSummaryStatsOk() (*ListAuthors200ResponseResultsInnerSummaryStats, bool) {
	if o == nil || IsNil(o.SummaryStats) {
		return nil, false
	}
	return o.SummaryStats, true
}

// HasSummaryStats returns a boolean if a field has been set.
func (o *ListAuthors200ResponseResultsInner) HasSummaryStats() bool {
	if o != nil && !IsNil(o.SummaryStats) {
		return true
	}

	return false
}

// SetSummaryStats gets a reference to the given ListAuthors200ResponseResultsInnerSummaryStats and assigns it to the SummaryStats field.
func (o *ListAuthors200ResponseResultsInner) SetSummaryStats(v ListAuthors200ResponseResultsInnerSummaryStats) {
	o.SummaryStats = &v
}

// GetUpdatedDate returns the UpdatedDate field value if set, zero value otherwise.
func (o *ListAuthors200ResponseResultsInner) GetUpdatedDate() string {
	if o == nil || IsNil(o.UpdatedDate) {
		var ret string
		return ret
	}
	return *o.UpdatedDate
}

// GetUpdatedDateOk returns a tuple with the UpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAuthors200ResponseResultsInner) GetUpdatedDateOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedDate) {
		return nil, false
	}
	return o.UpdatedDate, true
}

// HasUpdatedDate returns a boolean if a field has been set.
func (o *ListAuthors200ResponseResultsInner) HasUpdatedDate() bool {
	if o != nil && !IsNil(o.UpdatedDate) {
		return true
	}

	return false
}

// SetUpdatedDate gets a reference to the given string and assigns it to the UpdatedDate field.
func (o *ListAuthors200ResponseResultsInner) SetUpdatedDate(v string) {
	o.UpdatedDate = &v
}

// GetWorksApiUrl returns the WorksApiUrl field value if set, zero value otherwise.
func (o *ListAuthors200ResponseResultsInner) GetWorksApiUrl() string {
	if o == nil || IsNil(o.WorksApiUrl) {
		var ret string
		return ret
	}
	return *o.WorksApiUrl
}

// GetWorksApiUrlOk returns a tuple with the WorksApiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAuthors200ResponseResultsInner) GetWorksApiUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WorksApiUrl) {
		return nil, false
	}
	return o.WorksApiUrl, true
}

// HasWorksApiUrl returns a boolean if a field has been set.
func (o *ListAuthors200ResponseResultsInner) HasWorksApiUrl() bool {
	if o != nil && !IsNil(o.WorksApiUrl) {
		return true
	}

	return false
}

// SetWorksApiUrl gets a reference to the given string and assigns it to the WorksApiUrl field.
func (o *ListAuthors200ResponseResultsInner) SetWorksApiUrl(v string) {
	o.WorksApiUrl = &v
}

// GetWorksCount returns the WorksCount field value if set, zero value otherwise.
func (o *ListAuthors200ResponseResultsInner) GetWorksCount() int32 {
	if o == nil || IsNil(o.WorksCount) {
		var ret int32
		return ret
	}
	return *o.WorksCount
}

// GetWorksCountOk returns a tuple with the WorksCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAuthors200ResponseResultsInner) GetWorksCountOk() (*int32, bool) {
	if o == nil || IsNil(o.WorksCount) {
		return nil, false
	}
	return o.WorksCount, true
}

// HasWorksCount returns a boolean if a field has been set.
func (o *ListAuthors200ResponseResultsInner) HasWorksCount() bool {
	if o != nil && !IsNil(o.WorksCount) {
		return true
	}

	return false
}

// SetWorksCount gets a reference to the given int32 and assigns it to the WorksCount field.
func (o *ListAuthors200ResponseResultsInner) SetWorksCount(v int32) {
	o.WorksCount = &v
}

// GetXConcepts returns the XConcepts field value if set, zero value otherwise.
func (o *ListAuthors200ResponseResultsInner) GetXConcepts() []ListAuthors200ResponseResultsInnerXConceptsInner {
	if o == nil || IsNil(o.XConcepts) {
		var ret []ListAuthors200ResponseResultsInnerXConceptsInner
		return ret
	}
	return o.XConcepts
}

// GetXConceptsOk returns a tuple with the XConcepts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAuthors200ResponseResultsInner) GetXConceptsOk() ([]ListAuthors200ResponseResultsInnerXConceptsInner, bool) {
	if o == nil || IsNil(o.XConcepts) {
		return nil, false
	}
	return o.XConcepts, true
}

// HasXConcepts returns a boolean if a field has been set.
func (o *ListAuthors200ResponseResultsInner) HasXConcepts() bool {
	if o != nil && !IsNil(o.XConcepts) {
		return true
	}

	return false
}

// SetXConcepts gets a reference to the given []ListAuthors200ResponseResultsInnerXConceptsInner and assigns it to the XConcepts field.
func (o *ListAuthors200ResponseResultsInner) SetXConcepts(v []ListAuthors200ResponseResultsInnerXConceptsInner) {
	o.XConcepts = v
}

func (o ListAuthors200ResponseResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAuthors200ResponseResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Affiliations) {
		toSerialize["affiliations"] = o.Affiliations
	}
	if !IsNil(o.CitedByCount) {
		toSerialize["cited_by_count"] = o.CitedByCount
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
	if !IsNil(o.DisplayNameAlternatives) {
		toSerialize["display_name_alternatives"] = o.DisplayNameAlternatives
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !IsNil(o.LastKnownInstitution) {
		toSerialize["last_known_institution"] = o.LastKnownInstitution
	}
	if !IsNil(o.LastKnownInstitutions) {
		toSerialize["last_known_institutions"] = o.LastKnownInstitutions
	}
	if !IsNil(o.Orcid) {
		toSerialize["orcid"] = o.Orcid
	}
	if !IsNil(o.SummaryStats) {
		toSerialize["summary_stats"] = o.SummaryStats
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

type NullableListAuthors200ResponseResultsInner struct {
	value *ListAuthors200ResponseResultsInner
	isSet bool
}

func (v NullableListAuthors200ResponseResultsInner) Get() *ListAuthors200ResponseResultsInner {
	return v.value
}

func (v *NullableListAuthors200ResponseResultsInner) Set(val *ListAuthors200ResponseResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListAuthors200ResponseResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListAuthors200ResponseResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAuthors200ResponseResultsInner(val *ListAuthors200ResponseResultsInner) *NullableListAuthors200ResponseResultsInner {
	return &NullableListAuthors200ResponseResultsInner{value: val, isSet: true}
}

func (v NullableListAuthors200ResponseResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAuthors200ResponseResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


