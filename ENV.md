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

    "github.com/AhaSend/ahasend-go"
)

func main() {
    // Create client with configuration loaded from environment variables
    client := ahasend.NewAPIClientFromEnv()

    // Create context with API key from environment variables
    ctx := ahasend.ContextWithEnvAuth(context.Background())

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
client := ahasend.NewAPIClientFromEnv()
```

### Custom Configuration with Environment Override

```go
// Start with custom configuration
cfg := ahasend.NewConfiguration()
cfg.MaxRetries = 1

// Load environment variables into existing config
ahasend.LoadEnvIntoConfig(cfg)

// Create client
client := ahasend.NewAPIClient(cfg)
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
    "github.com/AhaSend/ahasend-go"
)

func init() {
    // Load .env file
    godotenv.Load()
}

func main() {
    // Environment variables are now loaded
    client := ahasend.NewAPIClientFromEnv()
}
```

#### With Viper

```go
import (
    "github.com/spf13/viper"
    "github.com/AhaSend/ahasend-go"
)

func main() {
    viper.AutomaticEnv()
    viper.SetEnvPrefix("AHASEND")

    // Get configuration
    cfg := ahasend.NewConfiguration()

    if viper.IsSet("API_KEY") {
        // Use viper values or fall back to environment
    }

    // Load remaining values from environment
    ahasend.LoadEnvIntoConfig(cfg)

    client := ahasend.NewAPIClient(cfg)
}
```

## Configuration Precedence

Configuration values are applied in the following order (later values override earlier ones):

1. **SDK defaults** (hardcoded values)
2. **Environment variables** (from `ENV.md` supported variables)
3. **Programmatic configuration** (values set in code)

```go
// This demonstrates precedence
cfg := ahasend.NewConfiguration()           // 1. SDK defaults
cfg.Debug = false                          // 1. SDK defaults

ahasend.LoadEnvIntoConfig(cfg)             // 2. Environment variables override defaults
// If AHASEND_DEBUG=true, cfg.Debug is now true

cfg.Debug = false                          // 3. Programmatic override
// cfg.Debug is now false regardless of environment
```

## Validation and Troubleshooting

### Validate Configuration

```go
// Check for configuration issues
issues := ahasend.ValidateEnvConfig()
if len(issues) > 0 {
    for _, issue := range issues {
        fmt.Printf("Configuration issue: %s\n", issue)
    }
}
```

### Get Documentation

```go
// Get list of all supported environment variables
docs := ahasend.GetEnvDocumentation()
for envVar, description := range docs {
    fmt.Printf("%s: %s\n", envVar, description)
}
```

### Common Issues

#### 1. Missing API Key

```
Error: No API key found in AHASEND_API_KEY or AHASEND_TOKEN
```

**Solution**: Set one of the API key environment variables:
```bash
export AHASEND_API_KEY="aha-sk-your-key"
```

#### 2. Invalid Boolean Values

```
Error: AHASEND_DEBUG must be a valid boolean (true/false, 1/0, yes/no)
```

**Solution**: Use accepted boolean values:
```bash
export AHASEND_DEBUG=true    # ✅ Good
export AHASEND_DEBUG=maybe   # ❌ Bad
```

#### 3. Invalid Scheme

```
Error: AHASEND_SCHEME must be 'http' or 'https'
```

**Solution**: Use valid scheme:
```bash
export AHASEND_SCHEME=https  # ✅ Good
export AHASEND_SCHEME=ftp    # ❌ Bad
```

### Debug Configuration Loading

```go
// Enable debug mode to see configuration loading
os.Setenv("AHASEND_DEBUG", "true")
cfg := ahasend.ConfigFromEnv()

// Check what values were loaded
fmt.Printf("Debug enabled: %v\n", cfg.Debug)
fmt.Printf("Max retries: %d\n", cfg.MaxRetries)
fmt.Printf("Host: %s\n", cfg.Host)
```

## Best Practices

### Security

1. **Never hardcode API keys** in your source code
2. **Use environment variables** for sensitive configuration
3. **Use secrets management** in production (K8s secrets, AWS Secrets Manager, etc.)

```go
// ❌ Don't do this
cfg := ahasend.NewConfiguration()
cfg.DefaultHeader["Authorization"] = "Bearer aha-sk-hardcoded-key"

