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

	// Configure rate limiting for statistics endpoints (they have lower limits)
	client.SetStatisticsRateLimit(1, 5) // 1 req/s, burst of 5

	// Create authentication context
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, apiKey)

	fmt.Println("=== Email Statistics Dashboard ===")

	// Get time range for statistics (last 30 days)
	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -30)

	fmt.Printf("Period: %s to %s\n\n",
		startDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"))

	// 1. Get deliverability statistics
	fmt.Println("1. Deliverability Statistics")
	fmt.Println("─────────────────────────────")
	getDeliverabilityStats(ctx, client, accountID, startDate, endDate)

	// 2. Get bounce statistics
	fmt.Println("\n2. Bounce Statistics")
	fmt.Println("─────────────────────────────")
	getBounceStats(ctx, client, accountID, startDate, endDate)

	// 3. Get delivery time statistics
	fmt.Println("\n3. Delivery Time Statistics")
	fmt.Println("─────────────────────────────")
	getDeliveryTimeStats(ctx, client, accountID, startDate, endDate)

	// 4. Get recent messages
	fmt.Println("\n4. Recent Messages")
	fmt.Println("─────────────────────────────")
	getRecentMessages(ctx, client, accountID)

	// 5. Calculate email performance metrics
	fmt.Println("\n5. Performance Summary")
	fmt.Println("─────────────────────────────")
	showPerformanceSummary()
}

func getDeliverabilityStats(ctx context.Context, client *api.APIClient, accountID uuid.UUID, startDate, endDate time.Time) {
	params := requests.GetDeliverabilityStatisticsParams{
		FromTime: &startDate,
		ToTime:   &endDate,
		GroupBy:  ahasend.String("day"), // Options: hour, day, week, month
	}
	response, httpResp, err := client.StatisticsAPI.GetDeliverabilityStatistics(ctx, accountID, params)

	if err != nil {
		handleError(err)
		return
	}

	if httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 && response.Data != nil && len(response.Data) > 0 {
		var totalSent, totalDelivered, totalBounced, totalFailed, totalSuppressed int

		for _, stat := range response.Data {
			totalSent += stat.ReceptionCount
			totalDelivered += stat.DeliveredCount
			totalBounced += stat.BouncedCount
			totalFailed += stat.FailedCount
			totalSuppressed += stat.SuppressedCount
		}

		fmt.Printf("  Total Sent:       %d\n", totalSent)
		fmt.Printf("  Total Delivered:  %d (%.1f%%)\n",
			totalDelivered,
			calculatePercentage(totalDelivered, totalSent))
		fmt.Printf("  Total Bounced:    %d (%.1f%%)\n",
			totalBounced,
			calculatePercentage(totalBounced, totalSent))
		fmt.Printf("  Total Failed:     %d (%.1f%%)\n",
			totalFailed,
			calculatePercentage(totalFailed, totalSent))
		fmt.Printf("  Total Suppressed: %d (%.1f%%)\n",
			totalSuppressed,
			calculatePercentage(totalSuppressed, totalSent))

		// Show daily breakdown for last 7 days
		fmt.Println("\n  Last 7 Days Breakdown:")
		dayCount := 7
		if len(response.Data) < dayCount {
			dayCount = len(response.Data)
		}

		for i := len(response.Data) - dayCount; i < len(response.Data); i++ {
			stat := response.Data[i]
			fmt.Printf("    %s: Sent=%d, Delivered=%d, Bounced=%d\n",
				stat.FromTimestamp.Format("Jan 02"),
				stat.ReceptionCount,
				stat.DeliveredCount,
				stat.BouncedCount)
		}
	} else {
		fmt.Println("  No statistics available for this period.")
	}
}

