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

	// Record that the contact unsubscribed from this list. Setting confirmed
	// would subscribe them again, so only do that when they asked for it.
	response, _, err := client.ListsAPI.UpsertListContact(ctx, accountID, listID, "user+tag@example.com", requests.UpsertListContactRequest{
		SubscriptionStatus: ahasend.String(requests.ListContactStatusUnsubscribed),
	})
	if err != nil {
		log.Fatalf("Error updating the list membership: %v", err)
	}
	fmt.Printf("%s is %s\n", response.Email, response.SubscriptionStatus)
}
