# AhaSend Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/AhaSend/ahasend-go.svg)](https://pkg.go.dev/github.com/AhaSend/ahasend-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/AhaSend/ahasend-go)](https://goreportcard.com/report/github.com/AhaSend/ahasend-go)

The official Go SDK for the [AhaSend](https://ahasend.com) API v2. Send transactional emails, manage domains, webhooks, routes, API keys, and view statistics with a developer-friendly interface that includes automatic rate limiting, idempotency, and intelligent retry logic.

## Features

- 🚀 **Simple & Fast**: Get started with just a few lines of code
- 🔒 **Secure by Default**: Built-in authentication and secure defaults
- 🛡️ **Automatic Protection**: Rate limiting, idempotency, and retries included
- 🪝 **Webhook Processing**: Standard Webhooks compliant verification and parsing
- ⚡ **Thread-Safe**: Safe for concurrent use from multiple goroutines
- 🎯 **Context-Aware**: Full support for Go contexts and cancellation
- 🔧 **Highly Configurable**: Environment variables, custom HTTP clients, and more
- 📖 **Comprehensive Examples**: [12+ real-world examples](examples/) included

## Quick Start

### Installation

```bash
go get github.com/AhaSend/ahasend-go
```

### Basic Usage

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/AhaSend/ahasend-go"
    "github.com/google/uuid"
)

func main() {
    // Set your API key (get one at https://ahasend.com)
    os.Setenv("AHASEND_API_KEY", "aha-sk-your-api-key-here")

    // Create client from environment variables (recommended)
    client := ahasend.NewAPIClientFromEnv()

    // Create authenticated context
    ctx := ahasend.ContextWithEnvAuth(context.Background())

    // Parse your account ID
    accountID := uuid.MustParse("your-account-id")

    // Send your first email
    message := ahasend.CreateMessageRequest{
        From: ahasend.SenderAddress{
            Email: "hello@yourdomain.com",
            Name:  ahasend.PtrString("Your App"),
        },
        Recipients: []ahasend.Recipient{
            {
                Email: "user@example.com",
                Name:  ahasend.PtrString("User"),
            },
        },
        Subject:     "Welcome to Your App!",
        TextContent: ahasend.PtrString("Thank you for signing up!"),
        HtmlContent: ahasend.PtrString("<h1>Welcome!</h1><p>Thank you for signing up!</p>"),
    }

    response, httpResp, err := client.MessagesAPI.
        CreateMessage(ctx, accountID).
        CreateMessageRequest(message).
        Execute()

    if err != nil {
        log.Fatalf("Error sending email: %v", err)
    }

    fmt.Printf("✅ Email sent successfully! (Status: %d)\n", httpResp.StatusCode)
    if response.Data != nil && len(response.Data) > 0 {
        fmt.Printf("Message ID: %s\n", *response.Data[0].Id)
    }
}
```

## Configuration

### Environment Variables (Recommended)

The SDK supports comprehensive configuration through environment variables, following the [12-Factor App](https://12factor.net/config) methodology:

```bash
# Required
export AHASEND_API_KEY="aha-sk-your-api-key-here"

# Optional configuration
export AHASEND_DEBUG=false
export AHASEND_MAX_RETRIES=3
export AHASEND_ENABLE_RATE_LIMIT=true
export AHASEND_IDEMPOTENCY_PREFIX="myapp"
```

**Quick Start with Environment:**
```go
// Automatically loads all configuration from environment
client := ahasend.NewAPIClientFromEnv()
ctx := ahasend.ContextWithEnvAuth(context.Background())

// Ready to use!
response, _, err := client.MessagesAPI.CreateMessage(ctx, accountID).Execute()
```

📖 **See [ENV.md](ENV.md) for complete environment variable documentation and examples**

### Programmatic Configuration

For more control, configure the client programmatically:

```go
cfg := ahasend.NewConfiguration()
cfg.Host = "api.ahasend.com"
cfg.Debug = true
cfg.MaxRetries = 5

// Configure rate limiting
client.SetSendMessageRateLimit(500, 1000) // 500 req/s, 1000 burst

// Configure idempotency
cfg.IdempotencyConfig = &ahasend.IdempotencyConfig{
    Enabled:           true,
    KeyPrefix:         "myapp",
    AutoGenerateForPOST: true,
}

