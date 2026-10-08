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

	// Versions come newest first, at most 50, with no pagination.
	response, _, err := client.TemplatesAPI.GetTemplateVersions(ctx, accountID, templateID)
	if err != nil {
		log.Fatalf("Error listing template versions: %v", err)
	}
	for _, version := range response.Data {
		publisher := "unknown"
		if version.PublishedBy != nil {
			publisher = version.PublishedBy.Type + " " + version.PublishedBy.ID.String()
		}
		fmt.Printf("v%d %s published by %s\n", version.Version, version.PublishedAt, publisher)
	}
}
