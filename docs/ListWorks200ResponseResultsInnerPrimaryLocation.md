# ListWorks200ResponseResultsInnerPrimaryLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAccepted** | Pointer to **bool** | True if this location&#39;s version is either acceptedVersion or publishedVersion; otherwise false. | [optional] 
**IsOa** | Pointer to **bool** | True if an Open Access (OA) version of this work is available at this location. OA is defined as having a URL where you can read the fulltext of this work without needing to pay money or log in. | [optional] 
**IsPublished** | Pointer to **bool** | True if this location&#39;s version is publishedVersion; otherwise false. | [optional] 
**LandingPageUrl** | Pointer to **string** | The landing page URL for this location. | [optional] 
**License** | Pointer to **string** | The location&#39;s publishing license. This can be a Creative Commons license such as cc0 or cc-by, a publisher-specific license, or null which means we are not able to determine a license for this location. | [optional] 
**Source** | Pointer to [**ListWorks200ResponseResultsInnerBestOaLocationSource**](ListWorks200ResponseResultsInnerBestOaLocationSource.md) |  | [optional] 
**PdfUrl** | Pointer to **string** | A URL where you can find this location as a PDF. | [optional] 
**Version** | Pointer to **string** | The version of the work, based on the DRIVER Guidelines versioning scheme. | [optional] 

## Methods

### NewListWorks200ResponseResultsInnerPrimaryLocation

`func NewListWorks200ResponseResultsInnerPrimaryLocation() *ListWorks200ResponseResultsInnerPrimaryLocation`

NewListWorks200ResponseResultsInnerPrimaryLocation instantiates a new ListWorks200ResponseResultsInnerPrimaryLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorks200ResponseResultsInnerPrimaryLocationWithDefaults

`func NewListWorks200ResponseResultsInnerPrimaryLocationWithDefaults() *ListWorks200ResponseResultsInnerPrimaryLocation`

NewListWorks200ResponseResultsInnerPrimaryLocationWithDefaults instantiates a new ListWorks200ResponseResultsInnerPrimaryLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAccepted

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) GetIsAccepted() bool`

GetIsAccepted returns the IsAccepted field if non-nil, zero value otherwise.

### GetIsAcceptedOk

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) GetIsAcceptedOk() (*bool, bool)`

GetIsAcceptedOk returns a tuple with the IsAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccepted

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) SetIsAccepted(v bool)`

SetIsAccepted sets IsAccepted field to given value.

### HasIsAccepted

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) HasIsAccepted() bool`

HasIsAccepted returns a boolean if a field has been set.

### GetIsOa

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) GetIsOa() bool`

GetIsOa returns the IsOa field if non-nil, zero value otherwise.

### GetIsOaOk

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) GetIsOaOk() (*bool, bool)`

GetIsOaOk returns a tuple with the IsOa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOa

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) SetIsOa(v bool)`

SetIsOa sets IsOa field to given value.

### HasIsOa

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) HasIsOa() bool`

HasIsOa returns a boolean if a field has been set.

### GetIsPublished

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetLandingPageUrl

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) GetLandingPageUrl() string`

GetLandingPageUrl returns the LandingPageUrl field if non-nil, zero value otherwise.

### GetLandingPageUrlOk

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) GetLandingPageUrlOk() (*string, bool)`

GetLandingPageUrlOk returns a tuple with the LandingPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingPageUrl

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) SetLandingPageUrl(v string)`

SetLandingPageUrl sets LandingPageUrl field to given value.

### HasLandingPageUrl

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) HasLandingPageUrl() bool`

HasLandingPageUrl returns a boolean if a field has been set.

### GetLicense

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetSource

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) GetSource() ListWorks200ResponseResultsInnerBestOaLocationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) GetSourceOk() (*ListWorks200ResponseResultsInnerBestOaLocationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) SetSource(v ListWorks200ResponseResultsInnerBestOaLocationSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPdfUrl

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) GetPdfUrl() string`

GetPdfUrl returns the PdfUrl field if non-nil, zero value otherwise.

### GetPdfUrlOk

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) GetPdfUrlOk() (*string, bool)`

GetPdfUrlOk returns a tuple with the PdfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfUrl

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) SetPdfUrl(v string)`

SetPdfUrl sets PdfUrl field to given value.

### HasPdfUrl

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) HasPdfUrl() bool`

HasPdfUrl returns a boolean if a field has been set.

### GetVersion

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ListWorks200ResponseResultsInnerPrimaryLocation) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


