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

	response, _, err := client.ListsAPI.CreateList(ctx, accountID, requests.CreateListRequest{
		Name:        "Product updates",
		Description: ahasend.String("Monthly release notes"),
		Tags:        []string{"product"},
	}, api.WithIdempotencyKey("list-create-0001"))
	if err != nil {
		log.Fatalf("Error creating list: %v", err)
	}
	fmt.Printf("Created list %s\n", response.ID)
}
