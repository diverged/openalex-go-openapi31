# ListWorks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListWorks200ResponseMeta**](ListWorks200ResponseMeta.md) |  | [optional] 
**Results** | Pointer to [**[]ListWorks200ResponseResultsInner**](ListWorks200ResponseResultsInner.md) |  | [optional] 

## Methods

### NewListWorks200Response

`func NewListWorks200Response() *ListWorks200Response`

NewListWorks200Response instantiates a new ListWorks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorks200ResponseWithDefaults

`func NewListWorks200ResponseWithDefaults() *ListWorks200Response`

NewListWorks200ResponseWithDefaults instantiates a new ListWorks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListWorks200Response) GetMeta() ListWorks200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListWorks200Response) GetMetaOk() (*ListWorks200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListWorks200Response) SetMeta(v ListWorks200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListWorks200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetResults

`func (o *ListWorks200Response) GetResults() []ListWorks200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListWorks200Response) GetResultsOk() (*[]ListWorks200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListWorks200Response) SetResults(v []ListWorks200ResponseResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListWorks200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


