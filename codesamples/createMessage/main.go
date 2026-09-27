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

	// The SDK sends a fresh Idempotency-Key with every POST and reuses it on its
	// own retries. Deriving the key from your own ID instead also makes a rerun
	// of the same job (after a crash, say) return the first send, not a second.
	orderID := "1234"

	response, _, err := client.MessagesAPI.CreateMessage(ctx, accountID, requests.CreateMessageRequest{
		From: common.SenderAddress{
			Email: "info@example.com",
			Name:  ahasend.String("Example Corp."),
		},
		Recipients: []common.Recipient{
			{
				Email: "john@example.com",
				Name:  ahasend.String("John Smith"),
			},
		},
		Subject:     "Hello",
		TextContent: ahasend.String("Hello world!"),
		Sandbox:     ahasend.Bool(true),
	}, api.WithIdempotencyKey("order-"+orderID+"-receipt"))
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
