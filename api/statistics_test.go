/*
AhaSend API v2

Testing StatisticsAPIService

*/

package api

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ahasend_StatisticsAPIService(t *testing.T) {
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
			assert.True(t, httpRes.StatusCode >= 200 && httpRes.StatusCode < 300, "Expected 2xx success status code, got %d", httpRes.StatusCode)
		} else {
			// Print error details for debugging
			t.Logf("API call failed with error: %v", err)
			if httpRes != nil {
				t.Logf("HTTP Status Code: %d", httpRes.StatusCode)
				// If we got a 2xx status code but still have an error, this is a real problem that should fail the test
				if httpRes.StatusCode >= 200 && httpRes.StatusCode < 300 {
					t.Errorf("Response parsing failed despite 2xx status code (%d) - this indicates a schema mismatch or parsing issue: %v", httpRes.StatusCode, err)
					return
				}
				assert.True(t, httpRes.StatusCode >= 400 && httpRes.StatusCode < 500, "Expected 4xx error status code, got %d", httpRes.StatusCode)
			}
		}
	}

	t.Run("Test StatisticsAPIService GetBounceStatistics", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()

		params := requests.GetBounceStatisticsParams{}

		resp, httpRes, err := apiClient.StatisticsAPI.GetBounceStatistics(auth, accountId, params)

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test StatisticsAPIService GetDeliverabilityStatistics", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()

		resp, httpRes, err := apiClient.StatisticsAPI.GetDeliverabilityStatistics(auth, accountId, requests.GetDeliverabilityStatisticsParams{})

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test StatisticsAPIService GetDeliveryTimeStatistics", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()

		resp, httpRes, err := apiClient.StatisticsAPI.GetDeliveryTimeStatistics(auth, accountId, requests.GetDeliveryTimeStatisticsParams{})

		validatePrismResponse(t, resp, httpRes, err)

	})

}
