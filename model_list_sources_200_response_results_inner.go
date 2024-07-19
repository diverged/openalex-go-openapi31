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

// checks if the ListSources200ResponseResultsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSources200ResponseResultsInner{}

// ListSources200ResponseResultsInner struct for ListSources200ResponseResultsInner
type ListSources200ResponseResultsInner struct {
	// An abbreviated title obtained from the ISSN Centre.
	AbbreviatedTitle *string `json:"abbreviated_title,omitempty"`
	// Alternate titles for this source, as obtained from the ISSN Centre and individual work records, like Crossref DOIs, that carry the source name as a string. These are commonly abbreviations or translations of the source's canonical name.
	AlternateTitles []string `json:"alternate_titles,omitempty"`
	// Article processing charge information, taken directly from DOAJ.
	ApcPrices []ListSources200ResponseResultsInnerApcPricesInner `json:"apc_prices,omitempty"`
	// The source's article processing charge in US Dollars, if available from DOAJ. The apc_usd value is calculated by taking the APC price with a currency of USD if it is available. If it's not available, we convert the first available value from apc_prices into USD, using recent exchange rates.
	ApcUsd *int32 `json:"apc_usd,omitempty"`
	// The total number of Works that cite a Work hosted in this source.
	CitedByCount *int32 `json:"cited_by_count,omitempty"`
	// The country that this source is associated with, represented as an ISO two-letter country code.
	CountryCode *string `json:"country_code,omitempty"`
	// works_count and cited_by_count for each of the last ten years, binned by year. If the source was founded less than ten years ago, there will naturally be fewer than ten years in this list. Years with zero citations and zero works have been removed.
	CountsByYear []ListAuthors200ResponseResultsInnerCountsByYearInner `json:"counts_by_year,omitempty"`
	// The date this Source object was created in the OpenAlex dataset, expressed as an ISO 8601 date string.
	CreatedDate *string `json:"created_date,omitempty"`
	// The name of the source.
	DisplayName *string `json:"display_name,omitempty"`
	// The starting page for navigating the contents of this source; the homepage for this source's website.
	HomepageUrl *string `json:"homepage_url,omitempty"`
	// The host organization for this source as an OpenAlex ID. This will be an Institution.id if the source is a repository, and a Publisher.id if the source is a journal, conference, or eBook platform (based on the type field).
	HostOrganization *string `json:"host_organization,omitempty"`
	// OpenAlex IDs of the host organization's lineage. This will only be included if the host_organization is a publisher (and not if the host_organization is an institution).
	HostOrganizationLineage []string `json:"host_organization_lineage,omitempty"`
	// The display_name from the host_organization, shown for convenience.
	HostOrganizationName *string `json:"host_organization_name,omitempty"`
	// The OpenAlex ID for this source.
	Id *string `json:"id,omitempty"`
	Ids *ListSources200ResponseResultsInnerIds `json:"ids,omitempty"`
	// Whether this source is identified as a \"core source\" by [CWTS](https://www.cwts.nl/), used in the [Open Leiden Ranking](https://open.leidenranking.com/) of universities around the world. The list of core sources can be found [here](https://zenodo.org/records/10949671).
	IsCore *bool `json:"is_core,omitempty"`
	// Whether this is a journal listed in the Directory of Open Access Journals (DOAJ).
	IsInDoaj *bool `json:"is_in_doaj,omitempty"`
	// Whether this is currently fully-open-access source. This could be true for a preprint repository where everything uploaded is free to read, or for a Gold or Diamond open access journal, where all newly published Works are available for free under an open license.
	IsOa *bool `json:"is_oa,omitempty"`
	// The ISSNs used by this source. Many publications have multiple ISSNs, so ISSN-L should be used when possible.
	Issn []string `json:"issn,omitempty"`
	// The ISSN-L identifying this source. This is the Canonical External ID for sources. ISSN-L or Linking ISSN solves the problem by designating a single canonical ISSN for all media versions of the title. It's usually the same as the print ISSN.
	IssnL *string `json:"issn_l,omitempty"`
	// Societies on whose behalf the source is published and maintained, obtained from our crowdsourced list.
	Societies []ListSources200ResponseResultsInnerSocietiesInner `json:"societies,omitempty"`
	SummaryStats *ListSources200ResponseResultsInnerSummaryStats `json:"summary_stats,omitempty"`
	// The type of source.
	Type *string `json:"type,omitempty"`
	// The last time anything in this author object changed. Formatted as ISO 8601 extended format without time zone designator.
	UpdatedDate *string `json:"updated_date,omitempty"`
	// A URL that will get you a list of all this source's Works. We express this as an API URL (instead of just listing the works themselves) because sometimes a source's publication list is too long to reasonably fit into a single Source object.
	WorksApiUrl *string `json:"works_api_url,omitempty"`
	// The number of Works this source hosts.
	WorksCount *int32 `json:"works_count,omitempty"`
	// The Concepts most frequently applied to works hosted by this source.
	XConcepts []map[string]interface{} `json:"x_concepts,omitempty"`
}

