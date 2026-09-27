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

	listID := uuid.MustParse("11111111-1111-4111-8111-111111111111")

	response, _, err := client.ListsAPI.GetList(ctx, accountID, listID)
	if err != nil {
		log.Fatalf("Error getting list: %v", err)
	}
	fmt.Printf("%s (%d contacts)\n", response.Name, response.ContactCount)
}
