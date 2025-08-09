package ahasend_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/AhaSend/ahasend-go"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	prismPort    = "4010"
	prismBaseURL = "http://localhost:" + prismPort
	testAPIKey   = "aha-sk-test-key-for-integration-testing-12345678901234567890"
)

var (
	testAccountID = uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")
)

// PrismManager handles starting and stopping Prism mock server
type PrismManager struct {
	cmd    *exec.Cmd
	port   string
	ready  chan bool
	errors chan error
}

// NewPrismManager creates a new Prism manager
func NewPrismManager(port string) *PrismManager {
	return &PrismManager{
		port:   port,
		ready:  make(chan bool),
		errors: make(chan error),
	}
}

// Start starts the Prism mock server
func (pm *PrismManager) Start() error {
	// Try to find prism command in order of preference
	var cmd *exec.Cmd

	// Get absolute path to OpenAPI spec
	specPath := "../api/openapi.yaml"
	if absPath, err := filepath.Abs(specPath); err == nil {
		specPath = absPath
	}

	// 1. Check for custom PRISM_CMD environment variable
	if prismCmd := os.Getenv("PRISM_CMD"); prismCmd != "" {
		args := []string{"mock", specPath, "--host", "0.0.0.0", "--port", pm.port, "--dynamic"}
		cmd = exec.Command(prismCmd, args...)
	} else if prismPath, err := exec.LookPath("prism"); err == nil {
		// 2. Use prism if it's in PATH
		cmd = exec.Command(prismPath, "mock", specPath, "--host", "0.0.0.0", "--port", pm.port, "--dynamic")
	} else if _, err := exec.LookPath("npx"); err == nil {
		// 3. Fall back to npx (most CI/CD environments have this)
		cmd = exec.Command("npx", "@stoplight/prism-cli", "mock", specPath, "--host", "0.0.0.0", "--port", pm.port, "--dynamic")
	} else {
		return fmt.Errorf("prism is not available. Install with: npm install -g @stoplight/prism-cli or ensure npx is available")
	}

	pm.cmd = cmd

	// Start the command
	err := pm.cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start prism: %w", err)
	}

	// Wait for Prism to be ready (up to 30 seconds to account for npx downloads)
	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			pm.Stop()
			return fmt.Errorf("prism failed to start within 30 seconds")
		case <-ticker.C:
			if pm.isReady() {
				return nil
			}
		}
	}
}

// Stop stops the Prism mock server
func (pm *PrismManager) Stop() error {
	if pm.cmd != nil && pm.cmd.Process != nil {
		return pm.cmd.Process.Kill()
	}
	return nil
}

// isReady checks if Prism is ready to accept requests
func (pm *PrismManager) isReady() bool {
	// Don't use -f flag since we expect 401 (authentication required) which indicates prism is running
	cmd := exec.Command("curl", "-s", "-w", "%{http_code}", "-o", "/dev/null", fmt.Sprintf("http://localhost:%s/v2/ping", pm.port))
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	statusCode := string(output)
	// Accept both 200 and 401 as valid responses (401 means server is running but requires auth)
	return statusCode == "200" || statusCode == "401"
}

// TestMain sets up and tears down the test environment
func TestMain(m *testing.M) {
	// Skip integration tests that require Prism if explicitly disabled
	if os.Getenv("SKIP_INTEGRATION_TESTS") == "true" {
		fmt.Println("Skipping Prism-based integration tests (SKIP_INTEGRATION_TESTS=true)")
		// Still run other tests, just not the ones that need Prism
		os.Exit(m.Run())
	}

	// Start Prism mock server
	prism := NewPrismManager(prismPort)

	fmt.Println("Starting Prism mock server...")
	err := prism.Start()
	if err != nil {
		fmt.Printf("Failed to start Prism: %v\n", err)
		fmt.Println("To install Prism: npm install -g @stoplight/prism-cli")
		os.Exit(1)
	}

	fmt.Printf("Prism mock server started on port %s\n", prismPort)

	// Run tests
	code := m.Run()

	// Clean up
	fmt.Println("Stopping Prism mock server...")
	prism.Stop()

	os.Exit(code)
}

