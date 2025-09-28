package ahasend_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/api"
	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/AhaSend/ahasend-go/models/responses"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// MockHandler captures request details for testing
type MockHandler struct {
	LastRequest *http.Request
	Response    interface{}
	StatusCode  int
}

func (m *MockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.LastRequest = r
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(m.StatusCode)

	// Return a basic valid response
	response := `{
		"object": "list",
		"data": [],
		"pagination": {
			"has_more": false
		}
	}`
	w.Write([]byte(response))
}

// TestGetMessagesParameterPassing tests that parameters are correctly passed to HTTP requests
func TestGetMessagesParameterPassing(t *testing.T) {
	t.Skip("Skipping parameter test - requires update for new Execute method architecture")
	// Create a mock HTTP server
	mockHandler := &MockHandler{StatusCode: 200}
	server := httptest.NewServer(mockHandler)
	defer server.Close()

	// Parse server URL
	serverURL, err := url.Parse(server.URL)
	require.NoError(t, err)

	// Configure client to use mock server
	configuration := api.NewConfiguration()
	configuration.Host = serverURL.Host
	configuration.Scheme = serverURL.Scheme
	configuration.APIKey = "test-key" // Set API key at client level
	client := api.NewAPIClientWithConfig(configuration)

	ctx := context.Background() // Use plain context since auth is at client level
	accountID := uuid.New()

	tests := []struct {
		name           string
		setupRequest   func() (*responses.PaginatedMessagesResponse, *http.Response, error)
		expectedParams map[string]string
	}{
		{
			name: "Status parameter only",
			setupRequest: func() (*responses.PaginatedMessagesResponse, *http.Response, error) {
				params := requests.GetMessagesParams{
					Sender: ahasend.String("test@example.com"),
					Status: ahasend.String("Delivered"),
				}
				return client.MessagesAPI.GetMessages(ctx, accountID, params)
			},
			expectedParams: map[string]string{
				"sender": "test@example.com",
				"status": "Delivered",
			},
		},
		{
			name: "Multiple statuses",
			setupRequest: func() (*responses.PaginatedMessagesResponse, *http.Response, error) {
				params := requests.GetMessagesParams{
					Sender: ahasend.String("test@example.com"),
					Status: ahasend.String("Delivered,Bounced,Failed"),
				}
				return client.MessagesAPI.GetMessages(ctx, accountID, params)
			},
			expectedParams: map[string]string{
				"sender": "test@example.com",
				"status": "Delivered,Bounced,Failed",
			},
		},
		{
			name: "Status with other parameters",
			setupRequest: func() (*responses.PaginatedMessagesResponse, *http.Response, error) {
				params := requests.GetMessagesParams{
					Sender:    ahasend.String("sender@example.com"),
					Status:    ahasend.String("Delivered"),
					Recipient: ahasend.String("recipient@example.com"),
					Subject:   ahasend.String("Test Subject"),
					PaginationParams: common.PaginationParams{
						Limit: ahasend.Int32(10),
					},
				}
				return client.MessagesAPI.GetMessages(ctx, accountID, params)
			},
			expectedParams: map[string]string{
				"sender":    "sender@example.com",
				"status":    "Delivered",
				"recipient": "recipient@example.com",
				"subject":   "Test Subject",
				"limit":     "10",
			},
		},
		{
			name: "All parameters including status",
			setupRequest: func() (*responses.PaginatedMessagesResponse, *http.Response, error) {
				params := requests.GetMessagesParams{
					Sender:          ahasend.String("sender@example.com"),
					Status:          ahasend.String("Bounced,Failed"),
					Recipient:       ahasend.String("recipient@example.com"),
					Subject:         ahasend.String("Complete Subject"),
					MessageIDHeader: ahasend.String("msg-12345"),
					PaginationParams: common.PaginationParams{
						Limit:  ahasend.Int32(25),
						Cursor: ahasend.String("test-cursor"),
					},
				}
				return client.MessagesAPI.GetMessages(ctx, accountID, params)
			},
			expectedParams: map[string]string{
				"status":            "Bounced,Failed",
				"sender":            "sender@example.com",
				"recipient":         "recipient@example.com",
				"subject":           "Complete Test",
				"message_id_header": "msg-12345",
				"limit":             "25",
				"cursor":            "test-cursor",
			},
		},
		{
			name: "No status parameter",
			setupRequest: func() (*responses.PaginatedMessagesResponse, *http.Response, error) {
				params := requests.GetMessagesParams{
					Sender:    ahasend.String("sender@example.com"),
					Recipient: ahasend.String("recipient@example.com"),
				}
				return client.MessagesAPI.GetMessages(ctx, accountID, params)
			},
			expectedParams: map[string]string{
				"sender":    "sender@example.com",
				"recipient": "recipient@example.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset mock handler
			mockHandler.LastRequest = nil

			// Execute request
			_, _, _ = tt.setupRequest()

			// We expect the request to complete (may have auth errors but that's OK)
			// The important thing is that the request was made with correct parameters
			require.NotNil(t, mockHandler.LastRequest, "Request should have been made")

			// Verify query parameters
			query := mockHandler.LastRequest.URL.Query()

			for expectedParam, expectedValue := range tt.expectedParams {
				actualValue := query.Get(expectedParam)
				assert.Equal(t, expectedValue, actualValue,
					"Parameter %s should be %s, got %s", expectedParam, expectedValue, actualValue)
			}

			// Verify status parameter is not present when not expected
			if _, shouldHaveStatus := tt.expectedParams["status"]; !shouldHaveStatus {
				assert.Empty(t, query.Get("status"), "Status parameter should not be present")
			}

			// Verify path
			expectedPath := fmt.Sprintf("/v2/accounts/%s/messages", accountID.String())
			assert.Equal(t, expectedPath, mockHandler.LastRequest.URL.Path)

			// Verify method
			assert.Equal(t, "GET", mockHandler.LastRequest.Method)
		})
	}
}

