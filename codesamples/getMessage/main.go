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

	// The Message-ID returned by Create Message, or its bare UUID.
	response, _, err := client.MessagesAPI.GetMessageByAPIID(ctx, accountID, "<8b1f2e8c-5d6a-4c3b-9e7f-1a2b3c4d5e6f@example.com>")
	if err != nil {
		log.Fatalf("Error getting message: %v", err)
	}
	fmt.Printf("Subject: %s\n", response.Subject)
	fmt.Printf("Status: %s\n", response.Status)
	fmt.Printf("Recipient: %s\n", response.Recipient)
	fmt.Printf("Opens: %d, clicks: %d\n", response.OpenCount, response.ClickCount)
}
