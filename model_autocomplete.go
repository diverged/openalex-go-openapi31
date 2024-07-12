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

// checks if the Autocomplete type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Autocomplete{}

// Autocomplete struct for Autocomplete
type Autocomplete struct {
	Meta *AutocompleteMeta `json:"meta,omitempty"`
	// A list of up to ten autocomplete results for the query, sorted by citation count.
	Results []AutocompleteResultsInner `json:"results,omitempty"`
}

// NewAutocomplete instantiates a new Autocomplete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutocomplete() *Autocomplete {
	this := Autocomplete{}
	return &this
}

// NewAutocompleteWithDefaults instantiates a new Autocomplete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutocompleteWithDefaults() *Autocomplete {
	this := Autocomplete{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Autocomplete) GetMeta() AutocompleteMeta {
	if o == nil || IsNil(o.Meta) {
		var ret AutocompleteMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Autocomplete) GetMetaOk() (*AutocompleteMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Autocomplete) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given AutocompleteMeta and assigns it to the Meta field.
func (o *Autocomplete) SetMeta(v AutocompleteMeta) {
	o.Meta = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *Autocomplete) GetResults() []AutocompleteResultsInner {
	if o == nil || IsNil(o.Results) {
		var ret []AutocompleteResultsInner
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Autocomplete) GetResultsOk() ([]AutocompleteResultsInner, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *Autocomplete) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []AutocompleteResultsInner and assigns it to the Results field.
func (o *Autocomplete) SetResults(v []AutocompleteResultsInner) {
	o.Results = v
}

func (o Autocomplete) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Autocomplete) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableAutocomplete struct {
	value *Autocomplete
	isSet bool
}

func (v NullableAutocomplete) Get() *Autocomplete {
	return v.value
}

func (v *NullableAutocomplete) Set(val *Autocomplete) {
	v.value = val
	v.isSet = true
}

func (v NullableAutocomplete) IsSet() bool {
	return v.isSet
}

func (v *NullableAutocomplete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutocomplete(val *Autocomplete) *NullableAutocomplete {
	return &NullableAutocomplete{value: val, isSet: true}
}

func (v NullableAutocomplete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutocomplete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