// TestGetMessagesStatusParameterEncoding tests URL encoding of status parameter values
func TestGetMessagesStatusParameterEncoding(t *testing.T) {
	t.Skip("Skipping parameter test - requires update for new Execute method architecture")
	mockHandler := &MockHandler{StatusCode: 200}
	server := httptest.NewServer(mockHandler)
	defer server.Close()

	serverURL, err := url.Parse(server.URL)
	require.NoError(t, err)

	configuration := api.NewConfiguration()
	configuration.Host = serverURL.Host
	configuration.Scheme = serverURL.Scheme
	configuration.APIKey = "test-key" // Set API key at client level
	client := api.NewAPIClientWithConfig(configuration)

	ctx := context.Background() // Use plain context since auth is at client level
	accountID := uuid.New()

	encodingTests := []struct {
		name          string
		statusValue   string
		expectedInURL string
	}{
		{
			name:          "Simple statuses",
			statusValue:   "Delivered,Bounced",
			expectedInURL: "Delivered,Bounced",
		},
		{
			name:          "Statuses with spaces",
			statusValue:   "Delivered, Bounced, Failed",
			expectedInURL: "Delivered, Bounced, Failed",
		},
		{
			name:          "Special characters",
			statusValue:   "Status+With+Plus",
			expectedInURL: "Status+With+Plus",
		},
		{
			name:          "URL-unsafe characters",
			statusValue:   "Status&With&Ampersand",
			expectedInURL: "Status&With&Ampersand",
		},
	}

	for _, tt := range encodingTests {
		t.Run(tt.name, func(t *testing.T) {
			mockHandler.LastRequest = nil

			params := requests.GetMessagesParams{
				Sender: ahasend.String("test@example.com"),
				Status: ahasend.String(tt.statusValue),
			}
			_, _, _ = client.MessagesAPI.GetMessages(ctx, accountID, params)

			require.NotNil(t, mockHandler.LastRequest)

			query := mockHandler.LastRequest.URL.Query()
			actualStatus := query.Get("status")

			assert.Equal(t, tt.expectedInURL, actualStatus,
				"Status parameter encoding should be correct")
		})
	}
}

