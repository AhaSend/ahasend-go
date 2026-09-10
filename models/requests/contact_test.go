package requests

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateContactRequestRejectsNullAttributeValues(t *testing.T) {
	tests := []struct {
		name       string
		attributes map[string]any
		wantError  bool
	}{
		{name: "canonical scalars", attributes: map[string]any{"name": "Ada", "score": 12.5, "active": true}},
		{name: "null", attributes: map[string]any{"obsolete": nil}, wantError: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := (CreateContactRequest{Email: "person@example.com", Attributes: tt.attributes}).Validate()
			if tt.wantError {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "must not be null")
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestContactMutationRequestsPreserveNullAttributes(t *testing.T) {
	tests := []struct {
		name    string
		request any
	}{
		{
			name: "update",
			request: UpdateContactRequest{
				Attributes: map[string]any{"obsolete": nil},
			},
		},
		{
			name: "batch",
			request: BatchUpsertContactsRequest{Data: []BatchUpsertContactInput{
				{Email: "person@example.com", Attributes: map[string]any{"obsolete": nil}},
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded, err := json.Marshal(tt.request)
			require.NoError(t, err)
			assert.Contains(t, string(encoded), `"obsolete":null`)
		})
	}
}
