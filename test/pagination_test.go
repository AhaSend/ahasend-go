package ahasend_test

import (
	"encoding/json"
	"testing"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/AhaSend/ahasend-go/models/responses"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPaginationInfo(t *testing.T) {
	t.Run("Complete pagination info", func(t *testing.T) {
		cursor := "eyJpZCI6MTIzNH0="
		info := common.PaginationInfo{
			HasMore:    true,
			NextCursor: &cursor,
		}

		// Test JSON marshaling
		data, err := json.Marshal(info)
		require.NoError(t, err)

		var decoded common.PaginationInfo
		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)

		assert.Equal(t, info.HasMore, decoded.HasMore)
		assert.NotNil(t, decoded.NextCursor)
		assert.Equal(t, cursor, *decoded.NextCursor)
	})

	t.Run("No next cursor", func(t *testing.T) {
		info := common.PaginationInfo{
			HasMore:    false,
			NextCursor: nil,
		}

		// Test JSON marshaling
		data, err := json.Marshal(info)
		require.NoError(t, err)

		var decoded common.PaginationInfo
		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)

		assert.Equal(t, info.HasMore, decoded.HasMore)
		assert.Nil(t, decoded.NextCursor)
	})

	t.Run("JSON omitempty behavior", func(t *testing.T) {
		info := common.PaginationInfo{
			HasMore:    false,
			NextCursor: nil,
		}

		data, err := json.Marshal(info)
		require.NoError(t, err)

		// Should not include next_cursor field when nil
		assert.Contains(t, string(data), `"has_more":false`)
		assert.NotContains(t, string(data), "next_cursor")
	})
}

func TestPaginatedResponse(t *testing.T) {
	t.Run("Generic paginated response with messages", func(t *testing.T) {
		cursor := "next_page_cursor"
		response := common.PaginatedResponse[responses.Message]{
			Object: "list",
			Data: []responses.Message{
				{
					ApiID:     uuid.MustParse("00000000-0000-0000-0000-000000000123"),
					Subject:   "Test Message",
					Status:    "delivered",
					Sender:    "test@example.com",
					Recipient: "user@example.com",
				},
				{
					ApiID:     uuid.MustParse("00000000-0000-0000-0000-000000000124"),
					Subject:   "Another Message",
					Status:    "queued",
					Sender:    "test@example.com",
					Recipient: "user2@example.com",
				},
			},
			Pagination: common.PaginationInfo{
				HasMore:    true,
				NextCursor: &cursor,
			},
		}

		// Test JSON marshaling
		data, err := json.Marshal(response)
		require.NoError(t, err)

		var decoded common.PaginatedResponse[responses.Message]
		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)

		assert.Equal(t, "list", decoded.Object)
		assert.Len(t, decoded.Data, 2)
		assert.Equal(t, uuid.MustParse("00000000-0000-0000-0000-000000000123"), decoded.Data[0].ApiID)
		assert.Equal(t, "Test Message", decoded.Data[0].Subject)
		assert.Equal(t, true, decoded.Pagination.HasMore)
		assert.NotNil(t, decoded.Pagination.NextCursor)
		assert.Equal(t, cursor, *decoded.Pagination.NextCursor)
	})

	t.Run("Generic paginated response with domains", func(t *testing.T) {
		response := common.PaginatedResponse[responses.Domain]{
			Object: "list",
			Data: []responses.Domain{
				{
					ID:     uuid.MustParse("00000000-0000-0000-0000-000000000123"),
					Domain: "example.com",
				},
			},
			Pagination: common.PaginationInfo{
				HasMore:    false,
				NextCursor: nil,
			},
		}

		// Test JSON marshaling
		data, err := json.Marshal(response)
		require.NoError(t, err)

		var decoded common.PaginatedResponse[responses.Domain]
		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)

		assert.Equal(t, "list", decoded.Object)
		assert.Len(t, decoded.Data, 1)
		assert.Equal(t, uuid.MustParse("00000000-0000-0000-0000-000000000123"), decoded.Data[0].ID)
		assert.Equal(t, "example.com", decoded.Data[0].Domain)
		assert.Equal(t, false, decoded.Pagination.HasMore)
		assert.Nil(t, decoded.Pagination.NextCursor)
	})

	t.Run("Empty data array", func(t *testing.T) {
		response := common.PaginatedResponse[responses.Message]{
			Object: "list",
			Data:   []responses.Message{},
			Pagination: common.PaginationInfo{
				HasMore:    false,
				NextCursor: nil,
			},
		}

		// Test JSON marshaling
		data, err := json.Marshal(response)
		require.NoError(t, err)

		var decoded common.PaginatedResponse[responses.Message]
		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)

		assert.Equal(t, "list", decoded.Object)
		assert.Empty(t, decoded.Data)
		assert.Equal(t, false, decoded.Pagination.HasMore)
	})
}

