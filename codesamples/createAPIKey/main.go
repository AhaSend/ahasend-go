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

	// The SDK sends a fresh Idempotency-Key with this POST and reuses it on
	// its own retries, so a retried request is not applied twice.
	response, _, err := client.APIKeysAPI.CreateAPIKey(ctx, accountID, requests.CreateAPIKeyRequest{
		Label: "My API Key",
		Scopes: []string{
			"messages:read:all",
			"domains:read",
		},
		// Optional: restrict this key to specific source IPs (CIDR blocks or
		// bare IPv4/IPv6 addresses). Omit it to allow any IP.
		IPAllowList: []string{"203.0.113.0/24", "198.51.100.7"},
	})
	if err != nil {
		log.Fatalf("Error creating API key: %v", err)
	}
	// response.SecretKey is returned only on create: put it in your secret
	// store now, and never log it.
	fmt.Printf("Created API key %s (%s)\n", response.ID, response.Label)
}
