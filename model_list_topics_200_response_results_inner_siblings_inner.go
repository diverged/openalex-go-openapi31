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

// checks if the ListTopics200ResponseResultsInnerSiblingsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTopics200ResponseResultsInnerSiblingsInner{}

// ListTopics200ResponseResultsInnerSiblingsInner struct for ListTopics200ResponseResultsInnerSiblingsInner
type ListTopics200ResponseResultsInnerSiblingsInner struct {
	// The ID of the sibling topic.
	Id *string `json:"id,omitempty"`
	// The name of the sibling topic.
	DisplayName *string `json:"display_name,omitempty"`
}

// NewListTopics200ResponseResultsInnerSiblingsInner instantiates a new ListTopics200ResponseResultsInnerSiblingsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTopics200ResponseResultsInnerSiblingsInner() *ListTopics200ResponseResultsInnerSiblingsInner {
	this := ListTopics200ResponseResultsInnerSiblingsInner{}
	return &this
}

// NewListTopics200ResponseResultsInnerSiblingsInnerWithDefaults instantiates a new ListTopics200ResponseResultsInnerSiblingsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTopics200ResponseResultsInnerSiblingsInnerWithDefaults() *ListTopics200ResponseResultsInnerSiblingsInner {
	this := ListTopics200ResponseResultsInnerSiblingsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListTopics200ResponseResultsInnerSiblingsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTopics200ResponseResultsInnerSiblingsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListTopics200ResponseResultsInnerSiblingsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListTopics200ResponseResultsInnerSiblingsInner) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ListTopics200ResponseResultsInnerSiblingsInner) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTopics200ResponseResultsInnerSiblingsInner) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ListTopics200ResponseResultsInnerSiblingsInner) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ListTopics200ResponseResultsInnerSiblingsInner) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o ListTopics200ResponseResultsInnerSiblingsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTopics200ResponseResultsInnerSiblingsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	return toSerialize, nil
}

type NullableListTopics200ResponseResultsInnerSiblingsInner struct {
	value *ListTopics200ResponseResultsInnerSiblingsInner
	isSet bool
}

func (v NullableListTopics200ResponseResultsInnerSiblingsInner) Get() *ListTopics200ResponseResultsInnerSiblingsInner {
	return v.value
}

func (v *NullableListTopics200ResponseResultsInnerSiblingsInner) Set(val *ListTopics200ResponseResultsInnerSiblingsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListTopics200ResponseResultsInnerSiblingsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListTopics200ResponseResultsInnerSiblingsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTopics200ResponseResultsInnerSiblingsInner(val *ListTopics200ResponseResultsInnerSiblingsInner) *NullableListTopics200ResponseResultsInnerSiblingsInner {
	return &NullableListTopics200ResponseResultsInnerSiblingsInner{value: val, isSet: true}
}

func (v NullableListTopics200ResponseResultsInnerSiblingsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTopics200ResponseResultsInnerSiblingsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


