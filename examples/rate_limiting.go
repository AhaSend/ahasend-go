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

	fmt.Println("=== Rate Limiting Examples ===")

	// Example 1: Default rate limiting
	fmt.Println("1. Default Rate Limiting")
	fmt.Println("─────────────────────────")
	demonstrateDefaultRateLimiting(client)

	// Example 2: Custom rate limiting configuration
	fmt.Println("\n2. Custom Rate Limiting Configuration")
	fmt.Println("─────────────────────────────────────")
	demonstrateCustomRateLimiting(client)

	// Example 3: Rate limiting with bursts
	fmt.Println("\n3. Burst Traffic Handling")
	fmt.Println("──────────────────────────")
	demonstrateBurstHandling(ctx, client, accountID)

	// Example 4: Per-endpoint rate limiting
	fmt.Println("\n4. Per-Endpoint Rate Limiting")
	fmt.Println("──────────────────────────────")
	demonstratePerEndpointRateLimiting(ctx, client, accountID)

	// Example 5: Monitor rate limit status
	fmt.Println("\n5. Monitoring Rate Limit Status")
	fmt.Println("─────────────────────────────────")
	monitorRateLimitStatus(client)

	// Example 6: Graceful degradation
	fmt.Println("\n6. Graceful Degradation Strategy")
	fmt.Println("──────────────────────────────────")
	demonstrateGracefulDegradation(ctx, client, accountID)
}

func demonstrateDefaultRateLimiting(client *api.APIClient) {
	fmt.Println("The SDK includes intelligent rate limiting by default:")
	fmt.Println("  • General API: 100 requests/second")
	fmt.Println("  • Send Message: 100 requests/second")
	fmt.Println("  • Statistics: 1 request/second")
	fmt.Println("")
	fmt.Println("These limits protect you from 429 (Too Many Requests) errors.")
	fmt.Println("The SDK automatically queues and throttles requests.")
}

func demonstrateCustomRateLimiting(client *api.APIClient) {
	// Configure custom rate limits based on your plan
	fmt.Println("Configuring custom rate limits...")

	// Example: You have a plan that allows 500 emails/second
	client.SetSendMessageRateLimit(500, 1000) // 500 req/s, burst of 1000
	fmt.Println("  ✓ Send Message: 500 req/s (burst: 1000)")

	// Statistics endpoints are usually more limited
	client.SetStatisticsRateLimit(5, 10) // 5 req/s, burst of 10
	fmt.Println("  ✓ Statistics: 5 req/s (burst: 10)")

	// General API endpoints
	client.SetGeneralRateLimit(200, 400) // 200 req/s, burst of 400
	fmt.Println("  ✓ General API: 200 req/s (burst: 400)")

	// You can also configure all at once
	client.ConfigureCustomerRateLimits(api.CustomerRateLimitConfig{
		SendMessage: &api.RateLimitConfig{
			RequestsPerSecond: 500,
			BurstCapacity:     1000,
			Enabled:           true,
		},
		Statistics: &api.RateLimitConfig{
			RequestsPerSecond: 5,
			BurstCapacity:     10,
			Enabled:           true,
		},
		General: &api.RateLimitConfig{
			RequestsPerSecond: 200,
			BurstCapacity:     400,
			Enabled:           true,
		},
	})
	fmt.Println("\n  ✅ All rate limits configured via batch configuration")
}

func demonstrateBurstHandling(ctx context.Context, client *api.APIClient, accountID uuid.UUID) {
	// Configure for burst demonstration
	client.SetSendMessageRateLimit(5, 10) // Low limit to demonstrate throttling

	fmt.Println("Sending 20 requests with rate limit of 5/second...")
	fmt.Println("Watch how the SDK handles the burst:")

	var wg sync.WaitGroup
	startTime := time.Now()

	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()

			requestStart := time.Now()

			message := requests.CreateMessageRequest{
				From: common.SenderAddress{
					Email: "test@example.com",
				},
				Recipients: []common.Recipient{
					{Email: fmt.Sprintf("user%d@example.com", num)},
				},
				Subject:     fmt.Sprintf("Test Message %d", num),
				TextContent: ahasend.String("Rate limiting test"),
			}

			_, _, err := client.MessagesAPI.CreateMessage(ctx, accountID, message)

			elapsed := time.Since(requestStart)
			totalElapsed := time.Since(startTime)

			if err != nil {
				fmt.Printf("  Request %2d: ❌ Error after %v (total: %v)\n",
					num, elapsed, totalElapsed)
			} else {
				fmt.Printf("  Request %2d: ✓ Completed after %v (total: %v)\n",
					num, elapsed, totalElapsed)
			}
		}(i)

		// Small delay to control goroutine creation
		time.Sleep(50 * time.Millisecond)
	}

	wg.Wait()
	totalTime := time.Since(startTime)

	fmt.Printf("\nTotal time for 20 requests: %v\n", totalTime)
	fmt.Printf("Average time per request: %v\n", totalTime/20)
	fmt.Println("\nNotice how the SDK automatically throttled requests to respect the rate limit!")
}