// NewListSources200ResponseResultsInner instantiates a new ListSources200ResponseResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSources200ResponseResultsInner() *ListSources200ResponseResultsInner {
	this := ListSources200ResponseResultsInner{}
	return &this
}

// NewListSources200ResponseResultsInnerWithDefaults instantiates a new ListSources200ResponseResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSources200ResponseResultsInnerWithDefaults() *ListSources200ResponseResultsInner {
	this := ListSources200ResponseResultsInner{}
	return &this
}

// GetAbbreviatedTitle returns the AbbreviatedTitle field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetAbbreviatedTitle() string {
	if o == nil || IsNil(o.AbbreviatedTitle) {
		var ret string
		return ret
	}
	return *o.AbbreviatedTitle
}

// GetAbbreviatedTitleOk returns a tuple with the AbbreviatedTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetAbbreviatedTitleOk() (*string, bool) {
	if o == nil || IsNil(o.AbbreviatedTitle) {
		return nil, false
	}
	return o.AbbreviatedTitle, true
}

// HasAbbreviatedTitle returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasAbbreviatedTitle() bool {
	if o != nil && !IsNil(o.AbbreviatedTitle) {
		return true
	}

	return false
}

// SetAbbreviatedTitle gets a reference to the given string and assigns it to the AbbreviatedTitle field.
func (o *ListSources200ResponseResultsInner) SetAbbreviatedTitle(v string) {
	o.AbbreviatedTitle = &v
}

// GetAlternateTitles returns the AlternateTitles field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetAlternateTitles() []string {
	if o == nil || IsNil(o.AlternateTitles) {
		var ret []string
		return ret
	}
	return o.AlternateTitles
}

// GetAlternateTitlesOk returns a tuple with the AlternateTitles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetAlternateTitlesOk() ([]string, bool) {
	if o == nil || IsNil(o.AlternateTitles) {
		return nil, false
	}
	return o.AlternateTitles, true
}

// HasAlternateTitles returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasAlternateTitles() bool {
	if o != nil && !IsNil(o.AlternateTitles) {
		return true
	}

	return false
}

// SetAlternateTitles gets a reference to the given []string and assigns it to the AlternateTitles field.
func (o *ListSources200ResponseResultsInner) SetAlternateTitles(v []string) {
	o.AlternateTitles = v
}

// GetApcPrices returns the ApcPrices field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetApcPrices() []ListSources200ResponseResultsInnerApcPricesInner {
	if o == nil || IsNil(o.ApcPrices) {
		var ret []ListSources200ResponseResultsInnerApcPricesInner
		return ret
	}
	return o.ApcPrices
}

// GetApcPricesOk returns a tuple with the ApcPrices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetApcPricesOk() ([]ListSources200ResponseResultsInnerApcPricesInner, bool) {
	if o == nil || IsNil(o.ApcPrices) {
		return nil, false
	}
	return o.ApcPrices, true
}

// HasApcPrices returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasApcPrices() bool {
	if o != nil && !IsNil(o.ApcPrices) {
		return true
	}

	return false
}

// SetApcPrices gets a reference to the given []ListSources200ResponseResultsInnerApcPricesInner and assigns it to the ApcPrices field.
func (o *ListSources200ResponseResultsInner) SetApcPrices(v []ListSources200ResponseResultsInnerApcPricesInner) {
	o.ApcPrices = v
}

// GetApcUsd returns the ApcUsd field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetApcUsd() int32 {
	if o == nil || IsNil(o.ApcUsd) {
		var ret int32
		return ret
	}
	return *o.ApcUsd
}

// GetApcUsdOk returns a tuple with the ApcUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetApcUsdOk() (*int32, bool) {
	if o == nil || IsNil(o.ApcUsd) {
		return nil, false
	}
	return o.ApcUsd, true
}

