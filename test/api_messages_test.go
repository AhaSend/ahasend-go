/*
AhaSend API v2

Testing MessagesAPIService

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

func Test_ahasend_MessagesAPIService(t *testing.T) {
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
			assert.True(t, httpRes.StatusCode >= 200 && httpRes.StatusCode < 500, "Expected valid HTTP status code, got %d", httpRes.StatusCode)
		} else {
			if httpRes != nil {
				assert.True(t, httpRes.StatusCode >= 400 && httpRes.StatusCode < 500, "Expected 4xx error status code, got %d", httpRes.StatusCode)
			}
		}
	}

	t.Run("Test MessagesAPIService CancelMessage", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		messageId := uuid.New().String()

		resp, httpRes, err := apiClient.MessagesAPI.CancelMessage(auth, accountId, messageId).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test MessagesAPIService CreateMessage", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		createMessageRequest := ahasend.CreateMessageRequest{
			From: ahasend.SenderAddress{
				Email: "sender@example.com",
				Name:  ahasend.PtrString("Test Sender"),
			},
			Recipients: []ahasend.Recipient{
				{
					Email: "recipient@example.com",
					Name:  ahasend.PtrString("Test Recipient"),
				},
			},
			Subject:     "Test Subject",
			TextContent: ahasend.PtrString("Test message content"),
		}

		resp, httpRes, err := apiClient.MessagesAPI.CreateMessage(auth, accountId).CreateMessageRequest(createMessageRequest).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test MessagesAPIService GetMessages", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		senderEmail := "test@example.com"

		// Test 1: Basic GetMessages with required sender parameter
		t.Run("Basic GetMessages", func(t *testing.T) {
			resp, httpRes, err := apiClient.MessagesAPI.GetMessages(auth, accountId).
				Sender(senderEmail).
				Execute()

			validatePrismResponse(t, resp, httpRes, err)
		})

		// Test 2: GetMessages with status filter (single status)
		t.Run("GetMessages with single status", func(t *testing.T) {
			resp, httpRes, err := apiClient.MessagesAPI.GetMessages(auth, accountId).
				Sender(senderEmail).
				Status("Delivered").
				Execute()

			validatePrismResponse(t, resp, httpRes, err)
		})

		// Test 3: GetMessages with multiple statuses
		t.Run("GetMessages with multiple statuses", func(t *testing.T) {
			resp, httpRes, err := apiClient.MessagesAPI.GetMessages(auth, accountId).
				Sender(senderEmail).
				Status("Bounced,Failed").
				Execute()

			validatePrismResponse(t, resp, httpRes, err)
		})

		// Test 4: GetMessages with status and other filters
		t.Run("GetMessages with status and recipient filter", func(t *testing.T) {
			resp, httpRes, err := apiClient.MessagesAPI.GetMessages(auth, accountId).
				Sender(senderEmail).
				Status("Delivered").
				Recipient("user@example.com").
				Limit(10).
				Execute()

			validatePrismResponse(t, resp, httpRes, err)
		})

		// Test 5: GetMessages with all possible statuses
		t.Run("GetMessages with all statuses", func(t *testing.T) {
			resp, httpRes, err := apiClient.MessagesAPI.GetMessages(auth, accountId).
				Sender(senderEmail).
				Status("Delivered,Bounced,Failed,Queued,Processing,Suppressed").
				Execute()

			validatePrismResponse(t, resp, httpRes, err)
		})

	})

}
