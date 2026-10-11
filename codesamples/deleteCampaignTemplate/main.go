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

	templateID := uuid.MustParse("11111111-1111-4111-8111-111111111111")

	// Campaigns made from the template keep their own copies.
	response, _, err := client.CampaignTemplatesAPI.DeleteCampaignTemplate(ctx, accountID, templateID)
	if err != nil {
		log.Fatalf("Error deleting campaign template: %v", err)
	}
	fmt.Println(response.Message)
}
