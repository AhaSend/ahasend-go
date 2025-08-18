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

	// Create authentication context
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, apiKey)

	fmt.Println("=== Scheduled Email Examples ===")

	// Example 1: Schedule email for later today
	scheduleForLaterToday(ctx, client, accountID)

	// Example 2: Schedule weekly newsletter
	scheduleWeeklyNewsletter(ctx, client, accountID)

	// Example 3: Schedule birthday emails
	scheduleBirthdayEmails(ctx, client, accountID)

	// Example 4: Schedule email campaigns
	scheduleEmailCampaign(ctx, client, accountID)

	// Example 5: Cancel scheduled email
	cancelScheduledEmail(ctx, client, accountID)
}

func scheduleForLaterToday(ctx context.Context, client *api.APIClient, accountID uuid.UUID) {
	fmt.Println("1. Schedule Email for Later Today")
	fmt.Println("──────────────────────────────────")

	// Schedule for 2 hours from now
	sendTime := time.Now().Add(2 * time.Hour)
	fmt.Printf("Scheduling email for: %s\n", sendTime.Format("3:04 PM MST"))

	message := requests.CreateMessageRequest{
		From: common.SenderAddress{
			Email: "reminders@yourdomain.com",
			Name:  ahasend.String("Reminder Service"),
		},
		Recipients: []common.Recipient{
			{
				Email: "user@example.com",
				Name:  ahasend.String("User"),
			},
		},
		Subject: "Reminder: Meeting at 3 PM",
		HtmlContent: ahasend.String(`
			<h2>Meeting Reminder</h2>
			<p>This is a friendly reminder about your meeting scheduled for 3 PM today.</p>
			<p><strong>Topic:</strong> Quarterly Review</p>
			<p><strong>Location:</strong> Conference Room A</p>
			<p>See you there!</p>
		`),
		Tags: []string{"reminder", "meeting"},
		Schedule: &common.MessageSchedule{
			FirstAttempt: &sendTime,
		},
	}

	response, httpResp, err := client.MessagesAPI.CreateMessage(ctx, accountID, message)

	if err != nil {
		handleError(err)
	} else {
		fmt.Printf("✅ Email scheduled successfully! (Status: %d)\n", httpResp.StatusCode)
		if len(response.Data) > 0 {
			fmt.Printf("   Message ID: %s\n", *response.Data[0].ID)
			fmt.Printf("   Status: %s\n", response.Data[0].Status)
		}
	}
}

func scheduleWeeklyNewsletter(ctx context.Context, client *api.APIClient, accountID uuid.UUID) {
	fmt.Println("\n2. Schedule Weekly Newsletter")
	fmt.Println("──────────────────────────────")

	// Schedule for next Friday at 9 AM
	now := time.Now()
	daysUntilFriday := (5 - int(now.Weekday()) + 7) % 7
	if daysUntilFriday == 0 && now.Weekday() == time.Friday {
		daysUntilFriday = 7 // Next Friday, not today
	}

	nextFriday := now.AddDate(0, 0, daysUntilFriday)
	sendTime := time.Date(nextFriday.Year(), nextFriday.Month(), nextFriday.Day(),
		9, 0, 0, 0, now.Location())

	fmt.Printf("Scheduling newsletter for: %s\n", sendTime.Format("Monday, January 2, 2006 at 3:04 PM MST"))

	// Newsletter content
	message := requests.CreateMessageRequest{
		From: common.SenderAddress{
			Email: "newsletter@yourdomain.com",
			Name:  ahasend.String("Weekly Newsletter"),
		},
		Recipients: []common.Recipient{
			{Email: "subscriber1@example.com"},
			{Email: "subscriber2@example.com"},
			{Email: "subscriber3@example.com"},
		},
		Subject: fmt.Sprintf("Weekly Newsletter - %s", sendTime.Format("Jan 2, 2006")),
		HtmlContent: ahasend.String(fmt.Sprintf(`
			<html>
				<body>
					<h1>Weekly Newsletter</h1>
					<h2>Week of %s</h2>
					<hr>
					<h3>🚀 Product Updates</h3>
					<p>New features released this week...</p>

					<h3>📈 Company News</h3>
					<p>What's happening at the company...</p>

					<h3>🎯 Tips & Tricks</h3>
					<p>Pro tips to get the most out of our service...</p>

					<hr>
					<p><small>You're receiving this because you subscribed to our newsletter.</small></p>
				</body>
			</html>
		`, sendTime.Format("January 2, 2006"))),
		Tags: []string{"newsletter", "weekly", "marketing"},
		Schedule: &common.MessageSchedule{
			FirstAttempt: &sendTime,
		},
	}

	response, httpResp, err := client.MessagesAPI.CreateMessage(ctx, accountID, message)

	if err != nil {
		handleError(err)
	} else {
		fmt.Printf("✅ Newsletter scheduled! (Status: %d)\n", httpResp.StatusCode)
		if len(response.Data) > 0 {
			fmt.Printf("   Will be sent to %d recipients\n", len(message.Recipients))
		}
	}
}

