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

	keyID := uuid.MustParse("c5a32c40-b351-439f-8230-779daed3e42c")

	response, _, err := client.APIKeysAPI.UpdateAPIKey(ctx, accountID, keyID, requests.UpdateAPIKeyRequest{
		Label: ahasend.String("My API Key"),
		Scopes: &[]string{
			"messages:read:all",
			"domains:read",
		},
		// Replaces the allowed source IPs. &[]string{} clears the list (any IP);
		// leaving the field nil keeps the current list.
		IPAllowList: &[]string{"203.0.113.0/24"},
	})
	if err != nil {
		log.Fatalf("Error updating API key: %v", err)
	}
	fmt.Printf("Updated API key %s\n", response.ID)
}