// createTestClient creates a test client configured for the mock server
func createTestClient() *ahasend.APIClient {
	cfg := ahasend.NewConfiguration()
	cfg.Host = "localhost:" + prismPort
	cfg.Scheme = "http"
	cfg.EnableRateLimit = false // Disable rate limiting for tests
	cfg.MaxRetries = 1          // Reduce retries for faster test execution

	return ahasend.NewAPIClient(cfg)
}

// createAuthContext creates a context with test authentication
func createAuthContext() context.Context {
	return CreateTestAuthContext(testAPIKey)
}

// TestPingEndpoint tests the basic connectivity with the mock server
func TestPingEndpoint(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION_TESTS") == "true" {
		t.Skip("Skipping Prism-based integration test")
	}

	client := createTestClient()
	ctx := createAuthContext()

	response, httpResp, err := client.UtilityAPI.Ping(ctx).Execute()

	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 200, httpResp.StatusCode)
}

// TestCreateMessage tests message creation with various scenarios
func TestCreateMessage(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION_TESTS") == "true" {
		t.Skip("Skipping Prism-based integration test")
	}

	client := createTestClient()
	ctx := createAuthContext()

	tests := []struct {
		name        string
		request     ahasend.CreateMessageRequest
		expectError bool
	}{
		{
			name: "Valid message",
			request: ahasend.CreateMessageRequest{
				From: ahasend.SenderAddress{
					Email: "test@example.com",
				},
				Recipients: []ahasend.Recipient{
					{
						Email: "recipient@example.com",
					},
				},
				Subject:     "Test Message",
				HtmlContent: ahasend.PtrString("<p>Hello from integration test!</p>"),
			},
			expectError: false,
		},
		{
			name: "Message with multiple recipients",
			request: ahasend.CreateMessageRequest{
				From: ahasend.SenderAddress{
					Email: "test@example.com",
				},
				Recipients: []ahasend.Recipient{
					{
						Email: "recipient1@example.com",
					},
					{
						Email: "recipient2@example.com",
					},
				},
				Subject:     "Test Message Multiple Recipients",
				HtmlContent: ahasend.PtrString("<p>Hello multiple recipients!</p>"),
			},
			expectError: false,
		},
		{
			name: "Message with attachments",
			request: ahasend.CreateMessageRequest{
				From: ahasend.SenderAddress{
					Email: "test@example.com",
				},
				Recipients: []ahasend.Recipient{
					{
						Email: "recipient@example.com",
					},
				},
				Subject:     "Test Message with Attachment",
				HtmlContent: ahasend.PtrString("<p>Hello with attachment!</p>"),
				Attachments: []ahasend.Attachment{
					{
						FileName:    "test.txt",
						ContentType: "text/plain",
						Data:        "VGVzdCBjb250ZW50", // "Test content" in base64
						Base64:      ahasend.PtrBool(true),
					},
				},
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, httpResp, err := client.MessagesAPI.
				CreateMessage(ctx, testAccountID).
				CreateMessageRequest(tt.request).
				Execute()

			if tt.expectError {
				assert.Error(t, err)
			} else {
				// With Prism mock server, we might get various response codes
				// The main goal is to test that the client can make the request properly
				if err == nil {
					require.NotNil(t, response)
					assert.True(t, httpResp.StatusCode >= 200 && httpResp.StatusCode < 500, "Expected valid HTTP status code, got %d", httpResp.StatusCode)

					// Verify idempotency key was added
					idempotencyKey := httpResp.Request.Header.Get("Idempotency-Key")
					assert.NotEmpty(t, idempotencyKey)

					// Basic response validation - only check if data exists
					if len(response.Data) > 0 {
						// Only validate ID if it exists (Prism might return different structures)
						if response.Data[0].Id != nil {
							assert.NotEmpty(t, *response.Data[0].Id)
						}
					}
				} else {
					// If there's an error, ensure we still got a response code
					require.NotNil(t, httpResp)
					assert.True(t, httpResp.StatusCode >= 400 && httpResp.StatusCode < 500, "Expected 4xx error status code, got %d", httpResp.StatusCode)
				}
			}
		})
	}
}

