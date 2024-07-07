# ListKeywords200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListAuthors200ResponseMeta**](ListAuthors200ResponseMeta.md) |  | [optional] 
**Results** | Pointer to [**[]ListKeywords200ResponseResultsInner**](ListKeywords200ResponseResultsInner.md) |  | [optional] 
**GroupBy** | Pointer to [**[]ListAuthors200ResponseGroupByInner**](ListAuthors200ResponseGroupByInner.md) |  | [optional] 

## Methods

### NewListKeywords200Response

`func NewListKeywords200Response() *ListKeywords200Response`

NewListKeywords200Response instantiates a new ListKeywords200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListKeywords200ResponseWithDefaults

`func NewListKeywords200ResponseWithDefaults() *ListKeywords200Response`

NewListKeywords200ResponseWithDefaults instantiates a new ListKeywords200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListKeywords200Response) GetMeta() ListAuthors200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListKeywords200Response) GetMetaOk() (*ListAuthors200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListKeywords200Response) SetMeta(v ListAuthors200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListKeywords200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetResults

`func (o *ListKeywords200Response) GetResults() []ListKeywords200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListKeywords200Response) GetResultsOk() (*[]ListKeywords200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListKeywords200Response) SetResults(v []ListKeywords200ResponseResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListKeywords200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetGroupBy

`func (o *ListKeywords200Response) GetGroupBy() []ListAuthors200ResponseGroupByInner`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *ListKeywords200Response) GetGroupByOk() (*[]ListAuthors200ResponseGroupByInner, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *ListKeywords200Response) SetGroupBy(v []ListAuthors200ResponseGroupByInner)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *ListKeywords200Response) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


