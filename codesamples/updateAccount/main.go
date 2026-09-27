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

	// Update the account name and open/click tracking settings
	response, _, err := client.AccountsAPI.UpdateAccount(ctx, accountID, requests.UpdateAccountRequest{
		Name:        ahasend.String("Acme Corp"),
		TrackOpens:  ahasend.Bool(true),
		TrackClicks: ahasend.Bool(true),
	})
	if err != nil {
		log.Fatalf("Error updating account: %v", err)
	}
	fmt.Printf("Updated account %s\n", response.Name)
}