func scheduleBirthdayEmails(ctx context.Context, client *api.APIClient, accountID uuid.UUID) {
	fmt.Println("\n3. Schedule Birthday Emails")
	fmt.Println("─────────────────────────────")

	// Simulate user birthdays (in a real app, you'd query your database)
	birthdays := []struct {
		email    string
		name     string
		birthday time.Time
	}{
		{"alice@example.com", "Alice Johnson", time.Date(2024, time.March, 15, 9, 0, 0, 0, time.UTC)},
		{"bob@example.com", "Bob Smith", time.Date(2024, time.March, 20, 9, 0, 0, 0, time.UTC)},
	}

	for _, user := range birthdays {
		// Schedule birthday email for 9 AM on their birthday
		sendTime := user.birthday

		message := requests.CreateMessageRequest{
			From: common.SenderAddress{
				Email: "birthday@yourdomain.com",
				Name:  ahasend.String("Birthday Team"),
			},
			Recipients: []common.Recipient{
				{
					Email: user.email,
					Name:  ahasend.String(user.name),
				},
			},
			Subject: fmt.Sprintf("🎉 Happy Birthday, %s!", user.name),
			HtmlContent: ahasend.String(fmt.Sprintf(`
				<html>
					<body style="font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto;">
						<div style="text-align: center; padding: 40px;">
							<h1 style="color: #ff6b6b;">🎉 Happy Birthday! 🎂</h1>
							<h2>Dear %s,</h2>
							<p style="font-size: 18px; line-height: 1.6;">
								We hope your special day is filled with happiness, laughter,
								and all your favorite things!
							</p>
							<div style="background: #f8f9fa; padding: 20px; border-radius: 8px; margin: 20px 0;">
								<h3>🎁 Birthday Special!</h3>
								<p>Enjoy 20%% off your next order with code: <strong>BIRTHDAY20</strong></p>
								<p><small>Valid until %s</small></p>
							</div>
							<p>Have a wonderful year ahead!</p>
							<p style="color: #666;"><small>The Team at Your Company</small></p>
						</div>
					</body>
				</html>
			`, user.name, sendTime.AddDate(0, 0, 30).Format("January 2, 2006"))),
			Tags: []string{"birthday", "special-offer", "customer-engagement"},
			Schedule: &common.MessageSchedule{
				FirstAttempt: &sendTime,
			},
		}

		fmt.Printf("Scheduling birthday email for %s on %s... ",
			user.name, sendTime.Format("Jan 2, 2006"))

		response, httpResp, err := client.MessagesAPI.CreateMessage(ctx, accountID, message)

		if err != nil {
			fmt.Printf("❌ Error\n")
			handleError(err)
		} else {
			fmt.Printf("✅ Scheduled (Status: %d)\n", httpResp.StatusCode)
			if len(response.Data) > 0 {
				fmt.Printf("   Message ID: %s\n", *response.Data[0].ID)
			}
		}
	}
}

