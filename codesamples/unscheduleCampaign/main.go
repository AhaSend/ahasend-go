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

	// A validating or scheduled campaign goes back to draft. To change the
	// send time, unschedule it and send it again.
	response, _, err := client.CampaignsAPI.UnscheduleCampaign(ctx, accountID, campaignID)
	if err != nil {
		log.Fatalf("Error unscheduling campaign: %v", err)
	}
	fmt.Printf("Campaign %s is %s\n", response.ID, response.Status)
}
