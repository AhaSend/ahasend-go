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

	subAccountID := uuid.MustParse("7d3c9f5e-2a41-4b8e-9c6d-0f1e2a3b4c5d")

	// List responses never carry secret keys; only create returns one.
	response, _, err := client.SubAccountsAPI.ListSubAccountAPIKeys(ctx, accountID, subAccountID, &common.PaginationParams{
		Limit: ahasend.Int32(50),
	})
	if err != nil {
		log.Fatalf("Error listing sub-account API keys: %v", err)
	}
	for _, key := range response.Data {
		fmt.Printf("%s %s\n", key.ID, key.Label)
	}
}
