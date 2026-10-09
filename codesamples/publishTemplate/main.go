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

	templateID := uuid.MustParse("11111111-1111-4111-8111-111111111111")

	// The whole draft is published, including changes made in the dashboard.
	// The SDK sends a fresh Idempotency-Key with this POST and reuses it on
	// its own retries.
	response, _, err := client.TemplatesAPI.PublishTemplate(ctx, accountID, templateID)
	if err != nil {
		log.Fatalf("Error publishing template: %v", err)
	}
	fmt.Printf("Published template %s\n", response.Name)
}
