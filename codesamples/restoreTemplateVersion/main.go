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

	templateID := uuid.MustParse("11111111-1111-4111-8111-111111111111")
	versionID := uuid.MustParse("22222222-2222-4222-8222-222222222222")

	// The version goes into the draft; Publish then publishes it, so sends use
	// it again.
	response, _, err := client.TemplatesAPI.RestoreTemplateVersion(ctx, accountID, templateID, versionID, requests.RestoreTemplateVersionRequest{
		Publish: true,
	})
	if err != nil {
		log.Fatalf("Error restoring template version: %v", err)
	}
	fmt.Printf("Restored template %s, has draft: %t\n", response.Name, response.HasDraft)
}
