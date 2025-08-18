//go:build ignore

package main

import (
	"context"
	"encoding/base64"
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

	// Create a new API client
	client := api.NewAPIClient(
		api.WithAPIKey(apiKey),
	)

	// Create authentication context
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, apiKey)

	// Example: Create a simple text attachment
	textContent := "This is the content of the text file attachment."
	textBase64 := base64.StdEncoding.EncodeToString([]byte(textContent))

	// Example: Read a file from disk (uncomment to use)
	// fileData, err := readFileAsBase64("path/to/your/file.pdf")
	// if err != nil {
	//     log.Fatalf("Error reading file: %v", err)
	// }

	// Create the email message with attachments
	message := requests.CreateMessageRequest{
		From: common.SenderAddress{
			Email: "sender@yourdomain.com",
			Name:  ahasend.String("Your Name"),
		},
		Recipients: []common.Recipient{
			{
				Email: "recipient@example.com",
				Name:  ahasend.String("Recipient Name"),
			},
		},
		Subject: "Email with Attachments",
		HtmlContent: ahasend.String(`
			<html>
				<body>
					<h1>Email with Attachments</h1>
					<p>This email includes file attachments.</p>
					<p>Please find the attached files:</p>
					<ul>
						<li>document.txt - A text document</li>
						<li>report.csv - A CSV report (example)</li>
					</ul>
					<p>Best regards,<br>The AhaSend Team</p>
				</body>
			</html>
		`),
		Attachments: []common.Attachment{
			{
				FileName:    "document.txt",
				ContentType: "text/plain",
				Data:        textBase64,
				Base64:      true,
			},
			{
				FileName:    "report.csv",
				ContentType: "text/csv",
				Data:        base64.StdEncoding.EncodeToString([]byte("Name,Email,Status\nJohn Doe,john@example.com,Active\nJane Smith,jane@example.com,Active")),
				Base64:      true,
			},
			// Uncomment to add a PDF or other binary file
			// {
			//     FileName:    "report.pdf",
			//     ContentType: "application/pdf",
			//     Data:        fileData,
			//     Base64:      true,
			// },
		},
	}

	// Send the email
	fmt.Println("Sending email with attachments...")
	response, httpResp, err := client.MessagesAPI.CreateMessage(ctx, accountID, message)

	if err != nil {
		// Check if it's an API error with details
		if apiErr, ok := err.(*api.APIError); ok {
			log.Fatalf("API Error: %s\nStatus Code: %d\nResponse Body: %s",
				apiErr.Error(), apiErr.StatusCode, string(apiErr.Raw))
		}
		log.Fatalf("Error sending email: %v", err)
	}

	// Check HTTP status
	if httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 {
		fmt.Printf("Email with attachments sent successfully!\n")
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

	// Show attachment size limits
	fmt.Println("\nNote: Maximum attachment size is typically 10MB per file.")
	fmt.Println("Total email size (including all attachments) should not exceed 25MB.")
}
