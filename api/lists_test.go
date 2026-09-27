/*
AhaSend API v2

Testing ListsAPIService against the OpenAPI Prism mock.
*/

package api

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/internal/prismmock"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ahasend_ListsAPIService(t *testing.T) {
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
	listID := uuid.MustParse("22222222-2222-4222-8222-222222222222")
	contactID := uuid.MustParse("11111111-1111-4111-8111-111111111111")
	rawEmail := "User+Tag@Example.COM"

	// Prism answers each operation with its first documented 2xx example.
	validatePrismResponse := func(t *testing.T, wantStatus int, response any, httpResponse *http.Response, err error) {
		t.Helper()
		require.NoError(t, err)
		require.NotNil(t, httpResponse)
		assert.Equal(t, wantStatus, httpResponse.StatusCode)
		assert.NotNil(t, response)
	}

	t.Run("GetLists", func(t *testing.T) {
		response, httpResponse, err := apiClient.ListsAPI.GetLists(auth, accountID, requests.GetListsParams{})
		validatePrismResponse(t, http.StatusOK, response, httpResponse, err)
	})

	t.Run("CreateList", func(t *testing.T) {
		response, httpResponse, err := apiClient.ListsAPI.CreateList(auth, accountID, requests.CreateListRequest{
			Name: "Newsletter",
			Tags: []string{"news"},
		}, WithIdempotencyKey("prism-create-list"))
		validatePrismResponse(t, http.StatusCreated, response, httpResponse, err)
	})

	t.Run("GetList", func(t *testing.T) {
		response, httpResponse, err := apiClient.ListsAPI.GetList(auth, accountID, listID)
		validatePrismResponse(t, http.StatusOK, response, httpResponse, err)
	})

	t.Run("UpdateList", func(t *testing.T) {
		response, httpResponse, err := apiClient.ListsAPI.UpdateList(auth, accountID, listID, requests.UpdateListRequest{
			Description: ahasend.String(""),
			Tags:        &[]string{},
		})
		validatePrismResponse(t, http.StatusOK, response, httpResponse, err)
	})

	t.Run("DeleteList", func(t *testing.T) {
		response, httpResponse, err := apiClient.ListsAPI.DeleteList(auth, accountID, listID)
		validatePrismResponse(t, http.StatusOK, response, httpResponse, err)
	})

	t.Run("GetListContacts", func(t *testing.T) {
		response, httpResponse, err := apiClient.ListsAPI.GetListContacts(auth, accountID, listID, requests.GetListContactsParams{
			SubscriptionStatus: ahasend.String(requests.ListContactStatusConfirmed),
			IncludeContacts:    ahasend.Bool(true),
		})
		validatePrismResponse(t, http.StatusOK, response, httpResponse, err)
	})

	t.Run("BatchAddListContacts", func(t *testing.T) {
		response, httpResponse, err := apiClient.ListsAPI.BatchAddListContacts(auth, accountID, listID, requests.BatchAddListContactsRequest{
			Data: []requests.BatchAddListContactInput{{Email: ahasend.String(rawEmail)}, {ID: &contactID}},
		}, WithIdempotencyKey("prism-batch-list-contacts"))
		validatePrismResponse(t, http.StatusOK, response, httpResponse, err)
	})

	t.Run("UpsertListContact", func(t *testing.T) {
		response, httpResponse, err := apiClient.ListsAPI.UpsertListContact(auth, accountID, listID, rawEmail, requests.UpsertListContactRequest{
			SubscriptionStatus: ahasend.String(requests.ListContactStatusUnsubscribed),
		})
		validatePrismResponse(t, http.StatusOK, response, httpResponse, err)
	})

	t.Run("DeleteListContact", func(t *testing.T) {
		response, httpResponse, err := apiClient.ListsAPI.DeleteListContact(auth, accountID, listID, rawEmail)
		validatePrismResponse(t, http.StatusOK, response, httpResponse, err)
	})

	t.Run("GetContactLists", func(t *testing.T) {
		response, httpResponse, err := apiClient.ListsAPI.GetContactLists(auth, accountID, rawEmail, requests.GetContactListsParams{})
		validatePrismResponse(t, http.StatusOK, response, httpResponse, err)
	})
}
