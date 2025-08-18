# Idempotency in AhaSend Go SDK

The AhaSend Go SDK provides comprehensive idempotency support to ensure safe request retries and prevent duplicate operations. This is especially important for email sending and other critical operations where duplicates can cause issues.

## Table of Contents

- [Overview](#overview)
- [Quick Start](#quick-start)
- [Configuration](#configuration)
- [Basic Usage](#basic-usage)
- [Advanced Features](#advanced-features)
- [API Reference](#api-reference)
- [Best Practices](#best-practices)
- [Examples](#examples)
- [Environment Configuration](#environment-configuration)

## Overview

Idempotency ensures that performing the same operation multiple times has the same effect as performing it once. The AhaSend SDK implements this using:

- **Automatic UUID generation** for idempotency keys
- **Configurable key prefixes** for organization
- **Builder patterns** for related operations
- **Helper utilities** for common scenarios
- **Environment variable configuration**

### Why Idempotency Matters

- **Network Failures**: Safe to retry failed requests
- **Duplicate Prevention**: Avoid sending the same email twice
- **Distributed Systems**: Handle race conditions gracefully
- **User Experience**: Prevent double-charging or duplicate actions

## Quick Start

The simplest way to use idempotency is to let the SDK handle everything automatically:

```go
package main

import (
    "context"
    "log"

    "github.com/AhaSend/ahasend-go/api"
    "github.com/google/uuid"
)

func main() {
    client := api.NewAPIClient(api.NewConfiguration())
    ctx := api.ContextWithEnvAuth(context.Background())
    
    // Create message request
    message := api.CreateMessageRequest{
        From: api.SenderAddress{Email: "sender@yourdomain.com"},
        Recipients: []api.Recipient{
            {Email: "recipient@example.com"},
        },
        Subject:     "Test Email",
        HtmlContent: api.PtrString("<h1>Hello!</h1>"),
    }

    accountID := uuid.MustParse("your-account-id")
    
    // SDK automatically generates idempotency key
    response, httpResp, err := client.MessagesAPI.CreateMessage(ctx, accountID, message)
    if err != nil {
        log.Fatal(err)
    }
    
    // View the auto-generated idempotency key
    log.Printf("Idempotency key: %s", httpResp.Request.Header.Get("Idempotency-Key"))
}
```

## Configuration

### Idempotency Configuration Structure

```go
type IdempotencyConfig struct {
    // AutoGenerate automatically generates UUID idempotency keys when none provided
    AutoGenerate bool
    // KeyPrefix adds a prefix to all generated idempotency keys (optional)
    KeyPrefix string
}
```

### Default Configuration

```go
// Default configuration enables auto-generation with no prefix
config := api.DefaultIdempotencyConfig()
// config.AutoGenerate = true
// config.KeyPrefix = ""
```

### Custom Configuration

```go
// Create custom idempotency configuration
config := api.IdempotencyConfig{
    AutoGenerate: true,
    KeyPrefix:    "myapp",  // Keys will be "myapp-uuid"
}

// Apply to client
client := api.NewAPIClient(api.NewConfiguration())
client.SetIdempotencyConfig(config)
```

## Basic Usage

### Automatic Key Generation

The SDK automatically generates idempotency keys for all requests:

```go
// No idempotency key specified - SDK generates one
response, httpResp, err := client.MessagesAPI.CreateMessage(ctx, accountID, message)

// Check the generated key
generatedKey := httpResp.Request.Header.Get("Idempotency-Key")
log.Printf("Auto-generated key: %s", generatedKey)
```

### Manual Key Specification

You can provide your own idempotency keys:

```go
// Generate your own key
customKey := "user-123-email-welcome"

// Add to context
ctx = context.WithValue(ctx, api.ContextIdempotencyKey, customKey)

// Make request with custom key
response, _, err := client.MessagesAPI.CreateMessage(ctx, accountID, message)
```

### Using Helper Functions

```go
// Generate a simple UUID key
key := api.GenerateIdempotencyKey()
// Example: "550e8400-e29b-41d4-a716-446655440000"

// Generate with prefix
prefixedKey := api.GenerateIdempotencyKeyWithPrefix("welcome")
// Example: "welcome-550e8400-e29b-41d4-a716-446655440000"
```

## Advanced Features

### Idempotency Key Builder

For related operations that need coordinated idempotency keys:

```go
// Create builder with base key
builder := api.NewIdempotencyKeyBuilder("user-123-onboarding")

// First operation uses the base key
welcomeKey := builder.Next()  // "user-123-onboarding"

// Subsequent operations get unique variants
followupKey := builder.Next() // "user-123-onboarding-a1b2c3d4"
reminderKey := builder.Next() // "user-123-onboarding-e5f6g7h8"

// Create keys with specific suffixes
verifyKey := builder.WithSuffix("verify")   // "user-123-onboarding-verify"
completeKey := builder.WithSuffix("complete") // "user-123-onboarding-complete"
```

### Idempotency Helper

For more complex scenarios with configuration:

```go
// Create helper with custom configuration
config := api.IdempotencyConfig{
    AutoGenerate: true,
    KeyPrefix:    "batch",
}
helper := api.NewIdempotencyHelper(config)

// Generate keys with the helper
key1 := helper.GenerateKey() // "batch-550e8400-e29b-41d4-a716-446655440000"

// Ensure key exists (generate if empty, return if provided)
key2 := helper.EnsureKey("")        // Generates new key
key3 := helper.EnsureKey("custom")  // Returns "custom"
```

### Generic Idempotent Operations

Execute operations with automatic idempotency:

```go
// Define an idempotent operation
sendEmail := func(key string) (*api.CreateMessageResponse, error) {
    ctx := context.WithValue(context.Background(), api.ContextIdempotencyKey, key)
    ctx = api.ContextWithEnvAuth(ctx)
    
    return client.MessagesAPI.CreateMessage(ctx, accountID, message)
}

// Execute with automatic key generation
response, err := api.ExecuteIdempotent(sendEmail)

// Execute with custom key
response, err = api.ExecuteIdempotent(sendEmail, "custom-key")
```

## API Reference

### Core Functions

```go
// Generate a new UUID-based idempotency key
func GenerateIdempotencyKey() string

// Generate key with prefix
func GenerateIdempotencyKeyWithPrefix(prefix string) string

// Execute an operation with idempotency
func ExecuteIdempotent[T any](op IdempotentOperation[T], customKey ...string) (T, error)
```

### Configuration

```go
// Get default configuration
func DefaultIdempotencyConfig() IdempotencyConfig

// Create helper with optional configuration
func NewIdempotencyHelper(config ...IdempotencyConfig) *IdempotencyHelper
```

### Builder Pattern

```go
// Create builder with optional base key
func NewIdempotencyKeyBuilder(baseKey string) *IdempotencyKeyBuilder

// Get next key in sequence
func (b *IdempotencyKeyBuilder) Next() string

// Get key with specific suffix
func (b *IdempotencyKeyBuilder) WithSuffix(suffix string) string
```

### Helper Methods

```go
// Generate key based on helper configuration
func (h *IdempotencyHelper) GenerateKey() string

// Ensure key exists (generate if empty when auto-generation enabled)
func (h *IdempotencyHelper) EnsureKey(key string) string
```

### Client Integration

```go
// Generate idempotency key using client configuration
func (c *APIClient) GenerateIdempotencyKey() string

// Get current idempotency configuration
func (c *APIClient) GetIdempotencyConfig() IdempotencyConfig

// Update idempotency configuration
func (c *APIClient) SetIdempotencyConfig(config IdempotencyConfig)

// Create builder with client configuration
func (c *APIClient) NewIdempotencyKeyBuilder(baseKey ...string) *IdempotencyKeyBuilder
```

## Best Practices

### 1. Use Meaningful Keys for Important Operations

```go
// ❌ Generic keys are hard to debug
key := api.GenerateIdempotencyKey()

// ✅ Descriptive keys help with monitoring and debugging
key := api.GenerateIdempotencyKeyWithPrefix("user-signup-welcome")
```

### 2. Use Builders for Related Operations

```go
// ✅ Related operations with coordinated keys
builder := api.NewIdempotencyKeyBuilder("order-12345")

confirmationEmail := builder.WithSuffix("confirmation")
receiptEmail := builder.WithSuffix("receipt")
shippingEmail := builder.WithSuffix("shipping")
```

### 3. Configure Prefixes for Different Environments

```go
// ✅ Environment-specific prefixes
var prefix string
switch os.Getenv("ENVIRONMENT") {
case "production":
    prefix = "prod"
case "staging":
    prefix = "staging"
default:
    prefix = "dev"
}

config := api.IdempotencyConfig{
    AutoGenerate: true,
    KeyPrefix:    prefix,
}
```

### 4. Store Keys for Audit and Debugging

```go
// ✅ Log idempotency keys for troubleshooting
response, httpResp, err := client.MessagesAPI.CreateMessage(ctx, accountID, message)
if err != nil {
    key := httpResp.Request.Header.Get("Idempotency-Key")
    log.Printf("Failed email send with idempotency key: %s", key)
    return err
}
```

### 5. Use Context for Request-Specific Keys

```go
// ✅ Pass keys through context for clean code
func sendWelcomeEmail(userID string, email string) error {
    key := fmt.Sprintf("welcome-%s-%d", userID, time.Now().Unix())
    ctx := context.WithValue(context.Background(), api.ContextIdempotencyKey, key)
    ctx = api.ContextWithEnvAuth(ctx)
    
    _, _, err := client.MessagesAPI.CreateMessage(ctx, accountID, message)
    return err
}
```

## Examples

### User Onboarding Flow

```go
func onboardUser(userID string, email string) error {
    // Create builder for related emails
    builder := api.NewIdempotencyKeyBuilder(fmt.Sprintf("onboard-%s", userID))
    
    ctx := api.ContextWithEnvAuth(context.Background())
    
    // Welcome email
    welcomeCtx := context.WithValue(ctx, api.ContextIdempotencyKey, builder.WithSuffix("welcome"))
    _, _, err := client.MessagesAPI.CreateMessage(welcomeCtx, accountID, welcomeMessage)
    if err != nil {
        return fmt.Errorf("failed to send welcome email: %w", err)
    }
    
    // Verification email  
    verifyCtx := context.WithValue(ctx, api.ContextIdempotencyKey, builder.WithSuffix("verify"))
    _, _, err = client.MessagesAPI.CreateMessage(verifyCtx, accountID, verificationMessage)
    if err != nil {
        return fmt.Errorf("failed to send verification email: %w", err)
    }
    
    return nil
}
```

### Batch Email Processing

```go
func processBatchEmails(batchID string, emails []EmailRequest) {
    helper := api.NewIdempotencyHelper(api.IdempotencyConfig{
        AutoGenerate: true,
        KeyPrefix:    fmt.Sprintf("batch-%s", batchID),
    })
    
    for i, email := range emails {
        // Create unique key for each email in the batch
        key := helper.GenerateKey()
        
        ctx := context.WithValue(context.Background(), api.ContextIdempotencyKey, key)
        ctx = api.ContextWithEnvAuth(ctx)
        
        go func(e EmailRequest, k string) {
            _, _, err := client.MessagesAPI.CreateMessage(ctx, accountID, e.ToMessage())
            if err != nil {
                log.Printf("Failed to send email %s with key %s: %v", e.Recipient, k, err)
            }
        }(email, key)
    }
}
```

### Retry with Same Key

```go
func sendEmailWithRetry(message api.CreateMessageRequest, maxRetries int) error {
    // Generate key once for all retry attempts
    key := api.GenerateIdempotencyKeyWithPrefix("retry")
    
    ctx := context.WithValue(context.Background(), api.ContextIdempotencyKey, key)
    ctx = api.ContextWithEnvAuth(ctx)
    
    var lastErr error
    for attempt := 1; attempt <= maxRetries; attempt++ {
        _, _, err := client.MessagesAPI.CreateMessage(ctx, accountID, message)
        if err == nil {
            log.Printf("Email sent successfully on attempt %d with key %s", attempt, key)
            return nil
        }
        
        lastErr = err
        log.Printf("Attempt %d failed with key %s: %v", attempt, key, err)
        
        // Wait before retry
        time.Sleep(time.Duration(attempt) * time.Second)
    }
    
    return fmt.Errorf("failed after %d attempts with key %s: %w", maxRetries, key, lastErr)
}
```

### Custom Key Generation Strategy

```go
// Custom strategy for deterministic keys based on content
func generateContentBasedKey(subject, recipient string) string {
    h := sha256.Sum256([]byte(subject + recipient + time.Now().Format("2006-01-02")))
    return fmt.Sprintf("content-%x", h[:8]) // Use first 8 bytes of hash
}

func sendDailyNewsletter(subject, recipient string) error {
    // Same content to same recipient on same day = same key
    key := generateContentBasedKey(subject, recipient)
    
    ctx := context.WithValue(context.Background(), api.ContextIdempotencyKey, key)
    ctx = api.ContextWithEnvAuth(ctx)
    
    message := api.CreateMessageRequest{
        // ... newsletter content
    }
    
    _, _, err := client.MessagesAPI.CreateMessage(ctx, accountID, message)
    return err
}
```

## Environment Configuration

You can configure idempotency behavior using environment variables:

### Environment Variables

```bash
# Enable/disable automatic key generation
export AHASEND_IDEMPOTENCY_AUTO_GENERATE=true

# Set prefix for all generated keys
export AHASEND_IDEMPOTENCY_PREFIX=myapp
```

### Using Environment Configuration

```go
// Configuration automatically loaded from environment
client := api.NewAPIClientFromEnv()

// Or load into existing configuration
config := api.NewConfiguration()
api.LoadEnvIntoConfig(config)

client := api.NewAPIClient(config)
```

### Environment Examples

```bash
# Development environment
export AHASEND_IDEMPOTENCY_PREFIX=dev
export AHASEND_IDEMPOTENCY_AUTO_GENERATE=true

# Production environment  
export AHASEND_IDEMPOTENCY_PREFIX=prod
export AHASEND_IDEMPOTENCY_AUTO_GENERATE=true

# Testing environment (disable auto-generation for predictable keys)
export AHASEND_IDEMPOTENCY_AUTO_GENERATE=false
```

## Troubleshooting

### Common Issues

#### 1. Duplicate Requests with Different Keys

```
Problem: Same operation being performed multiple times with different keys
```

**Solution**: Ensure you're using the same key for retries:

```go
// ❌ Don't generate new keys for retries
for i := 0; i < retries; i++ {
    key := api.GenerateIdempotencyKey() // New key each time!
    // This will not prevent duplicates
}

// ✅ Use same key for all retry attempts
key := api.GenerateIdempotencyKey()
for i := 0; i < retries; i++ {
    // Same key ensures idempotency
}
```

#### 2. Missing Idempotency Headers

```
Problem: Requests not including idempotency keys
```

**Solution**: Verify auto-generation is enabled:

```go
config := client.GetIdempotencyConfig()
if !config.AutoGenerate {
    log.Println("Auto-generation is disabled!")
}
```

#### 3. Key Conflicts in Distributed Systems

```
Problem: Different services generating the same keys
```

**Solution**: Use service-specific prefixes:

```go
config := api.IdempotencyConfig{
    AutoGenerate: true,
    KeyPrefix:    "email-service", // Unique per service
}
```

### Debug Information

```go
// Check if key was auto-generated
response, httpResp, err := client.MessagesAPI.CreateMessage(ctx, accountID, message)
key := httpResp.Request.Header.Get("Idempotency-Key")

if key != "" {
    log.Printf("Request used idempotency key: %s", key)
} else {
    log.Println("Warning: No idempotency key was set!")
}
```

## See Also

- [Environment Configuration](ENV.md) - Configure idempotency via environment variables
- [Configuration Guide](README.md#configuration) - General SDK configuration
- [Examples Directory](examples/) - Complete working examples
- [AhaSend API Documentation](https://ahasend.com/docs) - API-level idempotency documentation