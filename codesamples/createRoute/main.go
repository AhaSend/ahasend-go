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

	response, _, err := client.RoutesAPI.CreateRoute(ctx, accountID, requests.CreateRouteRequest{
		Name:         "My Route",
		URL:          "https://mystartup.com/tickets",
		Recipient:    "ticket-*@mystartup.com",
		Enabled:      ahasend.Bool(true),
		Attachments:  true,
		Headers:      true,
		StripReplies: true,
	})
	if err != nil {
		log.Fatalf("Error creating route: %v", err)
	}
	fmt.Printf("Created route %s\n", response.ID)
}