// HasApcUsd returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasApcUsd() bool {
	if o != nil && !IsNil(o.ApcUsd) {
		return true
	}

	return false
}

// SetApcUsd gets a reference to the given int32 and assigns it to the ApcUsd field.
func (o *ListSources200ResponseResultsInner) SetApcUsd(v int32) {
	o.ApcUsd = &v
}

// GetCitedByCount returns the CitedByCount field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetCitedByCount() int32 {
	if o == nil || IsNil(o.CitedByCount) {
		var ret int32
		return ret
	}
	return *o.CitedByCount
}

// GetCitedByCountOk returns a tuple with the CitedByCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetCitedByCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CitedByCount) {
		return nil, false
	}
	return o.CitedByCount, true
}

// HasCitedByCount returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasCitedByCount() bool {
	if o != nil && !IsNil(o.CitedByCount) {
		return true
	}

	return false
}

// SetCitedByCount gets a reference to the given int32 and assigns it to the CitedByCount field.
func (o *ListSources200ResponseResultsInner) SetCitedByCount(v int32) {
	o.CitedByCount = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *ListSources200ResponseResultsInner) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetCountsByYear returns the CountsByYear field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetCountsByYear() []ListAuthors200ResponseResultsInnerCountsByYearInner {
	if o == nil || IsNil(o.CountsByYear) {
		var ret []ListAuthors200ResponseResultsInnerCountsByYearInner
		return ret
	}
	return o.CountsByYear
}

// GetCountsByYearOk returns a tuple with the CountsByYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetCountsByYearOk() ([]ListAuthors200ResponseResultsInnerCountsByYearInner, bool) {
	if o == nil || IsNil(o.CountsByYear) {
		return nil, false
	}
	return o.CountsByYear, true
}

// HasCountsByYear returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasCountsByYear() bool {
	if o != nil && !IsNil(o.CountsByYear) {
		return true
	}

	return false
}

// SetCountsByYear gets a reference to the given []ListAuthors200ResponseResultsInnerCountsByYearInner and assigns it to the CountsByYear field.
func (o *ListSources200ResponseResultsInner) SetCountsByYear(v []ListAuthors200ResponseResultsInnerCountsByYearInner) {
	o.CountsByYear = v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetCreatedDate() string {
	if o == nil || IsNil(o.CreatedDate) {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetCreatedDateOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *ListSources200ResponseResultsInner) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ListSources200ResponseResultsInner) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetHomepageUrl returns the HomepageUrl field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetHomepageUrl() string {
	if o == nil || IsNil(o.HomepageUrl) {
		var ret string
		return ret
	}
	return *o.HomepageUrl
}

// GetHomepageUrlOk returns a tuple with the HomepageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetHomepageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HomepageUrl) {
		return nil, false
	}
	return o.HomepageUrl, true
}

// HasHomepageUrl returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasHomepageUrl() bool {
	if o != nil && !IsNil(o.HomepageUrl) {
		return true
	}

	return false
}

// SetHomepageUrl gets a reference to the given string and assigns it to the HomepageUrl field.
func (o *ListSources200ResponseResultsInner) SetHomepageUrl(v string) {
	o.HomepageUrl = &v
}

// GetHostOrganization returns the HostOrganization field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetHostOrganization() string {
	if o == nil || IsNil(o.HostOrganization) {
		var ret string
		return ret
	}
	return *o.HostOrganization
}

// GetHostOrganizationOk returns a tuple with the HostOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetHostOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.HostOrganization) {
		return nil, false
	}
	return o.HostOrganization, true
}

// HasHostOrganization returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasHostOrganization() bool {
	if o != nil && !IsNil(o.HostOrganization) {
		return true
	}

	return false
}

// SetHostOrganization gets a reference to the given string and assigns it to the HostOrganization field.
func (o *ListSources200ResponseResultsInner) SetHostOrganization(v string) {
	o.HostOrganization = &v
}

// GetHostOrganizationLineage returns the HostOrganizationLineage field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetHostOrganizationLineage() []string {
	if o == nil || IsNil(o.HostOrganizationLineage) {
		var ret []string
		return ret
	}
	return o.HostOrganizationLineage
}

// GetHostOrganizationLineageOk returns a tuple with the HostOrganizationLineage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetHostOrganizationLineageOk() ([]string, bool) {
	if o == nil || IsNil(o.HostOrganizationLineage) {
		return nil, false
	}
	return o.HostOrganizationLineage, true
}

