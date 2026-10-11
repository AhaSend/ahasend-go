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

	templateID := uuid.MustParse("11111111-1111-4111-8111-111111111111")

	// Only the fields you set change, at once. Campaigns already made from
	// the template keep their own copies.
	response, _, err := client.CampaignTemplatesAPI.UpdateCampaignTemplate(ctx, accountID, templateID, requests.UpdateCampaignTemplateRequest{
		Name: ahasend.String("Monthly newsletter, 2026"),
	})
	if err != nil {
		log.Fatalf("Error updating campaign template: %v", err)
	}
	fmt.Printf("Updated campaign template %s\n", response.Name)
}