client := ahasend.NewAPIClient(cfg)
```

## Core Features

### Rate Limiting

The SDK includes intelligent, automatic rate limiting that protects against 429 errors:

- **Enabled by default** - No configuration needed
- **Three endpoint types**: General API (100 req/s), Statistics (1 req/s), Send Message (100 req/s)
- **Smart detection** - Automatically categorizes endpoints by URL pattern
- **Token bucket algorithm** - Allows burst traffic up to configured limits

```go
// Configure custom rate limits based on your plan
client.SetSendMessageRateLimit(500, 1000) // 500 req/s, 1000 burst capacity
client.SetStatisticsRateLimit(10, 20)     // 10 req/s, 20 burst capacity

// Monitor rate limit status
status := client.GetRateLimitStatus(ahasend.SendMessageAPI)
fmt.Printf("Tokens available: %d/%d\n", status.TokensAvailable, status.BurstCapacity)
```

### Idempotency

Prevent duplicate operations with automatic or manual idempotency keys:

```go
// Automatic idempotency (enabled by default)
response, _, err := client.MessagesAPI.CreateMessage(ctx, accountID).
    CreateMessageRequest(message).
    Execute() // SDK automatically adds idempotency key

// Manual idempotency key
response, _, err := client.MessagesAPI.CreateMessage(ctx, accountID).
    CreateMessageRequest(message).
    IdempotencyKey("order-confirmation-12345").
    Execute()
```

📖 **See [IDEMPOTENCY.md](IDEMPOTENCY.md) for comprehensive idempotency documentation**

### Thread Safety

The SDK is **thread-safe** and designed for concurrent use:

```go
// ✅ Recommended: Create one client, share across goroutines
client := ahasend.NewAPIClient(cfg)

// ✅ Safe: Concurrent API calls from multiple goroutines
var wg sync.WaitGroup
for i := 0; i < 100; i++ {
    wg.Add(1)
    go func(index int) {
        defer wg.Done()
        // All goroutines can safely use the same client
        response, _, err := client.MessagesAPI.CreateMessage(ctx, accountID).Execute()
    }(i)
}
wg.Wait()
```

### Context Cancellation

Full support for Go contexts and request cancellation:

```go
// Request timeout
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

response, _, err := client.MessagesAPI.CreateMessage(ctx, accountID).Execute()
if errors.Is(err, context.DeadlineExceeded) {
    fmt.Println("Request timed out")
}

// Request cancellation
ctx, cancel := context.WithCancel(context.Background())
go func() {
    time.Sleep(5*time.Second)
    cancel() // Cancel after 5 seconds
}()

// This request will be cancelled if it takes longer than 5 seconds
response, _, err := client.MessagesAPI.CreateMessage(ctx, accountID).Execute()
```

### Webhook Processing

The SDK includes built-in support for verifying and processing AhaSend webhooks using the [Standard Webhooks](https://www.standardwebhooks.com/) specification:

```go
// Create webhook verifier with your webhook secret
verifier, err := ahasend.NewWebhookVerifier("whsec_your_webhook_secret")
if err != nil {
    log.Fatal(err)
}

