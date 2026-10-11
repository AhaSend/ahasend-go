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

	// HTML content makes an html campaign template; MJML would make an
	// advanced one. The text is made from the HTML when it is left out.
	response, _, err := client.CampaignTemplatesAPI.CreateCampaignTemplate(ctx, accountID, requests.CreateCampaignTemplateRequest{
		Name: "Monthly newsletter",
		Content: &requests.TemplateContentInput{
			HTML: ahasend.String(`<p>Hi {{ first_name }},</p><p>Here is our news.</p><p><a href="{{ unsubscribe_url }}">Unsubscribe</a></p>`),
		},
	})
	if err != nil {
		log.Fatalf("Error creating campaign template: %v", err)
	}
	fmt.Printf("Created %s campaign template %s\n", response.Editor, response.ID)
}
