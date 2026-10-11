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

	response, _, err := client.CampaignTemplatesAPI.GetCampaignTemplate(ctx, accountID, templateID)
	if err != nil {
		log.Fatalf("Error getting campaign template: %v", err)
	}
	// A campaign can be made from the template only once it has HTML.
	hasHTML := response.Content != nil && response.Content.HTML != ""
	fmt.Printf("%s (%s), has HTML: %t\n", response.Name, response.Editor, hasHTML)
}
