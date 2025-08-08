/*
AhaSend API v2

Testing WebhooksAPIService

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

func Test_ahasend_WebhooksAPIService(t *testing.T) {

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

	t.Run("Test WebhooksAPIService CreateWebhook", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		createWebhookRequest := ahasend.CreateWebhookRequest{
			Name: "Test Webhook",
			Url:  "https://example.com/webhook",
		}

		resp, httpRes, err := apiClient.WebhooksAPI.CreateWebhook(auth, accountId).CreateWebhookRequest(createWebhookRequest).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test WebhooksAPIService DeleteWebhook", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		webhookId := uuid.New()

		resp, httpRes, err := apiClient.WebhooksAPI.DeleteWebhook(auth, accountId, webhookId).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test WebhooksAPIService GetWebhook", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		webhookId := uuid.New()

		resp, httpRes, err := apiClient.WebhooksAPI.GetWebhook(auth, accountId, webhookId).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test WebhooksAPIService GetWebhooks", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()

		resp, httpRes, err := apiClient.WebhooksAPI.GetWebhooks(auth, accountId).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test WebhooksAPIService UpdateWebhook", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		webhookId := uuid.New()
		updateWebhookRequest := ahasend.UpdateWebhookRequest{
			Name: ahasend.PtrString("Updated Test Webhook"),
		}

		resp, httpRes, err := apiClient.WebhooksAPI.UpdateWebhook(auth, accountId, webhookId).UpdateWebhookRequest(updateWebhookRequest).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

}
