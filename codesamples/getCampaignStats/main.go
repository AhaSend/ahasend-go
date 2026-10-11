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

	response, _, err := client.CampaignsAPI.GetCampaignStats(ctx, accountID, campaignID)
	if err != nil {
		log.Fatalf("Error getting campaign stats: %v", err)
	}
	fmt.Printf("Delivered %d of %d\n", response.Delivered, response.Recipients)
	// A rate is nil until the report is counted and its base is not zero.
	if response.Rates.OpenRate != nil {
		fmt.Printf("Open rate: %.2f%%\n", *response.Rates.OpenRate)
	}
	if response.Report != nil {
		for status, count := range response.Report.ByStatus {
			fmt.Printf("%s: %d\n", status, count)
		}
	}
}
