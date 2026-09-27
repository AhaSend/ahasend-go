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

	userID := uuid.MustParse("5b9a1c1e-8f7d-4a55-9d0e-2c1f3b7a6e42")

	response, _, err := client.AccountsAPI.RemoveAccountMember(ctx, accountID, userID)
	if err != nil {
		log.Fatalf("Error removing account member: %v", err)
	}
	fmt.Println(response.Message)
}
