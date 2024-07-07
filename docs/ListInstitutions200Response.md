# ListInstitutions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListAuthors200ResponseMeta**](ListAuthors200ResponseMeta.md) |  | [optional] 
**Results** | Pointer to [**[]ListInstitutions200ResponseResultsInner**](ListInstitutions200ResponseResultsInner.md) |  | [optional] 
**GroupBy** | Pointer to [**[]ListAuthors200ResponseGroupByInner**](ListAuthors200ResponseGroupByInner.md) |  | [optional] 

## Methods

### NewListInstitutions200Response

`func NewListInstitutions200Response() *ListInstitutions200Response`

NewListInstitutions200Response instantiates a new ListInstitutions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstitutions200ResponseWithDefaults

`func NewListInstitutions200ResponseWithDefaults() *ListInstitutions200Response`

NewListInstitutions200ResponseWithDefaults instantiates a new ListInstitutions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListInstitutions200Response) GetMeta() ListAuthors200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListInstitutions200Response) GetMetaOk() (*ListAuthors200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListInstitutions200Response) SetMeta(v ListAuthors200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListInstitutions200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetResults

`func (o *ListInstitutions200Response) GetResults() []ListInstitutions200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListInstitutions200Response) GetResultsOk() (*[]ListInstitutions200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListInstitutions200Response) SetResults(v []ListInstitutions200ResponseResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListInstitutions200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetGroupBy

`func (o *ListInstitutions200Response) GetGroupBy() []ListAuthors200ResponseGroupByInner`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *ListInstitutions200Response) GetGroupByOk() (*[]ListAuthors200ResponseGroupByInner, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *ListInstitutions200Response) SetGroupBy(v []ListAuthors200ResponseGroupByInner)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *ListInstitutions200Response) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


