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

	response, _, err := client.ListsAPI.GetContactLists(ctx, accountID, "user+tag@example.com", requests.GetContactListsParams{
		SubscriptionStatus: ahasend.String(requests.ListContactStatusConfirmed),
	})
	if err != nil {
		log.Fatalf("Error getting the contact's lists: %v", err)
	}
	for _, membership := range response.Data {
		if membership.List != nil {
			fmt.Printf("%s %s\n", membership.List.Name, membership.SubscriptionStatus)
		}
	}
}
