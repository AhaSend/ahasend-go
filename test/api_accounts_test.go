/*
AhaSend API v2

Testing AccountsAPIService

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

func Test_ahasend_AccountsAPIService(t *testing.T) {
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
			require.NotNil(t, httpRes)
			assert.True(t, httpRes.StatusCode >= 400 && httpRes.StatusCode < 500, "Expected 4xx error status code, got %d", httpRes.StatusCode)
		}
	}

	t.Run("Test AccountsAPIService AddAccountMember", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		addMemberRequest := ahasend.AddMemberRequest{
			Email: "test@example.com",
			Role:  "member",
		}

		resp, httpRes, err := apiClient.AccountsAPI.AddAccountMember(auth, accountId).AddMemberRequest(addMemberRequest).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test AccountsAPIService GetAccount", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()

		resp, httpRes, err := apiClient.AccountsAPI.GetAccount(auth, accountId).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test AccountsAPIService GetAccountMembers", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()

		resp, httpRes, err := apiClient.AccountsAPI.GetAccountMembers(auth, accountId).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test AccountsAPIService RemoveAccountMember", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		userId := uuid.New()

		resp, httpRes, err := apiClient.AccountsAPI.RemoveAccountMember(auth, accountId, userId).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

	t.Run("Test AccountsAPIService UpdateAccount", func(t *testing.T) {

		// Skip test when not running against a real API
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}

		accountId := uuid.New()
		updateAccountRequest := ahasend.UpdateAccountRequest{
			Name: ahasend.PtrString("Updated Account Name"),
		}

		resp, httpRes, err := apiClient.AccountsAPI.UpdateAccount(auth, accountId).UpdateAccountRequest(updateAccountRequest).Execute()

		validatePrismResponse(t, resp, httpRes, err)

	})

}
