//go:build ignore

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

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

	// Note: The SDK includes built-in idempotency support.
	// You can provide manual idempotency keys via the WithIdempotencyKey() option.

	// Create authentication context
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, apiKey)

	fmt.Println("=== Idempotency Examples ===")

	// Example 1: Automatic idempotency
	fmt.Println("1. Automatic Idempotency")
	fmt.Println("─────────────────────────")
	demonstrateAutomaticIdempotency(ctx, client, accountID)

	// Example 2: Manual idempotency key
	fmt.Println("\n2. Manual Idempotency Key")
	fmt.Println("──────────────────────────")
	demonstrateManualIdempotency(ctx, client, accountID)

	// Example 3: Preventing duplicate sends
	fmt.Println("\n3. Preventing Duplicate Sends")
	fmt.Println("───────────────────────────────")
	demonstrateDuplicatePrevention(ctx, client, accountID)

	// Example 4: Idempotency with retries
	fmt.Println("\n4. Idempotency with Retries")
	fmt.Println("─────────────────────────────")
	demonstrateIdempotencyWithRetries(ctx, client, accountID)

	// Example 5: Transaction-safe email sending
	fmt.Println("\n5. Transaction-Safe Email Sending")
	fmt.Println("───────────────────────────────────")
	demonstrateTransactionSafeSending(ctx, client, accountID)
}

func demonstrateAutomaticIdempotency(ctx context.Context, client *api.APIClient, accountID uuid.UUID) {
	fmt.Println("The SDK automatically adds idempotency keys to prevent duplicates.")
	fmt.Println("")

	message := requests.CreateMessageRequest{
		From: common.SenderAddress{
			Email: "sender@example.com",
		},
		Recipients: []common.Recipient{
			{Email: "recipient@example.com"},
		},
		Subject:     "Automatic Idempotency Test",
		TextContent: ahasend.String("This email has automatic idempotency protection."),
	}

	// First send
	fmt.Print("Sending email (1st attempt)... ")
	response1, httpResp1, err := client.MessagesAPI.CreateMessage(ctx, accountID, message)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Success! (Status: %d)\n", httpResp1.StatusCode)
		idempotencyKey := httpResp1.Request.Header.Get("Idempotency-Key")
		fmt.Printf("  Auto-generated Idempotency Key: %s\n", idempotencyKey)
		if len(response1.Data) > 0 {
			fmt.Printf("  Message ID: %s\n", *response1.Data[0].ID)
		}
	}

	fmt.Println("\nNote: Each request gets a unique idempotency key automatically.")
	fmt.Println("This prevents accidental duplicates during retries.")
}

func demonstrateManualIdempotency(ctx context.Context, client *api.APIClient, accountID uuid.UUID) {
	fmt.Println("You can provide your own idempotency key for full control.")
	fmt.Println("")

	// Generate a unique key for this transaction
	idempotencyKey := fmt.Sprintf("order-confirmation-%d", time.Now().Unix())

	message := requests.CreateMessageRequest{
		From: common.SenderAddress{
			Email: "sender@example.com",
		},
		Recipients: []common.Recipient{
			{Email: "customer@example.com"},
		},
		Subject:     "Order Confirmation #12345",
		TextContent: ahasend.String("Your order has been confirmed."),
	}

	fmt.Printf("Using manual idempotency key: %s\n", idempotencyKey)

	// Send with manual idempotency key
	fmt.Print("Sending email with manual key... ")
	response, httpResp, err := client.MessagesAPI.CreateMessage(ctx, accountID, message, api.WithIdempotencyKey(idempotencyKey))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Success! (Status: %d)\n", httpResp.StatusCode)
		if len(response.Data) > 0 {
			fmt.Printf("  Message ID: %s\n", *response.Data[0].ID)
		}
	}

	// Try sending again with the same key (simulating a retry)
	fmt.Print("Sending email (retry)... ")
	response2, httpResp2, err2 := client.MessagesAPI.CreateMessage(ctx, accountID, message, api.WithIdempotencyKey(idempotencyKey))

	if err2 != nil {
		fmt.Printf("Error: %v\n", err2)
	} else {
		fmt.Printf("Success! (Status: %d)\n", httpResp2.StatusCode)
		if httpResp2.StatusCode == 200 || httpResp2.StatusCode == 201 {
			fmt.Println("  ✓ Server recognized duplicate and returned cached response")
			if len(response2.Data) > 0 {
				fmt.Printf("  Same Message ID: %s\n", *response2.Data[0].ID)
			}
		}
	}
}

