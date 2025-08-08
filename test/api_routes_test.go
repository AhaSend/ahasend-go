/*
AhaSend API v2

Testing RoutesAPIService

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

func Test_ahasend_RoutesAPIService(t *testing.T) {

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
		createRouteRequest := ahasend.CreateRouteRequest{
			Name: "Test Route",
			Url:  "https://example.com/webhook",
		}

		resp, httpRes, err := apiClient.RoutesAPI.CreateRoute(auth, accountId).CreateRouteRequest(createRouteRequest).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test RoutesAPIService DeleteRoute", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		routeId := uuid.New()

		resp, httpRes, err := apiClient.RoutesAPI.DeleteRoute(auth, accountId, routeId).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test RoutesAPIService GetRoute", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		routeId := uuid.New()

		resp, httpRes, err := apiClient.RoutesAPI.GetRoute(auth, accountId, routeId).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test RoutesAPIService GetRoutes", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()

		resp, httpRes, err := apiClient.RoutesAPI.GetRoutes(auth, accountId).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test RoutesAPIService UpdateRoute", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		routeId := uuid.New()
		updateRouteRequest := ahasend.UpdateRouteRequest{
			Name: ahasend.PtrString("Updated Test Route"),
		}

		resp, httpRes, err := apiClient.RoutesAPI.UpdateRoute(auth, accountId, routeId).UpdateRouteRequest(updateRouteRequest).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

}
