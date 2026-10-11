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

	campaignID := uuid.MustParse("11111111-1111-4111-8111-111111111111")

	// The copy is a new draft. Leave Name nil to name it "Copy of" and the
	// source's name.
	response, _, err := client.CampaignsAPI.DuplicateCampaign(ctx, accountID, campaignID, requests.DuplicateCampaignRequest{
		Name: ahasend.String("November newsletter"),
	})
	if err != nil {
		log.Fatalf("Error duplicating campaign: %v", err)
	}
	fmt.Printf("Created draft %s %s\n", response.ID, response.Name)
}