func getBounceStats(ctx context.Context, client *api.APIClient, accountID uuid.UUID, startDate, endDate time.Time) {
	params := requests.GetBounceStatisticsParams{
		FromTime: &startDate,
		ToTime:   &endDate,
	}
	response, httpResp, err := client.StatisticsAPI.GetBounceStatistics(ctx, accountID, params)

	if err != nil {
		handleError(err)
		return
	}

	if httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 && response.Data != nil && len(response.Data) > 0 {
		var totalBounces int
		bouncesByClassification := make(map[string]int)

		for _, stat := range response.Data {
			for _, bounce := range stat.Bounces {
				totalBounces += bounce.Count
				bouncesByClassification[bounce.Classification] += bounce.Count
			}
		}

		fmt.Printf("  Total Bounces:    %d\n", totalBounces)

		// Show breakdown by classification
		fmt.Println("\n  Bounce Breakdown by Classification:")
		for classification, count := range bouncesByClassification {
			fmt.Printf("    %s: %d (%.1f%%)\n",
				classification, count,
				calculatePercentage(count, totalBounces))
		}

		// Show detailed breakdown
		fmt.Println("\n  Detailed Breakdown:")
		for _, stat := range response.Data {
			if stat.Bounces != nil {
				for _, bounce := range stat.Bounces {
					if bounce.Count > 0 {
						fmt.Printf("    %s - %s: %d bounces\n",
							stat.FromTimestamp.Format("Jan 02"),
							bounce.Classification,
							bounce.Count)
					}
				}
			}
		}

		// Common bounce reasons
		fmt.Println("\n  Common Bounce Classifications:")
		fmt.Println("    • hard: Invalid email, domain doesn't exist")
		fmt.Println("    • soft: Mailbox full, server temporarily unavailable")
		fmt.Println("    • admin: Blocked by admin policies")

		if bouncesByClassification["hard"] > 0 {
			fmt.Println("\n  ⚠️  Action Required: Review and remove hard bounced emails from your list")
		}
	} else {
		fmt.Println("  No bounce statistics available for this period.")
	}
}

func getDeliveryTimeStats(ctx context.Context, client *api.APIClient, accountID uuid.UUID, startDate, endDate time.Time) {
	params := requests.GetDeliveryTimeStatisticsParams{
		FromTime: &startDate,
		ToTime:   &endDate,
	}
	response, httpResp, err := client.StatisticsAPI.GetDeliveryTimeStatistics(ctx, accountID, params)

	if err != nil {
		handleError(err)
		return
	}

	if httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 && response.Data != nil && len(response.Data) > 0 {
		var totalMessages int
		avgDeliveryTime := 0.0

		// Calculate weighted average delivery time
		for _, stat := range response.Data {
			if stat.DeliveredCount > 0 {
				totalMessages += stat.DeliveredCount
				avgDeliveryTime += float64(stat.DeliveredCount) * stat.AvgDeliveryTime
			}
		}

		if totalMessages > 0 {
			avgDeliveryTime = avgDeliveryTime / float64(totalMessages)
		}

		fmt.Printf("  Average Delivery Time: %.2f seconds\n", avgDeliveryTime)
		fmt.Printf("  Total Messages:        %d\n", totalMessages)

		// Show delivery time breakdown by time bucket
		fmt.Println("\n  Delivery Time Breakdown:")
		for _, stat := range response.Data {
			if stat.DeliveredCount > 0 {
				percentage := calculatePercentage(stat.DeliveredCount, totalMessages)
				fmt.Printf("    %s: %d messages, avg %.2fs (%.1f%%)\n",
					stat.FromTimestamp.Format("Jan 02"),
					stat.DeliveredCount,
					stat.AvgDeliveryTime,
					percentage)
			}

			// Show breakdown by domain if available
			if len(stat.DeliveryTimes) > 0 {
				for _, dt := range stat.DeliveryTimes {
					if dt.RecipientDomain != nil && dt.DeliveryTime != nil {
						fmt.Printf("      Domain %s: avg %.2fs\n",
							*dt.RecipientDomain,
							*dt.DeliveryTime)
					}
				}
			}
		}

		if avgDeliveryTime < 5 {
			fmt.Println("\n  ✅ Excellent delivery speed!")
		} else if avgDeliveryTime < 30 {
			fmt.Println("\n  ✓ Good delivery speed")
		} else {
			fmt.Println("\n  ⚠️  Delivery speed could be improved")
		}
	} else {
		fmt.Println("  No delivery time statistics available for this period.")
	}
}

