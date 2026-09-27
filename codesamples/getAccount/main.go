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

	response, _, err := client.AccountsAPI.GetAccount(ctx, accountID)
	if err != nil {
		log.Fatalf("Error getting account: %v", err)
	}
	fmt.Printf("%s %s\n", response.ID, response.Name)
}
