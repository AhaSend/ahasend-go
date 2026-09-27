package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/api"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/google/uuid"
)

func main() {
	client := api.NewAPIClient(api.WithAPIKey(os.Getenv("AHASEND_API_KEY")))
	accountID := uuid.MustParse(os.Getenv("AHASEND_ACCOUNT_ID"))
	ctx := context.Background()

	response, _, err := client.SuppressionsAPI.CreateSuppression(ctx, accountID, requests.CreateSuppressionRequest{
		Email:     "test@example.com",
		Reason:    ahasend.String("Inbox full"),
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
	})
	if err != nil {
		log.Fatalf("Error creating suppression: %v", err)
	}
	fmt.Printf("Created %d suppression(s)\n", len(response.Data))
}