// In your HTTP handler
func webhookHandler(w http.ResponseWriter, r *http.Request) {
    // Verify and parse the webhook
    event, err := verifier.ParseRequest(r)
    if err != nil {
        http.Error(w, "Invalid webhook", http.StatusBadRequest)
        return
    }

    // Process different event types
    switch e := event.(type) {
    case *ahasend.MessageDeliveredEvent:
        log.Printf("Email delivered to %s", e.Data.Recipient)
    case *ahasend.MessageBouncedEvent:
        log.Printf("Email bounced for %s", e.Data.Recipient)
    case *ahasend.MessageOpenedEvent:
        log.Printf("Email opened by %s", e.Data.Recipient)
    case *ahasend.MessageClickedEvent:
        log.Printf("Link clicked by %s: %s", e.Data.Recipient, e.Data.URL)
    }

    w.WriteHeader(http.StatusOK)
}
```

**Supported Webhook Events:**
- **Message Events**: `reception`, `delivered`, `transient_error`, `failed`, `bounced`, `suppressed`, `opened`, `clicked`
- **Suppression Events**: `created`
- **Domain Events**: `dns_error`
- **Route Events**: `message` (for inbound emails)

📖 **See [webhook_processing.go](examples/webhook_processing.go) for a complete webhook server example**

## Examples

We provide comprehensive, real-world examples for all major use cases:

| Example | Description |
|---------|-------------|
| [**send_email.go**](examples/send_email.go) | Basic transactional email sending |
| [**send_with_attachments.go**](examples/send_with_attachments.go) | Sending emails with file attachments |
| [**batch_send.go**](examples/batch_send.go) | Concurrent batch email sending |
| [**error_handling.go**](examples/error_handling.go) | Comprehensive error handling patterns |
| [**domain_management.go**](examples/domain_management.go) | Domain verification and management |
| [**webhook_management.go**](examples/webhook_management.go) | Webhook setup and event handling |
| [**webhook_processing.go**](examples/webhook_processing.go) | Webhook verification and event processing |
| [**statistics.go**](examples/statistics.go) | Email analytics and performance metrics |
| [**rate_limiting.go**](examples/rate_limiting.go) | Rate limit configuration and monitoring |
| [**idempotency.go**](examples/idempotency.go) | Preventing duplicate sends |
| [**scheduled_send.go**](examples/scheduled_send.go) | Scheduling emails for future delivery |

📖 **See [examples/README.md](examples/README.md) for detailed usage instructions**

### Running Examples

```bash
# Set your credentials
export AHASEND_API_KEY="aha-sk-your-api-key"
export AHASEND_ACCOUNT_ID="your-account-id"

# Run any example
go run examples/send_email.go
go run examples/batch_send.go
go run examples/webhook_management.go
```

## Authentication & Scopes

All API requests must be authenticated using a Bearer token:

```go
// Using environment variables (recommended)
ctx := ahasend.ContextWithEnvAuth(context.Background())

// Or programmatically
ctx := context.WithValue(context.Background(), ahasend.ContextAccessToken, "aha-sk-your-key")
```

### API Key Scopes

API keys have specific scopes that control access to different resources:

#### Message Scopes
- `messages:send:all` - Send messages from any domain in the account
- `messages:send:{domain}` - Send messages from a specific domain
- `messages:cancel:all` - Cancel messages from any domain
- `messages:read:all` - Read messages from any domain

#### Domain Scopes
- `domains:read` - Read all domains
- `domains:write` - Create and update domains
- `domains:delete:all` - Delete any domain

#### Account Scopes
- `accounts:read` - Read account information
- `accounts:write` - Update account settings
- `accounts:members:read` - Read account members

#### Webhook Scopes
- `webhooks:read:all` - Read all webhooks
- `webhooks:write:all` - Create and update webhooks
- `webhooks:delete:all` - Delete any webhook

#### Other Scopes
- `routes:read:all`, `routes:write:all` - Route management
- `suppressions:read`, `suppressions:write` - Suppression list management
- `smtp-credentials:read:all`, `smtp-credentials:write:all` - SMTP credential management
- `statistics-transactional:read:all` - Access email statistics
- `api-keys:read`, `api-keys:write`, `api-keys:delete` - API key management

## Advanced Configuration

### Custom HTTP Client

Configure the underlying HTTP client for advanced networking requirements:

```go
// Corporate proxy configuration
proxyURL, _ := url.Parse("http://corporate-proxy:8080")
transport := &http.Transport{
    Proxy: http.ProxyURL(proxyURL),
    MaxIdleConns: 100,
    MaxIdleConnsPerHost: 10,
    IdleConnTimeout: 90 * time.Second,
}

cfg := ahasend.NewConfiguration()
cfg.HTTPClient = &http.Client{
    Transport: transport,
    Timeout:   30 * time.Second,
}

client := ahasend.NewAPIClient(cfg)
```

### Production Optimization

```go
// Optimized production configuration
cfg := ahasend.NewConfiguration()
cfg.Debug = false
cfg.MaxRetries = 5

// High-performance HTTP settings
transport := &http.Transport{
    MaxIdleConns:        200,
    MaxIdleConnsPerHost: 20,
    MaxConnsPerHost:     50,
    IdleConnTimeout:     90 * time.Second,
    TLSHandshakeTimeout: 10 * time.Second,
}