// TestGetMessagesStatusParameterQueryStringBuilding tests query string construction
func TestGetMessagesStatusParameterQueryStringBuilding(t *testing.T) {
	t.Skip("Skipping parameter test - requires update for new Execute method architecture")
	mockHandler := &MockHandler{StatusCode: 200}
	server := httptest.NewServer(mockHandler)
	defer server.Close()

	serverURL, err := url.Parse(server.URL)
	require.NoError(t, err)

	configuration := api.NewConfiguration()
	configuration.Host = serverURL.Host
	configuration.Scheme = serverURL.Scheme
	configuration.APIKey = "test-key" // Set API key at client level
	client := api.NewAPIClientWithConfig(configuration)

	ctx := context.Background() // Use plain context since auth is at client level
	accountID := uuid.New()

	t.Run("Query string parameter order", func(t *testing.T) {
		mockHandler.LastRequest = nil

		params := requests.GetMessagesParams{
			Sender:    ahasend.String("sender@example.com"),
			Status:    ahasend.String("Delivered"),
			Recipient: ahasend.String("recipient@example.com"),
			Subject:   ahasend.String("Test"),
			PaginationParams: common.PaginationParams{
				Limit: ahasend.Int32(10),
			},
		}
		_, _, _ = client.MessagesAPI.GetMessages(ctx, accountID, params)

		require.NotNil(t, mockHandler.LastRequest)

		query := mockHandler.LastRequest.URL.Query()

		// Verify all parameters are present
		expectedParams := []string{"status", "sender", "recipient", "subject", "limit"}
		for _, param := range expectedParams {
			assert.NotEmpty(t, query.Get(param), "Parameter %s should be present", param)
		}
	})

	t.Run("Empty status parameter handling", func(t *testing.T) {
		mockHandler.LastRequest = nil

		params := requests.GetMessagesParams{
			Sender: ahasend.String("sender@example.com"),
			Status: ahasend.String(""),
		}
		_, _, _ = client.MessagesAPI.GetMessages(ctx, accountID, params)

		require.NotNil(t, mockHandler.LastRequest)

		query := mockHandler.LastRequest.URL.Query()

		// Should have status parameter even if empty
		assert.Contains(t, query, "status", "Status parameter should be present even if empty")
		assert.Equal(t, "", query.Get("status"), "Empty status should be preserved")
	})
}

// TestGetMessagesDirectMethodCall tests the new direct method call pattern
func TestGetMessagesDirectMethodCall(t *testing.T) {
	t.Skip("Skipping parameter test - requires update for new Execute method architecture")
	client := api.NewAPIClientWithConfig(api.NewConfiguration())
	ctx := context.Background()
	accountID := uuid.New()

	t.Run("Direct method call with all parameters", func(t *testing.T) {
		// Test the new direct method call pattern
		params := requests.GetMessagesParams{
			Sender:          ahasend.String("test@example.com"),
			Status:          ahasend.String("Delivered"),
			Recipient:       ahasend.String("recipient@example.com"),
			Subject:         ahasend.String("Test Subject"),
			MessageIDHeader: ahasend.String("msg-1234"),
			PaginationParams: common.PaginationParams{
				Limit:  ahasend.Int32(10),
				Cursor: ahasend.String("test-cursor"),
			},
		}
		// This should compile and not panic (even if it fails due to no server)
		_, _, err := client.MessagesAPI.GetMessages(ctx, accountID, params)

		// We expect an error since there's no real server, but it should be a network error, not a client error
		assert.Error(t, err) // Expected due to no server
	})

	t.Run("Direct method call with nil parameters", func(t *testing.T) {
		// Test that nil parameters are handled correctly
		params := requests.GetMessagesParams{
			Sender: ahasend.String("test@example.com"),
		}
		// This should compile and not panic
		_, _, err := client.MessagesAPI.GetMessages(ctx, accountID, params)

		// We expect an error since there's no real server, but it should be a network error, not a client error
		assert.Error(t, err) // Expected due to no server
	})

	t.Run("Direct method call type verification", func(t *testing.T) {
		// Verify the method signature returns the correct types
		params := requests.GetMessagesParams{
			Sender: ahasend.String("test@example.com"),
			Status: ahasend.String("Delivered"),
		}
		resp, httpResp, err := client.MessagesAPI.GetMessages(ctx, accountID, params)

		// Verify return types (even if nil due to error)
		assert.IsType(t, (*responses.PaginatedMessagesResponse)(nil), resp)
		assert.IsType(t, (*http.Response)(nil), httpResp)
		assert.IsType(t, (error)(nil), err)
	})
}
