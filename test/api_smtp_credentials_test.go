/*
AhaSend API v2

Testing SMTPCredentialsAPIService

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

func Test_ahasend_SMTPCredentialsAPIService(t *testing.T) {
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

	t.Run("Test SMTPCredentialsAPIService CreateSMTPCredential", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		createSMTPCredentialRequest := ahasend.CreateSMTPCredentialRequest{
			Name:     "Test SMTP Credential",
			Username: "test-user",
			Password: "test-password",
			Scope:    "global",
		}

		resp, httpRes, err := apiClient.SMTPCredentialsAPI.CreateSMTPCredential(auth, accountId).CreateSMTPCredentialRequest(createSMTPCredentialRequest).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test SMTPCredentialsAPIService DeleteSMTPCredential", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		smtpCredentialId := uuid.New()

		resp, httpRes, err := apiClient.SMTPCredentialsAPI.DeleteSMTPCredential(auth, accountId, smtpCredentialId).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test SMTPCredentialsAPIService GetSMTPCredential", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		smtpCredentialId := uuid.New()

		resp, httpRes, err := apiClient.SMTPCredentialsAPI.GetSMTPCredential(auth, accountId, smtpCredentialId).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test SMTPCredentialsAPIService GetSMTPCredentials", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()

		resp, httpRes, err := apiClient.SMTPCredentialsAPI.GetSMTPCredentials(auth, accountId).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

}
