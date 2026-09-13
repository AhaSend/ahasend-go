package requests

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateContactRequestRejectsNullAttributeValues(t *testing.T) {
	var nilString *string
	var nilMap map[string]any
	var nilSlice []any

	tests := []struct {
		name       string
		attributes map[string]any
		wantError  bool
	}{
		{name: "canonical scalars", attributes: map[string]any{"name": "Ada", "score": 12.5, "active": true}},
		{name: "untyped null", attributes: map[string]any{"obsolete": nil}, wantError: true},
		{name: "typed nil pointer", attributes: map[string]any{"obsolete": nilString}, wantError: true},
		{name: "typed nil map", attributes: map[string]any{"obsolete": nilMap}, wantError: true},
		{name: "typed nil slice", attributes: map[string]any{"obsolete": nilSlice}, wantError: true},
		{name: "raw JSON null", attributes: map[string]any{"obsolete": json.RawMessage("null")}, wantError: true},
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
