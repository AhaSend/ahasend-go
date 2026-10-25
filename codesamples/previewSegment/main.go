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

	// Nothing is saved. Count is nil when the contacts cannot be counted
	// within 10 seconds.
	response, _, err := client.SegmentsAPI.PreviewSegment(ctx, accountID, requests.PreviewSegmentRequest{
		Definition: common.SegmentDefinition{
			SchemaVersion: common.SegmentSchemaVersion,
			Match:         common.SegmentMatchAll,
			Conditions: []common.SegmentRule{
				{Condition: &common.SegmentCondition{
					Field: common.SegmentFieldCreatedAt,
					Op:    common.SegmentOpAfter,
					Value: &common.SegmentValue{String: ahasend.String("30 days ago")},
				}},
			},
		},
	})
	if err != nil {
		log.Fatalf("Error previewing segment: %v", err)
	}
	if response.Count != nil {
		fmt.Printf("%d contacts match\n", *response.Count)
	}
	for _, contact := range response.Contacts {
		fmt.Println(contact.Email)
	}
}
