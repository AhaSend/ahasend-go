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

	response, _, err := client.DomainsAPI.UpdateDomain(ctx, accountID, "example.com", requests.UpdateDomainRequest{
		TrackingSubdomain: ahasend.String("click"),
	})
	if err != nil {
		log.Fatalf("Error updating domain: %v", err)
	}
	fmt.Printf("Updated domain %s\n", response.Domain)
}