func scheduleEmailCampaign(ctx context.Context, client *api.APIClient, accountID uuid.UUID) {
	fmt.Println("\n4. Schedule Email Campaign")
	fmt.Println("─────────────────────────────")

	// Schedule a product launch email for next Monday at 10 AM
	now := time.Now()
	daysUntilMonday := (1 - int(now.Weekday()) + 7) % 7
	if daysUntilMonday == 0 && now.Weekday() == time.Monday {
		daysUntilMonday = 7 // Next Monday
	}

	launchDate := now.AddDate(0, 0, daysUntilMonday)
	sendTime := time.Date(launchDate.Year(), launchDate.Month(), launchDate.Day(),
		10, 0, 0, 0, now.Location())

	fmt.Printf("Scheduling product launch campaign for: %s\n",
		sendTime.Format("Monday, January 2, 2006 at 3:04 PM MST"))

	message := requests.CreateMessageRequest{
		From: common.SenderAddress{
			Email: "marketing@yourdomain.com",
			Name:  ahasend.String("Product Team"),
		},
		Recipients: []common.Recipient{
			{Email: "customer1@example.com", Name: ahasend.String("Customer 1")},
			{Email: "customer2@example.com", Name: ahasend.String("Customer 2")},
			{Email: "customer3@example.com", Name: ahasend.String("Customer 3")},
		},
		Subject: "🚀 Introducing Our Latest Innovation!",
		HtmlContent: ahasend.String(`
			<html>
				<body style="font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto;">
					<header style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); color: white; padding: 40px; text-align: center;">
						<h1>🚀 Product Launch!</h1>
						<p style="font-size: 20px; margin: 0;">The wait is over</p>
					</header>

					<main style="padding: 40px;">
						<h2>Introducing ProductName 2.0</h2>
						<p style="font-size: 16px; line-height: 1.6;">
							After months of development and customer feedback, we're excited to
							announce the launch of ProductName 2.0 with incredible new features:
						</p>

						<ul style="font-size: 16px; line-height: 1.8;">
							<li>⚡ 10x faster performance</li>
							<li>🎨 Redesigned user interface</li>
							<li>🔧 Advanced customization options</li>
							<li>📱 Mobile-first experience</li>
						</ul>

						<div style="text-align: center; margin: 40px 0;">
							<a href="#" style="background: #667eea; color: white; padding: 15px 30px; text-decoration: none; border-radius: 5px; font-weight: bold;">
								Try It Now
							</a>
						</div>

						<p style="color: #666; font-size: 14px;">
							Early bird pricing: 30% off for the first 100 customers!
						</p>
					</main>

					<footer style="background: #f8f9fa; padding: 20px; text-align: center; color: #666;">
						<p>Thanks for being an amazing customer!</p>
					</footer>
				</body>
			</html>
		`),
		Tags: []string{"product-launch", "campaign", "announcement"},
		Schedule: &common.MessageSchedule{
			FirstAttempt: &sendTime,
		},
		Tracking: &common.Tracking{
			Open:  ahasend.Bool(true),
			Click: ahasend.Bool(true),
		},
	}

	_, httpResp, err := client.MessagesAPI.CreateMessage(ctx, accountID, message)

	if err != nil {
		handleError(err)
	} else {
		fmt.Printf("✅ Campaign scheduled! (Status: %d)\n", httpResp.StatusCode)
		fmt.Printf("   Recipients: %d\n", len(message.Recipients))
		fmt.Printf("   Tracking: Opens and clicks enabled\n")
	}
}

func cancelScheduledEmail(ctx context.Context, client *api.APIClient, accountID uuid.UUID) {
	fmt.Println("\n5. Cancel Scheduled Email")
	fmt.Println("────────────────────────────")

	// Note: This is a conceptual example. The actual cancellation API
	// depends on your specific AhaSend API implementation.

	fmt.Println("To cancel a scheduled email:")
	fmt.Println("1. Keep track of Message IDs when scheduling")
	fmt.Println("2. Use the Cancel Message API before the scheduled send time")
	fmt.Println("3. Emails can only be cancelled before they're sent")

	// Example of how you might cancel (adjust based on actual API)
	messageID := "example-message-id" // You would get this from the schedule response

	fmt.Printf("Attempting to cancel message: %s\n", messageID)

	// This is pseudocode - adjust based on your actual API
	response, httpResp, err := client.MessagesAPI.CancelMessage(ctx, accountID, messageID)

	if err != nil {
		fmt.Printf("❌ Could not cancel: %v\n", err)
		fmt.Println("   Message may have already been sent or does not exist")
	} else {
		fmt.Printf("✅ Message cancelled! (Status: %d)\n", httpResp.StatusCode)
		if response != nil {
			fmt.Printf("   Response: %s\n", response.Message)
		}
	}

	fmt.Println("\n=== Best Practices for Scheduled Emails ===")
	fmt.Println("1. Always specify timezone to avoid confusion")
	fmt.Println("2. Use RFC3339 format for timestamps")
	fmt.Println("3. Test with short delays first (minutes, not days)")
	fmt.Println("4. Store message IDs if you need to cancel later")
	fmt.Println("5. Consider user's local timezone for better engagement")
	fmt.Println("6. Add buffer time for campaign reviews")
	fmt.Println("7. Use scheduling for time-sensitive content only")
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
