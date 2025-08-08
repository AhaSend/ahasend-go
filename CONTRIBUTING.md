# Contributing to AhaSend Go SDK

Thank you for your interest in contributing to the AhaSend Go SDK! This document provides comprehensive guidelines and information for contributors. We welcome all types of contributions that help improve the SDK for the Go community.

## 🎯 Types of Contributions

We welcome all types of contributions to improve the SDK:

### ✅ **Code Contributions**
- **🐛 Bug fixes** - Fix issues with SDK functionality, API compatibility, or edge cases
- **✨ Feature enhancements** - Add new features or improve existing ones
- **⚡ Performance improvements** - Optimize SDK performance, memory usage, or network efficiency
- **🛡️ Error handling** - Better error types, handling patterns, and user experience
- **🧪 Testing** - Unit tests, integration tests, benchmarks, and testing utilities
- **🔧 Code quality** - Refactoring, linting improvements, type safety enhancements

### 📚 **Documentation Contributions**
- **📖 Documentation improvements** - API docs, guides, tutorials, and inline comments
- **💡 Examples** - Real-world usage examples and best practices
- **🚀 Getting started guides** - Onboarding improvements and quick-start tutorials
- **📋 README enhancements** - Better explanations, formatting, and organization

### 🛠️ **Infrastructure Contributions**
- **🔄 CI/CD improvements** - GitHub Actions, testing automation, release workflows
- **📦 Build and tooling** - Development scripts, linters, formatters, build optimizations
- **🌐 Developer experience** - Configuration management, logging, debugging tools
- **📊 Monitoring and observability** - Metrics, tracing, and debugging capabilities

## 🚀 Getting Started

### Prerequisites

Ensure you have the following tools installed:
- **Go 1.19+** (latest stable version recommended)
- **Git** (for version control)
- **Make** (optional, for convenient build commands)

**Optional but recommended:**
- **Node.js 16+** and **npm** (for integration testing with Prism mock server)
- **golangci-lint** (for comprehensive code linting)
- **gocover-cobertura** (for code coverage reports)

### 🔧 Development Setup

1. **Fork and clone the repository**
   ```bash
   # Fork on GitHub first, then clone your fork
   git clone https://github.com/YOUR_USERNAME/ahasend-go.git
   cd ahasend-go

   # Add upstream remote
   git remote add upstream https://github.com/AhaSend/ahasend-go.git
   ```

2. **Install dependencies**
   ```bash
   # Install Go dependencies
   go mod tidy

   # Install development tools (optional)
   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

   # Install Prism for integration testing (optional)
   npm install -g @stoplight/prism-cli
   ```

3. **Verify your setup**
   ```bash
   # Build the SDK
   go build .

   # Run basic tests
   SKIP_INTEGRATION_TESTS=true go test -v ./test/

   # Check code quality
   go vet ./...
   go fmt ./...
   ```

### 🧪 Running Tests

The SDK has comprehensive test coverage with multiple types of tests:

```bash
# Run all tests (recommended for contributors)
make test

# Or manually run specific test types:

# 1. Unit tests only (fast, no external dependencies)
SKIP_INTEGRATION_TESTS=true go test -v ./test/

# 2. Integration tests with Prism mock server
go test -v ./test/

# 3. Run tests with coverage
go test -cover ./test/

# 4. Run specific test files
go test -v ./test/client_test.go
go test -v ./test/rate_limiter_test.go

# 5. Benchmark tests
go test -bench=. ./test/
```

### 🔍 Code Quality Checks

Ensure your code meets our quality standards:

```bash
# Format code (required)
go fmt ./...

# Lint code (recommended)
golangci-lint run

# Vet code for issues (required)
go vet ./...

# Check for race conditions in concurrent code
go test -race ./test/

# Check module dependencies
go mod tidy
go mod verify
```

## 📝 Contribution Guidelines

### 🐛 Reporting Issues

When reporting bugs or requesting features, please provide detailed information:

**For Bug Reports:**
```markdown
**Environment:**
- Go version: `go version`
- OS: macOS/Linux/Windows + version
- SDK version: v1.x.x
- Architecture: amd64/arm64

**Description:**
Brief description of the issue

**Steps to Reproduce:**
1. Step one
2. Step two
3. Step three

**Expected Behavior:**
What you expected to happen

**Actual Behavior:**
What actually happened

**Code Example:**
```go
// Minimal reproducible example
package main
// ...
```

**Error Output:**
```
Paste any error messages or logs here
```
```

