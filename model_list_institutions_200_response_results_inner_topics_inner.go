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

// checks if the ListInstitutions200ResponseResultsInnerTopicsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListInstitutions200ResponseResultsInnerTopicsInner{}

// ListInstitutions200ResponseResultsInnerTopicsInner A simplified representation of a topic, used when a topic is referenced by another entity.
type ListInstitutions200ResponseResultsInnerTopicsInner struct {
	// The number of works associated with this topic.
	Count *int32 `json:"count,omitempty"`
	// The name of the topic.
	DisplayName *string `json:"display_name,omitempty"`
	Domain *ListInstitutions200ResponseResultsInnerTopicsInnerDomain `json:"domain,omitempty"`
	Field *ListInstitutions200ResponseResultsInnerTopicsInnerDomain `json:"field,omitempty"`
	// The OpenAlex ID for this topic.
	Id *string `json:"id,omitempty"`
	// The strength of the connection between a work and this topic.
	Score *float32 `json:"score,omitempty"`
	Subfield *ListInstitutions200ResponseResultsInnerTopicsInnerDomain `json:"subfield,omitempty"`
	Value *float32 `json:"value,omitempty"`
}

// NewListInstitutions200ResponseResultsInnerTopicsInner instantiates a new ListInstitutions200ResponseResultsInnerTopicsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInstitutions200ResponseResultsInnerTopicsInner() *ListInstitutions200ResponseResultsInnerTopicsInner {
	this := ListInstitutions200ResponseResultsInnerTopicsInner{}
	return &this
}

// NewListInstitutions200ResponseResultsInnerTopicsInnerWithDefaults instantiates a new ListInstitutions200ResponseResultsInnerTopicsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInstitutions200ResponseResultsInnerTopicsInnerWithDefaults() *ListInstitutions200ResponseResultsInnerTopicsInner {
	this := ListInstitutions200ResponseResultsInnerTopicsInner{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) SetCount(v int32) {
	o.Count = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) GetDomain() ListInstitutions200ResponseResultsInnerTopicsInnerDomain {
	if o == nil || IsNil(o.Domain) {
		var ret ListInstitutions200ResponseResultsInnerTopicsInnerDomain
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) GetDomainOk() (*ListInstitutions200ResponseResultsInnerTopicsInnerDomain, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given ListInstitutions200ResponseResultsInnerTopicsInnerDomain and assigns it to the Domain field.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) SetDomain(v ListInstitutions200ResponseResultsInnerTopicsInnerDomain) {
	o.Domain = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) GetField() ListInstitutions200ResponseResultsInnerTopicsInnerDomain {
	if o == nil || IsNil(o.Field) {
		var ret ListInstitutions200ResponseResultsInnerTopicsInnerDomain
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) GetFieldOk() (*ListInstitutions200ResponseResultsInnerTopicsInnerDomain, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given ListInstitutions200ResponseResultsInnerTopicsInnerDomain and assigns it to the Field field.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) SetField(v ListInstitutions200ResponseResultsInnerTopicsInnerDomain) {
	o.Field = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) SetId(v string) {
	o.Id = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) GetScore() float32 {
	if o == nil || IsNil(o.Score) {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) GetScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) SetScore(v float32) {
	o.Score = &v
}

// GetSubfield returns the Subfield field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) GetSubfield() ListInstitutions200ResponseResultsInnerTopicsInnerDomain {
	if o == nil || IsNil(o.Subfield) {
		var ret ListInstitutions200ResponseResultsInnerTopicsInnerDomain
		return ret
	}
	return *o.Subfield
}

// GetSubfieldOk returns a tuple with the Subfield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) GetSubfieldOk() (*ListInstitutions200ResponseResultsInnerTopicsInnerDomain, bool) {
	if o == nil || IsNil(o.Subfield) {
		return nil, false
	}
	return o.Subfield, true
}

// HasSubfield returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) HasSubfield() bool {
	if o != nil && !IsNil(o.Subfield) {
		return true
	}

	return false
}

// SetSubfield gets a reference to the given ListInstitutions200ResponseResultsInnerTopicsInnerDomain and assigns it to the Subfield field.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) SetSubfield(v ListInstitutions200ResponseResultsInnerTopicsInnerDomain) {
	o.Subfield = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *ListInstitutions200ResponseResultsInnerTopicsInner) SetValue(v float32) {
	o.Value = &v
}

func (o ListInstitutions200ResponseResultsInnerTopicsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListInstitutions200ResponseResultsInnerTopicsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !IsNil(o.Subfield) {
		toSerialize["subfield"] = o.Subfield
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableListInstitutions200ResponseResultsInnerTopicsInner struct {
	value *ListInstitutions200ResponseResultsInnerTopicsInner
	isSet bool
}

func (v NullableListInstitutions200ResponseResultsInnerTopicsInner) Get() *ListInstitutions200ResponseResultsInnerTopicsInner {
	return v.value
}

func (v *NullableListInstitutions200ResponseResultsInnerTopicsInner) Set(val *ListInstitutions200ResponseResultsInnerTopicsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListInstitutions200ResponseResultsInnerTopicsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListInstitutions200ResponseResultsInnerTopicsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInstitutions200ResponseResultsInnerTopicsInner(val *ListInstitutions200ResponseResultsInnerTopicsInner) *NullableListInstitutions200ResponseResultsInnerTopicsInner {
	return &NullableListInstitutions200ResponseResultsInnerTopicsInner{value: val, isSet: true}
}

func (v NullableListInstitutions200ResponseResultsInnerTopicsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInstitutions200ResponseResultsInnerTopicsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


