//go:build ignore

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/api"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/AhaSend/ahasend-go/models/responses"
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

	// Create a new API client. The key needs contacts:write plus lists:read,
	// lists:write, and lists:delete.
	client := api.NewAPIClient(
		api.WithAPIKey(apiKey),
	)
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, apiKey)

	fmt.Println("=== List Management Example ===")

	// Lists hold existing contacts, so make sure the contacts exist first
	fmt.Println("\n1. Upserting contacts...")
	emails := []string{"alice@example.com", "bob@example.com"}
	contactInputs := make([]requests.BatchUpsertContactInput, len(emails))
	for i, email := range emails {
		contactInputs[i] = requests.BatchUpsertContactInput{Email: email}
	}
	contacts, _, err := client.ContactsAPI.BatchUpsertContacts(ctx, accountID, requests.BatchUpsertContactsRequest{
		Data: contactInputs,
	}, api.WithIdempotencyKey("list-example-contacts"))
	if err != nil {
		log.Fatalf("Failed to upsert contacts: %v", err)
	}
	fmt.Printf("Created %d, updated %d, failed %d\n", contacts.Created, contacts.Updated, contacts.Failed)

	fmt.Println("\n2. Creating a list...")
	list, _, err := client.ListsAPI.CreateList(ctx, accountID, requests.CreateListRequest{
		Name:        "Product Updates",
		Description: ahasend.String("Monthly product news"),
		Tags:        []string{"newsletter"},
	}, api.WithIdempotencyKey("list-example-create"))
	if err != nil {
		log.Fatalf("Failed to create list: %v", err)
	}
	fmt.Printf("Created list %s (%s)\n", list.Name, list.ID)

	// A batch add returns 200 even when some entries fail, so check every outcome
	fmt.Println("\n3. Adding contacts to the list...")
	listInputs := make([]requests.BatchAddListContactInput, len(emails))
	for i := range emails {
		listInputs[i] = requests.BatchAddListContactInput{Email: ahasend.String(emails[i])}
	}
	batch, _, err := client.ListsAPI.BatchAddListContacts(ctx, accountID, list.ID, requests.BatchAddListContactsRequest{
		Data: listInputs,
	}, api.WithIdempotencyKey("list-example-batch-add"))
	if err != nil {
		log.Fatalf("Failed to add contacts to list: %v", err)
	}
	fmt.Printf("Added %d, skipped %d, failed %d\n", batch.Added, batch.Skipped, batch.Failed)
	for _, result := range batch.Data {
		if result.Outcome == responses.BatchListContactOutcomeNotFound || result.Outcome == responses.BatchListContactOutcomeInvalid {
			fmt.Printf("  entry %d: %s (%s)\n", result.Position, result.Outcome, result.Reason)
		}
	}

	// Unsubscribing from one list keeps the record and can be reversed later
	fmt.Println("\n4. Unsubscribing bob from the list...")
	membership, _, err := client.ListsAPI.UpsertListContact(ctx, accountID, list.ID, "bob@example.com", requests.UpsertListContactRequest{
		SubscriptionStatus: ahasend.String(requests.ListContactStatusUnsubscribed),
	})
	if err != nil {
		log.Fatalf("Failed to unsubscribe contact: %v", err)
	}
	fmt.Printf("%s is now %s\n", membership.Email, membership.SubscriptionStatus)

	fmt.Println("\n5. Listing confirmed members...")
	members, _, err := client.ListsAPI.GetListContacts(ctx, accountID, list.ID, requests.GetListContactsParams{
		SubscriptionStatus: ahasend.String(requests.ListContactStatusConfirmed),
	})
	if err != nil {
		log.Fatalf("Failed to list members: %v", err)
	}
	for _, member := range members.Data {
		fmt.Printf("  %s\n", member.Email)
	}

	fmt.Println("\n6. Listing alice's lists...")
	memberships, _, err := client.ListsAPI.GetContactLists(ctx, accountID, "alice@example.com", requests.GetContactListsParams{})
	if err != nil {
		log.Fatalf("Failed to list contact's lists: %v", err)
	}
	for _, m := range memberships.Data {
		fmt.Printf("  %s: %s\n", m.List.Name, m.SubscriptionStatus)
	}

	// contact_count is the number of members a campaign to the list would reach
	fmt.Println("\n7. Reading the list's reach...")
	list, _, err = client.ListsAPI.GetList(ctx, accountID, list.ID)
	if err != nil {
		log.Fatalf("Failed to get list: %v", err)
	}
	fmt.Printf("%s reaches %d contacts\n", list.Name, list.ContactCount)

	fmt.Println("\n8. Clearing the description and tags...")
	list, _, err = client.ListsAPI.UpdateList(ctx, accountID, list.ID, requests.UpdateListRequest{
		Description: ahasend.String(""),
		Tags:        &[]string{},
	})
	if err != nil {
		log.Fatalf("Failed to update list: %v", err)
	}
	fmt.Printf("Description %q, tags %v\n", list.Description, list.Tags)

	// Deleting a list keeps its memberships, so unsubscribes survive it
	fmt.Println("\n9. Deleting the list...")
	if _, _, err := client.ListsAPI.DeleteList(ctx, accountID, list.ID); err != nil {
		log.Fatalf("Failed to delete list: %v", err)
	}
	fmt.Println("List deleted")
}
