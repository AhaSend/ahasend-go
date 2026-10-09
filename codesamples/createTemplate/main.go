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
	client := api.NewAPIClient(api.WithAPIKey(os.Getenv("AHASEND_API_KEY")))
	accountID := uuid.MustParse(os.Getenv("AHASEND_ACCOUNT_ID"))
	ctx := context.Background()

	// HTML content makes an html template; MJML would make an advanced one.
	// The text is made from the HTML when it is left out. Publish makes the
	// template ready to send at once; without it, the fields stay in a draft.
	response, _, err := client.TemplatesAPI.CreateTemplate(ctx, accountID, requests.CreateTemplateRequest{
		Name:    "Password reset",
		Subject: ahasend.String("Reset your password, {{ first_name }}"),
		From: &common.SenderAddress{
			Email: "hello@yourdomain.com",
			Name:  ahasend.String("Example"),
		},
		Content: &requests.TemplateContentInput{
			HTML: ahasend.String(`<p>Hi {{ first_name }},</p><p><a href="{{ reset_url }}">Reset your password</a></p>`),
		},
		Publish: true,
	})
	if err != nil {
		log.Fatalf("Error creating template: %v", err)
	}
	fmt.Printf("Created %s template %s\n", response.Editor, response.ID)
}
