/*
AhaSend API v2

Testing MessagesAPIService

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

func Test_ahasend_MessagesAPIService(t *testing.T) {

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

		resp, httpRes, err := apiClient.MessagesAPI.GetMessages(auth, accountId).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

}
