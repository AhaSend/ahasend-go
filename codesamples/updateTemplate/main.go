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

	templateID := uuid.MustParse("11111111-1111-4111-8111-111111111111")

	// Only the fields you set change. They go to the template's draft, the
	// same draft the dashboard edits; Publish publishes the whole draft. An
	// empty ReplyTo clears the reply-to address.
	response, _, err := client.TemplatesAPI.UpdateTemplate(ctx, accountID, templateID, requests.UpdateTemplateRequest{
		Subject: ahasend.String("Your password reset link"),
		ReplyTo: &common.SenderAddress{},
		Publish: true,
	})
	if err != nil {
		log.Fatalf("Error updating template: %v", err)
	}
	fmt.Printf("Updated template %s, has draft: %t\n", response.Name, response.HasDraft)
}
