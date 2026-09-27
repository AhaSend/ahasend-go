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

	// The SDK sends a fresh Idempotency-Key with this POST and reuses it on
	// its own retries, so a retried request is not applied twice.
	response, _, err := client.AccountsAPI.AddAccountMember(ctx, accountID, requests.AddMemberRequest{
		Email: "teammate@example.com",
		Name:  ahasend.String("Jordan Lee"),
		Role:  "Developer",
	})
	if err != nil {
		log.Fatalf("Error adding account member: %v", err)
	}
	fmt.Printf("Added user %s as %s\n", response.UserID, response.Role)
}
