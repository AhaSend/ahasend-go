package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/AhaSend/ahasend-go/api"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/google/uuid"
)

func main() {
	client := api.NewAPIClient(api.WithAPIKey(os.Getenv("AHASEND_API_KEY")))
	accountID := uuid.MustParse(os.Getenv("AHASEND_ACCOUNT_ID"))
	ctx := context.Background()

	// Pass the raw email; the SDK encodes it once in the path.
	response, _, err := client.SegmentsAPI.GetContactSegments(ctx, accountID, "User+Tag@Example.COM", requests.GetContactSegmentsParams{})
	if err != nil {
		log.Fatalf("Error getting the contact's segments: %v", err)
	}
	for _, segment := range response.Data {
		fmt.Println(segment.Name)
	}
}
