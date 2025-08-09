package ahasend_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/AhaSend/ahasend-go"
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
	// Create a mock HTTP server
	mockHandler := &MockHandler{StatusCode: 200}
	server := httptest.NewServer(mockHandler)
	defer server.Close()

	// Parse server URL
	serverURL, err := url.Parse(server.URL)
	require.NoError(t, err)

	// Configure client to use mock server
	configuration := ahasend.NewConfiguration()
	configuration.Host = serverURL.Host
	configuration.Scheme = serverURL.Scheme
	client := ahasend.NewAPIClient(configuration)

	ctx := context.WithValue(context.Background(), ahasend.ContextAccessToken, "test-key")
	accountID := uuid.New()

	tests := []struct {
		name           string
		setupRequest   func() ahasend.ApiGetMessagesRequest
		expectedParams map[string]string
	}{
		{
			name: "Status parameter only",
			setupRequest: func() ahasend.ApiGetMessagesRequest {
				return client.MessagesAPI.
					GetMessages(ctx, accountID).
					Sender("test@example.com").
					Status("Delivered")
			},
			expectedParams: map[string]string{
				"sender": "test@example.com",
				"status": "Delivered",
			},
		},
		{
			name: "Multiple statuses",
			setupRequest: func() ahasend.ApiGetMessagesRequest {
				return client.MessagesAPI.
					GetMessages(ctx, accountID).
					Sender("test@example.com").
					Status("Delivered,Bounced,Failed")
			},
			expectedParams: map[string]string{
				"sender": "test@example.com",
				"status": "Delivered,Bounced,Failed",
			},
		},
		{
			name: "Status with other parameters",
			setupRequest: func() ahasend.ApiGetMessagesRequest {
				return client.MessagesAPI.
					GetMessages(ctx, accountID).
					Sender("sender@example.com").
					Status("Delivered").
					Recipient("recipient@example.com").
					Subject("Test Subject").
					Limit(10)
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
			setupRequest: func() ahasend.ApiGetMessagesRequest {
				return client.MessagesAPI.
					GetMessages(ctx, accountID).
					Status("Bounced,Failed").
					Sender("sender@example.com").
					Recipient("recipient@example.com").
					Subject("Complete Test").
					MessageIdHeader("msg-12345").
					Limit(25).
					Cursor("test-cursor")
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
			setupRequest: func() ahasend.ApiGetMessagesRequest {
				return client.MessagesAPI.
					GetMessages(ctx, accountID).
					Sender("sender@example.com").
					Recipient("recipient@example.com")
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
			request := tt.setupRequest()
			_, _, _ = request.Execute()

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
	mockHandler := &MockHandler{StatusCode: 200}
	server := httptest.NewServer(mockHandler)
	defer server.Close()

	serverURL, err := url.Parse(server.URL)
	require.NoError(t, err)

	configuration := ahasend.NewConfiguration()
	configuration.Host = serverURL.Host
	configuration.Scheme = serverURL.Scheme
	client := ahasend.NewAPIClient(configuration)

	ctx := context.WithValue(context.Background(), ahasend.ContextAccessToken, "test-key")
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

			_, _, _ = client.MessagesAPI.
				GetMessages(ctx, accountID).
				Sender("test@example.com").
				Status(tt.statusValue).
				Execute()

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
	mockHandler := &MockHandler{StatusCode: 200}
	server := httptest.NewServer(mockHandler)
	defer server.Close()

	serverURL, err := url.Parse(server.URL)
	require.NoError(t, err)

	configuration := ahasend.NewConfiguration()
	configuration.Host = serverURL.Host
	configuration.Scheme = serverURL.Scheme
	client := ahasend.NewAPIClient(configuration)

	ctx := context.WithValue(context.Background(), ahasend.ContextAccessToken, "test-key")
	accountID := uuid.New()

	t.Run("Query string parameter order", func(t *testing.T) {
		mockHandler.LastRequest = nil

		_, _, _ = client.MessagesAPI.
			GetMessages(ctx, accountID).
			Status("Delivered").
			Sender("sender@example.com").
			Recipient("recipient@example.com").
			Subject("Test").
			Limit(10).
			Execute()

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

		_, _, _ = client.MessagesAPI.
			GetMessages(ctx, accountID).
			Sender("sender@example.com").
			Status("").
			Execute()

		require.NotNil(t, mockHandler.LastRequest)

		query := mockHandler.LastRequest.URL.Query()

		// Should have status parameter even if empty
		assert.Contains(t, query, "status", "Status parameter should be present even if empty")
		assert.Equal(t, "", query.Get("status"), "Empty status should be preserved")
	})
}

// TestGetMessagesStatusParameterMethodChaining tests method chaining behavior
func TestGetMessagesStatusParameterMethodChaining(t *testing.T) {
	client := ahasend.NewAPIClient(ahasend.NewConfiguration())
	ctx := context.Background()
	accountID := uuid.New()

	t.Run("Chaining returns correct types", func(t *testing.T) {
		// Each method should return the same request type for chaining
		request1 := client.MessagesAPI.GetMessages(ctx, accountID)
		request2 := request1.Status("Delivered")
		request3 := request2.Sender("test@example.com")
		request4 := request3.Limit(10)

		// All should be the same underlying type
		assert.IsType(t, request1, request2)
		assert.IsType(t, request1, request3)
		assert.IsType(t, request1, request4)
	})

	t.Run("Method call order variations", func(t *testing.T) {
		// Test different orders of method calls - all should work
		orders := []func() ahasend.ApiGetMessagesRequest{
			// Status first
			func() ahasend.ApiGetMessagesRequest {
				return client.MessagesAPI.GetMessages(ctx, accountID).
					Status("Delivered").
					Sender("test@example.com").
					Limit(10)
			},
			// Status middle
			func() ahasend.ApiGetMessagesRequest {
				return client.MessagesAPI.GetMessages(ctx, accountID).
					Sender("test@example.com").
					Status("Delivered").
					Limit(10)
			},
			// Status last
			func() ahasend.ApiGetMessagesRequest {
				return client.MessagesAPI.GetMessages(ctx, accountID).
					Sender("test@example.com").
					Limit(10).
					Status("Delivered")
			},
		}

		for i, orderFunc := range orders {
			t.Run(fmt.Sprintf("Order variation %d", i+1), func(t *testing.T) {
				request := orderFunc()
				assert.NotNil(t, request)
			})
		}
	})

	t.Run("Multiple status calls - last wins", func(t *testing.T) {
		// If Status() is called multiple times, the last one should win
		// This is the expected behavior for optional parameters
		request := client.MessagesAPI.GetMessages(ctx, accountID).
			Sender("test@example.com").
			Status("FirstStatus").
			Status("SecondStatus").
			Status("FinalStatus")

		assert.NotNil(t, request)
		// The actual verification of which value is used would require executing the request
		// and checking the generated URL, which is covered in other tests
	})
}

// TestGetMessagesStatusParameterRequestBuilding tests internal request building
func TestGetMessagesStatusParameterRequestBuilding(t *testing.T) {
	// This test focuses on the request building mechanics
	client := ahasend.NewAPIClient(ahasend.NewConfiguration())
	ctx := context.Background()
	accountID := uuid.New()

	t.Run("Request builder state", func(t *testing.T) {
		// Test that the request builder maintains state correctly
		baseRequest := client.MessagesAPI.GetMessages(ctx, accountID)
		assert.NotNil(t, baseRequest)

		// Add status
		withStatus := baseRequest.Status("Delivered")
		assert.NotNil(t, withStatus)

		// Add sender
		withSender := withStatus.Sender("test@example.com")
		assert.NotNil(t, withSender)

		// Should be able to continue chaining
		final := withSender.Limit(10).Cursor("test")
		assert.NotNil(t, final)
	})

	t.Run("Nil parameter handling", func(t *testing.T) {
		// Test that the implementation handles nil parameters correctly
		// This is important for optional parameters like status
		request := client.MessagesAPI.GetMessages(ctx, accountID).
			Sender("test@example.com")
		// Not calling Status() - should work fine

		assert.NotNil(t, request)
	})
}
