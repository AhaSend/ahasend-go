package ahasend_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"testing"
	"time"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/api"
	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/AhaSend/ahasend-go/models/requests"
)

// MockServerConfig holds configuration for the mock server
type MockServerConfig struct {
	Host            string
	Port            string
	SpecPath        string
	LogLevel        string
	DynamicExamples bool
}

// DefaultMockServerConfig returns default configuration for integration tests
func DefaultMockServerConfig() MockServerConfig {
	return MockServerConfig{
		Host:            "localhost",
		Port:            "4010",
		SpecPath:        "api/openapi.yaml",
		LogLevel:        "warn",
		DynamicExamples: true,
	}
}

// MockServer manages the Prism mock server lifecycle
type MockServer struct {
	config MockServerConfig
	cmd    *exec.Cmd
	mutex  sync.Mutex

	// Channel for server readiness
	ready chan struct{}
	// Channel for errors during startup
	errCh chan error
}

// NewMockServer creates a new mock server instance
func NewMockServer(config MockServerConfig) *MockServer {
	return &MockServer{
		config: config,
		ready:  make(chan struct{}),
		errCh:  make(chan error, 1),
	}
}

// Start starts the mock server and waits for it to be ready
func (ms *MockServer) Start() error {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	// Determine prism command to use
	var prismCmd string
	var prismArgs []string

	// 1. Check for custom PRISM_CMD environment variable
	if customCmd := os.Getenv("PRISM_CMD"); customCmd != "" {
		prismCmd = customCmd
		prismArgs = []string{
			"mock",
			ms.config.SpecPath,
			"--host", ms.config.Host,
			"--port", ms.config.Port,
			"--log-level", ms.config.LogLevel,
		}
	} else if prismPath, err := exec.LookPath("prism"); err == nil {
		// 2. Use prism if it's in PATH
		prismCmd = prismPath
		prismArgs = []string{
			"mock",
			ms.config.SpecPath,
			"--host", ms.config.Host,
			"--port", ms.config.Port,
			"--log-level", ms.config.LogLevel,
		}
	} else if _, err := exec.LookPath("npx"); err == nil {
		// 3. Fall back to npx (most CI/CD environments have this)
		prismCmd = "npx"
		prismArgs = []string{
			"@stoplight/prism-cli",
			"mock",
			ms.config.SpecPath,
			"--host", ms.config.Host,
			"--port", ms.config.Port,
			"--log-level", ms.config.LogLevel,
		}
	} else {
		return fmt.Errorf("prism is not available. Install with: npm install -g @stoplight/prism-cli or ensure npx is available")
	}

	// Check if OpenAPI spec exists
	if _, err := os.Stat(ms.config.SpecPath); os.IsNotExist(err) {
		return fmt.Errorf("OpenAPI spec not found at %s", ms.config.SpecPath)
	}

	// Add dynamic examples flag if enabled
	if ms.config.DynamicExamples {
		prismArgs = append(prismArgs, "--dynamic")
	}

	// Create command
	ms.cmd = exec.Command(prismCmd, prismArgs...)

	// Capture output for debugging
	var stdout, stderr bytes.Buffer
	ms.cmd.Stdout = &stdout
	ms.cmd.Stderr = &stderr

	// Start the process
	err := ms.cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start prism: %w", err)
	}

	// Start readiness check in goroutine
	go ms.waitForReadiness(&stdout, &stderr)

	// Wait for server to be ready or error
	select {
	case <-ms.ready:
		return nil
	case err := <-ms.errCh:
		ms.Stop()
		return fmt.Errorf("mock server failed to start: %w", err)
	case <-time.After(15 * time.Second):
		ms.Stop()
		return fmt.Errorf("mock server did not become ready within 15 seconds.\nStdout: %s\nStderr: %s",
			stdout.String(), stderr.String())
	}
}

// waitForReadiness polls the server until it's ready
func (ms *MockServer) waitForReadiness(stdout, stderr *bytes.Buffer) {
	baseURL := fmt.Sprintf("http://%s:%s", ms.config.Host, ms.config.Port)
	pingURL := baseURL + "/v2/ping"

	client := &http.Client{Timeout: 1 * time.Second}
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	maxAttempts := 75 // 15 seconds total
	attempts := 0

	for range ticker.C {
		attempts++
		if attempts > maxAttempts {
			ms.errCh <- fmt.Errorf("server readiness check timed out after %d attempts", maxAttempts)
			return
		}

		resp, err := client.Get(pingURL)
		if err == nil && resp.StatusCode < 500 {
			resp.Body.Close()
			close(ms.ready)
			return
		}
		if resp != nil {
			resp.Body.Close()
		}

		// Check if process died
		if ms.cmd.ProcessState != nil && ms.cmd.ProcessState.Exited() {
			ms.errCh <- fmt.Errorf("prism process exited unexpectedly.\nStdout: %s\nStderr: %s",
				stdout.String(), stderr.String())
			return
		}
	}
}

