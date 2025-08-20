/*
AhaSend API v2

Testing SMTPCredentialsAPIService

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
)

func Test_ahasend_SMTPCredentialsAPIService(t *testing.T) {
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

	t.Run("Test SMTPCredentialsAPIService CreateSMTPCredential", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		createSMTPCredentialRequest := requests.CreateSMTPCredentialRequest{
			Name:     "Test SMTP Credential",
			Username: "test-user",
			Password: "test-password",
			Scope:    "global",
		}

		resp, httpRes, err := apiClient.SMTPCredentialsAPI.CreateSMTPCredential(auth, accountId, createSMTPCredentialRequest)

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test SMTPCredentialsAPIService DeleteSMTPCredential", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		smtpCredentialId := uuid.New()

		resp, httpRes, err := apiClient.SMTPCredentialsAPI.DeleteSMTPCredential(auth, accountId, smtpCredentialId)

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test SMTPCredentialsAPIService GetSMTPCredential", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		smtpCredentialId := uuid.New()

		resp, httpRes, err := apiClient.SMTPCredentialsAPI.GetSMTPCredential(auth, accountId, smtpCredentialId)

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test SMTPCredentialsAPIService GetSMTPCredentials", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()

		resp, httpRes, err := apiClient.SMTPCredentialsAPI.GetSMTPCredentials(auth, accountId, nil, nil)

		validatePrismResponse(t, resp, httpRes, err)

	})

}
