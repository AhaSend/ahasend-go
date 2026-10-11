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

	// Only the fields you set change, and only while the campaign is a
	// draft. A ReplyTo that points at "" clears the reply-to address.
	response, _, err := client.CampaignsAPI.UpdateCampaign(ctx, accountID, campaignID, requests.UpdateCampaignRequest{
		Subject: ahasend.String("Our news for October"),
		ReplyTo: ahasend.String(""),
	})
	if err != nil {
		log.Fatalf("Error updating campaign: %v", err)
	}
	fmt.Printf("Updated campaign %s, %d issues left\n", response.Name, len(response.Issues))
}
