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

	listID := uuid.MustParse("22222222-2222-4222-8222-222222222222")
	templateID := uuid.MustParse("33333333-3333-4333-8333-333333333333")

	// The campaign copies the campaign template's design; later changes to
	// the template do not reach it. Every field but Name can be set later.
	response, _, err := client.CampaignsAPI.CreateCampaign(ctx, accountID, requests.CreateCampaignRequest{
		Name:      "October newsletter",
		Subject:   ahasend.String("Our news for October, {{ first_name }}"),
		Preheader: ahasend.String("Three things we shipped this month"),
		From: &common.SenderAddress{
			Email: "news@yourdomain.com",
			Name:  ahasend.String("Example"),
		},
		ListIDs:    []uuid.UUID{listID},
		TemplateID: &templateID,
	})
	if err != nil {
		log.Fatalf("Error creating campaign: %v", err)
	}
	fmt.Printf("Created campaign %s, ready: %t\n", response.ID, response.Ready != nil && *response.Ready)
}
