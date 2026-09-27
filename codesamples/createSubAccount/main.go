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
	response, _, err := client.SubAccountsAPI.CreateSubAccount(ctx, accountID, requests.CreateSubAccountRequest{
		Name:          "Acme Subsidiary",
		Website:       "acme.example.com",
		MonthlyCredit: ahasend.Int64(0),
	})
	if err != nil {
		log.Fatalf("Error creating sub account: %v", err)
	}
	fmt.Printf("Created sub account %s\n", response.ID)
}
