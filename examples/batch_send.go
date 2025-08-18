//go:build ignore

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/api"
	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/google/uuid"
)

// Recipient represents a person to send email to
type Recipient struct {
	Email     string
	Name      string
	Variables map[string]string // For template personalization
}

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

	// Create a new API client with custom rate limiting for batch sending
	client := api.NewAPIClient(
		api.WithAPIKey(apiKey),
	)

	// Configure rate limiting for batch sends (adjust based on your plan)
	client.SetSendMessageRateLimit(50, 100) // 50 requests/second, burst of 100

	// Create authentication context
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, apiKey)

	// Sample list of recipients
	recipients := []Recipient{
		{Email: "user1@example.com", Name: "User One", Variables: map[string]string{"code": "ABC123"}},
		{Email: "user2@example.com", Name: "User Two", Variables: map[string]string{"code": "DEF456"}},
		{Email: "user3@example.com", Name: "User Three", Variables: map[string]string{"code": "GHI789"}},
		// Add more recipients as needed...
	}

	// Send emails concurrently with controlled parallelism
	fmt.Printf("Sending %d emails in batch...\n", len(recipients))

	// Use a WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup

	// Use a semaphore channel to limit concurrent sends
	semaphore := make(chan struct{}, 10) // Max 10 concurrent sends

	// Track results
	successCount := 0
	failCount := 0
	var mu sync.Mutex

	startTime := time.Now()

	for _, recipient := range recipients {
		wg.Add(1)

		// Acquire semaphore
		semaphore <- struct{}{}

		go func(r Recipient) {
			defer wg.Done()
			defer func() { <-semaphore }() // Release semaphore

			// Create personalized message for this recipient
			message := requests.CreateMessageRequest{
				From: common.SenderAddress{
					Email: "noreply@yourdomain.com",
					Name:  ahasend.String("Your Company"),
				},
				Recipients: []common.Recipient{
					{
						Email: r.Email,
						Name:  ahasend.String(r.Name),
					},
				},
				Subject: fmt.Sprintf("Your verification code, %s", r.Name),
				HtmlContent: ahasend.String(fmt.Sprintf(`
					<html>
						<body>
							<h2>Hello %s!</h2>
							<p>Your verification code is: <strong>%s</strong></p>
							<p>This code will expire in 24 hours.</p>
							<p>If you didn't request this code, please ignore this email.</p>
							<p>Best regards,<br>Your Company</p>
						</body>
					</html>
				`, r.Name, r.Variables["code"])),
				TextContent: ahasend.String(fmt.Sprintf(`
Hello %s!

Your verification code is: %s

This code will expire in 24 hours.

If you didn't request this code, please ignore this email.

Best regards,
Your Company
				`, r.Name, r.Variables["code"])),
				Tags: []string{"verification", "batch"},
			}

			// Send the email with automatic retry on failure
			_, httpResp, err := client.MessagesAPI.CreateMessage(ctx, accountID, message)

			mu.Lock()
			if err != nil {
				failCount++
				log.Printf("Failed to send to %s: %v", r.Email, err)
			} else if httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 {
				successCount++
				fmt.Printf("✓ Sent to %s\n", r.Email)
			} else {
				failCount++
				log.Printf("Failed to send to %s: HTTP %d", r.Email, httpResp.StatusCode)
			}
			mu.Unlock()
		}(recipient)
	}

	// Wait for all sends to complete
	wg.Wait()

	duration := time.Since(startTime)

	// Print summary
	fmt.Println("\n=== Batch Send Complete ===")
	fmt.Printf("Total: %d emails\n", len(recipients))
	fmt.Printf("Successful: %d\n", successCount)
	fmt.Printf("Failed: %d\n", failCount)
	fmt.Printf("Duration: %v\n", duration)
	fmt.Printf("Rate: %.2f emails/second\n", float64(len(recipients))/duration.Seconds())

	// Best practices notes
	fmt.Println("\n=== Best Practices for Batch Sending ===")
	fmt.Println("1. Use rate limiting to avoid overwhelming the API")
	fmt.Println("2. Implement proper error handling and retries")
	fmt.Println("3. Use goroutines for concurrent sending with semaphore control")
	fmt.Println("4. Track success/failure metrics")
	fmt.Println("5. Consider using webhooks to track delivery status")
	fmt.Println("6. For very large batches (>10,000), consider using bulk email services")
}
