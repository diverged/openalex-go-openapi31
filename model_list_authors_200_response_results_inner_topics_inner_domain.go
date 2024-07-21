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

// checks if the ListAuthors200ResponseResultsInnerTopicsInnerDomain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAuthors200ResponseResultsInnerTopicsInnerDomain{}

// ListAuthors200ResponseResultsInnerTopicsInnerDomain struct for ListAuthors200ResponseResultsInnerTopicsInnerDomain
type ListAuthors200ResponseResultsInnerTopicsInnerDomain struct {
	DisplayName *string `json:"display_name,omitempty"`
	Id *string `json:"id,omitempty"`
}

// NewListAuthors200ResponseResultsInnerTopicsInnerDomain instantiates a new ListAuthors200ResponseResultsInnerTopicsInnerDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAuthors200ResponseResultsInnerTopicsInnerDomain() *ListAuthors200ResponseResultsInnerTopicsInnerDomain {
	this := ListAuthors200ResponseResultsInnerTopicsInnerDomain{}
	return &this
}

// NewListAuthors200ResponseResultsInnerTopicsInnerDomainWithDefaults instantiates a new ListAuthors200ResponseResultsInnerTopicsInnerDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAuthors200ResponseResultsInnerTopicsInnerDomainWithDefaults() *ListAuthors200ResponseResultsInnerTopicsInnerDomain {
	this := ListAuthors200ResponseResultsInnerTopicsInnerDomain{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ListAuthors200ResponseResultsInnerTopicsInnerDomain) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAuthors200ResponseResultsInnerTopicsInnerDomain) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ListAuthors200ResponseResultsInnerTopicsInnerDomain) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ListAuthors200ResponseResultsInnerTopicsInnerDomain) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListAuthors200ResponseResultsInnerTopicsInnerDomain) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAuthors200ResponseResultsInnerTopicsInnerDomain) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListAuthors200ResponseResultsInnerTopicsInnerDomain) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListAuthors200ResponseResultsInnerTopicsInnerDomain) SetId(v string) {
	o.Id = &v
}

func (o ListAuthors200ResponseResultsInnerTopicsInnerDomain) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAuthors200ResponseResultsInnerTopicsInnerDomain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableListAuthors200ResponseResultsInnerTopicsInnerDomain struct {
	value *ListAuthors200ResponseResultsInnerTopicsInnerDomain
	isSet bool
}

func (v NullableListAuthors200ResponseResultsInnerTopicsInnerDomain) Get() *ListAuthors200ResponseResultsInnerTopicsInnerDomain {
	return v.value
}

func (v *NullableListAuthors200ResponseResultsInnerTopicsInnerDomain) Set(val *ListAuthors200ResponseResultsInnerTopicsInnerDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableListAuthors200ResponseResultsInnerTopicsInnerDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableListAuthors200ResponseResultsInnerTopicsInnerDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAuthors200ResponseResultsInnerTopicsInnerDomain(val *ListAuthors200ResponseResultsInnerTopicsInnerDomain) *NullableListAuthors200ResponseResultsInnerTopicsInnerDomain {
	return &NullableListAuthors200ResponseResultsInnerTopicsInnerDomain{value: val, isSet: true}
}

func (v NullableListAuthors200ResponseResultsInnerTopicsInnerDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAuthors200ResponseResultsInnerTopicsInnerDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


