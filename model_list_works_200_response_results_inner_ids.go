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

// checks if the ListWorks200ResponseResultsInnerIds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListWorks200ResponseResultsInnerIds{}

// ListWorks200ResponseResultsInnerIds All the external identifiers that we know about for this work. IDs are expressed as URIs whenever possible.
type ListWorks200ResponseResultsInnerIds struct {
	// The DOI for this work.
	Doi *string `json:"doi,omitempty"`
	// The Microsoft Academic Graph ID for this work.
	Mag *string `json:"mag,omitempty"`
	// The OpenAlex ID for this work.
	Openalex *string `json:"openalex,omitempty"`
	// The PubMed Central ID for this work.
	Pmcid *string `json:"pmcid,omitempty"`
	// The PubMed ID for this work.
	Pmid *string `json:"pmid,omitempty"`
}

// NewListWorks200ResponseResultsInnerIds instantiates a new ListWorks200ResponseResultsInnerIds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWorks200ResponseResultsInnerIds() *ListWorks200ResponseResultsInnerIds {
	this := ListWorks200ResponseResultsInnerIds{}
	return &this
}

// NewListWorks200ResponseResultsInnerIdsWithDefaults instantiates a new ListWorks200ResponseResultsInnerIds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWorks200ResponseResultsInnerIdsWithDefaults() *ListWorks200ResponseResultsInnerIds {
	this := ListWorks200ResponseResultsInnerIds{}
	return &this
}

// GetDoi returns the Doi field value if set, zero value otherwise.
func (o *ListWorks200ResponseResultsInnerIds) GetDoi() string {
	if o == nil || IsNil(o.Doi) {
		var ret string
		return ret
	}
	return *o.Doi
}

// GetDoiOk returns a tuple with the Doi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorks200ResponseResultsInnerIds) GetDoiOk() (*string, bool) {
	if o == nil || IsNil(o.Doi) {
		return nil, false
	}
	return o.Doi, true
}

// HasDoi returns a boolean if a field has been set.
func (o *ListWorks200ResponseResultsInnerIds) HasDoi() bool {
	if o != nil && !IsNil(o.Doi) {
		return true
	}

	return false
}

// SetDoi gets a reference to the given string and assigns it to the Doi field.
func (o *ListWorks200ResponseResultsInnerIds) SetDoi(v string) {
	o.Doi = &v
}

// GetMag returns the Mag field value if set, zero value otherwise.
func (o *ListWorks200ResponseResultsInnerIds) GetMag() string {
	if o == nil || IsNil(o.Mag) {
		var ret string
		return ret
	}
	return *o.Mag
}

// GetMagOk returns a tuple with the Mag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorks200ResponseResultsInnerIds) GetMagOk() (*string, bool) {
	if o == nil || IsNil(o.Mag) {
		return nil, false
	}
	return o.Mag, true
}

// HasMag returns a boolean if a field has been set.
func (o *ListWorks200ResponseResultsInnerIds) HasMag() bool {
	if o != nil && !IsNil(o.Mag) {
		return true
	}

	return false
}

// SetMag gets a reference to the given string and assigns it to the Mag field.
func (o *ListWorks200ResponseResultsInnerIds) SetMag(v string) {
	o.Mag = &v
}

// GetOpenalex returns the Openalex field value if set, zero value otherwise.
func (o *ListWorks200ResponseResultsInnerIds) GetOpenalex() string {
	if o == nil || IsNil(o.Openalex) {
		var ret string
		return ret
	}
	return *o.Openalex
}

// GetOpenalexOk returns a tuple with the Openalex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorks200ResponseResultsInnerIds) GetOpenalexOk() (*string, bool) {
	if o == nil || IsNil(o.Openalex) {
		return nil, false
	}
	return o.Openalex, true
}

// HasOpenalex returns a boolean if a field has been set.
func (o *ListWorks200ResponseResultsInnerIds) HasOpenalex() bool {
	if o != nil && !IsNil(o.Openalex) {
		return true
	}

	return false
}

// SetOpenalex gets a reference to the given string and assigns it to the Openalex field.
func (o *ListWorks200ResponseResultsInnerIds) SetOpenalex(v string) {
	o.Openalex = &v
}

// GetPmcid returns the Pmcid field value if set, zero value otherwise.
func (o *ListWorks200ResponseResultsInnerIds) GetPmcid() string {
	if o == nil || IsNil(o.Pmcid) {
		var ret string
		return ret
	}
	return *o.Pmcid
}

// GetPmcidOk returns a tuple with the Pmcid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorks200ResponseResultsInnerIds) GetPmcidOk() (*string, bool) {
	if o == nil || IsNil(o.Pmcid) {
		return nil, false
	}
	return o.Pmcid, true
}

// HasPmcid returns a boolean if a field has been set.
func (o *ListWorks200ResponseResultsInnerIds) HasPmcid() bool {
	if o != nil && !IsNil(o.Pmcid) {
		return true
	}

	return false
}

// SetPmcid gets a reference to the given string and assigns it to the Pmcid field.
func (o *ListWorks200ResponseResultsInnerIds) SetPmcid(v string) {
	o.Pmcid = &v
}

// GetPmid returns the Pmid field value if set, zero value otherwise.
func (o *ListWorks200ResponseResultsInnerIds) GetPmid() string {
	if o == nil || IsNil(o.Pmid) {
		var ret string
		return ret
	}
	return *o.Pmid
}

// GetPmidOk returns a tuple with the Pmid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorks200ResponseResultsInnerIds) GetPmidOk() (*string, bool) {
	if o == nil || IsNil(o.Pmid) {
		return nil, false
	}
	return o.Pmid, true
}

// HasPmid returns a boolean if a field has been set.
func (o *ListWorks200ResponseResultsInnerIds) HasPmid() bool {
	if o != nil && !IsNil(o.Pmid) {
		return true
	}

	return false
}

// SetPmid gets a reference to the given string and assigns it to the Pmid field.
func (o *ListWorks200ResponseResultsInnerIds) SetPmid(v string) {
	o.Pmid = &v
}

func (o ListWorks200ResponseResultsInnerIds) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListWorks200ResponseResultsInnerIds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Doi) {
		toSerialize["doi"] = o.Doi
	}
	if !IsNil(o.Mag) {
		toSerialize["mag"] = o.Mag
	}
	if !IsNil(o.Openalex) {
		toSerialize["openalex"] = o.Openalex
	}
	if !IsNil(o.Pmcid) {
		toSerialize["pmcid"] = o.Pmcid
	}
	if !IsNil(o.Pmid) {
		toSerialize["pmid"] = o.Pmid
	}
	return toSerialize, nil
}

type NullableListWorks200ResponseResultsInnerIds struct {
	value *ListWorks200ResponseResultsInnerIds
	isSet bool
}

func (v NullableListWorks200ResponseResultsInnerIds) Get() *ListWorks200ResponseResultsInnerIds {
	return v.value
}

func (v *NullableListWorks200ResponseResultsInnerIds) Set(val *ListWorks200ResponseResultsInnerIds) {
	v.value = val
	v.isSet = true
}

func (v NullableListWorks200ResponseResultsInnerIds) IsSet() bool {
	return v.isSet
}

func (v *NullableListWorks200ResponseResultsInnerIds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWorks200ResponseResultsInnerIds(val *ListWorks200ResponseResultsInnerIds) *NullableListWorks200ResponseResultsInnerIds {
	return &NullableListWorks200ResponseResultsInnerIds{value: val, isSet: true}
}

func (v NullableListWorks200ResponseResultsInnerIds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWorks200ResponseResultsInnerIds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


