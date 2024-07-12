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

// checks if the ListKeywords200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListKeywords200Response{}

// ListKeywords200Response struct for ListKeywords200Response
type ListKeywords200Response struct {
	Meta *ListAuthors200ResponseMeta `json:"meta,omitempty"`
	Results []ListKeywords200ResponseResultsInner `json:"results,omitempty"`
	GroupBy []ListAuthors200ResponseGroupByInner `json:"group_by,omitempty"`
}

// NewListKeywords200Response instantiates a new ListKeywords200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListKeywords200Response() *ListKeywords200Response {
	this := ListKeywords200Response{}
	return &this
}

// NewListKeywords200ResponseWithDefaults instantiates a new ListKeywords200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListKeywords200ResponseWithDefaults() *ListKeywords200Response {
	this := ListKeywords200Response{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListKeywords200Response) GetMeta() ListAuthors200ResponseMeta {
	if o == nil || IsNil(o.Meta) {
		var ret ListAuthors200ResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListKeywords200Response) GetMetaOk() (*ListAuthors200ResponseMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListKeywords200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ListAuthors200ResponseMeta and assigns it to the Meta field.
func (o *ListKeywords200Response) SetMeta(v ListAuthors200ResponseMeta) {
	o.Meta = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListKeywords200Response) GetResults() []ListKeywords200ResponseResultsInner {
	if o == nil || IsNil(o.Results) {
		var ret []ListKeywords200ResponseResultsInner
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListKeywords200Response) GetResultsOk() ([]ListKeywords200ResponseResultsInner, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListKeywords200Response) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ListKeywords200ResponseResultsInner and assigns it to the Results field.
func (o *ListKeywords200Response) SetResults(v []ListKeywords200ResponseResultsInner) {
	o.Results = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *ListKeywords200Response) GetGroupBy() []ListAuthors200ResponseGroupByInner {
	if o == nil || IsNil(o.GroupBy) {
		var ret []ListAuthors200ResponseGroupByInner
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListKeywords200Response) GetGroupByOk() ([]ListAuthors200ResponseGroupByInner, bool) {
	if o == nil || IsNil(o.GroupBy) {
		return nil, false
	}
	return o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *ListKeywords200Response) HasGroupBy() bool {
	if o != nil && !IsNil(o.GroupBy) {
		return true
	}

	return false
}

// SetGroupBy gets a reference to the given []ListAuthors200ResponseGroupByInner and assigns it to the GroupBy field.
func (o *ListKeywords200Response) SetGroupBy(v []ListAuthors200ResponseGroupByInner) {
	o.GroupBy = v
}

func (o ListKeywords200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListKeywords200Response) ToMap() (map[string]interface{}, error) {
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

type NullableListKeywords200Response struct {
	value *ListKeywords200Response
	isSet bool
}

func (v NullableListKeywords200Response) Get() *ListKeywords200Response {
	return v.value
}

func (v *NullableListKeywords200Response) Set(val *ListKeywords200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListKeywords200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListKeywords200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListKeywords200Response(val *ListKeywords200Response) *NullableListKeywords200Response {
	return &NullableListKeywords200Response{value: val, isSet: true}
}

func (v NullableListKeywords200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListKeywords200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

