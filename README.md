# 📧 AhaSend Go SDK

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.18-blue.svg)](https://golang.org/)
[![Go Reference](https://pkg.go.dev/badge/github.com/AhaSend/ahasend-go.svg)](https://pkg.go.dev/github.com/AhaSend/ahasend-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/AhaSend/ahasend-go)](https://goreportcard.com/report/github.com/AhaSend/ahasend-go)
[![API Documentation](https://img.shields.io/badge/docs-api-green.svg)](https://ahasend.com/docs/api-reference)
[![License: MIT](https://img.shields.io/github/license/ahasend/ahasend-go)](https://opensource.org/licenses/MIT)

The official Go SDK for [AhaSend](https://ahasend.com) 🚀 - a powerful transactional email service with high deliverability, comprehensive tracking, and developer-friendly APIs.

## ✨ Features

- **📦 Complete API Coverage**: Send emails, manage domains, webhooks, routes, suppressions, and more
- **🔒 Type Safety**: Full Go type system with pointer utilities for optional fields
- **⚡ Built-in Rate Limiting**: Automatic protection against 429 errors with configurable limits
- **🔄 Intelligent Retries**: Exponential backoff with jitter for failed requests
- **🔗 Webhook Processing**: Standard Webhooks compliant verification and parsing
- **📊 Comprehensive Tracking**: Opens, clicks, bounces, deliveries, and more
- **🛡️ Automatic Idempotency**: Prevent duplicate API calls (including email sends) automatically
- **📚 Rich Examples**: 11+ production-ready examples covering all major use cases

## Quick Start

### Installation

```bash
go get github.com/AhaSend/ahasend-go
```

### Send Your First Email

```go
package main

import (
    "context"
    "log"
    "os"

    "github.com/AhaSend/ahasend-go"
    "github.com/AhaSend/ahasend-go/api"
    "github.com/AhaSend/ahasend-go/models/common"
    "github.com/AhaSend/ahasend-go/models/requests"
    "github.com/google/uuid"
)

func main() {
    apiKey := os.Getenv("AHASEND_API_KEY")
    if apiKey == "" {
        log.Fatal("AHASEND_API_KEY environment variable is required")
    }

    accountID, err := uuid.Parse(os.Getenv("AHASEND_ACCOUNT_ID"))
    if err != nil {
        log.Fatalf("Invalid AHASEND_ACCOUNT_ID: %v", err)
    }

    client := api.NewAPIClient(api.WithAPIKey(apiKey))
    ctx := context.Background()

    // Send email
    message := requests.CreateMessageRequest{
        From: common.SenderAddress{Email: "sender@yourdomain.com"},
        Recipients: []common.Recipient{
            {Email: "recipient@example.com"},
        },
        Subject:     "Hello from AhaSend!",
        HtmlContent: ahasend.String("<h1>Welcome!</h1>"),
        TextContent: ahasend.String("Welcome!"),
    }

    response, _, err := client.MessagesAPI.CreateMessage(ctx, accountID, message)
    if err != nil {
        log.Fatal(err)
    }

    if len(response.Data) > 0 && response.Data[0].ID != nil {
        log.Printf("Email sent! Message ID: %s", *response.Data[0].ID)
    }
}
```

## Authentication & API Keys

All API requests require a Bearer token. There are three ways to authenticate:

### Environment Variable (Recommended)
```bash
# Set environment variable
export AHASEND_API_KEY="aha-sk-your-64-character-key"
```

```go
client := api.NewAPIClientFromEnv()
ctx := context.Background()
```

### Client-wide Configuration
```go
// Set API key when creating client
client := api.NewAPIClient(
    api.WithAPIKey(apiKey),
)

// or:
// cfg := api.NewConfiguration()
// cfg.APIKey = "aha-sk-..."
// client := api.NewAPIClientWithConfig(cfg)

// All subsequent API calls will use this key automatically
response, _, err := client.MessagesAPI.CreateMessage(ctx, accountID, message)
```

### Context Override (Per-request)
```go
// Override API key for specific requests
ctx := context.WithValue(context.Background(),
    api.ContextAccessToken, "aha-sk-your-64-character-key")

response, _, err := client.MessagesAPI.CreateMessage(ctx, accountID, message)
```

Get your API key from the [AhaSend Dashboard](https://dashboard.ahasend.com).

## Core Functionality

### Email Operations
- **Send Emails**: HTML/text content, attachments, scheduling
- **Batch Operations**: Efficient bulk sending
- **Message Management**: Cancel, retrieve status, view history

### Domain & Infrastructure
- **Domain Management**: Add, verify, and configure sending domains
- **DNS Validation**: Automated DNS record verification
- **Route Management**: Handle inbound email processing
- **SMTP Credentials**: Generate credentials for legacy applications

### Monitoring & Analytics
- **Delivery Statistics**: Track sends, deliveries, bounces, opens, clicks
- **Real-time Events**: Webhook notifications for all email events
- **Suppression Management**: Handle bounces and unsubscribes automatically

### Developer Experience
- **Automatic Rate Limiting**: Three endpoint categories with smart detection
- **Retry Configuration**: Multiple backoff strategies (exponential, linear, constant)
- **Error Handling**: Structured error types with detailed context
- **Comprehensive Testing**: Unit and integration tests with mock server

## API Reference

| Service | Description | Key Methods |
|---------|-------------|-------------|
| **MessagesAPI** | Send and manage emails | `CreateMessage`, `GetMessage`, `CancelMessage` |
| **DomainsAPI** | Domain verification & management | `CreateDomain`, `CheckDomainDNS`, `GetDomain` |
| **WebhooksAPI** | Event notifications | `CreateWebhook`, `UpdateWebhook`, `GetWebhooks` |
| **StatisticsAPI** | Email analytics | `GetDeliverabilityStatistics`, `GetBounceStatistics` |
| **SuppressionsAPI** | Manage block lists | `CreateSuppression`, `DeleteSuppression`, `GetSuppressions` |
| **RoutesAPI** | Inbound email handling | `CreateRoute`, `UpdateRoute` |
| **AccountsAPI** | Account & member management | `GetAccount`, `AddAccountMember` |
| **APIKeysAPI** | API key management | `CreateAPIKey`, `UpdateAPIKey` |

## Examples

Explore our [comprehensive examples](./examples/):

- **[send_email.go](./examples/send_email.go)** - Basic email sending
- **[send_with_attachments.go](./examples/send_with_attachments.go)** - File attachments
- **[batch_send.go](./examples/batch_send.go)** - Bulk email operations
- **[scheduled_send.go](./examples/scheduled_send.go)** - Schedule future delivery
- **[webhook_processing.go](./examples/webhook_processing.go)** - Handle webhook events
- **[webhook_management.go](./examples/webhook_management.go)** - Create and manage webhooks
- **[domain_management.go](./examples/domain_management.go)** - Domain setup & verification
- **[statistics.go](./examples/statistics.go)** - Analytics and reporting
- **[error_handling.go](./examples/error_handling.go)** - Robust error handling
- **[rate_limiting.go](./examples/rate_limiting.go)** - Rate limit configuration
- **[idempotency.go](./examples/idempotency.go)** - Prevent duplicate sends

Run any example:
```bash
# Set your credentials
export AHASEND_API_KEY="your-api-key"
export AHASEND_ACCOUNT_ID="your-account-id"

# Run example
go run examples/send_email.go
```

## Webhook Processing

The SDK includes Standard Webhooks compliant processing with HMAC-SHA256 verification:

```go
import "github.com/AhaSend/ahasend-go/webhooks"

// Create verifier
verifier, _ := webhooks.NewWebhookVerifier("your-webhook-secret")

// In your HTTP handler
func webhookHandler(w http.ResponseWriter, r *http.Request) {
    event, err := verifier.ParseRequest(r)
    if err != nil {
        http.Error(w, "Invalid webhook", 400)
        return
    }

    // Handle different event types
    switch e := event.(type) {
    case *webhooks.MessageDeliveredEvent:
        log.Printf("Email delivered to %s", e.Data.Recipient)
    case *webhooks.MessageBouncedEvent:
        log.Printf("Email bounced: %s", e.Data.Reason)
    }

    w.WriteHeader(200)
}
```

**Supported Events**: `message.*` (delivered, bounced, opened, clicked), `suppression.*`, `domain.*`, `route.*`

## Configuration

### Rate Limiting
```go
client := api.NewAPIClient(api.WithAPIKey(apiKey))

// Configure for high-volume sending
client.SetSendMessageRateLimit(500, 1000) // 500 req/s, 1000 burst

// Configure statistics polling
client.SetStatisticsRateLimit(10, 20) // 10 req/s, 20 burst
```

### Retry Configuration
```go
retryConfig := api.RetryConfig{
    Enabled:           true,
    MaxRetries:        3,
    BackoffStrategy:   api.BackoffExponential,
    BaseDelay:         time.Second,
    MaxDelay:          30 * time.Second,
}

client := api.NewAPIClient(
    api.WithAPIKey(apiKey),
    api.WithRetryConfig(retryConfig),
)
```

## Development

This project includes a comprehensive [Makefile](./Makefile) for all development tasks:

```bash
# Set up development environment
make setup

# Quick development cycle (format, lint, test)
make dev-test

# Full test suite with coverage
make full-test

# Show all available commands
make help
```

### Testing
- **Unit Tests**: `make test-unit`
- **Integration Tests**: `make test-integration` (requires Prism mock server)
- **Coverage Reports**: `make test-coverage`

## Related Projects

- **[AhaSend CLI](https://github.com/AhaSend/ahasend-cli)** - Command-line tool built on this SDK

## Documentation & Support

- 📚 [API Documentation](https://ahasend.com/docs)
- 🔗 [Go Package Documentation](https://pkg.go.dev/github.com/AhaSend/ahasend-go)
- 💬 [Support](mailto:support@ahasend.com)
- 🐛 [Issues](https://github.com/AhaSend/ahasend-go/issues)

## Requirements

- **Go**: 1.18 or later
- **Runtime dependencies**: `github.com/google/uuid`
- **Test dependencies**: `github.com/stretchr/testify`

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Built with ❤️ by the [AhaSend](https://ahasend.com) team
