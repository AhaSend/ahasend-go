package main

import (
	"context"
	"errors"
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

	templateID := uuid.MustParse("11111111-1111-4111-8111-111111111111")

	// Only the fields you set change. They go to the template's draft, the
	// same draft the dashboard edits; Publish publishes the whole draft. The
	// design of a simple template is written as AhaSend Markup (AhaML).
	markup := `<Email version="1">
  <Heading level="2">Reset your password</Heading>
  <Paragraph>Hi <Var name="first_name" fallback="there"/>, use the button below to reset your password.</Paragraph>
  <Button url="{{ reset_url }}">Reset password</Button>
</Email>
`
	response, _, err := client.TemplatesAPI.UpdateTemplate(ctx, accountID, templateID, requests.UpdateTemplateRequest{
		Subject: ahasend.String("Your password reset link"),
		Content: &requests.TemplateContentInput{Markup: &markup},
		Publish: true,
	})
	if err != nil {
		var apiErr *api.APIError
		if errors.As(err, &apiErr) {
			// A refused design lists each problem, with its place in the
			// AhaML when it has one.
			for _, problem := range apiErr.Errors {
				if problem.Line > 0 {
					fmt.Printf("line %d, column %d: %s\n", problem.Line, problem.Column, problem.Message)
				} else {
					fmt.Println(problem.Message)
				}
			}
		}
		log.Fatalf("Error updating template: %v", err)
	}
	for _, warning := range response.MarkupWarnings {
		fmt.Printf("Warning: %s\n", warning.Message)
	}
	fmt.Printf("Updated template %s, has draft: %t\n", response.Name, response.HasDraft)
}
