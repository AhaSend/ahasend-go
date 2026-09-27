package requests

import (
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListRequestsEncodeOnlyNamedFields(t *testing.T) {
	empty := ""
	noTags := []string{}
	status := ListContactStatusUnsubscribed
	email := "person@example.com"
	id := uuid.MustParse("11111111-1111-4111-8111-111111111111")

	tests := []struct {
		name    string
		request any
		want    string
	}{
		{name: "create with name only", request: CreateListRequest{Name: "Newsletter"}, want: `{"name":"Newsletter"}`},
		{name: "update naming nothing", request: UpdateListRequest{}, want: `{}`},
		{name: "update clearing description and tags", request: UpdateListRequest{Description: &empty, Tags: &noTags}, want: `{"description":"","tags":[]}`},
		{name: "upsert naming no status", request: UpsertListContactRequest{}, want: `{}`},
		{name: "upsert naming a status", request: UpsertListContactRequest{SubscriptionStatus: &status}, want: `{"subscription_status":"unsubscribed"}`},
		{name: "batch entry by email", request: BatchAddListContactInput{Email: &email}, want: `{"email":"person@example.com"}`},
		{name: "batch entry by id", request: BatchAddListContactInput{ID: &id}, want: `{"id":"11111111-1111-4111-8111-111111111111"}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded, err := json.Marshal(tt.request)

			require.NoError(t, err)
			assert.JSONEq(t, tt.want, string(encoded))
		})
	}
}
