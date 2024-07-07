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

// checks if the ListInstitutions200ResponseResultsInnerIds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListInstitutions200ResponseResultsInnerIds{}

// ListInstitutions200ResponseResultsInnerIds External identifiers for this institution.
type ListInstitutions200ResponseResultsInnerIds struct {
	Openalex *string `json:"openalex,omitempty"`
	Ror *string `json:"ror,omitempty"`
	Grid *string `json:"grid,omitempty"`
	Wikipedia *string `json:"wikipedia,omitempty"`
	Wikidata *string `json:"wikidata,omitempty"`
	Mag *int32 `json:"mag,omitempty"`
}

// NewListInstitutions200ResponseResultsInnerIds instantiates a new ListInstitutions200ResponseResultsInnerIds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInstitutions200ResponseResultsInnerIds() *ListInstitutions200ResponseResultsInnerIds {
	this := ListInstitutions200ResponseResultsInnerIds{}
	return &this
}

// NewListInstitutions200ResponseResultsInnerIdsWithDefaults instantiates a new ListInstitutions200ResponseResultsInnerIds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInstitutions200ResponseResultsInnerIdsWithDefaults() *ListInstitutions200ResponseResultsInnerIds {
	this := ListInstitutions200ResponseResultsInnerIds{}
	return &this
}

// GetOpenalex returns the Openalex field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInnerIds) GetOpenalex() string {
	if o == nil || IsNil(o.Openalex) {
		var ret string
		return ret
	}
	return *o.Openalex
}

// GetOpenalexOk returns a tuple with the Openalex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInnerIds) GetOpenalexOk() (*string, bool) {
	if o == nil || IsNil(o.Openalex) {
		return nil, false
	}
	return o.Openalex, true
}

// HasOpenalex returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInnerIds) HasOpenalex() bool {
	if o != nil && !IsNil(o.Openalex) {
		return true
	}

	return false
}

// SetOpenalex gets a reference to the given string and assigns it to the Openalex field.
func (o *ListInstitutions200ResponseResultsInnerIds) SetOpenalex(v string) {
	o.Openalex = &v
}

// GetRor returns the Ror field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInnerIds) GetRor() string {
	if o == nil || IsNil(o.Ror) {
		var ret string
		return ret
	}
	return *o.Ror
}

// GetRorOk returns a tuple with the Ror field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInnerIds) GetRorOk() (*string, bool) {
	if o == nil || IsNil(o.Ror) {
		return nil, false
	}
	return o.Ror, true
}

// HasRor returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInnerIds) HasRor() bool {
	if o != nil && !IsNil(o.Ror) {
		return true
	}

	return false
}

// SetRor gets a reference to the given string and assigns it to the Ror field.
func (o *ListInstitutions200ResponseResultsInnerIds) SetRor(v string) {
	o.Ror = &v
}

// GetGrid returns the Grid field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInnerIds) GetGrid() string {
	if o == nil || IsNil(o.Grid) {
		var ret string
		return ret
	}
	return *o.Grid
}

// GetGridOk returns a tuple with the Grid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInnerIds) GetGridOk() (*string, bool) {
	if o == nil || IsNil(o.Grid) {
		return nil, false
	}
	return o.Grid, true
}

// HasGrid returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInnerIds) HasGrid() bool {
	if o != nil && !IsNil(o.Grid) {
		return true
	}

	return false
}

// SetGrid gets a reference to the given string and assigns it to the Grid field.
func (o *ListInstitutions200ResponseResultsInnerIds) SetGrid(v string) {
	o.Grid = &v
}

// GetWikipedia returns the Wikipedia field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInnerIds) GetWikipedia() string {
	if o == nil || IsNil(o.Wikipedia) {
		var ret string
		return ret
	}
	return *o.Wikipedia
}

// GetWikipediaOk returns a tuple with the Wikipedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInnerIds) GetWikipediaOk() (*string, bool) {
	if o == nil || IsNil(o.Wikipedia) {
		return nil, false
	}
	return o.Wikipedia, true
}

// HasWikipedia returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInnerIds) HasWikipedia() bool {
	if o != nil && !IsNil(o.Wikipedia) {
		return true
	}

	return false
}

// SetWikipedia gets a reference to the given string and assigns it to the Wikipedia field.
func (o *ListInstitutions200ResponseResultsInnerIds) SetWikipedia(v string) {
	o.Wikipedia = &v
}

// GetWikidata returns the Wikidata field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInnerIds) GetWikidata() string {
	if o == nil || IsNil(o.Wikidata) {
		var ret string
		return ret
	}
	return *o.Wikidata
}

// GetWikidataOk returns a tuple with the Wikidata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInnerIds) GetWikidataOk() (*string, bool) {
	if o == nil || IsNil(o.Wikidata) {
		return nil, false
	}
	return o.Wikidata, true
}

// HasWikidata returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInnerIds) HasWikidata() bool {
	if o != nil && !IsNil(o.Wikidata) {
		return true
	}

	return false
}

// SetWikidata gets a reference to the given string and assigns it to the Wikidata field.
func (o *ListInstitutions200ResponseResultsInnerIds) SetWikidata(v string) {
	o.Wikidata = &v
}

// GetMag returns the Mag field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInnerIds) GetMag() int32 {
	if o == nil || IsNil(o.Mag) {
		var ret int32
		return ret
	}
	return *o.Mag
}

// GetMagOk returns a tuple with the Mag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInnerIds) GetMagOk() (*int32, bool) {
	if o == nil || IsNil(o.Mag) {
		return nil, false
	}
	return o.Mag, true
}

// HasMag returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInnerIds) HasMag() bool {
	if o != nil && !IsNil(o.Mag) {
		return true
	}

	return false
}

// SetMag gets a reference to the given int32 and assigns it to the Mag field.
func (o *ListInstitutions200ResponseResultsInnerIds) SetMag(v int32) {
	o.Mag = &v
}

func (o ListInstitutions200ResponseResultsInnerIds) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListInstitutions200ResponseResultsInnerIds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Openalex) {
		toSerialize["openalex"] = o.Openalex
	}
	if !IsNil(o.Ror) {
		toSerialize["ror"] = o.Ror
	}
	if !IsNil(o.Grid) {
		toSerialize["grid"] = o.Grid
	}
	if !IsNil(o.Wikipedia) {
		toSerialize["wikipedia"] = o.Wikipedia
	}
	if !IsNil(o.Wikidata) {
		toSerialize["wikidata"] = o.Wikidata
	}
	if !IsNil(o.Mag) {
		toSerialize["mag"] = o.Mag
	}
	return toSerialize, nil
}

type NullableListInstitutions200ResponseResultsInnerIds struct {
	value *ListInstitutions200ResponseResultsInnerIds
	isSet bool
}

func (v NullableListInstitutions200ResponseResultsInnerIds) Get() *ListInstitutions200ResponseResultsInnerIds {
	return v.value
}

func (v *NullableListInstitutions200ResponseResultsInnerIds) Set(val *ListInstitutions200ResponseResultsInnerIds) {
	v.value = val
	v.isSet = true
}

func (v NullableListInstitutions200ResponseResultsInnerIds) IsSet() bool {
	return v.isSet
}

func (v *NullableListInstitutions200ResponseResultsInnerIds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInstitutions200ResponseResultsInnerIds(val *ListInstitutions200ResponseResultsInnerIds) *NullableListInstitutions200ResponseResultsInnerIds {
	return &NullableListInstitutions200ResponseResultsInnerIds{value: val, isSet: true}
}

func (v NullableListInstitutions200ResponseResultsInnerIds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInstitutions200ResponseResultsInnerIds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


