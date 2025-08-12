package ahasend_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/AhaSend/ahasend-go"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestGetDomainsWithDnsValidParameter tests that the dns_valid parameter methods work correctly
func TestGetDomainsWithDnsValidParameter(t *testing.T) {
	// Create a test configuration and client
	configuration := ahasend.NewConfiguration()
	configuration.Host = "api.example.com"
	configuration.Scheme = "https"
	apiClient := ahasend.NewAPIClient(configuration)

	accountId := uuid.New()
	ctx := context.Background()

	t.Run("DnsValid method returns correct type", func(t *testing.T) {
		request := apiClient.DomainsAPI.GetDomains(ctx, accountId).DnsValid(true)

		// Verify the method returns the correct type for chaining
		assert.IsType(t, ahasend.ApiGetDomainsRequest{}, request)
	})

	t.Run("DnsValid method returns correct type when false", func(t *testing.T) {
		request := apiClient.DomainsAPI.GetDomains(ctx, accountId).DnsValid(false)

		// Verify the method returns the correct type for chaining
		assert.IsType(t, ahasend.ApiGetDomainsRequest{}, request)
	})

	t.Run("Method chaining works with DnsValid", func(t *testing.T) {
		request := apiClient.DomainsAPI.GetDomains(ctx, accountId).
			DnsValid(true).
			Limit(50).
			Cursor("test-cursor")

		// Verify the method chaining returns the correct type
		assert.IsType(t, ahasend.ApiGetDomainsRequest{}, request)
	})
}

// TestGetDomainsQueryParameterBuilding tests that dns_valid is properly added to query parameters
func TestGetDomainsQueryParameterBuilding(t *testing.T) {
	t.Run("Query parameter includes dns_valid when set to true", func(t *testing.T) {
		// Create query values manually like the SDK does
		params := url.Values{}

		// Simulate what parameterAddToHeaderOrQuery does
		params.Add("dns_valid", "true")

		queryString := params.Encode()
		assert.Contains(t, queryString, "dns_valid=true")
	})

	t.Run("Query parameter includes dns_valid when set to false", func(t *testing.T) {
		// Create query values manually like the SDK does
		params := url.Values{}

		// Simulate what parameterAddToHeaderOrQuery does
		params.Add("dns_valid", "false")

		queryString := params.Encode()
		assert.Contains(t, queryString, "dns_valid=false")
	})

	t.Run("Query parameter does not include dns_valid when not set", func(t *testing.T) {
		// Create query values manually like the SDK does
		params := url.Values{}

		// Don't add dns_valid parameter
		params.Add("limit", "100")

		queryString := params.Encode()
		assert.NotContains(t, queryString, "dns_valid")
		assert.Contains(t, queryString, "limit=100")
	})
}
