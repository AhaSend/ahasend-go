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

	response, _, err := client.ListsAPI.GetListContacts(ctx, accountID, listID, requests.GetListContactsParams{
		SubscriptionStatus: ahasend.String(requests.ListContactStatusConfirmed),
		IncludeContacts:    ahasend.Bool(true),
	})
	if err != nil {
		log.Fatalf("Error listing list contacts: %v", err)
	}
	for _, member := range response.Data {
		fmt.Printf("%s %s\n", member.Email, member.SubscriptionStatus)
	}
}
