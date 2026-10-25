package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/AhaSend/ahasend-go/api"
	"github.com/google/uuid"
)

func main() {
	client := api.NewAPIClient(api.WithAPIKey(os.Getenv("AHASEND_API_KEY")))
	accountID := uuid.MustParse(os.Getenv("AHASEND_ACCOUNT_ID"))
	ctx := context.Background()

	segmentID := uuid.MustParse("11111111-1111-4111-8111-111111111111")

	response, _, err := client.SegmentsAPI.GetSegment(ctx, accountID, segmentID)
	if err != nil {
		log.Fatalf("Error getting segment: %v", err)
	}
	// MarketableCount is nil while the segment is not ready.
	if response.MarketableCount != nil {
		fmt.Printf("%s: %d members, %d a campaign would reach\n", response.Name, response.MemberCount, *response.MarketableCount)
	} else {
		fmt.Printf("%s: %d members (%s)\n", response.Name, response.MemberCount, response.State)
	}
}