// HasHostOrganizationLineage returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasHostOrganizationLineage() bool {
	if o != nil && !IsNil(o.HostOrganizationLineage) {
		return true
	}

	return false
}

// SetHostOrganizationLineage gets a reference to the given []string and assigns it to the HostOrganizationLineage field.
func (o *ListSources200ResponseResultsInner) SetHostOrganizationLineage(v []string) {
	o.HostOrganizationLineage = v
}

// GetHostOrganizationName returns the HostOrganizationName field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetHostOrganizationName() string {
	if o == nil || IsNil(o.HostOrganizationName) {
		var ret string
		return ret
	}
	return *o.HostOrganizationName
}

// GetHostOrganizationNameOk returns a tuple with the HostOrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetHostOrganizationNameOk() (*string, bool) {
	if o == nil || IsNil(o.HostOrganizationName) {
		return nil, false
	}
	return o.HostOrganizationName, true
}

// HasHostOrganizationName returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasHostOrganizationName() bool {
	if o != nil && !IsNil(o.HostOrganizationName) {
		return true
	}

	return false
}

// SetHostOrganizationName gets a reference to the given string and assigns it to the HostOrganizationName field.
func (o *ListSources200ResponseResultsInner) SetHostOrganizationName(v string) {
	o.HostOrganizationName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListSources200ResponseResultsInner) SetId(v string) {
	o.Id = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetIds() ListSources200ResponseResultsInnerIds {
	if o == nil || IsNil(o.Ids) {
		var ret ListSources200ResponseResultsInnerIds
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetIdsOk() (*ListSources200ResponseResultsInnerIds, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given ListSources200ResponseResultsInnerIds and assigns it to the Ids field.
func (o *ListSources200ResponseResultsInner) SetIds(v ListSources200ResponseResultsInnerIds) {
	o.Ids = &v
}

// GetIsCore returns the IsCore field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetIsCore() bool {
	if o == nil || IsNil(o.IsCore) {
		var ret bool
		return ret
	}
	return *o.IsCore
}

// GetIsCoreOk returns a tuple with the IsCore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetIsCoreOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCore) {
		return nil, false
	}
	return o.IsCore, true
}

// HasIsCore returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasIsCore() bool {
	if o != nil && !IsNil(o.IsCore) {
		return true
	}

	return false
}

// SetIsCore gets a reference to the given bool and assigns it to the IsCore field.
func (o *ListSources200ResponseResultsInner) SetIsCore(v bool) {
	o.IsCore = &v
}

// GetIsInDoaj returns the IsInDoaj field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetIsInDoaj() bool {
	if o == nil || IsNil(o.IsInDoaj) {
		var ret bool
		return ret
	}
	return *o.IsInDoaj
}

// GetIsInDoajOk returns a tuple with the IsInDoaj field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetIsInDoajOk() (*bool, bool) {
	if o == nil || IsNil(o.IsInDoaj) {
		return nil, false
	}
	return o.IsInDoaj, true
}

// HasIsInDoaj returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasIsInDoaj() bool {
	if o != nil && !IsNil(o.IsInDoaj) {
		return true
	}

	return false
}

// SetIsInDoaj gets a reference to the given bool and assigns it to the IsInDoaj field.
func (o *ListSources200ResponseResultsInner) SetIsInDoaj(v bool) {
	o.IsInDoaj = &v
}

// GetIsOa returns the IsOa field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetIsOa() bool {
	if o == nil || IsNil(o.IsOa) {
		var ret bool
		return ret
	}
	return *o.IsOa
}

// GetIsOaOk returns a tuple with the IsOa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetIsOaOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOa) {
		return nil, false
	}
	return o.IsOa, true
}

// HasIsOa returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasIsOa() bool {
	if o != nil && !IsNil(o.IsOa) {
		return true
	}

	return false
}

// SetIsOa gets a reference to the given bool and assigns it to the IsOa field.
func (o *ListSources200ResponseResultsInner) SetIsOa(v bool) {
	o.IsOa = &v
}

// GetIssn returns the Issn field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetIssn() []string {
	if o == nil || IsNil(o.Issn) {
		var ret []string
		return ret
	}
	return o.Issn
}

// GetIssnOk returns a tuple with the Issn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetIssnOk() ([]string, bool) {
	if o == nil || IsNil(o.Issn) {
		return nil, false
	}
	return o.Issn, true
}

