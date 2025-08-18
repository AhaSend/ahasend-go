//go:build ignore

package main

import (
	"context"
	"fmt"
	"log"
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

	// Create authentication context
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, apiKey)

	// Create the email message
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
		Subject: "Welcome to AhaSend!",
		HtmlContent: ahasend.String(`
			<html>
				<body>
					<h1>Welcome!</h1>
					<p>Thank you for using AhaSend. This is a test email sent using the Go SDK.</p>
					<p>Best regards,<br>The AhaSend Team</p>
				</body>
			</html>
		`),
		TextContent: ahasend.String(`
Welcome!

Thank you for using AhaSend. This is a test email sent using the Go SDK.

Best regards,
The AhaSend Team
		`),
		Tags: []string{"welcome", "test"},
		Tracking: &common.Tracking{
			Open:  ahasend.Bool(true),
			Click: ahasend.Bool(true),
		},
	}

	// Send the email
	fmt.Println("Sending email...")
	response, httpResp, err := client.MessagesAPI.CreateMessage(ctx, accountID, message)

	if err != nil {
		// Check if it's an API error with details
		if apiErr, ok := err.(*api.APIError); ok {
			log.Fatalf("API Error: %s\nStatus Code: %d\nResponse Body: %s",
				apiErr.Error(), apiErr.StatusCode, string(apiErr.Raw))
		}
		log.Fatalf("Error sending email: %v", err)
	}

	// Check HTTP status
	if httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 {
		fmt.Printf("Email sent successfully!\n")
		if len(response.Data) > 0 {
			if response.Data[0].ID != nil {
				fmt.Printf("Message ID: %s\n", *response.Data[0].ID)
			}
			fmt.Printf("Status: %s\n", response.Data[0].Status)
			fmt.Printf("Recipient: %s\n", response.Data[0].Recipient.Email)
		}
	} else {
		fmt.Printf("Unexpected status code: %d\n", httpResp.StatusCode)
	}

	// Note about idempotency
	fmt.Printf("\nIdempotency Key: %s\n", httpResp.Request.Header.Get("Idempotency-Key"))
	fmt.Println("(The SDK automatically adds idempotency keys to prevent duplicate sends)")
}
