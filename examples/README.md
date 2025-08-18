# AhaSend Go SDK Examples

This directory contains example code demonstrating how to use the AhaSend Go SDK for various common tasks.

## Prerequisites

Before running these examples, you need:

1. An AhaSend account and API key
2. Go 1.18 or later installed
3. Set your API credentials as environment variables:
   ```bash
   export AHASEND_API_KEY="your-api-key-here"
   export AHASEND_ACCOUNT_ID="your-account-id-here"
   ```

## Available Examples

### Basic Operations
- [send_email.go](./send_email.go) - Send a simple transactional email  
- [send_with_attachments.go](./send_with_attachments.go) - Send email with file attachments

### Advanced Features
- [batch_send.go](./batch_send.go) - Send emails to multiple recipients efficiently
- [scheduled_send.go](./scheduled_send.go) - Schedule emails for future delivery
- [idempotency.go](./idempotency.go) - Ensure emails are sent only once using idempotency keys

### Domain Management
- [domain_management.go](./domain_management.go) - Add, verify, and manage sending domains

### Webhook Management
- [webhook_management.go](./webhook_management.go) - Set up and manage webhooks for email events
- [webhook_processing.go](./webhook_processing.go) - Complete webhook receiver server with event processing

### Error Handling & Retries
- [error_handling.go](./error_handling.go) - Proper error handling and retry strategies
- [rate_limiting.go](./rate_limiting.go) - Configure and handle rate limiting

### Statistics & Analytics
- [statistics.go](./statistics.go) - Retrieve delivery statistics and analytics

## Running the Examples

Each example file contains a complete, runnable program. To run any example:

```bash
go run examples/send_email.go
```

Or build and run:

```bash
go build -o send_email examples/send_email.go
./send_email
```

### Build Constraints

All example files use `//go:build ignore` build constraints, which means:

- ✅ **Individual examples can be run** with `go run examples/filename.go`
- ✅ **The main package builds cleanly** without conflicts from multiple `main()` functions
- ✅ **Examples are excluded from normal builds** like `go build ./...`
- ✅ **IDE tooling works properly** without compilation errors

This is the standard Go pattern for executable example files that should be runnable but not part of the main package build.

## Special Examples

### Webhook Processing
The `webhook_processing.go` example is a complete webhook receiver server that:

- Verifies webhook signatures using HMAC-SHA256
- Handles all AhaSend webhook event types 
- Includes proper error handling and security measures
- Can be run as a standalone server

Set the webhook secret environment variable:
```bash
export AHASEND_WEBHOOK_SECRET="your-webhook-secret-here"
```

Run the webhook server:
```bash
go run examples/webhook_processing.go
```

## Important Notes

- Replace placeholder values (API keys, account IDs, email addresses) with your actual values
- Some examples require specific account permissions or features
- Check the AhaSend API documentation for detailed information about each endpoint
- The SDK includes automatic rate limiting to prevent 429 errors
- All examples follow Go best practices and include proper error handling

## Need Help?

- [AhaSend API Documentation](https://ahasend.com/docs)
- [Go SDK Documentation](https://pkg.go.dev/github.com/AhaSend/ahasend-go)
- [Support](https://ahasend.com/support)