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

	response, _, err := client.WebhooksAPI.CreateWebhook(ctx, accountID, requests.CreateWebhookRequest{
		Name:             "Failures",
		URL:              "https://mystartup.com/webhook",
		Scope:            "global",
		Enabled:          ahasend.Bool(true),
		OnBounced:        true,
		OnTransientError: true,
		OnFailed:         true,
	})
	if err != nil {
		log.Fatalf("Error creating webhook: %v", err)
	}
	// Use the secret to verify the signature of each delivery.
	fmt.Printf("Created webhook %s, secret: %s\n", response.ID, response.Secret)
}
