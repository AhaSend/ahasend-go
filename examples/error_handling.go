//go:build ignore

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/api"
	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/google/uuid"
)

func main() {
	// Get API credentials from environment variables
	apiKey := os.Getenv("AHASEND_API_KEY")
	if apiKey == "" {
		log.Fatal("AHASEND_API_KEY environment variable is required")
	}

	accountIDStr := os.Getenv("AHASEND_ACCOUNT_ID")
	if accountIDStr == "" {
		log.Fatal("AHASEND_ACCOUNT_ID environment variable is required")
	}

	accountID, err := uuid.Parse(accountIDStr)
	if err != nil {
		log.Fatalf("Invalid account ID: %v", err)
	}

	// Create a new API client
	client := api.NewAPIClient(
		api.WithAPIKey(apiKey),
	)

	// Note: The SDK includes built-in rate limiting and automatic retries
	// for 429 (Too Many Requests) responses with exponential backoff

	// Create authentication context
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, apiKey)

	// Example 1: Handle validation errors
	fmt.Println("=== Example 1: Validation Error ===")
	handleValidationError(ctx, client, accountID)

	// Example 2: Handle authentication errors
	fmt.Println("\n=== Example 2: Authentication Error ===")
	handleAuthenticationError(client, accountID)

	// Example 3: Handle rate limiting
	fmt.Println("\n=== Example 3: Rate Limiting ===")
	handleRateLimiting(ctx, client, accountID)

	// Example 4: Handle network errors
	fmt.Println("\n=== Example 4: Network Error ===")
	handleNetworkError(ctx, client, accountID)

	// Example 5: Comprehensive error handling
	fmt.Println("\n=== Example 5: Comprehensive Error Handling ===")
	comprehensiveErrorHandling(ctx, client, accountID)
}

func handleValidationError(ctx context.Context, client *api.APIClient, accountID uuid.UUID) {
	// Create an invalid message (missing required fields)
	message := requests.CreateMessageRequest{
		// Missing From field - will cause validation error
		Recipients: []common.Recipient{
			{
				Email: "invalid-email", // Invalid email format
			},
		},
		Subject: "", // Empty subject
	}

	_, _, err := client.MessagesAPI.CreateMessage(ctx, accountID, message)

	if err != nil {
		handleError(err, "Validation Error")
	}
}

func handleAuthenticationError(client *api.APIClient, accountID uuid.UUID) {
	// Create context with invalid API key
	invalidCtx := context.WithValue(context.Background(), api.ContextAccessToken, "invalid-api-key")

	message := requests.CreateMessageRequest{
		From: common.SenderAddress{
			Email: "sender@example.com",
		},
		Recipients: []common.Recipient{
			{
				Email: "recipient@example.com",
			},
		},
		Subject:     "Test",
		TextContent: ahasend.String("Test message"),
	}

	_, _, err := client.MessagesAPI.CreateMessage(invalidCtx, accountID, message)

	if err != nil {
		handleError(err, "Authentication Error")
	}
}

func handleRateLimiting(ctx context.Context, client *api.APIClient, accountID uuid.UUID) {
	// Note: The SDK includes automatic rate limiting and retry logic
	fmt.Println("The SDK automatically handles rate limiting with:")
	fmt.Println("- Token bucket algorithm for request throttling")
	fmt.Println("- Automatic retry with exponential backoff for 429 errors")
	fmt.Println("- Configurable rate limits per endpoint type")

	// Demonstrate rate limit configuration
	client.SetSendMessageRateLimit(10, 20) // 10 req/s, burst of 20

	fmt.Println("\nRate limits configured:")
	status := client.GetRateLimitStatus(api.SendMessageAPI)
	fmt.Printf("Send Message API: %d tokens available (capacity: %d)\n",
		status.TokensAvailable, status.BurstCapacity)
}

