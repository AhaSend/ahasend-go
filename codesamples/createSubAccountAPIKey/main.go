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

	subAccountID := uuid.MustParse("7d3c9f5e-2a41-4b8e-9c6d-0f1e2a3b4c5d")

	// Bootstrap an API key for the sub account; the idempotency key makes
	// retries safe.
	response, _, err := client.SubAccountsAPI.CreateSubAccountAPIKey(ctx, accountID, subAccountID, requests.CreateAPIKeyRequest{
		Label:  "Bootstrap key",
		Scopes: []string{"messages:send:all", "domains:read"},
		// Optional: restrict the key to specific source IPs.
		IPAllowList: []string{"203.0.113.0/24"},
	}, api.WithIdempotencyKey("child-bootstrap-key-20240101-acme"))
	if err != nil {
		log.Fatalf("Error creating sub-account API key: %v", err)
	}
	// SecretKey is returned only on create (and on an exact idempotent
	// replay): store it now.
	fmt.Printf("Created API key %s, secret key: %s\n", response.ID, *response.SecretKey)
}
