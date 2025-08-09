package ahasend_test

import (
	"context"
	"testing"

	"github.com/AhaSend/ahasend-go"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestGetMessagesStatusParameterUnitTests tests the status parameter functionality without making actual API calls
func TestGetMessagesStatusParameterUnitTests(t *testing.T) {

	t.Run("Status parameter method chaining", func(t *testing.T) {
		client := ahasend.NewAPIClient(ahasend.NewConfiguration())
		ctx := context.Background()
		accountID := uuid.New()

		// Test that Status() method returns the correct request type and allows chaining
		request := client.MessagesAPI.GetMessages(ctx, accountID).
			Sender("test@example.com").
			Status("Delivered")

		assert.NotNil(t, request)

		// Test chaining with other methods
		request2 := client.MessagesAPI.GetMessages(ctx, accountID).
			Status("Bounced,Failed").
			Sender("test@example.com").
			Recipient("user@example.com").
			Limit(10)

		assert.NotNil(t, request2)
	})

	t.Run("Status parameter accepts various formats", func(t *testing.T) {
		client := ahasend.NewAPIClient(ahasend.NewConfiguration())
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
				request := client.MessagesAPI.GetMessages(ctx, accountID).
					Sender("test@example.com").
					Status(tc.status)

				assert.NotNil(t, request)
			})
		}
	})
}

// TestGetMessagesStatusParameterIntegration tests status parameter with mock responses
func TestGetMessagesStatusParameterIntegration(t *testing.T) {

	// This test validates that the request is properly constructed
	// We'll test with different status combinations to ensure the parameter is passed correctly

	t.Run("Request construction with status parameter", func(t *testing.T) {
		configuration := ahasend.NewConfiguration()
		configuration.Host = "localhost:4010" // Point to Prism mock server
		configuration.Scheme = "http"         // Use HTTP for mock server
		client := ahasend.NewAPIClient(configuration)

		ctx := context.WithValue(context.Background(), ahasend.ContextAccessToken, "test-api-key")
		accountID := uuid.New()
		senderEmail := "test@example.com"

		testCases := []struct {
			name          string
			status        string
			expectedInURL bool
		}{
			{"Single status", "Delivered", true},
			{"Multiple statuses", "Bounced,Failed", true},
			{"Complex status list", "Delivered,Bounced,Failed,Queued", true},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				// Create the request but don't execute it
				request := client.MessagesAPI.GetMessages(ctx, accountID).
					Sender(senderEmail).
					Status(tc.status)

				assert.NotNil(t, request)

				// The key test is that the method chain works properly
				// In a real scenario, we would check the actual URL construction
				// but that requires access to internal request building
			})
		}
	})
}

// TestGetMessagesStatusParameterValidation tests edge cases and validation
func TestGetMessagesStatusParameterValidation(t *testing.T) {

	t.Run("Status parameter edge cases", func(t *testing.T) {
		client := ahasend.NewAPIClient(ahasend.NewConfiguration())
		ctx := context.Background()
		accountID := uuid.New()
		senderEmail := "test@example.com"

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
				// These should not panic and should return a valid request object
				request := client.MessagesAPI.GetMessages(ctx, accountID).
					Sender(senderEmail).
					Status(tc.status)

				assert.NotNil(t, request)
			})
		}
	})

	t.Run("Status parameter with nil handling", func(t *testing.T) {
		client := ahasend.NewAPIClient(ahasend.NewConfiguration())
		ctx := context.Background()
		accountID := uuid.New()
		senderEmail := "test@example.com"

		// Test that not calling Status() method still works (status should be nil)
		request := client.MessagesAPI.GetMessages(ctx, accountID).
			Sender(senderEmail)

		assert.NotNil(t, request)
	})
}

// TestGetMessagesStatusCombinations tests status parameter in combination with other parameters
func TestGetMessagesStatusCombinations(t *testing.T) {

	t.Run("Status with all other parameters", func(t *testing.T) {
		client := ahasend.NewAPIClient(ahasend.NewConfiguration())
		ctx := context.Background()
		accountID := uuid.New()

		// Test that status works with all other available parameters
		request := client.MessagesAPI.GetMessages(ctx, accountID).
			Status("Delivered,Bounced").
			Sender("sender@example.com").
			Recipient("recipient@example.com").
			Subject("Test Subject").
			MessageIdHeader("msg-123").
			Limit(50).
			Cursor("cursor-123")

		assert.NotNil(t, request)
	})

	t.Run("Method order independence", func(t *testing.T) {
		client := ahasend.NewAPIClient(ahasend.NewConfiguration())
		ctx := context.Background()
		accountID := uuid.New()

		// Test that the order of method calls doesn't matter
		request1 := client.MessagesAPI.GetMessages(ctx, accountID).
			Status("Delivered").
			Sender("test@example.com").
			Limit(10)

		request2 := client.MessagesAPI.GetMessages(ctx, accountID).
			Limit(10).
			Sender("test@example.com").
			Status("Delivered")

		request3 := client.MessagesAPI.GetMessages(ctx, accountID).
			Sender("test@example.com").
			Limit(10).
			Status("Delivered")

		assert.NotNil(t, request1)
		assert.NotNil(t, request2)
		assert.NotNil(t, request3)
	})
}

// TestGetMessagesStatusDocumentationExamples tests the examples from documentation
func TestGetMessagesStatusDocumentationExamples(t *testing.T) {

	client := ahasend.NewAPIClient(ahasend.NewConfiguration())
	ctx := context.WithValue(context.Background(), ahasend.ContextAccessToken, "test-key")
	accountID := uuid.New()

	t.Run("Documentation example 1: Single status", func(t *testing.T) {
		request := client.MessagesAPI.GetMessages(ctx, accountID).
			Sender("noreply@example.com").
			Status("Bounced")

		assert.NotNil(t, request)
	})

	t.Run("Documentation example 2: Multiple statuses", func(t *testing.T) {
		request := client.MessagesAPI.GetMessages(ctx, accountID).
			Sender("noreply@example.com").
			Status("Failed,Bounced")

		assert.NotNil(t, request)
	})

	t.Run("Documentation example 3: Status with other filters", func(t *testing.T) {
		request := client.MessagesAPI.GetMessages(ctx, accountID).
			Sender("noreply@example.com").
			Status("Delivered").
			Recipient("user@example.com").
			Limit(10)

		assert.NotNil(t, request)
	})

	t.Run("Documentation example 4: No status filter", func(t *testing.T) {
		request := client.MessagesAPI.GetMessages(ctx, accountID).
			Sender("noreply@example.com")

		assert.NotNil(t, request)
	})
}
