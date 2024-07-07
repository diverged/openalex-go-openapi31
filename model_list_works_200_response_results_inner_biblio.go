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

// checks if the ListWorks200ResponseResultsInnerBiblio type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListWorks200ResponseResultsInnerBiblio{}

// ListWorks200ResponseResultsInnerBiblio Old-timey bibliographic info for this work.
type ListWorks200ResponseResultsInnerBiblio struct {
	// The first page number for this work.
	FirstPage *string `json:"first_page,omitempty"`
	// The issue number for this work.
	Issue *string `json:"issue,omitempty"`
	// The last page number for this work.
	LastPage *string `json:"last_page,omitempty"`
	// The volume number for this work.
	Volume *string `json:"volume,omitempty"`
}

// NewListWorks200ResponseResultsInnerBiblio instantiates a new ListWorks200ResponseResultsInnerBiblio object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWorks200ResponseResultsInnerBiblio() *ListWorks200ResponseResultsInnerBiblio {
	this := ListWorks200ResponseResultsInnerBiblio{}
	return &this
}

// NewListWorks200ResponseResultsInnerBiblioWithDefaults instantiates a new ListWorks200ResponseResultsInnerBiblio object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWorks200ResponseResultsInnerBiblioWithDefaults() *ListWorks200ResponseResultsInnerBiblio {
	this := ListWorks200ResponseResultsInnerBiblio{}
	return &this
}

// GetFirstPage returns the FirstPage field value if set, zero value otherwise.
func (o *ListWorks200ResponseResultsInnerBiblio) GetFirstPage() string {
	if o == nil || IsNil(o.FirstPage) {
		var ret string
		return ret
	}
	return *o.FirstPage
}

// GetFirstPageOk returns a tuple with the FirstPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorks200ResponseResultsInnerBiblio) GetFirstPageOk() (*string, bool) {
	if o == nil || IsNil(o.FirstPage) {
		return nil, false
	}
	return o.FirstPage, true
}

// HasFirstPage returns a boolean if a field has been set.
func (o *ListWorks200ResponseResultsInnerBiblio) HasFirstPage() bool {
	if o != nil && !IsNil(o.FirstPage) {
		return true
	}

	return false
}

// SetFirstPage gets a reference to the given string and assigns it to the FirstPage field.
func (o *ListWorks200ResponseResultsInnerBiblio) SetFirstPage(v string) {
	o.FirstPage = &v
}

// GetIssue returns the Issue field value if set, zero value otherwise.
func (o *ListWorks200ResponseResultsInnerBiblio) GetIssue() string {
	if o == nil || IsNil(o.Issue) {
		var ret string
		return ret
	}
	return *o.Issue
}

// GetIssueOk returns a tuple with the Issue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorks200ResponseResultsInnerBiblio) GetIssueOk() (*string, bool) {
	if o == nil || IsNil(o.Issue) {
		return nil, false
	}
	return o.Issue, true
}

// HasIssue returns a boolean if a field has been set.
func (o *ListWorks200ResponseResultsInnerBiblio) HasIssue() bool {
	if o != nil && !IsNil(o.Issue) {
		return true
	}

	return false
}

// SetIssue gets a reference to the given string and assigns it to the Issue field.
func (o *ListWorks200ResponseResultsInnerBiblio) SetIssue(v string) {
	o.Issue = &v
}

// GetLastPage returns the LastPage field value if set, zero value otherwise.
func (o *ListWorks200ResponseResultsInnerBiblio) GetLastPage() string {
	if o == nil || IsNil(o.LastPage) {
		var ret string
		return ret
	}
	return *o.LastPage
}

// GetLastPageOk returns a tuple with the LastPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorks200ResponseResultsInnerBiblio) GetLastPageOk() (*string, bool) {
	if o == nil || IsNil(o.LastPage) {
		return nil, false
	}
	return o.LastPage, true
}

// HasLastPage returns a boolean if a field has been set.
func (o *ListWorks200ResponseResultsInnerBiblio) HasLastPage() bool {
	if o != nil && !IsNil(o.LastPage) {
		return true
	}

	return false
}

// SetLastPage gets a reference to the given string and assigns it to the LastPage field.
func (o *ListWorks200ResponseResultsInnerBiblio) SetLastPage(v string) {
	o.LastPage = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *ListWorks200ResponseResultsInnerBiblio) GetVolume() string {
	if o == nil || IsNil(o.Volume) {
		var ret string
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorks200ResponseResultsInnerBiblio) GetVolumeOk() (*string, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *ListWorks200ResponseResultsInnerBiblio) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given string and assigns it to the Volume field.
func (o *ListWorks200ResponseResultsInnerBiblio) SetVolume(v string) {
	o.Volume = &v
}

func (o ListWorks200ResponseResultsInnerBiblio) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListWorks200ResponseResultsInnerBiblio) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstPage) {
		toSerialize["first_page"] = o.FirstPage
	}
	if !IsNil(o.Issue) {
		toSerialize["issue"] = o.Issue
	}
	if !IsNil(o.LastPage) {
		toSerialize["last_page"] = o.LastPage
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	return toSerialize, nil
}

type NullableListWorks200ResponseResultsInnerBiblio struct {
	value *ListWorks200ResponseResultsInnerBiblio
	isSet bool
}

func (v NullableListWorks200ResponseResultsInnerBiblio) Get() *ListWorks200ResponseResultsInnerBiblio {
	return v.value
}

func (v *NullableListWorks200ResponseResultsInnerBiblio) Set(val *ListWorks200ResponseResultsInnerBiblio) {
	v.value = val
	v.isSet = true
}

func (v NullableListWorks200ResponseResultsInnerBiblio) IsSet() bool {
	return v.isSet
}

func (v *NullableListWorks200ResponseResultsInnerBiblio) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWorks200ResponseResultsInnerBiblio(val *ListWorks200ResponseResultsInnerBiblio) *NullableListWorks200ResponseResultsInnerBiblio {
	return &NullableListWorks200ResponseResultsInnerBiblio{value: val, isSet: true}
}

func (v NullableListWorks200ResponseResultsInnerBiblio) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWorks200ResponseResultsInnerBiblio) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


