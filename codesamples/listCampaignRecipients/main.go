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

	// Set exactly one of Status and BounceClassification, each a key of the
	// campaign's stats report.
	response, _, err := client.CampaignsAPI.ListCampaignRecipients(ctx, accountID, campaignID, requests.ListCampaignRecipientsParams{
		Status: ahasend.String("Bounced"),
		Limit:  ahasend.Int32(100),
	})
	if err != nil {
		log.Fatalf("Error listing campaign recipients: %v", err)
	}
	for _, recipient := range response.Data {
		reason := ""
		if recipient.BounceReason != nil {
			reason = *recipient.BounceReason
		}
		fmt.Printf("%s (%s) %s\n", recipient.Email, recipient.ListName, reason)
	}
}