func getRecentMessages(ctx context.Context, client *api.APIClient, accountID uuid.UUID) {
	params := requests.GetMessagesParams{
		PaginationParams: common.PaginationParams{
			Limit: ahasend.Int32(5),
		},
	}
	response, httpResp, err := client.MessagesAPI.GetMessages(ctx, accountID, params)

	if err != nil {
		handleError(err)
		return
	}

	if httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 && response.Data != nil && len(response.Data) > 0 {
		for i, msg := range response.Data {
			fmt.Printf("  %d. Message API ID: %s\n", i+1, msg.ID.String())
			fmt.Printf("     Message-ID header: %s\n", msg.MessageID)
			fmt.Printf("     Direction: %s\n", msg.Direction)
			fmt.Printf("     Subject: %s\n", msg.Subject)
			fmt.Printf("     From: %s → To: %s\n", msg.Sender, msg.Recipient)
			fmt.Printf("     Status: %s\n", msg.Status)
			fmt.Printf("     Created: %s\n", msg.CreatedAt.Format("2006-01-02 15:04:05"))

			if msg.SentAt != nil {
				fmt.Printf("     Sent: %s\n", msg.SentAt.Format("2006-01-02 15:04:05"))
			}

			if msg.DeliveredAt != nil {
				fmt.Printf("     Delivered: %s\n", msg.DeliveredAt.Format("2006-01-02 15:04:05"))
			}

			fmt.Printf("     Delivery Attempts: %d\n", msg.NumAttempts)
			fmt.Printf("     Opens: %d, Clicks: %d\n", msg.OpenCount, msg.ClickCount)

			if len(msg.Tags) > 0 {
				fmt.Printf("     Tags: %v\n", msg.Tags)
			}

			if msg.IsBounceNotification {
				fmt.Printf("     Bounce: %s\n", getStringValue(msg.BounceClassification))
				if msg.ReferenceMessageID != nil {
					fmt.Printf("     Reference Message: %d\n", *msg.ReferenceMessageID)
				}
			}

			if len(msg.DeliveryAttempts) > 0 {
				fmt.Printf("     Latest Delivery Attempt: %s - %s\n",
					msg.DeliveryAttempts[len(msg.DeliveryAttempts)-1].Status,
					msg.DeliveryAttempts[len(msg.DeliveryAttempts)-1].Log)
			}
			fmt.Println()
		}
	} else {
		fmt.Println("  No recent messages found.")
	}
}

func showPerformanceSummary() {
	fmt.Println("  Email Performance Benchmarks:")
	fmt.Println("  ┌─────────────────┬────────────┬─────────────┐")
	fmt.Println("  │ Metric          │ Your Rate  │ Industry Avg│")
	fmt.Println("  ├─────────────────┼────────────┼─────────────┤")
	fmt.Println("  │ Delivery Rate   │ Calculate  │ 95-98%      │")
	fmt.Println("  │ Open Rate       │ from stats │ 15-25%      │")
	fmt.Println("  │ Click Rate      │ above      │ 2-5%        │")
	fmt.Println("  │ Bounce Rate     │            │ <2%         │")
	fmt.Println("  │ Complaint Rate  │            │ <0.1%       │")
	fmt.Println("  └─────────────────┴────────────┴─────────────┘")

	fmt.Println("\n  Tips to Improve Performance:")
	fmt.Println("  • Maintain list hygiene by removing bounced addresses")
	fmt.Println("  • Use double opt-in to ensure engaged subscribers")
	fmt.Println("  • Personalize subject lines and content")
	fmt.Println("  • Test different send times for your audience")
	fmt.Println("  • Monitor and act on engagement metrics")
	fmt.Println("  • Implement proper authentication (SPF, DKIM, DMARC)")
}

func calculatePercentage(value, total int) float64 {
	if total == 0 {
		return 0
	}
	return float64(value) / float64(total) * 100
}

func getStringValue(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return "N/A"
}

func handleError(err error) {
	if apiErr, ok := err.(*api.APIError); ok {
		log.Printf("API Error: %s (Status: %d)", apiErr.Error(), apiErr.StatusCode)
		if len(apiErr.Raw) > 0 {
			log.Printf("Response Body: %s", string(apiErr.Raw))
		}
	} else {
		log.Printf("Error: %v", err)
	}
}
