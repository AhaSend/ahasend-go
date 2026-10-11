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

	campaignID := uuid.MustParse("11111111-1111-4111-8111-111111111111")

	response, _, err := client.CampaignsAPI.GetCampaign(ctx, accountID, campaignID)
	if err != nil {
		log.Fatalf("Error getting campaign: %v", err)
	}
	fmt.Printf("%s is %s, sending to %d lists\n", response.Name, response.Status, len(response.ListIDs))
	// A draft lists what keeps it from being sent.
	for _, issue := range response.Issues {
		fmt.Printf("%s: %s\n", issue.Code, issue.Message)
	}
}
