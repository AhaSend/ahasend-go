# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is the **AhaSend Go SDK** for the AhaSend API v2. The AhaSend API provides transactional email services, domain management, webhook management, routing, API keys, and email statistics.

The SDK was originally generated from the OpenAPI specification located at `api/openapi.yaml`, but is now actively maintained and enhanced with additional features, better error handling, comprehensive testing, and improved developer experience.

## Development Commands

This repository includes a comprehensive `Makefile` that provides shortcuts for all common development tasks. **Always prefer using the Makefile commands over direct Go commands** as they include additional setup, validation, and consistency checks.

### Quick Reference
```bash
make help        # Show all available commands with descriptions
make setup       # Set up development environment (run this first)
make dev-test    # Quick development test cycle (fmt, lint, vet, test)
make test        # Run all tests
make build       # Build and validate the SDK
make clean       # Clean up build artifacts
```

### Makefile Overview

The Makefile is organized into several categories:

#### **Setup & Dependencies**
- `make setup` - Complete development environment setup (installs tools + deps)
- `make install-tools` - Install required development tools (golangci-lint, prism, etc.)
- `make deps` - Download and verify dependencies
- `make check-deps` - Check if dependencies are up to date
- `make update-deps` - Update all dependencies

#### **Code Quality**
- `make fmt` - Format Go code (gofmt + goimports)
- `make lint` - Run comprehensive linting with golangci-lint
- `make vet` - Run go vet for static analysis
- `make dev-test` - Quick development cycle (fmt + lint + vet + unit tests)
- `make security` - Run security scans (gosec + govulncheck)

#### **Testing**
- `make test` - Run all tests (currently just unit tests)
- `make test-unit` - Run unit tests only (skips integration tests)
- `make test-integration` - Run integration tests (requires Prism mock server)
- `make test-coverage` - Generate test coverage reports (HTML + text)
- `make benchmark` - Run performance benchmarks
- `make perf-test` - Detailed performance testing

#### **Build & Validation**
- `make build` - Build the SDK for validation
- `make mock-server` - Start Prism mock server for manual testing
- `make mock-server-validate` - Validate OpenAPI specification
- `make ci` - Simulate complete CI pipeline locally

#### **Development Workflow**
- `make full-test` - Complete test cycle (all quality checks + tests + coverage)
- `make watch` - Watch for file changes and auto-run tests (requires inotify-tools)
- `make setup-hooks` - Set up git pre-commit hooks
- `make release-check` - Verify readiness for release

#### **Reporting & Debugging**
- `make quality-report` - Generate HTML reports (lint, coverage, security)
- `make stats` - Show project statistics (files, lines, packages)
- `make version` - Display version and build information
- `make debug-info` - Show debugging information and environment

### Testing Framework Details

**Test Organization**:
- All test files are in the `test/` directory
- Tests use the testify framework (`github.com/stretchr/testify`)
- Some auto-generated tests are skipped by default (using `t.Skip("skip test")`)
- Webhook processing tests are fully active and comprehensive

**Running Tests**:
```bash
# Quick unit tests during development
make dev-test

# Run unit tests only
make test-unit

# Run integration tests (requires Prism)
make test-integration

# Full test suite with coverage
make full-test
```

**Integration Testing**:
- Uses Prism mock server for API mocking
- Automatically started/stopped during test runs
- Validates against the OpenAPI specification
- Install Prism with: `npm install -g @stoplight/prism-cli`

### Recommended Development Workflow

1. **Initial Setup**: `make setup` (installs tools and dependencies)
2. **Development Cycle**: `make dev-test` (quick validation during coding)
3. **Before Committing**: `make full-test` (complete validation)
4. **Before Release**: `make release-check` (final validation)

## Code Architecture

### Core Structure

The SDK follows OpenAPI Generator's standard Go client structure:

**Configuration & Client (`configuration.go`, `client.go`)**
- `Configuration`: Holds client configuration including servers, authentication, HTTP client settings
- `APIClient`: Main client struct that manages all API services and HTTP requests
- Authentication is handled via Bearer tokens in the `Authorization` header

**API Services (`api_*.go` files)**
Each API service corresponds to a resource type:
- `MessagesAPIService`: Send, cancel, and retrieve messages
- `DomainsAPIService`: Manage email domains
- `AccountsAPIService`: Account management and members
- `APIKeysAPIService`: API key CRUD operations
- `WebhooksAPIService`: Webhook management
- `RoutesAPIService`: Email routing rules
- `SMTPCredentialsAPIService`: SMTP credential management
- `SuppressionsAPIService`: Email suppression list management
- `StatisticsAPIService`: Email delivery statistics
- `UtilityAPIService`: Utility endpoints like ping

**Data Models (`model_*.go` files)**
- All API request/response structures
- Includes utility pointer functions for optional fields
- Models use pointer fields to handle optional/nullable values

