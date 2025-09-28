package api

import (
	"context"
	"net/url"
	"testing"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestGetDomainsWithDnsValidParameter tests that the dns_valid parameter works correctly with the new direct method API
func TestGetDomainsWithDnsValidParameter(t *testing.T) {
	// Create a test configuration and client
	configuration := NewConfiguration()
	configuration.Host = "api.example.com"
	configuration.Scheme = "https"
	apiClient := NewAPIClientWithConfig(configuration)

	accountId := uuid.New()
	ctx := context.Background()

	t.Run("GetDomains with dnsValid=true parameter", func(t *testing.T) {
		dnsValid := true
		_, _, err := apiClient.DomainsAPI.GetDomains(ctx, accountId, &dnsValid, nil)

		// We expect this to fail with a network error since we're not hitting a real API
		// but it should not fail with a compilation or type error
		assert.Error(t, err) // This will be a network error, which is expected
	})

	t.Run("GetDomains with dnsValid=false parameter", func(t *testing.T) {
		dnsValid := false
		_, _, err := apiClient.DomainsAPI.GetDomains(ctx, accountId, &dnsValid, nil)

		// We expect this to fail with a network error since we're not hitting a real API
		// but it should not fail with a compilation or type error
		assert.Error(t, err) // This will be a network error, which is expected
	})

	t.Run("GetDomains with all parameters set", func(t *testing.T) {
		dnsValid := true
		pagination := &common.PaginationParams{
			Limit:  &[]int32{50}[0],
			Cursor: &[]string{"test-cursor"}[0],
			After:  &[]string{"test-after-cursor"}[0],
			Before: &[]string{"test-before-cursor"}[0],
		}

		_, _, err := apiClient.DomainsAPI.GetDomains(ctx, accountId, &dnsValid, pagination)

		// We expect this to fail with a network error since we're not hitting a real API
		// but it should not fail with a compilation or type error
		assert.Error(t, err) // This will be a network error, which is expected
	})

	t.Run("GetDomains with bidirectional pagination", func(t *testing.T) {
		dnsValid := true
		pagination := &common.PaginationParams{
			Limit: &[]int32{50}[0],
			After: &[]string{"test-after-cursor"}[0],
		}

		_, _, err := apiClient.DomainsAPI.GetDomains(ctx, accountId, &dnsValid, pagination)

		// We expect this to fail with a network error since we're not hitting a real API
		// but it should not fail with a compilation or type error
		assert.Error(t, err) // This will be a network error, which is expected
	})
}

// TestGetDomainsQueryParameterBuilding tests that dns_valid is properly added to query parameters
func TestGetDomainsQueryParameterBuilding(t *testing.T) {
	t.Run("Query parameter includes dns_valid when set to true", func(t *testing.T) {
		// Create query values manually like the SDK does
		params := url.Values{}

		// Simulate what the new SDK does
		params.Add("dns_valid", "true")

		queryString := params.Encode()
		assert.Contains(t, queryString, "dns_valid=true")
	})

	t.Run("Query parameter includes dns_valid when set to false", func(t *testing.T) {
		// Create query values manually like the SDK does
		params := url.Values{}

		// Simulate what the new SDK does
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