cfg.HTTPClient = &http.Client{
    Transport: transport,
    Timeout:   30 * time.Second,
}

// Configure rate limits based on your plan
client := ahasend.NewAPIClient(cfg)
client.SetSendMessageRateLimit(1000, 2000) // High-volume plan
```

### Configuration Validation

```go
// Validate configuration before use
cfg := ahasend.NewConfiguration()
result := ahasend.ValidateConfiguration(cfg)

if result.HasErrors() {
    log.Fatalf("Configuration errors: %s", result.Error())
}

// Check production readiness
ready, issues := ahasend.IsProductionReady(cfg)
if !ready {
    log.Printf("Production readiness issues: %v", issues)
    ahasend.OptimizeForProduction(cfg) // Automatically optimize
}
```

## Error Handling

The SDK provides structured error handling:

```go
response, httpResp, err := client.MessagesAPI.CreateMessage(ctx, accountID).Execute()

if err != nil {
    // Check for API errors
    if apiErr, ok := err.(*ahasend.GenericOpenAPIError); ok {
        fmt.Printf("API Error: %s (Status: %d)\n", apiErr.Error(), apiErr.StatusCode())

        // Parse structured error response
        if len(apiErr.Body()) > 0 {
            var errorResp ahasend.ErrorResponse
            if json.Unmarshal(apiErr.Body(), &errorResp) == nil {
                fmt.Printf("Error Code: %s\n", *errorResp.Error)
                fmt.Printf("Message: %s\n", *errorResp.Message)
            }
        }
    }

    // Check for context errors
    if errors.Is(err, context.DeadlineExceeded) {
        fmt.Println("Request timed out")
    } else if errors.Is(err, context.Canceled) {
        fmt.Println("Request was cancelled")
    }
}
```

## Rate Limiting

The SDK includes built-in rate limiting that protects against 429 errors:

- **General API endpoints**: 100 requests per second, 200 burst capacity
- **Statistics endpoints**: 1 request per second, 1 burst capacity
- **Send Message endpoints**: 100 requests per second, 200 burst capacity

Rate limits are automatically detected and enforced per endpoint type.

## Pagination

List endpoints use cursor-based pagination:

```go
// Get paginated results
response, _, err := client.MessagesAPI.
    GetMessages(ctx, accountID).
    Limit(50).
    Execute()

if err != nil {
    log.Fatalf("Error: %v", err)
}

// Process current page
for _, message := range response.Data {
    fmt.Printf("Message: %s\n", *message.Subject)
}

// Get next page if available
if response.Pagination != nil && response.Pagination.NextCursor != nil {
    nextResponse, _, err := client.MessagesAPI.
        GetMessages(ctx, accountID).
        Cursor(*response.Pagination.NextCursor).
        Execute()
    // Process next page...
}
```

## Time Formats

All timestamps must be in RFC3339 format:

```go
// Scheduling emails
sendTime := time.Now().Add(2 * time.Hour)
message.Schedule = &ahasend.MessageSchedule{
    ScheduledAt: ahasend.PtrString(sendTime.Format(time.RFC3339)),
    Timezone:    ahasend.PtrString("America/New_York"),
}

// Date filtering for statistics
startDate := time.Now().AddDate(0, 0, -30) // 30 days ago
response, _, err := client.StatisticsAPI.
    GetDeliverabilityStatistics(ctx, accountID).
    StartDate(startDate.Format(time.RFC3339)).
    EndDate(time.Now().Format(time.RFC3339)).
    Execute()
```

## Deployment Examples

### Docker

```dockerfile
FROM golang:1.21
WORKDIR /app

# Environment variables
ENV AHASEND_DEBUG=false
ENV AHASEND_MAX_RETRIES=5
ENV AHASEND_ENABLE_RATE_LIMIT=true

COPY . .
RUN go build -o app

