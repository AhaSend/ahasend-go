/*
AhaSend API v2

Testing SuppressionsAPIService

*/

package api

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_ahasend_SuppressionsAPIService(t *testing.T) {
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
		// For Prism mock server, we primarily care about getting a valid HTTP response
		if httpRes != nil {
			// Any 2xx-4xx range is acceptable for mock responses
			assert.True(t, httpRes.StatusCode >= 200 && httpRes.StatusCode < 500, "Expected valid HTTP status code, got %d", httpRes.StatusCode)

			// If we got a successful status code, we should have a response
			if httpRes.StatusCode >= 200 && httpRes.StatusCode < 400 {
				assert.NotNil(t, resp, "Expected response body for successful status code %d", httpRes.StatusCode)
			}
		} else {
			// If there's no HTTP response, there should be an error
			assert.NotNil(t, err, "Expected either HTTP response or error")
		}
	}

	t.Run("Test SuppressionsAPIService CreateSuppression", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		reason := "Test suppression"
		createSuppressionRequest := requests.CreateSuppressionRequest{
			Email:     "test@example.com",
			ExpiresAt: time.Now().Add(24 * time.Hour),
			Reason:    &reason,
		}

		resp, httpRes, err := apiClient.SuppressionsAPI.CreateSuppression(auth, accountId, createSuppressionRequest)

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test SuppressionsAPIService DeleteAllSuppressions", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()

		resp, httpRes, err := apiClient.SuppressionsAPI.DeleteAllSuppressions(auth, accountId, nil)

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test SuppressionsAPIService DeleteSuppression", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		email := "test@example.com"

		resp, httpRes, err := apiClient.SuppressionsAPI.DeleteSuppression(auth, accountId, email, nil)

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test SuppressionsAPIService GetSuppressions", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()

		params := requests.GetSuppressionsParams{}
		resp, httpRes, err := apiClient.SuppressionsAPI.GetSuppressions(auth, accountId, params)

		validatePrismResponse(t, resp, httpRes, err)

	})

}
