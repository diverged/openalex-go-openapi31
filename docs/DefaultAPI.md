# \DefaultAPI

All URIs are relative to *https://api.openalex.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutocompleteEntities**](DefaultAPI.md#AutocompleteEntities) | **Get** /autocomplete/{entity_type} | Autocomplete entities
[**ListAuthors**](DefaultAPI.md#ListAuthors) | **Get** /authors | List Authors
[**ListConcepts**](DefaultAPI.md#ListConcepts) | **Get** /concepts | List Concepts
[**ListFunders**](DefaultAPI.md#ListFunders) | **Get** /funders | List Funders
[**ListInstitutions**](DefaultAPI.md#ListInstitutions) | **Get** /institutions | List Institutions
[**ListKeywords**](DefaultAPI.md#ListKeywords) | **Get** /keywords | List Keywords
[**ListPublishers**](DefaultAPI.md#ListPublishers) | **Get** /publishers | List Publishers
[**ListSources**](DefaultAPI.md#ListSources) | **Get** /sources | List Sources
[**ListTopics**](DefaultAPI.md#ListTopics) | **Get** /topics | List Topics
[**ListWorks**](DefaultAPI.md#ListWorks) | **Get** /works | List Works



## AutocompleteEntities

> Autocomplete AutocompleteEntities(ctx, entityType).Q(q).Execute()

Autocomplete entities



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	entityType := "entityType_example" // string | 
	q := "flori" // string | The search string supplied by the user for autocomplete suggestions.  Additional parameters: - filter: Optional filter to apply to the autocomplete results. Uses the same syntax as entity filters. - search: Optional search query to apply to the autocomplete results. Uses the same syntax as entity search. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AutocompleteEntities(context.Background(), entityType).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AutocompleteEntities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutocompleteEntities`: Autocomplete
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AutocompleteEntities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutocompleteEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | The search string supplied by the user for autocomplete suggestions.  Additional parameters: - filter: Optional filter to apply to the autocomplete results. Uses the same syntax as entity filters. - search: Optional search query to apply to the autocomplete results. Uses the same syntax as entity search.  | 

### Return type

[**Autocomplete**](Autocomplete.md)

### Authorization

[PoliteEmail](../README.md#PoliteEmail)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthors

> ListAuthors200Response ListAuthors(ctx).Filter(filter).Search(search).Sort(sort).Cursor(cursor).PerPage(perPage).Select_(select_).GroupBy(groupBy).Sample(sample).Execute()

List Authors



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	filter := "type:book" // string | Filters narrow the list down to just entities that meet particular conditions. Multiple filters can be combined using commas. Filters support logical expressions including inequality, negation, intersection, and addition.  (optional)
	search := "dna" // string | Text search query string to find entities that match the given text. (optional)
	sort := "cited_by_count:desc" // string | Specify the property to sort the results by. Multiple sort keys can be provided, separated by commas. Append ':desc' to sort in descending order. Default sort direction is ascending.  (optional)
	cursor := "*" // string | A cursor value for paginating through results. Use '*' to start cursor pagination, or use the 'next_cursor' value from the previous response to get the next page.  Warning: Don't use cursor paging to download the whole dataset. It's inefficient and puts a massive load on the servers. Instead, use the OpenAlex snapshot for downloading all data at once.  (optional)
	perPage := int32(25) // int32 | The number of results to return per page. Minimum is 1, maximum is 200, default is 25. (optional) (default to 25)
	select_ := "id,display_name" // string | Comma-separated list of fields to return in the results. Only root-level fields can be selected. (optional)
	groupBy := "authorships.institutions.id" // string | Groups the results by a specified property. The response will include a `group_by` object with counts for each group.  You can group by most of the same properties that you can filter by.  To include the \"unknown\" group in the response, add `:include_unknown` after the group-by parameter.  Example: `group_by=authorships.institutions.id` or `group_by=authorships.institutions.id:include_unknown`  The response will contain: - `key`: The OpenAlex ID or raw value of the `group_by` parameter for members of this group. - `key_display_name`: The `display_name` or raw value of the `group_by` parameter for members of this group. - `count`: The number of entities in the group.  (optional)
	sample := int32(25) // int32 | Get a random sample of results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListAuthors(context.Background()).Filter(filter).Search(search).Sort(sort).Cursor(cursor).PerPage(perPage).Select_(select_).GroupBy(groupBy).Sample(sample).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListAuthors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAuthors`: ListAuthors200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListAuthors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAuthorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filters narrow the list down to just entities that meet particular conditions. Multiple filters can be combined using commas. Filters support logical expressions including inequality, negation, intersection, and addition.  | 
 **search** | **string** | Text search query string to find entities that match the given text. | 
 **sort** | **string** | Specify the property to sort the results by. Multiple sort keys can be provided, separated by commas. Append &#39;:desc&#39; to sort in descending order. Default sort direction is ascending.  | 
 **cursor** | **string** | A cursor value for paginating through results. Use &#39;*&#39; to start cursor pagination, or use the &#39;next_cursor&#39; value from the previous response to get the next page.  Warning: Don&#39;t use cursor paging to download the whole dataset. It&#39;s inefficient and puts a massive load on the servers. Instead, use the OpenAlex snapshot for downloading all data at once.  | 
 **perPage** | **int32** | The number of results to return per page. Minimum is 1, maximum is 200, default is 25. | [default to 25]
 **select_** | **string** | Comma-separated list of fields to return in the results. Only root-level fields can be selected. | 
 **groupBy** | **string** | Groups the results by a specified property. The response will include a &#x60;group_by&#x60; object with counts for each group.  You can group by most of the same properties that you can filter by.  To include the \&quot;unknown\&quot; group in the response, add &#x60;:include_unknown&#x60; after the group-by parameter.  Example: &#x60;group_by&#x3D;authorships.institutions.id&#x60; or &#x60;group_by&#x3D;authorships.institutions.id:include_unknown&#x60;  The response will contain: - &#x60;key&#x60;: The OpenAlex ID or raw value of the &#x60;group_by&#x60; parameter for members of this group. - &#x60;key_display_name&#x60;: The &#x60;display_name&#x60; or raw value of the &#x60;group_by&#x60; parameter for members of this group. - &#x60;count&#x60;: The number of entities in the group.  | 
 **sample** | **int32** | Get a random sample of results | 

### Return type

[**ListAuthors200Response**](ListAuthors200Response.md)

### Authorization

[PoliteEmail](../README.md#PoliteEmail)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConcepts

> ListConcepts200Response ListConcepts(ctx).Filter(filter).Search(search).Sort(sort).Page(page).PerPage(perPage).Select_(select_).GroupBy(groupBy).Sample(sample).Execute()

List Concepts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	filter := "type:book" // string | Filters narrow the list down to just entities that meet particular conditions. Multiple filters can be combined using commas. Filters support logical expressions including inequality, negation, intersection, and addition.  (optional)
	search := "dna" // string | Text search query string to find entities that match the given text. (optional)
	sort := "cited_by_count:desc" // string | Specify the property to sort the results by. Multiple sort keys can be provided, separated by commas. Append ':desc' to sort in descending order. Default sort direction is ascending.  (optional)
	page := int32(1) // int32 | The page number of results to return. Used for basic paging. (optional) (default to 1)
	perPage := int32(25) // int32 | The number of results to return per page. Minimum is 1, maximum is 200, default is 25. (optional) (default to 25)
	select_ := "id,display_name" // string | Comma-separated list of fields to return in the results. Only root-level fields can be selected. (optional)
	groupBy := "authorships.institutions.id" // string | Groups the results by a specified property. The response will include a `group_by` object with counts for each group.  You can group by most of the same properties that you can filter by.  To include the \"unknown\" group in the response, add `:include_unknown` after the group-by parameter.  Example: `group_by=authorships.institutions.id` or `group_by=authorships.institutions.id:include_unknown`  The response will contain: - `key`: The OpenAlex ID or raw value of the `group_by` parameter for members of this group. - `key_display_name`: The `display_name` or raw value of the `group_by` parameter for members of this group. - `count`: The number of entities in the group.  (optional)
	sample := int32(10) // int32 | Get a random sample of results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListConcepts(context.Background()).Filter(filter).Search(search).Sort(sort).Page(page).PerPage(perPage).Select_(select_).GroupBy(groupBy).Sample(sample).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListConcepts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConcepts`: ListConcepts200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListConcepts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConceptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filters narrow the list down to just entities that meet particular conditions. Multiple filters can be combined using commas. Filters support logical expressions including inequality, negation, intersection, and addition.  | 
 **search** | **string** | Text search query string to find entities that match the given text. | 
 **sort** | **string** | Specify the property to sort the results by. Multiple sort keys can be provided, separated by commas. Append &#39;:desc&#39; to sort in descending order. Default sort direction is ascending.  | 
 **page** | **int32** | The page number of results to return. Used for basic paging. | [default to 1]
 **perPage** | **int32** | The number of results to return per page. Minimum is 1, maximum is 200, default is 25. | [default to 25]
 **select_** | **string** | Comma-separated list of fields to return in the results. Only root-level fields can be selected. | 
 **groupBy** | **string** | Groups the results by a specified property. The response will include a &#x60;group_by&#x60; object with counts for each group.  You can group by most of the same properties that you can filter by.  To include the \&quot;unknown\&quot; group in the response, add &#x60;:include_unknown&#x60; after the group-by parameter.  Example: &#x60;group_by&#x3D;authorships.institutions.id&#x60; or &#x60;group_by&#x3D;authorships.institutions.id:include_unknown&#x60;  The response will contain: - &#x60;key&#x60;: The OpenAlex ID or raw value of the &#x60;group_by&#x60; parameter for members of this group. - &#x60;key_display_name&#x60;: The &#x60;display_name&#x60; or raw value of the &#x60;group_by&#x60; parameter for members of this group. - &#x60;count&#x60;: The number of entities in the group.  | 
 **sample** | **int32** | Get a random sample of results | 

### Return type

[**ListConcepts200Response**](ListConcepts200Response.md)

### Authorization

[PoliteEmail](../README.md#PoliteEmail)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFunders

> ListFunders200Response ListFunders(ctx).Filter(filter).Search(search).Sort(sort).Page(page).PerPage(perPage).Select_(select_).GroupBy(groupBy).Sample(sample).Execute()

List Funders



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	filter := "type:book" // string | Filters narrow the list down to just entities that meet particular conditions. Multiple filters can be combined using commas. Filters support logical expressions including inequality, negation, intersection, and addition.  (optional)
	search := "dna" // string | Text search query string to find entities that match the given text. (optional)
	sort := "cited_by_count:desc" // string | Specify the property to sort the results by. Multiple sort keys can be provided, separated by commas. Append ':desc' to sort in descending order. Default sort direction is ascending.  (optional)
	page := int32(1) // int32 | The page number of results to return. Used for basic paging. (optional) (default to 1)
	perPage := int32(25) // int32 | The number of results to return per page. Minimum is 1, maximum is 200, default is 25. (optional) (default to 25)
	select_ := "id,display_name" // string | Comma-separated list of fields to return in the results. Only root-level fields can be selected. (optional)
	groupBy := "authorships.institutions.id" // string | Groups the results by a specified property. The response will include a `group_by` object with counts for each group.  You can group by most of the same properties that you can filter by.  To include the \"unknown\" group in the response, add `:include_unknown` after the group-by parameter.  Example: `group_by=authorships.institutions.id` or `group_by=authorships.institutions.id:include_unknown`  The response will contain: - `key`: The OpenAlex ID or raw value of the `group_by` parameter for members of this group. - `key_display_name`: The `display_name` or raw value of the `group_by` parameter for members of this group. - `count`: The number of entities in the group.  (optional)
	sample := int32(10) // int32 | Get a random sample of results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListFunders(context.Background()).Filter(filter).Search(search).Sort(sort).Page(page).PerPage(perPage).Select_(select_).GroupBy(groupBy).Sample(sample).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListFunders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFunders`: ListFunders200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListFunders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFundersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filters narrow the list down to just entities that meet particular conditions. Multiple filters can be combined using commas. Filters support logical expressions including inequality, negation, intersection, and addition.  | 
 **search** | **string** | Text search query string to find entities that match the given text. | 
 **sort** | **string** | Specify the property to sort the results by. Multiple sort keys can be provided, separated by commas. Append &#39;:desc&#39; to sort in descending order. Default sort direction is ascending.  | 
 **page** | **int32** | The page number of results to return. Used for basic paging. | [default to 1]
 **perPage** | **int32** | The number of results to return per page. Minimum is 1, maximum is 200, default is 25. | [default to 25]
 **select_** | **string** | Comma-separated list of fields to return in the results. Only root-level fields can be selected. | 
 **groupBy** | **string** | Groups the results by a specified property. The response will include a &#x60;group_by&#x60; object with counts for each group.  You can group by most of the same properties that you can filter by.  To include the \&quot;unknown\&quot; group in the response, add &#x60;:include_unknown&#x60; after the group-by parameter.  Example: &#x60;group_by&#x3D;authorships.institutions.id&#x60; or &#x60;group_by&#x3D;authorships.institutions.id:include_unknown&#x60;  The response will contain: - &#x60;key&#x60;: The OpenAlex ID or raw value of the &#x60;group_by&#x60; parameter for members of this group. - &#x60;key_display_name&#x60;: The &#x60;display_name&#x60; or raw value of the &#x60;group_by&#x60; parameter for members of this group. - &#x60;count&#x60;: The number of entities in the group.  | 
 **sample** | **int32** | Get a random sample of results | 

### Return type

[**ListFunders200Response**](ListFunders200Response.md)

### Authorization

[PoliteEmail](../README.md#PoliteEmail)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstitutions

> ListInstitutions200Response ListInstitutions(ctx).Filter(filter).Search(search).Sort(sort).Page(page).PerPage(perPage).Select_(select_).GroupBy(groupBy).Sample(sample).Execute()

List Institutions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	filter := "type:book" // string | Filters narrow the list down to just entities that meet particular conditions. Multiple filters can be combined using commas. Filters support logical expressions including inequality, negation, intersection, and addition.  (optional)
	search := "dna" // string | Text search query string to find entities that match the given text. (optional)
	sort := "cited_by_count:desc" // string | Specify the property to sort the results by. Multiple sort keys can be provided, separated by commas. Append ':desc' to sort in descending order. Default sort direction is ascending.  (optional)
	page := int32(1) // int32 | The page number of results to return. Used for basic paging. (optional) (default to 1)
	perPage := int32(25) // int32 | The number of results to return per page. Minimum is 1, maximum is 200, default is 25. (optional) (default to 25)
	select_ := "id,display_name" // string | Comma-separated list of fields to return in the results. Only root-level fields can be selected. (optional)
	groupBy := "authorships.institutions.id" // string | Groups the results by a specified property. The response will include a `group_by` object with counts for each group.  You can group by most of the same properties that you can filter by.  To include the \"unknown\" group in the response, add `:include_unknown` after the group-by parameter.  Example: `group_by=authorships.institutions.id` or `group_by=authorships.institutions.id:include_unknown`  The response will contain: - `key`: The OpenAlex ID or raw value of the `group_by` parameter for members of this group. - `key_display_name`: The `display_name` or raw value of the `group_by` parameter for members of this group. - `count`: The number of entities in the group.  (optional)
	sample := int32(10) // int32 | Get a random sample of results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListInstitutions(context.Background()).Filter(filter).Search(search).Sort(sort).Page(page).PerPage(perPage).Select_(select_).GroupBy(groupBy).Sample(sample).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListInstitutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInstitutions`: ListInstitutions200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListInstitutions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInstitutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filters narrow the list down to just entities that meet particular conditions. Multiple filters can be combined using commas. Filters support logical expressions including inequality, negation, intersection, and addition.  | 
 **search** | **string** | Text search query string to find entities that match the given text. | 
 **sort** | **string** | Specify the property to sort the results by. Multiple sort keys can be provided, separated by commas. Append &#39;:desc&#39; to sort in descending order. Default sort direction is ascending.  | 
 **page** | **int32** | The page number of results to return. Used for basic paging. | [default to 1]
 **perPage** | **int32** | The number of results to return per page. Minimum is 1, maximum is 200, default is 25. | [default to 25]
 **select_** | **string** | Comma-separated list of fields to return in the results. Only root-level fields can be selected. | 
 **groupBy** | **string** | Groups the results by a specified property. The response will include a &#x60;group_by&#x60; object with counts for each group.  You can group by most of the same properties that you can filter by.  To include the \&quot;unknown\&quot; group in the response, add &#x60;:include_unknown&#x60; after the group-by parameter.  Example: &#x60;group_by&#x3D;authorships.institutions.id&#x60; or &#x60;group_by&#x3D;authorships.institutions.id:include_unknown&#x60;  The response will contain: - &#x60;key&#x60;: The OpenAlex ID or raw value of the &#x60;group_by&#x60; parameter for members of this group. - &#x60;key_display_name&#x60;: The &#x60;display_name&#x60; or raw value of the &#x60;group_by&#x60; parameter for members of this group. - &#x60;count&#x60;: The number of entities in the group.  | 
 **sample** | **int32** | Get a random sample of results | 

### Return type

[**ListInstitutions200Response**](ListInstitutions200Response.md)

### Authorization

[PoliteEmail](../README.md#PoliteEmail)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKeywords

> ListKeywords200Response ListKeywords(ctx).Filter(filter).Search(search).Sort(sort).Page(page).PerPage(perPage).Select_(select_).GroupBy(groupBy).Sample(sample).Execute()

List Keywords



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	filter := "type:book" // string | Filters narrow the list down to just entities that meet particular conditions. Multiple filters can be combined using commas. Filters support logical expressions including inequality, negation, intersection, and addition.  (optional)
	search := "dna" // string | Text search query string to find entities that match the given text. (optional)
	sort := "cited_by_count:desc" // string | Specify the property to sort the results by. Multiple sort keys can be provided, separated by commas. Append ':desc' to sort in descending order. Default sort direction is ascending.  (optional)
	page := int32(1) // int32 | The page number of results to return. Used for basic paging. (optional) (default to 1)
	perPage := int32(25) // int32 | The number of results to return per page. Minimum is 1, maximum is 200, default is 25. (optional) (default to 25)
	select_ := "id,display_name" // string | Comma-separated list of fields to return in the results. Only root-level fields can be selected. (optional)
	groupBy := "authorships.institutions.id" // string | Groups the results by a specified property. The response will include a `group_by` object with counts for each group.  You can group by most of the same properties that you can filter by.  To include the \"unknown\" group in the response, add `:include_unknown` after the group-by parameter.  Example: `group_by=authorships.institutions.id` or `group_by=authorships.institutions.id:include_unknown`  The response will contain: - `key`: The OpenAlex ID or raw value of the `group_by` parameter for members of this group. - `key_display_name`: The `display_name` or raw value of the `group_by` parameter for members of this group. - `count`: The number of entities in the group.  (optional)
	sample := int32(10) // int32 | Get a random sample of results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListKeywords(context.Background()).Filter(filter).Search(search).Sort(sort).Page(page).PerPage(perPage).Select_(select_).GroupBy(groupBy).Sample(sample).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListKeywords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKeywords`: ListKeywords200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListKeywords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListKeywordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filters narrow the list down to just entities that meet particular conditions. Multiple filters can be combined using commas. Filters support logical expressions including inequality, negation, intersection, and addition.  | 
 **search** | **string** | Text search query string to find entities that match the given text. | 
 **sort** | **string** | Specify the property to sort the results by. Multiple sort keys can be provided, separated by commas. Append &#39;:desc&#39; to sort in descending order. Default sort direction is ascending.  | 
 **page** | **int32** | The page number of results to return. Used for basic paging. | [default to 1]
 **perPage** | **int32** | The number of results to return per page. Minimum is 1, maximum is 200, default is 25. | [default to 25]
 **select_** | **string** | Comma-separated list of fields to return in the results. Only root-level fields can be selected. | 
 **groupBy** | **string** | Groups the results by a specified property. The response will include a &#x60;group_by&#x60; object with counts for each group.  You can group by most of the same properties that you can filter by.  To include the \&quot;unknown\&quot; group in the response, add &#x60;:include_unknown&#x60; after the group-by parameter.  Example: &#x60;group_by&#x3D;authorships.institutions.id&#x60; or &#x60;group_by&#x3D;authorships.institutions.id:include_unknown&#x60;  The response will contain: - &#x60;key&#x60;: The OpenAlex ID or raw value of the &#x60;group_by&#x60; parameter for members of this group. - &#x60;key_display_name&#x60;: The &#x60;display_name&#x60; or raw value of the &#x60;group_by&#x60; parameter for members of this group. - &#x60;count&#x60;: The number of entities in the group.  | 
 **sample** | **int32** | Get a random sample of results | 

### Return type

[**ListKeywords200Response**](ListKeywords200Response.md)

### Authorization

[PoliteEmail](../README.md#PoliteEmail)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPublishers

> ListPublishers200Response ListPublishers(ctx).Filter(filter).Search(search).Sort(sort).Page(page).PerPage(perPage).Select_(select_).GroupBy(groupBy).Sample(sample).Execute()

List Publishers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	filter := "type:book" // string | Filters narrow the list down to just entities that meet particular conditions. Multiple filters can be combined using commas. Filters support logical expressions including inequality, negation, intersection, and addition.  (optional)
	search := "dna" // string | Text search query string to find entities that match the given text. (optional)
	sort := "cited_by_count:desc" // string | Specify the property to sort the results by. Multiple sort keys can be provided, separated by commas. Append ':desc' to sort in descending order. Default sort direction is ascending.  (optional)
	page := int32(1) // int32 | The page number of results to return. Used for basic paging. (optional) (default to 1)
	perPage := int32(25) // int32 | The number of results to return per page. Minimum is 1, maximum is 200, default is 25. (optional) (default to 25)
	select_ := "id,display_name" // string | Comma-separated list of fields to return in the results. Only root-level fields can be selected. (optional)
	groupBy := "authorships.institutions.id" // string | Groups the results by a specified property. The response will include a `group_by` object with counts for each group.  You can group by most of the same properties that you can filter by.  To include the \"unknown\" group in the response, add `:include_unknown` after the group-by parameter.  Example: `group_by=authorships.institutions.id` or `group_by=authorships.institutions.id:include_unknown`  The response will contain: - `key`: The OpenAlex ID or raw value of the `group_by` parameter for members of this group. - `key_display_name`: The `display_name` or raw value of the `group_by` parameter for members of this group. - `count`: The number of entities in the group.  (optional)
	sample := int32(10) // int32 | Get a random sample of results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListPublishers(context.Background()).Filter(filter).Search(search).Sort(sort).Page(page).PerPage(perPage).Select_(select_).GroupBy(groupBy).Sample(sample).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListPublishers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPublishers`: ListPublishers200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListPublishers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPublishersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filters narrow the list down to just entities that meet particular conditions. Multiple filters can be combined using commas. Filters support logical expressions including inequality, negation, intersection, and addition.  | 
 **search** | **string** | Text search query string to find entities that match the given text. | 
 **sort** | **string** | Specify the property to sort the results by. Multiple sort keys can be provided, separated by commas. Append &#39;:desc&#39; to sort in descending order. Default sort direction is ascending.  | 
 **page** | **int32** | The page number of results to return. Used for basic paging. | [default to 1]
 **perPage** | **int32** | The number of results to return per page. Minimum is 1, maximum is 200, default is 25. | [default to 25]
 **select_** | **string** | Comma-separated list of fields to return in the results. Only root-level fields can be selected. | 
 **groupBy** | **string** | Groups the results by a specified property. The response will include a &#x60;group_by&#x60; object with counts for each group.  You can group by most of the same properties that you can filter by.  To include the \&quot;unknown\&quot; group in the response, add &#x60;:include_unknown&#x60; after the group-by parameter.  Example: &#x60;group_by&#x3D;authorships.institutions.id&#x60; or &#x60;group_by&#x3D;authorships.institutions.id:include_unknown&#x60;  The response will contain: - &#x60;key&#x60;: The OpenAlex ID or raw value of the &#x60;group_by&#x60; parameter for members of this group. - &#x60;key_display_name&#x60;: The &#x60;display_name&#x60; or raw value of the &#x60;group_by&#x60; parameter for members of this group. - &#x60;count&#x60;: The number of entities in the group.  | 
 **sample** | **int32** | Get a random sample of results | 

### Return type

[**ListPublishers200Response**](ListPublishers200Response.md)

### Authorization

[PoliteEmail](../README.md#PoliteEmail)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSources

> ListSources200Response ListSources(ctx).Filter(filter).Search(search).Sort(sort).Page(page).PerPage(perPage).Select_(select_).GroupBy(groupBy).Sample(sample).Execute()

List Sources



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	filter := "type:book" // string | Filters narrow the list down to just entities that meet particular conditions. Multiple filters can be combined using commas. Filters support logical expressions including inequality, negation, intersection, and addition.  (optional)
	search := "dna" // string | Text search query string to find entities that match the given text. (optional)
	sort := "cited_by_count:desc" // string | Specify the property to sort the results by. Multiple sort keys can be provided, separated by commas. Append ':desc' to sort in descending order. Default sort direction is ascending.  (optional)
	page := int32(1) // int32 | The page number of results to return. Used for basic paging. (optional) (default to 1)
	perPage := int32(25) // int32 | The number of results to return per page. Minimum is 1, maximum is 200, default is 25. (optional) (default to 25)
	select_ := "id,display_name" // string | Comma-separated list of fields to return in the results. Only root-level fields can be selected. (optional)
	groupBy := "authorships.institutions.id" // string | Groups the results by a specified property. The response will include a `group_by` object with counts for each group.  You can group by most of the same properties that you can filter by.  To include the \"unknown\" group in the response, add `:include_unknown` after the group-by parameter.  Example: `group_by=authorships.institutions.id` or `group_by=authorships.institutions.id:include_unknown`  The response will contain: - `key`: The OpenAlex ID or raw value of the `group_by` parameter for members of this group. - `key_display_name`: The `display_name` or raw value of the `group_by` parameter for members of this group. - `count`: The number of entities in the group.  (optional)
	sample := int32(10) // int32 | Get a random sample of results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListSources(context.Background()).Filter(filter).Search(search).Sort(sort).Page(page).PerPage(perPage).Select_(select_).GroupBy(groupBy).Sample(sample).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSources`: ListSources200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filters narrow the list down to just entities that meet particular conditions. Multiple filters can be combined using commas. Filters support logical expressions including inequality, negation, intersection, and addition.  | 
 **search** | **string** | Text search query string to find entities that match the given text. | 
 **sort** | **string** | Specify the property to sort the results by. Multiple sort keys can be provided, separated by commas. Append &#39;:desc&#39; to sort in descending order. Default sort direction is ascending.  | 
 **page** | **int32** | The page number of results to return. Used for basic paging. | [default to 1]
 **perPage** | **int32** | The number of results to return per page. Minimum is 1, maximum is 200, default is 25. | [default to 25]
 **select_** | **string** | Comma-separated list of fields to return in the results. Only root-level fields can be selected. | 
 **groupBy** | **string** | Groups the results by a specified property. The response will include a &#x60;group_by&#x60; object with counts for each group.  You can group by most of the same properties that you can filter by.  To include the \&quot;unknown\&quot; group in the response, add &#x60;:include_unknown&#x60; after the group-by parameter.  Example: &#x60;group_by&#x3D;authorships.institutions.id&#x60; or &#x60;group_by&#x3D;authorships.institutions.id:include_unknown&#x60;  The response will contain: - &#x60;key&#x60;: The OpenAlex ID or raw value of the &#x60;group_by&#x60; parameter for members of this group. - &#x60;key_display_name&#x60;: The &#x60;display_name&#x60; or raw value of the &#x60;group_by&#x60; parameter for members of this group. - &#x60;count&#x60;: The number of entities in the group.  | 
 **sample** | **int32** | Get a random sample of results | 

### Return type

[**ListSources200Response**](ListSources200Response.md)

### Authorization

[PoliteEmail](../README.md#PoliteEmail)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTopics

> ListTopics200Response ListTopics(ctx).Filter(filter).Search(search).Sort(sort).Page(page).PerPage(perPage).Select_(select_).GroupBy(groupBy).Sample(sample).Execute()

List Topics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	filter := "type:book" // string | Filters narrow the list down to just entities that meet particular conditions. Multiple filters can be combined using commas. Filters support logical expressions including inequality, negation, intersection, and addition.  (optional)
	search := "dna" // string | Text search query string to find entities that match the given text. (optional)
	sort := "cited_by_count:desc" // string | Specify the property to sort the results by. Multiple sort keys can be provided, separated by commas. Append ':desc' to sort in descending order. Default sort direction is ascending.  (optional)
	page := int32(1) // int32 | The page number of results to return. Used for basic paging. (optional) (default to 1)
	perPage := int32(25) // int32 | The number of results to return per page. Minimum is 1, maximum is 200, default is 25. (optional) (default to 25)
	select_ := "id,display_name" // string | Comma-separated list of fields to return in the results. Only root-level fields can be selected. (optional)
	groupBy := "authorships.institutions.id" // string | Groups the results by a specified property. The response will include a `group_by` object with counts for each group.  You can group by most of the same properties that you can filter by.  To include the \"unknown\" group in the response, add `:include_unknown` after the group-by parameter.  Example: `group_by=authorships.institutions.id` or `group_by=authorships.institutions.id:include_unknown`  The response will contain: - `key`: The OpenAlex ID or raw value of the `group_by` parameter for members of this group. - `key_display_name`: The `display_name` or raw value of the `group_by` parameter for members of this group. - `count`: The number of entities in the group.  (optional)
	sample := int32(10) // int32 | Get a random sample of results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListTopics(context.Background()).Filter(filter).Search(search).Sort(sort).Page(page).PerPage(perPage).Select_(select_).GroupBy(groupBy).Sample(sample).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListTopics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTopics`: ListTopics200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListTopics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filters narrow the list down to just entities that meet particular conditions. Multiple filters can be combined using commas. Filters support logical expressions including inequality, negation, intersection, and addition.  | 
 **search** | **string** | Text search query string to find entities that match the given text. | 
 **sort** | **string** | Specify the property to sort the results by. Multiple sort keys can be provided, separated by commas. Append &#39;:desc&#39; to sort in descending order. Default sort direction is ascending.  | 
 **page** | **int32** | The page number of results to return. Used for basic paging. | [default to 1]
 **perPage** | **int32** | The number of results to return per page. Minimum is 1, maximum is 200, default is 25. | [default to 25]
 **select_** | **string** | Comma-separated list of fields to return in the results. Only root-level fields can be selected. | 
 **groupBy** | **string** | Groups the results by a specified property. The response will include a &#x60;group_by&#x60; object with counts for each group.  You can group by most of the same properties that you can filter by.  To include the \&quot;unknown\&quot; group in the response, add &#x60;:include_unknown&#x60; after the group-by parameter.  Example: &#x60;group_by&#x3D;authorships.institutions.id&#x60; or &#x60;group_by&#x3D;authorships.institutions.id:include_unknown&#x60;  The response will contain: - &#x60;key&#x60;: The OpenAlex ID or raw value of the &#x60;group_by&#x60; parameter for members of this group. - &#x60;key_display_name&#x60;: The &#x60;display_name&#x60; or raw value of the &#x60;group_by&#x60; parameter for members of this group. - &#x60;count&#x60;: The number of entities in the group.  | 
 **sample** | **int32** | Get a random sample of results | 

### Return type

[**ListTopics200Response**](ListTopics200Response.md)

### Authorization

[PoliteEmail](../README.md#PoliteEmail)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorks

> ListWorks200Response ListWorks(ctx).Filter(filter).Search(search).Sort(sort).Page(page).PerPage(perPage).Select_(select_).GroupBy(groupBy).Sample(sample).Execute()

List Works



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	filter := "type:book" // string | Filters narrow the list down to just entities that meet particular conditions. Multiple filters can be combined using commas. Filters support logical expressions including inequality, negation, intersection, and addition.  (optional)
	search := "dna" // string | Text search query string to find entities that match the given text. (optional)
	sort := "cited_by_count:desc" // string | Specify the property to sort the results by. Multiple sort keys can be provided, separated by commas. Append ':desc' to sort in descending order. Default sort direction is ascending.  (optional)
	page := int32(1) // int32 | The page number of results to return. Used for basic paging. (optional) (default to 1)
	perPage := int32(25) // int32 | The number of results to return per page. Minimum is 1, maximum is 200, default is 25. (optional) (default to 25)
	select_ := "id,display_name" // string | Comma-separated list of fields to return in the results. Only root-level fields can be selected. (optional)
	groupBy := "authorships.institutions.id" // string | Groups the results by a specified property. The response will include a `group_by` object with counts for each group.  You can group by most of the same properties that you can filter by.  To include the \"unknown\" group in the response, add `:include_unknown` after the group-by parameter.  Example: `group_by=authorships.institutions.id` or `group_by=authorships.institutions.id:include_unknown`  The response will contain: - `key`: The OpenAlex ID or raw value of the `group_by` parameter for members of this group. - `key_display_name`: The `display_name` or raw value of the `group_by` parameter for members of this group. - `count`: The number of entities in the group.  (optional)
	sample := int32(10) // int32 | Get a random sample of results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListWorks(context.Background()).Filter(filter).Search(search).Sort(sort).Page(page).PerPage(perPage).Select_(select_).GroupBy(groupBy).Sample(sample).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListWorks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorks`: ListWorks200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListWorks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filters narrow the list down to just entities that meet particular conditions. Multiple filters can be combined using commas. Filters support logical expressions including inequality, negation, intersection, and addition.  | 
 **search** | **string** | Text search query string to find entities that match the given text. | 
 **sort** | **string** | Specify the property to sort the results by. Multiple sort keys can be provided, separated by commas. Append &#39;:desc&#39; to sort in descending order. Default sort direction is ascending.  | 
 **page** | **int32** | The page number of results to return. Used for basic paging. | [default to 1]
 **perPage** | **int32** | The number of results to return per page. Minimum is 1, maximum is 200, default is 25. | [default to 25]
 **select_** | **string** | Comma-separated list of fields to return in the results. Only root-level fields can be selected. | 
 **groupBy** | **string** | Groups the results by a specified property. The response will include a &#x60;group_by&#x60; object with counts for each group.  You can group by most of the same properties that you can filter by.  To include the \&quot;unknown\&quot; group in the response, add &#x60;:include_unknown&#x60; after the group-by parameter.  Example: &#x60;group_by&#x3D;authorships.institutions.id&#x60; or &#x60;group_by&#x3D;authorships.institutions.id:include_unknown&#x60;  The response will contain: - &#x60;key&#x60;: The OpenAlex ID or raw value of the &#x60;group_by&#x60; parameter for members of this group. - &#x60;key_display_name&#x60;: The &#x60;display_name&#x60; or raw value of the &#x60;group_by&#x60; parameter for members of this group. - &#x60;count&#x60;: The number of entities in the group.  | 
 **sample** | **int32** | Get a random sample of results | 

### Return type

[**ListWorks200Response**](ListWorks200Response.md)

### Authorization

[PoliteEmail](../README.md#PoliteEmail)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