func TestTypeAliases(t *testing.T) {
	t.Run("Type aliases work correctly", func(t *testing.T) {
		// Test that the type aliases are correctly defined
		var messagesResponse responses.PaginatedMessagesResponse
		var domainsResponse responses.PaginatedDomainsResponse
		var webhooksResponse responses.PaginatedWebhooksResponse
		var routesResponse responses.PaginatedRoutesResponse
		var smtpResponse responses.PaginatedSMTPCredentialsResponse
		var suppressionsResponse responses.PaginatedSuppressionsResponse
		var apiKeysResponse responses.PaginatedAPIKeysResponse

		// Verify these can be assigned and used
		messagesResponse = responses.PaginatedMessagesResponse{
			Object: "list",
			Data:   []responses.Message{},
			Pagination: common.PaginationInfo{
				HasMore: false,
			},
		}

		assert.Equal(t, "list", messagesResponse.Object)
		assert.Empty(t, messagesResponse.Data)
		assert.False(t, messagesResponse.Pagination.HasMore)

		// Test that they are indeed generic types
		_ = messagesResponse
		_ = domainsResponse
		_ = webhooksResponse
		_ = routesResponse
		_ = smtpResponse
		_ = suppressionsResponse
		_ = apiKeysResponse
	})
}

func TestPaginationJSONFormat(t *testing.T) {
	t.Run("JSON format matches OpenAPI specification", func(t *testing.T) {
		cursor := "eyJpZCI6MTIzNH0="
		response := common.PaginatedResponse[responses.Message]{
			Object: "list",
			Data: []responses.Message{
				{
					ApiID:     uuid.MustParse("00000000-0000-0000-0000-000000000123"),
					AccountID: uuid.MustParse("00000000-0000-0000-0000-000000000123"),
					Subject:   "Test Message",
					Status:    "delivered",
				},
			},
			Pagination: common.PaginationInfo{
				HasMore:    true,
				NextCursor: &cursor,
			},
		}

		data, err := json.Marshal(response)
		require.NoError(t, err)

		// Parse back to verify structure
		var raw map[string]interface{}
		err = json.Unmarshal(data, &raw)
		require.NoError(t, err)

		// Verify top-level structure
		assert.Equal(t, "list", raw["object"])
		assert.Contains(t, raw, "data")
		assert.Contains(t, raw, "pagination")

		// Verify pagination structure
		pagination := raw["pagination"].(map[string]interface{})
		assert.Equal(t, true, pagination["has_more"])
		assert.Equal(t, cursor, pagination["next_cursor"])

		// Verify data array structure
		dataArray := raw["data"].([]interface{})
		assert.Len(t, dataArray, 1)

		message := dataArray[0].(map[string]interface{})
		assert.Equal(t, "00000000-0000-0000-0000-000000000123", message["api_id"])
		assert.Equal(t, "00000000-0000-0000-0000-000000000123", message["account_id"])
		assert.Equal(t, "Test Message", message["subject"])
		assert.Equal(t, "delivered", message["status"])
	})
}
