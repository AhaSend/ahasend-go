/*
AhaSend API v2

Testing SuppressionsAPIService

*/

package ahasend_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/AhaSend/ahasend-go"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ahasend_SuppressionsAPIService(t *testing.T) {

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

	t.Run("Test SuppressionsAPIService CreateSuppression", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		createSuppressionRequest := ahasend.CreateSuppressionRequest{
			Email:     "test@example.com",
			ExpiresAt: time.Now().Add(24 * time.Hour),
			Reason:    ahasend.PtrString("Test suppression"),
		}

		resp, httpRes, err := apiClient.SuppressionsAPI.CreateSuppression(auth, accountId).CreateSuppressionRequest(createSuppressionRequest).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test SuppressionsAPIService DeleteAllSuppressions", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()

		resp, httpRes, err := apiClient.SuppressionsAPI.DeleteAllSuppressions(auth, accountId).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test SuppressionsAPIService DeleteSuppression", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		email := "test@example.com"

		resp, httpRes, err := apiClient.SuppressionsAPI.DeleteSuppression(auth, accountId, email).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test SuppressionsAPIService GetSuppressions", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()

		resp, httpRes, err := apiClient.SuppressionsAPI.GetSuppressions(auth, accountId).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

}
