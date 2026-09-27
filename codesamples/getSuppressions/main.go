package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/AhaSend/ahasend-go/api"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/google/uuid"
)

func main() {
	client := api.NewAPIClient(api.WithAPIKey(os.Getenv("AHASEND_API_KEY")))
	accountID := uuid.MustParse(os.Getenv("AHASEND_ACCOUNT_ID"))
	ctx := context.Background()

	response, _, err := client.SuppressionsAPI.GetSuppressions(ctx, accountID, requests.GetSuppressionsParams{})
	if err != nil {
		log.Fatalf("Error getting suppressions: %v", err)
	}
	for _, suppression := range response.Data {
		fmt.Printf("%s until %s: %s\n", suppression.Email, suppression.ExpiresAt, suppression.Reason)
	}
}
