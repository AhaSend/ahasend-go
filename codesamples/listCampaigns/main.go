package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/api"
	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/google/uuid"
)

func main() {
	client := api.NewAPIClient(api.WithAPIKey(os.Getenv("AHASEND_API_KEY")))
	accountID := uuid.MustParse(os.Getenv("AHASEND_ACCOUNT_ID"))
	ctx := context.Background()

	// Leave Status nil to list campaigns in every status.
	response, _, err := client.CampaignsAPI.ListCampaigns(ctx, accountID, requests.ListCampaignsParams{
		Status: ahasend.String(common.CampaignStatusScheduled),
		Limit:  ahasend.Int32(100),
	})
	if err != nil {
		log.Fatalf("Error listing campaigns: %v", err)
	}
	for _, campaign := range response.Data {
		if campaign.SendAt != nil {
			fmt.Printf("%s %s sends at %s\n", campaign.ID, campaign.Name, campaign.SendAt)
		}
	}
}