// ✅ Do this instead
client := ahasend.NewAPIClientFromEnv()
ctx := ahasend.ContextWithEnvAuth(context.Background())
```

### Environment Separation

1. **Use different API keys** for different environments
2. **Use environment-specific configurations**
3. **Validate configuration** at startup

```go
func main() {
    // Validate configuration at startup
    if issues := ahasend.ValidateEnvConfig(); len(issues) > 0 {
        log.Fatalf("Configuration errors: %v", issues)
    }

    client := ahasend.NewAPIClientFromEnv()

    // Your application logic
}
```

### Development Workflow

1. **Use `.env` files** for local development (with godotenv)
2. **Document required variables** in your README
3. **Provide example configurations**

```bash
# .env.example (commit this to version control)
AHASEND_API_KEY=aha-sk-your-development-key-here
AHASEND_DEBUG=true
AHASEND_IDEMPOTENCY_PREFIX=dev
```

### Production Deployment

1. **Use container environment variables**
2. **Set appropriate timeouts** and retry counts
3. **Disable debug logging** in production

```bash
# Production environment variables
export AHASEND_API_KEY="aha-sk-prod-key"
export AHASEND_DEBUG=false
export AHASEND_MAX_RETRIES=5
export AHASEND_TIMEOUT=30
export AHASEND_ENABLE_RATE_LIMIT=true
```

### Monitoring and Debugging

1. **Log configuration issues** at startup
2. **Monitor API key validity**
3. **Track retry rates** and adjust `AHASEND_MAX_RETRIES`

```go
func main() {
    // Validate and log configuration
    if apiKey := ahasend.GetAPIKeyFromEnv(); apiKey == "" {
        log.Fatal("AHASEND_API_KEY is required")
    }

    cfg := ahasend.ConfigFromEnv()
    log.Printf("Using host: %s, debug: %v, max_retries: %d",
        cfg.Host, cfg.Debug, cfg.MaxRetries)

    client := ahasend.NewAPIClient(cfg)
}
```

## Integration Examples

### Web Server

```go
package main

import (
    "context"
    "net/http"

    "github.com/AhaSend/ahasend-go"
)

var ahasendClient *ahasend.APIClient

func init() {
    // Initialize AhaSend client from environment
    ahasendClient = ahasend.NewAPIClientFromEnv()
}

func sendEmailHandler(w http.ResponseWriter, r *http.Request) {
    ctx := ahasend.ContextWithEnvAuth(r.Context())

    // Use the client to send email
    // ...
}
```

### Background Worker

```go
package main

import (
    "context"
    "time"

    "github.com/AhaSend/ahasend-go"
)

func main() {
    client := ahasend.NewAPIClientFromEnv()
    ctx := ahasend.ContextWithEnvAuth(context.Background())

    // Process messages every minute
    ticker := time.NewTicker(time.Minute)
    defer ticker.Stop()

    for range ticker.C {
        // Process email queue
        processEmailQueue(client, ctx)
    }
}
```

### CLI Application

```go
package main

import (
    "flag"
    "os"

    "github.com/AhaSend/ahasend-go"
)

func main() {
    var apiKey = flag.String("api-key", "", "AhaSend API key (or set AHASEND_API_KEY)")
    flag.Parse()

    // Use CLI flag or fall back to environment
    if *apiKey != "" {
        os.Setenv("AHASEND_API_KEY", *apiKey)
    }

    client := ahasend.NewAPIClientFromEnv()

    // CLI logic here...
}
```

## See Also

- [Configuration Guide](README.md#configuration)
- [Idempotency Documentation](IDEMPOTENCY.md)
- [AhaSend API Documentation](https://ahasend.com/docs)