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

	response, _, err := client.APIKeysAPI.GetAPIKeys(ctx, accountID, &common.PaginationParams{
		Limit: ahasend.Int32(20),
	})
	if err != nil {
		log.Fatalf("Error getting API keys: %v", err)
	}
	for _, key := range response.Data {
		fmt.Printf("%s %s\n", key.ID, key.Label)
	}
}