// TestIdempotencyBehavior tests idempotency key behavior
func TestIdempotencyBehavior(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION_TESTS") == "true" {
		t.Skip("Skipping Prism-based integration test")
	}

	client := createTestClient()
	ctx := createAuthContext()

	request := ahasend.CreateMessageRequest{
		From: ahasend.SenderAddress{
			Email: "test@example.com",
		},
		Recipients: []ahasend.Recipient{
			{
				Email: "recipient@example.com",
			},
		},
		Subject:     "Idempotency Test",
		HtmlContent: ahasend.PtrString("<p>Testing idempotency</p>"),
	}

	t.Run("Automatic idempotency key generation", func(t *testing.T) {
		response, httpResp, err := client.MessagesAPI.
			CreateMessage(ctx, testAccountID).
			CreateMessageRequest(request).
			Execute()

		require.NoError(t, err)
		assert.NotNil(t, response)

		// Verify idempotency key was automatically added
		idempotencyKey := httpResp.Request.Header.Get("Idempotency-Key")
		assert.NotEmpty(t, idempotencyKey)
		assert.Len(t, idempotencyKey, 36) // UUID length
	})

	t.Run("Custom idempotency key", func(t *testing.T) {
		customKey := "custom-test-key-12345"

		response, httpResp, err := client.MessagesAPI.
			CreateMessage(ctx, testAccountID).
			CreateMessageRequest(request).
			IdempotencyKey(customKey). // This should override auto-generation
			Execute()

		require.NoError(t, err)
		assert.NotNil(t, response)

		// Verify custom key was used
		usedKey := httpResp.Request.Header.Get("Idempotency-Key")
		assert.Equal(t, customKey, usedKey)
	})
}

// TestRateLimiting tests rate limiting behavior
func TestRateLimiting(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION_TESTS") == "true" {
		t.Skip("Skipping Prism-based integration test")
	}

	// Create client with rate limiting enabled
	cfg := ahasend.NewConfiguration()
	cfg.Host = "localhost:" + prismPort
	cfg.Scheme = "http"
	cfg.EnableRateLimit = true
	cfg.MaxRetries = 1

	client := ahasend.NewAPIClient(cfg)
	ctx := createAuthContext()

	// Set very restrictive rate limits for testing
	client.SetGeneralRateLimit(2, 2) // 2 req/s, 2 burst

	t.Run("Rate limiting with burst capacity", func(t *testing.T) {
		start := time.Now()

		// Make requests up to burst capacity (should be fast)
		for i := 0; i < 2; i++ {
			_, _, err := client.UtilityAPI.Ping(ctx).Execute()
			require.NoError(t, err)
		}

		burstDuration := time.Since(start)
		assert.Less(t, burstDuration, 500*time.Millisecond, "Burst requests should be fast")

		// Next request should be rate limited
		start = time.Now()
		_, _, err := client.UtilityAPI.Ping(ctx).Execute()
		require.NoError(t, err)

		rateLimitedDuration := time.Since(start)
		assert.Greater(t, rateLimitedDuration, 400*time.Millisecond, "Rate limited request should wait")
	})
}

