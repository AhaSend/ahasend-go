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

	response, _, err := client.ListsAPI.GetLists(ctx, accountID, requests.GetListsParams{
		Limit: ahasend.Int32(100),
		Name:  ahasend.String("newsletter"),
	})
	if err != nil {
		log.Fatalf("Error listing lists: %v", err)
	}
	for _, list := range response.Data {
		fmt.Printf("%s %s (%d contacts)\n", list.ID, list.Name, list.ContactCount)
	}
}
