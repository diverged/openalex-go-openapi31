# ListWorks200ResponseResultsInnerOpenAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsOa** | **bool** | True if this work is Open Access (OA). | 
**OaStatus** | **string** | The Open Access (OA) status of this work. | 
**OaUrl** | Pointer to **string** | The best Open Access (OA) URL for this work. | [optional] 
**AnyRepositoryHasFulltext** | Pointer to **bool** | True if any of this work&#39;s locations has is_oa&#x3D;true and source.type&#x3D;repository. This indicates that a full-text version of the work is available in an open access repository. | [optional] 

## Methods

### NewListWorks200ResponseResultsInnerOpenAccess

`func NewListWorks200ResponseResultsInnerOpenAccess(isOa bool, oaStatus string, ) *ListWorks200ResponseResultsInnerOpenAccess`

NewListWorks200ResponseResultsInnerOpenAccess instantiates a new ListWorks200ResponseResultsInnerOpenAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorks200ResponseResultsInnerOpenAccessWithDefaults

`func NewListWorks200ResponseResultsInnerOpenAccessWithDefaults() *ListWorks200ResponseResultsInnerOpenAccess`

NewListWorks200ResponseResultsInnerOpenAccessWithDefaults instantiates a new ListWorks200ResponseResultsInnerOpenAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsOa

`func (o *ListWorks200ResponseResultsInnerOpenAccess) GetIsOa() bool`

GetIsOa returns the IsOa field if non-nil, zero value otherwise.

### GetIsOaOk

`func (o *ListWorks200ResponseResultsInnerOpenAccess) GetIsOaOk() (*bool, bool)`

GetIsOaOk returns a tuple with the IsOa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOa

`func (o *ListWorks200ResponseResultsInnerOpenAccess) SetIsOa(v bool)`

SetIsOa sets IsOa field to given value.


### GetOaStatus

`func (o *ListWorks200ResponseResultsInnerOpenAccess) GetOaStatus() string`

GetOaStatus returns the OaStatus field if non-nil, zero value otherwise.

### GetOaStatusOk

`func (o *ListWorks200ResponseResultsInnerOpenAccess) GetOaStatusOk() (*string, bool)`

GetOaStatusOk returns a tuple with the OaStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOaStatus

`func (o *ListWorks200ResponseResultsInnerOpenAccess) SetOaStatus(v string)`

SetOaStatus sets OaStatus field to given value.


### GetOaUrl

`func (o *ListWorks200ResponseResultsInnerOpenAccess) GetOaUrl() string`

GetOaUrl returns the OaUrl field if non-nil, zero value otherwise.

### GetOaUrlOk

`func (o *ListWorks200ResponseResultsInnerOpenAccess) GetOaUrlOk() (*string, bool)`

GetOaUrlOk returns a tuple with the OaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOaUrl

`func (o *ListWorks200ResponseResultsInnerOpenAccess) SetOaUrl(v string)`

SetOaUrl sets OaUrl field to given value.

### HasOaUrl

`func (o *ListWorks200ResponseResultsInnerOpenAccess) HasOaUrl() bool`

HasOaUrl returns a boolean if a field has been set.

### GetAnyRepositoryHasFulltext

`func (o *ListWorks200ResponseResultsInnerOpenAccess) GetAnyRepositoryHasFulltext() bool`

GetAnyRepositoryHasFulltext returns the AnyRepositoryHasFulltext field if non-nil, zero value otherwise.

### GetAnyRepositoryHasFulltextOk

`func (o *ListWorks200ResponseResultsInnerOpenAccess) GetAnyRepositoryHasFulltextOk() (*bool, bool)`

GetAnyRepositoryHasFulltextOk returns a tuple with the AnyRepositoryHasFulltext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyRepositoryHasFulltext

`func (o *ListWorks200ResponseResultsInnerOpenAccess) SetAnyRepositoryHasFulltext(v bool)`

SetAnyRepositoryHasFulltext sets AnyRepositoryHasFulltext field to given value.

### HasAnyRepositoryHasFulltext

`func (o *ListWorks200ResponseResultsInnerOpenAccess) HasAnyRepositoryHasFulltext() bool`

HasAnyRepositoryHasFulltext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


