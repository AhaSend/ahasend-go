/*
AhaSend API v2

Testing DomainsAPIService

*/

package ahasend_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/AhaSend/ahasend-go"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ahasend_DomainsAPIService(t *testing.T) {

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
			assert.True(t, httpRes.StatusCode >= 200 && httpRes.StatusCode < 500, "Expected valid HTTP status code, got %d", httpRes.StatusCode)
		} else {
			require.NotNil(t, httpRes)
			assert.True(t, httpRes.StatusCode >= 400 && httpRes.StatusCode < 500, "Expected 4xx error status code, got %d", httpRes.StatusCode)
		}
	}

	t.Run("Test DomainsAPIService CreateDomain", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		createDomainRequest := ahasend.CreateDomainRequest{
			Domain: "example.com",
		}

		resp, httpRes, err := apiClient.DomainsAPI.CreateDomain(auth, accountId).CreateDomainRequest(createDomainRequest).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test DomainsAPIService DeleteDomain", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		domain := "example.com"

		resp, httpRes, err := apiClient.DomainsAPI.DeleteDomain(auth, accountId, domain).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test DomainsAPIService GetDomain", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		domain := "example.com"

		resp, httpRes, err := apiClient.DomainsAPI.GetDomain(auth, accountId, domain).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test DomainsAPIService GetDomains", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()

		resp, httpRes, err := apiClient.DomainsAPI.GetDomains(auth, accountId).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

}