CMD ["./app"]
```

```bash
# Run with API key
docker run -e AHASEND_API_KEY="aha-sk-your-key" your-app
```

### Kubernetes

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: ahasend-secret
data:
  AHASEND_API_KEY: <aha-sk-api-key>

---
apiVersion: v1
kind: ConfigMap
metadata:
  name: ahasend-config
data:
  AHASEND_DEBUG: "false"
  AHASEND_MAX_RETRIES: "5"
  AHASEND_ENABLE_RATE_LIMIT: "true"

---
apiVersion: apps/v1
kind: Deployment
spec:
  template:
    spec:
      containers:
      - name: app
        envFrom:
        - secretRef:
            name: ahasend-secret
        - configMapRef:
            name: ahasend-config
```

## Testing

Run the test suite:

```bash
# Install dependencies
go mod tidy

# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run integration tests (requires Prism mock server)
go test ./test/...
```

### Setting up Integration Tests

The integration tests use Prism to mock the AhaSend API:

```bash
# Install Prism (if not already installed)
npm install -g @stoplight/prism-cli

# Run integration tests
export AHASEND_API_KEY="test-key"
export AHASEND_ACCOUNT_ID="test-account-id"
go test ./test/integration_test.go
```

## API Reference

API reference is available at [AhaSend Docs](https://ahasend.com/docs/api-reference).

## Utility Functions

The SDK provides utility functions for working with pointer fields:

```go
// Create pointers to basic types
message := ahasend.CreateMessageRequest{
    Subject:     "Hello World",
    TextContent: ahasend.PtrString("Email content"),
    Priority:    ahasend.PtrInt(1),
    SendAt:      ahasend.PtrTime(time.Now().Add(time.Hour)),
    Testing:     ahasend.PtrBool(true),
}

// Available utility functions:
// PtrBool, PtrInt, PtrInt32, PtrInt64
// PtrFloat, PtrFloat32, PtrFloat64
// PtrString, PtrTime
```

### Debug Mode

Enable debug logging to troubleshoot issues:

```bash
export AHASEND_DEBUG=true
```

```go
cfg := ahasend.NewConfiguration()
cfg.Debug = true
client := ahasend.NewAPIClient(cfg)
```

### Getting Help

- 📖 **Documentation**: [ENV.md](ENV.md) | [IDEMPOTENCY.md](IDEMPOTENCY.md) | [examples/](examples/)
- 🐛 **Issues**: [GitHub Issues](https://github.com/AhaSend/ahasend-go/issues)
- 💬 **Support**: support@ahasend.com
- 🌐 **API Docs**: https://ahasend.com/docs

## Contributing

We welcome contributions to the AhaSend Go SDK! 🎉

### Quick Start for Contributors

1. **Fork and setup**: Fork the repository and run `go mod tidy`
2. **Create a feature branch**: `git checkout -b feature/your-feature-name`
3. **Make your changes**: Follow our coding standards and add comprehensive tests
4. **Test thoroughly**: Run `go fmt ./...`, `go vet ./...`, and `SKIP_INTEGRATION_TESTS=true go test -v ./test/`
5. **Commit with style**: Use [Conventional Commits](https://www.conventionalcommits.org/) format
6. **Submit a PR**: Open a detailed pull request with our PR template

### What We're Looking For

- 🐛 **Bug fixes** - Rate limiter issues, context cancellation, memory leaks
- ✨ **New features** - Additional API endpoints, performance optimizations
- 📚 **Documentation** - Real-world examples, troubleshooting guides, best practices
- 🧪 **Testing** - Edge case coverage, integration tests, benchmarks

### Contribution Areas

Whether you're fixing a typo or adding a major feature, all contributions are valued! Check our comprehensive **[CONTRIBUTING.md](CONTRIBUTING.md)** for detailed guidelines including:

- Development environment setup with Prism mock server
- Code quality standards and architectural guidelines
- Testing strategies (unit, integration, race condition detection)
- PR templates and commit message conventions
- Examples of different contribution types

**Get started**: Read [CONTRIBUTING.md](CONTRIBUTING.md) • Browse [open issues](https://github.com/AhaSend/ahasend-go/issues) • Join [GitHub Discussions](https://github.com/AhaSend/ahasend-go/discussions)

## License

This SDK is released under the MIT License. See [LICENSE](LICENSE) for details.

## Support

- **Email**: support@ahasend.com
- **Documentation**: https://ahasend.com/docs
- **Status Page**: https://status.ahasend.com

---

Made with ❤️ by the AhaSend team