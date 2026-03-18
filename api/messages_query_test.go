package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMessagesAPIGetMessagesSerializesTagsAsCommaSeparatedQueryParam(t *testing.T) {
	var lastRequest *http.Request

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lastRequest = r
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"object":"list","data":[],"pagination":{"has_more":false}}`))
	}))
	defer server.Close()

	serverURL, err := url.Parse(server.URL)
	require.NoError(t, err)

	cfg := NewConfiguration()
	cfg.Host = serverURL.Host
	cfg.Scheme = serverURL.Scheme
	cfg.APIKey = "test-key"

	client := NewAPIClientWithConfig(cfg)
	accountID := uuid.New()
	status := "Delivered"

	resp, httpResp, err := client.MessagesAPI.GetMessages(context.Background(), accountID, requests.GetMessagesParams{
		Status: &status,
		Tags:   []string{"billing", "urgent"},
		PaginationParams: common.PaginationParams{
			Limit: ahasend.Int32(5),
		},
	})

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, httpResp)
	require.NotNil(t, lastRequest)

	query := lastRequest.URL.Query()
	assert.Equal(t, "Delivered", query.Get("status"))
	assert.Equal(t, "billing,urgent", query.Get("tags"))
	assert.Equal(t, "5", query.Get("limit"))
}

func TestMessagesAPIGetMessagesOmitsEmptyTagsQueryParam(t *testing.T) {
	var lastRequest *http.Request

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lastRequest = r
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"object":"list","data":[],"pagination":{"has_more":false}}`))
	}))
	defer server.Close()

	serverURL, err := url.Parse(server.URL)
	require.NoError(t, err)

	cfg := NewConfiguration()
	cfg.Host = serverURL.Host
	cfg.Scheme = serverURL.Scheme
	cfg.APIKey = "test-key"

	client := NewAPIClientWithConfig(cfg)

	_, _, err = client.MessagesAPI.GetMessages(context.Background(), uuid.New(), requests.GetMessagesParams{})

	require.NoError(t, err)
	require.NotNil(t, lastRequest)
	assert.Empty(t, lastRequest.URL.Query().Get("tags"))
}