**Webhook Processing (`webhooks.go`, `webhook_events.go`)**
- `webhooks.go`: Standard Webhooks compliant verification
  - `WebhookVerifier`: HMAC-SHA256 signature verification
  - Timestamp validation to prevent replay attacks
  - Support for multiple signatures (key rotation)
- `webhook_events.go`: Strongly typed webhook event models
  - Message events (reception, delivered, failed, bounced, opened, clicked, etc.)
  - Suppression events (created)
  - Domain events (DNS errors)
  - Route events (inbound email processing)
  - Helper functions for event type checking

**Utilities (`utils.go`, `response.go`)**
- `utils.go`: Pointer utility functions (`PtrString`, `PtrBool`, etc.) and nullable type implementations
- `response.go`: HTTP response handling utilities

### Key Patterns

**Request Building Pattern**
All API methods return request builders that allow method chaining:
```go
resp, httpRes, err := client.MessagesAPI.CreateMessage(ctx, accountId).
    CreateMessageRequest(req).
    IdempotencyKey("unique-key").
    Execute()
```

**Authentication**
Bearer token authentication is set via context:
```go
auth := context.WithValue(context.Background(), ahasend.ContextAccessToken, "aha-sk-...")
```

**Pointer Utilities**
Since all model fields are pointers for optional handling, use utility functions:
```go
request := ahasend.CreateMessageRequest{
    Subject: ahasend.PtrString("Email Subject"),
    TextBody: ahasend.PtrString("Email content"),
}
```

**Error Handling**
API errors are returned as `ErrorResponse` models with structured error information.

## Rate Limiting (NEW!)

The SDK now includes intelligent, automatic rate limiting that protects against 429 errors:

### Automatic Protection
- **Enabled by default** - No configuration needed
- **Three endpoint types**: General API (100 req/s), Statistics (1 req/s), Send Message (100 req/s)
- **Smart detection** - Automatically categorizes endpoints by URL pattern
- **Token bucket algorithm** - Allows burst traffic up to configured limits

### Customer Configuration
```go
client := ahasend.NewAPIClient(ahasend.NewConfiguration())

// High-volume email sender
client.SetSendMessageRateLimit(500, 1000)  // 500 req/s, 1000 burst

// Analytics platform
client.SetStatisticsRateLimit(10, 20)      // 10 req/s, 20 burst

// Batch configuration
client.ConfigureCustomerRateLimits(ahasend.CustomerRateLimitConfig{
    SendMessage: &ahasend.RateLimitConfig{RequestsPerSecond: 1000, BurstCapacity: 2000, Enabled: true},
})

// Monitor status
status := client.GetRateLimitStatus(ahasend.SendMessageAPI)
fmt.Printf("Tokens available: %d\n", status.TokensAvailable)
```

### 429 Handling
- **Automatic retries** with exponential backoff + jitter
- **Configurable retry attempts** (default: 3)
- **Non-blocking** - Uses token bucket pre-request limiting

## Webhook Processing (NEW!)