// Stop stops the mock server
func (ms *MockServer) Stop() error {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	if ms.cmd == nil || ms.cmd.Process == nil {
		return nil
	}

	// Try graceful shutdown first
	err := ms.cmd.Process.Signal(os.Interrupt)
	if err != nil {
		// Force kill if graceful shutdown fails
		return ms.cmd.Process.Kill()
	}

	// Wait for process to exit
	done := make(chan error, 1)
	go func() {
		done <- ms.cmd.Wait()
	}()

	select {
	case <-done:
		return nil
	case <-time.After(5 * time.Second):
		// Force kill after timeout
		return ms.cmd.Process.Kill()
	}
}

// GetURL returns the base URL of the mock server
func (ms *MockServer) GetURL() string {
	return fmt.Sprintf("http://%s:%s", ms.config.Host, ms.config.Port)
}

// TestClientConfig creates a test client configuration for integration tests
type TestClientConfig struct {
	MockServerURL    string
	DisableRateLimit bool
	MaxRetries       int
	Debug            bool
}

// NewTestClient creates an AhaSend client configured for testing
func NewTestClient(config TestClientConfig) *api.APIClient {
	cfg := api.NewConfiguration()

	// Parse mock server URL
	if config.MockServerURL != "" {
		// Extract host and scheme from URL
		if config.MockServerURL[:7] == "http://" {
			cfg.Scheme = "http"
			cfg.Host = config.MockServerURL[7:]
		} else if config.MockServerURL[:8] == "https://" {
			cfg.Scheme = "https"
			cfg.Host = config.MockServerURL[8:]
		} else {
			cfg.Host = config.MockServerURL
			cfg.Scheme = "http"
		}
	}

	cfg.EnableRateLimit = !config.DisableRateLimit
	cfg.Debug = config.Debug

	if config.MaxRetries > 0 {
		cfg.RetryConfig.MaxRetries = config.MaxRetries
	} else {
		cfg.RetryConfig.MaxRetries = 1 // Faster tests
	}

	return api.NewAPIClientWithConfig(cfg)
}

// CreateTestAuthContext creates a context with test authentication
func CreateTestAuthContext(apiKey string) context.Context {
	if apiKey == "" {
		apiKey = "aha-sk-test-key-for-integration-testing-12345678901234567890"
	}
	return context.WithValue(context.Background(), api.ContextAccessToken, apiKey)
}

// IntegrationTestData contains test data generators for various scenarios
type IntegrationTestData struct{}

// NewIntegrationTestData creates a new test data generator
func NewIntegrationTestData() *IntegrationTestData {
	return &IntegrationTestData{}
}

// CreateMessageRequest generates a test message request
func (td *IntegrationTestData) CreateMessageRequest(overrides ...func(*requests.CreateMessageRequest)) requests.CreateMessageRequest {
	req := requests.CreateMessageRequest{
		From: common.SenderAddress{
			Email: "test@example.com",
		},
		Recipients: []common.Recipient{
			{
				Email: "recipient@example.com",
			},
		},
		Subject:     "Test Email",
		HtmlContent: ahasend.String("<p>This is a test email from integration tests.</p>"),
		TextContent: ahasend.String("This is a test email from integration tests."),
	}

	// Apply overrides
	for _, override := range overrides {
		override(&req)
	}

	return req
}

// CreateDomainRequest generates a test domain creation request
func (td *IntegrationTestData) CreateDomainRequest(domain string) requests.CreateDomainRequest {
	if domain == "" {
		domain = "test-example.com"
	}

	return requests.CreateDomainRequest{
		Domain: domain,
	}
}

// CreateWebhookRequest generates a test webhook creation request
func (td *IntegrationTestData) CreateWebhookRequest() requests.CreateWebhookRequest {
	return requests.CreateWebhookRequest{
		Name: "Integration Test Webhook",
		URL:  "https://example.com/webhook",
	}
}

// IntegrationTestRunner provides utilities for running integration tests
type IntegrationTestRunner struct {
	mockServer *MockServer
	client     *api.APIClient
	ctx        context.Context
	testData   *IntegrationTestData
}

// NewIntegrationTestRunner creates a new test runner with mock server
func NewIntegrationTestRunner(t *testing.T) *IntegrationTestRunner {
	config := DefaultMockServerConfig()
	mockServer := NewMockServer(config)

	err := mockServer.Start()
	if err != nil {
		t.Fatalf("Failed to start mock server: %v", err)
	}

	client := NewTestClient(TestClientConfig{
		MockServerURL:    mockServer.GetURL(),
		DisableRateLimit: true,
		MaxRetries:       1,
		Debug:            false,
	})

	ctx := CreateTestAuthContext("")

	// Ensure cleanup
	t.Cleanup(func() {
		mockServer.Stop()
	})

	return &IntegrationTestRunner{
		mockServer: mockServer,
		client:     client,
		ctx:        ctx,
		testData:   NewIntegrationTestData(),
	}
}

