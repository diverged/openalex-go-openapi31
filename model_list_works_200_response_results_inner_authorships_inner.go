/*
OpenAlex API

An OpenAPI specification for the OpenAlex API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ListWorks200ResponseResultsInnerAuthorshipsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListWorks200ResponseResultsInnerAuthorshipsInner{}

// ListWorks200ResponseResultsInnerAuthorshipsInner struct for ListWorks200ResponseResultsInnerAuthorshipsInner
type ListWorks200ResponseResultsInnerAuthorshipsInner struct {
	// A summarized description of this author's position in the work's author list.
	AuthorPosition string `json:"author_position"`
	Author ListWorks200ResponseResultsInnerAuthorshipsInnerAuthor `json:"author"`
	// The institutional affiliations this author claimed in the context of this work, as dehydrated Institution objects.
	Institutions []ListAuthors200ResponseResultsInnerAffiliationsInnerInstitution `json:"institutions"`
	// If true, this is a corresponding author for this work.
	IsCorresponding *bool `json:"is_corresponding,omitempty"`
	// This author's affiliation as it originally came to us, as a list of raw unformatted strings.
	RawAffiliationStrings []string `json:"raw_affiliation_strings"`
	// This author's name as it originally came to us, as a raw unformatted string.
	RawAuthorName string `json:"raw_author_name"`
	// The country or countries for this author, determined using a combination of matched institutions and parsing of the raw affiliation strings.
	Countries []string `json:"countries,omitempty"`
}

type _ListWorks200ResponseResultsInnerAuthorshipsInner ListWorks200ResponseResultsInnerAuthorshipsInner

// NewListWorks200ResponseResultsInnerAuthorshipsInner instantiates a new ListWorks200ResponseResultsInnerAuthorshipsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWorks200ResponseResultsInnerAuthorshipsInner(authorPosition string, author ListWorks200ResponseResultsInnerAuthorshipsInnerAuthor, institutions []ListAuthors200ResponseResultsInnerAffiliationsInnerInstitution, rawAffiliationStrings []string, rawAuthorName string) *ListWorks200ResponseResultsInnerAuthorshipsInner {
	this := ListWorks200ResponseResultsInnerAuthorshipsInner{}
	this.AuthorPosition = authorPosition
	this.Author = author
	this.Institutions = institutions
	this.RawAffiliationStrings = rawAffiliationStrings
	this.RawAuthorName = rawAuthorName
	return &this
}

// NewListWorks200ResponseResultsInnerAuthorshipsInnerWithDefaults instantiates a new ListWorks200ResponseResultsInnerAuthorshipsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWorks200ResponseResultsInnerAuthorshipsInnerWithDefaults() *ListWorks200ResponseResultsInnerAuthorshipsInner {
	this := ListWorks200ResponseResultsInnerAuthorshipsInner{}
	return &this
}

// GetAuthorPosition returns the AuthorPosition field value
func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetAuthorPosition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorPosition
}

// GetAuthorPositionOk returns a tuple with the AuthorPosition field value
// and a boolean to check if the value has been set.
func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetAuthorPositionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorPosition, true
}

// SetAuthorPosition sets field value
func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) SetAuthorPosition(v string) {
	o.AuthorPosition = v
}

// GetAuthor returns the Author field value
func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetAuthor() ListWorks200ResponseResultsInnerAuthorshipsInnerAuthor {
	if o == nil {
		var ret ListWorks200ResponseResultsInnerAuthorshipsInnerAuthor
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetAuthorOk() (*ListWorks200ResponseResultsInnerAuthorshipsInnerAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) SetAuthor(v ListWorks200ResponseResultsInnerAuthorshipsInnerAuthor) {
	o.Author = v
}

// GetInstitutions returns the Institutions field value
func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetInstitutions() []ListAuthors200ResponseResultsInnerAffiliationsInnerInstitution {
	if o == nil {
		var ret []ListAuthors200ResponseResultsInnerAffiliationsInnerInstitution
		return ret
	}

	return o.Institutions
}

// GetInstitutionsOk returns a tuple with the Institutions field value
// and a boolean to check if the value has been set.
func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetInstitutionsOk() ([]ListAuthors200ResponseResultsInnerAffiliationsInnerInstitution, bool) {
	if o == nil {
		return nil, false
	}
	return o.Institutions, true
}

// SetInstitutions sets field value
func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) SetInstitutions(v []ListAuthors200ResponseResultsInnerAffiliationsInnerInstitution) {
	o.Institutions = v
}

// GetIsCorresponding returns the IsCorresponding field value if set, zero value otherwise.
func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetIsCorresponding() bool {
	if o == nil || IsNil(o.IsCorresponding) {
		var ret bool
		return ret
	}
	return *o.IsCorresponding
}

// GetIsCorrespondingOk returns a tuple with the IsCorresponding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetIsCorrespondingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCorresponding) {
		return nil, false
	}
	return o.IsCorresponding, true
}

// HasIsCorresponding returns a boolean if a field has been set.
func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) HasIsCorresponding() bool {
	if o != nil && !IsNil(o.IsCorresponding) {
		return true
	}

	return false
}

// SetIsCorresponding gets a reference to the given bool and assigns it to the IsCorresponding field.
func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) SetIsCorresponding(v bool) {
	o.IsCorresponding = &v
}

// GetRawAffiliationStrings returns the RawAffiliationStrings field value
func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetRawAffiliationStrings() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RawAffiliationStrings
}

// GetRawAffiliationStringsOk returns a tuple with the RawAffiliationStrings field value
// and a boolean to check if the value has been set.
func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetRawAffiliationStringsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RawAffiliationStrings, true
}

// SetRawAffiliationStrings sets field value
func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) SetRawAffiliationStrings(v []string) {
	o.RawAffiliationStrings = v
}

// GetRawAuthorName returns the RawAuthorName field value
func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetRawAuthorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RawAuthorName
}

// GetRawAuthorNameOk returns a tuple with the RawAuthorName field value
// and a boolean to check if the value has been set.
func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetRawAuthorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RawAuthorName, true
}

// SetRawAuthorName sets field value
func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) SetRawAuthorName(v string) {
	o.RawAuthorName = v
}

// GetCountries returns the Countries field value if set, zero value otherwise.
func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetCountries() []string {
	if o == nil || IsNil(o.Countries) {
		var ret []string
		return ret
	}
	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetCountriesOk() ([]string, bool) {
	if o == nil || IsNil(o.Countries) {
		return nil, false
	}
	return o.Countries, true
}

// HasCountries returns a boolean if a field has been set.
func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) HasCountries() bool {
	if o != nil && !IsNil(o.Countries) {
		return true
	}

	return false
}

// SetCountries gets a reference to the given []string and assigns it to the Countries field.
func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) SetCountries(v []string) {
	o.Countries = v
}

func (o ListWorks200ResponseResultsInnerAuthorshipsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListWorks200ResponseResultsInnerAuthorshipsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["author_position"] = o.AuthorPosition
	toSerialize["author"] = o.Author
	toSerialize["institutions"] = o.Institutions
	if !IsNil(o.IsCorresponding) {
		toSerialize["is_corresponding"] = o.IsCorresponding
	}
	toSerialize["raw_affiliation_strings"] = o.RawAffiliationStrings
	toSerialize["raw_author_name"] = o.RawAuthorName
	if !IsNil(o.Countries) {
		toSerialize["countries"] = o.Countries
	}
	return toSerialize, nil
}

func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"author_position",
		"author",
		"institutions",
		"raw_affiliation_strings",
		"raw_author_name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varListWorks200ResponseResultsInnerAuthorshipsInner := _ListWorks200ResponseResultsInnerAuthorshipsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListWorks200ResponseResultsInnerAuthorshipsInner)

	if err != nil {
		return err
	}

	*o = ListWorks200ResponseResultsInnerAuthorshipsInner(varListWorks200ResponseResultsInnerAuthorshipsInner)

	return err
}

type NullableListWorks200ResponseResultsInnerAuthorshipsInner struct {
	value *ListWorks200ResponseResultsInnerAuthorshipsInner
	isSet bool
}

func (v NullableListWorks200ResponseResultsInnerAuthorshipsInner) Get() *ListWorks200ResponseResultsInnerAuthorshipsInner {
	return v.value
}

func (v *NullableListWorks200ResponseResultsInnerAuthorshipsInner) Set(val *ListWorks200ResponseResultsInnerAuthorshipsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListWorks200ResponseResultsInnerAuthorshipsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListWorks200ResponseResultsInnerAuthorshipsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWorks200ResponseResultsInnerAuthorshipsInner(val *ListWorks200ResponseResultsInnerAuthorshipsInner) *NullableListWorks200ResponseResultsInnerAuthorshipsInner {
	return &NullableListWorks200ResponseResultsInnerAuthorshipsInner{value: val, isSet: true}
}

func (v NullableListWorks200ResponseResultsInnerAuthorshipsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWorks200ResponseResultsInnerAuthorshipsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


