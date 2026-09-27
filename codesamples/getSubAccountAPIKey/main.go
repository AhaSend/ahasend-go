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

	subAccountID := uuid.MustParse("7d3c9f5e-2a41-4b8e-9c6d-0f1e2a3b4c5d")
	keyID := uuid.MustParse("c5a32c40-b351-439f-8230-779daed3e42c")

	response, _, err := client.SubAccountsAPI.GetSubAccountAPIKey(ctx, accountID, subAccountID, keyID)
	if err != nil {
		log.Fatalf("Error getting sub-account API key: %v", err)
	}
	fmt.Printf("%s %s (%d scopes)\n", response.ID, response.Label, len(response.Scopes))
}