The SDK includes comprehensive webhook processing capabilities that are fully compliant with the [Standard Webhooks specification](https://www.standardwebhooks.com/).

### Core Components

**WebhookVerifier (`webhooks.go`)**
- Creates verifier instances with webhook secrets
- Handles HMAC-SHA256 signature verification
- Validates timestamps to prevent replay attacks
- Supports multiple signatures for key rotation
- Constant-time comparison for security

**Event Models (`webhook_events.go`)**
- Strongly typed models for all AhaSend webhook events
- Type-safe parsing with automatic event detection
- Helper functions for event categorization
- Common data extraction utilities

### Usage Pattern

```go
// Create webhook verifier
verifier, err := ahasend.NewWebhookVerifier("whsec_your_webhook_secret")
if err != nil {
    log.Fatal(err)
}

// In HTTP handler
func webhookHandler(w http.ResponseWriter, r *http.Request) {
    // Parse and verify webhook
    event, err := verifier.ParseRequest(r)
    if err != nil {
        http.Error(w, "Invalid webhook", http.StatusBadRequest)
        return
    }

    // Process event by type
    switch e := event.(type) {
    case *ahasend.MessageDeliveredEvent:
        // Handle delivery confirmation
        log.Printf("Email delivered to %s", e.Data.Recipient)
    case *ahasend.MessageBouncedEvent:
        // Handle bounce
        log.Printf("Email bounced for %s", e.Data.Recipient)
    case *ahasend.MessageOpenedEvent:
        // Handle open tracking
        log.Printf("Email opened by %s", e.Data.Recipient)
    case *ahasend.MessageClickedEvent:
        // Handle click tracking
        log.Printf("Link clicked: %s", e.Data.URL)
    case *ahasend.SuppressionCreatedEvent:
        // Handle suppression
        log.Printf("Email suppressed: %s", e.Data.Recipient)
    case *ahasend.DomainDNSErrorEvent:
        // Handle DNS issues
        log.Printf("DNS error for domain: %s", e.Data.Domain)
    case *ahasend.RouteMessageEvent:
        // Handle inbound emails
        log.Printf("Inbound email from: %s", e.Data.From)
    }

    w.WriteHeader(http.StatusOK)
}
```

### Supported Event Types

**Message Events** (outbound email tracking):
- `message.reception` - Email queued for delivery
- `message.delivered` - Email successfully delivered
- `message.transient_error` - Temporary delivery issue (will retry)
- `message.failed` - Permanent delivery failure
- `message.bounced` - Bounce notification received
- `message.suppressed` - Email not sent (recipient suppressed)
- `message.opened` - Email opened by recipient (includes user agent, IP)
- `message.clicked` - Link clicked in email (includes URL, user agent, IP)

**Suppression Events**:
- `suppression.created` - Email address added to suppression list

**Domain Events**:
- `domain.dns_error` - DNS configuration issues detected

**Route Events** (inbound email processing):
- `route.message` - Inbound email received and processed

### Helper Functions

```go
// Check event categories
if ahasend.IsMessageEvent(event) {
    // Handle message-related event
}
if ahasend.IsSuppressionEvent(event) {
    // Handle suppression event
}

// Extract common message data
if messageData := ahasend.GetMessageEventData(event); messageData != nil {
    log.Printf("Message ID: %s", messageData.ID)
    log.Printf("From: %s", messageData.From)
    log.Printf("To: %s", messageData.Recipient)
}
```

### Security Features

- **HMAC-SHA256** signature verification
- **Constant-time comparison** to prevent timing attacks
- **Timestamp validation** with configurable tolerance (default: 5 minutes)
- **Multiple signature support** for seamless key rotation
- **Standard Webhooks compliance** for interoperability

### Configuration Options

```go
// Custom tolerance for timestamp validation
verifier.SetTolerance(10 * time.Minute)

// The verifier automatically handles:
// - whsec_ prefix removal
// - Multiple signatures in space-delimited format
// - Version prefixes (v1=signature)
```

## Retry Configuration (NEW!)

The SDK now includes sophisticated retry configuration with multiple backoff strategies:

### RetryConfig Structure
```go
type RetryConfig struct {
    Enabled               bool            // Controls whether retries are enabled
    MaxRetries            int             // Maximum number of retry attempts
    RetryClientErrors     bool            // Whether to retry 4xx client errors (default: false)
    RetryValidationErrors bool            // Whether to retry validation errors (default: false)
    BackoffStrategy       BackoffStrategy // Exponential, Linear, or Constant
    BaseDelay             time.Duration   // Initial delay for retries
    MaxDelay              time.Duration   // Maximum delay between retries
}
```

### Backoff Strategies
- **Exponential**: Exponential backoff with jitter (recommended for production)
- **Linear**: Linear increase in delay between retries
- **Constant**: Fixed delay between all retry attempts

### Default Behavior
- Automatically retries: Network errors, 429 rate limits, 5xx server errors
- Never retries (by default): 4xx client errors, validation errors, authentication errors
- Uses exponential backoff with 1-30 second delays
- Backward compatible with legacy `MaxRetries` field

## Important Notes

- **Always use Makefile commands** - They include proper setup, validation, and consistency checks
- **Rate limiting is ON by default** - Protects users from 429 errors immediately
- **Intelligent retry logic is ON by default** - Prevents unnecessary retries and uses smart backoff
- **Webhook processing is Standard Webhooks compliant** - Secure verification with HMAC-SHA256
- The OpenAPI spec at `api/openapi.yaml` provides the foundation for the API structure
- The webhook specification at `api/webhooks.yaml` defines all webhook event schemas
- Use the provided pointer utility functions (`PtrString`, `PtrBool`, etc.) for optional fields
- All timestamps must be in RFC3339 format
- The API uses cursor-based pagination with `limit` and `cursor` parameters
- Idempotency is supported via the `Idempotency-Key` header
- Rate limits: 100 req/sec general, 1 req/sec statistics, 100 req/sec send message (all customizable)
- Retry configuration: Exponential backoff by default, fully customizable strategies
- Webhook events include comprehensive data for email tracking and inbound processing
- Development tools are automatically installed via `make setup` (golangci-lint, Prism, etc.)
- Integration tests require Prism mock server for API validation
- The SDK is actively maintained - all code can be modified and enhanced as needed

## API Scopes

The AhaSend API uses fine-grained scopes for access control. Key scope patterns:
- `{resource}:read:all` - Read all resources of a type
- `{resource}:read:{domain}` - Read resources for a specific domain
- `{resource}:write:all` - Create/update resources
- `{resource}:delete:all` - Delete any resource of a type

Refer to the README.md for the complete list of available scopes.