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

	// Campaign templates are listed without their content; Get Campaign
	// Template returns it.
	response, _, err := client.CampaignTemplatesAPI.ListCampaignTemplates(ctx, accountID, requests.ListCampaignTemplatesParams{
		Limit: ahasend.Int32(100),
	})
	if err != nil {
		log.Fatalf("Error listing campaign templates: %v", err)
	}
	for _, template := range response.Data {
		fmt.Printf("%s %s (%s)\n", template.ID, template.Name, template.Editor)
	}
}
