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

	templateID := uuid.MustParse("11111111-1111-4111-8111-111111111111")

	// The template supplies the body, and the subject and sender when the
	// request leaves them out. Each recipient gives its own variable values.
	// For a template with no default sender, set From to an address on one of
	// your domains.
	response, _, err := client.MessagesAPI.CreateTemplateMessage(ctx, accountID, requests.CreateTemplateMessageRequest{
		TemplateID: templateID,
		Recipients: []common.Recipient{
			{
				Email: "john@example.com",
				Name:  ahasend.String("John Smith"),
				Substitutions: map[string]interface{}{
					"first_name": "John",
					"order_id":   "12345",
				},
			},
		},
		Sandbox: ahasend.Bool(true),
	})
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}
	// Each recipient gets its own message, or an error saying why it was not sent.
	for _, message := range response.Data {
		if message.ID != nil {
			fmt.Printf("%s: %s\n", message.Recipient.Email, *message.ID)
		} else if message.Error != nil {
			fmt.Printf("%s: %s\n", message.Recipient.Email, *message.Error)
		}
	}
}
