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

	response, _, err := client.SMTPCredentialsAPI.GetSMTPCredentials(ctx, accountID, nil)
	if err != nil {
		log.Fatalf("Error getting SMTP credentials: %v", err)
	}
	for _, credential := range response.Data {
		fmt.Printf("%s %s %s\n", credential.ID, credential.Name, credential.Username)
	}
}
