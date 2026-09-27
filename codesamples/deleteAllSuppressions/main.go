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

	// Pass a domain instead of nil to limit the wipe to that domain.
	response, _, err := client.SuppressionsAPI.DeleteAllSuppressions(ctx, accountID, nil)
	if err != nil {
		log.Fatalf("Error deleting all suppressions: %v", err)
	}
	fmt.Println(response.Message)
}
