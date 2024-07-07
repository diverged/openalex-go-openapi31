# ListAuthors200ResponseResultsInnerLastKnownInstitution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The OpenAlex ID for this institution. | 
**DisplayName** | **string** | The primary name of the institution. | 
**Ror** | Pointer to **string** | The ROR ID for this institution. | [optional] 
**CountryCode** | Pointer to **string** | The country where this institution is located, represented as an ISO two-letter country code. | [optional] 
**Type** | **string** | The institution&#39;s primary type, using the ROR \&quot;type\&quot; controlled vocabulary. | 
**Lineage** | Pointer to **[]string** | OpenAlex IDs of this institution and its parent institutions. | [optional] 

## Methods

### NewListAuthors200ResponseResultsInnerLastKnownInstitution

`func NewListAuthors200ResponseResultsInnerLastKnownInstitution(id string, displayName string, type_ string, ) *ListAuthors200ResponseResultsInnerLastKnownInstitution`

NewListAuthors200ResponseResultsInnerLastKnownInstitution instantiates a new ListAuthors200ResponseResultsInnerLastKnownInstitution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAuthors200ResponseResultsInnerLastKnownInstitutionWithDefaults

`func NewListAuthors200ResponseResultsInnerLastKnownInstitutionWithDefaults() *ListAuthors200ResponseResultsInnerLastKnownInstitution`

NewListAuthors200ResponseResultsInnerLastKnownInstitutionWithDefaults instantiates a new ListAuthors200ResponseResultsInnerLastKnownInstitution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListAuthors200ResponseResultsInnerLastKnownInstitution) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListAuthors200ResponseResultsInnerLastKnownInstitution) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListAuthors200ResponseResultsInnerLastKnownInstitution) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayName

`func (o *ListAuthors200ResponseResultsInnerLastKnownInstitution) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListAuthors200ResponseResultsInnerLastKnownInstitution) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListAuthors200ResponseResultsInnerLastKnownInstitution) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetRor

`func (o *ListAuthors200ResponseResultsInnerLastKnownInstitution) GetRor() string`

GetRor returns the Ror field if non-nil, zero value otherwise.

### GetRorOk

`func (o *ListAuthors200ResponseResultsInnerLastKnownInstitution) GetRorOk() (*string, bool)`

GetRorOk returns a tuple with the Ror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRor

`func (o *ListAuthors200ResponseResultsInnerLastKnownInstitution) SetRor(v string)`

SetRor sets Ror field to given value.

### HasRor

`func (o *ListAuthors200ResponseResultsInnerLastKnownInstitution) HasRor() bool`

HasRor returns a boolean if a field has been set.

### GetCountryCode

`func (o *ListAuthors200ResponseResultsInnerLastKnownInstitution) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ListAuthors200ResponseResultsInnerLastKnownInstitution) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ListAuthors200ResponseResultsInnerLastKnownInstitution) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *ListAuthors200ResponseResultsInnerLastKnownInstitution) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetType

`func (o *ListAuthors200ResponseResultsInnerLastKnownInstitution) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListAuthors200ResponseResultsInnerLastKnownInstitution) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListAuthors200ResponseResultsInnerLastKnownInstitution) SetType(v string)`

SetType sets Type field to given value.


### GetLineage

`func (o *ListAuthors200ResponseResultsInnerLastKnownInstitution) GetLineage() []string`

GetLineage returns the Lineage field if non-nil, zero value otherwise.

### GetLineageOk

`func (o *ListAuthors200ResponseResultsInnerLastKnownInstitution) GetLineageOk() (*[]string, bool)`

GetLineageOk returns a tuple with the Lineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineage

`func (o *ListAuthors200ResponseResultsInnerLastKnownInstitution) SetLineage(v []string)`

SetLineage sets Lineage field to given value.

### HasLineage

`func (o *ListAuthors200ResponseResultsInnerLastKnownInstitution) HasLineage() bool`

HasLineage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


