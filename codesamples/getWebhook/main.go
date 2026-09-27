package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/AhaSend/ahasend-go/api"
	"github.com/google/uuid"
)

func main() {
	client := api.NewAPIClient(api.WithAPIKey(os.Getenv("AHASEND_API_KEY")))
	accountID := uuid.MustParse(os.Getenv("AHASEND_ACCOUNT_ID"))
	ctx := context.Background()

	webhookID := uuid.MustParse("c5a32c40-b351-439f-8230-779daed3e42c")

	response, _, err := client.WebhooksAPI.GetWebhook(ctx, accountID, webhookID)
	if err != nil {
		log.Fatalf("Error getting webhook: %v", err)
	}
	fmt.Printf("%s -> %s (enabled: %t)\n", response.Name, response.URL, response.Enabled)
}
