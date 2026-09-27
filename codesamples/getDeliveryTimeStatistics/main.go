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

	fromTime := time.Now().Add(-30 * 24 * time.Hour)

	response, _, err := client.StatisticsAPI.GetDeliveryTimeStatistics(ctx, accountID, requests.GetDeliveryTimeStatisticsParams{
		FromTime:         &fromTime,
		SenderDomain:     ahasend.String("example.com"),
		RecipientDomains: ahasend.String("gmail.com"),
		GroupBy:          ahasend.String("week"),
	})
	if err != nil {
		log.Fatalf("Error getting delivery time statistics: %v", err)
	}
	fmt.Printf("Got %d weekly buckets\n", len(response.Data))
}
