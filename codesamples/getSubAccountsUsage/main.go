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

	response, _, err := client.SubAccountsAPI.GetSubAccountsUsage(ctx, accountID)
	if err != nil {
		log.Fatalf("Error getting sub-accounts usage: %v", err)
	}
	fmt.Printf("Allocation method: %s\n", response.AllocationMethod)
	fmt.Printf("Total: %d messages, %.2f %s\n", response.Total.ReceptionCount, response.Total.AllocatedCost, response.Currency)
	for _, sub := range response.SubAccounts {
		name := "(unnamed)"
		if sub.Name != nil {
			name = *sub.Name
		} else if sub.AccountID != nil {
			name = sub.AccountID.String()
		}
		fmt.Printf("  %s: %d messages, %.2f %s\n", name, sub.ReceptionCount, sub.AllocatedCost, response.Currency)
	}
}
