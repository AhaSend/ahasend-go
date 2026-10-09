//go:build ignore

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/api"
	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/google/uuid"
)

func main() {
	// Get API credentials from environment variables
	apiKey := os.Getenv("AHASEND_API_KEY")
	if apiKey == "" {
		log.Fatal("AHASEND_API_KEY environment variable is required")
	}

	accountIDStr := os.Getenv("AHASEND_ACCOUNT_ID")
	if accountIDStr == "" {
		log.Fatal("AHASEND_ACCOUNT_ID environment variable is required")
	}

	accountID, err := uuid.Parse(accountIDStr)
	if err != nil {
		log.Fatalf("Invalid account ID: %v", err)
	}

	templateIDStr := os.Getenv("AHASEND_TEMPLATE_ID")
	if templateIDStr == "" {
		log.Fatal("AHASEND_TEMPLATE_ID environment variable is required")
	}

	templateID, err := uuid.Parse(templateIDStr)
	if err != nil {
		log.Fatalf("Invalid template ID: %v", err)
	}

	// Create a new API client
	client := api.NewAPIClient(
		api.WithAPIKey(apiKey),
	)

	// Create authentication context
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, apiKey)

	// Read the template to see which variables a send has to supply.
	// This needs the templates:read scope; the send itself does not.
	template, _, err := client.TemplatesAPI.GetTemplate(ctx, accountID, templateID)
	if err != nil {
		log.Fatalf("Error reading template: %v", err)
	}

	fmt.Printf("Template: %s\n", template.Name)
	for _, variable := range template.Variables {
		fmt.Printf("  %s (required: %t)\n", variable.Name, variable.Required)
	}

	// The template supplies the subject, preview text and both bodies, so the
	// request carries no content. Leave Subject empty to use the template's.
	// Each recipient gives its own values for the template's variables.
	message := requests.CreateTemplateMessageRequest{
		TemplateID: templateID,
		Recipients: []common.Recipient{
			{
				Email:         "recipient@example.com",
				Name:          ahasend.String("Recipient Name"),
				Substitutions: map[string]interface{}{"first_name": "Pat", "company": "Your Company"},
			},
		},
		Tags: []string{"template", "test"},
	}

	// A template with a default sender supplies it when From is nil. A
	// template without one needs a sender in the request. The template's
	// reply-to applies whenever the request sets no ReplyTo and no reply-to
	// header, From or not.
	if template.From != nil {
		fmt.Printf("Sending from the template's sender: %s\n", template.From.Email)
	} else {
		message.From = &common.SenderAddress{
			Email: "sender@yourdomain.com",
			Name:  ahasend.String("Your Name"),
		}
	}

	// Send the email
	fmt.Println("Sending email from template...")
	response, httpResp, err := client.MessagesAPI.CreateTemplateMessage(ctx, accountID, message)

	if err != nil {
		// A missing required variable rejects the whole request, naming each
		// variable and the recipient it is missing for
		if apiErr, ok := err.(*api.APIError); ok {
			log.Fatalf("API Error: %s\nStatus Code: %d\nResponse Body: %s",
				apiErr.Error(), apiErr.StatusCode, string(apiErr.Raw))
		}
		log.Fatalf("Error sending email: %v", err)
	}

	// Check HTTP status
	if httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 {
		fmt.Printf("Email sent successfully!\n")
		if len(response.Data) > 0 {
			if response.Data[0].ID != nil {
				fmt.Printf("Message ID: %s\n", *response.Data[0].ID)
			}
			fmt.Printf("Status: %s\n", response.Data[0].Status)
			fmt.Printf("Recipient: %s\n", response.Data[0].Recipient.Email)
		}
	} else {
		fmt.Printf("Unexpected status code: %d\n", httpResp.StatusCode)
	}
}
