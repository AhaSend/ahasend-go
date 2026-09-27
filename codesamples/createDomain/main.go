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
	response, _, err := client.DomainsAPI.CreateDomain(ctx, accountID, requests.CreateDomainRequest{
		Domain: "example.com",
	})
	if err != nil {
		log.Fatalf("Error creating domain: %v", err)
	}
	// Publish these records at your DNS provider.
	for _, record := range response.DNSRecords {
		fmt.Printf("%s %s %s\n", record.Type, record.Host, record.Content)
	}
}
