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

	response, _, err := client.SegmentsAPI.GetSegments(ctx, accountID, requests.GetSegmentsParams{
		Limit: ahasend.Int32(50),
	})
	if err != nil {
		log.Fatalf("Error listing segments: %v", err)
	}
	for _, segment := range response.Data {
		fmt.Printf("%s %s (%d members)\n", segment.Name, segment.State, segment.MemberCount)
	}
}
