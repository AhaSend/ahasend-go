package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

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

	// Leave SendAt nil to send now.
	response, _, err := client.CampaignsAPI.SendCampaign(ctx, accountID, campaignID, requests.SendCampaignRequest{
		SendAt: ahasend.Time(time.Now().Add(24 * time.Hour)),
	})
	if notReady, ok := api.AsCampaignNotReady(err); ok {
		for _, issue := range notReady.Issues {
			fmt.Printf("%s: %s\n", issue.Code, issue.Message)
		}
		os.Exit(1)
	}
	if err != nil {
		log.Fatalf("Error sending campaign: %v", err)
	}
	fmt.Printf("Campaign %s is %s\n", response.ID, response.Status)
	for _, warning := range response.Warnings {
		fmt.Println("Warning:", warning)
	}
}
