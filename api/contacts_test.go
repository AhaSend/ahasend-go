/*
AhaSend API v2

Testing ContactsAPIService against the OpenAPI Prism mock.
*/

package api

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/AhaSend/ahasend-go/internal/prismmock"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ahasend_ContactsAPIService(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION_TESTS") == "true" {
		t.Skip("Skipping API integration tests (SKIP_INTEGRATION_TESTS=true)")
	}
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	configuration := NewConfiguration()
	configuration.Host = prismmock.Addr()
	configuration.Scheme = "http"
	apiClient := NewAPIClientWithConfig(configuration)
	auth := context.WithValue(context.Background(), ContextAccessToken, "test-api-key")
	accountID := uuid.MustParse("aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa")
	rawEmail := "User+Tag@Example.COM"

	// Prism answers each operation with its first documented 2xx example.
	validatePrismResponse := func(t *testing.T, wantStatus int, response any, httpResponse *http.Response, err error) {
		t.Helper()
		require.NoError(t, err)
		require.NotNil(t, httpResponse)
		assert.Equal(t, wantStatus, httpResponse.StatusCode)
		assert.NotNil(t, response)
	}

	t.Run("GetContacts", func(t *testing.T) {
		response, httpResponse, err := apiClient.ContactsAPI.GetContacts(auth, accountID, requests.GetContactsParams{})
		validatePrismResponse(t, http.StatusOK, response, httpResponse, err)
	})

	t.Run("GetContact", func(t *testing.T) {
		response, httpResponse, err := apiClient.ContactsAPI.GetContact(auth, accountID, rawEmail)
		validatePrismResponse(t, http.StatusOK, response, httpResponse, err)
	})

	t.Run("CreateContact", func(t *testing.T) {
		response, httpResponse, err := apiClient.ContactsAPI.CreateContact(auth, accountID, requests.CreateContactRequest{
			Email: rawEmail,
		}, WithIdempotencyKey("prism-create-contact"))
		validatePrismResponse(t, http.StatusCreated, response, httpResponse, err)
	})

	t.Run("UpdateContact", func(t *testing.T) {
		response, httpResponse, err := apiClient.ContactsAPI.UpdateContact(auth, accountID, rawEmail, requests.UpdateContactRequest{
			Attributes: map[string]any{"obsolete": nil},
		})
		validatePrismResponse(t, http.StatusOK, response, httpResponse, err)
	})

	t.Run("DeleteContact", func(t *testing.T) {
		response, httpResponse, err := apiClient.ContactsAPI.DeleteContact(auth, accountID, rawEmail)
		validatePrismResponse(t, http.StatusOK, response, httpResponse, err)
	})

	t.Run("BatchUpsertContacts", func(t *testing.T) {
		response, httpResponse, err := apiClient.ContactsAPI.BatchUpsertContacts(auth, accountID, requests.BatchUpsertContactsRequest{
			Data: []requests.BatchUpsertContactInput{{Email: rawEmail, Attributes: map[string]any{"obsolete": nil}}},
		}, WithIdempotencyKey("prism-batch-contacts"))
		validatePrismResponse(t, http.StatusOK, response, httpResponse, err)
	})
}
