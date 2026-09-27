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

	// The SDK sends a fresh Idempotency-Key with this POST and reuses it on
	// its own retries, so a retried request is not applied twice.
	response, _, err := client.SMTPCredentialsAPI.CreateSMTPCredential(ctx, accountID, requests.CreateSMTPCredentialRequest{
		Name:    "My SMTP Credential",
		Scope:   "global",
		Sandbox: true,
	})
	if err != nil {
		log.Fatalf("Error creating SMTP credential: %v", err)
	}
	// The password is returned only on create: store it now.
	fmt.Printf("Username: %s, password: %s\n", response.Username, response.Password)
}
