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

	// Only a draft can be deleted.
	response, _, err := client.CampaignsAPI.DeleteCampaign(ctx, accountID, campaignID)
	if err != nil {
		log.Fatalf("Error deleting campaign: %v", err)
	}
	fmt.Println(response.Message)
}
