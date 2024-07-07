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

// checks if the AutocompleteMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutocompleteMeta{}

// AutocompleteMeta Information about the autocomplete request and response.
type AutocompleteMeta struct {
	// The total number of matching results.
	Count *int32 `json:"count,omitempty"`
	// The time taken by the database to respond, in milliseconds.
	DbResponseTimeMs *int32 `json:"db_response_time_ms,omitempty"`
	// The current page number (always 1 for autocomplete).
	Page *int32 `json:"page,omitempty"`
	// The number of results per page (always 10 for autocomplete).
	PerPage *int32 `json:"per_page,omitempty"`
}

// NewAutocompleteMeta instantiates a new AutocompleteMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutocompleteMeta() *AutocompleteMeta {
	this := AutocompleteMeta{}
	return &this
}

// NewAutocompleteMetaWithDefaults instantiates a new AutocompleteMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutocompleteMetaWithDefaults() *AutocompleteMeta {
	this := AutocompleteMeta{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *AutocompleteMeta) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutocompleteMeta) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *AutocompleteMeta) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *AutocompleteMeta) SetCount(v int32) {
	o.Count = &v
}

// GetDbResponseTimeMs returns the DbResponseTimeMs field value if set, zero value otherwise.
func (o *AutocompleteMeta) GetDbResponseTimeMs() int32 {
	if o == nil || IsNil(o.DbResponseTimeMs) {
		var ret int32
		return ret
	}
	return *o.DbResponseTimeMs
}

// GetDbResponseTimeMsOk returns a tuple with the DbResponseTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutocompleteMeta) GetDbResponseTimeMsOk() (*int32, bool) {
	if o == nil || IsNil(o.DbResponseTimeMs) {
		return nil, false
	}
	return o.DbResponseTimeMs, true
}

// HasDbResponseTimeMs returns a boolean if a field has been set.
func (o *AutocompleteMeta) HasDbResponseTimeMs() bool {
	if o != nil && !IsNil(o.DbResponseTimeMs) {
		return true
	}

	return false
}

// SetDbResponseTimeMs gets a reference to the given int32 and assigns it to the DbResponseTimeMs field.
func (o *AutocompleteMeta) SetDbResponseTimeMs(v int32) {
	o.DbResponseTimeMs = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *AutocompleteMeta) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutocompleteMeta) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *AutocompleteMeta) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *AutocompleteMeta) SetPage(v int32) {
	o.Page = &v
}

// GetPerPage returns the PerPage field value if set, zero value otherwise.
func (o *AutocompleteMeta) GetPerPage() int32 {
	if o == nil || IsNil(o.PerPage) {
		var ret int32
		return ret
	}
	return *o.PerPage
}

// GetPerPageOk returns a tuple with the PerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutocompleteMeta) GetPerPageOk() (*int32, bool) {
	if o == nil || IsNil(o.PerPage) {
		return nil, false
	}
	return o.PerPage, true
}

// HasPerPage returns a boolean if a field has been set.
func (o *AutocompleteMeta) HasPerPage() bool {
	if o != nil && !IsNil(o.PerPage) {
		return true
	}

	return false
}

// SetPerPage gets a reference to the given int32 and assigns it to the PerPage field.
func (o *AutocompleteMeta) SetPerPage(v int32) {
	o.PerPage = &v
}

func (o AutocompleteMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutocompleteMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.DbResponseTimeMs) {
		toSerialize["db_response_time_ms"] = o.DbResponseTimeMs
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.PerPage) {
		toSerialize["per_page"] = o.PerPage
	}
	return toSerialize, nil
}

type NullableAutocompleteMeta struct {
	value *AutocompleteMeta
	isSet bool
}

func (v NullableAutocompleteMeta) Get() *AutocompleteMeta {
	return v.value
}

func (v *NullableAutocompleteMeta) Set(val *AutocompleteMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableAutocompleteMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableAutocompleteMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutocompleteMeta(val *AutocompleteMeta) *NullableAutocompleteMeta {
	return &NullableAutocompleteMeta{value: val, isSet: true}
}

func (v NullableAutocompleteMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutocompleteMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