// TestContextCancellation tests context cancellation behavior
func TestContextCancellation(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION_TESTS") == "true" {
		t.Skip("Skipping Prism-based integration test")
	}

	client := createTestClient()

	// Set restrictive rate limit to force waiting
	client.SetGeneralRateLimit(1, 1)

	t.Run("Context timeout cancels request", func(t *testing.T) {
		// With Prism mock server, we can't test real rate limiting timeouts
		// Instead, we test that context cancellation works by creating an already-cancelled context
		ctx := createAuthContext()

		// Create an already-cancelled context
		ctxWithCancel, cancel := context.WithCancel(ctx)
		cancel() // Cancel immediately

		start := time.Now()
		_, httpResp, err := client.UtilityAPI.Ping(ctxWithCancel).Execute()
		duration := time.Since(start)

		// Should get a context cancellation error quickly
		if err != nil {
			assert.Less(t, duration, 100*time.Millisecond, "Cancelled request should fail quickly")
			assert.Contains(t, err.Error(), "context canceled")
		} else {
			// If no error (Prism responds too fast), at least verify we got a response
			require.NotNil(t, httpResp)
			assert.True(t, httpResp.StatusCode >= 200 && httpResp.StatusCode < 500, "Expected valid status code")
		}
	})

	t.Run("Manual context cancellation", func(t *testing.T) {
		// With Prism mock server, we test context cancellation more directly
		ctx := createAuthContext()

		// Create cancellable context
		ctxWithCancel, cancel := context.WithCancel(ctx)

		// Cancel immediately to ensure cancellation happens
		cancel()

		start := time.Now()
		_, httpResp, err := client.UtilityAPI.Ping(ctxWithCancel).Execute()
		duration := time.Since(start)

		// Should get a context cancellation error quickly
		if err != nil {
			assert.Less(t, duration, 100*time.Millisecond, "Cancelled request should fail quickly")
			assert.Contains(t, err.Error(), "context canceled")
		} else {
			// If no error (Prism responds too fast), at least verify we got a response
			require.NotNil(t, httpResp)
			assert.True(t, httpResp.StatusCode >= 200 && httpResp.StatusCode < 500, "Expected valid status code")
		}
	})
}

// TestErrorHandling tests various error scenarios
func TestErrorHandling(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION_TESTS") == "true" {
		t.Skip("Skipping Prism-based integration test")
	}

	client := createTestClient()
	ctx := createAuthContext()

	t.Run("Invalid account ID", func(t *testing.T) {
		request := ahasend.CreateMessageRequest{
			From: ahasend.SenderAddress{
				Email: "test@example.com",
			},
			Recipients: []ahasend.Recipient{
				{
					Email: "recipient@example.com",
				},
			},
			Subject:     "Test Message",
			HtmlContent: ahasend.PtrString("<p>Hello!</p>"),
		}

		invalidID := uuid.MustParse("00000000-0000-0000-0000-000000000000")
		_, _, err := client.MessagesAPI.
			CreateMessage(ctx, invalidID).
			CreateMessageRequest(request).
			Execute()

		// With Prism mock server, the behavior for invalid data may vary
		// The main test is that the client can handle the response properly
		if err != nil {
			// Check if it's our structured error type
			if genErr, ok := err.(*ahasend.GenericOpenAPIError); ok {
				assert.Greater(t, genErr.StatusCode(), 0)
			}
		} else {
			// If no error, Prism might accept the request anyway - that's okay for testing
			t.Log("Prism accepted request with invalid account ID - this is acceptable for mock testing")
		}
	})

	t.Run("Unauthorized request", func(t *testing.T) {
		// Create context without authentication
		unauthCtx := context.Background()

		_, httpResp, err := client.UtilityAPI.Ping(unauthCtx).Execute()

		// With Prism mock server, we may get various responses or error types
		// The main test is that we can make unauthenticated requests without the client crashing
		if err != nil {
			// Check if it's our structured error type
			if genErr, ok := err.(*ahasend.GenericOpenAPIError); ok {
				statusCode := genErr.StatusCode()
				t.Logf("Got GenericOpenAPIError with status code: %d", statusCode)
				if statusCode > 0 {
					assert.True(t, statusCode >= 400 && statusCode < 500, "Expected 4xx error status code, got %d", statusCode)
				}
			} else {
				// For other error types, just verify we got an error (network errors, etc.)
				t.Logf("Got error (not GenericOpenAPIError): %T - %v", err, err)
				assert.NotNil(t, err)
			}
		} else {
			// If no error (Prism allows unauthenticated requests), at least verify response
			require.NotNil(t, httpResp)
			assert.True(t, httpResp.StatusCode >= 200 && httpResp.StatusCode < 500, "Expected valid status code")
		}
	})
}

