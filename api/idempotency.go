// Idempotency utilities for the AhaSend Go SDK.
//
// This file provides enhanced idempotency support with automatic UUID generation
// and utilities for safe request retries.

package api

import (
	"github.com/google/uuid"
)

// IdempotencyConfig configures idempotency behavior for the SDK
type IdempotencyConfig struct {
	// AutoGenerate automatically generates UUID idempotency keys when none provided
	AutoGenerate bool
	// KeyPrefix adds a prefix to all generated idempotency keys (optional)
	KeyPrefix string
}

// DefaultIdempotencyConfig returns the default idempotency configuration
func DefaultIdempotencyConfig() IdempotencyConfig {
	return IdempotencyConfig{
		AutoGenerate: true,
		KeyPrefix:    "",
	}
}

// GenerateIdempotencyKey generates a new UUID-based idempotency key
func GenerateIdempotencyKey() string {
	return uuid.New().String()
}

// GenerateIdempotencyKeyWithPrefix generates a new idempotency key with the given prefix
func GenerateIdempotencyKeyWithPrefix(prefix string) string {
	if prefix == "" {
		return GenerateIdempotencyKey()
	}
	return prefix + "-" + GenerateIdempotencyKey()
}

// IdempotencyKeyBuilder helps build deterministic idempotency keys for related operations
type IdempotencyKeyBuilder struct {
	baseKey string
	counter int
}

// NewIdempotencyKeyBuilder creates a new builder with a base key
func NewIdempotencyKeyBuilder(baseKey string) *IdempotencyKeyBuilder {
	if baseKey == "" {
		baseKey = GenerateIdempotencyKey()
	}
	return &IdempotencyKeyBuilder{
		baseKey: baseKey,
		counter: 0,
	}
}

// Next returns the next idempotency key in the sequence
func (b *IdempotencyKeyBuilder) Next() string {
	b.counter++
	if b.counter == 1 {
		return b.baseKey
	}
	return b.baseKey + "-" + uuid.New().String()[:8]
}

// WithSuffix returns an idempotency key with a specific suffix
func (b *IdempotencyKeyBuilder) WithSuffix(suffix string) string {
	return b.baseKey + "-" + suffix
}

// IdempotentOperation represents an operation that can be safely retried
type IdempotentOperation[T any] func(key string) (T, error)

// ExecuteIdempotent executes an idempotent operation with automatic key generation
func ExecuteIdempotent[T any](op IdempotentOperation[T], customKey ...string) (T, error) {
	key := GenerateIdempotencyKey()
	if len(customKey) > 0 && customKey[0] != "" {
		key = customKey[0]
	}
	return op(key)
}

// IdempotencyHelper provides convenience methods for idempotent operations
type IdempotencyHelper struct {
	config IdempotencyConfig
}

// NewIdempotencyHelper creates a new idempotency helper with the given config
func NewIdempotencyHelper(config ...IdempotencyConfig) *IdempotencyHelper {
	cfg := DefaultIdempotencyConfig()
	if len(config) > 0 {
		cfg = config[0]
	}
	return &IdempotencyHelper{config: cfg}
}

// GenerateKey generates a new idempotency key based on the helper's configuration
func (h *IdempotencyHelper) GenerateKey() string {
	return GenerateIdempotencyKeyWithPrefix(h.config.KeyPrefix)
}

// EnsureKey returns the provided key or generates a new one if empty
func (h *IdempotencyHelper) EnsureKey(key string) string {
	if key != "" {
		return key
	}
	if h.config.AutoGenerate {
		return h.GenerateKey()
	}
	return ""
}
