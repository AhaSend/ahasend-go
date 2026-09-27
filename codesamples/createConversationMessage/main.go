package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/api"
	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/google/uuid"
)

func main() {
	client := api.NewAPIClient(api.WithAPIKey(os.Getenv("AHASEND_API_KEY")))
	accountID := uuid.MustParse(os.Getenv("AHASEND_ACCOUNT_ID"))
	ctx := context.Background()

	// A conversational message sets the To/CC/BCC headers directly, so
	// recipients can see each other (BCC excluded). The combined
	// To + CC + BCC count must not exceed 50. Unlike Create Message,
	// this endpoint does not support template substitutions.
	response, _, err := client.MessagesAPI.CreateConversationMessage(ctx, accountID, requests.CreateConversationMessageRequest{
		From: common.SenderAddress{
			Email: "info@example.com",
			Name:  ahasend.String("Example Corp."),
		},
		To: []common.SenderAddress{
			{Email: "john@example.com", Name: ahasend.String("John Smith")},
			{Email: "jane@example.com", Name: ahasend.String("Jane Doe")},
		},
		CC: []common.SenderAddress{
			{Email: "hr@example.com", Name: ahasend.String("Example Corp HR")},
		},
		BCC: []common.SenderAddress{
			{Email: "archive@example.com"},
		},
		Subject:     "Welcome to Example Corp",
		HtmlContent: ahasend.String("<h1>Dear John and Jane, welcome to Example Corp!</h1>"),
		TextContent: ahasend.String("Dear John and Jane, welcome to Example Corp!"),
		Sandbox:     ahasend.Bool(true),
	})
	if err != nil {
		log.Fatalf("Error sending conversational message: %v", err)
	}
	for _, message := range response.Data {
		if message.ID != nil {
			fmt.Printf("%s: %s\n", message.Recipient.Email, *message.ID)
		}
	}
}