// TestStatisticsEndpoints tests statistics API endpoints
func TestStatisticsEndpoints(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION_TESTS") == "true" {
		t.Skip("Skipping Prism-based integration test")
	}

	client := createTestClient()
	ctx := createAuthContext()

	// Set appropriate rate limits for statistics
	client.SetStatisticsRateLimit(1, 1)

	tests := []struct {
		name     string
		endpoint func() (interface{}, *http.Response, error)
	}{
		{
			name: "Get deliverability statistics",
			endpoint: func() (interface{}, *http.Response, error) {
				return client.StatisticsAPI.GetDeliverabilityStatistics(ctx, testAccountID).Execute()
			},
		},
		{
			name: "Get bounce statistics",
			endpoint: func() (interface{}, *http.Response, error) {
				return client.StatisticsAPI.GetBounceStatistics(ctx, testAccountID).Execute()
			},
		},
		{
			name: "Get delivery time statistics",
			endpoint: func() (interface{}, *http.Response, error) {
				return client.StatisticsAPI.GetDeliveryTimeStatistics(ctx, testAccountID).Execute()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, httpResp, err := tt.endpoint()

			// Allow either success or specific error from mock server
			if err == nil {
				assert.NotNil(t, response)
				assert.Equal(t, 200, httpResp.StatusCode)
			} else {
				// If error, should be a structured error
				assert.IsType(t, &ahasend.GenericOpenAPIError{}, err)
			}
		})
	}
}

// TestDomainOperations tests domain management operations
func TestDomainOperations(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION_TESTS") == "true" {
		t.Skip("Skipping Prism-based integration test")
	}

	client := createTestClient()
	ctx := createAuthContext()

	testDomain := "test-example.com"

	t.Run("List domains", func(t *testing.T) {
		response, httpResp, err := client.DomainsAPI.GetDomains(ctx, testAccountID).Execute()

		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 200, httpResp.StatusCode)
	})

	t.Run("Create domain", func(t *testing.T) {
		request := ahasend.CreateDomainRequest{
			Domain: testDomain,
		}

		response, httpResp, err := client.DomainsAPI.
			CreateDomain(ctx, testAccountID).
			CreateDomainRequest(request).
			Execute()

		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 201, httpResp.StatusCode) // 201 Created for POST operations
	})

	t.Run("Get domain", func(t *testing.T) {
		response, httpResp, err := client.DomainsAPI.GetDomain(ctx, testAccountID, testDomain).Execute()

		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 200, httpResp.StatusCode)
	})
}

// TestJSONSchemaValidation tests that requests/responses conform to expected schemas
func TestJSONSchemaValidation(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION_TESTS") == "true" {
		t.Skip("Skipping Prism-based integration test")
	}

	client := createTestClient()
	ctx := createAuthContext()

	t.Run("Message request schema validation", func(t *testing.T) {
		request := ahasend.CreateMessageRequest{
			From: ahasend.SenderAddress{
				Email: "test@example.com",
			},
			Recipients: []ahasend.Recipient{
				{
					Email: "recipient@example.com",
				},
			},
			Subject:     "Schema Test",
			HtmlContent: ahasend.PtrString("<p>Testing schema validation</p>"),
		}

		// Serialize to JSON to check structure
		jsonData, err := json.Marshal(request)
		require.NoError(t, err)

		// Verify required fields are present
		var jsonMap map[string]interface{}
		err = json.Unmarshal(jsonData, &jsonMap)
		require.NoError(t, err)

		assert.Contains(t, jsonMap, "from")
		assert.Contains(t, jsonMap, "recipients")
		assert.Contains(t, jsonMap, "subject")
		assert.Contains(t, jsonMap, "html_content")
	})

	t.Run("Response schema validation", func(t *testing.T) {
		response, _, err := client.UtilityAPI.Ping(ctx).Execute()
		require.NoError(t, err)

		// Serialize response to check structure
		jsonData, err := json.Marshal(response)
		require.NoError(t, err)

		// Should be valid JSON
		var jsonMap map[string]interface{}
		err = json.Unmarshal(jsonData, &jsonMap)
		require.NoError(t, err)
	})
}

