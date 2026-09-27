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

	// Only the fields you set change. A nil attribute value deletes that key.
	response, _, err := client.ContactsAPI.UpdateContact(ctx, accountID, "user+tag@example.com", requests.UpdateContactRequest{
		FirstName:  ahasend.String("Jordan"),
		Attributes: map[string]any{"obsolete_key": nil},
	})
	if err != nil {
		log.Fatalf("Error updating contact: %v", err)
	}
	fmt.Printf("Updated contact %s\n", response.ID)
}