func demonstratePerEndpointRateLimiting(ctx context.Context, client *api.APIClient, accountID uuid.UUID) {
	// Different endpoints have different rate limits
	fmt.Println("Different endpoints have different rate limits:")
	fmt.Println("")

	// Fast endpoint - Messages
	fmt.Print("Getting 3 message lists (high rate limit)... ")
	start := time.Now()
	for i := 0; i < 3; i++ {
		params := requests.GetMessagesParams{
			Limit: ahasend.Int32(1),
		}
		client.MessagesAPI.GetMessages(ctx, accountID, params)
	}
	fmt.Printf("Done in %v\n", time.Since(start))

	// Slow endpoint - Statistics (rate limited to 1/second)
	client.SetStatisticsRateLimit(1, 1)
	fmt.Print("Getting statistics 3 times (low rate limit)... ")
	start = time.Now()
	for i := 0; i < 3; i++ {
		params := requests.GetDeliverabilityStatisticsParams{}
		client.StatisticsAPI.GetDeliverabilityStatistics(ctx, accountID, params)
	}
	fmt.Printf("Done in %v\n", time.Since(start))

	fmt.Println("\nNotice the statistics requests took longer due to stricter rate limiting!")
}

func monitorRateLimitStatus(client *api.APIClient) {
	// Monitor current rate limit status
	endpoints := []api.EndpointType{
		api.GeneralAPI,
		api.SendMessageAPI,
		api.StatisticsAPI,
	}

	fmt.Println("Current Rate Limit Status:")
	fmt.Println("┌────────────────┬──────────────┬──────────┬─────────┐")
	fmt.Println("│ Endpoint       │ Tokens Avail │ Capacity │ Enabled │")
	fmt.Println("├────────────────┼──────────────┼──────────┼─────────┤")

	for _, endpoint := range endpoints {
		status := client.GetRateLimitStatus(endpoint)
		enabled := "Yes"
		if !status.Enabled {
			enabled = "No"
		}
		fmt.Printf("│ %-14s │ %12d │ %8d │ %-7s │\n",
			endpoint,
			status.TokensAvailable,
			status.BurstCapacity,
			enabled)
	}
	fmt.Println("└────────────────┴──────────────┴──────────┴─────────┘")

	fmt.Println("\nTokens are consumed when making requests and replenish over time.")
	fmt.Println("When tokens = 0, requests will wait until tokens are available.")
}

func demonstrateGracefulDegradation(ctx context.Context, client *api.APIClient, accountID uuid.UUID) {
	fmt.Println("Implementing graceful degradation for high load:")
	fmt.Println("")

	// Check rate limit before sending
	status := client.GetRateLimitStatus(api.SendMessageAPI)

	if status.TokensAvailable < 10 {
		fmt.Println("⚠️  Low on rate limit tokens!")
		fmt.Println("   Implementing degradation strategy:")
		fmt.Println("   • Prioritizing transactional emails")
		fmt.Println("   • Queuing marketing emails for later")
		fmt.Println("   • Reducing statistics polling frequency")

		// Example: Disable non-critical features
		client.EnableRateLimit(api.StatisticsAPI, false)
		fmt.Println("   • Disabled rate limiting for statistics (use with caution!)")

		// Re-enable after recovery
		time.Sleep(2 * time.Second)
		client.EnableRateLimit(api.StatisticsAPI, true)
		fmt.Println("   • Re-enabled rate limiting after recovery")
	} else {
		fmt.Printf("✅ Sufficient tokens available: %d/%d\n",
			status.TokensAvailable, status.BurstCapacity)
	}

	// Best practices
	fmt.Println("\nBest Practices for Rate Limiting:")
	fmt.Println("1. Set appropriate limits based on your plan")
	fmt.Println("2. Use burst capacity for traffic spikes")
	fmt.Println("3. Monitor token availability")
	fmt.Println("4. Implement queue systems for large batches")
	fmt.Println("5. Use webhooks instead of polling for status")
	fmt.Println("6. Cache frequently accessed data")
	fmt.Println("7. Implement exponential backoff for retries")
}