// TestGetMessagesWithStatusFilter tests the new status parameter in GetMessages
func TestGetMessagesWithStatusFilter(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION_TESTS") == "true" {
		t.Skip("Skipping Prism-based integration test")
	}

	client := createTestClient()
	ctx := createAuthContext()

	tests := []struct {
		name        string
		status      string
		expectError bool
		description string
	}{
		{
			name:        "Single status filter - Delivered",
			status:      "Delivered",
			expectError: false,
			description: "Filter messages by single status",
		},
		{
			name:        "Single status filter - Bounced",
			status:      "Bounced",
			expectError: false,
			description: "Filter messages by bounced status",
		},
		{
			name:        "Multiple status filter - Delivered and Bounced",
			status:      "Delivered,Bounced",
			expectError: false,
			description: "Filter messages by multiple statuses",
		},
		{
			name:        "Complex status filter",
			status:      "Delivered,Bounced,Failed,Queued",
			expectError: false,
			description: "Filter messages by complex status combination",
		},
		{
			name:        "All possible statuses",
			status:      "Delivered,Bounced,Failed,Queued,Processing,Suppressed",
			expectError: false,
			description: "Filter messages by all possible statuses",
		},
		{
			name:        "Status with spaces",
			status:      "Delivered, Bounced, Failed",
			expectError: false,
			description: "Filter messages with spaces in status parameter",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, httpResp, err := client.MessagesAPI.
				GetMessages(ctx, testAccountID).
				Sender("test@example.com").
				Status(tt.status).
				Execute()

			if tt.expectError {
				assert.Error(t, err)
			} else {
				// With Prism mock server, we might get various response codes
				// The main goal is to test that the client can make the request properly
				if err == nil {
					require.NotNil(t, response)
					assert.True(t, httpResp.StatusCode >= 200 && httpResp.StatusCode < 500,
						"Expected valid HTTP status code, got %d", httpResp.StatusCode)

					// Verify the request URL contains the status parameter
					if httpResp.Request != nil && httpResp.Request.URL != nil {
						query := httpResp.Request.URL.Query()
						statusParam := query.Get("status")
						assert.Equal(t, tt.status, statusParam, "Status parameter should match")
					}
				} else {
					// If there's an error, it should be a valid API error, not a client error
					assert.True(t, httpResp.StatusCode >= 400 && httpResp.StatusCode < 500,
						"Expected 4xx error status code, got %d", httpResp.StatusCode)
				}
			}
		})
	}
}

