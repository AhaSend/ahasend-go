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

	fromTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	toTime := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)

	response, _, err := client.StatisticsAPI.GetEmailStatistics(ctx, accountID, requests.GetEmailStatisticsParams{
		FromTime:  &fromTime,
		ToTime:    &toTime,
		GroupBy:   ahasend.String("month"),
		Breakdown: ahasend.String("domain"),
	})
	if err != nil {
		log.Fatalf("Error getting email statistics: %v", err)
	}
	fmt.Printf("Accepted %d messages in 2025\n", response.Totals.ReceptionCount)
	for _, row := range response.Data {
		fmt.Printf("%s %s: %d accepted, %d delivered\n", *row.SenderDomain, row.FromTimestamp.Format("2006-01"), row.ReceptionCount, row.DeliveredCount)
	}
	if response.Pagination.HasMore {
		fmt.Printf("Next page of domains: After = %q\n", *response.Pagination.NextCursor)
	}
}
