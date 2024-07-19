# ListWorks200ResponseResultsInnerAuthorshipsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affiliations** | Pointer to [**[]ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner**](ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner.md) | Each institutional affiliation that this author has claimed will be listed here: the raw affiliation string that we found, along with the OpenAlex Institution ID or IDs that we matched it to. This information will be redundant with institutions below, but is useful if you need to know about what we used to match institutions. | [optional] 
**Author** | [**ListWorks200ResponseResultsInnerAuthorshipsInnerAuthor**](ListWorks200ResponseResultsInnerAuthorshipsInnerAuthor.md) |  | 
**AuthorPosition** | **string** | A summarized description of this author&#39;s position in the work&#39;s author list. | 
**Countries** | Pointer to **[]string** | The country or countries for this author, determined using a combination of matched institutions and parsing of the raw affiliation strings. | [optional] 
**Institutions** | [**[]ListAuthors200ResponseResultsInnerAffiliationsInnerInstitution**](ListAuthors200ResponseResultsInnerAffiliationsInnerInstitution.md) | The institutional affiliations this author claimed in the context of this work, as dehydrated Institution objects. | 
**IsCorresponding** | Pointer to **bool** | If true, this is a corresponding author for this work. | [optional] 
**RawAffiliationStrings** | **[]string** | This author&#39;s affiliation as it originally came to us, as a list of raw unformatted strings. | 
**RawAuthorName** | **string** | This author&#39;s name as it originally came to us, as a raw unformatted string. | 

## Methods

### NewListWorks200ResponseResultsInnerAuthorshipsInner

`func NewListWorks200ResponseResultsInnerAuthorshipsInner(author ListWorks200ResponseResultsInnerAuthorshipsInnerAuthor, authorPosition string, institutions []ListAuthors200ResponseResultsInnerAffiliationsInnerInstitution, rawAffiliationStrings []string, rawAuthorName string, ) *ListWorks200ResponseResultsInnerAuthorshipsInner`

NewListWorks200ResponseResultsInnerAuthorshipsInner instantiates a new ListWorks200ResponseResultsInnerAuthorshipsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorks200ResponseResultsInnerAuthorshipsInnerWithDefaults

`func NewListWorks200ResponseResultsInnerAuthorshipsInnerWithDefaults() *ListWorks200ResponseResultsInnerAuthorshipsInner`

NewListWorks200ResponseResultsInnerAuthorshipsInnerWithDefaults instantiates a new ListWorks200ResponseResultsInnerAuthorshipsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffiliations

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetAffiliations() []ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner`

GetAffiliations returns the Affiliations field if non-nil, zero value otherwise.

### GetAffiliationsOk

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetAffiliationsOk() (*[]ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner, bool)`

GetAffiliationsOk returns a tuple with the Affiliations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliations

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) SetAffiliations(v []ListWorks200ResponseResultsInnerAuthorshipsInnerAffiliationsInner)`

SetAffiliations sets Affiliations field to given value.

### HasAffiliations

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) HasAffiliations() bool`

HasAffiliations returns a boolean if a field has been set.

### GetAuthor

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetAuthor() ListWorks200ResponseResultsInnerAuthorshipsInnerAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetAuthorOk() (*ListWorks200ResponseResultsInnerAuthorshipsInnerAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) SetAuthor(v ListWorks200ResponseResultsInnerAuthorshipsInnerAuthor)`

SetAuthor sets Author field to given value.


### GetAuthorPosition

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetAuthorPosition() string`

GetAuthorPosition returns the AuthorPosition field if non-nil, zero value otherwise.

### GetAuthorPositionOk

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetAuthorPositionOk() (*string, bool)`

GetAuthorPositionOk returns a tuple with the AuthorPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorPosition

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) SetAuthorPosition(v string)`

SetAuthorPosition sets AuthorPosition field to given value.


### GetCountries

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetCountries() []string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetCountriesOk() (*[]string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) SetCountries(v []string)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetInstitutions

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetInstitutions() []ListAuthors200ResponseResultsInnerAffiliationsInnerInstitution`

GetInstitutions returns the Institutions field if non-nil, zero value otherwise.

### GetInstitutionsOk

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetInstitutionsOk() (*[]ListAuthors200ResponseResultsInnerAffiliationsInnerInstitution, bool)`

GetInstitutionsOk returns a tuple with the Institutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutions

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) SetInstitutions(v []ListAuthors200ResponseResultsInnerAffiliationsInnerInstitution)`

SetInstitutions sets Institutions field to given value.


### GetIsCorresponding

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetIsCorresponding() bool`

GetIsCorresponding returns the IsCorresponding field if non-nil, zero value otherwise.

### GetIsCorrespondingOk

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetIsCorrespondingOk() (*bool, bool)`

GetIsCorrespondingOk returns a tuple with the IsCorresponding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCorresponding

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) SetIsCorresponding(v bool)`

SetIsCorresponding sets IsCorresponding field to given value.

### HasIsCorresponding

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) HasIsCorresponding() bool`

HasIsCorresponding returns a boolean if a field has been set.

### GetRawAffiliationStrings

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetRawAffiliationStrings() []string`

GetRawAffiliationStrings returns the RawAffiliationStrings field if non-nil, zero value otherwise.

### GetRawAffiliationStringsOk

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetRawAffiliationStringsOk() (*[]string, bool)`

GetRawAffiliationStringsOk returns a tuple with the RawAffiliationStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawAffiliationStrings

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) SetRawAffiliationStrings(v []string)`

SetRawAffiliationStrings sets RawAffiliationStrings field to given value.


### GetRawAuthorName

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetRawAuthorName() string`

GetRawAuthorName returns the RawAuthorName field if non-nil, zero value otherwise.

### GetRawAuthorNameOk

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) GetRawAuthorNameOk() (*string, bool)`

GetRawAuthorNameOk returns a tuple with the RawAuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawAuthorName

`func (o *ListWorks200ResponseResultsInnerAuthorshipsInner) SetRawAuthorName(v string)`

SetRawAuthorName sets RawAuthorName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


