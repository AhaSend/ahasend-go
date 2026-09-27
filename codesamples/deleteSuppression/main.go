package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/api"
	"github.com/google/uuid"
)

func main() {
	client := api.NewAPIClient(api.WithAPIKey(os.Getenv("AHASEND_API_KEY")))
	accountID := uuid.MustParse(os.Getenv("AHASEND_ACCOUNT_ID"))
	ctx := context.Background()

	// Removes the suppression for one address; pass nil instead of a domain
	// to remove it for every domain.
	response, _, err := client.SuppressionsAPI.DeleteSuppression(
		ctx,
		accountID,
		"info@example.com",
		ahasend.String("notifications.example.com"),
	)
	if err != nil {
		log.Fatalf("Error deleting suppression: %v", err)
	}
	fmt.Println(response.Message)
}