**For Feature Requests:**
- **Use case**: Describe your specific need
- **Proposed solution**: How you envision it working
- **Alternatives considered**: Other approaches you've tried
- **Additional context**: Any other relevant information

### 🚀 Submitting Changes

#### 1. **Prepare Your Changes**

```bash
# Stay up to date with upstream
git checkout main
git pull upstream main

# Create a descriptive feature branch
git checkout -b feature/add-webhook-validation
# or
git checkout -b fix/rate-limiter-race-condition
# or
git checkout -b docs/improve-getting-started
```

#### 2. **Development Process**

**Follow our standards:**
- ✅ **Code style**: Use `go fmt` and follow Go conventions
- ✅ **Testing**: Add comprehensive tests for new functionality
- ✅ **Documentation**: Update docs, examples, and inline comments
- ✅ **Backwards compatibility**: Don't break existing APIs unless necessary
- ✅ **Performance**: Consider performance implications of your changes

**For different types of contributions:**

```bash
# 🐛 Bug fixes
git checkout -b fix/description-of-bug

# ✨ New features
git checkout -b feature/description-of-feature

# 📚 Documentation
git checkout -b docs/improve-section-name

# 🧪 Tests
git checkout -b test/add-coverage-for-feature

# ⚡ Performance improvements
git checkout -b perf/optimize-rate-limiter

# 🔧 Refactoring
git checkout -b refactor/simplify-error-handling
```

#### 3. **Testing Your Changes**

Before submitting, ensure all tests pass:

```bash
# 1. Code quality checks (required)
go fmt ./...
go vet ./...
golangci-lint run  # if available

# 2. Run the full test suite
SKIP_INTEGRATION_TESTS=true go test -v ./test/

# 3. Integration tests (optional but recommended)
go test -v ./test/

# 4. Test race conditions (for concurrent code)
go test -race ./test/

# 5. Build examples to ensure they still work
go build examples/send_email.go
```

#### 4. **Commit Your Changes**

