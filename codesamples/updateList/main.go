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

	listID := uuid.MustParse("11111111-1111-4111-8111-111111111111")

	// Only the fields you set change; an empty tags slice clears the tags.
	response, _, err := client.ListsAPI.UpdateList(ctx, accountID, listID, requests.UpdateListRequest{
		Name: ahasend.String("Product updates"),
		Tags: &[]string{},
	})
	if err != nil {
		log.Fatalf("Error updating list: %v", err)
	}
	fmt.Printf("Updated list %s\n", response.Name)
}
