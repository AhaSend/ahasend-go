//go:build ignore

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
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

	// Create authentication context for parent-account operations
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, apiKey)

	const exampleSlug = "acme-sandbox"

	subAccountRequest := requests.CreateSubAccountRequest{
		Name:          "Acme Sandbox",
		Website:       "acme.example.com",
		MonthlyCredit: ahasend.Int64(10000),
	}

	subAccountIdempotencyKey := fmt.Sprintf("subacct-bootstrap-%s", exampleSlug)
	apiKeyIdempotencyKey := fmt.Sprintf("subacct-bootstrap-key-%s", exampleSlug)

	fmt.Println("=== Sub Account Bootstrap Example ===")
	fmt.Println("This example creates a child account, creates a bootstrap child API key,")
	fmt.Println("stores the one-time secret_key immediately, and uses it on a normal child route.")

	fmt.Println("\n1. Creating sub account...")
	childAccountID, err := createSubAccount(ctx, client, accountID, subAccountRequest, subAccountIdempotencyKey)
	if err != nil {
		handleError(err)
		os.Exit(1)
	}

	fmt.Println("\n2. Creating child API key...")
	childSecret, err := createChildAPIKey(ctx, client, accountID, childAccountID, apiKeyIdempotencyKey)
	if err != nil {
		handleError(err)
		os.Exit(1)
	}

	fmt.Println("\n3. Calling a normal child-account route with the child secret...")
	if err := listChildDomains(ctx, client, childAccountID, childSecret); err != nil {
		handleError(err)
		os.Exit(1)
	}

	fmt.Println("\nBootstrap complete.")
}

func createSubAccount(
	ctx context.Context,
	client *api.APIClient,
	accountID uuid.UUID,
	request requests.CreateSubAccountRequest,
	idempotencyKey string,
) (uuid.UUID, error) {
	fmt.Printf("Using stable sub-account idempotency key: %s\n", idempotencyKey)

	response, httpResp, err := client.SubAccountsAPI.CreateSubAccount(
		ctx,
		accountID,
		request,
		api.WithIdempotencyKey(idempotencyKey),
	)
	if err != nil {
		return uuid.Nil, err
	}
	if response == nil {
		return uuid.Nil, fmt.Errorf("create sub account returned an empty response (status %d)", responseStatus(httpResp))
	}

	fmt.Printf("Created sub account (status %d)\n", responseStatus(httpResp))
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Name: %s\n", response.Name)
	fmt.Printf("  Website: %s\n", response.Website)
	fmt.Printf("  Status: %s\n", response.Status)

	return response.ID, nil
}

func createChildAPIKey(
	ctx context.Context,
	client *api.APIClient,
	parentAccountID uuid.UUID,
	childAccountID uuid.UUID,
	idempotencyKey string,
) (string, error) {
	fmt.Printf("Using stable child API-key idempotency key: %s\n", idempotencyKey)

	response, httpResp, err := client.SubAccountsAPI.CreateSubAccountAPIKey(
		ctx,
		parentAccountID,
		childAccountID,
		requests.CreateAPIKeyRequest{
			Label: "Bootstrap key",
			Scopes: []string{
				"messages:send:all",
				"domains:read",
			},
		},
		api.WithIdempotencyKey(idempotencyKey),
	)
	if err != nil {
		return "", err
	}
	if response == nil {
		return "", fmt.Errorf("create child API key returned an empty response (status %d)", responseStatus(httpResp))
	}
	if response.SecretKey == nil || *response.SecretKey == "" {
		return "", fmt.Errorf("create child API key response did not include secret_key; API-key create secrets are returned only on create responses and exact idempotent replays within the replay window")
	}

	fmt.Printf("Created child API key (status %d)\n", responseStatus(httpResp))
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Label: %s\n", response.Label)
	fmt.Printf("  Public key: %s\n", response.PublicKey)
	fmt.Printf("\nStore this one-time child SecretKey securely now: %s\n", *response.SecretKey)

	return *response.SecretKey, nil
}

func listChildDomains(
	ctx context.Context,
	client *api.APIClient,
	childAccountID uuid.UUID,
	childSecret string,
) error {
	response, httpResp, err := client.DomainsAPI.GetDomains(
		ctx,
		childAccountID,
		nil,
		&common.PaginationParams{
			Limit: ahasend.Int32(10),
		},
		api.WithRequestAPIKey(childSecret),
	)
	if err != nil {
		return err
	}

	fmt.Printf("Child-authenticated domains request succeeded (status %d)\n", responseStatus(httpResp))
	if response == nil || len(response.Data) == 0 {
		fmt.Println("No child domains found yet.")
		return nil
	}

	fmt.Printf("Found %d child domains:\n", len(response.Data))
	for _, domain := range response.Data {
		fmt.Printf("  - %s (dns_valid=%t)\n", domain.Domain, domain.DNSValid)
	}

	return nil
}

func responseStatus(resp *http.Response) int {
	if resp == nil {
		return 0
	}
	return resp.StatusCode
}

func handleError(err error) {
	if apiErr, ok := err.(*api.APIError); ok {
		log.Printf("API Error: %s (Status: %d)", apiErr.Error(), apiErr.StatusCode)
		if len(apiErr.Raw) > 0 {
			log.Printf("Response Body: %s", string(apiErr.Raw))
		}
	} else {
		log.Printf("Error: %v", err)
	}
}
