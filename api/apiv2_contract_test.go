package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newContractTestClient(t *testing.T, handler http.HandlerFunc) (*APIClient, func()) {
	t.Helper()

	server := httptest.NewServer(handler)
	serverURL, err := url.Parse(server.URL)
	require.NoError(t, err)

	cfg := NewConfiguration()
	cfg.Host = serverURL.Host
	cfg.Scheme = serverURL.Scheme
	cfg.APIKey = "test-key"

	return NewAPIClientWithConfig(cfg), server.Close
}

func TestSuppressionsAPIGetSuppressionsUsesTimeQueryNames(t *testing.T) {
	var lastRequest *http.Request
	client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		lastRequest = r
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"object":"list","data":[],"pagination":{"has_more":false}}`))
	})
	defer cleanup()

	fromTime := time.Date(2026, 4, 18, 10, 0, 0, 0, time.UTC)
	toTime := fromTime.Add(time.Hour)
	email := "user@example.com"
	domain := "example.com"

	_, _, err := client.SuppressionsAPI.GetSuppressions(context.Background(), uuid.New(), requests.GetSuppressionsParams{
		Email:    &email,
		Domain:   &domain,
		FromTime: &fromTime,
		ToTime:   &toTime,
	})

	require.NoError(t, err)
	require.NotNil(t, lastRequest)
	query := lastRequest.URL.Query()
	assert.Equal(t, email, query.Get("email"))
	assert.Equal(t, domain, query.Get("domain"))
	assert.Equal(t, fromTime.Format(time.RFC3339), query.Get("from_time"))
	assert.Equal(t, toTime.Format(time.RFC3339), query.Get("to_time"))
	assert.Empty(t, query.Get("from_date"))
	assert.Empty(t, query.Get("to_date"))
}

func TestRoutesAPIGetRoutesWithParamsSerializesDomain(t *testing.T) {
	var lastRequest *http.Request
	client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		lastRequest = r
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"object":"list","data":[],"pagination":{"has_more":false}}`))
	})
	defer cleanup()

	_, _, err := client.RoutesAPI.GetRoutesWithParams(context.Background(), uuid.New(), requests.GetRoutesParams{
		Domain: ahasend.String("example.com"),
		PaginationParams: common.PaginationParams{
			Limit: ahasend.Int32(25),
		},
	})

	require.NoError(t, err)
	require.NotNil(t, lastRequest)
	assert.Equal(t, "example.com", lastRequest.URL.Query().Get("domain"))
	assert.Equal(t, "25", lastRequest.URL.Query().Get("limit"))
}

func TestSMTPCredentialsAPIUsesUUIDPathIDs(t *testing.T) {
	accountID := uuid.New()
	credentialID := uuid.New()
	var lastPath string

	client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		lastPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"object":"credential_smtp","id":"` + credentialID.String() + `","created_at":"2026-04-18T10:00:00Z","updated_at":"2026-04-18T10:00:00Z","name":"prod","username":"user","sandbox":false,"scope":"global","domains":[]}`))
	})
	defer cleanup()

	resp, _, err := client.SMTPCredentialsAPI.GetSMTPCredential(context.Background(), accountID, credentialID)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, credentialID, resp.ID)
	assert.Equal(t, "/v2/accounts/"+accountID.String()+"/smtp-credentials/"+credentialID.String(), lastPath)
}

func TestMessagesAPICreateConversationMessageUsesToField(t *testing.T) {
	accountID := uuid.New()
	var body map[string]interface{}
	client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/v2/accounts/"+accountID.String()+"/messages/conversation", r.URL.Path)
		err := json.NewDecoder(r.Body).Decode(&body)
		require.NoError(t, err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte(`{"object":"list","data":[]}`))
	})
	defer cleanup()

	_, _, err := client.MessagesAPI.CreateConversationMessage(context.Background(), accountID, requests.CreateConversationMessageRequest{
		From:    common.SenderAddress{Email: "sender@example.com"},
		To:      []common.SenderAddress{{Email: "recipient@example.com"}},
		Subject: "Hello",
	})

	require.NoError(t, err)
	require.Contains(t, body, "to")
	assert.NotContains(t, body, "recipients")
}
