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

	// The SDK sends a fresh Idempotency-Key with this POST and reuses it on
	// its own retries, so a retried request is not applied twice.
	response, _, err := client.ContactsAPI.BatchUpsertContacts(ctx, accountID, requests.BatchUpsertContactsRequest{
		Data: []requests.BatchUpsertContactInput{
			{Email: "one@example.com"},
			{Email: "two@example.com", Unsubscribed: ahasend.Bool(true)},
		},
	})
	if err != nil {
		log.Fatalf("Error upserting contacts: %v", err)
	}
	// Each item succeeds or fails on its own.
	fmt.Printf("created=%d updated=%d failed=%d\n", response.Created, response.Updated, response.Failed)
	for _, result := range response.Data {
		if result.Reason != "" {
			fmt.Printf("%s: %s (%s)\n", result.Email, result.Outcome, result.Reason)
		}
	}
}
