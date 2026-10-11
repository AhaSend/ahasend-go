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

	// A template with no draft returns 404; its HasDraft is false.
	response, _, err := client.TemplatesAPI.GetTemplateDraft(ctx, accountID, templateID)
	if err != nil {
		log.Fatalf("Error getting template draft: %v", err)
	}
	fmt.Printf("Draft subject: %s, changed at %s\n", response.Subject, response.UpdatedAt)
}
