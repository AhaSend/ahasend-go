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

	fmt.Println("=== Webhook Management Example ===")

	// 1. List existing webhooks
	fmt.Println("1. Listing existing webhooks...")
	listWebhooks(ctx, client, accountID)

	// 2. Create a new webhook
	fmt.Println("\n2. Creating a new webhook...")
	webhookID := createWebhook(ctx, client, accountID)

	// 3. Get webhook details
	if webhookID != "" {
		fmt.Printf("\n3. Getting webhook details for ID: %s\n", webhookID)
		getWebhookDetails(ctx, client, accountID, webhookID)

		// 4. Update webhook
		fmt.Printf("\n4. Updating webhook ID: %s\n", webhookID)
		updateWebhook(ctx, client, accountID, webhookID)

		// 5. Delete webhook (commented out to prevent accidental deletion)
		// fmt.Printf("\n5. Deleting webhook ID: %s\n", webhookID)
		// deleteWebhook(ctx, client, accountID, webhookID)
	}

	// 6. Example webhook receiver server
	fmt.Println("\n6. Example Webhook Receiver Server")
	fmt.Println("Run this to test receiving webhooks:")
	fmt.Println("go run webhook_receiver.go")
	showWebhookReceiverExample()
}

func listWebhooks(ctx context.Context, client *api.APIClient, accountID uuid.UUID) {
	params := api.GetWebhooksParams{
		PaginationParams: common.PaginationParams{
			Limit: ahasend.Int32(10),
		},
	}
	response, httpResp, err := client.WebhooksAPI.GetWebhooks(ctx, accountID, params)

	if err != nil {
		handleError(err)
		return
	}

	if httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 {
		if len(response.Data) > 0 {
			fmt.Printf("Found %d webhooks:\n", len(response.Data))
			for _, webhook := range response.Data {
				fmt.Printf("  - Name: %s\n", webhook.Name)
				fmt.Printf("    URL: %s\n", webhook.URL)
				fmt.Printf("    Enabled: %v\n", webhook.Enabled)
				fmt.Printf("    Event Triggers: ")
				var triggers []string
				if webhook.OnReception {
					triggers = append(triggers, "reception")
				}
				if webhook.OnDelivered {
					triggers = append(triggers, "delivered")
				}
				if webhook.OnBounced {
					triggers = append(triggers, "bounced")
				}
				if webhook.OnOpened {
					triggers = append(triggers, "opened")
				}
				if webhook.OnClicked {
					triggers = append(triggers, "clicked")
				}
				fmt.Printf("%v\n", triggers)
			}
		} else {
			fmt.Println("No webhooks found.")
		}
	}
}

func createWebhook(ctx context.Context, client *api.APIClient, accountID uuid.UUID) string {
	request := requests.CreateWebhookRequest{
		Name:        "Email Event Webhook",
		URL:         "https://your-domain.com/webhooks/ahasend",
		Enabled:     true,
		OnReception: true,
		OnDelivered: true,
		OnBounced:   true,
		OnOpened:    true,
		OnClicked:   true,
	}

	response, httpResp, err := client.WebhooksAPI.CreateWebhook(ctx, accountID, request)

	if err != nil {
		handleError(err)
		return ""
	}

	if httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 {
		fmt.Println("Webhook created successfully!")
		if response != nil {
			fmt.Printf("  ID: %s\n", response.ID)
			fmt.Printf("  Name: %s\n", response.Name)
			fmt.Printf("  URL: %s\n", response.URL)
			return response.ID.String()
		}
	}
	return ""
}

func getWebhookDetails(ctx context.Context, client *api.APIClient, accountID uuid.UUID, webhookID string) {
	webhookUUID, err := uuid.Parse(webhookID)
	if err != nil {
		log.Printf("Invalid webhook ID: %v", err)
		return
	}

	response, httpResp, err := client.WebhooksAPI.GetWebhook(ctx, accountID, webhookUUID)

	if err != nil {
		handleError(err)
		return
	}

	if httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 && response != nil {
		webhook := response

		fmt.Println("Webhook Details:")
		fmt.Printf("  ID: %s\n", webhook.ID)
		fmt.Printf("  Name: %s\n", webhook.Name)
		fmt.Printf("  URL: %s\n", webhook.URL)
		fmt.Printf("  Enabled: %v\n", webhook.Enabled)
		fmt.Printf("  Created: %s\n", webhook.CreatedAt)
		fmt.Printf("  Updated: %s\n", webhook.UpdatedAt)

		fmt.Println("  Event Triggers:")
		fmt.Printf("    On Reception: %v\n", webhook.OnReception)
		fmt.Printf("    On Delivered: %v\n", webhook.OnDelivered)
		fmt.Printf("    On Bounced: %v\n", webhook.OnBounced)
		fmt.Printf("    On Opened: %v\n", webhook.OnOpened)
		fmt.Printf("    On Clicked: %v\n", webhook.OnClicked)
		fmt.Printf("    On Failed: %v\n", webhook.OnFailed)
		fmt.Printf("    On Suppressed: %v\n", webhook.OnSuppressed)
	}
}

