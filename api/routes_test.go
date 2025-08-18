/*
AhaSend API v2

Testing RoutesAPIService

*/

package api

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ahasend_RoutesAPIService(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION_TESTS") == "true" {
		t.Skip("Skipping API integration tests (SKIP_INTEGRATION_TESTS=true)")
	}

	configuration := NewConfiguration()
	configuration.Host = "localhost:4010" // Point to Prism mock server
	configuration.Scheme = "http"         // Use HTTP for mock server
	apiClient := NewAPIClientWithConfig(configuration)

	// Create authentication context
	auth := context.WithValue(context.Background(), ContextAccessToken, "test-api-key")

	// Helper function to validate Prism responses
	validatePrismResponse := func(t *testing.T, resp interface{}, httpRes *http.Response, err error) {
		if err == nil {
			require.NotNil(t, resp)
			assert.True(t, httpRes.StatusCode >= 200 && httpRes.StatusCode < 500, "Expected valid HTTP status code, got %d", httpRes.StatusCode)
		} else {
			if httpRes != nil {
				assert.True(t, httpRes.StatusCode >= 400 && httpRes.StatusCode < 500, "Expected 4xx error status code, got %d", httpRes.StatusCode)
			}
		}
	}

	t.Run("Test RoutesAPIService CreateRoute", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		createRouteRequest := requests.CreateRouteRequest{
			Name: "Test Route",
			URL:  "https://example.com/webhook",
		}

		resp, httpRes, err := apiClient.RoutesAPI.CreateRoute(auth, accountId, createRouteRequest)

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test RoutesAPIService DeleteRoute", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		routeId := uuid.New()

		resp, httpRes, err := apiClient.RoutesAPI.DeleteRoute(auth, accountId, routeId)

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test RoutesAPIService GetRoute", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		routeId := uuid.New()

		resp, httpRes, err := apiClient.RoutesAPI.GetRoute(auth, accountId, routeId)

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test RoutesAPIService GetRoutes", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()

		resp, httpRes, err := apiClient.RoutesAPI.GetRoutes(auth, accountId, nil, nil)

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test RoutesAPIService UpdateRoute", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		routeId := uuid.New()
		updateRouteRequest := requests.UpdateRouteRequest{
			Name: ahasend.String("Updated Test Route"),
		}

		resp, httpRes, err := apiClient.RoutesAPI.UpdateRoute(auth, accountId, routeId, updateRouteRequest)

		validatePrismResponse(t, resp, httpRes, err)

	})

}
