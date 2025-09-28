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

	// Create a new API client
	client := api.NewAPIClient(
		api.WithAPIKey(apiKey),
	)

	// Create authentication context
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, apiKey)

	// Example domain to add
	domainName := "mail.yourdomain.com"

	fmt.Println("=== Domain Management Example ===")

	// 1. List existing domains
	fmt.Println("1. Listing existing domains...")
	listDomains(ctx, client, accountID)

	// 2. Add a new domain
	fmt.Printf("\n2. Adding domain: %s\n", domainName)
	addDomain(ctx, client, accountID, domainName)

	// 3. Get domain details and verification records
	fmt.Printf("\n3. Getting domain details for: %s\n", domainName)
	getDomainDetails(ctx, client, accountID, domainName)

	// 4. Check domain verification status
	fmt.Printf("\n4. Checking verification status for: %s\n", domainName)
	checkDomainStatus(ctx, client, accountID, domainName)

	// 5. Show DNS record type explanations
	fmt.Println("\n5. DNS Record Types Explained")
	showDNSRecordTypes()

	// 6. Delete a domain (commented out to prevent accidental deletion)
	// fmt.Printf("\n6. Deleting domain: %s\n", domainName)
	// deleteDomain(ctx, client, accountID, domainName)
}

func listDomains(ctx context.Context, client *api.APIClient, accountID uuid.UUID) {
	response, httpResp, err := client.DomainsAPI.GetDomains(
		ctx,
		accountID,
		nil,
		&common.PaginationParams{
			Limit: ahasend.Int32(10),
		},
	)

	if err != nil {
		handleError(err)
		return
	}

	if httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 {
		if len(response.Data) > 0 {
			fmt.Printf("Found %d domains:\n", len(response.Data))
			for _, domain := range response.Data {
				status := "❌ Not Verified"
				if domain.DNSValid {
					status = "✅ Verified"
				}

				fmt.Printf("  - %s (%s)\n", domain.Domain, status)
				fmt.Printf("    ID: %s\n", domain.ID)
				fmt.Printf("    Created: %s\n", domain.CreatedAt.Format("2006-01-02 15:04:05"))
				fmt.Printf("    DNS Records: %d\n", len(domain.DNSRecords))

				if domain.LastDNSCheckAt != nil {
					fmt.Printf("    Last DNS Check: %s\n", domain.LastDNSCheckAt.Format("2006-01-02 15:04:05"))
				}
				fmt.Println()
			}
		} else {
			fmt.Println("No domains found.")
		}
	}
}

func addDomain(ctx context.Context, client *api.APIClient, accountID uuid.UUID, domainName string) {
	request := requests.CreateDomainRequest{
		Domain: domainName,
	}

	response, httpResp, err := client.DomainsAPI.CreateDomain(ctx, accountID, request)

	if err != nil {
		// Check if domain already exists
		if apiErr, ok := err.(*api.APIError); ok {
			if apiErr.StatusCode == 409 {
				fmt.Println("Domain already exists.")
				return
			}
		}
		handleError(err)
		return
	}

	if httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 {
		fmt.Println("Domain added successfully!")
		if response != nil {
			fmt.Printf("Domain ID: %s\n", response.ID)
			fmt.Printf("Domain: %s\n", response.Domain)
			fmt.Printf("DNS Valid: %t\n", response.DNSValid)

			if len(response.DNSRecords) > 0 {
				fmt.Printf("\nNext: Add %d DNS records to complete domain verification\n", len(response.DNSRecords))
				fmt.Println("Run the domain details command to see the required DNS records.")
			}
		}
	}
}

func getDomainDetails(ctx context.Context, client *api.APIClient, accountID uuid.UUID, domainName string) {
	response, httpResp, err := client.DomainsAPI.GetDomain(ctx, accountID, domainName)

	if err != nil {
		handleError(err)
		return
	}

	if httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 && response != nil {
		domain := response

		fmt.Println("Domain Details:")
		fmt.Printf("  Domain: %s\n", domain.Domain)
		fmt.Printf("  ID: %s\n", domain.ID)
		fmt.Printf("  Created: %s\n", domain.CreatedAt.Format("2006-01-02 15:04:05"))
		fmt.Printf("  Updated: %s\n", domain.UpdatedAt.Format("2006-01-02 15:04:05"))
		fmt.Printf("  DNS Valid: %t\n", domain.DNSValid)

		if domain.LastDNSCheckAt != nil {
			fmt.Printf("  Last DNS Check: %s\n", domain.LastDNSCheckAt.Format("2006-01-02 15:04:05"))
		}

		// Show DNS records required for verification
		if len(domain.DNSRecords) > 0 {
			fmt.Println("\n  DNS Records Required for Verification:")
			for i, record := range domain.DNSRecords {
				fmt.Printf("    %d. %s Record\n", i+1, record.Type)
				fmt.Printf("       Host: %s\n", record.Host)
				fmt.Printf("       Content: %s\n", record.Content)
				fmt.Printf("       Required: %t\n", record.Required)
				fmt.Printf("       Propagated: %t\n", record.Propagated)
				if record.Required && !record.Propagated {
					fmt.Printf("       ⚠️  Action needed: Add this DNS record to your domain\n")
				} else if record.Propagated {
					fmt.Printf("       ✅ Verified\n")
				}
				fmt.Println()
			}

			// Show setup instructions
			fmt.Println("  Setup Instructions:")
			fmt.Println("  1. Log in to your DNS provider (e.g., Cloudflare, Route 53, GoDaddy)")
			fmt.Println("  2. Add the DNS records shown above to your domain")
			fmt.Println("  3. Wait for DNS propagation (can take up to 48 hours)")
			fmt.Println("  4. Check back later - AhaSend will automatically verify the records")
		} else {
			fmt.Println("\n  No DNS records found for this domain.")
		}
	}
}

