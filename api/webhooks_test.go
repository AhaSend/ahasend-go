/*
AhaSend API v2

Testing WebhooksAPIService

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

func Test_ahasend_WebhooksAPIService(t *testing.T) {
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

	t.Run("Test WebhooksAPIService CreateWebhook", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		createWebhookRequest := requests.CreateWebhookRequest{
			Name: "Test Webhook",
			URL:  "https://example.com/webhook",
		}

		resp, httpRes, err := apiClient.WebhooksAPI.CreateWebhook(auth, accountId, createWebhookRequest)

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test WebhooksAPIService DeleteWebhook", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		webhookId := uuid.New()

		resp, httpRes, err := apiClient.WebhooksAPI.DeleteWebhook(auth, accountId, webhookId)

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test WebhooksAPIService GetWebhook", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		webhookId := uuid.New()

		resp, httpRes, err := apiClient.WebhooksAPI.GetWebhook(auth, accountId, webhookId)

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test WebhooksAPIService GetWebhooks", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()

		params := GetWebhooksParams{}
		resp, httpRes, err := apiClient.WebhooksAPI.GetWebhooks(auth, accountId, params)

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test WebhooksAPIService UpdateWebhook", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		webhookId := uuid.New()
		name := "Updated Test Webhook"
		updateWebhookRequest := requests.UpdateWebhookRequest{
			Name: &name,
		}

		resp, httpRes, err := apiClient.WebhooksAPI.UpdateWebhook(auth, accountId, webhookId, updateWebhookRequest)

		validatePrismResponse(t, resp, httpRes, err)

	})

}
