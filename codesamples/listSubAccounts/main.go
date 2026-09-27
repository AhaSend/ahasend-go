package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/api"
	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/google/uuid"
)

func main() {
	client := api.NewAPIClient(api.WithAPIKey(os.Getenv("AHASEND_API_KEY")))
	accountID := uuid.MustParse(os.Getenv("AHASEND_ACCOUNT_ID"))
	ctx := context.Background()

	response, _, err := client.SubAccountsAPI.ListSubAccounts(ctx, accountID, &common.PaginationParams{
		Limit: ahasend.Int32(50),
	})
	if err != nil {
		log.Fatalf("Error listing sub accounts: %v", err)
	}
	for _, sub := range response.Data {
		fmt.Printf("%s %s %s\n", sub.ID, sub.Name, sub.Status)
	}
}