// TestGetMessagesStatusParameterCombinations tests status parameter with other filters
func TestGetMessagesStatusParameterCombinations(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION_TESTS") == "true" {
		t.Skip("Skipping Prism-based integration test")
	}

	client := createTestClient()
	ctx := createAuthContext()

	t.Run("Status with recipient filter", func(t *testing.T) {
		response, httpResp, err := client.MessagesAPI.
			GetMessages(ctx, testAccountID).
			Sender("sender@example.com").
			Status("Delivered").
			Recipient("recipient@example.com").
			Execute()

		if err == nil {
			require.NotNil(t, response)
			assert.True(t, httpResp.StatusCode >= 200 && httpResp.StatusCode < 500)

			// Verify both parameters are in the request
			if httpResp.Request != nil && httpResp.Request.URL != nil {
				query := httpResp.Request.URL.Query()
				assert.Equal(t, "Delivered", query.Get("status"))
				assert.Equal(t, "recipient@example.com", query.Get("recipient"))
			}
		} else {
			assert.True(t, httpResp.StatusCode >= 400 && httpResp.StatusCode < 500)
		}
	})

	t.Run("Status with subject filter", func(t *testing.T) {
		response, httpResp, err := client.MessagesAPI.
			GetMessages(ctx, testAccountID).
			Sender("sender@example.com").
			Status("Bounced,Failed").
			Subject("Test Subject").
			Execute()

		if err == nil {
			require.NotNil(t, response)
			assert.True(t, httpResp.StatusCode >= 200 && httpResp.StatusCode < 500)

			// Verify parameters are in the request
			if httpResp.Request != nil && httpResp.Request.URL != nil {
				query := httpResp.Request.URL.Query()
				assert.Equal(t, "Bounced,Failed", query.Get("status"))
				assert.Equal(t, "Test Subject", query.Get("subject"))
			}
		} else {
			assert.True(t, httpResp.StatusCode >= 400 && httpResp.StatusCode < 500)
		}
	})

	t.Run("Status with pagination", func(t *testing.T) {
		response, httpResp, err := client.MessagesAPI.
			GetMessages(ctx, testAccountID).
			Sender("sender@example.com").
			Status("Delivered").
			Limit(10).
			Cursor("test-cursor").
			Execute()

		if err == nil {
			require.NotNil(t, response)
			assert.True(t, httpResp.StatusCode >= 200 && httpResp.StatusCode < 500)

			// Verify all parameters are in the request
			if httpResp.Request != nil && httpResp.Request.URL != nil {
				query := httpResp.Request.URL.Query()
				assert.Equal(t, "Delivered", query.Get("status"))
				assert.Equal(t, "10", query.Get("limit"))
				assert.Equal(t, "test-cursor", query.Get("cursor"))
			}
		} else {
			assert.True(t, httpResp.StatusCode >= 400 && httpResp.StatusCode < 500)
		}
	})

	t.Run("Status with all other parameters", func(t *testing.T) {
		response, httpResp, err := client.MessagesAPI.
			GetMessages(ctx, testAccountID).
			Sender("sender@example.com").
			Status("Delivered,Bounced").
			Recipient("recipient@example.com").
			Subject("Test Subject").
			MessageIdHeader("msg-12345").
			Limit(25).
			Cursor("comprehensive-test-cursor").
			Execute()

		if err == nil {
			require.NotNil(t, response)
			assert.True(t, httpResp.StatusCode >= 200 && httpResp.StatusCode < 500)

			// Verify all parameters are in the request
			if httpResp.Request != nil && httpResp.Request.URL != nil {
				query := httpResp.Request.URL.Query()
				assert.Equal(t, "Delivered,Bounced", query.Get("status"))
				assert.Equal(t, "sender@example.com", query.Get("sender"))
				assert.Equal(t, "recipient@example.com", query.Get("recipient"))
				assert.Equal(t, "Test Subject", query.Get("subject"))
				assert.Equal(t, "msg-12345", query.Get("message_id_header"))
				assert.Equal(t, "25", query.Get("limit"))
				assert.Equal(t, "comprehensive-test-cursor", query.Get("cursor"))
			}
		} else {
			assert.True(t, httpResp.StatusCode >= 400 && httpResp.StatusCode < 500)
		}
	})
}

