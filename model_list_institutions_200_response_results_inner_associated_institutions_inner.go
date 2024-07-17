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

// checks if the ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner{}

// ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner A representation of an associated institution, extending the dehydrated institution with a relationship field.
type ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner struct {
	// The OpenAlex ID for this institution.
	Id string `json:"id"`
	// The primary name of the institution.
	DisplayName string `json:"display_name"`
	// The ROR ID for this institution.
	Ror *string `json:"ror,omitempty"`
	// The country where this institution is located, represented as an ISO two-letter country code.
	CountryCode *string `json:"country_code,omitempty"`
	// The institution's primary type, using the ROR \"type\" controlled vocabulary.
	Type string `json:"type"`
	// OpenAlex IDs of this institution and its parent institutions.
	Lineage []string `json:"lineage,omitempty"`
	// The relationship between this institution and the main institution. Possible values are `parent`, `child`, and `related`. Institution associations and the relationship vocabulary come from [ROR's Relationships](https://ror.readme.io/docs/ror-data-structure#relationships)
	Relationship *string `json:"relationship,omitempty"`
}

type _ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner

// NewListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner instantiates a new ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner(id string, displayName string, type_ string) *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner {
	this := ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner{}
	this.Id = id
	this.DisplayName = displayName
	this.Type = type_
	return &this
}

// NewListInstitutions200ResponseResultsInnerAssociatedInstitutionsInnerWithDefaults instantiates a new ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInstitutions200ResponseResultsInnerAssociatedInstitutionsInnerWithDefaults() *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner {
	this := ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner{}
	return &this
}

// GetId returns the Id field value
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) SetId(v string) {
	o.Id = v
}

// GetDisplayName returns the DisplayName field value
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetRor returns the Ror field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetRor() string {
	if o == nil || IsNil(o.Ror) {
		var ret string
		return ret
	}
	return *o.Ror
}

// GetRorOk returns a tuple with the Ror field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetRorOk() (*string, bool) {
	if o == nil || IsNil(o.Ror) {
		return nil, false
	}
	return o.Ror, true
}

// HasRor returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) HasRor() bool {
	if o != nil && !IsNil(o.Ror) {
		return true
	}

	return false
}

// SetRor gets a reference to the given string and assigns it to the Ror field.
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) SetRor(v string) {
	o.Ror = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetType returns the Type field value
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) SetType(v string) {
	o.Type = v
}

// GetLineage returns the Lineage field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetLineage() []string {
	if o == nil || IsNil(o.Lineage) {
		var ret []string
		return ret
	}
	return o.Lineage
}

// GetLineageOk returns a tuple with the Lineage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetLineageOk() ([]string, bool) {
	if o == nil || IsNil(o.Lineage) {
		return nil, false
	}
	return o.Lineage, true
}

// HasLineage returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) HasLineage() bool {
	if o != nil && !IsNil(o.Lineage) {
		return true
	}

	return false
}

// SetLineage gets a reference to the given []string and assigns it to the Lineage field.
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) SetLineage(v []string) {
	o.Lineage = v
}

// GetRelationship returns the Relationship field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetRelationship() string {
	if o == nil || IsNil(o.Relationship) {
		var ret string
		return ret
	}
	return *o.Relationship
}

// GetRelationshipOk returns a tuple with the Relationship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetRelationshipOk() (*string, bool) {
	if o == nil || IsNil(o.Relationship) {
		return nil, false
	}
	return o.Relationship, true
}

// HasRelationship returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) HasRelationship() bool {
	if o != nil && !IsNil(o.Relationship) {
		return true
	}

	return false
}

// SetRelationship gets a reference to the given string and assigns it to the Relationship field.
func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) SetRelationship(v string) {
	o.Relationship = &v
}

func (o ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["display_name"] = o.DisplayName
	if !IsNil(o.Ror) {
		toSerialize["ror"] = o.Ror
	}
	if !IsNil(o.CountryCode) {
		toSerialize["country_code"] = o.CountryCode
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Lineage) {
		toSerialize["lineage"] = o.Lineage
	}
	if !IsNil(o.Relationship) {
		toSerialize["relationship"] = o.Relationship
	}
	return toSerialize, nil
}

func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"display_name",
		"type",
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

	varListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner := _ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner)

	if err != nil {
		return err
	}

	*o = ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner(varListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner)

	return err
}

type NullableListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner struct {
	value *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner
	isSet bool
}

func (v NullableListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) Get() *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner {
	return v.value
}

func (v *NullableListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) Set(val *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner(val *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) *NullableListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner {
	return &NullableListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner{value: val, isSet: true}
}

func (v NullableListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