// Client returns the configured AhaSend client
func (tr *IntegrationTestRunner) Client() *api.APIClient {
	return tr.client
}

// Context returns the authenticated context
func (tr *IntegrationTestRunner) Context() context.Context {
	return tr.ctx
}

// TestData returns the test data generator
func (tr *IntegrationTestRunner) TestData() *IntegrationTestData {
	return tr.testData
}

// MockServerURL returns the URL of the running mock server
func (tr *IntegrationTestRunner) MockServerURL() string {
	return tr.mockServer.GetURL()
}

// WaitForMockServer waits for the mock server to be ready
func (tr *IntegrationTestRunner) WaitForMockServer() error {
	// The mock server should already be ready from NewIntegrationTestRunner
	return nil
}

// PerformanceTest runs a performance test with the given function
func (tr *IntegrationTestRunner) PerformanceTest(t *testing.T, name string, iterations int, fn func()) {
	t.Run(name, func(t *testing.T) {
		// Disable rate limiting for performance tests
		tr.client.SetGlobalRateLimit(false)

		start := time.Now()

		for i := 0; i < iterations; i++ {
			fn()
		}

		duration := time.Since(start)
		avgDuration := duration / time.Duration(iterations)

		t.Logf("Performance Test: %s", name)
		t.Logf("Total time: %v", duration)
		t.Logf("Iterations: %d", iterations)
		t.Logf("Average per iteration: %v", avgDuration)
		t.Logf("Requests per second: %.2f", float64(iterations)/duration.Seconds())
	})
}

// ValidateJSONResponse validates that a response can be marshaled/unmarshaled correctly
func ValidateJSONResponse(t *testing.T, response interface{}) {
	// Marshal to JSON
	jsonData, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("Failed to marshal response to JSON: %v", err)
	}

	// Unmarshal back to map to validate structure
	var jsonMap map[string]interface{}
	err = json.Unmarshal(jsonData, &jsonMap)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON response: %v", err)
	}

	// Verify it's not empty
	if len(jsonMap) == 0 {
		t.Fatal("Response appears to be empty after JSON roundtrip")
	}
}

// AssertNoRateLimit asserts that the operation completes quickly (no rate limiting)
func AssertNoRateLimit(t *testing.T, operation func() error, maxDuration time.Duration) {
	start := time.Now()
	err := operation()
	duration := time.Since(start)

	if err != nil {
		t.Fatalf("Operation failed: %v", err)
	}

	if duration > maxDuration {
		t.Fatalf("Operation took too long (%v), suggesting rate limiting is active", duration)
	}
}

// AssertRateLimit asserts that the operation is rate limited (takes longer than expected)
func AssertRateLimit(t *testing.T, operation func() error, minDuration time.Duration) {
	start := time.Now()
	err := operation()
	duration := time.Since(start)

	if err != nil {
		t.Fatalf("Operation failed: %v", err)
	}

	if duration < minDuration {
		t.Fatalf("Operation completed too quickly (%v), rate limiting may not be working", duration)
	}
}

// CaptureHTTPRequest captures the HTTP request details for verification
func CaptureHTTPRequest(client *api.APIClient) *HTTPRequestCapture {
	return &HTTPRequestCapture{
		client: client,
	}
}

// HTTPRequestCapture helps capture and analyze HTTP requests
type HTTPRequestCapture struct {
	client      *api.APIClient
	lastRequest *http.Request
}

// GetLastRequest returns the last captured request
func (hrc *HTTPRequestCapture) GetLastRequest() *http.Request {
	return hrc.lastRequest
}

// HasHeader checks if the last request had a specific header
func (hrc *HTTPRequestCapture) HasHeader(name string) bool {
	if hrc.lastRequest == nil {
		return false
	}
	return hrc.lastRequest.Header.Get(name) != ""
}

// GetHeader returns the value of a specific header from the last request
func (hrc *HTTPRequestCapture) GetHeader(name string) string {
	if hrc.lastRequest == nil {
		return ""
	}
	return hrc.lastRequest.Header.Get(name)
}

// ReadRequestBody reads and returns the body of the last request
func (hrc *HTTPRequestCapture) ReadRequestBody() ([]byte, error) {
	if hrc.lastRequest == nil || hrc.lastRequest.Body == nil {
		return nil, fmt.Errorf("no request or request body available")
	}

	return io.ReadAll(hrc.lastRequest.Body)
}
