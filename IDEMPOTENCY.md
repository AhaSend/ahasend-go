# Idempotency and Retry Logic

The AhaSend Go SDK provides robust idempotency support and automatic retry logic to ensure reliable API operations even in the face of network issues or transient failures.

## Table of Contents

- [Idempotency](#idempotency)
  - [Automatic Idempotency](#automatic-idempotency)
  - [Manual Idempotency Keys](#manual-idempotency-keys)
  - [Disabling Automatic Idempotency](#disabling-automatic-idempotency)
  - [Idempotency Key Generation](#idempotency-key-generation)
  - [Advanced Idempotency Usage](#advanced-idempotency-usage)
- [Retry Logic](#retry-logic)
  - [Retryable Conditions](#retryable-conditions)
  - [Retry Configuration](#retry-configuration)
  - [Exponential Backoff](#exponential-backoff)
- [Best Practices](#best-practices)

## Idempotency

Idempotency ensures that API requests can be safely retried without causing duplicate operations. The AhaSend API supports idempotency through the `Idempotency-Key` header on POST requests.

### Automatic Idempotency

By default, the SDK automatically generates and adds idempotency keys to all POST requests:

```go
// Automatic idempotency is enabled by default
cfg := ahasend.NewConfiguration()
client := ahasend.NewAPIClient(cfg)

// This POST request will automatically have an idempotency key
response, _, err := client.MessagesAPI.
    CreateMessage(ctx, accountID).
    CreateMessageRequest(messageReq).
    Execute()
```

The SDK uses UUID v4 for generating unique idempotency keys, ensuring they are:
- Globally unique
- Cryptographically random
- Collision-resistant

### Manual Idempotency Keys

You can override the automatic idempotency key with your own:

```go
// Set a custom idempotency key
response, _, err := client.MessagesAPI.
    CreateMessage(ctx, accountID).
    CreateMessageRequest(messageReq).
    IdempotencyKey("my-custom-key-123").
    Execute()
```

When you provide a manual idempotency key:
- It overrides any automatic key generation
- The same key can be used for retries
- Keys must be unique per logical operation
- Keys expire after 24 hours on the server side

### Disabling Automatic Idempotency

To disable automatic idempotency key generation:

```go
cfg := ahasend.NewConfiguration()
cfg.IdempotencyConfig.AutoGenerate = false
client := ahasend.NewAPIClient(cfg)

// Now POST requests won't have automatic idempotency keys
// You must set them manually when needed
```

### Idempotency Key Generation

The SDK provides utilities for generating idempotency keys:

```go
// Generate a single idempotency key
key := client.GenerateIdempotencyKey()

// Use with prefix for organization
cfg.IdempotencyConfig.KeyPrefix = "msg"
client := ahasend.NewAPIClient(cfg)
// Generated keys will be like: "msg-550e8400-e29b-41d4-a716-446655440000"

// Builder for related operations
builder := client.NewIdempotencyKeyBuilder("batch-123")
key1 := builder.Next() // "batch-123"
key2 := builder.Next() // "batch-123-a1b2c3d4"
key3 := builder.WithSuffix("retry") // "batch-123-retry"
```

### Advanced Idempotency Usage

For complex workflows, you can use the idempotency helper:

```go
// Execute an idempotent operation
result, err := ahasend.ExecuteIdempotent(func(key string) (*ahasend.Message, error) {
    return client.MessagesAPI.
        CreateMessage(ctx, accountID).
        CreateMessageRequest(messageReq).
        IdempotencyKey(key).
        Execute()
})
```

## Retry Logic

The SDK automatically retries requests that fail due to transient errors, ensuring your application remains resilient.

### Retryable Conditions

The SDK will automatically retry requests in the following situations:

1. **Network Errors**: Connection refused, timeouts, DNS failures
2. **HTTP Status Codes**:
   - `429 Too Many Requests` - Rate limiting
   - `502 Bad Gateway` - Proxy/gateway errors
   - `503 Service Unavailable` - Temporary service issues
   - `504 Gateway Timeout` - Upstream timeout

**Note**: The SDK does NOT retry on `500 Internal Server Error` as these often indicate non-transient issues.

### Retry Configuration

Configure retry behavior when creating the client:

```go
cfg := ahasend.NewConfiguration()

// Set maximum number of retries (default: 3)
cfg.MaxRetries = 5

// Disable retries entirely
cfg.MaxRetries = 0

client := ahasend.NewAPIClient(cfg)
```

### Exponential Backoff

The SDK uses exponential backoff with jitter to avoid thundering herd problems:

| Attempt | Base Delay | Jitter Range | Total Wait Time |
|---------|------------|--------------|-----------------|
| 1       | 1s         | 0-1000ms     | 1-2s           |
| 2       | 2s         | 0-1000ms     | 2-3s           |
| 3       | 4s         | 0-1000ms     | 4-5s           |

The formula: `delay = 2^attempt + random(0, 1000ms)`

## Best Practices

### 1. Use Automatic Idempotency

Let the SDK handle idempotency automatically for most use cases:

```go
// Good: Automatic idempotency ensures safety
response, _, err := client.MessagesAPI.CreateMessage(ctx, accountID).
    CreateMessageRequest(messageReq).Execute()
```

### 2. Custom Keys for Business Logic

Use custom idempotency keys when you have specific business requirements:

```go
// Use order ID as idempotency key to prevent duplicate emails
orderID := "order-12345"
response, _, err := client.MessagesAPI.CreateMessage(ctx, accountID).
    CreateMessageRequest(messageReq).
    IdempotencyKey(fmt.Sprintf("order-confirmation-%s", orderID)).
    Execute()
```

### 3. Handle Retry Exhaustion

Always handle the case where all retries are exhausted:

```go
response, httpResp, err := client.MessagesAPI.CreateMessage(ctx, accountID).
    CreateMessageRequest(messageReq).Execute()

if err != nil {
    if httpResp != nil && httpResp.StatusCode >= 500 {
        // Server error after all retries
        log.Printf("Failed after retries: %v", err)
        // Implement your fallback logic
    }
    return err
}
```

### 4. Idempotency for Batch Operations

When processing batches, use related idempotency keys:

```go
builder := client.NewIdempotencyKeyBuilder(fmt.Sprintf("batch-%d", time.Now().Unix()))

for i, recipient := range recipients {
    messageReq.To = []string{recipient}

    response, _, err := client.MessagesAPI.CreateMessage(ctx, accountID).
        CreateMessageRequest(messageReq).
        IdempotencyKey(builder.Next()).
        Execute()

    if err != nil {
        // The same idempotency key will be used if you retry this specific message
        log.Printf("Failed to send to recipient %d: %v", i, err)
    }
}
```

### 5. Monitor Idempotent Replays

The AhaSend API returns special headers for idempotent requests:

- `Idempotent-Replayed: true` - Request was a duplicate and cached response was returned
- `Idempotent-Replayed: false` with 409 - Request is currently being processed
- `Idempotent-Replayed: false` with 412 - Previous request with this key failed

### 6. Testing Idempotency

Test your idempotency logic:

```go
// Send the same request twice with the same key
key := "test-idempotency-" + uuid.New().String()

// First request
resp1, _, err1 := client.MessagesAPI.CreateMessage(ctx, accountID).
    CreateMessageRequest(messageReq).
    IdempotencyKey(key).
    Execute()

// Second request with same key - should return cached response
resp2, httpResp, err2 := client.MessagesAPI.CreateMessage(ctx, accountID).
    CreateMessageRequest(messageReq).
    IdempotencyKey(key).
    Execute()

// Check if it was replayed
if httpResp.Header.Get("Idempotent-Replayed") == "true" {
    fmt.Println("Request was idempotent replay")
}
```

## Configuration Reference

```go
type IdempotencyConfig struct {
    // AutoGenerate automatically generates UUID idempotency keys for POST requests
    // Default: true
    AutoGenerate bool

    // KeyPrefix adds a prefix to all generated idempotency keys
    // Default: "" (no prefix)
    KeyPrefix string
}

type Configuration struct {
    // MaxRetries sets the maximum number of retry attempts
    // Default: 3
    MaxRetries int

    // IdempotencyConfig controls idempotency behavior
    IdempotencyConfig IdempotencyConfig

    // ... other configuration options
}
```

## Troubleshooting

### Duplicate Operations Despite Idempotency

If you're seeing duplicate operations:
1. Ensure you're using the same idempotency key for retries
2. Check that keys aren't expired (24-hour TTL)
3. Verify the API endpoint supports idempotency (currently only POST requests)

### Requests Not Being Retried

If requests aren't being retried:
1. Check `MaxRetries` configuration (must be > 0)
2. Verify the error is retryable (network error or 429/502/503/504 status)
3. Ensure the context hasn't been cancelled

### Performance Impact

Retries with exponential backoff can add latency:
- Consider reducing `MaxRetries` for time-sensitive operations
- Implement circuit breakers for persistent failures
- Use appropriate timeouts on your HTTP client

## See Also

- [AhaSend API Documentation on Idempotency](https://ahasend.com/docs/api-reference/idempotency)
- [Rate Limiting Documentation](https://ahasend.com/docs/api-reference/rate-limits)
- [Configuration Guide](README.md#configuration)
