# Makefile for AhaSend Go SDK
# Provides common development tasks and shortcuts

# Variables
GO_VERSION := 1.21
BINARY_NAME := ahasend-sdk
PKG_LIST := $(shell go list ./... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
COVERAGE_OUT := coverage.out
COVERAGE_HTML := coverage.html

# Colors for output
RED := \033[31m
GREEN := \033[32m
YELLOW := \033[33m
BLUE := \033[34m
RESET := \033[0m

.PHONY: all build clean test test-unit test-integration test-coverage lint fmt vet check-deps deps help install-tools benchmark security mock-server setup ci

# Default target
all: check-deps fmt lint vet test build

# Help target
help: ## Show this help message
	@echo "$(BLUE)AhaSend Go SDK - Development Commands$(RESET)"
	@echo ""
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "$(GREEN)%-20s$(RESET) %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Setup development environment
setup: install-tools deps ## Set up development environment
	@echo "$(GREEN)Development environment set up successfully!$(RESET)"

# Install development tools
install-tools: ## Install required development tools
	@echo "$(BLUE)Installing development tools...$(RESET)"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install github.com/securecodewarrior/github-action-gosec/v2/cmd/gosec@latest
	@which npm > /dev/null || (echo "$(RED)npm is required but not installed$(RESET)" && exit 1)
	npm install -g @stoplight/prism-cli@5
	@echo "$(GREEN)Development tools installed successfully!$(RESET)"

# Dependency management
deps: ## Download and verify dependencies
	@echo "$(BLUE)Downloading dependencies...$(RESET)"
	go mod download
	go mod verify
	go mod tidy

check-deps: ## Check if dependencies are up to date
	@echo "$(BLUE)Checking dependencies...$(RESET)"
	@go list -u -m all 2>/dev/null | grep -v "$(shell go list -m)" | grep "=>" || echo "$(GREEN)All dependencies are up to date$(RESET)"

# Code quality
fmt: ## Format Go code
	@echo "$(BLUE)Formatting code...$(RESET)"
	gofmt -s -w .
	goimports -w .

lint: ## Run linting
	@echo "$(BLUE)Running linter...$(RESET)"
	golangci-lint run ./...

vet: ## Run go vet
	@echo "$(BLUE)Running go vet...$(RESET)"
	go vet ./...

# Testing
test: test-unit ## Run all tests
	@echo "$(GREEN)All tests completed!$(RESET)"

test-unit: ## Run unit tests
	@echo "$(BLUE)Running unit tests...$(RESET)"
	SKIP_INTEGRATION_TESTS=true go test -v -race -short ./...

test-integration: ## Run integration tests (requires Prism)
	@echo "$(BLUE)Running integration tests...$(RESET)"
	@which prism > /dev/null || (echo "$(RED)Prism CLI is required for integration tests$(RESET)" && echo "$(YELLOW)Install with: npm install -g @stoplight/prism-cli$(RESET)" && exit 1)
	@test -f openapi/openapi.yaml || (echo "$(RED)OpenAPI spec not found at api/openapi.yaml$(RESET)" && exit 1)
	SKIP_INTEGRATION_TESTS=false go test -v -timeout=10m -tags=integration ./test/

test-coverage: ## Run tests with coverage
	@echo "$(BLUE)Running tests with coverage...$(RESET)"
	SKIP_INTEGRATION_TESTS=true go test -v -race -coverprofile=$(COVERAGE_OUT) -covermode=atomic ./...
	go tool cover -html=$(COVERAGE_OUT) -o $(COVERAGE_HTML)
	go tool cover -func=$(COVERAGE_OUT)
	@echo "$(GREEN)Coverage report generated: $(COVERAGE_HTML)$(RESET)"

benchmark: ## Run benchmarks
	@echo "$(BLUE)Running benchmarks...$(RESET)"
	@which prism > /dev/null || (echo "$(RED)Prism CLI is required for benchmarks$(RESET)" && exit 1)
	SKIP_INTEGRATION_TESTS=false go test -bench=. -benchmem -run=^$ ./test/

# Security
security: ## Run security checks
	@echo "$(BLUE)Running security checks...$(RESET)"
	govulncheck ./...

# Build
build: ## Build the SDK (validation)
	@echo "$(BLUE)Building SDK...$(RESET)"
	go build -v ./...

# Clean
clean: ## Clean build artifacts and test files
	@echo "$(BLUE)Cleaning up...$(RESET)"
	go clean -testcache
	rm -f $(COVERAGE_OUT) $(COVERAGE_HTML)
	rm -f gosec-report.json
	rm -f benchmark.txt
	rm -f build-info.txt

# Mock server management
mock-server: ## Start Prism mock server (for manual testing)
	@echo "$(BLUE)Starting Prism mock server...$(RESET)"
	@which prism > /dev/null || (echo "$(RED)Prism CLI is required$(RESET)" && echo "$(YELLOW)Install with: npm install -g @stoplight/prism-cli$(RESET)" && exit 1)
	@test -f api/openapi.yaml || (echo "$(RED)OpenAPI spec not found at api/openapi.yaml$(RESET)" && exit 1)
	@echo "$(GREEN)Mock server will start on http://localhost:4010$(RESET)"
	@echo "$(YELLOW)Press Ctrl+C to stop the server$(RESET)"
	prism mock api/openapi.yaml --host 0.0.0.0 --port 4010 --dynamic

mock-server-validate: ## Validate OpenAPI spec with Prism
	@echo "$(BLUE)Validating OpenAPI spec...$(RESET)"
	@which prism > /dev/null || (echo "$(RED)Prism CLI is required$(RESET)" && exit 1)
	@test -f api/openapi.yaml || (echo "$(RED)OpenAPI spec not found at api/openapi.yaml$(RESET)" && exit 1)
	npx @apidevtools/swagger-parser validate api/openapi.yaml
	@echo "$(GREEN)OpenAPI spec is valid!$(RESET)"

# Development workflow
dev-test: fmt lint vet test-unit ## Quick development test cycle
	@echo "$(GREEN)Development tests passed!$(RESET)"

full-test: fmt lint vet test-unit test-integration test-coverage security ## Full test cycle
	@echo "$(GREEN)All tests and checks passed!$(RESET)"

ci: deps full-test build ## Simulate CI pipeline locally
	@echo "$(GREEN)CI simulation completed successfully!$(RESET)"

# Documentation
docs: ## Generate documentation (placeholder)
	@echo "$(BLUE)Generating documentation...$(RESET)"
	@echo "$(YELLOW)Documentation generation not implemented yet$(RESET)"
	@echo "$(GREEN)Available documentation:$(RESET)"
	@echo "  - README.md: Main documentation"
	@echo "  - ENV.md: Environment variable configuration"
	@echo "  - IDEMPOTENCY.md: Idempotency and retry documentation"
	@echo "  - docs/: Auto-generated API documentation"

# Version management
version: ## Show current version info
	@echo "$(BLUE)Version Information:$(RESET)"
	@echo "Go version: $(shell go version)"
	@echo "Git commit: $(shell git rev-parse HEAD 2>/dev/null || echo 'not a git repository')"
	@echo "Git branch: $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null || echo 'not a git repository')"
	@echo "Build date: $(shell date -u +%Y-%m-%dT%H:%M:%SZ)"

# Performance testing
perf-test: ## Run performance tests
	@echo "$(BLUE)Running performance tests...$(RESET)"
	@which prism > /dev/null || (echo "$(RED)Prism CLI is required$(RESET)" && exit 1)
	SKIP_INTEGRATION_TESTS=false go test -run=XXX -bench=. -benchtime=10s -benchmem ./test/

# Code quality reports
quality-report: ## Generate code quality reports
	@echo "$(BLUE)Generating code quality reports...$(RESET)"
	mkdir -p reports

	# Linting report
	golangci-lint run --out-format=html > reports/lint-report.html 2>/dev/null || true

	# Test coverage
	go test -coverprofile=reports/coverage.out ./...
	go tool cover -html=reports/coverage.out -o reports/coverage.html

	# Security report
	gosec -fmt=html -out=reports/security-report.html ./... || true

	@echo "$(GREEN)Quality reports generated in ./reports/$(RESET)"

# Release preparation
release-check: ## Check if ready for release
	@echo "$(BLUE)Checking release readiness...$(RESET)"
	@echo "Running full test suite..."
	@$(MAKE) full-test
	@echo "$(BLUE)Checking for uncommitted changes...$(RESET)"
	@git diff --quiet || (echo "$(RED)There are uncommitted changes$(RESET)" && exit 1)
	@echo "$(BLUE)Checking if on main branch...$(RESET)"
	@test "$(shell git rev-parse --abbrev-ref HEAD)" = "main" || (echo "$(YELLOW)Not on main branch$(RESET)")
	@echo "$(GREEN)Release check completed!$(RESET)"

# Debugging helpers
debug-info: ## Show debugging information
	@echo "$(BLUE)Debug Information:$(RESET)"
	@echo "GOPATH: $(GOPATH)"
	@echo "GOROOT: $(GOROOT)"
	@echo "Go version: $(shell go version)"
	@echo "Go env GOOS: $(shell go env GOOS)"
	@echo "Go env GOARCH: $(shell go env GOARCH)"
	@echo "Current directory: $(PWD)"
	@echo "Available make targets:"
	@$(MAKE) help

# Install SDK locally (for testing)
install-local: ## Install SDK locally for testing
	@echo "$(BLUE)Installing SDK locally...$(RESET)"
	go install ./...

# Update dependencies
update-deps: ## Update all dependencies
	@echo "$(BLUE)Updating dependencies...$(RESET)"
	go get -u all
	go mod tidy
	@echo "$(GREEN)Dependencies updated!$(RESET)"

# Git hooks setup
setup-hooks: ## Set up git hooks
	@echo "$(BLUE)Setting up git hooks...$(RESET)"
	@mkdir -p .git/hooks
	@echo "#!/bin/sh\nmake dev-test" > .git/hooks/pre-commit
	@chmod +x .git/hooks/pre-commit
	@echo "$(GREEN)Git hooks set up! Tests will run before commits.$(RESET)"

# Show project statistics
stats: ## Show project statistics
	@echo "$(BLUE)Project Statistics:$(RESET)"
	@echo "Go files: $(shell find . -name '*.go' | wc -l)"
	@echo "Lines of code: $(shell find . -name '*.go' -exec cat {} \; | wc -l)"
	@echo "Test files: $(shell find . -name '*_test.go' | wc -l)"
	@echo "Packages: $(shell go list ./... | wc -l)"
	@echo "Dependencies: $(shell go list -m all | wc -l)"

# Watch for changes and run tests
watch: ## Watch for changes and run tests (requires inotify-tools)
	@echo "$(BLUE)Watching for changes...$(RESET)"
	@echo "$(YELLOW)Press Ctrl+C to stop watching$(RESET)"
	@which inotifywait > /dev/null || (echo "$(RED)inotify-tools is required$(RESET)" && echo "$(YELLOW)Install with: apt-get install inotify-tools (Ubuntu/Debian)$(RESET)" && exit 1)
	@while inotifywait -r -e modify,create,delete --include '\.go$$' .; do \
		clear; \
		echo "$(BLUE)Files changed, running tests...$(RESET)"; \
		$(MAKE) dev-test; \
	done