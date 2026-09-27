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

	response, _, err := client.MessagesAPI.GetMessages(ctx, accountID, requests.GetMessagesParams{
		Status: ahasend.String("Bounced,Failed"),
		Sender: ahasend.String("info@example.com"),
		Tags:   []string{"billing", "urgent"},
	})
	if err != nil {
		log.Fatalf("Error getting messages: %v", err)
	}
	for _, message := range response.Data {
		fmt.Printf("%s %s %s\n", message.MessageID, message.Recipient, message.Status)
	}
}