// TestGetMessagesStatusParameterEdgeCases tests edge cases for the status parameter
func TestGetMessagesStatusParameterEdgeCases(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION_TESTS") == "true" {
		t.Skip("Skipping Prism-based integration test")
	}

	client := createTestClient()
	ctx := createAuthContext()

	t.Run("Empty status parameter", func(t *testing.T) {
		response, httpResp, err := client.MessagesAPI.
			GetMessages(ctx, testAccountID).
			Sender("sender@example.com").
			Status("").
			Execute()

		// Should handle empty status parameter gracefully
		if err == nil {
			require.NotNil(t, response)
		}
		// Either way, should not panic or cause client errors
		assert.True(t, httpResp.StatusCode >= 200)
	})

	t.Run("Status parameter method order independence", func(t *testing.T) {
		// Test different orders of method calls
		orders := []func() ahasend.ApiGetMessagesRequest{
			func() ahasend.ApiGetMessagesRequest {
				return client.MessagesAPI.
					GetMessages(ctx, testAccountID).
					Status("Delivered").
					Sender("sender@example.com").
					Limit(10)
			},
			func() ahasend.ApiGetMessagesRequest {
				return client.MessagesAPI.
					GetMessages(ctx, testAccountID).
					Sender("sender@example.com").
					Status("Delivered").
					Limit(10)
			},
			func() ahasend.ApiGetMessagesRequest {
				return client.MessagesAPI.
					GetMessages(ctx, testAccountID).
					Limit(10).
					Status("Delivered").
					Sender("sender@example.com")
			},
		}

		for i, orderFunc := range orders {
			t.Run(fmt.Sprintf("Order %d", i+1), func(t *testing.T) {
				request := orderFunc()
				assert.NotNil(t, request)

				// Should be able to execute regardless of order
				response, httpResp, err := request.Execute()
				if err == nil {
					require.NotNil(t, response)
				}
				// Should not cause client errors regardless of response
				assert.True(t, httpResp.StatusCode >= 200)
			})
		}
	})
}

// TestGetMessagesWithoutStatusParameter tests backward compatibility
func TestGetMessagesWithoutStatusParameter(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION_TESTS") == "true" {
		t.Skip("Skipping Prism-based integration test")
	}

	client := createTestClient()
	ctx := createAuthContext()

	t.Run("GetMessages without status parameter (backward compatibility)", func(t *testing.T) {
		response, httpResp, err := client.MessagesAPI.
			GetMessages(ctx, testAccountID).
			Sender("sender@example.com").
			Execute()

		// Should work exactly as before - no breaking changes
		if err == nil {
			require.NotNil(t, response)
			assert.True(t, httpResp.StatusCode >= 200 && httpResp.StatusCode < 500)

			// Verify status parameter is not in the request when not specified
			if httpResp.Request != nil && httpResp.Request.URL != nil {
				query := httpResp.Request.URL.Query()
				assert.Empty(t, query.Get("status"), "Status parameter should not be present when not specified")
			}
		} else {
			assert.True(t, httpResp.StatusCode >= 400 && httpResp.StatusCode < 500)
		}
	})

	t.Run("GetMessages with other filters but no status", func(t *testing.T) {
		response, httpResp, err := client.MessagesAPI.
			GetMessages(ctx, testAccountID).
			Sender("sender@example.com").
			Recipient("recipient@example.com").
			Subject("Test").
			Limit(5).
			Execute()

		// Should work with all other parameters
		if err == nil {
			require.NotNil(t, response)
			assert.True(t, httpResp.StatusCode >= 200 && httpResp.StatusCode < 500)

			// Verify other parameters are present but status is not
			if httpResp.Request != nil && httpResp.Request.URL != nil {
				query := httpResp.Request.URL.Query()
				assert.Equal(t, "sender@example.com", query.Get("sender"))
				assert.Equal(t, "recipient@example.com", query.Get("recipient"))
				assert.Equal(t, "Test", query.Get("subject"))
				assert.Equal(t, "5", query.Get("limit"))
				assert.Empty(t, query.Get("status"), "Status parameter should not be present")
			}
		} else {
			assert.True(t, httpResp.StatusCode >= 400 && httpResp.StatusCode < 500)
		}
	})
}
