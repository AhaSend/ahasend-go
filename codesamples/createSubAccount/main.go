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

	response, _, err := client.SubAccountsAPI.CreateSubAccount(ctx, accountID, requests.CreateSubAccountRequest{
		Name:          "Acme Subsidiary",
		Website:       "acme.example.com",
		MonthlyCredit: ahasend.Int64(0),
	}, api.WithIdempotencyKey("subacct-20240101-acme"))
	if err != nil {
		log.Fatalf("Error creating sub account: %v", err)
	}
	fmt.Printf("Created sub account %s\n", response.ID)
}