// HasIssn returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasIssn() bool {
	if o != nil && !IsNil(o.Issn) {
		return true
	}

	return false
}

// SetIssn gets a reference to the given []string and assigns it to the Issn field.
func (o *ListSources200ResponseResultsInner) SetIssn(v []string) {
	o.Issn = v
}

// GetIssnL returns the IssnL field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetIssnL() string {
	if o == nil || IsNil(o.IssnL) {
		var ret string
		return ret
	}
	return *o.IssnL
}

// GetIssnLOk returns a tuple with the IssnL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetIssnLOk() (*string, bool) {
	if o == nil || IsNil(o.IssnL) {
		return nil, false
	}
	return o.IssnL, true
}

// HasIssnL returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasIssnL() bool {
	if o != nil && !IsNil(o.IssnL) {
		return true
	}

	return false
}

// SetIssnL gets a reference to the given string and assigns it to the IssnL field.
func (o *ListSources200ResponseResultsInner) SetIssnL(v string) {
	o.IssnL = &v
}

// GetSocieties returns the Societies field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetSocieties() []ListSources200ResponseResultsInnerSocietiesInner {
	if o == nil || IsNil(o.Societies) {
		var ret []ListSources200ResponseResultsInnerSocietiesInner
		return ret
	}
	return o.Societies
}

// GetSocietiesOk returns a tuple with the Societies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetSocietiesOk() ([]ListSources200ResponseResultsInnerSocietiesInner, bool) {
	if o == nil || IsNil(o.Societies) {
		return nil, false
	}
	return o.Societies, true
}

// HasSocieties returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasSocieties() bool {
	if o != nil && !IsNil(o.Societies) {
		return true
	}

	return false
}

// SetSocieties gets a reference to the given []ListSources200ResponseResultsInnerSocietiesInner and assigns it to the Societies field.
func (o *ListSources200ResponseResultsInner) SetSocieties(v []ListSources200ResponseResultsInnerSocietiesInner) {
	o.Societies = v
}

// GetSummaryStats returns the SummaryStats field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetSummaryStats() ListSources200ResponseResultsInnerSummaryStats {
	if o == nil || IsNil(o.SummaryStats) {
		var ret ListSources200ResponseResultsInnerSummaryStats
		return ret
	}
	return *o.SummaryStats
}

// GetSummaryStatsOk returns a tuple with the SummaryStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetSummaryStatsOk() (*ListSources200ResponseResultsInnerSummaryStats, bool) {
	if o == nil || IsNil(o.SummaryStats) {
		return nil, false
	}
	return o.SummaryStats, true
}

// HasSummaryStats returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasSummaryStats() bool {
	if o != nil && !IsNil(o.SummaryStats) {
		return true
	}

	return false
}

// SetSummaryStats gets a reference to the given ListSources200ResponseResultsInnerSummaryStats and assigns it to the SummaryStats field.
func (o *ListSources200ResponseResultsInner) SetSummaryStats(v ListSources200ResponseResultsInnerSummaryStats) {
	o.SummaryStats = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ListSources200ResponseResultsInner) SetType(v string) {
	o.Type = &v
}

// GetUpdatedDate returns the UpdatedDate field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetUpdatedDate() string {
	if o == nil || IsNil(o.UpdatedDate) {
		var ret string
		return ret
	}
	return *o.UpdatedDate
}

// GetUpdatedDateOk returns a tuple with the UpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetUpdatedDateOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedDate) {
		return nil, false
	}
	return o.UpdatedDate, true
}

// HasUpdatedDate returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasUpdatedDate() bool {
	if o != nil && !IsNil(o.UpdatedDate) {
		return true
	}

	return false
}

// SetUpdatedDate gets a reference to the given string and assigns it to the UpdatedDate field.
func (o *ListSources200ResponseResultsInner) SetUpdatedDate(v string) {
	o.UpdatedDate = &v
}

// GetWorksApiUrl returns the WorksApiUrl field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetWorksApiUrl() string {
	if o == nil || IsNil(o.WorksApiUrl) {
		var ret string
		return ret
	}
	return *o.WorksApiUrl
}

// GetWorksApiUrlOk returns a tuple with the WorksApiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetWorksApiUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WorksApiUrl) {
		return nil, false
	}
	return o.WorksApiUrl, true
}

