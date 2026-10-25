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

	segmentID := uuid.MustParse("11111111-1111-4111-8111-111111111111")

	// Only the fields you set change. A new name or description keeps the
	// members; a new definition or time zone builds them again.
	response, _, err := client.SegmentsAPI.UpdateSegment(ctx, accountID, segmentID, requests.UpdateSegmentRequest{
		Name:        ahasend.String("Engaged customers"),
		Description: ahasend.String("Opened a campaign recently"),
	})
	if err != nil {
		log.Fatalf("Error updating segment: %v", err)
	}
	fmt.Printf("Updated segment %s (%s)\n", response.Name, response.State)
}
