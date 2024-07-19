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

// checks if the ListWorks200ResponseResultsInnerPrimaryTopicDomain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListWorks200ResponseResultsInnerPrimaryTopicDomain{}

// ListWorks200ResponseResultsInnerPrimaryTopicDomain struct for ListWorks200ResponseResultsInnerPrimaryTopicDomain
type ListWorks200ResponseResultsInnerPrimaryTopicDomain struct {
	DisplayName *string `json:"display_name,omitempty"`
	Id *string `json:"id,omitempty"`
}

// NewListWorks200ResponseResultsInnerPrimaryTopicDomain instantiates a new ListWorks200ResponseResultsInnerPrimaryTopicDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWorks200ResponseResultsInnerPrimaryTopicDomain() *ListWorks200ResponseResultsInnerPrimaryTopicDomain {
	this := ListWorks200ResponseResultsInnerPrimaryTopicDomain{}
	return &this
}

// NewListWorks200ResponseResultsInnerPrimaryTopicDomainWithDefaults instantiates a new ListWorks200ResponseResultsInnerPrimaryTopicDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWorks200ResponseResultsInnerPrimaryTopicDomainWithDefaults() *ListWorks200ResponseResultsInnerPrimaryTopicDomain {
	this := ListWorks200ResponseResultsInnerPrimaryTopicDomain{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ListWorks200ResponseResultsInnerPrimaryTopicDomain) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorks200ResponseResultsInnerPrimaryTopicDomain) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ListWorks200ResponseResultsInnerPrimaryTopicDomain) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ListWorks200ResponseResultsInnerPrimaryTopicDomain) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListWorks200ResponseResultsInnerPrimaryTopicDomain) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorks200ResponseResultsInnerPrimaryTopicDomain) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListWorks200ResponseResultsInnerPrimaryTopicDomain) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListWorks200ResponseResultsInnerPrimaryTopicDomain) SetId(v string) {
	o.Id = &v
}

func (o ListWorks200ResponseResultsInnerPrimaryTopicDomain) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListWorks200ResponseResultsInnerPrimaryTopicDomain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableListWorks200ResponseResultsInnerPrimaryTopicDomain struct {
	value *ListWorks200ResponseResultsInnerPrimaryTopicDomain
	isSet bool
}

func (v NullableListWorks200ResponseResultsInnerPrimaryTopicDomain) Get() *ListWorks200ResponseResultsInnerPrimaryTopicDomain {
	return v.value
}

func (v *NullableListWorks200ResponseResultsInnerPrimaryTopicDomain) Set(val *ListWorks200ResponseResultsInnerPrimaryTopicDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableListWorks200ResponseResultsInnerPrimaryTopicDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableListWorks200ResponseResultsInnerPrimaryTopicDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWorks200ResponseResultsInnerPrimaryTopicDomain(val *ListWorks200ResponseResultsInnerPrimaryTopicDomain) *NullableListWorks200ResponseResultsInnerPrimaryTopicDomain {
	return &NullableListWorks200ResponseResultsInnerPrimaryTopicDomain{value: val, isSet: true}
}

func (v NullableListWorks200ResponseResultsInnerPrimaryTopicDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWorks200ResponseResultsInnerPrimaryTopicDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


