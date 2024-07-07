# ListAuthors200ResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of entities (this will be all entities if no filter is applied) | [optional] 
**DbResponseTimeMs** | Pointer to **int32** | The time taken by the database to respond, in milliseconds | [optional] 
**Page** | Pointer to **int32** | The current page number | [optional] 
**PerPage** | Pointer to **int32** | The number of results per page | [optional] 
**GroupsCount** | Pointer to **NullableInt32** | The count of groups in the current page (null if there are no groups in the response) | [optional] 

## Methods

### NewListAuthors200ResponseMeta

`func NewListAuthors200ResponseMeta() *ListAuthors200ResponseMeta`

NewListAuthors200ResponseMeta instantiates a new ListAuthors200ResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAuthors200ResponseMetaWithDefaults

`func NewListAuthors200ResponseMetaWithDefaults() *ListAuthors200ResponseMeta`

NewListAuthors200ResponseMetaWithDefaults instantiates a new ListAuthors200ResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ListAuthors200ResponseMeta) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListAuthors200ResponseMeta) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListAuthors200ResponseMeta) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListAuthors200ResponseMeta) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDbResponseTimeMs

`func (o *ListAuthors200ResponseMeta) GetDbResponseTimeMs() int32`

GetDbResponseTimeMs returns the DbResponseTimeMs field if non-nil, zero value otherwise.

### GetDbResponseTimeMsOk

`func (o *ListAuthors200ResponseMeta) GetDbResponseTimeMsOk() (*int32, bool)`

GetDbResponseTimeMsOk returns a tuple with the DbResponseTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbResponseTimeMs

`func (o *ListAuthors200ResponseMeta) SetDbResponseTimeMs(v int32)`

SetDbResponseTimeMs sets DbResponseTimeMs field to given value.

### HasDbResponseTimeMs

`func (o *ListAuthors200ResponseMeta) HasDbResponseTimeMs() bool`

HasDbResponseTimeMs returns a boolean if a field has been set.

### GetPage

`func (o *ListAuthors200ResponseMeta) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListAuthors200ResponseMeta) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListAuthors200ResponseMeta) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListAuthors200ResponseMeta) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPerPage

`func (o *ListAuthors200ResponseMeta) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *ListAuthors200ResponseMeta) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *ListAuthors200ResponseMeta) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *ListAuthors200ResponseMeta) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetGroupsCount

`func (o *ListAuthors200ResponseMeta) GetGroupsCount() int32`

GetGroupsCount returns the GroupsCount field if non-nil, zero value otherwise.

### GetGroupsCountOk

`func (o *ListAuthors200ResponseMeta) GetGroupsCountOk() (*int32, bool)`

GetGroupsCountOk returns a tuple with the GroupsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsCount

`func (o *ListAuthors200ResponseMeta) SetGroupsCount(v int32)`

SetGroupsCount sets GroupsCount field to given value.

### HasGroupsCount

`func (o *ListAuthors200ResponseMeta) HasGroupsCount() bool`

HasGroupsCount returns a boolean if a field has been set.

### SetGroupsCountNil

`func (o *ListAuthors200ResponseMeta) SetGroupsCountNil(b bool)`

 SetGroupsCountNil sets the value for GroupsCount to be an explicit nil

### UnsetGroupsCount
`func (o *ListAuthors200ResponseMeta) UnsetGroupsCount()`

UnsetGroupsCount ensures that no value is present for GroupsCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


