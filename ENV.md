# Environment Variable Configuration

The AhaSend Go SDK supports configuration through environment variables, making it easy to deploy your application across different environments without code changes. This follows the [12-Factor App](https://12factor.net/config) methodology for configuration management.

## Table of Contents

- [Quick Start](#quick-start)
- [Supported Environment Variables](#supported-environment-variables)
- [Usage Examples](#usage-examples)
- [Configuration Precedence](#configuration-precedence)
- [Validation and Troubleshooting](#validation-and-troubleshooting)
- [Best Practices](#best-practices)

## Quick Start

The easiest way to get started is to set your API key in an environment variable and use the convenience constructor:

```bash
export AHASEND_API_KEY="aha-sk-your-api-key-here"
```

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/AhaSend/ahasend-go/api"
)

func main() {
    // Create client with configuration loaded from environment variables
    client := api.NewAPIClientFromEnv()

    // Create context with API key from environment variables
    ctx := api.ContextWithEnvAuth(context.Background())

    // Your API calls here...
}
```

## Supported Environment Variables

### API Authentication

| Variable | Description | Example |
|----------|-------------|---------|
| `AHASEND_API_KEY` | Primary API key for authentication | `aha-sk-64-character-string` |
| `AHASEND_TOKEN` | Alternative to `AHASEND_API_KEY` | `aha-sk-64-character-string` |

**Note**: `AHASEND_API_KEY` takes precedence over `AHASEND_TOKEN` if both are set.

**Helper Functions**:
- `api.GetAPIKeyFromEnv()` - Retrieve API key from environment variables
- `api.ContextWithEnvAuth(ctx)` - Create an authenticated context
- `api.ConfigFromEnv()` - Create configuration from environment variables
- `api.NewConfigurationFromEnv()` - Alternative to ConfigFromEnv()
- `api.NewAPIClientFromEnv()` - Create client with environment configuration
- `api.LoadEnvIntoConfig(cfg)` - Load environment variables into existing configuration

### Server Configuration

| Variable | Description | Example | Default |
|----------|-------------|---------|---------|
| `AHASEND_BASE_URL` | Full base URL including scheme | `https://api.ahasend.com` | `https://api.ahasend.com` |
| `AHASEND_HOST` | API host without scheme | `api.staging.ahasend.com` | `api.ahasend.com` |
| `AHASEND_SCHEME` | URL scheme | `https` or `http` | `https` |

**Note**: `AHASEND_BASE_URL` will be parsed to extract host and scheme. Individual `AHASEND_HOST` and `AHASEND_SCHEME` values override the base URL.

### Debug and Logging

| Variable | Description | Accepted Values | Default |
|----------|-------------|----------------|---------|
| `AHASEND_DEBUG` | Enable debug logging | `true`, `false`, `1`, `0`, `yes`, `no`, `on`, `off`, `enable`, `disable` | `false` |
| `AHASEND_USER_AGENT` | Custom user agent string | Any string | `AhaSend-Go-SDK/1.0` |

### Rate Limiting

| Variable | Description | Accepted Values | Default |
|----------|-------------|----------------|---------|
| `AHASEND_ENABLE_RATE_LIMIT` | Enable rate limiting | Boolean values (see Debug section) | `true` |
| `AHASEND_MAX_RETRIES` | Maximum number of retries | Integer ≥ 0 | `3` |

### Timeouts

| Variable | Description | Accepted Values | Default |
|----------|-------------|----------------|---------|
| `AHASEND_TIMEOUT` | Request timeout in seconds | Integer > 0 | Not set |
| `AHASEND_CONNECT_TIMEOUT` | Connection timeout in seconds | Integer > 0 | Not set |

### Idempotency

| Variable | Description | Accepted Values | Default |
|----------|-------------|----------------|---------|
| `AHASEND_IDEMPOTENCY_AUTO_GENERATE` | Auto-generate idempotency keys | Boolean values | `true` |
| `AHASEND_IDEMPOTENCY_PREFIX` | Prefix for generated keys | Any string | Empty |

## Usage Examples

### Basic Usage

```go
// Set environment variables
os.Setenv("AHASEND_API_KEY", "aha-sk-your-key")
os.Setenv("AHASEND_DEBUG", "true")

// Create client from environment
client := api.NewAPIClientFromEnv()
```

### Custom Configuration with Environment Override

```go
// Start with custom configuration
cfg := api.NewConfiguration()
cfg.RetryConfig.MaxRetries = 1

// Load environment variables into existing config
api.LoadEnvIntoConfig(cfg)

// Create client
client := api.NewAPIClient(cfg)
```

### Environment-Specific Setup

#### Development Environment

```bash
# .env.development
export AHASEND_API_KEY="aha-sk-dev-key"
export AHASEND_BASE_URL="https://api.staging.ahasend.com"
export AHASEND_DEBUG=true
export AHASEND_IDEMPOTENCY_PREFIX="dev"
export AHASEND_MAX_RETRIES=1
```

#### Production Environment

```bash
# .env.production
export AHASEND_API_KEY="aha-sk-prod-key"
export AHASEND_DEBUG=false
export AHASEND_MAX_RETRIES=5
export AHASEND_TIMEOUT=30
```

### Docker Usage

```dockerfile
# Dockerfile
FROM golang:1.21

# Set environment variables
ENV AHASEND_DEBUG=false
ENV AHASEND_MAX_RETRIES=3
ENV AHASEND_ENABLE_RATE_LIMIT=true

# API key should be passed at runtime
# ENV AHASEND_API_KEY will be set via docker run -e

COPY . .
RUN go build -o app

CMD ["./app"]
```

```bash
# Run with API key
docker run -e AHASEND_API_KEY="aha-sk-your-key" your-app
```

### Kubernetes ConfigMap and Secret

```yaml
# configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: ahasend-config
data:
  AHASEND_DEBUG: "false"
  AHASEND_MAX_RETRIES: "5"
  AHASEND_ENABLE_RATE_LIMIT: "true"
  AHASEND_IDEMPOTENCY_PREFIX: "k8s"

---
# secret.yaml
apiVersion: v1
kind: Secret
metadata:
  name: ahasend-secret
data:
  AHASEND_API_KEY: <aha-sk-api-key>
```

```yaml
# deployment.yaml
apiVersion: apps/v1
kind: Deployment
spec:
  template:
    spec:
      containers:
      - name: app
        envFrom:
        - configMapRef:
            name: ahasend-config
        - secretRef:
            name: ahasend-secret
```

### Using with Popular Tools

#### With godotenv

```go
import (
    "github.com/joho/godotenv"
    "github.com/AhaSend/ahasend-go/api"
)

func init() {
    // Load .env file
    godotenv.Load()
}

func main() {
    // Environment variables are now loaded
    client := api.NewAPIClientFromEnv()
}
```

#### With Viper

```go
import (
    "github.com/spf13/viper"
    "github.com/AhaSend/ahasend-go/api"
)

func main() {
    viper.AutomaticEnv()
    viper.SetEnvPrefix("AHASEND")

    // Get configuration
    cfg := api.NewConfiguration()

    if viper.IsSet("API_KEY") {
        // Use viper values or fall back to environment
    }

    // Load remaining values from environment
    api.LoadEnvIntoConfig(cfg)

    client := api.NewAPIClient(cfg)
}
```

## Configuration Precedence

Configuration values are applied in the following order (later values override earlier ones):

1. **SDK defaults** (hardcoded values)
2. **Environment variables** (from `ENV.md` supported variables)
3. **Programmatic configuration** (values set in code)

```go
// This demonstrates precedence
cfg := api.NewConfiguration()           // 1. SDK defaults
cfg.Debug = false                       // 1. SDK defaults

api.LoadEnvIntoConfig(cfg)              // 2. Environment variables override defaults
// If AHASEND_DEBUG=true, cfg.Debug is now true

cfg.Debug = false                       // 3. Programmatic override
// cfg.Debug is now false regardless of environment
```

## Validation and Troubleshooting

### Validate Configuration

```go
// Check for configuration issues
issues := api.ValidateEnvConfig()
if len(issues) > 0 {
    for _, issue := range issues {
        fmt.Printf("Configuration issue: %s\n", issue)
    }
}
```

## See Also

- [Configuration Guide](README.md#configuration)
- [Idempotency Documentation](IDEMPOTENCY.md)
- [AhaSend API Documentation](https://ahasend.com/docs)