/*
AhaSend API v2

Testing APIKeysAPIService

*/

package ahasend_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/AhaSend/ahasend-go"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ahasend_APIKeysAPIService(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION_TESTS") == "true" {
		t.Skip("Skipping API integration tests (SKIP_INTEGRATION_TESTS=true)")
	}

	configuration := ahasend.NewConfiguration()
	configuration.Host = "localhost:4010" // Point to Prism mock server
	configuration.Scheme = "http"         // Use HTTP for mock server
	apiClient := ahasend.NewAPIClient(configuration)

	// Create authentication context
	auth := context.WithValue(context.Background(), ahasend.ContextAccessToken, "test-api-key")

	// Helper function to validate Prism responses
	validatePrismResponse := func(t *testing.T, resp interface{}, httpRes *http.Response, err error) {
		if err == nil {
			require.NotNil(t, resp)
			assert.True(t, httpRes.StatusCode >= 200 && httpRes.StatusCode < 299, "Expected valid HTTP status code, got %d", httpRes.StatusCode)
		} else {
			require.NotNil(t, httpRes)
			assert.True(t, httpRes.StatusCode >= 400 && httpRes.StatusCode < 500, "Expected 4xx error status code, got %d", httpRes.StatusCode)
		}
	}

	t.Run("Test APIKeysAPIService CreateAPIKey", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		createAPIKeyRequest := ahasend.CreateAPIKeyRequest{
			Label:  "Test API Key",
			Scopes: []string{"messages:send"},
		}

		resp, httpRes, err := apiClient.APIKeysAPI.CreateAPIKey(auth, accountId).CreateAPIKeyRequest(createAPIKeyRequest).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test APIKeysAPIService DeleteAPIKey", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		keyId := uuid.New()

		resp, httpRes, err := apiClient.APIKeysAPI.DeleteAPIKey(auth, accountId, keyId).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test APIKeysAPIService GetAPIKey", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		keyId := uuid.New()

		resp, httpRes, err := apiClient.APIKeysAPI.GetAPIKey(auth, accountId, keyId).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test APIKeysAPIService GetAPIKeys", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()

		resp, httpRes, err := apiClient.APIKeysAPI.GetAPIKeys(auth, accountId).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test APIKeysAPIService UpdateAPIKey", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		keyId := uuid.New()
		updateAPIKeyRequest := ahasend.UpdateAPIKeyRequest{
			Label: ahasend.PtrString("Updated API Key"),
		}

		resp, httpRes, err := apiClient.APIKeysAPI.UpdateAPIKey(auth, accountId, keyId).UpdateAPIKeyRequest(updateAPIKeyRequest).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

}
