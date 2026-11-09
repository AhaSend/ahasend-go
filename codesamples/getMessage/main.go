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

	// The id that Create Message returned for the recipient.
	response, _, err := client.MessagesAPI.GetMessageByAPIID(ctx, accountID, "0199f1c2-7d3a-7b4e-9a1f-2c3d4e5f6a7b")
	if err != nil {
		log.Fatalf("Error getting message: %v", err)
	}
	fmt.Printf("Subject: %s\n", response.Subject)
	fmt.Printf("Status: %s\n", response.Status)
	fmt.Printf("Recipient: %s\n", response.Recipient)
	fmt.Printf("Opens: %d, clicks: %d\n", response.OpenCount, response.ClickCount)
}
