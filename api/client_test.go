package api

import (
	"testing"

	"github.com/AhaSend/ahasend-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestClientCreation tests that we can create API client instances
func TestClientCreation(t *testing.T) {
	config := NewConfiguration()
	require.NotNil(t, config)

	client := NewAPIClientWithConfig(config)
	require.NotNil(t, client)

	// Test that all API services are initialized
	assert.NotNil(t, client.AccountsAPI)
	assert.NotNil(t, client.APIKeysAPI)
	assert.NotNil(t, client.DomainsAPI)
	assert.NotNil(t, client.MessagesAPI)
	assert.NotNil(t, client.RoutesAPI)
	assert.NotNil(t, client.SMTPCredentialsAPI)
	assert.NotNil(t, client.StatisticsAPI)
	assert.NotNil(t, client.SuppressionsAPI)
	assert.NotNil(t, client.UtilityAPI)
	assert.NotNil(t, client.WebhooksAPI)
}

// TestUtilityFunctions tests the pointer utility functions
func TestUtilityFunctions(t *testing.T) {
	// Test PtrString
	str := "test"
	ptrStr := ahasend.String(str)
	assert.NotNil(t, ptrStr)
	assert.Equal(t, str, *ptrStr)

	// Test PtrBool
	b := true
	ptrBool := ahasend.Bool(b)
	assert.NotNil(t, ptrBool)
	assert.Equal(t, b, *ptrBool)

	// Test PtrInt
	i := 42
	ptrInt := ahasend.Int(i)
	assert.NotNil(t, ptrInt)
	assert.Equal(t, i, *ptrInt)
}
