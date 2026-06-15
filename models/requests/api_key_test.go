package requests

import (
	"strings"
	"testing"

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
