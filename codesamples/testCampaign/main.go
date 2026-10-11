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

	campaignID := uuid.MustParse("11111111-1111-4111-8111-111111111111")

	response, _, err := client.CampaignsAPI.TestCampaign(ctx, accountID, campaignID, requests.TestCampaignRequest{
		To: []string{"Ann <ann@yourdomain.com>", "bob@yourdomain.com"},
	})
	if err != nil {
		log.Fatalf("Error sending test: %v", err)
	}
	if response.ResultUnknown {
		fmt.Println("The test may have gone out; check the inboxes before you send another")
		return
	}
	fmt.Printf("Sent to %v\n", response.SentTo)
	for _, failure := range response.Failed {
		fmt.Printf("Refused %s: %s\n", failure.Email, failure.Message)
	}
}