func checkDomainStatus(ctx context.Context, client *api.APIClient, accountID uuid.UUID, domainName string) {
	response, httpResp, err := client.DomainsAPI.GetDomain(ctx, accountID, domainName)

	if err != nil {
		handleError(err)
		return
	}

	if httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 && response != nil {
		domain := response

		fmt.Println("Domain Verification Status:")
		fmt.Printf("  Domain: %s\n", domain.Domain)
		fmt.Printf("  Overall Status: ")
		if domain.DNSValid {
			fmt.Printf("✅ VERIFIED\n")
		} else {
			fmt.Printf("⚠️  PENDING VERIFICATION\n")
		}

		if domain.LastDNSCheckAt != nil {
			fmt.Printf("  Last Checked: %s\n", domain.LastDNSCheckAt.Format("2006-01-02 15:04:05"))
		}

		// Show detailed verification status for each DNS record
		if len(domain.DNSRecords) > 0 {
			fmt.Println("\n  Detailed DNS Record Status:")
			requiredCount := 0
			verifiedCount := 0

			for _, record := range domain.DNSRecords {
				if record.Required {
					requiredCount++
					if record.Propagated {
						verifiedCount++
					}
				}

				status := "Optional"
				if record.Required {
					if record.Propagated {
						status = "✅ Verified"
					} else {
						status = "❌ Not Found"
					}
				}

				fmt.Printf("    %s (%s): %s\n", record.Type, record.Host, status)
			}

			fmt.Printf("\n  Verification Progress: %d/%d required records verified\n", verifiedCount, requiredCount)

			if !domain.DNSValid {
				fmt.Println("\n  📝 Next Steps:")
				fmt.Println("  1. Add any missing DNS records to your domain")
				fmt.Println("  2. Wait for DNS propagation (usually 5-30 minutes)")
				fmt.Println("  3. AhaSend will automatically re-check and verify your domain")
				fmt.Println("  4. Once verified, you can start sending emails from this domain")
			} else {
				fmt.Println("\n  🚀 Your domain is ready to send emails!")
			}
		}
	}
}

func deleteDomain(ctx context.Context, client *api.APIClient, accountID uuid.UUID, domainName string) {
	response, httpResp, err := client.DomainsAPI.DeleteDomain(ctx, accountID, domainName)

	if err != nil {
		handleError(err)
		return
	}

	if httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 {
		fmt.Println("Domain deleted successfully!")
		if response != nil {
			fmt.Printf("Message: %s\n", response.Message)
		}
	}
}

func showDNSRecordTypes() {
	fmt.Println("Understanding DNS Records for Email Authentication:")
	fmt.Println()

	fmt.Println("🗺️ DNS Record Types:")
	fmt.Println("  SPF (TXT Record)")
	fmt.Println("    • Purpose: Specifies which servers are authorized to send email for your domain")
	fmt.Println("    • Example: 'v=spf1 include:_spf.ahasend.com ~all'")
	fmt.Println("    • Prevents email spoofing and improves deliverability")
	fmt.Println()

	fmt.Println("  DKIM (TXT Record)")
	fmt.Println("    • Purpose: Adds a digital signature to your emails")
	fmt.Println("    • Host: Usually something like 'ahasend._domainkey'")
	fmt.Println("    • Allows receiving servers to verify email authenticity")
	fmt.Println()

	fmt.Println("  DMARC (TXT Record)")
	fmt.Println("    • Purpose: Tells receiving servers what to do with unauthenticated emails")
	fmt.Println("    • Host: '_dmarc'")
	fmt.Println("    • Example: 'v=DMARC1; p=quarantine; rua=mailto:dmarc@yourdomain.com'")
	fmt.Println()

	fmt.Println("  MX (Mail Exchange Record)")
	fmt.Println("    • Purpose: Directs incoming email to the correct mail servers")
	fmt.Println("    • Required for receiving emails (optional for sending-only domains)")
	fmt.Println("    • Example: Priority 10, Value: 'mail.ahasend.com'")
	fmt.Println()

	fmt.Println("📈 Why These Records Matter:")
	fmt.Println("  1. Deliverability: Properly configured DNS improves email delivery rates")
	fmt.Println("  2. Security: Prevents others from spoofing emails from your domain")
	fmt.Println("  3. Reputation: Builds trust with email providers like Gmail, Outlook")
	fmt.Println("  4. Compliance: Required by many email providers for bulk sending")
	fmt.Println()

	fmt.Println("🛠️ Setup Tips:")
	fmt.Println("  • DNS changes can take 5 minutes to 48 hours to propagate")
	fmt.Println("  • Use DNS checking tools to verify your records are correct")
	fmt.Println("  • Start with a DMARC policy of 'p=none' for monitoring")
	fmt.Println("  • Keep TTL values reasonable (300-3600 seconds)")
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