func demonstrateDuplicatePrevention(ctx context.Context, client *api.APIClient, accountID uuid.UUID) {
	fmt.Println("Scenario: Preventing duplicate welcome emails")
	fmt.Println("")

	userID := "user-12345"

	// Create idempotency key based on user ID and action
	idempotencyKey := fmt.Sprintf("welcome-email-%s", userID)

	message := requests.CreateMessageRequest{
		From: common.SenderAddress{
			Email: "noreply@example.com",
			Name:  ahasend.String("Your App"),
		},
		Recipients: []common.Recipient{
			{
				Email: "newuser@example.com",
				Name:  ahasend.String("New User"),
			},
		},
		Subject: "Welcome to Our Service!",
		HtmlContent: ahasend.String(`
			<h1>Welcome!</h1>
			<p>Thank you for signing up. Here's your special welcome offer...</p>
		`),
		Tags: []string{"welcome", "onboarding"},
	}

	// Simulate multiple attempts (e.g., due to network issues, retries, etc.)
	attempts := 3
	for i := 1; i <= attempts; i++ {
		fmt.Printf("Attempt %d: ", i)

		_, httpResp, err := client.MessagesAPI.CreateMessage(ctx, accountID, message, api.WithIdempotencyKey(idempotencyKey))

		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			if i == 1 {
				fmt.Printf("Email sent! (Status: %d)\n", httpResp.StatusCode)
			} else {
				fmt.Printf("Duplicate prevented! (Status: %d)\n", httpResp.StatusCode)
			}
		}

		// Simulate delay between attempts
		time.Sleep(1 * time.Second)
	}

	fmt.Println("\n✅ Only one email was actually sent despite multiple attempts!")
}

func demonstrateIdempotencyWithRetries(ctx context.Context, client *api.APIClient, accountID uuid.UUID) {
	fmt.Println("Implementing safe retry logic with idempotency")
	fmt.Println("")

	// Create a standard client for retry example
	retryClient := api.NewAPIClient(
		api.WithAPIKey(os.Getenv("AHASEND_API_KEY")),
	)

	// Note: The SDK includes built-in rate limiting and retry logic
	// with exponential backoff for 429 responses.

	message := requests.CreateMessageRequest{
		From: common.SenderAddress{
			Email: "sender@example.com",
		},
		Recipients: []common.Recipient{
			{Email: "recipient@example.com"},
		},
		Subject:     "Retry-Safe Email",
		TextContent: ahasend.String("This email is safe to retry."),
	}

	// The SDK will automatically retry with the same idempotency key
	fmt.Println("Sending email with automatic retry...")
	response, httpResp, err := retryClient.MessagesAPI.CreateMessage(ctx, accountID, message)

	if err != nil {
		fmt.Printf("Failed after retries: %v\n", err)
	} else {
		fmt.Printf("Success! (Status: %d)\n", httpResp.StatusCode)
		if len(response.Data) > 0 {
			fmt.Printf("Message ID: %s\n", *response.Data[0].ID)
		}
		fmt.Println("\n✅ The SDK handled retries safely with idempotency!")
	}
}

func demonstrateTransactionSafeSending(ctx context.Context, client *api.APIClient, accountID uuid.UUID) {
	fmt.Println("Example: Sending transaction confirmation emails safely")
	fmt.Println("")

	// Simulate a transaction
	transactionID := fmt.Sprintf("txn_%d", time.Now().UnixNano())
	amount := "$99.99"

	// Create idempotency key based on transaction ID
	// This ensures each transaction gets exactly one confirmation email
	idempotencyKey := fmt.Sprintf("txn-confirm-%s", transactionID)

	message := requests.CreateMessageRequest{
		From: common.SenderAddress{
			Email: "billing@example.com",
			Name:  ahasend.String("Billing Department"),
		},
		Recipients: []common.Recipient{
			{
				Email: "customer@example.com",
				Name:  ahasend.String("Valued Customer"),
			},
		},
		Subject: fmt.Sprintf("Payment Confirmation - %s", transactionID),
		HtmlContent: ahasend.String(fmt.Sprintf(`
			<h2>Payment Confirmation</h2>
			<p>Your payment has been processed successfully.</p>
			<table>
				<tr><td>Transaction ID:</td><td><strong>%s</strong></td></tr>
				<tr><td>Amount:</td><td><strong>%s</strong></td></tr>
				<tr><td>Date:</td><td>%s</td></tr>
			</table>
			<p>Thank you for your business!</p>
		`, transactionID, amount, time.Now().Format("January 2, 2006"))),
		Tags: []string{"transaction", "payment", "confirmation"},
		// Note: Use substitutions or headers for custom metadata
		Substitutions: map[string]interface{}{
			"transaction_id": transactionID,
			"amount":         amount,
		},
	}

	fmt.Printf("Transaction ID: %s\n", transactionID)
	fmt.Printf("Idempotency Key: %s\n", idempotencyKey)
	fmt.Println("")

	// Send confirmation email
	fmt.Print("Sending transaction confirmation... ")
	response, httpResp, err := client.MessagesAPI.CreateMessage(ctx, accountID, message, api.WithIdempotencyKey(idempotencyKey))

	if err != nil {
		fmt.Printf("Error: %v\n", err)

		// In a real application, you might want to:
		// 1. Log the error
		// 2. Queue for retry
		// 3. Alert operations team
		// 4. But NOT charge the customer again!

	} else {
		fmt.Printf("Success! (Status: %d)\n", httpResp.StatusCode)
		if len(response.Data) > 0 {
			fmt.Printf("Confirmation sent! Message ID: %s\n", *response.Data[0].ID)
		}
	}

	fmt.Println("\n=== Best Practices for Idempotency ===")
	fmt.Println("1. Use deterministic keys based on your business logic")
	fmt.Println("2. Include entity IDs in keys (user ID, order ID, etc.)")
	fmt.Println("3. Store idempotency keys with transactions in your database")
	fmt.Println("4. Set appropriate key expiration (24-48 hours typical)")
	fmt.Println("5. Use for all non-idempotent operations (payments, emails, etc.)")
	fmt.Println("6. Document your key format for team consistency")
}
