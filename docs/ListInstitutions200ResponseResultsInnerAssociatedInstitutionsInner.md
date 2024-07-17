# ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The OpenAlex ID for this institution. | 
**DisplayName** | **string** | The primary name of the institution. | 
**Ror** | Pointer to **string** | The ROR ID for this institution. | [optional] 
**CountryCode** | Pointer to **string** | The country where this institution is located, represented as an ISO two-letter country code. | [optional] 
**Type** | **string** | The institution&#39;s primary type, using the ROR \&quot;type\&quot; controlled vocabulary. | 
**Lineage** | Pointer to **[]string** | OpenAlex IDs of this institution and its parent institutions. | [optional] 
**Relationship** | Pointer to **string** | The relationship between this institution and the main institution. Possible values are &#x60;parent&#x60;, &#x60;child&#x60;, and &#x60;related&#x60;. Institution associations and the relationship vocabulary come from [ROR&#39;s Relationships](https://ror.readme.io/docs/ror-data-structure#relationships) | [optional] 

## Methods

### NewListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner

`func NewListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner(id string, displayName string, type_ string, ) *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner`

NewListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner instantiates a new ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstitutions200ResponseResultsInnerAssociatedInstitutionsInnerWithDefaults

`func NewListInstitutions200ResponseResultsInnerAssociatedInstitutionsInnerWithDefaults() *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner`

NewListInstitutions200ResponseResultsInnerAssociatedInstitutionsInnerWithDefaults instantiates a new ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayName

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetRor

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetRor() string`

GetRor returns the Ror field if non-nil, zero value otherwise.

### GetRorOk

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetRorOk() (*string, bool)`

GetRorOk returns a tuple with the Ror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRor

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) SetRor(v string)`

SetRor sets Ror field to given value.

### HasRor

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) HasRor() bool`

HasRor returns a boolean if a field has been set.

### GetCountryCode

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetType

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) SetType(v string)`

SetType sets Type field to given value.


### GetLineage

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetLineage() []string`

GetLineage returns the Lineage field if non-nil, zero value otherwise.

### GetLineageOk

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetLineageOk() (*[]string, bool)`

GetLineageOk returns a tuple with the Lineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineage

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) SetLineage(v []string)`

SetLineage sets Lineage field to given value.

### HasLineage

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) HasLineage() bool`

HasLineage returns a boolean if a field has been set.

### GetRelationship

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetRelationship() string`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) GetRelationshipOk() (*string, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) SetRelationship(v string)`

SetRelationship sets Relationship field to given value.

### HasRelationship

`func (o *ListInstitutions200ResponseResultsInnerAssociatedInstitutionsInner) HasRelationship() bool`

HasRelationship returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


