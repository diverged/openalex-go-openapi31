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

// checks if the ListInstitutions200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListInstitutions200Response{}

// ListInstitutions200Response struct for ListInstitutions200Response
type ListInstitutions200Response struct {
	Meta *ListAuthors200ResponseMeta `json:"meta,omitempty"`
	Results []ListInstitutions200ResponseResultsInner `json:"results,omitempty"`
	GroupBy []ListAuthors200ResponseGroupByInner `json:"group_by,omitempty"`
}

// NewListInstitutions200Response instantiates a new ListInstitutions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInstitutions200Response() *ListInstitutions200Response {
	this := ListInstitutions200Response{}
	return &this
}

// NewListInstitutions200ResponseWithDefaults instantiates a new ListInstitutions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInstitutions200ResponseWithDefaults() *ListInstitutions200Response {
	this := ListInstitutions200Response{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListInstitutions200Response) GetMeta() ListAuthors200ResponseMeta {
	if o == nil || IsNil(o.Meta) {
		var ret ListAuthors200ResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200Response) GetMetaOk() (*ListAuthors200ResponseMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListInstitutions200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ListAuthors200ResponseMeta and assigns it to the Meta field.
func (o *ListInstitutions200Response) SetMeta(v ListAuthors200ResponseMeta) {
	o.Meta = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListInstitutions200Response) GetResults() []ListInstitutions200ResponseResultsInner {
	if o == nil || IsNil(o.Results) {
		var ret []ListInstitutions200ResponseResultsInner
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200Response) GetResultsOk() ([]ListInstitutions200ResponseResultsInner, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListInstitutions200Response) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ListInstitutions200ResponseResultsInner and assigns it to the Results field.
func (o *ListInstitutions200Response) SetResults(v []ListInstitutions200ResponseResultsInner) {
	o.Results = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *ListInstitutions200Response) GetGroupBy() []ListAuthors200ResponseGroupByInner {
	if o == nil || IsNil(o.GroupBy) {
		var ret []ListAuthors200ResponseGroupByInner
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200Response) GetGroupByOk() ([]ListAuthors200ResponseGroupByInner, bool) {
	if o == nil || IsNil(o.GroupBy) {
		return nil, false
	}
	return o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *ListInstitutions200Response) HasGroupBy() bool {
	if o != nil && !IsNil(o.GroupBy) {
		return true
	}

	return false
}

// SetGroupBy gets a reference to the given []ListAuthors200ResponseGroupByInner and assigns it to the GroupBy field.
func (o *ListInstitutions200Response) SetGroupBy(v []ListAuthors200ResponseGroupByInner) {
	o.GroupBy = v
}

func (o ListInstitutions200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListInstitutions200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.GroupBy) {
		toSerialize["group_by"] = o.GroupBy
	}
	return toSerialize, nil
}

type NullableListInstitutions200Response struct {
	value *ListInstitutions200Response
	isSet bool
}

func (v NullableListInstitutions200Response) Get() *ListInstitutions200Response {
	return v.value
}

func (v *NullableListInstitutions200Response) Set(val *ListInstitutions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListInstitutions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListInstitutions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInstitutions200Response(val *ListInstitutions200Response) *NullableListInstitutions200Response {
	return &NullableListInstitutions200Response{value: val, isSet: true}
}

func (v NullableListInstitutions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInstitutions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

