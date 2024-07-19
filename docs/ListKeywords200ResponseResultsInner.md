# ListKeywords200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The OpenAlex ID for this keyword. | 
**DisplayName** | **string** | The name of the keyword. | 
**Score** | Pointer to **float32** | The similarity score of the keyword to the work&#39;s title and abstract text. Higher scores indicate greater relevance. | [optional] 
**WorksCount** | Pointer to **int32** | The number of works in OpenAlex that have this keyword. | [optional] 
**WorksApiUrl** | Pointer to **string** | An URL that will get you a list of all the works with this keyword. | [optional] 
**CitedByCount** | Pointer to **int32** | The total number of citations received by all works with this keyword. | [optional] 
**UpdatedDate** | Pointer to **string** | The last time anything in this keyword object changed. Formatted as ISO 8601 extended format without time zone designator. | [optional] 
**CreatedDate** | Pointer to **string** | The date this keyword object was created in the OpenAlex dataset. | [optional] 

## Methods

### NewListKeywords200ResponseResultsInner

`func NewListKeywords200ResponseResultsInner(id string, displayName string, ) *ListKeywords200ResponseResultsInner`

NewListKeywords200ResponseResultsInner instantiates a new ListKeywords200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListKeywords200ResponseResultsInnerWithDefaults

`func NewListKeywords200ResponseResultsInnerWithDefaults() *ListKeywords200ResponseResultsInner`

NewListKeywords200ResponseResultsInnerWithDefaults instantiates a new ListKeywords200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListKeywords200ResponseResultsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListKeywords200ResponseResultsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListKeywords200ResponseResultsInner) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayName

`func (o *ListKeywords200ResponseResultsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListKeywords200ResponseResultsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListKeywords200ResponseResultsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetScore

`func (o *ListKeywords200ResponseResultsInner) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ListKeywords200ResponseResultsInner) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ListKeywords200ResponseResultsInner) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ListKeywords200ResponseResultsInner) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetWorksCount

`func (o *ListKeywords200ResponseResultsInner) GetWorksCount() int32`

GetWorksCount returns the WorksCount field if non-nil, zero value otherwise.

### GetWorksCountOk

`func (o *ListKeywords200ResponseResultsInner) GetWorksCountOk() (*int32, bool)`

GetWorksCountOk returns a tuple with the WorksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksCount

`func (o *ListKeywords200ResponseResultsInner) SetWorksCount(v int32)`

SetWorksCount sets WorksCount field to given value.

### HasWorksCount

`func (o *ListKeywords200ResponseResultsInner) HasWorksCount() bool`

HasWorksCount returns a boolean if a field has been set.

### GetWorksApiUrl

`func (o *ListKeywords200ResponseResultsInner) GetWorksApiUrl() string`

GetWorksApiUrl returns the WorksApiUrl field if non-nil, zero value otherwise.

### GetWorksApiUrlOk

`func (o *ListKeywords200ResponseResultsInner) GetWorksApiUrlOk() (*string, bool)`

GetWorksApiUrlOk returns a tuple with the WorksApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksApiUrl

`func (o *ListKeywords200ResponseResultsInner) SetWorksApiUrl(v string)`

SetWorksApiUrl sets WorksApiUrl field to given value.

### HasWorksApiUrl

`func (o *ListKeywords200ResponseResultsInner) HasWorksApiUrl() bool`

HasWorksApiUrl returns a boolean if a field has been set.

### GetCitedByCount

`func (o *ListKeywords200ResponseResultsInner) GetCitedByCount() int32`

GetCitedByCount returns the CitedByCount field if non-nil, zero value otherwise.

### GetCitedByCountOk

`func (o *ListKeywords200ResponseResultsInner) GetCitedByCountOk() (*int32, bool)`

GetCitedByCountOk returns a tuple with the CitedByCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitedByCount

`func (o *ListKeywords200ResponseResultsInner) SetCitedByCount(v int32)`

SetCitedByCount sets CitedByCount field to given value.

### HasCitedByCount

`func (o *ListKeywords200ResponseResultsInner) HasCitedByCount() bool`

HasCitedByCount returns a boolean if a field has been set.

### GetUpdatedDate

`func (o *ListKeywords200ResponseResultsInner) GetUpdatedDate() string`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *ListKeywords200ResponseResultsInner) GetUpdatedDateOk() (*string, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *ListKeywords200ResponseResultsInner) SetUpdatedDate(v string)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *ListKeywords200ResponseResultsInner) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.

### GetCreatedDate

`func (o *ListKeywords200ResponseResultsInner) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ListKeywords200ResponseResultsInner) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ListKeywords200ResponseResultsInner) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ListKeywords200ResponseResultsInner) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


