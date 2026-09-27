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

	listID := uuid.MustParse("11111111-1111-4111-8111-111111111111")

	// Name each contact by email or by ID.
	contactID := uuid.MustParse("22222222-2222-4222-8222-222222222222")
	response, _, err := client.ListsAPI.BatchAddListContacts(ctx, accountID, listID, requests.BatchAddListContactsRequest{
		Data: []requests.BatchAddListContactInput{
			{Email: ahasend.String("one@example.com")},
			{ID: &contactID},
		},
	}, api.WithIdempotencyKey("list-contacts-batch-0001"))
	if err != nil {
		log.Fatalf("Error adding contacts to the list: %v", err)
	}
	// Each item succeeds or fails on its own.
	fmt.Printf("added=%d skipped=%d failed=%d\n", response.Added, response.Skipped, response.Failed)
}
