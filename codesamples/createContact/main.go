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

	response, _, err := client.ContactsAPI.CreateContact(ctx, accountID, requests.CreateContactRequest{
		Email:      "person@example.com",
		FirstName:  ahasend.String("Pat"),
		Attributes: map[string]any{"customer": true},
	}, api.WithIdempotencyKey("contact-create-0001"))
	if err != nil {
		log.Fatalf("Error creating contact: %v", err)
	}
	fmt.Printf("Created contact %s\n", response.ID)
}
