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

// checks if the ListSources200ResponseResultsInnerSocietiesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSources200ResponseResultsInnerSocietiesInner{}

// ListSources200ResponseResultsInnerSocietiesInner struct for ListSources200ResponseResultsInnerSocietiesInner
type ListSources200ResponseResultsInnerSocietiesInner struct {
	Url *string `json:"url,omitempty"`
	Organization *string `json:"organization,omitempty"`
}

// NewListSources200ResponseResultsInnerSocietiesInner instantiates a new ListSources200ResponseResultsInnerSocietiesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSources200ResponseResultsInnerSocietiesInner() *ListSources200ResponseResultsInnerSocietiesInner {
	this := ListSources200ResponseResultsInnerSocietiesInner{}
	return &this
}

// NewListSources200ResponseResultsInnerSocietiesInnerWithDefaults instantiates a new ListSources200ResponseResultsInnerSocietiesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSources200ResponseResultsInnerSocietiesInnerWithDefaults() *ListSources200ResponseResultsInnerSocietiesInner {
	this := ListSources200ResponseResultsInnerSocietiesInner{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInnerSocietiesInner) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInnerSocietiesInner) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInnerSocietiesInner) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ListSources200ResponseResultsInnerSocietiesInner) SetUrl(v string) {
	o.Url = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ListSources200ResponseResultsInnerSocietiesInner) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSources200ResponseResultsInnerSocietiesInner) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ListSources200ResponseResultsInnerSocietiesInner) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *ListSources200ResponseResultsInnerSocietiesInner) SetOrganization(v string) {
	o.Organization = &v
}

func (o ListSources200ResponseResultsInnerSocietiesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSources200ResponseResultsInnerSocietiesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	return toSerialize, nil
}

type NullableListSources200ResponseResultsInnerSocietiesInner struct {
	value *ListSources200ResponseResultsInnerSocietiesInner
	isSet bool
}

func (v NullableListSources200ResponseResultsInnerSocietiesInner) Get() *ListSources200ResponseResultsInnerSocietiesInner {
	return v.value
}

func (v *NullableListSources200ResponseResultsInnerSocietiesInner) Set(val *ListSources200ResponseResultsInnerSocietiesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListSources200ResponseResultsInnerSocietiesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListSources200ResponseResultsInnerSocietiesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSources200ResponseResultsInnerSocietiesInner(val *ListSources200ResponseResultsInnerSocietiesInner) *NullableListSources200ResponseResultsInnerSocietiesInner {
	return &NullableListSources200ResponseResultsInnerSocietiesInner{value: val, isSet: true}
}

func (v NullableListSources200ResponseResultsInnerSocietiesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSources200ResponseResultsInnerSocietiesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