// HasWorksApiUrl returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasWorksApiUrl() bool {
	if o != nil && !IsNil(o.WorksApiUrl) {
		return true
	}

	return false
}

// SetWorksApiUrl gets a reference to the given string and assigns it to the WorksApiUrl field.
func (o *ListSources200ResponseResultsInner) SetWorksApiUrl(v string) {
	o.WorksApiUrl = &v
}

// GetWorksCount returns the WorksCount field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetWorksCount() int32 {
	if o == nil || IsNil(o.WorksCount) {
		var ret int32
		return ret
	}
	return *o.WorksCount
}

// GetWorksCountOk returns a tuple with the WorksCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetWorksCountOk() (*int32, bool) {
	if o == nil || IsNil(o.WorksCount) {
		return nil, false
	}
	return o.WorksCount, true
}

// HasWorksCount returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasWorksCount() bool {
	if o != nil && !IsNil(o.WorksCount) {
		return true
	}

	return false
}

// SetWorksCount gets a reference to the given int32 and assigns it to the WorksCount field.
func (o *ListSources200ResponseResultsInner) SetWorksCount(v int32) {
	o.WorksCount = &v
}

// GetXConcepts returns the XConcepts field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInner) GetXConcepts() []map[string]interface{} {
	if o == nil || IsNil(o.XConcepts) {
		var ret []map[string]interface{}
		return ret
	}
	return o.XConcepts
}

// GetXConceptsOk returns a tuple with the XConcepts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInner) GetXConceptsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.XConcepts) {
		return nil, false
	}
	return o.XConcepts, true
}

// HasXConcepts returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInner) HasXConcepts() bool {
	if o != nil && !IsNil(o.XConcepts) {
		return true
	}

	return false
}

// SetXConcepts gets a reference to the given []map[string]interface{} and assigns it to the XConcepts field.
func (o *ListSources200ResponseResultsInner) SetXConcepts(v []map[string]interface{}) {
	o.XConcepts = v
}

func (o ListSources200ResponseResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSources200ResponseResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AbbreviatedTitle) {
		toSerialize["abbreviated_title"] = o.AbbreviatedTitle
	}
	if !IsNil(o.AlternateTitles) {
		toSerialize["alternate_titles"] = o.AlternateTitles
	}
	if !IsNil(o.ApcPrices) {
		toSerialize["apc_prices"] = o.ApcPrices
	}
	if !IsNil(o.ApcUsd) {
		toSerialize["apc_usd"] = o.ApcUsd
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
	if !IsNil(o.HomepageUrl) {
		toSerialize["homepage_url"] = o.HomepageUrl
	}
	if !IsNil(o.HostOrganization) {
		toSerialize["host_organization"] = o.HostOrganization
	}
	if !IsNil(o.HostOrganizationLineage) {
		toSerialize["host_organization_lineage"] = o.HostOrganizationLineage
	}
	if !IsNil(o.HostOrganizationName) {
		toSerialize["host_organization_name"] = o.HostOrganizationName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !IsNil(o.IsCore) {
		toSerialize["is_core"] = o.IsCore
	}
	if !IsNil(o.IsInDoaj) {
		toSerialize["is_in_doaj"] = o.IsInDoaj
	}
	if !IsNil(o.IsOa) {
		toSerialize["is_oa"] = o.IsOa
	}
	if !IsNil(o.Issn) {
		toSerialize["issn"] = o.Issn
	}
	if !IsNil(o.IssnL) {
		toSerialize["issn_l"] = o.IssnL
	}
	if !IsNil(o.Societies) {
		toSerialize["societies"] = o.Societies
	}
	if !IsNil(o.SummaryStats) {
		toSerialize["summary_stats"] = o.SummaryStats
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

type NullableListSources200ResponseResultsInner struct {
	value *ListSources200ResponseResultsInner
	isSet bool
}

func (v NullableListSources200ResponseResultsInner) Get() *ListSources200ResponseResultsInner {
	return v.value
}

func (v *NullableListSources200ResponseResultsInner) Set(val *ListSources200ResponseResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListSources200ResponseResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListSources200ResponseResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSources200ResponseResultsInner(val *ListSources200ResponseResultsInner) *NullableListSources200ResponseResultsInner {
	return &NullableListSources200ResponseResultsInner{value: val, isSet: true}
}

func (v NullableListSources200ResponseResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSources200ResponseResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