Use clear, descriptive commit messages following [Conventional Commits](https://www.conventionalcommits.org/):

```bash
# Format: <type>(<scope>): <description>

# Examples:
git commit -m "feat: add webhook signature validation"
git commit -m "fix: resolve race condition in rate limiter"
git commit -m "docs: improve getting started guide with examples"
git commit -m "test: add coverage for message sending edge cases"
git commit -m "refactor: simplify error handling in client"
git commit -m "perf: optimize memory allocation in batch sending"

# For breaking changes:
git commit -m "feat!: change rate limit configuration API"
```

**Commit Types:**
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `test`: Adding or fixing tests
- `refactor`: Code refactoring without functional changes
- `perf`: Performance improvements
- `style`: Code style changes (formatting, etc.)
- `chore`: Build process or auxiliary tool changes

#### 5. **Submit a Pull Request**

Create a well-structured pull request:

**PR Title Format:**
```
<type>(<scope>): <description>

Example: feat(rate-limiter): add custom rate limit configuration per endpoint
```

**PR Description Template:**
```markdown
## Summary
Brief description of the changes and motivation.

## Type of Change
- [ ] 🐛 Bug fix
- [ ] ✨ New feature
- [ ] 📚 Documentation update
- [ ] 🧪 Test improvement
- [ ] ⚡ Performance improvement
- [ ] 🔧 Code refactoring

## Changes Made
- List the specific changes
- Include any breaking changes
- Mention new dependencies if any

## Testing
- [ ] Unit tests pass
- [ ] Integration tests pass (if applicable)
- [ ] Manual testing performed
- [ ] No race conditions detected

## Documentation
- [ ] README updated (if needed)
- [ ] Examples updated (if needed)
- [ ] Inline documentation added
- [ ] CHANGELOG updated (for significant changes)

## Checklist
- [ ] Code follows Go conventions and passes `go fmt`
- [ ] Code passes `go vet` without warnings
- [ ] Tests have been added or updated
- [ ] All tests pass
- [ ] Documentation has been updated
- [ ] Commit messages follow conventional format
```

### 🎯 Code Standards and Best Practices

#### **Go Code Standards**
- **Formatting**: All code must pass `go fmt`
- **Linting**: Code should pass `go vet` and `golangci-lint`
- **Naming**: Follow Go naming conventions (PascalCase for exports, camelCase for internal)
- **Error handling**: Always handle errors appropriately, use proper Go error patterns
- **Documentation**: All exported functions, types, and constants must have godoc comments
- **Testing**: Aim for high test coverage, especially for new functionality

#### **Architecture Guidelines**
- **Thread safety**: All public APIs must be thread-safe
- **Context support**: All long-running operations must accept and respect context.Context
- **Backwards compatibility**: Don't break existing APIs without major version bump
- **Rate limiting**: Respect and work within the SDK's rate limiting system
- **Error types**: Use structured error types that provide useful information

#### **Performance Considerations**
- **Memory allocation**: Minimize unnecessary allocations
- **Network efficiency**: Batch operations where possible
- **Concurrent safety**: Use appropriate synchronization primitives
- **Resource cleanup**: Ensure proper cleanup of resources (connections, goroutines)

#### **Testing Standards**
- **Unit tests**: Test individual components in isolation
- **Integration tests**: Test real API interactions (with mocks)
- **Table-driven tests**: Use table-driven tests for multiple scenarios
- **Edge cases**: Test error conditions and edge cases
- **Race detection**: Run tests with `-race` flag for concurrent code

## 🏗️ Project Structure

Understanding the SDK architecture helps you contribute effectively:

### **Core SDK Files**
```
├── api_*.go              # API service implementations (MessagesAPI, DomainsAPI, etc.)
├── model_*.go            # Data models and structures (Message, Domain, etc.)
├── client.go             # Main HTTP client and configuration
├── configuration.go      # SDK configuration options and validation
├── rate_limiter.go       # Rate limiting implementation
├── idempotency.go        # Idempotency key management
├── errors.go             # Error types and handling
├── env.go                # Environment variable configuration
├── utils.go              # Utility functions and helpers
└── response.go           # HTTP response handling logic
```

### **Supporting Files**
```
├── test/                 # Comprehensive test suite
│   ├── *_test.go        # Unit and integration tests
│   └── integration_*.go # Integration test helpers
├── examples/            # Usage examples and tutorials
│   ├── send_email.go    # Basic email sending
│   ├── batch_send.go    # Batch operations
│   └── *.go             # Other real-world examples
├── docs/                # Documentation files
│   ├── ENV.md           # Environment variables guide
│   ├── IDEMPOTENCY.md   # Idempotency documentation
│   └── *.md             # Other guides
└── .github/             # CI/CD and issue templates
```

## 🛠️ Development Workflow

### **Adding New Features**

1. **Plan the feature:**
   ```bash
   # Create an issue first to discuss the feature
   # Get feedback from maintainers before starting
   ```

2. **API Design:**
   ```go
   // Follow existing patterns for consistency
   // Example: New API service
   type NewAPIService service

   func (s *NewAPIService) NewOperation(ctx context.Context, accountID uuid.UUID) *NewOperationRequest {
       return &NewOperationRequest{
           apiService: s,
           ctx: ctx,
           accountID: accountID,
       }
   }
   ```

3. **Add tests:**
   ```go
   // test/new_feature_test.go
   func TestNewFeature(t *testing.T) {
       tests := []struct {
           name string
           input interface{}
           want interface{}
           wantErr bool
       }{
           // Add comprehensive test cases
       }

       for _, tt := range tests {
           t.Run(tt.name, func(t *testing.T) {
               // Test implementation
           })
       }
   }
   ```

### **Adding Examples**

Create comprehensive, real-world examples:

```go
//go:build ignore

package main

import (
    // Standard imports
    "context"
    "fmt"
    "log"
    "os"

    // SDK import
    "github.com/AhaSend/ahasend-go"
    "github.com/google/uuid"
)

func main() {
    // 1. Environment variable setup with clear instructions
    apiKey := os.Getenv("AHASEND_API_KEY")
    if apiKey == "" {
        log.Fatal("AHASEND_API_KEY environment variable is required")
    }

    // 2. Clear, step-by-step implementation

    // 3. Comprehensive error handling

    // 4. Educational comments explaining each step
}
```

**Example Guidelines:**
- Start with `//go:build ignore` for executable examples
- Include comprehensive error handling
- Add educational comments
- Show real-world usage patterns
- Test that examples build and run correctly

### **Adding Tests**

#### **Unit Tests**
```go
// test/feature_test.go
func TestFeatureName(t *testing.T) {
    // Setup
    client := ahasend.NewAPIClient(ahasend.NewConfiguration())

    // Table-driven tests
    tests := []struct {
        name    string
        input   interface{}
        want    interface{}
        wantErr bool
    }{
        {
            name: "success case",
            input: validInput,
            want: expectedOutput,
            wantErr: false,
        },
        {
            name: "error case",
            input: invalidInput,
            want: nil,
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := client.MethodUnderTest(tt.input)

            if (err != nil) != tt.wantErr {
                t.Errorf("MethodUnderTest() error = %v, wantErr %v", err, tt.wantErr)
                return
            }

            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("MethodUnderTest() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

#### **Integration Tests**
```go
// test/integration_feature_test.go
func TestFeatureIntegration(t *testing.T) {
    if os.Getenv("SKIP_INTEGRATION_TESTS") == "true" {
        t.Skip("Skipping integration test")
    }

    // Use Prism mock server for integration tests
    client := createTestClient()
    ctx := createAuthContext()

    // Test real API interactions
}
```

### **Documentation Updates**

1. **README.md**: User-facing features and changes
2. **CLAUDE.md**: Development and AI assistant guidance
3. **Examples**: Always include working code examples
4. **Inline docs**: Godoc comments for all public APIs
5. **Guides**: Comprehensive guides for complex features (ENV.md, IDEMPOTENCY.md)

## 🎯 Contribution Areas

### **High-Impact Contributions**

1. **🐛 Bug Fixes**
   - Rate limiter edge cases
   - Context cancellation issues
   - Memory leaks or race conditions
   - API compatibility issues

2. **✨ New Features**
   - Additional API endpoints
   - Enhanced error handling
   - Performance optimizations
   - Developer experience improvements

3. **📚 Documentation**
   - More real-world examples
   - Performance tuning guides
   - Troubleshooting documentation
   - Best practices guides

4. **🧪 Testing**
   - Edge case coverage
   - Performance benchmarks
   - Integration test scenarios
   - Mock improvements

## 🚀 Release Process

Releases are handled by maintainers but contributors can help:

1. **Version Planning**: Semantic versioning (v1.x.x)
2. **CHANGELOG.md**: Document all changes
3. **Testing**: Comprehensive testing before release
4. **Documentation**: Ensure docs are up to date
5. **GitHub Release**: Automated via GitHub Actions

**Release Types:**
- **Patch (v1.0.X)**: Bug fixes, documentation
- **Minor (v1.X.0)**: New features, backwards compatible
- **Major (vX.0.0)**: Breaking changes

## 🤝 Getting Help

### **Before You Start**
- 📖 Read the [documentation](https://ahasend.com/docs)
- 🔍 Search existing [issues](https://github.com/AhaSend/ahasend-go/issues)
- 💬 Check [GitHub Discussions](https://github.com/AhaSend/ahasend-go/discussions)

### **During Development**
- **Questions**: Use GitHub Discussions for general questions
- **Issues**: Create detailed bug reports or feature requests
- **Code Review**: Engage constructively in pull request reviews
- **API Questions**: Contact support@ahasend.com for API-specific questions

### **Community Guidelines**
- **Be respectful and inclusive** in all interactions
- **Focus on constructive feedback** that helps improve the project
- **Help others** by answering questions and reviewing code
- **Follow the code of conduct** and maintain a welcoming environment
- **Report issues** to maintainers if you encounter problems

## 📋 Contributor Checklist

Before submitting your contribution:

**Code Quality:**
- [ ] Code passes `go fmt ./...`
- [ ] Code passes `go vet ./...`
- [ ] Code passes `golangci-lint run` (if available)
- [ ] No race conditions detected with `go test -race`

**Testing:**
- [ ] Unit tests added/updated for new functionality
- [ ] All tests pass: `SKIP_INTEGRATION_TESTS=true go test -v ./test/`
- [ ] Integration tests pass: `go test -v ./test/` (if applicable)
- [ ] Examples still build and run correctly

**Documentation:**
- [ ] Public APIs have godoc comments
- [ ] README updated for user-facing changes
- [ ] Examples added/updated as needed
- [ ] Inline comments explain complex logic

**Git:**
- [ ] Commit messages follow conventional format
- [ ] Branch name is descriptive
- [ ] Changes are focused and atomic
- [ ] No merge commits in feature branch

## 📄 License

By contributing to AhaSend Go SDK, you agree that your contributions will be licensed under the same [MIT License](LICENSE) that covers the project.