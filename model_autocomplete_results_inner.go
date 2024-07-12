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

// checks if the AutocompleteResultsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutocompleteResultsInner{}

// AutocompleteResultsInner struct for AutocompleteResultsInner
type AutocompleteResultsInner struct {
	// The OpenAlex ID for this result entity.
	Id *string `json:"id,omitempty"`
	// The Canonical External ID for this result entity.
	ExternalId *string `json:"external_id,omitempty"`
	// The entity's display_name property.
	DisplayName *string `json:"display_name,omitempty"`
	// The entity's type (author, concept, institution, source, publisher, funder, or work).
	EntityType *string `json:"entity_type,omitempty"`
	// The entity's cited_by_count property. For works, this is the number of incoming citations. For other entities, it's the sum of incoming citations for all linked works.
	CitedByCount *int32 `json:"cited_by_count,omitempty"`
	// The number of works associated with the entity. For entity type 'work' it's always null.
	WorksCount *int32 `json:"works_count,omitempty"`
	// Extra information to help identify the right item. Content varies by entity type.
	Hint *string `json:"hint,omitempty"`
}

// NewAutocompleteResultsInner instantiates a new AutocompleteResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutocompleteResultsInner() *AutocompleteResultsInner {
	this := AutocompleteResultsInner{}
	return &this
}

// NewAutocompleteResultsInnerWithDefaults instantiates a new AutocompleteResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutocompleteResultsInnerWithDefaults() *AutocompleteResultsInner {
	this := AutocompleteResultsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AutocompleteResultsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutocompleteResultsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AutocompleteResultsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AutocompleteResultsInner) SetId(v string) {
	o.Id = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *AutocompleteResultsInner) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutocompleteResultsInner) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *AutocompleteResultsInner) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *AutocompleteResultsInner) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AutocompleteResultsInner) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutocompleteResultsInner) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AutocompleteResultsInner) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AutocompleteResultsInner) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *AutocompleteResultsInner) GetEntityType() string {
	if o == nil || IsNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutocompleteResultsInner) GetEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *AutocompleteResultsInner) HasEntityType() bool {
	if o != nil && !IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *AutocompleteResultsInner) SetEntityType(v string) {
	o.EntityType = &v
}

// GetCitedByCount returns the CitedByCount field value if set, zero value otherwise.
func (o *AutocompleteResultsInner) GetCitedByCount() int32 {
	if o == nil || IsNil(o.CitedByCount) {
		var ret int32
		return ret
	}
	return *o.CitedByCount
}

// GetCitedByCountOk returns a tuple with the CitedByCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutocompleteResultsInner) GetCitedByCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CitedByCount) {
		return nil, false
	}
	return o.CitedByCount, true
}

// HasCitedByCount returns a boolean if a field has been set.
func (o *AutocompleteResultsInner) HasCitedByCount() bool {
	if o != nil && !IsNil(o.CitedByCount) {
		return true
	}

	return false
}

// SetCitedByCount gets a reference to the given int32 and assigns it to the CitedByCount field.
func (o *AutocompleteResultsInner) SetCitedByCount(v int32) {
	o.CitedByCount = &v
}

// GetWorksCount returns the WorksCount field value if set, zero value otherwise.
func (o *AutocompleteResultsInner) GetWorksCount() int32 {
	if o == nil || IsNil(o.WorksCount) {
		var ret int32
		return ret
	}
	return *o.WorksCount
}

// GetWorksCountOk returns a tuple with the WorksCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutocompleteResultsInner) GetWorksCountOk() (*int32, bool) {
	if o == nil || IsNil(o.WorksCount) {
		return nil, false
	}
	return o.WorksCount, true
}

// HasWorksCount returns a boolean if a field has been set.
func (o *AutocompleteResultsInner) HasWorksCount() bool {
	if o != nil && !IsNil(o.WorksCount) {
		return true
	}

	return false
}

// SetWorksCount gets a reference to the given int32 and assigns it to the WorksCount field.
func (o *AutocompleteResultsInner) SetWorksCount(v int32) {
	o.WorksCount = &v
}

// GetHint returns the Hint field value if set, zero value otherwise.
func (o *AutocompleteResultsInner) GetHint() string {
	if o == nil || IsNil(o.Hint) {
		var ret string
		return ret
	}
	return *o.Hint
}

// GetHintOk returns a tuple with the Hint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutocompleteResultsInner) GetHintOk() (*string, bool) {
	if o == nil || IsNil(o.Hint) {
		return nil, false
	}
	return o.Hint, true
}

// HasHint returns a boolean if a field has been set.
func (o *AutocompleteResultsInner) HasHint() bool {
	if o != nil && !IsNil(o.Hint) {
		return true
	}

	return false
}

// SetHint gets a reference to the given string and assigns it to the Hint field.
func (o *AutocompleteResultsInner) SetHint(v string) {
	o.Hint = &v
}

func (o AutocompleteResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutocompleteResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ExternalId) {
		toSerialize["external_id"] = o.ExternalId
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.EntityType) {
		toSerialize["entity_type"] = o.EntityType
	}
	if !IsNil(o.CitedByCount) {
		toSerialize["cited_by_count"] = o.CitedByCount
	}
	if !IsNil(o.WorksCount) {
		toSerialize["works_count"] = o.WorksCount
	}
	if !IsNil(o.Hint) {
		toSerialize["hint"] = o.Hint
	}
	return toSerialize, nil
}

type NullableAutocompleteResultsInner struct {
	value *AutocompleteResultsInner
	isSet bool
}

func (v NullableAutocompleteResultsInner) Get() *AutocompleteResultsInner {
	return v.value
}

func (v *NullableAutocompleteResultsInner) Set(val *AutocompleteResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAutocompleteResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAutocompleteResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutocompleteResultsInner(val *AutocompleteResultsInner) *NullableAutocompleteResultsInner {
	return &NullableAutocompleteResultsInner{value: val, isSet: true}
}

func (v NullableAutocompleteResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutocompleteResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

