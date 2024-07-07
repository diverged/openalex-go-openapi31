# ListTopics200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListAuthors200ResponseMeta**](ListAuthors200ResponseMeta.md) |  | [optional] 
**Results** | Pointer to [**[]ListTopics200ResponseResultsInner**](ListTopics200ResponseResultsInner.md) |  | [optional] 
**GroupBy** | Pointer to [**[]ListAuthors200ResponseGroupByInner**](ListAuthors200ResponseGroupByInner.md) |  | [optional] 

## Methods

### NewListTopics200Response

`func NewListTopics200Response() *ListTopics200Response`

NewListTopics200Response instantiates a new ListTopics200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTopics200ResponseWithDefaults

`func NewListTopics200ResponseWithDefaults() *ListTopics200Response`

NewListTopics200ResponseWithDefaults instantiates a new ListTopics200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListTopics200Response) GetMeta() ListAuthors200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListTopics200Response) GetMetaOk() (*ListAuthors200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListTopics200Response) SetMeta(v ListAuthors200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListTopics200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetResults

`func (o *ListTopics200Response) GetResults() []ListTopics200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListTopics200Response) GetResultsOk() (*[]ListTopics200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListTopics200Response) SetResults(v []ListTopics200ResponseResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListTopics200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetGroupBy

`func (o *ListTopics200Response) GetGroupBy() []ListAuthors200ResponseGroupByInner`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *ListTopics200Response) GetGroupByOk() (*[]ListAuthors200ResponseGroupByInner, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *ListTopics200Response) SetGroupBy(v []ListAuthors200ResponseGroupByInner)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *ListTopics200Response) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


