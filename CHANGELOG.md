# Changelog

All notable changes to the AhaSend Go SDK will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2025-01-08

### 🎉 Initial Release

The official Go SDK for the AhaSend API v2, providing a comprehensive, production-ready solution for transactional email delivery and management.

### Added

#### Core Features
- **Full API Coverage**: Complete implementation of all AhaSend API v2 endpoints
  - Messages API: Send, cancel, and retrieve messages
  - Domains API: Manage email sending domains
  - Webhooks API: Configure webhook endpoints
  - Routes API: Set up email routing rules
  - API Keys API: Manage API keys programmatically
  - SMTP Credentials API: Create and manage SMTP credentials
  - Suppressions API: Manage email suppression lists
  - Statistics API: Access detailed email analytics
  - Accounts API: Account and member management
  - Utility API: Health checks and service status

#### 🛡️ Protection & Reliability
- **Automatic Rate Limiting**: Built-in intelligent rate limiting with token bucket algorithm
  - Prevents 429 errors automatically
  - Three endpoint categories with different limits
  - Customizable limits based on your plan
  - Real-time status monitoring
- **Idempotency Support**: Prevent duplicate operations
  - Automatic key generation for POST requests
  - Manual key specification support
  - Configurable prefix for multi-environment setups
- **Intelligent Retry Logic**: Automatic retries with exponential backoff
  - Configurable retry attempts (default: 3)
  - Smart jitter to prevent thundering herd
- **Context Cancellation**: Full support for Go contexts
  - Request timeouts
  - Graceful cancellation
  - Deadline propagation

#### 🪝 Webhook Processing (Standard Webhooks Compliant)
- **Signature Verification**: Secure webhook validation using HMAC-SHA256
- **Event Parsing**: Strongly typed models for all webhook events
  - Message events (reception, delivered, failed, bounced, opened, clicked, etc.)
  - Suppression events (created)
  - Domain events (DNS errors)
  - Route events (inbound email processing)
- **Helper Functions**: Utilities for event type checking and data extraction
- **Timestamp Validation**: Prevent replay attacks with configurable tolerance
- **Multiple Signature Support**: Handle key rotation seamlessly

#### 🔧 Configuration & Flexibility
- **Environment Variable Support**: 12-Factor App compliant configuration
  - `AHASEND_API_KEY`: API authentication
  - `AHASEND_DEBUG`: Debug logging
  - `AHASEND_MAX_RETRIES`: Retry configuration
  - `AHASEND_ENABLE_RATE_LIMIT`: Rate limiting control
  - `AHASEND_IDEMPOTENCY_PREFIX`: Multi-environment support
  - And many more (see ENV.md)
- **Custom HTTP Client**: Support for proxies and advanced networking
- **Production Optimizations**: Pre-configured settings for high-volume usage

#### 📚 Developer Experience
- **Type Safety**: Strongly typed models for all API operations
- **Thread Safety**: Safe for concurrent use from multiple goroutines
- **Comprehensive Examples**: 12+ real-world examples included
  - Basic email sending
  - Attachments handling
  - Batch operations
  - Webhook processing
  - Error handling patterns
  - Domain management
  - Statistics retrieval
  - Rate limiting configuration
- **Extensive Documentation**:
  - Complete README with quick start guide
  - ENV.md for environment configuration
  - IDEMPOTENCY.md for duplicate prevention
  - CONTRIBUTING.md for contributors
  - In-code documentation for all public APIs

#### 🧪 Quality Assurance
- **Comprehensive Test Suite**: Unit and integration tests
- **OpenAPI Compliance**: Generated from and validated against OpenAPI spec
- **Error Handling**: Structured error types with detailed messages
- **Validation**: Configuration validation and production readiness checks

### Technical Details

- **Go Version**: 1.18+
- **Dependencies**: Minimal external dependencies for maximum stability
- **License**: MIT
- **API Version**: AhaSend API v2
- **Standard Compliance**:
  - Standard Webhooks specification
  - RFC3339 timestamp format
  - OAuth 2.0 Bearer token authentication

### Migration Notes

This is the first stable release of the AhaSend Go SDK. If you're migrating from direct API calls:

1. Install the SDK: `go get github.com/AhaSend/ahasend-go`
2. Replace HTTP client code with SDK client initialization
3. Update API calls to use SDK methods
4. Implement webhook verification using the built-in verifier
5. Enable automatic features (rate limiting, idempotency, retries)

### Examples

#### Quick Start
```go
// Initialize client from environment
client := ahasend.NewAPIClientFromEnv()
ctx := ahasend.ContextWithEnvAuth(context.Background())

// Send an email
message := ahasend.CreateMessageRequest{
    From: ahasend.SenderAddress{
        Email: "hello@yourdomain.com",
        Name:  ahasend.PtrString("Your App"),
    },
    Recipients: []ahasend.Recipient{{
        Email: "user@example.com",
        Name:  ahasend.PtrString("User"),
    }},
    Subject:     "Welcome!",
    HtmlContent: ahasend.PtrString("<h1>Welcome to our service!</h1>"),
}

response, _, err := client.MessagesAPI.
    CreateMessage(ctx, accountID).
    CreateMessageRequest(message).
    Execute()
```

#### Webhook Processing
```go
// Create webhook verifier
verifier, _ := ahasend.NewWebhookVerifier("whsec_your_secret")

// In your HTTP handler
func webhookHandler(w http.ResponseWriter, r *http.Request) {
    event, err := verifier.ParseRequest(r)
    if err != nil {
        http.Error(w, "Invalid webhook", http.StatusBadRequest)
        return
    }

    switch e := event.(type) {
    case *ahasend.MessageDeliveredEvent:
        log.Printf("Email delivered to %s", e.Data.Recipient)
    case *ahasend.MessageBouncedEvent:
        log.Printf("Email bounced for %s", e.Data.Recipient)
    }

    w.WriteHeader(http.StatusOK)
}
```

### Contributors

- AhaSend Team

### Support

- **Documentation**: https://ahasend.com/docs/api-reference
- **Issues**: https://github.com/AhaSend/ahasend-go/issues
- **Email**: support@ahasend.com

---

[1.0.0]: https://github.com/AhaSend/ahasend-go/releases/tag/v1.0.0