/*
OpenAlex API

Testing DefaultAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/diverged/openalex-go-openapi31"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_DefaultAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefaultAPIService AutocompleteEntities", func(t *testing.T) {

		// 		t.Skip("skip test")  // remove to run test

		// var entityType string
		var entityType string = "authors" // Example entity type
		var query string = "example"      // Example query

		resp, httpRes, err := apiClient.DefaultAPI.AutocompleteEntities(context.Background(), entityType).Q(query).Execute()

		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if httpRes == nil {
			t.Error("Received nil HTTP response")
		} else {
			t.Logf("HTTP Status Code: %d", httpRes.StatusCode)
		}
		if resp == nil {
			t.Error("Received nil response")
		}

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListAuthors", func(t *testing.T) {

		// 		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.ListAuthors(context.Background()).Execute()

		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if httpRes == nil {
			t.Error("Received nil HTTP response")
		} else {
			t.Logf("HTTP Status Code: %d", httpRes.StatusCode)
		}
		if resp == nil {
			t.Error("Received nil response")
		}

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListConcepts", func(t *testing.T) {

		// 		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.ListConcepts(context.Background()).Execute()

		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if httpRes == nil {
			t.Error("Received nil HTTP response")
		} else {
			t.Logf("HTTP Status Code: %d", httpRes.StatusCode)
		}
		if resp == nil {
			t.Error("Received nil response")
		}

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListFunders", func(t *testing.T) {

		// 		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.ListFunders(context.Background()).Execute()

		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if httpRes == nil {
			t.Error("Received nil HTTP response")
		} else {
			t.Logf("HTTP Status Code: %d", httpRes.StatusCode)
		}
		if resp == nil {
			t.Error("Received nil response")
		}

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListInstitutions", func(t *testing.T) {

		// 		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.ListInstitutions(context.Background()).Execute()

		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if httpRes == nil {
			t.Error("Received nil HTTP response")
		} else {
			t.Logf("HTTP Status Code: %d", httpRes.StatusCode)
		}
		if resp == nil {
			t.Error("Received nil response")
		}

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListKeywords", func(t *testing.T) {

		// 		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.ListKeywords(context.Background()).Execute()

		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if httpRes == nil {
			t.Error("Received nil HTTP response")
		} else {
			t.Logf("HTTP Status Code: %d", httpRes.StatusCode)
		}
		if resp == nil {
			t.Error("Received nil response")
		}

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListPublishers", func(t *testing.T) {

		// 		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.ListPublishers(context.Background()).Execute()

		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if httpRes == nil {
			t.Error("Received nil HTTP response")
		} else {
			t.Logf("HTTP Status Code: %d", httpRes.StatusCode)
		}
		if resp == nil {
			t.Error("Received nil response")
		}

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListSources", func(t *testing.T) {

		// 		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.ListSources(context.Background()).Execute()

		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if httpRes == nil {
			t.Error("Received nil HTTP response")
		} else {
			t.Logf("HTTP Status Code: %d", httpRes.StatusCode)
		}
		if resp == nil {
			t.Error("Received nil response")
		}

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListTopics", func(t *testing.T) {

		// 		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.ListTopics(context.Background()).Execute()

		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if httpRes == nil {
			t.Error("Received nil HTTP response")
		} else {
			t.Logf("HTTP Status Code: %d", httpRes.StatusCode)
		}
		if resp == nil {
			t.Error("Received nil response")
		}

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListWorks", func(t *testing.T) {

		// 		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.ListWorks(context.Background()).Execute()

		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if httpRes == nil {
			t.Error("Received nil HTTP response")
		} else {
			t.Logf("HTTP Status Code: %d", httpRes.StatusCode)
		}
		if resp == nil {
			t.Error("Received nil response")
		}

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
