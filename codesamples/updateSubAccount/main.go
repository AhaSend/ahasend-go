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

	subAccountID := uuid.MustParse("7d3c9f5e-2a41-4b8e-9c6d-0f1e2a3b4c5d")

	response, _, err := client.SubAccountsAPI.UpdateSubAccount(ctx, accountID, subAccountID, requests.UpdateSubAccountRequest{
		Name:          ahasend.String("Acme Subsidiary"),
		MonthlyCredit: ahasend.Int64(50000),
	})
	if err != nil {
		log.Fatalf("Error updating sub account: %v", err)
	}
	fmt.Printf("%s monthly credit: %d\n", response.Name, response.MonthlyCredit)
}
