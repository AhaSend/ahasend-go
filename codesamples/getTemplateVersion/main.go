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
	versionID := uuid.MustParse("22222222-2222-4222-8222-222222222222")

	response, _, err := client.TemplatesAPI.GetTemplateVersion(ctx, accountID, templateID, versionID)
	if err != nil {
		log.Fatalf("Error getting template version: %v", err)
	}
	fmt.Printf("v%d subject: %s\n", response.Version, response.Subject)
	if response.Content != nil {
		fmt.Println(response.Content.Text)
	}
}
