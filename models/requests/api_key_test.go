package requests

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateAPIKeyRequest_Validate(t *testing.T) {
	label := "Production API Key"
	blankLabel := ""
	overlongLabel := strings.Repeat("a", maxAPIKeyLabelLength+1)
	scopes := []string{"messages:send:all"}

	tests := []struct {
		name    string
		request UpdateAPIKeyRequest
		wantErr bool
	}{
		{
			name:    "all nil",
			request: UpdateAPIKeyRequest{},
			wantErr: true,
		},
		{
			name:    "valid label only",
			request: UpdateAPIKeyRequest{Label: &label},
		},
		{
			name:    "valid scopes only",
			request: UpdateAPIKeyRequest{Scopes: &scopes},
		},
		{
			name:    "empty label",
			request: UpdateAPIKeyRequest{Label: &blankLabel},
			wantErr: true,
		},
		{
			name:    "overlong label",
			request: UpdateAPIKeyRequest{Label: &overlongLabel},
			wantErr: true,
		},
		{
			name:    "provided empty scopes",
			request: UpdateAPIKeyRequest{Scopes: &[]string{}},
			wantErr: true,
		},
		{
			name:    "valid ip_allow_list only",
			request: UpdateAPIKeyRequest{IPAllowList: &[]string{"203.0.113.0/24"}},
		},
		{
			name:    "empty ip_allow_list clears the list",
			request: UpdateAPIKeyRequest{IPAllowList: &[]string{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestCreateAPIKeyRequest_IPAllowListMarshaling(t *testing.T) {
	t.Run("omits ip_allow_list when empty", func(t *testing.T) {
		data, err := json.Marshal(CreateAPIKeyRequest{Label: "key", Scopes: []string{"messages:read:all"}})
		require.NoError(t, err)
		assert.NotContains(t, string(data), "ip_allow_list")
	})

	t.Run("includes ip_allow_list when set", func(t *testing.T) {
		data, err := json.Marshal(CreateAPIKeyRequest{
			Label:       "key",
			Scopes:      []string{"messages:read:all"},
			IPAllowList: []string{"203.0.113.0/24", "198.51.100.7"},
		})
		require.NoError(t, err)
		assert.Contains(t, string(data), `"ip_allow_list":["203.0.113.0/24","198.51.100.7"]`)
	})
}

// TestUpdateAPIKeyRequest_IPAllowListMarshaling locks in the tri-state contract:
// nil -> field omitted (leave unchanged); &[]string{} -> "[]" (clear);
// non-empty pointer -> the array (replace).
func TestUpdateAPIKeyRequest_IPAllowListMarshaling(t *testing.T) {
	t.Run("nil pointer omits the field (leave unchanged)", func(t *testing.T) {
		data, err := json.Marshal(UpdateAPIKeyRequest{Label: strPtr("key")})
		require.NoError(t, err)
		assert.NotContains(t, string(data), "ip_allow_list")
	})

	t.Run("empty slice clears the list", func(t *testing.T) {
		data, err := json.Marshal(UpdateAPIKeyRequest{IPAllowList: &[]string{}})
		require.NoError(t, err)
		assert.Contains(t, string(data), `"ip_allow_list":[]`)
	})

	t.Run("non-empty slice replaces the list", func(t *testing.T) {
		data, err := json.Marshal(UpdateAPIKeyRequest{IPAllowList: &[]string{"203.0.113.0/24"}})
		require.NoError(t, err)
		assert.Contains(t, string(data), `"ip_allow_list":["203.0.113.0/24"]`)
	})
}

func strPtr(s string) *string { return &s }
