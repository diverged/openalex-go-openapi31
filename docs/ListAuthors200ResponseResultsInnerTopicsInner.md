# ListAuthors200ResponseResultsInnerTopicsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The number of works associated with this topic. | [optional] 
**DisplayName** | Pointer to **string** | The name of the topic. | [optional] 
**Domain** | Pointer to [**ListAuthors200ResponseResultsInnerTopicsInnerDomain**](ListAuthors200ResponseResultsInnerTopicsInnerDomain.md) |  | [optional] 
**Field** | Pointer to [**ListAuthors200ResponseResultsInnerTopicsInnerDomain**](ListAuthors200ResponseResultsInnerTopicsInnerDomain.md) |  | [optional] 
**Id** | Pointer to **string** | The OpenAlex ID for this topic. | [optional] 
**Score** | Pointer to **float32** | The strength of the connection between a work and this topic. | [optional] 
**Subfield** | Pointer to [**ListAuthors200ResponseResultsInnerTopicsInnerDomain**](ListAuthors200ResponseResultsInnerTopicsInnerDomain.md) |  | [optional] 
**Value** | Pointer to **float32** |  | [optional] 

## Methods

### NewListAuthors200ResponseResultsInnerTopicsInner

`func NewListAuthors200ResponseResultsInnerTopicsInner() *ListAuthors200ResponseResultsInnerTopicsInner`

NewListAuthors200ResponseResultsInnerTopicsInner instantiates a new ListAuthors200ResponseResultsInnerTopicsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAuthors200ResponseResultsInnerTopicsInnerWithDefaults

`func NewListAuthors200ResponseResultsInnerTopicsInnerWithDefaults() *ListAuthors200ResponseResultsInnerTopicsInner`

NewListAuthors200ResponseResultsInnerTopicsInnerWithDefaults instantiates a new ListAuthors200ResponseResultsInnerTopicsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDisplayName

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDomain

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) GetDomain() ListAuthors200ResponseResultsInnerTopicsInnerDomain`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) GetDomainOk() (*ListAuthors200ResponseResultsInnerTopicsInnerDomain, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) SetDomain(v ListAuthors200ResponseResultsInnerTopicsInnerDomain)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetField

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) GetField() ListAuthors200ResponseResultsInnerTopicsInnerDomain`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) GetFieldOk() (*ListAuthors200ResponseResultsInnerTopicsInnerDomain, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) SetField(v ListAuthors200ResponseResultsInnerTopicsInnerDomain)`

SetField sets Field field to given value.

### HasField

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) HasField() bool`

HasField returns a boolean if a field has been set.

### GetId

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScore

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSubfield

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) GetSubfield() ListAuthors200ResponseResultsInnerTopicsInnerDomain`

GetSubfield returns the Subfield field if non-nil, zero value otherwise.

### GetSubfieldOk

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) GetSubfieldOk() (*ListAuthors200ResponseResultsInnerTopicsInnerDomain, bool)`

GetSubfieldOk returns a tuple with the Subfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubfield

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) SetSubfield(v ListAuthors200ResponseResultsInnerTopicsInnerDomain)`

SetSubfield sets Subfield field to given value.

### HasSubfield

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) HasSubfield() bool`

HasSubfield returns a boolean if a field has been set.

### GetValue

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ListAuthors200ResponseResultsInnerTopicsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


