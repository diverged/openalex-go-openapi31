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

// checks if the ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner{}

// ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner struct for ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner
type ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner struct {
	RawAffiliationString *string `json:"raw_affiliation_string,omitempty"`
	InstitutionIds []string `json:"institution_ids,omitempty"`
}

// NewListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner instantiates a new ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner() *ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner {
	this := ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner{}
	return &this
}

// NewListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInnerWithDefaults instantiates a new ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInnerWithDefaults() *ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner {
	this := ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner{}
	return &this
}

// GetRawAffiliationString returns the RawAffiliationString field value if set, zero value otherwise.
func (o *ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner) GetRawAffiliationString() string {
	if o == nil || IsNil(o.RawAffiliationString) {
		var ret string
		return ret
	}
	return *o.RawAffiliationString
}

// GetRawAffiliationStringOk returns a tuple with the RawAffiliationString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner) GetRawAffiliationStringOk() (*string, bool) {
	if o == nil || IsNil(o.RawAffiliationString) {
		return nil, false
	}
	return o.RawAffiliationString, true
}

// HasRawAffiliationString returns a boolean if a field has been set.
func (o *ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner) HasRawAffiliationString() bool {
	if o != nil && !IsNil(o.RawAffiliationString) {
		return true
	}

	return false
}

// SetRawAffiliationString gets a reference to the given string and assigns it to the RawAffiliationString field.
func (o *ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner) SetRawAffiliationString(v string) {
	o.RawAffiliationString = &v
}

// GetInstitutionIds returns the InstitutionIds field value if set, zero value otherwise.
func (o *ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner) GetInstitutionIds() []string {
	if o == nil || IsNil(o.InstitutionIds) {
		var ret []string
		return ret
	}
	return o.InstitutionIds
}

// GetInstitutionIdsOk returns a tuple with the InstitutionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner) GetInstitutionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.InstitutionIds) {
		return nil, false
	}
	return o.InstitutionIds, true
}

// HasInstitutionIds returns a boolean if a field has been set.
func (o *ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner) HasInstitutionIds() bool {
	if o != nil && !IsNil(o.InstitutionIds) {
		return true
	}

	return false
}

// SetInstitutionIds gets a reference to the given []string and assigns it to the InstitutionIds field.
func (o *ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner) SetInstitutionIds(v []string) {
	o.InstitutionIds = v
}

func (o ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RawAffiliationString) {
		toSerialize["raw_affiliation_string"] = o.RawAffiliationString
	}
	if !IsNil(o.InstitutionIds) {
		toSerialize["institution_ids"] = o.InstitutionIds
	}
	return toSerialize, nil
}

type NullableListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner struct {
	value *ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner
	isSet bool
}

func (v NullableListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner) Get() *ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner {
	return v.value
}

func (v *NullableListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner) Set(val *ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner(val *ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner) *NullableListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner {
	return &NullableListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner{value: val, isSet: true}
}

func (v NullableListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


