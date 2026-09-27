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

	routeID := uuid.MustParse("c5a32c40-b351-439f-8230-779daed3e42c")

	response, _, err := client.RoutesAPI.UpdateRoute(ctx, accountID, routeID, requests.UpdateRouteRequest{
		Name: ahasend.String("Updated Name"),
		URL:  ahasend.String("https://example.com/new-tickets"),
	})
	if err != nil {
		log.Fatalf("Error updating route: %v", err)
	}
	fmt.Printf("Updated route %s\n", response.ID)
}
