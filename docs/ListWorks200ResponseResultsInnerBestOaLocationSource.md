# ListWorks200ResponseResultsInnerBestOaLocationSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The OpenAlex ID for this source. | 
**DisplayName** | **string** | The name of the source. | 
**HostOrganization** | Pointer to **string** | The OpenAlex ID of the host organization for this source. | [optional] 
**HostOrganizationLineage** | Pointer to **[]string** | OpenAlex IDs of the host organization&#39;s lineage, from most specific to most general. | [optional] 
**HostOrganizationLineageNames** | Pointer to **[]string** | The display names of the host organization&#39;s lineage, from most specific to most general. | [optional] 
**HostOrganizationName** | Pointer to **string** | The display_name of the host organization, shown for convenience. | [optional] 
**IsCore** | Pointer to **bool** | Whether this source is identified as a \&quot;core source\&quot; by [CWTS](https://www.cwts.nl/), used in the [Open Leiden Ranking](https://open.leidenranking.com/) of universities around the world. The list of core sources can be found [here](https://zenodo.org/records/10949671). | [optional] 
**IsInDoaj** | Pointer to **bool** | Whether this source is listed in the Directory of Open Access Journals (DOAJ). | [optional] 
**IsOa** | Pointer to **bool** | Whether this is currently a fully open-access source. | [optional] 
**Issn** | Pointer to **[]string** | The ISSNs (International Standard Serial Numbers) used by this source. | [optional] 
**IssnL** | Pointer to **string** | The ISSN-L (linking ISSN) identifying this source. | [optional] 
**Type** | **string** | The type of source. | 

## Methods

### NewListWorks200ResponseResultsInnerBestOaLocationSource

`func NewListWorks200ResponseResultsInnerBestOaLocationSource(id string, displayName string, type_ string, ) *ListWorks200ResponseResultsInnerBestOaLocationSource`

NewListWorks200ResponseResultsInnerBestOaLocationSource instantiates a new ListWorks200ResponseResultsInnerBestOaLocationSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorks200ResponseResultsInnerBestOaLocationSourceWithDefaults

`func NewListWorks200ResponseResultsInnerBestOaLocationSourceWithDefaults() *ListWorks200ResponseResultsInnerBestOaLocationSource`

NewListWorks200ResponseResultsInnerBestOaLocationSourceWithDefaults instantiates a new ListWorks200ResponseResultsInnerBestOaLocationSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayName

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetHostOrganization

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) GetHostOrganization() string`

GetHostOrganization returns the HostOrganization field if non-nil, zero value otherwise.

### GetHostOrganizationOk

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) GetHostOrganizationOk() (*string, bool)`

GetHostOrganizationOk returns a tuple with the HostOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOrganization

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) SetHostOrganization(v string)`

SetHostOrganization sets HostOrganization field to given value.

### HasHostOrganization

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) HasHostOrganization() bool`

HasHostOrganization returns a boolean if a field has been set.

### GetHostOrganizationLineage

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) GetHostOrganizationLineage() []string`

GetHostOrganizationLineage returns the HostOrganizationLineage field if non-nil, zero value otherwise.

### GetHostOrganizationLineageOk

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) GetHostOrganizationLineageOk() (*[]string, bool)`

GetHostOrganizationLineageOk returns a tuple with the HostOrganizationLineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOrganizationLineage

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) SetHostOrganizationLineage(v []string)`

SetHostOrganizationLineage sets HostOrganizationLineage field to given value.

### HasHostOrganizationLineage

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) HasHostOrganizationLineage() bool`

HasHostOrganizationLineage returns a boolean if a field has been set.

### GetHostOrganizationLineageNames

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) GetHostOrganizationLineageNames() []string`

GetHostOrganizationLineageNames returns the HostOrganizationLineageNames field if non-nil, zero value otherwise.

### GetHostOrganizationLineageNamesOk

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) GetHostOrganizationLineageNamesOk() (*[]string, bool)`

GetHostOrganizationLineageNamesOk returns a tuple with the HostOrganizationLineageNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOrganizationLineageNames

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) SetHostOrganizationLineageNames(v []string)`

SetHostOrganizationLineageNames sets HostOrganizationLineageNames field to given value.

### HasHostOrganizationLineageNames

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) HasHostOrganizationLineageNames() bool`

HasHostOrganizationLineageNames returns a boolean if a field has been set.

### GetHostOrganizationName

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) GetHostOrganizationName() string`

GetHostOrganizationName returns the HostOrganizationName field if non-nil, zero value otherwise.

### GetHostOrganizationNameOk

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) GetHostOrganizationNameOk() (*string, bool)`

GetHostOrganizationNameOk returns a tuple with the HostOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOrganizationName

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) SetHostOrganizationName(v string)`

SetHostOrganizationName sets HostOrganizationName field to given value.

### HasHostOrganizationName

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) HasHostOrganizationName() bool`

HasHostOrganizationName returns a boolean if a field has been set.

### GetIsCore

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) GetIsCore() bool`

GetIsCore returns the IsCore field if non-nil, zero value otherwise.

### GetIsCoreOk

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) GetIsCoreOk() (*bool, bool)`

GetIsCoreOk returns a tuple with the IsCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCore

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) SetIsCore(v bool)`

SetIsCore sets IsCore field to given value.

### HasIsCore

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) HasIsCore() bool`

HasIsCore returns a boolean if a field has been set.

### GetIsInDoaj

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) GetIsInDoaj() bool`

GetIsInDoaj returns the IsInDoaj field if non-nil, zero value otherwise.

### GetIsInDoajOk

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) GetIsInDoajOk() (*bool, bool)`

GetIsInDoajOk returns a tuple with the IsInDoaj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInDoaj

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) SetIsInDoaj(v bool)`

SetIsInDoaj sets IsInDoaj field to given value.

### HasIsInDoaj

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) HasIsInDoaj() bool`

HasIsInDoaj returns a boolean if a field has been set.

### GetIsOa

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) GetIsOa() bool`

GetIsOa returns the IsOa field if non-nil, zero value otherwise.

### GetIsOaOk

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) GetIsOaOk() (*bool, bool)`

GetIsOaOk returns a tuple with the IsOa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOa

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) SetIsOa(v bool)`

SetIsOa sets IsOa field to given value.

### HasIsOa

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) HasIsOa() bool`

HasIsOa returns a boolean if a field has been set.

### GetIssn

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) GetIssn() []string`

GetIssn returns the Issn field if non-nil, zero value otherwise.

### GetIssnOk

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) GetIssnOk() (*[]string, bool)`

GetIssnOk returns a tuple with the Issn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssn

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) SetIssn(v []string)`

SetIssn sets Issn field to given value.

### HasIssn

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) HasIssn() bool`

HasIssn returns a boolean if a field has been set.

### GetIssnL

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) GetIssnL() string`

GetIssnL returns the IssnL field if non-nil, zero value otherwise.

### GetIssnLOk

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) GetIssnLOk() (*string, bool)`

GetIssnLOk returns a tuple with the IssnL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssnL

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) SetIssnL(v string)`

SetIssnL sets IssnL field to given value.

### HasIssnL

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) HasIssnL() bool`

HasIssnL returns a boolean if a field has been set.

### GetType

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListWorks200ResponseResultsInnerBestOaLocationSource) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


