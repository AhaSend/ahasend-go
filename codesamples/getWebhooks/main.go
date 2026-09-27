package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/api"
	"github.com/google/uuid"
)

func main() {
	client := api.NewAPIClient(api.WithAPIKey(os.Getenv("AHASEND_API_KEY")))
	accountID := uuid.MustParse(os.Getenv("AHASEND_ACCOUNT_ID"))
	ctx := context.Background()

	params := api.GetWebhooksParams{Enabled: ahasend.Bool(true)}
	params.Limit = ahasend.Int32(100)

	response, _, err := client.WebhooksAPI.GetWebhooks(ctx, accountID, params)
	if err != nil {
		log.Fatalf("Error getting webhooks: %v", err)
	}
	for _, webhook := range response.Data {
		fmt.Printf("%s %s -> %s\n", webhook.ID, webhook.Name, webhook.URL)
	}
}
