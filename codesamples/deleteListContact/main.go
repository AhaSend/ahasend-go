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

	listID := uuid.MustParse("11111111-1111-4111-8111-111111111111")

	// Removing drops the membership. To stop sending to the contact while
	// keeping the record, set its status to unsubscribed instead.
	response, _, err := client.ListsAPI.DeleteListContact(ctx, accountID, listID, "user+tag@example.com")
	if err != nil {
		log.Fatalf("Error removing the contact from the list: %v", err)
	}
	fmt.Println(response.Message)
}
