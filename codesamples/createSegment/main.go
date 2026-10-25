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

	listID := uuid.MustParse("11111111-1111-4111-8111-111111111111")
	campaignID := uuid.MustParse("22222222-2222-4222-8222-222222222222")

	// Contacts in Germany, confirmed on the list, who opened the campaign or
	// opened at least 3 campaigns in the last 90 days. A group holds
	// conditions only.
	definition := common.SegmentDefinition{
		SchemaVersion: common.SegmentSchemaVersion,
		Match:         common.SegmentMatchAll,
		Conditions: []common.SegmentRule{
			{Condition: &common.SegmentCondition{
				Field: "attributes.country",
				Op:    common.SegmentOpEqual,
				Value: &common.SegmentValue{String: ahasend.String("DE")},
			}},
			{Condition: &common.SegmentCondition{
				Field: common.SegmentFieldList,
				Op:    common.SegmentOpIn,
				Value: &common.SegmentValue{String: ahasend.String(listID.String())},
			}},
			{Group: &common.SegmentGroup{
				Match: common.SegmentMatchAny,
				Conditions: []common.SegmentCondition{
					{
						Field: common.SegmentFieldCampaign,
						Op:    common.SegmentOpOpened,
						Value: &common.SegmentValue{String: ahasend.String(campaignID.String())},
					},
					{
						Field: common.SegmentFieldCampaignCount,
						Op:    common.SegmentOpOpenedAtLeast,
						Value: &common.SegmentValue{Count: &common.SegmentCampaignCount{Count: 3, Since: "90 days ago"}},
					},
				},
			}},
		},
	}

	// The SDK sends a fresh Idempotency-Key with this POST and reuses it on
	// its own retries, so a retried request creates no second segment.
	response, _, err := client.SegmentsAPI.CreateSegment(ctx, accountID, requests.CreateSegmentRequest{
		Name:       "Engaged German customers",
		Definition: definition,
		Timezone:   ahasend.String("Europe/Berlin"),
	})
	if err != nil {
		log.Fatalf("Error creating segment: %v", err)
	}
	// A new segment is building until its members are found.
	fmt.Printf("Created segment %s (%s)\n", response.ID, response.State)
}
