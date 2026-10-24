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

	// The published copy is the one a send uses, so build a send from its
	// variables. GetTemplate returns the draft instead while one exists.
	response, _, err := client.TemplatesAPI.GetPublishedTemplate(ctx, accountID, templateID)
	if err != nil {
		log.Fatalf("Error getting published template: %v", err)
	}
	fmt.Printf("Published subject: %s, draft waiting: %t\n", response.Subject, response.HasDraft)
	for _, variable := range response.Variables {
		fmt.Printf("%s required=%t kind=%s\n", variable.Name, variable.Required, variable.Kind)
	}
}
