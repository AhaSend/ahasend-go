package api

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdempotencyConfig(t *testing.T) {
	config := DefaultIdempotencyConfig()
	assert.True(t, config.AutoGenerate, "Should auto-generate by default")
	assert.Empty(t, config.KeyPrefix, "Should have no prefix by default")
}

func TestGenerateIdempotencyKey(t *testing.T) {
	key := GenerateIdempotencyKey()
	assert.NotEmpty(t, key, "Should generate a non-empty key")
	assert.Len(t, key, 36, "Should be UUID length (36 characters)")

	// Test uniqueness
	key2 := GenerateIdempotencyKey()
	assert.NotEqual(t, key, key2, "Should generate unique keys")
}

func TestGenerateIdempotencyKeyWithPrefix(t *testing.T) {
	prefix := "test"
	key := GenerateIdempotencyKeyWithPrefix(prefix)

	assert.NotEmpty(t, key, "Should generate a non-empty key")
	assert.True(t, strings.HasPrefix(key, prefix+"-"), "Should start with prefix")

	// Test without prefix
	keyWithoutPrefix := GenerateIdempotencyKeyWithPrefix("")
	assert.Len(t, keyWithoutPrefix, 36, "Should be normal UUID when prefix is empty")
}

func TestIdempotencyKeyBuilder(t *testing.T) {
	// Test with custom base key
	builder := NewIdempotencyKeyBuilder("base-key")

	first := builder.Next()
	assert.Equal(t, "base-key", first, "First key should be the base key")

	second := builder.Next()
	assert.True(t, strings.HasPrefix(second, "base-key-"), "Second key should have base key prefix")
	assert.NotEqual(t, first, second, "Keys should be unique")

	// Test with suffix
	suffixed := builder.WithSuffix("retry")
	assert.Equal(t, "base-key-retry", suffixed, "Should append suffix correctly")

	// Test with auto-generated base
	autoBuilder := NewIdempotencyKeyBuilder("")
	autoKey := autoBuilder.Next()
	assert.Len(t, autoKey, 36, "Auto-generated base should be UUID length")
}

func TestIdempotencyHelper(t *testing.T) {
	// Test with default config
	helper := NewIdempotencyHelper()

	key := helper.GenerateKey()
	assert.NotEmpty(t, key, "Should generate key")
	assert.Len(t, key, 36, "Should be UUID length")

	// Test EnsureKey with auto-generation
	ensuredKey := helper.EnsureKey("")
	assert.NotEmpty(t, ensuredKey, "Should generate key when empty")

	existingKey := helper.EnsureKey("existing")
	assert.Equal(t, "existing", existingKey, "Should return existing key")

	// Test with prefix config
	configWithPrefix := IdempotencyConfig{
		AutoGenerate: true,
		KeyPrefix:    "prefix",
	}
	helperWithPrefix := NewIdempotencyHelper(configWithPrefix)

	prefixedKey := helperWithPrefix.GenerateKey()
	assert.True(t, strings.HasPrefix(prefixedKey, "prefix-"), "Should include prefix")

	// Test with auto-generation disabled
	configNoAuto := IdempotencyConfig{
		AutoGenerate: false,
		KeyPrefix:    "",
	}
	helperNoAuto := NewIdempotencyHelper(configNoAuto)

	noAutoKey := helperNoAuto.EnsureKey("")
	assert.Empty(t, noAutoKey, "Should not generate key when auto-generation is disabled")
}

func TestExecuteIdempotent(t *testing.T) {
	// Test with custom key
	customKey := "test-key"
	result, err := ExecuteIdempotent(func(key string) (string, error) {
		assert.Equal(t, customKey, key, "Should use custom key")
		return "success", nil
	}, customKey)

	assert.NoError(t, err, "Should not error")
	assert.Equal(t, "success", result, "Should return operation result")

	// Test with auto-generated key
	result2, err := ExecuteIdempotent(func(key string) (string, error) {
		assert.NotEmpty(t, key, "Should generate key")
		assert.Len(t, key, 36, "Should be UUID length")
		return "auto-generated", nil
	})

	assert.NoError(t, err, "Should not error")
	assert.Equal(t, "auto-generated", result2, "Should return operation result")
}

func TestAPIClientIdempotencyIntegration(t *testing.T) {
	config := NewConfiguration()
	client := NewAPIClientWithConfig(config)

	// Test key generation
	key := client.GenerateIdempotencyKey()
	assert.NotEmpty(t, key, "Should generate key")
	assert.Len(t, key, 36, "Should be UUID length")

	// Test configuration
	currentConfig := client.GetIdempotencyConfig()
	assert.True(t, currentConfig.AutoGenerate, "Should have auto-generate enabled by default")

	// Test configuration update
	newConfig := IdempotencyConfig{
		AutoGenerate: false,
		KeyPrefix:    "test",
	}
	client.SetIdempotencyConfig(newConfig)

	updatedConfig := client.GetIdempotencyConfig()
	assert.False(t, updatedConfig.AutoGenerate, "Should update auto-generate setting")
	assert.Equal(t, "test", updatedConfig.KeyPrefix, "Should update prefix setting")

	// Test builder creation
	builder := client.NewIdempotencyKeyBuilder()
	builderKey := builder.Next()
	assert.NotEmpty(t, builderKey, "Should create working builder")

	customBuilder := client.NewIdempotencyKeyBuilder("custom-base")
	customKey := customBuilder.Next()
	assert.Equal(t, "custom-base", customKey, "Should use custom base key")
}

func TestIdempotencyConfigInConfiguration(t *testing.T) {
	config := NewConfiguration()

	// Should have default idempotency config
	assert.True(t, config.IdempotencyConfig.AutoGenerate, "Should enable auto-generation by default")
	assert.Empty(t, config.IdempotencyConfig.KeyPrefix, "Should have empty prefix by default")

	// Test custom configuration
	config.IdempotencyConfig.KeyPrefix = "custom"
	config.IdempotencyConfig.AutoGenerate = false

	client := NewAPIClientWithConfig(config)
	clientConfig := client.GetIdempotencyConfig()

	assert.False(t, clientConfig.AutoGenerate, "Should use custom auto-generate setting")
	assert.Equal(t, "custom", clientConfig.KeyPrefix, "Should use custom prefix")
}
