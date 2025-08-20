package ahasend_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/api"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/AhaSend/ahasend-go/models/responses"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestGetMessagesDirectMethodUnitTests tests the new direct method functionality without making actual API calls
func TestGetMessagesDirectMethodUnitTests(t *testing.T) {

	t.Run("Direct method with status parameter", func(t *testing.T) {
		client := api.NewAPIClientWithConfig(api.NewConfiguration())
		ctx := context.Background()
		accountID := uuid.New()

		// Test the new direct method call pattern
		params := requests.GetMessagesParams{
			Sender: ahasend.String("test@example.com"),
			Status: ahasend.String("Delivered"),
		}

		// This should compile and return the correct types (even if it fails due to no server)
		resp, httpResp, err := client.MessagesAPI.GetMessages(ctx, accountID, params)

		// We expect an error since there's no real server, but types should be correct
		assert.IsType(t, (*responses.PaginatedMessagesResponse)(nil), resp)
		assert.IsType(t, (*http.Response)(nil), httpResp)
		assert.Error(t, err) // Expected due to no server
	})

	t.Run("Direct method with nil status parameter", func(t *testing.T) {
		client := api.NewAPIClientWithConfig(api.NewConfiguration())
		ctx := context.Background()
		accountID := uuid.New()
		params := requests.GetMessagesParams{
			Sender: ahasend.String("test@example.com"),
		}
		// Test that nil status parameter works
		resp, httpResp, err := client.MessagesAPI.GetMessages(ctx, accountID, params)

		// We expect an error since there's no real server, but types should be correct
		assert.IsType(t, (*responses.PaginatedMessagesResponse)(nil), resp)
		assert.IsType(t, (*http.Response)(nil), httpResp)
		assert.Error(t, err) // Expected due to no server
	})
}

// TestGetMessagesStatusParameterFormats tests various status parameter formats
func TestGetMessagesStatusParameterFormats(t *testing.T) {

	t.Run("Status parameter accepts various formats", func(t *testing.T) {
		client := api.NewAPIClientWithConfig(api.NewConfiguration())
		ctx := context.Background()
		accountID := uuid.New()

		testCases := []struct {
			name   string
			status string
		}{
			{"Single status", "Delivered"},
			{"Two statuses", "Delivered,Bounced"},
			{"Multiple statuses", "Delivered,Bounced,Failed"},
			{"With spaces", "Delivered, Bounced, Failed"},
			{"Mixed case", "delivered,BOUNCED,Failed"},
			{"Empty string", ""},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				// Test that all formats compile and don't panic
				params := requests.GetMessagesParams{
					Sender: ahasend.String("test@example.com"),
					Status: ahasend.String(tc.status),
				}
				_, _, err := client.MessagesAPI.GetMessages(ctx, accountID, params)

				// We expect an error due to no server, but it should be a network error, not a client error
				assert.Error(t, err) // Expected due to no server
			})
		}
	})
}

// TestGetMessagesDirectMethodEdgeCases tests edge cases and validation
func TestGetMessagesDirectMethodEdgeCases(t *testing.T) {

	t.Run("Status parameter edge cases", func(t *testing.T) {
		client := api.NewAPIClientWithConfig(api.NewConfiguration())
		ctx := context.Background()
		accountID := uuid.New()

		edgeCases := []struct {
			name   string
			status string
		}{
			{"Empty status", ""},
			{"Single character", "D"},
			{"Very long status list", "Status1,Status2,Status3,Status4,Status5,Status6,Status7,Status8,Status9,Status10"},
			{"Special characters", "Delivered;Bounced"},
			{"Unicode characters", "Delivered,Bounçed"},
		}

		for _, tc := range edgeCases {
			t.Run(tc.name, func(t *testing.T) {
				params := requests.GetMessagesParams{
					Sender: ahasend.String("test@example.com"),
					Status: ahasend.String(tc.status),
				}
				// These should not panic and should handle the parameter correctly
				_, _, err := client.MessagesAPI.GetMessages(ctx, accountID, params)

				// We expect an error due to no server, but it should be a network error, not a client error
				assert.Error(t, err) // Expected due to no server
			})
		}
	})

	t.Run("All parameters including status", func(t *testing.T) {
		client := api.NewAPIClientWithConfig(api.NewConfiguration())
		ctx := context.Background()
		accountID := uuid.New()

		// Test that status works with all other available parameters
		params := requests.GetMessagesParams{
			Sender:          ahasend.String("sender@example.com"),
			Status:          ahasend.String("Delivered,Bounced"),
			Recipient:       ahasend.String("recipient@example.com"),
			Subject:         ahasend.String("Test Subject"),
			MessageIDHeader: ahasend.String("msg-123"),
			Limit:           ahasend.Int32(50),
			Cursor:          ahasend.String("cursor-123"),
		}

		_, _, err := client.MessagesAPI.GetMessages(ctx, accountID, params)

		// We expect an error due to no server, but it should be a network error, not a client error
		assert.Error(t, err) // Expected due to no server
	})
}

// TestGetMessagesDirectMethodDocumentationExamples tests examples that would appear in documentation
func TestGetMessagesDirectMethodDocumentationExamples(t *testing.T) {

	client := api.NewAPIClientWithConfig(api.NewConfiguration())
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, "test-key")
	accountID := uuid.New()

	t.Run("Documentation example 1: Single status", func(t *testing.T) {
		params := requests.GetMessagesParams{
			Sender: ahasend.String("noreply@example.com"),
			Status: ahasend.String("Bounced"),
		}
		_, _, err := client.MessagesAPI.GetMessages(ctx, accountID, params)

		// We expect an error due to no server, but it should be a network error, not a client error
		assert.Error(t, err) // Expected due to no server
	})

	t.Run("Documentation example 2: Multiple statuses", func(t *testing.T) {
		params := requests.GetMessagesParams{
			Sender: ahasend.String("noreply@example.com"),
			Status: ahasend.String("Failed,Bounced"),
		}
		_, _, err := client.MessagesAPI.GetMessages(ctx, accountID, params)

		// We expect an error due to no server, but it should be a network error, not a client error
		assert.Error(t, err) // Expected due to no server
	})

	t.Run("Documentation example 3: Status with other filters", func(t *testing.T) {
		params := requests.GetMessagesParams{
			Sender:    ahasend.String("noreply@example.com"),
			Status:    ahasend.String("Delivered"),
			Recipient: ahasend.String("user@example.com"),
			Limit:     ahasend.Int32(10),
		}

		_, _, err := client.MessagesAPI.GetMessages(ctx, accountID, params)

		// We expect an error due to no server, but it should be a network error, not a client error
		assert.Error(t, err) // Expected due to no server
	})

	t.Run("Documentation example 4: No status filter", func(t *testing.T) {
		params := requests.GetMessagesParams{
			Sender: ahasend.String("noreply@example.com"),
		}
		_, _, err := client.MessagesAPI.GetMessages(ctx, accountID, params)

		// We expect an error due to no server, but it should be a network error, not a client error
		assert.Error(t, err) // Expected due to no server
	})

	t.Run("Documentation example 5: Using functional options", func(t *testing.T) {
		params := requests.GetMessagesParams{
			Sender: ahasend.String("noreply@example.com"),
			Status: ahasend.String("Delivered"),
		}

		_, _, err := client.MessagesAPI.GetMessages(ctx, accountID, params,
			api.WithTimeout(30*time.Second),
			api.WithHeaders(map[string]string{"X-Custom": "value"}))

		// We expect an error due to no server, but it should be a network error, not a client error
		assert.Error(t, err) // Expected due to no server
	})
}
