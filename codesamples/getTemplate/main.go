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

	response, _, err := client.TemplatesAPI.GetTemplate(ctx, accountID, templateID)
	if err != nil {
		log.Fatalf("Error getting template: %v", err)
	}
	// These are the working copy's variables: the draft's while one exists.
	// A send uses the published copy's, which GetPublishedTemplate returns.
	// Kind says what shape of value to send; a list variable takes a slice of
	// elements with the fields in ItemFields.
	for _, variable := range response.Variables {
		fmt.Printf("%s required=%t kind=%s item_fields=%v\n", variable.Name, variable.Required, variable.Kind, variable.ItemFields)
	}
}