func handleNetworkError(ctx context.Context, client *api.APIClient, accountID uuid.UUID) {
	// Simulate network error by using invalid host
	badClient := api.NewAPIClient(
		api.WithAPIKey("test-key"),
	)
	// Override configuration with invalid host to simulate network error
	badClient.GetConfig().Host = "invalid.host.that.does.not.exist:9999"

	message := requests.CreateMessageRequest{
		From: common.SenderAddress{
			Email: "sender@example.com",
		},
		Recipients: []common.Recipient{
			{
				Email: "recipient@example.com",
			},
		},
		Subject:     "Test",
		TextContent: ahasend.String("Test message"),
	}

	_, _, err := badClient.MessagesAPI.CreateMessage(ctx, accountID, message)

	if err != nil {
		handleError(err, "Network Error")
	}
}

func comprehensiveErrorHandling(ctx context.Context, client *api.APIClient, accountID uuid.UUID) {
	message := requests.CreateMessageRequest{
		From: common.SenderAddress{
			Email: "sender@yourdomain.com",
			Name:  ahasend.String("Your Name"),
		},
		Recipients: []common.Recipient{
			{
				Email: "recipient@example.com",
				Name:  ahasend.String("Recipient Name"),
			},
		},
		Subject:     "Test Email",
		HtmlContent: ahasend.String("<p>This is a test email.</p>"),
	}

	response, httpResp, err := client.MessagesAPI.CreateMessage(ctx, accountID, message)

	if err != nil {
		// Comprehensive error handling
		var apiErr *api.APIError
		if errors.As(err, &apiErr) {
			fmt.Printf("API Error Details:\n")
			fmt.Printf("  Status Code: %d\n", apiErr.StatusCode)
			fmt.Printf("  Error Message: %s\n", apiErr.Error())
			fmt.Printf("  Response Body: %s\n", string(apiErr.Raw))

			// Handle specific status codes
			switch apiErr.StatusCode {
			case http.StatusBadRequest:
				fmt.Println("  → Fix: Check request parameters and validation")
			case http.StatusUnauthorized:
				fmt.Println("  → Fix: Check API key and authentication")
			case http.StatusForbidden:
				fmt.Println("  → Fix: Check account permissions and API key scopes")
			case http.StatusNotFound:
				fmt.Println("  → Fix: Check account ID and endpoint URL")
			case http.StatusTooManyRequests:
				fmt.Println("  → Fix: Implement rate limiting or reduce request frequency")
			case http.StatusInternalServerError:
				fmt.Println("  → Fix: Retry request or contact support")
			default:
				fmt.Printf("  → Unexpected error code: %d\n", apiErr.StatusCode)
			}
		} else {
			// Handle non-API errors (network, parsing, etc.)
			fmt.Printf("Non-API Error: %v\n", err)
			fmt.Println("  → This might be a network issue or client-side error")
		}
		return
	}

	// Success handling
	if httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 {
		fmt.Println("Email sent successfully!")
		if response != nil && len(response.Data) > 0 && response.Data[0].ID != nil {
			fmt.Printf("Message ID: %s\n", *response.Data[0].ID)
		}
	}
}

func handleError(err error, context string) {
	fmt.Printf("%s occurred:\n", context)

	// Check if it's an API error
	var apiErr *api.APIError
	if errors.As(err, &apiErr) {
		fmt.Printf("  Type: API Error\n")
		fmt.Printf("  Status Code: %d\n", apiErr.StatusCode)
		fmt.Printf("  Message: %s\n", apiErr.Error())
		if apiErr.StatusCode > 0 && len(apiErr.Raw) > 0 {
			fmt.Printf("  Response Body: %s\n", string(apiErr.Raw))
		}
	} else {
		fmt.Printf("  Type: Client Error\n")
		fmt.Printf("  Message: %s\n", err.Error())
	}

	// Determine if error is retryable
	if isRetryableError(err) {
		fmt.Println("  → This error is retryable")
	} else {
		fmt.Println("  → This error is not retryable")
	}
}

func isRetryableError(err error) bool {
	var apiErr *api.APIError
	if errors.As(err, &apiErr) {
		retryableStatuses := []int{429, 502, 503, 504}
		for _, status := range retryableStatuses {
			if apiErr.StatusCode == status {
				return true
			}
		}
	}
	return false
}