func updateWebhook(ctx context.Context, client *api.APIClient, accountID uuid.UUID, webhookID string) {
	webhookUUID, err := uuid.Parse(webhookID)
	if err != nil {
		log.Printf("Invalid webhook ID: %v", err)
		return
	}

	request := requests.UpdateWebhookRequest{
		Name:                 ahasend.String("Updated Email Event Webhook"),
		Enabled:              ahasend.Bool(true),
		OnReception:          ahasend.Bool(true),
		OnDelivered:          ahasend.Bool(true),
		OnBounced:            ahasend.Bool(true),
		OnOpened:             ahasend.Bool(true),
		OnClicked:            ahasend.Bool(true),
		OnSuppressionCreated: ahasend.Bool(true), // Added new event
	}

	response, httpResp, err := client.WebhooksAPI.UpdateWebhook(ctx, accountID, webhookUUID, request)

	if err != nil {
		handleError(err)
		return
	}

	if httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 {
		fmt.Println("Webhook updated successfully!")
		if response != nil {
			fmt.Printf("  Name: %s\n", response.Name)
			fmt.Println("  Updated Event Triggers:")
			fmt.Printf("    On Reception: %v\n", response.OnReception)
			fmt.Printf("    On Delivered: %v\n", response.OnDelivered)
			fmt.Printf("    On Bounced: %v\n", response.OnBounced)
			fmt.Printf("    On Opened: %v\n", response.OnOpened)
			fmt.Printf("    On Clicked: %v\n", response.OnClicked)
			fmt.Printf("    On Suppression Created: %v\n", response.OnSuppressionCreated)
		}
	}
}

func deleteWebhook(ctx context.Context, client *api.APIClient, accountID uuid.UUID, webhookID string) {
	webhookUUID, err := uuid.Parse(webhookID)
	if err != nil {
		log.Printf("Invalid webhook ID: %v", err)
		return
	}

	response, httpResp, err := client.WebhooksAPI.DeleteWebhook(ctx, accountID, webhookUUID)

	if err != nil {
		handleError(err)
		return
	}

	if httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 {
		fmt.Println("Webhook deleted successfully!")
		if response != nil {
			fmt.Printf("Message: %s\n", response.Message)
		}
	}
}

func showWebhookReceiverExample() {
	text := `
=== Example Webhook Receiver Server ===

Create a file 'webhook_receiver.go' with this content:

package main

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
)

type WebhookEvent struct {
    Event     string                 ` + "`json:\"event\"`" + `
    Timestamp string                 ` + "`json:\"timestamp\"`" + `
    Data      map[string]interface{} ` + "`json:\"data\"`" + `
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
    // Read the request body
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Failed to read body", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // Verify webhook signature (if using webhook secret)
    signature := r.Header.Get("X-AhaSend-Signature")
    webhookSecret := "your-webhook-secret" // Store this securely

    if signature != "" {
        expectedSig := computeSignature(body, webhookSecret)
        if signature != expectedSig {
            http.Error(w, "Invalid signature", http.StatusUnauthorized)
            return
        }
    }

    // Parse the webhook event
    var event WebhookEvent
    if err := json.Unmarshal(body, &event); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Process the event based on type
    switch event.Event {
    case "message.sent":
        fmt.Printf("Email sent: %v\n", event.Data["message_id"])
    case "message.delivered":
        fmt.Printf("Email delivered: %v\n", event.Data["message_id"])
    case "message.bounced":
        fmt.Printf("Email bounced: %v, Reason: %v\n",
            event.Data["message_id"], event.Data["reason"])
    case "message.complained":
        fmt.Printf("Spam complaint: %v\n", event.Data["message_id"])
    case "message.opened":
        fmt.Printf("Email opened: %v\n", event.Data["message_id"])
    case "message.clicked":
        fmt.Printf("Link clicked: %v, URL: %v\n",
            event.Data["message_id"], event.Data["url"])
    default:
        fmt.Printf("Unknown event: %s\n", event.Event)
    }

    // Send success response
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}

func computeSignature(payload []byte, secret string) string {
    h := hmac.New(sha256.New, []byte(secret))
    h.Write(payload)
    return hex.EncodeToString(h.Sum(nil))
}

func main() {
    http.HandleFunc("/webhooks/ahasend", webhookHandler)

    fmt.Println("Webhook receiver listening on :8080...")
    fmt.Println("Configure your webhook URL as: http://your-domain.com:8080/webhooks/ahasend")

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

Then run: go run webhook_receiver.go
`
	fmt.Print(text)
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
