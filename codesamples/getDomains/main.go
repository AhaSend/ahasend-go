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

	// Pass a *bool instead of nil to filter by DNS validity.
	response, _, err := client.DomainsAPI.GetDomains(ctx, accountID, nil, nil)
	if err != nil {
		log.Fatalf("Error getting domains: %v", err)
	}
	for _, domain := range response.Data {
		fmt.Printf("%s dns_valid=%t\n", domain.Domain, domain.DNSValid)
	}
}
