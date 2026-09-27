package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/api"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/google/uuid"
)

func main() {
	client := api.NewAPIClient(api.WithAPIKey(os.Getenv("AHASEND_API_KEY")))
	accountID := uuid.MustParse(os.Getenv("AHASEND_ACCOUNT_ID"))
	ctx := context.Background()

	webhookID := uuid.MustParse("c5a32c40-b351-439f-8230-779daed3e42c")

	response, _, err := client.WebhooksAPI.UpdateWebhook(ctx, accountID, webhookID, requests.UpdateWebhookRequest{
		Name:      ahasend.String("Failures"),
		URL:       ahasend.String("https://example.com/webhook"),
		Enabled:   ahasend.Bool(true),
		OnBounced: ahasend.Bool(true),
	})
	if err != nil {
		log.Fatalf("Error updating webhook: %v", err)
	}
	fmt.Printf("Updated webhook %s\n", response.ID)
}
