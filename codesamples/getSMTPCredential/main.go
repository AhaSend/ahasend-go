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

	smtpCredentialID := uuid.MustParse("0194f7ac-2f6a-7d29-9f0c-9a27b19f2d1c")

	response, _, err := client.SMTPCredentialsAPI.GetSMTPCredential(ctx, accountID, smtpCredentialID)
	if err != nil {
		log.Fatalf("Error getting SMTP credential: %v", err)
	}
	fmt.Printf("%s %s (%s)\n", response.Name, response.Username, response.Scope)
}
